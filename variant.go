package com

import (
	"fmt"
	"math"
	"math/big"
	"strings"
	"time"
	"unsafe"
)

type Vartype uint16

const (
	VT_EMPTY           Vartype = 0x0
	VT_NULL                    = 0x1
	VT_I2                      = 0x2
	VT_I4                      = 0x3
	VT_R4                      = 0x4
	VT_R8                      = 0x5
	VT_CY                      = 0x6
	VT_DATE                    = 0x7
	VT_BSTR                    = 0x8
	VT_DISPATCH                = 0x9
	VT_ERROR                   = 0xa
	VT_BOOL                    = 0xb
	VT_VARIANT                 = 0xc
	VT_UNKNOWN                 = 0xd
	VT_DECIMAL                 = 0xe
	VT_I1                      = 0x10
	VT_UI1                     = 0x11
	VT_UI2                     = 0x12
	VT_UI4                     = 0x13
	VT_I8                      = 0x14
	VT_UI8                     = 0x15
	VT_INT                     = 0x16
	VT_UINT                    = 0x17
	VT_VOID                    = 0x18
	VT_HRESULT                 = 0x19
	VT_PTR                     = 0x1a
	VT_SAFEARRAY               = 0x1b
	VT_CARRAY                  = 0x1c
	VT_USERDEFINED             = 0x1d
	VT_LPSTR                   = 0x1e
	VT_LPWSTR                  = 0x1f
	VT_RECORD                  = 0x24
	VT_INT_PTR                 = 0x25
	VT_UINT_PTR                = 0x26
	VT_FILETIME                = 0x40
	VT_BLOB                    = 0x41
	VT_STREAM                  = 0x42
	VT_STORAGE                 = 0x43
	VT_STREAMED_OBJECT         = 0x44
	VT_STORED_OBJECT           = 0x45
	VT_BLOB_OBJECT             = 0x46
	VT_CF                      = 0x47
	VT_CLSID                   = 0x48
	VT_BSTR_BLOB               = 0xfff
	VT_VECTOR                  = 0x1000
	VT_ARRAY                   = 0x2000
	VT_BYREF                   = 0x4000
	VT_RESERVED                = 0x8000
	VT_ILLEGAL                 = 0xffff
	VT_ILLEGALMASKED           = 0xfff
	VT_TYPEMASK                = 0xfff
)

type Variant struct {
	VT        Vartype
	Reserved1 uint16
	Reserved2 uint16
	Reserved3 uint16
	Val       uint64
}

