package d2d

import (
	"syscall"
	"unsafe"

	"github.com/richardwilkes/win32/com"
)

// ResourceGUID holds the GUID for the Resource class.
var ResourceGUID = com.GUID{Data1: 0x2cd90691, Data2: 0x12e2, Data3: 0x11dc, Data4: [8]byte{0x9f, 0xed, 0x00, 0x11, 0x43, 0xa0, 0x55, 0xf9}}

// Resource https://docs.microsoft.com/en-us/windows/win32/api/d2d1/nn-d2d1-id2d1resource
type Resource struct {
	com.IUnknown
}

type vmtResource struct {
	com.VMTIUnknown
	GetFactory uintptr
}

func (obj *Resource) vmt() *vmtResource {
	return (*vmtResource)(obj.UnsafeVirtualMethodTable)
}

// Factory https://docs.microsoft.com/en-us/windows/win32/api/d2d1/nf-d2d1-id2d1resource-getfactory
func (obj *Resource) Factory() *Factory {
	var factory *Factory
	syscall.Syscall(obj.vmt().GetFactory, 2, uintptr(unsafe.Pointer(obj)), uintptr(unsafe.Pointer(&factory)), 0)
	return factory
}
