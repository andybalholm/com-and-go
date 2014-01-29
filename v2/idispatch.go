package com

import (
	"fmt"
	"unsafe"
)

type ITypeInfo struct {
	IUnknown
}

type ExcepInfo struct {
	Code           uint16
	WReserved      uint16
	Source         BStr
	Description    BStr
	HelpFile       BStr
	HelpContext    uint32
	PvReserved     uintptr
	DeferredFillIn Func
	Scode          int32
}

func (e *ExcepInfo) Error() string {
	if e.Description.P != nil {
		return e.Description.String()
	}
	if e.DeferredFillIn != 0 {
		e.DeferredFillIn.Call(unsafe.Pointer(&e), 1)
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

var IID_IDispatch = NewGUID("{00020400-0000-0000-C000-000000000046}")

type IDispatch struct {
	IUnknown
}

func (d *IDispatch) GetTypeInfoCount() (int, error) {
	var n uint32
	args := &struct {
		d       *IDispatch
		pctinfo *uint32
	}{d, &n}
	err := d.VTable[3].CallHR(unsafe.Pointer(args), 2)
	return int(n), err
}

func (d *IDispatch) GetTypeInfo(iTInfo int, lcid int) (ti *ITypeInfo, err error) {
	args := &struct {
		d       *IDispatch
		iTInfo  int
		lcid    int
		ppTInfo **ITypeInfo
	}{d, iTInfo, lcid, &ti}
	err = d.VTable[4].CallHR(unsafe.Pointer(args), 4)
	return ti, err
}

func (d *IDispatch) GetIDsOfNames(names []string, lcid int) (ids []uint32, err error) {
	wNames := make([]*uint16, len(names))
	for i, s := range names {
		wNames[i] = WideString(s)
	}
	ids = make([]uint32, len(names))
	args := &struct {
		d         *IDispatch
		reserved  *GUID
		rgszNames **uint16
		cNames    int
		lcid      int
		rgDispId  *uint32
	}{d, IID_NULL, &wNames[0], len(names), lcid, &ids[0]}

	err = d.VTable[5].CallHR(unsafe.Pointer(args), 6)
	return
}

func (d *IDispatch) Invoke(dispIdMember uint32, riid *GUID, lcid int, wFlags int, pDispParams *DispParams, pVarResult *Variant, pExcepInfo *ExcepInfo, puArgErr *uint32) error {
	return d.VTable[6].CallHR(unsafe.Pointer(&d), 9)
}

func (d *IDispatch) GetIDOfName(name string) (id uint32, err error) {
	ids, err := d.GetIDsOfNames([]string{name}, 0)
	return ids[0], err
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

// CallErr calls a method by name, returning an error on failure.
func (d *IDispatch) CallErr(methodName string, params ...interface{}) (interface{}, error) {
	methodID, err := d.GetIDOfName(methodName)
	if err != nil {
		return nil, err
	}

	// If any of the parameters is a string, convert it to a BStr and free it
	// later.
	for i, p := range params {
		if s, ok := p.(string); ok {
			b := ToBStr(s)
			params[i] = b
			defer b.Free()
		}
	}

	var result Variant
	var excepInfo ExcepInfo
	err = d.Invoke(methodID, IID_NULL, 0, DISPATCH_METHOD, NewDispParams(params...), &result, &excepInfo, nil)
	if err == HResult(0x80020009) {
		err = &excepInfo
	}

	r := result.ToInterface()
	return r, err
}

// Call calls a method by name. It panics if an error is encountered.
// Use Call when the only possible errors are programming mistakes; use CallErr
// when I/O errors and such are possible.
func (d *IDispatch) Call(methodName string, params ...interface{}) interface{} {
	res, err := d.CallErr(methodName, params...)
	if err != nil {
		panic(err)
	}
	return res
}

// GetErr returns the value of a property of d, returning an error on failure.
func (d *IDispatch) GetErr(propertyName string) (interface{}, error) {
	id, err := d.GetIDOfName(propertyName)
	if err != nil {
		return nil, err
	}
	var result Variant
	var excepInfo ExcepInfo
	err = d.Invoke(id, IID_NULL, 0, DISPATCH_PROPERTYGET, new(DispParams), &result, &excepInfo, nil)
	if err == HResult(0x80020009) {
		err = &excepInfo
	}

	r := result.ToInterface()
	return r, err
}

// Get returns the value of a property of d. It panics if an error is
// encountered.
func (d *IDispatch) Get(propertyName string) interface{} {
	res, err := d.GetErr(propertyName)
	if err != nil {
		panic(err)
	}
	return res
}

// PutErr sets a property of d, returning an error on failure.
func (d *IDispatch) PutErr(propertyName string, value interface{}) (err error) {
	id, err := d.GetIDOfName(propertyName)
	if err != nil {
		return
	}
	flags := DISPATCH_PROPERTYPUT
	switch v := value.(type) {
	case *IUnknown, *IDispatch:
		flags = DISPATCH_PROPERTYPUTREF
	case string:
		b := ToBStr(v)
		value = b
		defer b.Free()
	}

	v := ToVariant(value)
	pp := int32(DISPID_PROPERTYPUT)
	dp := &DispParams{
		Args:            &v,
		DispIDNamedArgs: (*uint32)(unsafe.Pointer(&pp)),
		CArgs:           1,
		CNamedArgs:      1,
	}
	var excepInfo ExcepInfo
	err = d.Invoke(id, IID_NULL, 0, flags, dp, nil, &excepInfo, nil)
	if err == HResult(0x80020009) {
		err = &excepInfo
	}
	return
}

// Put sets a property of d. It panics if an error is encoutered.
func (d *IDispatch) Put(propertyName string, value interface{}) {
	err := d.PutErr(propertyName, value)
	if err != nil {
		panic(err)
	}
}

// NewIDispatch returns a new object of the specified class.
func NewIDispatch(class string) (*IDispatch, error) {
	clsID, err := CLSIDFromProgID(class)
	if err != nil {
		return nil, err
	}
	var res *IDispatch
	err = CoCreateInstance(clsID, nil, 21, IID_IDispatch, unsafe.Pointer(&res))
	if err != nil {
		return nil, err
	}
	return res, nil
}
