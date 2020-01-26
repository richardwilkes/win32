package d2d

import (
	"syscall"
	"unsafe"

	"github.com/richardwilkes/win32"
	"github.com/richardwilkes/win32/com"
)

// RenderTargetGUID holds the GUID for the RenderTarget Class.
var RenderTargetGUID = com.GUID{Data1: 0x2cd90694, Data2: 0x12e2, Data3: 0x11dc, Data4: [8]byte{0x9f, 0xed, 0x00, 0x11, 0x43, 0xa0, 0x55, 0xf9}}

// RenderTarget https://docs.microsoft.com/en-us/windows/win32/api/d2d1/nn-d2d1-id2d1rendertarget
type RenderTarget struct {
	Resource
}

type vmtRenderTarget struct {
	vmtResource
	CreateBitmap                 uintptr
	CreateBitmapFromWicBitmap    uintptr
	CreateSharedBitmap           uintptr
	CreateBitmapBrush            uintptr
	CreateSolidColorBrush        uintptr
	CreateGradientStopCollection uintptr
	CreateLinearGradientBrush    uintptr
	CreateRadialGradientBrush    uintptr
	CreateCompatibleRenderTarget uintptr
	CreateLayer                  uintptr
	CreateMesh                   uintptr
	DrawLine                     uintptr
	DrawRectangle                uintptr
	FillRectangle                uintptr
	DrawRoundedRectangle         uintptr
	FillRoundedRectangle         uintptr
	DrawEllipse                  uintptr
	FillEllipse                  uintptr
	DrawGeometry                 uintptr
	FillGeometry                 uintptr
	FillMesh                     uintptr
	FillOpacityMask              uintptr
	DrawBitmap                   uintptr
	DrawText                     uintptr
	DrawTextLayout               uintptr
	DrawGlyphRun                 uintptr
	SetTransform                 uintptr
	GetTransform                 uintptr
	SetAntialiasMode             uintptr
	GetAntialiasMode             uintptr
	SetTextAntialiasMode         uintptr
	GetTextAntialiasMode         uintptr
	SetTextRenderingParams       uintptr
	GetTextRenderingParams       uintptr
	SetTags                      uintptr
	GetTags                      uintptr
	PushLayer                    uintptr
	PopLayer                     uintptr
	Flush                        uintptr
	SaveDrawingState             uintptr
	RestoreDrawingState          uintptr
	PushAxisAlignedClip          uintptr
	PopAxisAlignedClip           uintptr
	Clear                        uintptr
	BeginDraw                    uintptr
	EndDraw                      uintptr
	GetPixelFormat               uintptr
	SetDpi                       uintptr
	GetDpi                       uintptr
	GetSize                      uintptr
	GetPixelSize                 uintptr
	GetMaximumBitmapSize         uintptr
	IsSupported                  uintptr
}

func (obj *RenderTarget) vmt() *vmtRenderTarget {
	return (*vmtRenderTarget)(obj.UnsafeVirtualMethodTable)
}

// CreateBitmap https://docs.microsoft.com/en-us/windows/win32/api/d2d1/nf-d2d1-id2d1rendertarget-createbitmap%28d2d1_size_u_constvoid_uint32_constd2d1_bitmap_properties_id2d1bitmap%29
func (obj *RenderTarget) CreateBitmap(size SizeU, srcData unsafe.Pointer, pitch uint32, bitmapProperties *BitmapProperties) *Bitmap {
	var bitmap *Bitmap
	if ret, _, _ := syscall.Syscall6(obj.vmt().CreateBitmap, 6, uintptr(unsafe.Pointer(obj)), uintptr(*(*uint64)(unsafe.Pointer(&size))), uintptr(srcData), uintptr(pitch), uintptr(unsafe.Pointer(bitmapProperties)), uintptr(unsafe.Pointer(&bitmap))); ret != win32.S_OK {
		return nil
	}
	return bitmap
}

// CreateBitmapFromWicBitmap https://docs.microsoft.com/en-us/windows/win32/api/d2d1/nf-d2d1-id2d1rendertarget-createbitmapfromwicbitmap%28iwicbitmapsource_constd2d1_bitmap_properties_id2d1bitmap%29
func (obj *RenderTarget) CreateBitmapFromWicBitmap(wicBitmapSource *WicBitmapSource, bitmapProperties *BitmapProperties) *Bitmap {
	var bitmap *Bitmap
	if ret, _, _ := syscall.Syscall6(obj.vmt().CreateBitmapFromWicBitmap, 4, uintptr(unsafe.Pointer(obj)), uintptr(unsafe.Pointer(wicBitmapSource)), uintptr(unsafe.Pointer(bitmapProperties)), uintptr(unsafe.Pointer(&bitmap)), 0, 0); ret != win32.S_OK {
		return nil
	}
	return bitmap
}

