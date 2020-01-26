package d2d

import (
	"fmt"
	"syscall"
	"unsafe"

	"github.com/richardwilkes/win32"
	"github.com/richardwilkes/win32/com"
)

// PathGeometryGUID holds the GUID for teh PathGeometry class.
var PathGeometryGUID = com.GUID{Data1: 0x2cd906a5, Data2: 0x12e2, Data3: 0x11dc, Data4: [8]byte{0x9f, 0xed, 0x00, 0x11, 0x43, 0xa0, 0x55, 0xf9}}

// PathGeometry https://docs.microsoft.com/en-us/windows/win32/api/d2d1/nn-d2d1-id2d1pathgeometry
type PathGeometry struct {
	Geometry
}

type vmtPathGeometry struct {
	vmtGeometry
	Open            uintptr
	Stream          uintptr
	GetSegmentCount uintptr
	GetFigureCount  uintptr
}

func (obj *PathGeometry) vmt() *vmtPathGeometry {
	return (*vmtPathGeometry)(obj.UnsafeVirtualMethodTable)
}

// Open https://docs.microsoft.com/en-us/windows/win32/api/d2d1/nf-d2d1-id2d1pathgeometry-open
// May only be called once.
func (obj *PathGeometry) Open() (*GeometrySink, error) {
	var geometrySink *GeometrySink
	if ret, _, _ := syscall.Syscall(obj.vmt().Open, 2, uintptr(unsafe.Pointer(obj)), uintptr(unsafe.Pointer(&geometrySink)), 0); ret != win32.S_OK {
		return nil, fmt.Errorf("call to Open failed: %#x", ret)
	}
	return geometrySink, nil
}

// Stream https://docs.microsoft.com/en-us/windows/win32/api/d2d1/nf-d2d1-id2d1pathgeometry-stream
func (obj *PathGeometry) Stream(geometrySink *GeometrySink) error {
	if ret, _, _ := syscall.Syscall(obj.vmt().Stream, 2, uintptr(unsafe.Pointer(obj)), uintptr(unsafe.Pointer(geometrySink)), 0); ret != win32.S_OK {
		return fmt.Errorf("call to Stream failed: %#x", ret)
	}
	return nil
}

// SegmentCount https://docs.microsoft.com/en-us/windows/win32/api/d2d1/nf-d2d1-id2d1pathgeometry-getsegmentcount
func (obj *PathGeometry) SegmentCount() uint32 {
	var count uint32
	syscall.Syscall(obj.vmt().GetSegmentCount, 2, uintptr(unsafe.Pointer(obj)), uintptr(unsafe.Pointer(&count)), 0)
	return count
}

// FigureCount https://docs.microsoft.com/en-us/windows/win32/api/d2d1/nf-d2d1-id2d1pathgeometry-getfigurecount
func (obj *PathGeometry) FigureCount() uint32 {
	var count uint32
	syscall.Syscall(obj.vmt().GetFigureCount, 2, uintptr(unsafe.Pointer(obj)), uintptr(unsafe.Pointer(&count)), 0)
	return count
}
