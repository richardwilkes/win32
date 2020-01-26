package win32

import (
	"syscall"
	"unsafe"
)

var (
	user32                        = syscall.NewLazyDLL("user32.dll")
	adjustWindowRectEx            = user32.NewProc("AdjustWindowRectEx")
	beginPaint                    = user32.NewProc("BeginPaint")
	closeClipboard                = user32.NewProc("CloseClipboard")
	createAcceleratorTableW       = user32.NewProc("CreateAcceleratorTableW")
	createMenu                    = user32.NewProc("CreateMenu")
	createPopupMenu               = user32.NewProc("CreatePopupMenu")
	createWindowExW               = user32.NewProc("CreateWindowExW")
	defWindowProcW                = user32.NewProc("DefWindowProcW")
	deleteMenu                    = user32.NewProc("DeleteMenu")
	destroyAcceleratorTable       = user32.NewProc("DestroyAcceleratorTable")
	destroyMenu                   = user32.NewProc("DestroyMenu")
	destroyWindow                 = user32.NewProc("DestroyWindow")
	dispatchMessageW              = user32.NewProc("DispatchMessageW")
	drawMenuBar                   = user32.NewProc("DrawMenuBar")
	emptyClipboard                = user32.NewProc("EmptyClipboard")
	enableMenuItem                = user32.NewProc("EnableMenuItem")
	enableWindow                  = user32.NewProc("EnableWindow")
	endPaint                      = user32.NewProc("EndPaint")
	enumClipboardFormats          = user32.NewProc("EnumClipboardFormats")
	enumDisplayDevicesW           = user32.NewProc("EnumDisplayDevicesW")
	enumDisplayMonitors           = user32.NewProc("EnumDisplayMonitors")
	enumDisplaySettingsExW        = user32.NewProc("EnumDisplaySettingsExW")
	enumWindows                   = user32.NewProc("EnumWindows")
	getActiveWindow               = user32.NewProc("GetActiveWindow")
	getClientRect                 = user32.NewProc("GetClientRect")
	getClipboardData              = user32.NewProc("GetClipboardData")
	getClipboardFormatNameA       = user32.NewProc("getClipboardFormatNameA")
	getClipboardSequenceNumber    = user32.NewProc("GetClipboardSequenceNumber")
	getDC                         = user32.NewProc("GetDC")
	getDpiForSystem               = user32.NewProc("GetDpiForSystem")
	getFocus                      = user32.NewProc("GetFocus")
	getForegroundWindow           = user32.NewProc("GetForegroundWindow")
	getMenu                       = user32.NewProc("GetMenu")
	getMenuItemCount              = user32.NewProc("GetMenuItemCount")
	getMenuItemInfoW              = user32.NewProc("GetMenuItemInfoW")
	getMessageW                   = user32.NewProc("GetMessageW")
	getMonitorInfoW               = user32.NewProc("GetMonitorInfoW")
	getSystemMetrics              = user32.NewProc("GetSystemMetrics")
	getWindow                     = user32.NewProc("GetWindow")
	getWindowRect                 = user32.NewProc("GetWindowRect")
	getWindowTextW                = user32.NewProc("GetWindowTextW")
	insertMenuItemW               = user32.NewProc("InsertMenuItemW")
	invalidateRect                = user32.NewProc("InvalidateRect")
	isIconic                      = user32.NewProc("IsIconic")
	isZoomed                      = user32.NewProc("IsZoomed")
	loadCursorW                   = user32.NewProc("LoadCursorW")
	mapWindowPoints               = user32.NewProc("MapWindowPoints")
	moveWindow                    = user32.NewProc("MoveWindow")
	openClipboard                 = user32.NewProc("OpenClipboard")
	postMessageW                  = user32.NewProc("PostMessageW")
	postQuitMessage               = user32.NewProc("PostQuitMessage")
	postThreadMessageW            = user32.NewProc("PostThreadMessageW")
	registerClassExW              = user32.NewProc("RegisterClassExW")
	registerWindowMessageW        = user32.NewProc("RegisterWindowMessageW")
	releaseDC                     = user32.NewProc("ReleaseDC")
	setActiveWindow               = user32.NewProc("SetActiveWindow")
	setClipboardData              = user32.NewProc("SetClipboardData")
	setFocus                      = user32.NewProc("SetFocus")
	setForegroundWindow           = user32.NewProc("SetForegroundWindow")
	setMenu                       = user32.NewProc("SetMenu")
	setMenuItemInfoW              = user32.NewProc("SetMenuItemInfoW")
	setProcessDpiAwarenessContext = user32.NewProc("SetProcessDpiAwarenessContext")
	setWindowPos                  = user32.NewProc("SetWindowPos")
	setWindowTextW                = user32.NewProc("SetWindowTextW")
	showWindow                    = user32.NewProc("ShowWindow")
	translateMessage              = user32.NewProc("TranslateMessage")
	validateRect                  = user32.NewProc("ValidateRect")
)