// CreateSharedBitmap https://docs.microsoft.com/en-us/windows/win32/api/d2d1/nf-d2d1-id2d1rendertarget-createsharedbitmap
func (obj *RenderTarget) CreateSharedBitmap(riid *com.GUID, data unsafe.Pointer, bitmapProperties *BitmapProperties) *Bitmap {
	var bitmap *Bitmap
	if ret, _, _ := syscall.Syscall6(obj.vmt().CreateSharedBitmap, 5, uintptr(unsafe.Pointer(obj)), uintptr(unsafe.Pointer(riid)), uintptr(data), uintptr(unsafe.Pointer(bitmapProperties)), uintptr(unsafe.Pointer(&bitmap)), 0); ret != win32.S_OK {
		return nil
	}
	return bitmap
}

// CreateBitmapBrush https://docs.microsoft.com/en-us/windows/win32/api/d2d1/nf-d2d1-id2d1rendertarget-createbitmapbrush%28id2d1bitmap_id2d1bitmapbrush%29
func (obj *RenderTarget) CreateBitmapBrush(bitmap *Bitmap, bitmapBrushProperties *BitmapBrushProperties, brushProperties *BrushProperties) *BitmapBrush {
	var bitmapBrush *BitmapBrush
	if ret, _, _ := syscall.Syscall6(obj.vmt().CreateBitmapBrush, 5, uintptr(unsafe.Pointer(obj)), uintptr(unsafe.Pointer(bitmap)), uintptr(unsafe.Pointer(bitmapBrushProperties)), uintptr(unsafe.Pointer(brushProperties)), uintptr(unsafe.Pointer(&bitmapBrush)), 0); ret != win32.S_OK {
		return nil
	}
	return bitmapBrush
}

// CreateSolidColorBrush https://docs.microsoft.com/en-us/windows/win32/api/d2d1/nf-d2d1-id2d1rendertarget-createsolidcolorbrush%28constd2d1_color_f_constd2d1_brush_properties_id2d1solidcolorbrush%29
func (obj *RenderTarget) CreateSolidColorBrush(color Color, brushProperties *BrushProperties) *SolidColorBrush {
	var solidColorBrush *SolidColorBrush
	if ret, _, _ := syscall.Syscall6(obj.vmt().CreateSolidColorBrush, 4, uintptr(unsafe.Pointer(obj)), uintptr(unsafe.Pointer(&color)), uintptr(unsafe.Pointer(brushProperties)), uintptr(unsafe.Pointer(&solidColorBrush)), 0, 0); ret != win32.S_OK {
		return nil
	}
	return solidColorBrush
}

// CreateGradientStopCollection https://docs.microsoft.com/en-us/windows/win32/api/d2d1/nf-d2d1-id2d1rendertarget-creategradientstopcollection(constd2d1_gradient_stop_uint32_d2d1_gamma_d2d1_extend_mode_id2d1gradientstopcollection)
func (obj *RenderTarget) CreateGradientStopCollection(gradientStops []GradientStop, useGamma1 bool, extendMode ExtendMode) *GradientStopCollection {
	var gradientStopCollection *GradientStopCollection
	var colorInterpolationGamma uint32
	if useGamma1 {
		colorInterpolationGamma = 1
	}
	if ret, _, _ := syscall.Syscall6(obj.vmt().CreateGradientStopCollection, 6, uintptr(unsafe.Pointer(obj)), uintptr(unsafe.Pointer(&(gradientStops[0]))), uintptr(len(gradientStops)), uintptr(colorInterpolationGamma), uintptr(extendMode), uintptr(unsafe.Pointer(&gradientStopCollection))); ret != win32.S_OK {
		return nil
	}
	return gradientStopCollection
}

// CreateLinearGradientBrush https://docs.microsoft.com/en-us/windows/win32/api/d2d1/nf-d2d1-id2d1rendertarget-createlineargradientbrush%28constd2d1_linear_gradient_brush_properties_constd2d1_brush_properties_id2d1gradientstopcollection_id2d1lineargradientbrush%29
func (obj *RenderTarget) CreateLinearGradientBrush(linearGradientBrushProperties *LinearGradientBrushProperties, brushProperties *BrushProperties, gradientStopCollection *GradientStopCollection) *LinearGradientBrush {
	var linearGradientBrush *LinearGradientBrush
	if ret, _, _ := syscall.Syscall6(obj.vmt().CreateLinearGradientBrush, 5, uintptr(unsafe.Pointer(obj)), uintptr(unsafe.Pointer(linearGradientBrushProperties)), uintptr(unsafe.Pointer(brushProperties)), uintptr(unsafe.Pointer(gradientStopCollection)), uintptr(unsafe.Pointer(&linearGradientBrush)), 0); ret != win32.S_OK {
		return nil
	}
	return linearGradientBrush
}

