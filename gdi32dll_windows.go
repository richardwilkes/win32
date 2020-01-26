package win32

import (
	"syscall"
	"unsafe"
)

var (
	gdi32                 = syscall.NewLazyDLL("gdi32.dll")
	angleArc              = gdi32.NewProc("AngleArc")
	arc                   = gdi32.NewProc("Arc")
	arcTo                 = gdi32.NewProc("ArcTo")
	beginPath             = gdi32.NewProc("BeginPath")
	closeFigure           = gdi32.NewProc("CloseFigure")
	createDCW             = gdi32.NewProc("CreateDCW")
	createDIBSection      = gdi32.NewProc("CreateDIBSection")
	createFontW           = gdi32.NewProc("CreateFontW")
	createPen             = gdi32.NewProc("CreatePen")
	createSolidBrush      = gdi32.NewProc("CreateSolidBrush")
	deleteDC              = gdi32.NewProc("DeleteDC")
	deleteObject          = gdi32.NewProc("DeleteObject")
	endPath               = gdi32.NewProc("EndPath")
	enumFontFamiliesExW   = gdi32.NewProc("EnumFontFamiliesExW")
	fillPath              = gdi32.NewProc("FillPath")
	gdiFlush              = gdi32.NewProc("GdiFlush")
	getClipBox            = gdi32.NewProc("GetClipBox")
	getCurrentPositionEx  = gdi32.NewProc("GetCurrentPositionEx")
	getDeviceCaps         = gdi32.NewProc("GetDeviceCaps")
	getTextExtentPoint32W = gdi32.NewProc("GetTextExtentPoint32W")
	getTextMetricsW       = gdi32.NewProc("GetTextMetricsW")
	getWorldTransform     = gdi32.NewProc("GetWorldTransform")
	intersectClipRect     = gdi32.NewProc("IntersectClipRect")
	lineTo                = gdi32.NewProc("LineTo")
	moveToEx              = gdi32.NewProc("MoveToEx")
	polyBezier            = gdi32.NewProc("PolyBezier")
	polyBezierTo          = gdi32.NewProc("PolyBezierTo")
	rectangle             = gdi32.NewProc("Rectangle")
	restoreDC             = gdi32.NewProc("RestoreDC")
	saveDC                = gdi32.NewProc("SaveDC")
	selectClipPath        = gdi32.NewProc("SelectClipPath")
	selectObject          = gdi32.NewProc("SelectObject")
	setArcDirection       = gdi32.NewProc("SetArcDirection")
	setDCBrushColor       = gdi32.NewProc("SetDCBrushColor")
	setDCPenColor         = gdi32.NewProc("SetDCPenColor")
	setGraphicsMode       = gdi32.NewProc("SetGraphicsMode")
	setPolyFillMode       = gdi32.NewProc("SetPolyFillMode")
	setTextAlign          = gdi32.NewProc("SetTextAlign")
	setWorldTransform     = gdi32.NewProc("SetWorldTransform")
	strokePath            = gdi32.NewProc("StrokePath")
	textOutW              = gdi32.NewProc("TextOutW")
)

// AngleArc https://docs.microsoft.com/en-us/windows/desktop/api/wingdi/nf-wingdi-anglearc
func AngleArc(hdc HDC, x, y, r int, startAngle, sweepAngle float32) {
	angleArc.Call(uintptr(hdc), uintptr(x), uintptr(y), uintptr(r), uintptr(startAngle), uintptr(sweepAngle))
}

// Arc https://docs.microsoft.com/en-us/windows/desktop/api/wingdi/nf-wingdi-arc
func Arc(hdc HDC, x1, y1, x2, y2, x3, y3, x4, y4 int) {
	arc.Call(uintptr(hdc), uintptr(x1), uintptr(y1), uintptr(x2), uintptr(y2), uintptr(x3), uintptr(y3), uintptr(x4), uintptr(y4))
}

