package win32

import (
	"syscall"
	"unsafe"
)

var (
	kernel32           = syscall.NewLazyDLL("kernel32.dll")
	getCurrentThreadId = kernel32.NewProc("GetCurrentThreadId")
	getModuleHandleW   = kernel32.NewProc("GetModuleHandleW")
	globalAlloc        = kernel32.NewProc("GlobalAlloc")
	globalFree         = kernel32.NewProc("GlobalFree")
	globalLock         = kernel32.NewProc("GlobalLock")
	globalUnlock       = kernel32.NewProc("GlobalUnlock")
	moveMemory         = kernel32.NewProc("MoveMemory")
)

// GetCurrentThreadID https://docs.microsoft.com/en-us/windows/desktop/api/processthreadsapi/nf-processthreadsapi-getcurrentthreadid
func GetCurrentThreadID() DWORD {
	ret, _, _ := getCurrentThreadId.Call()
	return DWORD(ret)
}

// GetModuleHandleS https://docs.microsoft.com/en-us/windows/desktop/api/libloaderapi/nf-libloaderapi-getmodulehandlew
func GetModuleHandleS(moduleName string) HINSTANCE {
	return GetModuleHandle(ToWin32Str(moduleName, true))
}

// GetModuleHandle https://docs.microsoft.com/en-us/windows/desktop/api/libloaderapi/nf-libloaderapi-getmodulehandlew
func GetModuleHandle(moduleName LPCWSTR) HINSTANCE {
	h, _, _ := getModuleHandleW.Call(uintptr(unsafe.Pointer(moduleName))) //nolint:errcheck
	return HINSTANCE(h)
}

// GlobalAlloc https://docs.microsoft.com/en-us/windows/desktop/api/winbase/nf-winbase-globalalloc
func GlobalAlloc(flags uint, size int) HGLOBAL {
	ret, _, _ := globalAlloc.Call(uintptr(flags), uintptr(size))
	return HGLOBAL(ret)
}

// GlobalFree https://docs.microsoft.com/en-us/windows/desktop/api/winbase/nf-winbase-globalfree
func GlobalFree(mem HGLOBAL) HGLOBAL {
	ret, _, _ := globalFree.Call(uintptr(mem))
	return HGLOBAL(ret)
}

// GlobalLock https://docs.microsoft.com/en-us/windows/desktop/api/winbase/nf-winbase-globallock
func GlobalLock(mem HGLOBAL) unsafe.Pointer {
	ret, _, _ := globalLock.Call(uintptr(mem))
	return unsafe.Pointer(ret)
}

// GlobalUnlock https://docs.microsoft.com/en-us/windows/desktop/api/winbase/nf-winbase-globalunlock
func GlobalUnlock(mem HGLOBAL) bool {
	ret, _, _ := globalUnlock.Call(uintptr(mem))
	return ret != 0
}

// MoveMemory https://msdn.microsoft.com/en-us/library/windows/desktop/aa366788(v=vs.85).aspx
func MoveMemory(dst unsafe.Pointer, src unsafe.Pointer, length int) {
	moveMemory.Call(uintptr(dst), uintptr(src), uintptr(length))
}
