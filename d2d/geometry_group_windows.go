package d2d

import (
	"syscall"
	"unsafe"

	"github.com/richardwilkes/win32/com"
)

// GeometryGroupGUID holds the GUID for the GeometryGroup class.
var GeometryGroupGUID = com.GUID{Data1: 0x2cd906a6, Data2: 0x12e2, Data3: 0x11dc, Data4: [8]byte{0x9f, 0xed, 0x00, 0x11, 0x43, 0xa0, 0x55, 0xf9}}

// GeometryGroup https://docs.microsoft.com/en-us/windows/win32/api/d2d1/nn-d2d1-id2d1geometrygroup
type GeometryGroup struct {
	Geometry
}

type vmtGeometryGroup struct {
	vmtGeometry
	GetFillMode            uintptr
	GetSourceGeometryCount uintptr
	GetSourceGeometries    uintptr
}

func (obj *GeometryGroup) vmt() *vmtGeometryGroup {
	return (*vmtGeometryGroup)(obj.UnsafeVirtualMethodTable)
}

// IsWinding https://docs.microsoft.com/en-us/windows/win32/api/d2d1/nf-d2d1-id2d1geometrygroup-getfillmode
func (obj *GeometryGroup) IsWindingFillMode() bool {
	ret, _, _ := syscall.Syscall(obj.vmt().GetFillMode, 1, uintptr(unsafe.Pointer(obj)), 0, 0)
	return ret == 1
}

// IsAlternating https://docs.microsoft.com/en-us/windows/win32/api/d2d1/nf-d2d1-id2d1geometrygroup-getfillmode
func (obj *GeometryGroup) IsAlternatingFillMode() bool {
	ret, _, _ := syscall.Syscall(obj.vmt().GetFillMode, 1, uintptr(unsafe.Pointer(obj)), 0, 0)
	return ret == 0
}

// SourceGeometryCount https://docs.microsoft.com/en-us/windows/win32/api/d2d1/nf-d2d1-id2d1geometrygroup-getsourcegeometrycount
func (obj *GeometryGroup) SourceGeometryCount() uint32 {
	ret, _, _ := syscall.Syscall(obj.vmt().GetSourceGeometryCount, 1, uintptr(unsafe.Pointer(obj)), 0, 0)
	return uint32(ret)
}

// SourceGeometries https://docs.microsoft.com/en-us/windows/win32/api/d2d1/nf-d2d1-id2d1geometrygroup-getsourcegeometries
func (obj *GeometryGroup) SourceGeometries(geometries []*Geometry) {
	syscall.Syscall(obj.vmt().GetSourceGeometries, 3, uintptr(unsafe.Pointer(obj)), uintptr(unsafe.Pointer(&(geometries[0]))), uintptr(len(geometries)))
}
