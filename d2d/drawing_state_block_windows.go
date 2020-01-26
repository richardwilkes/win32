package d2d

import (
	"syscall"
	"unsafe"

	"github.com/richardwilkes/win32/com"
)

// DrawingStateBlockGUID holds the GUID for the DrawingStateBlock class.
var DrawingStateBlockGUID = com.GUID{Data1: 0x28506e39, Data2: 0xebf6, Data3: 0x46a1, Data4: [8]byte{0xbb, 0x47, 0xfd, 0x85, 0x56, 0x5a, 0xb9, 0x57}}

// DrawingStateBlock https://docs.microsoft.com/en-us/windows/win32/api/d2d1/nn-d2d1-id2d1drawingstateblock
type DrawingStateBlock struct {
	Resource
}

type vmtDrawingStateBlock struct {
	vmtResource
	GetDescription         uintptr
	SetDescription         uintptr
	SetTextRenderingParams uintptr
	GetTextRenderingParams uintptr
}

func (obj *DrawingStateBlock) vmt() *vmtDrawingStateBlock {
	return (*vmtDrawingStateBlock)(obj.UnsafeVirtualMethodTable)
}

// Description https://docs.microsoft.com/en-us/windows/win32/api/d2d1/nf-d2d1-id2d1drawingstateblock-getdescription
func (obj *DrawingStateBlock) Description() *DrawingStateDescription {
	var desc DrawingStateDescription
	syscall.Syscall(obj.vmt().GetDescription, 2, uintptr(unsafe.Pointer(obj)), uintptr(unsafe.Pointer(&desc)), 0)
	return &desc
}

// SetDescription https://docs.microsoft.com/en-us/windows/win32/api/d2d1/nf-d2d1-id2d1drawingstateblock-setdescription(constd2d1_drawing_state_description)
func (obj *DrawingStateBlock) SetDescription(stateDescription *DrawingStateDescription) {
	syscall.Syscall(obj.vmt().SetDescription, 2, uintptr(unsafe.Pointer(obj)), uintptr(unsafe.Pointer(stateDescription)), 0)
}

// TextRenderingParams https://docs.microsoft.com/en-us/windows/win32/api/d2d1/nf-d2d1-id2d1drawingstateblock-gettextrenderingparams
func (obj *DrawingStateBlock) TextRenderingParams() *TextRenderingParams {
	var params *TextRenderingParams
	syscall.Syscall(obj.vmt().GetTextRenderingParams, 2, uintptr(unsafe.Pointer(obj)), uintptr(unsafe.Pointer(&params)), 0)
	return params
}

// SetTextRenderingParams https://docs.microsoft.com/en-us/windows/win32/api/d2d1/nf-d2d1-id2d1drawingstateblock-settextrenderingparams
func (obj *DrawingStateBlock) SetTextRenderingParams(textRenderingParams *TextRenderingParams) {
	syscall.Syscall(obj.vmt().SetTextRenderingParams, 2, uintptr(unsafe.Pointer(obj)), uintptr(unsafe.Pointer(textRenderingParams)), 0)
}
