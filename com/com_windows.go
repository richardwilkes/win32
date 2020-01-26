package com

import (
	"syscall"
	"unsafe"

	"github.com/richardwilkes/win32"
)

// GUID https://docs.microsoft.com/en-us/windows/win32/api/guiddef/ns-guiddef-guid
type GUID struct {
	Data1 uint32
	Data2 uint16
	Data3 uint16
	Data4 [8]byte
}

// IUnknown https://docs.microsoft.com/en-us/windows/win32/api/unknwn/nn-unknwn-iunknown
type IUnknown struct {
	UnsafeVirtualMethodTable unsafe.Pointer
}

// VMTIUnknown holds the virtual dispatch table entries for IUnknown.
type VMTIUnknown struct {
	QueryInterface uintptr
	AddRef         uintptr
	Release        uintptr
}

func (obj *IUnknown) vmt() *VMTIUnknown {
	return (*VMTIUnknown)(obj.UnsafeVirtualMethodTable)
}

// QueryInterface https://docs.microsoft.com/en-us/windows/win32/api/unknwn/nf-unknwn-iunknown-queryinterface(refiid_void)
func (obj *IUnknown) QueryInterface(guid *GUID) unsafe.Pointer {
	var dest unsafe.Pointer
	if ret, _, _ := syscall.Syscall(obj.vmt().QueryInterface, 3, uintptr(unsafe.Pointer(obj)), uintptr(unsafe.Pointer(guid)), uintptr(unsafe.Pointer(&dest))); ret != win32.S_OK {
		return nil
	}
	return dest
}

// AddRef https://docs.microsoft.com/en-us/windows/win32/api/unknwn/nf-unknwn-iunknown-addref
func (obj *IUnknown) AddRef() {
	syscall.Syscall(obj.vmt().AddRef, 1, uintptr(unsafe.Pointer(obj)), 0, 0)
}

// Release https://docs.microsoft.com/en-us/windows/win32/api/unknwn/nf-unknwn-iunknown-release
func (obj *IUnknown) Release() {
	syscall.Syscall(obj.vmt().Release, 1, uintptr(unsafe.Pointer(obj)), 0, 0)
}
