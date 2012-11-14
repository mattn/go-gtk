package gtksourceview

/*
#include <gtksourceview/gtksourceview.h>
#include <gtksourceview/gtksourcebuffer.h>
#include <gtksourceview/gtksourcelanguage.h>
#include <gtksourceview/gtksourcelanguagemanager.h>
#include <stdlib.h>

static inline gchar** make_strings(int count) {
	return (gchar**)malloc(sizeof(gchar*) * count);
}

static inline void destroy_strings(gchar** strings) {
	free(strings);
}

static inline void set_string(gchar** strings, int n, gchar* str) {
	strings[n] = str;
}

static inline GObject* to_GObject(void* o) { return G_OBJECT(o); }
static inline gchar* to_gcharptr(const char* s) { return (gchar*)s; }
static inline char* to_charptr(const gchar* s) { return (char*)s; }
static inline gchar** next_gcharptr(gchar** s) { return (s+1); }
static inline void free_string(char* s) { free(s); }
static GtkSourceView* to_GtkSourceView(void* w) { return GTK_SOURCE_VIEW(w); }
*/
// #cgo pkg-config: gtksourceview-2.0
import "C"
import "github.com/mattn/go-gtk/gtk"
import "unsafe"

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

//-----------------------------------------------------------------------
// GtkSourceBuffer
//-----------------------------------------------------------------------
type GtkSourceBuffer struct {
	SourceBuffer *C.GtkSourceBuffer
	gtk.GtkTextBuffer
}

func SourceBuffer() *GtkSourceBuffer {
	v := C.gtk_source_buffer_new(nil)
	return &GtkSourceBuffer{v, gtk.TextBufferFromPointer(unsafe.Pointer(v))}
}
func SourceBufferWithLanguage(lang *GtkSourceLanguage) *GtkSourceBuffer {
	v := C.gtk_source_buffer_new_with_language(lang.SourceLanguage)
	return &GtkSourceBuffer{v, gtk.TextBufferFromPointer(unsafe.Pointer(v))}
}
func (v *GtkSourceBuffer) GetNativeBuffer() unsafe.Pointer {
	return unsafe.Pointer(v.SourceBuffer)
}
func (v *GtkSourceBuffer) SetHighlightSyntax(highlight bool) {
	C.gtk_source_buffer_set_highlight_syntax(v.SourceBuffer, bool2gboolean(highlight))
}
func (v *GtkSourceBuffer) GetHighlightSyntax() bool {
	return gboolean2bool(C.gtk_source_buffer_get_highlight_syntax(v.SourceBuffer))
}
func (v *GtkSourceBuffer) SetLanguage(lang *GtkSourceLanguage) {
	C.gtk_source_buffer_set_language(v.SourceBuffer, lang.SourceLanguage)
}
func (v *GtkSourceBuffer) GetLanguage() *GtkSourceLanguage {
	return &GtkSourceLanguage{C.gtk_source_buffer_get_language(v.SourceBuffer)}
}
func (v *GtkSourceBuffer) BeginNotUndoableAction() {
	C.gtk_source_buffer_begin_not_undoable_action(v.SourceBuffer)
}
func (v *GtkSourceBuffer) EndNotUndoableAction() {
	C.gtk_source_buffer_end_not_undoable_action(v.SourceBuffer)
}

//-----------------------------------------------------------------------
// GtkSourceView
//-----------------------------------------------------------------------
type GtkSourceView struct {
	gtk.GtkTextView
}

