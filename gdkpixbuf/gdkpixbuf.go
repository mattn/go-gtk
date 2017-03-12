// +build !cgocheck

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

func panic_if_version_older(major int, minor int, micro int, function string) {
	if C._check_version(C.int(major), C.int(minor), C.int(micro)) == 0 {
		log.Panicf("%s is not provided on your Glib, version %d.%d is required\n", function, major, minor)
	}
}

func panic_if_version_older_auto(major, minor, micro int) {
	if C._check_version(C.int(major), C.int(minor), C.int(micro)) != 0 {
		return
	}
	formatStr := "%s is not provided on your Glib, version %d.%d is required\n"
	if pc, _, _, ok := runtime.Caller(1); ok {
		log.Panicf(formatStr, runtime.FuncForPC(pc).Name(), major, minor)
	} else {
		log.Panicf("Glib version %d.%d is required (unknown caller, see stack)\n",
			major, minor)
	}
}

func deprecated_since(major int, minor int, micro int, function string) {
	if C._check_version(C.int(major), C.int(minor), C.int(micro)) != 0 {
		log.Printf("\nWarning: %s is deprecated since glib %d.%d\n", function, major, minor)
	}
}

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
	*GdkPixbuf
	*glib.GObject
}

type GdkPixbuf struct {
	GPixbuf *C.GdkPixbuf
}

func NewGdkPixbuf(p unsafe.Pointer) *GdkPixbuf {
	return &GdkPixbuf{(*C.GdkPixbuf)(p)}
}

// File Loading
// GdkPixbuf * gdk_pixbuf_new (GdkColorspace colorspace, gboolean has_alpha, int bits_per_sample, int width, int height);
func NewPixbuf(colorspace Colorspace, hasAlpha bool, bitsPerSample, width, height int) *Pixbuf {
	gpixbuf := C.gdk_pixbuf_new(
		C.GdkColorspace(colorspace),
		gbool(hasAlpha),
		C.int(bitsPerSample),
		C.int(width),
		C.int(height),
	)

	return &Pixbuf{
		GdkPixbuf: &GdkPixbuf{gpixbuf},
		GObject:   glib.ObjectFromNative(unsafe.Pointer(gpixbuf)),
	}
}

func NewPixbufFromFile(filename string) (*Pixbuf, *glib.Error) {
	var err *C.GError
	ptr := C.CString(filename)
	defer cfree(ptr)
	gpixbuf := C.gdk_pixbuf_new_from_file(ptr, &err)
	if err != nil {
		return nil, glib.ErrorFromNative(unsafe.Pointer(err))
	}
	return &Pixbuf{
		GdkPixbuf: &GdkPixbuf{gpixbuf},
		GObject:   glib.ObjectFromNative(unsafe.Pointer(gpixbuf)),
	}, nil
}

func NewPixbufFromFileAtSize(filename string, width, heigth int) (*Pixbuf, *glib.Error) {
	var err *C.GError
	ptr := C.CString(filename)
	defer cfree(ptr)
	gpixbuf := C.gdk_pixbuf_new_from_file_at_size(ptr, C.int(width), C.int(heigth), &err)
	if err != nil {
		return nil, glib.ErrorFromNative(unsafe.Pointer(err))
	}
	return &Pixbuf{
		GdkPixbuf: &GdkPixbuf{gpixbuf},
		GObject:   glib.ObjectFromNative(unsafe.Pointer(gpixbuf)),
	}, nil
}

func NewPixbufFromFileAtScale(filename string, width, height int, preserve_aspect_ratio bool) (*Pixbuf, *glib.Error) {
	var err *C.GError
	ptr := C.CString(filename)
	defer cfree(ptr)
	gpixbuf := C.gdk_pixbuf_new_from_file_at_scale(ptr, C.int(width), C.int(height), gbool(preserve_aspect_ratio), &err)
	if err != nil {
		return nil, glib.ErrorFromNative(unsafe.Pointer(err))
	}
	return &Pixbuf{
		GdkPixbuf: &GdkPixbuf{gpixbuf},
		GObject:   glib.ObjectFromNative(unsafe.Pointer(gpixbuf)),
	}, nil
}

// NewPixbufFromData creates a Pixbuf from image data in a byte array
//
// Can be used for reading Base64 encoded images easily with the output from base64.StdEncoding.DecodeString("...")
func NewPixbufFromData(buffer []byte) (*Pixbuf, *glib.Error) {
	var err *C.GError
	loader := C.gdk_pixbuf_loader_new()
	C.gdk_pixbuf_loader_write(loader, C.to_gucharptr(unsafe.Pointer(&buffer[0])), C.gsize(len(buffer)), &err)
	gpixbuf := C.gdk_pixbuf_loader_get_pixbuf(loader)

	if err != nil {
		return nil, glib.ErrorFromNative(unsafe.Pointer(err))
	}
	return &Pixbuf{
		GdkPixbuf: &GdkPixbuf{gpixbuf},
		GObject:   glib.ObjectFromNative(unsafe.Pointer(gpixbuf)),
	}, nil
}

// NewPixbufFromBytes creates a Pixbuf from image data in a byte array
//
// Can be used for reading Base64 encoded images easily with the output from base64.StdEncoding.DecodeString("...")
func NewPixbufFromBytes(buffer []byte) (*Pixbuf, *glib.Error) {
	var err *C.GError
	loader := C.gdk_pixbuf_loader_new()
	C.gdk_pixbuf_loader_write(loader, C.to_gucharptr(unsafe.Pointer(&buffer[0])), C.gsize(len(buffer)), &err)
	gpixbuf := C.gdk_pixbuf_loader_get_pixbuf(loader)

	if err != nil {
		return nil, glib.ErrorFromNative(unsafe.Pointer(err))
	}
	return &Pixbuf{
		GdkPixbuf: &GdkPixbuf{gpixbuf},
		GObject:   glib.ObjectFromNative(unsafe.Pointer(gpixbuf)),
	}, nil
}

