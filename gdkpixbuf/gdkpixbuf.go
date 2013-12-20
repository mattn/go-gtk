package gdkpixbuf

// #include "gdkpixbuf.go.h"
// #cgo pkg-config: gdk-pixbuf-2.0
import "C"
import "github.com/mattn/go-gtk/glib"
import (
	"log"
	"runtime"
	"unsafe"
)

func gstring(s *C.char) *C.gchar { return C.toGstr(s) }
func cstring(s *C.gchar) *C.char { return C.toCstr(s) }
func gostring(s *C.gchar) string { return C.GoString(cstring(s)) }

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

func cfree(s *C.char) { C.freeCstr(s) }

func argumentPanic(message string) {
	if pc, _, _, ok := runtime.Caller(2); ok {
		log.Panicf("Arguments error: %s : %s\n",
			runtime.FuncForPC(pc).Name(), message)
	} else {
		log.Panicln("Arguments error: (unknown caller, see stack):", message)
	}
}

//-----------------------------------------------------------------------
// Pixbuf
//-----------------------------------------------------------------------
type Pixbuf struct {
	GPixbuf *C.GdkPixbuf
	*glib.GObject
}

// File Loading

func NewFromFile(filename string) (*Pixbuf, *glib.Error) {
	var error *C.GError
	ptr := C.CString(filename)
	defer cfree(ptr)
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
	defer cfree(ptr)
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
	defer cfree(ptr)
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
	defer cfree(ptr)

	var w, h C.gint
	format := &Format{C.gdk_pixbuf_get_file_info(gstring(ptr), &w, &h)}
	*width = int(w)
	*height = int(h)
	return format
}

// Scaling

type InterpType int

const (
	INTERP_NEAREST InterpType = iota
	INTERP_TILES
	INTERP_BILINEAR
	INTERP_HYPER
)

func ScaleSimple(p *Pixbuf, width, height int, interp InterpType) *Pixbuf {
	gpixbuf := C.gdk_pixbuf_scale_simple(p.GPixbuf, C.int(width), C.int(height), C.GdkInterpType(interp))
	return &Pixbuf{
		GPixbuf: gpixbuf,
		GObject: glib.ObjectFromNative(unsafe.Pointer(gpixbuf)),
	}
}

