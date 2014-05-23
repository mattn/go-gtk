package glib

import (
	// #include "error.go.h"
	// #cgo pkg-config: glib-2.0 gobject-2.0
	"C"
	"unsafe"
)

//-----------------------------------------------------------------------
// g_error
//-----------------------------------------------------------------------
type Error struct {
	GError *C.GError
}

func (v *Error) Error() string {
	return v.Message()
}

func (v *Error) Message() string {
	if unsafe.Pointer(v.GError) == nil || unsafe.Pointer(v.GError.message) == nil {
		return ""
	}
	return C.GoString(C.to_charptr(v.GError.message))
}

func ErrorFromNative(err unsafe.Pointer) *Error {
	return &Error{
		C.to_error(err)}
}
