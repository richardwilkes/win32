package d2d

import (
	"syscall"
	"unsafe"

	"github.com/richardwilkes/win32/com"
)

// RectangleGeometryGUID holds the GUID for the RectangleGeometry class.
var RectangleGeometryGUID = com.GUID{Data1: 0x2cd906a2, Data2: 0x12e2, Data3: 0x11dc, Data4: [8]byte{0x9f, 0xed, 0x00, 0x11, 0x43, 0xa0, 0x55, 0xf9}}

// RectangleGeometry https://docs.microsoft.com/en-us/windows/win32/api/d2d1/nn-d2d1-id2d1rectanglegeometry
type RectangleGeometry struct {
	Geometry
}

type vmtRectangleGeometry struct {
	vmtGeometry
	GetRect uintptr
}

func (obj *RectangleGeometry) vmt() *vmtRectangleGeometry {
	return (*vmtRectangleGeometry)(obj.UnsafeVirtualMethodTable)
}

// Rect https://docs.microsoft.com/en-us/windows/win32/api/d2d1/nf-d2d1-id2d1rectanglegeometry-getrect
func (obj *RectangleGeometry) Rect() Rect {
	var rect Rect
	syscall.Syscall(obj.vmt().GetRect, 2, uintptr(unsafe.Pointer(obj)), uintptr(unsafe.Pointer(&rect)), 0)
	return rect
}