// AdjustWindowRectEx https://docs.microsoft.com/en-us/windows/desktop/api/winuser/nf-winuser-adjustwindowrectex
func AdjustWindowRectEx(rect *RECT, style DWORD, hasMenu bool, exStyle DWORD) {
	adjustWindowRectEx.Call(uintptr(unsafe.Pointer(rect)), uintptr(style), ToSysBool(hasMenu), uintptr(exStyle)) //nolint:errcheck
}

// BeginPaint https://docs.microsoft.com/en-us/windows/desktop/api/Winuser/nf-winuser-beginpaint
func BeginPaint(hwnd HWND, ps *PAINTSTRUCT) HDC {
	ret, _, _ := beginPaint.Call(uintptr(hwnd), uintptr(unsafe.Pointer(ps)))
	return HDC(ret)
}

// CloseClipboard https://docs.microsoft.com/en-us/windows/desktop/api/winuser/nf-winuser-closeclipboard
func CloseClipboard() bool {
	ret, _, _ := closeClipboard.Call()
	return ret != 0
}

// CreateAcceleratorTable https://docs.microsoft.com/en-us/windows/desktop/api/winuser/nf-winuser-createacceleratortablew
func CreateAcceleratorTable(accelList LPACCEL, count int) HACCEL {
	ret, _, _ := createAcceleratorTableW.Call(uintptr(accelList), uintptr(count)) //nolint:errcheck
	return HACCEL(ret)
}

// CreateMenu https://docs.microsoft.com/en-us/windows/desktop/api/winuser/nf-winuser-createmenu
func CreateMenu() HMENU {
	ret, _, _ := createMenu.Call() //nolint:errcheck
	return HMENU(ret)
}

// CreatePopupMenu https://docs.microsoft.com/en-us/windows/desktop/api/winuser/nf-winuser-createpopupmenu
func CreatePopupMenu() HMENU {
	ret, _, _ := createPopupMenu.Call() //nolint:errcheck
	return HMENU(ret)
}

// CreateWindowExS https://docs.microsoft.com/en-us/windows/desktop/api/winuser/nf-winuser-createwindowexw
func CreateWindowExS(exStyle DWORD, className, windowName string, style DWORD, x, y, width, height int32, parent HWND, menu HMENU, instance HINSTANCE, param LPVOID) HWND {
	return CreateWindowEx(exStyle, ToWin32Str(className, true), ToWin32Str(windowName, true), style, x, y, width, height, parent, menu, instance, param)
}

// CreateWindowEx https://docs.microsoft.com/en-us/windows/desktop/api/winuser/nf-winuser-createwindowexw
func CreateWindowEx(exStyle DWORD, className, windowName LPCWSTR, style DWORD, x, y, width, height int32, parent HWND, menu HMENU, instance HINSTANCE, param LPVOID) HWND {
	ret, _, _ := createWindowExW.Call(uintptr(exStyle), uintptr(unsafe.Pointer(className)), uintptr(unsafe.Pointer(windowName)), uintptr(style), uintptr(x), uintptr(y), uintptr(width), uintptr(height), uintptr(parent), uintptr(menu), uintptr(instance), uintptr(param)) //nolint:errcheck
	return HWND(ret)
}

