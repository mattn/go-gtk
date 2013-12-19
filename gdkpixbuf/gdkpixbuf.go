package gdkpixbuf

// #include "gdkpixbuf.go.h"
// #cgo pkg-config: gdk-pixbuf-2.0
import "C"
import "github.com/mattn/go-gtk/glib"
import "unsafe"

func gbool(b bool) C.gboolean {
	if b {
		return C.gboolean(1)
	}
	return C.gboolean(0)
}

func gobool(b C.gboolean) bool {
	if b != 0 {
		return true
	}
	return false
}

//-----------------------------------------------------------------------
// Pixbuf
//-----------------------------------------------------------------------
type InterpType int

const (
	INTERP_NEAREST InterpType = iota
	INTERP_TILES
	INTERP_BILINEAR
	INTERP_HYPER
)

type Pixbuf struct {
	GPixbuf *C.GdkPixbuf
	*glib.GObject
}

func NewFromFile(filename string) (*Pixbuf, *glib.Error) {
	var error *C.GError
	ptr := C.CString(filename)
	defer C.free_string(ptr)
	gpixbuf := C.gdk_pixbuf_new_from_file(ptr, &error)
	if error != nil {
		err := glib.ErrorFromNative(unsafe.Pointer(error))
		return nil, err
	}
	pixbuf := &Pixbuf{
		GPixbuf: gpixbuf,
		GObject: glib.ObjectFromNative(unsafe.Pointer(gpixbuf)),
	}
	return pixbuf, nil
}

func NewFromFileAtSize(filename string, width, heigth int) (*Pixbuf, *glib.Error) {
	var error *C.GError
	ptr := C.CString(filename)
	defer C.free_string(ptr)
	gpixbuf := C.gdk_pixbuf_new_from_file_at_size(ptr, C.int(width), C.int(heigth), &error)
	if error != nil {
		err := glib.ErrorFromNative(unsafe.Pointer(error))
		return nil, err
	}
	pixbuf := &Pixbuf{
		GPixbuf: gpixbuf,
		GObject: glib.ObjectFromNative(unsafe.Pointer(gpixbuf)),
	}
	return pixbuf, nil
}

func NewFromFileAtScale(filename string, width, height int, preserve_aspect_ratio bool) (*Pixbuf, *glib.Error) {
	var error *C.GError
	ptr := C.CString(filename)
	defer C.free_string(ptr)
	gpixbuf := C.gdk_pixbuf_new_from_file_at_scale(ptr, C.int(width), C.int(height), gbool(preserve_aspect_ratio), &error)
	if error != nil {
		err := glib.ErrorFromNative(unsafe.Pointer(error))
		return nil, err
	}
	pixbuf := &Pixbuf{
		GPixbuf: gpixbuf,
		GObject: glib.ObjectFromNative(unsafe.Pointer(gpixbuf)),
	}
	return pixbuf, nil
}

func GetType() int {
	return int(C.gdk_pixbuf_get_type())
}

func GetFileInfo(filename string, width, height *int) *Format {
	ptr := C.CString(filename)
	defer C.free_string(ptr)

	var w, h C.gint
	format := &Format{C.gdk_pixbuf_get_file_info(C.to_gcharptr(ptr), &w, &h)}
	*width = int(w)
	*height = int(h)
	return format
}

func Scale(p *Pixbuf, x, y, width, height int, offsetX, offsetY, scaleX, scaleY float64, interp InterpType) *Pixbuf {
	var gpixbuf *C.GdkPixbuf
	C.gdk_pixbuf_scale(
		C.toGdkPixbuf(unsafe.Pointer(p.GPixbuf)),
		gpixbuf,
		C.int(x), C.int(y),
		C.int(width), C.int(height),
		C.double(offsetX), C.double(offsetY),
		C.double(scaleX), C.double(scaleY),
		C.GdkInterpType(interp))
	return &Pixbuf{
		GPixbuf: gpixbuf,
		GObject: glib.ObjectFromNative(unsafe.Pointer(gpixbuf)),
	}
}

func ScaleSimple(p *Pixbuf, width, height int, interp InterpType) *Pixbuf {
	gpixbuf := C.gdk_pixbuf_scale_simple(C.toGdkPixbuf(unsafe.Pointer(p.GPixbuf)), C.int(width), C.int(height), C.GdkInterpType(interp))
	return &Pixbuf{
		GPixbuf: gpixbuf,
		GObject: glib.ObjectFromNative(unsafe.Pointer(gpixbuf)),
	}
}