// CreateRadialGradientBrush https://docs.microsoft.com/en-us/windows/win32/api/d2d1/nf-d2d1-id2d1rendertarget-createradialgradientbrush%28constd2d1_radial_gradient_brush_properties_constd2d1_brush_properties_id2d1gradientstopcollection_id2d1radialgradientbrush%29
func (obj *RenderTarget) CreateRadialGradientBrush(radialGradientBrushProperties *RadialGradientBrushProperties, brushProperties *BrushProperties, gradientStopCollection *GradientStopCollection) *RadialGradientBrush {
	var radialGradientBrush *RadialGradientBrush
	if ret, _, _ := syscall.Syscall6(obj.vmt().CreateRadialGradientBrush, 5, uintptr(unsafe.Pointer(obj)), uintptr(unsafe.Pointer(radialGradientBrushProperties)), uintptr(unsafe.Pointer(brushProperties)), uintptr(unsafe.Pointer(gradientStopCollection)), uintptr(unsafe.Pointer(&radialGradientBrush)), 0); ret != win32.S_OK {
		return nil
	}
	return radialGradientBrush
}

// CreateCompatibleRenderTarget https://docs.microsoft.com/en-us/windows/win32/api/d2d1/nf-d2d1-id2d1rendertarget-createcompatiblerendertarget(d2d1_size_f_d2d1_size_u_d2d1_pixel_format_id2d1bitmaprendertarget)
func (obj *RenderTarget) CreateCompatibleRenderTarget(desiredSize *Size, desiredPixelSize *SizeU, desiredFormat *PixelFormat, GDICompatible bool) *BitmapRenderTarget {
	var bitmapRenderTarget *BitmapRenderTarget
	var options uint32
	if GDICompatible {
		options = 1
	}
	if ret, _, _ := syscall.Syscall6(obj.vmt().CreateCompatibleRenderTarget, 6, uintptr(unsafe.Pointer(obj)), uintptr(unsafe.Pointer(&desiredSize)), uintptr(unsafe.Pointer(desiredPixelSize)), uintptr(unsafe.Pointer(desiredFormat)), uintptr(options), uintptr(unsafe.Pointer(&bitmapRenderTarget))); ret != win32.S_OK {
		return nil
	}
	return bitmapRenderTarget
}

// CreateLayer https://docs.microsoft.com/en-us/windows/win32/api/d2d1/nf-d2d1-id2d1rendertarget-createlayer(d2d1_size_f_id2d1layer)
func (obj *RenderTarget) CreateLayer(size Size) *Layer {
	var layer *Layer
	if ret, _, _ := syscall.Syscall(obj.vmt().CreateLayer, 3, uintptr(unsafe.Pointer(obj)), uintptr(unsafe.Pointer(&size)), uintptr(unsafe.Pointer(&layer))); ret != win32.S_OK {
		return nil
	}
	return layer
}

// CreateMesh https://docs.microsoft.com/en-us/windows/win32/api/d2d1/nf-d2d1-id2d1rendertarget-createmesh
func (obj *RenderTarget) CreateMesh() *Mesh {
	var mesh *Mesh
	if ret, _, _ := syscall.Syscall(obj.vmt().CreateMesh, 2, uintptr(unsafe.Pointer(obj)), uintptr(unsafe.Pointer(&mesh)), 0); ret != win32.S_OK {
		return nil
	}
	return mesh
}

// DrawLine https://docs.microsoft.com/en-us/windows/win32/api/d2d1/nf-d2d1-id2d1rendertarget-drawline
func (obj *RenderTarget) DrawLine(p0 Point, p1 Point, brush *Brush, strokeWidth float32, strokeStyle *StrokeStyle) {
	syscall.Syscall6(obj.vmt().DrawLine, 6, uintptr(unsafe.Pointer(obj)), uintptr(*(*uint64)(unsafe.Pointer(&p0))), uintptr(*(*uint64)(unsafe.Pointer(&p1))), uintptr(unsafe.Pointer(brush)), uintptr(*(*uint32)(unsafe.Pointer(&strokeWidth))), uintptr(unsafe.Pointer(strokeStyle)))
}

