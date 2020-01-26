package d2d

import (
	"syscall"
	"unsafe"

	"github.com/richardwilkes/win32"
	"github.com/richardwilkes/win32/com"
)

var (
	d2d1          = syscall.NewLazyDLL("d2d1.dll")
	createFactory = d2d1.NewProc("D2D1CreateFactory")
	// FactoryGUID holds the GUID for the Factory class.
	FactoryGUID   = com.GUID{Data1: 0x06152247, Data2: 0x6f50, Data3: 0x465a, Data4: [8]byte{0x92, 0x45, 0x11, 0x8b, 0xfd, 0x3b, 0x60, 0x07}}
)

// Factory https://docs.microsoft.com/en-us/windows/win32/api/d2d1/nn-d2d1-id2d1factory
type Factory struct {
	com.IUnknown
}

type vmtFactory struct {
	com.VMTIUnknown
	ReloadSystemMetrics            uintptr
	GetDesktopDPI                  uintptr
	CreateRectangleGeometry        uintptr
	CreateRoundedRectangleGeometry uintptr
	CreateEllipseGeometry          uintptr
	CreateGeometryGroup            uintptr
	CreateTransformedGeometry      uintptr
	CreatePathGeometry             uintptr
	CreateStrokeStyle              uintptr
	CreateDrawingStateBlock        uintptr
	CreateWicBitmapRenderTarget    uintptr
	CreateHwndRenderTarget         uintptr
	CreateDxgiSurfaceRenderTarget  uintptr
	CreateDCRenderTarget           uintptr
}

// CreateFactory https://docs.microsoft.com/en-us/windows/win32/api/d2d1/nf-d2d1-d2d1createfactory
func CreateFactory(multiThreaded bool, debugLevel DebugLevel) *Factory {
	var factory *Factory
	var factoryType uint32
	if multiThreaded {
		factoryType = 1
	}
	if ret, _, _ := createFactory.Call(uintptr(factoryType), uintptr(unsafe.Pointer(&FactoryGUID)), uintptr(unsafe.Pointer(&debugLevel)), uintptr(unsafe.Pointer(&factory))); ret != win32.S_OK {
		return nil
	}
	return factory
}

func (f *Factory) vmt() *vmtFactory {
	return (*vmtFactory)(f.UnsafeVirtualMethodTable)
}

// ReloadSystemMetrics https://docs.microsoft.com/en-us/windows/win32/api/d2d1/nf-d2d1-id2d1factory-reloadsystemmetrics
func (f *Factory) ReloadSystemMetrics() {
	syscall.Syscall(f.vmt().ReloadSystemMetrics, 1, uintptr(unsafe.Pointer(f)), 0, 0)
}

// DesktopDPI https://docs.microsoft.com/en-us/windows/win32/api/d2d1/nf-d2d1-id2d1factory-getdesktopdpi
func (f *Factory) DesktopDPI() (dpiX, dpiY float32) {
	syscall.Syscall(f.vmt().GetDesktopDPI, 3, uintptr(unsafe.Pointer(f)), uintptr(unsafe.Pointer(&dpiX)), uintptr(unsafe.Pointer(&dpiY)))
	return
}

// CreateRectangleGeometry https://docs.microsoft.com/en-us/windows/win32/api/d2d1/nf-d2d1-id2d1factory-createrectanglegeometry%28constd2d1_rect_f_id2d1rectanglegeometry%29
func (f *Factory) CreateRectangleGeometry(rectangle *Rect) *RectangleGeometry {
	var rectangleGeometry *RectangleGeometry
	if ret, _, _ := syscall.Syscall(f.vmt().CreateRectangleGeometry, 3, uintptr(unsafe.Pointer(f)), uintptr(unsafe.Pointer(rectangle)), uintptr(unsafe.Pointer(&rectangleGeometry))); ret != win32.S_OK {
		return nil
	}
	return rectangleGeometry
}

