package win32

import (
	"syscall"
	"unsafe"
)

// ToWin32Str converts a Go string to a Windows string.
func ToWin32Str(in string, emptyReturnsNil bool) *uint16 {
	if in == "" {
		return nil
	}
	out, err := syscall.UTF16PtrFromString(in)
	if err != nil {
		var empty [1]uint16
		out = &empty[0]
	}
	return out
}

// ToSysWin32Str converts a Go string to a Windows string suitable for passing
// via a syscall.
func ToSysWin32Str(in string, emptyReturnsNil bool) uintptr {
	return uintptr(unsafe.Pointer(ToWin32Str(in, emptyReturnsNil)))
}

// FromBOOL converts a Windows BOOL to a Go bool.
func FromBOOL(in BOOL) bool {
	return in != 0
}

// ToBOOL converts a Go bool to a Windows BOOL.
func ToBOOL(in bool) BOOL {
	if in {
		return 1
	}
	return 0
}

// ToSysBool converts a Go bool to a Windows BOOL suitable for passing via a
// syscall.
func ToSysBool(in bool) uintptr {
	if in {
		return 1
	}
	return 0
}