func NewPixbufFromXpmData(data **byte) (*Pixbuf, *glib.Error) {
	var err *C.GError
	gpixbuf := C.gdk_pixbuf_new_from_xpm_data(
		(**C.char)(unsafe.Pointer(data)),
	)
	if err != nil {
		return nil, glib.ErrorFromNative(unsafe.Pointer(err))
	}
	return &Pixbuf{
		GdkPixbuf: &GdkPixbuf{gpixbuf},
		GObject:   glib.ObjectFromNative(unsafe.Pointer(gpixbuf)),
	}, nil
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

func (p *Pixbuf) ScaleSimple(width, height int, interp InterpType) *Pixbuf {
	gpixbuf := C.gdk_pixbuf_scale_simple(p.GPixbuf, C.int(width), C.int(height), C.GdkInterpType(interp))
	return &Pixbuf{
		GdkPixbuf: &GdkPixbuf{gpixbuf},
		GObject:   glib.ObjectFromNative(unsafe.Pointer(gpixbuf)),
	}
}

func (p *Pixbuf) Scale(x, y, width, height int, offsetX, offsetY, scaleX, scaleY float64, interp InterpType) *Pixbuf {
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
		GdkPixbuf: &GdkPixbuf{gpixbuf},
		GObject:   glib.ObjectFromNative(unsafe.Pointer(gpixbuf)),
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

func (p *Pixbuf) RotateSimple(angle PixbufRotation) *Pixbuf {
	gpixbuf := C.gdk_pixbuf_rotate_simple(p.GPixbuf, C.GdkPixbufRotation(angle))
	return &Pixbuf{
		GdkPixbuf: &GdkPixbuf{gpixbuf},
		GObject:   glib.ObjectFromNative(unsafe.Pointer(gpixbuf)),
	}
}

func (p *Pixbuf) Flip(horizontal bool) *Pixbuf {
	gpixbuf := C.gdk_pixbuf_flip(p.GPixbuf, gbool(horizontal))
	return &Pixbuf{
		GdkPixbuf: &GdkPixbuf{gpixbuf},
		GObject:   glib.ObjectFromNative(unsafe.Pointer(gpixbuf)),
	}
}

func (p *Pixbuf) Fill(pixel uint32) {
	C.gdk_pixbuf_fill(p.GPixbuf, C.guint32(pixel))
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

func (p *Pixbuf) GetPixels() []byte {
	ptr := C.gdk_pixbuf_get_pixels(
		p.GPixbuf,
	)
	return (*[1 << 30]byte)(unsafe.Pointer(ptr))[:]
}

// guchar * gdk_pixbuf_get_pixels_with_length (const GdkPixbuf *pixbuf, guint *length);
//
// Retuns a slice of byte backed by a C array of pixbuf data.
func (p *Pixbuf) GetPixelsWithLength() []byte {
	panic_if_version_older(2, 26, 0, "gdk_pixbuf_get_pixels_with_length()")
	length := C.guint(0)
	ptr := C._gdk_pixbuf_get_pixels_with_length(
		p.GPixbuf,
		&length,
	)
	return (*[1 << 30]byte)(unsafe.Pointer(ptr))[:length]
}

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

func NewLoaderWithType(image_type string) (loader *Loader, err *glib.Error) {
	var gerr *C.GError
	ptr := C.CString(image_type)
	defer cfree(ptr)
	loader = &Loader{
		C.gdk_pixbuf_loader_new_with_type(ptr, &gerr)}
	if gerr != nil {
		err = glib.ErrorFromNative(unsafe.Pointer(gerr))
	}
	return
}

func NewLoaderWithMimeType(mime_type string) (loader *Loader, err *glib.Error) {
	var error *C.GError
	ptr := C.CString(mime_type)
	defer cfree(ptr)
	loader = &Loader{
		C.gdk_pixbuf_loader_new_with_mime_type(ptr, &error)}
	err = glib.ErrorFromNative(unsafe.Pointer(error))
	return
}

func (v Loader) GetPixbuf() *Pixbuf {
	gpixbuf := C.gdk_pixbuf_loader_get_pixbuf(v.GPixbufLoader)
	return &Pixbuf{
		GdkPixbuf: &GdkPixbuf{gpixbuf},
		GObject:   glib.ObjectFromNative(unsafe.Pointer(gpixbuf)),
	}
}

func (v Loader) Write(buf []byte) (bool, *glib.Error) {
	var err *C.GError
	var pbuf *byte
	pbuf = &buf[0]
	ret := gobool(C.gdk_pixbuf_loader_write(v.GPixbufLoader, C.to_gucharptr(unsafe.Pointer(pbuf)), C.gsize(len(buf)), &err))
	if err != nil {
		return ret, glib.ErrorFromNative(unsafe.Pointer(err))
	}
	return ret, nil
}

func (v Loader) Close() (bool, *glib.Error) {
	var err *C.GError
	ret := gobool(C.gdk_pixbuf_loader_close(v.GPixbufLoader, &err))
	if err != nil {
		return ret, glib.ErrorFromNative(unsafe.Pointer(err))
	}
	return ret, nil
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
