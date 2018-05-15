package gio

// #include "gio.go.h"
// #cgo pkg-config: gio-2.0
import "C"
import (
	"unsafe"

	"github.com/mattn/go-gtk/glib"
)

func cfree(s *C.char)            { C.freeCstr(s) }
func gstring(s *C.char) *C.gchar { return C.toGstr(s) }
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

// NewGFileForURI is g_file_new_for_uri
func NewGFileForURI(uriFilename string) *GFile {
	ptrFilename := C.CString(uriFilename)
	defer cfree(ptrFilename)

	return &GFile{C.g_file_new_for_uri(ptrFilename)}
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

// GetPath is `g_file_get_path`, return real, absolute, canonical path. It might contain symlinks.
func (f *GFile) GetPath() string {
	return gostring(gstring(C.g_file_get_path(f.GFile)))
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

type GInputStream struct {
	GInputStream *C.GInputStream
}

func NewMemoryInputStreamFromBytes(buffer []byte) *GInputStream {
	pbyt := &buffer[0]
	gbytes := C.g_bytes_new_take(C.gpointer(unsafe.Pointer(pbyt)), C.gsize(len(buffer)))

	stream := C.g_memory_input_stream_new_from_bytes(gbytes)
	return &GInputStream{stream}
}

func NewCancellable() *C.GCancellable {
	return C.g_cancellable_new()
}
