package d2d

import (
	"fmt"
	"syscall"
	"unsafe"

	"github.com/richardwilkes/win32"
	"github.com/richardwilkes/win32/com"
)

// SimplifiedGeometrySinkGUID holds the GUID for the SimplifiedGeometrySink class.
var SimplifiedGeometrySinkGUID = com.GUID{Data1: 0x2cd9069e, Data2: 0x12e2, Data3: 0x11dc, Data4: [8]byte{0x9f, 0xed, 0x00, 0x11, 0x43, 0xa0, 0x55, 0xf9}}

// SimplifiedGeometrySink https://docs.microsoft.com/en-us/windows/win32/api/d2d1/nn-d2d1-id2d1simplifiedgeometrysink
type SimplifiedGeometrySink struct {
	com.IUnknown
}

type vmtSimplifiedGeometrySink struct {
	com.VMTIUnknown
	SetFillMode     uintptr
	SetSegmentFlags uintptr
	BeginFigure     uintptr
	AddLines        uintptr
	AddBeziers      uintptr
	EndFigure       uintptr
	Close           uintptr
}

func (obj *SimplifiedGeometrySink) vmt() *vmtSimplifiedGeometrySink {
	return (*vmtSimplifiedGeometrySink)(obj.UnsafeVirtualMethodTable)
}

// SetFillMode https://docs.microsoft.com/en-us/windows/win32/api/d2d1/nf-d2d1-id2d1simplifiedgeometrysink-setfillmode
func (obj *SimplifiedGeometrySink) SetFillMode(winding bool) {
	var fillMode uint32
	if winding {
		fillMode = 1
	}
	syscall.Syscall(obj.vmt().SetFillMode, 2, uintptr(unsafe.Pointer(obj)), uintptr(fillMode), 0)
}

// SetSegmentFlags https://docs.microsoft.com/en-us/windows/win32/api/d2d1/nf-d2d1-id2d1simplifiedgeometrysink-setsegmentflags
func (obj *SimplifiedGeometrySink) SetSegmentFlags(vertexFlags PathSegment) {
	syscall.Syscall(obj.vmt().SetSegmentFlags, 2, uintptr(unsafe.Pointer(obj)), uintptr(vertexFlags), 0)
}

// BeginFigure https://docs.microsoft.com/en-us/windows/win32/api/d2d1/nf-d2d1-id2d1simplifiedgeometrysink-beginfigure
// Must be balanced with a call to EndFigure.
func (obj *SimplifiedGeometrySink) BeginFigure(startPoint Point, line bool) {
	var figureBegin uint32
	if line {
		figureBegin = 1
	}
	syscall.Syscall(obj.vmt().BeginFigure, 3, uintptr(unsafe.Pointer(obj)), uintptr(*(*uint64)(unsafe.Pointer(&startPoint))), uintptr(figureBegin))
}

// AddLines https://docs.microsoft.com/en-us/windows/win32/api/d2d1/nf-d2d1-id2d1simplifiedgeometrysink-addlines
func (obj *SimplifiedGeometrySink) AddLines(points []Point) {
	syscall.Syscall(obj.vmt().AddLines, 3, uintptr(unsafe.Pointer(obj)), uintptr(unsafe.Pointer(&(points[0]))), uintptr(len(points)))
}

// AddBeziers https://docs.microsoft.com/en-us/windows/win32/api/d2d1/nf-d2d1-id2d1simplifiedgeometrysink-addbeziers
func (obj *SimplifiedGeometrySink) AddBeziers(beziers []BezierSegment) {
	syscall.Syscall(obj.vmt().AddBeziers, 3, uintptr(unsafe.Pointer(obj)), uintptr(unsafe.Pointer(&(beziers[0]))), uintptr(len(beziers)))
}

// EndFigure https://docs.microsoft.com/en-us/windows/win32/api/d2d1/nf-d2d1-id2d1simplifiedgeometrysink-endfigure
func (obj *SimplifiedGeometrySink) EndFigure(close bool) {
	var figureEnd uint32
	if close {
		figureEnd = 1
	}
	syscall.Syscall(obj.vmt().EndFigure, 2, uintptr(unsafe.Pointer(obj)), uintptr(figureEnd), 0)
}

// Close https://docs.microsoft.com/en-us/windows/win32/api/d2d1/nf-d2d1-id2d1simplifiedgeometrysink-close
// Must be called if a figure is still in progress.
func (obj *SimplifiedGeometrySink) Close() error {
	if ret, _, _ := syscall.Syscall(obj.vmt().Close, 1, uintptr(unsafe.Pointer(obj)), 0, 0); ret != win32.S_OK {
		return fmt.Errorf("call to Close failed: %#x", ret)
	}
	return nil
}
