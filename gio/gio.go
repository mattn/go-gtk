// +build !cgocheck

package gio

// #include "gio.go.h"
// #cgo pkg-config: gtk+-2.0
import "C"

type GFile struct {
	GFile *C.GFile
}

func cfree(s *C.char)            { C.freeCstr(s) }
func cstring(s *C.gchar) *C.char { return C.toCstr(s) }
func gostring(s *C.gchar) string { return C.GoString(cstring(s)) }

func NewGFileForPath(filename string) *GFile {
	ptrFilename := C.CString(filename)
	defer cfree(ptrFilename)

	return &GFile{C.g_file_new_for_path(ptrFilename)}
}

func (f *GFile) QueryInfo() *GFileInfo {
	ptr := C.CString("standard::*")
	defer cfree(ptr)

	// C.CString(C.G_FILE_ATTRIBUTE_STANDARD_TYPE)
	return &GFileInfo{C.g_file_query_info(f.GFile, ptr, C.G_FILE_QUERY_INFO_NONE, nil, nil)}
}

//
//
type GFileInfo struct {
	GFileInfo *C.GFileInfo
}

func (fi *GFileInfo) GetSymbolicIcon() *GIcon {
	return &GIcon{C.g_file_info_get_symbolic_icon(fi.GFileInfo)}
}

func (fi *GFileInfo) GetIcon() *GIcon {
	return &GIcon{C.g_file_info_get_icon(fi.GFileInfo)}
}

type GIcon struct {
	GIcon *C.GIcon
}

func (icon *GIcon) ToString() string {
	return gostring(C.g_icon_to_string(icon.GIcon))
}
