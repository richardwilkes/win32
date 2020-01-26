package d2d

import (
	"fmt"
	"syscall"
	"unsafe"

	"github.com/richardwilkes/win32"
	"github.com/richardwilkes/win32/com"
)

// TessellationSinkGUID holds the GUID for the TessellationSink class.
var TessellationSinkGUID = com.GUID{Data1: 0x2cd906c1, Data2: 0x12e2, Data3: 0x11dc, Data4: [8]byte{0x9f, 0xed, 0x00, 0x11, 0x43, 0xa0, 0x55, 0xf9}}

// TessellationSink https://docs.microsoft.com/en-us/windows/win32/api/d2d1/nn-d2d1-id2d1tessellationsink
type TessellationSink struct {
	com.IUnknown
}

type vmtTessellationSink struct {
	com.VMTIUnknown
	AddTriangles uintptr
	Close        uintptr
}

func (obj *TessellationSink) vmt() *vmtTessellationSink {
	return (*vmtTessellationSink)(obj.UnsafeVirtualMethodTable)
}

// AddTriangles https://docs.microsoft.com/en-us/windows/win32/api/d2d1/nf-d2d1-id2d1tessellationsink-addtriangles
func (obj *TessellationSink) AddTriangles(triangles []Triangle) {
	syscall.Syscall(obj.vmt().AddTriangles, 3, uintptr(unsafe.Pointer(obj)), uintptr(unsafe.Pointer(&(triangles[0]))), uintptr(len(triangles)))
}

// Close https://docs.microsoft.com/en-us/windows/win32/api/d2d1/nf-d2d1-id2d1tessellationsink-close
func (obj *TessellationSink) Close() error {
	if ret, _, _ := syscall.Syscall(obj.vmt().Close, 1, uintptr(unsafe.Pointer(obj)), 0, 0); ret != win32.S_OK {
		return fmt.Errorf("call to Close failed: %#x", ret)
	}
	return nil
}
