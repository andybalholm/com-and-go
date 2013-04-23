package com

import (
	"syscall"
	"unsafe"
)

var (
	modole32 = syscall.NewLazyDLL("ole32.dll")

	procCLSIDFromProgID  = modole32.NewProc("CLSIDFromProgID")
	procCoCreateInstance = modole32.NewProc("CoCreateInstance")
	procCoInitialize     = modole32.NewProc("CoInitialize")
)

func CLSIDFromProgID(progID string) (clsID *GUID, err error) {
	var _p0 *uint16
	_p0, err = syscall.UTF16PtrFromString(progID)
	if err != nil {
		return nil, err
	}
	clsID = new(GUID)
	r1, _, _ := syscall.Syscall(procCLSIDFromProgID.Addr(), 2, uintptr(unsafe.Pointer(_p0)), uintptr(unsafe.Pointer(clsID)), 0)
	if r1 != 0 {
		err = HResult(r1)
	}
	return
}

func CreateInstance(clsID *GUID) (*IUnknown, error) {
	var p *IUnknown
	r1, _, _ := procCoCreateInstance.Call(
		uintptr(unsafe.Pointer(clsID)),
		0,
		21, // any server context
		uintptr(unsafe.Pointer(IID_IUnknown)),
		uintptr(unsafe.Pointer(&p)))
	if r1 != 0 {
		return nil, HResult(r1)
	}
	return p, nil
}

func Initialize() error {
	r1, _, _ := procCoInitialize.Call(0)
	if r1 != 0 {
		return HResult(r1)
	}
	return nil
}
