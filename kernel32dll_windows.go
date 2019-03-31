package win32

import (
	"syscall"
	"unsafe"
)

var (
	kernel32         = syscall.NewLazyDLL("kernel32.dll")
	getModuleHandleW = kernel32.NewProc("GetModuleHandleW")
)

// GetModuleHandleS https://docs.microsoft.com/en-us/windows/desktop/api/libloaderapi/nf-libloaderapi-getmodulehandlew
func GetModuleHandleS(moduleName string) HINSTANCE {
	return GetModuleHandle(ToWin32Str(moduleName, true))
}

// GetModuleHandle https://docs.microsoft.com/en-us/windows/desktop/api/libloaderapi/nf-libloaderapi-getmodulehandlew
func GetModuleHandle(moduleName LPCWSTR) HINSTANCE {
	h, _, _ := getModuleHandleW.Call(uintptr(unsafe.Pointer(moduleName))) //nolint:errcheck
	return HINSTANCE(h)
}