// DrawRectangle https://docs.microsoft.com/en-us/windows/win32/api/d2d1/nf-d2d1-id2d1rendertarget-drawrectangle%28constd2d1_rect_f_id2d1brush_float_id2d1strokestyle%29
func (obj *RenderTarget) DrawRectangle(rect Rect, brush *Brush, strokeWidth float32, strokeStyle *StrokeStyle) {
	syscall.Syscall6(obj.vmt().DrawRectangle, 5, uintptr(unsafe.Pointer(obj)), uintptr(unsafe.Pointer(&rect)), uintptr(unsafe.Pointer(brush)), uintptr(*(*uint32)(unsafe.Pointer(&strokeWidth))), uintptr(unsafe.Pointer(strokeStyle)), 0)
}

// FillRectangle https://docs.microsoft.com/en-us/windows/win32/api/d2d1/nf-d2d1-id2d1rendertarget-fillrectangle%28constd2d1_rect_f_id2d1brush%29
func (obj *RenderTarget) FillRectangle(rect Rect, brush *Brush) {
	syscall.Syscall(obj.vmt().FillRectangle, 3, uintptr(unsafe.Pointer(obj)), uintptr(unsafe.Pointer(&rect)), uintptr(unsafe.Pointer(brush)))
}

// DrawRoundedRectangle https://docs.microsoft.com/en-us/windows/win32/api/d2d1/nf-d2d1-id2d1rendertarget-drawroundedrectangle%28constd2d1_rounded_rect__id2d1brush_float_id2d1strokestyle%29
func (obj *RenderTarget) DrawRoundedRectangle(roundedRect RoundedRect, brush *Brush, strokeWidth float32, strokeStyle *StrokeStyle) {
	syscall.Syscall6(obj.vmt().DrawRoundedRectangle, 5, uintptr(unsafe.Pointer(obj)), uintptr(unsafe.Pointer(&roundedRect)), uintptr(unsafe.Pointer(brush)), uintptr(*(*uint32)(unsafe.Pointer(&strokeWidth))), uintptr(unsafe.Pointer(strokeStyle)), 0)
}

// FillRoundedRectangle https://docs.microsoft.com/en-us/windows/win32/api/d2d1/nf-d2d1-id2d1rendertarget-fillroundedrectangle%28constd2d1_rounded_rect_id2d1brush%29
func (obj *RenderTarget) FillRoundedRectangle(roundedRect RoundedRect, brush *Brush) {
	syscall.Syscall(obj.vmt().FillRoundedRectangle, 3, uintptr(unsafe.Pointer(obj)), uintptr(unsafe.Pointer(&roundedRect)), uintptr(unsafe.Pointer(brush)))
}

// DrawEllipse https://docs.microsoft.com/en-us/windows/win32/api/d2d1/nf-d2d1-id2d1rendertarget-drawellipse%28constd2d1_ellipse__id2d1brush_float_id2d1strokestyle%29
func (obj *RenderTarget) DrawEllipse(ellipse Ellipse, brush *Brush, strokeWidth float32, strokeStyle *StrokeStyle) {
	syscall.Syscall6(obj.vmt().DrawEllipse, 5, uintptr(unsafe.Pointer(obj)), uintptr(unsafe.Pointer(&ellipse)), uintptr(unsafe.Pointer(brush)), uintptr(*(*uint32)(unsafe.Pointer(&strokeWidth))), uintptr(unsafe.Pointer(strokeStyle)), 0)
}

// FillEllipse https://docs.microsoft.com/en-us/windows/win32/api/d2d1/nf-d2d1-id2d1rendertarget-fillellipse%28constd2d1_ellipse__id2d1brush%29
func (obj *RenderTarget) FillEllipse(ellipse Ellipse, brush *Brush) {
	syscall.Syscall(obj.vmt().FillEllipse, 3, uintptr(unsafe.Pointer(obj)), uintptr(unsafe.Pointer(&ellipse)), uintptr(unsafe.Pointer(brush)))
}

// DrawGeometry https://docs.microsoft.com/en-us/windows/win32/api/d2d1/nf-d2d1-id2d1rendertarget-drawgeometry
func (obj *RenderTarget) DrawGeometry(geometry *Geometry, brush *Brush, strokeWidth float32, strokeStyle *StrokeStyle) {
	syscall.Syscall6(obj.vmt().DrawGeometry, 5, uintptr(unsafe.Pointer(obj)), uintptr(unsafe.Pointer(geometry)), uintptr(unsafe.Pointer(brush)), uintptr(*(*uint32)(unsafe.Pointer(&strokeWidth))), uintptr(unsafe.Pointer(strokeStyle)), 0)
}

