package d2d

import (
	"syscall"
	"unsafe"

	"github.com/richardwilkes/win32/com"
)

// GeometrySinkGUID holds the GUID for the GeometrySink class.
var GeometrySinkGUID = com.GUID{Data1: 0x2cd9069f, Data2: 0x12e2, Data3: 0x11dc, Data4: [8]byte{0x9f, 0xed, 0x00, 0x11, 0x43, 0xa0, 0x55, 0xf9}}

// GeometrySink https://docs.microsoft.com/en-us/windows/win32/api/d2d1/nn-d2d1-id2d1geometrysink
type GeometrySink struct {
	SimplifiedGeometrySink
}

type vmtGeometrySink struct {
	vmtSimplifiedGeometrySink
	AddLine             uintptr
	AddBezier           uintptr
	AddQuadraticBezier  uintptr
	AddQuadraticBeziers uintptr
	AddArc              uintptr
}

func (obj *GeometrySink) vmt() *vmtGeometrySink {
	return (*vmtGeometrySink)(obj.UnsafeVirtualMethodTable)
}

// AddLine https://docs.microsoft.com/en-us/windows/win32/api/d2d1/nf-d2d1-id2d1geometrysink-addline
func (obj *GeometrySink) AddLine(point Point) {
	syscall.Syscall(obj.vmt().AddLine, 2, uintptr(unsafe.Pointer(obj)), uintptr(*(*uint64)(unsafe.Pointer(&point))), 0)
}

// AddBezier https://docs.microsoft.com/en-us/windows/win32/api/d2d1/nf-d2d1-id2d1geometrysink-addbezier%28constd2d1_bezier_segment%29
func (obj *GeometrySink) AddBezier(bezier BezierSegment) {
	syscall.Syscall(obj.vmt().AddBezier, 2, uintptr(unsafe.Pointer(obj)), uintptr(unsafe.Pointer(&bezier)), 0)
}

// AddQuadraticBezier https://docs.microsoft.com/en-us/windows/win32/api/d2d1/nf-d2d1-id2d1geometrysink-addquadraticbezier%28constd2d1_quadratic_bezier_segment%29
func (obj *GeometrySink) AddQuadraticBezier(bezier QuadraticBezierSegment) {
	syscall.Syscall(obj.vmt().AddQuadraticBezier, 2, uintptr(unsafe.Pointer(obj)), uintptr(unsafe.Pointer(&bezier)), 0)
}

// AddQuadraticBeziers https://docs.microsoft.com/en-us/windows/win32/api/d2d1/nf-d2d1-id2d1geometrysink-addquadraticbeziers
func (obj *GeometrySink) AddQuadraticBeziers(beziers []QuadraticBezierSegment) {
	syscall.Syscall(obj.vmt().AddQuadraticBeziers, 3, uintptr(unsafe.Pointer(obj)), uintptr(unsafe.Pointer(&(beziers[0]))), uintptr(len(beziers)))
}

// AddArc https://docs.microsoft.com/en-us/windows/win32/api/d2d1/nf-d2d1-id2d1geometrysink-addarc%28constd2d1_arc_segment%29
func (obj *GeometrySink) AddArc(arc ArcSegment) {
	syscall.Syscall(obj.vmt().AddArc, 2, uintptr(unsafe.Pointer(obj)), uintptr(unsafe.Pointer(&arc)), 0)
}
