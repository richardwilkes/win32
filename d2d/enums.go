package d2d

// DebugLevel https://docs.microsoft.com/en-us/windows/win32/api/d2d1/ne-d2d1-d2d1_debug_level
type DebugLevel uint32

// Possible values for DebugLevel.
const (
	DebugLevelNone DebugLevel = iota
	DebugLevelError
	DebugLevelWarning
	DebugLevelInfo
)

// CapStyle https://docs.microsoft.com/en-us/windows/win32/api/d2d1/ne-d2d1-d2d1_cap_style
type CapStyle uint32

// Possible values for CapStyle.
const (
	CapStyleFlat CapStyle = iota
	CapStyleSquare
	CapStyleRound
	CapStyleTriangle
)

// LineJoin https://docs.microsoft.com/en-us/windows/win32/api/d2d1/ne-d2d1-d2d1_line_join
type LineJoin uint32

// Possible values for LineJoin.
const (
	LineJoinMiter LineJoin = iota
	LineJoinBevel
	LineJoinRound
	LineJoinMiterOrBevel
)

// DashStyle https://docs.microsoft.com/en-us/windows/win32/api/d2d1/ne-d2d1-d2d1_dash_style
type DashStyle uint32

// Possible values for DashStyle.
const (
	DashStyleSolid DashStyle = iota
	DashStyleDash
	DashStyleDot
	DashStyleDashDot
	DashStyleDashDotDot
	DashStyleCustom
)

// GeometryRelation https://docs.microsoft.com/en-us/windows/win32/api/d2d1/ne-d2d1-d2d1_geometry_relation
type GeometryRelation uint32

// Possible values for GeometryRelation.
const (
	GeometryRelationUnknown GeometryRelation = iota
	GeometryRelationDisjoint
	GeometryRelationIsContained
	GeometryRelationContains
	GeometryRelationOverlap
)

// CombineMode https://docs.microsoft.com/en-us/windows/win32/api/d2d1/ne-d2d1-d2d1_combine_mode
type CombineMode uint32

// Possible values for CombineMode.
const (
	CombineModeUnion CombineMode = iota
	CombineModeIntersect
	CombineModeXor
	CombineModeExclude
)

// PathSegment https://docs.microsoft.com/en-us/windows/win32/api/d2d1/ne-d2d1-d2d1_path_segment
type PathSegment uint32

// Possible values for PathSegment.
const (
	PathSegmentNone PathSegment = iota
	PathSegmentForceUnstroked
	PathSegmentForceRoundLineJoin
)

// TextAntialiasMode https://docs.microsoft.com/en-us/windows/win32/api/d2d1/ne-d2d1-d2d1_text_antialias_mode
type TextAntialiasMode uint32

// Possible values for TextAntialiasMode.
const (
	TextAntialiasModeDefault TextAntialiasMode = iota
	TextAntialiasModeCleartype
	TextAntialiasModeGrayscale
	TextAntialiasModeAliased
)

// Tag https://docs.microsoft.com/en-us/windows/win32/direct2d/d2d1-tag
type Tag uint64

// RenderTargetType https://docs.microsoft.com/en-us/windows/win32/api/d2d1/ne-d2d1-d2d1_render_target_type
type RenderTargetType uint32

// Possible values for RenderTargetType.
const (
	RenderTargetTypeDefault RenderTargetType = iota
	RenderTargetTypeSoftware
	RenderTargetTypeHardware
)

// RenderTargetUsage https://docs.microsoft.com/en-us/windows/win32/api/d2d1/ne-d2d1-d2d1_render_target_usage
type RenderTargetUsage uint32

// Possible values for RenderTargetUsage.
const (
	RenderTargetUsageNone RenderTargetUsage = iota
	RenderTargetUsageForceBitmapRemoting
	RenderTargetUsageGDICompatible
)

// FeatureLevel https://docs.microsoft.com/en-us/windows/win32/api/d2d1/ne-d2d1-d2d1_feature_level
type FeatureLevel uint32

// Possible values for FeatureLevel.
const (
	FeatureLevelDefault FeatureLevel = 0
	FeatureLevel9       FeatureLevel = 0x9100
	FeatureLevel10      FeatureLevel = 0xa100
)

// PresentOptions https://docs.microsoft.com/en-us/windows/win32/api/d2d1/ne-d2d1-d2d1_present_options
type PresentOptions uint32

// Possible values for PresentOptions.
const (
	PresentOptionsNone PresentOptions = iota
	PresentOptionsRetainContents
	PresentOptionsImmediately
)

// Format https://docs.microsoft.com/en-us/windows/win32/api/dxgiformat/ne-dxgiformat-dxgi_format
type Format uint32

