package win32

import (
	"syscall"
	"unsafe"
)

var (
	shcore           = syscall.NewLazyDLL("shcore.dll")
	getDpiForMonitor = shcore.NewProc("GetDpiForMonitor")
)

// GetDpiForMonitor https://docs.microsoft.com/en-us/windows/desktop/api/shellscalingapi/nf-shellscalingapi-getdpiformonitor
func GetDpiForMonitor(monitor HMONITOR, dpiType int32, dpiX, dpiY *uint32) bool {
	if err := getDpiForMonitor.Find(); err != nil {
		return false
	}
	ret, _, _ := getDpiForMonitor.Call(uintptr(monitor), uintptr(dpiType), uintptr(unsafe.Pointer(dpiX)), uintptr(unsafe.Pointer(dpiY))) //nolint:errcheck
	return ret == S_OK
}
