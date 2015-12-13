// +build !with-x11

package gdk

/*
#cgo pkg-config: gdk-2.0 gthread-2.0
#cgo CFLAGS: -x objective-c

#include <gdk/gdk.h>
#include <gdk/gdkquartz.h>

// Must return void* to avoid "struct size calculation error off=8 bytesize=0"
// See:
// - https://github.com/golang/go/issues/12065
// - http://thread0.me/2015/07/gogoa-cocoa-bindings-for-go/
void* getNsWindow(GdkWindow *w) {
	return (void*) gdk_quartz_window_get_nswindow(w);
}

*/
import "C"

import "unsafe"

type NSWindow struct {
	ID unsafe.Pointer
}

func (v *Window) GetNativeWindow() *NSWindow {
	return &NSWindow{
		unsafe.Pointer(C.getNsWindow(v.GWindow)),
	}
}