// CreateRoundedRectangleGeometry https://docs.microsoft.com/en-us/windows/win32/api/d2d1/nf-d2d1-id2d1factory-createroundedrectanglegeometry(constd2d1_rounded_rect__id2d1roundedrectanglegeometry)
func (f *Factory) CreateRoundedRectangleGeometry(roundedRectangle *RoundedRect) *RoundedRectangleGeometry {
	var roundedRectangleGeometry *RoundedRectangleGeometry
	if ret, _, _ := syscall.Syscall(f.vmt().CreateRoundedRectangleGeometry, 3, uintptr(unsafe.Pointer(f)), uintptr(unsafe.Pointer(roundedRectangle)), uintptr(unsafe.Pointer(&roundedRectangleGeometry))); ret != win32.S_OK {
		return nil
	}
	return roundedRectangleGeometry
}

// CreateEllipseGeometry https://docs.microsoft.com/en-us/windows/win32/api/d2d1/nf-d2d1-id2d1factory-createellipsegeometry(constd2d1_ellipse_id2d1ellipsegeometry)
func (f *Factory) CreateEllipseGeometry(ellipse *Ellipse) *EllipseGeometry {
	var ellipseGeometry *EllipseGeometry
	if ret, _, _ := syscall.Syscall(f.vmt().CreateEllipseGeometry, 3, uintptr(unsafe.Pointer(f)), uintptr(unsafe.Pointer(ellipse)), uintptr(unsafe.Pointer(&ellipseGeometry))); ret != win32.S_OK {
		return nil
	}
	return ellipseGeometry
}

// CreateGeometryGroup https://docs.microsoft.com/en-us/windows/win32/api/d2d1/nf-d2d1-id2d1factory-creategeometrygroup
func (f *Factory) CreateGeometryGroup(winding bool, geometries []*Geometry) *GeometryGroup {
	var geometryGroup *GeometryGroup
	var fillMode uint32
	if winding {
		fillMode = 1
	}
	if ret, _, _ := syscall.Syscall6(f.vmt().CreateGeometryGroup, 5, uintptr(unsafe.Pointer(f)), uintptr(fillMode), uintptr(unsafe.Pointer(&(geometries[0]))), uintptr(len(geometries)), uintptr(unsafe.Pointer(&geometryGroup)), 0); ret != win32.S_OK {
		return nil
	}
	return geometryGroup
}

// CreateTransformedGeometry https://docs.microsoft.com/en-us/windows/win32/api/d2d1/nf-d2d1-id2d1factory-createtransformedgeometry(id2d1geometry_constd2d1_matrix_3x2_f__id2d1transformedgeometry)
func (f *Factory) CreateTransformedGeometry(sourceGeometry *Geometry, transform *Matrix3x2) *TransformedGeometry {
	var transformedGeometry *TransformedGeometry
	if ret, _, _ := syscall.Syscall6(f.vmt().CreateTransformedGeometry, 4, uintptr(unsafe.Pointer(f)), uintptr(unsafe.Pointer(sourceGeometry)), uintptr(unsafe.Pointer(transform)), uintptr(unsafe.Pointer(&transformedGeometry)), 0, 0); ret != win32.S_OK {
		return nil
	}
	return transformedGeometry
}

// CreatePathGeometry https://docs.microsoft.com/en-us/windows/win32/api/d2d1/nf-d2d1-id2d1factory-createpathgeometry
func (f *Factory) CreatePathGeometry() *PathGeometry {
	var pathGeometry *PathGeometry
	if ret, _, _ := syscall.Syscall(f.vmt().CreatePathGeometry, 2, uintptr(unsafe.Pointer(f)), uintptr(unsafe.Pointer(&pathGeometry)), 0); ret != win32.S_OK {
		return nil
	}
	return pathGeometry
}

// CreateStrokeStyle https://docs.microsoft.com/en-us/windows/win32/api/d2d1/nf-d2d1-id2d1factory-createstrokestyle(constd2d1_stroke_style_properties__constfloat_uint32_id2d1strokestyle)
func (f *Factory) CreateStrokeStyle(strokeStyleProperties *StrokeStyleProperties, dashes []float32) *StrokeStyle {
	var strokeStyle *StrokeStyle
	if ret, _, _ := syscall.Syscall6(f.vmt().CreateStrokeStyle, 5, uintptr(unsafe.Pointer(f)), uintptr(unsafe.Pointer(strokeStyleProperties)), uintptr(unsafe.Pointer(&(dashes[0]))), uintptr(len(dashes)), uintptr(unsafe.Pointer(&strokeStyle)), 0); ret != win32.S_OK {
		return nil
	}
	return strokeStyle
}