// DefWindowProc https://docs.microsoft.com/en-us/windows/desktop/api/winuser/nf-winuser-defwindowprocw
func DefWindowProc(hwnd HWND, msg uint32, wparam WPARAM, lparam LPARAM) LRESULT {
	ret, _, _ := defWindowProcW.Call(uintptr(hwnd), uintptr(msg), uintptr(wparam), uintptr(lparam)) //nolint:errcheck
	return LRESULT(ret)
}

// DeleteMenu https://docs.microsoft.com/en-us/windows/desktop/api/winuser/nf-winuser-deletemenu
func DeleteMenu(hmenu HMENU, position, flags uint32) {
	deleteMenu.Call(uintptr(hmenu), uintptr(position), uintptr(flags)) //nolint:errcheck
}

// DestroyAcceleratorTable https://docs.microsoft.com/en-us/windows/desktop/api/winuser/nf-winuser-destroyacceleratortable
func DestroyAcceleratorTable(accel HACCEL) bool {
	ret, _, _ := destroyAcceleratorTable.Call(uintptr(accel)) //nolint:errcheck
	return ret != 0
}

// DestroyMenu https://docs.microsoft.com/en-us/windows/desktop/api/winuser/nf-winuser-destroymenu
func DestroyMenu(menu HMENU) {
	destroyMenu.Call(uintptr(menu)) //nolint:errcheck
}

// DestroyWindow https://docs.microsoft.com/en-us/windows/desktop/api/winuser/nf-winuser-destroywindow
func DestroyWindow(hwnd HWND) {
	destroyWindow.Call(uintptr(hwnd)) //nolint:errcheck
}

// DispatchMessage https://docs.microsoft.com/en-us/windows/desktop/api/winuser/nf-winuser-dispatchmessage
func DispatchMessage(msg *MSG) LRESULT {
	ret, _, _ := dispatchMessageW.Call(uintptr(unsafe.Pointer(msg))) //nolint:errcheck
	return LRESULT(ret)
}

// DrawMenuBar https://docs.microsoft.com/en-us/windows/desktop/api/Winuser/nf-winuser-drawmenubar
func DrawMenuBar(hwnd HWND) {
	drawMenuBar.Call(uintptr(hwnd)) //nolint:errcheck
}

// EmptyClipboard https://docs.microsoft.com/en-us/windows/desktop/api/winuser/nf-winuser-emptyclipboard
func EmptyClipboard() bool {
	ret, _, _ := emptyClipboard.Call()
	return ret != 0
}

// EnableMenuItem https://docs.microsoft.com/en-us/windows/desktop/api/winuser/nf-winuser-enablemenuitem
func EnableMenuItem(hmenu HMENU, idEnableItem, enable int) int {
	ret, _, _ := enableMenuItem.Call(uintptr(hmenu), uintptr(idEnableItem), uintptr(enable)) //nolint:errcheck
	return int(ret)
}

// EnableWindow https://docs.microsoft.com/en-us/windows/desktop/api/winuser/nf-winuser-enablewindow
func EnableWindow(hwnd HWND, enable bool) bool {
	ret, _, _ := enableWindow.Call(uintptr(hwnd), ToSysBool(enable)) //nolint:errcheck
	return ret != 0
}

// EndPaint https://docs.microsoft.com/en-us/windows/desktop/api/Winuser/nf-winuser-endpaint
func EndPaint(hwnd HWND, ps *PAINTSTRUCT) {
	endPaint.Call(uintptr(hwnd), uintptr(unsafe.Pointer(ps)))
}

// EnumClipboardFormats https://docs.microsoft.com/en-us/windows/desktop/api/winuser/nf-winuser-enumclipboardformats
func EnumClipboardFormats(format uint) uint {
	ret, _, _ := enumClipboardFormats.Call(uintptr(format))
	return uint(ret)
}

