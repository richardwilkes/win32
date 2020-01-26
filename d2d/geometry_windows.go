package d2d

import (
	"fmt"
	"syscall"
	"unsafe"

	"github.com/richardwilkes/win32"
	"github.com/richardwilkes/win32/com"
)

// GeometryGUID holds the GUID for the Geometry class.
var GeometryGUID = com.GUID{Data1: 0x2cd906a1, Data2: 0x12e2, Data3: 0x11dc, Data4: [8]byte{0x9f, 0xed, 0x00, 0x11, 0x43, 0xa0, 0x55, 0xf9}}

// Geometry https://docs.microsoft.com/en-us/windows/win32/api/d2d1/nn-d2d1-id2d1geometry
type Geometry struct {
	Resource
}

type vmtGeometry struct {
	vmtResource
	GetBounds            uintptr
	GetWidenedBounds     uintptr
	StrokeContainsPoint  uintptr
	FillContainsPoint    uintptr
	CompareWithGeometry  uintptr
	Simplify             uintptr
	Tessellate           uintptr
	CombineWithGeometry  uintptr
	Outline              uintptr
	ComputeArea          uintptr
	ComputeLength        uintptr
	ComputePointAtLength uintptr
	Widen                uintptr
}

func (obj *Geometry) vmt() *vmtGeometry {
	return (*vmtGeometry)(obj.UnsafeVirtualMethodTable)
}

// Bounds https://docs.microsoft.com/en-us/windows/win32/api/d2d1/nf-d2d1-id2d1geometry-getbounds%28constd2d1_matrix_3x2_f_d2d1_rect_f%29
func (obj *Geometry) Bounds(worldTransform *Matrix3x2) Rect {
	var bounds Rect
	syscall.Syscall(obj.vmt().GetBounds, 3, uintptr(unsafe.Pointer(obj)), uintptr(unsafe.Pointer(worldTransform)), uintptr(unsafe.Pointer(&bounds)))
	return bounds
}

// WidenedBounds https://docs.microsoft.com/en-us/windows/win32/api/d2d1/nf-d2d1-id2d1geometry-getwidenedbounds(float_id2d1strokestyle_constd2d1_matrix_3x2_f__d2d1_rect_f)
func (obj *Geometry) WidenedBounds(strokeWidth float32, strokeStyle *StrokeStyle, worldTransform *Matrix3x2, flatteningTolerance float32) Rect {
	var bounds Rect
	syscall.Syscall6(obj.vmt().GetWidenedBounds, 6, uintptr(unsafe.Pointer(obj)), uintptr(*(*uint32)(unsafe.Pointer(&strokeWidth))), uintptr(unsafe.Pointer(strokeStyle)), uintptr(unsafe.Pointer(worldTransform)), uintptr(*(*uint32)(unsafe.Pointer(&flatteningTolerance))), uintptr(unsafe.Pointer(&bounds)))
	return bounds
}

// StrokeContainsPoint https://docs.microsoft.com/en-us/windows/win32/api/d2d1/nf-d2d1-id2d1geometry-strokecontainspoint(d2d1_point_2f_float_id2d1strokestyle_constd2d1_matrix_3x2_f__float_bool)
func (obj *Geometry) StrokeContainsPoint(point Point, strokeWidth float32, strokeStyle *StrokeStyle, worldTransform *Matrix3x2, flatteningTolerance float32) bool {
	var contains int32
	syscall.Syscall9(obj.vmt().StrokeContainsPoint, 7, uintptr(unsafe.Pointer(obj)), uintptr(*(*uint64)(unsafe.Pointer(&point))), uintptr(*(*uint32)(unsafe.Pointer(&strokeWidth))), uintptr(unsafe.Pointer(strokeStyle)), uintptr(unsafe.Pointer(worldTransform)), uintptr(*(*uint32)(unsafe.Pointer(&flatteningTolerance))), uintptr(unsafe.Pointer(&contains)), 0, 0)
	return contains != 0
}