// Possible values for Format.
const (
	FormatUnknown Format = iota
	FormatR32G32B32A32Typeless
	FormatR32G32B32A32Float
	FormatR32G32B32A32Uint
	FormatR32G32B32A32Sint
	FormatR32G32B32Typeless
	FormatR32G32B32Float
	FormatR32G32B32Uint
	FormatR32G32B32Sint
	FormatR16G16B16A16Typeless
	FormatR16G16B16A16Float
	FormatR16G16B16A16Unorm
	FormatR16G16B16A16Uint
	FormatR16G16B16A16Snorm
	FormatR16G16B16A16Sint
	FormatR32G32Typeless
	FormatR32G32Float
	FormatR32G32Uint
	FormatR32G32Sint
	FormatR32G8X24Typeless
	FormatD32FloatS8X24Uint
	FormatR32FloatX8X24Typeless
	FormatX32TypelessG8X24Uint
	FormatR10G10B10A2Typeless
	FormatR10G10B10A2Unorm
	FormatR10G10B10A2Uint
	FormatR11G11B10Float
	FormatR8G8B8A8Typeless
	FormatR8G8B8A8Unorm
	R8G8B8A8UnormSrgb
	FormatR8G8B8A8Uint
	FormatR8G8B8A8Snorm
	FormatR8G8B8A8Sint
	FormatR16G16Typeless
	FormatR16G16Float
	FormatR16G16Unorm
	FormatR16G16Uint
	FormatR16G16Snorm
	FormatR16G16Sint
	FormatR32Typeless
	FormatD32Float
	FormatR32Float
	FormatR32Uint
	FormatR32Sint
	FormatR24G8Typeless
	FormatD24UnormS8Uint
	FormatR24UnormX8Typeless
	FormatX24TypelessG8Uint
	FormatR8G8Typeless
	FormatR8G8Unorm
	FormatR8G8Uint
	FormatR8G8Snorm
	FormatR8G8Sint
	FormatR16Typeless
	FormatR16Float
	FormatD16Unorm
	FormatR16Unorm
	FormatR16Uint
	FormatR16Snorm
	FormatR16Sint
	FormatR8Typeless
	FormatR8Unorm
	FormatR8Uint
	FormatR8Snorm
	FormatR8Sint
	FormatA8Unorm
	FormatR1Unorm
	FormatR9G9B9E5Sharedexp
	FormatR8G8B8G8Unorm
	FormatG8R8G8B8Unorm
	FormatBc1Typeless
	FormatBc1Unorm
	FormatBc1UnormSrgb
	FormatBc2Typeless
	FormatBc2Unorm
	FormatBc2UnormSrgb
	FormatBc3Typeless
	FormatBc3Unorm
	FormatBc3UnormSrgb
	FormatBc4Typeless
	FormatBc4Unorm
	FormatBc4Snorm
	FormatBc5Typeless
	FormatBc5Unorm
	FormatBc5Snorm
	FormatB5G6R5Unorm
	FormatB5G5R5A1Unorm
	FormatB8G8R8A8Unorm
	FormatB8G8R8X8Unorm
	FormatR10G10B10XrBiasA2Unorm
	FormatB8G8R8A8Typeless
	FormatB8G8R8A8UnormSrgb
	FormatB8G8R8X8Typeless
	FormatB8G8R8X8UnormSrgb
	FormatBc6HTypeless
	FormatBc6HUf16
	FormatBc6HSf16
	FormatBc7Typeless
	FormatBc7Unorm
	FormatBc7UnormSrgb
)

// AlphaMode https://docs.microsoft.com/en-us/windows/win32/api/dcommon/ne-dcommon-d2d1_alpha_mode
type AlphaMode uint32

// Possible values for AlphaMode.
const (
	AlphaModeUnknown AlphaMode = iota
	AlphaModePremultiplied
	AlphaModeStraight
	AlphaModeIgnore
)

// ExtendMode https://docs.microsoft.com/en-us/windows/win32/api/d2d1/ne-d2d1-d2d1_extend_mode
type ExtendMode uint32

// Possible values for ExtendMode.
const (
	ExtendModeClamp ExtendMode = iota
	ExtendModeWrap
	ExtendModeMirror
)

// DrawTextOptions https://docs.microsoft.com/en-us/windows/win32/api/d2d1/ne-d2d1-d2d1_draw_text_options
type DrawTextOptions uint32

// Possible values for DrawTextOptions.
const (
	DrawTextOptionsNone DrawTextOptions = iota
	DrawTextOptionsNoSnap
	DrawTextOptionsClip
)

// OpacityMaskContent https://docs.microsoft.com/en-us/windows/win32/api/d2d1/ne-d2d1-d2d1_opacity_mask_content
type OpacityMaskContent uint32

// Possible values for OpacityMaskContent.
const (
	OpacityMaskContentGraphics OpacityMaskContent = iota
	OpacityMaskContentTextNatural
	OpacityMaskContentTextGDICompatible
)

// MeasuringMode https://docs.microsoft.com/en-us/windows/win32/api/dcommon/ne-dcommon-dwrite_measuring_mode
type MeasuringMode uint32

// Possible values for MeasuringMode.
const (
	MeasuringModeNatural MeasuringMode = iota
	MeasuringModeGDIClassic
	MeasuringModeGDINatural
)