func Scale(p *Pixbuf, x, y, width, height int, offsetX, offsetY, scaleX, scaleY float64, interp InterpType) *Pixbuf {
	var gpixbuf *C.GdkPixbuf
	C.gdk_pixbuf_scale(
		p.GPixbuf,
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

// gdk_pixbuf_composite_color_simple
// gdk_pixbuf_composite
// gdk_pixbuf_composite_color

type PixbufRotation int

const (
	PIXBUF_ROTATE_NONE             PixbufRotation = 0
	PIXBUF_ROTATE_COUNTERCLOCKWISE PixbufRotation = 90
	PIXBUF_ROTATE_UPSIDEDOWN       PixbufRotation = 180
	PIXBUF_ROTATE_CLOCKWISE        PixbufRotation = 270
)

func RotateSimple(p *Pixbuf, angle PixbufRotation) *Pixbuf {
	gpixbuf := C.gdk_pixbuf_rotate_simple(p.GPixbuf, C.GdkPixbufRotation(angle))
	return &Pixbuf{
		GPixbuf: gpixbuf,
		GObject: glib.ObjectFromNative(unsafe.Pointer(gpixbuf)),
	}
}

func Flip(p *Pixbuf, horizontal bool) *Pixbuf {
	gpixbuf := C.gdk_pixbuf_flip(p.GPixbuf, gbool(horizontal))
	return &Pixbuf{
		GPixbuf: gpixbuf,
		GObject: glib.ObjectFromNative(unsafe.Pointer(gpixbuf)),
	}
}

// The GdkPixbuf Structure

type Colorspace int

const (
	GDK_COLORSPACE_RGB Colorspace = iota
)

type PixbufAlphaMode int

const (
	GDK_PIXBUF_ALPHA_BILEVEL PixbufAlphaMode = iota
	GDK_PIXBUF_ALPHA_FULL
)

func (p *Pixbuf) GetColorspace() Colorspace {
	return Colorspace(C.gdk_pixbuf_get_colorspace(p.GPixbuf))
}

func (p *Pixbuf) GetNChannels() int {
	return int(C.gdk_pixbuf_get_n_channels(p.GPixbuf))
}

func (p *Pixbuf) GetHasAlpha() bool {
	return gobool(C.gdk_pixbuf_get_has_alpha(p.GPixbuf))
}

func (p *Pixbuf) GetBitsPerSample() int {
	return int(C.gdk_pixbuf_get_bits_per_sample(p.GPixbuf))
}

// gdk_pixbuf_get_pixels
// gdk_pixbuf_get_pixels_with_length

func (p *Pixbuf) GetWidth() int {
	return int(C.gdk_pixbuf_get_width(p.GPixbuf))
}

func (p *Pixbuf) GetHeight() int {
	return int(C.gdk_pixbuf_get_height(p.GPixbuf))
}

func (p *Pixbuf) GetRowstride() int {
	return int(C.gdk_pixbuf_get_rowstride(p.GPixbuf))
}

// gdk_pixbuf_get_byte_length
// gdk_pixbuf_get_option

// File saving

func (p *Pixbuf) Save(filename, savetype string, options ...string) *glib.Error {
	if len(options)%2 != 0 {
		argumentPanic("Save options must be even (key and value)")
	}

	pfilename := C.CString(filename)
	defer cfree(pfilename)
	psavetype := C.CString(savetype)
	defer cfree(psavetype)

	klen := len(options) / 2
	keys := C.makeCstrv(C.int(klen + 1))
	vals := C.makeCstrv(C.int(klen + 1))
	for i := 0; i < klen; i++ {
		C.setCstr(keys, C.int(i), C.CString(options[2*i]))
		C.setCstr(vals, C.int(i), C.CString(options[2*i+1]))
	}
	C.setCstr(keys, C.int(klen), nil)
	C.setCstr(vals, C.int(klen), nil)
	defer func() {
		for i := 0; i < klen; i++ {
			cfree(C.getCstr(keys, C.int(i)))
			cfree(C.getCstr(vals, C.int(i)))
		}
		C.freeCstrv(keys)
		C.freeCstrv(vals)
	}()

	var err *C.GError
	C.gdk_pixbuf_savev(p.GPixbuf, pfilename, psavetype, keys, vals, &err)
	if err != nil {
		return glib.ErrorFromNative(unsafe.Pointer(err))
	}
	return nil
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

//  gdk_pixbuf_get_formats

func (v *Format) GetName() string {
	return gostring(C.gdk_pixbuf_format_get_name(v.GPixbufFormat))
}

func (v *Format) GetDescription() string {
	return gostring(C.gdk_pixbuf_format_get_description(v.GPixbufFormat))
}

func (v *Format) GetMimeTypes() []string {
	gstrv := C.gdk_pixbuf_format_get_mime_types(v.GPixbufFormat)
	defer C.g_strfreev(gstrv)
	s := make([]string, 0)
	for i := 0; C.getGstr(gstrv, C.int(i)) != nil; i++ {
		s = append(s, gostring(C.getGstr(gstrv, C.int(i))))
	}
	return s
}

func (v *Format) GetExtensions() []string {
	gstrv := C.gdk_pixbuf_format_get_extensions(v.GPixbufFormat)
	defer C.g_strfreev(gstrv)
	s := make([]string, 0)
	for i := 0; C.getGstr(gstrv, C.int(i)) != nil; i++ {
		s = append(s, gostring(C.getGstr(gstrv, C.int(i))))
	}
	return s
}

func (v *Format) IsWritable() bool {
	return gobool(C.gdk_pixbuf_format_is_writable(v.GPixbufFormat))
}

func (v *Format) IsScalable() bool {
	return gobool(C.gdk_pixbuf_format_is_scalable(v.GPixbufFormat))
}

func (v *Format) IsDisabled() bool {
	return gobool(C.gdk_pixbuf_format_is_disabled(v.GPixbufFormat))
}

func (v *Format) SetDisabled(disabled bool) {
	C.gdk_pixbuf_format_set_disabled(v.GPixbufFormat, gbool(disabled))
}

func (v *Format) GetLicense() string {
	return gostring(C.gdk_pixbuf_format_get_license(v.GPixbufFormat))
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
	defer cfree(ptr)
	loader = &Loader{
		C.gdk_pixbuf_loader_new_with_type(ptr, &error)}
	err = error
	return
}
func NewLoaderWithMimeType(mime_type string) (loader *Loader, err *C.GError) {
	var error *C.GError
	ptr := C.CString(mime_type)
	defer cfree(ptr)
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
