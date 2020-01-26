package d2d

import (
	"syscall"
	"unsafe"

	"github.com/richardwilkes/win32/com"
)

// StrokeStyleGUID holds the GUID for the StrokeStyle class.
var StrokeStyleGUID = com.GUID{Data1: 0x2cd9069d, Data2: 0x12e2, Data3: 0x11dc, Data4: [8]byte{0x9f, 0xed, 0x00, 0x11, 0x43, 0xa0, 0x55, 0xf9}}

// StrokeStyle https://docs.microsoft.com/en-us/windows/win32/api/d2d1/nn-d2d1-id2d1strokestyle
type StrokeStyle struct {
	Resource
}

type vmtStrokeStyle struct {
	vmtResource
	GetStartCap    uintptr
	GetEndCap      uintptr
	GetDashCap     uintptr
	GetMiterLimit  uintptr
	GetLineJoin    uintptr
	GetDashOffset  uintptr
	GetDashStyle   uintptr
	GetDashesCount uintptr
	GetDashes      uintptr
}

func (obj *StrokeStyle) vmt() *vmtStrokeStyle {
	return (*vmtStrokeStyle)(obj.UnsafeVirtualMethodTable)
}

// StartCap https://docs.microsoft.com/en-us/windows/win32/api/d2d1/nf-d2d1-id2d1strokestyle-getstartcap
func (obj *StrokeStyle) StartCap() CapStyle {
	ret, _, _ := syscall.Syscall(obj.vmt().GetStartCap, 1, uintptr(unsafe.Pointer(obj)), 0, 0)
	return CapStyle(ret)
}

// EndCap https://docs.microsoft.com/en-us/windows/win32/api/d2d1/nf-d2d1-id2d1strokestyle-getendcap
func (obj *StrokeStyle) EndCap() CapStyle {
	ret, _, _ := syscall.Syscall(obj.vmt().GetEndCap, 1, uintptr(unsafe.Pointer(obj)), 0, 0)
	return CapStyle(ret)
}

// DashCap https://docs.microsoft.com/en-us/windows/win32/api/d2d1/nf-d2d1-id2d1strokestyle-getdashcap
func (obj *StrokeStyle) DashCap() CapStyle {
	ret, _, _ := syscall.Syscall(obj.vmt().GetDashCap, 1, uintptr(unsafe.Pointer(obj)), 0, 0)
	return CapStyle(ret)
}

// MiterLimit https://docs.microsoft.com/en-us/windows/win32/api/d2d1/nf-d2d1-id2d1strokestyle-getmiterlimit
func (obj *StrokeStyle) MiterLimit() float32 {
	ret, _, _ := syscall.Syscall(obj.vmt().GetMiterLimit, 1, uintptr(unsafe.Pointer(obj)), 0, 0)
	return *(*float32)(unsafe.Pointer(ret))
}

// LineJoin https://docs.microsoft.com/en-us/windows/win32/api/d2d1/nf-d2d1-id2d1strokestyle-getlinejoin
func (obj *StrokeStyle) LineJoin() LineJoin {
	ret, _, _ := syscall.Syscall(obj.vmt().GetLineJoin, 1, uintptr(unsafe.Pointer(obj)), 0, 0)
	return LineJoin(ret)
}

// DashOffset https://docs.microsoft.com/en-us/windows/win32/api/d2d1/nf-d2d1-id2d1strokestyle-getdashoffset
func (obj *StrokeStyle) DashOffset() float32 {
	ret, _, _ := syscall.Syscall(obj.vmt().GetDashOffset, 1, uintptr(unsafe.Pointer(obj)), 0, 0)
	return *(*float32)(unsafe.Pointer(ret))
}

// DashStyle https://docs.microsoft.com/en-us/windows/win32/api/d2d1/nf-d2d1-id2d1strokestyle-getdashstyle
func (obj *StrokeStyle) DashStyle() DashStyle {
	ret, _, _ := syscall.Syscall(obj.vmt().GetDashStyle, 1, uintptr(unsafe.Pointer(obj)), 0, 0)
	return DashStyle(ret)
}

// DashesCount https://docs.microsoft.com/en-us/windows/win32/api/d2d1/nf-d2d1-id2d1strokestyle-getdashescount
func (obj *StrokeStyle) DashesCount() uint32 {
	ret, _, _ := syscall.Syscall(obj.vmt().GetDashesCount, 1, uintptr(unsafe.Pointer(obj)), 0, 0)
	return uint32(ret)
}

// Dashes https://docs.microsoft.com/en-us/windows/win32/api/d2d1/nf-d2d1-id2d1strokestyle-getdashes
func (obj *StrokeStyle) Dashes(dashes []float32) {
	syscall.Syscall(obj.vmt().GetDashes, 3, uintptr(unsafe.Pointer(obj)), uintptr(unsafe.Pointer(&(dashes[0]))), uintptr(len(dashes)))
}