type PixbufRotation int

const (
	PIXBUF_ROTATE_NONE             PixbufRotation = 0
	PIXBUF_ROTATE_COUNTERCLOCKWISE PixbufRotation = 90
	PIXBUF_ROTATE_UPSIDEDOWN       PixbufRotation = 180
	PIXBUF_ROTATE_CLOCKWISE        PixbufRotation = 270
)

func RotateSimple(p *Pixbuf, angle PixbufRotation) *Pixbuf {
	gpixbuf := C.gdk_pixbuf_rotate_simple(C.toGdkPixbuf(unsafe.Pointer(p.GPixbuf)), C.GdkPixbufRotation(angle))
	return &Pixbuf{
		GPixbuf: gpixbuf,
		GObject: glib.ObjectFromNative(unsafe.Pointer(gpixbuf)),
	}
}

func Flip(p *Pixbuf, horizontal bool) *Pixbuf {
	gpixbuf := C.gdk_pixbuf_flip(C.toGdkPixbuf(unsafe.Pointer(p.GPixbuf)), gbool(horizontal))
	return &Pixbuf{
		GPixbuf: gpixbuf,
		GObject: glib.ObjectFromNative(unsafe.Pointer(gpixbuf)),
	}
}

func (p *Pixbuf) GetWidth() int {
	return int(C.gdk_pixbuf_get_width(C.toGdkPixbuf(unsafe.Pointer(p.GPixbuf))))
}

func (p *Pixbuf) GetHeight() int {
	return int(C.gdk_pixbuf_get_height(C.toGdkPixbuf(unsafe.Pointer(p.GPixbuf))))
}

//-----------------------------------------------------------------------
// Animation
//-----------------------------------------------------------------------
type Animation struct {
	GPixbufAnimation *C.GdkPixbufAnimation
}

//-----------------------------------------------------------------------
// Format
//-----------------------------------------------------------------------
type Format struct {
	GPixbufFormat *C.GdkPixbufFormat
}

//-----------------------------------------------------------------------
// Loader
//-----------------------------------------------------------------------
type Loader struct {
	GPixbufLoader *C.GdkPixbufLoader
}

func NewLoader() *Loader {
	return &Loader{
		C.gdk_pixbuf_loader_new()}
}
func NewLoaderWithType(image_type string) (loader *Loader, err *C.GError) {
	var error *C.GError
	ptr := C.CString(image_type)
	defer C.free_string(ptr)
	loader = &Loader{
		C.gdk_pixbuf_loader_new_with_type(ptr, &error)}
	err = error
	return
}
func NewLoaderWithMimeType(mime_type string) (loader *Loader, err *C.GError) {
	var error *C.GError
	ptr := C.CString(mime_type)
	defer C.free_string(ptr)
	loader = &Loader{
		C.gdk_pixbuf_loader_new_with_mime_type(ptr, &error)}
	err = error
	return
}
func (v Loader) GetPixbuf() *Pixbuf {
	gpixbuf := C.gdk_pixbuf_loader_get_pixbuf(v.GPixbufLoader)
	return &Pixbuf{
		GPixbuf: gpixbuf,
		GObject: glib.ObjectFromNative(unsafe.Pointer(gpixbuf)),
	}
}
func (v Loader) Write(buf []byte) (ret bool, err *C.GError) {
	var error *C.GError
	var pbuf *byte
	pbuf = &buf[0]
	ret = gobool(C.gdk_pixbuf_loader_write(v.GPixbufLoader, C.to_gucharptr(unsafe.Pointer(pbuf)), C.gsize(len(buf)), &error))
	err = error
	return
}
func (v Loader) Close() (ret bool, err *C.GError) {
	var error *C.GError
	ret = gobool(C.gdk_pixbuf_loader_close(v.GPixbufLoader, &error))
	err = error
	return
}

//func (v Loader) GetPixbufAnimation() *Animation {
//	return &Animation {
//		C.gdk_pixbuf_loader_get_animation(v.GPixbufLoader) };
//}
func (v Loader) SetSize(width int, height int) {
	C.gdk_pixbuf_loader_set_size(v.GPixbufLoader, C.int(width), C.int(height))
}
func (v Loader) GetFormat() *Format {
	return &Format{
		C.gdk_pixbuf_loader_get_format(v.GPixbufLoader)}
}

// FINISH