// FillGeometry https://docs.microsoft.com/en-us/windows/win32/api/d2d1/nf-d2d1-id2d1rendertarget-fillgeometry
func (obj *RenderTarget) FillGeometry(geometry *Geometry, brush *Brush, opacityBrush *Brush) {
	syscall.Syscall6(obj.vmt().FillGeometry, 4, uintptr(unsafe.Pointer(obj)), uintptr(unsafe.Pointer(geometry)), uintptr(unsafe.Pointer(brush)), uintptr(unsafe.Pointer(opacityBrush)), 0, 0)
}

// FillMesh https://docs.microsoft.com/en-us/windows/win32/api/d2d1/nf-d2d1-id2d1rendertarget-fillmesh
func (obj *RenderTarget) FillMesh(mesh *Mesh, brush *Brush) {
	syscall.Syscall(obj.vmt().FillMesh, 3, uintptr(unsafe.Pointer(obj)), uintptr(unsafe.Pointer(mesh)), uintptr(unsafe.Pointer(brush)))
}

// FillOpacityMask https://docs.microsoft.com/en-us/windows/win32/api/d2d1/nf-d2d1-id2d1rendertarget-fillopacitymask%28id2d1bitmap_id2d1brush_d2d1_opacity_mask_content_constd2d1_rect_f__constd2d1_rect_f_%29
func (obj *RenderTarget) FillOpacityMask(opacityMask *Bitmap, brush *Brush, content OpacityMaskContent, destinationRectangle, sourceRectangle Rect) {
	syscall.Syscall6(obj.vmt().FillOpacityMask, 6, uintptr(unsafe.Pointer(obj)), uintptr(unsafe.Pointer(opacityMask)), uintptr(unsafe.Pointer(brush)), uintptr(content), uintptr(unsafe.Pointer(&destinationRectangle)), uintptr(unsafe.Pointer(&sourceRectangle)))
}

// DrawBitmap https://docs.microsoft.com/en-us/windows/win32/api/d2d1/nf-d2d1-id2d1rendertarget-drawbitmap(id2d1bitmap_constd2d1_rect_f__float_d2d1_bitmap_interpolation_mode_constd2d1_rect_f_)
func (obj *RenderTarget) DrawBitmap(bitmap *Bitmap, destinationRectangle Rect, opacity float32, useLinearInterpolation bool, sourceRectangle Rect) {
	var interpolationMode uint32
	if useLinearInterpolation {
		interpolationMode = 1
	}
	syscall.Syscall6(obj.vmt().DrawBitmap, 6, uintptr(unsafe.Pointer(obj)), uintptr(unsafe.Pointer(bitmap)), uintptr(unsafe.Pointer(&destinationRectangle)), uintptr(*(*uint32)(unsafe.Pointer(&opacity))), uintptr(interpolationMode), uintptr(unsafe.Pointer(&sourceRectangle)))
}

// DrawText https://docs.microsoft.com/en-us/windows/win32/api/d2d1/nf-d2d1-id2d1rendertarget-drawtext%28constwchar_uint32_idwritetextformat_constd2d1_rect_f_id2d1brush_d2d1_draw_text_options_dwrite_measuring_mode%29
func (obj *RenderTarget) DrawText(str string, textFormat *TextFormat, layoutRect Rect, defaultForegroundBrush *Brush, options DrawTextOptions, measuringMode MeasuringMode) {
	s, err := syscall.UTF16FromString(str)
	if err != nil {
		s = []uint16{0}
	}
	syscall.Syscall9(obj.vmt().DrawText, 8, uintptr(unsafe.Pointer(obj)), uintptr(unsafe.Pointer(&(s[0]))), uintptr(len(s)), uintptr(unsafe.Pointer(textFormat)), uintptr(unsafe.Pointer(&layoutRect)), uintptr(unsafe.Pointer(defaultForegroundBrush)), uintptr(options), uintptr(measuringMode), 0)
}

// DrawTextLayout https://docs.microsoft.com/en-us/windows/win32/api/d2d1/nf-d2d1-id2d1rendertarget-drawtextlayout
func (obj *RenderTarget) DrawTextLayout(origin Point, textLayout *TextLayout, defaultForegroundBrush *Brush, options DrawTextOptions) {
	syscall.Syscall6(obj.vmt().DrawTextLayout, 5, uintptr(unsafe.Pointer(obj)), uintptr(*(*uint64)(unsafe.Pointer(&origin))), uintptr(unsafe.Pointer(textLayout)), uintptr(unsafe.Pointer(defaultForegroundBrush)), uintptr(options), 0)
}

