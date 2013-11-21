package com

import (
	"fmt"
	"syscall"
	"unsafe"
)

type ITypeInfo uintptr // TODO

type ExcepInfo struct {
	Code           uint16
	WReserved      uint16
	Source         BStr
	Description    BStr
	HelpFile       BStr
	HelpContext    uint32
	PvReserved     uintptr
	DeferredFillIn uintptr
	Scode          int32
}

func (e *ExcepInfo) Error() string {
	if e.Description.P != nil {
		return e.Description.String()
	}
	if e.DeferredFillIn != 0 {
		Syscall(e.DeferredFillIn, e)
	}
	if e.Description.P != nil {
		return e.Description.String()
	}
	if e.Source.P != nil {
		return fmt.Sprintf("%s exception %d", e.Source, e.Code)
	}
	return fmt.Sprintf("COM exception %d", e.Code)
}

var IID_NULL = new(GUID)

const (
	DISPATCH_METHOD         = 1
	DISPATCH_PROPERTYGET    = 2
	DISPATCH_PROPERTYPUT    = 4
	DISPATCH_PROPERTYPUTREF = 8
)

const (
	DISPID_UNKNOWN     = -1
	DISPID_VALUE       = 0
	DISPID_PROPERTYPUT = -3
	DISPID_NEWENUM     = -4
	DISPID_EVALUATE    = -5
	DISPID_CONSTRUCTOR = -6
	DISPID_DESTRUCTOR  = -7
	DISPID_COLLECT     = -8
)

/* com
// IID {00020400-0000-0000-C000-000000000046}
type IDispatch interface {
	IUnknown
	GetTypeInfoCount() (count uint32, err error)
	GetTypeInfo(i uint, localeID uint32) (info *ITypeInfo, err error)
	GetIDsOfNames(iid *GUID, names []*uint16, localeID uint32, dispIDs *uint32) (err error)
	Invoke(member uint32, iid *GUID, localeID uint32, flags uint16, params *DispParams) (result Variant, excepInfo ExcepInfo, argErr uint32, err error)
}
*/

func (d *IDispatch) GetIDOfName(name string) (id uint32, err error) {
	name16, err := syscall.UTF16PtrFromString(name)
	if err != nil {
		return
	}
	err = d.GetIDsOfNames(IID_NULL, []*uint16{name16}, 0, &id)
	return
}

type DispParams struct {
	Args            *Variant
	DispIDNamedArgs *uint32
	CArgs           uint32
	CNamedArgs      uint32
}

// NewDispParams returns a pointer to a DispParams structure containing params.
func NewDispParams(params ...interface{}) *DispParams {
	dp := new(DispParams)
	if len(params) == 0 {
		return dp
	}

	variants := make([]Variant, len(params))
	for i, p := range params {
		// The parameters are in reverse order.
		variants[len(params)-i-1] = ToVariant(p)
	}

	dp.Args = &variants[0]
	dp.CArgs = uint32(len(params))
	return dp
}

func (d *IDispatch) Call(methodName string, params ...interface{}) (interface{}, error) {
	methodID, err := d.GetIDOfName(methodName)
	if err != nil {
		return nil, err
	}
	result, excepInfo, _, err := d.Invoke(methodID, IID_NULL, 0, DISPATCH_METHOD, NewDispParams(params...))
	if err == HResult(0x80020009) {
		err = &excepInfo
	}
	return result.ToInterface(), err
}

func (d *IDispatch) Get(propertyName string) (interface{}, error) {
	id, err := d.GetIDOfName(propertyName)
	if err != nil {
		return nil, err
	}
	result, excepInfo, _, err := d.Invoke(id, IID_NULL, 0, DISPATCH_PROPERTYGET, new(DispParams))
	if err == HResult(0x80020009) {
		err = &excepInfo
	}
	return result.ToInterface(), err
}

func (d *IDispatch) Put(propertyName string, value interface{}) (err error) {
	id, err := d.GetIDOfName(propertyName)
	if err != nil {
		return
	}
	var flags uint16
	if _, ok := value.(*IDispatch); ok {
		flags = DISPATCH_PROPERTYPUTREF
	} else {
		flags = DISPATCH_PROPERTYPUT
	}
	v := ToVariant(value)
	pp := int32(DISPID_PROPERTYPUT)
	dp := &DispParams{
		Args:            &v,
		DispIDNamedArgs: (*uint32)(unsafe.Pointer(&pp)),
		CArgs:           1,
		CNamedArgs:      1,
	}
	_, excepInfo, _, err := d.Invoke(id, IID_NULL, 0, flags, dp)
	if err == HResult(0x80020009) {
		err = &excepInfo
	}
	return
}

// NewIDispatch returns a new object of the specified class.
func NewIDispatch(class string) (*IDispatch, error) {
	clsID, err := CLSIDFromProgID(class)
	if err != nil {
		return nil, err
	}
	u, err := CoCreateInstance(&clsID, nil, 21, IID_IDispatch)
	if err != nil {
		return nil, err
	}
	return (*IDispatch)(u), nil
}