// FillContainsPoint https://docs.microsoft.com/en-us/windows/win32/api/d2d1/nf-d2d1-id2d1geometry-fillcontainspoint(d2d1_point_2f_constd2d1_matrix_3x2_f__float_bool)
func (obj *Geometry) FillContainsPoint(point Point, worldTransform *Matrix3x2, flatteningTolerance float32) bool {
	var contains int32
	syscall.Syscall6(obj.vmt().FillContainsPoint, 5, uintptr(unsafe.Pointer(obj)), uintptr(*(*uint64)(unsafe.Pointer(&point))), uintptr(unsafe.Pointer(worldTransform)), uintptr(*(*uint32)(unsafe.Pointer(&flatteningTolerance))), uintptr(unsafe.Pointer(&contains)), 0)
	return contains != 0
}

// CompareWithGeometry https://docs.microsoft.com/en-us/windows/win32/api/d2d1/nf-d2d1-id2d1geometry-comparewithgeometry(id2d1geometry_constd2d1_matrix_3x2_f__d2d1_geometry_relation)
func (obj *Geometry) CompareWithGeometry(inputGeometry *Geometry, inputGeometryTransform *Matrix3x2, flatteningTolerance float32) GeometryRelation {
	var relation GeometryRelation
	syscall.Syscall6(obj.vmt().CompareWithGeometry, 5, uintptr(unsafe.Pointer(obj)), uintptr(unsafe.Pointer(inputGeometry)), uintptr(unsafe.Pointer(inputGeometryTransform)), uintptr(*(*uint32)(unsafe.Pointer(&flatteningTolerance))), uintptr(unsafe.Pointer(&relation)), 0)
	return relation
}

// Simplify https://docs.microsoft.com/en-us/windows/win32/api/d2d1/nf-d2d1-id2d1geometry-simplify(d2d1_geometry_simplification_option_constd2d1_matrix_3x2_f_float_id2d1simplifiedgeometrysink)
func (obj *Geometry) Simplify(onlyLines bool, worldTransform *Matrix3x2, flatteningTolerance float32, geometrySink *SimplifiedGeometrySink) {
	var simplificationOption uint32
	if onlyLines {
		simplificationOption = 1
	}
	syscall.Syscall6(obj.vmt().Simplify, 5, uintptr(unsafe.Pointer(obj)), uintptr(simplificationOption), uintptr(unsafe.Pointer(worldTransform)), uintptr(*(*uint32)(unsafe.Pointer(&flatteningTolerance))), uintptr(unsafe.Pointer(geometrySink)), 0)
}

// Tessellate https://docs.microsoft.com/en-us/windows/win32/api/d2d1/nf-d2d1-id2d1geometry-tessellate(constd2d1_matrix_3x2_f_id2d1tessellationsink)
func (obj *Geometry) Tessellate(worldTransform *Matrix3x2, flatteningTolerance float32, tessellationSink *TessellationSink) {
	syscall.Syscall6(obj.vmt().Tessellate, 4, uintptr(unsafe.Pointer(obj)), uintptr(unsafe.Pointer(worldTransform)), uintptr(*(*uint32)(unsafe.Pointer(&flatteningTolerance))), uintptr(unsafe.Pointer(tessellationSink)), 0, 0)
}

// CombineWithGeometry https://docs.microsoft.com/en-us/windows/win32/api/d2d1/nf-d2d1-id2d1geometry-combinewithgeometry(id2d1geometry_d2d1_combine_mode_constd2d1_matrix_3x2_f__id2d1simplifiedgeometrysink)
func (obj *Geometry) CombineWithGeometry(inputGeometry *Geometry, combineMode CombineMode, inputGeometryTransform *Matrix3x2, flatteningTolerance float32, geometrySink *SimplifiedGeometrySink) error {
	if ret, _, _ := syscall.Syscall6(obj.vmt().CombineWithGeometry, 6, uintptr(unsafe.Pointer(obj)), uintptr(unsafe.Pointer(inputGeometry)), uintptr(combineMode), uintptr(unsafe.Pointer(inputGeometryTransform)), uintptr(*(*uint32)(unsafe.Pointer(&flatteningTolerance))), uintptr(unsafe.Pointer(geometrySink))); ret != win32.S_OK {
		return fmt.Errorf("call to CombineWithGeometry failed: %#x", ret)
	}
	return nil
}