// DrawGlyphRun https://docs.microsoft.com/en-us/windows/win32/api/d2d1/nf-d2d1-id2d1rendertarget-drawglyphrun
func (obj *RenderTarget) DrawGlyphRun(baselineOrigin Point, glyphRun *GlyphRun, foregroundBrush *Brush, measuringMode MeasuringMode) {
	syscall.Syscall6(obj.vmt().DrawGlyphRun, 5, uintptr(unsafe.Pointer(obj)), uintptr(*(*uint64)(unsafe.Pointer(&baselineOrigin))), uintptr(unsafe.Pointer(glyphRun)), uintptr(unsafe.Pointer(foregroundBrush)), uintptr(measuringMode), 0)
}

// SetTransform https://docs.microsoft.com/en-us/windows/win32/api/d2d1/nf-d2d1-id2d1rendertarget-settransform%28constd2d1_matrix_3x2_f%29
func (obj *RenderTarget) SetTransform(transform *Matrix3x2) {
	syscall.Syscall(obj.vmt().SetTransform, 2, uintptr(unsafe.Pointer(obj)), uintptr(unsafe.Pointer(transform)), 0)
}

// Transform https://docs.microsoft.com/en-us/windows/win32/api/d2d1/nf-d2d1-id2d1rendertarget-gettransform
func (obj *RenderTarget) Transform() *Matrix3x2 {
	var transform Matrix3x2
	syscall.Syscall(obj.vmt().GetTransform, 2, uintptr(unsafe.Pointer(obj)), uintptr(unsafe.Pointer(&transform)), 0)
	return &transform
}

// SetAntialiasMode https://docs.microsoft.com/en-us/windows/win32/api/d2d1/nf-d2d1-id2d1rendertarget-setantialiasmode
func (obj *RenderTarget) SetAntialiasMode(aliased bool) {
	var antialiasMode uint32
	if aliased {
		antialiasMode = 1
	}
	syscall.Syscall(obj.vmt().SetAntialiasMode, 2, uintptr(unsafe.Pointer(obj)), uintptr(antialiasMode), 0)
}

// IsAntialiasModePerPrimitive https://docs.microsoft.com/en-us/windows/win32/api/d2d1/nf-d2d1-id2d1rendertarget-getantialiasmode
func (obj *RenderTarget) IsAntialiasModePerPrimitive() bool {
	ret, _, _ := syscall.Syscall(obj.vmt().GetAntialiasMode, 1, uintptr(unsafe.Pointer(obj)), 0, 0)
	return ret == 0
}

// IsAntialiasModeAliased https://docs.microsoft.com/en-us/windows/win32/api/d2d1/nf-d2d1-id2d1rendertarget-getantialiasmode
func (obj *RenderTarget) IsAntialiasModeAliased() bool {
	ret, _, _ := syscall.Syscall(obj.vmt().GetAntialiasMode, 1, uintptr(unsafe.Pointer(obj)), 0, 0)
	return ret == 1
}

// SetTextAntialiasMode https://docs.microsoft.com/en-us/windows/win32/api/d2d1/nf-d2d1-id2d1rendertarget-settextantialiasmode
func (obj *RenderTarget) SetTextAntialiasMode(textAntialiasMode TextAntialiasMode) {
	syscall.Syscall(obj.vmt().SetTextAntialiasMode, 2, uintptr(unsafe.Pointer(obj)), uintptr(textAntialiasMode), 0)
}

// TextAntialiasMode https://docs.microsoft.com/en-us/windows/win32/api/d2d1/nf-d2d1-id2d1rendertarget-gettextantialiasmode
func (obj *RenderTarget) TextAntialiasMode() TextAntialiasMode {
	ret, _, _ := syscall.Syscall(obj.vmt().GetTextAntialiasMode, 1, uintptr(unsafe.Pointer(obj)), 0, 0)
	return TextAntialiasMode(ret)
}

// SetTextRenderingParams https://docs.microsoft.com/en-us/windows/win32/api/d2d1/nf-d2d1-id2d1rendertarget-settextrenderingparams
func (obj *RenderTarget) SetTextRenderingParams(textRenderingParams *TextRenderingParams) {
	syscall.Syscall(obj.vmt().SetTextRenderingParams, 2, uintptr(unsafe.Pointer(obj)), uintptr(unsafe.Pointer(textRenderingParams)), 0)
}

