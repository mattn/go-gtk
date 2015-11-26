// +build !cgocheck

package gtkgl

// #include "gtkgl.go.h"
// #cgo pkg-config: gtkgl-2.0
import "C"

import "github.com/mattn/go-gtk/gtk"
import "unsafe"

type GLArea struct {
	gtk.DrawingArea
}

func NewGLArea(attrList []int) *GLArea {
	return &GLArea{gtk.DrawingArea{
		*gtk.WidgetFromNative(C.make_area(C.int(len(attrList)), (*C.int)(unsafe.Pointer(&attrList[0]))))}}
}

func (v *GLArea) MakeCurrent() {
	C.gtk_gl_area_make_current((*C.GtkGLArea)(unsafe.Pointer(v.GWidget)))
}

func (v *GLArea) SwapBuffers() {
	C.gtk_gl_area_swap_buffers((*C.GtkGLArea)(unsafe.Pointer(v.GWidget)))
}
