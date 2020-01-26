package d2d

import (
	"github.com/richardwilkes/win32"
)

// Rect https://docs.microsoft.com/en-us/windows/win32/direct2d/d2d1-rect-f
type Rect struct {
	Left   float32
	Top    float32
	Right  float32
	Bottom float32
}

// Size https://docs.microsoft.com/en-us/windows/win32/direct2d/d2d1-size-f
type Size struct {
	Width  float32
	Height float32
}

// SizeU https://docs.microsoft.com/en-us/windows/win32/direct2d/d2d1-size-u
type SizeU struct {
	Width  uint32
	Height uint32
}

// RoundedRect https://docs.microsoft.com/en-us/windows/win32/api/d2d1/ns-d2d1-d2d1_rounded_rect
type RoundedRect struct {
	Rect    Rect
	RadiusX float32
	RadiusY float32
}

// Ellipse https://docs.microsoft.com/en-us/windows/win32/api/d2d1/ns-d2d1-d2d1_ellipse
type Ellipse struct {
	Point   Point
	RadiusX float32
	RadiusY float32
}

// Point https://docs.microsoft.com/en-us/windows/win32/direct2d/d2d1-point-2f
type Point struct {
	X float32
	Y float32
}

// Matrix3x2 https://docs.microsoft.com/en-us/windows/win32/api/dcommon/ns-dcommon-d2d_matrix_3x2_f
type Matrix3x2 struct {
	A11 float32
	A12 float32
	A21 float32
	A22 float32
	A31 float32
	A32 float32
}

// Triangle https://docs.microsoft.com/en-us/windows/win32/api/d2d1/ns-d2d1-d2d1_triangle
type Triangle struct {
	Point1 Point
	Point2 Point
	Point3 Point
}

// BezierSegment https://docs.microsoft.com/en-us/windows/win32/api/d2d1/ns-d2d1-d2d1_bezier_segment
type BezierSegment struct {
	Point1 Point
	Point2 Point
	Point3 Point
}

// QuadraticBezierSegment https://docs.microsoft.com/en-us/windows/win32/api/d2d1/ns-d2d1-d2d1_quadratic_bezier_segment
type QuadraticBezierSegment struct {
	Point1 Point
	Point2 Point
}

// ArcSegment https://docs.microsoft.com/en-us/windows/win32/api/d2d1/ns-d2d1-d2d1_arc_segment
type ArcSegment struct {
	Point          Point
	Size           Size
	RotationAngle  float32
	SweepClockwise uint32 // 0 = Counter Clockwise, 1 = Clockwise
	ArcSize        uint32 // 0 = small, 1 = large
}

// StrokeStyleProperties https://docs.microsoft.com/en-us/windows/win32/api/d2d1/ns-d2d1-d2d1_stroke_style_properties
type StrokeStyleProperties struct {
	StartCap   CapStyle
	EndCap     CapStyle
	DashCap    CapStyle
	LineJoin   LineJoin
	MiterLimit float32
	DashStyle  DashStyle
	DashOffset float32
}

// RenderTargetProperties https://docs.microsoft.com/en-us/windows/win32/api/d2d1/ns-d2d1-d2d1_render_target_properties
type RenderTargetProperties struct {
	Type        RenderTargetType
	PixelFormat PixelFormat
	DpiX        float32
	DpiY        float32
	Usage       RenderTargetUsage
	MinLevel    FeatureLevel
}

// HWNDRenderTargetProperties https://docs.microsoft.com/en-us/windows/win32/api/d2d1/ns-d2d1-d2d1_hwnd_render_target_properties
type HWNDRenderTargetProperties struct {
	HWND           win32.HWND
	PixelSize      SizeU
	PresentOptions PresentOptions
}

// PixelFormat https://docs.microsoft.com/en-us/windows/win32/api/dcommon/ns-dcommon-d2d1_pixel_format
type PixelFormat struct {
	Format    Format
	AlphaMode AlphaMode
}

// Color https://docs.microsoft.com/en-us/windows/win32/direct2d/d2d-color-f
type Color struct {
	R float32
	G float32
	B float32
	A float32
}

// BrushProperties https://docs.microsoft.com/en-us/windows/win32/api/d2d1/ns-d2d1-d2d1_brush_properties
type BrushProperties struct {
	Opacity   float32
	Transform Matrix3x2
}

// GradientStop https://docs.microsoft.com/en-us/windows/win32/api/d2d1/ns-d2d1-d2d1_gradient_stop
type GradientStop struct {
	Position float32
	Color    Color
}

// LinearGradientBrushProperties https://docs.microsoft.com/en-us/windows/win32/api/d2d1/ns-d2d1-d2d1_linear_gradient_brush_properties
type LinearGradientBrushProperties struct {
	StartPoint Point
	EndPoint   Point
}

// RadialGradientBrushProperties https://docs.microsoft.com/en-us/windows/win32/api/d2d1/ns-d2d1-d2d1_radial_gradient_brush_properties
type RadialGradientBrushProperties struct {
	Center               Point
	GradientOriginOffset Point
	RadiusX              float32
	RadiusY              float32
}

// BitmapProperties https://docs.microsoft.com/en-us/windows/win32/api/d2d1/ns-d2d1-d2d1_bitmap_properties
type BitmapProperties struct {
	PixelFormat PixelFormat
	DpiX        float32
	DpiY        float32
}

// BitmapBrushProperties https://docs.microsoft.com/en-us/windows/win32/api/d2d1/ns-d2d1-d2d1_bitmap_brush_properties
type BitmapBrushProperties struct {
	ExtendModeX       ExtendMode
	ExtendModeY       ExtendMode
	InterpolationMode uint32 // 0 = Nearest Neighbor, 1 = Linear
}

// LayerParameters https://docs.microsoft.com/en-us/windows/win32/api/d2d1/ns-d2d1-d2d1_layer_parameters
type LayerParameters struct {
	ContentBounds     Rect
	GeometricMask     *Geometry
	MaskAntialiasMode uint32 // 0 = Per Primitive, 1 = Aliased
	MaskTransform     Matrix3x2
	Opacity           float32
	OpacityBrush      *Brush
	LayerOptions      uint32 // 0 = None, 1 = Initialize for Cleartype
}

// GlyphOffset https://docs.microsoft.com/en-us/windows/win32/api/dwrite/ns-dwrite-dwrite_glyph_offset
type GlyphOffset struct {
	AdvanceOffset  float32
	AscenderOffset float32
}

// GlyphRun https://docs.microsoft.com/en-us/windows/win32/api/dwrite/ns-dwrite-dwrite_glyph_run
type GlyphRun struct {
	FontFace      *FontFace
	FontEmSize    float32
	GlyphCount    uint32
	GlyphIndices  *uint16
	GlyphAdvances *float32
	GlyphOffsets  *GlyphOffset
	IsSideways    int32
	BidiLevel     uint32
}