// ToVariant returns x as a Variant. If x has an unsupported
// type, it panics.
func ToVariant(x interface{}) Variant {
	switch v := x.(type) {
	case nil:
		return Variant{VT: VT_NULL}
	case int16:
		return Variant{VT_I2, 0, 0, 0, uint64(v)}
	case *int16:
		return Variant{VT_I2 | VT_BYREF, 0, 0, 0, uint64(uintptr(unsafe.Pointer(v)))}
	case int32:
		return Variant{VT_I4, 0, 0, 0, uint64(v)}
	case *int32:
		return Variant{VT_I4 | VT_BYREF, 0, 0, 0, uint64(uintptr(unsafe.Pointer(v)))}
	case float32:
		return Variant{VT_R4, 0, 0, 0, uint64(math.Float32bits(v))}
	case *float32:
		return Variant{VT_R4 | VT_BYREF, 0, 0, 0, uint64(uintptr(unsafe.Pointer(v)))}
	case float64:
		return Variant{VT_R8, 0, 0, 0, math.Float64bits(v)}
	case *float64:
		return Variant{VT_R8 | VT_BYREF, 0, 0, 0, uint64(uintptr(unsafe.Pointer(v)))}
	case string:
		return Variant{VT_BSTR, 0, 0, 0, uint64(uintptr(unsafe.Pointer(BStrFromString(v).P)))}
	case BStr:
		return Variant{VT_BSTR, 0, 0, 0, uint64(uintptr(unsafe.Pointer(v.P)))}
	case *BStr:
		return Variant{VT_BSTR | VT_BYREF, 0, 0, 0, uint64(uintptr(unsafe.Pointer(v)))}
	case *IDispatch:
		return Variant{VT_DISPATCH, 0, 0, 0, uint64(uintptr(unsafe.Pointer(v)))}
	case **IDispatch:
		return Variant{VT_DISPATCH | VT_BYREF, 0, 0, 0, uint64(uintptr(unsafe.Pointer(v)))}
	case HResult:
		return Variant{VT_ERROR, 0, 0, 0, uint64(v)}
	case *HResult:
		return Variant{VT_ERROR | VT_BYREF, 0, 0, 0, uint64(uintptr(unsafe.Pointer(v)))}
	case bool:
		b := uint64(0)
		if v {
			b = 0xffff
		}
		return Variant{VT_BOOL, 0, 0, 0, b}
	case Variant:
		return v
	case *Variant:
		return Variant{VT_VARIANT | VT_BYREF, 0, 0, 0, uint64(uintptr(unsafe.Pointer(v)))}
	case *IUnknown:
		return Variant{VT_UNKNOWN, 0, 0, 0, uint64(uintptr(unsafe.Pointer(v)))}
	case **IUnknown:
		return Variant{VT_UNKNOWN | VT_BYREF, 0, 0, 0, uint64(uintptr(unsafe.Pointer(v)))}
	case int8:
		return Variant{VT_I1, 0, 0, 0, uint64(v)}
	case *int8:
		return Variant{VT_I1 | VT_BYREF, 0, 0, 0, uint64(uintptr(unsafe.Pointer(v)))}
	case uint8:
		return Variant{VT_UI1, 0, 0, 0, uint64(v)}
	case *uint8:
		return Variant{VT_UI1 | VT_BYREF, 0, 0, 0, uint64(uintptr(unsafe.Pointer(v)))}
	case uint16:
		return Variant{VT_UI2, 0, 0, 0, uint64(v)}
	case *uint16:
		return Variant{VT_UI2 | VT_BYREF, 0, 0, 0, uint64(uintptr(unsafe.Pointer(v)))}
	case uint32:
		return Variant{VT_UI4, 0, 0, 0, uint64(v)}
	case *uint32:
		return Variant{VT_UI4 | VT_BYREF, 0, 0, 0, uint64(uintptr(unsafe.Pointer(v)))}
	case int64:
		return Variant{VT_I8, 0, 0, 0, uint64(v)}
	case *int64:
		return Variant{VT_I8 | VT_BYREF, 0, 0, 0, uint64(uintptr(unsafe.Pointer(v)))}
	case uint64:
		return Variant{VT_UI8, 0, 0, 0, v}
	case *uint64:
		return Variant{VT_UI8 | VT_BYREF, 0, 0, 0, uint64(uintptr(unsafe.Pointer(v)))}
	case int:
		return Variant{VT_INT, 0, 0, 0, uint64(v)}
	case uint:
		return Variant{VT_UINT, 0, 0, 0, uint64(v)}
	case uintptr:
		return Variant{VT_UINT_PTR, 0, 0, 0, uint64(v)}
	case *uintptr:
		return Variant{VT_UINT_PTR | VT_BYREF, 0, 0, 0, uint64(uintptr(unsafe.Pointer(v)))}
	case time.Time:
		if v.IsZero() {
			return Variant{VT_DATE, 0, 0, 0, 0}
		}
		hour, min, sec := v.Clock()
		t := float64(hour*3600+min*60+sec) / 86400
		year, month, day := v.Date()
		if year == 0 && month == 1 && day == 1 {
			return Variant{VT_DATE, 0, 0, 0, math.Float64bits(t)}
		}
		// Find the number of days between this date and Dec. 30 1899.
		diff := time.Date(year, month, day, 0, 0, 0, 0, time.UTC).Sub(time.Date(1899, 12, 30, 0, 0, 0, 0, time.UTC))
		d := float64(diff / (24 * time.Hour))
		if d >= 0 {
			d += t
		} else {
			d -= t
		}
		return Variant{VT_DATE, 0, 0, 0, math.Float64bits(d)}
	}

	panic(fmt.Errorf("converting %T to Variant is not implemented", x))
}