func SourceView() *GtkSourceView {
	return &GtkSourceView{gtk.GtkTextView{gtk.GtkContainer{
		*gtk.WidgetFromNative(unsafe.Pointer(C.gtk_source_view_new()))}}}
}
func SourceViewWithBuffer(buf *GtkSourceBuffer) *GtkSourceView {
	return &GtkSourceView{gtk.GtkTextView{gtk.GtkContainer{
		*gtk.WidgetFromNative(unsafe.Pointer(C.gtk_source_view_new_with_buffer(buf.SourceBuffer)))}}}
}
func (v *GtkSourceView) ToNativeSourceView() *C.GtkSourceView {
	return C.to_GtkSourceView(unsafe.Pointer(v.Widget))
}
func (v *GtkSourceView) SetAutoIndent(enable bool) {
	C.gtk_source_view_set_auto_indent(v.ToNativeSourceView(), bool2gboolean(enable))
}
func (v *GtkSourceView) GetAutoIndent() bool {
	return gboolean2bool(C.gtk_source_view_get_auto_indent(v.ToNativeSourceView()))
}
func (v *GtkSourceView) SetHighlightCurrentLine(enable bool) {
	C.gtk_source_view_set_highlight_current_line(v.ToNativeSourceView(), bool2gboolean(enable))
}
func (v *GtkSourceView) GetHighlightCurrentLine() bool {
	return gboolean2bool(C.gtk_source_view_get_highlight_current_line(v.ToNativeSourceView()))
}
func (v *GtkSourceView) SetShowLineNumbers(enable bool) {
	C.gtk_source_view_set_show_line_numbers(v.ToNativeSourceView(), bool2gboolean(enable))
}
func (v *GtkSourceView) GetShowLineNumbers() bool {
	return gboolean2bool(C.gtk_source_view_get_show_line_numbers(v.ToNativeSourceView()))
}
func (v *GtkSourceView) SetRightMarginPosition(pos uint) {
	C.gtk_source_view_set_right_margin_position(v.ToNativeSourceView(), C.guint(pos))
}
func (v *GtkSourceView) GetRightMarginPosition() uint {
	return uint(C.gtk_source_view_get_right_margin_position(v.ToNativeSourceView()))
}
func (v *GtkSourceView) SetIndentWidth(width int) {
	C.gtk_source_view_set_indent_width(v.ToNativeSourceView(), C.gint(width))
}
func (v *GtkSourceView) GetIndentWidth() int {
	return int(C.gtk_source_view_get_indent_width(v.ToNativeSourceView()))
}
func (v *GtkSourceView) SetShowRightMargin(enable bool) {
	C.gtk_source_view_set_show_right_margin(v.ToNativeSourceView(), bool2gboolean(enable))
}
func (v *GtkSourceView) GetShowRightMargin() bool {
	return gboolean2bool(C.gtk_source_view_get_show_right_margin(v.ToNativeSourceView()))
}
func (v *GtkSourceView) SetInsertSpacesInsteadOfTabs(enable bool) {
	C.gtk_source_view_set_insert_spaces_instead_of_tabs(v.ToNativeSourceView(), bool2gboolean(enable))
}
func (v *GtkSourceView) GetInsertSpacesInsteadOfTabs() bool {
	return gboolean2bool(C.gtk_source_view_get_insert_spaces_instead_of_tabs(v.ToNativeSourceView()))
}

type GtkSourceDrawSpacesFlags int

const (
	GTK_SOURCE_DRAW_SPACES_SPACE    GtkSourceDrawSpacesFlags = 1 << 0
	GTK_SOURCE_DRAW_SPACES_TAB      GtkSourceDrawSpacesFlags = 1 << 1
	GTK_SOURCE_DRAW_SPACES_NEWLINE  GtkSourceDrawSpacesFlags = 1 << 2
	GTK_SOURCE_DRAW_SPACES_NBSP     GtkSourceDrawSpacesFlags = 1 << 3
	GTK_SOURCE_DRAW_SPACES_LEADING  GtkSourceDrawSpacesFlags = 1 << 4
	GTK_SOURCE_DRAW_SPACES_TEXT     GtkSourceDrawSpacesFlags = 1 << 5
	GTK_SOURCE_DRAW_SPACES_TRAILING GtkSourceDrawSpacesFlags = 1 << 6
	GTK_SOURCE_DRAW_SPACES_ALL      GtkSourceDrawSpacesFlags = (GTK_SOURCE_DRAW_SPACES_SPACE |
		GTK_SOURCE_DRAW_SPACES_TAB |
		GTK_SOURCE_DRAW_SPACES_NEWLINE |
		GTK_SOURCE_DRAW_SPACES_NBSP |
		GTK_SOURCE_DRAW_SPACES_LEADING |
		GTK_SOURCE_DRAW_SPACES_TEXT |
		GTK_SOURCE_DRAW_SPACES_TRAILING)
)

func (v *GtkSourceView) SetDrawSpaces(flags GtkSourceDrawSpacesFlags) {
	C.gtk_source_view_set_draw_spaces(v.ToNativeSourceView(),
		C.GtkSourceDrawSpacesFlags(flags))
}
func (v *GtkSourceView) GetDrawSpaces() GtkSourceDrawSpacesFlags {
	return GtkSourceDrawSpacesFlags(C.gtk_source_view_get_draw_spaces(v.ToNativeSourceView()))
}
func (v *GtkSourceView) SetTabWidth(width uint) {
	C.gtk_source_view_set_tab_width(v.ToNativeSourceView(),
		C.guint(width))
}
func (v *GtkSourceView) GetTabWidth() uint {
	return uint(C.gtk_source_view_get_tab_width(v.ToNativeSourceView()))
}

type GtkSourceSmartHomeEndType int

const (
	GTK_SOURCE_SMART_HOME_END_DISABLED GtkSourceSmartHomeEndType = 0
	GTK_SOURCE_SMART_HOME_END_BEFORE   GtkSourceSmartHomeEndType = 1
	GTK_SOURCE_SMART_HOME_END_AFTER    GtkSourceSmartHomeEndType = 2
	GTK_SOURCE_SMART_HOME_END_ALWAYS   GtkSourceSmartHomeEndType = 3
)

func (v *GtkSourceView) SetSmartHomeEnd(flags GtkSourceSmartHomeEndType) {
	C.gtk_source_view_set_smart_home_end(v.ToNativeSourceView(),
		C.GtkSourceSmartHomeEndType(flags))
}
func (v *GtkSourceView) GetSmartHomeEnd() GtkSourceSmartHomeEndType {
	return GtkSourceSmartHomeEndType(C.gtk_source_view_get_smart_home_end(v.ToNativeSourceView()))
}

