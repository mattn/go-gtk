package gio

// #include "gio.go.h"
// #cgo pkg-config: gtk+-2.0
import "C"
import (
	"unsafe"

	"github.com/mattn/go-gtk/glib"
)

func cfree(s *C.char)            { C.freeCstr(s) }
func cstring(s *C.gchar) *C.char { return C.toCstr(s) }
func gostring(s *C.gchar) string { return C.GoString(cstring(s)) }

//-----------------------------------------------------------------------
// GFile
//-----------------------------------------------------------------------
type GFile struct {
	GFile *C.GFile
}

// NewGFileForPath is g_file_new_for_path
func NewGFileForPath(filename string) *GFile {
	ptrFilename := C.CString(filename)
	defer cfree(ptrFilename)

	return &GFile{C.g_file_new_for_path(ptrFilename)}
}

type GFileQueryInfoFlags C.GFileQueryInfoFlags

var (
	GFileQueryInfoNone             GFileQueryInfoFlags = C.G_FILE_QUERY_INFO_NONE
	GFileQueryInfoNoFollowSymlinks GFileQueryInfoFlags = C.G_FILE_QUERY_INFO_NOFOLLOW_SYMLINKS
)

// QueryInfo is g_file_query_info
func (f *GFile) QueryInfo(attributes string, flags GFileQueryInfoFlags) (*GFileInfo, error) {
	ptr := C.CString(attributes)
	defer cfree(ptr)

	var gerror *C.GError

	gfileinfo := C.g_file_query_info(
		f.GFile,
		ptr,
		C.GFileQueryInfoFlags(flags),
		nil, // nil is GCancellable, not yet implemented
		&gerror,
	)
	if gerror != nil {
		return nil, glib.ErrorFromNative(unsafe.Pointer(gerror))
	}

	return &GFileInfo{gfileinfo}, nil
}

//-----------------------------------------------------------------------
// GFileInfo
//-----------------------------------------------------------------------
type GFileInfo struct {
	GFileInfo *C.GFileInfo
}

// GetSymbolicIcon is g_file_info_get_symbolic_icon
func (fi *GFileInfo) GetSymbolicIcon() *GIcon {
	return &GIcon{C.g_file_info_get_symbolic_icon(fi.GFileInfo)}
}

// GetIcon is g_file_info_get_icon
func (fi *GFileInfo) GetIcon() *GIcon {
	return &GIcon{C.g_file_info_get_icon(fi.GFileInfo)}
}

//-----------------------------------------------------------------------
// GIcon
//-----------------------------------------------------------------------
type GIcon struct {
	GIcon *C.GIcon
}

//ToString is g_icon_to_string
func (icon *GIcon) ToString() string {
	return gostring(C.g_icon_to_string(icon.GIcon))
}
