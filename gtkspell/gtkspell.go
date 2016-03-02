// +build !cgocheck

package gtkspell

/*
#include <gtk/gtk.h>
#include <gtkspell/gtkspell.h>
#include <stdlib.h>

static GtkTextView* to_GtkTextView(void* w) { return GTK_TEXT_VIEW(w); }
static inline gchar* to_gcharptr(const char* s) { return (gchar*)s; }
*/
// #cgo pkg-config: gtkspell-2.0
import "C"
import "unsafe"

import "github.com/mattn/go-gtk/glib"
import "github.com/mattn/go-gtk/gtk"

//-----------------------------------------------------------------------
// GtkSpell
//-----------------------------------------------------------------------
type GtkSpell struct {
	Spell *C.GtkSpell
}

func New(textview *gtk.TextView, language string) (*GtkSpell, *glib.Error) {
	var lang *C.char
	if len(language) > 0 {
		lang = C.CString(language)
		defer C.free(unsafe.Pointer(lang))
	}

	var gerror *C.GError
	v := C.gtkspell_new_attach(C.to_GtkTextView(unsafe.Pointer(&textview.Widget)), C.to_gcharptr(lang), &gerror)
	if gerror != nil {
		return nil, glib.ErrorFromNative(unsafe.Pointer(gerror))
	}
	return &GtkSpell{v}, nil
}

func (spell *GtkSpell) SetLanguage(language string) *glib.Error {
	lang := C.CString(language)
	defer C.free(unsafe.Pointer(lang))

	var gerror *C.GError
	C.gtkspell_set_language(spell.Spell, C.to_gcharptr(lang), &gerror)
	if gerror != nil {
		return glib.ErrorFromNative(unsafe.Pointer(gerror))
	}
	return nil
}

func (spell *GtkSpell) Recheck() {
	C.gtkspell_recheck_all(spell.Spell)
}