//-----------------------------------------------------------------------
// GtkSourceLanguage
//-----------------------------------------------------------------------
type GtkSourceLanguage struct {
	SourceLanguage *C.GtkSourceLanguage
}

func (v *GtkSourceLanguage) GetId() string {
	return C.GoString(C.to_charptr(C.gtk_source_language_get_id(v.SourceLanguage)))
}
func (v *GtkSourceLanguage) GetName() string {
	return C.GoString(C.to_charptr(C.gtk_source_language_get_name(v.SourceLanguage)))
}
func (v *GtkSourceLanguage) GetSection() string {
	return C.GoString(C.to_charptr(C.gtk_source_language_get_section(v.SourceLanguage)))
}
func (v *GtkSourceLanguage) GetHidden() bool {
	return gboolean2bool(C.gtk_source_language_get_hidden(v.SourceLanguage))
}
func (v *GtkSourceLanguage) GetMetadata(name string) string {
	cname := C.CString(name)
	defer C.free_string(cname)
	return C.GoString(C.to_charptr(C.gtk_source_language_get_metadata(v.SourceLanguage, C.to_gcharptr(cname))))
}
func (v *GtkSourceLanguage) GetMimeTypes() []string {
	var types []string
	ctypes := C.gtk_source_language_get_mime_types(v.SourceLanguage)
	for {
		types = append(types, C.GoString(C.to_charptr(*ctypes)))
		ctypes = C.next_gcharptr(ctypes)
		if *ctypes == nil {
			break
		}
	}
	return types
}
func (v *GtkSourceLanguage) GetGlobs() []string {
	var globs []string
	cglobs := C.gtk_source_language_get_globs(v.SourceLanguage)
	for {
		globs = append(globs, C.GoString(C.to_charptr(*cglobs)))
		cglobs = C.next_gcharptr(cglobs)
		if *cglobs == nil {
			break
		}
	}
	return globs
}
func (v *GtkSourceLanguage) GetStyleName(styleId string) string {
	cstyleId := C.CString(styleId)
	defer C.free_string(cstyleId)
	return C.GoString(C.to_charptr(C.gtk_source_language_get_metadata(v.SourceLanguage, C.to_gcharptr(cstyleId))))
}
func (v *GtkSourceLanguage) GetStyleIds() []string {
	var ids []string
	cids := C.gtk_source_language_get_globs(v.SourceLanguage)
	for {
		ids = append(ids, C.GoString(C.to_charptr(*cids)))
		cids = C.next_gcharptr(cids)
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
type GtkSourceLanguageManager struct {
	LanguageManager *C.GtkSourceLanguageManager
}

func SourceLanguageManager() *GtkSourceLanguageManager {
	return &GtkSourceLanguageManager{C.gtk_source_language_manager_new()}
}
func SourceLanguageManagerGetDefault() *GtkSourceLanguageManager {
	return &GtkSourceLanguageManager{C.gtk_source_language_manager_get_default()}
}
func (v *GtkSourceLanguageManager) SetSearchPath(paths []string) {
	cpaths := C.make_strings(C.int(len(paths) + 1))
	for i, path := range paths {
		ptr := C.CString(path)
		defer C.free_string(ptr)
		C.set_string(cpaths, C.int(i), C.to_gcharptr(ptr))
	}
	C.set_string(cpaths, C.int(len(paths)), nil)
	C.gtk_source_language_manager_set_search_path(v.LanguageManager, cpaths)
	C.destroy_strings(cpaths)
}
func (v *GtkSourceLanguageManager) GetSearchPath() []string {
	var dirs []string
	cdirs := C.gtk_source_language_manager_get_search_path(v.LanguageManager)
	for {
		dirs = append(dirs, C.GoString(C.to_charptr(*cdirs)))
		cdirs = C.next_gcharptr(cdirs)
		if *cdirs == nil {
			break
		}
	}
	return dirs
}
func (v *GtkSourceLanguageManager) GetLanguageIds() []string {
	var ids []string
	cids := C.gtk_source_language_manager_get_language_ids(v.LanguageManager)
	for {
		ids = append(ids, C.GoString(C.to_charptr(*cids)))
		cids = C.next_gcharptr(cids)
		if *cids == nil {
			break
		}
	}
	return ids
}
func (v *GtkSourceLanguageManager) GetLanguage(id string) *GtkSourceLanguage {
	cid := C.CString(id)
	defer C.free_string(cid)
	return &GtkSourceLanguage{C.gtk_source_language_manager_get_language(v.LanguageManager,
		C.to_gcharptr(cid))}
}
func (v *GtkSourceLanguageManager) GuessLanguage(filename string, contentType string) *GtkSourceLanguage {
	if filename == "" {
		cct := C.CString(contentType)
		defer C.free_string(cct)
		return &GtkSourceLanguage{C.gtk_source_language_manager_guess_language(v.LanguageManager, nil, C.to_gcharptr(cct))}
	}
	cfn := C.CString(filename)
	defer C.free_string(cfn)
	return &GtkSourceLanguage{C.gtk_source_language_manager_guess_language(v.LanguageManager, C.to_gcharptr(cfn), nil)}
}

//FINISH
