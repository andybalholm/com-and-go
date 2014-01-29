// The ado package is a database/sql driver for ADO connections to databases.
package ado

import (
	"code.google.com/p/com-and-go/v2"
	"database/sql"
	"database/sql/driver"
	"errors"
	"fmt"
	"io"
	"time"
)

func init() {
	sql.Register("ado", adoDriver{})
}

type adoDriver struct{}

func (adoDriver) Open(name string) (driver.Conn, error) {
	db, err := com.NewIDispatch("ADODB.Connection")
	if err != nil {
		return nil, err
	}

	c := &conn{
		db: db,
	}
	_, err = db.CallErr("open", name)
	if err != nil {
		db.Release()
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
	_, err := c.db.CallErr("close")
	c.db.Release()
	return err
}

func (c *conn) Begin() (driver.Tx, error) {
	return nil, errors.New("ado: transactions aren't supported yet.")
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
	_, err := s.Query(args)
	return nil, err
}

func (s *stmt) Query(args []driver.Value) (driver.Rows, error) {
	cmd, err := com.NewIDispatch("ADODB.Command")
	if err != nil {
		return nil, err
	}
	defer cmd.Release()

	cmd.Put("ActiveConnection", s.c.db)
	cmd.Put("CommandText", s.query)
	cmd.Put("CommandType", 1)
	params := cmd.Get("Parameters").(*com.IDispatch)
	defer params.Release()

	for _, a := range args {
		var param interface{}
		switch a := a.(type) {
		case int64:
			a32 := int32(a)
			if int64(a32) == a {
				param = cmd.Call("CreateParameter", "", Integer, ParamInput, 4, a32)
			} else {
				param = cmd.Call("CreateParameter", "", BigInt, ParamInput, 8, a)
			}
		case float64:
			param = cmd.Call("CreateParameter", "", Double, ParamInput, 8, a)
		case bool:
			param = cmd.Call("CreateParameter", "", Boolean, ParamInput, 1, a)
		case []byte:
			param = cmd.Call("CreateParameter", "", BSTR, ParamInput, len(a), string(a))
		case string:
			param = cmd.Call("CreateParameter", "", BSTR, ParamInput, len(a), a)
		case time.Time:
			param = cmd.Call("CreateParameter", "", Date, ParamInput, 8, a)
		default:
			return nil, fmt.Errorf("ado: parameters of type %T are not supported", a)
		}
		params.Call("Append", param)
		param.(*com.IDispatch).Release()
	}

	recordset, err := cmd.CallErr("Execute")
	if err != nil {
		return nil, err
	}
	return &rows{recordset.(*com.IDispatch)}, nil
}

type rows struct {
	rs *com.IDispatch
}

func (r *rows) Columns() []string {
	fields := r.rs.Get("Fields").(*com.IDispatch)
	defer fields.Release()
	n := fields.Get("Count").(int32)

	cols := make([]string, n)
	for i := int32(0); i < n; i++ {
		item := fields.Call("Item", i).(*com.IDispatch)
		cols[i] = item.Get("Name").(string)
		item.Release()
	}
	return cols
}

func (r *rows) Close() error {
	r.rs.Release()
	r.rs = nil
	return nil
}

func (r *rows) Next(dest []driver.Value) error {
	if r.rs.Get("EOF") == true {
		return io.EOF
	}

	fields := r.rs.Get("Fields").(*com.IDispatch)
	defer fields.Release()

	for i := range dest {
		item := fields.Call("Item", int32(i)).(*com.IDispatch)
		defer item.Release()
		v := item.Get("Value")
		switch v := v.(type) {
		case string:
			dest[i] = v
		case int32:
			dest[i] = int64(v)
		case bool:
			dest[i] = v
		case com.Decimal:
			dest[i] = v.String()
		case time.Time:
			dest[i] = v
		case float64:
			dest[i] = v
		case nil:
			dest[i] = nil
		case []byte:
			dest[i] = v
		case com.Variant:
			return fmt.Errorf("ado: result variant with VT=%d not supported yet", v.VT)
		default:
			return fmt.Errorf("ado: result type %T not supported yet", v)
		}
	}

	r.rs.Call("MoveNext")
	return nil
}
