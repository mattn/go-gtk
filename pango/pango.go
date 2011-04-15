package pango

/*
#ifndef uintptr
#define uintptr unsigned int*
#endif
#include <pango/pango.h>
#include <unistd.h>
#include <stdlib.h>
#include <stdint.h>
#include <stdarg.h>
#include <string.h>

//static const gchar* to_gcharptr(const char* s) { return (const gchar*)s; }
//static guchar* to_gucharptr(void* s) { return (guchar*)s; }

//static void free_string(char* s) { free(s); }

//static gchar* to_gcharptr(char* s) { return (gchar*)s; }

//static void free_string(char* s) { free(s); }
*/
import "C"

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

type PangoWrapMode int

const (
	PANGO_WRAP_WORD      PangoWrapMode = 0
	PANGO_WRAP_CHAR      PangoWrapMode = 1
	PANGO_WRAP_WORD_CHAR PangoWrapMode = 2
)

type PangoEllipsizeMode int

const (
	PANGO_ELLIPSIZE_NONE   PangoEllipsizeMode = 0
	PANGO_ELLIPSIZE_START  PangoEllipsizeMode = 1
	PANGO_ELLIPSIZE_MIDDLE PangoEllipsizeMode = 2
	PANGO_ELLIPSIZE_END    PangoEllipsizeMode = 3
)
