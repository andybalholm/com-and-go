// The foxpro package is a database/sql driver for ADO connections to FoxPro
// databases.
package foxpro

import (
	"com-and-go"
	"database/sql"
	"database/sql/driver"
	"errors"
	"fmt"
	"io"
	"strings"
)

func init() {
	sql.Register("foxpro", foxDriver{})
}

type foxDriver struct{}

func (d foxDriver) Open(name string) (driver.Conn, error) {
	err := com.CoInitializeEx(0, 0)
	if err != nil {
		return nil, err
	}

	db, err := com.NewIDispatch("ADODB.Connection")
	if err != nil {
		return nil, err
	}

	c := &conn{
		db: db,
	}
	dsn := fmt.Sprintf("Provider=vfpoledb;Data Source=%s;", name)
	_, err = db.Call("Open", dsn)
	if err != nil {
		return nil, err
	}

	return c, nil
}

type conn struct {
	db *com.IDispatch
}

func (c *conn) Prepare(query string) (driver.Stmt, error) {
	return &stmt{c, query}, nil
}

func (c *conn) Close() error {
	_, err := c.db.Call("Close")
	return err
}

func (c *conn) Begin() (driver.Tx, error) {
	return nil, errors.New("foxpro: transactions aren't supported yet.")
}

type stmt struct {
	c     *conn
	query string
}

func (s *stmt) Close() error {
	return nil
}

func (s *stmt) NumInput() int {
	return -1
}

func (s *stmt) Exec(args []driver.Value) (driver.Result, error) {
	return nil, errors.New("foxpro: Exec isn't supported yet.")
}

func (s *stmt) Query(args []driver.Value) (driver.Rows, error) {
	cmd, err := com.NewIDispatch("ADODB.Command")
	if err != nil {
		return nil, err
	}
	defer cmd.Release()
	err = cmd.Put("ActiveConnection", s.c.db)
	if err != nil {
		return nil, err
	}

	err = cmd.Put("CommandText", s.query)
	if err != nil {
		return nil, err
	}
	err = cmd.Put("CommandType", 1)
	if err != nil {
		return nil, err
	}

	x, err := cmd.Get("Parameters")
	if err != nil {
		return nil, err
	}
	params := x.(*com.IDispatch)
	defer params.Release()

	for _, a := range args {
		var p interface{} // the parameter object
		switch v := a.(type) {
		case int64:
			v32 := int32(v)
			if int64(v32) != v {
				return nil, fmt.Errorf("integer too large to pass to FoxPro: %d", v)
			}
			p, err = cmd.Call("CreateParameter", "", 3 /* adInteger */, 1, 4, v32)
		case float64:
			p, err = cmd.Call("CreateParameter", "", 5 /* adDouble */, 1, 8, v)
		case bool:
			p, err = cmd.Call("CreateParameter", "", 11 /* adBoolean */, 1, 1, v)
		case []byte:
			p, err = cmd.Call("CreateParameter", "", 8 /* adBSTR */, 1, len(v), string(v))
		case string:
			p, err = cmd.Call("CreateParameter", "", 8 /* adBSTR */, 1, len(v), v)
		default:
			err = fmt.Errorf("foxpro: parameters of type %T are not supported", a)
		}
		if err != nil {
			return nil, err
		}
		param := p.(*com.IDispatch)
		defer param.Release()
		_, err = params.Call("Append", param)
		if err != nil {
			return nil, err
		}
	}

	x, err = cmd.Call("Execute")
	if err != nil {
		return nil, err
	}
	recordset := x.(*com.IDispatch)
	return &rows{recordset}, nil
}

type rows struct {
	rs *com.IDispatch
}

func (r *rows) Columns() []string {
	x, err := r.rs.Get("Fields")
	if err != nil {
		panic(err)
	}
	fields := x.(*com.IDispatch)
	defer fields.Release()

	x, err = fields.Get("Count")
	if err != nil {
		panic(err)
	}
	n := x.(int32)

	cols := make([]string, n)
	for i := int32(0); i < n; i++ {
		x, err = fields.Call("Item", i)
		if err != nil {
			panic(err)
		}
		item := x.(*com.IDispatch)
		defer item.Release()

		x, err = item.Get("Name")
		if err != nil {
			panic(err)
		}
		cols[i] = x.(string)
	}
	return cols
}

func (r *rows) Close() error {
	r.rs.Release()
	r.rs = nil
	return nil
}

func (r *rows) Next(dest []driver.Value) error {
	x, err := r.rs.Get("EOF")
	if err != nil {
		return err
	}
	if x == true {
		return io.EOF
	}

	x, err = r.rs.Get("Fields")
	if err != nil {
		return err
	}
	fields := x.(*com.IDispatch)
	defer fields.Release()

	for i := range dest {
		x, err = fields.Call("Item", int32(i))
		if err != nil {
			return err
		}
		item := x.(*com.IDispatch)
		defer item.Release()

		x, err = item.Get("Value")
		if err != nil {
			return err
		}
		switch v := x.(type) {
		case string:
			dest[i] = strings.TrimRight(v, " ")
		default:
			return fmt.Errorf("foxpro: result type %T not supported yet", v)
		}
	}

	_, err = r.rs.Call("MoveNext")
	return err
}
