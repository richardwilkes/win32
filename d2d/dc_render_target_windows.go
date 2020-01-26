package d2d

import (
	"fmt"
	"syscall"
	"unsafe"

	"github.com/richardwilkes/win32"
	"github.com/richardwilkes/win32/com"
)

// DCRenderTargetGUID holds the GUID for the DCRenderTarget class.
var DCRenderTargetGUID = com.GUID{Data1: 0x1c51bc64, Data2: 0xde61, Data3: 0x46fd, Data4: [8]byte{0x98, 0x99, 0x63, 0xa5, 0xd8, 0xf0, 0x39, 0x50}}

// DCRenderTarget https://docs.microsoft.com/en-us/windows/win32/api/d2d1/nn-d2d1-id2d1dcrendertarget
type DCRenderTarget struct {
	RenderTarget
}

type vmtDCRenderTarget struct {
	vmtRenderTarget
	BindDC uintptr
}

func (obj *DCRenderTarget) vmt() *vmtDCRenderTarget {
	return (*vmtDCRenderTarget)(obj.UnsafeVirtualMethodTable)
}

// BindDC https://docs.microsoft.com/en-us/windows/win32/api/d2d1/nf-d2d1-id2d1dcrendertarget-binddc
func (obj *DCRenderTarget) BindDC(hDC win32.HDC, rect win32.RECT) error {
	if ret, _, _ := syscall.Syscall(obj.vmt().BindDC, 3, uintptr(unsafe.Pointer(obj)), uintptr(hDC), uintptr(unsafe.Pointer(&rect))); ret != win32.S_OK {
		return fmt.Errorf("call to BindDC failed: %#x", ret)
	}
	return nil
}