// Outline https://docs.microsoft.com/en-us/windows/win32/api/d2d1/nf-d2d1-id2d1geometry-outline%28constd2d1_matrix_3x2_f_float_id2d1simplifiedgeometrysink%29
func (obj *Geometry) Outline(worldTransform *Matrix3x2, flatteningTolerance float32, geometrySink *SimplifiedGeometrySink) {
	syscall.Syscall6(obj.vmt().Outline, 4, uintptr(unsafe.Pointer(obj)), uintptr(unsafe.Pointer(worldTransform)), uintptr(*(*uint32)(unsafe.Pointer(&flatteningTolerance))), uintptr(unsafe.Pointer(geometrySink)), 0, 0)
}

// ComputeArea https://docs.microsoft.com/en-us/windows/win32/api/d2d1/nf-d2d1-id2d1geometry-computearea%28constd2d1_matrix_3x2_f__float%29
func (obj *Geometry) ComputeArea(worldTransform *Matrix3x2, flatteningTolerance float32) float32 {
	var area float32
	syscall.Syscall6(obj.vmt().ComputeArea, 4, uintptr(unsafe.Pointer(obj)), uintptr(unsafe.Pointer(worldTransform)), uintptr(*(*uint32)(unsafe.Pointer(&flatteningTolerance))), uintptr(unsafe.Pointer(&area)), 0, 0)
	return area
}

// ComputeLength https://docs.microsoft.com/en-us/windows/win32/api/d2d1/nf-d2d1-id2d1geometry-computelength%28constd2d1_matrix_3x2_f_float_float%29
func (obj *Geometry) ComputeLength(worldTransform *Matrix3x2, flatteningTolerance float32) float32 {
	var length float32
	syscall.Syscall6(obj.vmt().ComputeLength, 4, uintptr(unsafe.Pointer(obj)), uintptr(unsafe.Pointer(worldTransform)), uintptr(*(*uint32)(unsafe.Pointer(&flatteningTolerance))), uintptr(unsafe.Pointer(&length)), 0, 0)
	return length
}

// ComputePointAtLength https://docs.microsoft.com/en-us/windows/win32/api/d2d1/nf-d2d1-id2d1geometry-computepointatlength%28float_constd2d1_matrix_3x2_f_float_d2d1_point_2f_d2d1_point_2f%29
func (obj *Geometry) ComputePointAtLength(length float32, worldTransform *Matrix3x2, flatteningTolerance float32) (point Point, unitTangentVector Point, err error) {
	if ret, _, _ := syscall.Syscall6(obj.vmt().ComputePointAtLength, 6, uintptr(unsafe.Pointer(obj)), uintptr(*(*uint32)(unsafe.Pointer(&length))), uintptr(unsafe.Pointer(worldTransform)), uintptr(*(*uint32)(unsafe.Pointer(&flatteningTolerance))), uintptr(unsafe.Pointer(&point)), uintptr(unsafe.Pointer(&unitTangentVector))); ret != win32.S_OK {
		err = fmt.Errorf("call to ComputePointAtLength failed: %#x", ret)
	}
	return
}

// Widen https://docs.microsoft.com/en-us/windows/win32/api/d2d1/nf-d2d1-id2d1geometry-widen%28float_id2d1strokestyle_constd2d1_matrix_3x2_f_id2d1simplifiedgeometrysink%29
func (obj *Geometry) Widen(strokeWidth float32, strokeStyle *StrokeStyle, worldTransform *Matrix3x2, flatteningTolerance float32, geometrySink *SimplifiedGeometrySink) {
	syscall.Syscall6(obj.vmt().Widen, 6, uintptr(unsafe.Pointer(obj)), uintptr(*(*uint32)(unsafe.Pointer(&strokeWidth))), uintptr(unsafe.Pointer(strokeStyle)), uintptr(unsafe.Pointer(worldTransform)), uintptr(*(*uint32)(unsafe.Pointer(&flatteningTolerance))), uintptr(unsafe.Pointer(geometrySink)))
}