// EnumDisplayDevicesS https://docs.microsoft.com/en-us/windows/desktop/api/winuser/nf-winuser-enumdisplaydevicesw
func EnumDisplayDevicesS(device string, devNum DWORD, displayDevice *DISPLAY_DEVICE, flags DWORD) {
	EnumDisplayDevices(ToWin32Str(device, true), devNum, displayDevice, flags)
}

// EnumDisplayDevices https://docs.microsoft.com/en-us/windows/desktop/api/winuser/nf-winuser-enumdisplaydevicesw
func EnumDisplayDevices(device LPCWSTR, devNum DWORD, displayDevice *DISPLAY_DEVICE, flags DWORD) {
	enumDisplayDevicesW.Call(uintptr(unsafe.Pointer(device)), uintptr(devNum), uintptr(unsafe.Pointer(displayDevice)), uintptr(flags)) //nolint:errcheck
}

// EnumDisplayMonitors https://docs.microsoft.com/en-us/windows/desktop/api/Winuser/nf-winuser-enumdisplaymonitors
func EnumDisplayMonitors(hdc HDC, clip *RECT, callback func(monitor HMONITOR, dc HDC, rect *RECT, param LPARAM) BOOL, data LPARAM) {
	enumDisplayMonitors.Call(uintptr(hdc), uintptr(unsafe.Pointer(clip)), syscall.NewCallback(func(monitor HMONITOR, dc HDC, rect *RECT, param LPARAM) uintptr { //nolint:errcheck
		return uintptr(callback(monitor, dc, rect, param))
	}), uintptr(data))
}

// EnumDisplaySettingsExS https://docs.microsoft.com/en-us/windows/desktop/api/winuser/nf-winuser-enumdisplaysettingsexw
func EnumDisplaySettingsExS(deviceName string, modeNum DWORD, devMode *DEVMODE, flags DWORD) {
	EnumDisplaySettingsEx(ToWin32Str(deviceName, true), modeNum, devMode, flags)
}

// EnumDisplaySettingsEx https://docs.microsoft.com/en-us/windows/desktop/api/winuser/nf-winuser-enumdisplaysettingsexw
func EnumDisplaySettingsEx(deviceName LPCWSTR, modeNum DWORD, devMode *DEVMODE, flags DWORD) {
	enumDisplaySettingsExW.Call(uintptr(unsafe.Pointer(deviceName)), uintptr(modeNum), uintptr(unsafe.Pointer(devMode)), uintptr(flags)) //nolint:errcheck
}

// EnumWindows https://docs.microsoft.com/en-us/windows/desktop/api/winuser/nf-winuser-enumwindows
func EnumWindows(callback func(hwnd HWND, data LPARAM) BOOL, param LPARAM) {
	enumWindows.Call(syscall.NewCallback(func(hwnd HWND, data LPARAM) uintptr { return uintptr(callback(hwnd, data)) }), uintptr(param)) //nolint:errcheck
}

// GetActiveWindow https://docs.microsoft.com/en-us/windows/desktop/api/winuser/nf-winuser-getactivewindow
func GetActiveWindow() HWND {
	ret, _, _ := getActiveWindow.Call() //nolint:errcheck
	return HWND(ret)
}

// GetClientRect https://docs.microsoft.com/en-us/windows/desktop/api/winuser/nf-winuser-getclientrect
func GetClientRect(hwnd HWND, rect *RECT) {
	getClientRect.Call(uintptr(hwnd), uintptr(unsafe.Pointer(rect))) //nolint:errcheck
}

// GetClipboardData https://docs.microsoft.com/en-us/windows/desktop/api/Winuser/nf-winuser-getclipboarddata
func GetClipboardData(format uint) HANDLE {
	ret, _, _ := getClipboardData.Call(uintptr(format))
	return HANDLE(ret)
}

