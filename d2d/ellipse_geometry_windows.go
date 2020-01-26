package d2d

import (
	"syscall"
	"unsafe"

	"github.com/richardwilkes/win32/com"
)

// EllipseGeometryGUID holds the GUID for the EllipseGeometry class.
var EllipseGeometryGUID = com.GUID{Data1: 0x2cd906a4, Data2: 0x12e2, Data3: 0x11dc, Data4: [8]byte{0x9f, 0xed, 0x00, 0x11, 0x43, 0xa0, 0x55, 0xf9}}

// EllipseGeometry https://docs.microsoft.com/en-us/windows/win32/api/d2d1/nn-d2d1-id2d1ellipsegeometry
type EllipseGeometry struct {
	Geometry
}

type vmtEllipseGeometry struct {
	vmtGeometry
	GetEllipse uintptr
}

func (obj *EllipseGeometry) vmt() *vmtEllipseGeometry {
	return (*vmtEllipseGeometry)(obj.UnsafeVirtualMethodTable)
}

// Ellipse https://docs.microsoft.com/en-us/windows/win32/api/d2d1/nf-d2d1-id2d1ellipsegeometry-getellipse
func (obj *EllipseGeometry) Ellipse() Ellipse {
	var ellipse Ellipse
	syscall.Syscall(obj.vmt().GetEllipse, 2, uintptr(unsafe.Pointer(obj)), uintptr(unsafe.Pointer(&ellipse)), 0)
	return ellipse
}
