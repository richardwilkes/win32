package win32

import "syscall"

// RECT https://msdn.microsoft.com/en-us/9439cb6c-f2f7-4c27-b1d7-8ddf16d81fe8
type RECT struct {
	Left   int32
	Top    int32
	Right  int32
	Bottom int32
}

// POINT https://docs.microsoft.com/en-us/previous-versions/dd162805(v=vs.85)
type POINT struct {
	X int32
	Y int32
}

// SIZE https://docs.microsoft.com/en-us/previous-versions/dd145106(v=vs.85)
type SIZE struct {
	CX int32
	CY int32
}

// MSG https://docs.microsoft.com/en-us/windows/desktop/api/winuser/ns-winuser-tagmsg
type MSG struct {
	HWND    HWND
	Message uint32
	WParam  WPARAM
	LParam  LPARAM
	Time    DWORD
	Pt      POINT
	Private DWORD
}

// CIEXYZ https://docs.microsoft.com/en-us/windows/desktop/api/wingdi/ns-wingdi-tagciexyz
type CIEXYZ struct {
	CiexyzX int32
	CiexyzY int32
	CiexyzZ int32
}

// CIEXYZTRIPLE https://docs.microsoft.com/en-us/windows/desktop/api/wingdi/ns-wingdi-tagicexyztriple
type CIEXYZTRIPLE struct {
	CiexyzRed   CIEXYZ
	CiexyzGreen CIEXYZ
	CiexyzBlue  CIEXYZ
}

// BITMAPINFOHEADER https://docs.microsoft.com/en-us/previous-versions/dd183376(v=vs.85)
type BITMAPINFOHEADER struct {
	BiSize          uint32
	BiWidth         int32
	BiHeight        int32
	BiPlanes        uint16
	BiBitCount      uint16
	BiCompression   uint32
	BiSizeImage     uint32
	BiXPelsPerMeter int32
	BiYPelsPerMeter int32
	BiClrUsed       uint32
	BiClrImportant  uint32
}

// BITMAPV4HEADER https://docs.microsoft.com/en-us/windows/desktop/api/wingdi/ns-wingdi-bitmapv4header
type BITMAPV4HEADER struct {
	BITMAPINFOHEADER
	BV4RedMask    uint32
	BV4GreenMask  uint32
	BV4BlueMask   uint32
	BV4AlphaMask  uint32
	BV4CSType     uint32
	BV4Endpoints  CIEXYZTRIPLE
	BV4GammaRed   uint32
	BV4GammaGreen uint32
	BV4GammaBlue  uint32
}

// BITMAPV5HEADER https://docs.microsoft.com/en-us/windows/desktop/api/wingdi/ns-wingdi-bitmapv5header
type BITMAPV5HEADER struct {
	BITMAPV4HEADER
	BV5Intent      uint32
	BV5ProfileData uint32
	BV5ProfileSize uint32
	BV5Reserved    uint32
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

// WINDOWPOS https://docs.microsoft.com/en-us/windows/desktop/api/winuser/ns-winuser-tagwindowpos
type WINDOWPOS struct {
	HwndInsertAfter HWND
	Hwnd            HWND
	X               int32
	Y               int32
	CX              int32
	CY              int32
	Flags           uint32
}

// PAINTSTRUCT https://docs.microsoft.com/en-us/windows/desktop/api/winuser/ns-winuser-tagpaintstruct
type PAINTSTRUCT struct {
	Hdc         HDC
	FErase      BOOL
	RcPaint     RECT
	FRestore    BOOL
	FIncUpdate  BOOL
	RgbReserved [32]byte
}

// XFORM https://docs.microsoft.com/en-us/windows/desktop/api/wingdi/ns-wingdi-tagxform
type XFORM struct {
	EM11 float32
	EM12 float32
	EM21 float32
	EM22 float32
	EDx  float32
	EDy  float32
}

// LOGFONT https://docs.microsoft.com/en-us/windows/desktop/api/wingdi/ns-wingdi-taglogfontw
type LOGFONT struct {
	LfHeight         int32
	LfWidth          int32
	LfEscapement     int32
	LfOrientation    int32
	LfWeight         int32
	LfItalic         byte
	LfUnderline      byte
	LfStrikeOut      byte
	LfCharSet        byte
	LfOutPrecision   byte
	LfClipPrecision  byte
	LfQuality        byte
	LfPitchAndFamily byte
	LfFaceName       [LF_FACESIZE]uint16
}

// TEXTMETRIC https://docs.microsoft.com/en-us/windows/desktop/api/wingdi/ns-wingdi-tagtextmetrica
type TEXTMETRIC struct {
	TmHeight           int32
	TmAscent           int32
	TmDescent          int32
	TmInternalLeading  int32
	TmExternalLeading  int32
	TmAveCharWidth     int32
	TmMaxCharWidth     int32
	TmWeight           int32
	TmOverhang         int32
	TmDigitizedAspectX int32
	TmDigitizedAspectY int32
	TmFirstChar        uint16
	TmLastChar         uint16
	TmDefaultChar      uint16
	TmBreakChar        uint16
	TmItalic           byte
	TmUnderlined       byte
	TmStruckOut        byte
	TmPitchAndFamily   byte
	TmCharSet          byte
}

// NEWTEXTMETRIC https://docs.microsoft.com/en-us/windows/desktop/api/wingdi/ns-wingdi-tagnewtextmetrica
type NEWTEXTMETRIC struct {
	TEXTMETRIC
	NtmFlags      uint32
	NtmSizeEM     uint16
	NtmCellHeight uint16
	NtmAvgWidth   uint16
}
