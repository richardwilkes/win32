package d2d

import (
	"syscall"
	"unsafe"

	"github.com/richardwilkes/win32/com"
)

// RoundedRectangleGeometryGUID holds the GUID for the RoundedRectangleGeometry class.
var RoundedRectangleGeometryGUID = com.GUID{Data1: 0x2cd906a3, Data2: 0x12e2, Data3: 0x11dc, Data4: [8]byte{0x9f, 0xed, 0x00, 0x11, 0x43, 0xa0, 0x55, 0xf9}}

// RoundedRectangleGeometry https://docs.microsoft.com/en-us/windows/win32/api/d2d1/nn-d2d1-id2d1roundedrectanglegeometry
type RoundedRectangleGeometry struct {
	Geometry
}

type vmtRoundedRectangleGeometry struct {
	vmtGeometry
	GetRoundedRect uintptr
}

func (obj *RoundedRectangleGeometry) vmt() *vmtRoundedRectangleGeometry {
	return (*vmtRoundedRectangleGeometry)(obj.UnsafeVirtualMethodTable)
}

// RoundedRect https://docs.microsoft.com/en-us/windows/win32/api/d2d1/nf-d2d1-id2d1roundedrectanglegeometry-getroundedrect
func (obj *RoundedRectangleGeometry) RoundedRect() RoundedRect {
	var roundedRect RoundedRect
	syscall.Syscall(obj.vmt().GetRoundedRect, 2, uintptr(unsafe.Pointer(obj)), uintptr(unsafe.Pointer(&roundedRect)), 0)
	return roundedRect
}
