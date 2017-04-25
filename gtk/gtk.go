// +build !cgocheck

/*
Go Bindings for Gtk+ 2. Support version 2.16 and later.
*/
package gtk

// #include "gtk.go.h"
// #cgo pkg-config: gtk+-2.0
import "C"
import (
	"fmt"
	"log"
	"reflect"
	"runtime"
	"strings"
	"unsafe"

	"github.com/mattn/go-gtk/gdk"
	"github.com/mattn/go-gtk/gdkpixbuf"
	"github.com/mattn/go-gtk/glib"
	"github.com/mattn/go-gtk/pango"
	"github.com/mattn/go-pointer"
)

func gint(v int) C.gint           { return C.gint(v) }
func guint(v uint) C.guint        { return C.guint(v) }
func guint16(v uint16) C.guint16  { return C.guint16(v) }
func guint32(v uint32) C.guint32  { return C.guint32(v) }
func glong(v int32) C.glong       { return C.glong(v) }
func gdouble(v float64) C.gdouble { return C.gdouble(v) }
func gsize_t(v C.size_t) C.gint   { return C.gint(v) }

func gstring(s *C.char) *C.gchar { return C.toGstr(s) }
func cstring(s *C.gchar) *C.char { return C.toCstr(s) }
func gostring(s *C.gchar) string { return C.GoString(cstring(s)) }

func gslist(l *glib.SList) *C.GSList {
	if l == nil {
		return nil
	}
	return C.to_gslist(unsafe.Pointer(l.ToSList()))
}

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

func WINDOW(p *Window) *C.GtkWindow                 { return C.toGWindow(p.GWidget) }
func DIALOG(p *Dialog) *C.GtkDialog                 { return C.toGDialog(p.GWidget) }
func ABOUT_DIALOG(p *AboutDialog) *C.GtkAboutDialog { return C.toGAboutDialog(p.GWidget) }
func CONTAINER(p *Container) *C.GtkContainer        { return C.toGContainer(p.GWidget) }
func FILE_CHOOSER(p *Widget) *C.GtkFileChooser      { return C.toGFileChooser(p.GWidget) }
func FONT_SELECTION_DIALOG(p *FontSelectionDialog) *C.GtkFontSelectionDialog {
	return C.toGFontSelectionDialog(p.GWidget)
}

func MISC(m *Misc) *C.GtkMisc                              { return C.toGMisc(m.GWidget) }
func LABEL(p *Label) *C.GtkLabel                           { return C.toGLabel(p.GWidget) }
func BUTTON(p *Button) *C.GtkButton                        { return C.toGButton(p.GWidget) }
func SPIN_BUTTON(p *SpinButton) *C.GtkSpinButton           { return C.toGSpinButton(p.GWidget) }
func RADIO_BUTTON(p *RadioButton) *C.GtkRadioButton        { return C.toGRadioButton(p.GWidget) }
func FONT_BUTTON(p *FontButton) *C.GtkFontButton           { return C.toGFontButton(p.GWidget) }
func LINK_BUTTON(p *LinkButton) *C.GtkLinkButton           { return C.toGLinkButton(p.GWidget) }
func COMBO_BOX(p *ComboBox) *C.GtkComboBox                 { return C.toGComboBox(p.GWidget) }
func COMBO_BOX_ENTRY(p *ComboBoxEntry) *C.GtkComboBoxEntry { return C.toGComboBoxEntry(p.GWidget) }
func MESSAGE_DIALOG(p *MessageDialog) *C.GtkMessageDialog  { return C.toGMessageDialog(p.GWidget) }
func COMBO_BOX_TEXT(p *ComboBoxText) *C.GtkComboBoxText    { return C.toGComboBoxText(p.GWidget) }
func ACCESSIBLE(p *Accessible) *C.GtkAccessible            { return C.toGAccessible(unsafe.Pointer(p.Object)) }
func BIN(p *Bin) *C.GtkBin                                 { return C.toGBin(p.GWidget) }
func STATUSBAR(p *Statusbar) *C.GtkStatusbar               { return C.toGStatusbar(p.GWidget) }
func INFO_BAR(p *InfoBar) *C.GtkInfoBar                    { return C.toGInfoBar(p.GWidget) }
func FRAME(p *Frame) *C.GtkFrame                           { return C.toGFrame(p.GWidget) }
func BOX(p *Box) *C.GtkBox                                 { return C.toGBox(p.GWidget) }
func PANED(p *Paned) *C.GtkPaned                           { return C.toGPaned(p.GWidget) }
func TOGGLE_BUTTON(p *ToggleButton) *C.GtkToggleButton     { return C.toGToggleButton(p.GWidget) }
func ACCEL_LABEL(p *AccelLabel) *C.GtkAccelLabel           { return C.toGAccelLabel(p.GWidget) }
func SCALEBUTTON(p *ScaleButton) *C.GtkScaleButton         { return C.toGScaleButton(p.GWidget) }
func STYLE(p *Style) *C.GtkStyle                           { return p.GStyle }
func ENTRY(p *Entry) *C.GtkEntry                           { return C.toGEntry(p.GWidget) }
func ADJUSTMENT(p *Adjustment) *C.GtkAdjustment            { return p.GAdjustment }
func ARROW(p *Arrow) *C.GtkArrow                           { return C.toGArrow(p.GWidget) }
func TEXT_VIEW(p *TextView) *C.GtkTextView                 { return C.toGTextView(p.GWidget) }
func TEXT_BUFFER(p unsafe.Pointer) *C.GtkTextBuffer        { return C.toGTextBuffer(p) }
func TEXT_TAG(p unsafe.Pointer) *C.GtkTextTag              { return C.toGTextTag(p) }
func MENU(p *Menu) *C.GtkMenu                              { return C.toGMenu(p.GWidget) }
func MENU_BAR(p *MenuBar) *C.GtkMenuBar                    { return C.toGMenuBar(p.GWidget) }
func MENU_SHELL(p *Menu) *C.GtkMenuShell                   { return C.toGMenuShell(p.GWidget) } // TODO (GtkMenuShell receiver)
func MENU_BAR_SHELL(p *MenuBar) *C.GtkMenuShell            { return C.toGMenuShell(p.GWidget) } // TODO
func MENU_ITEM(p *MenuItem) *C.GtkMenuItem                 { return C.toGMenuItem(p.GWidget) }
func ITEM(p *Item) *C.GtkItem                              { return C.toGItem(p.GWidget) }
func TOOLBAR(p *Toolbar) *C.GtkToolbar                     { return C.toGToolbar(p.GWidget) }
func TOOL_ITEM(p *ToolItem) *C.GtkToolItem                 { return C.toGToolItem(p.GWidget) }
func SEPARATOR_TOOL_ITEM(p *SeparatorToolItem) *C.GtkSeparatorToolItem {
	return C.toGSeparatorToolItem(p.GWidget)
}

func TOOL_BUTTON(p *ToolButton) *C.GtkToolButton              { return C.toGToolButton(p.GWidget) }
func TOOL_PALETTE(p *ToolPalette) *C.GtkToolPalette           { return C.toGToolPalette(p.GWidget) }
func TOOL_ITEM_GROUP(p *ToolItemGroup) *C.GtkToolItemGroup    { return C.toGToolItemGroup(p.GWidget) }
func MENU_TOOL_BUTTON(p *MenuToolButton) *C.GtkMenuToolButton { return C.toGMenuToolButton(p.GWidget) }
func TOGGLE_TOOL_BUTTON(p *ToggleToolButton) *C.GtkToggleToolButton {
	return C.toGToggleToolButton(p.GWidget)
}

func SCROLLED_WINDOW(p *ScrolledWindow) *C.GtkScrolledWindow { return C.toGScrolledWindow(p.GWidget) }
func VIEWPORT(p *Viewport) *C.GtkViewport                    { return C.toGViewport(p.GWidget) }
func WIDGET(p *Widget) *C.GtkWidget                          { return C.toGWidget(unsafe.Pointer(p.GWidget)) }
func TREE_VIEW(p *TreeView) *C.GtkTreeView                   { return C.toGTreeView(p.GWidget) }
func ICON_VIEW(p *IconView) *C.GtkIconView                   { return C.toGIconView(p.GWidget) }
func CELL_RENDERER_TEXT(p *CellRendererText) *C.GtkCellRendererText {
	return C.toGCellRendererText(p.GCellRenderer)
}

func CELL_RENDERER_TOGGLE(p *CellRendererToggle) *C.GtkCellRendererToggle {
	return C.toGCellRendererToggle(p.GCellRenderer)
}

func SCALE(p *Scale) *C.GtkScale                           { return C.toGScale(p.GWidget) }
func RANGE(p *Range) *C.GtkRange                           { return C.toGRange(p.GWidget) }
func IMAGE(p *Image) *C.GtkImage                           { return C.toGImage(p.GWidget) }
func NOTEBOOK(p *Notebook) *C.GtkNotebook                  { return C.toGNotebook(p.GWidget) }
func TABLE(p *Table) *C.GtkTable                           { return C.toGTable(p.GWidget) }
func DRAWING_AREA(p *DrawingArea) *C.GtkDrawingArea        { return C.toGDrawingArea(p.GWidget) }
func SPINNER(s *Spinner) *C.GtkSpinner                     { return C.toGSpinner(s.GWidget) }
func ASSISTANT(p *Assistant) *C.GtkAssistant               { return C.toGAssistant(p.GWidget) }
func EXPANDER(p *Expander) *C.GtkExpander                  { return C.toGExpander(p.GWidget) }
func ALIGNMENT(p *Alignment) *C.GtkAlignment               { return C.toGAlignment(p.GWidget) }
func PROGRESS_BAR(p *ProgressBar) *C.GtkProgressBar        { return C.toGProgressBar(p.GWidget) }
func FIXED(p *Fixed) *C.GtkFixed                           { return C.toGFixed(p.GWidget) }
func CHECK_MENU_ITEM(p *CheckMenuItem) *C.GtkCheckMenuItem { return C.toGCheckMenuItem(p.GWidget) }
func RADIO_MENU_ITEM(p *RadioMenuItem) *C.GtkRadioMenuItem { return C.toGRadioMenuItem(p.GWidget) }
func LAYOUT(p *Layout) *C.GtkLayout                        { return C.toGLayout(p.GWidget) }
func COLOR_BUTTON(p *ColorButton) *C.GtkColorButton        { return C.toGColorButton(p.GWidget) }
func IMAGE_MENU_ITEM(p *ImageMenuItem) *C.GtkImageMenuItem { return C.toGImageMenuItem(p.GWidget) }
func ACTION(p *Action) *C.GtkAction                        { return C.toGAction(p.Object) }
func TOGGLE_ACTION(p *ToggleAction) *C.GtkToggleAction     { return C.toGToggleAction(p.Object) }
func RADIO_ACTION(p *RadioAction) *C.GtkRadioAction        { return C.toGRadioAction(p.Object) }
func RECENT_ACTION(p *RecentAction) *C.GtkRecentAction     { return C.toGRecentAction(p.Object) }
func ACTION_GROUP(p *ActionGroup) *C.GtkActionGroup        { return C.toGActionGroup(p.Object) }
func ACTIVATABLE(p *Activatable) *C.GtkActivatable         { return C.toGActivatable(p.GWidget) }
func AS_GWIDGET(p unsafe.Pointer) *C.GtkWidget             { return C.toGWidget(p) }
func UI_MANAGER(p *UIManager) *C.GtkUIManager              { return C.toGUIManager(p.Object) }
func FONT_SELECTION(p *FontSelection) *C.GtkFontSelection  { return C.toGFontSelection(p.GWidget) }

//static inline GtkFileFilter* toGFileFilter(gpointer p) { return GTK_FILE_FILTER(p); }

func panic_if_version_older(major int, minor int, micro int, function string) {
	if C._check_version(C.int(major), C.int(minor), C.int(micro)) == 0 {
		log.Panicf("%s is not provided on your GTK, version %d.%d is required\n", function, major, minor)
	}
}

func panic_if_version_older_auto(major, minor, micro int) {
	if C._check_version(C.int(major), C.int(minor), C.int(micro)) != 0 {
		return
	}
	formatStr := "%s is not provided on your GTK, version %d.%d is required\n"
	if pc, _, _, ok := runtime.Caller(1); ok {
		log.Panicf(formatStr, runtime.FuncForPC(pc).Name(), major, minor)
	} else {
		log.Panicf("GTK version %d.%d is required (unknown caller, see stack)\n",
			major, minor)
	}
}

func deprecated_since(major int, minor int, micro int, function string) {
	if C._check_version(C.int(major), C.int(minor), C.int(micro)) != 0 {
		log.Printf("\nWarning: %s is deprecated since gtk %d.%d\n", function, major, minor)
	}
}

func variadicButtonsToArrays(buttons []interface{}) ([]string, []ResponseType) {
	if len(buttons)%2 != 0 {
		argumentPanic("variadic parameter must be even (couples of string-ResponseType (button label - button response)")
	}
	text := make([]string, len(buttons)/2)
	res := make([]ResponseType, len(buttons)/2)
	for i := 0; i < len(text); i++ {
		btext, ok := buttons[2*i].(string)
		if !ok {
			argumentPanic("button text must be a string")
		}
		bresponse, ok := buttons[2*i+1].(ResponseType)
		if !ok {
			argumentPanic("button response must be an ResponseType")
		}
		text[i] = btext
		res[i] = bresponse
	}
	return text, res
}

func argumentPanic(message string) {
	if pc, _, _, ok := runtime.Caller(2); ok {
		log.Panicf("Arguments error: %s : %s\n",
			runtime.FuncForPC(pc).Name(), message)
	} else {
		log.Panicln("Arguments error: (unknown caller, see stack):", message)
	}
}

//-----------------------------------------------------------------------
// Main Loop and Events
//-----------------------------------------------------------------------

//Deprecated since 2.24. Use setlocale() directly.
//(see http://developer.gnome.org/gtk/stable/gtk-General.html#gtk-set-locale)
func SetLocale() {
	C.gtk_set_locale()
}

// gtk_disable_setlocale
// gtk_get_default_language
// gtk_parse_args

func Init(args *[]string) {
	if args != nil {
		argc := C.int(len(*args))
		argv := C.make_strings(argc)
		defer C.destroy_strings(argv)

		for i, arg := range *args {
			cstr := C.CString(arg)
			C.set_string(argv, C.int(i), (*C.gchar)(cstr))
		}

		C.gtk_init((*C.int)(unsafe.Pointer(&argc)),
			(***C.char)(unsafe.Pointer(&argv)))

		unhandled := make([]string, argc)
		for i := 0; i < int(argc); i++ {
			cstr := C.get_string(argv, C.int(i))
			unhandled[i] = C.GoString((*C.char)(cstr))
			C.free(unsafe.Pointer(cstr))
		}
		*args = unhandled
	} else {
		C.gtk_init(nil, nil)
	}
}

// gtk_init_check
// gtk_init_with_args
// gtk_get_option_group
// gtk_event_pending

func Main() {
	C.gtk_main()
}

// gtk_main_level

func MainQuit() {
	C.gtk_main_quit()
}

func MainIteration() bool {
	return gobool(C.gtk_main_iteration())
}

func MainIterationDo(blocking bool) bool {
	return gobool(C.gtk_main_iteration_do(gbool(blocking)))
}

func EventsPending() bool {
	return gobool(C.gtk_events_pending())
}

// gtk_main_do_event
// gtk_grab_add
// gtk_grab_get_current
// gtk_grab_remove
// gtk_key_snooper_install
// gtk_key_snooper_remove
// gtk_get_current_event

func GetCurrentEventTime() uint32 {
	return uint32(C.gtk_get_current_event_time())
}

// gtk_get_current_event_state
// gtk_get_event_widget
// gtk_propagate_event

//-----------------------------------------------------------------------
// GtkAccelGroup
//-----------------------------------------------------------------------
type AccelGroup struct {
	GAccelGroup *C.GtkAccelGroup
}

func NewAccelGroup() *AccelGroup {
	return &AccelGroup{C.gtk_accel_group_new()}
}

// gtk_accel_group_connect
// gtk_accel_group_connect_by_path
// gtk_accel_group_disconnect
// gtk_accel_group_disconnect_key
// gtk_accel_group_query
// gtk_accel_group_activate
// gtk_accel_group_lock
// gtk_accel_group_unlock
// gtk_accel_group_from_accel_closure
// gtk_accel_groups_activate
// gtk_accel_groups_from_object
// gtk_accel_group_find
// gtk_accelerator_valid

// Parse string representing an accelerator
// and return the key code and the modifier masks.
func AcceleratorParse(accelerator string) (uint, gdk.ModifierType) {
	ptrn := C.CString(accelerator)
	defer cfree(ptrn)

	var key C.guint
	var mods C.GdkModifierType
	C.gtk_accelerator_parse(gstring(ptrn), &key, &mods)
	return uint(key), gdk.ModifierType(mods)
}

// gtk_accelerator_name
// gtk_accelerator_get_label
// gtk_accelerator_set_default_mod_mask

func AcceleratorGetDefaultModMask() uint {
	return uint(C.gtk_accelerator_get_default_mod_mask())
}

//-----------------------------------------------------------------------
// GtkAccelMap
//-----------------------------------------------------------------------

// gtk_accel_map_add_entry
// gtk_accel_map_lookup_entry
// gtk_accel_map_change_entry
// gtk_accel_map_load
// gtk_accel_map_save
// gtk_accel_map_foreach
// gtk_accel_map_load_fd
// gtk_accel_map_save_fd
// gtk_accel_map_load_scanner
// gtk_accel_map_add_filter
// gtk_accel_map_foreach_unfiltered
// gtk_accel_map_get
// gtk_accel_map_lock_path
// gtk_accel_map_unlock_path

//-----------------------------------------------------------------------
// GtkClipboard
//-----------------------------------------------------------------------
type Clipboard struct {
	GClipboard *C.GtkClipboard
}

func NewClipboardGetForDisplay(display *gdk.Display, selection gdk.Atom) *Clipboard {
	var cdisplay unsafe.Pointer
	if display != nil {
		cdisplay = display.GDisplay
	}
	return &Clipboard{C._gtk_clipboard_get_for_display(cdisplay, unsafe.Pointer(uintptr(selection)))}
}

func (v *Clipboard) Clear() {
	C.gtk_clipboard_clear(v.GClipboard)
}

func (v *Clipboard) SetText(text string) {
	//ptr := C.toCstrV(unsafe.Pointer(&([]byte(text))[0])) // FIXME
	p := C.CString(text)
	defer cfree(p)
	C.gtk_clipboard_set_text(v.GClipboard, gstring(p), gint(-1))
}

func (v *Clipboard) SetImage(pixbuf *gdkpixbuf.Pixbuf) {
	C.gtk_clipboard_set_image(v.GClipboard, (*C.GdkPixbuf)(unsafe.Pointer(pixbuf.GPixbuf)))
}

func (v *Clipboard) Store() {
	C.gtk_clipboard_store(v.GClipboard)
}

func (v *Clipboard) WaitForImage() *gdkpixbuf.Pixbuf {
	gpixbuf := C.gtk_clipboard_wait_for_image(v.GClipboard)
	return &gdkpixbuf.Pixbuf{
		GdkPixbuf: gdkpixbuf.NewGdkPixbuf(unsafe.Pointer(gpixbuf)),
		GObject:   glib.ObjectFromNative(unsafe.Pointer(gpixbuf)),
	}
}

func (v *Clipboard) WaitForText() string {
	return gostring(C.gtk_clipboard_wait_for_text(v.GClipboard))
}

func (v *Clipboard) Connect(s string, f interface{}, datas ...interface{}) int {
	return glib.ObjectFromNative(unsafe.Pointer(C.toGObject(unsafe.Pointer(v.GClipboard)))).Connect(s, f, datas...)
}

// gtk_clipboard_get
// gtk_clipboard_set_with_data
// gtk_clipboard_set_with_owner
// gtk_clipboard_get_owner
// gtk_clipboard_set_image
// gtk_clipboard_request_contents
// gtk_clipboard_request_text
// gtk_clipboard_request_image
// gtk_clipboard_request_targets
// gtk_clipboard_request_rich_text
// gtk_clipboard_request_uris
// gtk_clipboard_wait_for_contents
// gtk_clipboard_wait_for_text
// gtk_clipboard_wait_for_image
// gtk_clipboard_wait_for_rich_text
// gtk_clipboard_wait_for_uris
// gtk_clipboard_wait_is_text_available
// gtk_clipboard_wait_is_image_available
// gtk_clipboard_wait_is_rich_text_available
// gtk_clipboard_wait_is_uris_available
// gtk_clipboard_wait_for_targets
// gtk_clipboard_wait_is_target_available
// gtk_clipboard_set_can_store

//-----------------------------------------------------------------------
// Drag and Drop
//-----------------------------------------------------------------------
type DestDefaults int

const (
	DEST_DEFAULT_MOTION    DestDefaults = 1 << iota /* respond to "drag_motion" */
	DEST_DEFAULT_HIGHLIGHT                          /* auto-highlight */
	DEST_DEFAULT_DROP                               /* respond to "drag_drop" */
	DEST_DEFAULT_ALL       = 0x07
)

type TargetEntry struct {
	Target string
	Flags  uint
	Info   uint
}

func (v *Widget) DragDestSet(flags DestDefaults, targets []TargetEntry, actions gdk.DragAction) {
	ctargets := make([]C.GtkTargetEntry, len(targets))
	for i, target := range targets {
		ptr := C.CString(target.Target)
		defer cfree(ptr)
		ctargets[i].target = gstring(ptr)
		ctargets[i].flags = guint(target.Flags)
		ctargets[i].info = guint(target.Info)
	}
	C.gtk_drag_dest_set(v.GWidget, C.GtkDestDefaults(flags), &ctargets[0], gint(len(targets)), C.GdkDragAction(actions))
}

func (v *Widget) DragSourceSet(start_button_mask gdk.ModifierType, targets []TargetEntry, actions gdk.DragAction) {
	ctargets := make([]C.GtkTargetEntry, len(targets))
	for i, target := range targets {
		ptr := C.CString(target.Target)
		defer cfree(ptr)
		ctargets[i].target = gstring(ptr)
		ctargets[i].flags = guint(target.Flags)
		ctargets[i].info = guint(target.Info)
	}
	C.gtk_drag_source_set(v.GWidget, C.GdkModifierType(start_button_mask), &ctargets[0], gint(len(targets)), C.GdkDragAction(actions))
}

func (v *Widget) DragFinish(context *gdk.DragContext, success bool, del bool, time uint32) {
	C._gtk_drag_finish(unsafe.Pointer(context.DragContext), gbool(success), gbool(del), guint32(time))
}

func (v *Widget) DragDestAddUriTargets() {
	C.gtk_drag_dest_add_uri_targets(v.GWidget)
}

// gtk_drag_dest_set_proxy
// gtk_drag_dest_unset
// gtk_drag_dest_find_target
// gtk_drag_dest_get_target_list
// gtk_drag_dest_set_target_list
// gtk_drag_dest_add_text_targets
// gtk_drag_dest_add_image_targets
// gtk_drag_dest_set_track_motion
// gtk_drag_dest_get_track_motion
// gtk_drag_get_data
// gtk_drag_get_source_widget
// gtk_drag_highlight
// gtk_drag_unhighlight
// gtk_drag_begin
// gtk_drag_set_icon_widget
// gtk_drag_set_icon_pixmap
// gtk_drag_set_icon_pixbuf
// gtk_drag_set_icon_stock
// gtk_drag_set_icon_name
// gtk_drag_set_icon_default
// gtk_drag_set_default_icon
// gtk_drag_check_threshold
// gtk_drag_source_set_icon
// gtk_drag_source_set_icon_pixbuf
// gtk_drag_source_set_icon_stock
// gtk_drag_source_set_icon_name
// gtk_drag_source_unset
// gtk_drag_source_set_target_list
// gtk_drag_source_get_target_list
// gtk_drag_source_add_text_targets
// gtk_drag_source_add_image_targets

//-----------------------------------------------------------------------
// GtkIconTheme
//-----------------------------------------------------------------------

// gtk_icon_theme_new
// gtk_icon_theme_get_default
// gtk_icon_theme_get_for_screen
// gtk_icon_theme_set_screen
// gtk_icon_theme_set_search_path
// gtk_icon_theme_get_search_path
// gtk_icon_theme_append_search_path
// gtk_icon_theme_prepend_search_path
// gtk_icon_theme_set_custom_theme
// gtk_icon_theme_has_icon
// gtk_icon_theme_lookup_icon
// gtk_icon_theme_choose_icon
// gtk_icon_theme_lookup_by_gicon
// gtk_icon_theme_load_icon
// gtk_icon_theme_list_contexts
// gtk_icon_theme_list_icons
// gtk_icon_theme_get_icon_sizes
// gtk_icon_theme_get_example_icon_name
// gtk_icon_theme_rescan_if_needed
// gtk_icon_theme_add_builtin_icon
// gtk_icon_info_copy
// gtk_icon_info_free
// gtk_icon_info_new_for_pixbuf
// gtk_icon_info_get_base_size
// gtk_icon_info_get_filename
// gtk_icon_info_get_builtin_pixbuf
// gtk_icon_info_load_icon
// gtk_icon_info_set_raw_coordinates
// gtk_icon_info_get_embedded_rect
// gtk_icon_info_get_attach_points
// gtk_icon_info_get_display_name

//-----------------------------------------------------------------------
// GtkStockItem
//-----------------------------------------------------------------------
const (
	STOCK_ABOUT                         = "gtk-about"
	STOCK_ADD                           = "gtk-add"
	STOCK_APPLY                         = "gtk-apply"
	STOCK_BOLD                          = "gtk-bold"
	STOCK_CANCEL                        = "gtk-cancel"
	STOCK_CAPS_LOCK_WARNING             = "gtk-caps-lock-warning"
	STOCK_CDROM                         = "gtk-cdrom"
	STOCK_CLEAR                         = "gtk-clear"
	STOCK_CLOSE                         = "gtk-close"
	STOCK_COLOR_PICKER                  = "gtk-color-picker"
	STOCK_CONVERT                       = "gtk-convert"
	STOCK_CONNECT                       = "gtk-connect"
	STOCK_COPY                          = "gtk-copy"
	STOCK_CUT                           = "gtk-cut"
	STOCK_DELETE                        = "gtk-delete"
	STOCK_DIALOG_AUTHENTICATION         = "gtk-dialog-authentication"
	STOCK_DIALOG_INFO                   = "gtk-dialog-info"
	STOCK_DIALOG_WARNING                = "gtk-dialog-warning"
	STOCK_DIALOG_ERROR                  = "gtk-dialog-error"
	STOCK_DIALOG_QUESTION               = "gtk-dialog-question"
	STOCK_DIRECTORY                     = "gtk-directory"
	STOCK_DISCARD                       = "gtk-discard"
	STOCK_DISCONNECT                    = "gtk-disconnect"
	STOCK_DND                           = "gtk-dnd"
	STOCK_DND_MULTIPLE                  = "gtk-dnd-multiple"
	STOCK_EDIT                          = "gtk-edit"
	STOCK_EXECUTE                       = "gtk-execute"
	STOCK_FILE                          = "gtk-file"
	STOCK_FIND                          = "gtk-find"
	STOCK_FIND_AND_REPLACE              = "gtk-find-and-replace"
	STOCK_FLOPPY                        = "gtk-floppy"
	STOCK_FULLSCREEN                    = "gtk-fullscreen"
	STOCK_GOTO_BOTTOM                   = "gtk-goto-bottom"
	STOCK_GOTO_FIRST                    = "gtk-goto-first"
	STOCK_GOTO_LAST                     = "gtk-goto-last"
	STOCK_GOTO_TOP                      = "gtk-goto-top"
	STOCK_GO_BACK                       = "gtk-go-back"
	STOCK_GO_DOWN                       = "gtk-go-down"
	STOCK_GO_FORWARD                    = "gtk-go-forward"
	STOCK_GO_UP                         = "gtk-go-up"
	STOCK_HARDDISK                      = "gtk-harddisk"
	STOCK_HELP                          = "gtk-help"
	STOCK_HOME                          = "gtk-home"
	STOCK_INDEX                         = "gtk-index"
	STOCK_INDENT                        = "gtk-indent"
	STOCK_INFO                          = "gtk-info"
	STOCK_UNINDENT                      = "gtk-unindent"
	STOCK_ITALIC                        = "gtk-italic"
	STOCK_JUMP_TO                       = "gtk-jump-to"
	STOCK_JUSTIFY_CENTER                = "gtk-justify-center"
	STOCK_JUSTIFY_FILL                  = "gtk-justify-fill"
	STOCK_JUSTIFY_LEFT                  = "gtk-justify-left"
	STOCK_JUSTIFY_RIGHT                 = "gtk-justify-right"
	STOCK_LEAVE_FULLSCREEN              = "gtk-leave-fullscreen"
	STOCK_MISSING_IMAGE                 = "gtk-missing-image"
	STOCK_MEDIA_FORWARD                 = "gtk-media-forward"
	STOCK_MEDIA_NEXT                    = "gtk-media-next"
	STOCK_MEDIA_PAUSE                   = "gtk-media-pause"
	STOCK_MEDIA_PLAY                    = "gtk-media-play"
	STOCK_MEDIA_PREVIOUS                = "gtk-media-previous"
	STOCK_MEDIA_RECORD                  = "gtk-media-record"
	STOCK_MEDIA_REWIND                  = "gtk-media-rewind"
	STOCK_MEDIA_STOP                    = "gtk-media-stop"
	STOCK_NETWORK                       = "gtk-network"
	STOCK_NEW                           = "gtk-new"
	STOCK_NO                            = "gtk-no"
	STOCK_OK                            = "gtk-ok"
	STOCK_OPEN                          = "gtk-open"
	STOCK_ORIENTATION_PORTRAIT          = "gtk-orientation-portrait"
	STOCK_ORIENTATION_LANDSCAPE         = "gtk-orientation-landscape"
	STOCK_ORIENTATION_REVERSE_LANDSCAPE = "gtk-orientation-reverse-landscape"
	STOCK_ORIENTATION_REVERSE_PORTRAIT  = "gtk-orientation-reverse-portrait"
	STOCK_PAGE_SETUP                    = "gtk-page-setup"
	STOCK_PASTE                         = "gtk-paste"
	STOCK_PREFERENCES                   = "gtk-preferences"
	STOCK_PRINT                         = "gtk-print"
	STOCK_PRINT_ERROR                   = "gtk-print-error"
	STOCK_PRINT_PAUSED                  = "gtk-print-paused"
	STOCK_PRINT_PREVIEW                 = "gtk-print-preview"
	STOCK_PRINT_REPORT                  = "gtk-print-report"
	STOCK_PRINT_WARNING                 = "gtk-print-warning"
	STOCK_PROPERTIES                    = "gtk-properties"
	STOCK_QUIT                          = "gtk-quit"
	STOCK_REDO                          = "gtk-redo"
	STOCK_REFRESH                       = "gtk-refresh"
	STOCK_REMOVE                        = "gtk-remove"
	STOCK_REVERT_TO_SAVED               = "gtk-revert-to-saved"
	STOCK_SAVE                          = "gtk-save"
	STOCK_SAVE_AS                       = "gtk-save-as"
	STOCK_SELECT_ALL                    = "gtk-select-all"
	STOCK_SELECT_COLOR                  = "gtk-select-color"
	STOCK_SELECT_FONT                   = "gtk-select-font"
	STOCK_SORT_ASCENDING                = "gtk-sort-ascending"
	STOCK_SORT_DESCENDING               = "gtk-sort-descending"
	STOCK_SPELL_CHECK                   = "gtk-spell-check"
	STOCK_STOP                          = "gtk-stop"
	STOCK_STRIKETHROUGH                 = "gtk-strikethrough"
	STOCK_UNDELETE                      = "gtk-undelete"
	STOCK_UNDERLINE                     = "gtk-underline"
	STOCK_UNDO                          = "gtk-undo"
	STOCK_YES                           = "gtk-yes"
	STOCK_ZOOM_100                      = "gtk-zoom-100"
	STOCK_ZOOM_FIT                      = "gtk-zoom-fit"
	STOCK_ZOOM_IN                       = "gtk-zoom-in"
	STOCK_ZOOM_OUT                      = "gtk-zoom-out"
)

type StockItem struct {
	GStockItem *C.GtkStockItem
}

func (v *StockItem) Add(nitems uint) {
	C.gtk_stock_add(v.GStockItem, guint(nitems))
}

func (v *StockItem) AddStatic(nitems uint) {
	C.gtk_stock_add_static(v.GStockItem, guint(nitems))
}

func StockLookup(stock_id string, item *StockItem) bool {
	ptr := C.CString(stock_id)
	defer cfree(ptr)
	return gobool(C.gtk_stock_lookup(gstring(ptr), item.GStockItem))
}

// gtk_stock_item_copy
// gtk_stock_item_free

func StockListIDs() *glib.SList {
	return glib.SListFromNative(unsafe.Pointer(C.gtk_stock_list_ids()))
}

//-----------------------------------------------------------------------
// Themeable Stock Images
//-----------------------------------------------------------------------

// gtk_icon_source_copy
// gtk_icon_source_free
// gtk_icon_factory_add
// gtk_icon_factory_add_default
// gtk_icon_factory_lookup
// gtk_icon_factory_lookup_default
// gtk_icon_factory_new
// gtk_icon_factory_remove_default
// gtk_icon_set_add_source
// gtk_icon_set_copy
// gtk_icon_set_new
// gtk_icon_set_new_from_pixbuf
// gtk_icon_set_ref
// gtk_icon_set_render_icon
// gtk_icon_set_unref
// gtk_icon_size_lookup
// gtk_icon_size_lookup_for_settings
// gtk_icon_size_register
// gtk_icon_size_register_alias
// gtk_icon_size_from_name
// gtk_icon_size_get_name
// gtk_icon_set_get_sizes
// gtk_icon_source_get_direction
// gtk_icon_source_get_direction_wildcarded
// gtk_icon_source_get_filename
// gtk_icon_source_get_pixbuf
// gtk_icon_source_get_icon_name
// gtk_icon_source_get_size
// gtk_icon_source_get_size_wildcarded
// gtk_icon_source_get_state
// gtk_icon_source_get_state_wildcarded
// gtk_icon_source_new
// gtk_icon_source_set_direction
// gtk_icon_source_set_direction_wildcarded
// gtk_icon_source_set_filename
// gtk_icon_source_set_pixbuf
// gtk_icon_source_set_icon_name
// gtk_icon_source_set_size
// gtk_icon_source_set_size_wildcarded
// gtk_icon_source_set_state
// gtk_icon_source_set_state_wildcarded

//-----------------------------------------------------------------------
// Resource Files
//-----------------------------------------------------------------------

// gtk_rc_scanner_new
// gtk_rc_get_style
// gtk_rc_get_style_by_paths
// gtk_rc_add_widget_name_style
// gtk_rc_add_widget_class_style
// gtk_rc_add_class_style

func RCParse(file string) {
	ptr := C.CString(file)
	defer cfree(ptr)
	C.gtk_rc_parse((*C.gchar)(ptr))
}

func RCParseString(style string) {
	ptr := C.CString(style)
	defer cfree(ptr)
	C.gtk_rc_parse_string((*C.gchar)(ptr))
}

func RCReparseAll() bool {
	return gobool(C.gtk_rc_reparse_all())
}

// gtk_rc_reparse_all_for_settings

func RCResetStyles(settings *Settings) {
	C.gtk_rc_reset_styles(settings.GSettings)
}

// gtk_rc_add_default_file
// gtk_rc_get_default_files
// gtk_rc_set_default_files
// gtk_rc_parse_color
// gtk_rc_parse_color_full
// gtk_rc_parse_state
// gtk_rc_parse_priority
// gtk_rc_find_module_in_path
// gtk_rc_find_pixmap_in_path
// gtk_rc_get_module_dir
// gtk_rc_get_im_module_path
// gtk_rc_get_im_module_file
// gtk_rc_get_theme_dir
// gtk_rc_style_new
// gtk_rc_style_copy
// gtk_rc_style_ref
// gtk_rc_style_unref

//-----------------------------------------------------------------------
// GtkSettings
//-----------------------------------------------------------------------
type Settings struct {
	GSettings *C.GtkSettings
}

// gtk_settings_get_for_screen
func GetSettingsForScreen(screen gdk.Screen) *Settings {
	gScreen := C.toGdkScreen(unsafe.Pointer(screen.GScreen))
	return &Settings{GSettings: C.gtk_settings_get_for_screen(gScreen)}
}

// gtk_settings_install_property
// gtk_settings_install_property_parser
// gtk_rc_property_parse_color
// gtk_rc_property_parse_enum
// gtk_rc_property_parse_flags
// gtk_rc_property_parse_requisition
// gtk_rc_property_parse_border
// gtk_settings_set_property_value

func (s *Settings) ToGObject() *glib.GObject {
	return &glib.GObject{unsafe.Pointer(s.GSettings)}
}

func SettingsGetDefault() *Settings {
	return &Settings{C.gtk_settings_get_default()}
}

func (s *Settings) SetStringProperty(name string, v_string string, origin string) {
	ptrn := C.CString(name)
	defer cfree(ptrn)
	ptrv := C.CString(v_string)
	defer cfree(ptrv)
	prts := C.CString(origin)
	defer cfree(prts)
	C.gtk_settings_set_string_property(s.GSettings, gstring(ptrn), gstring(ptrv), gstring(prts))
}

func (s *Settings) SetLongProperty(name string, v_long int32, origin string) {
	ptrn := C.CString(name)
	defer cfree(ptrn)
	prts := C.CString(origin)
	defer cfree(prts)
	C.gtk_settings_set_long_property(s.GSettings, gstring(ptrn), glong(v_long), gstring(prts))
}

func (s *Settings) SetDoubleProperty(name string, v_double float64, origin string) {
	ptrn := C.CString(name)
	defer cfree(ptrn)
	prts := C.CString(origin)
	defer cfree(prts)
	C.gtk_settings_set_double_property(s.GSettings, gstring(ptrn), gdouble(v_double), gstring(prts))
}

//-----------------------------------------------------------------------
// GtkBinding
//-----------------------------------------------------------------------

// gtk_binding_entry_add
// gtk_binding_entry_add_signall
// gtk_binding_entry_clear
// gtk_binding_parse_binding
// gtk_binding_set_new
// gtk_binding_set_by_class
// gtk_binding_set_find
// gtk_bindings_activate
// gtk_bindings_activate_event
// gtk_binding_set_activate
// gtk_binding_entry_add_signal
// gtk_binding_entry_skip
// gtk_binding_entry_remove
// gtk_binding_set_add_path

//-----------------------------------------------------------------------
// Graphics Contexts
//-----------------------------------------------------------------------

// gtk_gc_get
// gtk_gc_release

//-----------------------------------------------------------------------
// GtkStyle
//-----------------------------------------------------------------------

type Style struct {
	GStyle *C.GtkStyle
}

func NewStyle() *Style {
	return &Style{C.gtk_style_new()}
}

func (v *Style) Copy() *Style {
	return &Style{C.gtk_style_copy(STYLE(v))}
}

func (v *Style) Attach(window *Window) *Style {
	return &Style{C.gtk_style_attach(STYLE(v), C.toGdkWindow(unsafe.Pointer(window)))}
}

func (v *Style) Detach() {
	C.gtk_style_detach(STYLE(v))
}

func (v *Style) SetBackground(window *Window, state StateType) {
	C.gtk_style_set_background(STYLE(v), C.toGdkWindow(unsafe.Pointer(window)), C.GtkStateType(state))
}

// gtk_style_apply_default_background

func (v *Style) LookupColor(colorName string) (*gdk.Color, bool) {
	color_name := C.CString(colorName)
	defer cfree(color_name)
	color := new(gdk.Color)
	b := C.gtk_style_lookup_color(v.GStyle, gstring(color_name), (*C.GdkColor)(unsafe.Pointer(&color.GColor)))
	return color, gobool(b)
}

// gtk_style_lookup_icon_set
// gtk_style_render_icon
// gtk_style_get_style_property

// gtk_style_get_valist
// gtk_style_get
// gtk_paint_arrow
// gtk_paint_box
// gtk_paint_box_gap
// gtk_paint_check
// gtk_paint_diamond
// gtk_paint_extension
// gtk_paint_flat_box
// gtk_paint_focus
// gtk_paint_handle
// gtk_paint_hline
// gtk_paint_option
// gtk_paint_polygon
// gtk_paint_shadow
// gtk_paint_shadow_gap
// gtk_paint_slider
// gtk_paint_spinner
// gtk_paint_tab
// gtk_paint_vline
// gtk_paint_expander
// gtk_paint_layout
// gtk_paint_resize_grip
// gtk_draw_insertion_cursor
// gtk_border_new
// gtk_border_copy
// gtk_border_free

// Deprecated:
// gtk_style_ref
// gtk_style_unref
// gtk_style_apply_default_pixmap
// gtk_style_get_font
// gtk_style_set_font
// gtk_draw_hline
// gtk_draw_vline
// gtk_draw_shadow
// gtk_draw_polygon
// gtk_draw_arrow
// gtk_draw_diamond
// gtk_draw_string
// gtk_draw_box
// gtk_draw_box_gap
// gtk_draw_check
// gtk_draw_extension
// gtk_draw_flat_box
// gtk_draw_focus
// gtk_draw_handle
// gtk_draw_option
// gtk_draw_shadow_gap
// gtk_draw_slider
// gtk_draw_tab
// gtk_draw_expander
// gtk_draw_layout
// gtk_draw_resize_grip
// gtk_paint_string

//-----------------------------------------------------------------------
// Selections
//-----------------------------------------------------------------------

type SelectionData struct {
	GSelectionData unsafe.Pointer
}

func NewSelectionDataFromNative(l unsafe.Pointer) *SelectionData {
	return &SelectionData{l}
}

func (v *SelectionData) GetLength() int {
	return int(C._gtk_selection_data_get_length(v.GSelectionData))
}

func (v *SelectionData) GetData() unsafe.Pointer {
	return unsafe.Pointer(C._gtk_selection_data_get_data(v.GSelectionData))
}

func (v *SelectionData) GetText() string {
	return C.GoString(C._gtk_selection_data_get_text(v.GSelectionData))
}

// gtk_target_list_new
// gtk_target_list_ref
// gtk_target_list_unref
// gtk_target_list_add
// gtk_target_list_add_table
// gtk_target_list_add_text_targets
// gtk_target_list_add_image_targets
// gtk_target_list_add_uri_targets
// gtk_target_list_add_rich_text_targets
// gtk_target_list_remove
// gtk_target_list_find
// gtk_target_table_free
// gtk_target_table_new_from_list
// gtk_selection_owner_set
// gtk_selection_owner_set_for_display
// gtk_selection_add_target
// gtk_selection_add_targets
// gtk_selection_clear_targets
// gtk_selection_convert
// gtk_selection_data_set
// gtk_selection_data_set_text
// gtk_selection_data_set_pixbuf
// gtk_selection_data_get_pixbuf
// gtk_selection_data_set_uris
// gtk_selection_data_get_uris
// gtk_selection_data_get_targets
// gtk_selection_data_targets_include_image
// gtk_selection_data_targets_include_text
// gtk_selection_data_targets_include_uri
// gtk_selection_data_targets_include_rich_text
// gtk_selection_data_get_selection
// gtk_selection_data_get_data_type
// gtk_selection_data_get_display
// gtk_selection_data_get_format
// gtk_selection_data_get_target
// gtk_targets_include_image
// gtk_targets_include_text
// gtk_targets_include_uri
// gtk_targets_include_rich_text
// gtk_selection_remove_all
// gtk_selection_clear
// gtk_selection_data_copy
// gtk_selection_data_free

//-----------------------------------------------------------------------
// Version Information
//-----------------------------------------------------------------------

// gtk_major_version
// gtk_minor_version
// gtk_micro_version
// gtk_binary_age
// gtk_interface_age
// gtk_check_version

//-----------------------------------------------------------------------
// Testing
//-----------------------------------------------------------------------

// gtk_test_create_simple_window
// gtk_test_create_widget
// gtk_test_display_button_window
// gtk_test_find_label
// gtk_test_find_sibling
// gtk_test_find_widget
// gtk_test_init
// gtk_test_list_all_types
// gtk_test_register_all_types
// gtk_test_slider_get_value
// gtk_test_slider_set_perc
// gtk_test_spin_button_click
// gtk_test_text_get
// gtk_test_text_set
// gtk_test_widget_click
// gtk_test_widget_send_key

//-----------------------------------------------------------------------
// Filesystem Utilities
//-----------------------------------------------------------------------
// gtk_mount_operation_new
// gtk_mount_operation_is_showing
// gtk_mount_operation_set_parent
// gtk_mount_operation_get_parent
// gtk_mount_operation_set_screen
// gtk_mount_operation_get_screen
// gtk_show_uri

//-----------------------------------------------------------------------
// GtkDialog
//-----------------------------------------------------------------------
type DialogFlags int

const (
	DIALOG_MODAL               DialogFlags = 1 << iota /* call gtk_window_set_modal (win, TRUE) */
	DIALOG_DESTROY_WITH_PARENT                         /* call gtk_window_set_destroy_with_parent () */
	DIALOG_NO_SEPARATOR                                /* no separator bar above buttons */
)

type ResponseType int

const (
	RESPONSE_NONE ResponseType = -iota - 1
	RESPONSE_REJECT
	RESPONSE_ACCEPT
	RESPONSE_DELETE_EVENT
	RESPONSE_OK
	RESPONSE_CANCEL
	RESPONSE_CLOSE
	RESPONSE_YES
	RESPONSE_NO
	RESPONSE_APPLY
	RESPONSE_HELP
)

/*type IDialog interface {
	IWidget
	Run() int
	Response(interface{}, ...interface{})
}*/

type Dialog struct {
	Window
	wfr *Widget // WidgetForResponse
}

func NewDialog() *Dialog {
	return &Dialog{Window{Bin{Container{Widget{C.gtk_dialog_new()}}}}, nil}
}

// gtk_dialog_new_with_buttons

func (v *Dialog) Run() ResponseType {
	return ResponseType(C.gtk_dialog_run(DIALOG(v)))
}

func (v *Dialog) Response(response interface{}, datas ...interface{}) {
	if id, ok := response.(ResponseType); ok {
		C.gtk_dialog_response(DIALOG(v), gint(int(id)))
		return
	}
	v.Connect("response", response, datas...)
}

func (v *Dialog) AddButton(button_text string, response_id ResponseType) *Button {
	ptr := C.CString(button_text)
	defer cfree(ptr)
	return newButtonInternal(C.gtk_dialog_add_button(
		DIALOG(v), gstring(ptr), gint(int(response_id))))
}

// gtk_dialog_add_buttons
// gtk_dialog_add_action_widget
// gtk_dialog_get_has_separator //deprecated since 2.22

//Deprecated since 2.22.
func (v *Dialog) SetHasSeparator(f bool) {
	deprecated_since(2, 22, 0, "gtk_dialog_set_has_separator()")
	C.gtk_dialog_set_has_separator(DIALOG(v), gbool(f))
}

func (v *Dialog) SetDefaultResponse(id ResponseType) {
	C.gtk_dialog_set_default_response(DIALOG(v), gint(int(id)))
}

// gtk_dialog_set_has_separator //deprecated since 2.22
// gtk_dialog_set_response_sensitive

func (v *Dialog) GetResponseForWidget(w *Widget) ResponseType {
	return ResponseType(int(C.gtk_dialog_get_response_for_widget(DIALOG(v), w.GWidget)))
}

func (v *Dialog) GetWidgetForResponse(id int) *Widget {
	panic_if_version_older(2, 20, 0, "gtk_dialog_get_widget_for_response()")
	w := C._gtk_dialog_get_widget_for_response(DIALOG(v), gint(id))
	if v.wfr == nil {
		v.wfr = &Widget{w}
	} else {
		v.wfr.GWidget = w
	}
	return v.wfr
}

// gtk_dialog_get_action_area
// gtk_dialog_get_content_area
// gtk_alternative_dialog_button_order
// gtk_dialog_set_alternative_button_order
// gtk_dialog_set_alternative_button_order_from_array

func (v *Dialog) GetVBox() *VBox {
	return &VBox{Box{Container{Widget{C._gtk_dialog_get_vbox(v.GWidget)}}}}
}

//-----------------------------------------------------------------------
// GtkMessageDialog
//-----------------------------------------------------------------------
type MessageType int

const (
	MESSAGE_INFO MessageType = iota
	MESSAGE_WARNING
	MESSAGE_QUESTION
	MESSAGE_ERROR
	MESSAGE_OTHER
)

type ButtonsType int

const (
	BUTTONS_NONE ButtonsType = iota
	BUTTONS_OK
	BUTTONS_CLOSE
	BUTTONS_CANCEL
	BUTTONS_YES_NO
	BUTTONS_OK_CANCEL
)

type MessageDialog struct {
	Dialog
}

// TODO should be variadic function
func NewMessageDialog(parent *Window, flag DialogFlags, t MessageType, buttons ButtonsType, format string, args ...interface{}) *MessageDialog {
	ptr := C.CString(strings.Replace(fmt.Sprintf(format, args...), "%", "%%", -1))
	defer cfree(ptr)
	return &MessageDialog{Dialog{Window{Bin{Container{Widget{
		C._gtk_message_dialog_new(
			ToNative(parent),
			C.GtkDialogFlags(flag),
			C.GtkMessageType(t),
			C.GtkButtonsType(buttons),
			gstring(ptr))}}}}, nil}}
}

func NewMessageDialogWithMarkup(parent *Window, flag DialogFlags, t MessageType, buttons ButtonsType, format string, args ...interface{}) *MessageDialog {
	r := &MessageDialog{Dialog{Window{Bin{Container{Widget{
		C._gtk_message_dialog_new_with_markup(
			ToNative(parent),
			C.GtkDialogFlags(flag),
			C.GtkMessageType(t),
			C.GtkButtonsType(buttons),
			nil)}}}}, nil}}
	r.SetMarkup(fmt.Sprintf(format, args...))
	return r
}

func (v *MessageDialog) SetMarkup(markup string) {
	ptr := C.CString(markup)
	defer cfree(ptr)
	C.gtk_message_dialog_set_markup(MESSAGE_DIALOG(v), gstring(ptr))
}

func (v *MessageDialog) SetImage(image IWidget) {
	C.gtk_message_dialog_set_image(MESSAGE_DIALOG(v), ToNative(image))
}

func (v *MessageDialog) GetImage() *Image {
	return &Image{Misc{Widget{C.gtk_message_dialog_get_image(MESSAGE_DIALOG(v))}}}
}

// gtk_message_dialog_get_message_area //since 2.22
// gtk_message_dialog_format_secondary_text
// gtk_message_dialog_format_secondary_markup

//-----------------------------------------------------------------------
// GtkWindow
//-----------------------------------------------------------------------
type WindowType int

const (
	WINDOW_TOPLEVEL WindowType = iota
	WINDOW_POPUP
)

type WindowPosition int

const (
	WIN_POS_NONE WindowPosition = iota
	WIN_POS_CENTER
	WIN_POS_MOUSE
	WIN_POS_CENTER_ALWAYS
	WIN_POS_CENTER_ON_PARENT
)

/*type IWindow interface {
	IContainer
	SetTransientFor(parent IWindow)
	GetTitle() string
	SetTitle(title string)
}*/

type Window struct {
	Bin
}

func NewWindow(t WindowType) *Window {
	return &Window{Bin{Container{Widget{
		C.gtk_window_new(C.GtkWindowType(t))}}}}
}

func (v *Window) SetTitle(title string) {
	ptr := C.CString(title)
	defer cfree(ptr)
	C.gtk_window_set_title(WINDOW(v), gstring(ptr))
}

func (v *Window) SetWMClass(name string, class string) {
	ptr1 := C.CString(name)
	defer cfree(ptr1)
	ptr2 := C.CString(class)
	defer cfree(ptr2)
	C.gtk_window_set_wmclass(WINDOW(v), gstring(ptr1), gstring(ptr2))
}

func (v *Window) SetResizable(resizable bool) {
	C.gtk_window_set_resizable(WINDOW(v), gbool(resizable))
}

func (v *Window) GetResizable() bool {
	return gobool(C.gtk_window_get_resizable(WINDOW(v)))
}

func (v *Window) AddAccelGroup(agroup *AccelGroup) {
	C.gtk_window_add_accel_group(WINDOW(v), agroup.GAccelGroup)
}

func (v *Window) RemoveAccelGroup(agroup *AccelGroup) {
	C.gtk_window_remove_accel_group(WINDOW(v), agroup.GAccelGroup)
}

func (v *Window) ActivateFocus() bool {
	return gobool(C.gtk_window_activate_focus(WINDOW(v)))
}

func (v *Window) ActivateDefault() bool {
	return gobool(C.gtk_window_activate_default(WINDOW(v)))
}

func (v *Window) SetModal(modal bool) {
	C.gtk_window_set_modal(WINDOW(v), gbool(modal))
}

func (v *Window) SetDefaultSize(width int, height int) {
	C.gtk_window_set_default_size(WINDOW(v), gint(width), gint(height))
}

func (v *Window) SetGravity(gravity gdk.Gravity) {
	C.gtk_window_set_gravity(WINDOW(v), C.GdkGravity(gravity))
}

// gtk_window_set_geometry_hints
// gtk_window_get_gravity

func (v *Window) SetPosition(position WindowPosition) {
	C.gtk_window_set_position(WINDOW(v), C.GtkWindowPosition(position))
}

func (v *Window) SetTransientFor(parent *Window) {
	C.gtk_window_set_transient_for(WINDOW(v), WINDOW(parent))
}

func (v *Window) SetDestroyWithParent(setting bool) {
	C.gtk_window_set_destroy_with_parent(WINDOW(v), gbool(setting))
}

func (v *Window) SetScreen(screen *gdk.Screen) {
	gScreen := C.toGdkScreen(unsafe.Pointer(screen.GScreen))
	C.gtk_window_set_screen(WINDOW(v), gScreen)
}

func (v *Window) GetScreen() *gdk.Screen {
	return gdk.ScreenFromUnsafe(unsafe.Pointer(C.gtk_window_get_screen(WINDOW(v))))
}

func (v *Window) IsActive() bool {
	return gobool(C.gtk_window_is_active(WINDOW(v)))
}

func (v *Window) HasToplevelFocus() bool {
	return gobool(C.gtk_window_has_toplevel_focus(WINDOW(v)))
}

func (v *Window) ListToplevels() *glib.List {
	return glib.ListFromNative(unsafe.Pointer(C.gtk_window_list_toplevels()))
}

// gtk_window_add_mnemonic
// gtk_window_remove_mnemonic
// gtk_window_mnemonic_activate
// gtk_window_activate_key
// gtk_window_propagate_key_event
// gtk_window_get_focus
// gtk_window_set_focus
// gtk_window_get_default_widget

func (v *Window) SetDefault(w *Widget) {
	C.gtk_window_set_default(WINDOW(v), w.GWidget)
}

func (v *Window) Present() {
	C.gtk_window_present(WINDOW(v))
}

// gtk_window_present_with_time

func (v *Window) Stick() {
	C.gtk_window_stick(WINDOW(v))
}

func (v *Window) Unstick() {
	C.gtk_window_unstick(WINDOW(v))
}

func (v *Window) Iconify() {
	C.gtk_window_iconify(WINDOW(v))
}

func (v *Window) Deiconify() {
	C.gtk_window_deiconify(WINDOW(v))
}

func (v *Window) Maximize() {
	C.gtk_window_maximize(WINDOW(v))
}

func (v *Window) Unmaximize() {
	C.gtk_window_unmaximize(WINDOW(v))
}

func (v *Window) Fullscreen() {
	C.gtk_window_fullscreen(WINDOW(v))
}

func (v *Window) Unfullscreen() {
	C.gtk_window_unfullscreen(WINDOW(v))
}

func (v *Window) SetKeepAbove(setting bool) {
	C.gtk_window_set_keep_above(WINDOW(v), gbool(setting))
}

func (v *Window) SetKeepBelow(setting bool) {
	C.gtk_window_set_keep_below(WINDOW(v), gbool(setting))
}

func (v *Window) SetDecorated(setting bool) {
	C.gtk_window_set_decorated(WINDOW(v), gbool(setting))
}

func (v *Window) SetDeletable(setting bool) {
	C.gtk_window_set_deletable(WINDOW(v), gbool(setting))
}

func (v *Window) SetTypeHint(hint gdk.WindowTypeHint) {
	C.gtk_window_set_type_hint(WINDOW(v), C.GdkWindowTypeHint(hint))
}

// gtk_window_begin_resize_drag
// gtk_window_begin_move_drag
// gtk_window_set_frame_dimensions //deprecated since 2.24
// gtk_window_set_has_frame  //deprecated since 2.24
// gtk_window_set_mnemonic_modifier
// gtk_window_set_skip_pager_hint
// gtk_window_set_urgency_hint

func (v *Window) SetAcceptFocus(setting bool) {
	C.gtk_window_set_accept_focus(WINDOW(v), gbool(setting))
}

func (v *Window) SetSkipTaskbarHint(setting bool) {
	C.gtk_window_set_skip_taskbar_hint(WINDOW(v), gbool(setting))
}

// gtk_window_set_focus_on_map
// gtk_window_set_startup_id
// gtk_window_set_role
// gtk_window_get_decorated
// gtk_window_get_deletable
// gtk_window_get_default_icon_list
// gtk_window_get_default_icon_name

func (v *Window) GetDefaultSize(width *int, height *int) {
	var cwidth, cheight C.gint
	C.gtk_window_get_default_size(WINDOW(v), &cwidth, &cheight)
	*width = int(cwidth)
	*height = int(cheight)
}

func (v *Window) GetDestroyWithParent() bool {
	return gobool(C.gtk_window_get_destroy_with_parent(WINDOW(v)))
}

// gtk_window_get_frame_dimensions //deprecated since 2.24
// gtk_window_get_has_frame  //deprecated since 2.24
// gtk_window_get_icon
// gtk_window_get_icon_list

func (v *Window) GetIconName() string {
	return gostring(C.gtk_window_get_icon_name(WINDOW(v)))
}

// gtk_window_get_mnemonic_modifier

func (v *Window) GetModal() bool {
	return gobool(C.gtk_window_get_modal(WINDOW(v)))
}

func (v *Window) GetPosition(root_x *int, root_y *int) {
	var croot_x, croot_y C.gint
	C.gtk_window_get_position(WINDOW(v), &croot_x, &croot_y)
	*root_x = int(croot_x)
	*root_y = int(croot_y)
}

// gtk_window_get_role

func (v *Window) GetSize(width *int, height *int) {
	var cwidth, cheight C.gint
	C.gtk_window_get_size(WINDOW(v), &cwidth, &cheight)
	*width = int(cwidth)
	*height = int(cheight)
}

func (v *Window) GetTitle() string {
	return gostring(C.gtk_window_get_title(WINDOW(v)))
}

func (v *Window) GetTypeHint() gdk.WindowTypeHint {
	return gdk.WindowTypeHint(C.gtk_window_get_type_hint(WINDOW(v)))
}

// gtk_window_get_transient_for
// gtk_window_get_skip_pager_hint
// gtk_window_get_urgency_hint

func (v *Window) GetAcceptFocus() bool {
	return gobool(C.gtk_window_get_accept_focus(WINDOW(v)))
}

func (v *Window) GetSkipTaskbarHint() bool {
	return gobool(C.gtk_window_get_skip_taskbar_hint(WINDOW(v)))
}

// gtk_window_get_focus_on_map
// gtk_window_get_group
// gtk_window_has_group //since 2.22
// gtk_window_get_window_type //since 2.20

func (v *Window) Move(x int, y int) {
	C.gtk_window_move(WINDOW(v), gint(x), gint(y))
}

// gtk_window_parse_geometry
// gtk_window_reshow_with_initial_size

func (v *Window) Resize(width int, height int) {
	C.gtk_window_resize(WINDOW(v), gint(width), gint(height))
}

// gtk_window_set_default_icon_list
// gtk_window_set_default_icon
// gtk_window_set_default_icon_from_file
// gtk_window_set_default_icon_name
func (v *Window) SetIcon(icon *gdkpixbuf.Pixbuf) {
	C.gtk_window_set_icon(WINDOW(v), (*C.GdkPixbuf)(unsafe.Pointer(icon.GPixbuf)))
}

// gtk_window_set_icon_list

func (v *Window) SetIconFromFile(file string) {
	ptr := C.CString(file)
	defer cfree(ptr)
	C.gtk_window_set_icon_from_file(WINDOW(v), gstring(ptr), nil) // last arg : GError **err
}

func (v *Window) SetIconName(name string) {
	ptr := C.CString(name)
	defer cfree(ptr)
	C.gtk_window_set_icon_name(WINDOW(v), gstring(ptr))
}

func (v *Window) SetAutoStartupNotification(setting bool) {
	C.gtk_window_set_auto_startup_notification(gbool(setting))
}

func (v *Window) GetOpacity() float64 {
	return float64(C.gtk_window_get_opacity(WINDOW(v)))
}

func (v *Window) SetOpacity(opacity float64) {
	C.gtk_window_set_opacity(WINDOW(v), gdouble(opacity))
}

// gtk_window_get_mnemonics_visible //since 2.20
// gtk_window_set_mnemonics_visible //since 2.20

//-----------------------------------------------------------------------
// GtkWindowGroup
//-----------------------------------------------------------------------

// gtk_window_group_new
// gtk_window_group_add_window
// gtk_window_group_remove_window
// gtk_window_group_list_windows
// gtk_window_group_get_current_grab

//-----------------------------------------------------------------------
// GtkAboutDialog
//-----------------------------------------------------------------------
type AboutDialog struct {
	Dialog
}

func NewAboutDialog() *AboutDialog {
	return &AboutDialog{Dialog{Window{Bin{Container{Widget{C.gtk_about_dialog_new()}}}}, nil}}
}

func (v *AboutDialog) GetProgramName() string {
	return gostring(C.gtk_about_dialog_get_program_name(ABOUT_DIALOG(v)))
}

func (v *AboutDialog) SetProgramName(name string) {
	ptr := C.CString(name)
	defer cfree(ptr)
	C.gtk_about_dialog_set_program_name(ABOUT_DIALOG(v), gstring(ptr))
}

func (v *AboutDialog) GetVersion() string {
	return gostring(C.gtk_about_dialog_get_version(ABOUT_DIALOG(v)))
}

func (v *AboutDialog) SetVersion(version string) {
	ptr := C.CString(version)
	defer cfree(ptr)
	C.gtk_about_dialog_set_version(ABOUT_DIALOG(v), gstring(ptr))
}

func (v *AboutDialog) GetCopyright() string {
	return gostring(C.gtk_about_dialog_get_copyright(ABOUT_DIALOG(v)))
}

func (v *AboutDialog) SetCopyright(copyright string) {
	ptr := C.CString(copyright)
	defer cfree(ptr)
	C.gtk_about_dialog_set_copyright(ABOUT_DIALOG(v), gstring(ptr))
}

func (v *AboutDialog) GetComments() string {
	return gostring(C.gtk_about_dialog_get_comments(ABOUT_DIALOG(v)))
}

func (v *AboutDialog) SetComments(comments string) {
	ptr := C.CString(comments)
	defer cfree(ptr)
	C.gtk_about_dialog_set_comments(ABOUT_DIALOG(v), gstring(ptr))
}

func (v *AboutDialog) GetLicense() string {
	return gostring(C.gtk_about_dialog_get_license(ABOUT_DIALOG(v)))
}

func (v *AboutDialog) SetLicense(license string) {
	ptr := C.CString(license)
	defer cfree(ptr)
	C.gtk_about_dialog_set_license(ABOUT_DIALOG(v), gstring(ptr))
}

func (v *AboutDialog) GetWrapLicense() bool {
	return gobool(C.gtk_about_dialog_get_wrap_license(ABOUT_DIALOG(v)))
}

func (v *AboutDialog) SetWrapLicense(wrap_license bool) {
	C.gtk_about_dialog_set_wrap_license(ABOUT_DIALOG(v), gbool(wrap_license))
}

func (v *AboutDialog) GetWebsite() string {
	return gostring(C.gtk_about_dialog_get_website(ABOUT_DIALOG(v)))
}

func (v *AboutDialog) SetWebsite(website string) {
	ptr := C.CString(website)
	defer cfree(ptr)
	C.gtk_about_dialog_set_website(ABOUT_DIALOG(v), gstring(ptr))
}

func (v *AboutDialog) GetWebsiteLabel() string {
	return gostring(C.gtk_about_dialog_get_website_label(ABOUT_DIALOG(v)))
}

func (v *AboutDialog) SetWebsiteLabel(website_label string) {
	var ptr *C.char
	if len(website_label) > 0 {
		ptr = C.CString(website_label)
		defer cfree(ptr)
	}
	C.gtk_about_dialog_set_website_label(ABOUT_DIALOG(v), gstring(ptr))
}

func (v *AboutDialog) GetAuthors() []string {
	var authors []string
	cauthors := C.gtk_about_dialog_get_authors(ABOUT_DIALOG(v))
	for {
		authors = append(authors, gostring(*cauthors))
		cauthors = C.nextGstr(cauthors)
		if *cauthors == nil {
			break
		}
	}
	return authors
}

func (v *AboutDialog) SetAuthors(authors []string) {
	cauthors := C.make_strings(C.int(len(authors) + 1))
	for i, author := range authors {
		ptr := C.CString(author)
		defer cfree(ptr)
		C.set_string(cauthors, C.int(i), gstring(ptr))
	}
	C.set_string(cauthors, C.int(len(authors)), nil)
	C.gtk_about_dialog_set_authors(ABOUT_DIALOG(v), cauthors)
	C.destroy_strings(cauthors)
}

func (v *AboutDialog) GetArtists() []string {
	var artists []string
	cartists := C.gtk_about_dialog_get_artists(ABOUT_DIALOG(v))
	for {
		artists = append(artists, gostring(*cartists))
		cartists = C.nextGstr(cartists)
		if *cartists == nil {
			break
		}
	}
	return artists
}

func (v *AboutDialog) SetArtists(artists []string) {
	cartists := C.make_strings(C.int(len(artists)) + 1)
	for i, author := range artists {
		ptr := C.CString(author)
		defer cfree(ptr)
		C.set_string(cartists, C.int(i), gstring(ptr))
	}
	C.set_string(cartists, C.int(len(artists)), nil)
	C.gtk_about_dialog_set_artists(ABOUT_DIALOG(v), cartists)
	C.destroy_strings(cartists)
}

func (v *AboutDialog) GetDocumenters() []string {
	var documenters []string
	cdocumenters := C.gtk_about_dialog_get_documenters(ABOUT_DIALOG(v))
	for {
		documenters = append(documenters, gostring(*cdocumenters))
		cdocumenters = C.nextGstr(cdocumenters)
		if *cdocumenters == nil {
			break
		}
	}
	return documenters
}

func (v *AboutDialog) SetDocumenters(documenters []string) {
	cdocumenters := C.make_strings(C.int(len(documenters)) + 1)
	for i, author := range documenters {
		ptr := C.CString(author)
		defer cfree(ptr)
		C.set_string(cdocumenters, C.int(i), gstring(ptr))
	}
	C.set_string(cdocumenters, C.int(len(documenters)), nil)
	C.gtk_about_dialog_set_documenters(ABOUT_DIALOG(v), cdocumenters)
	C.destroy_strings(cdocumenters)
}

func (v *AboutDialog) GetTranslatorCredits() string {
	return gostring(C.gtk_about_dialog_get_translator_credits(ABOUT_DIALOG(v)))
}

func (v *AboutDialog) SetTranslatorCredits(translator_credits string) {
	ptr := C.CString(translator_credits)
	defer cfree(ptr)
	C.gtk_about_dialog_set_translator_credits(ABOUT_DIALOG(v), gstring(ptr))
}

func (v *AboutDialog) GetLogo() *gdkpixbuf.Pixbuf {
	gpixbuf := C.gtk_about_dialog_get_logo(ABOUT_DIALOG(v))
	return &gdkpixbuf.Pixbuf{
		GdkPixbuf: gdkpixbuf.NewGdkPixbuf(unsafe.Pointer(gpixbuf)),
		GObject:   glib.ObjectFromNative(unsafe.Pointer(gpixbuf)),
	}
}

func (v *AboutDialog) SetLogo(logo *gdkpixbuf.Pixbuf) {
	C.gtk_about_dialog_set_logo(ABOUT_DIALOG(v), (*C.GdkPixbuf)(unsafe.Pointer(logo.GPixbuf)))
}

func (v *AboutDialog) GetLogoIconName() string {
	return gostring(C.gtk_about_dialog_get_logo_icon_name(ABOUT_DIALOG(v)))
}

func (v *AboutDialog) SetLogoIconName(icon_name string) {
	ptr := C.CString(icon_name)
	defer cfree(ptr)
	C.gtk_about_dialog_set_logo_icon_name(ABOUT_DIALOG(v), gstring(ptr))
}

// gtk_about_dialog_set_email_hook //deprecated since 2.24
// gtk_about_dialog_set_url_hook //deprecated since 2.24
// gtk_show_about_dialog

//-----------------------------------------------------------------------
// GtkAssistant
//-----------------------------------------------------------------------
type AssistantPageType int

const (
	ASSISTANT_PAGE_CONTENT AssistantPageType = iota
	ASSISTANT_PAGE_INTRO
	ASSISTANT_PAGE_CONFIRM
	ASSISTANT_PAGE_SUMMARY
	ASSISTANT_PAGE_PROGRESS
)

type Assistant struct {
	Widget
}

func NewAssistant() *Assistant {
	return &Assistant{Widget{C.gtk_assistant_new()}}
}

func (v *Assistant) GetCurrentPage() int {
	return int(C.gtk_assistant_get_current_page(ASSISTANT(v)))
}

func (v *Assistant) SetCurrentPage(page_num int) {
	C.gtk_assistant_set_current_page(ASSISTANT(v), gint(page_num))
}

func (v *Assistant) GetNPages() int {
	return int(C.gtk_assistant_get_n_pages(ASSISTANT(v)))
}

func (v *Assistant) GetNthPage(page_num int) *Widget {
	return &Widget{
		C.gtk_assistant_get_nth_page(ASSISTANT(v), gint(page_num))}
}

func (v *Assistant) PrependPage(page IWidget) int {
	return int(C.gtk_assistant_prepend_page(ASSISTANT(v), ToNative(page)))
}

func (v *Assistant) AppendPage(page IWidget) int {
	return int(C.gtk_assistant_prepend_page(ASSISTANT(v), ToNative(page)))
}

func (v *Assistant) InsertPage(page IWidget, position int) int {
	return int(C.gtk_assistant_insert_page(ASSISTANT(v), ToNative(page), gint(position)))
}

// void gtk_assistant_set_forward_page_func (GtkAssistant *assistant, GtkAssistantPageFunc page_func, gpointer data, GDestroyNotify destroy);

func (v *Assistant) SetPageType(page IWidget, t AssistantPageType) {
	C.gtk_assistant_set_page_type(ASSISTANT(v), ToNative(page), C.GtkAssistantPageType(t))
}

func (v *Assistant) GetPageType(page IWidget) AssistantPageType {
	return AssistantPageType(C.gtk_assistant_get_page_type(ASSISTANT(v), ToNative(page)))
}

func (v *Assistant) SetPageTitle(page IWidget, title string) {
	ptr := C.CString(title)
	defer cfree(ptr)
	C.gtk_assistant_set_page_title(ASSISTANT(v), ToNative(page), gstring(ptr))
}

func (v *Assistant) GetPageTitle(page IWidget) string {
	return gostring(C.gtk_assistant_get_page_title(ASSISTANT(v), ToNative(page)))
}

func (v *Assistant) SetPageHeaderImage(page IWidget, pixbuf *gdkpixbuf.Pixbuf) {
	C.gtk_assistant_set_page_header_image(ASSISTANT(v), ToNative(page), (*C.GdkPixbuf)(unsafe.Pointer(pixbuf.GPixbuf)))
}

func (v *Assistant) GetPageHeaderImage(page IWidget) *gdkpixbuf.Pixbuf {
	gpixbuf := C.gtk_assistant_get_page_header_image(ASSISTANT(v), ToNative(page))
	return &gdkpixbuf.Pixbuf{
		GdkPixbuf: gdkpixbuf.NewGdkPixbuf(unsafe.Pointer(gpixbuf)),
		GObject:   glib.ObjectFromNative(unsafe.Pointer(gpixbuf)),
	}
}

func (v *Assistant) SetPageSideImage(page IWidget, pixbuf *gdkpixbuf.Pixbuf) {
	C.gtk_assistant_set_page_side_image(ASSISTANT(v), ToNative(page), (*C.GdkPixbuf)(unsafe.Pointer(pixbuf.GPixbuf)))
}

func (v *Assistant) GetPageSideImage(page IWidget) *gdkpixbuf.Pixbuf {
	gpixbuf := C.gtk_assistant_get_page_side_image(ASSISTANT(v), ToNative(page))
	return &gdkpixbuf.Pixbuf{
		GdkPixbuf: gdkpixbuf.NewGdkPixbuf(unsafe.Pointer(gpixbuf)),
		GObject:   glib.ObjectFromNative(unsafe.Pointer(gpixbuf)),
	}
}

func (v *Assistant) SetPageComplete(page IWidget, complete bool) {
	C.gtk_assistant_set_page_complete(ASSISTANT(v), ToNative(page), gbool(complete))
}

func (v *Assistant) GetPageComplete(page IWidget) bool {
	return gobool(C.gtk_assistant_get_page_complete(ASSISTANT(v), ToNative(page)))
}

func (v *Assistant) AddActionWidget(child IWidget) {
	C.gtk_assistant_add_action_widget(ASSISTANT(v), ToNative(child))
}

func (v *Assistant) RemoveActionWidget(child IWidget) {
	C.gtk_assistant_remove_action_widget(ASSISTANT(v), ToNative(child))
}

func (v *Assistant) UpdateButtonsState() {
	C.gtk_assistant_update_buttons_state(ASSISTANT(v))
}

// gtk_assistant_commit //since 2.22

//-----------------------------------------------------------------------
// GtkOffscreenWindow
//-----------------------------------------------------------------------

// gtk_offscreen_window_new
// gtk_offscreen_window_get_pixmap
// gtk_offscreen_window_get_pixbuf

//-----------------------------------------------------------------------
// GtkAccelLabel
//-----------------------------------------------------------------------
/*type IAccelLabel interface {
	IWidget
	GetAccelWidget() GtkWidget
	SetAccelWidget(GtkWidget)
}*/
type AccelLabel struct {
	Widget
}

func NewAccelLabel(label string) *AccelLabel {
	var ptr *C.char
	if len(label) > 0 {
		ptr = C.CString(label)
		defer cfree(ptr)
	}
	return &AccelLabel{Widget{C.gtk_accel_label_new(gstring(ptr))}}
}

// gtk_accel_label_set_accel_closure

func (v *AccelLabel) GetAccelWidget() Widget {
	return Widget{C.gtk_accel_label_get_accel_widget(ACCEL_LABEL(v))}
}

func (v *AccelLabel) SetAccelWidget(w IWidget) {
	C.gtk_accel_label_set_accel_widget(ACCEL_LABEL(v), ToNative(w))
}

func (v *AccelLabel) GetAccelWidth() uint {
	return uint(C.gtk_accel_label_get_accel_width(ACCEL_LABEL(v)))
}

func (v *AccelLabel) Refetch() bool {
	return gobool(C.gtk_accel_label_refetch(ACCEL_LABEL(v)))
}

//-----------------------------------------------------------------------
// GtkImage
//-----------------------------------------------------------------------
type IconSize int

const (
	ICON_SIZE_INVALID IconSize = iota
	ICON_SIZE_MENU
	ICON_SIZE_SMALL_TOOLBAR
	ICON_SIZE_LARGE_TOOLBAR
	ICON_SIZE_BUTTON
	ICON_SIZE_DND
	ICON_SIZE_DIALOG
)

/*type IImage interface {
	IWidget
}*/
type Image struct {
	Misc
}

func NewImage() *Image {
	return &Image{Misc{Widget{C.gtk_image_new()}}}
}

func NewImageFromFile(filename string) *Image {
	ptr := C.CString(filename)
	defer cfree(ptr)
	return &Image{Misc{Widget{C.gtk_image_new_from_file(gstring(ptr))}}}
}

// gtk_image_new_from_icon_set
// gtk_image_new_from_image

func NewImageFromPixbuf(pixbuf *gdkpixbuf.Pixbuf) *Image {
	return &Image{Misc{Widget{C.gtk_image_new_from_pixbuf((*C.GdkPixbuf)(unsafe.Pointer(pixbuf.GPixbuf)))}}}
}

// gtk_image_new_from_pixmap

func NewImageFromStock(stock_id string, size IconSize) *Image {
	ptr := C.CString(stock_id)
	defer cfree(ptr)
	return &Image{Misc{Widget{C.gtk_image_new_from_stock(gstring(ptr), C.GtkIconSize(size))}}}
}

// gtk_image_new_from_animation

func NewImageFromIconName(stock_id string, size IconSize) *Image {
	ptr := C.CString(stock_id)
	defer cfree(ptr)
	return &Image{Misc{Widget{C.gtk_image_new_from_icon_name(gstring(ptr), C.GtkIconSize(size))}}}
}

// gtk_image_new_from_gicon

func (v *Image) GetPixbuf() *gdkpixbuf.Pixbuf {
	gpixbuf := C.gtk_image_get_pixbuf(IMAGE(v))
	return &gdkpixbuf.Pixbuf{
		GdkPixbuf: gdkpixbuf.NewGdkPixbuf(unsafe.Pointer(gpixbuf)),
		GObject:   glib.ObjectFromNative(unsafe.Pointer(gpixbuf)),
	}
}

func (v *Image) SetFromFile(filename string) {
	ptr := C.CString(filename)
	defer cfree(ptr)
	C.gtk_image_set_from_file(IMAGE(v), gstring(ptr))
}

// gtk_image_set_from_icon_set
// gtk_image_set_from_image

func (v *Image) SetFromPixbuf(pixbuf *gdkpixbuf.Pixbuf) {
	C.gtk_image_set_from_pixbuf(IMAGE(v), (*C.GdkPixbuf)(unsafe.Pointer(pixbuf.GPixbuf)))
}

// gtk_image_set_from_pixmap

func (v *Image) SetFromStock(stock_id string, size IconSize) {
	ptr := C.CString(stock_id)
	defer cfree(ptr)
	C.gtk_image_set_from_stock(IMAGE(v), gstring(ptr), C.GtkIconSize(size))
}

// gtk_image_set_from_animation
// gtk_image_set_from_icon_name
// gtk_image_set_from_gicon

// gtk_image_get_icon_set
// gtk_image_get_image

// gtk_image_get_pixmap
// gtk_image_get_stock
// gtk_image_get_animation
// gtk_image_get_icon_name
// gtk_image_get_gicon
// gtk_image_get_storage_type

func (v *Image) Clear() {
	C.gtk_image_clear(IMAGE(v))
}

// gtk_image_set_pixel_size
// gtk_image_get_pixel_size

//-----------------------------------------------------------------------
// GtkLabel
//-----------------------------------------------------------------------
type Justification int

const (
	JUSTIFY_LEFT Justification = iota
	JUSTIFY_RIGHT
	JUSTIFY_CENTER
	JUSTIFY_FILL
)

type ILabel interface {
	IWidget
	isILabel()
	GetLabel() string
	SetLabel(label string)
}
type Label struct {
	Misc
}

func (Label) isILabel() {}

func NewLabel(label string) *Label {
	var ptr *C.char
	if len(label) > 0 {
		ptr = C.CString(label)
		defer cfree(ptr)
	}
	return &Label{Misc{Widget{C.gtk_label_new(gstring(ptr))}}}
}

func (v *Label) SetText(label string) {
	var ptr *C.char
	if len(label) > 0 {
		ptr = C.CString(label)
		defer cfree(ptr)
	}
	C.gtk_label_set_text(LABEL(v), gstring(ptr))
}

func (v *Label) SetMnemonicWidget(widget IWidget) {
	C.gtk_label_set_mnemonic_widget(LABEL(v), ToNative(widget))
}

func (v *Label) SetMarkup(markup string) {
	ptr := C.CString(markup)
	defer cfree(ptr)
	C.gtk_label_set_markup(LABEL(v), gstring(ptr))
}

func (v *Label) GetMnemonicWidget() *Widget {
	return &Widget{C.gtk_label_get_mnemonic_widget(LABEL(v))}
}

func (v *Label) SetPattern(pattern string) {
	ptr := C.CString(pattern)
	defer cfree(ptr)
	C.gtk_label_set_pattern(LABEL(v), gstring(ptr))
}

func (v *Label) SetJustify(jtype Justification) {
	C.gtk_label_set_justify(LABEL(v), C.GtkJustification(jtype))
}

func (v *Label) SetEllipsize(ellipsize pango.EllipsizeMode) {
	C.gtk_label_set_ellipsize(LABEL(v), C.PangoEllipsizeMode(ellipsize))
}

func (v *Label) SetWidthChars(n_chars int) {
	C.gtk_label_set_width_chars(LABEL(v), gint(n_chars))
}

func (v *Label) SetMaxWidthChars(n_chars int) {
	C.gtk_label_set_max_width_chars(LABEL(v), gint(n_chars))
}

func (v *Label) SetLineWrap(setting bool) {
	C.gtk_label_set_line_wrap(LABEL(v), gbool(setting))
}

func (v *Label) SetUseLineWrapMode(wrap_mode pango.WrapMode) {
	C.gtk_label_set_line_wrap_mode(LABEL(v), C.PangoWrapMode(wrap_mode))
}

// gtk_label_get_layout_offsets
// gtk_label_get_mnemonic_keyval

func (v *Label) GetSelectable() bool {
	return gobool(C.gtk_label_get_selectable(LABEL(v)))
}

func (v *Label) GetText() string {
	return gostring(C.gtk_label_get_text(LABEL(v)))
}

func LabelWithMnemonic(label string) *Label {
	var ptr *C.char
	if len(label) > 0 {
		ptr = C.CString(label)
		defer cfree(ptr)
	}
	return &Label{Misc{Widget{C.gtk_label_new_with_mnemonic(gstring(ptr))}}}
}

func (v *Label) SelectRegion(start_offset int, end_offset int) {
	C.gtk_label_select_region(LABEL(v), gint(start_offset), gint(end_offset))
}

// gtk_label_set_mnemonic_widget

func (v *Label) SetSelectable(setting bool) {
	C.gtk_label_set_selectable(LABEL(v), gbool(setting))
}

func (v *Label) SetTextWithMnemonic(str string) {
	ptr := C.CString(str)
	defer cfree(ptr)
	C.gtk_label_set_text_with_mnemonic(LABEL(v), gstring(ptr))
}

// gtk_label_get_attributes
func (v *Label) GetJustify() Justification {
	return Justification(C.gtk_label_get_justify(LABEL(v)))
}

func (v *Label) GetEllipsize() pango.EllipsizeMode {
	return pango.EllipsizeMode(C.gtk_label_get_ellipsize(LABEL(v)))
}

func (v *Label) GetWidthChars() int {
	return int(C.gtk_label_get_width_chars(LABEL(v)))
}

func (v *Label) GetMaxWidthChars() int {
	return int(C.gtk_label_get_max_width_chars(LABEL(v)))
}

func (v *Label) GetLabel() string {
	return gostring(C.gtk_label_get_label(LABEL(v)))
}

// gtk_label_get_layout

func (v *Label) GetLineWrap() bool {
	return gobool(C.gtk_label_get_line_wrap(LABEL(v)))
}

func (v *Label) GetLineWrapMode() pango.WrapMode {
	return pango.WrapMode(C.gtk_label_get_line_wrap_mode(LABEL(v)))
}

// gtk_label_get_mnemonic_widget
func (v *Label) GetSelectionBounds(start *int, end *int) {
	var cstart, cend C.gint
	C.gtk_label_get_selection_bounds(LABEL(v), &cstart, &cend)
	*start = int(cstart)
	*end = int(cend)
}

func (v *Label) GetUseMarkup() bool {
	return gobool(C.gtk_label_get_use_markup(LABEL(v)))
}

func (v *Label) GetUseUnderline() bool {
	return gobool(C.gtk_label_get_use_underline(LABEL(v)))
}

func (v *Label) GetSingleLineMode() bool {
	return gobool(C.gtk_label_get_single_line_mode(LABEL(v)))
}

func (v *Label) GetAngle() float64 {
	return float64(C.gtk_label_get_angle(LABEL(v)))
}

func (v *Label) SetLabel(label string) {
	var ptr *C.char
	if len(label) > 0 {
		ptr = C.CString(label)
		defer cfree(ptr)
	}
	C.gtk_label_set_label(LABEL(v), gstring(ptr))
}

func (v *Label) SetUseMarkup(setting bool) {
	C.gtk_label_set_use_markup(LABEL(v), gbool(setting))
}

func (v *Label) SetUseUnderline(setting bool) {
	C.gtk_label_set_use_underline(LABEL(v), gbool(setting))
}

func (v *Label) SetSingleLineMode(single_line bool) {
	C.gtk_label_set_single_line_mode(LABEL(v), gbool(single_line))
}

func (v *Label) SetAngle(angle float64) {
	C.gtk_label_set_angle(LABEL(v), gdouble(angle))
}

func (v *Label) GetCurrentUri() string {
	panic_if_version_older(2, 18, 0, "gtk_label_get_current_uri()")
	return gostring(C.gtk_label_get_current_uri(LABEL(v)))
}

func (v *Label) SetTrackVisitedLinks(track_links bool) {
	panic_if_version_older(2, 18, 0, "gtk_label_set_track_visited_links()")
	C.gtk_label_set_track_visited_links(LABEL(v), gbool(track_links))
}

func (v *Label) GetTrackVisitedLinks() bool {
	panic_if_version_older(2, 18, 0, "gtk_label_get_track_visited_links()")
	return gobool(C.gtk_label_get_track_visited_links(LABEL(v)))
}

//-----------------------------------------------------------------------
// GtkProgressBar
//-----------------------------------------------------------------------
type ProgressBarOrientation int

const (
	PROGRESS_LEFT_TO_RIGHT ProgressBarOrientation = iota
	PROGRESS_RIGHT_TO_LEFT
	PROGRESS_BOTTOM_TO_TOP
	PROGRESS_TOP_TO_BOTTOM
)

type ProgressBar struct {
	Widget
}

func NewProgressBar() *ProgressBar {
	return &ProgressBar{Widget{C.gtk_progress_bar_new()}}
}

func (v *ProgressBar) Pulse() {
	C.gtk_progress_bar_pulse(PROGRESS_BAR(v))
}

func (v *ProgressBar) SetText(show_text string) {
	ptr := C.CString(show_text)
	defer cfree(ptr)
	C.gtk_progress_bar_set_text(PROGRESS_BAR(v), gstring(ptr))
}

func (v *ProgressBar) SetFraction(fraction float64) {
	C.gtk_progress_bar_set_fraction(PROGRESS_BAR(v), gdouble(fraction))
}

func (v *ProgressBar) SetPulseStep(fraction float64) {
	C.gtk_progress_bar_set_pulse_step(PROGRESS_BAR(v), gdouble(fraction))
}

func (v *ProgressBar) SetOrientation(i ProgressBarOrientation) {
	C.gtk_progress_bar_set_orientation(PROGRESS_BAR(v), C.GtkProgressBarOrientation(i))
}

func (v *ProgressBar) SetEllipsize(ellipsize pango.EllipsizeMode) {
	C.gtk_progress_bar_set_ellipsize(PROGRESS_BAR(v), C.PangoEllipsizeMode(ellipsize))
}

func (v *ProgressBar) GetText() string {
	return gostring(C.gtk_progress_bar_get_text(PROGRESS_BAR(v)))
}

func (v *ProgressBar) GetFraction() float64 {
	r := C.gtk_progress_bar_get_fraction(PROGRESS_BAR(v))
	return float64(r)
}

func (v *ProgressBar) GetPulseStep() float64 {
	r := C.gtk_progress_bar_get_pulse_step(PROGRESS_BAR(v))
	return float64(r)
}

func (v *ProgressBar) GetOrientation() ProgressBarOrientation {
	return ProgressBarOrientation(C.gtk_progress_bar_get_orientation(PROGRESS_BAR(v)))
}

func (v *ProgressBar) GetEllipsize() pango.EllipsizeMode {
	return pango.EllipsizeMode(C.gtk_progress_bar_get_ellipsize(PROGRESS_BAR(v)))
}

//-----------------------------------------------------------------------
// GtkStatusbar
//-----------------------------------------------------------------------
type Statusbar struct {
	HBox
}

func NewStatusbar() *Statusbar {
	return &Statusbar{HBox{Box{Container{Widget{C.gtk_statusbar_new()}}}}}
}

func (v *Statusbar) GetContextId(content_description string) uint {
	ptr := C.CString(content_description)
	defer cfree(ptr)
	return uint(C.gtk_statusbar_get_context_id(STATUSBAR(v), gstring(ptr)))
}

func (v *Statusbar) Push(context_id uint, text string) {
	ptr := C.CString(text)
	defer cfree(ptr)
	C.gtk_statusbar_push(STATUSBAR(v), guint(context_id), gstring(ptr))
}

func (v *Statusbar) Pop(context_id uint) {
	C.gtk_statusbar_pop(STATUSBAR(v), guint(context_id))
}

func (v *Statusbar) Remove(context_id uint, message_id uint) {
	C.gtk_statusbar_remove(STATUSBAR(v), guint(context_id), guint(message_id))
}

// gtk_statusbar_remove_all //since 2.22

func (v *Statusbar) SetHasResizeGrip(add_tearoffs bool) {
	C.gtk_statusbar_set_has_resize_grip(STATUSBAR(v), gbool(add_tearoffs))
}

func (v *Statusbar) GetHasResizeGrip() bool {
	return gobool(C.gtk_statusbar_get_has_resize_grip(STATUSBAR(v)))
}

// gtk_statusbar_get_message_area //since 2.20

//-----------------------------------------------------------------------
// GtkInfoBar
//-----------------------------------------------------------------------
type InfoBar struct {
	HBox
}

func NewInfoBar() *InfoBar {
	panic_if_version_older_auto(2, 18, 0)
	return &InfoBar{HBox{Box{Container{Widget{C._gtk_info_bar_new()}}}}}
}

func NewInfoBarWithButtons(buttons ...interface{}) *InfoBar {
	panic_if_version_older_auto(2, 18, 0)
	infobar := NewInfoBar()
	text, res := variadicButtonsToArrays(buttons)
	for i := range text {
		infobar.AddButton(text[i], res[i])
	}
	return infobar
}

func (v *InfoBar) AddActionWidget(child IWidget, responseId ResponseType) {
	panic_if_version_older_auto(2, 18, 0)
	C._gtk_info_bar_add_action_widget(INFO_BAR(v), ToNative(child), gint(int(responseId)))
}

func (v *InfoBar) AddButton(buttonText string, responseId ResponseType) *Widget {
	panic_if_version_older_auto(2, 18, 0)
	ptr := C.CString(buttonText)
	defer cfree(ptr)
	return &Widget{C._gtk_info_bar_add_button(INFO_BAR(v), gstring(ptr), gint(int(responseId)))}
}

func (v *InfoBar) AddButtons(buttons ...interface{}) {
	panic_if_version_older_auto(2, 18, 0)
	text, res := variadicButtonsToArrays(buttons)
	for i := range text {
		v.AddButton(text[i], res[i])
	}
}

func (v *InfoBar) SetResponseSensitive(responseId ResponseType, setting bool) {
	panic_if_version_older_auto(2, 18, 0)
	C._gtk_info_bar_set_response_sensitive(INFO_BAR(v), gint(int(responseId)), gbool(setting))
}

func (v *InfoBar) SetDefaultResponse(responseId ResponseType) {
	panic_if_version_older_auto(2, 18, 0)
	C._gtk_info_bar_set_default_response(INFO_BAR(v), gint(int(responseId)))
}

func (v *InfoBar) Response(responseId ResponseType) {
	panic_if_version_older_auto(2, 18, 0)
	C._gtk_info_bar_response(INFO_BAR(v), gint(int(responseId)))
}

func (v *InfoBar) SetMessageType(messageType MessageType) {
	panic_if_version_older_auto(2, 18, 0)
	C._gtk_info_bar_set_message_type(INFO_BAR(v), C.GtkMessageType(messageType))
}

func (v *InfoBar) GetMessageType() MessageType {
	panic_if_version_older_auto(2, 18, 0)
	return MessageType(C._gtk_info_bar_get_message_type(INFO_BAR(v)))
}

func (v *InfoBar) GetActionArea() *Widget {
	panic_if_version_older_auto(2, 18, 0)
	return &Widget{C._gtk_info_bar_get_action_area(INFO_BAR(v))}
}

func (v *InfoBar) GetContentArea() *Widget {
	panic_if_version_older_auto(2, 18, 0)
	return &Widget{C._gtk_info_bar_get_content_area(INFO_BAR(v))}
}

//-----------------------------------------------------------------------
// GtkStatusIcon
//-----------------------------------------------------------------------
type StatusIcon struct {
	GStatusIcon *C.GtkStatusIcon
}

func NewStatusIcon() *StatusIcon {
	return &StatusIcon{C.gtk_status_icon_new()}
}

func NewStatusIconFromPixbuf(pixbuf *gdkpixbuf.Pixbuf) *StatusIcon {
	return &StatusIcon{C.gtk_status_icon_new_from_pixbuf((*C.GdkPixbuf)(unsafe.Pointer(pixbuf.GPixbuf)))}
}

func NewStatusIconFromFile(filename string) *StatusIcon {
	ptr := C.CString(filename)
	defer cfree(ptr)
	return &StatusIcon{C.gtk_status_icon_new_from_file(gstring(ptr))}
}

func NewStatusIconFromStock(stock_id string) *StatusIcon {
	ptr := C.CString(stock_id)
	defer cfree(ptr)
	return &StatusIcon{C.gtk_status_icon_new_from_stock(gstring(ptr))}
}

func NewStatusIconFromIconName(icon_name string) *StatusIcon {
	ptr := C.CString(icon_name)
	defer cfree(ptr)
	return &StatusIcon{C.gtk_status_icon_new_from_icon_name(gstring(ptr))}
}

//GtkStatusIcon *gtk_status_icon_new_from_gicon(GIcon *icon);

func (v *StatusIcon) SetFromPixbuf(pixbuf *gdkpixbuf.Pixbuf) {
	C.gtk_status_icon_set_from_pixbuf(v.GStatusIcon, (*C.GdkPixbuf)(unsafe.Pointer(pixbuf.GPixbuf)))
}

func (v *StatusIcon) SetFromFile(filename string) {
	ptr := C.CString(filename)
	defer cfree(ptr)
	C.gtk_status_icon_set_from_file(v.GStatusIcon, gstring(ptr))
}

func (v *StatusIcon) SetFromStock(stock_id string) {
	ptr := C.CString(stock_id)
	defer cfree(ptr)
	C.gtk_status_icon_set_from_stock(v.GStatusIcon, gstring(ptr))
}

func (v *StatusIcon) SetFromIconName(icon_name string) {
	ptr := C.CString(icon_name)
	defer cfree(ptr)
	C.gtk_status_icon_set_from_icon_name(v.GStatusIcon, gstring(ptr))
}

//void gtk_status_icon_set_from_gicon (GtkStatusIcon *status_icon, GIcon *icon);
//GtkImageType gtk_status_icon_get_storage_type (GtkStatusIcon *status_icon);

func (v *StatusIcon) GetPixbuf() *gdkpixbuf.Pixbuf {
	gpixbuf := C.gtk_status_icon_get_pixbuf(v.GStatusIcon)
	return &gdkpixbuf.Pixbuf{
		GdkPixbuf: gdkpixbuf.NewGdkPixbuf(unsafe.Pointer(gpixbuf)),
		GObject:   glib.ObjectFromNative(unsafe.Pointer(gpixbuf)),
	}
}

func (v *StatusIcon) GetStock() string {
	return gostring(C.gtk_status_icon_get_stock(v.GStatusIcon))
}

func (v *StatusIcon) GetIconName() string {
	return gostring(C.gtk_status_icon_get_icon_name(v.GStatusIcon))
}

func (v *StatusIcon) SetName(name string) {
	panic_if_version_older(2, 20, 0, "gtk_status_icon_set_name()")
	ptr := C.CString(name)
	defer cfree(ptr)
	C._gtk_status_icon_set_name(v.GStatusIcon, gstring(ptr))
}

func (v *StatusIcon) SetTitle(title string) {
	panic_if_version_older(2, 18, 0, "gtk_status_icon_set_title()")
	ptr := C.CString(title)
	defer cfree(ptr)
	C._gtk_status_icon_set_title(v.GStatusIcon, gstring(ptr))
}

func (v *StatusIcon) GetTitle() string {
	panic_if_version_older(2, 18, 0, "gtk_status_icon_get_title()")
	return gostring(C._gtk_status_icon_get_title(v.GStatusIcon))
}

func (v *StatusIcon) SetTooltipText(text string) {
	ptr := C.CString(text)
	defer cfree(ptr)
	C.gtk_status_icon_set_tooltip_text(v.GStatusIcon, gstring(ptr))
}

func (v *StatusIcon) GetTooltipText() string {
	return gostring(C.gtk_status_icon_get_tooltip_text(v.GStatusIcon))
}

func (v *StatusIcon) SetTooltipMarkup(markup string) {
	ptr := C.CString(markup)
	defer cfree(ptr)
	C.gtk_status_icon_set_tooltip_markup(v.GStatusIcon, gstring(ptr))
}

func (v *StatusIcon) GetTooltipMarkup() string {
	return gostring(C.gtk_status_icon_get_tooltip_markup(v.GStatusIcon))
}

func (v *StatusIcon) GetHasTooltip() bool {
	return gobool(C.gtk_status_icon_get_has_tooltip(v.GStatusIcon))
}

func (v *StatusIcon) SetHasTooltip(setting bool) {
	C.gtk_status_icon_set_has_tooltip(v.GStatusIcon, gbool(setting))
}

func (v *StatusIcon) GetVisible() bool {
	return gobool(C.gtk_status_icon_get_visible(v.GStatusIcon))
}

func (v *StatusIcon) SetVisible(setting bool) {
	C.gtk_status_icon_set_visible(v.GStatusIcon, gbool(setting))
}

func StatusIconPositionMenu(menu *Menu, px, py *int, push_in *bool, data interface{}) {
	x := gint(*px)
	y := gint(*py)
	pi := gbool(*push_in)
	var pdata C.gpointer
	if sm, ok := data.(*StatusIcon); ok {
		pdata = C.gpointer(unsafe.Pointer(sm.GStatusIcon))
	}
	C.gtk_status_icon_position_menu(MENU(menu), &x, &y, &pi, pdata)
	*px = int(x)
	*py = int(y)
	*push_in = gobool(pi)
}

func (v *StatusIcon) Connect(s string, f interface{}, datas ...interface{}) int {
	return glib.ObjectFromNative(unsafe.Pointer(C.toGObject(unsafe.Pointer(v.GStatusIcon)))).Connect(s, f, datas...)
}

func PrintContextFromNative(l unsafe.Pointer) *PrintContext {
	return &PrintContext{(*C.GtkPrintContext)(l)}
}

//GIcon *gtk_status_icon_get_gicon (GtkStatusIcon *status_icon);
//gint gtk_status_icon_get_size (GtkStatusIcon *status_icon);
//void gtk_status_icon_set_screen (GtkStatusIcon *status_icon, GdkScreen *screen);
//GdkScreen *gtk_status_icon_get_screen (GtkStatusIcon *status_icon);
// gtk_status_icon_set_blinking //deprecated since 2.22
// gtk_status_icon_get_blinking //deprecated since 2.22
// gtk_status_icon_is_embedded
// gtk_status_icon_get_geometry
// gtk_status_icon_get_x11_window_id

//-----------------------------------------------------------------------
// GtkSpinner
//-----------------------------------------------------------------------

type Spinner struct {
	DrawingArea
}

func NewSpinner() *Spinner {
	panic_if_version_older(2, 20, 0, "NewSpinner()")
	return &Spinner{DrawingArea{Widget{C.gtk_spinner_new()}}}
}

func (s *Spinner) Start() {
	panic_if_version_older(2, 20, 0, "SpinnerStart()")
	C.gtk_spinner_start(SPINNER(s))
}

func (s *Spinner) Stop() {
	panic_if_version_older(2, 20, 0, "SpinnerStop()")
	C.gtk_spinner_stop(SPINNER(s))
}

//-----------------------------------------------------------------------
// GtkButton
//-----------------------------------------------------------------------
/*type IButton interface { // Buttons are ILabel Widgets!
	ILabel
	// the following should be just Clickable; ...
	Clicked(interface{}, ...interface{}) // this is a very simple interface...
}*/
/*type Clickable interface {
	IWidget
	Clicked(interface{}, ...interface{}) // this is a very simple interface...
}*/

type ReliefStyle int

const (
	RELIEF_NORMAL ReliefStyle = iota
	RELIEF_HALF
	RELIEF_NONE
)

type Button struct {
	Bin
	Activatable // implement GtkActivatable interface
}

func (Button) isILabel() {} // TODO

func newButtonInternal(widget *C.GtkWidget) *Button {
	return &Button{Bin{Container{Widget{widget}}}, Activatable{widget}}
}

func NewButton() *Button {
	return newButtonInternal(C.gtk_button_new())
}

func NewButtonWithLabel(label string) *Button {
	var ptr *C.char
	if len(label) > 0 {
		ptr = C.CString(label)
		defer cfree(ptr)
	}
	return newButtonInternal(C.gtk_button_new_with_label(gstring(ptr)))
}

func NewButtonWithMnemonic(label string) *Button {
	var ptr *C.char
	if len(label) > 0 {
		ptr = C.CString(label)
		defer cfree(ptr)
	}
	return newButtonInternal(C.gtk_button_new_with_mnemonic(gstring(ptr)))
}

func NewButtonFromStock(stock_id string) *Button {
	p_stock_id := C.CString(stock_id)
	defer cfree(p_stock_id)
	return newButtonInternal(C.gtk_button_new_from_stock(gstring(p_stock_id)))
}

func (v *Button) Pressed() {
	deprecated_since(2, 20, 0, "gtk_button_pressed()")
	C.gtk_button_pressed(BUTTON(v))
}

func (v *Button) Released() {
	deprecated_since(2, 20, 0, "gtk_button_released()")
	C.gtk_button_released(BUTTON(v))
}

func (v *Button) Clicked(onclick interface{}, datas ...interface{}) int {
	return v.Connect("clicked", onclick, datas...)
}

func (v *Button) Enter() {
	deprecated_since(2, 20, 0, "gtk_button_enter()")
	C.gtk_button_enter(BUTTON(v))
}

func (v *Button) Leave() {
	deprecated_since(2, 20, 0, "gtk_button_leave()")
	C.gtk_button_leave(BUTTON(v))
}

func (v *Button) GetRelief() ReliefStyle {
	return ReliefStyle(C.gtk_button_get_relief(BUTTON(v)))
}

func (v *Button) SetRelief(relief ReliefStyle) {
	C.gtk_button_set_relief(BUTTON(v), C.GtkReliefStyle(relief))
}

func (v *Button) GetLabel() string {
	return gostring(C.gtk_button_get_label(BUTTON(v)))
}

func (v *Button) SetLabel(label string) {
	var ptr *C.char
	if len(label) > 0 {
		ptr = C.CString(label)
		defer cfree(ptr)
	}
	C.gtk_button_set_label(BUTTON(v), gstring(ptr))
}

func (v *Button) GetUseStock() bool {
	return gobool(C.gtk_button_get_use_stock(BUTTON(v)))
}

func (v *Button) SetUseStock(use bool) {
	C.gtk_button_set_use_stock(BUTTON(v), gbool(use))
}

func (v *Button) GetUseUnderline() bool {
	return gobool(C.gtk_button_get_use_underline(BUTTON(v)))
}

func (v *Button) SetUseUnderline(setting bool) {
	C.gtk_button_set_use_underline(BUTTON(v), gbool(setting))
}

func (v *Button) GetFocusOnClick() bool {
	return gobool(C.gtk_button_get_focus_on_click(BUTTON(v)))
}

func (v *Button) SetFocusOnClick(setting bool) {
	C.gtk_button_set_focus_on_click(BUTTON(v), gbool(setting))
}

func (v *Button) SetAlignment(xalign, yalign float64) {
	C.gtk_button_set_alignment(BUTTON(v), C.gfloat(xalign), C.gfloat(yalign))
}

func (v *Button) GetAlignment() (xalign, yalign float64) {
	var xalign_, yalign_ C.gfloat
	C.gtk_button_get_alignment(BUTTON(v), &xalign_, &yalign_)
	return float64(xalign_), float64(yalign_)
}

func (v *Button) SetImage(image IWidget) {
	C.gtk_button_set_image(BUTTON(v), ToNative(image))
}

func (v *Button) GetImage() *Image {
	return &Image{Misc{Widget{C.gtk_button_get_image(BUTTON(v))}}}
}

func (v *Button) SetImagePosition(pos PositionType) {
	C.gtk_button_set_image_position(BUTTON(v), C.GtkPositionType(pos))
}

func (v *Button) GetImagePosition() PositionType {
	return PositionType(C.gtk_button_get_image_position(BUTTON(v)))
}

func (v *Button) GetEventWindow() *gdk.Window {
	return gdk.WindowFromUnsafe(unsafe.Pointer(C.gtk_button_get_event_window(BUTTON(v))))
}

//-----------------------------------------------------------------------
// GtkCheckButton
//-----------------------------------------------------------------------
type CheckButton struct {
	ToggleButton
}

func NewCheckButton() *CheckButton {
	return &CheckButton{ToggleButton{*newButtonInternal(
		C.gtk_check_button_new())}}
}

func NewCheckButtonWithLabel(label string) *CheckButton {
	var ptr *C.char
	if len(label) > 0 {
		ptr = C.CString(label)
		defer cfree(ptr)
	}
	return &CheckButton{ToggleButton{*newButtonInternal(
		C.gtk_check_button_new_with_label(gstring(ptr)))}}
}

func NewCheckButtonWithMnemonic(label string) *CheckButton {
	var ptr *C.char
	if len(label) > 0 {
		ptr = C.CString(label)
		defer cfree(ptr)
	}
	return &CheckButton{ToggleButton{*newButtonInternal(
		C.gtk_check_button_new_with_mnemonic(gstring(ptr)))}}
}

//-----------------------------------------------------------------------
// GtkRadioButton
//-----------------------------------------------------------------------
type RadioButton struct {
	CheckButton
}

func NewRadioButton(group *glib.SList) *RadioButton {
	return &RadioButton{CheckButton{ToggleButton{*newButtonInternal(
		C.gtk_radio_button_new(gslist(group)))}}}
}

func NewRadioButtonFromWidget(w *RadioButton) *RadioButton {
	var widget *C.GtkRadioButton
	if w != nil {
		widget = RADIO_BUTTON(w)
	}
	return &RadioButton{CheckButton{ToggleButton{*newButtonInternal(
		C.gtk_radio_button_new_from_widget(widget))}}}
}

func NewRadioButtonWithLabel(group *glib.SList, label string) *RadioButton {
	var ptr *C.char
	if len(label) > 0 {
		ptr = C.CString(label)
		defer cfree(ptr)
	}
	return &RadioButton{CheckButton{ToggleButton{*newButtonInternal(
		C.gtk_radio_button_new_with_label(gslist(group), gstring(ptr)))}}}
}

func NewRadioButtonWithLabelFromWidget(w *RadioButton, label string) *RadioButton {
	var ptr *C.char
	if len(label) > 0 {
		ptr = C.CString(label)
		defer cfree(ptr)
	}
	var widget *C.GtkRadioButton
	if w != nil {
		widget = RADIO_BUTTON(w)
	}
	return &RadioButton{CheckButton{ToggleButton{*newButtonInternal(
		C.gtk_radio_button_new_with_label_from_widget(widget, gstring(ptr)))}}}
}

func NewRadioButtonWithMnemonic(group *glib.SList, label string) *RadioButton {
	var ptr *C.char
	if len(label) > 0 {
		ptr = C.CString(label)
		defer cfree(ptr)
	}
	return &RadioButton{CheckButton{ToggleButton{*newButtonInternal(
		C.gtk_radio_button_new_with_mnemonic(gslist(group), gstring(ptr)))}}}
}

func NewRadioButtonWithMnemonicFromWidget(w *RadioButton, label string) *RadioButton {
	var ptr *C.char
	if len(label) > 0 {
		ptr = C.CString(label)
		defer cfree(ptr)
	}
	var widget *C.GtkRadioButton
	if w != nil {
		widget = RADIO_BUTTON(w)
	}
	return &RadioButton{CheckButton{ToggleButton{*newButtonInternal(
		C.gtk_radio_button_new_with_mnemonic_from_widget(widget, gstring(ptr)))}}}
}

func (v *RadioButton) GetGroup() *glib.SList {
	return glib.SListFromNative(unsafe.Pointer(C.gtk_radio_button_get_group(RADIO_BUTTON(v))))
}

func (v *RadioButton) SetGroup(group *glib.SList) {
	C.gtk_radio_button_set_group(RADIO_BUTTON(v), gslist(group))
}

//-----------------------------------------------------------------------
// GtkToggleButton
//-----------------------------------------------------------------------
type ToggleButton struct {
	Button
}

func NewToggleButton() *ToggleButton {
	return &ToggleButton{*newButtonInternal(C.gtk_toggle_button_new())}
}

func NewToggleButtonWithLabel(label string) *ToggleButton {
	var ptr *C.char
	if len(label) > 0 {
		ptr = C.CString(label)
		defer cfree(ptr)
	}
	return &ToggleButton{*newButtonInternal(C.gtk_toggle_button_new_with_label(gstring(ptr)))}
}

func NewToggleButtonWithMnemonic(label string) *ToggleButton {
	var ptr *C.char
	if len(label) > 0 {
		ptr = C.CString(label)
		defer cfree(ptr)
	}
	return &ToggleButton{*newButtonInternal(C.gtk_toggle_button_new_with_mnemonic(gstring(ptr)))}
}

func (v *ToggleButton) SetMode(draw_indicator bool) {
	C.gtk_toggle_button_set_mode(TOGGLE_BUTTON(v), gbool(draw_indicator))
}

func (v *ToggleButton) GetMode() bool {
	return gobool(C.gtk_toggle_button_get_mode(TOGGLE_BUTTON(v)))
}

func (v *ToggleButton) GetActive() bool {
	return gobool(C.gtk_toggle_button_get_active(TOGGLE_BUTTON(v)))
}

func (v *ToggleButton) SetActive(is_active bool) {
	C.gtk_toggle_button_set_active(TOGGLE_BUTTON(v), gbool(is_active))
}

func (v *ToggleButton) GetInconsistent() bool {
	return gobool(C.gtk_toggle_button_get_inconsistent(TOGGLE_BUTTON(v)))
}

func (v *ToggleButton) SetInconsistent(setting bool) {
	C.gtk_toggle_button_set_inconsistent(TOGGLE_BUTTON(v), gbool(setting))
}

//-----------------------------------------------------------------------
// GtkLinkButton
//-----------------------------------------------------------------------
type LinkButton struct {
	Button
}

func NewLinkButton(uri string) *LinkButton {
	var ptr *C.char
	if len(uri) > 0 {
		ptr = C.CString(uri)
		defer cfree(ptr)
	}
	return &LinkButton{*newButtonInternal(
		C.gtk_link_button_new(gstring(ptr)))}
}

func NewLinkButtonWithLabel(uri string, label string) *LinkButton {
	var puri *C.char
	if len(uri) > 0 {
		puri = C.CString(uri)
		defer cfree(puri)
	}
	var plabel *C.char
	if len(label) > 0 {
		plabel = C.CString(label)
		defer cfree(plabel)
	}
	return &LinkButton{*newButtonInternal(C.gtk_link_button_new_with_label(gstring(puri), gstring(plabel)))}
}

func (v *LinkButton) GetUri() string {
	return gostring(C.gtk_link_button_get_uri(LINK_BUTTON(v)))
}

func (v *LinkButton) SetUri(uri string) {
	ptr := C.CString(uri)
	defer cfree(ptr)
	C.gtk_link_button_set_uri(LINK_BUTTON(v), gstring(ptr))
}

//gtk_link_button_set_uri_hook has been deprecated since 2.24. Use clicked signal instead. //TODO
//func (v GtkLinkButton) SetUriHook(f func(button *GtkLinkButton, link string, user_data unsafe.Pointer), ) {
// GtkLinkButtonUriFunc gtk_link_button_set_uri_hook (GtkLinkButtonUriFunc func, gpointer data, GDestroyNotify destroy);
//}
func (v *LinkButton) GetVisited() bool {
	return gobool(C.gtk_link_button_get_visited(LINK_BUTTON(v)))
}

func (v *LinkButton) SetVisited(visited bool) {
	C.gtk_link_button_set_visited(LINK_BUTTON(v), gbool(visited))
}

//-----------------------------------------------------------------------
// GtkScaleButton
//-----------------------------------------------------------------------

const (
	GTK_ICON_SIZE_INVALID       IconSize = C.GTK_ICON_SIZE_INVALID
	GTK_ICON_SIZE_MENU          IconSize = C.GTK_ICON_SIZE_MENU
	GTK_ICON_SIZE_SMALL_TOOLBAR IconSize = C.GTK_ICON_SIZE_SMALL_TOOLBAR
	GTK_ICON_SIZE_LARGE_TOOLBAR IconSize = C.GTK_ICON_SIZE_LARGE_TOOLBAR
	GTK_ICON_SIZE_BUTTON        IconSize = C.GTK_ICON_SIZE_BUTTON
	GTK_ICON_SIZE_DND           IconSize = C.GTK_ICON_SIZE_DND
	GTK_ICON_SIZE_DIALOG        IconSize = C.GTK_ICON_SIZE_DIALOG
)

type ScaleButton struct {
	Bin
}

// TODO: wrapper around icons** C.gchar
func NewScaleButton(size IconSize, min, max, step float64, icons **C.gchar) *ScaleButton {
	return &ScaleButton{Bin{Container{Widget{
		C.gtk_scale_button_new(C.GtkIconSize(size), gdouble(min), gdouble(max), gdouble(step), icons)}}}}
}

func (v *ScaleButton) SetAdjustment(a *Adjustment) {
	C.gtk_scale_button_set_adjustment(SCALEBUTTON(v), a.GAdjustment)
}

func (v *ScaleButton) SetIcons(icons **C.gchar) {
	C.gtk_scale_button_set_icons(SCALEBUTTON(v), icons)
}

func (v *ScaleButton) SetValue(value float64) {
	C.gtk_scale_button_set_value(SCALEBUTTON(v), gdouble(value))
}

func (v *ScaleButton) GetAdjustment() *Adjustment {
	return &Adjustment{C.gtk_scale_button_get_adjustment(SCALEBUTTON(v))}
}

func (v *ScaleButton) GetValue() float64 {
	return float64(C.gtk_scale_button_get_value(SCALEBUTTON(v)))
}

func (v *ScaleButton) GetPopup() *Widget {
	return &Widget{C.gtk_scale_button_get_popup(SCALEBUTTON(v))}
}

func (v *ScaleButton) GetPlusButton() *Widget {
	return &Widget{C.gtk_scale_button_get_plus_button(SCALEBUTTON(v))}
}

func (v *ScaleButton) GetMinusButton() *Widget {
	return &Widget{C.gtk_scale_button_get_minus_button(SCALEBUTTON(v))}
}

//-----------------------------------------------------------------------
// GtkVolumeButton
//-----------------------------------------------------------------------

type VolumeButton struct {
	Bin
}

func NewVolumeButton() *VolumeButton {
	return &VolumeButton{Bin{Container{Widget{C.gtk_volume_button_new()}}}}
}

//-----------------------------------------------------------------------
// GtkEntry
//-----------------------------------------------------------------------
/*type ITextInput interface {
	IWidget
	GetText() string
	SetText(label string)
}*/

type Entry struct {
	Widget
	Editable
}

func NewEntry() *Entry {
	w := Widget{C.gtk_entry_new()}
	return &Entry{w, Editable{C.toGEditable(w.GWidget)}}
}

func NewEntryWithBuffer(buffer *EntryBuffer) *Entry {
	panic_if_version_older_auto(2, 18, 0)
	w := Widget{C._gtk_entry_new_with_buffer(buffer.GEntryBuffer)}
	return &Entry{w, Editable{C.toGEditable(w.GWidget)}}
}

func (v *Entry) GetBuffer() *EntryBuffer {
	panic_if_version_older_auto(2, 18, 0)
	return &EntryBuffer{C._gtk_entry_get_buffer(ENTRY(v))}
}

func (v *Entry) SetBuffer(buffer *EntryBuffer) {
	panic_if_version_older_auto(2, 18, 0)
	C._gtk_entry_set_buffer(ENTRY(v), buffer.GEntryBuffer)
}

func (v *Entry) SetText(text string) {
	ptr := C.CString(text)
	defer cfree(ptr)
	C.gtk_entry_set_text(ENTRY(v), gstring(ptr))
}

func (v *Entry) GetText() string {
	return gostring(C.gtk_entry_get_text(ENTRY(v)))
}

func (v *Entry) GetTextLength() int {
	return int(C.gtk_entry_get_text_length(ENTRY(v)))
}

func (v *Entry) SetVisibility(setting bool) {
	C.gtk_entry_set_visibility(ENTRY(v), gbool(setting))
}

func (v *Entry) SetInvisibleChar(ch rune) {
	C.gtk_entry_set_invisible_char(ENTRY(v), C.gunichar(uint8(ch)))
}

func (v *Entry) UnsetInvisibleChar() {
	C.gtk_entry_unset_invisible_char(ENTRY(v))
}

func (v *Entry) SetMaxLength(i int) {
	C.gtk_entry_set_max_length(ENTRY(v), gint(i))
}

func (v *Entry) GetActivatesDefault() bool {
	return gobool(C.gtk_entry_get_activates_default(ENTRY(v)))
}

func (v *Entry) GetHasFrame() bool {
	return gobool(C.gtk_entry_get_has_frame(ENTRY(v)))
}

// gtk_entry_get_inner_border

func (v *Entry) GetWidthChars() int {
	return int(C.gtk_entry_get_width_chars(ENTRY(v)))
}

func (v *Entry) SetActivatesDefault(setting bool) {
	C.gtk_entry_set_activates_default(ENTRY(v), gbool(setting))
}

func (v *Entry) SetHasFrame(setting bool) {
	C.gtk_entry_set_has_frame(ENTRY(v), gbool(setting))
}

// gtk_entry_set_inner_border

func (v *Entry) SetWidthChars(i int) {
	C.gtk_entry_set_width_chars(ENTRY(v), gint(i))
}

func (v *Entry) GetInvisibleChar() rune {
	return rune(C.gtk_entry_get_invisible_char(ENTRY(v)))
}

func (v *Entry) SetAlignment(xalign float64) {
	C.gtk_entry_set_alignment(ENTRY(v), C.gfloat(xalign))
}

func (v *Entry) GetAlignment() float64 {
	return float64(C.gtk_entry_get_alignment(ENTRY(v)))
}

func (v *Entry) SetOverwriteMode(mode bool) {
	C.gtk_entry_set_overwrite_mode(ENTRY(v), gbool(mode))
}

func (v *Entry) GetOverwriteMode() bool {
	return gobool(C.gtk_entry_get_overwrite_mode(ENTRY(v)))
}

// gtk_entry_get_layout
// gtk_entry_get_layout_offsets
// gtk_entry_layout_index_to_text_index
// gtk_entry_text_index_to_layout_index

func (v *Entry) GetMaxLength() int {
	return int(C.gtk_entry_get_max_length(ENTRY(v)))
}

func (v *Entry) GetVisibility() bool {
	return gobool(C.gtk_entry_get_visibility(ENTRY(v)))
}

func (v *Entry) SetCompletion(completion *EntryCompletion) {
	C.gtk_entry_set_completion(ENTRY(v), completion.GEntryCompletion)
}

func (v *Entry) GetCompletion() *EntryCompletion {
	return newEntryCompletion(C.gtk_entry_get_completion(ENTRY(v)))
}

// gtk_entry_set_cursor_hadjustment
// gtk_entry_get_cursor_hadjustment
// gtk_entry_set_progress_fraction
// gtk_entry_get_progress_fraction
// gtk_entry_set_progress_pulse_step
// gtk_entry_get_progress_pulse_step
// gtk_entry_progress_pulse
// gtk_entry_im_context_filter_keypresse //since 2.22
// gtk_entry_reset_im_context //since 2.22
// gtk_entry_set_icon_from_pixbuf
// gtk_entry_set_icon_from_stock
// gtk_entry_set_icon_from_icon_name
// gtk_entry_set_icon_from_gicon
// gtk_entry_get_icon_storage_type
// gtk_entry_get_icon_pixbuf
// gtk_entry_get_icon_stock
// gtk_entry_get_icon_name
// gtk_entry_get_icon_gicon
// gtk_entry_set_icon_activatable
// gtk_entry_get_icon_activatable
// gtk_entry_set_icon_sensitive
// gtk_entry_get_icon_sensitive
// gtk_entry_get_icon_at_pos
// gtk_entry_set_icon_tooltip_text
// gtk_entry_get_icon_tooltip_text
// gtk_entry_set_icon_tooltip_markup
// gtk_entry_get_icon_tooltip_markup
// gtk_entry_set_icon_drag_source
// gtk_entry_get_current_icon_drag_source
// gtk_entry_get_icon_window //since 2.20
// gtk_entry_get_text_window //since 2.20

//-----------------------------------------------------------------------
// GtkEntryBuffer
//-----------------------------------------------------------------------
type EntryBuffer struct {
	GEntryBuffer *C.GtkEntryBuffer
}

func NewEntryBuffer(initialText string) *EntryBuffer {
	panic_if_version_older_auto(2, 18, 0)
	if len(initialText) == 0 {
		return &EntryBuffer{C._gtk_entry_buffer_new(nil, gint(-1))}
	}
	ptr := C.CString(initialText)
	defer cfree(ptr)
	return &EntryBuffer{C._gtk_entry_buffer_new(gstring(ptr), gint(len(initialText)))}
}

func (v *EntryBuffer) GetText() string {
	panic_if_version_older_auto(2, 18, 0)
	return gostring(C._gtk_entry_buffer_get_text(v.GEntryBuffer))
}

func (v *EntryBuffer) SetText(text string) {
	panic_if_version_older_auto(2, 18, 0)
	if len(text) == 0 {
		C._gtk_entry_buffer_set_text(v.GEntryBuffer, nil, gint(-1))
	}
	ptr := C.CString(text)
	defer cfree(ptr)
	C._gtk_entry_buffer_set_text(v.GEntryBuffer, gstring(ptr), gint(len(text)))
}

// gtk_entry_buffer_get_bytes //since 2.18
/*func (v *GtkEntryBuffer) GetBytes() ? {
	panic_if_version_older_auto(2, 18, 0)
	//TODO(any) what is the equivalent type for gsize in go?
	return ?(C._gtk_entry_buffer_get_bytes(v.GEntryBuffer))
}*/

func (v *EntryBuffer) GetLength() uint {
	panic_if_version_older_auto(2, 18, 0)
	return uint(C._gtk_entry_buffer_get_length(v.GEntryBuffer))
}

func (v *EntryBuffer) GetMaxLength() int {
	panic_if_version_older_auto(2, 18, 0)
	return int(C._gtk_entry_buffer_get_max_length(v.GEntryBuffer))
}

func (v *EntryBuffer) SetMaxLength(maxLength int) {
	panic_if_version_older_auto(2, 18, 0)
	C._gtk_entry_buffer_set_max_length(v.GEntryBuffer, gint(maxLength))
}

func (v *EntryBuffer) InsertText(position uint, text string) uint {
	panic_if_version_older_auto(2, 18, 0)
	ptr := C.CString(text)
	defer cfree(ptr)
	return uint(C._gtk_entry_buffer_insert_text(v.GEntryBuffer, guint(position), gstring(ptr), gint(len(text))))
}

func (v *EntryBuffer) DeleteText(position uint, nChars int) uint {
	panic_if_version_older_auto(2, 18, 0)
	return uint(C._gtk_entry_buffer_delete_text(v.GEntryBuffer, guint(position), gint(nChars)))
}

// gtk_entry_buffer_emit_deleted_text //since 2.18
// gtk_entry_buffer_emit_inserted_text //since 2.18

//-----------------------------------------------------------------------
// GtkEntryCompletion
//-----------------------------------------------------------------------

type EntryCompletion struct {
	GEntryCompletion *C.GtkEntryCompletion
	*glib.GObject
}

func newEntryCompletion(entryCompletion *C.GtkEntryCompletion) *EntryCompletion { // TODO
	return &EntryCompletion{
		GEntryCompletion: entryCompletion,
		GObject:          glib.ObjectFromNative(unsafe.Pointer(entryCompletion)),
	}
}

func NewEntryCompletion() *EntryCompletion {
	return newEntryCompletion(C.gtk_entry_completion_new())
}

func (v *EntryCompletion) GetEntry() *Widget {
	return &Widget{C.gtk_entry_completion_get_entry(v.GEntryCompletion)}
}

func (v *EntryCompletion) SetModel(model *TreeModel) {
	C.gtk_entry_completion_set_model(v.GEntryCompletion, model.GTreeModel)
}

func (v *EntryCompletion) GetModel() *TreeModel {
	return &TreeModel{C.gtk_entry_completion_get_model(v.GEntryCompletion)}
}

type EntryCompletionMatchFunc func(completion *EntryCompletion, key string, iter *TreeIter, data interface{})

// gtk_entry_completion_set_match_func

func (v *EntryCompletion) SetMinimumKeyLength(length int) {
	C.gtk_entry_completion_set_minimum_key_length(v.GEntryCompletion, gint(length))
}

func (v *EntryCompletion) GetMinimumKeyLength() int {
	return int(C.gtk_entry_completion_get_minimum_key_length(v.GEntryCompletion))
}

func (v *EntryCompletion) Complete() {
	C.gtk_entry_completion_complete(v.GEntryCompletion)
}

func (v *EntryCompletion) GetCompletionPrefix() string {
	return gostring(C.gtk_entry_completion_get_completion_prefix(v.GEntryCompletion))
}

func (v *EntryCompletion) InsertPrefix() {
	C.gtk_entry_completion_insert_prefix(v.GEntryCompletion)
}

func (v *EntryCompletion) InsertActionText(index int, text string) {
	ptr := C.CString(text)
	defer cfree(ptr)
	C.gtk_entry_completion_insert_action_text(v.GEntryCompletion, gint(index), gstring(ptr))
}

func (v *EntryCompletion) InsertActionMarkup(index int, markup string) {
	ptr := C.CString(markup)
	defer cfree(ptr)
	C.gtk_entry_completion_insert_action_markup(
		v.GEntryCompletion, gint(index), gstring(ptr))
}

func (v *EntryCompletion) DeleteAction(index int) {
	C.gtk_entry_completion_delete_action(v.GEntryCompletion, gint(index))
}

func (v *EntryCompletion) SetTextColumn(column int) {
	C.gtk_entry_completion_set_text_column(v.GEntryCompletion, gint(column))
}

func (v *EntryCompletion) GetTextColumn() int {
	return int(C.gtk_entry_completion_get_text_column(v.GEntryCompletion))
}

func (v *EntryCompletion) SetInlineCompletion(inlineCompletion bool) {
	C.gtk_entry_completion_set_inline_completion(v.GEntryCompletion,
		gbool(inlineCompletion))
}

func (v *EntryCompletion) GetInlineCompletion() bool {
	return gobool(C.gtk_entry_completion_get_inline_completion(v.GEntryCompletion))
}

func (v *EntryCompletion) SetInlineSelection(inlineSelection bool) {
	C.gtk_entry_completion_set_inline_selection(v.GEntryCompletion,
		gbool(inlineSelection))
}

func (v *EntryCompletion) GetInlineSelection() bool {
	return gobool(C.gtk_entry_completion_get_inline_selection(v.GEntryCompletion))
}

func (v *EntryCompletion) SetPopupCompletion(popupCompletion bool) {
	C.gtk_entry_completion_set_popup_completion(v.GEntryCompletion,
		gbool(popupCompletion))
}

func (v *EntryCompletion) GetPopupCompletion() bool {
	return gobool(C.gtk_entry_completion_get_popup_completion(v.GEntryCompletion))
}

func (v *EntryCompletion) SetPopupSetWidth(popupSetWidth bool) {
	C.gtk_entry_completion_set_popup_set_width(v.GEntryCompletion,
		gbool(popupSetWidth))
}

func (v *EntryCompletion) GetPopupSetWidth() bool {
	return gobool(C.gtk_entry_completion_get_popup_set_width(v.GEntryCompletion))
}

func (v *EntryCompletion) SetPopupSingleMatch(popupSingleMatch bool) {
	C.gtk_entry_completion_set_popup_single_match(v.GEntryCompletion,
		gbool(popupSingleMatch))
}

func (v *EntryCompletion) GetPopupSingleMatch() bool {
	return gobool(C.gtk_entry_completion_get_popup_single_match(v.GEntryCompletion))
}

//-----------------------------------------------------------------------
// GtkHScale
//-----------------------------------------------------------------------
func NewHScale(adjustment *Adjustment) *Scale {
	return &Scale{Range{Widget{C.gtk_hscale_new(adjustment.GAdjustment)}}}
}

func NewHScaleWithRange(min, max, step float64) *Scale {
	return &Scale{Range{Widget{
		C.gtk_hscale_new_with_range(gdouble(min), gdouble(max), gdouble(step))}}}
}

//-----------------------------------------------------------------------
// GtkVScale
//-----------------------------------------------------------------------
func NewVScale(a *Adjustment) *Scale {
	return &Scale{Range{Widget{C.gtk_vscale_new(a.GAdjustment)}}}
}

func NewVScaleWithRange(min, max, step float64) *Scale {
	return &Scale{Range{Widget{C.gtk_vscale_new_with_range(gdouble(min), gdouble(max), gdouble(step))}}}
}

//-----------------------------------------------------------------------
// GtkSpinButton
//-----------------------------------------------------------------------
type SpinButtonUpdatePolicy int

const (
	UPDATE_ALWAYS SpinButtonUpdatePolicy = iota
	UPDATE_IF_VALID
)

type SpinType int

const (
	SPIN_STEP_FORWARD SpinType = iota
	SPIN_STEP_BACKWARD
	SPIN_PAGE_FORWARD
	SPIN_PAGE_BACKWARD
	SPIN_HOME
	SPIN_END
	SPIN_USER_DEFINED
)

type SpinButton struct {
	Entry
}

func NewSpinButton(a *Adjustment, climb float64, digits uint) *SpinButton {
	w := Widget{C.gtk_spin_button_new(a.GAdjustment, gdouble(climb), guint(digits))}
	return &SpinButton{Entry{w, Editable{C.toGEditable(w.GWidget)}}}
}

func NewSpinButtonWithRange(min, max, step float64) *SpinButton {
	w := Widget{C.gtk_spin_button_new_with_range(gdouble(min), gdouble(max), gdouble(step))}
	return &SpinButton{Entry{w, Editable{C.toGEditable(w.GWidget)}}}
}

func (v *SpinButton) OnChangeValue(onclick interface{}, datas ...interface{}) int {
	return v.Connect("change-value", onclick, datas...)
}

func (v *SpinButton) OnInput(onclick interface{}, datas ...interface{}) int {
	return v.Connect("input", onclick, datas...)
}

func (v *SpinButton) OnOutput(onclick interface{}, datas ...interface{}) int {
	return v.Connect("output", onclick, datas...)
}

func (v *SpinButton) OnValueChanged(onclick interface{}, datas ...interface{}) int {
	return v.Connect("value-changed", onclick, datas...)
}

func (v *SpinButton) OnWrapped(onclick interface{}, datas ...interface{}) int {
	return v.Connect("wrapped", onclick, datas...)
}

func (v *SpinButton) Configure(a *Adjustment, climb_rate float64, digits uint) {
	C.gtk_spin_button_configure(SPIN_BUTTON(v), a.GAdjustment, gdouble(climb_rate), guint(digits))
}

func (v *SpinButton) SetAdjustment(a *Adjustment) {
	C.gtk_spin_button_set_adjustment(SPIN_BUTTON(v), a.GAdjustment)
}

func (v *SpinButton) GetAdjustment() *Adjustment {
	return &Adjustment{C.gtk_spin_button_get_adjustment(SPIN_BUTTON(v))}
}

func (v *SpinButton) SetDigits(digits uint) {
	C.gtk_spin_button_set_digits(SPIN_BUTTON(v), guint(digits))
}

func (v *SpinButton) SetIncrements(step, page float64) {
	C.gtk_spin_button_set_increments(SPIN_BUTTON(v), gdouble(step), gdouble(page))
}

func (v *SpinButton) SetRange(min, max float64) {
	C.gtk_spin_button_set_range(SPIN_BUTTON(v), gdouble(min), gdouble(max))
}

func (v *SpinButton) GetValueAsFloat() float64 {
	return float64(C.gtk_spin_button_get_value_as_float(SPIN_BUTTON(v)))
}

func (v *SpinButton) GetValueAsInt() int {
	return int(C.gtk_spin_button_get_value_as_int(SPIN_BUTTON(v)))
}

func (v *SpinButton) SetValue(val float64) {
	C.gtk_spin_button_set_value(SPIN_BUTTON(v), gdouble(val))
}

func (v *SpinButton) SetUpdatePolicy(policy SpinButtonUpdatePolicy) {
	C.gtk_spin_button_set_update_policy(SPIN_BUTTON(v), C.GtkSpinButtonUpdatePolicy(policy))
}

func (v *SpinButton) SetNumeric(numeric bool) {
	C.gtk_spin_button_set_numeric(SPIN_BUTTON(v), gbool(numeric))
}

func (v *SpinButton) Spin(direction SpinType, increment float64) {
	C.gtk_spin_button_spin(SPIN_BUTTON(v), C.GtkSpinType(direction), gdouble(increment))
}

func (v *SpinButton) SetWrap(wrap bool) {
	C.gtk_spin_button_set_wrap(SPIN_BUTTON(v), gbool(wrap))
}

func (v *SpinButton) SetSnapToTicks(snap_to_ticks bool) {
	C.gtk_spin_button_set_snap_to_ticks(SPIN_BUTTON(v), gbool(snap_to_ticks))
}

func (v *SpinButton) Update() {
	C.gtk_spin_button_update(SPIN_BUTTON(v))
}

func (v *SpinButton) GetDigits() uint {
	return uint(C.gtk_spin_button_get_digits(SPIN_BUTTON(v)))
}

func (v *SpinButton) GetIncrements() (float64, float64) {
	var step, page C.gdouble
	C.gtk_spin_button_get_increments(SPIN_BUTTON(v), &step, &page)
	return float64(step), float64(page)
}

func (v *SpinButton) GetNumeric() bool {
	return gobool(C.gtk_spin_button_get_numeric(SPIN_BUTTON(v)))
}

func (v *SpinButton) GetRange() ( /*min*/ float64 /*max*/, float64) {
	var min, max C.gdouble
	C.gtk_spin_button_get_range(SPIN_BUTTON(v), &min, &max)
	return float64(min), float64(max)
}

func (v *SpinButton) GetSnapToTicks() bool {
	return gobool(C.gtk_spin_button_get_snap_to_ticks(SPIN_BUTTON(v)))
}

func (v *SpinButton) GetUpdatePolicy() SpinButtonUpdatePolicy {
	return SpinButtonUpdatePolicy(C.gtk_spin_button_get_update_policy(SPIN_BUTTON(v)))
}

func (v *SpinButton) GetValue() float64 {
	return float64(C.gtk_spin_button_get_value(SPIN_BUTTON(v)))
}

func (v *SpinButton) GetWrap() bool {
	return gobool(C.gtk_spin_button_get_wrap(SPIN_BUTTON(v)))
}

//-----------------------------------------------------------------------
// GtkEditable
//-----------------------------------------------------------------------
type Editable struct {
	GEditable *C.GtkEditable
}

func (v *Editable) SelectRegion(startPos int, endPos int) {
	C.gtk_editable_select_region(v.GEditable, gint(startPos), gint(endPos))
}

func (v *Editable) GetSelectionBounds() (isSelected bool,
	startPos int, endPos int) {
	var s, e C.gint
	return gobool(C.gtk_editable_get_selection_bounds(v.GEditable, &s, &e)), int(s), int(e)
}

func (v *Editable) InsertText(newText string, position int) int {
	ptr := C.CString(newText)
	defer cfree(ptr)
	gpos := (C.gint)(position)
	C.gtk_editable_insert_text(v.GEditable, gstring(ptr), gint(len(newText)), &gpos)
	return int(gpos)
}

func (v *Editable) DeleteText(startPos int, endPos int) {
	C.gtk_editable_delete_text(v.GEditable, gint(startPos), gint(endPos))
}

func (v *Editable) GetChars(startPos int, endPos int) string {
	return gostring(C.gtk_editable_get_chars(v.GEditable, gint(startPos), gint(endPos)))
}

func (v *Editable) CutClipboard() {
	C.gtk_editable_cut_clipboard(v.GEditable)
}

func (v *Editable) CopyClipboard() {
	C.gtk_editable_copy_clipboard(v.GEditable)
}

func (v *Editable) PasteClipboard() {
	C.gtk_editable_paste_clipboard(v.GEditable)
}

func (v *Editable) DeleteSelection() {
	C.gtk_editable_delete_selection(v.GEditable)
}

func (v *Editable) SetPosition(position int) {
	C.gtk_editable_set_position(v.GEditable, gint(position))
}

func (v *Editable) GetPosition() int {
	return int(C.gtk_editable_get_position(v.GEditable))
}

func (v *Editable) SetEditable(isEditable bool) {
	C.gtk_editable_set_editable(v.GEditable, gbool(isEditable))
}

func (v *Editable) GetEditable() bool {
	return gobool(C.gtk_editable_get_editable(v.GEditable))
}

//-----------------------------------------------------------------------
// GtkTextIter
//-----------------------------------------------------------------------
type TextIter struct {
	GTextIter C.GtkTextIter
}

type TextSearchFlags int

const (
	TEXT_SEARCH_VISIBLE_ONLY TextSearchFlags = 1 << iota
	TEXT_SEARCH_TEXT_ONLY
	TEXT_SEARCH_CASE_INSENSITIVE
)

func (v *TextIter) GetBuffer() *TextBuffer {
	return newTextBuffer(C.gtk_text_iter_get_buffer(&v.GTextIter))
}

func (v *TextIter) Copy() *TextIter {
	return &TextIter{*C.gtk_text_iter_copy(&v.GTextIter)}
}

func (v *TextIter) Free() {
	C.gtk_text_iter_free(&v.GTextIter)
}

func (v *TextIter) GetOffset() int {
	return int(C.gtk_text_iter_get_offset(&v.GTextIter))
}

func (v *TextIter) GetLine() int {
	return int(C.gtk_text_iter_get_line(&v.GTextIter))
}

func (v *TextIter) GetLineOffset() int {
	return int(C.gtk_text_iter_get_line_offset(&v.GTextIter))
}

func (v *TextIter) GetLineIndex() int {
	return int(C.gtk_text_iter_get_line_index(&v.GTextIter))
}

func (v *TextIter) GetVisibleLineIndex() int {
	return int(C.gtk_text_iter_get_visible_line_index(&v.GTextIter))
}

func (v *TextIter) GetVisibleLineOffset() int {
	return int(C.gtk_text_iter_get_visible_line_offset(&v.GTextIter))
}

func (v *TextIter) GetChar() int {
	return int(C.gtk_text_iter_get_char(&v.GTextIter))
}

func (v *TextIter) GetSlice(end *TextIter) string {
	pchar := cstring(C.gtk_text_iter_get_slice(&v.GTextIter, &end.GTextIter))
	defer cfree(pchar)
	return C.GoString(pchar)
}

func (v *TextIter) GetText(end *TextIter) string {
	pchar := cstring(C.gtk_text_iter_get_text(&v.GTextIter, &end.GTextIter))
	defer cfree(pchar)
	return C.GoString(pchar)
}

func (v *TextIter) GetVisibleSlice(end *TextIter) string {
	return gostring(C.gtk_text_iter_get_visible_slice(&v.GTextIter, &end.GTextIter))
}

func (v *TextIter) GetVisibleText(end *TextIter) string {
	return gostring(C.gtk_text_iter_get_visible_text(&v.GTextIter, &end.GTextIter))
}

// gtk_text_iter_get_pixbuf

func (v *TextIter) GetMarks() *glib.SList {
	return glib.SListFromNative(unsafe.Pointer(C.gtk_text_iter_get_marks(&v.GTextIter)))
}

func (v *TextIter) GetTags() *glib.SList {
	return glib.SListFromNative(unsafe.Pointer(C.gtk_text_iter_get_tags(&v.GTextIter)))
}

// gtk_text_iter_get_toggled_tags
// gtk_text_iter_get_child_anchor
// gtk_text_iter_begins_tag
// gtk_text_iter_ends_tag
// gtk_text_iter_toggles_tag
// gtk_text_iter_has_tag
// gtk_text_iter_editable
// gtk_text_iter_can_insert
// gtk_text_iter_starts_word
// gtk_text_iter_ends_word
// gtk_text_iter_inside_word
// gtk_text_iter_starts_line
// gtk_text_iter_ends_line
// gtk_text_iter_starts_sentence
// gtk_text_iter_ends_sentence
// gtk_text_iter_inside_sentence
// gtk_text_iter_is_cursor_position
// gtk_text_iter_get_chars_in_line
// gtk_text_iter_get_bytes_in_line
// gtk_text_iter_get_attributes
// gtk_text_iter_get_language
// gtk_text_iter_is_end
// gtk_text_iter_is_start

func (v *TextIter) ForwardChar() bool {
	return gobool(C.gtk_text_iter_forward_char(&v.GTextIter))
}

func (v *TextIter) BackwardChar() bool {
	return gobool(C.gtk_text_iter_backward_char(&v.GTextIter))
}

func (v *TextIter) ForwardChars(count int) bool {
	return gobool(C.gtk_text_iter_forward_chars(&v.GTextIter, gint(count)))
}

func (v *TextIter) BackwardChars(count int) bool {
	return gobool(C.gtk_text_iter_backward_chars(&v.GTextIter, gint(count)))
}

// gtk_text_iter_forward_line
// gtk_text_iter_backward_line
// gtk_text_iter_forward_lines
// gtk_text_iter_backward_lines
// gtk_text_iter_forward_word_ends
// gtk_text_iter_backward_word_starts
// gtk_text_iter_forward_word_end
// gtk_text_iter_backward_word_start
// gtk_text_iter_forward_cursor_position
// gtk_text_iter_backward_cursor_position
// gtk_text_iter_forward_cursor_positions
// gtk_text_iter_backward_cursor_positions
// gtk_text_iter_backward_sentence_start
// gtk_text_iter_backward_sentence_starts
// gtk_text_iter_forward_sentence_end
// gtk_text_iter_forward_sentence_ends
// gtk_text_iter_forward_visible_word_ends
// gtk_text_iter_backward_visible_word_starts
// gtk_text_iter_forward_visible_word_end
// gtk_text_iter_backward_visible_word_start
// gtk_text_iter_forward_visible_cursor_position
// gtk_text_iter_backward_visible_cursor_position
// gtk_text_iter_forward_visible_cursor_positions
// gtk_text_iter_backward_visible_cursor_positions
// gtk_text_iter_forward_visible_line
// gtk_text_iter_backward_visible_line
// gtk_text_iter_forward_visible_lines
// gtk_text_iter_backward_visible_lines
// gtk_text_iter_set_offset
// gtk_text_iter_set_line
// gtk_text_iter_set_line_offset
// gtk_text_iter_set_line_index
// gtk_text_iter_set_visible_line_index
// gtk_text_iter_set_visible_line_offset
// gtk_text_iter_forward_to_end
// gtk_text_iter_forward_to_line_end
// gtk_text_iter_forward_to_tag_toggle
// gtk_text_iter_backward_to_tag_toggle
// gtk_text_iter_forward_find_char
// gtk_text_iter_backward_find_char

func (v *TextIter) ForwardSearch(str string, flags TextSearchFlags, start *TextIter, end *TextIter, limit *TextIter) bool {
	cstr := C.CString(str)
	defer cfree(cstr)
	return gobool(C.gtk_text_iter_forward_search(&v.GTextIter, gstring(cstr), C.GtkTextSearchFlags(flags),
		&start.GTextIter, &end.GTextIter, &limit.GTextIter))
}

func (v *TextIter) BackwardSearch(str string, flags TextSearchFlags, start *TextIter, end *TextIter, limit *TextIter) bool {
	cstr := C.CString(str)
	defer cfree(cstr)
	return gobool(C.gtk_text_iter_backward_search(&v.GTextIter, gstring(cstr), C.GtkTextSearchFlags(flags),
		&start.GTextIter, &end.GTextIter, &limit.GTextIter))
}

// gtk_text_iter_equal
// gtk_text_iter_compare
// gtk_text_iter_in_range
// gtk_text_iter_order

func (v *TextIter) Assign(iter *TextIter) {
	C._gtk_text_iter_assign(&v.GTextIter, &iter.GTextIter)
}

//-----------------------------------------------------------------------
// GtkTextMark
//-----------------------------------------------------------------------
type TextMark struct {
	GTextMark *C.GtkTextMark
}

// gtk_text_mark_new
// gtk_text_mark_set_visible
// gtk_text_mark_get_visible
// gtk_text_mark_get_deleted
// gtk_text_mark_get_name
// gtk_text_mark_get_buffer
// gtk_text_mark_get_left_gravity

//-----------------------------------------------------------------------
// GtkTextBuffer
//-----------------------------------------------------------------------
type ITextBuffer interface {
	GetNativeBuffer() unsafe.Pointer
}

type TextBuffer struct {
	GTextBuffer *C.GtkTextBuffer
	*glib.GObject
}

func newTextBuffer(buffer *C.GtkTextBuffer) *TextBuffer { // TODO
	return &TextBuffer{
		GTextBuffer: buffer,
		GObject:     glib.ObjectFromNative(unsafe.Pointer(buffer)),
	}
}

func NewTextBufferFromPointer(v unsafe.Pointer) *TextBuffer {
	return newTextBuffer(TEXT_BUFFER(v))
}

func NewTextBuffer(tagtable *TextTagTable) *TextBuffer {
	return newTextBuffer(C.gtk_text_buffer_new(tagtable.GTextTagTable))
}

func (v *TextBuffer) GetNativeBuffer() unsafe.Pointer {
	return unsafe.Pointer(v.GTextBuffer)
}

func (v *TextBuffer) GetLineCount() int {
	return int(C.gtk_text_buffer_get_line_count(v.GTextBuffer))
}

func (v *TextBuffer) GetCharCount() int {
	return int(C.gtk_text_buffer_get_char_count(v.GTextBuffer))
}

func (v *TextBuffer) GetTagTable() *TextTagTable {
	return &TextTagTable{C.gtk_text_buffer_get_tag_table(v.GTextBuffer)}
}

func (v *TextBuffer) Insert(iter *TextIter, text string) {
	ptr := C.CString(text)
	defer cfree(ptr)
	l := C.strlen(ptr)
	C.gtk_text_buffer_insert(v.GTextBuffer, &iter.GTextIter, gstring(ptr), gsize_t(l))
}

func (v *TextBuffer) InsertAtCursor(text string) {
	ptr := C.CString(text)
	defer cfree(ptr)
	l := C.strlen(ptr)
	C.gtk_text_buffer_insert_at_cursor(v.GTextBuffer, gstring(ptr), gsize_t(l))
}

func (v *TextBuffer) InsertInteractive(iter *TextIter, text string, default_editable bool) bool {
	ptr := C.CString(text)
	defer cfree(ptr)
	l := C.strlen(ptr)
	return gobool(C.gtk_text_buffer_insert_interactive(v.GTextBuffer, &iter.GTextIter, gstring(ptr), gsize_t(l), gbool(default_editable)))
}

func (v *TextBuffer) InsertInteractiveAtCursor(text string, default_editable bool) bool {
	ptr := C.CString(text)
	defer cfree(ptr)
	l := C.strlen(ptr)
	return gobool(C.gtk_text_buffer_insert_interactive_at_cursor(v.GTextBuffer, gstring(ptr), gsize_t(l), gbool(default_editable)))
}

func (v *TextBuffer) InsertRange(iter *TextIter, start *TextIter, end *TextIter) {
	C.gtk_text_buffer_insert_range(v.GTextBuffer, &iter.GTextIter, &start.GTextIter, &end.GTextIter)
}

func (v *TextBuffer) InsertRangeInteractive(iter *TextIter, start *TextIter, end *TextIter, default_editable bool) bool {
	return gobool(C.gtk_text_buffer_insert_range_interactive(v.GTextBuffer, &iter.GTextIter, &start.GTextIter, &end.GTextIter, gbool(default_editable)))
}

func (v *TextBuffer) InsertWithTag(iter *TextIter, text string, tag *TextTag) {
	ptr := C.CString(text)
	defer cfree(ptr)
	l := C.strlen(ptr)
	C._gtk_text_buffer_insert_with_tag(v.GTextBuffer, &iter.GTextIter, gstring(ptr), gsize_t(l), tag.GTextTag)
}

//func (v GtkTextBuffer) InsertWithTags(iter *GtkTextIter, start *GtkTextIter, end *GtkTextIter, default_editable bool) bool {
//	return gobool(C._gtk_text_buffer_insert_range_interactive(v.GTextBuffer, &iter.GTextIter, &start.GTextIter, &end.GTextIter, gbool(default_editable)));
//}
// gtk_text_buffer_insert_with_tags_by_name

func (v *TextBuffer) Delete(start *TextIter, end *TextIter) {
	C.gtk_text_buffer_delete(v.GTextBuffer, &start.GTextIter, &end.GTextIter)
}

func (v *TextBuffer) DeleteInteractive(start *TextIter, end *TextIter, default_editable bool) bool {
	return gobool(C.gtk_text_buffer_delete_interactive(v.GTextBuffer, &start.GTextIter, &end.GTextIter, gbool(default_editable)))
}

func (v *TextBuffer) Backspace(iter *TextIter, interactive bool, default_editable bool) bool {
	return gobool(C.gtk_text_buffer_backspace(v.GTextBuffer, &iter.GTextIter, gbool(interactive), gbool(default_editable)))
}

func (v *TextBuffer) SetText(text string) {
	ptr := C.CString(text)
	defer cfree(ptr)
	l := C.strlen(ptr)
	C.gtk_text_buffer_set_text(v.GTextBuffer, gstring(ptr), gsize_t(l))
}

func (v *TextBuffer) GetText(start *TextIter, end *TextIter, include_hidden_chars bool) string {
	pchar := cstring(C.gtk_text_buffer_get_text(v.GTextBuffer, &start.GTextIter, &end.GTextIter, gbool(include_hidden_chars)))
	defer cfree(pchar)
	return C.GoString(pchar)
}

func (v *TextBuffer) GetSlice(start *TextIter, end *TextIter, include_hidden_chars bool) string {
	pchar := cstring(C.gtk_text_buffer_get_slice(v.GTextBuffer, &start.GTextIter, &end.GTextIter, gbool(include_hidden_chars)))
	defer cfree(pchar)
	return C.GoString(pchar)
}

func (v *TextBuffer) InsertPixbuf(iter *TextIter, pixbuf *gdkpixbuf.Pixbuf) {
	C.gtk_text_buffer_insert_pixbuf(v.GTextBuffer, &iter.GTextIter, (*C.GdkPixbuf)(unsafe.Pointer(pixbuf.GPixbuf)))
}

// gtk_text_buffer_insert_child_anchor
// gtk_text_buffer_create_child_anchor

func (v *TextBuffer) CreateMark(mark_name string, where *TextIter, left_gravity bool) *TextMark {
	ptr := C.CString(mark_name)
	defer cfree(ptr)
	return &TextMark{C.gtk_text_buffer_create_mark(v.GTextBuffer, gstring(ptr), &where.GTextIter, gbool(left_gravity))}
}

func (v *TextBuffer) MoveMark(mark *TextMark, where *TextIter) {
	C.gtk_text_buffer_move_mark(v.GTextBuffer, mark.GTextMark, &where.GTextIter)
}

func (v *TextBuffer) MoveMarkByName(name string, where *TextIter) {
	ptr := C.CString(name)
	C.gtk_text_buffer_move_mark_by_name(v.GTextBuffer, gstring(ptr), &where.GTextIter)
}

func (v *TextBuffer) AddMark(mark *TextMark, where *TextIter) {
	C.gtk_text_buffer_add_mark(v.GTextBuffer, mark.GTextMark, &where.GTextIter)
}

func (v *TextBuffer) DeleteMark(mark *TextMark) {
	C.gtk_text_buffer_delete_mark(v.GTextBuffer, mark.GTextMark)
}

func (v *TextBuffer) DeleteMarkByName(name string) {
	ptr := C.CString(name)
	defer cfree(ptr)
	C.gtk_text_buffer_delete_mark_by_name(v.GTextBuffer, gstring(ptr))
}

func (v *TextBuffer) GetMark(name string) *TextMark {
	ptr := C.CString(name)
	defer cfree(ptr)
	return &TextMark{C.gtk_text_buffer_get_mark(v.GTextBuffer, gstring(ptr))}
}

func (v *TextBuffer) GetInsert() *TextMark {
	return &TextMark{C.gtk_text_buffer_get_insert(v.GTextBuffer)}
}

func (v *TextBuffer) GetSelectionBound() *TextMark {
	return &TextMark{C.gtk_text_buffer_get_selection_bound(v.GTextBuffer)}
}

func (v *TextBuffer) GetHasSelection() bool {
	return gobool(C.gtk_text_buffer_get_has_selection(v.GTextBuffer))
}

func (v *TextBuffer) PlaceCursor(where *TextIter) {
	C.gtk_text_buffer_place_cursor(v.GTextBuffer, &where.GTextIter)
}

func (v *TextBuffer) SelectRange(ins *TextIter, bound *TextIter) {
	C.gtk_text_buffer_select_range(v.GTextBuffer, &ins.GTextIter, &bound.GTextIter)
}

func (v *TextBuffer) ApplyTag(tag *TextTag, start *TextIter, end *TextIter) {
	C.gtk_text_buffer_apply_tag(v.GTextBuffer, tag.GTextTag, &start.GTextIter, &end.GTextIter)
}

func (v *TextBuffer) RemoveTag(tag *TextTag, start *TextIter, end *TextIter) {
	C.gtk_text_buffer_remove_tag(v.GTextBuffer, tag.GTextTag, &start.GTextIter, &end.GTextIter)
}

func (v *TextBuffer) ApplyTagByName(name string, start *TextIter, end *TextIter) {
	ptr := C.CString(name)
	defer cfree(ptr)
	C.gtk_text_buffer_apply_tag_by_name(v.GTextBuffer, gstring(ptr), &start.GTextIter, &end.GTextIter)
}

func (v *TextBuffer) RemoveTagByName(name string, start *TextIter, end *TextIter) {
	ptr := C.CString(name)
	defer cfree(ptr)
	C.gtk_text_buffer_remove_tag_by_name(v.GTextBuffer, gstring(ptr), &start.GTextIter, &end.GTextIter)
}

func (v *TextBuffer) RemoveAllTags(start *TextIter, end *TextIter) {
	C.gtk_text_buffer_remove_all_tags(v.GTextBuffer, &start.GTextIter, &end.GTextIter)
}

func (v *TextBuffer) CreateTag(tag_name string, props map[string]string) *TextTag {
	ptr := C.CString(tag_name)
	defer cfree(ptr)
	tag := C._gtk_text_buffer_create_tag(v.GTextBuffer, gstring(ptr))
	for prop, val := range props {
		pprop := C.CString(prop)
		pval := C.CString(val)
		C._apply_property(unsafe.Pointer(tag), gstring(pprop), gstring(pval))
		cfree(pprop)
		cfree(pval)
	}
	return newTextTag(tag)
}

func (v *TextBuffer) GetIterAtLineOffset(iter *TextIter, line_number int, char_offset int) {
	C.gtk_text_buffer_get_iter_at_line_offset(v.GTextBuffer, &iter.GTextIter, gint(line_number), gint(char_offset))
}

func (v *TextBuffer) GetIterAtOffset(iter *TextIter, char_offset int) {
	C.gtk_text_buffer_get_iter_at_offset(v.GTextBuffer, &iter.GTextIter, gint(char_offset))
}

func (v *TextBuffer) GetIterAtLine(iter *TextIter, line_number int) {
	C.gtk_text_buffer_get_iter_at_line(v.GTextBuffer, &iter.GTextIter, gint(line_number))
}

func (v *TextBuffer) GetIterAtLineIndex(iter *TextIter, line_number int, byte_index int) {
	C.gtk_text_buffer_get_iter_at_line_index(v.GTextBuffer, &iter.GTextIter, gint(line_number), gint(byte_index))
}

func (v *TextBuffer) GetIterAtMark(iter *TextIter, mark *TextMark) {
	C.gtk_text_buffer_get_iter_at_mark(v.GTextBuffer, &iter.GTextIter, mark.GTextMark)
}

func (v *TextBuffer) GetIterAtChildAnchor(i *TextIter, a *TextChildAnchor) {
	C.gtk_text_buffer_get_iter_at_child_anchor(v.GTextBuffer, &i.GTextIter, a.GTextChildAnchor)
}

func (v *TextBuffer) GetStartIter(iter *TextIter) {
	var i C.GtkTextIter
	C.gtk_text_buffer_get_start_iter(v.GTextBuffer, &i)
	iter.GTextIter = i
}

func (v *TextBuffer) GetEndIter(iter *TextIter) {
	var i C.GtkTextIter
	C.gtk_text_buffer_get_end_iter(v.GTextBuffer, &i)
	iter.GTextIter = i
}

func (v *TextBuffer) GetBounds(start *TextIter, end *TextIter) {
	var i1, i2 C.GtkTextIter
	C.gtk_text_buffer_get_bounds(v.GTextBuffer, &i1, &i2)
	start.GTextIter = i1
	end.GTextIter = i2
}

func (v *TextBuffer) GetModified() bool {
	return gobool(C.gtk_text_buffer_get_modified(v.GTextBuffer))
}

func (v *TextBuffer) SetModified(setting bool) {
	C.gtk_text_buffer_set_modified(v.GTextBuffer, gbool(setting))
}

func (v *TextBuffer) DeleteSelection(interactive bool, default_editable bool) {
	C.gtk_text_buffer_delete_selection(v.GTextBuffer, gbool(interactive), gbool(default_editable))
}

// gtk_text_buffer_paste_clipboard
// gtk_text_buffer_copy_clipboard
// gtk_text_buffer_cut_clipboard

func (v *TextBuffer) GetSelectionBounds(be, en *TextIter) bool {
	return gobool(C.gtk_text_buffer_get_selection_bounds(v.GTextBuffer, &be.GTextIter, &en.GTextIter))
}

// gtk_text_buffer_begin_user_action
// gtk_text_buffer_end_user_action
// gtk_text_buffer_add_selection_clipboard
// gtk_text_buffer_remove_selection_clipboard
// gtk_text_buffer_deserialize
// gtk_text_buffer_deserialize_get_can_create_tags
// gtk_text_buffer_deserialize_set_can_create_tags
// gtk_text_buffer_get_copy_target_list
// gtk_text_buffer_get_deserialize_formats
// gtk_text_buffer_get_paste_target_list
// gtk_text_buffer_get_serialize_formats
// gtk_text_buffer_register_deserialize_format
// gtk_text_buffer_register_deserialize_tagset
// gtk_text_buffer_register_serialize_format
// gtk_text_buffer_register_serialize_tagset
// gtk_text_buffer_serialize
// gtk_text_buffer_unregister_deserialize_format
// gtk_text_buffer_unregister_serialize_format

//-----------------------------------------------------------------------
// GtkTextTag
//-----------------------------------------------------------------------
type TextTag struct {
	GTextTag *C.GtkTextTag
	*glib.GObject
}

func NewTextTag(name string) *TextTag {
	ptr := C.CString(name)
	defer cfree(ptr)
	return newTextTag(C.gtk_text_tag_new(gstring(ptr)))
}

func (v *TextTag) SetPriority(priority int) {
	C.gtk_text_tag_set_priority(v.GTextTag, gint(priority))
}

func (v *TextTag) GetPriority() int {
	return int(C.gtk_text_tag_get_priority(v.GTextTag))
}

// gtk_text_tag_event

func newTextTag(tag *C.GtkTextTag) *TextTag { // TODO
	return &TextTag{
		GTextTag: tag,
		GObject:  glib.ObjectFromNative(unsafe.Pointer(tag)),
	}
}

func NewTextTagFromPointer(v unsafe.Pointer) *TextTag {
	return newTextTag(TEXT_TAG(v))
}

//-----------------------------------------------------------------------
// GtkTextAttributes
//-----------------------------------------------------------------------
type TextAttributes struct {
	GTextAttributes *C.GtkTextAttributes
}

func NewTextAttributes() *TextAttributes {
	return &TextAttributes{C.gtk_text_attributes_new()}
}

func (v *TextAttributes) Copy() *TextAttributes {
	return &TextAttributes{C.gtk_text_attributes_copy(v.GTextAttributes)}
}

func (v *TextAttributes) CopyValues(ta *TextAttributes) {
	C.gtk_text_attributes_copy_values(v.GTextAttributes, ta.GTextAttributes)
}

func (v *TextAttributes) Unref() {
	C.gtk_text_attributes_unref(v.GTextAttributes)
}

func (v *TextAttributes) Ref() *TextAttributes {
	return &TextAttributes{C.gtk_text_attributes_ref(v.GTextAttributes)}
}

//-----------------------------------------------------------------------
// GtkTextTagTable
//-----------------------------------------------------------------------
type TextTagTable struct {
	GTextTagTable *C.GtkTextTagTable
}

func NewTextTagTable() *TextTagTable {
	return &TextTagTable{C.gtk_text_tag_table_new()}
}

func (v *TextTagTable) Add(tag *TextTag) {
	C.gtk_text_tag_table_add(v.GTextTagTable, tag.GTextTag)
}

func (v *TextTagTable) Remove(tag *TextTag) {
	C.gtk_text_tag_table_remove(v.GTextTagTable, tag.GTextTag)
}

func (v *TextTagTable) Lookup(name string) *TextTag {
	ptr := C.CString(name)
	defer cfree(ptr)
	return newTextTag(C.gtk_text_tag_table_lookup(v.GTextTagTable, gstring(ptr)))
}

// gtk_text_tag_table_foreach

func (v *TextTagTable) GetSize() int {
	return int(C.gtk_text_tag_table_get_size(v.GTextTagTable))
}

//-----------------------------------------------------------------------
// GtkTextView
//-----------------------------------------------------------------------
type WrapMode int

const (
	WRAP_NONE WrapMode = iota
	WRAP_CHAR
	WRAP_WORD
	WRAP_WORD_CHAR
)

type TextChildAnchor struct {
	GTextChildAnchor *C.GtkTextChildAnchor
}

// gtk_text_child_anchor_new
// gtk_text_child_anchor_get_widgets
// gtk_text_child_anchor_get_deleted
type TextView struct {
	Container
}

func NewTextView() *TextView {
	return &TextView{Container{Widget{C.gtk_text_view_new()}}}
}

func NewTextViewWithBuffer(b TextBuffer) *TextView {
	return &TextView{Container{Widget{C.gtk_text_view_new_with_buffer(b.GTextBuffer)}}}
}

func (v *TextView) SetBuffer(b ITextBuffer) {
	C.gtk_text_view_set_buffer(TEXT_VIEW(v), TEXT_BUFFER(b.GetNativeBuffer()))
}

func (v *TextView) GetBuffer() *TextBuffer {
	return newTextBuffer(C.gtk_text_view_get_buffer(TEXT_VIEW(v)))
}

func (v *TextView) ScrollToMark(mark *TextMark, wm float64, ua bool, xa float64, ya float64) {
	C.gtk_text_view_scroll_to_mark(TEXT_VIEW(v),
		mark.GTextMark, gdouble(wm), gbool(ua), gdouble(xa), gdouble(ya))
}

func (v *TextView) ScrollToIter(iter *TextIter, wm float64, ua bool, xa float64, ya float64) bool {
	return gobool(C.gtk_text_view_scroll_to_iter(TEXT_VIEW(v),
		&iter.GTextIter, gdouble(wm), gbool(ua), gdouble(xa), gdouble(ya)))
}

// void gtk_text_view_scroll_mark_onscreen(GtkTextView* text_view, GtkTextMark* mark);
// gboolean gtk_text_view_move_mark_onscreen(GtkTextView* text_view, GtkTextMark* mark);
// gboolean gtk_text_view_place_cursor_onscreen(GtkTextView* text_view);
// gtk_text_view_get_visible_rect
// void gtk_text_view_get_iter_location(GtkTextView* text_view, const GtkTextIter* iter, GdkRectangle* location);
// void gtk_text_view_get_line_at_y(GtkTextView* text_view, GtkTextIter* target_iter, gint y, gint* line_top);

func (v *TextView) GetLineYrange(iter *TextIter, y *int, h *int) {
	var yy, hh C.gint
	C.gtk_text_view_get_line_yrange(TEXT_VIEW(v), &iter.GTextIter, &yy, &hh)
	*y = int(yy)
	*h = int(hh)
}

func (v *TextView) GetIterAtLocation(iter *TextIter, x int, y int) {
	C.gtk_text_view_get_iter_at_location(TEXT_VIEW(v), &iter.GTextIter, gint(x), gint(y))
}

func (v *TextView) GetIterAtPosition(iter *TextIter, trailing *int, x int, y int) {
	if nil != trailing {
		var tt C.gint
		C.gtk_text_view_get_iter_at_position(TEXT_VIEW(v), &iter.GTextIter, &tt, gint(x), gint(y))
		*trailing = int(tt)
	} else {
		C.gtk_text_view_get_iter_at_position(TEXT_VIEW(v), &iter.GTextIter, nil, gint(x), gint(y))
	}
}

// void gtk_text_view_buffer_to_window_coords(GtkTextView* text_view, GtkTextWindowType win, gint buffer_x, gint buffer_y, gint* window_x, gint* window_y);
// void gtk_text_view_window_to_buffer_coords(GtkTextView* text_view, GtkTextWindowType win, gint window_x, gint window_y, gint* buffer_x, gint* buffer_y);
// GdkWindow* gtk_text_view_get_window(GtkTextView* text_view, GtkTextWindowType win);
// GtkTextWindowType gtk_text_view_get_window_type(GtkTextView* text_view, GdkWindow* window);
// void gtk_text_view_set_border_window_size(GtkTextView* text_view, GtkTextWindowType type, gint size);
// gint gtk_text_view_get_border_window_size(GtkTextView* text_view, GtkTextWindowType type);
// gboolean gtk_text_view_forward_display_line(GtkTextView* text_view, GtkTextIter* iter);
// gboolean gtk_text_view_backward_display_line(GtkTextView* text_view, GtkTextIter* iter);
// gboolean gtk_text_view_forward_display_line_end(GtkTextView* text_view, GtkTextIter* iter);
// gboolean gtk_text_view_backward_display_line_start(GtkTextView* text_view, GtkTextIter* iter);
// gboolean gtk_text_view_starts_display_line(GtkTextView* text_view, const GtkTextIter* iter);
// gboolean gtk_text_view_move_visually(GtkTextView* text_view, GtkTextIter* iter, gint count);
// void gtk_text_view_add_child_at_anchor(GtkTextView* text_view, GtkWidget* child, GtkTextChildAnchor* anchor);
// void gtk_text_view_add_child_in_window(GtkTextView* text_view, GtkWidget* child, GtkTextWindowType which_window, gint xpos, gint ypos);
// void gtk_text_view_move_child(GtkTextView* text_view, GtkWidget* child, gint xpos, gint ypos);

func (v *TextView) SetWrapMode(mode WrapMode) {
	C.gtk_text_view_set_wrap_mode(TEXT_VIEW(v), C.GtkWrapMode(mode))
}

func (v *TextView) GetWrapMode() WrapMode {
	return WrapMode(C.gtk_text_view_get_wrap_mode(TEXT_VIEW(v)))
}

func (v *TextView) SetEditable(setting bool) {
	C.gtk_text_view_set_editable(TEXT_VIEW(v), gbool(setting))
}

func (v *TextView) GetEditable() bool {
	return gobool(C.gtk_text_view_get_editable(TEXT_VIEW(v)))
}

func (v *TextView) SetCursorVisible(setting bool) {
	C.gtk_text_view_set_cursor_visible(TEXT_VIEW(v), gbool(setting))
}

func (v *TextView) GetCursorVisible() bool {
	return gobool(C.gtk_text_view_get_cursor_visible(TEXT_VIEW(v)))
}

func (v *TextView) SetOverwrite(overwrite bool) {
	C.gtk_text_view_set_overwrite(TEXT_VIEW(v), gbool(overwrite))
}

func (v *TextView) GetOverwrite() bool {
	return gobool(C.gtk_text_view_get_overwrite(TEXT_VIEW(v)))
}

// void gtk_text_view_set_pixels_above_lines(GtkTextView* text_view, gint pixels_above_lines);
// gint gtk_text_view_get_pixels_above_lines(GtkTextView* text_view);
// void gtk_text_view_set_pixels_below_lines(GtkTextView* text_view, gint pixels_below_lines);
// gint gtk_text_view_get_pixels_below_lines(GtkTextView* text_view);
// void gtk_text_view_set_pixels_inside_wrap(GtkTextView* text_view, gint pixels_inside_wrap);
// gint gtk_text_view_get_pixels_inside_wrap(GtkTextView* text_view);
// void gtk_text_view_set_justification(GtkTextView* text_view, GtkJustification justification);
// GtkJustification gtk_text_view_get_justification(GtkTextView* text_view);
// void gtk_text_view_set_left_margin(GtkTextView* text_view, gint left_margin);
// gint gtk_text_view_get_left_margin(GtkTextView* text_view);
// void gtk_text_view_set_right_margin(GtkTextView* text_view, gint right_margin);
// gint gtk_text_view_get_right_margin(GtkTextView* text_view);
// void gtk_text_view_set_indent(GtkTextView* text_view, gint indent);
// gint gtk_text_view_get_indent(GtkTextView* text_view);
// void gtk_text_view_set_tabs(GtkTextView* text_view, PangoTabArray* tabs);
// PangoTabArray* gtk_text_view_get_tabs(GtkTextView* text_view);

func (v *TextView) SetAcceptsTab(accepts_tab bool) {
	C.gtk_text_view_set_accepts_tab(TEXT_VIEW(v), gbool(accepts_tab))
}

func (v *TextView) GetAcceptsTab() bool {
	return gobool(C.gtk_text_view_get_accepts_tab(TEXT_VIEW(v)))
}

// GtkTextAttributes* gtk_text_view_get_default_attributes(GtkTextView* text_view);
// gtk_text_view_im_context_filter_keypress //since 2.22
// gtk_text_view_reset_im_context //since 2.22
// gtk_text_view_get_hadjustment //since 2.22
// gtk_text_view_get_vadjustment //since 2.22

//-----------------------------------------------------------------------
// GtkTreePath
//-----------------------------------------------------------------------
type TreePath struct {
	GTreePath *C.GtkTreePath
}

func NewTreePath() *TreePath {
	return &TreePath{C.gtk_tree_path_new()}
}

func NewTreePathFromString(path string) *TreePath {
	ptr := C.CString(path)
	defer cfree(ptr)
	return &TreePath{C.gtk_tree_path_new_from_string(gstring(ptr))}
}

func TreePathFromNative(value unsafe.Pointer) *TreePath {
	return &TreePath{C.to_GTreePath((C.gpointer)(value))}
}

func NewTreePathNewFirst() *TreePath {
	return &TreePath{C.gtk_tree_path_new_first()}
}

// gtk_tree_path_new_from_indices

func (v *TreePath) String() string {
	return gostring(C.gtk_tree_path_to_string(v.GTreePath))
}

func (v *TreePath) AppendIndex(index int) {
	C.gtk_tree_path_append_index(v.GTreePath, gint(index))
}

func (v *TreePath) PrependIndex(index int) {
	C.gtk_tree_path_prepend_index(v.GTreePath, gint(index))
}

func (v *TreePath) GetDepth() int {
	return int(C.gtk_tree_path_get_depth(v.GTreePath))
}

func (v *TreePath) GetIndices() []int {
	depth := v.GetDepth()
	idx := make([]int, depth)

	var gint_size_help C.gint
	addr := uintptr(unsafe.Pointer(C.gtk_tree_path_get_indices(v.GTreePath)))
	size := unsafe.Sizeof(gint_size_help)

	for i := 0; i < depth; i++ {
		idx[i] = int(*((*C.gint)(unsafe.Pointer(addr))))
		addr += size
	}

	return idx
}

// gtk_tree_path_get_indices_with_depth //since 2.22

func (v *TreePath) Free() {
	C.gtk_tree_path_free(v.GTreePath)
}

func (v *TreePath) Copy() *TreePath {
	return &TreePath{C.gtk_tree_path_copy(v.GTreePath)}
}

func (v *TreePath) Compare(w TreePath) int {
	return int(C.gtk_tree_path_compare(v.GTreePath, w.GTreePath))
}

func (v *TreePath) Next() {
	C.gtk_tree_path_next(v.GTreePath)
}

func (v *TreePath) Prev() bool {
	return gobool(C.gtk_tree_path_prev(v.GTreePath))
}

func (v *TreePath) Up() bool {
	return gobool(C.gtk_tree_path_up(v.GTreePath))
}

func (v *TreePath) Down() {
	C.gtk_tree_path_down(v.GTreePath)
}

func (v *TreePath) IsAncestor(descendant TreePath) bool {
	return gobool(C.gtk_tree_path_is_ancestor(v.GTreePath, descendant.GTreePath))
}

func (v *TreePath) IsDescendant(ancestor TreePath) bool {
	return gobool(C.gtk_tree_path_is_descendant(v.GTreePath, ancestor.GTreePath))
}

//-----------------------------------------------------------------------
// GtkTreeRowReference
//-----------------------------------------------------------------------
type TreeRowReference struct {
	GTreeRowReference *C.GtkTreeRowReference
}

func NewTreeRowReference(model ITreeModel, path *TreePath) *TreeRowReference {
	var tm *C.GtkTreeModel
	if model != nil {
		tm = model.cTreeModel()
	}
	return &TreeRowReference{C.gtk_tree_row_reference_new(tm, path.GTreePath)}
}

// gtk_tree_row_reference_new_proxy

func (r *TreeRowReference) GetModel() *TreeModel {
	return &TreeModel{C.gtk_tree_row_reference_get_model(r.GTreeRowReference)}
}

func (r *TreeRowReference) GetPath() *TreePath {
	return &TreePath{C.gtk_tree_row_reference_get_path(r.GTreeRowReference)}
}

func (r *TreeRowReference) Valid() bool {
	return gobool(C.gtk_tree_row_reference_valid(r.GTreeRowReference))
}

func (r *TreeRowReference) Free() {
	C.gtk_tree_row_reference_free(r.GTreeRowReference)
}

func (r *TreeRowReference) Copy() *TreeRowReference {
	return &TreeRowReference{C.gtk_tree_row_reference_copy(r.GTreeRowReference)}
}

// gtk_tree_row_reference_inserted
// gtk_tree_row_reference_deleted
// gtk_tree_row_reference_reordered

//-----------------------------------------------------------------------
// GtkTreeIter
//-----------------------------------------------------------------------
type TreeIter struct {
	GTreeIter *C.GtkTreeIter
}

func (v *TreeIter) ready() {
	if v.GTreeIter == nil {
		v.GTreeIter = C._gtk_tree_iter_new()
	}
}

func (v *TreeIter) Assign(to *TreeIter) {
	v.ready()
	C._gtk_tree_iter_assign(unsafe.Pointer(v.GTreeIter), unsafe.Pointer(to.GTreeIter))
}

//-----------------------------------------------------------------------
// GtkTreeModel
//-----------------------------------------------------------------------
type TreeModelFlags int

const (
	TREE_MODEL_ITERS_PERSIST TreeModelFlags = 1 << iota
	TREE_MODEL_LIST_ONLY
)

type ITreeModel interface {
	cTreeModel() *C.GtkTreeModel
}

type TreeModel struct {
	GTreeModel *C.GtkTreeModel
}

func (v TreeModel) cTreeModel() *C.GtkTreeModel {
	return v.GTreeModel
}

func (v *TreeModel) GetFlags() TreeModelFlags {
	return TreeModelFlags(C.gtk_tree_model_get_flags(v.GTreeModel))
}

func (v *TreeModel) GetNColumns() int {
	return int(C.gtk_tree_model_get_n_columns(v.GTreeModel))
}

// gtk_tree_model_get_column_type
func (v *TreeModel) GetIter(iter *TreeIter, path *TreePath) bool {
	iter.ready()
	return gobool(C.gtk_tree_model_get_iter(v.GTreeModel, iter.GTreeIter, path.GTreePath))
}

func (v *TreeModel) GetIterFromString(iter *TreeIter, path_string string) bool {
	iter.ready()
	ptr := C.CString(path_string)
	defer cfree(ptr)
	return gobool(C.gtk_tree_model_get_iter_from_string(v.GTreeModel, iter.GTreeIter, gstring(ptr)))
}

func (v *TreeModel) GetIterFirst(iter *TreeIter) bool {
	iter.ready()
	return gobool(C.gtk_tree_model_get_iter_first(v.GTreeModel, iter.GTreeIter))
}

func (v *TreeModel) GetPath(iter *TreeIter) *TreePath {
	iter.ready()
	return &TreePath{C._gtk_tree_model_get_path(v.GTreeModel, iter.GTreeIter)}
}

func (v *TreeModel) GetValue(iter *TreeIter, col int, val *glib.GValue) {
	iter.ready()
	C.gtk_tree_model_get_value(v.GTreeModel, iter.GTreeIter, gint(col), C.toGValue(unsafe.Pointer(&val.Value)))
}

func (v *TreeModel) IterNext(iter *TreeIter) bool {
	iter.ready()
	return gobool(C.gtk_tree_model_iter_next(v.GTreeModel, iter.GTreeIter))
}

func (v *TreeModel) IterChildren(iter *TreeIter, parent *TreeIter) bool {
	iter.ready()
	return gobool(C.gtk_tree_model_iter_children(v.GTreeModel, iter.GTreeIter, parent.GTreeIter))
}

func (v *TreeModel) IterHasChild(iter *TreeIter) bool {
	iter.ready()
	return gobool(C.gtk_tree_model_iter_has_child(v.GTreeModel, iter.GTreeIter))
}

func (v *TreeModel) IterNChildren(iter *TreeIter) int {
	iter.ready()
	return int(C.gtk_tree_model_iter_n_children(v.GTreeModel, iter.GTreeIter))
}

func (v *TreeModel) IterNthChild(iter *TreeIter, parent *TreeIter, n int) bool {
	iter.ready()
	return gobool(C.gtk_tree_model_iter_nth_child(v.GTreeModel, iter.GTreeIter, parent.GTreeIter, gint(n)))
}

func (v *TreeModel) IterParent(iter *TreeIter, child *TreeIter) bool {
	iter.ready()
	child.ready()
	return gobool(C.gtk_tree_model_iter_parent(v.GTreeModel, iter.GTreeIter, child.GTreeIter))
}

func (v *TreeModel) GetStringFromIter(iter *TreeIter) string {
	iter.ready()
	return gostring(C.gtk_tree_model_get_string_from_iter(v.GTreeModel, iter.GTreeIter))
}

// gtk_tree_model_ref_node
// gtk_tree_model_unref_node
// gtk_tree_model_get
// gtk_tree_model_get_valist
// gtk_tree_model_foreach
// gtk_tree_model_row_changed
// gtk_tree_model_row_inserted
// gtk_tree_model_row_has_child_toggled
// gtk_tree_model_row_deleted
// gtk_tree_model_rows_reordered

//-----------------------------------------------------------------------
// GtkTreeSelection
//-----------------------------------------------------------------------
type TreeSelection struct {
	GTreeSelection *C.GtkTreeSelection
}

type SelectionMode int

const (
	SELECTION_NONE SelectionMode = iota
	SELECTION_SINGLE
	SELECTION_BROWSE
	SELECTION_MULTIPLE
	SELECTION_EXTENDED = SELECTION_MULTIPLE
)

func (v *TreeSelection) Connect(s string, f interface{}, datas ...interface{}) int {
	return glib.ObjectFromNative(unsafe.Pointer(v.GTreeSelection)).Connect(s, f, datas...)
}

func (v *TreeSelection) SetMode(m SelectionMode) {
	C.gtk_tree_selection_set_mode(v.GTreeSelection, C.GtkSelectionMode(m))
}

func (v *TreeSelection) GetMode() SelectionMode {
	return SelectionMode(C.gtk_tree_selection_get_mode(v.GTreeSelection))
}

//gtk_tree_selection_set_select_function (GtkTreeSelection *selection, GtkTreeSelectionFunc func, gpointer data, GDestroyNotify destroy);
//gtk_tree_selection_get_select_function (GtkTreeSelection *selection);
//gtk_tree_selection_get_tree_view (GtkTreeSelection *selection);

type GtkTreeSelecter interface {
	Select(selection *TreeSelection, model *TreeModel, path *TreePath, selected bool) bool
}

//export _go_gtk_tree_selection_select_func
func _go_gtk_tree_selection_select_func(selection unsafe.Pointer, model unsafe.Pointer, path unsafe.Pointer, selected C.gboolean, payload unsafe.Pointer) C.gboolean {
	cb := (*GtkTreeSelecter)(payload)
	rv := (*cb).Select(&TreeSelection{C.to_GTreeSelection(selection)},
		&TreeModel{C.to_GTreeModel(model)},
		&TreePath{C.to_GTreePath((C.gpointer)(path))},
		gobool(selected))
	return gbool(rv)
}

// Caller is responsible for ensuring that *selecter does not get reaped
// by garbage collection during the time that gtk maintains a reference to
// it (e.g. pass the address of a package global variable).
func (v *TreeSelection) SetSelectFunction(selecter *GtkTreeSelecter) {
	C._go_gtk_tree_selection_set_select_function(v.GTreeSelection, unsafe.Pointer(selecter))
}

func (v *TreeSelection) GetSelected(iter *TreeIter) bool {
	iter.ready()
	return gobool(C.gtk_tree_selection_get_selected(v.GTreeSelection, nil, iter.GTreeIter))
}

//gtk_tree_selection_selected_foreach (GtkTreeSelection *selection, GtkTreeSelectionForeachFunc func, gpointer data);

func (v *TreeSelection) GetSelectedRows(model *TreeModel) *C.GList {
	return C.gtk_tree_selection_get_selected_rows(v.GTreeSelection, &model.GTreeModel)
}

func (v *TreeSelection) CountSelectedRows() int {
	return int(C.gtk_tree_selection_count_selected_rows(v.GTreeSelection))
}

func (v *TreeSelection) SelectPath(path *TreePath) {
	C.gtk_tree_selection_select_path(v.GTreeSelection, path.GTreePath)
}

func (v *TreeSelection) UnselectPath(path *TreePath) {
	C.gtk_tree_selection_unselect_path(v.GTreeSelection, path.GTreePath)
}

func (v *TreeSelection) PathIsSelected(path *TreePath) bool {
	return gobool(C.gtk_tree_selection_path_is_selected(v.GTreeSelection, path.GTreePath))
}

func (v *TreeSelection) SelectIter(iter *TreeIter) {
	iter.ready()
	C.gtk_tree_selection_select_iter(v.GTreeSelection, iter.GTreeIter)
}

func (v *TreeSelection) UnselectIter(iter *TreeIter) {
	iter.ready()
	C.gtk_tree_selection_unselect_iter(v.GTreeSelection, iter.GTreeIter)
}

func (v *TreeSelection) IterIsSelected(iter *TreeIter) bool {
	iter.ready()
	return gobool(C.gtk_tree_selection_iter_is_selected(v.GTreeSelection, iter.GTreeIter))
}

func (v *TreeSelection) SelectAll() {
	C.gtk_tree_selection_select_all(v.GTreeSelection)
}

func (v *TreeSelection) UnselectAll() {
	C.gtk_tree_selection_unselect_all(v.GTreeSelection)
}

func (v *TreeSelection) SelectRange(start_path *TreePath, end_path *TreePath) {
	C.gtk_tree_selection_select_range(v.GTreeSelection, start_path.GTreePath, end_path.GTreePath)
}

func (v *TreeSelection) UnselectRange(start_path *TreePath, end_path *TreePath) {
	C.gtk_tree_selection_unselect_range(v.GTreeSelection, start_path.GTreePath, end_path.GTreePath)
}

//-----------------------------------------------------------------------
// GtkTreeViewColumn
//-----------------------------------------------------------------------
type TreeViewColumnSizing int

const (
	TREE_VIEW_COLUMN_GROW_ONLY TreeViewColumnSizing = iota
	TREE_VIEW_COLUMN_AUTOSIZE
	TREE_VIEW_COLUMN_FIXED
)

type TreeViewColumn struct {
	GTreeViewColumn *C.GtkTreeViewColumn
	*glib.GObject
}

func newTreeViewColumn(column *C.GtkTreeViewColumn) *TreeViewColumn {
	return &TreeViewColumn{
		GTreeViewColumn: column,
		GObject:         glib.ObjectFromNative(unsafe.Pointer(column)),
	}
}

func NewTreeViewColumn() *TreeViewColumn {
	return newTreeViewColumn(C.gtk_tree_view_column_new())
}

func NewTreeViewColumnWithAttribute(title string, cell ICellRenderer) *TreeViewColumn {
	ptitle := C.CString(title)
	defer cfree(ptitle)
	return newTreeViewColumn(
		C._gtk_tree_view_column_new_with_attribute(gstring(ptitle), cell.ToCellRenderer()))
}

/*
func NewTreeViewColumnWithAttributes(title string, cell ICellRenderer, attribute string, column int) *TreeViewColumn {
	ptitle := C.CString(title)
	defer cfree(ptitle)
	pattribute := C.CString(attribute)
	defer cfree(pattribute)
	return newTreeViewColumn(
		C._gtk_tree_view_column_new_with_attributes(gstring(ptitle), cell.ToCellRenderer(), gstring(pattribute), gint(column)))
}
*/
func NewTreeViewColumnWithAttributes(title string, cell ICellRenderer, attributes ...interface{}) *TreeViewColumn {
	if len(attributes)%2 != 0 {
		log.Panic("Error in gtk.TreeViewColumnWithAttributes: last attribute isn't associated to a value, len(attributes) must be even")
	}
	ptrTitle := C.CString(title)
	defer cfree(ptrTitle)
	ret := newTreeViewColumn(C._gtk_tree_view_column_new_with_attribute(
		gstring(ptrTitle), cell.ToCellRenderer()))
	for i := 0; i < len(attributes)/2; i++ {
		attribute, ok := attributes[2*i].(string)
		if !ok {
			log.Panic("Error calling gtk.TreeViewColumnWithAttributes: property name must be a string")
		}
		ptrAttribute := C.CString(attribute)
		column, ok := attributes[2*i+1].(int)
		if !ok {
			log.Panic("Error calling gtk.TreeViewColumnWithAttributes: attributes column must be an int")
		}
		C.gtk_tree_view_column_add_attribute(ret.GTreeViewColumn,
			cell.ToCellRenderer(), gstring(ptrAttribute), gint(column))
	}
	return ret
}

func (v *TreeViewColumn) PackStart(cell ICellRenderer, expand bool) {
	C.gtk_tree_view_column_pack_start(v.GTreeViewColumn, cell.ToCellRenderer(), gbool(expand))
}

func (v *TreeViewColumn) PackEnd(cell ICellRenderer, expand bool) {
	C.gtk_tree_view_column_pack_end(v.GTreeViewColumn, cell.ToCellRenderer(), gbool(expand))
}

func (v *TreeViewColumn) Clear() {
	C.gtk_tree_view_column_clear(v.GTreeViewColumn)
}

// gtk_tree_view_column_get_cell_renderers //deprecated since 2.18

func (v *TreeViewColumn) AddAttribute(cell ICellRenderer, attribute string, column int) {
	ptr := C.CString(attribute)
	defer cfree(ptr)
	C.gtk_tree_view_column_add_attribute(v.GTreeViewColumn, cell.ToCellRenderer(), gstring(ptr), gint(column))
}

//void gtk_tree_view_column_set_attributes (GtkTreeViewColumn *tree_column, GtkCellRenderer *cell_renderer, ...) G_GNUC_NULL_TERMINATED;
//void gtk_tree_view_column_set_cell_data_func (GtkTreeViewColumn *tree_column, GtkCellRenderer *cell_renderer, GtkTreeCellDataFunc func, gpointer func_data, GDestroyNotify destroy);
//void gtk_tree_view_column_clear_attributes (GtkTreeViewColumn *tree_column, GtkCellRenderer *cell_renderer);

func (v *TreeViewColumn) SetSpacing(spacing int) {
	C.gtk_tree_view_column_set_spacing(v.GTreeViewColumn, gint(spacing))
}

func (v *TreeViewColumn) GetSpacing() int {
	return int(C.gtk_tree_view_column_get_spacing(v.GTreeViewColumn))
}

func (v *TreeViewColumn) SetVisible(resizable bool) {
	C.gtk_tree_view_column_set_visible(v.GTreeViewColumn, gbool(resizable))
}

func (v *TreeViewColumn) GetVisible() bool {
	return gobool(C.gtk_tree_view_column_get_visible(v.GTreeViewColumn))
}

func (v *TreeViewColumn) SetResizable(resizable bool) {
	C.gtk_tree_view_column_set_resizable(v.GTreeViewColumn, gbool(resizable))
}

func (v *TreeViewColumn) GetResizable() bool {
	return gobool(C.gtk_tree_view_column_get_resizable(v.GTreeViewColumn))
}

func (v *TreeViewColumn) SetSizing(s TreeViewColumnSizing) {
	C.gtk_tree_view_column_set_sizing(v.GTreeViewColumn, C.GtkTreeViewColumnSizing(s))
}

func (v *TreeViewColumn) GetSizing() TreeViewColumnSizing {
	return TreeViewColumnSizing(C.gtk_tree_view_column_get_sizing(v.GTreeViewColumn))
}

func (v *TreeViewColumn) GetWidth() int {
	return int(C.gtk_tree_view_column_get_width(v.GTreeViewColumn))
}

func (v *TreeViewColumn) GetFixedWidth() int {
	return int(C.gtk_tree_view_column_get_fixed_width(v.GTreeViewColumn))
}

func (v *TreeViewColumn) SetFixedWidth(w int) {
	C.gtk_tree_view_column_set_fixed_width(v.GTreeViewColumn, gint(w))
}

func (v *TreeViewColumn) SetMinWidth(w int) {
	C.gtk_tree_view_column_set_min_width(v.GTreeViewColumn, gint(w))
}

func (v *TreeViewColumn) GetMinWidth() int {
	return int(C.gtk_tree_view_column_get_min_width(v.GTreeViewColumn))
}

func (v *TreeViewColumn) SetMaxWidth(w int) {
	C.gtk_tree_view_column_set_max_width(v.GTreeViewColumn, gint(w))
}

func (v *TreeViewColumn) GetMaxWidth() int {
	return int(C.gtk_tree_view_column_get_max_width(v.GTreeViewColumn))
}

func (v *TreeViewColumn) Clicked() {
	C.gtk_tree_view_column_clicked(v.GTreeViewColumn)
}

func (v *TreeViewColumn) SetTitle(title string) {
	ptr := C.CString(title)
	defer cfree(ptr)
	C.gtk_tree_view_column_set_title(v.GTreeViewColumn, gstring(ptr))

}

func (v *TreeViewColumn) GetTitle() string {
	return gostring(C.gtk_tree_view_column_get_title(v.GTreeViewColumn))
}

func (v *TreeViewColumn) SetExpand(expand bool) {
	C.gtk_tree_view_column_set_expand(v.GTreeViewColumn, gbool(expand))
}

func (v *TreeViewColumn) GetExpand() bool {
	return gobool(C.gtk_tree_view_column_get_expand(v.GTreeViewColumn))
}

func (v *TreeViewColumn) SetClickable(clickable bool) {
	C.gtk_tree_view_column_set_clickable(v.GTreeViewColumn, gbool(clickable))
}

func (v *TreeViewColumn) GetClickable() bool {
	return gobool(C.gtk_tree_view_column_get_clickable(v.GTreeViewColumn))
}

//void gtk_tree_view_column_set_widget (GtkTreeViewColumn *tree_column, GtkWidget *widget);
//GtkWidget *gtk_tree_view_column_get_widget (GtkTreeViewColumn *tree_column);
//void gtk_tree_view_column_set_alignment (GtkTreeViewColumn *tree_column, gfloat xalign);
//gfloat gtk_tree_view_column_get_alignment (GtkTreeViewColumn *tree_column);

func (v *TreeViewColumn) SetReorderable(reorderable bool) {
	C.gtk_tree_view_column_set_reorderable(v.GTreeViewColumn, gbool(reorderable))
}

func (v *TreeViewColumn) GetReorderable() bool {
	return gobool(C.gtk_tree_view_column_get_reorderable(v.GTreeViewColumn))
}

//void gtk_tree_view_column_set_sort_column_id (GtkTreeViewColumn *tree_column, gint sort_column_id);
func (v *TreeViewColumn) SetSortColumnId(col int) {
	C.gtk_tree_view_column_set_sort_column_id(v.GTreeViewColumn, gint(col))
}

//gint gtk_tree_view_column_get_sort_column_id (GtkTreeViewColumn *tree_column);
//void gtk_tree_view_column_set_sort_indicator (GtkTreeViewColumn *tree_column, gboolean setting);
//gboolean gtk_tree_view_column_get_sort_indicator (GtkTreeViewColumn *tree_column);
//void gtk_tree_view_column_set_sort_order (GtkTreeViewColumn *tree_column, GtkSortType order);
//GtkSortType gtk_tree_view_column_get_sort_order (GtkTreeViewColumn *tree_column);
//void gtk_tree_view_column_cell_set_cell_data (GtkTreeViewColumn *tree_column, GtkTreeModel *tree_model, GtkTreeIter *iter, gboolean is_expander, gboolean is_expanded);
//void gtk_tree_view_column_cell_get_size (GtkTreeViewColumn *tree_column, const GdkRectangle *cell_area, gint *x_offset, gint *y_offset, gint *width, gint *height);
//gboolean gtk_tree_view_column_cell_get_position (GtkTreeViewColumn *tree_column, GtkCellRenderer *cell_renderer, gint *start_pos, gint *width);
//gboolean gtk_tree_view_column_cell_is_visible (GtkTreeViewColumn *tree_column);
//void gtk_tree_view_column_focus_cell (GtkTreeViewColumn *tree_column, GtkCellRenderer *cell);
//void gtk_tree_view_column_queue_resize (GtkTreeViewColumn *tree_column);
//GtkWidget *gtk_tree_view_column_get_tree_view (GtkTreeViewColumn *tree_column);

//-----------------------------------------------------------------------
// GtkTreeView
//-----------------------------------------------------------------------
type TreeView struct {
	Container
}

func NewTreeView() *TreeView {
	return &TreeView{Container{Widget{C.gtk_tree_view_new()}}}
}

//gint gtk_tree_view_get_level_indentation (GtkTreeView *tree_view);
//gboolean gtk_tree_view_get_show_expanders (GtkTreeView *tree_view);
//void gtk_tree_view_set_level_indentation (GtkTreeView *tree_view, gint indentation);
//void gtk_tree_view_set_show_expanders (GtkTreeView *tree_view, gboolean enabled);
//GtkWidget *gtk_tree_view_new_with_model (GtkTreeModel *model);

func (v *TreeView) GetModel() *TreeModel {
	var tm = C.gtk_tree_view_get_model(TREE_VIEW(v))
	return &TreeModel{GTreeModel: tm}
}

func (v *TreeView) SetModel(model ITreeModel) {
	var tm *C.GtkTreeModel
	if model != nil {
		tm = model.cTreeModel()
	}
	C.gtk_tree_view_set_model(TREE_VIEW(v), tm)
}

func (v *TreeView) GetSelection() *TreeSelection {
	return &TreeSelection{C.gtk_tree_view_get_selection(TREE_VIEW(v))}
}

//GtkAdjustment *gtk_tree_view_get_hadjustment (GtkTreeView *tree_view);
//void gtk_tree_view_set_hadjustment (GtkTreeView *tree_view, GtkAdjustment *adjustment);
//GtkAdjustment *gtk_tree_view_get_vadjustment (GtkTreeView *tree_view);
//void gtk_tree_view_set_vadjustment (GtkTreeView *tree_view, GtkAdjustment *adjustment);
//gboolean gtk_tree_view_get_headers_visible (GtkTreeView *tree_view);

func (v *TreeView) SetHeadersVisible(flag bool) {
	C.gtk_tree_view_set_headers_visible(TREE_VIEW(v), gbool(flag))
}

//void gtk_tree_view_columns_autosize (GtkTreeView *tree_view);
//gboolean gtk_tree_view_get_headers_clickable (GtkTreeView *tree_view);
//void gtk_tree_view_set_headers_clickable (GtkTreeView *tree_view, gboolean setting);
func (v *TreeView) SetHeadersClickable(flag bool) {
	C.gtk_tree_view_set_headers_clickable(TREE_VIEW(v), gbool(flag))
}

//void gtk_tree_view_set_rules_hint (GtkTreeView *tree_view, gboolean setting);
//gboolean gtk_tree_view_get_rules_hint (GtkTreeView *tree_view);

func (v *TreeView) AppendColumn(c *TreeViewColumn) int {
	return int(C.gtk_tree_view_append_column(TREE_VIEW(v), c.GTreeViewColumn))
}

//gint gtk_tree_view_insert_column (GtkTreeView *tree_view, GtkTreeViewColumn *column, gint position);
//gint gtk_tree_view_insert_column_with_attributes (GtkTreeView *tree_view, gint position, const gchar *title, GtkCellRenderer *cell, ...) G_GNUC_NULL_TERMINATED;
//gint gtk_tree_view_insert_column_with_data_func (GtkTreeView *tree_view, gint position, const gchar *title, GtkCellRenderer *cell, GtkTreeCellDataFunc func, gpointer data, GDestroyNotify dnotify);

func (v *TreeView) GetColumn(n int) *TreeViewColumn {
	return newTreeViewColumn(C.gtk_tree_view_get_column(TREE_VIEW(v), gint(n)))
}

func (v *TreeView) GetColumns() []*TreeViewColumn {
	var columns []*TreeViewColumn
	raw_columns := glib.ListFromNative(unsafe.Pointer(C.gtk_tree_view_get_columns(TREE_VIEW(v))))
	raw_columns.ForEach(func(p unsafe.Pointer, i interface{}) {
		columns = append(columns, newTreeViewColumn((*C.GtkTreeViewColumn)(unsafe.Pointer(uintptr(p)))))
	})
	return columns
}

// Remove column from TreeView and return number of existing columns
func (v *TreeView) RemoveColumn(c *TreeViewColumn) int {
	return int(C.gtk_tree_view_remove_column(TREE_VIEW(v), c.GTreeViewColumn))
}

//void gtk_tree_view_move_column_after (GtkTreeView *tree_view, GtkTreeViewColumn *column, GtkTreeViewColumn *base_column);
//void gtk_tree_view_set_expander_column (GtkTreeView *tree_view, GtkTreeViewColumn *column);
//GtkTreeViewColumn *gtk_tree_view_get_expander_column (GtkTreeView *tree_view);
//void gtk_tree_view_set_column_drag_function (GtkTreeView *tree_view, GtkTreeViewColumnDropFunc func, gpointer user_data, GDestroyNotify destroy);
//void gtk_tree_view_scroll_to_point (GtkTreeView *tree_view, gint tree_x, gint tree_y);

func (v *TreeView) ScrollToCell(path *TreePath, col *TreeViewColumn, use bool, ra, ca float64) {
	var pcol *C.GtkTreeViewColumn
	if nil == col {
		pcol = nil
	} else {
		pcol = col.GTreeViewColumn
	}
	C.gtk_tree_view_scroll_to_cell(TREE_VIEW(v), path.GTreePath, pcol, gbool(use), C.gfloat(ra), C.gfloat(ca))
}

func (v *TreeView) SetCursor(path *TreePath, col *TreeViewColumn, se bool) {
	var pcol *C.GtkTreeViewColumn
	if nil == col {
		pcol = nil
	} else {
		pcol = col.GTreeViewColumn
	}
	C.gtk_tree_view_set_cursor(TREE_VIEW(v), path.GTreePath, pcol, gbool(se))
}

//void gtk_tree_view_set_cursor_on_cell (GtkTreeView *tree_view, GtkTreePath *path, GtkTreeViewColumn *focus_column, GtkCellRenderer *focus_cell, gboolean start_editing);

func (v *TreeView) GetCursor(path **TreePath, focus_column **TreeViewColumn) {
	*path = &TreePath{nil}
	if nil != focus_column {
		*focus_column = &TreeViewColumn{nil, nil}
		C.gtk_tree_view_get_cursor(TREE_VIEW(v), &(*path).GTreePath, &(*focus_column).GTreeViewColumn)
	} else {
		C.gtk_tree_view_get_cursor(TREE_VIEW(v), &(*path).GTreePath, nil)
	}
}

//void gtk_tree_view_row_activated (GtkTreeView *tree_view, GtkTreePath *path, GtkTreeViewColumn *column);

func (v *TreeView) ExpandAll() {
	C.gtk_tree_view_expand_all(TREE_VIEW(v))
}

func (v *TreeView) CollapseAll() {
	C.gtk_tree_view_collapse_all(TREE_VIEW(v))
}

//void gtk_tree_view_expand_to_path (GtkTreeView *tree_view, GtkTreePath *path);

func (v *TreeView) ExpandRow(path *TreePath, openall bool) bool {
	return gobool(C.gtk_tree_view_expand_row(TREE_VIEW(v), path.GTreePath, gbool(openall)))
}

func (v *TreeView) CollapseRow(path *TreePath) bool {
	return gobool(C.gtk_tree_view_collapse_row(TREE_VIEW(v), path.GTreePath))
}

//void gtk_tree_view_map_expanded_rows (GtkTreeView *tree_view, GtkTreeViewMappingFunc func, gpointer data);

func (v *TreeView) RowExpanded(path *TreePath) bool {
	return gobool(C.gtk_tree_view_row_expanded(TREE_VIEW(v), path.GTreePath))
}

//void gtk_tree_view_set_reorderable (GtkTreeView *tree_view, gboolean reorderable);
//gboolean gtk_tree_view_get_reorderable (GtkTreeView *tree_view);
//gboolean gtk_tree_view_get_path_at_pos (GtkTreeView *tree_view, gint x, gint y, GtkTreePath **path, GtkTreeViewColumn **column, gint *cell_x, gint *cell_y);
//void gtk_tree_view_get_cell_area (GtkTreeView *tree_view, GtkTreePath *path, GtkTreeViewColumn *column, GdkRectangle *rect);
//void gtk_tree_view_get_background_area (GtkTreeView *tree_view, GtkTreePath *path, GtkTreeViewColumn *column, GdkRectangle *rect);
//void gtk_tree_view_get_visible_rect (GtkTreeView *tree_view, GdkRectangle *visible_rect);
//gboolean gtk_tree_view_get_visible_range (GtkTreeView *tree_view, GtkTreePath **start_path, GtkTreePath **end_path);
//GdkWindow *gtk_tree_view_get_bin_window (GtkTreeView *tree_view);
//void gtk_tree_view_convert_bin_window_to_tree_coords (GtkTreeView *tree_view, gint bx, gint by, gint *tx, gint *ty);
//void gtk_tree_view_convert_bin_window_to_widget_coords (GtkTreeView *tree_view, gint bx, gint by, gint *wx, gint *wy);
//void gtk_tree_view_convert_tree_to_bin_window_coords (GtkTreeView *tree_view, gint tx, gint ty, gint *bx, gint *by);
//void gtk_tree_view_convert_tree_to_widget_coords (GtkTreeView *tree_view, gint tx, gint ty, gint *wx, gint *wy);
//void gtk_tree_view_convert_widget_to_bin_window_coords (GtkTreeView *tree_view, gint wx, gint wy, gint *bx, gint *by);
//void gtk_tree_view_convert_widget_to_tree_coords (GtkTreeView *tree_view, gint wx, gint wy, gint *tx, gint *ty);
//void gtk_tree_view_enable_model_drag_dest (GtkTreeView *tree_view, const GtkTargetEntry *targets, gint n_targets, GdkDragAction actions);
//void gtk_tree_view_enable_model_drag_source (GtkTreeView *tree_view, GdkModifierType start_button_mask, const GtkTargetEntry *targets, gint n_targets, GdkDragAction actions);
//void gtk_tree_view_unset_rows_drag_source (GtkTreeView *tree_view);
//void gtk_tree_view_unset_rows_drag_dest (GtkTreeView *tree_view);
//void gtk_tree_view_set_drag_dest_row (GtkTreeView *tree_view, GtkTreePath *path, GtkTreeViewDropPosition pos);
//void gtk_tree_view_get_drag_dest_row (GtkTreeView *tree_view, GtkTreePath **path, GtkTreeViewDropPosition *pos);
//gboolean gtk_tree_view_get_dest_row_at_pos (GtkTreeView *tree_view, gint drag_x, gint drag_y, GtkTreePath **path, GtkTreeViewDropPosition *pos);
//GdkPixmap *gtk_tree_view_create_row_drag_icon (GtkTreeView *tree_view, GtkTreePath *path);
//void gtk_tree_view_set_enable_search (GtkTreeView *tree_view, gboolean enable_search);
//gboolean gtk_tree_view_get_enable_search (GtkTreeView *tree_view);
//gint gtk_tree_view_get_search_column (GtkTreeView *tree_view);
//void gtk_tree_view_set_search_column (GtkTreeView *tree_view, gint column);
//GtkTreeViewSearchEqualFunc gtk_tree_view_get_search_equal_func (GtkTreeView *tree_view);
//void gtk_tree_view_set_search_equal_func (GtkTreeView *tree_view, GtkTreeViewSearchEqualFunc search_equal_func, gpointer search_user_data, GDestroyNotify search_destroy);
//GtkEntry *gtk_tree_view_get_search_entry (GtkTreeView *tree_view);
//void gtk_tree_view_set_search_entry (GtkTreeView *tree_view, GtkEntry *entry);
//GtkTreeViewSearchPositionFunc gtk_tree_view_get_search_position_func (GtkTreeView *tree_view);
//void gtk_tree_view_set_search_position_func (GtkTreeView *tree_view, GtkTreeViewSearchPositionFunc func, gpointer data, GDestroyNotify destroy);
//void gtk_tree_view_set_fixed_height_mode (GtkTreeView *tree_view, gboolean enable);
//gboolean gtk_tree_view_get_fixed_height_mode (GtkTreeView *tree_view);
//gboolean gtk_tree_view_get_hover_selection (GtkTreeView *tree_view);
//void gtk_tree_view_set_hover_selection (GtkTreeView *tree_view, gboolean hover);
//gboolean gtk_tree_view_get_hover_expand (GtkTreeView *tree_view);
//void gtk_tree_view_set_hover_expand (GtkTreeView *tree_view, gboolean expand);
//void gtk_tree_view_set_destroy_count_func (GtkTreeView *tree_view, GtkTreeDestroyCountFunc func, gpointer data, GDestroyNotify destroy);
//GtkTreeViewRowSeparatorFunc gtk_tree_view_get_row_separator_func (GtkTreeView *tree_view);
//void gtk_tree_view_set_row_separator_func (GtkTreeView *tree_view, GtkTreeViewRowSeparatorFunc func, gpointer data, GDestroyNotify destroy);
//gboolean gtk_tree_view_get_rubber_banding (GtkTreeView *tree_view);
//void gtk_tree_view_set_rubber_banding (GtkTreeView *tree_view, gboolean enable);
//gboolean gtk_tree_view_is_rubber_banding_active (GtkTreeView *tree_view);
//gboolean gtk_tree_view_get_enable_tree_lines (GtkTreeView *tree_view);
//void gtk_tree_view_set_enable_tree_lines (GtkTreeView *tree_view, gboolean enabled);
//GtkTreeViewGridLines gtk_tree_view_get_grid_lines (GtkTreeView *tree_view);
//void gtk_tree_view_set_grid_lines (GtkTreeView *tree_view, GtkTreeViewGridLines grid_lines);
//void gtk_tree_view_set_tooltip_row (GtkTreeView *tree_view, GtkTooltip *tooltip, GtkTreePath *path);
//void gtk_tree_view_set_tooltip_cell (GtkTreeView *tree_view, GtkTooltip *tooltip, GtkTreePath *path, GtkTreeViewColumn *column, GtkCellRenderer *cell);
//gboolean gtk_tree_view_get_tooltip_context(GtkTreeView *tree_view, gint *x, gint *y, gboolean keyboard_tip, GtkTreeModel **model, GtkTreePath **path, GtkTreeIter *iter);
//gint gtk_tree_view_get_tooltip_column (GtkTreeView *tree_view);
//void gtk_tree_view_set_tooltip_column (GtkTreeView *tree_view, gint column);

//-----------------------------------------------------------------------
// GtkTreeView drag-and-drop
//-----------------------------------------------------------------------

// gtk_tree_drag_source_drag_data_delete
// gtk_tree_drag_source_drag_data_get
// gtk_tree_drag_source_row_draggable
// gtk_tree_drag_dest_drag_data_received
// gtk_tree_drag_dest_row_drop_possible
// gtk_tree_set_row_drag_data
// gtk_tree_get_row_drag_data

//-----------------------------------------------------------------------
// GtkCellView
//-----------------------------------------------------------------------

// gtk_cell_view_new
// gtk_cell_view_new_with_text
// gtk_cell_view_new_with_markup
// gtk_cell_view_new_with_pixbuf
// gtk_cell_view_set_model
// gtk_cell_view_get_model
// gtk_cell_view_set_displayed_row
// gtk_cell_view_get_displayed_row
// gtk_cell_view_get_size_of_row
// gtk_cell_view_set_background_color
// gtk_cell_view_get_cell_renderers

//-----------------------------------------------------------------------
// GtkIconView
//-----------------------------------------------------------------------
type IconView struct {
	Container
}

func NewIconView() *IconView {
	return &IconView{Container{Widget{
		C.gtk_icon_view_new()}}}
}

func NewIconViewWithModel(model ITreeModel) *IconView {
	var tm *C.GtkTreeModel
	if model != nil {
		tm = model.cTreeModel()
	}
	return &IconView{Container{Widget{C.gtk_icon_view_new_with_model(tm)}}}
}

func (v *IconView) GetModel() *TreeModel {
	return &TreeModel{C.gtk_icon_view_get_model(ICON_VIEW(v))}
}

func (v *IconView) SetModel(model ITreeModel) {
	var tm *C.GtkTreeModel
	if model != nil {
		tm = model.cTreeModel()
	}
	C.gtk_icon_view_set_model(ICON_VIEW(v), tm)
}

func (v *IconView) GetTextColumn() int {
	return int(C.gtk_icon_view_get_text_column(ICON_VIEW(v)))
}

func (v *IconView) SetTextColumn(text_column int) {
	C.gtk_icon_view_set_text_column(ICON_VIEW(v), gint(text_column))
}

func (v *IconView) GetMarkupColumn() int {
	return int(C.gtk_icon_view_get_markup_column(ICON_VIEW(v)))
}

func (v *IconView) SetMarkupColumn(markup_column int) {
	C.gtk_icon_view_set_markup_column(ICON_VIEW(v), gint(markup_column))
}

func (v *IconView) GetPixbufColumn() int {
	return int(C.gtk_icon_view_get_pixbuf_column(ICON_VIEW(v)))
}

func (v *IconView) SetPixbufColumn(pixbuf_column int) {
	C.gtk_icon_view_set_pixbuf_column(ICON_VIEW(v), gint(pixbuf_column))
}

// gtk_icon_view_get_path_at_pos
// gtk_icon_view_get_item_at_pos
// gtk_icon_view_convert_widget_to_bin_window_coords
// gtk_icon_view_set_cursor
// gtk_icon_view_get_cursor
// gtk_icon_view_selected_foreach
// gtk_icon_view_set_selection_mode
// gtk_icon_view_get_selection_mode
// gtk_icon_view_set_orientation
// gtk_icon_view_get_orientation
// gtk_icon_view_set_item_orientation //since 2.22
// gtk_icon_view_get_item_orientation //since 2.22
// gtk_icon_view_set_columns
// gtk_icon_view_get_columns
// gtk_icon_view_set_item_width
// gtk_icon_view_get_item_width
// gtk_icon_view_set_spacing
// gtk_icon_view_get_spacing
// gtk_icon_view_set_row_spacing
// gtk_icon_view_get_row_spacing
// gtk_icon_view_set_column_spacing
// gtk_icon_view_get_column_spacing
// gtk_icon_view_set_margin
// gtk_icon_view_get_margin
// gtk_icon_view_set_item_padding //since 2.18
// gtk_icon_view_get_item_padding //since 2.18
// gtk_icon_view_select_path
// gtk_icon_view_unselect_path
// gtk_icon_view_path_is_selected
// gtk_icon_view_get_selected_items
// gtk_icon_view_select_all
// gtk_icon_view_unselect_all
// gtk_icon_view_item_activated

func (v *IconView) ScrollToPath(path *TreePath, use bool, ra float64, ca float64) {
	C.gtk_icon_view_scroll_to_path(ICON_VIEW(v), path.GTreePath,
		gbool(use), C.gfloat(ra), C.gfloat(ca))
}

// gtk_icon_view_get_visible_range
// gtk_icon_view_set_tooltip_item
// gtk_icon_view_set_tooltip_cell
// gtk_icon_view_get_tooltip_context
// gtk_icon_view_set_tooltip_column
// gtk_icon_view_get_tooltip_column
// gtk_icon_view_get_item_row //since 2.22
// gtk_icon_view_get_item_column //since 2.22
// gtk_icon_view_enable_model_drag_source
// gtk_icon_view_enable_model_drag_dest
// gtk_icon_view_unset_model_drag_source
// gtk_icon_view_unset_model_drag_dest
// gtk_icon_view_set_reorderable
// gtk_icon_view_get_reorderable
// gtk_icon_view_set_drag_dest_item
// gtk_icon_view_get_drag_dest_item
// gtk_icon_view_get_dest_item_at_pos
// gtk_icon_view_create_drag_icon

//-----------------------------------------------------------------------
// GtkTreeSortable
//-----------------------------------------------------------------------
type SortType int
type SortFunc func(m *TreeModel, a *TreeIter, b *TreeIter) int

const (
	SORT_ASCENDING SortType = iota
	SORT_DESCENDING
)

const (
	TREE_SORTABLE_DEFAULT_SORT_COLUMN_ID  int = -1
	TREE_SORTABLE_UNSORTED_SORT_COLUMN_ID int = -2
)

type TreeSortable struct {
	GTreeSortable *C.GtkTreeSortable
	sortFuncs     map[int]SortFunc
}

func NewTreeSortable(model ITreeModel) *TreeSortable {
	var tm *C.GtkTreeModel
	if model != nil {
		tm = model.cTreeModel()
	}
	return &TreeSortable{C.toGTreeSortable(tm), make(map[int]SortFunc)}
}

// gtk_tree_sortable_sort_column_changed

func (ts *TreeSortable) GetSortColumnId(id int, order SortType) bool {
	idg := gint(id)
	orderg := C.GtkSortType(order)
	return gobool(C.gtk_tree_sortable_get_sort_column_id(ts.GTreeSortable, &idg, &orderg))
}

func (ts *TreeSortable) SetSortColumnId(id int, order SortType) {
	C.gtk_tree_sortable_set_sort_column_id(ts.GTreeSortable, gint(id), C.GtkSortType(order))
}

func (ts *TreeSortable) SetSortFunc(col int, fun SortFunc) {
	ts.sortFuncs[col] = fun
	C._gtk_tree_sortable_set_sort_func(ts.GTreeSortable, gint(col), pointer.Save(&ts))
}

//export _go_call_sort_func
func _go_call_sort_func(gsfi *C._gtk_sort_func_info) {
	if gsfi == nil {
		return
	}
	gots := *(pointer.Restore(unsafe.Pointer(gsfi.gots)).(**TreeSortable))
	if gots.sortFuncs[int(gsfi.columnNum)] == nil {
		return
	}
	var a, b TreeIter
	a.GTreeIter = gsfi.a
	b.GTreeIter = gsfi.b
	gsfi.ret = C.int(gots.sortFuncs[int(gsfi.columnNum)](&TreeModel{gsfi.model}, &a, &b))
}

// gtk_tree_sortable_set_default_sort_func
// gtk_tree_sortable_has_default_sort_func

//-----------------------------------------------------------------------
// GtkTreeModelSort
//-----------------------------------------------------------------------

// gtk_tree_model_sort_new_with_model
// gtk_tree_model_sort_get_model
// gtk_tree_model_sort_convert_child_path_to_path
// gtk_tree_model_sort_convert_child_iter_to_iter
// gtk_tree_model_sort_convert_path_to_child_path
// gtk_tree_model_sort_convert_iter_to_child_iter
// gtk_tree_model_sort_reset_default_sort_func
// gtk_tree_model_sort_clear_cache
// gtk_tree_model_sort_iter_is_valid

//-----------------------------------------------------------------------
// GtkTreeModelFilter
//-----------------------------------------------------------------------

// gtk_tree_model_filter_new
// gtk_tree_model_filter_set_visible_func
// gtk_tree_model_filter_set_modify_func
// gtk_tree_model_filter_set_visible_column
// gtk_tree_model_filter_get_model
// gtk_tree_model_filter_convert_child_iter_to_iter
// gtk_tree_model_filter_convert_iter_to_child_iter
// gtk_tree_model_filter_convert_child_path_to_path
// gtk_tree_model_filter_convert_path_to_child_path
// gtk_tree_model_filter_refilter
// gtk_tree_model_filter_clear_cache

//-----------------------------------------------------------------------
// GtkCellLayout
//-----------------------------------------------------------------------

// gtk_cell_layout_pack_start
// gtk_cell_layout_pack_end
// gtk_cell_layout_get_cells
// gtk_cell_layout_reorder
// gtk_cell_layout_clear
// gtk_cell_layout_set_attributes
// gtk_cell_layout_add_attribute
// gtk_cell_layout_set_cell_data_func
// gtk_cell_layout_clear_attributes

//-----------------------------------------------------------------------
// GtkCellRenderer
//-----------------------------------------------------------------------
type ICellRenderer interface {
	ToCellRenderer() *C.GtkCellRenderer
}
type CellRenderer struct {
	GCellRenderer *C.GtkCellRenderer
	ICellRenderer
}

func (v *CellRenderer) ToCellRenderer() *C.GtkCellRenderer {
	return v.GCellRenderer
}

func (v *CellRenderer) Connect(s string, f interface{}, datas ...interface{}) int {
	return glib.ObjectFromNative(unsafe.Pointer(v.GCellRenderer)).Connect(s, f, datas...)
}

//-----------------------------------------------------------------------
// GtkCellEditable
//-----------------------------------------------------------------------

// gtk_cell_editable_start_editing
// gtk_cell_editable_editing_done
// gtk_cell_editable_remove_widget

//-----------------------------------------------------------------------
// GtkCellRendererAccel
//-----------------------------------------------------------------------
type CellRendererAccel struct {
	CellRenderer
}

func NewCellRendererAccel() *CellRendererAccel {
	return &CellRendererAccel{CellRenderer{C.gtk_cell_renderer_accel_new(), nil}}
}

//-----------------------------------------------------------------------
// GtkCellRendererCombo
//-----------------------------------------------------------------------
type CellRendererCombo struct {
	CellRenderer
}

func NewCellRendererCombo() *CellRendererCombo {
	return &CellRendererCombo{CellRenderer{C.gtk_cell_renderer_combo_new(), nil}}
}

//-----------------------------------------------------------------------
// GtkCellRendererPixbuf
//-----------------------------------------------------------------------
type CellRendererPixbuf struct {
	CellRenderer
}

func NewCellRendererPixbuf() *CellRendererPixbuf {
	return &CellRendererPixbuf{CellRenderer{C.gtk_cell_renderer_pixbuf_new(), nil}}
}

//-----------------------------------------------------------------------
// GtkCellRendererProgress
//-----------------------------------------------------------------------
type CellRendererProgress struct {
	CellRenderer
}

func NewCellRendererProgress() *CellRendererProgress {
	return &CellRendererProgress{CellRenderer{C.gtk_cell_renderer_progress_new(), nil}}
}

//-----------------------------------------------------------------------
// GtkCellRendererSpin
//-----------------------------------------------------------------------
type CellRendererSpin struct {
	CellRenderer
}

func NewCellRendererSpin() *CellRendererSpin {
	return &CellRendererSpin{CellRenderer{C.gtk_cell_renderer_spin_new(), nil}}
}

//-----------------------------------------------------------------------
// GtkCellRendererText
//-----------------------------------------------------------------------
type CellRendererText struct {
	CellRenderer
}

func NewCellRendererText() *CellRendererText {
	return &CellRendererText{CellRenderer{C.gtk_cell_renderer_text_new(), nil}}
}

func (v *CellRendererText) SetFixedHeightFromFont(number_of_rows int) {
	C.gtk_cell_renderer_text_set_fixed_height_from_font(CELL_RENDERER_TEXT(v), gint(number_of_rows))
}

//-----------------------------------------------------------------------
// GtkCellRendererToggle
//-----------------------------------------------------------------------
type CellRendererToggle struct {
	CellRenderer
}

func NewCellRendererToggle() *CellRendererToggle {
	return &CellRendererToggle{CellRenderer{C.gtk_cell_renderer_toggle_new(), nil}}
}

func (v *CellRendererToggle) GetRadio() bool {
	return gobool(C.gtk_cell_renderer_toggle_get_radio(CELL_RENDERER_TOGGLE(v)))
}

func (v *CellRendererToggle) SetRadio(radio bool) {
	C.gtk_cell_renderer_toggle_set_radio(CELL_RENDERER_TOGGLE(v), gbool(radio))
}

func (v *CellRendererToggle) GetActive() bool {
	return gobool(C.gtk_cell_renderer_toggle_get_active(CELL_RENDERER_TOGGLE(v)))
}

func (v *CellRendererToggle) SetActive(active bool) {
	C.gtk_cell_renderer_toggle_set_active(CELL_RENDERER_TOGGLE(v), gbool(active))
}

func (v *CellRendererToggle) GetActivatable() bool {
	panic_if_version_older(2, 18, 0, "gtk_cell_renderer_toggle_get_activatable()")
	return gobool(C._gtk_cell_renderer_toggle_get_activatable(CELL_RENDERER_TOGGLE(v)))
}

func (v *CellRendererToggle) SetActivatable(activatable bool) {
	panic_if_version_older(2, 18, 0, "gtk_cell_renderer_toggle_set_activatable()")
	C._gtk_cell_renderer_toggle_set_activatable(CELL_RENDERER_TOGGLE(v), gbool(activatable))
}

//-----------------------------------------------------------------------
// GtkCellRendererSpinner
//-----------------------------------------------------------------------
type CellRendererSpinner struct {
	CellRenderer
}

func NewCellRendererSpinner() *CellRendererSpinner {
	panic_if_version_older(2, 20, 0, "gtk_cell_renderer_spinner_new()")
	return &CellRendererSpinner{CellRenderer{C._gtk_cell_renderer_spinner_new(), nil}}
}

//-----------------------------------------------------------------------
// GtkListStore
//-----------------------------------------------------------------------
const (
	TYPE_CHAR    = glib.G_TYPE_CHAR
	TYPE_UCHAR   = glib.G_TYPE_UCHAR
	TYPE_BOOL    = glib.G_TYPE_BOOL
	TYPE_INT     = glib.G_TYPE_INT
	TYPE_UINT    = glib.G_TYPE_UINT
	TYPE_LONG    = glib.G_TYPE_LONG
	TYPE_ULONG   = glib.G_TYPE_ULONG
	TYPE_FLOAT   = glib.G_TYPE_FLOAT
	TYPE_DOUBLE  = glib.G_TYPE_DOUBLE
	TYPE_STRING  = glib.G_TYPE_STRING
	TYPE_BOXED   = glib.G_TYPE_BOXED
	TYPE_POINTER = glib.G_TYPE_POINTER
	TYPE_PIXBUF  = TYPE_POINTER
)

type ListStore struct {
	TreeModel
	GListStore *C.GtkListStore
}

func NewListStore(v ...interface{}) *ListStore {
	types := C.make_gtypes(C.int(len(v)))
	for n := range v {
		C.set_gtype(types, C.int(n), C.int(v[n].(int)))
	}
	defer C.destroy_gtypes(types)
	cliststore := C.gtk_list_store_newv(gint(len(v)), types)
	return &ListStore{TreeModel{C.toGTreeModelFromListStore(cliststore)}, cliststore}
}

func NewListStoreFromNative(l unsafe.Pointer) *ListStore {
	cliststore := C.toGListStore(l)
	return &ListStore{TreeModel{C.toGTreeModelFromListStore(cliststore)}, cliststore}
}

//void gtk_list_store_set_column_types (GtkListStore *list_store, gint n_columns, GType *types);

func (v *ListStore) Set(iter *TreeIter, a ...interface{}) {
	for i := 0; i < len(a); i += 2 {
		v.SetValue(iter, a[i].(int), a[i+1])
	}
}

func (v *ListStore) SetValue(iter *TreeIter, column int, a interface{}) {
	iter.ready()
	gv := glib.GValueFromNative(a)
	if gv != nil {
		C.gtk_list_store_set_value(v.GListStore, iter.GTreeIter, gint(column), C.toGValue(unsafe.Pointer(gv)))
	} else {
		if pv, ok := a.(*[0]uint8); ok {
			C._gtk_list_store_set_ptr(v.GListStore, iter.GTreeIter, gint(column), unsafe.Pointer(pv))
		} else {
			av := reflect.ValueOf(a)
			if av.Kind() == reflect.Ptr {
				C._gtk_list_store_set_ptr(v.GListStore, iter.GTreeIter, gint(column), unsafe.Pointer(av.Pointer()))
			} else if av.CanAddr() {
				C._gtk_list_store_set_addr(v.GListStore, iter.GTreeIter, gint(column), unsafe.Pointer(av.UnsafeAddr()))
			} else {
				C._gtk_list_store_set_addr(v.GListStore, iter.GTreeIter, gint(column), unsafe.Pointer(&a))
			}
		}
	}
}

func (v *ListStore) Remove(iter *TreeIter) bool {
	iter.ready()
	return gobool(C.gtk_list_store_remove(v.GListStore, iter.GTreeIter))
}

func (v *ListStore) Insert(iter *TreeIter, position int) {
	iter.ready()
	C.gtk_list_store_insert(v.GListStore, iter.GTreeIter, gint(position))
}

func (v *ListStore) InsertBefore(iter *TreeIter, sibling *TreeIter) {
	iter.ready()
	sibling.ready()
	C.gtk_list_store_insert_before(v.GListStore, iter.GTreeIter, sibling.GTreeIter)
}

func (v *ListStore) InsertAfter(iter *TreeIter, sibling *TreeIter) {
	iter.ready()
	sibling.ready()
	C.gtk_list_store_insert_after(v.GListStore, iter.GTreeIter, sibling.GTreeIter)
}

//void gtk_list_store_insert_with_values (GtkListStore *list_store, GtkTreeIter *iter, gint position, ...);
//void gtk_list_store_insert_with_valuesv (GtkListStore *list_store, GtkTreeIter *iter, gint position, gint *columns, GValue *values, gint n_values);

func (v *ListStore) Prepend(iter *TreeIter) {
	iter.ready()
	C.gtk_list_store_prepend(v.GListStore, iter.GTreeIter)
}

func (v *ListStore) Append(iter *TreeIter) {
	iter.ready()
	C.gtk_list_store_append(v.GListStore, iter.GTreeIter)
}

func (v *ListStore) Clear() {
	C.gtk_list_store_clear(v.GListStore)
}

func (v *ListStore) IterIsValid(iter *TreeIter) bool {
	log.Println("Warning: ListStore.IterIsValid: This function is slow. Only use it for debugging and/or testing purposes.")
	iter.ready()
	return gobool(C.gtk_list_store_iter_is_valid(v.GListStore, iter.GTreeIter))
}

func (v *ListStore) Reorder(i *int) {
	gi := gint(*i)
	C.gtk_list_store_reorder(v.GListStore, &gi)
	*i = int(gi)
}

func (v *ListStore) Swap(a *TreeIter, b *TreeIter) {
	a.ready()
	b.ready()
	C.gtk_list_store_swap(v.GListStore, a.GTreeIter, b.GTreeIter)
}

func (v *ListStore) MoveBefore(iter *TreeIter, position *TreeIter) {
	iter.ready()
	position.ready()
	C.gtk_list_store_move_before(v.GListStore, iter.GTreeIter, position.GTreeIter)
}

func (v *ListStore) MoveAfter(iter *TreeIter, position *TreeIter) {
	iter.ready()
	position.ready()
	C.gtk_list_store_move_after(v.GListStore, iter.GTreeIter, position.GTreeIter)
}

//TODO instead of using this methods to change between treemodel and liststore, is better to usa an interface ITreeModel
//nb: ListStore e TreeStore sono un TreeModel (implementano GtkTreeModel!)
/*func (v *GtkListStore) ToTreeModel() *GtkTreeModel {
	return &TreeModel{
		C.toGTreeModelFromListStore(v.GListStore)}
}*/
/*func (v *GtkTreeModel) ToListStore() *GtkListStore {
	return &ListStore{
		C.toGListStoreFromTreeModel(v.GTreeModel)}
}*/

//-----------------------------------------------------------------------
// GtkTreeStore
//-----------------------------------------------------------------------
type TreeStore struct {
	TreeModel
	GTreeStore *C.GtkTreeStore
}

func NewTreeStore(v ...interface{}) *TreeStore {
	types := C.make_gtypes(C.int(len(v)))
	for n := range v {
		C.set_gtype(types, C.int(n), C.int(v[n].(int)))
	}
	defer C.destroy_gtypes(types)
	ctreestore := C.gtk_tree_store_newv(gint(len(v)), types)
	return &TreeStore{TreeModel{C.toGTreeModelFromTreeStore(ctreestore)}, ctreestore}
}

// void gtk_tree_store_set_column_types (GtkTreeStore *tree_store, gint n_columns, GType *types); void gtk_tree_store_set_value (GtkTreeStore *tree_store, GtkTreeIter *iter, gint column, GValue *value);

func (v *TreeStore) Set(iter *TreeIter, a ...interface{}) {
	for r := range a {
		v.SetValue(iter, r, a[r])
	}
}

func (v *TreeStore) SetValue(iter *TreeIter, column int, a interface{}) {
	iter.ready()
	gv := glib.GValueFromNative(a)
	if gv != nil {
		C.gtk_tree_store_set_value(v.GTreeStore, iter.GTreeIter, gint(column), C.toGValue(unsafe.Pointer(gv)))
	} else {
		if pv, ok := a.(*[0]uint8); ok {
			C._gtk_tree_store_set_ptr(v.GTreeStore, iter.GTreeIter, gint(column), unsafe.Pointer(pv))
		} else {
			av := reflect.ValueOf(a)
			if av.Kind() == reflect.Ptr {
				C._gtk_tree_store_set_ptr(v.GTreeStore, iter.GTreeIter, gint(column), unsafe.Pointer(av.Pointer()))
			} else if av.CanAddr() {
				C._gtk_tree_store_set_addr(v.GTreeStore, iter.GTreeIter, gint(column), unsafe.Pointer(av.UnsafeAddr()))
			} else {
				C._gtk_tree_store_set_addr(v.GTreeStore, iter.GTreeIter, gint(column), unsafe.Pointer(&a))
			}
		}
	}
}

func (v *TreeStore) Remove(iter *TreeIter) bool {
	iter.ready()
	return gobool(C.gtk_tree_store_remove(v.GTreeStore, iter.GTreeIter))
}

func (v *TreeStore) Insert(iter *TreeIter, parent *TreeIter, position int) {
	iter.ready()
	parent.ready()
	C.gtk_tree_store_insert(v.GTreeStore, iter.GTreeIter, parent.GTreeIter, gint(position))
}

func (v *TreeStore) InsertBefore(iter *TreeIter, parent *TreeIter, sibling *TreeIter) {
	iter.ready()
	parent.ready()
	sibling.ready()
	C.gtk_tree_store_insert_before(v.GTreeStore, iter.GTreeIter, parent.GTreeIter, sibling.GTreeIter)
}

func (v *TreeStore) InsertAfter(iter *TreeIter, parent *TreeIter, sibling *TreeIter) {
	iter.ready()
	parent.ready()
	sibling.ready()
	C.gtk_tree_store_insert_after(v.GTreeStore, iter.GTreeIter, parent.GTreeIter, sibling.GTreeIter)
}

// void gtk_tree_store_insert_with_values (GtkTreeStore *tree_store, GtkTreeIter *iter, GtkTreeIter *parent, gint position, ...);
// void gtk_tree_store_insert_with_valuesv (GtkTreeStore *tree_store, GtkTreeIter *iter, GtkTreeIter *parent, gint position, gint *columns, GValue *values, gint n_values);

func (v *TreeStore) Prepend(iter *TreeIter, parent *TreeIter) {
	iter.ready()
	if parent == nil {
		C.gtk_tree_store_prepend(v.GTreeStore, iter.GTreeIter, nil)
	} else {
		parent.ready()
		C.gtk_tree_store_prepend(v.GTreeStore, iter.GTreeIter, parent.GTreeIter)
	}
}

func (v *TreeStore) Append(iter *TreeIter, parent *TreeIter) {
	iter.ready()
	if parent == nil {
		C.gtk_tree_store_append(v.GTreeStore, iter.GTreeIter, nil)
	} else {
		parent.ready()
		C.gtk_tree_store_append(v.GTreeStore, iter.GTreeIter, parent.GTreeIter)
	}
}

// gtk_tree_store_is_ancestor

func (v *TreeStore) ToTreeModel() *TreeModel {
	return &TreeModel{C.toGTreeModelFromTreeStore(v.GTreeStore)}
}

func (v *TreeStore) IterDepth(iter *TreeIter) int {
	iter.ready()
	return int(C.gtk_tree_store_iter_depth(v.GTreeStore, iter.GTreeIter))
}

func (v *TreeStore) Clear() {
	C.gtk_tree_store_clear(v.GTreeStore)
}

func (v *TreeStore) IterIsValid(iter *TreeIter) bool {
	log.Println("Warning: TreeStore.IterIsValid: This function is slow. Only use it for debugging and/or testing purposes.")
	iter.ready()
	return gobool(C.gtk_tree_store_iter_is_valid(v.GTreeStore, iter.GTreeIter))
}

func (v *TreeStore) Reorder(iter *TreeIter, i *int) {
	gi := gint(*i)
	iter.ready()
	C.gtk_tree_store_reorder(v.GTreeStore, iter.GTreeIter, &gi)
	*i = int(gi)
}

func (v *TreeStore) Swap(a *TreeIter, b *TreeIter) {
	a.ready()
	b.ready()
	C.gtk_tree_store_swap(v.GTreeStore, a.GTreeIter, b.GTreeIter)
}

func (v *TreeStore) MoveBefore(iter *TreeIter, position *TreeIter) {
	iter.ready()
	position.ready()
	C.gtk_tree_store_move_before(v.GTreeStore, iter.GTreeIter, position.GTreeIter)
}

func (v *TreeStore) MoveAfter(iter *TreeIter, position *TreeIter) {
	iter.ready()
	position.ready()
	C.gtk_tree_store_move_after(v.GTreeStore, iter.GTreeIter, position.GTreeIter)
}

//-----------------------------------------------------------------------
// GtkComboBox
//-----------------------------------------------------------------------
type ComboBox struct {
	Bin
}

func NewComboBox() *ComboBox {
	return &ComboBox{Bin{Container{Widget{C.gtk_combo_box_new()}}}}
}

func NewComboBoxWithEntry() *ComboBox {
	deprecated_since(2, 24, 0, "gtk_combo_box_new_with_entry()")
	return &ComboBox{Bin{Container{Widget{C._gtk_combo_box_new_with_entry()}}}}
}

func NewComboBoxWithModel(model *TreeModel) *ComboBox {
	return &ComboBox{Bin{Container{Widget{C.gtk_combo_box_new_with_model(model.GTreeModel)}}}}
}

func NewComboBoxWithModelAndEntry(model *TreeModel) *ComboBox {
	deprecated_since(2, 24, 0, "gtk_combo_box_new_with_model_and_entry()")
	return &ComboBox{Bin{Container{Widget{C._gtk_combo_box_new_with_model_and_entry(model.GTreeModel)}}}}
}

func (v *ComboBox) GetWrapWidth() int {
	return int(C.gtk_combo_box_get_wrap_width(COMBO_BOX(v)))
}

func (v *ComboBox) SetWrapWidth(width int) {
	C.gtk_combo_box_set_wrap_width(COMBO_BOX(v), gint(width))
}

func (v *ComboBox) GetRowSpanColumn() int {
	return int(C.gtk_combo_box_get_row_span_column(COMBO_BOX(v)))
}

func (v *ComboBox) SetRowSpanColumn(row_span int) {
	C.gtk_combo_box_set_row_span_column(COMBO_BOX(v), gint(row_span))
}

func (v *ComboBox) GetColumnSpanColumn() int {
	return int(C.gtk_combo_box_get_column_span_column(COMBO_BOX(v)))
}

func (v *ComboBox) SetColumnSpanColumn(column_span int) {
	C.gtk_combo_box_set_column_span_column(COMBO_BOX(v), gint(column_span))
}

func (v *ComboBox) GetActive() int {
	return int(C.gtk_combo_box_get_active(COMBO_BOX(v)))
}

func (v *ComboBox) SetActive(width int) {
	C.gtk_combo_box_set_active(COMBO_BOX(v), gint(width))
}

func (v *ComboBox) GetActiveIter(iter *TreeIter) bool {
	iter.ready()
	return gobool(C.gtk_combo_box_get_active_iter(COMBO_BOX(v), iter.GTreeIter))
}

func (v *ComboBox) SetActiveIter(iter *TreeIter) {
	iter.ready()
	C.gtk_combo_box_set_active_iter(COMBO_BOX(v), iter.GTreeIter)
}

func (v *ComboBox) GetModel() *TreeModel {
	return &TreeModel{
		C.gtk_combo_box_get_model(COMBO_BOX(v))}
}

func (v *ComboBox) SetModel(model *TreeModel) {
	C.gtk_combo_box_set_model(COMBO_BOX(v), model.GTreeModel)
}

//Deprecated since 2.24. Use GtkComboBoxText.
func NewComboBoxNewText() *ComboBox {
	deprecated_since(2, 24, 0, "gtk_combo_box_new_text()")
	return &ComboBox{Bin{Container{Widget{C.gtk_combo_box_new_text()}}}}
}

//Deprecated since 2.24. Use GtkComboBoxText.
func (v *ComboBox) AppendText(text string) {
	deprecated_since(2, 24, 0, "gtk_combo_box_append_text()")
	ptr := C.CString(text)
	defer cfree(ptr)
	C.gtk_combo_box_append_text(COMBO_BOX(v), gstring(ptr))
}

//Deprecated since 2.24. Use GtkComboBoxText.
func (v *ComboBox) InsertText(text string, position int) {
	deprecated_since(2, 24, 0, "gtk_combo_box_insert_text()")
	ptr := C.CString(text)
	defer cfree(ptr)
	C.gtk_combo_box_insert_text(COMBO_BOX(v), gint(position), gstring(ptr))
}

//Deprecated since 2.24. Use GtkComboBoxText.
func (v *ComboBox) PrependText(text string) {
	deprecated_since(2, 24, 0, "gtk_combo_box_prepend_text()")
	ptr := C.CString(text)
	defer cfree(ptr)
	C.gtk_combo_box_prepend_text(COMBO_BOX(v), gstring(ptr))
}

//Deprecated since 2.24. Use GtkComboBoxText.
func (v *ComboBox) RemoveText(position int) {
	deprecated_since(2, 24, 0, "gtk_combo_box_remove_text()")
	C.gtk_combo_box_remove_text(COMBO_BOX(v), gint(position))
}

//Deprecated since 2.24. Use GtkComboBoxText or, if combo box contains an entry,
// get text directly from GtkEntry.
func (v *ComboBox) GetActiveText() string {
	deprecated_since(2, 24, 0, "gtk_combo_box_get_active_text()")
	return gostring(C.gtk_combo_box_get_active_text(COMBO_BOX(v)))
}

func (v *ComboBox) Popup() {
	C.gtk_combo_box_popup(COMBO_BOX(v))
}

func (v *ComboBox) Popdown() {
	C.gtk_combo_box_popdown(COMBO_BOX(v))
}

// gtk_combo_box_get_popup_accessible
// gtk_combo_box_get_row_separator_func
// gtk_combo_box_set_row_separator_func

func (v *ComboBox) SetAddTearoffs(add_tearoffs bool) {
	C.gtk_combo_box_set_add_tearoffs(COMBO_BOX(v), gbool(add_tearoffs))
}

func (v *ComboBox) GetAddTearoffs() bool {
	return gobool(C.gtk_combo_box_get_add_tearoffs(COMBO_BOX(v)))
}

func (v *ComboBox) SetTitle(title string) {
	ptr := C.CString(title)
	defer cfree(ptr)
	C.gtk_combo_box_set_title(COMBO_BOX(v), gstring(ptr))
}

func (v *ComboBox) GetTitle() string {
	return gostring(C.gtk_combo_box_get_title(COMBO_BOX(v)))
}

func (v *ComboBox) SetFocusOnClick(focus_on_click bool) {
	C.gtk_combo_box_set_focus_on_click(COMBO_BOX(v), gbool(focus_on_click))
}

func (v *ComboBox) GetFocusOnClick() bool {
	return gobool(C.gtk_combo_box_get_focus_on_click(COMBO_BOX(v)))
}

// gtk_combo_box_set_button_sensitivity
// gtk_combo_box_get_button_sensitivity
// gtk_combo_box_get_has_entry //since 2.24
// gtk_combo_box_set_entry_text_column //since 2.24
// gtk_combo_box_get_entry_text_column //since 2.24

//-----------------------------------------------------------------------
// GtkComboBoxText
//-----------------------------------------------------------------------
type ComboBoxText struct {
	ComboBox
}

func NewComboBoxText() *ComboBoxText {
	return &ComboBoxText{ComboBox{Bin{Container{Widget{C._gtk_combo_box_text_new()}}}}}
}

func NewComboBoxTextWithEntry() *ComboBoxText {
	return &ComboBoxText{ComboBox{Bin{Container{Widget{C._gtk_combo_box_text_new_with_entry()}}}}}
}

func (v *ComboBoxText) AppendText(text string) {
	ptr := C.CString(text)
	defer cfree(ptr)
	C._gtk_combo_box_text_append_text(COMBO_BOX_TEXT(v), gstring(ptr))
}

func (v *ComboBoxText) InsertText(position int, text string) {
	ptr := C.CString(text)
	defer cfree(ptr)
	C._gtk_combo_box_text_insert_text(COMBO_BOX_TEXT(v), gint(position), gstring(ptr))
}

func (v *ComboBoxText) PrependText(text string) {
	ptr := C.CString(text)
	defer cfree(ptr)
	C._gtk_combo_box_text_prepend_text(COMBO_BOX_TEXT(v), gstring(ptr))
}

func (v *ComboBoxText) Remove(position int) {
	C._gtk_combo_box_text_remove(COMBO_BOX_TEXT(v), gint(position))
}

func (v *ComboBoxText) GetActiveText() string {
	return gostring(C._gtk_combo_box_text_get_active_text(COMBO_BOX_TEXT(v)))
}

//-----------------------------------------------------------------------
// GtkComboBoxEntry
//-----------------------------------------------------------------------
type ComboBoxEntry struct {
	ComboBox
}

func NewComboBoxEntry() *ComboBoxEntry {
	deprecated_since(2, 24, 0, "gtk_combo_box_entry_new()")
	return &ComboBoxEntry{ComboBox{Bin{Container{Widget{C.gtk_combo_box_entry_new()}}}}}
}

// gtk_combo_box_entry_new_with_model

func NewComboBoxEntryNewText() *ComboBoxEntry {
	deprecated_since(2, 24, 0, "gtk_combo_box_entry_new_text()")
	return &ComboBoxEntry{ComboBox{Bin{Container{Widget{C.gtk_combo_box_entry_new_text()}}}}}
}

func (v *ComboBoxEntry) GetTextColumn() int {
	deprecated_since(2, 24, 0, "gtk_combo_box_entry_get_text_column()")
	return int(C.gtk_combo_box_entry_get_text_column(COMBO_BOX_ENTRY(v)))
}

func (v *ComboBoxEntry) SetTextColumn(text_column int) {
	deprecated_since(2, 24, 0, "gtk_combo_box_entry_set_text_column()")
	C.gtk_combo_box_entry_set_text_column(COMBO_BOX_ENTRY(v), gint(text_column))
}

//-----------------------------------------------------------------------
// GtkMenu
//-----------------------------------------------------------------------
type Menu struct {
	Container
	//TODO GtkMenuShell
}

func NewMenu() *Menu {
	return &Menu{Container{Widget{C.gtk_menu_new()}}}
}

// void gtk_menu_set_screen (GtkMenu *menu, GdkScreen *screen);

//TODO remove when GtkMenuShell is done
func (v *Menu) Append(child IWidget) {
	C.gtk_menu_shell_append(MENU_SHELL(v), ToNative(child))
}

//TODO remove when GtkMenuShell is done
func (v *Menu) Prepend(child IWidget) {
	C.gtk_menu_shell_prepend(MENU_SHELL(v), ToNative(child))
}

//TODO remove when GtkMenuShell is done
func (v *Menu) Insert(child IWidget, position int) {
	C.gtk_menu_shell_insert(MENU_SHELL(v), ToNative(child), gint(position))
}

// void gtk_menu_reorder_child(GtkMenu *menu, GtkWidget *child, gint position);
// void gtk_menu_attach(GtkMenu *menu, GtkWidget *child, guint left_attach, guint right_attach, guint top_attach, guint bottom_attach);

func (v *Menu) Popup(parent_menu_shell, parent_menu_item IWidget, f MenuPositionFunc, data interface{}, button uint, active_item uint32) {
	var pms, pmi *C.GtkWidget
	if parent_menu_shell != nil {
		pms = ToNative(parent_menu_shell)
	}
	if parent_menu_item != nil {
		pmi = ToNative(parent_menu_item)
	}
	C._gtk_menu_popup(v.GWidget, pms, pmi, pointer.Save(&MenuPositionFuncInfo{v, f, data}), guint(button), guint32(active_item))
}

// void gtk_menu_set_accel_group (GtkMenu *menu, GtkAccelGroup *accel_group);
// GtkAccelGroup* gtk_menu_get_accel_group (GtkMenu *menu);
// void gtk_menu_set_accel_path(GtkMenu *menu, const gchar *accel_path);
// const gchar* gtk_menu_get_accel_path(GtkMenu *menu);
// void gtk_menu_set_title(GtkMenu *menu, const gchar *title);
// G_CONST_RETURN gchar *gtk_menu_get_title(GtkMenu *menu);
// void gtk_menu_set_monitor(GtkMenu *menu, gint monitor_num);
// gint gtk_menu_get_monitor(GtkMenu *menu);

func (v *Menu) GetTearoffState() bool {
	return gobool(C.gtk_menu_get_tearoff_state(MENU(v)))
}

func (v *Menu) SetReserveToggleSize(b bool) {
	panic_if_version_older(2, 18, 0, "gtk_menu_set_reserve_toggle_size()")
	C._gtk_menu_set_reserve_toggle_size(MENU(v), gbool(b))
}

func (v *Menu) GetReserveToggleSize() bool {
	panic_if_version_older(2, 18, 0, "gtk_menu_get_reserve_toggle_size()")
	return gobool(C._gtk_menu_get_reserve_toggle_size(MENU(v)))
}

func (v *Menu) Popdown() {
	C.gtk_menu_popdown(MENU(v))
}

func (v *Menu) Reposition() {
	C.gtk_menu_reposition(MENU(v))
}

func (v *Menu) GetActive() *Widget {
	return &Widget{C.gtk_menu_get_active(MENU(v))}
}

// void gtk_menu_set_active (GtkMenu *menu, guint index_);

func (v *Menu) SetTearoffState(b bool) {
	C.gtk_menu_set_tearoff_state(MENU(v), gbool(b))
}

// void gtk_menu_attach_to_widget (GtkMenu *menu, GtkWidget *attach_widget, GtkMenuDetachFunc detacher);

func (v *Menu) Detach() {
	C.gtk_menu_detach(MENU(v))
}

func (v *Menu) GetAttachWidget() *Widget {
	return &Widget{C.gtk_menu_get_attach_widget(MENU(v))}
}

// GList* gtk_menu_get_for_attach_widget(GtkWidget *widget);

type MenuPositionFunc func(menu *Menu, px, py *int, push_in *bool, data interface{})
type MenuPositionFuncInfo struct {
	menu *Menu
	f    MenuPositionFunc
	data interface{}
}

//export _go_gtk_menu_position_func
func _go_gtk_menu_position_func(gmpfi *C._gtk_menu_position_func_info) {
	if gmpfi == nil {
		return
	}
	gmpfigo := pointer.Restore(unsafe.Pointer(gmpfi.data)).(*MenuPositionFuncInfo)
	if gmpfigo.f == nil {
		return
	}
	x := int(gmpfi.x)
	y := int(gmpfi.y)
	push_in := gobool(gmpfi.push_in)
	gmpfigo.f(gmpfigo.menu, &x, &y, &push_in, gmpfigo.data)
	gmpfi.x = gint(x)
	gmpfi.y = gint(y)
	gmpfi.push_in = gbool(push_in)
}

//-----------------------------------------------------------------------
// GtkMenuBar
//-----------------------------------------------------------------------
type PackDirection int

const (
	PACK_DIRECTION_LTR PackDirection = iota
	PACK_DIRECTION_RTL
	PACK_DIRECTION_TTB
	PACK_DIRECTION_BTT
)

type MenuBar struct {
	Widget
}

func NewMenuBar() *MenuBar {
	return &MenuBar{Widget{C.gtk_menu_bar_new()}}
}

func (v *MenuBar) SetPackDirection(pack_dir PackDirection) {
	C.gtk_menu_bar_set_pack_direction(MENU_BAR(v), C.GtkPackDirection(pack_dir))
}

func (v *MenuBar) GetPackDirection() PackDirection {
	return PackDirection(C.gtk_menu_bar_get_pack_direction(MENU_BAR(v)))
}

func (v *MenuBar) SetChildPackDirection(pack_dir PackDirection) {
	C.gtk_menu_bar_set_child_pack_direction(MENU_BAR(v), C.GtkPackDirection(pack_dir))
}

func (v *MenuBar) GetChildPackDirection() PackDirection {
	return PackDirection(C.gtk_menu_bar_get_child_pack_direction(MENU_BAR(v)))
}

//TODO da rimuovere, creare GtkMenuShell e usarlo come anonymous field per GtkMenu
func (v *MenuBar) Append(child IWidget) {
	C.gtk_menu_shell_append(MENU_BAR_SHELL(v), ToNative(child))
}

//TODO da rimuovere, creare GtkMenuShell e usarlo come anonymous field per GtkMenu
func (v *MenuBar) Prepend(child IWidget) {
	C.gtk_menu_shell_prepend(MENU_BAR_SHELL(v), ToNative(child))
}

//TODO da rimuovere, creare GtkMenuShell e usarlo come anonymous field per GtkMenu
func (v *MenuBar) Insert(child IWidget, position int) {
	C.gtk_menu_shell_insert(MENU_BAR_SHELL(v), ToNative(child), gint(position))
}

//-----------------------------------------------------------------------
// GtkMenuItem
//-----------------------------------------------------------------------
type MenuItem struct {
	Item
	Activatable // implement GtkActivatable interface
}

func newMenuItemInternal(widget *C.GtkWidget) *MenuItem {
	return &MenuItem{Item{Bin{Container{Widget{widget}}}}, Activatable{widget}}
}

func NewMenuItem() *MenuItem {
	return newMenuItemInternal(C.gtk_menu_item_new())
}

func NewMenuItemWithLabel(label string) *MenuItem {
	var ptr *C.char
	if len(label) > 0 {
		ptr = C.CString(label)
		defer cfree(ptr)
	}
	return newMenuItemInternal(C.gtk_menu_item_new_with_label(
		gstring(ptr)))
}

func NewMenuItemWithMnemonic(label string) *MenuItem {
	var ptr *C.char
	if len(label) > 0 {
		ptr = C.CString(label)
		defer cfree(ptr)
	}
	return newMenuItemInternal(C.gtk_menu_item_new_with_mnemonic(
		gstring(ptr)))
}

func (v *MenuItem) SetRightJustified(b bool) {
	C.gtk_menu_item_set_right_justified(MENU_ITEM(v), gbool(b))
}

func (v *MenuItem) GetRightJustified() bool {
	return gobool(C.gtk_menu_item_get_right_justified(MENU_ITEM(v)))
}

func (v *MenuItem) GetLabel() string {
	return gostring(C.gtk_menu_item_get_label(MENU_ITEM(v)))
}

func (v *MenuItem) SetLabel(label string) {
	var ptr *C.char
	if len(label) > 0 {
		ptr = C.CString(label)
		defer cfree(ptr)
	}
	C.gtk_menu_item_set_label(MENU_ITEM(v), gstring(ptr))
}

func (v *MenuItem) GetUseUnderline() bool {
	return gobool(C.gtk_menu_item_get_use_underline(MENU_ITEM(v)))
}

func (v *MenuItem) SetUseUnderline(setting bool) {
	C.gtk_menu_item_set_use_underline(MENU_ITEM(v), gbool(setting))
}

func (v *MenuItem) SetSubmenu(w IWidget) {
	C.gtk_menu_item_set_submenu(MENU_ITEM(v), ToNative(w))
}

func (v *MenuItem) GetSubmenu() *Widget {
	return &Widget{C.gtk_menu_item_get_submenu(MENU_ITEM(v))}
}

//Deprecated since 2.12. Use SetSubmenu() instead.
func (v *MenuItem) RemoveSubmenu() {
	deprecated_since(2, 12, 0, "gtk_menu_item_remove_submenu()")
	C.gtk_menu_item_remove_submenu(MENU_ITEM(v))
}

// void gtk_menu_item_set_accel_path(GtkMenuItem *menu_item, const gchar *accel_path);
// G_CONST_RETURN gchar* gtk_menu_item_get_accel_path(GtkMenuItem *menu_item);

func (v *MenuItem) Select() {
	C.gtk_menu_item_select(MENU_ITEM(v))
}

func (v *MenuItem) Deselect() {
	C.gtk_menu_item_deselect(MENU_ITEM(v))
}

func (v *MenuItem) Activate() {
	C.gtk_menu_item_activate(MENU_ITEM(v))
}

func (v *MenuItem) ToggleSizeRequest(i *int) {
	gi := gint(*i)
	C.gtk_menu_item_toggle_size_request(MENU_ITEM(v), &gi)
	*i = int(gi)
}

func (v *MenuItem) ToggleSizeAllocate(i int) {
	C.gtk_menu_item_toggle_size_allocate(MENU_ITEM(v), gint(i))
}

//-----------------------------------------------------------------------
// GtkImageMenuItem
//-----------------------------------------------------------------------

type ImageMenuItem struct {
	MenuItem
}

func NewImageMenuItem() *ImageMenuItem {
	return &ImageMenuItem{*newMenuItemInternal(
		C.gtk_image_menu_item_new())}
}

func NewImageMenuItemFromStock(stock_id string, accel_group *AccelGroup) *ImageMenuItem {
	p := C.CString(stock_id)
	defer cfree(p)
	return &ImageMenuItem{*newMenuItemInternal(
		C.gtk_image_menu_item_new_from_stock(gstring(p), accel_group.GAccelGroup))}
}

func NewImageMenuItemWithLabel(label string) *ImageMenuItem {
	p := C.CString(label)
	defer cfree(p)
	return &ImageMenuItem{*newMenuItemInternal(
		C.gtk_image_menu_item_new_with_label(gstring(p)))}
}

func NewImageMenuItemWithMnemonic(label string) *ImageMenuItem {
	p := C.CString(label)
	defer cfree(p)
	return &ImageMenuItem{*newMenuItemInternal(
		C.gtk_image_menu_item_new_with_mnemonic(gstring(p)))}
}

func (v *ImageMenuItem) SetImage(image *Widget) {
	C.gtk_image_menu_item_set_image(IMAGE_MENU_ITEM(v), WIDGET(image))
}

func (v *ImageMenuItem) GetImage() *Widget {
	return &Widget{C.gtk_image_menu_item_get_image(IMAGE_MENU_ITEM(v))}
}

// gtk_image_menu_item_get_use_stock
// gtk_image_menu_item_set_use_stock
// gtk_image_menu_item_get_always_show_image
// gtk_image_menu_item_set_always_show_image
// gtk_image_menu_item_set_accel_group

//-----------------------------------------------------------------------
// GtkRadioMenuItem
//-----------------------------------------------------------------------
type RadioMenuItem struct {
	CheckMenuItem
}

func NewRadioMenuItem(group *glib.SList) *RadioMenuItem {
	return &RadioMenuItem{CheckMenuItem{*newMenuItemInternal(
		C.gtk_radio_menu_item_new(gslist(group)))}}
}

func NewRadioMenuItemWithLabel(group *glib.SList, label string) *RadioMenuItem {
	var ptr *C.char
	if len(label) > 0 {
		ptr = C.CString(label)
		defer cfree(ptr)
	}
	return &RadioMenuItem{CheckMenuItem{*newMenuItemInternal(
		C.gtk_radio_menu_item_new_with_label(gslist(group), gstring(ptr)))}}
}

// gtk_radio_menu_item_new_with_mnemonic
// gtk_radio_menu_item_new_from_widget
// gtk_radio_menu_item_new_with_label_from_widget
// gtk_radio_menu_item_new_with_mnemonic_from_widget
// gtk_radio_menu_item_group

func (v *RadioMenuItem) SetGroup(group *glib.SList) {
	C.gtk_radio_menu_item_set_group(RADIO_MENU_ITEM(v), gslist(group))
}

func (v *RadioMenuItem) GetGroup() *glib.SList {
	return glib.SListFromNative(unsafe.Pointer(
		C.gtk_radio_menu_item_get_group(RADIO_MENU_ITEM(v))))
}

//-----------------------------------------------------------------------
// GtkCheckMenuItem
//-----------------------------------------------------------------------
type CheckMenuItem struct {
	MenuItem
}

func NewCheckMenuItem() *CheckMenuItem {
	return &CheckMenuItem{*newMenuItemInternal(
		C.gtk_check_menu_item_new())}
}

func NewCheckMenuItemWithLabel(label string) *CheckMenuItem {
	var ptr *C.char
	if len(label) > 0 {
		ptr = C.CString(label)
		defer cfree(ptr)
	}
	return &CheckMenuItem{*newMenuItemInternal(
		C.gtk_check_menu_item_new_with_label(gstring(ptr)))}
}

func NewCheckMenuItemWithMnemonic(label string) *CheckMenuItem {
	var ptr *C.char
	if len(label) > 0 {
		ptr = C.CString(label)
		defer cfree(ptr)
	}
	return &CheckMenuItem{*newMenuItemInternal(
		C.gtk_check_menu_item_new_with_mnemonic(gstring(ptr)))}
}

func (v *CheckMenuItem) GetActive() bool {
	return gobool(C.gtk_check_menu_item_get_active(CHECK_MENU_ITEM(v)))
}

func (v *CheckMenuItem) SetActive(setting bool) {
	C.gtk_check_menu_item_set_active(CHECK_MENU_ITEM(v), gbool(setting))
}

func (v *CheckMenuItem) Toggled() {
	C.gtk_check_menu_item_toggled(CHECK_MENU_ITEM(v))
}

func (v *CheckMenuItem) GetInconsistent() bool {
	return gobool(C.gtk_check_menu_item_get_inconsistent(CHECK_MENU_ITEM(v)))
}

func (v *CheckMenuItem) SetInconsistent(setting bool) {
	C.gtk_check_menu_item_set_inconsistent(CHECK_MENU_ITEM(v), gbool(setting))
}

func (v *CheckMenuItem) SetDrawAsRadio(setting bool) {
	C.gtk_check_menu_item_set_draw_as_radio(CHECK_MENU_ITEM(v), gbool(setting))
}

func (v *CheckMenuItem) GetDrawAsRadio() bool {
	return gobool(C.gtk_check_menu_item_get_draw_as_radio(CHECK_MENU_ITEM(v)))
}

//-----------------------------------------------------------------------
// GtkSeparatorMenuItem
//-----------------------------------------------------------------------
type SeparatorMenuItem struct {
	MenuItem
}

func NewSeparatorMenuItem() *SeparatorMenuItem {
	return &SeparatorMenuItem{*newMenuItemInternal(
		C.gtk_separator_menu_item_new())}
}

//-----------------------------------------------------------------------
// GtkTearoffMenuItem
//-----------------------------------------------------------------------
type TearoffMenuItem struct {
	MenuItem
}

func NewTearoffMenuItem() *TearoffMenuItem {
	return &TearoffMenuItem{*newMenuItemInternal(
		C.gtk_tearoff_menu_item_new())}
}

//-----------------------------------------------------------------------
// GtkToolShell
//-----------------------------------------------------------------------

// gtk_tool_shell_get_ellipsize_mode
// gtk_tool_shell_get_icon_size
// gtk_tool_shell_get_orientation
// gtk_tool_shell_get_relief_style
// gtk_tool_shell_get_style
// gtk_tool_shell_get_text_alignment
// gtk_tool_shell_get_text_orientation
// gtk_tool_shell_rebuild_menu
// gtk_tool_shell_get_text_size_group

//-----------------------------------------------------------------------
// GtkToolbar
//-----------------------------------------------------------------------

type Orientation int

const (
	ORIENTATION_HORIZONTAL Orientation = iota
	ORIENTATION_VERTICAL
)

type ToolbarStyle int

const (
	TOOLBAR_ICONS ToolbarStyle = iota
	TOOLBAR_TEXT
	TOOLBAR_BOTH
	TOOLBAR_BOTH_HORIZ
)

type Toolbar struct {
	Container
	items map[*C.GtkToolItem]IWidget
}

func NewToolbar() *Toolbar {
	return &Toolbar{Container{Widget{C.gtk_toolbar_new()}}, make(map[*C.GtkToolItem]IWidget)}
}

func (v *Toolbar) OnFocusHomeOrEnd(onclick interface{}, datas ...interface{}) int {
	return v.Connect("focus-home-or-end", onclick, datas...)
}

func (v *Toolbar) OnOrientationChanged(onclick interface{}, datas ...interface{}) int {
	return v.Connect("orientation-changed", onclick, datas...)
}

func (v *Toolbar) OnPopupContextMenu(onclick interface{}, datas ...interface{}) int {
	return v.Connect("popup-context-menu", onclick, datas...)
}

func (v *Toolbar) OnStyleChanged(onclick interface{}, datas ...interface{}) int {
	return v.Connect("style-changed", onclick, datas...)
}

func (v *Toolbar) Insert(item IWidget, pos int) {
	p_tool_item := C.toGToolItem(ToNative(item))
	v.items[p_tool_item] = item
	C.gtk_toolbar_insert(TOOLBAR(v), p_tool_item, gint(pos))
}

func (v *Toolbar) GetItemIndex(item IWidget) int {
	return int(C.gtk_toolbar_get_item_index(TOOLBAR(v), C.toGToolItem(ToNative(item))))
}

func (v *Toolbar) GetNItems() int {
	return int(C.gtk_toolbar_get_n_items(TOOLBAR(v)))
}

func (v *Toolbar) GetNthItem(n int) IWidget {
	p_tool_item := C.gtk_toolbar_get_nth_item(TOOLBAR(v), gint(n))
	if p_tool_item == nil {
		panic("Toolbar.GetNthItem: index out of bounds")
	}
	if _, ok := v.items[p_tool_item]; !ok {
		panic("Toolbar.GetNthItem: interface not found in map")
	}
	return v.items[p_tool_item]
}

func (v *Toolbar) GetDropIndex(x, y int) int {
	return int(C.gtk_toolbar_get_drop_index(TOOLBAR(v), gint(x), gint(y)))
}

func (v *Toolbar) SetDropHighlightItem(item IWidget, index int) {
	C.gtk_toolbar_set_drop_highlight_item(TOOLBAR(v), C.toGToolItem(ToNative(item)), gint(index))
}

func (v *Toolbar) SetShowArrow(show_arrow bool) {
	C.gtk_toolbar_set_show_arrow(TOOLBAR(v), gbool(show_arrow))
}

func (v *Toolbar) SetOrientation(orientation Orientation) {
	C.gtk_toolbar_set_orientation(TOOLBAR(v), C.GtkOrientation(orientation))
}

func (v *Toolbar) SetTooltips(enable bool) {
	C.gtk_toolbar_set_tooltips(TOOLBAR(v), gbool(enable))
}

func (v *Toolbar) UnsetIconSize() {
	C.gtk_toolbar_unset_icon_size(TOOLBAR(v))
}

func (v *Toolbar) GetShowArrow() bool {
	return gobool(C.gtk_toolbar_get_show_arrow(TOOLBAR(v)))
}

func (v *Toolbar) GetOrientation() Orientation {
	return Orientation(C.gtk_toolbar_get_orientation(TOOLBAR(v)))
}

func (v *Toolbar) GetStyle() ToolbarStyle {
	return ToolbarStyle(C.gtk_toolbar_get_style(TOOLBAR(v)))
}

func (v *Toolbar) GetIconSize() IconSize {
	return IconSize(C.gtk_toolbar_get_icon_size(TOOLBAR(v)))
}

func (v *Toolbar) GetTooltips() bool {
	return gobool(C.gtk_toolbar_get_tooltips(TOOLBAR(v)))
}

func (v *Toolbar) GetReliefStyle() ReliefStyle {
	return ReliefStyle(C.gtk_toolbar_get_relief_style(TOOLBAR(v)))
}

// gtk_toolbar_append_item
// gtk_toolbar_prepend_item
// gtk_toolbar_insert_item
// gtk_toolbar_append_space
// gtk_toolbar_prepend_space
// gtk_toolbar_insert_space
// gtk_toolbar_append_element
// gtk_toolbar_prepend_element
// gtk_toolbar_insert_element
// gtk_toolbar_append_widget
// gtk_toolbar_prepend_widget
// gtk_toolbar_insert_widget
func (v *Toolbar) SetStyle(style ToolbarStyle) {
	C.gtk_toolbar_set_style(TOOLBAR(v), C.GtkToolbarStyle(style))
}

// gtk_toolbar_insert_stock
func (v *Toolbar) SetIconSize(icon_size IconSize) {
	C.gtk_toolbar_set_icon_size(TOOLBAR(v), C.GtkIconSize(icon_size))
}

// gtk_toolbar_remove_space
func (v *Toolbar) UnsetStyle() {
	C.gtk_toolbar_unset_style(TOOLBAR(v))
}

//-----------------------------------------------------------------------
// GtkToolItem
//-----------------------------------------------------------------------

type ToolItem struct {
	Bin
	Activatable // implement GtkActivatable interface
}

func newToolItemInternal(widget *C.GtkWidget) *ToolItem {
	return &ToolItem{Bin{Container{Widget{widget}}}, Activatable{widget}}
}

func NewToolItem() *ToolItem {
	return newToolItemInternal(AS_GWIDGET(
		unsafe.Pointer(C.gtk_tool_item_new())))
}

func (v *ToolItem) OnCreateMenuProxy(onclick interface{}, datas ...interface{}) int {
	return v.Connect("create-menu-proxy", onclick, datas...)
}

func (v *ToolItem) OnSetTooltip(onclick interface{}, datas ...interface{}) int {
	return v.Connect("set-tooltip", onclick, datas...)
}

func (v *ToolItem) OnToolbarReconfigured(onclick interface{}, datas ...interface{}) int {
	return v.Connect("toolbar-reconfigured", onclick, datas...)
}

func (v *ToolItem) SetHomogeneous(homogeneous bool) {
	C.gtk_tool_item_set_homogeneous(TOOL_ITEM(v), gbool(homogeneous))
}

func (v *ToolItem) GetHomogeneous() bool {
	return gobool(C.gtk_tool_item_get_homogeneous(TOOL_ITEM(v)))
}

func (v *ToolItem) SetExpand(expand bool) {
	C.gtk_tool_item_set_expand(TOOL_ITEM(v), gbool(expand))
}

func (v *ToolItem) GetExpand() bool {
	return gobool(C.gtk_tool_item_get_expand(TOOL_ITEM(v)))
}

// gtk_tool_item_set_tooltip (deprecated since 2.12)
func (v *ToolItem) SetArrowTooltipText(text string) {
	pt := C.CString(text)
	defer cfree(pt)
	C.gtk_tool_item_set_tooltip_text(TOOL_ITEM(v), gstring(pt))
}

func (v *ToolItem) SetArrowTooltipMarkup(markup string) {
	pm := C.CString(markup)
	defer cfree(pm)
	C.gtk_tool_item_set_tooltip_markup(TOOL_ITEM(v), gstring(pm))
}

func (v *ToolItem) SetTooltipMarkup(markup string) {
	p_markup := C.CString(markup)
	defer cfree(p_markup)
	C.gtk_tool_item_set_tooltip_markup(TOOL_ITEM(v), gstring(p_markup))
}

// Mark a tool item as important or non-important.
//
// When a gtk.Toolbar’s style is gtk.TOOLBAR_BOTH_HORIZ,
// labels are only displayed for tool item buttons considered important.
// This is an effect known as “priority text”.
func (v *ToolItem) SetIsImportant(b bool) {
	C.gtk_tool_item_set_is_important(TOOL_ITEM(v), gbool(b))
}

func (v *ToolItem) GetIsImportant() bool {
	return gobool(C.gtk_tool_item_get_is_important(TOOL_ITEM(v)))
}

func (v *ToolItem) GetToolbarStyle() ToolbarStyle {
	return ToolbarStyle(C.gtk_tool_item_get_toolbar_style(TOOL_ITEM(v)))
}

func (v *ToolItem) GetReliefStyle() ReliefStyle {
	return ReliefStyle(C.gtk_tool_item_get_relief_style(TOOL_ITEM(v)))
}

func (v *ToolItem) GetTextAlignment() float64 {
	return float64(C.gtk_tool_item_get_text_alignment(TOOL_ITEM(v)))
}

func (v *ToolItem) GetTextOrientation() Orientation {
	return Orientation(C.gtk_tool_item_get_text_orientation(TOOL_ITEM(v)))
}

// gtk_tool_item_retrieve_proxy_menu_item
// gtk_tool_item_get_proxy_menu_item
// gtk_tool_item_set_proxy_menu_item
func (v *ToolItem) RebuildMenu() {
	C.gtk_tool_item_rebuild_menu(TOOL_ITEM(v))
}

// gtk_tool_item_toolbar_reconfigured
// gtk_tool_item_get_text_size_group

//-----------------------------------------------------------------------
// GtkToolPalette
//-----------------------------------------------------------------------

type ToolPalette struct {
	Container
}

func NewToolPalette() *ToolPalette {
	return &ToolPalette{Container{Widget{C.gtk_tool_palette_new()}}}
}

func (v *ToolPalette) GetExclusive(group *ToolItemGroup) bool {
	return gobool(C.gtk_tool_palette_get_exclusive(TOOL_PALETTE(v), TOOL_ITEM_GROUP(group)))
}

func (v *ToolPalette) SetExclusive(group *ToolItemGroup, exclusive bool) {
	C.gtk_tool_palette_set_exclusive(TOOL_PALETTE(v), TOOL_ITEM_GROUP(group), gbool(exclusive))
}

func (v *ToolPalette) GetExpand(group *ToolItemGroup) bool {
	return gobool(C.gtk_tool_palette_get_expand(TOOL_PALETTE(v), TOOL_ITEM_GROUP(group)))
}

func (v *ToolPalette) SetExpand(group *ToolItemGroup, expand bool) {
	C.gtk_tool_palette_set_expand(TOOL_PALETTE(v), TOOL_ITEM_GROUP(group), gbool(expand))
}

func (v *ToolPalette) GetGroupPosition(group *ToolItemGroup) int {
	return int(C.gtk_tool_palette_get_group_position(TOOL_PALETTE(v), TOOL_ITEM_GROUP(group)))
}

func (v *ToolPalette) SetGroupPosition(group *ToolItemGroup, pos int) {
	C.gtk_tool_palette_set_group_position(TOOL_PALETTE(v), TOOL_ITEM_GROUP(group), gint(pos))
}

func (v *ToolPalette) GetIconSize() IconSize {
	return IconSize(C.gtk_tool_palette_get_icon_size(TOOL_PALETTE(v)))
}

func (v *ToolPalette) SetIconSize(size IconSize) {
	C.gtk_tool_palette_set_icon_size(TOOL_PALETTE(v), C.GtkIconSize(size))
}

func (v *ToolPalette) UnsetIconSize() {
	C.gtk_tool_palette_unset_icon_size(TOOL_PALETTE(v))
}

func (v *ToolPalette) GetStyle() ToolbarStyle {
	return ToolbarStyle(C.gtk_tool_palette_get_style(TOOL_PALETTE(v)))
}

func (v *ToolPalette) SetStyle(style ToolbarStyle) {
	C.gtk_tool_palette_set_style(TOOL_PALETTE(v), C.GtkToolbarStyle(style))
}

func (v *ToolPalette) UnsetStyle() {
	C.gtk_tool_palette_unset_style(TOOL_PALETTE(v))
}

// gtk_tool_palette_add_drag_dest
// gtk_tool_palette_get_drag_item
// gtk_tool_palette_get_drag_target_group
// gtk_tool_palette_get_drag_target_item
// gtk_tool_palette_get_drop_group
// gtk_tool_palette_get_drop_item
// gtk_tool_palette_set_drag_source
// gtk_tool_palette_get_hadjustment
// gtk_tool_palette_get_vadjustment

//-----------------------------------------------------------------------
// GtkToolItemGroup
//-----------------------------------------------------------------------

type ToolItemGroup struct {
	Container
	items map[*C.GtkToolItem]IWidget
}

func NewToolItemGroup(label string) *ToolItemGroup {
	var ptr *C.char
	if len(label) > 0 {
		ptr = C.CString(label)
		defer cfree(ptr)
	}
	return &ToolItemGroup{Container{Widget{C.gtk_tool_item_group_new(gstring(ptr))}},
		make(map[*C.GtkToolItem]IWidget)}
}

func (v *ToolItemGroup) Insert(item IWidget, pos int) {
	pitem := C.toGToolItem(ToNative(item))
	C.gtk_tool_item_group_insert(TOOL_ITEM_GROUP(v), pitem, gint(pos))
	v.items[pitem] = item
}

func (v *ToolItemGroup) GetCollapsed() bool {
	return gobool(C.gtk_tool_item_group_get_collapsed(TOOL_ITEM_GROUP(v)))
}

// gtk_tool_item_group_get_drop_item
// gtk_tool_item_group_get_ellipsize
// gtk_tool_item_group_get_item_position
// gtk_tool_item_group_get_n_items
func (v *ToolItemGroup) GetLabel() string {
	return gostring(C.gtk_tool_item_group_get_label(TOOL_ITEM_GROUP(v)))
}

// gtk_tool_item_group_get_label_widget
// gtk_tool_item_group_get_nth_item
func (v *ToolItemGroup) GetHeaderRelief() ReliefStyle {
	return ReliefStyle(C.gtk_tool_item_group_get_header_relief(TOOL_ITEM_GROUP(v)))
}

func (v *ToolItemGroup) SetCollapsed(collapsed bool) {
	C.gtk_tool_item_group_set_collapsed(TOOL_ITEM_GROUP(v), gbool(collapsed))
}

// gtk_tool_item_group_set_ellipsize
// gtk_tool_item_group_set_item_position
func (v *ToolItemGroup) SetLabel(label string) {
	var ptr *C.char
	if len(label) > 0 {
		ptr = C.CString(label)
		defer cfree(ptr)
	}
	C.gtk_tool_item_group_set_label(TOOL_ITEM_GROUP(v), gstring(ptr))
}

// gtk_tool_item_group_set_label_widget
func (v *ToolItemGroup) SetHeaderRelief(style ReliefStyle) {
	C.gtk_tool_item_group_set_header_relief(TOOL_ITEM_GROUP(v), C.GtkReliefStyle(style))
}

//-----------------------------------------------------------------------
// GtkSeparatorToolItem
//-----------------------------------------------------------------------
type SeparatorToolItem struct {
	ToolItem
}

func NewSeparatorToolItem() *SeparatorToolItem {
	return &SeparatorToolItem{*newToolItemInternal(
		AS_GWIDGET(unsafe.Pointer(C.gtk_separator_tool_item_new())))}
}

func (v *SeparatorToolItem) SetDraw(draw bool) {
	C.gtk_separator_tool_item_set_draw(SEPARATOR_TOOL_ITEM(v), gbool(draw))
}

func (v *SeparatorToolItem) GetDraw() bool {
	return gobool(C.gtk_separator_tool_item_get_draw(SEPARATOR_TOOL_ITEM(v)))
}

//-----------------------------------------------------------------------
// GtkToolButton
//-----------------------------------------------------------------------

type ToolButton struct {
	ToolItem
	iw, lw *Widget // Proxies for IconWidget and LabelWidget
}

func NewToolButton(icon IWidget, text string) *ToolButton {
	ctext := C.CString(text)
	defer cfree(ctext)
	tb := AS_GWIDGET(unsafe.Pointer(C.gtk_tool_button_new(
		ToNative(icon), gstring(ctext))))
	return &ToolButton{*newToolItemInternal(tb), nil, nil}
}

func NewToolButtonFromStock(stock_id string) *ToolButton {
	si := C.CString(stock_id)
	defer cfree(si)
	tb := AS_GWIDGET(unsafe.Pointer(
		C.gtk_tool_button_new_from_stock(gstring(si))))
	return &ToolButton{*newToolItemInternal(tb), nil, nil}
}

func (v *ToolButton) OnClicked(onclick interface{}, datas ...interface{}) int {
	return v.Connect("clicked", onclick, datas...)
}

func (v *ToolButton) SetLabel(label string) {
	var ptr *C.char
	if len(label) > 0 {
		ptr = C.CString(label)
		defer cfree(ptr)
	}
	C.gtk_tool_button_set_label(TOOL_BUTTON(v), gstring(ptr))
}

func (v *ToolButton) GetLabel() string {
	return gostring(C.gtk_tool_button_get_label(TOOL_BUTTON(v)))
}

func (v *ToolButton) SetUseUnderline(use_underline bool) {
	C.gtk_tool_button_set_use_underline(TOOL_BUTTON(v), gbool(use_underline))
}

func (v *ToolButton) GetUseUnderline() bool {
	return gobool(C.gtk_tool_button_get_use_underline(TOOL_BUTTON(v)))
}

func (v *ToolButton) SetStockId(stock_id string) {
	p_stock_id := C.CString(stock_id)
	defer cfree(p_stock_id)
	C.gtk_tool_button_set_stock_id(TOOL_BUTTON(v), gstring(p_stock_id))
}

func (v *ToolButton) GetStockId() string {
	return gostring(C.gtk_tool_button_get_stock_id(TOOL_BUTTON(v)))
}

func (v *ToolButton) SetIconName(icon_name string) {
	p_icon_name := C.CString(icon_name)
	defer cfree(p_icon_name)
	C.gtk_tool_button_set_icon_name(TOOL_BUTTON(v), gstring(p_icon_name))
}

func (v *ToolButton) GetIconName() string {
	return gostring(C.gtk_tool_button_get_icon_name(TOOL_BUTTON(v)))
}

func (v *ToolButton) SetIconWidget(icon_widget *Widget) {
	v.iw = icon_widget
	C.gtk_tool_button_set_icon_widget(TOOL_BUTTON(v), icon_widget.GWidget)
}

func (v *ToolButton) GetIconWidget() *Widget {
	if v.iw == nil {
		v.iw = &Widget{C.gtk_tool_button_get_icon_widget(TOOL_BUTTON(v))}
	}
	return v.iw
}

func (v *ToolButton) SetLabelWidget(label_widget *Widget) {
	v.lw = label_widget
	C.gtk_tool_button_set_label_widget(TOOL_BUTTON(v), label_widget.GWidget)
}

func (v *ToolButton) GetLabelWidget() *Widget {
	if v.lw == nil {
		v.lw = &Widget{C.gtk_tool_button_get_label_widget(TOOL_BUTTON(v))}
	}
	return v.lw
}

//-----------------------------------------------------------------------
// GtkMenuToolButton
//-----------------------------------------------------------------------

type MenuToolButton struct {
	ToolButton
	mw *Menu // Proxy for menu widget
}

func NewMenuToolButton(icon IWidget, text string) *MenuToolButton {
	ctext := C.CString(text)
	defer cfree(ctext)
	mtb := AS_GWIDGET(unsafe.Pointer(C.gtk_menu_tool_button_new(
		ToNative(icon), gstring(ctext))))
	return &MenuToolButton{ToolButton{*newToolItemInternal(mtb), nil, nil}, nil}
}

func NewMenuToolButtonFromStock(stock_id string) *MenuToolButton {
	si := C.CString(stock_id)
	defer cfree(si)
	mtb := AS_GWIDGET(unsafe.Pointer(
		C.gtk_menu_tool_button_new_from_stock(gstring(si))))
	return &MenuToolButton{ToolButton{*newToolItemInternal(mtb), nil, nil}, nil}
}

func (v *MenuToolButton) SetMenu(menu *Menu) {
	v.mw = menu
	C.gtk_menu_tool_button_set_menu(MENU_TOOL_BUTTON(v), menu.GWidget)
}

func (v *MenuToolButton) GetMenu() *Menu {
	if v.mw == nil {
		v.mw = &Menu{Container{Widget{C.gtk_menu_tool_button_get_menu(MENU_TOOL_BUTTON(v))}}}
	}
	return v.mw
}

// gtk_menu_tool_button_set_arrow_tooltip (deprecated since 2.12)
func (v *MenuToolButton) SetArrowTooltipText(text string) {
	pt := C.CString(text)
	defer cfree(pt)
	C.gtk_menu_tool_button_set_arrow_tooltip_text(MENU_TOOL_BUTTON(v), gstring(pt))
}

func (v *MenuToolButton) SetArrowTooltipMarkup(markup string) {
	pm := C.CString(markup)
	defer cfree(pm)
	C.gtk_menu_tool_button_set_arrow_tooltip_text(MENU_TOOL_BUTTON(v), gstring(pm))
}

//-----------------------------------------------------------------------
// GtkToggleToolButton
//-----------------------------------------------------------------------

type ToggleToolButton struct {
	ToolButton
}

func NewToggleToolButton() *ToggleToolButton {
	return &ToggleToolButton{ToolButton{*newToolItemInternal(
		AS_GWIDGET(unsafe.Pointer(C.gtk_toggle_tool_button_new()))), nil, nil}}
}

func NewToggleToolButtonFromStock(stock_id string) *ToggleToolButton {
	p_stock_id := C.CString(stock_id)
	defer cfree(p_stock_id)
	return &ToggleToolButton{ToolButton{*newToolItemInternal(
		AS_GWIDGET(unsafe.Pointer(C.gtk_toggle_tool_button_new_from_stock(
			gstring(p_stock_id))))), nil, nil}}
}

func (v *ToggleToolButton) OnToggled(onclick interface{}, datas ...interface{}) int {
	return v.Connect("toggled", onclick, datas...)
}

func (v *ToggleToolButton) SetActive(is_active bool) {
	C.gtk_toggle_tool_button_set_active(TOGGLE_TOOL_BUTTON(v), gbool(is_active))
}

func (v *ToggleToolButton) GetActive() bool {
	return gobool(C.gtk_toggle_tool_button_get_active(TOGGLE_TOOL_BUTTON(v)))
}

//-----------------------------------------------------------------------
// GtkRadioToolButton
//-----------------------------------------------------------------------

type RadioToolButton struct {
	ToggleToolButton
}

func NewRadioToolButton(group *glib.SList) *RadioToolButton {
	return &RadioToolButton{ToggleToolButton{ToolButton{*newToolItemInternal(
		AS_GWIDGET(unsafe.Pointer(C.gtk_radio_tool_button_new(
			gslist(group))))), nil, nil}}}
}

func NewRadioToolButtonFromStock(group *glib.SList, stock_id string) *RadioToolButton {
	p_stock_id := C.CString(stock_id)
	defer cfree(p_stock_id)
	return &RadioToolButton{ToggleToolButton{ToolButton{*newToolItemInternal(
		AS_GWIDGET(unsafe.Pointer(C.gtk_radio_tool_button_new_from_stock(
			gslist(group), gstring(p_stock_id))))), nil, nil}}}
}

// gtk_radio_tool_button_new_from_widget
// gtk_radio_tool_button_new_with_stock_from_widget
// gtk_radio_tool_button_get_group
// gtk_radio_tool_button_set_group

//-----------------------------------------------------------------------
// GtkUIManager
//-----------------------------------------------------------------------

type UIManager struct {
	glib.GObject
}

func NewUIManager() *UIManager {
	return &UIManager{glib.GObject{unsafe.Pointer(
		C.gtk_ui_manager_new())}}
}

// gtk_ui_manager_set_add_tearoffs
// gtk_ui_manager_get_add_tearoffs

func (v *UIManager) InsertActionGroup(action_group *ActionGroup, pos int) {
	C.gtk_ui_manager_insert_action_group(UI_MANAGER(v),
		ACTION_GROUP(action_group), gint(pos))
}

// gtk_ui_manager_remove_action_group
// gtk_ui_manager_get_action_groups

func (v *UIManager) GetAccelGroup() *AccelGroup {
	return &AccelGroup{C.gtk_ui_manager_get_accel_group(UI_MANAGER(v))}
}

func (v *UIManager) GetWidget(path string) *Widget {
	ptr := C.CString(path)
	defer cfree(ptr)
	return &Widget{C.gtk_ui_manager_get_widget(UI_MANAGER(v), gstring(ptr))}
}

// gtk_ui_manager_get_toplevels
// gtk_ui_manager_get_action

func (v *UIManager) AddUIFromString(buffer string) error {
	ptr := C.CString(buffer)
	defer cfree(ptr)
	var gerror *C.GError
	C.gtk_ui_manager_add_ui_from_string(UI_MANAGER(v), gstring(ptr), -1, &gerror)
	if gerror != nil {
		err := glib.ErrorFromNative(unsafe.Pointer(gerror))
		return err
	}
	return nil
}

// gtk_ui_manager_add_ui_from_file
// gtk_ui_manager_new_merge_id
// gtk_ui_manager_add_ui
// gtk_ui_manager_remove_ui
// gtk_ui_manager_get_ui
// gtk_ui_manager_ensure_update

//-----------------------------------------------------------------------
// GtkActionGroup
//-----------------------------------------------------------------------

type ActionGroup struct {
	glib.GObject
}

func NewActionGroup(name string) *ActionGroup {
	deprecated_since(3, 10, 0, "gtk_action_group_new()")
	ptr := C.CString(name)
	defer cfree(ptr)
	return &ActionGroup{glib.GObject{unsafe.Pointer(
		C.gtk_action_group_new(gstring(ptr)))}}
}

func (v *ActionGroup) GetName() string {
	deprecated_since(3, 10, 0, "gtk_action_group_get_name()")
	return gostring(C.gtk_action_group_get_name(ACTION_GROUP(v)))
}

func (v *ActionGroup) GetSensitive() bool {
	deprecated_since(3, 10, 0, "gtk_action_group_get_sensitive()")
	return gobool(C.gtk_action_group_get_sensitive(ACTION_GROUP(v)))
}

func (v *ActionGroup) SetSensitive(sensitive bool) {
	deprecated_since(3, 10, 0, "gtk_action_group_set_sensitive()")
	C.gtk_action_group_set_sensitive(ACTION_GROUP(v), gbool(sensitive))
}

func (v *ActionGroup) GetVisible() bool {
	deprecated_since(3, 10, 0, "gtk_action_group_get_visible()")
	return gobool(C.gtk_action_group_get_visible(ACTION_GROUP(v)))
}

func (v *ActionGroup) SetVisible(visible bool) {
	deprecated_since(3, 10, 0, "gtk_action_group_set_visible()")
	C.gtk_action_group_set_visible(ACTION_GROUP(v), gbool(visible))
}

func (v *ActionGroup) GetAction(action_name string) *Action {
	deprecated_since(3, 10, 0, "gtk_action_group_get_action()")
	ptr := C.CString(action_name)
	defer cfree(ptr)
	return &Action{glib.GObject{unsafe.Pointer(C.gtk_action_group_get_action(
		ACTION_GROUP(v), gstring(ptr)))}}
}

func (v *ActionGroup) ListActions() *glib.List {
	deprecated_since(3, 10, 0, "gtk_action_group_list_actions()")
	return glib.ListFromNative(unsafe.Pointer(
		C.gtk_action_group_list_actions(ACTION_GROUP(v))))
}

func (v *ActionGroup) AddAction(action interface{}) {
	deprecated_since(3, 10, 0, "gtk_action_group_add_action()")
	if a, ok := action.(IAction); ok {
		C.gtk_action_group_add_action(ACTION_GROUP(v), a.toGAction())
	} else {
		log.Panicf("Error calling gtk.AddAction: argument must be "+
			"an Action object, but %v type provided", reflect.TypeOf(action))
	}
}

func (v *ActionGroup) AddActionWithAccel(action *Action, accelerator string) {
	deprecated_since(3, 10, 0, "gtk_action_group_add_action_with_accel()")
	var ptr *C.char
	if len(accelerator) > 0 {
		ptr = C.CString(accelerator)
		defer cfree(ptr)
	}
	C.gtk_action_group_add_action_with_accel(ACTION_GROUP(v),
		ACTION(action), gstring(ptr))
}

func (v *ActionGroup) RemoveAction(action *Action) {
	deprecated_since(3, 10, 0, "gtk_action_group_remove_action()")
	C.gtk_action_group_remove_action(ACTION_GROUP(v), ACTION(action))
}

// gtk_action_group_add_actions
// gtk_action_group_add_actions_full
// gtk_action_group_add_toggle_actions
// gtk_action_group_add_toggle_actions_full
// gtk_action_group_add_radio_actions
// gtk_action_group_add_radio_actions_full
// gtk_action_group_set_translate_func
// gtk_action_group_set_translation_domain
// gtk_action_group_translate_string

//-----------------------------------------------------------------------
// GtkAction
//-----------------------------------------------------------------------

type IAction interface {
	toGAction() *C.GtkAction
}

type Action struct {
	glib.GObject
}

func NewAction(name string, label string, tooltip string, stock_id string) *Action {
	deprecated_since(3, 10, 0, "gtk_action_new()")
	name_ptr := C.CString(name)
	defer cfree(name_ptr)
	var label_ptr *C.char
	if len(label) > 0 {
		label_ptr = C.CString(label)
		defer cfree(label_ptr)
	}
	var tooltip_ptr *C.char
	if len(tooltip) > 0 {
		tooltip_ptr = C.CString(tooltip)
		defer cfree(tooltip_ptr)
	}
	var stockid_ptr *C.char
	if len(stock_id) > 0 {
		stockid_ptr = C.CString(stock_id)
		defer cfree(stockid_ptr)
	}
	return &Action{glib.GObject{unsafe.Pointer(C.gtk_action_new(gstring(name_ptr),
		gstring(label_ptr), gstring(tooltip_ptr), gstring(stockid_ptr)))}}
}

func (v *Action) toGAction() *C.GtkAction {
	return C.toGAction(v.Object)
}

func (v *Action) GetName() string {
	deprecated_since(3, 10, 0, "gtk_action_get_name()")
	return gostring(C.gtk_action_get_name(ACTION(v)))
}

func (v *Action) IsSensitive() bool {
	deprecated_since(3, 10, 0, "gtk_action_is_sensitive()")
	return gobool(C.gtk_action_is_sensitive(ACTION(v)))
}

func (v *Action) GetSensitive() bool {
	deprecated_since(3, 10, 0, "gtk_action_get_sensitive()")
	return gobool(C.gtk_action_get_sensitive(ACTION(v)))
}

func (v *Action) SetSensitive(sensitive bool) {
	deprecated_since(3, 10, 0, "gtk_action_set_sensitive()")
	C.gtk_action_set_sensitive(ACTION(v), gbool(sensitive))
}

func (v *Action) IsVisible() bool {
	deprecated_since(3, 10, 0, "gtk_action_is_visible()")
	return gobool(C.gtk_action_is_visible(ACTION(v)))
}

func (v *Action) GetVisible() bool {
	deprecated_since(3, 10, 0, "gtk_action_get_visible()")
	return gobool(C.gtk_action_get_visible(ACTION(v)))
}

func (v *Action) SetVisible(visible bool) {
	deprecated_since(3, 10, 0, "gtk_action_set_visible()")
	C.gtk_action_set_visible(ACTION(v), gbool(visible))
}

func (v *Action) Activate() {
	deprecated_since(3, 10, 0, "gtk_action_activate()")
	C.gtk_action_activate(ACTION(v))
}

func (v *Action) CreateIcon(icon_size IconSize) *Widget {
	deprecated_since(3, 10, 0, "gtk_action_create_icon()")
	return &Widget{C.gtk_action_create_icon(ACTION(v), C.GtkIconSize(icon_size))}
}

func (v *Action) CreateMenuItem() *Widget {
	deprecated_since(3, 10, 0, "gtk_action_create_menu_item()")
	return &Widget{C.gtk_action_create_menu_item(ACTION(v))}
}

func (v *Action) CreateToolItem() *Widget {
	deprecated_since(3, 10, 0, "gtk_action_create_tool_item()")
	return &Widget{C.gtk_action_create_tool_item(ACTION(v))}
}

func (v *Action) CreateMenu() *Widget {
	deprecated_since(3, 10, 0, "gtk_action_create_menu()")
	return &Widget{C.gtk_action_create_menu(ACTION(v))}
}

func (v *Action) ConnectProxy(proxy *Widget) {
	deprecated_since(2, 16, 0, "gtk_action_connect_proxy()")
	C.gtk_action_connect_proxy(ACTION(v), WIDGET(proxy))
}

func (v *Action) DisconnectProxy(proxy *Widget) {
	deprecated_since(2, 16, 0, "gtk_action_disconnect_proxy()")
	C.gtk_action_disconnect_proxy(ACTION(v), WIDGET(proxy))
}

func (v *Action) GetProxies() *glib.SList {
	deprecated_since(3, 10, 0, "gtk_action_get_proxies()")
	return glib.SListFromNative(unsafe.Pointer(
		C.gtk_action_get_proxies(ACTION(v))))
}

func (v *Action) ConnectAccelerator() {
	deprecated_since(3, 10, 0, "gtk_action_connect_accelerator()")
	C.gtk_action_connect_accelerator(ACTION(v))
}

func (v *Action) DisconnectAccelerator() {
	deprecated_since(3, 10, 0, "gtk_action_disconnect_accelerator()")
	C.gtk_action_disconnect_accelerator(ACTION(v))
}

func (v *Action) BlockActivate() {
	deprecated_since(3, 10, 0, "gtk_action_block_activate()")
	C.gtk_action_block_activate(ACTION(v))
}

func (v *Action) UnblockActivate() {
	deprecated_since(3, 10, 0, "gtk_action_unblock_activate()")
	C.gtk_action_unblock_activate(ACTION(v))
}

func (v *Action) BlockActivateFrom(proxy *Widget) {
	deprecated_since(2, 16, 0, "gtk_action_block_activate_from()")
	C.gtk_action_block_activate_from(ACTION(v), WIDGET(proxy))
}

func (v *Action) UnblockActivateFrom(proxy *Widget) {
	deprecated_since(2, 16, 0, "gtk_action_unblock_activate_from()")
	C.gtk_action_unblock_activate_from(ACTION(v), WIDGET(proxy))
}

func (v *Action) GetAlwaysShowImage() bool {
	deprecated_since(3, 10, 0, "gtk_action_get_always_show_image()")
	return gobool(C.gtk_action_get_always_show_image(ACTION(v)))
}

func (v *Action) SetAlwaysShowImage(always_show bool) {
	deprecated_since(3, 10, 0, "gtk_action_get_always_show_image()")
	C.gtk_action_set_always_show_image(ACTION(v), gbool(always_show))
}

func (v *Action) GetAccelPath() string {
	deprecated_since(3, 10, 0, "gtk_action_get_accel_path()")
	return gostring(C.gtk_action_get_accel_path(ACTION(v)))
}

func (v *Action) SetAccelPath(accel_path string) {
	deprecated_since(3, 10, 0, "gtk_action_set_accel_path()")
	ptr := C.CString(accel_path)
	C.gtk_action_set_accel_path(ACTION(v), gstring(ptr))
}

// gtk_action_get_accel_closure

func (v *Action) SetAccelGroup(accel_group *AccelGroup) {
	deprecated_since(3, 10, 0, "gtk_action_set_accel_group()")
	C.gtk_action_set_accel_group(ACTION(v), accel_group.GAccelGroup)
}

func (v *Action) SetLable(label string) {
	deprecated_since(3, 10, 0, "gtk_action_set_label()")
	ptr := C.CString(label)
	defer cfree(ptr)
	C.gtk_action_set_label(ACTION(v), gstring(ptr))
}

func (v *Action) GetLabel() string {
	deprecated_since(3, 10, 0, "gtk_action_get_label()")
	return gostring(C.gtk_action_get_label(ACTION(v)))
}

func (v *Action) SetShortLabel(short_label string) {
	deprecated_since(3, 10, 0, "gtk_action_set_short_label()")
	ptr := C.CString(short_label)
	defer cfree(ptr)
	C.gtk_action_set_short_label(ACTION(v), gstring(ptr))
}

func (v *Action) GetShortLabel() string {
	deprecated_since(3, 10, 0, "gtk_action_get_short_label()")
	return gostring(C.gtk_action_get_short_label(ACTION(v)))
}

func (v *Action) SetTooltip(tooltip string) {
	deprecated_since(3, 10, 0, "gtk_action_set_tooltip()")
	ptr := C.CString(tooltip)
	defer cfree(ptr)
	C.gtk_action_set_tooltip(ACTION(v), gstring(ptr))
}

func (v *Action) GetTooltip() string {
	deprecated_since(3, 10, 0, "gtk_action_get_tooltip()")
	return gostring(C.gtk_action_get_tooltip(ACTION(v)))
}

func (v *Action) SetStockId(stock_id string) {
	deprecated_since(3, 10, 0, "gtk_action_set_stock_id()")
	ptr := C.CString(stock_id)
	defer cfree(ptr)
	C.gtk_action_set_stock_id(ACTION(v), gstring(ptr))
}

func (v *Action) GetStockId() string {
	deprecated_since(3, 10, 0, "gtk_action_get_stock_id()")
	return gostring(C.gtk_action_get_stock_id(ACTION(v)))
}

// gtk_action_set_gicon
// gtk_action_get_gicon

func (v *Action) SetIconName(icon_name string) {
	deprecated_since(3, 10, 0, "gtk_action_set_icon_name()")
	ptr := C.CString(icon_name)
	defer cfree(ptr)
	C.gtk_action_set_icon_name(ACTION(v), gstring(ptr))
}

func (v *Action) GetIconName() string {
	deprecated_since(3, 10, 0, "gtk_action_get_icon_name()")
	return gostring(C.gtk_action_get_icon_name(ACTION(v)))
}

func (v *Action) SetVisibleHorizontal(visible_horizontal bool) {
	deprecated_since(3, 10, 0, "gtk_action_set_visible_horizontal()")
	C.gtk_action_set_visible_horizontal(ACTION(v), gbool(visible_horizontal))
}

func (v *Action) GetVisibleHorizontal() bool {
	deprecated_since(3, 10, 0, "gtk_action_get_visible_horizontal()")
	return gobool(C.gtk_action_get_visible_horizontal(ACTION(v)))
}

func (v *Action) SetVisibleVertical(visible_vertical bool) {
	deprecated_since(3, 10, 0, "gtk_action_set_visible_vertical()")
	C.gtk_action_set_visible_vertical(ACTION(v), gbool(visible_vertical))
}

func (v *Action) GetVisibleVertical() bool {
	deprecated_since(3, 10, 0, "gtk_action_get_visible_vertical()")
	return gobool(C.gtk_action_get_visible_vertical(ACTION(v)))
}

func (v *Action) SetIsImportant(is_important bool) {
	deprecated_since(3, 10, 0, "gtk_action_set_is_important()")
	C.gtk_action_set_is_important(ACTION(v), gbool(is_important))
}

func (v *Action) GetIsImportant() bool {
	deprecated_since(3, 10, 0, "gtk_action_get_is_important()")
	return gobool(C.gtk_action_get_is_important(ACTION(v)))
}

//-----------------------------------------------------------------------
// GtkToggleAction
//-----------------------------------------------------------------------

type ToggleAction struct {
	Action
}

func NewToggleAction(name string, label string, tooltip string,
	stock_id string) *ToggleAction {
	deprecated_since(3, 10, 0, "gtk_toggle_action_new()")
	name_ptr := C.CString(name)
	defer cfree(name_ptr)
	var label_ptr *C.char
	if len(label) > 0 {
		label_ptr = C.CString(label)
		defer cfree(label_ptr)
	}
	var tooltip_ptr *C.char
	if len(tooltip) > 0 {
		tooltip_ptr = C.CString(tooltip)
		defer cfree(tooltip_ptr)
	}
	var stockid_ptr *C.char
	if len(stock_id) > 0 {
		stockid_ptr = C.CString(stock_id)
		defer cfree(stockid_ptr)
	}
	return &ToggleAction{Action{glib.GObject{unsafe.Pointer(C.gtk_toggle_action_new(gstring(name_ptr),
		gstring(label_ptr), gstring(tooltip_ptr), gstring(stockid_ptr)))}}}
}

func (v *ToggleAction) Toggled() {
	deprecated_since(3, 10, 0, "gtk_toggle_action_toggled()")
	C.gtk_toggle_action_toggled(TOGGLE_ACTION(v))
}

func (v *ToggleAction) SetActive(is_active bool) {
	deprecated_since(3, 10, 0, "gtk_toggle_action_set_active()")
	C.gtk_toggle_action_set_active(TOGGLE_ACTION(v), gbool(is_active))
}

func (v *ToggleAction) GetActive() bool {
	deprecated_since(3, 10, 0, "gtk_toggle_action_get_active()")
	return gobool(C.gtk_toggle_action_get_active(TOGGLE_ACTION(v)))
}

func (v *ToggleAction) SetDrawAsRadio(draw_as_radio bool) {
	deprecated_since(3, 10, 0, "gtk_toggle_action_set_draw_as_radio()")
	C.gtk_toggle_action_set_draw_as_radio(TOGGLE_ACTION(v),
		gbool(draw_as_radio))
}

func (v *ToggleAction) GetDrawAsRadio() bool {
	deprecated_since(3, 10, 0, "gtk_toggle_action_get_draw_as_radio()")
	return gobool(C.gtk_toggle_action_get_draw_as_radio(TOGGLE_ACTION(v)))
}

//-----------------------------------------------------------------------
// GtkRadioAction
//-----------------------------------------------------------------------

type RadioAction struct {
	ToggleAction
}

func NewRadioAction(name string, label string, tooltip string,
	stock_id string, value int) *RadioAction {
	deprecated_since(3, 10, 0, "gtk_radio_action_new()")
	name_ptr := C.CString(name)
	defer cfree(name_ptr)
	var label_ptr *C.char
	if len(label) > 0 {
		label_ptr = C.CString(label)
		defer cfree(label_ptr)
	}
	var tooltip_ptr *C.char
	if len(tooltip) > 0 {
		tooltip_ptr = C.CString(tooltip)
		defer cfree(tooltip_ptr)
	}
	var stockid_ptr *C.char
	if len(stock_id) > 0 {
		stockid_ptr = C.CString(stock_id)
		defer cfree(stockid_ptr)
	}
	return &RadioAction{ToggleAction{Action{glib.GObject{unsafe.Pointer(
		C.gtk_radio_action_new(gstring(name_ptr), gstring(label_ptr),
			gstring(tooltip_ptr), gstring(stockid_ptr), gint(value)))}}}}
}

func (v *RadioAction) GetGroup() *glib.SList {
	deprecated_since(3, 10, 0, "gtk_radio_action_get_group()")
	return glib.SListFromNative(unsafe.Pointer(
		C.gtk_radio_action_get_group(RADIO_ACTION(v))))
}

func (v *RadioAction) SetGroup(group *glib.SList) {
	deprecated_since(3, 10, 0, "gtk_radio_action_set_group()")
	C.gtk_radio_action_set_group(RADIO_ACTION(v), gslist(group))
}

func (v *RadioAction) GetCurrentValue() int {
	deprecated_since(3, 10, 0, "gtk_radio_action_get_current_value()")
	return int(C.gtk_radio_action_get_current_value(RADIO_ACTION(v)))
}

func (v *RadioAction) SetCurrentValue(current_value int) {
	deprecated_since(3, 10, 0, "gtk_radio_action_set_current_value()")
	C.gtk_radio_action_set_current_value(RADIO_ACTION(v), gint(current_value))
}

//-----------------------------------------------------------------------
// GtkRecentAction
//-----------------------------------------------------------------------

type RecentAction struct {
	Action
}

func NewRecentAction(name string, label string, tooltip string,
	stock_id string) *RecentAction {
	deprecated_since(3, 10, 0, "gtk_recent_action_new()")
	name_ptr := C.CString(name)
	defer cfree(name_ptr)
	var label_ptr *C.char
	if len(label) > 0 {
		label_ptr = C.CString(label)
		defer cfree(label_ptr)
	}
	var tooltip_ptr *C.char
	if len(tooltip) > 0 {
		tooltip_ptr = C.CString(tooltip)
		defer cfree(tooltip_ptr)
	}
	var stockid_ptr *C.char
	if len(stock_id) > 0 {
		stockid_ptr = C.CString(stock_id)
		defer cfree(stockid_ptr)
	}
	return &RecentAction{Action{glib.GObject{unsafe.Pointer(
		C.gtk_recent_action_new(gstring(name_ptr), gstring(label_ptr),
			gstring(tooltip_ptr), gstring(stockid_ptr)))}}}
}

// gtk_recent_action_new_for_manager

func (v *RecentAction) GetShowNumbers() bool {
	deprecated_since(3, 10, 0, "gtk_recent_action_get_show_numbers()")
	return gobool(C.gtk_recent_action_get_show_numbers(RECENT_ACTION(v)))
}

func (v *RecentAction) SetShowNumbers(show_numbers bool) {
	deprecated_since(3, 10, 0, "gtk_recent_action_set_show_numbers()")
	C.gtk_recent_action_set_show_numbers(RECENT_ACTION(v), gbool(show_numbers))
}

//-----------------------------------------------------------------------
// GtkActivatable
//-----------------------------------------------------------------------

// Known Implementations for GtkActivatable interface:
// GtkActivatable is implemented by GtkButton, GtkCheckButton,
// GtkCheckMenuItem, GtkColorButton, GtkFontButton, GtkImageMenuItem,
// GtkLinkButton, GtkLockButton, GtkMenuItem, GtkMenuToolButton,
// GtkRadioButton, GtkRadioMenuItem, GtkRadioToolButton,
// GtkRecentChooserMenu, GtkScaleButton, GtkSeparatorMenuItem,
// GtkSeparatorToolItem, GtkSwitch, GtkTearoffMenuItem, GtkToggleButton,
// GtkToggleToolButton, GtkToolButton, GtkToolItem and GtkVolumeButton.
type Activatable struct {
	GWidget *C.GtkWidget
}

// gtk_activatable_do_set_related_action

func (v *Activatable) GetRelatedAction() *Action {
	deprecated_since(3, 10, 0, "gtk_activatable_get_related_action()")
	return &Action{glib.GObject{unsafe.Pointer(C.gtk_activatable_get_related_action(ACTIVATABLE(v)))}}
}

func (v *Activatable) GetUseActionAppearance() bool {
	deprecated_since(3, 10, 0, "gtk_activatable_get_use_action_appearance()")
	return gobool(C.gtk_activatable_get_use_action_appearance(ACTIVATABLE(v)))
}

// gtk_activatable_sync_action_properties

func (v *Activatable) SetRelatedAction(action *Action) {
	deprecated_since(3, 10, 0, "gtk_activatable_set_related_action()")
	C.gtk_activatable_set_related_action(ACTIVATABLE(v), ACTION(action))
}

func (v *Activatable) SetUseActionAppearance(use_appearance bool) {
	deprecated_since(3, 10, 0, "gtk_activatable_set_use_action_appearance()")
	C.gtk_activatable_set_use_action_appearance(ACTIVATABLE(v),
		gbool(use_appearance))
}

//-----------------------------------------------------------------------
// GtkColorButton
//-----------------------------------------------------------------------

type ColorButton struct {
	Button
}

func NewColorButton() *ColorButton {
	return &ColorButton{*newButtonInternal(
		AS_GWIDGET(unsafe.Pointer(C.gtk_color_button_new())))}
}

func NewColorButtonWithColor(color *gdk.Color) *ColorButton {
	v := NewColorButton()
	v.SetColor(color)
	return v
}

func (v *ColorButton) SetColor(color *gdk.Color) {
	C.gtk_color_button_set_color(COLOR_BUTTON(v), (*C.GdkColor)(unsafe.Pointer(&color.GColor)))
}

func (v *ColorButton) GetColor() *gdk.Color {
	c := new(gdk.Color)
	C.gtk_color_button_get_color(COLOR_BUTTON(v), (*C.GdkColor)(unsafe.Pointer(&c.GColor)))
	return c
}

func (v *ColorButton) SetAlpha(alpha uint16) {
	C.gtk_color_button_set_alpha(COLOR_BUTTON(v), guint16(alpha))
}

func (v *ColorButton) GetAlpha() uint16 {
	return uint16(C.gtk_color_button_get_alpha(COLOR_BUTTON(v)))
}

func (v *ColorButton) SetUseAlpha(use_alpha bool) {
	C.gtk_color_button_set_use_alpha(COLOR_BUTTON(v), gbool(use_alpha))
}

func (v *ColorButton) GetUseAlpha() bool {
	return gobool(C.gtk_color_button_get_use_alpha(COLOR_BUTTON(v)))
}

func (v *ColorButton) SetTitle(title string) {
	ptr := C.CString(title)
	defer cfree(ptr)
	C.gtk_color_button_set_title(COLOR_BUTTON(v), gstring(ptr))
}

func (v *ColorButton) GetTitle() string {
	return gostring(C.gtk_color_button_get_title(COLOR_BUTTON(v)))
}

//-----------------------------------------------------------------------
// GtkColorSelectionDialog
//-----------------------------------------------------------------------

// gtk_color_selection_dialog_new
// gtk_color_selection_dialog_get_color_selection

//-----------------------------------------------------------------------
// GtkColorSelection
//-----------------------------------------------------------------------

// gtk_color_selection_new
// gtk_color_selection_set_update_policy
// gtk_color_selection_set_has_opacity_control
// gtk_color_selection_get_has_opacity_control
// gtk_color_selection_set_has_palette
// gtk_color_selection_get_has_palette
// gtk_color_selection_get_current_alpha
// gtk_color_selection_set_current_alpha
// gtk_color_selection_get_current_color
// gtk_color_selection_set_current_color
// gtk_color_selection_get_previous_alpha
// gtk_color_selection_set_previous_alpha
// gtk_color_selection_get_previous_color
// gtk_color_selection_set_previous_color
// gtk_color_selection_is_adjusting
// gtk_color_selection_palette_from_string
// gtk_color_selection_palette_to_string
// gtk_color_selection_set_change_palette_hook
// gtk_color_selection_set_change_palette_with_screen_hook
// gtk_color_selection_set_color
// gtk_color_selection_get_color

//-----------------------------------------------------------------------
// GtkHSV
//-----------------------------------------------------------------------

// gtk_hsv_new
// gtk_hsv_set_color
// gtk_hsv_get_color
// gtk_hsv_set_metrics
// gtk_hsv_get_metrics
// gtk_hsv_is_adjusting
// gtk_hsv_to_rgb
// gtk_rgb_to_hsv

//-----------------------------------------------------------------------
// GtkFileChooser
//-----------------------------------------------------------------------
type FileChooserAction int

const (
	FILE_CHOOSER_ACTION_OPEN FileChooserAction = iota
	FILE_CHOOSER_ACTION_SAVE
	FILE_CHOOSER_ACTION_SELECT_FOLDER
	FILE_CHOOSER_ACTION_CREATE_FOLDER
)

type FileChooser struct {
	GFileChooser *C.GtkFileChooser
}

func (v *FileChooser) SetAction(action FileChooserAction) {
	C.gtk_file_chooser_set_action(v.GFileChooser, C.GtkFileChooserAction(action))
}

func (v *FileChooser) GetAction() FileChooserAction {
	return FileChooserAction(C.gtk_file_chooser_get_action(v.GFileChooser))
}

func (v *FileChooser) SetLocalOnly(b bool) {
	C.gtk_file_chooser_set_local_only(v.GFileChooser, gbool(b))
}

func (v *FileChooser) GetLocalOnly() bool {
	return gobool(C.gtk_file_chooser_get_local_only(v.GFileChooser))
}

// void gtk_file_chooser_set_select_multiple(GtkFileChooser* chooser, gboolean select_multiple);
// gboolean gtk_file_chooser_get_select_multiple(GtkFileChooser* chooser);
// void gtk_file_chooser_set_show_hidden(GtkFileChooser* chooser, gboolean show_hidden);
// gboolean gtk_file_chooser_get_show_hidden(GtkFileChooser* chooser);
// void gtk_file_chooser_set_do_overwrite_confirmation(GtkFileChooser* chooser, gboolean do_overwrite_confirmation);
// gboolean gtk_file_chooser_get_do_overwrite_confirmation(GtkFileChooser* chooser);
// void gtk_file_chooser_set_create_folders(GtkFileChooser* chooser, gboolean create_folders); //since 2.18
// gboolean gtk_file_chooser_get_create_folders(GtkFileChooser* chooser); //since 2.18
// void gtk_file_chooser_set_current_name(GtkFileChooser* chooser, const gchar* name);

func (v *FileChooser) GetFilename() string {
	return gostring(C.gtk_file_chooser_get_filename(v.GFileChooser))
}

func (v *FileChooser) SetFilename(filename string) {
	ptr := C.CString(filename)
	defer cfree(ptr)
	C.gtk_file_chooser_set_filename(v.GFileChooser, ptr)
}

// gboolean gtk_file_chooser_select_filename(GtkFileChooser* chooser, const char* filename);
// void gtk_file_chooser_unselect_filename(GtkFileChooser* chooser, const char* filename);
// void gtk_file_chooser_select_all(GtkFileChooser* chooser);
// void gtk_file_chooser_unselect_all(GtkFileChooser* chooser);
// GSList*  gtk_file_chooser_get_filenames(GtkFileChooser* chooser);

func (v *FileChooser) SetCurrentFolder(f string) bool {
	cf := C.CString(f)
	defer cfree(cf)
	return gobool(C.gtk_file_chooser_set_current_folder(
		v.GFileChooser, gstring(cf)))
}

func (v *FileChooser) GetCurrentFolder() string {
	return gostring(C.gtk_file_chooser_get_current_folder(v.GFileChooser))
}

// gchar*  gtk_file_chooser_get_uri(GtkFileChooser* chooser);
// gboolean gtk_file_chooser_set_uri(GtkFileChooser* chooser, const char* uri);
// gboolean gtk_file_chooser_select_uri(GtkFileChooser* chooser, const char* uri);
// void gtk_file_chooser_unselect_uri(GtkFileChooser* chooser, const char* uri);
// GSList*  gtk_file_chooser_get_uris(GtkFileChooser* chooser);
// gboolean gtk_file_chooser_set_current_folder_uri(GtkFileChooser* chooser, const gchar* uri);
// gchar*  gtk_file_chooser_get_current_folder_uri(GtkFileChooser* chooser);
// void gtk_file_chooser_set_preview_widget(GtkFileChooser* chooser, GtkWidget* preview_widget);
// GtkWidget* gtk_file_chooser_get_preview_widget(GtkFileChooser* chooser);
// void gtk_file_chooser_set_preview_widget_active(GtkFileChooser* chooser, gboolean active);
// gboolean gtk_file_chooser_get_preview_widget_active(GtkFileChooser* chooser);
// void gtk_file_chooser_set_use_preview_label(GtkFileChooser* chooser, gboolean use_label);
// gboolean gtk_file_chooser_get_use_preview_label(GtkFileChooser* chooser);
// char* gtk_file_chooser_get_preview_filename(GtkFileChooser* chooser);
// char* gtk_file_chooser_get_preview_uri(GtkFileChooser* chooser);
// void gtk_file_chooser_set_extra_widget(GtkFileChooser* chooser, GtkWidget* extra_widget);
// GtkWidget* gtk_file_chooser_get_extra_widget(GtkFileChooser* chooser);

func (v *FileChooser) AddFilter(f *FileFilter) {
	C.gtk_file_chooser_add_filter(v.GFileChooser, f.GFileFilter)
}

func (v *FileChooser) RemoveFilter(f *FileFilter) {
	C.gtk_file_chooser_remove_filter(v.GFileChooser, f.GFileFilter)
}

func (v *FileChooser) ListFilters() []*FileFilter {
	c_list := C.gtk_file_chooser_list_filters(v.GFileChooser)
	defer C.g_slist_free(c_list)
	n := int(C.g_slist_length(c_list))
	ret := make([]*FileFilter, n)
	for i := 0; i < n; i++ {
		ret[i] = &FileFilter{C.toGFileFilter(C.g_slist_nth_data(c_list, C.guint(i)))}
	}
	return ret
}

func (v *FileChooser) SetFilter(f *FileFilter) {
	C.gtk_file_chooser_set_filter(v.GFileChooser, f.GFileFilter)
}

func (v *FileChooser) GetFilter() *FileFilter {
	return &FileFilter{C.gtk_file_chooser_get_filter(v.GFileChooser)}
}

// gboolean gtk_file_chooser_add_shortcut_folder(GtkFileChooser* chooser, const char* folder, GError* *error);
// gboolean gtk_file_chooser_remove_shortcut_folder(GtkFileChooser* chooser, const char* folder, GError* *error);
// GSList* gtk_file_chooser_list_shortcut_folders(GtkFileChooser* chooser);
// gboolean gtk_file_chooser_add_shortcut_folder_uri(GtkFileChooser* chooser, const char* uri, GError* *error);
// gboolean gtk_file_chooser_remove_shortcut_folder_uri(GtkFileChooser* chooser, const char* uri, GError* *error);
// GSList* gtk_file_chooser_list_shortcut_folder_uris(GtkFileChooser* chooser);
// GFile*  gtk_file_chooser_get_current_folder_file(GtkFileChooser* chooser);
// GFile*  gtk_file_chooser_get_file(GtkFileChooser* chooser);
// GSList*  gtk_file_chooser_get_files(GtkFileChooser* chooser);
// GFile* gtk_file_chooser_get_preview_file(GtkFileChooser* chooser);
// gboolean gtk_file_chooser_select_file(GtkFileChooser* chooser, GFile* file, GError* *error);
// gboolean gtk_file_chooser_set_current_folder_file(GtkFileChooser* chooser, GFile* file, GError* *error);
// gboolean gtk_file_chooser_set_file(GtkFileChooser* chooser, GFile* file, GError* *error);
// void gtk_file_chooser_unselect_file(GtkFileChooser* chooser, GFile* file);

//-----------------------------------------------------------------------
// GtkFileChooserButton
//-----------------------------------------------------------------------
type FileChooserButton struct {
	HBox
	FileChooser
}

// gtk_file_chooser_button_new
func NewFileChooserButton(title string, action FileChooserAction) *FileChooserButton {
	ptitle := C.CString(title)
	defer cfree(ptitle)
	w := Widget{
		C.gtk_file_chooser_button_new(gstring(ptitle), C.GtkFileChooserAction(action)),
	}
	return &FileChooserButton{HBox{Box{Container{w}}},
		FileChooser{FILE_CHOOSER(&w)},
	}
}

// gtk_file_chooser_button_new_with_backend

// gtk_file_chooser_button_new_with_dialog
func NewFileChooserButtonWithDialog(dialog *FileChooserDialog) *FileChooserButton {
	w := Widget{
		C.gtk_file_chooser_button_new_with_dialog(dialog.GWidget),
	}
	return &FileChooserButton{
		HBox{Box{Container{w}}},
		FileChooser{FILE_CHOOSER(&w)},
	}
}

// gtk_file_chooser_button_get_title
// gtk_file_chooser_button_set_title
// gtk_file_chooser_button_get_width_chars
// gtk_file_chooser_button_set_width_chars
// gtk_file_chooser_button_get_focus_on_click
// gtk_file_chooser_button_set_focus_on_click

//-----------------------------------------------------------------------
// GtkFileChooserDialog
//-----------------------------------------------------------------------
type FileChooserDialog struct {
	Dialog
	FileChooser
}

func NewFileChooserDialog(title string, parent *Window, file_chooser_action FileChooserAction,
	button_text string, button_action ResponseType, buttons ...interface{}) *FileChooserDialog {
	ptitle := C.CString(title)
	defer cfree(ptitle)
	pbutton := C.CString(button_text)
	defer cfree(pbutton)
	widget := Widget{
		C._gtk_file_chooser_dialog_new(
			gstring(ptitle),
			ToNative(parent),
			C.int(file_chooser_action),
			C.int(button_action),
			gstring(pbutton))}
	ret := &FileChooserDialog{Dialog{Window{Bin{Container{widget}}}, nil}, FileChooser{FILE_CHOOSER(&widget)}}
	text, res := variadicButtonsToArrays(buttons)
	for i := range text {
		ret.AddButton(text[i], res[i])
	}
	return ret
}

//-----------------------------------------------------------------------
// GtkFileChooserWidget
//-----------------------------------------------------------------------
type FileChooserWidget struct {
	VBox
	FileChooser
}

func NewFileChooserWidget(file_chooser_action FileChooserAction) *FileChooserWidget {
	widget := Widget{C._gtk_file_chooser_widget_new(C.int(file_chooser_action))}
	return &FileChooserWidget{VBox{Box{Container{widget}}},
		FileChooser{FILE_CHOOSER(&widget)},
	}
}

// gtk_file_chooser_widget_new_with_backend

//-----------------------------------------------------------------------
// GtkFileFilter
//-----------------------------------------------------------------------
type FileFilter struct {
	GFileFilter *C.GtkFileFilter
}

func NewFileFilter() *FileFilter {
	return &FileFilter{C.gtk_file_filter_new()}
}

func (v *FileFilter) SetName(name string) {
	ptr := C.CString(name)
	defer cfree(ptr)
	C.gtk_file_filter_set_name(v.GFileFilter, gstring(ptr))
}

func (v *FileFilter) GetName() string {
	return gostring(C.gtk_file_filter_get_name(v.GFileFilter))
}

func (v *FileFilter) AddMimeType(mimetype string) {
	ptr := C.CString(mimetype)
	defer cfree(ptr)
	C.gtk_file_filter_add_mime_type(v.GFileFilter, gstring(ptr))
}

func (v *FileFilter) AddPattern(pattern string) {
	ptr := C.CString(pattern)
	defer cfree(ptr)
	C.gtk_file_filter_add_pattern(v.GFileFilter, gstring(ptr))
}

//void gtk_file_filter_add_pixbuf_formats (GtkFileFilter *filter);
//void gtk_file_filter_add_custom (GtkFileFilter *filter, GtkFileFilterFlags needed, GtkFileFilterFunc func, gpointer data, GDestroyNotify notify);
// gtk_file_filter_get_needed  //for use in the implementation of GtkFileChooser
// gtk_file_filter_filter  //for use in the implementation of GtkFileChooser

//-----------------------------------------------------------------------
// GtkFontButton
//-----------------------------------------------------------------------
type FontButton struct {
	Button
}

func NewFontButton() *FontButton {
	return &FontButton{*newButtonInternal(
		C.gtk_font_button_new())}
}

func NewFontButtonWithFont(fontname string) *FontButton {
	ptr := C.CString(fontname)
	defer cfree(ptr)
	return &FontButton{*newButtonInternal(
		C.gtk_font_button_new_with_font(gstring(ptr)))}
}

func (v *FontButton) SetFontName(fontname string) {
	ptr := C.CString(fontname)
	defer cfree(ptr)
	C.gtk_font_button_set_font_name(FONT_BUTTON(v), gstring(ptr))
}

func (v *FontButton) GetFontName() string {
	return gostring(C.gtk_font_button_get_font_name(FONT_BUTTON(v)))
}

func (v *FontButton) SetShowStyle(show bool) {
	C.gtk_font_button_set_show_style(FONT_BUTTON(v), gbool(show))
}

func (v *FontButton) GetShowStyle() bool {
	return gobool(C.gtk_font_button_get_show_style(FONT_BUTTON(v)))
}

func (v *FontButton) SetShowSize(show_size bool) {
	C.gtk_font_button_set_show_size(FONT_BUTTON(v), gbool(show_size))
}

func (v *FontButton) GetShowSize() bool {
	return gobool(C.gtk_font_button_get_show_size(FONT_BUTTON(v)))
}

func (v *FontButton) SetUseFont(use bool) {
	C.gtk_font_button_set_use_font(FONT_BUTTON(v), gbool(use))
}

func (v *FontButton) GetUseFont() bool {
	return gobool(C.gtk_font_button_get_use_font(FONT_BUTTON(v)))
}

func (v *FontButton) SetUseSize(use_size bool) {
	C.gtk_font_button_set_use_size(FONT_BUTTON(v), gbool(use_size))
}

func (v *FontButton) GetUseSize() bool {
	return gobool(C.gtk_font_button_get_use_size(FONT_BUTTON(v)))
}

func (v *FontButton) SetTitle(title string) {
	ptr := C.CString(title)
	defer cfree(ptr)
	C.gtk_font_button_set_title(FONT_BUTTON(v), gstring(ptr))
}

func (v *FontButton) GetTitle() string {
	return gostring(C.gtk_font_button_get_title(FONT_BUTTON(v)))
}

//-----------------------------------------------------------------------
// GtkFontSelection
//-----------------------------------------------------------------------
type FontSelection struct {
	VBox
}

func NewFontSelection() *FontSelection {
	return &FontSelection{VBox{Box{Container{Widget{C.gtk_font_selection_new()}}}}}
}

func (v *FontSelection) GetFont() *gdk.Font {
	return gdk.FontFromUnsafe(unsafe.Pointer(C.gtk_font_selection_get_font(FONT_SELECTION(v))))
}

func (v *FontSelection) GetFontName() string {
	return gostring(C.gtk_font_selection_get_font_name(FONT_SELECTION(v)))
}

func (v *FontSelection) SetFontName(name string) {
	ptr := C.CString(name)
	defer cfree(ptr)
	C.gtk_font_selection_set_font_name(FONT_SELECTION(v), gstring(ptr))
}

// gtk_font_selection_get_preview_text
// gtk_font_selection_set_preview_text
// gtk_font_selection_get_face
// gtk_font_selection_get_face_list
// gtk_font_selection_get_family
// gtk_font_selection_get_size
// gtk_font_selection_get_family_list
// gtk_font_selection_get_preview_entry
// gtk_font_selection_get_size_entry
// gtk_font_selection_get_size_list

//-----------------------------------------------------------------------
// GtkFontSelectionDialog
//-----------------------------------------------------------------------
type FontSelectionDialog struct {
	Dialog
}

func NewFontSelectionDialog(title string) *FontSelectionDialog {
	ptitle := C.CString(title)
	defer cfree(ptitle)
	return &FontSelectionDialog{Dialog{Window{Bin{Container{Widget{
		C.gtk_font_selection_dialog_new(gstring(ptitle))}}}}, nil}}
}

func (v *FontSelectionDialog) GetFontName() string {
	return gostring(C.gtk_font_selection_dialog_get_font_name(FONT_SELECTION_DIALOG(v)))
}

func (v *FontSelectionDialog) SetFontName(font string) {
	pfont := C.CString(font)
	defer cfree(pfont)
	C.gtk_font_selection_dialog_set_font_name(FONT_SELECTION_DIALOG(v), gstring(pfont))
}

func (v *FontSelectionDialog) GetPreviewText() string {
	return gostring(C.gtk_font_selection_dialog_get_preview_text(FONT_SELECTION_DIALOG(v)))
}

func (v *FontSelectionDialog) SetPreviewText(text string) {
	ptr := C.CString(text)
	defer cfree(ptr)
	C.gtk_font_selection_dialog_set_preview_text(FONT_SELECTION_DIALOG(v), gstring(ptr))
}

func (v *FontSelectionDialog) GetCancelButton() *Widget {
	return &Widget{C.gtk_font_selection_dialog_get_cancel_button(FONT_SELECTION_DIALOG(v))}
}

func (v *FontSelectionDialog) GetOkButton() *Widget {
	return &Widget{C.gtk_font_selection_dialog_get_ok_button(FONT_SELECTION_DIALOG(v))}
}

func (v *FontSelectionDialog) GetFontSelection() *FontSelection {
	return &FontSelection{VBox{Box{Container{Widget{
		C.gtk_font_selection_dialog_get_font_selection(FONT_SELECTION_DIALOG(v))}}}}}
}

//-----------------------------------------------------------------------
// GtkInputDialog
//-----------------------------------------------------------------------

// gtk_input_dialog_new //deprecated in 2.20

//-----------------------------------------------------------------------
// GtkAlignment
//-----------------------------------------------------------------------
type Alignment struct {
	Bin
}

func NewAlignment(xalign float64, yalign float64, xscale float64, yscale float64) *Alignment {
	return &Alignment{Bin{Container{Widget{
		C.gtk_alignment_new(C.gfloat(xalign), C.gfloat(yalign), C.gfloat(xscale), C.gfloat(yscale))}}}}
}

func (v *Alignment) Set(xalign float64, yalign float64, xscale float64, yscale float64) {
	C.gtk_alignment_set(ALIGNMENT(v), C.gfloat(xalign), C.gfloat(yalign), C.gfloat(xscale), C.gfloat(yscale))
}

func (v *Alignment) SetPadding(padding_top uint, padding_bottom uint, padding_left uint, padding_right uint) {
	C.gtk_alignment_set_padding(ALIGNMENT(v), guint(padding_top), guint(padding_bottom), guint(padding_left), guint(padding_right))
}

func (v *Alignment) GetPadding() (padding_top uint, padding_bottom uint, padding_left uint, padding_right uint) {
	var cpadding_top, cpadding_bottom, cpadding_left, cpadding_right C.guint
	C.gtk_alignment_get_padding(ALIGNMENT(v), &cpadding_top, &cpadding_bottom, &cpadding_left, &cpadding_right)
	padding_top = uint(cpadding_top)
	padding_bottom = uint(cpadding_bottom)
	padding_left = uint(cpadding_left)
	padding_right = uint(cpadding_right)
	return
}

//-----------------------------------------------------------------------
// GtkAspectFrame
//-----------------------------------------------------------------------

// gtk_aspect_frame_new
// gtk_aspect_frame_set

//-----------------------------------------------------------------------
// GtkHBox
//-----------------------------------------------------------------------
type HBox struct {
	Box
}

func NewHBox(homogeneous bool, spacing int) *HBox {
	return &HBox{Box{Container{Widget{C.gtk_hbox_new(gbool(homogeneous), gint(spacing))}}}}
}

//-----------------------------------------------------------------------
// GtkVBox
//-----------------------------------------------------------------------
type VBox struct {
	Box
}

func NewVBox(homogeneous bool, spacing int) *VBox {
	return &VBox{Box{Container{Widget{C.gtk_vbox_new(gbool(homogeneous), gint(spacing))}}}}
}

//-----------------------------------------------------------------------
// GtkHButtonBox
//-----------------------------------------------------------------------

// gtk_hbutton_box_new
// gtk_hbutton_box_get_spacing_default
// gtk_hbutton_box_get_layout_default
// gtk_hbutton_box_set_spacing_default
// gtk_hbutton_box_set_layout_default

//-----------------------------------------------------------------------
// GtkVButtonBox
//-----------------------------------------------------------------------

// gtk_vbutton_box_new
// gtk_vbutton_box_get_spacing_default
// gtk_vbutton_box_set_spacing_default
// gtk_vbutton_box_get_layout_default
// gtk_vbutton_box_set_layout_default

//-----------------------------------------------------------------------
// GtkFixed
//-----------------------------------------------------------------------
type Fixed struct {
	Container
}

func NewFixed() *Fixed {
	return &Fixed{Container{Widget{C.gtk_fixed_new()}}}
}

func (v *Fixed) Put(w IWidget, x, y int) {
	C.gtk_fixed_put(FIXED(v), ToNative(w), gint(x), gint(y))
}

func (v *Fixed) Move(w IWidget, x, y int) {
	C.gtk_fixed_move(FIXED(v), ToNative(w), gint(x), gint(y))
}

//Deprecated since 2.20. Use GtkWidget.GetHasWindow() instead.
/*GtkFixed gets GetHasWindow() from anonymous field so this method can be commented out.
func (v *GtkFixed) GetHasWindow() bool {
	deprecated_since(2,20,0,"gtk_fixed_get_has_window()")
	return gobool(C.gtk_fixed_get_has_window(FIXED(v)))
}*/
//Deprecated since 2.20. Use GtkWidget.SetHasWindow() instead.
/*GtkFixed gets SetHasWindow() from anonymous field so this method can be commented out.
func (v *GtkFixed) SetHasWindow(has_window bool) {
	deprecated_since(2,20,0,"gtk_fixed_set_has_window()")
	C.gtk_fixed_set_has_window(FIXED(v), gbool(has_window))
}*/

//-----------------------------------------------------------------------
// GtkHPaned
//-----------------------------------------------------------------------
type HPaned struct {
	Paned
}

func NewHPaned() *HPaned {
	return &HPaned{Paned{Container{Widget{C.gtk_hpaned_new()}}}}
}

//-----------------------------------------------------------------------
// GtkVPaned
//-----------------------------------------------------------------------
type VPaned struct {
	Paned
}

func NewVPaned() *VPaned {
	return &VPaned{Paned{Container{Widget{C.gtk_vpaned_new()}}}}
}

//-----------------------------------------------------------------------
// GtkLayout
//-----------------------------------------------------------------------

type Layout struct {
	Container
}

func NewLayout(hadjustment *Adjustment, vadjustment *Adjustment) *Layout {
	var had, vad *C.GtkAdjustment
	if hadjustment != nil {
		had = hadjustment.GAdjustment
	}
	if vadjustment != nil {
		vad = vadjustment.GAdjustment
	}
	return &Layout{Container{Widget{
		C.gtk_layout_new(had, vad)}}}
}

func (v *Layout) GetHAdjustment() *Adjustment {
	return &Adjustment{
		C.gtk_layout_get_hadjustment(LAYOUT(v))}
}

func (v *Layout) GetVAdjustment() *Adjustment {
	return &Adjustment{
		C.gtk_layout_get_vadjustment(LAYOUT(v))}
}

func (v *Layout) SetHAdjustment(hadjustment *Adjustment) {
	C.gtk_layout_set_hadjustment(LAYOUT(v), hadjustment.GAdjustment)
}

func (v *Layout) SetVAdjustment(vadjustment *Adjustment) {
	C.gtk_layout_set_vadjustment(LAYOUT(v), vadjustment.GAdjustment)
}

func (v *Layout) Move(child IWidget, x int, y int) {
	C.gtk_layout_move(LAYOUT(v), ToNative(child), C.gint(x), C.gint(y))
}

func (v *Layout) Put(child IWidget, x int, y int) {
	C.gtk_layout_put(LAYOUT(v), ToNative(child), C.gint(x), C.gint(y))
}

func (v *Layout) SetSize(width int, height int) {
	C.gtk_layout_set_size(LAYOUT(v), C.guint(width), C.guint(height))
}

func (v *Layout) Freeze() {
	C.gtk_layout_freeze(LAYOUT(v))
}

func (v *Layout) Thaw() {
	C.gtk_layout_thaw(LAYOUT(v))
}

func (v *Layout) GetSize(width *int, height *int) {
	var cwidth, cheight C.guint
	C.gtk_layout_get_size(LAYOUT(v), &cwidth, &cheight)
	*width = int(cwidth)
	*height = int(cheight)
}

func (v *Layout) GetBinWindow() *Window {
	return &Window{Bin{Container{Widget{
		C.toGWidget(unsafe.Pointer(C.gtk_layout_get_bin_window(LAYOUT(v))))}}}}
}

//-----------------------------------------------------------------------
// GtkNotebook
//-----------------------------------------------------------------------
type Notebook struct {
	Container
}

func NewNotebook() *Notebook {
	return &Notebook{Container{Widget{
		C.gtk_notebook_new()}}}
}

func (v *Notebook) AppendPage(child IWidget, tab_label IWidget) int {
	return int(C.gtk_notebook_append_page(NOTEBOOK(v), ToNative(child), ToNative(tab_label)))
}

func (v *Notebook) AppendPageMenu(child IWidget, tab_label IWidget, menu_label IWidget) int {
	return int(C.gtk_notebook_append_page_menu(NOTEBOOK(v), ToNative(child), ToNative(tab_label), ToNative(menu_label)))
}

func (v *Notebook) PrependPage(child IWidget, tab_label IWidget) int {
	return int(C.gtk_notebook_prepend_page(NOTEBOOK(v), ToNative(child), ToNative(tab_label)))
}

func (v *Notebook) PrependPageMenu(child IWidget, tab_label IWidget, menu_label IWidget) int {
	return int(C.gtk_notebook_prepend_page_menu(NOTEBOOK(v), ToNative(child), ToNative(tab_label), ToNative(menu_label)))
}

func (v *Notebook) InsertPage(child IWidget, tab_label IWidget, position int) int {
	return int(C.gtk_notebook_insert_page(NOTEBOOK(v), ToNative(child), ToNative(tab_label), gint(position)))
}

func (v *Notebook) InsertPageMenu(child IWidget, tab_label IWidget, menu_label IWidget, position int) int {
	return int(C.gtk_notebook_insert_page_menu(NOTEBOOK(v), ToNative(child), ToNative(tab_label), ToNative(menu_label), gint(position)))
}

func (v *Notebook) RemovePage(child IWidget, page_num int) {
	C.gtk_notebook_remove_page(NOTEBOOK(v), gint(page_num))
}

func (v *Notebook) PageNum(child IWidget) int {
	return int(C.gtk_notebook_page_num(NOTEBOOK(v), ToNative(child)))
}

func (v *Notebook) NextPage() {
	C.gtk_notebook_next_page(NOTEBOOK(v))
}

func (v *Notebook) PrevPage() {
	C.gtk_notebook_prev_page(NOTEBOOK(v))
}

func (v *Notebook) ReorderChild(child IWidget, position int) {
	C.gtk_notebook_reorder_child(NOTEBOOK(v), ToNative(child), gint(position))
}

func (v *Notebook) SetTabPos(pos PositionType) {
	C.gtk_notebook_set_tab_pos(NOTEBOOK(v), C.GtkPositionType(pos))
}

func (v *Notebook) SetShowTabs(show_tabs bool) {
	C.gtk_notebook_set_show_tabs(NOTEBOOK(v), gbool(show_tabs))
}

func (v *Notebook) SetShowBorder(show_border bool) {
	C.gtk_notebook_set_show_border(NOTEBOOK(v), gbool(show_border))
}

func (v *Notebook) SetScrollable(scrollable bool) {
	C.gtk_notebook_set_scrollable(NOTEBOOK(v), gbool(scrollable))
}

//Deprecated.
func (v *Notebook) SetTabBorder(border_width uint) {
	deprecated_since(2, 0, 0, "gtk_notebook_set_tab_border()")
	C.gtk_notebook_set_tab_border(NOTEBOOK(v), guint(border_width))
}

func (v *Notebook) PopupEnable() {
	C.gtk_notebook_popup_enable(NOTEBOOK(v))
}

func (v *Notebook) PopupDisable() {
	C.gtk_notebook_popup_disable(NOTEBOOK(v))
}

func (v *Notebook) GetCurrentPage() int {
	return int(C.gtk_notebook_get_current_page(NOTEBOOK(v)))
}

func (v *Notebook) GetMenuLabel(child IWidget) *Widget {
	return &Widget{
		C.gtk_notebook_get_menu_label(NOTEBOOK(v), ToNative(child))}
}

func (v *Notebook) GetNthPage(page_num int) *Widget {
	return &Widget{
		C.gtk_notebook_get_nth_page(NOTEBOOK(v), gint(page_num))}
}

func (v *Notebook) GetNPages() int {
	return int(C.gtk_notebook_get_n_pages(NOTEBOOK(v)))
}

func (v *Notebook) GetTabLabel(child IWidget) *Widget {
	return &Widget{
		C.gtk_notebook_get_tab_label(NOTEBOOK(v), ToNative(child))}
}

//Deprecated since 2.20. Modify the "tab-expand" and "tab-fill" child properties instead.
func (v *Notebook) QueryTabLabelPacking(child IWidget, expand *bool, fill *bool, pack_type *uint) {
	deprecated_since(2, 20, 0, "gtk_notebook_query_tab_label_packing()")
	var cexpand, cfill C.gboolean
	var cpack_type C.GtkPackType
	C.gtk_notebook_query_tab_label_packing(NOTEBOOK(v), ToNative(child), &cexpand, &cfill, &cpack_type)
	*expand = gobool(cexpand)
	*fill = gobool(cfill)
	*pack_type = uint(cpack_type)
}

//Deprecated.
func (v *Notebook) SetHomogeneousTabs(homogeneous bool) {
	deprecated_since(2, 0, 0, "gtk_notebook_set_homogeneous_tabs()")
	C.gtk_notebook_set_homogeneous_tabs(NOTEBOOK(v), gbool(homogeneous))
}

func (v *Notebook) SetMenuLabel(child IWidget, menu_label IWidget) {
	C.gtk_notebook_set_menu_label(NOTEBOOK(v), ToNative(child), ToNative(menu_label))
}

func (v *Notebook) SetMenuLabelText(child IWidget, name string) {
	ptr := C.CString(name)
	defer cfree(ptr)
	C.gtk_notebook_set_menu_label_text(NOTEBOOK(v), ToNative(child), gstring(ptr))
}

//Deprecated.
func (v *Notebook) SetTabHBorder(tab_hborder uint) {
	deprecated_since(2, 0, 0, "gtk_notebook_set_tab_hborder()")
	C.gtk_notebook_set_tab_hborder(NOTEBOOK(v), guint(tab_hborder))
}

func (v *Notebook) SetTabLabel(child IWidget, tab_label IWidget) {
	C.gtk_notebook_set_tab_label(NOTEBOOK(v), ToNative(child), ToNative(tab_label))
}

//Deprecated since 2.20. Modify the "tab-expand" and "tab-fill" child properties instead.
func (v *Notebook) SetTabLabelPacking(child IWidget, expand bool, fill bool, pack_type uint) {
	deprecated_since(2, 20, 0, "gtk_notebook_set_tab_label_packing()")
	C.gtk_notebook_set_tab_label_packing(NOTEBOOK(v), ToNative(child), gbool(expand), gbool(fill), C.GtkPackType(pack_type))
}

func (v *Notebook) SetTabLabelText(child IWidget, name string) {
	ptr := C.CString(name)
	defer cfree(ptr)
	C.gtk_notebook_set_tab_label_text(NOTEBOOK(v), ToNative(child), gstring(ptr))
}

//Deprecated.
func (v *Notebook) SetTabVBorder(tab_vborder uint) {
	deprecated_since(2, 0, 0, "gtk_notebook_set_tab_vborder()")
	C.gtk_notebook_set_tab_vborder(NOTEBOOK(v), guint(tab_vborder))
}

func (v *Notebook) SetReorderable(child IWidget, reorderable bool) {
	C.gtk_notebook_set_tab_reorderable(NOTEBOOK(v), ToNative(child), gbool(reorderable))
}

func (v *Notebook) SetTabDetachable(child IWidget, detachable bool) {
	C.gtk_notebook_set_tab_detachable(NOTEBOOK(v), ToNative(child), gbool(detachable))
}

func (v *Notebook) GetMenuLabelText(child IWidget) string {
	return gostring(C.gtk_notebook_get_menu_label_text(NOTEBOOK(v), ToNative(child)))
}

func (v *Notebook) GetScrollable() bool {
	return gobool(C.gtk_notebook_get_scrollable(NOTEBOOK(v)))
}

func (v *Notebook) GetShowBorder() bool {
	return gobool(C.gtk_notebook_get_show_border(NOTEBOOK(v)))
}

func (v *Notebook) GetShowTabs() bool {
	return gobool(C.gtk_notebook_get_show_tabs(NOTEBOOK(v)))
}

func (v *Notebook) GetTabLabelText(child IWidget) string {
	return gostring(C.gtk_notebook_get_tab_label_text(NOTEBOOK(v), ToNative(child)))
}

func (v *Notebook) GetTabPos() uint {
	return uint(C.gtk_notebook_get_tab_pos(NOTEBOOK(v)))
}

func (v *Notebook) GetTabReorderable(child IWidget) bool {
	return gobool(C.gtk_notebook_get_tab_reorderable(NOTEBOOK(v), ToNative(child)))
}

func (v *Notebook) GetTabDetachable(child IWidget) bool {
	return gobool(C.gtk_notebook_get_tab_detachable(NOTEBOOK(v), ToNative(child)))
}

// gtk_notebook_get_tab_hborder //since 2.22
// gtk_notebook_get_tab_vborder //since 2.22

func (v *Notebook) SetCurrentPage(pageNum int) {
	C.gtk_notebook_set_current_page(NOTEBOOK(v), gint(pageNum))
}

//Deprecated since 2.12, use SetGroupName() instead
func (v *Notebook) SetGroupId(group_id int) {
	deprecated_since(2, 12, 0, "gtk_notebook_set_group_id()")
	C.gtk_notebook_set_group_id(NOTEBOOK(v), gint(group_id))
}

//Deprecated since 2.12, use GetGroupName() instead
func (v *Notebook) GetGroupId() int {
	deprecated_since(2, 12, 0, "gtk_notebook_get_group_id()")
	return int(C.gtk_notebook_get_group_id(NOTEBOOK(v)))
}

//Deprecated since 2.24, use SetGroupName() instead
func (v *Notebook) SetGroup(group unsafe.Pointer) {
	deprecated_since(2, 24, 0, "gtk_notebook_set_group()")
	C.gtk_notebook_set_group(NOTEBOOK(v), C.gpointer(group))
}

//Deprecated since 2.24, use GetGroupName() instead
func (v *Notebook) GetGroup() unsafe.Pointer {
	deprecated_since(2, 24, 0, "gtk_notebook_get_group()")
	return unsafe.Pointer(C.gtk_notebook_get_group(NOTEBOOK(v)))
}

func (v *Notebook) SetGroupName(group string) {
	panic_if_version_older(2, 24, 0, "gtk_notebook_set_group_name()")
	ptr := C.CString(group)
	defer cfree(ptr)
	C._gtk_notebook_set_group_name(NOTEBOOK(v), gstring(ptr))
}

func (v *Notebook) GetGroupName() string {
	panic_if_version_older(2, 24, 0, "gtk_notebook_get_group_name()")
	return gostring(C._gtk_notebook_get_group_name(NOTEBOOK(v)))
}

// gtk_notebook_set_action_widget //since 2.20
// gtk_notebook_get_action_widget //since 2.20
// void gtk_notebook_set_window_creation_hook (GtkNotebookWindowCreationFunc func, gpointer data, GDestroyNotify destroy); //deprecated in 2.24

//-----------------------------------------------------------------------
// GtkTable
//-----------------------------------------------------------------------
type AttachOptions int

const (
	EXPAND AttachOptions = 1 << iota
	SHRINK
	FILL
)

type Table struct {
	Container
}

func NewTable(rows uint, columns uint, homogeneous bool) *Table {
	return &Table{Container{Widget{
		C.gtk_table_new(guint(rows), guint(columns), gbool(homogeneous))}}}
}

func (v *Table) Resize(rows uint, columns uint) {
	C.gtk_table_resize(TABLE(v), guint(rows), guint(columns))
}

func (v *Table) Attach(child IWidget, left_attach uint, right_attach uint, top_attach uint, bottom_attach uint, xoptions AttachOptions, yoptions AttachOptions, xpadding uint, ypadding uint) {
	C.gtk_table_attach(TABLE(v), ToNative(child), guint(left_attach), guint(right_attach), guint(top_attach), guint(bottom_attach), C.GtkAttachOptions(xoptions), C.GtkAttachOptions(yoptions), guint(xpadding), guint(ypadding))
}

func (v *Table) AttachDefaults(child IWidget, left_attach uint, right_attach uint, top_attach uint, bottom_attach uint) {
	C.gtk_table_attach_defaults(TABLE(v), ToNative(child), guint(left_attach), guint(right_attach), guint(top_attach), guint(bottom_attach))
}

func (v *Table) SetRowSpacing(row uint, spacing uint) {
	C.gtk_table_set_row_spacing(TABLE(v), guint(row), guint(spacing))
}

func (v *Table) SetColSpacing(column uint, spacing uint) {
	C.gtk_table_set_col_spacing(TABLE(v), guint(column), guint(spacing))
}

func (v *Table) SetRowSpacings(spacing uint) {
	C.gtk_table_set_row_spacings(TABLE(v), guint(spacing))
}

func (v *Table) SetColSpacings(spacing uint) {
	C.gtk_table_set_col_spacings(TABLE(v), guint(spacing))
}

func (v *Table) SetHomogeneous(homogeneous bool) {
	C.gtk_table_set_homogeneous(TABLE(v), gbool(homogeneous))
}

func (v *Table) GetDefaultRowSpacing() uint {
	return uint(C.gtk_table_get_default_row_spacing(TABLE(v)))
}

func (v *Table) GetHomogeneous() bool {
	return gobool(C.gtk_table_get_homogeneous(TABLE(v)))
}

func (v *Table) GetRowSpacing(row uint) uint {
	return uint(C.gtk_table_get_row_spacing(TABLE(v), guint(row)))
}

func (v *Table) GetColSpacing(column uint) uint {
	return uint(C.gtk_table_get_col_spacing(TABLE(v), guint(column)))
}

func (v *Table) GetDefaultColSpacing() uint {
	return uint(C.gtk_table_get_default_col_spacing(TABLE(v)))
}

// gtk_table_get_size //since 2.22

//-----------------------------------------------------------------------
// GtkExpander
//-----------------------------------------------------------------------
type Expander struct {
	Bin
}

func NewExpander(label string) *Expander {
	var ptr *C.char
	if len(label) > 0 {
		ptr = C.CString(label)
		defer cfree(ptr)
	}
	return &Expander{Bin{Container{Widget{C.gtk_expander_new(gstring(ptr))}}}}
}

func NewExpanderWithMnemonic(label string) *Expander {
	var ptr *C.char
	if len(label) > 0 {
		ptr = C.CString(label)
		defer cfree(ptr)
	}
	return &Expander{Bin{Container{Widget{C.gtk_expander_new_with_mnemonic(gstring(ptr))}}}}
}

func (v *Expander) SetExpanded(expanded bool) {
	C.gtk_expander_set_expanded(EXPANDER(v), gbool(expanded))
}

func (v *Expander) GetExpanded() bool {
	return gobool(C.gtk_expander_get_expanded(EXPANDER(v)))
}

func (v *Expander) SetSpacing(spacing int) {
	C.gtk_expander_set_spacing(EXPANDER(v), gint(spacing))
}

func (v *Expander) GetSpacing() int {
	return int(C.gtk_expander_get_spacing(EXPANDER(v)))
}

func (v *Expander) SetLabel(label string) {
	var ptr *C.char
	if len(label) > 0 {
		ptr = C.CString(label)
		defer cfree(ptr)
	}
	C.gtk_expander_set_label(EXPANDER(v), gstring(ptr))
}

func (v *Expander) GetLabel() string {
	return gostring(C.gtk_expander_get_label(EXPANDER(v)))
}

func (v *Expander) SetUseUnderline(setting bool) {
	C.gtk_expander_set_use_underline(EXPANDER(v), gbool(setting))
}

func (v *Expander) GetUseUnderline() bool {
	return gobool(C.gtk_expander_get_use_underline(EXPANDER(v)))
}

func (v *Expander) SetUseMarkup(setting bool) {
	C.gtk_expander_set_use_markup(EXPANDER(v), gbool(setting))
}

func (v *Expander) GetUseMarkup() bool {
	return gobool(C.gtk_expander_get_use_markup(EXPANDER(v)))
}

func (v *Expander) SetLabelWidget(label_widget ILabel) {
	C.gtk_expander_set_label_widget(EXPANDER(v), ToNative(label_widget))
}

func (v *Expander) GetLabelWidget() ILabel {
	return &Label{Misc{Widget{
		C.gtk_expander_get_label_widget(EXPANDER(v))}}}
}

// gtk_expander_set_label_fill //since 2.22
// gtk_expander_get_label_fill //since 2.22

//-----------------------------------------------------------------------
// GtkOrientable
//-----------------------------------------------------------------------

// gtk_orientable_get_orientation
// gtk_orientable_set_orientation

//-----------------------------------------------------------------------
// GtkFrame
//-----------------------------------------------------------------------
type ShadowType int

const (
	SHADOW_NONE ShadowType = iota
	SHADOW_IN
	SHADOW_OUT
	SHADOW_ETCHED_IN
	SHADOW_ETCHED_OUT
)

type Frame struct {
	Bin
}

func NewFrame(label string) *Frame {
	var ptr *C.char
	if len(label) > 0 {
		ptr = C.CString(label)
		defer cfree(ptr)
	}
	return &Frame{Bin{Container{Widget{C.gtk_frame_new(gstring(ptr))}}}}
}

func (v *Frame) SetLabel(label string) {
	var ptr *C.char
	if len(label) > 0 {
		ptr = C.CString(label)
		defer cfree(ptr)
	}
	C.gtk_frame_set_label(FRAME(v), gstring(ptr))
}

func (v *Frame) SetLabelWidget(label_widget ILabel) {
	C.gtk_frame_set_label_widget(FRAME(v), ToNative(label_widget))
}

func (v *Frame) SetLabelAlign(xalign, yalign float64) {
	C.gtk_frame_set_label_align(FRAME(v), C.gfloat(xalign), C.gfloat(yalign))
}

func (v *Frame) SetShadowType(shadow_type ShadowType) {
	C.gtk_frame_set_shadow_type(FRAME(v), C.GtkShadowType(shadow_type))
}

func (v *Frame) GetLabel() string {
	return gostring(C.gtk_frame_get_label(FRAME(v)))
}

func (v *Frame) GetLabelAlign() (xalign, yalign float64) {
	var xalign_, yalign_ C.gfloat
	C.gtk_frame_get_label_align(FRAME(v), &xalign_, &yalign_)
	return float64(xalign_), float64(yalign_)
}

func (v *Frame) GetLabelWidget() ILabel {
	return &Label{Misc{Widget{
		C.gtk_frame_get_label_widget(FRAME(v))}}}
}

func (v *Frame) GetShadowType() ShadowType {
	return ShadowType(C.gtk_frame_get_shadow_type(FRAME(v)))
}

//-----------------------------------------------------------------------
// GtkHSeparator
//-----------------------------------------------------------------------
type HSeparator struct {
	Separator
}

func NewHSeparator() *HSeparator {
	return &HSeparator{Separator{Widget{C.gtk_hseparator_new()}}}
}

//-----------------------------------------------------------------------
// GtkVSeparator
//-----------------------------------------------------------------------
type VSeparator struct {
	Separator
}

func NewVSeparator() *VSeparator {
	return &VSeparator{Separator{Widget{C.gtk_vseparator_new()}}}
}

//-----------------------------------------------------------------------
// GtkScrollbar
//-----------------------------------------------------------------------
type Scrollbar struct {
	Range
}

//-----------------------------------------------------------------------
// GtkHScrollbar
//-----------------------------------------------------------------------
type HScrollbar struct {
	Scrollbar
}

func NewHScrollbar(ha *Adjustment) *HScrollbar {
	return &HScrollbar{Scrollbar{Range{Widget{C.gtk_hscrollbar_new(ADJUSTMENT(ha))}}}}
}

//-----------------------------------------------------------------------
// GtkVScrollbar
//-----------------------------------------------------------------------
type VScrollbar struct {
	Scrollbar
}

func NewVScrollbar(va *Adjustment) *VScrollbar {
	return &VScrollbar{Scrollbar{Range{Widget{C.gtk_vscrollbar_new(ADJUSTMENT(va))}}}}
}

//-----------------------------------------------------------------------
// GtkScrolledWindow
//-----------------------------------------------------------------------
type PolicyType int

const (
	POLICY_ALWAYS PolicyType = iota
	POLICY_AUTOMATIC
	POLICY_NEVER
)

type CornerType int

const (
	CORNER_TOP_LEFT CornerType = iota
	CORNER_BOTTOM_LEFT
	CORNER_TOP_RIGHT
	CORNER_BOTTOM_RIGHT
)

type ScrolledWindow struct {
	Bin
}

func NewScrolledWindow(hadjustment *Adjustment, vadjustment *Adjustment) *ScrolledWindow {
	var had, vad *C.GtkAdjustment
	if hadjustment != nil {
		had = hadjustment.GAdjustment
	}
	if vadjustment != nil {
		vad = vadjustment.GAdjustment
	}
	return &ScrolledWindow{Bin{Container{Widget{
		C.gtk_scrolled_window_new(had, vad)}}}}
}

func (v *ScrolledWindow) GetHAdjustment() *Adjustment {
	return &Adjustment{
		C.gtk_scrolled_window_get_hadjustment(SCROLLED_WINDOW(v))}
}

func (v *ScrolledWindow) GetVAdjustment() *Adjustment {
	return &Adjustment{
		C.gtk_scrolled_window_get_vadjustment(SCROLLED_WINDOW(v))}
}

// gtk_scrolled_window_get_hscrollbar
// gtk_scrolled_window_get_vscrollbar

func (v *ScrolledWindow) SetPolicy(hscrollbar_policy PolicyType, vscrollbar_policy PolicyType) {
	C.gtk_scrolled_window_set_policy(SCROLLED_WINDOW(v), C.GtkPolicyType(hscrollbar_policy), C.GtkPolicyType(vscrollbar_policy))
}

func (v *ScrolledWindow) AddWithViewPort(w IWidget) {
	C.gtk_scrolled_window_add_with_viewport(SCROLLED_WINDOW(v), ToNative(w))
}

func (v *ScrolledWindow) SetPlacement(window_placement CornerType) {
	C.gtk_scrolled_window_set_placement(SCROLLED_WINDOW(v), C.GtkCornerType(window_placement))
}

func (v *ScrolledWindow) UnsetPlacement() {
	C.gtk_scrolled_window_unset_placement(SCROLLED_WINDOW(v))
}

func (v *ScrolledWindow) SetShadowType(typ ShadowType) {
	C.gtk_scrolled_window_set_shadow_type(SCROLLED_WINDOW(v), C.GtkShadowType(typ))
}

func (v *ScrolledWindow) SetHAdjustment(hadjustment *Adjustment) {
	C.gtk_scrolled_window_set_hadjustment(SCROLLED_WINDOW(v), hadjustment.GAdjustment)
}

func (v *ScrolledWindow) SetVAdjustment(vadjustment *Adjustment) {
	C.gtk_scrolled_window_set_vadjustment(SCROLLED_WINDOW(v), vadjustment.GAdjustment)
}

func (v *ScrolledWindow) GetPlacement() CornerType {
	return CornerType(C.gtk_scrolled_window_get_placement(SCROLLED_WINDOW(v)))
}

func (v *ScrolledWindow) GetPolicy(hscrollbar_policy *PolicyType, vscrollbar_policy *PolicyType) {
	var chscrollbar_policy, cvscrollbar_policy C.GtkPolicyType
	C.gtk_scrolled_window_get_policy(SCROLLED_WINDOW(v), &chscrollbar_policy, &cvscrollbar_policy)
	*hscrollbar_policy = PolicyType(chscrollbar_policy)
	*vscrollbar_policy = PolicyType(cvscrollbar_policy)
}

func (v *ScrolledWindow) GetShadowType() ShadowType {
	return ShadowType(C.gtk_scrolled_window_get_shadow_type(SCROLLED_WINDOW(v)))
}

//-----------------------------------------------------------------------
// GtkPrintOperation
//-----------------------------------------------------------------------

type PrintOperation struct {
	GPrintOperation *C.GtkPrintOperation
}

type PrintOperationResult int

const (
	PRINT_OPERATION_RESULT_ERROR PrintOperationResult = iota
	PRINT_OPERATION_RESULT_APPLY
	PRINT_OPERATION_RESULT_CANCEL
	PRINT_OPERATION_RESULT_IN_PROGRESS
)

type PrintOperationAction int

const (
	PRINT_OPERATION_ACTION_PRINT_DIALOG PrintOperationAction = iota
	PRINT_OPERATION_ACTION_PRINT
	PRINT_OPERATION_ACTION_PREVIEW
	PRINT_OPERATION_ACTION_EXPOR
)

func NewPrintOperation() *PrintOperation {
	return &PrintOperation{C.gtk_print_operation_new()}
}

func (v *PrintOperation) Run(action PrintOperationAction, parent *Window) (result PrintOperationResult, err error) {
	var gerror *C.GError
	ret := PrintOperationResult(
		C.gtk_print_operation_run(
			v.GPrintOperation,
			C.GtkPrintOperationAction(action),
			C.toGWindow(parent.GWidget),
			&gerror))
	if gerror != nil {
		err = glib.ErrorFromNative(unsafe.Pointer(gerror))
	}
	return PrintOperationResult(ret), err
}

func (v *PrintOperation) Cancel() {
	C.gtk_print_operation_cancel(v.GPrintOperation)
}

func (v *PrintOperation) IsFinished() bool {
	return gobool(C.gtk_print_operation_is_finished(v.GPrintOperation))
}

func (v *PrintOperation) Connect(s string, f interface{}, datas ...interface{}) int {
	return glib.ObjectFromNative(unsafe.Pointer(v.GPrintOperation)).Connect(s, f, datas...)
}

// gtk_print_operation_set_allow_async
// gtk_print_operation_get_error
// gtk_print_operation_set_default_page_setup
// gtk_print_operation_get_default_page_setup
// gtk_print_operation_set_print_settings
// gtk_print_operation_get_print_settings
// gtk_print_operation_set_job_name
// gtk_print_operation_set_n_pages
// gtk_print_operation_get_n_pages_to_print //since 2.18
// gtk_print_operation_set_current_page
// gtk_print_operation_set_use_full_page
// gtk_print_operation_set_unit
// gtk_print_operation_set_export_filename
// gtk_print_operation_set_show_progress
// gtk_print_operation_set_track_print_status
// gtk_print_operation_set_custom_tab_label
// gtk_print_operation_draw_page_finish
// gtk_print_operation_set_defer_drawing
// gtk_print_operation_get_status
// gtk_print_operation_get_status_string
// gtk_print_operation_set_support_selection //since 2.18
// gtk_print_operation_get_support_selection //since 2.18
// gtk_print_operation_set_has_selection //since 2.18
// gtk_print_operation_get_has_selection //since 2.18
// gtk_print_operation_set_embed_page_setup //since 2.18
// gtk_print_operation_get_embed_page_setup //since 2.18
// gtk_print_run_page_setup_dialog
// gtk_print_run_page_setup_dialog_async
// gtk_print_operation_preview_end_preview
// gtk_print_operation_preview_is_selected
// gtk_print_operation_preview_render_page

//-----------------------------------------------------------------------
// GtkPrintContext
//-----------------------------------------------------------------------
type PrintContext struct {
	GPrintContext *C.GtkPrintContext
}

func (v *PrintContext) GetCairoContext() *C.cairo_t {
	return C.gtk_print_context_get_cairo_context(v.GPrintContext)
}

func (v *PrintContext) SetCairoContext(cairo *C.cairo_t, dpi_x float64, dpi_y float64) {
	C.gtk_print_context_set_cairo_context(v.GPrintContext, cairo, C.double(dpi_x), C.double(dpi_y))
}

// gtk_print_context_get_page_setup
// gtk_print_context_get_width
// gtk_print_context_get_height
// gtk_print_context_get_dpi_x
// gtk_print_context_get_dpi_y
// gtk_print_context_get_pango_fontmap
// gtk_print_context_create_pango_context
// gtk_print_context_create_pango_layout
// gtk_print_context_get_hard_margins //since 2.20

//-----------------------------------------------------------------------
// GtkPrintSettings
//-----------------------------------------------------------------------

// gtk_print_settings_new
// gtk_print_settings_copy
// gtk_print_settings_has_key
// gtk_print_settings_get
// gtk_print_settings_set
// gtk_print_settings_unset
// gtk_print_settings_foreach
// gtk_print_settings_get_bool
// gtk_print_settings_set_bool
// gtk_print_settings_get_double
// gtk_print_settings_get_double_with_default
// gtk_print_settings_set_double
// gtk_print_settings_get_length
// gtk_print_settings_set_length
// gtk_print_settings_get_int
// gtk_print_settings_get_int_with_default
// gtk_print_settings_set_int
// gtk_print_settings_get_printer
// gtk_print_settings_set_printer
// gtk_print_settings_get_orientation
// gtk_print_settings_set_orientation
// gtk_print_settings_get_paper_size
// gtk_print_settings_set_paper_size
// gtk_print_settings_get_paper_width
// gtk_print_settings_set_paper_width
// gtk_print_settings_get_paper_height
// gtk_print_settings_set_paper_height
// gtk_print_settings_get_use_color
// gtk_print_settings_set_use_color
// gtk_print_settings_get_collate
// gtk_print_settings_set_collate
// gtk_print_settings_get_reverse
// gtk_print_settings_set_reverse
// gtk_print_settings_get_duplex
// gtk_print_settings_set_duplex
// gtk_print_settings_get_quality
// gtk_print_settings_set_quality
// gtk_print_settings_get_n_copies
// gtk_print_settings_set_n_copies
// gtk_print_settings_get_number_up
// gtk_print_settings_set_number_up
// gtk_print_settings_get_number_up_layout
// gtk_print_settings_set_number_up_layout
// gtk_print_settings_get_resolution
// gtk_print_settings_set_resolution
// gtk_print_settings_set_resolution_xy
// gtk_print_settings_get_resolution_x
// gtk_print_settings_get_resolution_y
// gtk_print_settings_get_printer_lpi
// gtk_print_settings_set_printer_lpi
// gtk_print_settings_get_scale
// gtk_print_settings_set_scale
// gtk_print_settings_get_print_pages
// gtk_print_settings_set_print_pages
// gtk_print_settings_get_page_ranges
// gtk_print_settings_set_page_ranges
// gtk_print_settings_get_page_set
// gtk_print_settings_set_page_set
// gtk_print_settings_get_default_source
// gtk_print_settings_set_default_source
// gtk_print_settings_get_media_type
// gtk_print_settings_set_media_type
// gtk_print_settings_get_dither
// gtk_print_settings_set_dither
// gtk_print_settings_get_finishings
// gtk_print_settings_set_finishings
// gtk_print_settings_get_output_bin
// gtk_print_settings_set_output_bin
// gtk_print_settings_new_from_file
// gtk_print_settings_new_from_key_file
// gtk_print_settings_load_file
// gtk_print_settings_load_key_file
// gtk_print_settings_to_file
// gtk_print_settings_to_key_file

//-----------------------------------------------------------------------
// GtkPageSetup
//-----------------------------------------------------------------------

// gtk_page_setup_new
// gtk_page_setup_copy
// gtk_page_setup_get_orientation
// gtk_page_setup_set_orientation
// gtk_page_setup_get_paper_size
// gtk_page_setup_set_paper_size
// gtk_page_setup_get_top_margin
// gtk_page_setup_set_top_margin
// gtk_page_setup_get_bottom_margin
// gtk_page_setup_set_bottom_margin
// gtk_page_setup_get_left_margin
// gtk_page_setup_set_left_margin
// gtk_page_setup_get_right_margin
// gtk_page_setup_set_right_margin
// gtk_page_setup_set_paper_size_and_default_margins
// gtk_page_setup_get_paper_width
// gtk_page_setup_get_paper_height
// gtk_page_setup_get_page_width
// gtk_page_setup_get_page_height
// gtk_page_setup_new_from_file
// gtk_page_setup_new_from_key_file
// gtk_page_setup_load_file
// gtk_page_setup_load_key_file
// gtk_page_setup_to_file
// gtk_page_setup_to_key_file

//-----------------------------------------------------------------------
// GtkPaperSize
//-----------------------------------------------------------------------

// gtk_paper_size_new
// gtk_paper_size_new_from_ppd
// gtk_paper_size_new_custom
// gtk_paper_size_copy
// gtk_paper_size_free
// gtk_paper_size_is_equal
// gtk_paper_size_get_paper_sizes
// gtk_paper_size_get_name
// gtk_paper_size_get_display_name
// gtk_paper_size_get_ppd_name
// gtk_paper_size_get_width
// gtk_paper_size_get_height
// gtk_paper_size_is_custom
// gtk_paper_size_set_size
// gtk_paper_size_get_default_top_margin
// gtk_paper_size_get_default_bottom_margin
// gtk_paper_size_get_default_left_margin
// gtk_paper_size_get_default_right_margin
// gtk_paper_size_get_default
// gtk_paper_size_new_from_key_file
// gtk_paper_size_to_key_file

//-----------------------------------------------------------------------
// GtkPrinter
//-----------------------------------------------------------------------

// gtk_printer_new
// gtk_printer_get_backend
// gtk_printer_get_name
// gtk_printer_get_state_message
// gtk_printer_get_description
// gtk_printer_get_location
// gtk_printer_get_icon_name
// gtk_printer_get_job_count
// gtk_printer_is_active
// gtk_printer_is_paused
// gtk_printer_is_accepting_jobs
// gtk_printer_is_virtual
// gtk_printer_is_default
// gtk_printer_accepts_ps
// gtk_printer_accepts_pdf
// gtk_printer_list_papers
// gtk_printer_compare
// gtk_printer_has_details
// gtk_printer_request_details
// gtk_printer_get_capabilities
// gtk_printer_get_default_page_size
// gtk_printer_get_hard_margins
// gtk_enumerate_printers

//-----------------------------------------------------------------------
// GtkPrintJob
//-----------------------------------------------------------------------

// gtk_print_job_new
// gtk_print_job_get_settings
// gtk_print_job_get_printer
// gtk_print_job_get_title
// gtk_print_job_get_status
// gtk_print_job_set_source_file
// gtk_print_job_get_surface
// gtk_print_job_send
// gtk_print_job_set_track_print_status
// gtk_print_job_get_track_print_status

//-----------------------------------------------------------------------
// GtkPrintUnixDialog
//-----------------------------------------------------------------------

// gtk_print_unix_dialog_new
// gtk_print_unix_dialog_set_page_setup
// gtk_print_unix_dialog_get_page_setup
// gtk_print_unix_dialog_set_current_page
// gtk_print_unix_dialog_get_current_page
// gtk_print_unix_dialog_set_settings
// gtk_print_unix_dialog_get_settings
// gtk_print_unix_dialog_get_selected_printer
// gtk_print_unix_dialog_add_custom_tab
// gtk_print_unix_dialog_set_support_selection
// gtk_print_unix_dialog_get_support_selection
// gtk_print_unix_dialog_set_has_selection
// gtk_print_unix_dialog_get_has_selection
// gtk_print_unix_dialog_set_embed_page_setup
// gtk_print_unix_dialog_get_embed_page_setup
// gtk_print_unix_dialog_get_page_setup_set
// gtk_print_unix_dialog_set_manual_capabilities
// gtk_print_unix_dialog_get_manual_capabilities

//-----------------------------------------------------------------------
// GtkPageSetupUnixDialog
//-----------------------------------------------------------------------

// gtk_page_setup_unix_dialog_new
// gtk_page_setup_unix_dialog_set_page_setup
// gtk_page_setup_unix_dialog_get_page_setup
// gtk_page_setup_unix_dialog_set_print_settings
// gtk_page_setup_unix_dialog_get_print_settings

//-----------------------------------------------------------------------
// GtkAdjustment
//-----------------------------------------------------------------------
type Adjustment struct {
	GAdjustment *C.GtkAdjustment
}

func NewAdjustment(value, lower, upper, step_increment, page_increment, page_size float64) *Adjustment {
	return &Adjustment{
		C.toGAdjustment(C.gtk_adjustment_new(gdouble(value), gdouble(lower), gdouble(upper),
			gdouble(step_increment), gdouble(page_increment), gdouble(page_size)))}
}

func (v *Adjustment) GetValue() float64 {
	return float64(C.gtk_adjustment_get_value(v.GAdjustment))
}

func (v *Adjustment) SetValue(value float64) {
	C.gtk_adjustment_set_value(v.GAdjustment, gdouble(value))
}

// gtk_adjustment_clamp_page
// gtk_adjustment_changed
// gtk_adjustment_value_changed

func (v *Adjustment) Configure(value, lower, upper, step_increment, page_increment, page_size float64) {
	panic_if_version_older(2, 14, 0, "gtk_adjustment_configure()")
	C._gtk_adjustment_configure(v.GAdjustment, gdouble(value), gdouble(lower), gdouble(upper),
		gdouble(step_increment), gdouble(page_increment), gdouble(page_size))
}

func (v *Adjustment) GetLower() float64 {
	panic_if_version_older(2, 14, 0, "gtk_adjustment_get_lower()")
	return float64(C._gtk_adjustment_get_lower(v.GAdjustment))
}

func (v *Adjustment) GetPageIncrement() float64 {
	panic_if_version_older(2, 14, 0, "gtk_adjustment_get_page_increment()")
	return float64(C._gtk_adjustment_get_page_increment(v.GAdjustment))
}

func (v *Adjustment) GetPageSize() float64 {
	panic_if_version_older(2, 14, 0, "gtk_adjustment_get_page_size()")
	return float64(C._gtk_adjustment_get_page_size(v.GAdjustment))
}

func (v *Adjustment) GetStepIncrement() float64 {
	panic_if_version_older(2, 14, 0, "gtk_adjustment_get_step_increment()")
	return float64(C._gtk_adjustment_get_step_increment(v.GAdjustment))
}

func (v *Adjustment) GetUpper() float64 {
	panic_if_version_older(2, 14, 0, "gtk_adjustment_get_upper()")
	return float64(C._gtk_adjustment_get_upper(v.GAdjustment))
}

func (v *Adjustment) SetLower(lower float64) {
	panic_if_version_older(2, 14, 0, "gtk_adjustment_set_lower()")
	C._gtk_adjustment_set_lower(v.GAdjustment, gdouble(lower))
}

func (v *Adjustment) SetPageIncrement(page_increment float64) {
	panic_if_version_older(2, 14, 0, "gtk_adjustment_set_page_increment()")
	C._gtk_adjustment_set_page_increment(v.GAdjustment, gdouble(page_increment))
}

func (v *Adjustment) SetPageSize(page_size float64) {
	panic_if_version_older(2, 14, 0, "gtk_adjustment_set_page_size()")
	C._gtk_adjustment_set_page_size(v.GAdjustment, gdouble(page_size))
}

func (v *Adjustment) SetStepIncrement(step_increment float64) {
	panic_if_version_older(2, 14, 0, "gtk_adjustment_set_step_increment()")
	C._gtk_adjustment_set_step_increment(v.GAdjustment, gdouble(step_increment))
}

func (v *Adjustment) SetUpper(upper float64) {
	panic_if_version_older(2, 14, 0, "gtk_adjustment_set_upper()")
	C._gtk_adjustment_set_upper(v.GAdjustment, gdouble(upper))
}

func (v *Adjustment) Connect(s string, f interface{}, datas ...interface{}) int {
	return glib.ObjectFromNative(unsafe.Pointer(v.GAdjustment)).Connect(s, f, datas...)
}

//-----------------------------------------------------------------------
// GtkArrow
//-----------------------------------------------------------------------

type Arrow struct {
	Widget
}

type ArrowType int

const (
	ARROW_UP ArrowType = iota
	ARROW_DOWN
	ARROW_LEFT
	ARROW_RIGHT
	ARROW_NONE
)

// Draw arrowhead facing in given direction with a shadow.
// Like gtk.Label it does not emit signals.
func NewArrow(at ArrowType, st ShadowType) *Arrow {
	return &Arrow{Widget{C.gtk_arrow_new(C.GtkArrowType(at), C.GtkShadowType(st))}}
}

// Change the arrows direction and shadow.
func (a *Arrow) Set(at ArrowType, st ShadowType) {
	C.gtk_arrow_set(ARROW(a), C.GtkArrowType(at), C.GtkShadowType(st))
}

//-----------------------------------------------------------------------
// GtkCalendar
//-----------------------------------------------------------------------

// gtk_calendar_new
// gtk_calendar_select_month
// gtk_calendar_select_day
// gtk_calendar_mark_day
// gtk_calendar_unmark_day
// gtk_calendar_clear_marks
// gtk_calendar_get_display_options
// gtk_calendar_set_display_options
// gtk_calendar_get_date
// gtk_calendar_set_detail_func
// gtk_calendar_get_detail_width_chars
// gtk_calendar_set_detail_width_chars
// gtk_calendar_get_detail_height_rows
// gtk_calendar_set_detail_height_rows
// gtk_calendar_display_options
// gtk_calendar_freeze
// gtk_calendar_thaw

//-----------------------------------------------------------------------
// GtkDrawingArea
//-----------------------------------------------------------------------
type DrawingArea struct {
	Widget
}

func NewDrawingArea() *DrawingArea {
	return &DrawingArea{Widget{C.gtk_drawing_area_new()}}
}

//Deprecated. Use GtkWidget.SetSizeRequest() instead.
func (v *DrawingArea) GetSizeRequest(width, height int) {
	deprecated_since(2, 0, 0, "gtk_drawing_area_size()")
	C.gtk_drawing_area_size(DRAWING_AREA(v), gint(width), gint(height))
}

//-----------------------------------------------------------------------
// GtkEventBox
//-----------------------------------------------------------------------
type EventBox struct {
	Bin
}

func NewEventBox() *EventBox {
	return &EventBox{Bin{Container{Widget{C.gtk_event_box_new()}}}}
}

// void gtk_event_box_set_above_child (GtkEventBox *event_box, gboolean above_child);
// gboolean gtk_event_box_get_above_child (GtkEventBox *event_box);
// void gtk_event_box_set_visible_window (GtkEventBox *event_box, gboolean visible_window);
// gboolean gtk_event_box_get_visible_window (GtkEventBox *event_box);

//-----------------------------------------------------------------------
// GtkHandleBox
//-----------------------------------------------------------------------

// gtk_handle_box_new
// gtk_handle_box_set_shadow_type
// gtk_handle_box_set_handle_position
// gtk_handle_box_set_snap_edge
// gtk_handle_box_get_handle_position
// gtk_handle_box_get_shadow_type
// gtk_handle_box_get_snap_edge
// gtk_handle_box_get_child_detached

//-----------------------------------------------------------------------
// GtkIMContextSimple
//-----------------------------------------------------------------------

// gtk_im_context_simple_new
// gtk_im_context_simple_add_table

//-----------------------------------------------------------------------
// GtkIMMulticontext
//-----------------------------------------------------------------------

// gtk_im_multicontext_new
// gtk_im_multicontext_append_menuitems
// gtk_im_multicontext_get_context_id
// gtk_im_multicontext_set_context_id

//-----------------------------------------------------------------------
// GtkSizeGroup
//-----------------------------------------------------------------------
type SizeGroupMode int

const (
	SIZE_GROUP_NONE SizeGroupMode = iota
	SIZE_GROUP_HORIZONTAL
	SIZE_GROUP_VERTICAL
	SIZE_GROUP_BOTH
)

type SizeGroup struct {
	GSizeGroup *C.GtkSizeGroup
}

func NewSizeGroup(mode SizeGroupMode) *SizeGroup {
	return &SizeGroup{C.gtk_size_group_new(C.GtkSizeGroupMode(mode))}
}

func (v *SizeGroup) SetMode(mode SizeGroupMode) {
	C.gtk_size_group_set_mode(v.GSizeGroup, C.GtkSizeGroupMode(mode))
}

func (v *SizeGroup) GetMode() SizeGroupMode {
	return SizeGroupMode(C.gtk_size_group_get_mode(v.GSizeGroup))
}

func (v *SizeGroup) SetIgnoreHidden(ignore_hidden bool) {
	C.gtk_size_group_set_ignore_hidden(v.GSizeGroup, gbool(ignore_hidden))
}

func (v *SizeGroup) GetIgnoreHidden() bool {
	return gobool(C.gtk_size_group_get_ignore_hidden(v.GSizeGroup))
}

func (v *SizeGroup) AddWidget(w IWidget) {
	C.gtk_size_group_add_widget(v.GSizeGroup, ToNative(w))
}

func (v *SizeGroup) RemoveWidget(w IWidget) {
	C.gtk_size_group_remove_widget(v.GSizeGroup, ToNative(w))
}

func (v *SizeGroup) GetWidgets() *glib.SList {
	return glib.SListFromNative(unsafe.Pointer(C.gtk_size_group_get_widgets(v.GSizeGroup)))
}

//-----------------------------------------------------------------------
// GtkTooltip
//-----------------------------------------------------------------------
type Tooltip struct {
	GTooltip *C.GtkTooltip
}

func (v *Tooltip) SetMarkup(markup string) {
	ptr := C.CString(markup)
	defer cfree(ptr)
	C.gtk_tooltip_set_markup(v.GTooltip, gstring(ptr))
}

func (v *Tooltip) SetText(text string) {
	ptr := C.CString(text)
	defer cfree(ptr)
	C.gtk_tooltip_set_text(v.GTooltip, gstring(ptr))
}

func (v *Tooltip) SetIcon(pixbuf *gdkpixbuf.Pixbuf) {
	C.gtk_tooltip_set_icon(v.GTooltip, (*C.GdkPixbuf)(unsafe.Pointer(pixbuf.GPixbuf)))
}

func (v *Tooltip) SetIconFromStock(stock_id string, size IconSize) {
	ptr := C.CString(stock_id)
	defer cfree(ptr)
	C.gtk_tooltip_set_icon_from_stock(v.GTooltip, gstring(ptr), C.GtkIconSize(size))
}

func (v *Tooltip) SetIconFromIconName(icon_name string, size IconSize) {
	ptr := C.CString(icon_name)
	defer cfree(ptr)
	C.gtk_tooltip_set_icon_from_icon_name(v.GTooltip, gstring(ptr), C.GtkIconSize(size))
}

// gtk_tooltip_set_icon_from_gicon //since 2.20
// gtk_tooltip_set_custom
// gtk_tooltip_trigger_tooltip_query
// gtk_tooltip_set_tip_area

func TooltipFromNative(l unsafe.Pointer) *Tooltip {
	return &Tooltip{(*C.GtkTooltip)(l)}
}

//-----------------------------------------------------------------------
// GtkViewport
//-----------------------------------------------------------------------
type Viewport struct {
	Bin
}

func NewViewport(ha, va *Adjustment) *Viewport {
	var had, vad *C.GtkAdjustment
	if ha != nil {
		had = ha.GAdjustment
	}
	if va != nil {
		vad = va.GAdjustment
	}
	return &Viewport{Bin{Container{Widget{C.gtk_viewport_new(had, vad)}}}}
}

func (v *Viewport) GetHAdjustment() *Adjustment {
	return &Adjustment{C.gtk_viewport_get_hadjustment(VIEWPORT(v))}
}

func (v *Viewport) GetVAdjustment() *Adjustment {
	return &Adjustment{C.gtk_viewport_get_vadjustment(VIEWPORT(v))}
}

func (v *Viewport) SetHAdjustment(ha *Adjustment) {
	C.gtk_viewport_set_hadjustment(VIEWPORT(v), ADJUSTMENT(ha))
}

func (v *Viewport) SetVAdjustment(va *Adjustment) {
	C.gtk_viewport_set_vadjustment(VIEWPORT(v), ADJUSTMENT(va))
}

func (v *Viewport) GetShadowType() ShadowType {
	return ShadowType(C.gtk_viewport_get_shadow_type(VIEWPORT(v)))
}

func (v *Viewport) SetShadowType(typ ShadowType) {
	C.gtk_viewport_set_shadow_type(VIEWPORT(v), C.GtkShadowType(typ))
}

func (v *Viewport) GetBinWindow() *Window {
	panic_if_version_older_auto(2, 20, 0)
	return &Window{Bin{Container{Widget{
		C.toGWidget(unsafe.Pointer(C._gtk_viewport_get_bin_window(VIEWPORT(v))))}}}}
}

func (v *Viewport) GetViewWindow() *Window {
	panic_if_version_older_auto(2, 22, 0)
	return &Window{Bin{Container{Widget{
		C.toGWidget(unsafe.Pointer(C._gtk_viewport_get_view_window(VIEWPORT(v))))}}}}
}

//-----------------------------------------------------------------------
// GtkAccessible
//-----------------------------------------------------------------------
type Accessible struct {
	glib.GObject
}

func (v *Accessible) ConnectWidgetDestroyed() {
	C.gtk_accessible_connect_widget_destroyed(ACCESSIBLE(v))
}

func (v *Accessible) SetWidget(w IWidget) {
	panic_if_version_older_auto(2, 22, 0)
	C._gtk_accessible_set_widget(ACCESSIBLE(v), ToNative(w))
}

func (v *Accessible) GetWidget() *Widget {
	panic_if_version_older_auto(2, 22, 0)
	return &Widget{C._gtk_accessible_get_widget(ACCESSIBLE(v))}
}

//-----------------------------------------------------------------------
// GtkBin
//-----------------------------------------------------------------------
type Bin struct {
	Container
}

func (v *Bin) GetChild() *Widget {
	return &Widget{C.gtk_bin_get_child(BIN(v))}
}

//-----------------------------------------------------------------------
// GtkBox
//-----------------------------------------------------------------------
type PackType int

const (
	PACK_START PackType = iota
	PACK_END
)

type IBox interface {
	IContainer
	PackStart(child IWidget, expand bool, fill bool, padding uint)
	PackEnd(child IWidget, expand bool, fill bool, padding uint)
}
type Box struct {
	Container
}

func (v *Box) PackStart(child IWidget, expand bool, fill bool, padding uint) {
	C.gtk_box_pack_start(BOX(v), ToNative(child), gbool(expand),
		gbool(fill), guint(padding))
}

func (v *Box) PackEnd(child IWidget, expand bool, fill bool, padding uint) {
	C.gtk_box_pack_end(BOX(v), ToNative(child), gbool(expand), gbool(fill), guint(padding))
}

//Deprecated since 2.14. Use PackStart() instead.
func (v *Box) PackStartDefaults(child IWidget) {
	deprecated_since(2, 14, 0, "gtk_box_pack_start_defaults()")
	C.gtk_box_pack_start_defaults(BOX(v), ToNative(child))
}

//Deprecated since 2.14. Use PackEnd() instead.
func (v *Box) PackEndDefaults(child IWidget) {
	deprecated_since(2, 14, 0, "gtk_box_pack_end_defaults()")
	C.gtk_box_pack_end_defaults(BOX(v), ToNative(child))
}

func (v *Box) GetHomogeneous() bool {
	return gobool(C.gtk_box_get_homogeneous(BOX(v)))
}

func (v *Box) SetHomogeneous(homogeneous bool) {
	C.gtk_box_set_homogeneous(BOX(v), gbool(homogeneous))
}

func (v *Box) GetSpacing() int {
	return int(C.gtk_box_get_spacing(BOX(v)))
}

func (v *Box) SetSpacing(spacing int) {
	C.gtk_box_set_spacing(BOX(v), gint(spacing))
}

func (v *Box) ReorderChild(child IWidget, position int) {
	C.gtk_box_reorder_child(BOX(v), ToNative(child), C.gint(position))
}

func (v *Box) QueryChildPacking(child IWidget, expand *bool, fill *bool, padding *uint, pack_type *PackType) {
	var cexpand, cfill C.gboolean
	var cpadding C.guint
	var cpack_type C.GtkPackType
	C.gtk_box_query_child_packing(BOX(v), ToNative(child), &cexpand, &cfill, &cpadding, &cpack_type)
	*expand = gobool(cexpand)
	*fill = gobool(cfill)
	*padding = uint(cpadding)
	*pack_type = PackType(cpack_type)
}

func (v *Box) SetChildPacking(child IWidget, expand, fill bool, padding uint, pack_type PackType) {
	C.gtk_box_set_child_packing(BOX(v), ToNative(child), gbool(expand), gbool(fill), guint(padding), C.GtkPackType(pack_type))
}

//-----------------------------------------------------------------------
// GtkButtonBox
//-----------------------------------------------------------------------

// gtk_button_box_get_spacing
// gtk_button_box_get_layout
// gtk_button_box_get_child_size
// gtk_button_box_get_child_ipadding
// gtk_button_box_get_child_secondary
// gtk_button_box_set_spacing
// gtk_button_box_set_layout
// gtk_button_box_set_child_size
// gtk_button_box_set_child_ipadding
// gtk_button_box_set_child_secondary

//-----------------------------------------------------------------------
// GtkContainer
//-----------------------------------------------------------------------
type IContainer interface {
	IWidget
	Add(w IWidget)
}
type Container struct {
	Widget
}

type ResizeMode int

const (
	RESIZE_PARENT    = ResizeMode(0)
	RESIZE_QUEUE     = ResizeMode(1)
	RESIZE_IMMEDIATE = ResizeMode(2)
)

func (v *Container) Add(w IWidget) {
	C.gtk_container_add(CONTAINER(v), ToNative(w))
}

func (v *Container) Remove(w IWidget) {
	C.gtk_container_remove(CONTAINER(v), ToNative(w))
}

// gtk_container_add_with_properties
func (v *Container) GetResizeMode() ResizeMode {
	return ResizeMode(C.gtk_container_get_resize_mode(CONTAINER(v)))
}

func (v *Container) SetResizeMode(resize_mode ResizeMode) {
	C.gtk_container_set_resize_mode(CONTAINER(v), C.GtkResizeMode(resize_mode))
}

func (v *Container) CheckResize() {
	C.gtk_container_check_resize(CONTAINER(v))
}

// gtk_container_foreach

func (v *Container) GetChildren() *glib.List {
	return glib.ListFromNative(unsafe.Pointer(C.gtk_container_get_children(CONTAINER(v))))
}

func (v *Container) SetReallocateRedraws(redraws bool) {
	C.gtk_container_set_reallocate_redraws(CONTAINER(v), gbool(redraws))
}

func (v *Container) GetFocusChild() *Widget {
	return &Widget{C.gtk_container_get_focus_child(CONTAINER(v))}
}

func (v *Container) SetFocusChild(child *Widget) {
	C.gtk_container_set_focus_child(CONTAINER(v), child.GWidget)
} 

func (v *Container) GetFocusVAdjustment() *Adjustment {
	return &Adjustment{C.gtk_container_get_focus_vadjustment(CONTAINER(v))}
}

func (v *Container) SetFocusVAdjustment(adjustment *Adjustment) {
	C.gtk_container_set_focus_vadjustment(CONTAINER(v), adjustment.GAdjustment)
}

func (v *Container) GetFocusHAdjustment() *Adjustment {
	return &Adjustment{C.gtk_container_get_focus_hadjustment(CONTAINER(v))}
}

func (v *Container) SetFocusHAdjustment(adjustment *Adjustment) {
	C.gtk_container_set_focus_hadjustment(CONTAINER(v), adjustment.GAdjustment)
}
// gtk_container_resize_children
// gtk_container_child_type
// gtk_container_child_get

func (v *Container) ChildSet(w IWidget, propName string, value interface{}) {

	ptr := C.CString(propName)
	defer cfree(ptr)

	switch value.(type) {
	case bool:
		C._gtk_container_child_set_bool(CONTAINER(v), ToNative(w), gstring(ptr), gbool(value.(bool)))
	case int:
		C._gtk_container_child_set_int(CONTAINER(v), ToNative(w), gstring(ptr), gint(value.(int)))
	}
}

// gtk_container_child_get_property
// gtk_container_child_set_property
// gtk_container_child_get_valist
// gtk_container_child_set_valist
// gtk_container_forall

func (v *Container) GetBorderWidth() uint {
	return uint(C.gtk_container_get_border_width(CONTAINER(v)))
}

func (v *Container) SetBorderWidth(border_width uint) {
	C.gtk_container_set_border_width(CONTAINER(v), guint(border_width))
}

// gtk_container_propagate_expose
// gtk_container_get_focus_chain
// gtk_container_set_focus_chain
// gtk_container_unset_focus_chain
// gtk_container_class_find_child_property
// gtk_container_class_install_child_property
// gtk_container_class_list_child_properties

//-----------------------------------------------------------------------
// GtkItem
//-----------------------------------------------------------------------
type Item struct {
	Bin
}

//Deprecated since 2.22. Use GtkMenuItem.Select() instead.
func (v *Item) Select() {
	deprecated_since(2, 22, 0, "gtk_item_select()")
	C.gtk_item_select(ITEM(v))
}

//Deprecated since 2.22. Use GtkMenuItem.Deselect() instead.
func (v *Item) Deselect() {
	deprecated_since(2, 22, 0, "gtk_item_deselect()")
	C.gtk_item_deselect(ITEM(v))
}

//Deprecated since 2.22.
func (v *Item) Toggle() {
	deprecated_since(2, 22, 0, "gtk_item_select()")
	C.gtk_item_toggle(ITEM(v))
}

//-----------------------------------------------------------------------
// GtkMenuShell
//-----------------------------------------------------------------------

// gtk_menu_shell_append
// gtk_menu_shell_prepend
// gtk_menu_shell_insert
// gtk_menu_shell_deactivate
// gtk_menu_shell_select_item
// gtk_menu_shell_select_first
// gtk_menu_shell_deselect
// gtk_menu_shell_activate_item
// gtk_menu_shell_cancel
// gtk_menu_shell_set_take_focus
// gtk_menu_shell_get_take_focus

//-----------------------------------------------------------------------
// GtkMisc
//-----------------------------------------------------------------------

type Misc struct {
	Widget
}

func (m *Misc) SetAlignment(xalign, yalign float64) {
	C.gtk_misc_set_alignment(MISC(m), C.gfloat(xalign), C.gfloat(yalign))
}

func (m *Misc) SetPadding(xpad, ypad int) {
	C.gtk_misc_set_padding(MISC(m), C.gint(xpad), C.gint(ypad))
}

func (m *Misc) GetAlignment(xalign, yalign *float64) {
	var gxalign, gyalign *C.gfloat
	C.gtk_misc_get_alignment(MISC(m), gxalign, gyalign)
	*xalign = float64(*gxalign)
	*yalign = float64(*gyalign)
}

func (m *Misc) GetPadding(xpad, ypad *int) {
	var gxpad, gypad *C.gint
	C.gtk_misc_get_padding(MISC(m), gxpad, gypad)
	*xpad = int(*gxpad)
	*ypad = int(*gypad)
}

//-----------------------------------------------------------------------
// GtkObject
//-----------------------------------------------------------------------
type Object struct {
	glib.GObject
}

func (v *Object) Ref() {
	C.g_object_ref(C.gpointer(v.GObject.Object))
}

func (v *Object) Unref() {
	C.g_object_unref(C.gpointer(v.GObject.Object))
}

//deprecated since 2.20

//-----------------------------------------------------------------------
// GtkPaned
//-----------------------------------------------------------------------
type IPaned interface {
	IContainer
	Add1(child IWidget)
	Add2(child IWidget)
	Pack1(child IWidget, expand bool, fill bool)
	Pack2(child IWidget, expand bool, fill bool)
}
type Paned struct {
	Container
}

func (v *Paned) Add1(child IWidget) {
	C.gtk_paned_add1(PANED(v), ToNative(child))
}

func (v *Paned) Add2(child IWidget) {
	C.gtk_paned_add2(PANED(v), ToNative(child))
}

func (v *Paned) Pack1(child IWidget, resize bool, shrink bool) {
	C.gtk_paned_pack1(PANED(v), ToNative(child), gbool(resize), gbool(shrink))
}

func (v *Paned) Pack2(child IWidget, resize bool, shrink bool) {
	C.gtk_paned_pack2(PANED(v), ToNative(child), gbool(resize), gbool(shrink))
}

func (v *Paned) GetChild1() *Widget {
	return &Widget{C.gtk_paned_get_child1(PANED(v))}
}

func (v *Paned) GetChild2() *Widget {
	return &Widget{C.gtk_paned_get_child2(PANED(v))}
}

func (v *Paned) SetPosition(position int) {
	C.gtk_paned_set_position(PANED(v), gint(position))
}

func (v *Paned) GetPosition() int {
	return int(C.gtk_paned_get_position(PANED(v)))
}

// gtk_paned_get_handle_window //since 2.20

//-----------------------------------------------------------------------
// GtkRange
//-----------------------------------------------------------------------
type Range struct {
	Widget
}

func (v *Range) GetFillLevel() float64 {
	r := C.gtk_range_get_fill_level(RANGE(v))
	return float64(r)
}

func (v *Range) GetRestrictToFillLevel() bool {
	return gobool(C.gtk_range_get_restrict_to_fill_level(RANGE(v)))
}

func (v *Range) GetShowFillLevel() bool {
	return gobool(C.gtk_range_get_show_fill_level(RANGE(v)))
}

func (v *Range) SetFillLevel(value float64) {
	C.gtk_range_set_fill_level(RANGE(v), gdouble(value))
}

func (v *Range) SetRestrictToFillLevel(b bool) {
	C.gtk_range_set_restrict_to_fill_level(RANGE(v), gbool(b))
}

func (v *Range) SetShowFillLevel(b bool) {
	C.gtk_range_set_show_fill_level(RANGE(v), gbool(b))
}

func (v *Range) GetAdjustment() *Adjustment {
	return &Adjustment{C.gtk_range_get_adjustment(RANGE(v))}
}

// void gtk_range_set_update_policy (GtkRange *range, GtkUpdateType policy); //deprecated in 2.24

func (v *Range) SetAdjustment(a *Adjustment) {
	C.gtk_range_set_adjustment(RANGE(v), a.GAdjustment)
}

func (v *Range) GetInverted() bool {
	return gobool(C.gtk_range_get_inverted(RANGE(v)))
}

func (v *Range) SetInverted(b bool) {
	C.gtk_range_set_inverted(RANGE(v), gbool(b))
}

// GtkUpdateType gtk_range_get_update_policy (GtkRange *range); //deprecated since 2.24

func (v *Range) SetIncrements(step, page float64) {
	C.gtk_range_set_increments(RANGE(v), gdouble(step), gdouble(page))
}

func (v *Range) SetRange(min, max float64) {
	C.gtk_range_set_range(RANGE(v), gdouble(min), gdouble(max))
}

func (v *Range) GetValue() float64 {
	return float64(C.gtk_range_get_value(RANGE(v))) //TODO test
	//var r C.gdouble
	//C._gtk_range_get_value(RANGE(v), &r)
	//return float64(r)
}

func (v *Range) SetValue(value float64) {
	C.gtk_range_set_value(RANGE(v), gdouble(value))
}

// gtk_range_get_round_digits //since 2.24
// gtk_range_set_round_digits //since 2.24
// void gtk_range_set_lower_stepper_sensitivity (GtkRange *range, GtkSensitivityType sensitivity);
// GtkSensitivityType gtk_range_get_lower_stepper_sensitivity (GtkRange *range);
// void gtk_range_set_upper_stepper_sensitivity (GtkRange *range, GtkSensitivityType sensitivity);
// GtkSensitivityType gtk_range_get_upper_stepper_sensitivity (GtkRange *range);

func (v *Range) GetFlippable() bool {
	panic_if_version_older(2, 18, 0, "gtk_range_get_flippable()")
	return gobool(C._gtk_range_get_flippable(RANGE(v)))
}

func (v *Range) SetFlippable(b bool) {
	panic_if_version_older(2, 18, 0, "gtk_range_set_flippable()")
	C._gtk_range_set_flippable(RANGE(v), gbool(b))
}

// gtk_range_get_min_slider_size //since 2.20
// gtk_range_get_range_rect //since 2.20
// gtk_range_get_slider_range //since 2.20
// gtk_range_get_slider_size_fixed //since 2.20
// gtk_range_set_min_slider_size //since 2.20
// gtk_range_set_slider_size_fixed //since 2.20

//-----------------------------------------------------------------------
// GtkScale
//-----------------------------------------------------------------------
type PositionType int

const (
	POS_LEFT PositionType = iota
	POS_RIGHT
	POS_TOP
	POS_BOTTOM
)

type Scale struct {
	Range
}

func (v *Scale) SetDigits(digits int) {
	C.gtk_scale_set_digits(SCALE(v), gint(digits))
}

func (v *Scale) SetDrawValue(draw_value bool) {
	C.gtk_scale_set_draw_value(SCALE(v), gbool(draw_value))
}

func (v *Scale) SetValuePos(pos PositionType) {
	C.gtk_scale_set_value_pos(SCALE(v), C.GtkPositionType(pos))
}

func (v *Scale) GetDigits() int {
	return int(C.gtk_scale_get_digits(SCALE(v)))
}

func (v *Scale) GetDrawValue() bool {
	return gobool(C.gtk_scale_get_draw_value(SCALE(v)))
}

func (v *Scale) GetValuePos() PositionType {
	return PositionType(C.gtk_scale_get_value_pos(SCALE(v)))
}

// PangoLayout * gtk_scale_get_layout (GtkScale *scale);

func (v *Scale) GetLayoutOffsets(x *int, y *int) {
	var xx, yy C.gint
	C.gtk_scale_get_layout_offsets(SCALE(v), &xx, &yy)
	*x = int(xx)
	*y = int(yy)
}

func (v *Scale) AddMark(value float64, position PositionType, markup string) {
	ptr := C.CString(markup)
	defer cfree(ptr)
	C.gtk_scale_add_mark(SCALE(v), gdouble(value), C.GtkPositionType(position), gstring(ptr))
}

func (v *Scale) ClearMarks() {
	C.gtk_scale_clear_marks(SCALE(v))
}

//-----------------------------------------------------------------------
// GtkSeparator
//-----------------------------------------------------------------------
type Separator struct {
	Widget
}

//-----------------------------------------------------------------------
// GtkWidget
//-----------------------------------------------------------------------
type Allocation gdk.Rectangle

type AccelFlags int

const (
	ACCEL_VISIBLE AccelFlags = 1 << iota
	ACCEL_LOCKED
	ACCEL_MASK = 0x07
)

type StateType int

const (
	STATE_NORMAL StateType = iota
	STATE_ACTIVE
	STATE_PRELIGHT
	STATE_SELECTED
	STATE_INSENSITIVE
)

type IWidget interface {
	ToNative() *C.GtkWidget
	Hide()
	HideAll()
	Show()
	ShowAll()
	ShowNow()
	Destroy()
	Connect(s string, f interface{}, data ...interface{}) int
	GetTopLevel() *Widget
	GetTopLevelAsWindow() *Window
	HideOnDelete()
	QueueResize()
	GetName() string
}

func ToNative(w IWidget) *C.GtkWidget {
	if w == nil {
		return nil
	}
	if reflect.ValueOf(w).IsNil() {
		return nil
	}
	return w.ToNative()
}

type Widget struct {
	GWidget *C.GtkWidget
}

func WidgetFromNative(p unsafe.Pointer) *Widget {
	return &Widget{C.toGWidget(p)}
}

//TODO GtkWidget will have GObject as anonymous field.
func WidgetFromObject(object *glib.GObject) *Widget {
	return &Widget{C.toGWidget(object.Object)}
}

func (v *Widget) ToNative() *C.GtkWidget {
	return v.GWidget
}

func (v *Widget) Connect(s string, f interface{}, datas ...interface{}) int {
	return glib.ObjectFromNative(unsafe.Pointer(v.GWidget)).Connect(s, f, datas...)
}

func (v *Widget) StopEmission(s string) {
	glib.ObjectFromNative(unsafe.Pointer(v.GWidget)).StopEmission(s)
}

func (v *Widget) Emit(s string) {
	glib.ObjectFromNative(unsafe.Pointer(v.GWidget)).Emit(s)
}

func (v *Widget) HandlerBlock(id int) {
	glib.ObjectFromNative(unsafe.Pointer(v.GWidget)).HandlerBlock(id)
}

func (v *Widget) HandlerUnblock(id int) {
	glib.ObjectFromNative(unsafe.Pointer(v.GWidget)).HandlerUnblock(id)
}

func (v *Widget) HandlerDisconnect(id int) {
	glib.ObjectFromNative(unsafe.Pointer(v.GWidget)).HandlerDisconnect(id)
}

// gtk_widget_new

func (v *Widget) Ref() {
	C.g_object_ref(C.gpointer(v.GWidget))
}

func (v *Widget) Unref() {
	C.g_object_unref(C.gpointer(v.GWidget))
}

func (v *Widget) Destroy() {
	C.gtk_widget_destroy(v.GWidget)
}

// gtk_widget_destroyed

func (v *Widget) Unparent() {
	C.gtk_widget_unparent(v.GWidget)
}

func (v *Widget) Show() {
	C.gtk_widget_show(v.GWidget)
}

func (v *Widget) ShowNow() {
	C.gtk_widget_show_now(v.GWidget)
}

func (v *Widget) Hide() {
	C.gtk_widget_hide(v.GWidget)
}

func (v *Widget) ShowAll() {
	C.gtk_widget_show_all(v.GWidget)
}

//Deprecated since 2.24. Use Hide() instead.
func (v *Widget) HideAll() {
	deprecated_since(2, 24, 0, "gtk_widget_hide_all()")
	C.gtk_widget_hide_all(v.GWidget)
}

func (v *Widget) Map() {
	C.gtk_widget_map(v.GWidget)
}

func (v *Widget) Unmap() {
	C.gtk_widget_unmap(v.GWidget)
}

func (v *Widget) Realize() {
	C.gtk_widget_realize(v.GWidget)
}

func (v *Widget) Unrealize() {
	C.gtk_widget_unrealize(v.GWidget)
}

func (v *Widget) QueueDraw() {
	C.gtk_widget_queue_draw(v.GWidget)
}

func (v *Widget) QueueResize() {
	C.gtk_widget_queue_resize(v.GWidget)
}

func (v *Widget) QueueResizeNoRedraw() {
	C.gtk_widget_queue_resize_no_redraw(v.GWidget)
}

// gtk_widget_size_request
// gtk_widget_get_child_requisition
// gtk_widget_size_allocate

func (v *Widget) AddAccelerator(signal string, group *AccelGroup, key uint, mods gdk.ModifierType, flags AccelFlags) {
	csignal := C.CString(signal)
	defer cfree(csignal)
	C.gtk_widget_add_accelerator(v.GWidget, gstring(csignal), group.GAccelGroup, guint(key),
		C.GdkModifierType(mods), C.GtkAccelFlags(flags))
}

// gtk_widget_remove_accelerator
// gtk_widget_set_accel_path
// gtk_widget_list_accel_closures

func (v *Widget) CanActivateAccel(signal_id uint) bool {
	return gobool(C.gtk_widget_can_activate_accel(v.GWidget, guint(signal_id)))
}

// gtk_widget_event

func (v *Widget) Activate() {
	C.gtk_widget_activate(v.GWidget)
}

func (v *Widget) Reparent(parent IWidget) {
	C.gtk_widget_reparent(v.GWidget, ToNative(parent))
}

// gtk_widget_intersect

func (v *Widget) IsFocus() bool {
	return gobool(C.gtk_widget_is_focus(v.GWidget))
}

func (v *Widget) GrabFocus() {
	C.gtk_widget_grab_focus(v.GWidget)
}

func (v *Widget) GrabDefault() {
	C.gtk_widget_grab_default(v.GWidget)
}

func (v *Widget) SetName(name string) {
	ptr := C.CString(name)
	defer cfree(ptr)
	C.gtk_widget_set_name(v.GWidget, gstring(ptr))
}

func (v *Widget) GetName() string {
	return gostring(C.gtk_widget_get_name(v.GWidget))
}

func (v *Widget) SetState(state StateType) {
	C.gtk_widget_set_state(v.GWidget, C.GtkStateType(state))
}

func (v *Widget) SetSensitive(setting bool) {
	C.gtk_widget_set_sensitive(v.GWidget, gbool(setting))
}

func (v *Widget) SetParent(parent IWidget) {
	C.gtk_widget_set_parent(v.GWidget, ToNative(parent))
}

func (v *Widget) SetParentWindow(parent *gdk.Window) {
	C.gtk_widget_set_parent_window(v.GWidget, C.toGdkWindow(unsafe.Pointer(parent.GWindow)))
}

func (v *Widget) GetParentWindow() *gdk.Window {
	return gdk.WindowFromUnsafe(unsafe.Pointer(C.gtk_widget_get_parent_window(v.GWidget)))
}

//Deprecated since 2.2, use SetSizeRequest() instead
func (v *Widget) SetUSize(width int, height int) {
	deprecated_since(2, 2, 0, "gtk_widget_set_usize()")
	C.gtk_widget_set_usize(v.GWidget, gint(width), gint(height))
}

func (v *Widget) SetEvents(events int) {
	C.gtk_widget_set_events(v.GWidget, gint(events))
}

func (v *Widget) AddEvents(events int) {
	C.gtk_widget_add_events(v.GWidget, gint(events))
}

// gtk_widget_set_extension_events
// gtk_widget_get_extension_events

func (v *Widget) GetTopLevel() *Widget {
	return &Widget{C.gtk_widget_get_toplevel(v.GWidget)}
}

// gtk_widget_get_ancestor

func (v *Widget) GetColormap() *gdk.Colormap {
	return gdk.ColormapFromUnsafe(unsafe.Pointer(C.gtk_widget_get_colormap(v.GWidget)))
}

// gtk_widget_get_colormap
// gtk_widget_set_colormap
// gtk_widget_get_visual
// gtk_widget_get_events
// gtk_widget_get_pointer
// gtk_widget_is_ancestor
// gtk_widget_translate_coordinates

func (v *Widget) HideOnDelete() {
	C._gtk_widget_hide_on_delete(v.GWidget)
}

// gtk_widget_set_style
// gtk_widget_ensure_style

func (v *Widget) GetStyle() *Style {
	return &Style{C.gtk_rc_get_style(v.GWidget)}
}

// gtk_widget_reset_rc_styles
// gtk_widget_push_colormap
// gtk_widget_pop_colormap
// gtk_widget_set_default_colormap
// gtk_widget_get_default_style
// gtk_widget_get_default_colormap
// gtk_widget_get_default_visual
// gtk_widget_set_direction
// gtk_widget_get_direction
// gtk_widget_set_default_direction
// gtk_widget_get_default_direction
// gtk_widget_shape_combine_mask
// gtk_widget_input_shape_combine_mask
// gtk_widget_path
// gtk_widget_class_path
// gtk_widget_get_composite_name
// gtk_widget_modify_style
// gtk_widget_get_modifier_style
// gtk_widget_modify_text
// gtk_widget_modify_base
// gtk_widget_modify_cursor
// gtk_widget_create_pango_context
// gtk_widget_get_pango_context
func (v *Widget) GetPangoContext() *pango.Context {
	return pango.ContextFromUnsafe(unsafe.Pointer(C.gtk_widget_get_pango_context((*C.GtkWidget)(v.GWidget))))
}

// gtk_widget_create_pango_layout

func (v *Widget) RenderIcon(stock_id string, size IconSize, detail string) *gdkpixbuf.Pixbuf {
	pstock_id := C.CString(stock_id)
	defer cfree(pstock_id)
	pdetail := C.CString(detail)
	defer cfree(pdetail)
	gpixbuf := C.gtk_widget_render_icon(v.GWidget, gstring(pstock_id), C.GtkIconSize(size), gstring(pdetail))
	return &gdkpixbuf.Pixbuf{
		GdkPixbuf: gdkpixbuf.NewGdkPixbuf(unsafe.Pointer(gpixbuf)),
		GObject:   glib.ObjectFromNative(unsafe.Pointer(gpixbuf)),
	}
}

// gtk_widget_pop_composite_child
// gtk_widget_push_composite_child

//Deprecated since 2.2. Use QueueDraw() instead.
func (v *Widget) QueueClear() {
	deprecated_since(2, 2, 0, "gtk_widget_queue_clear()")
	C.gtk_widget_queue_clear(v.GWidget)
}

// gtk_widget_queue_draw_area

func (v *Widget) SetAppPaintable(setting bool) {
	C.gtk_widget_set_app_paintable(v.GWidget, gbool(setting))
}

func (v *Widget) SetDoubleBuffered(setting bool) {
	C.gtk_widget_set_double_buffered(v.GWidget, gbool(setting))
}

func (v *Widget) SetRedrawOnAllocate(setting bool) {
	C.gtk_widget_set_redraw_on_allocate(v.GWidget, gbool(setting))
}

// gtk_widget_set_composite_name
// gtk_widget_set_scroll_adjustments

func (v *Widget) MnemonicActivate(group_cycling bool) bool {
	return gobool(C.gtk_widget_mnemonic_activate(v.GWidget, gbool(group_cycling)))
}

// gtk_widget_class_install_style_property
// gtk_widget_class_install_style_property_parser
// gtk_widget_class_find_style_property
// gtk_widget_class_list_style_properties
// gtk_widget_region_intersect
// gtk_widget_send_expose
// gtk_widget_send_focus_change //since 2.22
// gtk_widget_style_get
// gtk_widget_style_get_property
// gtk_widget_style_get_valist
// gtk_widget_style_attach //since 2.20
// gtk_widget_get_accessible
// gtk_widget_child_focus
// gtk_widget_child_notify
// gtk_widget_freeze_child_notify

func (v *Widget) GetChildVisible() bool {
	return gobool(C.gtk_widget_get_child_visible(v.GWidget))
}

func (v *Widget) GetParent() *Widget {
	return &Widget{C.gtk_widget_get_parent(v.GWidget)}
}

func (v *Widget) GetSettings() *Settings {
	return &Settings{C.gtk_widget_get_settings(v.GWidget)}
}

// gtk_widget_get_clipboard
// gtk_widget_get_display
// gtk_widget_get_root_window
// gtk_widget_get_screen
// gtk_widget_has_screen

//TODO go can have multiple return, adapt the function!
func (v *Widget) GetSizeRequest(width *int, height *int) {
	var w, h C.gint
	C.gtk_widget_get_size_request(v.GWidget, &w, &h)
	*width = int(w)
	*height = int(h)
}

func (v *Widget) SetChildVisible(setting bool) {
	C.gtk_widget_set_child_visible(v.GWidget, gbool(setting))
}

func (v *Widget) SetSizeRequest(width int, height int) {
	C.gtk_widget_set_size_request(v.GWidget, gint(width), gint(height))
}

func (v *Widget) SetNoShowAll(setting bool) {
	C.gtk_widget_set_no_show_all(v.GWidget, gbool(setting))
}

func (v *Widget) GetNoShowAll() bool {
	return gobool(C.gtk_widget_get_no_show_all(v.GWidget))
}

// gtk_widget_list_mnemonic_labels
// gtk_widget_add_mnemonic_label
// gtk_widget_remove_mnemonic_label

func (v *Widget) IsComposited() bool {
	return gobool(C.gtk_widget_is_composited(v.GWidget))
}

// gtk_widget_error_bell
// gtk_widget_keynav_failed

func (v *Widget) GetTooltipMarkup() string {
	return gostring(C.gtk_widget_get_tooltip_markup(v.GWidget))
}

func (v *Widget) SetTooltipMarkup(markup string) {
	ptr := C.CString(markup)
	defer cfree(ptr)
	C.gtk_widget_set_tooltip_markup(v.GWidget, gstring(ptr))
}

func (v *Widget) GetTooltipText() string {
	return gostring(C.gtk_widget_get_tooltip_text(v.GWidget))
}

func (v *Widget) SetTooltipText(text string) {
	ptr := C.CString(text)
	defer cfree(ptr)
	C.gtk_widget_set_tooltip_text(v.GWidget, gstring(ptr))
}

func (v *Widget) GetTooltipWindow() *Window {
	return &Window{Bin{Container{Widget{
		C.toGWidget(unsafe.Pointer(C.gtk_widget_get_tooltip_window(v.GWidget)))}}}}
}

func (v *Widget) SetTooltipWindow(w *Window) {
	C.gtk_widget_set_tooltip_window(v.GWidget, C.toGWindow(ToNative(w)))
}

func (v *Widget) GetHasTooltip() bool {
	return gobool(C.gtk_widget_get_has_tooltip(v.GWidget))
}

func (v *Widget) SetHasTooltip(setting bool) {
	C.gtk_widget_set_has_tooltip(v.GWidget, gbool(setting))
}

// gtk_widget_trigger_tooltip_query
// gtk_widget_get_snapshot

func (v *Widget) GetWindow() *gdk.Window {
	return gdk.WindowFromUnsafe(unsafe.Pointer(C.gtk_widget_get_window(v.GWidget)))
}

func (v *Widget) GetAllocation() *Allocation {
	panic_if_version_older(2, 18, 0, "gtk_widget_get_allocation()")
	var _allocation C.GtkAllocation
	C._gtk_widget_get_allocation(v.GWidget, &_allocation)
	return &Allocation{X: int(_allocation.x), Y: int(_allocation.y), Width: int(_allocation.width), Height: int(_allocation.height)}
}

func (v *Widget) SetAllocation(allocation *Allocation) {
	panic_if_version_older(2, 18, 0, "gtk_widget_set_allocation()")
	var _allocation C.GtkAllocation
	_allocation.x = gint(allocation.X)
	_allocation.y = gint(allocation.Y)
	_allocation.width = gint(allocation.Width)
	_allocation.height = gint(allocation.Height)
	C._gtk_widget_set_allocation(v.GWidget, &_allocation)
}

func (v *Widget) GetAppPaintable() bool {
	panic_if_version_older(2, 18, 0, "gtk_widget_get_app_paintable()")
	return gobool(C._gtk_widget_get_app_paintable(v.GWidget))
}

func (v *Widget) GetCanDefault() bool {
	panic_if_version_older(2, 18, 0, "gtk_widget_get_can_default()")
	return gobool(C._gtk_widget_get_can_default(v.GWidget))
}

func (v *Widget) SetCanDefault(setting bool) {
	panic_if_version_older(2, 18, 0, "gtk_widget_set_can_default()")
	C._gtk_widget_set_can_default(v.GWidget, gbool(setting))
}

func (v *Widget) GetCanFocus() bool {
	panic_if_version_older(2, 18, 0, "gtk_widget_get_can_focus()")
	return gobool(C._gtk_widget_get_can_focus(v.GWidget))
}

func (v *Widget) SetCanFocus(setting bool) {
	panic_if_version_older(2, 18, 0, "gtk_widget_set_can_focus()")
	C._gtk_widget_set_can_focus(v.GWidget, gbool(setting))
}

func (v *Widget) GetDoubleBuffered() bool {
	panic_if_version_older(2, 18, 0, "gtk_widget_get_double_buffered()")
	return gobool(C._gtk_widget_get_double_buffered(v.GWidget))
}

func (v *Widget) GetHasWindow() bool {
	panic_if_version_older(2, 18, 0, "gtk_widget_get_has_window()")
	return gobool(C._gtk_widget_get_has_window(v.GWidget))
}

func (v *Widget) SetHasWindow(setting bool) {
	panic_if_version_older(2, 18, 0, "gtk_widget_set_has_window()")
	C._gtk_widget_set_has_window(v.GWidget, gbool(setting))
}

func (v *Widget) GetSensitive() bool {
	panic_if_version_older(2, 18, 0, "gtk_widget_get_sensitive()")
	return gobool(C._gtk_widget_get_sensitive(v.GWidget))
}

func (v *Widget) IsSensitive() bool {
	panic_if_version_older(2, 18, 0, "gtk_widget_is_sensitive()")
	return gobool(C._gtk_widget_is_sensitive(v.GWidget))
}

func (v *Widget) GetState() StateType {
	panic_if_version_older(2, 18, 0, "gtk_widget_get_state()")
	return StateType(C._gtk_widget_get_state(v.GWidget))
}

func (v *Widget) GetVisible() bool {
	panic_if_version_older(2, 18, 0, "gtk_widget_get_visible()")
	return gobool(C._gtk_widget_get_visible(v.GWidget))
}

func (v *Widget) SetVisible(setting bool) {
	panic_if_version_older(2, 18, 0, "gtk_widget_set_visible()")
	C._gtk_widget_set_visible(v.GWidget, gbool(setting))
}

func (v *Widget) HasDefault() bool {
	panic_if_version_older(2, 18, 0, "gtk_widget_has_default()")
	return gobool(C._gtk_widget_has_default(v.GWidget))
}

func (v *Widget) HasFocus() bool {
	panic_if_version_older(2, 18, 0, "gtk_widget_has_focus()")
	return gobool(C._gtk_widget_has_focus(v.GWidget))
}

func (v *Widget) HasGrab() bool {
	panic_if_version_older(2, 18, 0, "gtk_widget_has_grab()")
	return gobool(C._gtk_widget_has_grab(v.GWidget))
}

// gtk_widget_has_rc_style //since 2.20

func (v *Widget) IsDrawable() bool {
	panic_if_version_older(2, 18, 0, "gtk_widget_is_drawable()")
	return gobool(C._gtk_widget_is_drawable(v.GWidget))
}

func (v *Widget) IsToplevel() bool {
	panic_if_version_older(2, 18, 0, "gtk_widget_is_toplevel()")
	return gobool(C._gtk_widget_is_toplevel(v.GWidget))
}

func (v *Widget) SetWindow(window *gdk.Window) {
	panic_if_version_older(2, 18, 0, "gtk_widget_set_window()")
	C._gtk_widget_set_window(v.GWidget, C.toGdkWindow(unsafe.Pointer(window.GWindow)))
}

func (v *Widget) SetReceivesDefault(setting bool) {
	panic_if_version_older(2, 18, 0, "gtk_widget_set_receives_default()")
	C._gtk_widget_set_receives_default(v.GWidget, gbool(setting))
}

func (v *Widget) GetReceivesDefault() bool {
	panic_if_version_older(2, 18, 0, "gtk_widget_get_receives_default()")
	return gobool(C._gtk_widget_get_receives_default(v.GWidget))
}

// gtk_widget_set_realized //since 2.20
// gtk_widget_get_realized //since 2.20
// gtk_widget_set_mapped //since 2.20
// gtk_widget_get_mapped //since 2.20
// gtk_widget_get_requisition //since 2.20
// gtk_requisition_copy
// gtk_requisition_free

//Convenience functions

func (v *Widget) GetTopLevelAsWindow() *Window {
	return &Window{Bin{Container{Widget{
		C.gtk_widget_get_toplevel(v.GWidget)}}}}
}

func (v *Widget) ModifyFontEasy(desc string) {
	pdesc := C.CString(desc)
	defer cfree(pdesc)
	C.gtk_widget_modify_font(v.GWidget, C.pango_font_description_from_string(pdesc))
}

func (v *Widget) ModifyFG(state StateType, color *gdk.Color) {
	C.gtk_widget_modify_fg(v.GWidget, C.GtkStateType(state), (*C.GdkColor)(unsafe.Pointer(&color.GColor)))
}

func (v *Widget) ModifyBG(state StateType, color *gdk.Color) {
	C.gtk_widget_modify_bg(v.GWidget, C.GtkStateType(state), (*C.GdkColor)(unsafe.Pointer(&color.GColor)))
}

func (v *Widget) ModifyText(state StateType, color *gdk.Color) {
	C.gtk_widget_modify_text(v.GWidget, C.GtkStateType(state), (*C.GdkColor)(unsafe.Pointer(&color.GColor)))
}

func (v *Widget) ModifyBase(state StateType, color *gdk.Color) {
	C.gtk_widget_modify_base(v.GWidget, C.GtkStateType(state), (*C.GdkColor)(unsafe.Pointer(&color.GColor)))
}

//-----------------------------------------------------------------------
// GtkIMContext
//-----------------------------------------------------------------------

// gtk_im_context_set_client_window
// gtk_im_context_get_preedit_string
// gtk_im_context_filter_keypress
// gtk_im_context_focus_in
// gtk_im_context_focus_out
// gtk_im_context_reset
// gtk_im_context_set_cursor_location
// gtk_im_context_set_use_preedit
// gtk_im_context_set_surrounding
// gtk_im_context_get_surrounding
// gtk_im_context_delete_surrounding

//-----------------------------------------------------------------------
// GtkPlug
//-----------------------------------------------------------------------

// gtk_plug_construct
// gtk_plug_construct_for_display
// gtk_plug_new
// gtk_plug_new_for_display
// gtk_plug_get_id
// gtk_plug_get_embedded
// gtk_plug_get_socket_window

//-----------------------------------------------------------------------
// GtkSocket
//-----------------------------------------------------------------------

// gtk_socket_new
// gtk_socket_steal
// gtk_socket_add_id
// gtk_socket_get_id
// gtk_socket_get_plug_window

//-----------------------------------------------------------------------
// GtkRecentManager
//-----------------------------------------------------------------------

// gtk_recent_manager_new
// gtk_recent_manager_get_default
// gtk_recent_manager_get_for_screen
// gtk_recent_manager_set_screen
// gtk_recent_manager_add_item
// gtk_recent_manager_add_full
// gtk_recent_manager_remove_item
// gtk_recent_manager_lookup_item
// gtk_recent_manager_has_item
// gtk_recent_manager_move_item
// gtk_recent_manager_get_limit
// gtk_recent_manager_set_limit
// gtk_recent_manager_get_items
// gtk_recent_manager_purge_items
// gtk_recent_info_ref
// gtk_recent_info_unref
// gtk_recent_info_get_uri
// gtk_recent_info_get_display_name
// gtk_recent_info_get_description
// gtk_recent_info_get_mime_type
// gtk_recent_info_get_added
// gtk_recent_info_get_modified
// gtk_recent_info_get_visited
// gtk_recent_info_get_private_hint
// gtk_recent_info_get_application_info
// gtk_recent_info_get_applications
// gtk_recent_info_last_application
// gtk_recent_info_get_groups
// gtk_recent_info_has_group
// gtk_recent_info_has_application
// gtk_recent_info_get_icon
// gtk_recent_info_get_short_name
// gtk_recent_info_get_uri_display
// gtk_recent_info_get_age
// gtk_recent_info_is_local
// gtk_recent_info_exists
// gtk_recent_info_match

//-----------------------------------------------------------------------
// GtkRecentChooser
//-----------------------------------------------------------------------

// gtk_recent_chooser_set_show_private
// gtk_recent_chooser_get_show_private
// gtk_recent_chooser_set_show_not_found
// gtk_recent_chooser_get_show_not_found
// gtk_recent_chooser_set_show_icons
// gtk_recent_chooser_get_show_icons
// gtk_recent_chooser_set_select_multiple
// gtk_recent_chooser_get_select_multiple
// gtk_recent_chooser_set_local_only
// gtk_recent_chooser_get_local_only
// gtk_recent_chooser_set_limit
// gtk_recent_chooser_get_limit
// gtk_recent_chooser_set_show_tips
// gtk_recent_chooser_get_show_tips
// gtk_recent_chooser_set_show_numbers
// gtk_recent_chooser_get_show_numbers
// gtk_recent_chooser_set_sort_type
// gtk_recent_chooser_get_sort_type
// gtk_recent_chooser_set_sort_func
// gtk_recent_chooser_set_current_uri
// gtk_recent_chooser_get_current_uri
// gtk_recent_chooser_get_current_item
// gtk_recent_chooser_select_uri
// gtk_recent_chooser_unselect_uri
// gtk_recent_chooser_select_all
// gtk_recent_chooser_unselect_all
// gtk_recent_chooser_get_items
// gtk_recent_chooser_get_uris
// gtk_recent_chooser_add_filter
// gtk_recent_chooser_remove_filter
// gtk_recent_chooser_list_filters
// gtk_recent_chooser_set_filter
// gtk_recent_chooser_get_filter

//-----------------------------------------------------------------------
// GtkRecentChooserDialog
//-----------------------------------------------------------------------

// gtk_recent_chooser_dialog_new
// gtk_recent_chooser_dialog_new_for_manager

//-----------------------------------------------------------------------
// GtkRecentChooserMenu
//-----------------------------------------------------------------------

// gtk_recent_chooser_menu_new
// gtk_recent_chooser_menu_new_for_manager
// gtk_recent_chooser_menu_get_show_numbers
// gtk_recent_chooser_menu_set_show_numbers

//-----------------------------------------------------------------------
// GtkRecentChooserWidget
//-----------------------------------------------------------------------

// gtk_recent_chooser_widget_new
// gtk_recent_chooser_widget_new_for_manager

//-----------------------------------------------------------------------
// GtkRecentFilter
//-----------------------------------------------------------------------

// gtk_recent_filter_new
// gtk_recent_filter_get_name
// gtk_recent_filter_set_name
// gtk_recent_filter_add_mime_type
// gtk_recent_filter_add_pattern
// gtk_recent_filter_add_pixbuf_formats
// gtk_recent_filter_add_application
// gtk_recent_filter_add_group
// gtk_recent_filter_add_age
// gtk_recent_filter_add_custom
// gtk_recent_filter_get_needed
// gtk_recent_filter_filter

//-----------------------------------------------------------------------
// GtkBuildable
//-----------------------------------------------------------------------

// gtk_buildable_set_name
// gtk_buildable_get_name
// gtk_buildable_add_child
// gtk_buildable_set_buildable_property
// gtk_buildable_construct_child
// gtk_buildable_custom_tag_start
// gtk_buildable_custom_tag_end
// gtk_buildable_custom_finished
// gtk_buildable_parser_finished
// gtk_buildable_get_internal_child

//-----------------------------------------------------------------------
// GtkBuilder
//-----------------------------------------------------------------------
type Builder struct {
	GBuilder *C.GtkBuilder
}

func NewBuilder() *Builder {
	return &Builder{
		C.gtk_builder_new()}
}

func (v *Builder) AddFromFile(filename string) (ret uint, error *glib.Error) {
	var gerror *C.GError
	ptr := C.CString(filename)
	defer cfree(ptr)
	ret = uint(C.gtk_builder_add_from_file(v.GBuilder, gstring(ptr), &gerror))
	if gerror != nil {
		error = glib.ErrorFromNative(unsafe.Pointer(gerror))
	} else {
		error = nil
	}
	return
}

func (v *Builder) AddFromString(buffer string) (ret uint, error *glib.Error) {
	var gerror *C.GError
	ptr := C.CString(buffer)
	defer cfree(ptr)
	ret = uint(C.gtk_builder_add_from_string(v.GBuilder, gstring(ptr), C.gsize(C.strlen(ptr)), &gerror))
	if gerror != nil {
		error = glib.ErrorFromNative(unsafe.Pointer(gerror))
	} else {
		error = nil
	}
	return
}

// guint gtk_builder_add_objects_from_file (GtkBuilder *builder, const gchar *filename, gchar **object_ids, GError **error);
// guint gtk_builder_add_objects_from_string (GtkBuilder *builder, const gchar *buffer, gsize length, gchar **object_ids, GError **error);

func (v *Builder) GetObject(name string) *glib.GObject {
	ptr := C.CString(name)
	defer cfree(ptr)
	return &glib.GObject{unsafe.Pointer(C.gtk_builder_get_object(v.GBuilder, gstring(ptr)))}
}

func (v *Builder) GetObjects() *glib.SList {
	return glib.SListFromNative(unsafe.Pointer(C.gtk_builder_get_objects(v.GBuilder)))
}

func (v *Builder) ConnectSignals(user_data interface{}) {
	C.gtk_builder_connect_signals(v.GBuilder, nil)
}

type BuilderConnectFunc func(builder *Builder, obj *glib.GObject, sig, handler string, conn *glib.GObject, flags glib.ConnectFlags, user_data interface{})

type BuilderConnectInfo struct {
	cb        BuilderConnectFunc
	user_data interface{}
}

func (v *Builder) ConnectSignalsFull(f BuilderConnectFunc, user_data interface{}) {
	C._gtk_builder_connect_signals_full(v.GBuilder, pointer.Save(&BuilderConnectInfo{f, user_data}))
}

//export _go_gtk_builder_connect_signals_full_cb
func _go_gtk_builder_connect_signals_full_cb(builder unsafe.Pointer, obj unsafe.Pointer, signal_name *C.gchar, handler_name *C.gchar, connect_object unsafe.Pointer, flags C.int, user_data unsafe.Pointer) {
	ci := pointer.Restore(user_data).(*BuilderConnectInfo)
	ci.cb(
		&Builder{(*C.GtkBuilder)(builder)},
		glib.ObjectFromNative(obj),
		gostring(signal_name),
		gostring(handler_name),
		glib.ObjectFromNative(connect_object),
		glib.ConnectFlags(flags),
		ci.user_data,
	)
}

func (v *Builder) SetTranslationDomain(domain string) {
	ptr := C.CString(domain)
	defer cfree(ptr)
	C.gtk_builder_set_translation_domain(v.GBuilder, gstring(ptr))
}

func (v *Builder) GetTranslationDomain() string {
	return gostring(C.gtk_builder_get_translation_domain(v.GBuilder))
}

func (v *Builder) GetTypeFromName(type_name string) int {
	ptr := C.CString(type_name)
	defer cfree(ptr)
	return int(C.gtk_builder_get_type_from_name(v.GBuilder, ptr))
}

// gboolean gtk_builder_value_from_string (GtkBuilder *builder, GParamSpec *pspec, const gchar *string, GValue *value, GError **error);
// gboolean gtk_builder_value_from_string_type (GtkBuilder *builder, GType type, const gchar *string, GValue *value, GError **error);
