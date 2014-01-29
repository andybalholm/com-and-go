package com

import (
	"syscall"
	"unsafe"
)

// A Func represents a C function from a DLL.
type Func uintptr

// A libCall structure represents a call to a DLL function.
type libCall struct {
	fn          Func
	n           uintptr
	args        unsafe.Pointer
	r1, r2, err uintptr
}

//go:noescape
func cSyscall(c *libCall)

// Call calls f with the argument list specified by argPtr and argLen.
// The argument list needs to have the same layout in memory as the arguments
// that f expects.
// argPtr points to the first argument in the list. argLen is the number of
// CPU words in the argument list. Normally this is the same as the number of
// arguments, but it is larger if any of the arguments is larger than a CPU
// word.
//
// There are two main options for how to construct the argument list.
// One is to use the argument list of a wrapper function; take the address of
// the first argument (or potentially the method receiver).
// The other is to create a custom struct type to hold the argument list.
func (f Func) Call(argPtr unsafe.Pointer, argLen uintptr) (r1, r2, err uintptr) {
	c := libCall{
		fn:   f,
		n:    argLen,
		args: argPtr,
	}
	cSyscall(&c)
	return c.r1, c.r2, c.err
}

// CallInt is like call, but for a function that returns an integer.
func (f Func) CallInt(argPtr unsafe.Pointer, argLen uintptr) int {
	r, _, _ := f.Call(argPtr, argLen)
	return int(r)
}

// CallIntErr is like Call, but for a function that returns an integer, with a
// return value of 0 indicating that an error has occurred.
func (f Func) CallIntErr(argPtr unsafe.Pointer, argLen uintptr) (int, error) {
	r1, _, e := f.Call(argPtr, argLen)
	if r1 == 0 {
		return 0, syscall.Errno(e)
	}
	return int(r1), nil
}

// CallHR is like Call, but for a function that returns an HResult.
func (f Func) CallHR(argPtr unsafe.Pointer, argLen uintptr) error {
	hr, _, _ := f.Call(argPtr, argLen)
	if hr == 0 {
		return nil
	}
	return HResult(hr)
}

type DLL struct {
	*syscall.DLL
}

// LoadDLL loads a DLL file into memory. It panics if the file is not found.
func LoadDLL(name string) DLL {
	lib, err := syscall.LoadDLL(name)
	if err != nil {
		panic(err)
	}
	return DLL{lib}
}

// Func returns the specified function from d. It panics if the function is not
// found.
func (d DLL) Func(name string) Func {
	f, err := d.FindProc(name)
	if err != nil {
		panic(err)
	}
	return Func(f.Addr())
}