// TextRenderingParams https://docs.microsoft.com/en-us/windows/win32/api/d2d1/nf-d2d1-id2d1rendertarget-gettextrenderingparams
func (obj *RenderTarget) TextRenderingParams() *TextRenderingParams {
	var textRenderingParams *TextRenderingParams
	syscall.Syscall(obj.vmt().GetTextRenderingParams, 2, uintptr(unsafe.Pointer(obj)), uintptr(unsafe.Pointer(&textRenderingParams)), 0)
	return textRenderingParams
}

// SetTags https://docs.microsoft.com/en-us/windows/win32/api/d2d1/nf-d2d1-id2d1rendertarget-settags
func (obj *RenderTarget) SetTags(tag1, tag2 Tag) {
	syscall.Syscall(obj.vmt().SetTags, 3, uintptr(unsafe.Pointer(obj)), uintptr(tag1), uintptr(tag2))
}

// Tags https://docs.microsoft.com/en-us/windows/win32/api/d2d1/nf-d2d1-id2d1rendertarget-gettags
func (obj *RenderTarget) Tags() (tag1, tag2 Tag) {
	syscall.Syscall(obj.vmt().GetTags, 3, uintptr(unsafe.Pointer(obj)), uintptr(unsafe.Pointer(&tag1)), uintptr(unsafe.Pointer(&tag2)))
	return
}

// PushLayer https://docs.microsoft.com/en-us/windows/win32/api/d2d1/nf-d2d1-id2d1rendertarget-pushlayer%28constd2d1_layer_parameters_id2d1layer%29
// Must have a matching PopLayer call.
func (obj *RenderTarget) PushLayer(layerParameters *LayerParameters, layer *Layer) {
	syscall.Syscall(obj.vmt().PushLayer, 3, uintptr(unsafe.Pointer(obj)), uintptr(unsafe.Pointer(layerParameters)), uintptr(unsafe.Pointer(layer)))
}

// PopLayer https://docs.microsoft.com/en-us/windows/win32/api/d2d1/nf-d2d1-id2d1rendertarget-poplayer
func (obj *RenderTarget) PopLayer() {
	syscall.Syscall(obj.vmt().PopLayer, 1, uintptr(unsafe.Pointer(obj)), 0, 0)
}

// Flush https://docs.microsoft.com/en-us/windows/win32/api/d2d1/nf-d2d1-id2d1rendertarget-flush
func (obj *RenderTarget) Flush() (tag1, tag2 Tag) {
	syscall.Syscall(obj.vmt().Flush, 3, uintptr(unsafe.Pointer(obj)), uintptr(unsafe.Pointer(&tag1)), uintptr(unsafe.Pointer(&tag2)))
	return
}

// SaveDrawingState https://docs.microsoft.com/en-us/windows/win32/api/d2d1/nf-d2d1-id2d1rendertarget-savedrawingstate
func (obj *RenderTarget) SaveDrawingState() *DrawingStateBlock {
	var drawingStateBlock DrawingStateBlock
	syscall.Syscall(obj.vmt().SaveDrawingState, 2, uintptr(unsafe.Pointer(obj)), uintptr(unsafe.Pointer(&drawingStateBlock)), 0)
	return &drawingStateBlock
}

// RestoreDrawingState https://docs.microsoft.com/en-us/windows/win32/api/d2d1/nf-d2d1-id2d1rendertarget-restoredrawingstate
func (obj *RenderTarget) RestoreDrawingState(drawingStateBlock *DrawingStateBlock) {
	syscall.Syscall(obj.vmt().RestoreDrawingState, 2, uintptr(unsafe.Pointer(obj)), uintptr(unsafe.Pointer(drawingStateBlock)), 0)
}

// PushAxisAlignedClip https://docs.microsoft.com/en-us/windows/win32/api/d2d1/nf-d2d1-id2d1rendertarget-pushaxisalignedclip(constd2d1_rect_f__d2d1_antialias_mode)
// Must have a matching PopAxisAlignedClip call.
func (obj *RenderTarget) PushAxisAlignedClip(clipRect Rect, aliased bool) {
	var antialiasMode uint32
	if aliased {
		antialiasMode = 1
	}
	syscall.Syscall(obj.vmt().PushAxisAlignedClip, 3, uintptr(unsafe.Pointer(obj)), uintptr(unsafe.Pointer(&clipRect)), uintptr(antialiasMode))
}

// PopAxisAlignedClip https://docs.microsoft.com/en-us/windows/win32/api/d2d1/nf-d2d1-id2d1rendertarget-popaxisalignedclip
func (obj *RenderTarget) PopAxisAlignedClip() {
	syscall.Syscall(obj.vmt().PopAxisAlignedClip, 1, uintptr(unsafe.Pointer(obj)), 0, 0)
}

