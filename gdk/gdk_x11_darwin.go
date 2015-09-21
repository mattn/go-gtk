// +build with-x11

package gdk

/*
#cgo pkg-config: gdk-2.0 gthread-2.0

#include <gdk/gdk.h>
#include <gdk/gdkx.h>
*/
import "C"
import "unsafe"

func (v *Window) GetNativeWindowID() int32 {
	return int32(C.gdk_x11_drawable_get_xid((*C.GdkDrawable)(unsafe.Pointer(v.GWindow))))
}