// GetClipboardFormatName https://docs.microsoft.com/en-us/windows/desktop/api/Winuser/nf-winuser-getclipboardformatnamea
func GetClipboardFormatName(format uint) string {
	var buffer [256]byte
	if ret, _, _ := getClipboardFormatNameA.Call(uintptr(format), uintptr(unsafe.Pointer(&buffer[0])), uintptr(len(buffer)-1)); ret != 0 {
		str := make([]byte, ret)
		copy(str, buffer[:])
		return string(str)
	}
	return ""
}

// GetClipboardSequenceNumber https://docs.microsoft.com/en-us/windows/desktop/api/winuser/nf-winuser-getclipboardsequencenumber
func GetClipboardSequenceNumber() int {
	ret, _, _ := getClipboardSequenceNumber.Call()
	return int(ret)
}

// GetDC https://docs.microsoft.com/en-us/windows/desktop/api/winuser/nf-winuser-getdc
func GetDC(hwnd HWND) HDC {
	ret, _, _ := getDC.Call(uintptr(hwnd))
	return HDC(ret)
}

// GetDpiForSystem https://docs.microsoft.com/en-us/windows/desktop/api/winuser/nf-winuser-getdpiforsystem
func GetDpiForSystem() int {
	ret, _, _ := getDpiForSystem.Call() //nolint:errcheck
	return int(ret)
}

// GetFocus https://docs.microsoft.com/en-us/windows/desktop/api/winuser/nf-winuser-getfocus
func GetFocus() HWND {
	ret, _, _ := getFocus.Call() //nolint:errcheck
	return HWND(ret)
}

// GetForegroundWindow https://docs.microsoft.com/en-us/windows/desktop/api/winuser/nf-winuser-getforegroundwindow
func GetForegroundWindow() HWND {
	ret, _, _ := getForegroundWindow.Call() //nolint:errcheck
	return HWND(ret)
}

// GetMenu https://docs.microsoft.com/en-us/windows/desktop/api/winuser/nf-winuser-getmenu
func GetMenu(hwnd HWND) HMENU {
	ret, _, _ := getMenu.Call(uintptr(hwnd)) //nolint:errcheck
	return HMENU(ret)
}

// GetMenuItemCount https://docs.microsoft.com/en-us/windows/desktop/api/Winuser/nf-winuser-getmenuitemcount
func GetMenuItemCount(hmenu HMENU) int {
	ret, _, _ := getMenuItemCount.Call(uintptr(hmenu)) //nolint:errcheck
	return int(ret)
}

// GetMenuItemInfo https://docs.microsoft.com/en-us/windows/desktop/api/winuser/nf-winuser-getmenuiteminfow
func GetMenuItemInfo(hmenu HMENU, item uint32, byPosition bool, lpmii *MENUITEMINFO) bool {
	ret, _, _ := getMenuItemInfoW.Call(uintptr(hmenu), uintptr(item), ToSysBool(byPosition), uintptr(unsafe.Pointer(lpmii))) //nolint:errcheck
	return ret != 0
}

// GetMessage https://docs.microsoft.com/en-us/windows/desktop/api/winuser/nf-winuser-getmessagew
func GetMessage(msg *MSG, hwnd HWND, msgFilterMin, msgFilterMax uint32) (quit bool, err error) {
	var ret uintptr
	ret, _, err = getMessageW.Call(uintptr(unsafe.Pointer(msg)), uintptr(hwnd), uintptr(msgFilterMin), uintptr(msgFilterMax)) //nolint:errcheck
	if ret == 0 {
		return true, nil
	}
	result := int32(ret)
	if result < 0 {
		return false, err
	}
	return false, nil
}

// GetMonitorInfo https://docs.microsoft.com/en-us/windows/desktop/api/winuser/nf-winuser-getmonitorinfow
func GetMonitorInfo(monitor HMONITOR, pmi *MONITORINFO) bool {
	ret, _, _ := getMonitorInfoW.Call(uintptr(monitor), uintptr(unsafe.Pointer(pmi))) //nolint:errcheck
	return ret != 0
}

