package d2d

import (
	"syscall"
	"unsafe"

	"github.com/richardwilkes/win32"
	"github.com/richardwilkes/win32/com"
)

// HWNDRenderTargetGUID holds the GUID for the HWNDRenderTarget class.
var HWNDRenderTargetGUID = com.GUID{Data1: 0x2cd90698, Data2: 0x12e2, Data3: 0x11dc, Data4: [8]byte{0x9f, 0xed, 0x00, 0x11, 0x43, 0xa0, 0x55, 0xf9}}

// HWNDRenderTarget https://docs.microsoft.com/en-us/windows/win32/api/d2d1/nn-d2d1-id2d1hwndrendertarget
type HWNDRenderTarget struct {
	RenderTarget
}

type vmtHWNDRenderTarget struct {
	vmtRenderTarget
	CheckWindowState uintptr
	Resize           uintptr
	GetHwnd          uintptr
}

func (obj *HWNDRenderTarget) vmt() *vmtHWNDRenderTarget {
	return (*vmtHWNDRenderTarget)(obj.UnsafeVirtualMethodTable)
}

// IsOccluded https://docs.microsoft.com/en-us/windows/win32/api/d2d1/nf-d2d1-id2d1hwndrendertarget-checkwindowstate
func (obj *HWNDRenderTarget) IsOccluded() bool {
	ret, _, _ := syscall.Syscall(obj.vmt().CheckWindowState, 1, uintptr(unsafe.Pointer(obj)), 0, 0)
	return ret == 1
}

// Resize https://docs.microsoft.com/en-us/windows/win32/api/d2d1/nf-d2d1-id2d1hwndrendertarget-resize%28constd2d1_size_u%29
func (obj *HWNDRenderTarget) Resize(pixelSize SizeU) {
	syscall.Syscall(obj.vmt().Resize, 2, uintptr(unsafe.Pointer(obj)), uintptr(unsafe.Pointer(&pixelSize)), 0)
}

// HWND https://docs.microsoft.com/en-us/windows/win32/api/d2d1/nf-d2d1-id2d1hwndrendertarget-gethwnd
func (obj *HWNDRenderTarget) HWND() win32.HWND {
	ret, _, _ := syscall.Syscall(obj.vmt().GetHwnd, 1, uintptr(unsafe.Pointer(obj)), 0, 0)
	return win32.HWND(ret)
}
