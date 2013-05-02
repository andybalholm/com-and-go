package com

// A VTable contains function pointers for a COM object's methods.
// The actual length varies; hopefully 1024 is enough to cover all we need.
type VTable [1024]uintptr

/* com
// IID {00000000-0000-0000-C000-000000000046}
type IUnknown interface {
	QueryInterface(iid *GUID) (object unsafe.Pointer, err error)
	AddRef() (newCount int)
	Release() (newCount int)
}
*/