// GetSystemMetrics https://docs.microsoft.com/en-us/windows/desktop/api/winuser/nf-winuser-getsystemmetrics
func GetSystemMetrics(index int) int {
	ret, _, _ := getSystemMetrics.Call(uintptr(index)) //nolint:errcheck
	return int(ret)
}

// GetWindow https://docs.microsoft.com/en-us/windows/desktop/api/winuser/nf-winuser-getwindow
func GetWindow(hwnd HWND, cmd int) HWND {
	ret, _, _ := getWindow.Call(uintptr(hwnd), uintptr(cmd)) //nolint:errcheck
	return HWND(ret)
}

// GetWindowRect https://docs.microsoft.com/en-us/windows/desktop/api/winuser/nf-winuser-getwindowrect
func GetWindowRect(hwnd HWND, rect *RECT) {
	getWindowRect.Call(uintptr(hwnd), uintptr(unsafe.Pointer(rect))) //nolint:errcheck
}

// GetWindowText https://docs.microsoft.com/en-us/windows/desktop/api/winuser/nf-winuser-getwindowtextw
func GetWindowText(hwnd HWND) string {
	var buffer [512]uint16
	getWindowTextW.Call(uintptr(hwnd), uintptr(unsafe.Pointer(&buffer[0])), uintptr(len(buffer)-1)) //nolint:errcheck
	return syscall.UTF16ToString(buffer[:])
}

// InsertMenuItem https://docs.microsoft.com/en-us/windows/desktop/api/winuser/nf-winuser-insertmenuitemw
func InsertMenuItem(hmenu HMENU, item uint32, byPosition bool, lpmi *MENUITEMINFO) {
	insertMenuItemW.Call(uintptr(hmenu), uintptr(item), ToSysBool(byPosition), uintptr(unsafe.Pointer(lpmi))) //nolint:errcheck
}

// InvalidateRect https://docs.microsoft.com/en-us/windows/desktop/api/Winuser/nf-winuser-invalidaterect
func InvalidateRect(hwnd HWND, rect *RECT, erase bool) {
	invalidateRect.Call(uintptr(hwnd), uintptr(unsafe.Pointer(rect)), ToSysBool(erase))
}

// IsIconic https://docs.microsoft.com/en-us/windows/desktop/api/winuser/nf-winuser-isiconic
func IsIconic(hwnd HWND) bool {
	ret, _, _ := isIconic.Call(uintptr(hwnd)) //nolint:errcheck
	return ret != 0
}

// IsZoomed https://docs.microsoft.com/en-us/windows/desktop/api/winuser/nf-winuser-iszoomed
func IsZoomed(hwnd HWND) bool {
	ret, _, _ := isZoomed.Call(uintptr(hwnd)) //nolint:errcheck
	return ret != 0
}

// LoadCursorS https://docs.microsoft.com/en-us/windows/desktop/api/winuser/nf-winuser-loadcursorw
func LoadCursorS(instance HINSTANCE, cursorName string) HCURSOR {
	return LoadCursor(instance, ToWin32Str(cursorName, false))
}

// LoadCursor https://docs.microsoft.com/en-us/windows/desktop/api/winuser/nf-winuser-loadcursorw
func LoadCursor(instance HINSTANCE, cursorName LPCWSTR) HCURSOR {
	h, _, _ := loadCursorW.Call(uintptr(instance), uintptr(unsafe.Pointer(cursorName))) //nolint:errcheck
	return HCURSOR(h)
}

// LoadSystemCursor https://docs.microsoft.com/en-us/windows/desktop/api/winuser/nf-winuser-loadcursorw
func LoadSystemCursor(cursorID int) HCURSOR {
	h, _, _ := loadCursorW.Call(0, uintptr(cursorID)) //nolint:errcheck
	return HCURSOR(h)
}

