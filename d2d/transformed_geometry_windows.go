package d2d

import (
	"syscall"
	"unsafe"

	"github.com/richardwilkes/win32/com"
)

// TransformedGeometryGUID holds the GUID for the TransformedGeometry class.
var TransformedGeometryGUID = com.GUID{Data1: 0x2cd906bb, Data2: 0x12e2, Data3: 0x11dc, Data4: [8]byte{0x9f, 0xed, 0x00, 0x11, 0x43, 0xa0, 0x55, 0xf9}}

// TransformedGeometry https://docs.microsoft.com/en-us/windows/win32/api/d2d1/nn-d2d1-id2d1transformedgeometry
type TransformedGeometry struct {
	Geometry
}

type vmtTransformedGeometry struct {
	vmtGeometry
	GetSourceGeometry uintptr
	GetTransform      uintptr
}

func (obj *TransformedGeometry) vmt() *vmtTransformedGeometry {
	return (*vmtTransformedGeometry)(obj.UnsafeVirtualMethodTable)
}

// SourceGeometry https://docs.microsoft.com/en-us/windows/win32/api/d2d1/nf-d2d1-id2d1transformedgeometry-getsourcegeometry
func (obj *TransformedGeometry) SourceGeometry() *Geometry {
	var sourceGeometry *Geometry
	syscall.Syscall(obj.vmt().GetSourceGeometry, 2, uintptr(unsafe.Pointer(obj)), uintptr(unsafe.Pointer(&sourceGeometry)), 0)
	return sourceGeometry
}

// Transform https://docs.microsoft.com/en-us/windows/win32/api/d2d1/nf-d2d1-id2d1transformedgeometry-gettransform
func (obj *TransformedGeometry) Transform() *Matrix3x2 {
	var transform Matrix3x2
	syscall.Syscall(obj.vmt().GetTransform, 2, uintptr(unsafe.Pointer(obj)), uintptr(unsafe.Pointer(&transform)), 0)
	return &transform
}
