// +build !cgocheck

package pango

// #include "pango.go.h"
// #cgo pkg-config: pango
import "C"
import "unsafe"

const (
	SCALE = C.PANGO_SCALE
)

func bool2gboolean(b bool) C.gboolean {
	if b {
		return C.gboolean(1)
	}
	return C.gboolean(0)
}

func gboolean2bool(b C.gboolean) bool {
	if b != 0 {
		return true
	}
	return false
}

type WrapMode int

const (
	WRAP_WORD      WrapMode = 0
	WRAP_CHAR      WrapMode = 1
	WRAP_WORD_CHAR WrapMode = 2
)

type EllipsizeMode int

const (
	ELLIPSIZE_NONE   EllipsizeMode = 0
	ELLIPSIZE_START  EllipsizeMode = 1
	ELLIPSIZE_MIDDLE EllipsizeMode = 2
	ELLIPSIZE_END    EllipsizeMode = 3
)

type Context struct {
	GContext *C.PangoContext
}

type Layout struct {
	GLayout *C.PangoLayout
}

type FontDescription struct {
	GFontDescription *C.PangoFontDescription
}

func ContextFromUnsafe(context unsafe.Pointer) *Context {
	return &Context{(*C.PangoContext)(context)}
}

func (v *Layout) Unref() {
	C.g_object_unref(C.gpointer(v.GLayout))
}

func (v *Layout) SetWidth(width int) {
	C.pango_layout_set_width(v.GLayout, C.int(width))
}

func (v *Layout) SetFontDescription(d *FontDescription) {
	C.pango_layout_set_font_description(v.GLayout, d.GFontDescription)
}

func (v *Layout) SetText(s string) {
	cs := C.CString(s)
	C.pango_layout_set_text(v.GLayout, cs, -1)
	C.free(unsafe.Pointer(cs))
}

func NewLayout(ctx *Context) *Layout {
	return &Layout{C.pango_layout_new(ctx.GContext)}
}

func NewFontDescription() *FontDescription {
	return &FontDescription{C.pango_font_description_new()}
}

func (v *FontDescription) Free() {
	C.pango_font_description_free(v.GFontDescription)
}

func (v *FontDescription) SetSize(size int) {
	C.pango_font_description_set_size(v.GFontDescription, C.gint(size))
}

func (v *FontDescription) Copy() *FontDescription {
	return &FontDescription{C.pango_font_description_copy(v.GFontDescription)}
}

func (f *FontDescription) GetSize() int {
	return int(C.pango_font_description_get_size(f.GFontDescription))
}
