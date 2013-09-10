// The foxpro package is a database/sql driver for ADO connections to FoxPro
// databases.
package foxpro

import (
	"code.google.com/p/com-and-go"
	"code.google.com/p/com-and-go/ado"
	"database/sql"
	"database/sql/driver"
	"errors"
	"fmt"
	"io"
	"strings"
	"time"
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

	db, err := ado.NewConnection()
	if err != nil {
		return nil, err
	}

	c := &conn{
		db: db,
	}
	dsn := fmt.Sprintf("Provider=vfpoledb;Data Source=%s;", name)
	err = db.Open(dsn, "", "", ado.ConnectUnspecified)
	if err != nil {
		return nil, err
	}

	return c, nil
}

type conn struct {
	db *ado.Connection
}

func (c *conn) Prepare(query string) (driver.Stmt, error) {
	return &stmt{c, query}, nil
}

func (c *conn) Close() error {
	err := c.db.Close()
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
	_, result, err := s.q(args)
	return result, err
}

func (s *stmt) Query(args []driver.Value) (driver.Rows, error) {
	rows, _, err := s.q(args)
	return rows, err
}

func (s *stmt) q(args []driver.Value) (driver.Rows, driver.Result, error) {
	cmd, err := ado.NewCommand()
	if err != nil {
		return nil, nil, err
	}
	defer cmd.Release()
	err = cmd.PutrefActiveADOConnection(s.c.db)
	if err != nil {
		return nil, nil, err
	}

	err = cmd.PutCommandText(s.query)
	if err != nil {
		return nil, nil, err
	}
	err = cmd.PutCommandType(1)
	if err != nil {
		return nil, nil, err
	}

	params, err := cmd.GetParameters()
	if err != nil {
		return nil, nil, err
	}
	defer params.Release()

	for _, a := range args {
		var param *ado.Parameter
		switch v := a.(type) {
		case int64:
			v32 := int32(v)
			if int64(v32) != v {
				return nil, nil, fmt.Errorf("integer too large to pass to FoxPro: %d", v)
			}
			param, err = cmd.CreateParameter("", ado.Integer, ado.ParamInput, 4, v32)
		case float64:
			param, err = cmd.CreateParameter("", ado.Double, ado.ParamInput, 8, v)
		case bool:
			param, err = cmd.CreateParameter("", ado.Boolean, ado.ParamInput, 1, v)
		case []byte:
			param, err = cmd.CreateParameter("", ado.BSTR, ado.ParamInput, uintptr(len(v)), string(v))
		case string:
			param, err = cmd.CreateParameter("", ado.BSTR, ado.ParamInput, uintptr(len(v)), v)
		case time.Time:
			param, err = cmd.CreateParameter("", ado.Date, ado.ParamInput, 8, v)
		default:
			err = fmt.Errorf("foxpro: parameters of type %T are not supported", a)
		}
		if err != nil {
			return nil, nil, err
		}
		defer param.Release()
		err = params.Append(param)
		if err != nil {
			return nil, nil, err
		}
	}

	var nRecords com.Variant
	recordset, err := cmd.Execute(&nRecords, nil, ado.OptionUnspecified)
	if err != nil {
		return nil, nil, err
	}
	return &rows{recordset}, driver.RowsAffected(nRecords.Val), nil
}

type rows struct {
	rs *ado.Recordset
}

func (r *rows) Columns() []string {
	fields, err := r.rs.GetFields()
	if err != nil {
		panic(err)
	}
	defer fields.Release()

	n, err := fields.GetCount()
	if err != nil {
		panic(err)
	}

	cols := make([]string, n)
	for i := int32(0); i < n; i++ {
		item, err := fields.GetItem(i)
		if err != nil {
			panic(err)
		}
		defer item.Release()

		name, err := item.GetName()
		if err != nil {
			panic(err)
		}
		cols[i] = name
	}
	return cols
}

func (r *rows) Close() error {
	r.rs.Release()
	r.rs = nil
	return nil
}

func (r *rows) Next(dest []driver.Value) error {
	eof, err := r.rs.GetEOF()
	if err != nil {
		return err
	}
	if eof {
		return io.EOF
	}

	fields, err := r.rs.GetFields()
	if err != nil {
		return err
	}
	defer fields.Release()

	for i := range dest {
		item, err := fields.GetItem(int32(i))
		if err != nil {
			return err
		}
		defer item.Release()

		v, err := item.GetValue()
		if err != nil {
			return err
		}
		switch v := v.(type) {
		case string:
			dest[i] = strings.TrimRight(v, " ")
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
		case com.Variant:
			return fmt.Errorf("foxpro: result variant with VT=%d not supported yet", v.VT)
		default:
			return fmt.Errorf("foxpro: result type %T not supported yet", v)
		}
	}

	err = r.rs.MoveNext()
	if err != nil {
		return err
	}
	return nil
}
