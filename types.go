package com

import (
	"fmt"
	"strings"
	"syscall"
	"unicode/utf16"
	"unsafe"
)

// http://msdn.microsoft.com/en-us/library/windows/desktop/aa373931.aspx
type GUID struct {
	Data1 uint32
	Data2 uint16
	Data3 uint16
	Data4 [8]byte
}

func (g *GUID) String() string {
	s := make([]uint16, 40)
	n := StringFromGUID2(g, s)
	return string(utf16.Decode(s[:n-1]))
}

func NewGUID(s string) *GUID {
	g, err := CLSIDFromString(s)
	if err != nil {
		panic(fmt.Errorf("error parsing GUID (%q): %s", s, err))
	}
	return &g
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

func BStrFromString(s string) BStr {
	n := 0
	for _, c := range s {
		n++
		if c >= 0x10000 {
			n++
		}
	}

	a := make([]uint16, n+3) // 2 words for the length + one for the terminator
	i := 2
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

	byteLen := n * 2
	a[0] = uint16(byteLen)
	a[1] = uint16(byteLen >> 16)

	return BStr{&a[2]}
}

func (b BStr) String() string {
	n := *(*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer(b.P)) - 4)) / 2
	a := (*[1 << 29]uint16)(unsafe.Pointer(b.P))[:n]
	return string(utf16.Decode(a))
}

func VariantBool(x bool) uintptr {
	if x {
		return ^uintptr(0)
	}
	return 0
}
