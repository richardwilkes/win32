package d2d

import (
	"syscall"
	"unsafe"

	"github.com/richardwilkes/win32/com"
)

// SolidColorBrushGUID holds the GUID for the SolidColorBrush class.
var SolidColorBrushGUID = com.GUID{Data1: 0x2cd906a9, Data2: 0x12e2, Data3: 0x11dc, Data4: [8]byte{0x9f, 0xed, 0x00, 0x11, 0x43, 0xa0, 0x55, 0xf9}}

// SolidColorBrush https://docs.microsoft.com/en-us/windows/win32/api/d2d1/nn-d2d1-id2d1solidcolorbrush
type SolidColorBrush struct {
	Brush
}

type vmtSolidColorBrush struct {
	vmtBrush
	SetColor uintptr
	GetColor uintptr
}

func (obj *SolidColorBrush) vmt() *vmtSolidColorBrush {
	return (*vmtSolidColorBrush)(obj.UnsafeVirtualMethodTable)
}

// Color https://docs.microsoft.com/en-us/windows/win32/api/d2d1/nf-d2d1-id2d1solidcolorbrush-getcolor
func (obj *SolidColorBrush) Color() Color {
	var color Color
	syscall.Syscall(obj.vmt().GetColor, 2, uintptr(unsafe.Pointer(obj)), uintptr(unsafe.Pointer(&color)), 0)
	return color
}

// SetColor https://docs.microsoft.com/en-us/windows/win32/api/d2d1/nf-d2d1-id2d1solidcolorbrush-setcolor(constd2d1_color_f_)
func (obj *SolidColorBrush) SetColor(color Color) {
	syscall.Syscall(obj.vmt().SetColor, 2, uintptr(unsafe.Pointer(obj)), uintptr(unsafe.Pointer(&color)), 0)
}
