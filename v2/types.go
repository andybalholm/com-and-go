package com

import (
	"fmt"
	"strings"
	"syscall"
	"unicode/utf16"
	"unsafe"
)

var (
	ole32    = LoadDLL("ole32.dll")
	oleaut32 = LoadDLL("oleaut32.dll")

	clsidFromProgID = ole32.Func("CLSIDFromProgID")
	clsidFromString = ole32.Func("CLSIDFromString")

	sysAllocStringLen = oleaut32.Func("SysAllocStringLen")
	sysFreeString     = oleaut32.Func("SysFreeString")
)

// http://msdn.microsoft.com/en-us/library/windows/desktop/aa373931.aspx
type GUID struct {
	Data1 uint32
	Data2 uint16
	Data3 uint16
	Data4 [8]byte
}

func (g *GUID) String() string {
	return fmt.Sprintf("{%08X-%04X-%04X-%X-%X}", g.Data1, g.Data2, g.Data3, g.Data4[:2], g.Data4[2:])
}

func CLSIDFromString(s string) (clsid *GUID, err error) {
	clsid = new(GUID)
	args := &struct {
		lpsz   *uint16
		pclsid *GUID
	}{WideString(s), clsid}
	err = clsidFromString.CallHR(unsafe.Pointer(args), 2)
	return
}

func CLSIDFromProgID(s string) (clsid *GUID, err error) {
	clsid = new(GUID)
	args := &struct {
		lpszProgID *uint16
		lpclsid    *GUID
	}{WideString(s), clsid}
	err = clsidFromProgID.CallHR(unsafe.Pointer(args), 2)
	return
}

func NewGUID(s string) *GUID {
	g, err := CLSIDFromString(s)
	if err != nil {
		panic(fmt.Errorf("error parsing GUID (%q): %s", s, err))
	}
	return g
}

type HResult uint32

func (hr HResult) Error() string {
	buf := make([]uint16, 300)
	n, err := syscall.FormatMessage(syscall.FORMAT_MESSAGE_FROM_SYSTEM|syscall.FORMAT_MESSAGE_ARGUMENT_ARRAY|syscall.FORMAT_MESSAGE_IGNORE_INSERTS,
		0, uint32(hr), 0, buf, nil)
	if err != nil {
		return fmt.Sprintf("COM error %08x", uint32(hr))
	}
	return strings.TrimSpace(string(utf16.Decode(buf[:n])))
}

// A BStr is a UTF-16 string with an null terminator and an explicit length.
// The length in bytes (not code units) is stored before the first code unit,
// as a 32-bit integer.
// P points at the first code unit.
type BStr struct {
	P *uint16
}

func SysAllocStringLen(strIn *uint16, ui int) BStr {
	r, _, _ := sysAllocStringLen.Call(unsafe.Pointer(&strIn), 2)
	return BStr{(*uint16)(unsafe.Pointer(r))}
}

func (b BStr) Free() {
	sysFreeString.Call(unsafe.Pointer(&b), 1)
}

// wideString converts s to a UTF-16 string. It will be terminated with a null
// character if terminate is true.
func wideString(s string, terminate bool) []uint16 {
	n := 0
	for _, c := range s {
		n++
		if c >= 0x10000 {
			n++
		}
	}
	if terminate {
		n++
	}

	a := make([]uint16, n)
	i := 0
	for _, c := range s {
		if c < 0x10000 {
			a[i] = uint16(c)
			i++
		} else {
			r1, r2 := utf16.EncodeRune(c)
			a[i] = uint16(r1)
			a[i+1] = uint16(r2)
			i += 2
		}
	}

	return a
}

// ToBStr returns s as a BStr (OLE Automation string). The result is allocated
// with SysAllocStringLen, so it should be freed with SysFreeString.
func ToBStr(s string) BStr {
	if s == "" {
		return BStr{nil}
	}
	ws := wideString(s, false)
	return SysAllocStringLen(&ws[0], len(ws))
}

func (b BStr) String() string {
	if b.P == nil {
		return ""
	}
	n := *(*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer(b.P)) - 4)) / 2
	a := (*[1 << 29]uint16)(unsafe.Pointer(b.P))[:n]
	return string(utf16.Decode(a))
}

// WideString returns s as a wide C string.
func WideString(s string) *uint16 {
	ws := wideString(s, true)
	return &ws[0]
}

// GoString returns s as a UTF-8 Go string.
func GoString(s *uint16) string {
	s1 := (*[1 << 29]uint16)(unsafe.Pointer(s))
	for i, c := range s1 {
		if c == 0 {
			return string(utf16.Decode(s1[:i]))
		}
	}
	panic("no null terminator found")
}
