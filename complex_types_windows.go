package win32

import "syscall"

// RECT https://msdn.microsoft.com/en-us/9439cb6c-f2f7-4c27-b1d7-8ddf16d81fe8
type RECT struct {
	Left   int32
	Top    int32
	Right  int32
	Bottom int32
}

// MENUITEMINFO https://docs.microsoft.com/en-us/windows/desktop/api/winuser/ns-winuser-tagmenuiteminfow
type MENUITEMINFO struct { //nolint:maligned
	Size         uint32
	Mask         uint32
	Type         uint32
	State        uint32
	ID           uint32
	SubMenu      HMENU
	BMPChecked   HBITMAP
	BMPUnchecked HBITMAP
	ItemData     uintptr
	TypeData     uintptr
	CCH          uint32
	BMPItem      HBITMAP
}

// WNDCLASSEX https://docs.microsoft.com/en-us/windows/desktop/api/winuser/ns-winuser-tagwndclassexw
type WNDCLASSEX struct {
	Size       uint32
	Style      uint32
	WndProc    uintptr
	ClsExtra   int32
	WndExtra   int32
	Instance   HINSTANCE
	Icon       syscall.Handle
	Cursor     HCURSOR
	Background syscall.Handle
	MenuName   *uint16
	ClassName  *uint16
	IconSm     syscall.Handle
}

// DISPLAY_DEVICE https://docs.microsoft.com/en-us/windows/desktop/api/wingdi/ns-wingdi-_display_devicew
type DISPLAY_DEVICE struct {
	Size         uint32
	DeviceName   [32]uint16
	DeviceString [128]uint16
	Flags        uint32
	DeviceID     [128]uint16
	DeviceKey    [128]uint16
}

// DEVMODE https://docs.microsoft.com/en-us/windows/desktop/api/wingdi/ns-wingdi-_devicemodew
type DEVMODE struct {
	DeviceName    [32]uint16
	SpecVersion   uint16
	DriverVersion uint16
	Size          uint16
	DriverExtra   uint16
	Fields        uint32
	X             int32
	Y             int32
	Orientation   uint32
	FixedOutput   uint32
	Color         int16
	Duplex        int16
	YResolution   int16
	TTOption      int16
	Collate       int16
	FormName      [32]uint16
	LogPixels     uint16
	BitsPerPixel  uint32
	PelsWidth     uint32
	PelsHeight    uint32
	Flags         uint32
	Frequency     uint32
	ICMMethod     uint32
	ICMIntent     uint32
	MediaType     uint32
	DitherType    uint32
	Reserved1     uint32
	Reserved2     uint32
	PanningWidth  uint32
	PanningHeight uint32
}

// MONITORINFO https://docs.microsoft.com/en-us/windows/desktop/api/winuser/ns-winuser-tagmonitorinfo
type MONITORINFO struct {
	Size          DWORD
	MonitorBounds RECT
	WorkBounds    RECT
	Flags         DWORD
}

// ACCEL https://docs.microsoft.com/en-us/windows/desktop/menurc/using-keyboard-accelerators#creating-a-run-time-accelerator-table
type ACCEL struct {
	Flags   byte
	padding byte //nolint:structcheck,unused
	Key     WORD
	Cmd     WORD
}
