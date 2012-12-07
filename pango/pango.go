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