// Clear https://docs.microsoft.com/en-us/windows/win32/api/d2d1/nf-d2d1-id2d1rendertarget-clear%28constd2d1_color_f_%29
func (obj *RenderTarget) Clear(color Color) {
	syscall.Syscall(obj.vmt().Clear, 2, uintptr(unsafe.Pointer(obj)), uintptr(unsafe.Pointer(&color)), 0)
}

// BeginDraw https://docs.microsoft.com/en-us/windows/win32/api/d2d1/nf-d2d1-id2d1rendertarget-begindraw
// Must have a matching EndDraw call.
func (obj *RenderTarget) BeginDraw() {
	syscall.Syscall(obj.vmt().BeginDraw, 1, uintptr(unsafe.Pointer(obj)), 0, 0)
}

// EndDraw https://docs.microsoft.com/en-us/windows/win32/api/d2d1/nf-d2d1-id2d1rendertarget-enddraw
func (obj *RenderTarget) EndDraw() (tag1, tag2 Tag) {
	syscall.Syscall(obj.vmt().EndDraw, 3, uintptr(unsafe.Pointer(obj)), uintptr(unsafe.Pointer(&tag1)), uintptr(unsafe.Pointer(&tag2)))
	return
}

// PixelFormat https://docs.microsoft.com/en-us/windows/win32/api/d2d1/nf-d2d1-id2d1rendertarget-getpixelformat
func (obj *RenderTarget) PixelFormat() PixelFormat {
	var format PixelFormat
	syscall.Syscall(obj.vmt().GetPixelFormat, 2, uintptr(unsafe.Pointer(obj)), uintptr(unsafe.Pointer(&format)), 0)
	return format
}

// SetDPI https://docs.microsoft.com/en-us/windows/win32/api/d2d1/nf-d2d1-id2d1rendertarget-setdpi
func (obj *RenderTarget) SetDPI(dpiX, dpiY float32) {
	syscall.Syscall(obj.vmt().SetDpi, 3, uintptr(unsafe.Pointer(obj)), uintptr(*(*uint32)(unsafe.Pointer(&dpiX))), uintptr(*(*uint32)(unsafe.Pointer(&dpiY))))
}

// DPI https://docs.microsoft.com/en-us/windows/win32/api/d2d1/nf-d2d1-id2d1rendertarget-getdpi
func (obj *RenderTarget) DPI() (dpiX, dpiY float32) {
	syscall.Syscall(obj.vmt().GetDpi, 3, uintptr(unsafe.Pointer(obj)), uintptr(unsafe.Pointer(&dpiX)), uintptr(unsafe.Pointer(&dpiY)))
	return
}

// Size https://docs.microsoft.com/en-us/windows/win32/api/d2d1/nf-d2d1-id2d1rendertarget-getsize
func (obj *RenderTarget) Size() Size {
	var size Size
	syscall.Syscall(obj.vmt().GetSize, 2, uintptr(unsafe.Pointer(obj)), uintptr(unsafe.Pointer(&size)), 0)
	return size
}

// PixelSize https://docs.microsoft.com/en-us/windows/win32/api/d2d1/nf-d2d1-id2d1rendertarget-getpixelsize
func (obj *RenderTarget) PixelSize() SizeU {
	var size SizeU
	syscall.Syscall(obj.vmt().GetPixelSize, 2, uintptr(unsafe.Pointer(obj)), uintptr(unsafe.Pointer(&size)), 0)
	return size
}

// MaximumBitmapSize https://docs.microsoft.com/en-us/windows/win32/api/d2d1/nf-d2d1-id2d1rendertarget-getmaximumbitmapsize
func (obj *RenderTarget) MaximumBitmapSize() uint32 {
	ret, _, _ := syscall.Syscall(obj.vmt().GetMaximumBitmapSize, 1, uintptr(unsafe.Pointer(obj)), 0, 0)
	return uint32(ret)
}

// IsSupported https://docs.microsoft.com/en-us/windows/win32/api/d2d1/nf-d2d1-id2d1rendertarget-issupported%28constd2d1_render_target_properties%29
func (obj *RenderTarget) IsSupported(renderTargetProperties *RenderTargetProperties) bool {
	ret, _, _ := syscall.Syscall(obj.vmt().IsSupported, 2, uintptr(unsafe.Pointer(obj)), uintptr(unsafe.Pointer(renderTargetProperties)), 0)
	return ret != 0
}
