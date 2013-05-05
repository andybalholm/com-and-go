package ado

import (
	"com-and-go"
)

type Recordset struct{ com.IDispatch } //TODO
type Errors struct{ com.IDispatch }

/* com
// IID {00000550-0000-0010-8000-00AA006D2EA4}
// CLSID {00000514-0000-0010-8000-00AA006D2EA4}
type Connection interface {
	_ADO
	ConnectionString() (str string, err error)
	PutConnectionString(str string) (err error)
	CommandTimeout() (timeout int32, err error)
	PutCommandTimeout(timeout int32) (err error)
	ConnectionTimeout() (timeout int32, err error)
	PutConnectionTimeout(timeout int32) (err error)
	Version() (version string, err error)
	Close() (err error)
	Execute(commandText string, recordsAffected *com.Variant, options int32) (rset *Recordset, err error)
	BeginTrans() (level int32, err error)
	CommitTrans() (err error)
	RolbackTrans() (err error)
	Open(connectionString string, userID string, password string, options int32) (err error)
	Errors() (obj *Errors, err error)
	DefaultDatabase() (str string, err error)
	PutDefaultDatabase(str string) (err error)
	IsolationLevel() (level int32, err error)
	PutIsolationLevel(level int32) (err error)
	Attributes() (attr int32, err error)
	PutAttributes(attr int32) (err error)
	CursorLocation() (cursorLoc int32, err error)
	PutCursorLocation(cursorLoc int32) (err error)
	Mode() (mode int32, err error)
	PutMode(mode int32) (err error)
	Provider() (provider string, err error)
	PutProvider(provider string) (err error)
	State() (objState int32, err error)
	_OpenSchema() // This doesn't work since I don't know how to pass Variants on the stack.
	Cancel() (err error)
}
*/
