package win32

import (
	"syscall"
	"unsafe"
)

var (
	gdi32         = syscall.NewLazyDLL("gdi32.dll")
	createDCW     = gdi32.NewProc("CreateDCW")
	deleteDC      = gdi32.NewProc("DeleteDC")
	getDeviceCaps = gdi32.NewProc("GetDeviceCaps")
)

// CreateDC https://docs.microsoft.com/en-us/windows/desktop/api/wingdi/nf-wingdi-createdcw
func CreateDC(driver, device, port LPCWSTR, pdm *DEVMODE) HDC {
	h, _, _ := createDCW.Call(uintptr(unsafe.Pointer(driver)), uintptr(unsafe.Pointer(device)), uintptr(unsafe.Pointer(port)), uintptr(unsafe.Pointer(pdm))) //nolint:errcheck
	return HDC(h)
}

// DeleteDC https://docs.microsoft.com/en-us/windows/desktop/api/wingdi/nf-wingdi-deletedc
func DeleteDC(hdc HDC) bool {
	ret, _, _ := deleteDC.Call(uintptr(hdc)) //nolint:errcheck
	return ret != 0
}

// GetDeviceCaps https://docs.microsoft.com/en-us/windows/desktop/api/wingdi/nf-wingdi-getdevicecaps
func GetDeviceCaps(hdc HDC, index int) int {
	ret, _, _ := getDeviceCaps.Call(uintptr(hdc), uintptr(index)) //nolint:errcheck
	return int(ret)
}