// MapWindowPoints https://docs.microsoft.com/en-us/windows/desktop/api/winuser/nf-winuser-mapwindowpoints
func MapWindowPoints(from, to HWND, points []POINT) int32 {
	ret, _, _ := mapWindowPoints.Call(uintptr(from), uintptr(to), uintptr(unsafe.Pointer(&points[0])), uintptr(len(points)))
	return int32(ret)
}

// MapWindowRect https://docs.microsoft.com/en-us/windows/desktop/api/winuser/nf-winuser-mapwindowpoints
func MapWindowRect(from, to HWND, rect *RECT) int32 {
	ret, _, _ := mapWindowPoints.Call(uintptr(from), uintptr(to), uintptr(unsafe.Pointer(rect)), 2)
	return int32(ret)
}

// MoveWindow https://docs.microsoft.com/en-us/windows/desktop/api/winuser/nf-winuser-movewindow
func MoveWindow(hwnd HWND, x, y, width, height int32, repaint bool) {
	moveWindow.Call(uintptr(hwnd), uintptr(x), uintptr(y), uintptr(width), uintptr(height), ToSysBool(repaint)) //nolint:errcheck
}

// OpenClipboard https://docs.microsoft.com/en-us/windows/desktop/api/winuser/nf-winuser-openclipboard
func OpenClipboard(hwnd HWND) bool {
	ret, _, _ := openClipboard.Call(uintptr(hwnd))
	return ret != 0
}

// PostMessage https://docs.microsoft.com/en-us/windows/desktop/api/winuser/nf-winuser-postmessagew
func PostMessage(hwnd HWND, msg uint32, wParam WPARAM, lParam LPARAM) error {
	if ret, _, err := postMessageW.Call(uintptr(hwnd), uintptr(msg), uintptr(wParam), uintptr(lParam)); ret == 0 {
		return err
	}
	return nil
}

// PostQuitMessage https://docs.microsoft.com/en-us/windows/desktop/api/winuser/nf-winuser-postquitmessage
func PostQuitMessage(exitCode int32) {
	postQuitMessage.Call(uintptr(exitCode)) //nolint:errcheck
}

// PostThreadMessage https://docs.microsoft.com/en-us/windows/desktop/api/winuser/nf-winuser-postthreadmessagew
func PostThreadMessage(threadID DWORD, msg uint32, wParam WPARAM, lParam LPARAM) error {
	if ret, _, err := postThreadMessageW.Call(uintptr(threadID), uintptr(msg), uintptr(wParam), uintptr(lParam)); ret == 0 {
		return err
	}
	return nil
}

// RegisterClassEx https://docs.microsoft.com/en-us/windows/desktop/api/winuser/nf-winuser-registerclassexw
func RegisterClassEx(wndcls *WNDCLASSEX) ATOM {
	h, _, _ := registerClassExW.Call(uintptr(unsafe.Pointer(wndcls))) //nolint:errcheck
	return ATOM(h)
}

// RegisterWindowMessageS https://docs.microsoft.com/en-us/windows/desktop/api/winuser/nf-winuser-registerwindowmessagew
func RegisterWindowMessageS(name string) uint32 {
	return RegisterWindowMessage(ToWin32Str(name, false))
}

// RegisterWindowMessage https://docs.microsoft.com/en-us/windows/desktop/api/winuser/nf-winuser-registerwindowmessagew
func RegisterWindowMessage(name LPCWSTR) uint32 {
	ret, _, _ := registerWindowMessageW.Call(uintptr(unsafe.Pointer(name))) //nolint:errcheck
	return uint32(ret)
}

// ReleaseDC https://docs.microsoft.com/en-us/windows/desktop/api/winuser/nf-winuser-releasedc
func ReleaseDC(hWnd HWND, hDC HDC) bool {
	ret, _, _ := releaseDC.Call(uintptr(hWnd), uintptr(hDC))
	return ret != 0
}

// SetActiveWindow https://docs.microsoft.com/en-us/windows/desktop/api/winuser/nf-winuser-setactivewindow
func SetActiveWindow(hwnd HWND) {
	setActiveWindow.Call(uintptr(hwnd)) //nolint:errcheck
}

