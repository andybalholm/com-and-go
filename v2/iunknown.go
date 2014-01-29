package com

import (
	"unsafe"
)

func init() {
	CoInitializeEx(nil, 0)
}

var IID_IUnknown = NewGUID("{00000000-0000-0000-C000-000000000046}")

type IUnknown struct {
	// VTable contains function pointers for a COM object's methods.
	// The actual length varies; hopefully 1024 is enough to cover all we need.
	VTable *[1024]Func
}

func (u *IUnknown) QueryInterface(riid *GUID, ppvObject unsafe.Pointer) error {
	return u.VTable[0].CallHR(unsafe.Pointer(&u), 3)
}

func (u *IUnknown) AddRef() int {
	return u.VTable[1].CallInt(unsafe.Pointer(&u), 1)
}

func (u *IUnknown) Release() int {
	return u.VTable[2].CallInt(unsafe.Pointer(&u), 1)
}

var (
	coCreateInstance = ole32.Func("CoCreateInstance")
	coInitializeEx   = ole32.Func("CoInitializeEx")
)

func CoInitializeEx(pvReserved *struct{}, dwCoInit int) error {
	return coInitializeEx.CallHR(unsafe.Pointer(&pvReserved), 2)
}

func CoCreateInstance(rclsid *GUID, pUnkOuter *IUnknown, dwClsContext int, riid *GUID, ppv unsafe.Pointer) error {
	return coCreateInstance.CallHR(unsafe.Pointer(&rclsid), 5)
}
