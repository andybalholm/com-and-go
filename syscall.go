package com

/* com
func(ole32) CoInitialize(reserved int) (err error)
func(ole32) CoInitializeEx(reserved uintptr, coInit int) (err error)
func(ole32) CLSIDFromProgID(progID string) (classID GUID, err error)
func(ole32) CoCreateInstance(classID *GUID, outer *IUnknown, clsContext uint32, iid *GUID) (instance unsafe.Pointer, err error)
func(ole32) StringFromGUID2(guid *GUID, str []uint16) (n int)
func(ole32) CLSIDFromString(s string) (clsID GUID, err error)
func(oleaut32) SysAllocString(s string) (bstr *uint16)
func(oleaut32) SysFreeString(s BStr)
*/