// SetClipboardData https://docs.microsoft.com/en-us/windows/desktop/api/Winuser/nf-winuser-setclipboarddata
func SetClipboardData(format uint, data HANDLE) bool {
	ret, _, _ := setClipboardData.Call(uintptr(format), uintptr(data))
	return ret != 0
}

// SetFocus https://docs.microsoft.com/en-us/windows/desktop/api/winuser/nf-winuser-setfocus
func SetFocus(hwnd HWND) HWND {
	ret, _, _ := setFocus.Call(uintptr(hwnd)) //nolint:errcheck
	return HWND(ret)
}

// SetForegroundWindow https://docs.microsoft.com/en-us/windows/desktop/api/winuser/nf-winuser-setforegroundwindow
func SetForegroundWindow(hwnd HWND) bool {
	ret, _, _ := setForegroundWindow.Call(uintptr(hwnd)) //nolint:errcheck
	return ret != 0
}

// SetMenu https://docs.microsoft.com/en-us/windows/desktop/api/winuser/nf-winuser-setmenu
func SetMenu(hwnd HWND, menu HMENU) {
	setMenu.Call(uintptr(hwnd), uintptr(menu)) //nolint:errcheck
}

// SetMenuItemInfo https://docs.microsoft.com/en-us/windows/desktop/api/winuser/nf-winuser-setmenuiteminfow
func SetMenuItemInfo(menu HMENU, item uint32, byPosition bool, info *MENUITEMINFO) bool {
	ret, _, _ := setMenuItemInfoW.Call(uintptr(menu), uintptr(item), ToSysBool(byPosition), uintptr(unsafe.Pointer(info))) //nolint:errcheck
	return ret != 0
}

// SetProcessDpiAwarenessContext https://docs.microsoft.com/en-us/windows/desktop/api/winuser/nf-winuser-setprocessdpiawarenesscontext
func SetProcessDpiAwarenessContext(value DPI_AWARENESS_CONTEXT) {
	setProcessDpiAwarenessContext.Call(uintptr(value)) //nolint:errcheck
}

// SetWindowPos https://docs.microsoft.com/en-us/windows/desktop/api/winuser/nf-winuser-setwindowpos
func SetWindowPos(hwnd, hwndInsertAfter HWND, x, y, width, height int32, flags uint32) {
	setWindowPos.Call(uintptr(hwnd), uintptr(hwndInsertAfter), uintptr(x), uintptr(y), uintptr(width), uintptr(height), uintptr(flags)) //nolint:errcheck
}

// SetWindowTextS https://docs.microsoft.com/en-us/windows/desktop/api/winuser/nf-winuser-setwindowtextw
func SetWindowTextS(hwnd HWND, title string) {
	SetWindowText(hwnd, ToWin32Str(title, false))
}

// SetWindowText https://docs.microsoft.com/en-us/windows/desktop/api/winuser/nf-winuser-setwindowtextw
func SetWindowText(hwnd HWND, title LPCWSTR) {
	setWindowTextW.Call(uintptr(hwnd), uintptr(unsafe.Pointer(title))) //nolint:errcheck
}

// ShowWindow https://docs.microsoft.com/en-us/windows/desktop/api/winuser/nf-winuser-showwindow
func ShowWindow(hwnd HWND, cmd int32) bool {
	ret, _, _ := showWindow.Call(uintptr(hwnd), uintptr(cmd)) //nolint:errcheck
	return ret != 0
}

// TranslateMessage https://docs.microsoft.com/en-us/windows/desktop/api/winuser/nf-winuser-translatemessage
func TranslateMessage(msg *MSG) {
	translateMessage.Call(uintptr(unsafe.Pointer(msg)))
}

// ValidateRect https://docs.microsoft.com/en-us/windows/desktop/api/Winuser/nf-winuser-validaterect
func ValidateRect(hwnd HWND, rect *RECT) {
	validateRect.Call(uintptr(hwnd), uintptr(unsafe.Pointer(rect)))
}