// CreateDrawingStateBlock https://docs.microsoft.com/en-us/windows/win32/api/d2d1/nf-d2d1-id2d1factory-createdrawingstateblock(id2d1drawingstateblock)
func (f *Factory) CreateDrawingStateBlock(drawingStateDescription *DrawingStateDescription, textRenderingParams *TextRenderingParams) *DrawingStateBlock {
	var drawingStateBlock *DrawingStateBlock
	if ret, _, _ := syscall.Syscall6(f.vmt().CreateDrawingStateBlock, 4, uintptr(unsafe.Pointer(f)), uintptr(unsafe.Pointer(drawingStateDescription)), uintptr(unsafe.Pointer(textRenderingParams)), uintptr(unsafe.Pointer(&drawingStateBlock)), 0, 0); ret != win32.S_OK {
		return nil
	}
	return drawingStateBlock
}

// CreateWicBitmapRenderTarget https://docs.microsoft.com/en-us/windows/win32/api/d2d1/nf-d2d1-id2d1factory-createwicbitmaprendertarget(iwicbitmap_constd2d1_render_target_properties__id2d1rendertarget)
func (f *Factory) CreateWicBitmapRenderTarget(target *WicBitmap, renderTargetProperties *RenderTargetProperties) *RenderTarget {
	var renderTarget *RenderTarget
	if ret, _, _ := syscall.Syscall6(f.vmt().CreateWicBitmapRenderTarget, 4, uintptr(unsafe.Pointer(f)), uintptr(unsafe.Pointer(target)), uintptr(unsafe.Pointer(renderTargetProperties)), uintptr(unsafe.Pointer(&renderTarget)), 0, 0); ret != win32.S_OK {
		return nil
	}
	return renderTarget
}

// CreateHwndRenderTarget https://docs.microsoft.com/en-us/windows/win32/api/d2d1/nf-d2d1-id2d1factory-createhwndrendertarget(constd2d1_render_target_properties_constd2d1_hwnd_render_target_properties_id2d1hwndrendertarget)
func (f *Factory) CreateHwndRenderTarget(renderTargetProperties *RenderTargetProperties, hwndRenderTargetProperties *HWNDRenderTargetProperties) *HWNDRenderTarget {
	var hwndRenderTarget *HWNDRenderTarget
	if ret, _, _ := syscall.Syscall6(f.vmt().CreateHwndRenderTarget, 4, uintptr(unsafe.Pointer(f)), uintptr(unsafe.Pointer(renderTargetProperties)), uintptr(unsafe.Pointer(hwndRenderTargetProperties)), uintptr(unsafe.Pointer(&hwndRenderTarget)), 0, 0); ret != win32.S_OK {
		return nil
	}
	return hwndRenderTarget
}

// CreateDxgiSurfaceRenderTarget https://docs.microsoft.com/en-us/windows/win32/api/d2d1/nf-d2d1-id2d1factory-createdxgisurfacerendertarget(idxgisurface_constd2d1_render_target_properties_id2d1rendertarget)
func (f *Factory) CreateDxgiSurfaceRenderTarget(dxgiSurface *Surface, renderTargetProperties *RenderTargetProperties) *RenderTarget {
	var renderTarget *RenderTarget
	if ret, _, _ := syscall.Syscall6(f.vmt().CreateDxgiSurfaceRenderTarget, 4, uintptr(unsafe.Pointer(f)), uintptr(unsafe.Pointer(dxgiSurface)), uintptr(unsafe.Pointer(renderTargetProperties)), uintptr(unsafe.Pointer(&renderTarget)), 0, 0); ret != win32.S_OK {
		return nil
	}
	return renderTarget
}

// CreateDCRenderTarget https://docs.microsoft.com/en-us/windows/win32/api/d2d1/nf-d2d1-id2d1factory-createdcrendertarget
func (f *Factory) CreateDCRenderTarget(renderTargetProperties *RenderTargetProperties) *DCRenderTarget {
	var dcRenderTarget *DCRenderTarget
	if ret, _, _ := syscall.Syscall(f.vmt().CreateDCRenderTarget, 3, uintptr(unsafe.Pointer(f)), uintptr(unsafe.Pointer(renderTargetProperties)), uintptr(unsafe.Pointer(&dcRenderTarget))); ret != win32.S_OK {
		return nil
	}
	return dcRenderTarget
}
