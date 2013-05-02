package com

import (
	"syscall"
	"unsafe"
)

// A VTable contains function pointers for a COM object's methods.
// The actual length varies; hopefully 1024 is enough to cover all we need.
type VTable [1024]uintptr

// IUnknown is the basic COM interface.
type IUnknown struct {
	vtable *VTable
}

var IID_IUnknown = NewGUID("{00000000-0000-0000-C000-000000000046}")

// QueryInterface converts u to the interface whose GUID is specified.
// If u does not implement that interface, it returns nil.
func (u *IUnknown) QueryInterface(id *GUID) unsafe.Pointer {
	var result unsafe.Pointer
	syscall.Syscall(u.vtable[0], 3,
		uintptr(unsafe.Pointer(u)),
		uintptr(unsafe.Pointer(id)),
		uintptr(unsafe.Pointer(&result)))
	return result
}

// AddRef increments u's reference count and returns the new count.
func (u *IUnknown) AddRef() uint32 {
	ret, _, _ := syscall.Syscall(u.vtable[1], 1,
		uintptr(unsafe.Pointer(u)),
		0,
		0)
	return uint32(ret)
}

// Release decrements u's reference count and returns the new count.
func (u *IUnknown) Release() uint32 {
	ret, _, _ := syscall.Syscall(u.vtable[2], 1,
		uintptr(unsafe.Pointer(u)),
		0,
		0)
	return uint32(ret)
}
