package com

import (
	"errors"
	"reflect"
	"unsafe"
)

var (
	safeArrayAccessData   = oleaut32.Func("SafeArrayAccessData")
	safeArrayDestroy      = oleaut32.Func("SafeArrayDestroy")
	safeArrayGetDim       = oleaut32.Func("SafeArrayGetDim")
	safeArrayGetElement   = oleaut32.Func("SafeArrayGetElement")
	safeArrayGetElemsize  = oleaut32.Func("SafeArrayGetElemsize")
	safeArrayGetLBound    = oleaut32.Func("SafeArrayGetLBound")
	safeArrayGetUBound    = oleaut32.Func("SafeArrayGetUBound")
	safeArrayGetVartype   = oleaut32.Func("SafeArrayGetVartype")
	safeArrayUnaccessData = oleaut32.Func("SafeArrayUnaccessData")
)

type SafeArrayBound struct {
	Elements uint32
	Lbound   int32
}

type SafeArray struct {
	cDims      uint16
	fFeatures  uint16
	cbElements uint32
	cLocks     uint32
	pvData     unsafe.Pointer
	rgsabound  SafeArrayBound
}

func (a *SafeArray) GetVartype() Vartype {
	var vt Vartype
	args := &struct {
		psa *SafeArray
		pvt *Vartype
	}{a, &vt}
	err := safeArrayGetVartype.CallHR(unsafe.Pointer(args), 2)
	if err != nil {
		panic(err)
	}
	return vt
}

func (a *SafeArray) GetDim() int {
	return safeArrayGetDim.CallInt(unsafe.Pointer(&a), 1)
}

func (a *SafeArray) GetLBound(dim int) int {
	var bound int32
	args := &struct {
		psa      *SafeArray
		nDim     uint32
		plLbound *int32
	}{a, uint32(dim), &bound}
	err := safeArrayGetLBound.CallHR(unsafe.Pointer(args), 3)
	if err != nil {
		panic(err)
	}
	return int(bound)
}

func (a *SafeArray) GetUBound(dim int) int {
	var bound int32
	args := &struct {
		psa      *SafeArray
		nDim     uint32
		plUbound *int32
	}{a, uint32(dim), &bound}
	err := safeArrayGetUBound.CallHR(unsafe.Pointer(args), 3)
	if err != nil {
		panic(err)
	}
	return int(bound)
}

func (a *SafeArray) GetElement(indices []int32, dest unsafe.Pointer) {
	args := &struct {
		psa       *SafeArray
		rgIndices *int32
		pv        unsafe.Pointer
	}{a, &indices[0], dest}
	err := safeArrayGetElement.CallHR(unsafe.Pointer(args), 3)
	if err != nil {
		panic(err)
	}
}

func (a *SafeArray) GetElemSize() int {
	return safeArrayGetElemsize.CallInt(unsafe.Pointer(&a), 1)
}

func (a *SafeArray) Destroy() {
	err := safeArrayDestroy.CallHR(unsafe.Pointer(&a), 1)
	if err != nil {
		panic(err)
	}
}

func (a *SafeArray) AccessData() unsafe.Pointer {
	var p unsafe.Pointer
	args := &struct {
		psa     *SafeArray
		ppvData *unsafe.Pointer
	}{a, &p}
	err := safeArrayAccessData.CallHR(unsafe.Pointer(args), 2)
	if err != nil {
		panic(err)
	}
	return p
}

func (a *SafeArray) UnaccessData() {
	err := safeArrayUnaccessData.CallHR(unsafe.Pointer(&a), 1)
	if err != nil {
		panic(err)
	}
}

func (a *SafeArray) ToSlice() interface{} {
	if a.GetDim() != 1 {
		panic(errors.New("SafeArray.ToSlice only supports one-dimensional arrays"))
	}
	if a.GetElemSize() > 8 {
		panic(errors.New("SafeArray.ToSlice only supports arrays with elements that are 8 bytes or less"))
	}

	lBound := a.GetLBound(1)
	uBound := a.GetUBound(1)
	n := uBound - lBound + 1

	var elem Variant
	elem.VT = a.GetVartype()

	switch elem.VT {
	case VT_UI1:
		b := make([]byte, n)
		data := (*[1 << 30]byte)(a.AccessData())[:n]
		copy(b, data)
		a.UnaccessData()
		return b
	}

	elemType := reflect.TypeOf(elem.ToInterface())
	resultType := reflect.SliceOf(elemType)

	result := reflect.MakeSlice(resultType, n, n)

	for i := 0; i < n; i++ {
		a.GetElement([]int32{int32(lBound + i)}, unsafe.Pointer(&elem.Val))
		result.Index(i).Set(reflect.ValueOf(elem.ToInterface()))
	}
	return result.Interface()
}
