package com

import (
	"unicode/utf16"
	"unsafe"
)

func UTF16PtrToString(s *uint16) string {
	s1 := (*[1 << 29]uint16)(unsafe.Pointer(s))
	for i, c := range s1 {
		if c == 0 {
			return string(utf16.Decode(s1[:i]))
		}
	}
	panic("no null terminator found")
}
