package gdkpixbuf

/*
#include <gdk-pixbuf/gdk-pixbuf.h>
#include <unistd.h>
#include <stdlib.h>
#include <stdint.h>
#include <stdarg.h>
#include <string.h>

//static const gchar* to_gcharptr(const char* s) { return (const gchar*)s; }
static guchar* to_gucharptr(void* s) { return (guchar*)s; }

static void free_string(char* s) { free(s); }

static gchar* to_gcharptr(char* s) { return (gchar*)s; }

//static void free_string(char* s) { free(s); }
*/
// #cgo pkg-config: gdk-pixbuf-2.0
import "C"
import "github.com/agl/go-gtk/glib"
import "unsafe"

func gboolean2bool(b C.gboolean) bool {
	if b != 0 {
		return true
	}
	return false
}

//-----------------------------------------------------------------------
// GdkPixbuf
//-----------------------------------------------------------------------
type GdkPixbuf struct {
	Pixbuf *C.GdkPixbuf
}

func PixbufFromFile(path string) (pixbuf *GdkPixbuf, err **glib.Error) {
	var error *C.GError
	ptr := C.CString(path)
	defer C.free_string(ptr)
	pixbuf = &GdkPixbuf{C.gdk_pixbuf_new_from_file(ptr, &error)}
	if err != nil && error != nil {
		*err = glib.ErrorFromNative(unsafe.Pointer(error))
	}
	return
}

func PixbufFromFileAtSize(path string, imgW int, imgH int) (pixbuf *GdkPixbuf, err **glib.Error) {
	var error *C.GError
	ptr := C.CString(path)
	defer C.free_string(ptr)
	pixbuf = &GdkPixbuf{C.gdk_pixbuf_new_from_file_at_size(ptr, C.int(imgW), C.int(imgH), &error)}
	if err != nil && error != nil {
		*err = glib.ErrorFromNative(unsafe.Pointer(error))
	}
	return
}

func GetGdkPixbufType() int {
	return int(C.gdk_pixbuf_get_type())
}

func GetFileInfo(path string, imgW *int, imgH *int) *GdkPixbufFormat {
	ptr := C.CString(path)
	defer C.free_string(ptr)

	var w, h C.gint
	format := &GdkPixbufFormat{C.gdk_pixbuf_get_file_info(C.to_gcharptr(ptr), &w, &h)}
	*imgW = int(w)
	*imgH = int(h)
	return format
}

//-----------------------------------------------------------------------
// GdkPixbufAnimation
//-----------------------------------------------------------------------
type GdkPixbufAnimation struct {
	PixbufAnimation *C.GdkPixbufAnimation
}

//-----------------------------------------------------------------------
// GdkPixbufFormat
//-----------------------------------------------------------------------
type GdkPixbufFormat struct {
	PixbufFormat *C.GdkPixbufFormat
}

//-----------------------------------------------------------------------
// GdkPixbufLoader
//-----------------------------------------------------------------------
type GdkPixbufLoader struct {
	PixbufLoader *C.GdkPixbufLoader
}

func PixbufLoader() *GdkPixbufLoader {
	return &GdkPixbufLoader{
		C.gdk_pixbuf_loader_new()}
}
func PixbufLoaderWithType(image_type string) (loader *GdkPixbufLoader, err *C.GError) {
	var error *C.GError
	ptr := C.CString(image_type)
	defer C.free_string(ptr)
	loader = &GdkPixbufLoader{
		C.gdk_pixbuf_loader_new_with_type(ptr, &error)}
	err = error
	return
}
func PixbufLoaderWithMimeType(mime_type string) (loader *GdkPixbufLoader, err *C.GError) {
	var error *C.GError
	ptr := C.CString(mime_type)
	defer C.free_string(ptr)
	loader = &GdkPixbufLoader{
		C.gdk_pixbuf_loader_new_with_mime_type(ptr, &error)}
	err = error
	return
}
func (v GdkPixbufLoader) GetPixbuf() *GdkPixbuf {
	return &GdkPixbuf{
		C.gdk_pixbuf_loader_get_pixbuf(v.PixbufLoader)}
}
func (v GdkPixbufLoader) Write(buf []byte) (ret bool, err *C.GError) {
	var error *C.GError
	var pbuf *byte
	pbuf = &buf[0]
	ret = gboolean2bool(C.gdk_pixbuf_loader_write(v.PixbufLoader, C.to_gucharptr(unsafe.Pointer(pbuf)), C.gsize(len(buf)), &error))
	err = error
	return
}
func (v GdkPixbufLoader) Close() (ret bool, err *C.GError) {
	var error *C.GError
	ret = gboolean2bool(C.gdk_pixbuf_loader_close(v.PixbufLoader, &error))
	err = error
	return
}

//func (v GdkPixbufLoader) GetPixbufAnimation() *GdkPixbufAnimation {
//	return &GdkPixbufAnimation {
//		C.gdk_pixbuf_loader_get_animation(v.PixbufLoader) };
//}
func (v GdkPixbufLoader) SetSize(width int, height int) {
	C.gdk_pixbuf_loader_set_size(v.PixbufLoader, C.int(width), C.int(height))
}
func (v GdkPixbufLoader) GetFormat() *GdkPixbufFormat {
	return &GdkPixbufFormat{
		C.gdk_pixbuf_loader_get_format(v.PixbufLoader)}
}

// FINISH