// ArcTo https://docs.microsoft.com/en-us/windows/desktop/api/wingdi/nf-wingdi-arcto
func ArcTo(hdc HDC, left, top, right, bottom, xr1, yr1, xr2, yr2 int) {
	arcTo.Call(uintptr(hdc), uintptr(left), uintptr(top), uintptr(right), uintptr(bottom), uintptr(xr1), uintptr(yr1), uintptr(xr2), uintptr(yr2))
}

// BeginPath https://docs.microsoft.com/en-us/windows/desktop/api/wingdi/nf-wingdi-beginpath
func BeginPath(hdc HDC) {
	beginPath.Call(uintptr(hdc))
}

// CloseFigure https://docs.microsoft.com/en-us/windows/desktop/api/wingdi/nf-wingdi-closefigure
func CloseFigure(hdc HDC) {
	closeFigure.Call(uintptr(hdc))
}

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
func CreateDIBSection(hdc HDC, pbmi *BITMAPINFOHEADER, usage uint32, ppvBits *unsafe.Pointer, hSection HANDLE, offset uint32) HBITMAP {
	ret, _, _ := createDIBSection.Call(uintptr(hdc), uintptr(unsafe.Pointer(pbmi)), uintptr(usage), uintptr(unsafe.Pointer(ppvBits)), uintptr(hSection), uintptr(offset))
	if ret == ERROR_INVALID_PARAMETER {
		ret = 0
	}
	return HBITMAP(ret)
}

// CreateFont https://docs.microsoft.com/en-us/windows/desktop/api/wingdi/nf-wingdi-createfontw
func CreateFont(height, width, escapement, orientation, weight, italic, underline, strikeOut, charSet, outPrecision, clipPrecision, quality, pitchAndFamily int, faceName string) HFONT {
	if str, err := syscall.UTF16FromString(faceName); err == nil {
		ret, _, _ := createFontW.Call(uintptr(height), uintptr(width), uintptr(escapement), uintptr(orientation), uintptr(weight), uintptr(italic), uintptr(underline), uintptr(strikeOut), uintptr(charSet), uintptr(outPrecision), uintptr(clipPrecision), uintptr(quality), uintptr(pitchAndFamily), uintptr(unsafe.Pointer(&str[0])))
		return HFONT(ret)
	}
	return 0
}

// CreateSolidBrush https://docs.microsoft.com/en-us/windows/desktop/api/wingdi/nf-wingdi-createsolidbrush
func CreateSolidBrush(color COLORREF) HBRUSH {
	ret, _, _ := createSolidBrush.Call(uintptr(color))
	return HBRUSH(ret)
}

