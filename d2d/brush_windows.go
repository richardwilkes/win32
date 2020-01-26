package d2d

import (
	"syscall"
	"unsafe"

	"github.com/richardwilkes/win32/com"
)

// BrushGUID holds the GUID for the Brush class.
var BrushGUID = com.GUID{Data1: 0x2cd906a8, Data2: 0x12e2, Data3: 0x11dc, Data4: [8]byte{0x9f, 0xed, 0x00, 0x11, 0x43, 0xa0, 0x55, 0xf9}}

// Brush https://docs.microsoft.com/en-us/windows/win32/api/d2d1/nn-d2d1-id2d1brush
type Brush struct {
	Resource
}

type vmtBrush struct {
	vmtResource
	SetOpacity   uintptr
	SetTransform uintptr
	GetOpacity   uintptr
	GetTransform uintptr
}

func (obj *Brush) vmt() *vmtBrush {
	return (*vmtBrush)(obj.UnsafeVirtualMethodTable)
}

// Opacity https://docs.microsoft.com/en-us/windows/win32/api/d2d1/nf-d2d1-id2d1brush-getopacity
func (obj *Brush) Opacity() float32 {
	var ret, _, _ = syscall.Syscall(obj.vmt().GetOpacity, 1, uintptr(unsafe.Pointer(obj)), 0, 0)
	return *(*float32)(unsafe.Pointer(ret))
}

// SetOpacity https://docs.microsoft.com/en-us/windows/win32/api/d2d1/nf-d2d1-id2d1brush-setopacity
func (obj *Brush) SetOpacity(opacity float32) {
	syscall.Syscall(obj.vmt().SetOpacity, 2, uintptr(unsafe.Pointer(obj)), uintptr(*(*uint32)(unsafe.Pointer(&opacity))), 0)
}

// Transform https://docs.microsoft.com/en-us/windows/win32/api/d2d1/nf-d2d1-id2d1brush-gettransform
func (obj *Brush) Transform() *Matrix3x2 {
	var transform Matrix3x2
	syscall.Syscall(obj.vmt().GetTransform, 2, uintptr(unsafe.Pointer(obj)), uintptr(unsafe.Pointer(&transform)), 0)
	return &transform
}

// SetTransform https://docs.microsoft.com/en-us/windows/win32/api/d2d1/nf-d2d1-id2d1brush-settransform(constd2d1_matrix_3x2_f_)
func (obj *Brush) SetTransform(transform *Matrix3x2) {
	syscall.Syscall(obj.vmt().SetTransform, 2, uintptr(unsafe.Pointer(obj)), uintptr(unsafe.Pointer(transform)), 0)
}
