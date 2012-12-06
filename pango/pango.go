package pango

// #include "pango.go.h"
// #cgo pkg-config: pango
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