// ToInterface returns v's value wrapped in an interface{} instead of in a
// Variant. If it can't convert the value, it wraps v in the interface{}
// instead.
func (v Variant) ToInterface() interface{} {
	// Avoid typing this conversion so many times.
	p := unsafe.Pointer(uintptr(v.Val))

	switch v.VT {
	case VT_NULL:
		return nil
	case VT_I2:
		return int16(v.Val)
	case VT_I2 | VT_BYREF:
		return (*int16)(p)
	case VT_I4:
		return int32(v.Val)
	case VT_I4 | VT_BYREF:
		return (*int32)(p)
	case VT_R4:
		return math.Float32frombits(uint32(v.Val))
	case VT_R4 | VT_BYREF:
		return (*float32)(p)
	case VT_R8:
		return math.Float64frombits(v.Val)
	case VT_R8 | VT_BYREF:
		return (*float64)(p)
	case VT_BSTR:
		b := BStr{(*uint16)(p)}
		s := b.String()
		SysFreeString(b)
		return s
	case VT_BSTR | VT_BYREF:
		return (*BStr)(p)
	case VT_DISPATCH:
		return (*IDispatch)(p)
	case VT_DISPATCH | VT_BYREF:
		return (**IDispatch)(p)
	case VT_ERROR:
		return HResult(v.Val)
	case VT_ERROR | VT_BYREF:
		return (*HResult)(p)
	case VT_BOOL:
		return v.Val != 0
	case VT_VARIANT | VT_BYREF:
		return (*Variant)(p)
	case VT_UNKNOWN:
		return (*IUnknown)(p)
	case VT_UNKNOWN | VT_BYREF:
		return (**IUnknown)(p)
	case VT_I1:
		return int8(v.Val)
	case VT_I1 | VT_BYREF:
		return (*int8)(p)
	case VT_UI1:
		return uint8(v.Val)
	case VT_UI1 | VT_BYREF:
		return (*uint8)(p)
	case VT_UI2:
		return uint16(v.Val)
	case VT_UI2 | VT_BYREF:
		return (*uint16)(p)
	case VT_UI4:
		return uint32(v.Val)
	case VT_UI4 | VT_BYREF:
		return (*uint32)(p)
	case VT_I8:
		return int64(v.Val)
	case VT_I8 | VT_BYREF:
		return (*int64)(p)
	case VT_UI8:
		return v.Val
	case VT_UI8 | VT_BYREF:
		return (*uint64)(p)
	case VT_INT:
		return int(int32(v.Val))
	case VT_UINT:
		return uint(v.Val)
	case VT_UINT_PTR:
		return uintptr(v.Val)
	case VT_UINT_PTR | VT_BYREF:
		return (*uintptr)(p)
	case VT_DECIMAL:
		d := (*Decimal)(unsafe.Pointer(&v))
		return *d
	case VT_DECIMAL | VT_BYREF:
		return (*Decimal)(p)
	case VT_DATE:
		// see http://blogs.msdn.com/b/ericlippert/archive/2003/09/16/eric-s-complete-guide-to-vt-date.aspx
		d, t := math.Modf(math.Float64frombits(uint64(v.Val)))
		t = math.Abs(t)
		if d == 0 {
			// We'll just have to hope that no one ever wants to actually refer to Dec. 30, 1899.
			if t == 0 {
				return time.Time{} // zero time
			}
			return time.Date(0, 1, 1, 0, 0, int(t*86400), 0, time.Local) // time without date
		}
		return time.Date(1899, 12, 30+int(d), 0, 0, int(t*86400), 0, time.Local)
	case VT_CY:
		return float64(int64(v.Val)) / 10000
	}

	return v
}

type Decimal struct {
	wReserved uint16
	Scale     byte
	Sign      byte
	Hi32      uint32
	Lo64      uint64
}

func (d Decimal) String() string {
	i := big.NewInt(0)
	i.SetUint64(d.Lo64)
	if d.Hi32 > 0 {
		hi := big.NewInt(int64(d.Hi32))
		two64 := big.NewInt(0)
		two64.SetBit(two64, 64, 1)
		hi.Mul(hi, two64)
		i.Add(i, hi)
	}
	s := i.String()
	if d.Scale > 0 {
		if zeroes := int(d.Scale) - len(s) + 1; zeroes > 0 {
			s = strings.Repeat("0", zeroes) + s
		}
		s = s[:len(s)-int(d.Scale)] + "." + s[len(s)-int(d.Scale):]
	}
	if d.Sign&0x80 == 0x80 {
		s = "-" + s
	}
	return s
}
