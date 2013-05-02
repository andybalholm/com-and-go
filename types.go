package com

import (
	"fmt"
	"strings"
	"syscall"
	"unicode/utf16"
)

// http://msdn.microsoft.com/en-us/library/windows/desktop/aa373931.aspx
type GUID struct {
	Data1 uint32
	Data2 uint16
	Data3 uint16
	Data4 [8]byte
}

func (g *GUID) String()string {
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

type HResult int32

func (hr HResult) Error() string {
	buf := make([]uint16, 300)
	n, err := syscall.FormatMessage(syscall.FORMAT_MESSAGE_FROM_SYSTEM|syscall.FORMAT_MESSAGE_ARGUMENT_ARRAY|syscall.FORMAT_MESSAGE_IGNORE_INSERTS,
		0, uint32(hr), 0, buf, nil)
	if err != nil {
		return fmt.Sprintf("COM error %08x", uint32(hr))
	}
	return strings.TrimSpace(string(utf16.Decode(buf[:n])))
}
