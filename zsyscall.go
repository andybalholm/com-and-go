package com

// generated by mkcomcall syscall.go iunknown.go idispatch.go

import (
	"syscall"
	"unsafe"
)

var _ unsafe.Pointer

var (
	modole32    = syscall.NewLazyDLL("ole32.dll")
	modoleaut32 = syscall.NewLazyDLL("oleaut32.dll")

	procCoInitialize     = modole32.NewProc("CoInitialize")
	procCLSIDFromProgID  = modole32.NewProc("CLSIDFromProgID")
	procCoCreateInstance = modole32.NewProc("CoCreateInstance")
	procStringFromGUID2  = modole32.NewProc("StringFromGUID2")
	procCLSIDFromString  = modole32.NewProc("CLSIDFromString")
	procSysAllocString   = modoleaut32.NewProc("SysAllocString")
)

func CoInitialize(reserved int) (err error) {
	_res, _, _ := procCoInitialize.Call(uintptr(reserved))
	if _res != 0 {
		err = HResult(_res)
	}
	return
}

func CLSIDFromProgID(progID string) (classID GUID, err error) {
	_res, _, _ := procCLSIDFromProgID.Call(uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(progID))),
		uintptr(unsafe.Pointer(&classID)))
	if _res != 0 {
		err = HResult(_res)
	}
	return
}

func CoCreateInstance(classID *GUID, outer *IUnknown, clsContext uint32, iid *GUID) (instance unsafe.Pointer, err error) {
	_res, _, _ := procCoCreateInstance.Call(uintptr(unsafe.Pointer(classID)),
		uintptr(unsafe.Pointer(outer)),
		uintptr(clsContext),
		uintptr(unsafe.Pointer(iid)),
		uintptr(unsafe.Pointer(&instance)))
	if _res != 0 {
		err = HResult(_res)
	}
	return
}

func StringFromGUID2(guid *GUID, str []uint16) (n int) {
	_res, _, _ := procStringFromGUID2.Call(uintptr(unsafe.Pointer(guid)),
		uintptr(unsafe.Pointer(&str[0])),
		uintptr(len(str)))
	n = int(_res)
	return
}

func CLSIDFromString(s string) (clsID GUID, err error) {
	_res, _, _ := procCLSIDFromString.Call(uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(s))),
		uintptr(unsafe.Pointer(&clsID)))
	if _res != 0 {
		err = HResult(_res)
	}
	return
}

func SysAllocString(s string) (bstr *uint16) {
	_res, _, _ := procSysAllocString.Call(uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(s))))
	bstr = (*uint16)(unsafe.Pointer(_res))
	return
}

var IID_IUnknown = NewGUID("{00000000-0000-0000-C000-000000000046}")

type IUnknown struct {
	*VTable
}

func (this *IUnknown) QueryInterface(iid *GUID) (object unsafe.Pointer, err error) {
	_res, _, _ := syscall.Syscall(this.VTable[0], 3,
		uintptr(unsafe.Pointer(this)),
		uintptr(unsafe.Pointer(iid)),
		uintptr(unsafe.Pointer(&object)))
	if _res != 0 {
		err = HResult(_res)
	}
	return
}

func (this *IUnknown) AddRef() (newCount int) {
	_res, _, _ := syscall.Syscall(this.VTable[1], 1,
		uintptr(unsafe.Pointer(this)),
		0,
		0)
	newCount = int(_res)
	return
}

func (this *IUnknown) Release() (newCount int) {
	_res, _, _ := syscall.Syscall(this.VTable[2], 1,
		uintptr(unsafe.Pointer(this)),
		0,
		0)
	newCount = int(_res)
	return
}

var IID_IDispatch = NewGUID("{00020400-0000-0000-C000-000000000046}")

type IDispatch struct {
	IUnknown
}

func (this *IDispatch) GetTypeInfoCount() (count uint32, err error) {
	_res, _, _ := syscall.Syscall(this.VTable[3], 2,
		uintptr(unsafe.Pointer(this)),
		uintptr(unsafe.Pointer(&count)),
		0)
	if _res != 0 {
		err = HResult(_res)
	}
	return
}

func (this *IDispatch) GetTypeInfo(i uint, localeID uint32) (info *ITypeInfo, err error) {
	_res, _, _ := syscall.Syscall6(this.VTable[4], 4,
		uintptr(unsafe.Pointer(this)),
		uintptr(i),
		uintptr(localeID),
		uintptr(unsafe.Pointer(&info)),
		0,
		0)
	if _res != 0 {
		err = HResult(_res)
	}
	return
}

func (this *IDispatch) GetIDsOfNames(iid *GUID, names []*uint16, localeID uint32, dispIDs *uint32) (err error) {
	_res, _, _ := syscall.Syscall6(this.VTable[5], 6,
		uintptr(unsafe.Pointer(this)),
		uintptr(unsafe.Pointer(iid)),
		uintptr(unsafe.Pointer(&names[0])),
		uintptr(len(names)),
		uintptr(localeID),
		uintptr(unsafe.Pointer(dispIDs)))
	if _res != 0 {
		err = HResult(_res)
	}
	return
}

func (this *IDispatch) Invoke(member uint32, iid *GUID, localeID uint32, flags uint16, params *DispParams) (result Variant, excepInfo ExcepInfo, argErr uint32, err error) {
	_res, _, _ := syscall.Syscall9(this.VTable[6], 9,
		uintptr(unsafe.Pointer(this)),
		uintptr(member),
		uintptr(unsafe.Pointer(iid)),
		uintptr(localeID),
		uintptr(flags),
		uintptr(unsafe.Pointer(params)),
		uintptr(unsafe.Pointer(&result)),
		uintptr(unsafe.Pointer(&excepInfo)),
		uintptr(unsafe.Pointer(&argErr)))
	if _res != 0 {
		err = HResult(_res)
	}
	return
}