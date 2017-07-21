// +build !cgocheck

package gtksourceview

// #include "gtksourceview.go.h"
// #cgo pkg-config: gtksourceview-2.0
import "C"
import (
	"unsafe"

	"github.com/mattn/go-gtk/gtk"
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

//-----------------------------------------------------------------------
// GtkSourceBuffer
//-----------------------------------------------------------------------
type SourceBuffer struct {
	GSourceBuffer *C.GtkSourceBuffer
	*gtk.TextBuffer
}

func NewSourceBuffer() *SourceBuffer {
	v := C.gtk_source_buffer_new(nil)
	return &SourceBuffer{v, gtk.NewTextBufferFromPointer(unsafe.Pointer(v))}
}

func NewSourceBufferWithLanguage(lang *SourceLanguage) *SourceBuffer {
	v := C.gtk_source_buffer_new_with_language(lang.GSourceLanguage)
	return &SourceBuffer{v, gtk.NewTextBufferFromPointer(unsafe.Pointer(v))}
}

func (v *SourceBuffer) GetNativeBuffer() unsafe.Pointer {
	return unsafe.Pointer(v.GSourceBuffer)
}

func (v *SourceBuffer) SetHighlightSyntax(highlight bool) {
	C.gtk_source_buffer_set_highlight_syntax(v.GSourceBuffer, gbool(highlight))
}

func (v *SourceBuffer) GetHighlightSyntax() bool {
	return gobool(C.gtk_source_buffer_get_highlight_syntax(v.GSourceBuffer))
}

func (v *SourceBuffer) SetHighlightMatchingBrackets(hl bool) {
	C.gtk_source_buffer_set_highlight_matching_brackets(v.GSourceBuffer, gbool(hl))
}

func (v *SourceBuffer) SetLanguage(lang *SourceLanguage) {
	C.gtk_source_buffer_set_language(v.GSourceBuffer, lang.GSourceLanguage)
}

func (v *SourceBuffer) GetLanguage() *SourceLanguage {
	return &SourceLanguage{C.gtk_source_buffer_get_language(v.GSourceBuffer)}
}

func (v *SourceBuffer) BeginNotUndoableAction() {
	C.gtk_source_buffer_begin_not_undoable_action(v.GSourceBuffer)
}

func (v *SourceBuffer) EndNotUndoableAction() {
	C.gtk_source_buffer_end_not_undoable_action(v.GSourceBuffer)
}

func (v *SourceBuffer) SetStyleScheme(scheme *SourceStyleScheme) {
	C.gtk_source_buffer_set_style_scheme(v.GSourceBuffer, scheme.GSourceStyleScheme)
}

//-----------------------------------------------------------------------
// GtkSourceView
//-----------------------------------------------------------------------
type SourceView struct {
	gtk.TextView
}

func NewSourceView() *SourceView {
	return &SourceView{gtk.TextView{gtk.Container{
		*gtk.WidgetFromNative(unsafe.Pointer(C.gtk_source_view_new()))}}}
}

func NewSourceViewWithBuffer(buf *SourceBuffer) *SourceView {
	return &SourceView{gtk.TextView{gtk.Container{
		*gtk.WidgetFromNative(unsafe.Pointer(C.gtk_source_view_new_with_buffer(buf.GSourceBuffer)))}}}
}

func (v *SourceView) ToNativeSourceView() *C.GtkSourceView {
	return C.toGtkSourceView(unsafe.Pointer(v.GWidget))
}

func (v *SourceView) SetAutoIndent(enable bool) {
	C.gtk_source_view_set_auto_indent(v.ToNativeSourceView(), gbool(enable))
}

func (v *SourceView) GetAutoIndent() bool {
	return gobool(C.gtk_source_view_get_auto_indent(v.ToNativeSourceView()))
}

func (v *SourceView) SetHighlightCurrentLine(enable bool) {
	C.gtk_source_view_set_highlight_current_line(v.ToNativeSourceView(), gbool(enable))
}

func (v *SourceView) GetHighlightCurrentLine() bool {
	return gobool(C.gtk_source_view_get_highlight_current_line(v.ToNativeSourceView()))
}

func (v *SourceView) SetShowLineNumbers(enable bool) {
	C.gtk_source_view_set_show_line_numbers(v.ToNativeSourceView(), gbool(enable))
}

func (v *SourceView) GetShowLineNumbers() bool {
	return gobool(C.gtk_source_view_get_show_line_numbers(v.ToNativeSourceView()))
}

func (v *SourceView) SetRightMarginPosition(pos uint) {
	C.gtk_source_view_set_right_margin_position(v.ToNativeSourceView(), C.guint(pos))
}

func (v *SourceView) GetRightMarginPosition() uint {
	return uint(C.gtk_source_view_get_right_margin_position(v.ToNativeSourceView()))
}

func (v *SourceView) SetIndentWidth(width int) {
	C.gtk_source_view_set_indent_width(v.ToNativeSourceView(), C.gint(width))
}

func (v *SourceView) GetIndentWidth() int {
	return int(C.gtk_source_view_get_indent_width(v.ToNativeSourceView()))
}

func (v *SourceView) SetShowRightMargin(enable bool) {
	C.gtk_source_view_set_show_right_margin(v.ToNativeSourceView(), gbool(enable))
}

func (v *SourceView) GetShowRightMargin() bool {
	return gobool(C.gtk_source_view_get_show_right_margin(v.ToNativeSourceView()))
}

func (v *SourceView) SetInsertSpacesInsteadOfTabs(enable bool) {
	C.gtk_source_view_set_insert_spaces_instead_of_tabs(v.ToNativeSourceView(), gbool(enable))
}

func (v *SourceView) GetInsertSpacesInsteadOfTabs() bool {
	return gobool(C.gtk_source_view_get_insert_spaces_instead_of_tabs(v.ToNativeSourceView()))
}

type SourceDrawSpacesFlags int

const (
	SOURCE_DRAW_SPACES_SPACE    SourceDrawSpacesFlags = 1 << 0
	SOURCE_DRAW_SPACES_TAB      SourceDrawSpacesFlags = 1 << 1
	SOURCE_DRAW_SPACES_NEWLINE  SourceDrawSpacesFlags = 1 << 2
	SOURCE_DRAW_SPACES_NBSP     SourceDrawSpacesFlags = 1 << 3
	SOURCE_DRAW_SPACES_LEADING  SourceDrawSpacesFlags = 1 << 4
	SOURCE_DRAW_SPACES_TEXT     SourceDrawSpacesFlags = 1 << 5
	SOURCE_DRAW_SPACES_TRAILING SourceDrawSpacesFlags = 1 << 6
	SOURCE_DRAW_SPACES_ALL      SourceDrawSpacesFlags = (SOURCE_DRAW_SPACES_SPACE |
		SOURCE_DRAW_SPACES_TAB |
		SOURCE_DRAW_SPACES_NEWLINE |
		SOURCE_DRAW_SPACES_NBSP |
		SOURCE_DRAW_SPACES_LEADING |
		SOURCE_DRAW_SPACES_TEXT |
		SOURCE_DRAW_SPACES_TRAILING)
)

func (v *SourceView) SetDrawSpaces(flags SourceDrawSpacesFlags) {
	C.gtk_source_view_set_draw_spaces(v.ToNativeSourceView(),
		C.GtkSourceDrawSpacesFlags(flags))
}

func (v *SourceView) GetDrawSpaces() SourceDrawSpacesFlags {
	return SourceDrawSpacesFlags(C.gtk_source_view_get_draw_spaces(v.ToNativeSourceView()))
}

func (v *SourceView) SetTabWidth(width uint) {
	C.gtk_source_view_set_tab_width(v.ToNativeSourceView(),
		C.guint(width))
}

func (v *SourceView) GetTabWidth() uint {
	return uint(C.gtk_source_view_get_tab_width(v.ToNativeSourceView()))
}

type SourceSmartHomeEndType int

const (
	SOURCE_SMART_HOME_END_DISABLED SourceSmartHomeEndType = 0
	SOURCE_SMART_HOME_END_BEFORE   SourceSmartHomeEndType = 1
	SOURCE_SMART_HOME_END_AFTER    SourceSmartHomeEndType = 2
	SOURCE_SMART_HOME_END_ALWAYS   SourceSmartHomeEndType = 3
)

func (v *SourceView) SetSmartHomeEnd(flags SourceSmartHomeEndType) {
	C.gtk_source_view_set_smart_home_end(v.ToNativeSourceView(),
		C.GtkSourceSmartHomeEndType(flags))
}

func (v *SourceView) GetSmartHomeEnd() SourceSmartHomeEndType {
	return SourceSmartHomeEndType(C.gtk_source_view_get_smart_home_end(v.ToNativeSourceView()))
}

//-----------------------------------------------------------------------
// GtkSourceLanguage
//-----------------------------------------------------------------------
type SourceLanguage struct {
	GSourceLanguage *C.GtkSourceLanguage
}

func (v *SourceLanguage) GetId() string {
	return gostring(C.gtk_source_language_get_id(v.GSourceLanguage))
}

func (v *SourceLanguage) GetName() string {
	return gostring(C.gtk_source_language_get_name(v.GSourceLanguage))
}

func (v *SourceLanguage) GetSection() string {
	return gostring(C.gtk_source_language_get_section(v.GSourceLanguage))
}

func (v *SourceLanguage) GetHidden() bool {
	return gobool(C.gtk_source_language_get_hidden(v.GSourceLanguage))
}

func (v *SourceLanguage) GetMetadata(name string) string {
	cname := C.CString(name)
	defer cfree(cname)
	return gostring(C.gtk_source_language_get_metadata(v.GSourceLanguage, gstring(cname)))
}

func (v *SourceLanguage) GetMimeTypes() []string {
	var types []string
	ctypes := C.gtk_source_language_get_mime_types(v.GSourceLanguage)
	for {
		types = append(types, gostring(*ctypes))
		ctypes = C.nextGstr(ctypes)
		if *ctypes == nil {
			break
		}
	}
	return types
}

func (v *SourceLanguage) GetGlobs() []string {
	var globs []string
	cglobs := C.gtk_source_language_get_globs(v.GSourceLanguage)
	for {
		globs = append(globs, gostring(*cglobs))
		cglobs = C.nextGstr(cglobs)
		if *cglobs == nil {
			break
		}
	}
	return globs
}

func (v *SourceLanguage) GetStyleName(styleId string) string {
	cstyleId := C.CString(styleId)
	defer cfree(cstyleId)
	return gostring(C.gtk_source_language_get_metadata(v.GSourceLanguage, gstring(cstyleId)))
}

func (v *SourceLanguage) GetStyleIds() []string {
	var ids []string
	cids := C.gtk_source_language_get_globs(v.GSourceLanguage)
	for {
		ids = append(ids, gostring(*cids))
		cids = C.nextGstr(cids)
		if *cids == nil {
			break
		}
	}
	return ids
}

//FINISH

//-----------------------------------------------------------------------
// GtkSourceLanguageManager
//-----------------------------------------------------------------------
type SourceLanguageManager struct {
	GSourceLanguageManager *C.GtkSourceLanguageManager
}

func NewSourceLanguageManager() *SourceLanguageManager {
	return &SourceLanguageManager{C.gtk_source_language_manager_new()}
}

func SourceLanguageManagerGetDefault() *SourceLanguageManager {
	return &SourceLanguageManager{C.gtk_source_language_manager_get_default()}
}

func (v *SourceLanguageManager) SetSearchPath(paths []string) {
	cpaths := C.make_strings(C.int(len(paths) + 1))
	for i, path := range paths {
		ptr := C.CString(path)
		defer cfree(ptr)
		C.set_string(cpaths, C.int(i), gstring(ptr))
	}
	C.set_string(cpaths, C.int(len(paths)), nil)
	C.gtk_source_language_manager_set_search_path(v.GSourceLanguageManager, cpaths)
	C.destroy_strings(cpaths)
}

func (v *SourceLanguageManager) GetSearchPath() []string {
	var dirs []string
	cdirs := C.gtk_source_language_manager_get_search_path(v.GSourceLanguageManager)
	for {
		dirs = append(dirs, gostring(*cdirs))
		cdirs = C.nextGstr(cdirs)
		if *cdirs == nil {
			break
		}
	}
	return dirs
}

func (v *SourceLanguageManager) GetLanguageIds() []string {
	var ids []string
	cids := C.gtk_source_language_manager_get_language_ids(v.GSourceLanguageManager)
	for {
		ids = append(ids, gostring(*cids))
		cids = C.nextGstr(cids)
		if *cids == nil {
			break
		}
	}
	return ids
}

func (v *SourceLanguageManager) GetLanguage(id string) *SourceLanguage {
	cid := C.CString(id)
	defer cfree(cid)
	return &SourceLanguage{C.gtk_source_language_manager_get_language(v.GSourceLanguageManager, gstring(cid))}
}

func (v *SourceLanguageManager) GuessLanguage(filename string, contentType string) *SourceLanguage {
	if filename == "" {
		cct := C.CString(contentType)
		defer cfree(cct)
		return &SourceLanguage{C.gtk_source_language_manager_guess_language(v.GSourceLanguageManager, nil, gstring(cct))}
	}
	cfn := C.CString(filename)
	defer cfree(cfn)
	return &SourceLanguage{C.gtk_source_language_manager_guess_language(v.GSourceLanguageManager, gstring(cfn), nil)}
}

//FINISH

//-----------------------------------------------------------------------
// GtkSourceStyle
//-----------------------------------------------------------------------
type SourceStyleScheme struct {
	GSourceStyleScheme *C.GtkSourceStyleScheme
}

// gtk_source_style_scheme_get_id
// gtk_source_style_scheme_get_name
// gtk_source_style_scheme_get_description
// gtk_source_style_scheme_get_authors
// gtk_source_style_scheme_get_filename
// gtk_source_style_scheme_get_style

//-----------------------------------------------------------------------
// GtkSourceStyleSchemeManager
//-----------------------------------------------------------------------
type SourceStyleSchemeManager struct {
	GSourceStyleSchemeManager *C.GtkSourceStyleSchemeManager
}

func NewSourceStyleSchemeManager() *SourceStyleSchemeManager {
	return &SourceStyleSchemeManager{C.gtk_source_style_scheme_manager_new()}
}

func SourceStyleSchemeManagerGetDefault() *SourceStyleSchemeManager {
	return &SourceStyleSchemeManager{C.gtk_source_style_scheme_manager_get_default()}
}

func (v *SourceStyleSchemeManager) GetScheme(scheme_id string) *SourceStyleScheme {
	cscheme_id := C.CString(scheme_id)
	defer cfree(cscheme_id)
	return &SourceStyleScheme{C.gtk_source_style_scheme_manager_get_scheme(v.GSourceStyleSchemeManager, gstring(cscheme_id))}
}

func (v *SourceStyleSchemeManager) SetSearchPath(paths []string) {
	cpaths := C.make_strings(C.int(len(paths) + 1))
	for i, path := range paths {
		ptr := C.CString(path)
		defer cfree(ptr)
		C.set_string(cpaths, C.int(i), gstring(ptr))
	}
	C.set_string(cpaths, C.int(len(paths)), nil)
	C.gtk_source_style_scheme_manager_set_search_path(v.GSourceStyleSchemeManager, cpaths)
}

func (v *SourceStyleSchemeManager) GetSearchPath() []string {
	var dirs []string
	cdirs := C.gtk_source_style_scheme_manager_get_search_path(v.GSourceStyleSchemeManager)
	for {
		dirs = append(dirs, gostring(*cdirs))
		cdirs = C.nextGstr(cdirs)
		if *cdirs == nil {
			break
		}
	}
	return dirs
}

func (v *SourceStyleSchemeManager) AppendSearchPath(path string) {
	cpath := C.CString(path)
	defer cfree(cpath)
	C.gtk_source_style_scheme_manager_append_search_path(v.GSourceStyleSchemeManager, gstring(cpath))
}

func (v *SourceStyleSchemeManager) PrepandSearchPath(path string) {
	cpath := C.CString(path)
	defer cfree(cpath)
	C.gtk_source_style_scheme_manager_prepend_search_path(v.GSourceStyleSchemeManager, gstring(cpath))
}

func (v *SourceStyleSchemeManager) GetSchemeIds() []string {
	var ids []string
	cids := C.gtk_source_style_scheme_manager_get_scheme_ids(v.GSourceStyleSchemeManager)
	for {
		ids = append(ids, gostring(*cids))
		cids = C.nextGstr(cids)
		if *cids == nil {
			break
		}
	}
	return ids
}

func (v *SourceStyleSchemeManager) ForseRescan() {
	C.gtk_source_style_scheme_manager_force_rescan(v.GSourceStyleSchemeManager)
}

//FINISH
