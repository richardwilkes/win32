package win32

import (
	"syscall"
	"unsafe"

	"github.com/richardwilkes/toolbox/errs"
)

var (
	gdi32            = syscall.NewLazyDLL("gdi32.dll")
	createDCW        = gdi32.NewProc("CreateDCW")
	createDIBSection = gdi32.NewProc("CreateDIBSection")
	deleteDC         = gdi32.NewProc("DeleteDC")
	deleteObject     = gdi32.NewProc("DeleteObject")
	gdiFlush         = gdi32.NewProc("GdiFlush")
	getDeviceCaps    = gdi32.NewProc("GetDeviceCaps")
)

// CreateDCS https://docs.microsoft.com/en-us/windows/desktop/api/wingdi/nf-wingdi-createdcw
func CreateDCS(driver, device, port string, pdm *DEVMODE) HDC {
	return CreateDC(ToWin32Str(driver, true), ToWin32Str(device, true), ToWin32Str(port, true), pdm)
}

// CreateDC https://docs.microsoft.com/en-us/windows/desktop/api/wingdi/nf-wingdi-createdcw
func CreateDC(driver, device, port LPCWSTR, pdm *DEVMODE) HDC {
	h, _, _ := createDCW.Call(uintptr(unsafe.Pointer(driver)), uintptr(unsafe.Pointer(device)), uintptr(unsafe.Pointer(port)), uintptr(unsafe.Pointer(pdm))) //nolint:errcheck
	return HDC(h)
}

// CreateDIBSection https://docs.microsoft.com/en-us/windows/desktop/api/wingdi/nf-wingdi-createdibsection
func CreateDIBSection(hdc HDC, pbmi *BITMAPINFOHEADER, usage uint32, ppvBits *unsafe.Pointer, hSection HANDLE, offset uint32) (HBITMAP, error) {
	ret, _, _ := createDIBSection.Call(uintptr(hdc), uintptr(unsafe.Pointer(pbmi)), uintptr(usage), uintptr(unsafe.Pointer(ppvBits)), uintptr(hSection), uintptr(offset))
	if ret == ERROR_INVALID_PARAMETER {
		return 0, errs.New("invalid parameter")
	}
	if ret == 0 {
		return 0, errs.New("unable to create DIB section")
	}
	return HBITMAP(ret), nil
}

// DeleteDC https://docs.microsoft.com/en-us/windows/desktop/api/wingdi/nf-wingdi-deletedc
func DeleteDC(hdc HDC) bool {
	ret, _, _ := deleteDC.Call(uintptr(hdc)) //nolint:errcheck
	return ret != 0
}

// DeleteObject https://docs.microsoft.com/en-us/windows/desktop/api/wingdi/nf-wingdi-deleteobject
func DeleteObject(obj HGDIOBJ) bool {
	ret, _, _ := deleteObject.Call(uintptr(obj)) //nolint:errcheck
	return ret != 0
}

// GdiFlush https://docs.microsoft.com/en-us/windows/desktop/api/wingdi/nf-wingdi-gdiflush
func GdiFlush() bool {
	ret, _, _ := gdiFlush.Call()
	return ret != 0
}

// GetDeviceCaps https://docs.microsoft.com/en-us/windows/desktop/api/wingdi/nf-wingdi-getdevicecaps
func GetDeviceCaps(hdc HDC, index int) int {
	ret, _, _ := getDeviceCaps.Call(uintptr(hdc), uintptr(index)) //nolint:errcheck
	return int(ret)
}