// CreatePen https://docs.microsoft.com/en-us/windows/desktop/api/wingdi/nf-wingdi-createpen
func CreatePen(style, width int, color COLORREF) HPEN {
	ret, _, _ := createPen.Call(uintptr(style), uintptr(width), uintptr(color))
	return HPEN(ret)
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

// EndPath https://docs.microsoft.com/en-us/windows/desktop/api/wingdi/nf-wingdi-endpath
func EndPath(hdc HDC) {
	endPath.Call(uintptr(hdc))
}

// EnumFontFamiliesExW https://docs.microsoft.com/en-us/windows/desktop/api/wingdi/nf-wingdi-enumfontfamiliesexw
func EnumFontFamiliesExW(hdc HDC, logFont *LOGFONT, proc func(lf *LOGFONT, tm *NEWTEXTMETRIC, fontType DWORD, parm LPARAM) int, param LPARAM, flags DWORD) int {
	ret, _, _ := enumFontFamiliesExW.Call(uintptr(hdc), uintptr(unsafe.Pointer(logFont)), syscall.NewCallback(func(lf *LOGFONT, tm *NEWTEXTMETRIC, fontType DWORD, parm LPARAM) uintptr {
		return uintptr(proc(lf, tm, fontType, parm))
	}), uintptr(param), uintptr(flags))
	return int(ret)
}

// FillPath https://docs.microsoft.com/en-us/windows/desktop/api/wingdi/nf-wingdi-fillpath
func FillPath(hdc HDC) {
	fillPath.Call(uintptr(hdc))
}

// GdiFlush https://docs.microsoft.com/en-us/windows/desktop/api/wingdi/nf-wingdi-gdiflush
func GdiFlush() bool {
	ret, _, _ := gdiFlush.Call()
	return ret != 0
}

// GetClipBox https://docs.microsoft.com/en-us/windows/desktop/api/wingdi/nf-wingdi-getclipbox
func GetClipBox(hdc HDC, rect *RECT) int {
	ret, _, _ := getClipBox.Call(uintptr(hdc), uintptr(unsafe.Pointer(rect)))
	return int(ret)
}

// GetCurrentPositionEx https://docs.microsoft.com/en-us/windows/desktop/api/wingdi/nf-wingdi-getcurrentpositionex
func GetCurrentPositionEx(hdc HDC, where *POINT) {
	getCurrentPositionEx.Call(uintptr(hdc), uintptr(unsafe.Pointer(where)))
}

// GetDeviceCaps https://docs.microsoft.com/en-us/windows/desktop/api/wingdi/nf-wingdi-getdevicecaps
func GetDeviceCaps(hdc HDC, index int) int {
	ret, _, _ := getDeviceCaps.Call(uintptr(hdc), uintptr(index)) //nolint:errcheck
	return int(ret)
}

// GetTextExtentPoint https://docs.microsoft.com/en-us/windows/desktop/api/wingdi/nf-wingdi-gettextextentpoint32w
func GetTextExtentPoint(hdc HDC, text string, size *SIZE) {
	if str, err := syscall.UTF16FromString(text); err == nil {
		getTextExtentPoint32W.Call(uintptr(hdc), uintptr(unsafe.Pointer(&str[0])), uintptr(len(str)), uintptr(unsafe.Pointer(size)))
	}
}

// GetTextMetrics https://docs.microsoft.com/en-us/windows/desktop/api/wingdi/nf-wingdi-gettextmetricsw
func GetTextMetrics(hdc HDC, tm *NEWTEXTMETRIC) {
	getTextMetricsW.Call(uintptr(hdc), uintptr(unsafe.Pointer(tm)))
}

// GetWorldTransform https://docs.microsoft.com/en-us/windows/desktop/api/Wingdi/nf-wingdi-getworldtransform
func GetWorldTransform(hdc HDC, matrix *XFORM) {
	getWorldTransform.Call(uintptr(hdc), uintptr(unsafe.Pointer(matrix)))
}

// IntersectClipRect https://docs.microsoft.com/en-us/windows/desktop/api/wingdi/nf-wingdi-intersectcliprect
func IntersectClipRect(hdc HDC, left, top, right, bottom int) int {
	ret, _, _ := intersectClipRect.Call(uintptr(hdc), uintptr(left), uintptr(top), uintptr(right), uintptr(bottom))
	return int(ret)
}

// LineTo https://docs.microsoft.com/en-us/windows/desktop/api/wingdi/nf-wingdi-lineto
func LineTo(hdc HDC, x, y int) {
	lineTo.Call(uintptr(hdc), uintptr(x), uintptr(y))
}

// MoveToEx https://docs.microsoft.com/en-us/windows/desktop/api/wingdi/nf-wingdi-movetoex
func MoveToEx(hdc HDC, x, y int, pt *POINT) {
	moveToEx.Call(uintptr(hdc), uintptr(x), uintptr(y), uintptr(unsafe.Pointer(pt)))
}

// PolyBezier https://docs.microsoft.com/en-us/windows/desktop/api/wingdi/nf-wingdi-polybezier
func PolyBezier(hdc HDC, pts []POINT) {
	polyBezier.Call(uintptr(hdc), uintptr(unsafe.Pointer(&pts[0])), uintptr(len(pts)))
}

// PolyBezierTo https://docs.microsoft.com/en-us/windows/desktop/api/wingdi/nf-wingdi-polybezierto
func PolyBezierTo(hdc HDC, pts []POINT) {
	polyBezierTo.Call(uintptr(hdc), uintptr(unsafe.Pointer(&pts[0])), uintptr(len(pts)))
}

// Rectangle https://docs.microsoft.com/en-us/windows/desktop/api/wingdi/nf-wingdi-rectangle
func Rectangle(hdc HDC, left, top, right, bottom int) {
	rectangle.Call(uintptr(hdc), uintptr(left), uintptr(top), uintptr(right), uintptr(bottom))
}

// RestoreDC https://docs.microsoft.com/en-us/windows/desktop/api/wingdi/nf-wingdi-restoredc
func RestoreDC(hdc HDC, nSavedDC int) {
	restoreDC.Call(uintptr(hdc), uintptr(nSavedDC))
}

// SaveDC https://docs.microsoft.com/en-us/windows/desktop/api/wingdi/nf-wingdi-savedc
func SaveDC(hdc HDC) {
	saveDC.Call(uintptr(hdc))
}

// SelectClipPath https://docs.microsoft.com/en-us/windows/desktop/api/wingdi/nf-wingdi-selectclippath
func SelectClipPath(hdc HDC, mode int) {
	selectClipPath.Call(uintptr(hdc), uintptr(mode))
}

// SelectObject https://docs.microsoft.com/en-us/windows/desktop/api/wingdi/nf-wingdi-selectobject
func SelectObject(hdc HDC, obj HGDIOBJ) HGDIOBJ {
	ret, _, _ := selectObject.Call(uintptr(hdc), uintptr(obj))
	return HGDIOBJ(ret)
}

// SetArcDirection https://docs.microsoft.com/en-us/windows/desktop/api/wingdi/nf-wingdi-setarcdirection
func SetArcDirection(hdc HDC, dir int) {
	setArcDirection.Call(uintptr(hdc), uintptr(dir))
}

// SetDCBrushColor https://docs.microsoft.com/en-us/windows/desktop/api/wingdi/nf-wingdi-setdcbrushcolor
func SetDCBrushColor(hdc HDC, color COLORREF) COLORREF {
	ret, _, _ := setDCBrushColor.Call(uintptr(hdc), uintptr(color))
	return COLORREF(ret)
}

// SetDCPenColor https://docs.microsoft.com/en-us/windows/desktop/api/wingdi/nf-wingdi-setdcpencolor
func SetDCPenColor(hdc HDC, color COLORREF) COLORREF {
	ret, _, _ := setDCPenColor.Call(uintptr(hdc), uintptr(color))
	return COLORREF(ret)
}

// SetGraphicsMode https://docs.microsoft.com/en-us/windows/desktop/api/wingdi/nf-wingdi-setgraphicsmode
func SetGraphicsMode(hdc HDC, mode int) int {
	ret, _, _ := setGraphicsMode.Call(uintptr(hdc), uintptr(mode))
	return int(ret)
}

// SetPolyFillMode https://docs.microsoft.com/en-us/windows/desktop/api/wingdi/nf-wingdi-setpolyfillmode
func SetPolyFillMode(hdc HDC, mode int) int {
	ret, _, _ := setPolyFillMode.Call(uintptr(hdc), uintptr(mode))
	return int(ret)
}

// SetTextAlign https://docs.microsoft.com/en-us/windows/desktop/api/wingdi/nf-wingdi-settextalign
func SetTextAlign(hdc HDC, align int) {
	setTextAlign.Call(uintptr(hdc), uintptr(align))
}

// SetWorldTransform https://docs.microsoft.com/en-us/windows/desktop/api/Wingdi/nf-wingdi-setworldtransform
func SetWorldTransform(hdc HDC, matrix *XFORM) {
	setWorldTransform.Call(uintptr(hdc), uintptr(unsafe.Pointer(matrix)))
}

// StrokePath https://docs.microsoft.com/en-us/windows/desktop/api/wingdi/nf-wingdi-strokepath
func StrokePath(hdc HDC) {
	strokePath.Call(uintptr(hdc))
}

// TextOut https://docs.microsoft.com/en-us/windows/desktop/api/wingdi/nf-wingdi-textoutw
func TextOut(hdc HDC, x, y int, text string) {
	if str, err := syscall.UTF16FromString(text); err == nil {
		textOutW.Call(uintptr(hdc), uintptr(x), uintptr(y), uintptr(unsafe.Pointer(&str[0])), uintptr(len(str)))
	}
}
