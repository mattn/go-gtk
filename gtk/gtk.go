/*
Go Bindings for Gtk+ 2. Support version 2.16 and later.
*/
package gtk

// #include "gtk.go.h"
// #cgo pkg-config: gtk+-2.0
import "C"
import (
	"fmt"
	"github.com/mattn/go-gtk/gdk"
	"github.com/mattn/go-gtk/gdkpixbuf"
	"github.com/mattn/go-gtk/glib"
	"github.com/mattn/go-gtk/pango"
	"log"
	"reflect"
	"runtime"
	"strings"
	"unsafe"
)

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

func variadicButtonsToArrays(buttons []interface{}) ([]string, []int) {
	if len(buttons)%2 != 0 {
		argumentPanic("variadic parameter must be even (couples of string-int (button label - button response)")
	}
	text := make([]string, len(buttons)/2)
	res := make([]int, len(buttons)/2)
	for i := 0; i < len(text); i++ {
		btext, ok := buttons[2*i].(string)
		if !ok {
			argumentPanic("button text must be a string")
		}
		bresponse := func() int {
			switch val := buttons[2*i+1].(type) {
			case int:
				return val
			case ResponseType:
				return int(val)
			default:
				argumentPanic("button response must be an int")
			}
			panic(nil)
		}()
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
		var argc C.int = C.int(len(*args))
		cargs := make([]*C.char, argc)
		for i, arg := range *args {
			cargs[i] = C.CString(arg)
		}
		C._gtk_init(unsafe.Pointer(&argc), unsafe.Pointer(&cargs))
		goargs := make([]string, argc)
		for i := 0; i < int(argc); i++ {
			goargs[i] = C.GoString(cargs[i])
		}
		for i := 0; i < int(argc); i++ {
			C.free_string(cargs[i])
		}
		*args = goargs
	} else {
		C._gtk_init(nil, nil)
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
	return gboolean2bool(C.gtk_main_iteration())
}
func MainIterationDo(blocking bool) bool {
	return gboolean2bool(C.gtk_main_iteration_do(bool2gboolean(blocking)))
}

func EventsPending() bool {
	return gboolean2bool(C.gtk_events_pending())
}

// gtk_main_do_event
// gtk_grab_add
// gtk_grab_get_current
// gtk_grab_remove
// gtk_key_snooper_install
// gtk_key_snooper_remove
// gtk_get_current_event
// gtk_get_current_event_time
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
// gtk_accelerator_parse
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

func NewClipboardGetForDisplay(display *gdk.GdkDisplay, selection gdk.GdkAtom) *Clipboard {
	var cdisplay unsafe.Pointer
	if display != nil {
		cdisplay = display.Display
	}
	return &Clipboard{C._gtk_clipboard_get_for_display(cdisplay, unsafe.Pointer(selection))}
}

func (v *Clipboard) Clear() {
	C.gtk_clipboard_clear(v.GClipboard)
}

func (v *Clipboard) SetText(text string) {
	ptr := C.to_charptr_voidp(unsafe.Pointer(&([]byte(text))[0]))
	C.gtk_clipboard_set_text(v.GClipboard, C.to_gcharptr(ptr), C.gint(-1))
}

func (v *Clipboard) SetImage(pixbuf *gdkpixbuf.GdkPixbuf) {
	C.gtk_clipboard_set_image(v.GClipboard, pixbuf.Pixbuf)
}

func (v *Clipboard) Store() {
	C.gtk_clipboard_store(v.GClipboard)
}

func (v *Clipboard) WaitForText() string {
	return C.GoString(C.to_charptr(C.gtk_clipboard_wait_for_text(v.GClipboard)))
}

func (v *Clipboard) Connect(s string, f interface{}, datas ...interface{}) int {
	return glib.ObjectFromNative(unsafe.Pointer(C.to_GObject(unsafe.Pointer(v.GClipboard)))).Connect(s, f, datas...)
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
	DEST_DEFAULT_MOTION    DestDefaults = 1 << 0 /* respond to "drag_motion" */
	DEST_DEFAULT_HIGHLIGHT DestDefaults = 1 << 1 /* auto-highlight */
	DEST_DEFAULT_DROP      DestDefaults = 1 << 2 /* respond to "drag_drop" */
	DEST_DEFAULT_ALL       DestDefaults = 0x07
)

type TargetEntry struct {
	Target string
	Flags  uint
	Info   uint
}

func (v *Widget) DragDestSet(flags DestDefaults, targets []TargetEntry, actions gdk.GdkDragAction) {
	ctargets := make([]C.GtkTargetEntry, len(targets))
	for i, target := range targets {
		ptr := C.CString(target.Target)
		defer C.free_string(ptr)
		ctargets[i].target = C.to_gcharptr(ptr)
		ctargets[i].flags = C.guint(target.Flags)
		ctargets[i].info = C.guint(target.Info)
	}
	C.gtk_drag_dest_set(v.GWidget, C.GtkDestDefaults(flags), &ctargets[0], C.gint(len(targets)), C.GdkDragAction(actions))
}

func (v *Widget) DragSourceSet(start_button_mask gdk.GdkModifierType, targets []TargetEntry, actions gdk.GdkDragAction) {
	ctargets := make([]C.GtkTargetEntry, len(targets))
	for i, target := range targets {
		ptr := C.CString(target.Target)
		defer C.free_string(ptr)
		ctargets[i].target = C.to_gcharptr(ptr)
		ctargets[i].flags = C.guint(target.Flags)
		ctargets[i].info = C.guint(target.Info)
	}
	C.gtk_drag_source_set(v.GWidget, C.GdkModifierType(start_button_mask), &ctargets[0], C.gint(len(targets)), C.GdkDragAction(actions))
}

func (v *Widget) DragFinish(context *gdk.GdkDragContext, success bool, del bool, time uint) {
	C._gtk_drag_finish(unsafe.Pointer(context.DragContext), bool2gboolean(success), bool2gboolean(del), C.guint32(time))
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
	C.gtk_stock_add(v.GStockItem, C.guint(nitems))
}
func (v *StockItem) AddStatic(nitems uint) {
	C.gtk_stock_add_static(v.GStockItem, C.guint(nitems))
}
func StockLookup(stock_id string, item *StockItem) bool {
	ptr := C.CString(stock_id)
	defer C.free_string(ptr)
	return gboolean2bool(C.gtk_stock_lookup(C.to_gcharptr(ptr), item.GStockItem))
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
// gtk_rc_parse
// gtk_rc_parse_string
// gtk_rc_reparse_all
// gtk_rc_reparse_all_for_settings
// gtk_rc_reset_styles
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

// gtk_settings_get_default
// gtk_settings_get_for_screen
// gtk_settings_install_property
// gtk_settings_install_property_parser
// gtk_rc_property_parse_color
// gtk_rc_property_parse_enum
// gtk_rc_property_parse_flags
// gtk_rc_property_parse_requisition
// gtk_rc_property_parse_border
// gtk_settings_set_property_value

func (s *Settings) SetStringProperty(name string, v_string string, origin string) {
	ptrn := C.CString(name)
	defer C.free_string(ptrn)
	ptrv := C.CString(v_string)
	defer C.free_string(ptrv)
	prts := C.CString(origin)
	defer C.free_string(prts)
	C.gtk_settings_set_string_property(s.GSettings, C.to_gcharptr(ptrn), C.to_gcharptr(ptrv), C.to_gcharptr(prts))
}
func (s *Settings) SetLongProperty(name string, v_long int32, origin string) {
	ptrn := C.CString(name)
	defer C.free_string(ptrn)
	prts := C.CString(origin)
	defer C.free_string(prts)
	C.gtk_settings_set_long_property(s.GSettings, C.to_gcharptr(ptrn), C.glong(v_long), C.to_gcharptr(prts))
}
func (s *Settings) SetDoubleProperty(name string, v_double float64, origin string) {
	ptrn := C.CString(name)
	defer C.free_string(ptrn)
	prts := C.CString(origin)
	defer C.free_string(prts)
	C.gtk_settings_set_double_property(s.GSettings, C.to_gcharptr(ptrn), C.gdouble(v_double), C.to_gcharptr(prts))
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

// gtk_style_new
// gtk_style_copy
// gtk_style_attach
// gtk_style_detach
// gtk_style_ref
// gtk_style_unref
// gtk_style_set_background
// gtk_style_apply_default_background
// gtk_style_apply_default_pixmap
// gtk_style_lookup_color
// gtk_style_lookup_icon_set
// gtk_style_render_icon
// gtk_style_get_font
// gtk_style_set_font
// gtk_style_get_style_property
// gtk_style_get_valist
// gtk_style_get
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
// gtk_paint_string
// gtk_paint_tab
// gtk_paint_vline
// gtk_paint_expander
// gtk_paint_layout
// gtk_paint_resize_grip
// gtk_draw_insertion_cursor
// gtk_border_new
// gtk_border_copy
// gtk_border_free

//-----------------------------------------------------------------------
// Selections
//-----------------------------------------------------------------------

type SelectionData struct {
	GSelectionData unsafe.Pointer
}

func SelectionDataFromNative(l unsafe.Pointer) *SelectionData {
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
	DIALOG_MODAL               DialogFlags = 1 << 0 /* call gtk_window_set_modal (win, TRUE) */
	DIALOG_DESTROY_WITH_PARENT DialogFlags = 1 << 1 /* call gtk_window_set_destroy_with_parent () */
	DIALOG_NO_SEPARATOR        DialogFlags = 1 << 2 /* no separator bar above buttons */
)

type ResponseType int

const (
	RESPONSE_NONE         ResponseType = -1
	RESPONSE_REJECT       ResponseType = -2
	RESPONSE_ACCEPT       ResponseType = -3
	RESPONSE_DELETE_EVENT ResponseType = -4
	RESPONSE_OK           ResponseType = -5
	RESPONSE_CANCEL       ResponseType = -6
	RESPONSE_CLOSE        ResponseType = -7
	RESPONSE_YES          ResponseType = -8
	RESPONSE_NO           ResponseType = -9
	RESPONSE_APPLY        ResponseType = -10
	RESPONSE_HELP         ResponseType = -11
)

/*type DialogLike interface {
	WidgetLike
	Run() int
	Response(interface{}, ...interface{})
}*/

type Dialog struct {
	Window
}

func NewDialog() *Dialog {
	return &Dialog{Window{Bin{Container{Widget{C.gtk_dialog_new()}}}}}
}

// gtk_dialog_new_with_buttons

func (v *Dialog) Run() ResponseType {
	return ResponseType(C.gtk_dialog_run(C.to_GtkDialog(v.GWidget)))
}
func (v *Dialog) Response(response interface{}, datas ...interface{}) {
	v.Connect("response", response, datas...)
}
func (v *Dialog) AddButton(button_text string, response_id int) *Button {
	ptr := C.CString(button_text)
	defer C.free_string(ptr)
	return &Button{Bin{Container{Widget{
		C.gtk_dialog_add_button(C.to_GtkDialog(v.GWidget), C.to_gcharptr(ptr), C.gint(response_id))}}}}
}

// gtk_dialog_add_buttons
// gtk_dialog_add_action_widget
// gtk_dialog_get_has_separator //deprecated since 2.22

//Deprecated since 2.22.
func (v *Dialog) SetHasSeparator(f bool) {
	deprecated_since(2, 22, 0, "gtk_dialog_set_has_separator()")
	C.gtk_dialog_set_has_separator(C.to_GtkDialog(v.GWidget), bool2gboolean(f))
}
func (v *Dialog) SetDefaultResponse(id int) {
	C.gtk_dialog_set_default_response(C.to_GtkDialog(v.GWidget), C.gint(id))
}

// gtk_dialog_set_has_separator //deprecated since 2.22
// gtk_dialog_set_response_sensitive

func (v *Dialog) GetResponseForWidget(w *Widget) int {
	return int(C.gtk_dialog_get_response_for_widget(C.to_GtkDialog(v.GWidget), w.GWidget))
}
func (v *Dialog) GetWidgetForResponse(id int) *Widget {
	panic_if_version_older(2, 20, 0, "gtk_dialog_get_widget_for_response()")
	return &Widget{C._gtk_dialog_get_widget_for_response(C.to_GtkDialog(v.GWidget), C.gint(id))}
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
	MESSAGE_INFO     = 0
	MESSAGE_WARNING  = 1
	MESSAGE_QUESTION = 2
	MESSAGE_ERROR    = 3
	MESSAGE_OTHER    = 4
)

type ButtonsType int

const (
	BUTTONS_NONE      = 0
	BUTTONS_OK        = 1
	BUTTONS_CLOSE     = 2
	BUTTONS_CANCEL    = 3
	BUTTONS_YES_NO    = 4
	BUTTONS_OK_CANCEL = 5
)

type MessageDialog struct {
	Dialog
}

// TODO should be variadic function
func NewMessageDialog(parent *Window, flag DialogFlags, t MessageType, buttons ButtonsType, format string, args ...interface{}) *MessageDialog {
	ptr := C.CString(strings.Replace(fmt.Sprintf(format, args...), "%", "%%", -1))
	defer C.free_string(ptr)
	return &MessageDialog{Dialog{Window{Bin{Container{Widget{
		C._gtk_message_dialog_new(
			parent.ToNative(),
			C.GtkDialogFlags(flag),
			C.GtkMessageType(t),
			C.GtkButtonsType(buttons),
			C.to_gcharptr(ptr))}}}}}}
}

func NewMessageDialogWithMarkup(parent *Window, flag DialogFlags, t MessageType, buttons ButtonsType, format string, args ...interface{}) *MessageDialog {
	r := &MessageDialog{Dialog{Window{Bin{Container{Widget{
		C._gtk_message_dialog_new_with_markup(
			parent.ToNative(),
			C.GtkDialogFlags(flag),
			C.GtkMessageType(t),
			C.GtkButtonsType(buttons),
			nil)}}}}}}
	r.SetMarkup(fmt.Sprintf(format, args...))
	return r
}

func (v *MessageDialog) SetMarkup(markup string) {
	ptr := C.CString(markup)
	defer C.free_string(ptr)
	C.gtk_message_dialog_set_markup(C.to_GtkMessageDialog(v.GWidget), C.to_gcharptr(ptr))
}
func (v *MessageDialog) SetImage(image WidgetLike) {
	C.gtk_message_dialog_set_image(C.to_GtkMessageDialog(v.GWidget), image.ToNative())
}
func (v *MessageDialog) GetImage() *Image {
	return &Image{Widget{C.gtk_message_dialog_get_image(C.to_GtkMessageDialog(v.GWidget))}}
}

// gtk_message_dialog_get_message_area //since 2.22
// gtk_message_dialog_format_secondary_text
// gtk_message_dialog_format_secondary_markup

//-----------------------------------------------------------------------
// GtkWindow
//-----------------------------------------------------------------------
type WindowType int

const (
	WINDOW_TOPLEVEL WindowType = 0
	WINDOW_POPUP    WindowType = 1
)

type WindowPosition int

const (
	WIN_POS_NONE             WindowPosition = 0
	WIN_POS_CENTER           WindowPosition = 1
	WIN_POS_MOUSE            WindowPosition = 2
	WIN_POS_CENTER_ALWAYS    WindowPosition = 3
	WIN_POS_CENTER_ON_PARENT WindowPosition = 4
)

/*type WindowLike interface {
	ContainerLike
	SetTransientFor(parent WindowLike)
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
	defer C.free_string(ptr)
	C.gtk_window_set_title(C.to_GtkWindow(v.GWidget), C.to_gcharptr(ptr))
}

// gtk_window_set_wmclass

func (v *Window) SetResizable(resizable bool) {
	C.gtk_window_set_resizable(C.to_GtkWindow(v.GWidget), bool2gboolean(resizable))
}
func (v *Window) GetResizable() bool {
	return gboolean2bool(C.gtk_window_get_resizable(C.to_GtkWindow(v.GWidget)))
}
func (v *Window) AddAccelGroup(group *AccelGroup) {
	C.gtk_window_add_accel_group(C.to_GtkWindow(v.GWidget), group.GAccelGroup)
}

// gtk_window_remove_accel_group
// gtk_window_activate_focus
// gtk_window_activate_default

func (v *Window) SetModal(modal bool) {
	C.gtk_window_set_modal(C.to_GtkWindow(v.GWidget), bool2gboolean(modal))
}
func (v *Window) SetDefaultSize(width int, height int) {
	C.gtk_window_set_default_size(C.to_GtkWindow(v.GWidget), C.gint(width), C.gint(height))
}

// gtk_window_set_geometry_hints
// gtk_window_set_gravity
// gtk_window_get_gravity

func (v *Window) SetPosition(position WindowPosition) {
	C.gtk_window_set_position(C.to_GtkWindow(v.GWidget), C.GtkWindowPosition(position))
}
func (v *Window) SetTransientFor(parent *Window) {
	C.gtk_window_set_transient_for(C.to_GtkWindow(v.GWidget), C.to_GtkWindow(parent.ToNative()))
}
func (v *Window) SetDestroyWithParent(setting bool) {
	C.gtk_window_set_destroy_with_parent(C.to_GtkWindow(v.GWidget), bool2gboolean(setting))
}

// gtk_window_set_screen
// gtk_window_get_screen
// gtk_window_is_active
// gtk_window_has_toplevel_focus
// gtk_window_list_toplevels
// gtk_window_add_mnemonic
// gtk_window_remove_mnemonic
// gtk_window_mnemonic_activate
// gtk_window_activate_key
// gtk_window_propagate_key_event
// gtk_window_get_focus
// gtk_window_set_focus
// gtk_window_get_default_widget

func (v *Window) SetDefault(w *Widget) {
	C.gtk_window_set_default(C.to_GtkWindow(v.GWidget), w.GWidget)
}
func (v *Window) Present() {
	C.gtk_window_present(C.to_GtkWindow(v.GWidget))
}

// gtk_window_present_with_time

func (v *Window) Stick() {
	C.gtk_window_stick(C.to_GtkWindow(v.GWidget))
}

func (v *Window) Unstick() {
	C.gtk_window_unstick(C.to_GtkWindow(v.GWidget))
}

func (v *Window) Iconify() {
	C.gtk_window_iconify(C.to_GtkWindow(v.GWidget))
}

func (v *Window) Deiconify() {
	C.gtk_window_deiconify(C.to_GtkWindow(v.GWidget))
}

func (v *Window) Maximize() {
	C.gtk_window_maximize(C.to_GtkWindow(v.GWidget))
}
func (v *Window) Unmaximize() {
	C.gtk_window_unmaximize(C.to_GtkWindow(v.GWidget))
}

func (v *Window) Fullscreen() {
	C.gtk_window_fullscreen(C.to_GtkWindow(v.GWidget))
}
func (v *Window) Unfullscreen() {
	C.gtk_window_unfullscreen(C.to_GtkWindow(v.GWidget))
}
func (v *Window) SetKeepAbove(setting bool){
	C.gtk_window_set_keep_above(C.to_GtkWindow(v.GWidget), bool2gboolean(setting))
}
func (v *Window) SetKeepBelow(setting bool){
	C.gtk_window_set_keep_below(C.to_GtkWindow(v.GWidget), bool2gboolean(setting))
}
func (v *Window) SetDecorated(setting bool){
	C.gtk_window_set_decorated(C.to_GtkWindow(v.GWidget), bool2gboolean(setting))
}
func (v *Window) SetDeletable(setting bool){
	C.gtk_window_set_deletable(C.to_GtkWindow(v.GWidget), bool2gboolean(setting))
}
func (v *GtkWindow) SetTypeHint(hint gdk.GdkWindowTypeHint) {
	C.gtk_window_set_type_hint(C.to_GtkWindow(v.Widget), C.GdkWindowTypeHint(hint))
}

// gtk_window_begin_resize_drag
// gtk_window_begin_move_drag
// gtk_window_set_frame_dimensions //deprecated since 2.24
// gtk_window_set_has_frame  //deprecated since 2.24
// gtk_window_set_mnemonic_modifier
// gtk_window_set_skip_taskbar_hint
// gtk_window_set_skip_pager_hint
// gtk_window_set_urgency_hint

func (v *Window) SetAcceptFocus(setting bool) {
	C.gtk_window_set_accept_focus(C.to_GtkWindow(v.GWidget), bool2gboolean(setting))
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
	C.gtk_window_get_default_size(C.to_GtkWindow(v.GWidget), &cwidth, &cheight)
	*width = int(cwidth)
	*height = int(cheight)
}
func (v *Window) GetDestroyWithParent() bool {
	return gboolean2bool(C.gtk_window_get_destroy_with_parent(C.to_GtkWindow(v.GWidget)))
}

// gtk_window_get_frame_dimensions //deprecated since 2.24
// gtk_window_get_has_frame  //deprecated since 2.24
// gtk_window_get_icon
// gtk_window_get_icon_list

func (v *Window) GetIconName() string {
	return C.GoString(C.to_charptr(C.gtk_window_get_icon_name(C.to_GtkWindow(v.GWidget))))
}
// gtk_window_get_mnemonic_modifier

func (v *Window) GetModal() bool {
	return gboolean2bool(C.gtk_window_get_modal(C.to_GtkWindow(v.GWidget)))
}
func (v *Window) GetPosition(root_x *int, root_y *int) {
	var croot_x, croot_y C.gint
	C.gtk_window_get_position(C.to_GtkWindow(v.GWidget), &croot_x, &croot_y)
	*root_x = int(croot_x)
	*root_y = int(croot_y)
}

// gtk_window_get_role

func (v *Window) GetSize(width *int, height *int) {
	var cwidth, cheight C.gint
	C.gtk_window_get_size(C.to_GtkWindow(v.GWidget), &cwidth, &cheight)
	*width = int(cwidth)
	*height = int(cheight)
}
func (v *Window) GetTitle() string {
	return C.GoString(C.to_charptr(C.gtk_window_get_title(C.to_GtkWindow(v.GWidget))))
}

func (v *GtkWindow) GetTypeHint() gdk.GdkWindowTypeHint {
	return gdk.GdkWindowTypeHint(C.gtk_window_get_type_hint(C.to_GtkWindow(v.Widget)))
}

// gtk_window_get_transient_for
// gtk_window_get_skip_taskbar_hint
// gtk_window_get_skip_pager_hint
// gtk_window_get_urgency_hint

func (v *Window) GetAcceptFocus() bool {
	return gboolean2bool(C.gtk_window_get_accept_focus(C.to_GtkWindow(v.GWidget)))
}

// gtk_window_get_focus_on_map
// gtk_window_get_group
// gtk_window_has_group //since 2.22
// gtk_window_get_window_type //since 2.20

func (v *Window) Move(x int, y int) {
	C.gtk_window_move(C.to_GtkWindow(v.GWidget), C.gint(x), C.gint(y))
}

// gtk_window_parse_geometry
// gtk_window_reshow_with_initial_size

func (v *Window) Resize(width int, height int) {
	C.gtk_window_resize(C.to_GtkWindow(v.GWidget), C.gint(width), C.gint(height))
}

func (v *Window) XID() int32 {
	return gdk.WindowFromUnsafe(unsafe.Pointer(v.GWidget.window)).GetNativeWindowID()
}

// gtk_window_set_default_icon_list
// gtk_window_set_default_icon
// gtk_window_set_default_icon_from_file
// gtk_window_set_default_icon_name
// gtk_window_set_icon
// gtk_window_set_icon_list

func (v *Window) SetIconFromFile(file string) {
	ptr := C.CString(file)
	defer C.free_string(ptr)
	C.gtk_window_set_icon_from_file(C.to_GtkWindow(v.GWidget), C.to_gcharptr(ptr), nil) // last arg : GError **err
}

func (v *Window) SetIconName(name string) {
	ptr := C.CString(name)
	defer C.free_string(ptr)
	C.gtk_window_set_icon_name(C.to_GtkWindow(v.GWidget), C.to_gcharptr(ptr))
}

// gtk_window_set_auto_startup_notification
// gtk_window_get_opacity
// gtk_window_set_opacity
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
	return &AboutDialog{Dialog{Window{Bin{Container{Widget{
		C.gtk_about_dialog_new()}}}}}}
}
func (v *AboutDialog) GetProgramName() string {
	return C.GoString(C.to_charptr(C.gtk_about_dialog_get_program_name(C.to_GtkAboutDialog(v.GWidget))))
}
func (v *AboutDialog) SetProgramName(name string) {
	ptr := C.CString(name)
	defer C.free_string(ptr)
	C.gtk_about_dialog_set_program_name(C.to_GtkAboutDialog(v.GWidget), C.to_gcharptr(ptr))
}
func (v *AboutDialog) GetVersion() string {
	return C.GoString(C.to_charptr(C.gtk_about_dialog_get_version(C.to_GtkAboutDialog(v.GWidget))))
}
func (v *AboutDialog) SetVersion(version string) {
	ptr := C.CString(version)
	defer C.free_string(ptr)
	C.gtk_about_dialog_set_version(C.to_GtkAboutDialog(v.GWidget), C.to_gcharptr(ptr))
}
func (v *AboutDialog) GetCopyright() string {
	return C.GoString(C.to_charptr(C.gtk_about_dialog_get_copyright(C.to_GtkAboutDialog(v.GWidget))))
}
func (v *AboutDialog) SetCopyright(copyright string) {
	ptr := C.CString(copyright)
	defer C.free_string(ptr)
	C.gtk_about_dialog_set_copyright(C.to_GtkAboutDialog(v.GWidget), C.to_gcharptr(ptr))
}
func (v *AboutDialog) GetComments() string {
	return C.GoString(C.to_charptr(C.gtk_about_dialog_get_comments(C.to_GtkAboutDialog(v.GWidget))))
}
func (v *AboutDialog) SetComments(comments string) {
	ptr := C.CString(comments)
	defer C.free_string(ptr)
	C.gtk_about_dialog_set_comments(C.to_GtkAboutDialog(v.GWidget), C.to_gcharptr(ptr))
}
func (v *AboutDialog) GetLicense() string {
	return C.GoString(C.to_charptr(C.gtk_about_dialog_get_license(C.to_GtkAboutDialog(v.GWidget))))
}
func (v *AboutDialog) SetLicense(license string) {
	ptr := C.CString(license)
	defer C.free_string(ptr)
	C.gtk_about_dialog_set_license(C.to_GtkAboutDialog(v.GWidget), C.to_gcharptr(ptr))
}
func (v *AboutDialog) GetWrapLicense() bool {
	return gboolean2bool(C.gtk_about_dialog_get_wrap_license(C.to_GtkAboutDialog(v.GWidget)))
}
func (v *AboutDialog) SetWrapLicense(wrap_license bool) {
	C.gtk_about_dialog_set_wrap_license(C.to_GtkAboutDialog(v.GWidget), bool2gboolean(wrap_license))
}
func (v *AboutDialog) GetWebsite() string {
	return C.GoString(C.to_charptr(C.gtk_about_dialog_get_website(C.to_GtkAboutDialog(v.GWidget))))
}
func (v *AboutDialog) SetWebsite(website string) {
	ptr := C.CString(website)
	defer C.free_string(ptr)
	C.gtk_about_dialog_set_website(C.to_GtkAboutDialog(v.GWidget), C.to_gcharptr(ptr))
}
func (v *AboutDialog) GetWebsiteLabel() string {
	return C.GoString(C.to_charptr(C.gtk_about_dialog_get_website_label(C.to_GtkAboutDialog(v.GWidget))))
}
func (v *AboutDialog) SetWebsiteLabel(website_label string) {
	ptr := C.CString(website_label)
	defer C.free_string(ptr)
	C.gtk_about_dialog_set_website_label(C.to_GtkAboutDialog(v.GWidget), C.to_gcharptr(ptr))
}
func (v *AboutDialog) GetAuthors() []string {
	var authors []string
	cauthors := C.gtk_about_dialog_get_authors(C.to_GtkAboutDialog(v.GWidget))
	for {
		authors = append(authors, C.GoString(C.to_charptr(*cauthors)))
		cauthors = C.next_gcharptr(cauthors)
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
		defer C.free_string(ptr)
		C.set_string(cauthors, C.int(i), C.to_gcharptr(ptr))
	}
	C.set_string(cauthors, C.int(len(authors)), nil)
	C.gtk_about_dialog_set_authors(C.to_GtkAboutDialog(v.GWidget), cauthors)
	C.destroy_strings(cauthors)
}
func (v *AboutDialog) GetArtists() []string {
	var artists []string
	cartists := C.gtk_about_dialog_get_artists(C.to_GtkAboutDialog(v.GWidget))
	for {
		artists = append(artists, C.GoString(C.to_charptr(*cartists)))
		cartists = C.next_gcharptr(cartists)
		if *cartists == nil {
			break
		}
	}
	return artists
}
func (v *AboutDialog) SetArtists(artists []string) {
	cartists := C.make_strings(C.int(len(artists)))
	for i, author := range artists {
		ptr := C.CString(author)
		defer C.free_string(ptr)
		C.set_string(cartists, C.int(i), C.to_gcharptr(ptr))
	}
	C.gtk_about_dialog_set_artists(C.to_GtkAboutDialog(v.GWidget), cartists)
	C.destroy_strings(cartists)
}
func (v *AboutDialog) GetDocumenters() []string {
	var documenters []string
	cdocumenters := C.gtk_about_dialog_get_documenters(C.to_GtkAboutDialog(v.GWidget))
	for {
		documenters = append(documenters, C.GoString(C.to_charptr(*cdocumenters)))
		cdocumenters = C.next_gcharptr(cdocumenters)
		if *cdocumenters == nil {
			break
		}
	}
	return documenters
}
func (v *AboutDialog) SetDocumenters(documenters []string) {
	cdocumenters := C.make_strings(C.int(len(documenters)))
	for i, author := range documenters {
		ptr := C.CString(author)
		defer C.free_string(ptr)
		C.set_string(cdocumenters, C.int(i), C.to_gcharptr(ptr))
	}
	C.gtk_about_dialog_set_documenters(C.to_GtkAboutDialog(v.GWidget), cdocumenters)
	C.destroy_strings(cdocumenters)
}
func (v *AboutDialog) GetTranslatorCredits() string {
	return C.GoString(C.to_charptr(C.gtk_about_dialog_get_translator_credits(C.to_GtkAboutDialog(v.GWidget))))
}
func (v *AboutDialog) SetTranslatorCredits(translator_credits string) {
	ptr := C.CString(translator_credits)
	defer C.free_string(ptr)
	C.gtk_about_dialog_set_translator_credits(C.to_GtkAboutDialog(v.GWidget), C.to_gcharptr(ptr))
}
func (v *AboutDialog) GetLogo() *gdkpixbuf.GdkPixbuf {
	return &gdkpixbuf.GdkPixbuf{
		C.gtk_about_dialog_get_logo(C.to_GtkAboutDialog(v.GWidget))}
}
func (v *AboutDialog) SetLogo(logo *gdkpixbuf.GdkPixbuf) {
	C.gtk_about_dialog_set_logo(C.to_GtkAboutDialog(v.GWidget), logo.Pixbuf)
}
func (v *AboutDialog) GetLogoIconName() string {
	return C.GoString(C.to_charptr(C.gtk_about_dialog_get_logo_icon_name(C.to_GtkAboutDialog(v.GWidget))))
}
func (v *AboutDialog) SetLogoIconName(icon_name string) {
	ptr := C.CString(icon_name)
	defer C.free_string(ptr)
	C.gtk_about_dialog_set_logo_icon_name(C.to_GtkAboutDialog(v.GWidget), C.to_gcharptr(ptr))
}

// gtk_about_dialog_set_email_hook //deprecated since 2.24
// gtk_about_dialog_set_url_hook //deprecated since 2.24
// gtk_show_about_dialog

//-----------------------------------------------------------------------
// GtkAssistant
//-----------------------------------------------------------------------
type AssistantPageType int

const (
	ASSISTANT_PAGE_CONTENT  AssistantPageType = 0
	ASSISTANT_PAGE_INTRO    AssistantPageType = 1
	ASSISTANT_PAGE_CONFIRM  AssistantPageType = 2
	ASSISTANT_PAGE_SUMMARY  AssistantPageType = 3
	ASSISTANT_PAGE_PROGRESS AssistantPageType = 4
)

type Assistant struct {
	Widget
}

func NewAssistant() *Assistant {
	return &Assistant{Widget{C.gtk_assistant_new()}}
}
func (v *Assistant) GetCurrentPage() int {
	return int(C.gtk_assistant_get_current_page(C.to_GtkAssistant(v.GWidget)))
}
func (v *Assistant) SetCurrentPage(page_num int) {
	C.gtk_assistant_set_current_page(C.to_GtkAssistant(v.GWidget), C.gint(page_num))
}
func (v *Assistant) GetNPages() int {
	return int(C.gtk_assistant_get_n_pages(C.to_GtkAssistant(v.GWidget)))
}
func (v *Assistant) GetNthPage(page_num int) *Widget {
	return &Widget{
		C.gtk_assistant_get_nth_page(C.to_GtkAssistant(v.GWidget), C.gint(page_num))}
}
func (v *Assistant) PrependPage(page WidgetLike) int {
	return int(C.gtk_assistant_prepend_page(C.to_GtkAssistant(v.GWidget), page.ToNative()))
}
func (v *Assistant) AppendPage(page WidgetLike) int {
	return int(C.gtk_assistant_prepend_page(C.to_GtkAssistant(v.GWidget), page.ToNative()))
}
func (v *Assistant) InsertPage(page WidgetLike, position int) int {
	return int(C.gtk_assistant_insert_page(C.to_GtkAssistant(v.GWidget), page.ToNative(), C.gint(position)))
}

// void gtk_assistant_set_forward_page_func (GtkAssistant *assistant, GtkAssistantPageFunc page_func, gpointer data, GDestroyNotify destroy);

func (v *Assistant) SetPageType(page WidgetLike, t AssistantPageType) {
	C.gtk_assistant_set_page_type(C.to_GtkAssistant(v.GWidget), page.ToNative(), C.GtkAssistantPageType(t))
}
func (v *Assistant) GetPageType(page WidgetLike) AssistantPageType {
	return AssistantPageType(C.gtk_assistant_get_page_type(C.to_GtkAssistant(v.GWidget), page.ToNative()))
}
func (v *Assistant) SetPageTitle(page WidgetLike, title string) {
	ptr := C.CString(title)
	defer C.free_string(ptr)
	C.gtk_assistant_set_page_title(C.to_GtkAssistant(v.GWidget), page.ToNative(), C.to_gcharptr(ptr))
}
func (v *Assistant) GetPageTitle(page WidgetLike) string {
	return C.GoString(C.to_charptr(C.gtk_assistant_get_page_title(C.to_GtkAssistant(v.GWidget), page.ToNative())))
}
func (v *Assistant) SetPageHeaderImage(page WidgetLike, pixbuf *gdkpixbuf.GdkPixbuf) {
	C.gtk_assistant_set_page_header_image(C.to_GtkAssistant(v.GWidget), page.ToNative(), pixbuf.Pixbuf)
}
func (v *Assistant) GetPageHeaderImage(page WidgetLike) *gdkpixbuf.GdkPixbuf {
	return &gdkpixbuf.GdkPixbuf{
		C.gtk_assistant_get_page_header_image(C.to_GtkAssistant(v.GWidget), page.ToNative())}
}
func (v *Assistant) SetPageSideImage(page WidgetLike, pixbuf *gdkpixbuf.GdkPixbuf) {
	C.gtk_assistant_set_page_side_image(C.to_GtkAssistant(v.GWidget), page.ToNative(), pixbuf.Pixbuf)
}
func (v *Assistant) GetPageSideImage(page WidgetLike) *gdkpixbuf.GdkPixbuf {
	return &gdkpixbuf.GdkPixbuf{
		C.gtk_assistant_get_page_side_image(C.to_GtkAssistant(v.GWidget), page.ToNative())}
}
func (v *Assistant) SetPageComplete(page WidgetLike, complete bool) {
	C.gtk_assistant_set_page_complete(C.to_GtkAssistant(v.GWidget), page.ToNative(), bool2gboolean(complete))
}
func (v *Assistant) GetPageComplete(page WidgetLike) bool {
	return gboolean2bool(C.gtk_assistant_get_page_complete(C.to_GtkAssistant(v.GWidget), page.ToNative()))
}
func (v *Assistant) AddActionWidget(child WidgetLike) {
	C.gtk_assistant_add_action_widget(C.to_GtkAssistant(v.GWidget), child.ToNative())
}
func (v *Assistant) RemoveActionWidget(child WidgetLike) {
	C.gtk_assistant_remove_action_widget(C.to_GtkAssistant(v.GWidget), child.ToNative())
}
func (v *Assistant) UpdateButtonsState() {
	C.gtk_assistant_update_buttons_state(C.to_GtkAssistant(v.GWidget))
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
/*type AccelLabelLike interface {
	WidgetLike
	GetAccelWidget() GtkWidget
	SetAccelWidget(GtkWidget)
}*/
type AccelLabel struct {
	Widget
}

func NewAccelLabel(label string) *AccelLabel {
	ptr := C.CString(label)
	defer C.free_string(ptr)
	return &AccelLabel{Widget{C.gtk_accel_label_new(C.to_gcharptr(ptr))}}
}

// gtk_accel_label_set_accel_closure

func (v *AccelLabel) GetAccelWidget() Widget {
	return Widget{C.gtk_accel_label_get_accel_widget(C.to_GtkAccelLabel(v.GWidget))}
}
func (v *AccelLabel) SetAccelWidget(w WidgetLike) {
	C.gtk_accel_label_set_accel_widget(C.to_GtkAccelLabel(v.GWidget), w.ToNative())
}
func (v *AccelLabel) GetAccelWidth() uint {
	return uint(C.gtk_accel_label_get_accel_width(C.to_GtkAccelLabel(v.GWidget)))
}
func (v *AccelLabel) Refetch() bool {
	return gboolean2bool(C.gtk_accel_label_refetch(C.to_GtkAccelLabel(v.GWidget)))
}

//-----------------------------------------------------------------------
// GtkImage
//-----------------------------------------------------------------------
type IconSize int

const (
	ICON_SIZE_INVALID       IconSize = 0
	ICON_SIZE_MENU          IconSize = 1
	ICON_SIZE_SMALL_TOOLBAR IconSize = 2
	ICON_SIZE_LARGE_TOOLBAR IconSize = 3
	ICON_SIZE_BUTTON        IconSize = 4
	ICON_SIZE_DND           IconSize = 5
	ICON_SIZE_DIALOG        IconSize = 6
)

/*type ImageLike interface {
	WidgetLike
}*/
type Image struct {
	Widget
}

func NewImage() *Image {
	return &Image{Widget{C.gtk_image_new()}}
}

func NewImageFromFile(filename string) *Image {
	ptr := C.CString(filename)
	defer C.free_string(ptr)
	return &Image{Widget{C.gtk_image_new_from_file(C.to_gcharptr(ptr))}}
}

// gtk_image_new_from_icon_set
// gtk_image_new_from_image

func NewImageFromPixbuf(pixbuf gdkpixbuf.GdkPixbuf) *Image {
	return &Image{Widget{
		C.gtk_image_new_from_pixbuf(pixbuf.Pixbuf)}}
}

// gtk_image_new_from_pixmap

func NewImageFromStock(stock_id string, size IconSize) *Image {
	ptr := C.CString(stock_id)
	defer C.free_string(ptr)
	return &Image{Widget{C.gtk_image_new_from_stock(C.to_gcharptr(ptr), C.GtkIconSize(size))}}
}

// gtk_image_new_from_animation
// gtk_image_new_from_icon_name
// gtk_image_new_from_gicon

func (v *Image) GetPixbuf() *gdkpixbuf.GdkPixbuf {
	return &gdkpixbuf.GdkPixbuf{
		C.gtk_image_get_pixbuf(C.to_GtkImage(v.GWidget))}
}

func (v *Image) SetFromFile(filename string) {
	ptr := C.CString(filename)
	defer C.free_string(ptr)
	C.gtk_image_set_from_file(C.to_GtkImage(v.GWidget), C.to_gcharptr(ptr))
}

// gtk_image_set_from_icon_set
// gtk_image_set_from_image

func (v *Image) SetFromPixbuf(pixbuf *gdkpixbuf.GdkPixbuf) {
	C.gtk_image_set_from_pixbuf(C.to_GtkImage(v.GWidget), pixbuf.Pixbuf)
}

// gtk_image_set_from_pixmap

func (v *Image) SetFromStock(stock_id string, size IconSize) {
	ptr := C.CString(stock_id)
	defer C.free_string(ptr)
	C.gtk_image_set_from_stock(C.to_GtkImage(v.GWidget), C.to_gcharptr(ptr), C.GtkIconSize(size))
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
	C.gtk_image_clear(C.to_GtkImage(v.GWidget))
}

// gtk_image_set_pixel_size
// gtk_image_get_pixel_size

//-----------------------------------------------------------------------
// GtkLabel
//-----------------------------------------------------------------------
type Justification int

const (
	JUSTIFY_LEFT   Justification = 0
	JUSTIFY_RIGHT  Justification = 1
	JUSTIFY_CENTER Justification = 2
	JUSTIFY_FILL   Justification = 3
)

type LabelLike interface {
	WidgetLike
	isLabelLike()
	GetLabel() string
	SetLabel(label string)
}
type Label struct {
	Widget
}

func (Label) isLabelLike() {}

func NewLabel(label string) *Label {
	ptr := C.CString(label)
	defer C.free_string(ptr)
	return &Label{Widget{C.gtk_label_new(C.to_gcharptr(ptr))}}
}
func (v *Label) SetText(label string) {
	ptr := C.CString(label)
	defer C.free_string(ptr)
	C.gtk_label_set_text(C.to_GtkLabel(v.GWidget), C.to_gcharptr(ptr))
}

func (v *Label) SetMnemonicWidget(widget WidgetLike) {
	C.gtk_label_set_mnemonic_widget(C.to_GtkLabel(v.GWidget), widget.ToNative())
}

func (v *Label) SetMarkup(markup string) {
	ptr := C.CString(markup)
	defer C.free_string(ptr)
	C.gtk_label_set_markup(C.to_GtkLabel(v.GWidget), C.to_gcharptr(ptr))
}

func (v *Label) GetMnemonicWidget() *Widget {
	return &Widget{C.gtk_label_get_mnemonic_widget(C.to_GtkLabel(v.GWidget))}
}

func (v *Label) SetPattern(pattern string) {
	ptr := C.CString(pattern)
	defer C.free_string(ptr)
	C.gtk_label_set_pattern(C.to_GtkLabel(v.GWidget), C.to_gcharptr(ptr))
}
func (v *Label) SetJustify(jtype Justification) {
	C.gtk_label_set_justify(C.to_GtkLabel(v.GWidget), C.GtkJustification(jtype))
}
func (v *Label) SetEllipsize(ellipsize pango.PangoEllipsizeMode) {
	C.gtk_label_set_ellipsize(C.to_GtkLabel(v.GWidget), C.PangoEllipsizeMode(ellipsize))
}
func (v *Label) SetWidthChars(n_chars int) {
	C.gtk_label_set_width_chars(C.to_GtkLabel(v.GWidget), C.gint(n_chars))
}
func (v *Label) SetMaxWidthChars(n_chars int) {
	C.gtk_label_set_max_width_chars(C.to_GtkLabel(v.GWidget), C.gint(n_chars))
}
func (v *Label) SetLineWrap(setting bool) {
	C.gtk_label_set_line_wrap(C.to_GtkLabel(v.GWidget), bool2gboolean(setting))
}
func (v *Label) SetUseLineWrapMode(wrap_mode pango.PangoWrapMode) {
	C.gtk_label_set_line_wrap_mode(C.to_GtkLabel(v.GWidget), C.PangoWrapMode(wrap_mode))
}

// gtk_label_get_layout_offsets
// gtk_label_get_mnemonic_keyval

func (v *Label) GetSelectable() bool {
	return gboolean2bool(C.gtk_label_get_selectable(C.to_GtkLabel(v.GWidget)))
}
func (v *Label) GetText() string {
	return C.GoString(C.to_charptr(C.gtk_label_get_text(C.to_GtkLabel(v.GWidget))))
}
func LabelWithMnemonic(label string) *Label {
	ptr := C.CString(label)
	defer C.free_string(ptr)
	return &Label{Widget{C.gtk_label_new_with_mnemonic(C.to_gcharptr(ptr))}}
}
func (v *Label) SelectRegion(start_offset int, end_offset int) {
	C.gtk_label_select_region(C.to_GtkLabel(v.GWidget), C.gint(start_offset), C.gint(end_offset))
}

// gtk_label_set_mnemonic_widget

func (v *Label) SetSelectable(setting bool) {
	C.gtk_label_set_selectable(C.to_GtkLabel(v.GWidget), bool2gboolean(setting))
}
func (v *Label) SetTextWithMnemonic(str string) {
	ptr := C.CString(str)
	defer C.free_string(ptr)
	C.gtk_label_set_text_with_mnemonic(C.to_GtkLabel(v.GWidget), C.to_gcharptr(ptr))
}

// gtk_label_get_attributes
func (v *Label) GetJustify() Justification {
	return Justification(C.gtk_label_get_justify(C.to_GtkLabel(v.GWidget)))
}
func (v *Label) GetEllipsize() pango.PangoEllipsizeMode {
	return pango.PangoEllipsizeMode(C.gtk_label_get_ellipsize(C.to_GtkLabel(v.GWidget)))
}
func (v *Label) GetWidthChars() int {
	return int(C.gtk_label_get_width_chars(C.to_GtkLabel(v.GWidget)))
}
func (v *Label) GetMaxWidthChars() int {
	return int(C.gtk_label_get_max_width_chars(C.to_GtkLabel(v.GWidget)))
}
func (v *Label) GetLabel() string {
	return C.GoString(C.to_charptr(C.gtk_label_get_label(C.to_GtkLabel(v.GWidget))))
}

// gtk_label_get_layout

func (v *Label) GetLineWrap() bool {
	return gboolean2bool(C.gtk_label_get_line_wrap(C.to_GtkLabel(v.GWidget)))
}
func (v *Label) GetLineWrapMode() pango.PangoWrapMode {
	return pango.PangoWrapMode(C.gtk_label_get_line_wrap_mode(C.to_GtkLabel(v.GWidget)))
}

// gtk_label_get_mnemonic_widget
func (v *Label) GetSelectionBounds(start *int, end *int) {
	var cstart, cend C.gint
	C.gtk_label_get_selection_bounds(C.to_GtkLabel(v.GWidget), &cstart, &cend)
	*start = int(cstart)
	*end = int(cend)
}
func (v *Label) GetUseMarkup() bool {
	return gboolean2bool(C.gtk_label_get_use_markup(C.to_GtkLabel(v.GWidget)))
}
func (v *Label) GetUseUnderline() bool {
	return gboolean2bool(C.gtk_label_get_use_underline(C.to_GtkLabel(v.GWidget)))
}
func (v *Label) GetSingleLineMode() bool {
	return gboolean2bool(C.gtk_label_get_single_line_mode(C.to_GtkLabel(v.GWidget)))
}
func (v *Label) GetAngle() float64 {
	r := C.gtk_label_get_angle(C.to_GtkLabel(v.GWidget))
	return float64(r)
}
func (v *Label) SetLabel(label string) {
	ptr := C.CString(label)
	defer C.free_string(ptr)
	C.gtk_label_set_label(C.to_GtkLabel(v.GWidget), C.to_gcharptr(ptr))
}
func (v *Label) SetUseMarkup(setting bool) {
	C.gtk_label_set_use_markup(C.to_GtkLabel(v.GWidget), bool2gboolean(setting))
}
func (v *Label) SetUseUnderline(setting bool) {
	C.gtk_label_set_use_underline(C.to_GtkLabel(v.GWidget), bool2gboolean(setting))
}
func (v *Label) SetSingleLineMode(single_line bool) {
	C.gtk_label_set_single_line_mode(C.to_GtkLabel(v.GWidget), bool2gboolean(single_line))
}
func (v *Label) SetAngle(angle float64) {
	C.gtk_label_set_angle(C.to_GtkLabel(v.GWidget), C.gdouble(angle))
}
func (v *Label) GetCurrentUri() string {
	panic_if_version_older(2, 18, 0, "gtk_label_get_current_uri()")
	return C.GoString(C.to_charptr(C.gtk_label_get_current_uri(C.to_GtkLabel(v.GWidget))))
}
func (v *Label) SetTrackVisitedLinks(track_links bool) {
	panic_if_version_older(2, 18, 0, "gtk_label_set_track_visited_links()")
	C.gtk_label_set_track_visited_links(C.to_GtkLabel(v.GWidget), bool2gboolean(track_links))
}
func (v *Label) GetTrackVisitedLinks() bool {
	panic_if_version_older(2, 18, 0, "gtk_label_get_track_visited_links()")
	return gboolean2bool(C.gtk_label_get_track_visited_links(C.to_GtkLabel(v.GWidget)))
}

//-----------------------------------------------------------------------
// GtkProgressBar
//-----------------------------------------------------------------------
type ProgressBarOrientation int

const (
	PROGRESS_LEFT_TO_RIGHT ProgressBarOrientation = 0
	PROGRESS_RIGHT_TO_LEFT ProgressBarOrientation = 1
	PROGRESS_BOTTOM_TO_TOP ProgressBarOrientation = 2
	PROGRESS_TOP_TO_BOTTOM ProgressBarOrientation = 3
)

type ProgressBar struct {
	Widget
}

func NewProgressBar() *ProgressBar {
	return &ProgressBar{Widget{C.gtk_progress_bar_new()}}
}
func (v *ProgressBar) Pulse() {
	C.gtk_progress_bar_pulse(C.to_GtkProgressBar(v.GWidget))
}
func (v *ProgressBar) SetText(show_text string) {
	ptr := C.CString(show_text)
	defer C.free_string(ptr)
	C.gtk_progress_bar_set_text(C.to_GtkProgressBar(v.GWidget), C.to_gcharptr(ptr))
}
func (v *ProgressBar) SetFraction(fraction float64) {
	C.gtk_progress_bar_set_fraction(C.to_GtkProgressBar(v.GWidget), C.gdouble(fraction))
}
func (v *ProgressBar) SetPulseStep(fraction float64) {
	C.gtk_progress_bar_set_pulse_step(C.to_GtkProgressBar(v.GWidget), C.gdouble(fraction))
}
func (v *ProgressBar) SetOrientation(i ProgressBarOrientation) {
	C.gtk_progress_bar_set_orientation(C.to_GtkProgressBar(v.GWidget), C.GtkProgressBarOrientation(i))
}

// gtk_progress_bar_set_ellipsize

func (v *ProgressBar) GetText() string {
	return C.GoString(C.to_charptr(C.gtk_progress_bar_get_text(C.to_GtkProgressBar(v.GWidget))))
}
func (v *ProgressBar) GetFraction() float64 {
	r := C.gtk_progress_bar_get_fraction(C.to_GtkProgressBar(v.GWidget))
	return float64(r)
}
func (v *ProgressBar) GetPulseStep() float64 {
	r := C.gtk_progress_bar_get_pulse_step(C.to_GtkProgressBar(v.GWidget))
	return float64(r)
}
func (v *ProgressBar) GetOrientation() ProgressBarOrientation {
	return ProgressBarOrientation(C.gtk_progress_bar_get_orientation(C.to_GtkProgressBar(v.GWidget)))
}

// gtk_progress_bar_get_ellipsize

//-----------------------------------------------------------------------
// GtkStatusbar
//-----------------------------------------------------------------------
type Statusbar struct {
	HBox
}

func NewStatusbar() *Statusbar {
	return &Statusbar{HBox{Box{Container{Widget{
		C.gtk_statusbar_new()}}}}}
}
func (v *Statusbar) GetContextId(content_description string) uint {
	ptr := C.CString(content_description)
	defer C.free_string(ptr)
	return uint(C.gtk_statusbar_get_context_id(C.to_GtkStatusbar(v.GWidget), C.to_gcharptr(ptr)))
}
func (v *Statusbar) Push(context_id uint, text string) {
	ptr := C.CString(text)
	defer C.free_string(ptr)
	C.gtk_statusbar_push(C.to_GtkStatusbar(v.GWidget), C.guint(context_id), C.to_gcharptr(ptr))
}
func (v *Statusbar) Pop(context_id uint) {
	C.gtk_statusbar_pop(C.to_GtkStatusbar(v.GWidget), C.guint(context_id))
}
func (v *Statusbar) Remove(context_id uint, message_id uint) {
	C.gtk_statusbar_remove(C.to_GtkStatusbar(v.GWidget), C.guint(context_id), C.guint(message_id))
}

// gtk_statusbar_remove_all //since 2.22

func (v *Statusbar) SetHasResizeGrip(add_tearoffs bool) {
	C.gtk_statusbar_set_has_resize_grip(C.to_GtkStatusbar(v.GWidget), bool2gboolean(add_tearoffs))
}
func (v *Statusbar) GetHasResizeGrip() bool {
	return gboolean2bool(C.gtk_statusbar_get_has_resize_grip(C.to_GtkStatusbar(v.GWidget)))
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
	return &InfoBar{HBox{Box{Container{Widget{
		C._gtk_info_bar_new()}}}}}
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

func (v *InfoBar) AddActionWidget(child WidgetLike, responseId int) {
	panic_if_version_older_auto(2, 18, 0)
	C._gtk_info_bar_add_action_widget(C.to_GtkInfoBar(v.GWidget),
		child.ToNative(), C.gint(responseId))
}

func (v *InfoBar) AddButton(buttonText string, responseId int) *Widget {
	panic_if_version_older_auto(2, 18, 0)
	ptr := C.CString(buttonText)
	defer C.free_string(ptr)
	return &Widget{C._gtk_info_bar_add_button(C.to_GtkInfoBar(v.GWidget),
		C.to_gcharptr(ptr), C.gint(responseId))}
}

func (v *InfoBar) AddButtons(buttons ...interface{}) {
	panic_if_version_older_auto(2, 18, 0)
	text, res := variadicButtonsToArrays(buttons)
	for i := range text {
		v.AddButton(text[i], res[i])
	}
}

func (v *InfoBar) SetResponseSensitive(responseId int, setting bool) {
	panic_if_version_older_auto(2, 18, 0)
	C._gtk_info_bar_set_response_sensitive(C.to_GtkInfoBar(v.GWidget),
		C.gint(responseId), bool2gboolean(setting))
}

func (v *InfoBar) SetDefaultResponse(responseId int) {
	panic_if_version_older_auto(2, 18, 0)
	C._gtk_info_bar_set_default_response(C.to_GtkInfoBar(v.GWidget),
		C.gint(responseId))
}

func (v *InfoBar) Response(responseId int) {
	panic_if_version_older_auto(2, 18, 0)
	C._gtk_info_bar_response(C.to_GtkInfoBar(v.GWidget), C.gint(responseId))
}

func (v *InfoBar) SetMessageType(messageType MessageType) {
	panic_if_version_older_auto(2, 18, 0)
	C._gtk_info_bar_set_message_type(C.to_GtkInfoBar(v.GWidget),
		C.GtkMessageType(messageType))
}

func (v *InfoBar) GetMessageType() MessageType {
	panic_if_version_older_auto(2, 18, 0)
	return MessageType(C._gtk_info_bar_get_message_type(C.to_GtkInfoBar(v.GWidget)))
}

func (v *InfoBar) GetActionArea() *Widget {
	panic_if_version_older_auto(2, 18, 0)
	return &Widget{C._gtk_info_bar_get_action_area(C.to_GtkInfoBar(v.GWidget))}
}

func (v *InfoBar) GetContentArea() *Widget {
	panic_if_version_older_auto(2, 18, 0)
	return &Widget{C._gtk_info_bar_get_content_area(C.to_GtkInfoBar(v.GWidget))}
}

//-----------------------------------------------------------------------
// GtkStatusIcon
//-----------------------------------------------------------------------
type StatusIcon struct {
	GStatusIcon *C.GtkStatusIcon
}

func NewStatusIcon() *StatusIcon {
	return &StatusIcon{
		C.gtk_status_icon_new()}
}
func NewStatusIconFromPixbuf(pixbuf *gdkpixbuf.GdkPixbuf) *StatusIcon {
	return &StatusIcon{
		C.gtk_status_icon_new_from_pixbuf(pixbuf.Pixbuf)}
}
func NewStatusIconFromFile(filename string) *StatusIcon {
	ptr := C.CString(filename)
	defer C.free_string(ptr)
	return &StatusIcon{
		C.gtk_status_icon_new_from_file(C.to_gcharptr(ptr))}
}
func NewStatusIconFromStock(stock_id string) *StatusIcon {
	ptr := C.CString(stock_id)
	defer C.free_string(ptr)
	return &StatusIcon{
		C.gtk_status_icon_new_from_stock(C.to_gcharptr(ptr))}
}
func NewStatusIconFromIconName(icon_name string) *StatusIcon {
	ptr := C.CString(icon_name)
	defer C.free_string(ptr)
	return &StatusIcon{
		C.gtk_status_icon_new_from_icon_name(C.to_gcharptr(ptr))}
}

//GtkStatusIcon *gtk_status_icon_new_from_gicon(GIcon *icon);

func (v *StatusIcon) SetFromPixbuf(pixbuf *gdkpixbuf.GdkPixbuf) {
	C.gtk_status_icon_set_from_pixbuf(v.GStatusIcon, pixbuf.Pixbuf)
}
func (v *StatusIcon) SetFromFile(filename string) {
	ptr := C.CString(filename)
	defer C.free_string(ptr)
	C.gtk_status_icon_set_from_file(v.GStatusIcon, C.to_gcharptr(ptr))
}
func (v *StatusIcon) SetFromStock(stock_id string) {
	ptr := C.CString(stock_id)
	defer C.free_string(ptr)
	C.gtk_status_icon_set_from_stock(v.GStatusIcon, C.to_gcharptr(ptr))
}
func (v *StatusIcon) SetFromIconName(icon_name string) {
	ptr := C.CString(icon_name)
	defer C.free_string(ptr)
	C.gtk_status_icon_set_from_icon_name(v.GStatusIcon, C.to_gcharptr(ptr))
}

//void gtk_status_icon_set_from_gicon (GtkStatusIcon *status_icon, GIcon *icon);
//GtkImageType gtk_status_icon_get_storage_type (GtkStatusIcon *status_icon);

func (v *StatusIcon) GetPixbuf() *gdkpixbuf.GdkPixbuf {
	return &gdkpixbuf.GdkPixbuf{
		C.gtk_status_icon_get_pixbuf(v.GStatusIcon)}
}
func (v *StatusIcon) GetStock() string {
	return C.GoString(C.to_charptr(C.gtk_status_icon_get_stock(v.GStatusIcon)))
}
func (v *StatusIcon) GetIconName() string {
	return C.GoString(C.to_charptr(C.gtk_status_icon_get_icon_name(v.GStatusIcon)))
}
func (v *StatusIcon) SetName(name string) {
	panic_if_version_older(2, 20, 0, "gtk_status_icon_set_name()")
	ptr := C.CString(name)
	defer C.free_string(ptr)
	C._gtk_status_icon_set_name(v.GStatusIcon, C.to_gcharptr(ptr))
}
func (v *StatusIcon) SetTitle(title string) {
	panic_if_version_older(2, 18, 0, "gtk_status_icon_set_title()")
	ptr := C.CString(title)
	defer C.free_string(ptr)
	C._gtk_status_icon_set_title(v.GStatusIcon, C.to_gcharptr(ptr))
}
func (v *StatusIcon) GetTitle() string {
	panic_if_version_older(2, 18, 0, "gtk_status_icon_get_title()")
	return C.GoString(C.to_charptr(C._gtk_status_icon_get_title(v.GStatusIcon)))
}
func (v *StatusIcon) SetTooltipText(text string) {
	ptr := C.CString(text)
	defer C.free_string(ptr)
	C.gtk_status_icon_set_tooltip_text(v.GStatusIcon, C.to_gcharptr(ptr))
}
func (v *StatusIcon) GetTooltipText() string {
	return C.GoString(C.to_charptr(C.gtk_status_icon_get_tooltip_text(v.GStatusIcon)))
}
func (v *StatusIcon) SetTooltipMarkup(markup string) {
	ptr := C.CString(markup)
	defer C.free_string(ptr)
	C.gtk_status_icon_set_tooltip_markup(v.GStatusIcon, C.to_gcharptr(ptr))
}
func (v *StatusIcon) GetTooltipMarkup() string {
	return C.GoString(C.to_charptr(C.gtk_status_icon_get_tooltip_markup(v.GStatusIcon)))
}
func (v *StatusIcon) GetHasTooltip() bool {
	return gboolean2bool(C.gtk_status_icon_get_has_tooltip(v.GStatusIcon))
}
func (v *StatusIcon) SetHasTooltip(setting bool) {
	C.gtk_status_icon_set_has_tooltip(v.GStatusIcon, bool2gboolean(setting))
}
func (v *StatusIcon) GetVisible() bool {
	return gboolean2bool(C.gtk_status_icon_get_visible(v.GStatusIcon))
}
func (v *StatusIcon) SetVisible(setting bool) {
	C.gtk_status_icon_set_visible(v.GStatusIcon, bool2gboolean(setting))
}
func StatusIconPositionMenu(menu *Menu, px, py *int, push_in *bool, data interface{}) {
	x := C.gint(*px)
	y := C.gint(*py)
	pi := bool2gboolean(*push_in)
	var pdata C.gpointer
	if sm, ok := data.(*StatusIcon); ok {
		pdata = C.gpointer(unsafe.Pointer(sm.GStatusIcon))
	}
	C.gtk_status_icon_position_menu(C.to_GtkMenu(menu.GWidget), &x, &y, &pi, pdata)
	*px = int(x)
	*py = int(y)
	*push_in = gboolean2bool(pi)
}
func (v *StatusIcon) Connect(s string, f interface{}, datas ...interface{}) int {
	return glib.ObjectFromNative(unsafe.Pointer(C.to_GObject(unsafe.Pointer(v.GStatusIcon)))).Connect(s, f, datas...)
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

// gtk_spinner_new //since 2.20
// gtk_spinner_start //since 2.20
// gtk_spinner_stop //since 2.20

//-----------------------------------------------------------------------
// GtkButton
//-----------------------------------------------------------------------
/*type ButtonLike interface { // Buttons are LabelLike Widgets!
	LabelLike
	// the following should be just Clickable; ...
	Clicked(interface{}, ...interface{}) // this is a very simple interface...
}*/
/*type Clickable interface {
	WidgetLike
	Clicked(interface{}, ...interface{}) // this is a very simple interface...
}*/

type ReliefStyle int

const (
	RELIEF_NORMAL ReliefStyle = 0
	RELIEF_HALF   ReliefStyle = 1
	RELIEF_NONE   ReliefStyle = 2
)

type Button struct {
	Bin
}

func (Button) isLabelLike() {} // TODO

func NewButton() *Button {
	return &Button{Bin{Container{Widget{C.gtk_button_new()}}}}
}
func NewButtonWithLabel(label string) *Button {
	ptr := C.CString(label)
	defer C.free_string(ptr)
	return &Button{Bin{Container{Widget{
		C.gtk_button_new_with_label(C.to_gcharptr(ptr))}}}}
}
func NewButtonWithMnemonic(label string) *Button {
	ptr := C.CString(label)
	defer C.free_string(ptr)
	return &Button{Bin{Container{Widget{
		C.gtk_button_new_with_mnemonic(C.to_gcharptr(ptr))}}}}
}

// gtk_button_new_from_stock

func (v *Button) Pressed() {
	deprecated_since(2, 20, 0, "gtk_button_pressed()")
	C.gtk_button_pressed(C.to_GtkButton(v.GWidget))
}

func (v *Button) Released() {
	deprecated_since(2, 20, 0, "gtk_button_released()")
	C.gtk_button_released(C.to_GtkButton(v.GWidget))
}

func (v *Button) Clicked(onclick interface{}, datas ...interface{}) int {
	return v.Connect("clicked", onclick, datas...)
}

func (v *Button) Enter() {
	deprecated_since(2, 20, 0, "gtk_button_enter()")
	C.gtk_button_enter(C.to_GtkButton(v.GWidget))
}

func (v *Button) Leave() {
	deprecated_since(2, 20, 0, "gtk_button_leave()")
	C.gtk_button_leave(C.to_GtkButton(v.GWidget))
}

func (v *Button) GetRelief() ReliefStyle {
	return ReliefStyle(C.gtk_button_get_relief(C.to_GtkButton(v.GWidget)))
}
func (v *Button) SetRelief(relief ReliefStyle) {
	C.gtk_button_set_relief(C.to_GtkButton(v.GWidget), C.GtkReliefStyle(relief))
}
func (v *Button) GetLabel() string {
	return C.GoString(C.to_charptr(C.gtk_button_get_label(C.to_GtkButton(v.GWidget))))
}
func (v *Button) SetLabel(label string) {
	ptr := C.CString(label)
	defer C.free_string(ptr)
	C.gtk_button_set_label(C.to_GtkButton(v.GWidget), C.to_gcharptr(ptr))
}
func (v *Button) GetUseStock() bool {
	return gboolean2bool(C.gtk_button_get_use_stock(C.to_GtkButton(v.GWidget)))
}
func (v *Button) SetUseStock(use bool) {
	C.gtk_button_set_use_stock(C.to_GtkButton(v.GWidget), bool2gboolean(use))
}
func (v *Button) GetUseUnderline() bool {
	return gboolean2bool(C.gtk_button_get_use_underline(C.to_GtkButton(v.GWidget)))
}
func (v *Button) SetUseUnderline(setting bool) {
	C.gtk_button_set_use_underline(C.to_GtkButton(v.GWidget), bool2gboolean(setting))
}
func (v *Button) GetFocusOnClick() bool {
	return gboolean2bool(C.gtk_button_get_focus_on_click(C.to_GtkButton(v.GWidget)))
}
func (v *Button) SetFocusOnClick(setting bool) {
	C.gtk_button_set_focus_on_click(C.to_GtkButton(v.GWidget), bool2gboolean(setting))
}
func (v *Button) SetAlignment(xalign, yalign float64) {
	C.gtk_button_set_alignment(C.to_GtkButton(v.GWidget), C.gfloat(xalign), C.gfloat(yalign))
}
func (v *Button) GetAlignment() (xalign, yalign float64) {
	var xalign_, yalign_ C.gfloat
	C.gtk_button_get_alignment(C.to_GtkButton(v.GWidget), &xalign_, &yalign_)
	return float64(xalign_), float64(yalign_)
}
func (v *Button) SetImage(image WidgetLike) {
	C.gtk_button_set_image(C.to_GtkButton(v.GWidget), image.ToNative())
}
func (v *Button) GetImage() *Image {
	return &Image{Widget{C.gtk_button_get_image(C.to_GtkButton(v.GWidget))}}
}

// gtk_button_set_image_position
// gtk_button_get_image_position
// gtk_button_get_event_window

//-----------------------------------------------------------------------
// GtkCheckButton
//-----------------------------------------------------------------------
type CheckButton struct {
	ToggleButton
}

func NewCheckButton() *CheckButton {
	return &CheckButton{ToggleButton{Button{Bin{Container{Widget{
		C.gtk_check_button_new()}}}}}}
}
func NewCheckButtonWithLabel(label string) *CheckButton {
	ptr := C.CString(label)
	defer C.free_string(ptr)
	return &CheckButton{ToggleButton{Button{Bin{Container{Widget{
		C.gtk_check_button_new_with_label(C.to_gcharptr(ptr))}}}}}}
}
func NewCheckButtonWithMnemonic(label string) *CheckButton {
	ptr := C.CString(label)
	defer C.free_string(ptr)
	return &CheckButton{ToggleButton{Button{Bin{Container{Widget{
		C.gtk_check_button_new_with_mnemonic(C.to_gcharptr(ptr))}}}}}}
}

//-----------------------------------------------------------------------
// GtkRadioButton
//-----------------------------------------------------------------------
type RadioButton struct {
	CheckButton
}

func NewRadioButton(group *glib.SList) *RadioButton {
	if group != nil {
		return &RadioButton{CheckButton{ToggleButton{Button{Bin{Container{Widget{
			C.gtk_radio_button_new(C.to_gslist(unsafe.Pointer(group.ToSList())))}}}}}}}
	}
	return &RadioButton{CheckButton{ToggleButton{Button{Bin{Container{Widget{
		C.gtk_radio_button_new(nil)}}}}}}}
}
func NewRadioButtonFromWidget(w *RadioButton) *RadioButton {
	return &RadioButton{CheckButton{ToggleButton{Button{Bin{Container{Widget{
		C.gtk_radio_button_new_from_widget(C.to_GtkRadioButton(w.GWidget))}}}}}}}
}
func NewRadioButtonWithLabel(group *glib.SList, label string) *RadioButton {
	ptr := C.CString(label)
	defer C.free_string(ptr)
	if group != nil {
		return &RadioButton{CheckButton{ToggleButton{Button{Bin{Container{Widget{
			C.gtk_radio_button_new_with_label(C.to_gslist(unsafe.Pointer(group.ToSList())), C.to_gcharptr(ptr))}}}}}}}
	}
	return &RadioButton{CheckButton{ToggleButton{Button{Bin{Container{Widget{
		C.gtk_radio_button_new_with_label(nil, C.to_gcharptr(ptr))}}}}}}}
}
func NewRadioButtonWithLabelFromWidget(w *RadioButton, label string) *RadioButton {
	ptr := C.CString(label)
	defer C.free_string(ptr)
	return &RadioButton{CheckButton{ToggleButton{Button{Bin{Container{Widget{
		C.gtk_radio_button_new_with_label_from_widget(C.to_GtkRadioButton(w.GWidget), C.to_gcharptr(ptr))}}}}}}}
}
func NewRadioButtonWithMnemonic(group *glib.SList, label string) *RadioButton {
	ptr := C.CString(label)
	defer C.free_string(ptr)
	if group != nil {
		return &RadioButton{CheckButton{ToggleButton{Button{Bin{Container{Widget{
			C.gtk_radio_button_new_with_mnemonic(C.to_gslist(unsafe.Pointer(group.ToSList())), C.to_gcharptr(ptr))}}}}}}}
	}
	return &RadioButton{CheckButton{ToggleButton{Button{Bin{Container{Widget{
		C.gtk_radio_button_new_with_mnemonic(nil, C.to_gcharptr(ptr))}}}}}}}
}
func NewRadioButtonWithMnemonicFromWidget(w *RadioButton, label string) *RadioButton {
	ptr := C.CString(label)
	defer C.free_string(ptr)
	return &RadioButton{CheckButton{ToggleButton{Button{Bin{Container{Widget{
		C.gtk_radio_button_new_with_mnemonic_from_widget(C.to_GtkRadioButton(w.GWidget), C.to_gcharptr(ptr))}}}}}}}
}
func (v *RadioButton) GetGroup() *glib.SList {
	return glib.SListFromNative(unsafe.Pointer(C.gtk_radio_button_get_group(C.to_GtkRadioButton(v.GWidget))))
}
func (v *RadioButton) SetGroup(group *glib.SList) {
	if group != nil {
		C.gtk_radio_button_set_group(C.to_GtkRadioButton(v.GWidget), C.to_gslist(unsafe.Pointer(group)))
	} else {
		C.gtk_radio_button_set_group(C.to_GtkRadioButton(v.GWidget), nil)
	}
}

//-----------------------------------------------------------------------
// GtkToggleButton
//-----------------------------------------------------------------------
type ToggleButton struct {
	Button
}

func NewToggleButton() *ToggleButton {
	return &ToggleButton{Button{Bin{Container{Widget{
		C.gtk_toggle_button_new()}}}}}
}
func NewToggleButtonWithLabel(label string) *ToggleButton {
	ptr := C.CString(label)
	defer C.free_string(ptr)
	return &ToggleButton{Button{Bin{Container{Widget{
		C.gtk_toggle_button_new_with_label(C.to_gcharptr(ptr))}}}}}
}
func NewToggleButtonWithMnemonic(label string) *ToggleButton {
	ptr := C.CString(label)
	defer C.free_string(ptr)
	return &ToggleButton{Button{Bin{Container{Widget{
		C.gtk_toggle_button_new_with_mnemonic(C.to_gcharptr(ptr))}}}}}
}
func (v *ToggleButton) SetMode(draw_indicator bool) {
	C.gtk_toggle_button_set_mode(C.to_GtkToggleButton(v.GWidget), bool2gboolean(draw_indicator))
}
func (v *ToggleButton) GetMode() bool {
	return gboolean2bool(C.gtk_toggle_button_get_mode(C.to_GtkToggleButton(v.GWidget)))
}
func (v *ToggleButton) GetActive() bool {
	return gboolean2bool(C.gtk_toggle_button_get_active(C.to_GtkToggleButton(v.GWidget)))
}
func (v *ToggleButton) SetActive(is_active bool) {
	C.gtk_toggle_button_set_active(C.to_GtkToggleButton(v.GWidget), bool2gboolean(is_active))
}
func (v *ToggleButton) GetInconsistent() bool {
	return gboolean2bool(C.gtk_toggle_button_get_inconsistent(C.to_GtkToggleButton(v.GWidget)))
}
func (v *ToggleButton) SetInconsistent(setting bool) {
	C.gtk_toggle_button_set_inconsistent(C.to_GtkToggleButton(v.GWidget), bool2gboolean(setting))
}

//-----------------------------------------------------------------------
// GtkLinkButton
//-----------------------------------------------------------------------
type LinkButton struct {
	Button
}

func NewLinkButton(uri string) *LinkButton {
	ptr := C.CString(uri)
	defer C.free_string(ptr)
	return &LinkButton{Button{Bin{Container{Widget{
		C.gtk_link_button_new(C.to_gcharptr(ptr))}}}}}
}
func NewLinkButtonWithLabel(uri string, label string) *LinkButton {
	puri := C.CString(uri)
	defer C.free_string(puri)
	plabel := C.CString(label)
	defer C.free_string(plabel)
	return &LinkButton{Button{Bin{Container{Widget{
		C.gtk_link_button_new_with_label(C.to_gcharptr(puri), C.to_gcharptr(plabel))}}}}}
}
func (v *LinkButton) GetUri() string {
	return C.GoString(C.to_charptr(C.gtk_link_button_get_uri(C.to_GtkLinkButton(v.GWidget))))
}
func (v *LinkButton) SetUri(uri string) {
	ptr := C.CString(uri)
	defer C.free_string(ptr)
	C.gtk_link_button_set_uri(C.to_GtkLinkButton(v.GWidget), C.to_gcharptr(ptr))
}

//gtk_link_button_set_uri_hook has been deprecated since 2.24. Use clicked signal instead. //TODO
//func (v GtkLinkButton) SetUriHook(f func(button *GtkLinkButton, link string, user_data unsafe.Pointer), ) {
// GtkLinkButtonUriFunc gtk_link_button_set_uri_hook (GtkLinkButtonUriFunc func, gpointer data, GDestroyNotify destroy);
//}
func (v *LinkButton) GetVisited() bool {
	return gboolean2bool(C.gtk_link_button_get_visited(C.to_GtkLinkButton(v.GWidget)))
}
func (v *LinkButton) SetVisited(visited bool) {
	C.gtk_link_button_set_visited(C.to_GtkLinkButton(v.GWidget), bool2gboolean(visited))
}

//-----------------------------------------------------------------------
// GtkScaleButton
//-----------------------------------------------------------------------

// gtk_scale_button_new
// gtk_scale_button_set_adjustment
// gtk_scale_button_set_icons
// gtk_scale_button_set_value
// gtk_scale_button_get_adjustment
// gtk_scale_button_get_value
// gtk_scale_button_get_popup
// gtk_scale_button_get_plus_button
// gtk_scale_button_get_minus_button

//-----------------------------------------------------------------------
// GtkVolumeButton
//-----------------------------------------------------------------------

// gtk_volume_button_new

//-----------------------------------------------------------------------
// GtkEntry
//-----------------------------------------------------------------------
/*type TextInputLike interface {
	WidgetLike
	GetText() string
	SetText(label string)
}*/

type Entry struct {
	Widget
	Editable
}

func NewEntry() *Entry {
	w := Widget{C.gtk_entry_new()}
	return &Entry{w, Editable{C.to_GtkEditable(w.GWidget)}}
}
func NewEntryWithBuffer(buffer *EntryBuffer) *Entry {
	panic_if_version_older_auto(2, 18, 0)
	w := Widget{C._gtk_entry_new_with_buffer(buffer.GEntryBuffer)}
	return &Entry{w, Editable{C.to_GtkEditable(w.GWidget)}}
}
func (v *Entry) GetBuffer() *EntryBuffer {
	panic_if_version_older_auto(2, 18, 0)
	return &EntryBuffer{C._gtk_entry_get_buffer(C.to_GtkEntry(v.GWidget))}
}
func (v *Entry) SetBuffer(buffer *EntryBuffer) {
	panic_if_version_older_auto(2, 18, 0)
	C._gtk_entry_set_buffer(C.to_GtkEntry(v.GWidget), buffer.GEntryBuffer)
}
func (v *Entry) SetText(text string) {
	ptr := C.CString(text)
	defer C.free_string(ptr)
	C.gtk_entry_set_text(C.to_GtkEntry(v.GWidget), C.to_gcharptr(ptr))
}
func (v *Entry) GetText() string {
	return C.GoString(C.to_charptr(C.gtk_entry_get_text(C.to_GtkEntry(v.GWidget))))
}
func (v *Entry) GetTextLength() int {
	return int(C.gtk_entry_get_text_length(C.to_GtkEntry(v.GWidget)))
}
func (v *Entry) SetVisibility(setting bool) {
	C.gtk_entry_set_visibility(C.to_GtkEntry(v.GWidget), bool2gboolean(setting))
}
func (v *Entry) SetInvisibleChar(ch uint8) {
	C.gtk_entry_set_invisible_char(C.to_GtkEntry(v.GWidget), C.gunichar(ch))
}
func (v *Entry) UnsetInvisibleChar() {
	C.gtk_entry_unset_invisible_char(C.to_GtkEntry(v.GWidget))
}
func (v *Entry) SetMaxLength(i int) {
	C.gtk_entry_set_max_length(C.to_GtkEntry(v.GWidget), C.gint(i))
}
func (v *Entry) GetActivatesDefault() bool {
	return gboolean2bool(C.gtk_entry_get_activates_default(C.to_GtkEntry(v.GWidget)))
}
func (v *Entry) GetHasFrame() bool {
	return gboolean2bool(C.gtk_entry_get_has_frame(C.to_GtkEntry(v.GWidget)))
}

// gtk_entry_get_inner_border

func (v *Entry) GetWidthChars() int {
	return int(C.gtk_entry_get_width_chars(C.to_GtkEntry(v.GWidget)))
}
func (v *Entry) SetActivatesDefault(setting bool) {
	C.gtk_entry_set_activates_default(C.to_GtkEntry(v.GWidget), bool2gboolean(setting))
}
func (v *Entry) SetHasFrame(setting bool) {
	C.gtk_entry_set_has_frame(C.to_GtkEntry(v.GWidget), bool2gboolean(setting))
}

// gtk_entry_set_inner_border

func (v *Entry) SetWidthChars(i int) {
	C.gtk_entry_set_width_chars(C.to_GtkEntry(v.GWidget), C.gint(i))
}
func (v *Entry) GetInvisibleChar() uint8 {
	return uint8(C.gtk_entry_get_invisible_char(C.to_GtkEntry(v.GWidget)))
}
func (v *Entry) SetAlignment(xalign float64) {
	C.gtk_entry_set_alignment(C.to_GtkEntry(v.GWidget), C.gfloat(xalign))
}
func (v *Entry) GetAlignment() float64 {
	return float64(C.gtk_entry_get_alignment(C.to_GtkEntry(v.GWidget)))
}

func (v *Entry) SetOverwriteMode(mode bool) {
	C.gtk_entry_set_overwrite_mode(C.to_GtkEntry(v.GWidget), bool2gboolean(mode))
}

func (v *Entry) GetOverwriteMode() bool {
	return gboolean2bool(C.gtk_entry_get_overwrite_mode(C.to_GtkEntry(v.GWidget)))
}

// gtk_entry_get_layout
// gtk_entry_get_layout_offsets
// gtk_entry_layout_index_to_text_index
// gtk_entry_text_index_to_layout_index

func (v *Entry) GetMaxLength() int {
	return int(C.gtk_entry_get_max_length(C.to_GtkEntry(v.GWidget)))
}
func (v *Entry) GetVisibility() bool {
	return gboolean2bool(C.gtk_entry_get_visibility(C.to_GtkEntry(v.GWidget)))
}
func (v *Entry) SetCompletion(completion *EntryCompletion) {
	C.gtk_entry_set_completion(C.to_GtkEntry(v.GWidget), completion.GEntryCompletion)
}
var g_Entry_EntryCompletion *EntryCompletion
func (v *Entry) GetCompletion() *EntryCompletion {
	if g_Entry_EntryCompletion == nil {
		g_Entry_EntryCompletion = &EntryCompletion{C.gtk_entry_get_completion(C.to_GtkEntry(v.GWidget))}
	} else {
		g_Entry_EntryCompletion.GEntryCompletion = C.gtk_entry_get_completion(C.to_GtkEntry(v.GWidget))
	}
	return g_Entry_EntryCompletion
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
		return &EntryBuffer{C._gtk_entry_buffer_new(nil, C.gint(-1))}
	}
	ptr := C.CString(initialText)
	defer C.free_string(ptr)
	return &EntryBuffer{
		C._gtk_entry_buffer_new(C.to_gcharptr(ptr), C.gint(len(initialText)))}
}
func (v *EntryBuffer) GetText() string {
	panic_if_version_older_auto(2, 18, 0)
	return C.GoString(C.to_charptr(C._gtk_entry_buffer_get_text(v.GEntryBuffer)))
}
func (v *EntryBuffer) SetText(text string) {
	panic_if_version_older_auto(2, 18, 0)
	if len(text) == 0 {
		C._gtk_entry_buffer_set_text(v.GEntryBuffer, nil, C.gint(-1))
	}
	ptr := C.CString(text)
	defer C.free_string(ptr)
	C._gtk_entry_buffer_set_text(v.GEntryBuffer,
		C.to_gcharptr(ptr), C.gint(len(text)))
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
	C._gtk_entry_buffer_set_max_length(v.GEntryBuffer, C.gint(maxLength))
}
func (v *EntryBuffer) InsertText(position uint, text string) uint {
	panic_if_version_older_auto(2, 18, 0)
	ptr := C.CString(text)
	defer C.free_string(ptr)
	return uint(C._gtk_entry_buffer_insert_text(v.GEntryBuffer,
		C.guint(position), C.to_gcharptr(ptr), C.gint(len(text))))
}
func (v *EntryBuffer) DeleteText(position uint, nChars int) uint {
	panic_if_version_older_auto(2, 18, 0)
	return uint(C._gtk_entry_buffer_delete_text(v.GEntryBuffer,
		C.guint(position), C.gint(nChars)))
}

// gtk_entry_buffer_emit_deleted_text //since 2.18
// gtk_entry_buffer_emit_inserted_text //since 2.18

//-----------------------------------------------------------------------
// GtkEntryCompletion
//-----------------------------------------------------------------------

type EntryCompletion struct {
	GEntryCompletion *C.GtkEntryCompletion
}

func NewEntryCompletion() *EntryCompletion {
	return &EntryCompletion{C.gtk_entry_completion_new()}
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
	C.gtk_entry_completion_set_minimum_key_length(v.GEntryCompletion, C.gint(length))
}
func (v *EntryCompletion) GetMinimumKeyLength() int {
	return int(C.gtk_entry_completion_get_minimum_key_length(v.GEntryCompletion))
}
func (v *EntryCompletion) Complete() {
	C.gtk_entry_completion_complete(v.GEntryCompletion)
}
func (v *EntryCompletion) GetCompletionPrefix() string {
	return C.GoString(C.to_charptr(C.gtk_entry_completion_get_completion_prefix(v.GEntryCompletion)))
}
func (v *EntryCompletion) InsertPrefix() {
	C.gtk_entry_completion_insert_prefix(v.GEntryCompletion)
}
func (v *EntryCompletion) InsertActionText(index int, text string) {
	ptr := C.CString(text)
	defer C.free_string(ptr)
	C.gtk_entry_completion_insert_action_text(
		v.GEntryCompletion, C.gint(index), C.to_gcharptr(ptr))
}
func (v *EntryCompletion) InsertActionMarkup(index int, markup string) {
	ptr := C.CString(markup)
	defer C.free_string(ptr)
	C.gtk_entry_completion_insert_action_markup(
		v.GEntryCompletion, C.gint(index), C.to_gcharptr(ptr))
}
func (v *EntryCompletion) DeleteAction(index int) {
	C.gtk_entry_completion_delete_action(v.GEntryCompletion, C.gint(index))
}
func (v *EntryCompletion) SetTextColumn(column int) {
	C.gtk_entry_completion_set_text_column(v.GEntryCompletion, C.gint(column))
}
func (v *EntryCompletion) GetTextColumn() int {
	return int(C.gtk_entry_completion_get_text_column(v.GEntryCompletion))
}
func (v *EntryCompletion) SetInlineCompletion(inlineCompletion bool) {
	C.gtk_entry_completion_set_inline_completion(v.GEntryCompletion,
		bool2gboolean(inlineCompletion))
}
func (v *EntryCompletion) GetInlineCompletion() bool {
	return gboolean2bool(C.gtk_entry_completion_get_inline_completion(v.GEntryCompletion))
}
func (v *EntryCompletion) SetInlineSelection(inlineSelection bool) {
	C.gtk_entry_completion_set_inline_selection(v.GEntryCompletion,
		bool2gboolean(inlineSelection))
}
func (v *EntryCompletion) GetInlineSelection() bool {
	return gboolean2bool(C.gtk_entry_completion_get_inline_selection(v.GEntryCompletion))
}
func (v *EntryCompletion) SetPopupCompletion(popupCompletion bool) {
	C.gtk_entry_completion_set_popup_completion(v.GEntryCompletion,
		bool2gboolean(popupCompletion))
}
func (v *EntryCompletion) GetPopupCompletion() bool {
	return gboolean2bool(C.gtk_entry_completion_get_popup_completion(v.GEntryCompletion))
}
func (v *EntryCompletion) SetPopupSetWidth(popupSetWidth bool) {
	C.gtk_entry_completion_set_popup_set_width(v.GEntryCompletion,
		bool2gboolean(popupSetWidth))
}
func (v *EntryCompletion) GetPopupSetWidth() bool {
	return gboolean2bool(C.gtk_entry_completion_get_popup_set_width(v.GEntryCompletion))
}
func (v *EntryCompletion) SetPopupSingleMatch(popupSingleMatch bool) {
	C.gtk_entry_completion_set_popup_single_match(v.GEntryCompletion,
		bool2gboolean(popupSingleMatch))
}
func (v *EntryCompletion) GetPopupSingleMatch() bool {
	return gboolean2bool(C.gtk_entry_completion_get_popup_single_match(v.GEntryCompletion))
}

//-----------------------------------------------------------------------
// GtkHScale
//-----------------------------------------------------------------------
func NewHScale(adjustment *Adjustment) *Scale {
	return &Scale{Range{Widget{C.gtk_hscale_new(adjustment.GAdjustment)}}}
}
func NewHScaleWithRange(min, max, step float64) *Scale {
	return &Scale{Range{Widget{
		C.gtk_hscale_new_with_range(C.gdouble(min), C.gdouble(max), C.gdouble(step))}}}
}

//-----------------------------------------------------------------------
// GtkVScale
//-----------------------------------------------------------------------
func NewVScale(a *Adjustment) *Scale {
	return &Scale{Range{Widget{C.gtk_vscale_new(a.GAdjustment)}}}
}
func NewVScaleWithRange(min, max, step float64) *Scale {
	return &Scale{Range{Widget{C.gtk_vscale_new_with_range(C.gdouble(min), C.gdouble(max), C.gdouble(step))}}}
}

//-----------------------------------------------------------------------
// GtkSpinButton
//-----------------------------------------------------------------------
type SpinButtonUpdatePolicy int

const (
	UPDATE_ALWAYS           = 0
	UPDATE_IF_VALID         = 1
)

type SpinType int

const (
	SPIN_STEP_FORWARD       = 0
	SPIN_STEP_BACKWARD      = 1
	SPIN_PAGE_FORWARD       = 2
	SPIN_PAGE_BACKWARD      = 3
	SPIN_HOME               = 4
	SPIN_END                = 5
	SPIN_USER_DEFINED       = 6
)

type SpinButton struct {
	Entry
}

func NewSpinButton(a *Adjustment, climb float64, digits uint) *SpinButton {
	w := Widget{C.gtk_spin_button_new(a.GAdjustment, C.gdouble(climb), C.guint(digits))}
	return &SpinButton{Entry{w, Editable{C.to_GtkEditable(w.GWidget)}}}
}
func NewSpinButtonWithRange(min, max, step float64) *SpinButton {
	w := Widget{C.gtk_spin_button_new_with_range(C.gdouble(min), C.gdouble(max), C.gdouble(step))}
	return &SpinButton{Entry{w, Editable{C.to_GtkEditable(w.GWidget)}}}
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
	C.gtk_spin_button_configure(C.to_GtkSpinButton(v.GWidget), a.GAdjustment, C.gdouble(climb_rate), C.guint(digits))
}
func (v *SpinButton) SetAdjustment(a *Adjustment) {
	C.gtk_spin_button_set_adjustment(C.to_GtkSpinButton(v.GWidget), a.GAdjustment)
}
func (v *SpinButton) GetAdjustment() *Adjustment {
	return &Adjustment{C.gtk_spin_button_get_adjustment(C.to_GtkSpinButton(v.GWidget))}
}
func (v *SpinButton) SetDigits(digits uint) {
	C.gtk_spin_button_set_digits(C.to_GtkSpinButton(v.GWidget), C.guint(digits))
}
func (v *SpinButton) SetIncrements(step, page float64) {
	C.gtk_spin_button_set_increments(C.to_GtkSpinButton(v.GWidget), C.gdouble(step), C.gdouble(page))
}
func (v *SpinButton) SetRange(min, max float64) {
	C.gtk_spin_button_set_range(C.to_GtkSpinButton(v.GWidget), C.gdouble(min), C.gdouble(max))
}
func (v *SpinButton) GetValueAsFloat() float64 {
	return float64(C.gtk_spin_button_get_value_as_float(C.to_GtkSpinButton(v.GWidget)))
}
func (v *SpinButton) GetValueAsInt() int {
	return int(C.gtk_spin_button_get_value_as_int(C.to_GtkSpinButton(v.GWidget)))
}
func (v *SpinButton) SetValue(val float64) {
	C.gtk_spin_button_set_value(C.to_GtkSpinButton(v.GWidget), C.gdouble(val))
}
func (v *SpinButton) SetUpdatePolicy(policy SpinButtonUpdatePolicy) {
	C.gtk_spin_button_set_update_policy(C.to_GtkSpinButton(v.GWidget), C.GtkSpinButtonUpdatePolicy(policy))
}
func (v *SpinButton) SetNumeric(numeric bool) {
	C.gtk_spin_button_set_numeric(C.to_GtkSpinButton(v.GWidget), bool2gboolean(numeric))
}
func (v *SpinButton) Spin(direction SpinType, increment float64) {
	C.gtk_spin_button_spin(C.to_GtkSpinButton(v.GWidget), C.GtkSpinType(direction), C.gdouble(increment))
}
func (v *SpinButton) SetWrap(wrap bool) {
	C.gtk_spin_button_set_wrap(C.to_GtkSpinButton(v.GWidget), bool2gboolean(wrap))
}
func (v *SpinButton) SetSnapToTicks(snap_to_ticks bool) {
	C.gtk_spin_button_set_snap_to_ticks(C.to_GtkSpinButton(v.GWidget), bool2gboolean(snap_to_ticks))
}
func (v *SpinButton) Update() {
	C.gtk_spin_button_update(C.to_GtkSpinButton(v.GWidget))
}
func (v *SpinButton) GetDigits() uint {
	return uint(C.gtk_spin_button_get_digits(C.to_GtkSpinButton(v.GWidget)))
}
func (v *SpinButton) GetIncrements() (float64, float64) {
	var step, page C.gdouble
	C.gtk_spin_button_get_increments(C.to_GtkSpinButton(v.GWidget), &step, &page)
	return float64(step), float64(page)
}
func (v *SpinButton) GetNumeric() bool {
	return gboolean2bool(C.gtk_spin_button_get_numeric(C.to_GtkSpinButton(v.GWidget)))
}
func (v *SpinButton) GetRange() (/*min*/ float64, /*max*/ float64) {
	var min, max C.gdouble
	C.gtk_spin_button_get_range(C.to_GtkSpinButton(v.GWidget), &min, &max)
	return float64(min), float64(max)
}
func (v *SpinButton) GetSnapToTicks() bool {
	return gboolean2bool(C.gtk_spin_button_get_snap_to_ticks(C.to_GtkSpinButton(v.GWidget)))
}
func (v *SpinButton) GetUpdatePolicy() SpinButtonUpdatePolicy {
	return SpinButtonUpdatePolicy(C.gtk_spin_button_get_update_policy(C.to_GtkSpinButton(v.GWidget)))
}
func (v *SpinButton) GetValue() float64 {
	return float64(C.gtk_spin_button_get_value(C.to_GtkSpinButton(v.GWidget)))
}
func (v *SpinButton) GetWrap() bool {
	return gboolean2bool(C.gtk_spin_button_get_wrap(C.to_GtkSpinButton(v.GWidget)))
}

//-----------------------------------------------------------------------
// GtkEditable
//-----------------------------------------------------------------------
type Editable struct {
	GEditable *C.GtkEditable
}

func (v *Editable) SelectRegion(startPos int, endPos int) {
	C.gtk_editable_select_region(v.GEditable, C.gint(startPos), C.gint(endPos))
}
func (v *Editable) GetSelectionBounds() (isSelected bool,
	startPos int, endPos int) {
	var s, e C.gint
	return gboolean2bool(C.gtk_editable_get_selection_bounds(v.GEditable, &s, &e)), int(s), int(e)
}
func (v *Editable) InsertText(newText string, position int) int {
	ptr := C.CString(newText)
	defer C.free_string(ptr)
	gpos := (C.gint)(position)
	C.gtk_editable_insert_text(v.GEditable, C.to_gcharptr(ptr), C.gint(len(newText)), &gpos)
	return int(gpos)
}
func (v *Editable) DeleteText(startPos int, endPos int) {
	C.gtk_editable_delete_text(v.GEditable, C.gint(startPos), C.gint(endPos))
}
func (v *Editable) GetChars(startPos int, endPos int) string {
	return C.GoString(C.to_charptr(C.gtk_editable_get_chars(v.GEditable, C.gint(startPos), C.gint(endPos))))
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
	C.gtk_editable_set_position(v.GEditable, C.gint(position))
}
func (v *Editable) GetPosition() int {
	return int(C.gtk_editable_get_position(v.GEditable))
}
func (v *Editable) SetEditable(isEditable bool) {
	C.gtk_editable_set_editable(v.GEditable, bool2gboolean(isEditable))
}
func (v *Editable) GetEditable() bool {
	return gboolean2bool(C.gtk_editable_get_editable(v.GEditable))
}

//-----------------------------------------------------------------------
// GtkTextIter
//-----------------------------------------------------------------------
type TextIter struct {
	GTextIter C.GtkTextIter
}

type TextSearchFlags int

const (
	TEXT_SEARCH_VISIBLE_ONLY     TextSearchFlags = 1 << 0
	TEXT_SEARCH_TEXT_ONLY        TextSearchFlags = 1 << 1
	TEXT_SEARCH_CASE_INSENSITIVE TextSearchFlags = 1 << 2
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
	pchar := C.to_charptr(C.gtk_text_iter_get_slice(&v.GTextIter, &end.GTextIter))
	defer C.free(unsafe.Pointer(pchar))
	return C.GoString(pchar)
}
func (v *TextIter) GetText(end *TextIter) string {
	pchar := C.to_charptr(C.gtk_text_iter_get_text(&v.GTextIter, &end.GTextIter))
	defer C.free(unsafe.Pointer(pchar))
	return C.GoString(pchar)
}
func (v *TextIter) GetVisibleSlice(end *TextIter) string {
	return C.GoString(C.to_charptr(C.gtk_text_iter_get_visible_slice(&v.GTextIter, &end.GTextIter)))
}
func (v *TextIter) GetVisibleText(end *TextIter) string {
	return C.GoString(C.to_charptr(C.gtk_text_iter_get_visible_text(&v.GTextIter, &end.GTextIter)))
}

// gtk_text_iter_get_pixbuf

func (v *TextIter) GetMarks() *glib.SList {
	return glib.SListFromNative(unsafe.Pointer(C.gtk_text_iter_get_marks(&v.GTextIter)))
}

// gtk_text_iter_get_toggled_tags
// gtk_text_iter_get_child_anchor
// gtk_text_iter_begins_tag
// gtk_text_iter_ends_tag
// gtk_text_iter_toggles_tag
// gtk_text_iter_has_tag
// gtk_text_iter_get_tags
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
	return gboolean2bool(C.gtk_text_iter_forward_char(&v.GTextIter))
}

// gtk_text_iter_backward_char
// gtk_text_iter_forward_chars
// gtk_text_iter_backward_chars
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
	defer C.free_string(cstr)
	return gboolean2bool(C.gtk_text_iter_forward_search(&v.GTextIter,
		C.to_gcharptr(cstr), C.GtkTextSearchFlags(flags), &start.GTextIter,
		&end.GTextIter, &limit.GTextIter))
}
func (v *TextIter) BackwardSearch(str string, flags TextSearchFlags, start *TextIter, end *TextIter, limit *TextIter) bool {
	cstr := C.CString(str)
	defer C.free_string(cstr)
	return gboolean2bool(C.gtk_text_iter_backward_search(&v.GTextIter,
		C.to_gcharptr(cstr), C.GtkTextSearchFlags(flags), &start.GTextIter,
		&end.GTextIter, &limit.GTextIter))
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
type TextBufferLike interface {
	GetNativeBuffer() unsafe.Pointer
}

type TextBuffer struct {
	GTextBuffer *C.GtkTextBuffer
	*glib.GObject
}

func newTextBuffer(buffer *C.GtkTextBuffer) *TextBuffer { // TODO
	return &TextBuffer{
		GTextBuffer: buffer,
		GObject:    glib.ObjectFromNative(unsafe.Pointer(buffer)),
	}
}
func NewTextBufferFromPointer(v unsafe.Pointer) TextBuffer {
	return *newTextBuffer(C.to_GtkTextBuffer(v))
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
	defer C.free_string(ptr)
	l := C.strlen(ptr)
	C.gtk_text_buffer_insert(v.GTextBuffer, &iter.GTextIter, C.to_gcharptr(ptr), C.gint(l))
}
func (v *TextBuffer) InsertAtCursor(text string) {
	ptr := C.CString(text)
	defer C.free_string(ptr)
	l := C.strlen(ptr)
	C.gtk_text_buffer_insert_at_cursor(v.GTextBuffer, C.to_gcharptr(ptr), C.gint(l))
}
func (v *TextBuffer) InsertInteractive(iter *TextIter, text string, default_editable bool) bool {
	ptr := C.CString(text)
	defer C.free_string(ptr)
	l := C.strlen(ptr)
	return gboolean2bool(C.gtk_text_buffer_insert_interactive(v.GTextBuffer, &iter.GTextIter, C.to_gcharptr(ptr), C.gint(l), bool2gboolean(default_editable)))
}
func (v *TextBuffer) InsertInteractiveAtCursor(text string, default_editable bool) bool {
	ptr := C.CString(text)
	defer C.free_string(ptr)
	l := C.strlen(ptr)
	return gboolean2bool(C.gtk_text_buffer_insert_interactive_at_cursor(v.GTextBuffer, C.to_gcharptr(ptr), C.gint(l), bool2gboolean(default_editable)))
}
func (v *TextBuffer) InsertRange(iter *TextIter, start *TextIter, end *TextIter) {
	C.gtk_text_buffer_insert_range(v.GTextBuffer, &iter.GTextIter, &start.GTextIter, &end.GTextIter)
}
func (v *TextBuffer) InsertRangeInteractive(iter *TextIter, start *TextIter, end *TextIter, default_editable bool) bool {
	return gboolean2bool(C.gtk_text_buffer_insert_range_interactive(v.GTextBuffer, &iter.GTextIter, &start.GTextIter, &end.GTextIter, bool2gboolean(default_editable)))
}
func (v *TextBuffer) InsertWithTag(iter *TextIter, text string, tag *TextTag) {
	ptr := C.CString(text)
	defer C.free_string(ptr)
	l := C.strlen(ptr)
	C._gtk_text_buffer_insert_with_tag(v.GTextBuffer, &iter.GTextIter, C.to_gcharptr(ptr), C.gint(l), tag.GTextTag)
}

//func (v GtkTextBuffer) InsertWithTags(iter *GtkTextIter, start *GtkTextIter, end *GtkTextIter, default_editable bool) bool {
//	return gboolean2bool(C._gtk_text_buffer_insert_range_interactive(v.GTextBuffer, &iter.GTextIter, &start.GTextIter, &end.GTextIter, bool2gboolean(default_editable)));
//}
// gtk_text_buffer_insert_with_tags_by_name

func (v *TextBuffer) Delete(start *TextIter, end *TextIter) {
	C.gtk_text_buffer_delete(v.GTextBuffer, &start.GTextIter, &end.GTextIter)
}
func (v *TextBuffer) DeleteInteractive(start *TextIter, end *TextIter, default_editable bool) bool {
	return gboolean2bool(C.gtk_text_buffer_delete_interactive(v.GTextBuffer, &start.GTextIter, &end.GTextIter, bool2gboolean(default_editable)))
}
func (v *TextBuffer) Backspace(iter *TextIter, interactive bool, default_editable bool) bool {
	return gboolean2bool(C.gtk_text_buffer_backspace(v.GTextBuffer, &iter.GTextIter, bool2gboolean(interactive), bool2gboolean(default_editable)))
}
func (v *TextBuffer) SetText(text string) {
	ptr := C.CString(text)
	defer C.free_string(ptr)
	l := C.strlen(ptr)
	C.gtk_text_buffer_set_text(v.GTextBuffer, C.to_gcharptr(ptr), C.gint(l))
}
func (v *TextBuffer) GetText(start *TextIter, end *TextIter, include_hidden_chars bool) string {
	pchar := C.to_charptr(C.gtk_text_buffer_get_text(v.GTextBuffer, &start.GTextIter, &end.GTextIter, bool2gboolean(include_hidden_chars)))
	defer C.free(unsafe.Pointer(pchar))
	return C.GoString(pchar)
}
func (v *TextBuffer) GetSlice(start *TextIter, end *TextIter, include_hidden_chars bool) string {
	pchar := C.to_charptr(C.gtk_text_buffer_get_slice(v.GTextBuffer, &start.GTextIter, &end.GTextIter, bool2gboolean(include_hidden_chars)))
	defer C.free(unsafe.Pointer(pchar))
	return C.GoString(pchar)
}
func (v *TextBuffer) InsertPixbuf(iter *TextIter, pixbuf *gdkpixbuf.GdkPixbuf) {
	C.gtk_text_buffer_insert_pixbuf(v.GTextBuffer, &iter.GTextIter, pixbuf.Pixbuf)
}

// gtk_text_buffer_insert_child_anchor
// gtk_text_buffer_create_child_anchor

func (v *TextBuffer) CreateMark(mark_name string, where *TextIter, left_gravity bool) *TextMark {
	ptr := C.CString(mark_name)
	defer C.free_string(ptr)
	return &TextMark{C.gtk_text_buffer_create_mark(v.GTextBuffer, C.to_gcharptr(ptr), &where.GTextIter, bool2gboolean(left_gravity))}
}
func (v *TextBuffer) MoveMark(mark *TextMark, where *TextIter) {
	C.gtk_text_buffer_move_mark(v.GTextBuffer, mark.GTextMark, &where.GTextIter)
}
func (v *TextBuffer) MoveMarkByName(name string, where *TextIter) {
	ptr := C.CString(name)
	defer C.free_string(ptr)
	C.gtk_text_buffer_move_mark_by_name(v.GTextBuffer, C.to_gcharptr(ptr), &where.GTextIter)
}
func (v *TextBuffer) AddMark(mark *TextMark, where *TextIter) {
	C.gtk_text_buffer_add_mark(v.GTextBuffer, mark.GTextMark, &where.GTextIter)
}
func (v *TextBuffer) DeleteMark(mark *TextMark) {
	C.gtk_text_buffer_delete_mark(v.GTextBuffer, mark.GTextMark)
}
func (v *TextBuffer) DeleteMarkByName(name string) {
	ptr := C.CString(name)
	defer C.free_string(ptr)
	C.gtk_text_buffer_delete_mark_by_name(v.GTextBuffer, C.to_gcharptr(ptr))
}
func (v *TextBuffer) GetMark(name string) *TextMark {
	ptr := C.CString(name)
	defer C.free_string(ptr)
	return &TextMark{
		C.gtk_text_buffer_get_mark(v.GTextBuffer, C.to_gcharptr(ptr))}
}
func (v *TextBuffer) GetInsert() *TextMark {
	return &TextMark{
		C.gtk_text_buffer_get_insert(v.GTextBuffer)}
}
func (v *TextBuffer) GetSelectionBound() *TextMark {
	return &TextMark{
		C.gtk_text_buffer_get_selection_bound(v.GTextBuffer)}
}
func (v *TextBuffer) GetHasSelection() bool {
	return gboolean2bool(C.gtk_text_buffer_get_has_selection(v.GTextBuffer))
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
	defer C.free_string(ptr)
	C.gtk_text_buffer_apply_tag_by_name(v.GTextBuffer, C.to_gcharptr(ptr), &start.GTextIter, &end.GTextIter)
}
func (v *TextBuffer) RemoveTagByName(name string, start *TextIter, end *TextIter) {
	ptr := C.CString(name)
	defer C.free_string(ptr)
	C.gtk_text_buffer_remove_tag_by_name(v.GTextBuffer, C.to_gcharptr(ptr), &start.GTextIter, &end.GTextIter)
}
func (v *TextBuffer) RemoveAllTags(start *TextIter, end *TextIter) {
	C.gtk_text_buffer_remove_all_tags(v.GTextBuffer, &start.GTextIter, &end.GTextIter)
}
func (v *TextBuffer) CreateTag(tag_name string, props map[string]string) *TextTag {
	ptr := C.CString(tag_name)
	defer C.free_string(ptr)
	tag := C._gtk_text_buffer_create_tag(v.GTextBuffer, C.to_gcharptr(ptr))
	for prop, val := range props {
		pprop := C.CString(prop)
		pval := C.CString(val)
		C._apply_property(unsafe.Pointer(tag), C.to_gcharptr(pprop), C.to_gcharptr(pval))
		C.free_string(pprop)
		C.free_string(pval)
	}
	return &TextTag{tag}
}
func (v *TextBuffer) GetIterAtLineOffset(iter *TextIter, line_number int, char_offset int) {
	C.gtk_text_buffer_get_iter_at_line_offset(v.GTextBuffer, &iter.GTextIter, C.gint(line_number), C.gint(char_offset))
}
func (v *TextBuffer) GetIterAtOffset(iter *TextIter, char_offset int) {
	C.gtk_text_buffer_get_iter_at_offset(v.GTextBuffer, &iter.GTextIter, C.gint(char_offset))
}
func (v *TextBuffer) GetIterAtLine(iter *TextIter, line_number int) {
	C.gtk_text_buffer_get_iter_at_line(v.GTextBuffer, &iter.GTextIter, C.gint(line_number))
}
func (v *TextBuffer) GetIterAtLineIndex(iter *TextIter, line_number int, byte_index int) {
	C.gtk_text_buffer_get_iter_at_line_index(v.GTextBuffer, &iter.GTextIter, C.gint(line_number), C.gint(byte_index))
}
func (v *TextBuffer) GetIterAtMark(iter *TextIter, mark *TextMark) {
	C.gtk_text_buffer_get_iter_at_mark(v.GTextBuffer, &iter.GTextIter, mark.GTextMark)
}
func (v *TextBuffer) GetIterAtChildAnchor(i *TextIter, a *TextChildAnchor) {
	C.gtk_text_buffer_get_iter_at_child_anchor(v.GTextBuffer, &i.GTextIter, a.GTextChildAnchor)
}
func (v *TextBuffer) GetStartIter(iter *TextIter) {
	C.gtk_text_buffer_get_start_iter(v.GTextBuffer, &iter.GTextIter)
}
func (v *TextBuffer) GetEndIter(iter *TextIter) {
	C.gtk_text_buffer_get_end_iter(v.GTextBuffer, &iter.GTextIter)
}
func (v *TextBuffer) GetBounds(start *TextIter, end *TextIter) {
	C.gtk_text_buffer_get_bounds(v.GTextBuffer, &start.GTextIter, &end.GTextIter)
}
func (v *TextBuffer) GetModified() bool {
	return gboolean2bool(C.gtk_text_buffer_get_modified(v.GTextBuffer))
}
func (v *TextBuffer) SetModified(setting bool) {
	C.gtk_text_buffer_set_modified(v.GTextBuffer, bool2gboolean(setting))
}
func (v *TextBuffer) DeleteSelection(interactive bool, default_editable bool) {
	C.gtk_text_buffer_delete_selection(v.GTextBuffer, bool2gboolean(interactive), bool2gboolean(default_editable))
}

// gtk_text_buffer_paste_clipboard
// gtk_text_buffer_copy_clipboard
// gtk_text_buffer_cut_clipboard

func (v *TextBuffer) GetSelectionBounds(be, en *TextIter) bool {
	return gboolean2bool(C.gtk_text_buffer_get_selection_bounds(v.GTextBuffer, &be.GTextIter, &en.GTextIter))
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
}

func NewTextTag(name string) *TextTag {
	ptr := C.CString(name)
	defer C.free_string(ptr)
	return &TextTag{
		C.gtk_text_tag_new(C.to_gcharptr(ptr))}
}
func (v *TextTag) SetPriority(priority int) {
	C.gtk_text_tag_set_priority(v.GTextTag, C.gint(priority))
}
func (v *TextTag) GetPriority() int {
	return int(C.gtk_text_tag_get_priority(v.GTextTag))
}

// gtk_text_tag_event

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
	defer C.free_string(ptr)
	return &TextTag{
		C.gtk_text_tag_table_lookup(v.GTextTagTable, C.to_gcharptr(ptr))}
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
	WRAP_NONE      WrapMode = 0
	WRAP_CHAR      WrapMode = 1
	WRAP_WORD      WrapMode = 2
	WRAP_WORD_CHAR WrapMode = 3
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
func (v *TextView) SetBuffer(b TextBufferLike) {
	C.gtk_text_view_set_buffer(C.to_GtkTextView(v.GWidget), C.to_GtkTextBuffer(b.GetNativeBuffer()))
}
func (v *TextView) GetBuffer() *TextBuffer {
	return newTextBuffer(C.gtk_text_view_get_buffer(C.to_GtkTextView(v.GWidget)))
}
func (v *TextView) ScrollToMark(mark *TextMark, wm float64, ua bool, xa float64, ya float64) {
	C.gtk_text_view_scroll_to_mark(C.to_GtkTextView(v.GWidget),
		mark.GTextMark, C.gdouble(wm), bool2gboolean(ua), C.gdouble(xa), C.gdouble(ya))
}
func (v *TextView) ScrollToIter(iter *TextIter, wm float64, ua bool, xa float64, ya float64) bool {
	return gboolean2bool(C.gtk_text_view_scroll_to_iter(C.to_GtkTextView(v.GWidget),
		&iter.GTextIter, C.gdouble(wm), bool2gboolean(ua), C.gdouble(xa), C.gdouble(ya)))
}

// void gtk_text_view_scroll_mark_onscreen(GtkTextView* text_view, GtkTextMark* mark);
// gboolean gtk_text_view_move_mark_onscreen(GtkTextView* text_view, GtkTextMark* mark);
// gboolean gtk_text_view_place_cursor_onscreen(GtkTextView* text_view);
// gtk_text_view_get_visible_rect
// void gtk_text_view_get_iter_location(GtkTextView* text_view, const GtkTextIter* iter, GdkRectangle* location);
// void gtk_text_view_get_line_at_y(GtkTextView* text_view, GtkTextIter* target_iter, gint y, gint* line_top);

func (v *TextView) GetLineYrange(iter *TextIter, y *int, h *int) {
	var yy, hh C.gint
	C.gtk_text_view_get_line_yrange(C.to_GtkTextView(v.GWidget), &iter.GTextIter, &yy, &hh)
	*y = int(yy)
	*h = int(hh)
}

// void gtk_text_view_get_iter_at_location(GtkTextView* text_view, GtkTextIter* iter, gint x, gint y);

func (v *TextView) GetIterAtPosition(iter *TextIter, trailing *int, x int, y int) {
	if nil != trailing {
		var tt C.gint
		C.gtk_text_view_get_iter_at_position(C.to_GtkTextView(v.GWidget), &iter.GTextIter, &tt, C.gint(x), C.gint(y))
		*trailing = int(tt)
	} else {
		C.gtk_text_view_get_iter_at_position(C.to_GtkTextView(v.GWidget), &iter.GTextIter, nil, C.gint(x), C.gint(y))
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
	C.gtk_text_view_set_wrap_mode(C.to_GtkTextView(v.GWidget), C.GtkWrapMode(mode))
}
func (v *TextView) GetWrapMode() WrapMode {
	return WrapMode(C.gtk_text_view_get_wrap_mode(C.to_GtkTextView(v.GWidget)))
}
func (v *TextView) SetEditable(setting bool) {
	C.gtk_text_view_set_editable(C.to_GtkTextView(v.GWidget), bool2gboolean(setting))
}
func (v *TextView) GetEditable() bool {
	return gboolean2bool(C.gtk_text_view_get_editable(C.to_GtkTextView(v.GWidget)))
}
func (v *TextView) SetCursorVisible(setting bool) {
	C.gtk_text_view_set_cursor_visible(C.to_GtkTextView(v.GWidget), bool2gboolean(setting))
}
func (v *TextView) GetCursorVisible() bool {
	return gboolean2bool(C.gtk_text_view_get_cursor_visible(C.to_GtkTextView(v.GWidget)))
}
func (v *TextView) SetOverwrite(overwrite bool) {
	C.gtk_text_view_set_overwrite(C.to_GtkTextView(v.GWidget), bool2gboolean(overwrite))
}
func (v *TextView) GetOverwrite() bool {
	return gboolean2bool(C.gtk_text_view_get_overwrite(C.to_GtkTextView(v.GWidget)))
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
	C.gtk_text_view_set_accepts_tab(C.to_GtkTextView(v.GWidget), bool2gboolean(accepts_tab))
}
func (v *TextView) GetAcceptsTab() bool {
	return gboolean2bool(C.gtk_text_view_get_accepts_tab(C.to_GtkTextView(v.GWidget)))
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
	defer C.free_string(ptr)
	return &TreePath{
		C.gtk_tree_path_new_from_string(C.to_gcharptr(ptr))}
}
func NewTreePathNewFirst() *TreePath {
	return &TreePath{C.gtk_tree_path_new_first()}
}
// gtk_tree_path_new_from_indices

func (v *TreePath) String() string {
	return C.GoString(C.to_charptr(C.gtk_tree_path_to_string(v.GTreePath)))
}
func (v *TreePath) AppendIndex(index int) {
	C.gtk_tree_path_append_index(v.GTreePath, C.gint(index))
}
func (v *TreePath) PrependIndex(index int) {
	C.gtk_tree_path_prepend_index(v.GTreePath, C.gint(index))
}
func (v *TreePath) GetDepth() int {
	return int(C.gtk_tree_path_get_depth(v.GTreePath))
}

// gtk_tree_path_get_indices
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
	return gboolean2bool(C.gtk_tree_path_prev(v.GTreePath))
}
func (v *TreePath) Up() bool {
	return gboolean2bool(C.gtk_tree_path_up(v.GTreePath))
}
func (v *TreePath) Down() {
	C.gtk_tree_path_down(v.GTreePath)
}
func (v *TreePath) IsAncestor(descendant TreePath) bool {
	return gboolean2bool(C.gtk_tree_path_is_ancestor(v.GTreePath, descendant.GTreePath))
}
func (v *TreePath) IsDescendant(ancestor TreePath) bool {
	return gboolean2bool(C.gtk_tree_path_is_descendant(v.GTreePath, ancestor.GTreePath))
}

//-----------------------------------------------------------------------
// GtkTreeRowReference
//-----------------------------------------------------------------------

// gtk_tree_row_reference_new
// gtk_tree_row_reference_new_proxy
// gtk_tree_row_reference_get_model
// gtk_tree_row_reference_get_path
// gtk_tree_row_reference_valid
// gtk_tree_row_reference_free
// gtk_tree_row_reference_copy
// gtk_tree_row_reference_inserted
// gtk_tree_row_reference_deleted
// gtk_tree_row_reference_reordered

//-----------------------------------------------------------------------
// GtkTreeIter
//-----------------------------------------------------------------------
type TreeIter struct {
	GTreeIter C.GtkTreeIter
}

func (v *TreeIter) Assign(to *TreeIter) {
	C._gtk_tree_iter_assign(unsafe.Pointer(&v.GTreeIter), unsafe.Pointer(&to.GTreeIter))
}

//-----------------------------------------------------------------------
// GtkTreeModel
//-----------------------------------------------------------------------
type TreeModelFlags int

const (
	TREE_MODEL_ITERS_PERSIST TreeModelFlags = 1 << 0
	TREE_MODEL_LIST_ONLY     TreeModelFlags = 1 << 1
)

type TreeModelLike interface {
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
	return gboolean2bool(C.gtk_tree_model_get_iter(v.GTreeModel, &iter.GTreeIter, path.GTreePath))
}
func (v *TreeModel) GetIterFromString(iter *TreeIter, path_string string) bool {
	ptr := C.CString(path_string)
	defer C.free_string(ptr)
	ret := gboolean2bool(C.gtk_tree_model_get_iter_from_string(v.GTreeModel, &iter.GTreeIter, C.to_gcharptr(ptr)))
	return ret
}
func (v *TreeModel) GetIterFirst(iter *TreeIter) bool {
	return gboolean2bool(C.gtk_tree_model_get_iter_first(v.GTreeModel, &iter.GTreeIter))
}
func (v *TreeModel) GetPath(iter *TreeIter) *TreePath {
	return &TreePath{C._gtk_tree_model_get_path(v.GTreeModel, &iter.GTreeIter)}
}
func (v *TreeModel) GetValue(iter *TreeIter, col int, val *glib.GValue) {
	C.gtk_tree_model_get_value(v.GTreeModel, &iter.GTreeIter, C.gint(col), C.to_GValueptr(unsafe.Pointer(&val.Value)))
}
func (v *TreeModel) IterNext(iter *TreeIter) bool {
	return gboolean2bool(C.gtk_tree_model_iter_next(v.GTreeModel, &iter.GTreeIter))
}
func (v *TreeModel) IterChildren(iter *TreeIter, parent *TreeIter) bool {
	return gboolean2bool(C.gtk_tree_model_iter_children(v.GTreeModel, &iter.GTreeIter, &parent.GTreeIter))
}
func (v *TreeModel) IterHasChild(iter *TreeIter) bool {
	return gboolean2bool(C.gtk_tree_model_iter_has_child(v.GTreeModel, &iter.GTreeIter))
}
func (v *TreeModel) IterNChildren(iter *TreeIter) int {
	return int(C.gtk_tree_model_iter_n_children(v.GTreeModel, &iter.GTreeIter))
}
func (v *TreeModel) IterNthChild(iter *TreeIter, parent *TreeIter, n int) bool {
	return gboolean2bool(C.gtk_tree_model_iter_nth_child(v.GTreeModel, &iter.GTreeIter, &parent.GTreeIter, C.gint(n)))
}
func (v *TreeModel) IterParent(iter *TreeIter, child *TreeIter) bool {
	return gboolean2bool(C.gtk_tree_model_iter_parent(v.GTreeModel, &iter.GTreeIter, &child.GTreeIter))
}
func (v *TreeModel) GetStringFromIter(i *TreeIter) string {
	return C.GoString(C.to_charptr(C.gtk_tree_model_get_string_from_iter(v.GTreeModel, &i.GTreeIter)))
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
	SELECTION_NONE     SelectionMode = 0
	SELECTION_SINGLE   SelectionMode = 1
	SELECTION_BROWSE   SelectionMode = 2
	SELECTION_MULTIPLE SelectionMode = 3
	SELECTION_EXTENDED                  = SELECTION_MULTIPLE
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

func (v *TreeSelection) GetSelected(i *TreeIter) bool {
	return gboolean2bool(C.gtk_tree_selection_get_selected(v.GTreeSelection, nil, &i.GTreeIter))
}

//gtk_tree_selection_selected_foreach (GtkTreeSelection *selection, GtkTreeSelectionForeachFunc func, gpointer data);
//gtk_tree_selection_get_selected_rows (GtkTreeSelection *selection, GtkTreeModel **model);

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
	return gboolean2bool(C.gtk_tree_selection_path_is_selected(v.GTreeSelection, path.GTreePath))
}
func (v *TreeSelection) SelectIter(iter *TreeIter) {
	C.gtk_tree_selection_select_iter(v.GTreeSelection, &iter.GTreeIter)
}
func (v *TreeSelection) UnselectIter(iter *TreeIter) {
	C.gtk_tree_selection_unselect_iter(v.GTreeSelection, &iter.GTreeIter)
}
func (v *TreeSelection) IterIsSelected(iter *TreeIter) bool {
	return gboolean2bool(C.gtk_tree_selection_iter_is_selected(v.GTreeSelection, &iter.GTreeIter))
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
	TREE_VIEW_COLUMN_GROW_ONLY TreeViewColumnSizing = 0
	TREE_VIEW_COLUMN_AUTOSIZE  TreeViewColumnSizing = 1
	TREE_VIEW_COLUMN_FIXED     TreeViewColumnSizing = 2
)

type TreeViewColumn struct {
	GTreeViewColumn *C.GtkTreeViewColumn
	*glib.GObject
}

func newTreeViewColumn(column *C.GtkTreeViewColumn) *TreeViewColumn {
	return &TreeViewColumn{
		GTreeViewColumn: column,
		GObject:        glib.ObjectFromNative(unsafe.Pointer(column)),
	}
}

func NewTreeViewColumn() *TreeViewColumn {
	return newTreeViewColumn(C.gtk_tree_view_column_new())
}

//TODO test
func NewTreeViewColumnWithAttributes2(title string, cell CellRendererLike, attributes ...interface{}) *TreeViewColumn {
	if len(attributes)%2 != 0 {
		log.Panic("Error in gtk.TreeViewColumnWithAttributes: last attribute isn't associated to a value, len(attributes) must be even")
	}
	ptrTitle := C.CString(title)
	defer C.free_string(ptrTitle)
	ret := newTreeViewColumn(C._gtk_tree_view_column_new_with_attribute(
		C.to_gcharptr(ptrTitle), cell.ToCellRenderer()))
	for i := 0; i < len(attributes)/2; i++ {
		attribute, ok := attributes[2*i].(string)
		if !ok {
			log.Panic("Error calling gtk.TreeViewColumnWithAttributes: property name must be a string")
		}
		ptrAttribute := C.CString(attribute)
		column, ok := attributes[2*i].(int)
		if !ok {
			log.Panic("Error calling gtk.TreeViewColumnWithAttributes: attributes column must be an int")
		}
		C.gtk_tree_view_column_add_attribute(ret.GTreeViewColumn,
			cell.ToCellRenderer(), C.to_gcharptr(ptrAttribute), C.gint(column))
	}
	return ret
}

func NewTreeViewColumnWithAttribute(title string, cell CellRendererLike) *TreeViewColumn {
	ptitle := C.CString(title)
	defer C.free_string(ptitle)
	return newTreeViewColumn(
		C._gtk_tree_view_column_new_with_attribute(C.to_gcharptr(ptitle), cell.ToCellRenderer()))
}
func NewTreeViewColumnWithAttributes(title string, cell CellRendererLike, attribute string, column int) *TreeViewColumn {
	ptitle := C.CString(title)
	defer C.free_string(ptitle)
	pattribute := C.CString(attribute)
	defer C.free_string(pattribute)
	return newTreeViewColumn(
		C._gtk_tree_view_column_new_with_attributes(C.to_gcharptr(ptitle), cell.ToCellRenderer(), C.to_gcharptr(pattribute), C.gint(column)))
}
func (v *TreeViewColumn) PackStart(cell CellRendererLike, expand bool) {
	C.gtk_tree_view_column_pack_start(v.GTreeViewColumn, cell.ToCellRenderer(), bool2gboolean(expand))
}
func (v *TreeViewColumn) PackEnd(cell CellRendererLike, expand bool) {
	C.gtk_tree_view_column_pack_end(v.GTreeViewColumn, cell.ToCellRenderer(), bool2gboolean(expand))
}
func (v *TreeViewColumn) Clear() {
	C.gtk_tree_view_column_clear(v.GTreeViewColumn)
}

// gtk_tree_view_column_get_cell_renderers //deprecated since 2.18

func (v *TreeViewColumn) AddAttribute(cell CellRendererLike, attribute string, column int) {
	ptr := C.CString(attribute)
	defer C.free_string(ptr)
	C.gtk_tree_view_column_add_attribute(v.GTreeViewColumn, cell.ToCellRenderer(), C.to_gcharptr(ptr), C.gint(column))
}

//void gtk_tree_view_column_set_attributes (GtkTreeViewColumn *tree_column, GtkCellRenderer *cell_renderer, ...) G_GNUC_NULL_TERMINATED;
//void gtk_tree_view_column_set_cell_data_func (GtkTreeViewColumn *tree_column, GtkCellRenderer *cell_renderer, GtkTreeCellDataFunc func, gpointer func_data, GDestroyNotify destroy);
//void gtk_tree_view_column_clear_attributes (GtkTreeViewColumn *tree_column, GtkCellRenderer *cell_renderer);

func (v *TreeViewColumn) SetSpacing(spacing int) {
	C.gtk_tree_view_column_set_spacing(v.GTreeViewColumn, C.gint(spacing))
}
func (v *TreeViewColumn) GetSpacing() int {
	return int(C.gtk_tree_view_column_get_spacing(v.GTreeViewColumn))
}
func (v *TreeViewColumn) SetVisible(resizable bool) {
	C.gtk_tree_view_column_set_visible(v.GTreeViewColumn, bool2gboolean(resizable))
}
func (v *TreeViewColumn) GetVisible() bool {
	return gboolean2bool(C.gtk_tree_view_column_get_visible (v.GTreeViewColumn))
}
func (v *TreeViewColumn) SetResizable(resizable bool) {
	C.gtk_tree_view_column_set_resizable(v.GTreeViewColumn, bool2gboolean(resizable))
}
func (v *TreeViewColumn) GetResizable() bool {
	return gboolean2bool(C.gtk_tree_view_column_get_resizable(v.GTreeViewColumn))
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
	C.gtk_tree_view_column_set_fixed_width(v.GTreeViewColumn, C.gint(w))
}
func (v *TreeViewColumn) SetMinWidth(w int) {
	C.gtk_tree_view_column_set_min_width(v.GTreeViewColumn, C.gint(w))
}
func (v *TreeViewColumn) GetMinWidth() int {
	return int(C.gtk_tree_view_column_get_min_width(v.GTreeViewColumn))
}
func (v *TreeViewColumn) SetMaxWidth(w int) {
	C.gtk_tree_view_column_set_max_width(v.GTreeViewColumn, C.gint(w))
}
func (v *TreeViewColumn) GetMaxWidth() int {
	return int(C.gtk_tree_view_column_get_max_width(v.GTreeViewColumn))
}
func (v *TreeViewColumn) Clicked() {
	C.gtk_tree_view_column_clicked(v.GTreeViewColumn)
}
func (v *TreeViewColumn) SetTitle(title string) {
	ptr := C.CString(title)
	defer C.free_string(ptr)
	C.gtk_tree_view_column_set_title(v.GTreeViewColumn, C.to_gcharptr(ptr))

}
func (v *TreeViewColumn) GetTitle() string {
	return C.GoString(C.to_charptr(C.gtk_tree_view_column_get_title(v.GTreeViewColumn)))
}

func (v *TreeViewColumn) SetExpand(expand bool) {
	C.gtk_tree_view_column_set_expand(v.GTreeViewColumn, bool2gboolean(expand))
}
func (v *TreeViewColumn) GetExpand() bool {
	return gboolean2bool(C.gtk_tree_view_column_get_expand(v.GTreeViewColumn))
}

func (v *TreeViewColumn) SetClickable(clickable bool) {
	C.gtk_tree_view_column_set_clickable(v.GTreeViewColumn, bool2gboolean(clickable))
}
func (v *TreeViewColumn) GetClickable() bool {
	return gboolean2bool(C.gtk_tree_view_column_get_clickable(v.GTreeViewColumn))
}
//void gtk_tree_view_column_set_widget (GtkTreeViewColumn *tree_column, GtkWidget *widget);
//GtkWidget *gtk_tree_view_column_get_widget (GtkTreeViewColumn *tree_column);
//void gtk_tree_view_column_set_alignment (GtkTreeViewColumn *tree_column, gfloat xalign);
//gfloat gtk_tree_view_column_get_alignment (GtkTreeViewColumn *tree_column);

func (v *TreeViewColumn) SetReorderable(reorderable bool) {
	C.gtk_tree_view_column_set_reorderable(v.GTreeViewColumn, bool2gboolean(reorderable))
}
func (v *TreeViewColumn) GetReorderable() bool {
	return gboolean2bool(C.gtk_tree_view_column_get_reorderable(v.GTreeViewColumn))
}

//void gtk_tree_view_column_set_sort_column_id (GtkTreeViewColumn *tree_column, gint sort_column_id);
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
//GtkTreeModel *gtk_tree_view_get_model (GtkTreeView *tree_view);

func (v *TreeView) SetModel(model TreeModelLike) {
	var tm *C.GtkTreeModel
	if model != nil {
		tm = model.cTreeModel()
	}
	C.gtk_tree_view_set_model(C.to_GtkTreeView(v.GWidget), tm)
}
func (v *TreeView) GetSelection() *TreeSelection {
	return &TreeSelection{C.gtk_tree_view_get_selection(C.to_GtkTreeView(v.GWidget))}
}

//GtkAdjustment *gtk_tree_view_get_hadjustment (GtkTreeView *tree_view);
//void gtk_tree_view_set_hadjustment (GtkTreeView *tree_view, GtkAdjustment *adjustment);
//GtkAdjustment *gtk_tree_view_get_vadjustment (GtkTreeView *tree_view);
//void gtk_tree_view_set_vadjustment (GtkTreeView *tree_view, GtkAdjustment *adjustment);
//gboolean gtk_tree_view_get_headers_visible (GtkTreeView *tree_view);

func (v *TreeView) SetHeadersVisible(flag bool) {
	C.gtk_tree_view_set_headers_visible(C.to_GtkTreeView(v.GWidget), bool2gboolean(flag))
}

//void gtk_tree_view_columns_autosize (GtkTreeView *tree_view);
//gboolean gtk_tree_view_get_headers_clickable (GtkTreeView *tree_view);
//void gtk_tree_view_set_headers_clickable (GtkTreeView *tree_view, gboolean setting);
//void gtk_tree_view_set_rules_hint (GtkTreeView *tree_view, gboolean setting);
//gboolean gtk_tree_view_get_rules_hint (GtkTreeView *tree_view);

func (v *TreeView) AppendColumn(c *TreeViewColumn) int {
	return int(C.gtk_tree_view_append_column(C.to_GtkTreeView(v.GWidget), c.GTreeViewColumn))
}

//gint gtk_tree_view_remove_column (GtkTreeView *tree_view, GtkTreeViewColumn *column);
//gint gtk_tree_view_insert_column (GtkTreeView *tree_view, GtkTreeViewColumn *column, gint position);
//gint gtk_tree_view_insert_column_with_attributes (GtkTreeView *tree_view, gint position, const gchar *title, GtkCellRenderer *cell, ...) G_GNUC_NULL_TERMINATED;
//gint gtk_tree_view_insert_column_with_data_func (GtkTreeView *tree_view, gint position, const gchar *title, GtkCellRenderer *cell, GtkTreeCellDataFunc func, gpointer data, GDestroyNotify dnotify);

func (v *TreeView) GetColumn(n int) *TreeViewColumn {
	return newTreeViewColumn(C.gtk_tree_view_get_column(C.to_GtkTreeView(v.GWidget), C.gint(n)))
}

//GList *gtk_tree_view_get_columns (GtkTreeView *tree_view);
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
	C.gtk_tree_view_scroll_to_cell(C.to_GtkTreeView(v.GWidget), path.GTreePath,
		pcol, bool2gboolean(use), C.gfloat(ra), C.gfloat(ca))
}
func (v *TreeView) SetCursor(path *TreePath, col *TreeViewColumn, se bool) {
	var pcol *C.GtkTreeViewColumn
	if nil == col {
		pcol = nil
	} else {
		pcol = col.GTreeViewColumn
	}
	C.gtk_tree_view_set_cursor(C.to_GtkTreeView(v.GWidget), path.GTreePath,
		pcol, bool2gboolean(se))
}

//void gtk_tree_view_set_cursor_on_cell (GtkTreeView *tree_view, GtkTreePath *path, GtkTreeViewColumn *focus_column, GtkCellRenderer *focus_cell, gboolean start_editing);

func (v *TreeView) GetCursor(path **TreePath, focus_column **TreeViewColumn) {
	*path = &TreePath{nil}
	if nil != focus_column {
		*focus_column = &TreeViewColumn{nil, nil}
		C.gtk_tree_view_get_cursor(C.to_GtkTreeView(v.GWidget), &(*path).GTreePath, &(*focus_column).GTreeViewColumn)
	} else {
		C.gtk_tree_view_get_cursor(C.to_GtkTreeView(v.GWidget), &(*path).GTreePath, nil)
	}
}

//void gtk_tree_view_row_activated (GtkTreeView *tree_view, GtkTreePath *path, GtkTreeViewColumn *column);

func (v *TreeView) ExpandAll() {
	C.gtk_tree_view_expand_all(C.to_GtkTreeView(v.GWidget))
}
func (v *TreeView) CollapseAll() {
	C.gtk_tree_view_collapse_all(C.to_GtkTreeView(v.GWidget))
}

//void gtk_tree_view_expand_to_path (GtkTreeView *tree_view, GtkTreePath *path);

func (v *TreeView) ExpandRow(path *TreePath, openall bool) bool {
	return gboolean2bool(C.gtk_tree_view_expand_row(C.to_GtkTreeView(v.GWidget), path.GTreePath, bool2gboolean(openall)))
}
func (v *TreeView) CollapseRow(path *TreePath) bool {
	return gboolean2bool(C.gtk_tree_view_collapse_row(C.to_GtkTreeView(v.GWidget), path.GTreePath))
}

//void gtk_tree_view_map_expanded_rows (GtkTreeView *tree_view, GtkTreeViewMappingFunc func, gpointer data);

func (v *TreeView) RowExpanded(path *TreePath) bool {
	return gboolean2bool(C.gtk_tree_view_row_expanded(C.to_GtkTreeView(v.GWidget), path.GTreePath))
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
func NewIconViewWithModel(model TreeModelLike) *IconView {
	var tm *C.GtkTreeModel
	if model != nil {
		tm = model.cTreeModel()
	}
	return &IconView{Container{Widget{
		C.gtk_icon_view_new_with_model(tm)}}}
}
func (v *IconView) GetModel() *TreeModel {
	return &TreeModel{
		C.gtk_icon_view_get_model(C.to_GtkIconView(v.GWidget))}
}
func (v *IconView) SetModel(model TreeModelLike) {
	var tm *C.GtkTreeModel
	if model != nil {
		tm = model.cTreeModel()
	}
	C.gtk_icon_view_set_model(C.to_GtkIconView(v.GWidget), tm)
}
func (v *IconView) GetTextColumn() int {
	return int(C.gtk_icon_view_get_text_column(C.to_GtkIconView(v.GWidget)))
}
func (v *IconView) SetTextColumn(text_column int) {
	C.gtk_icon_view_set_text_column(C.to_GtkIconView(v.GWidget), C.gint(text_column))
}
func (v *IconView) GetMarkupColumn() int {
	return int(C.gtk_icon_view_get_markup_column(C.to_GtkIconView(v.GWidget)))
}
func (v *IconView) SetMarkupColumn(markup_column int) {
	C.gtk_icon_view_set_markup_column(C.to_GtkIconView(v.GWidget), C.gint(markup_column))
}
func (v *IconView) GetPixbufColumn() int {
	return int(C.gtk_icon_view_get_pixbuf_column(C.to_GtkIconView(v.GWidget)))
}
func (v *IconView) SetPixbufColumn(pixbuf_column int) {
	C.gtk_icon_view_set_pixbuf_column(C.to_GtkIconView(v.GWidget), C.gint(pixbuf_column))
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
	C.gtk_icon_view_scroll_to_path(C.to_GtkIconView(v.GWidget), path.GTreePath,
		bool2gboolean(use), C.gfloat(ra), C.gfloat(ca))
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

// gtk_tree_sortable_sort_column_changed
// gtk_tree_sortable_get_sort_column_id
// gtk_tree_sortable_set_sort_column_id
// gtk_tree_sortable_set_sort_func
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
type CellRendererLike interface {
	ToCellRenderer() *C.GtkCellRenderer
}
type CellRenderer struct {
	GCellRenderer *C.GtkCellRenderer
	CellRendererLike
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
	return &CellRendererAccel{CellRenderer{
		C.gtk_cell_renderer_accel_new(), nil}}
}

//-----------------------------------------------------------------------
// GtkCellRendererCombo
//-----------------------------------------------------------------------
type CellRendererCombo struct {
	CellRenderer
}

func NewCellRendererCombo() *CellRendererCombo {
	return &CellRendererCombo{CellRenderer{
		C.gtk_cell_renderer_combo_new(), nil}}
}

//-----------------------------------------------------------------------
// GtkCellRendererPixbuf
//-----------------------------------------------------------------------
type CellRendererPixbuf struct {
	CellRenderer
}

func NewCellRendererPixbuf() *CellRendererPixbuf {
	return &CellRendererPixbuf{CellRenderer{
		C.gtk_cell_renderer_pixbuf_new(), nil}}
}

//-----------------------------------------------------------------------
// GtkCellRendererProgress
//-----------------------------------------------------------------------
type CellRendererProgress struct {
	CellRenderer
}

func NewCellRendererProgress() *CellRendererProgress {
	return &CellRendererProgress{CellRenderer{
		C.gtk_cell_renderer_progress_new(), nil}}
}

//-----------------------------------------------------------------------
// GtkCellRendererSpin
//-----------------------------------------------------------------------
type CellRendererSpin struct {
	CellRenderer
}

func NewCellRendererSpin() *CellRendererSpin {
	return &CellRendererSpin{CellRenderer{
		C.gtk_cell_renderer_spin_new(), nil}}
}

//-----------------------------------------------------------------------
// GtkCellRendererText
//-----------------------------------------------------------------------
type CellRendererText struct {
	CellRenderer
}

func NewCellRendererText() *CellRendererText {
	return &CellRendererText{CellRenderer{
		C.gtk_cell_renderer_text_new(), nil}}
}
func (v *CellRendererText) SetFixedHeightFromFont(number_of_rows int) {
	C.gtk_cell_renderer_text_set_fixed_height_from_font(C.to_GtkCellRendererText(v.GCellRenderer), C.gint(number_of_rows))
}

//-----------------------------------------------------------------------
// GtkCellRendererToggle
//-----------------------------------------------------------------------
type CellRendererToggle struct {
	CellRenderer
}

func NewCellRendererToggle() *CellRendererToggle {
	return &CellRendererToggle{CellRenderer{
		C.gtk_cell_renderer_toggle_new(), nil}}
}
func (v *CellRendererToggle) GetRadio() bool {
	return gboolean2bool(C.gtk_cell_renderer_toggle_get_radio(C.to_GtkCellRendererToggle(v.GCellRenderer)))
}
func (v *CellRendererToggle) SetRadio(radio bool) {
	C.gtk_cell_renderer_toggle_set_radio(C.to_GtkCellRendererToggle(v.GCellRenderer), bool2gboolean(radio))
}
func (v *CellRendererToggle) GetActive() bool {
	return gboolean2bool(C.gtk_cell_renderer_toggle_get_active(C.to_GtkCellRendererToggle(v.GCellRenderer)))
}
func (v *CellRendererToggle) SetActive(active bool) {
	C.gtk_cell_renderer_toggle_set_active(C.to_GtkCellRendererToggle(v.GCellRenderer), bool2gboolean(active))
}
func (v *CellRendererToggle) GetActivatable() bool {
	panic_if_version_older(2, 18, 0, "gtk_cell_renderer_toggle_get_activatable()")
	return gboolean2bool(C._gtk_cell_renderer_toggle_get_activatable(C.to_GtkCellRendererToggle(v.GCellRenderer)))
}
func (v *CellRendererToggle) SetActivatable(activatable bool) {
	panic_if_version_older(2, 18, 0, "gtk_cell_renderer_toggle_set_activatable()")
	C._gtk_cell_renderer_toggle_set_activatable(C.to_GtkCellRendererToggle(v.GCellRenderer), bool2gboolean(activatable))
}

//-----------------------------------------------------------------------
// GtkCellRendererSpinner
//-----------------------------------------------------------------------
type CellRendererSpinner struct {
	CellRenderer
}

func NewCellRendererSpinner() *CellRendererSpinner {
	panic_if_version_older(2, 20, 0, "gtk_cell_renderer_spinner_new()")
	return &CellRendererSpinner{CellRenderer{
		C._gtk_cell_renderer_spinner_new(), nil}}
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
	cliststore := C.gtk_list_store_newv(C.gint(len(v)), types)
	return &ListStore{TreeModel{
		C.to_GtkTreeModelFromListStore(cliststore)}, cliststore}
}

//void gtk_list_store_set_column_types (GtkListStore *list_store, gint n_columns, GType *types);

func (v *ListStore) Set(iter *TreeIter, a ...interface{}) {
	for r := range a {
		v.SetValue(iter, r, a[r])
	}
}
func (v *ListStore) SetValue(iter *TreeIter, column int, a interface{}) {
	gv := glib.GValueFromNative(a)
	if gv != nil {
		C.gtk_list_store_set_value(v.GListStore, &iter.GTreeIter, C.gint(column), C.to_GValueptr(unsafe.Pointer(gv)))
	} else {
		if pv, ok := a.(*[0]uint8); ok {
			C._gtk_list_store_set_ptr(v.GListStore, &iter.GTreeIter, C.gint(column), unsafe.Pointer(pv))
		} else {
			av := reflect.ValueOf(a)
			if av.CanAddr() {
				C._gtk_list_store_set_addr(v.GListStore, &iter.GTreeIter, C.gint(column), unsafe.Pointer(av.UnsafeAddr()))
			} else {
				C._gtk_list_store_set_addr(v.GListStore, &iter.GTreeIter, C.gint(column), unsafe.Pointer(&a))
			}
		}
	}
}
func (v *ListStore) Remove(iter *TreeIter) bool {
	return gboolean2bool(C.gtk_list_store_remove(v.GListStore, &iter.GTreeIter))
}
func (v *ListStore) Insert(iter *TreeIter, position int) {
	C.gtk_list_store_insert(v.GListStore, &iter.GTreeIter, C.gint(position))
}
func (v *ListStore) InsertBefore(iter *TreeIter, sibling *TreeIter) {
	C.gtk_list_store_insert_before(v.GListStore, &iter.GTreeIter, &sibling.GTreeIter)
}
func (v *ListStore) InsertAfter(iter *TreeIter, sibling *TreeIter) {
	C.gtk_list_store_insert_after(v.GListStore, &iter.GTreeIter, &sibling.GTreeIter)
}

//void gtk_list_store_insert_with_values (GtkListStore *list_store, GtkTreeIter *iter, gint position, ...);
//void gtk_list_store_insert_with_valuesv (GtkListStore *list_store, GtkTreeIter *iter, gint position, gint *columns, GValue *values, gint n_values);

func (v *ListStore) Prepend(iter *TreeIter) {
	C.gtk_list_store_prepend(v.GListStore, &iter.GTreeIter)
}
func (v *ListStore) Append(iter *TreeIter) {
	C.gtk_list_store_append(v.GListStore, &iter.GTreeIter)
}
func (v *ListStore) Clear() {
	C.gtk_list_store_clear(v.GListStore)
}
func (v *ListStore) IterIsValid(iter *TreeIter) bool {
	log.Println("Warning: ListStore.IterIsValid: This function is slow. Only use it for debugging and/or testing purposes.")
	return gboolean2bool(C.gtk_list_store_iter_is_valid(v.GListStore, &iter.GTreeIter))
}
func (v *ListStore) Reorder(i *int) {
	gi := C.gint(*i)
	C.gtk_list_store_reorder(v.GListStore, &gi)
	*i = int(gi)
}
func (v *ListStore) Swap(a *TreeIter, b *TreeIter) {
	C.gtk_list_store_swap(v.GListStore, &a.GTreeIter, &b.GTreeIter)
}
func (v *ListStore) MoveBefore(iter *TreeIter, position *TreeIter) {
	C.gtk_list_store_move_before(v.GListStore, &iter.GTreeIter, &position.GTreeIter)
}
func (v *ListStore) MoveAfter(iter *TreeIter, position *TreeIter) {
	C.gtk_list_store_move_after(v.GListStore, &iter.GTreeIter, &position.GTreeIter)
}

//TODO instead of using this methods to change between treemodel and liststore, is better to usa an interface GtkTreeModelLike
//nb: ListStore e TreeStore sono un TreeModel (implementano GtkTreeModel!)
/*func (v *GtkListStore) ToTreeModel() *GtkTreeModel {
	return &TreeModel{
		C.to_GtkTreeModelFromListStore(v.GListStore)}
}*/
/*func (v *GtkTreeModel) ToListStore() *GtkListStore {
	return &ListStore{
		C.to_GtkListStoreFromTreeModel(v.GTreeModel)}
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
	ctreestore := C.gtk_tree_store_newv(C.gint(len(v)), types)
	return &TreeStore{TreeModel{C.to_GtkTreeModelFromTreeStore(ctreestore)}, ctreestore}
}

// void gtk_tree_store_set_column_types (GtkTreeStore *tree_store, gint n_columns, GType *types); void gtk_tree_store_set_value (GtkTreeStore *tree_store, GtkTreeIter *iter, gint column, GValue *value);

func (v *TreeStore) Set(iter *TreeIter, a ...interface{}) {
	for r := range a {
		v.SetValue(iter, r, a[r])
	}
}
func (v *TreeStore) SetValue(iter *TreeIter, column int, a interface{}) {
	gv := glib.GValueFromNative(a)
	if gv != nil {
		C.gtk_tree_store_set_value(v.GTreeStore, &iter.GTreeIter, C.gint(column), C.to_GValueptr(unsafe.Pointer(gv)))
	} else {
		if pv, ok := a.(*[0]uint8); ok {
			C._gtk_tree_store_set_ptr(v.GTreeStore, &iter.GTreeIter, C.gint(column), unsafe.Pointer(pv))
		} else {
			av := reflect.ValueOf(a)
			if av.CanAddr() {
				C._gtk_tree_store_set_addr(v.GTreeStore, &iter.GTreeIter, C.gint(column), unsafe.Pointer(av.UnsafeAddr()))
			} else {
				C._gtk_tree_store_set_addr(v.GTreeStore, &iter.GTreeIter, C.gint(column), unsafe.Pointer(&a))
			}
		}
	}
}
func (v *TreeStore) Remove(iter *TreeIter) bool {
	return gboolean2bool(C.gtk_tree_store_remove(v.GTreeStore, &iter.GTreeIter))
}
func (v *TreeStore) Insert(iter *TreeIter, parent *TreeIter, position int) {
	C.gtk_tree_store_insert(v.GTreeStore, &iter.GTreeIter, &parent.GTreeIter, C.gint(position))
}
func (v *TreeStore) InsertBefore(iter *TreeIter, parent *TreeIter, sibling *TreeIter) {
	C.gtk_tree_store_insert_before(v.GTreeStore, &iter.GTreeIter, &parent.GTreeIter, &sibling.GTreeIter)
}
func (v *TreeStore) InsertAfter(iter *TreeIter, parent *TreeIter, sibling *TreeIter) {
	C.gtk_tree_store_insert_after(v.GTreeStore, &iter.GTreeIter, &parent.GTreeIter, &sibling.GTreeIter)
}

// void gtk_tree_store_insert_with_values (GtkTreeStore *tree_store, GtkTreeIter *iter, GtkTreeIter *parent, gint position, ...);
// void gtk_tree_store_insert_with_valuesv (GtkTreeStore *tree_store, GtkTreeIter *iter, GtkTreeIter *parent, gint position, gint *columns, GValue *values, gint n_values);

func (v *TreeStore) Prepend(iter *TreeIter, parent *TreeIter) {
	if parent == nil {
		C.gtk_tree_store_prepend(v.GTreeStore, &iter.GTreeIter, nil)
	} else {
		C.gtk_tree_store_prepend(v.GTreeStore, &iter.GTreeIter, &parent.GTreeIter)
	}
}
func (v *TreeStore) Append(iter *TreeIter, parent *TreeIter) {
	if parent == nil {
		C.gtk_tree_store_append(v.GTreeStore, &iter.GTreeIter, nil)
	} else {
		C.gtk_tree_store_append(v.GTreeStore, &iter.GTreeIter, &parent.GTreeIter)
	}
}

// gtk_tree_store_is_ancestor

func (v *TreeStore) ToTreeModel() *TreeModel {
	return &TreeModel{
		C.to_GtkTreeModelFromTreeStore(v.GTreeStore)}
}
func (v *TreeStore) IterDepth(iter *TreeIter) int {
	return int(C.gtk_tree_store_iter_depth(v.GTreeStore, &iter.GTreeIter))
}
func (v *TreeStore) Clear() {
	C.gtk_tree_store_clear(v.GTreeStore)
}
func (v *TreeStore) IterIsValid(iter *TreeIter) bool {
	log.Println("Warning: TreeStore.IterIsValid: This function is slow. Only use it for debugging and/or testing purposes.")
	return gboolean2bool(C.gtk_tree_store_iter_is_valid(v.GTreeStore, &iter.GTreeIter))
}
func (v *TreeStore) Reorder(iter *TreeIter, i *int) {
	gi := C.gint(*i)
	C.gtk_tree_store_reorder(v.GTreeStore, &iter.GTreeIter, &gi)
	*i = int(gi)
}
func (v *TreeStore) Swap(a *TreeIter, b *TreeIter) {
	C.gtk_tree_store_swap(v.GTreeStore, &a.GTreeIter, &b.GTreeIter)
}
func (v *TreeStore) MoveBefore(iter *TreeIter, position *TreeIter) {
	C.gtk_tree_store_move_before(v.GTreeStore, &iter.GTreeIter, &position.GTreeIter)
}
func (v *TreeStore) MoveAfter(iter *TreeIter, position *TreeIter) {
	C.gtk_tree_store_move_after(v.GTreeStore, &iter.GTreeIter, &position.GTreeIter)
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
	return &ComboBox{Bin{Container{Widget{
		C.gtk_combo_box_new_with_model(model.GTreeModel)}}}}
}
func NewComboBoxWithModelAndEntry(model *TreeModel) *ComboBox {
	deprecated_since(2, 24, 0, "gtk_combo_box_new_with_model_and_entry()")
	return &ComboBox{Bin{Container{Widget{
		C._gtk_combo_box_new_with_model_and_entry(model.GTreeModel)}}}}
}
func (v *ComboBox) GetWrapWidth() int {
	return int(C.gtk_combo_box_get_wrap_width(C.to_GtkComboBox(v.GWidget)))
}
func (v *ComboBox) SetWrapWidth(width int) {
	C.gtk_combo_box_set_wrap_width(C.to_GtkComboBox(v.GWidget), C.gint(width))
}
func (v *ComboBox) GetRowSpanColumn() int {
	return int(C.gtk_combo_box_get_row_span_column(C.to_GtkComboBox(v.GWidget)))
}
func (v *ComboBox) SetRowSpanColumn(row_span int) {
	C.gtk_combo_box_set_row_span_column(C.to_GtkComboBox(v.GWidget), C.gint(row_span))
}
func (v *ComboBox) GetColumnSpanColumn() int {
	return int(C.gtk_combo_box_get_column_span_column(C.to_GtkComboBox(v.GWidget)))
}
func (v *ComboBox) SetColumnSpanColumn(column_span int) {
	C.gtk_combo_box_set_column_span_column(C.to_GtkComboBox(v.GWidget), C.gint(column_span))
}
func (v *ComboBox) GetActive() int {
	return int(C.gtk_combo_box_get_active(C.to_GtkComboBox(v.GWidget)))
}
func (v *ComboBox) SetActive(width int) {
	C.gtk_combo_box_set_active(C.to_GtkComboBox(v.GWidget), C.gint(width))
}
func (v *ComboBox) GetActiveIter(iter *TreeIter) bool {
	return gboolean2bool(C.gtk_combo_box_get_active_iter(C.to_GtkComboBox(v.GWidget), &iter.GTreeIter))
}
func (v *ComboBox) SetActiveIter(iter *TreeIter) {
	C.gtk_combo_box_set_active_iter(C.to_GtkComboBox(v.GWidget), &iter.GTreeIter)
}
func (v *ComboBox) GetModel() *TreeModel {
	return &TreeModel{
		C.gtk_combo_box_get_model(C.to_GtkComboBox(v.GWidget))}
}
func (v *ComboBox) SetModel(model *TreeModel) {
	C.gtk_combo_box_set_model(C.to_GtkComboBox(v.GWidget), model.GTreeModel)
}

//Deprecated since 2.24. Use GtkComboBoxText.
func NewComboBoxNewText() *ComboBox {
	deprecated_since(2, 24, 0, "gtk_combo_box_new_text()")
	return &ComboBox{Bin{Container{Widget{
		C.gtk_combo_box_new_text()}}}}
}

//Deprecated since 2.24. Use GtkComboBoxText.
func (v *ComboBox) AppendText(text string) {
	deprecated_since(2, 24, 0, "gtk_combo_box_append_text()")
	ptr := C.CString(text)
	defer C.free_string(ptr)
	C.gtk_combo_box_append_text(C.to_GtkComboBox(v.GWidget), C.to_gcharptr(ptr))
}

//Deprecated since 2.24. Use GtkComboBoxText.
func (v *ComboBox) InsertText(text string, position int) {
	deprecated_since(2, 24, 0, "gtk_combo_box_insert_text()")
	ptr := C.CString(text)
	defer C.free_string(ptr)
	C.gtk_combo_box_insert_text(C.to_GtkComboBox(v.GWidget), C.gint(position), C.to_gcharptr(ptr))
}

//Deprecated since 2.24. Use GtkComboBoxText.
func (v *ComboBox) PrependText(text string) {
	deprecated_since(2, 24, 0, "gtk_combo_box_prepend_text()")
	ptr := C.CString(text)
	defer C.free_string(ptr)
	C.gtk_combo_box_prepend_text(C.to_GtkComboBox(v.GWidget), C.to_gcharptr(ptr))
}

//Deprecated since 2.24. Use GtkComboBoxText.
func (v *ComboBox) RemoveText(position int) {
	deprecated_since(2, 24, 0, "gtk_combo_box_remove_text()")
	C.gtk_combo_box_remove_text(C.to_GtkComboBox(v.GWidget), C.gint(position))
}

//Deprecated since 2.24. Use GtkComboBoxText or, if combo box contains an entry,
// get text directly from GtkEntry.
func (v *ComboBox) GetActiveText() string {
	deprecated_since(2, 24, 0, "gtk_combo_box_get_active_text()")
	return C.GoString(C.to_charptr(C.gtk_combo_box_get_active_text(C.to_GtkComboBox(v.GWidget))))
}
func (v *ComboBox) Popup() {
	C.gtk_combo_box_popup(C.to_GtkComboBox(v.GWidget))
}
func (v *ComboBox) Popdown() {
	C.gtk_combo_box_popdown(C.to_GtkComboBox(v.GWidget))
}

// gtk_combo_box_get_popup_accessible
// gtk_combo_box_get_row_separator_func
// gtk_combo_box_set_row_separator_func

func (v *ComboBox) SetAddTearoffs(add_tearoffs bool) {
	C.gtk_combo_box_set_add_tearoffs(C.to_GtkComboBox(v.GWidget), bool2gboolean(add_tearoffs))
}
func (v *ComboBox) GetAddTearoffs() bool {
	return gboolean2bool(C.gtk_combo_box_get_add_tearoffs(C.to_GtkComboBox(v.GWidget)))
}
func (v *ComboBox) SetTitle(title string) {
	ptr := C.CString(title)
	defer C.free_string(ptr)
	C.gtk_combo_box_set_title(C.to_GtkComboBox(v.GWidget), C.to_gcharptr(ptr))
}
func (v *ComboBox) GetTitle() string {
	return C.GoString(C.to_charptr(C.gtk_combo_box_get_title(C.to_GtkComboBox(v.GWidget))))
}
func (v *ComboBox) SetFocusOnClick(focus_on_click bool) {
	C.gtk_combo_box_set_focus_on_click(C.to_GtkComboBox(v.GWidget), bool2gboolean(focus_on_click))
}
func (v *ComboBox) GetFocusOnClick() bool {
	return gboolean2bool(C.gtk_combo_box_get_focus_on_click(C.to_GtkComboBox(v.GWidget)))
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
	return &ComboBoxText{ComboBox{Bin{Container{Widget{
		C._gtk_combo_box_text_new()}}}}}
}
func NewComboBoxTextWithEntry() *ComboBoxText {
	return &ComboBoxText{ComboBox{Bin{Container{Widget{
		C._gtk_combo_box_text_new_with_entry()}}}}}
}
func (v *ComboBoxText) AppendText(text string) {
	ptr := C.CString(text)
	defer C.free_string(ptr)
	C._gtk_combo_box_text_append_text(C.to_GtkComboBoxText(v.GWidget), C.to_gcharptr(ptr))
}
func (v *ComboBoxText) InsertText(position int, text string) {
	ptr := C.CString(text)
	defer C.free_string(ptr)
	C._gtk_combo_box_text_insert_text(C.to_GtkComboBoxText(v.GWidget), C.gint(position), C.to_gcharptr(ptr))
}
func (v *ComboBoxText) PrependText(text string) {
	ptr := C.CString(text)
	defer C.free_string(ptr)
	C._gtk_combo_box_text_prepend_text(C.to_GtkComboBoxText(v.GWidget), C.to_gcharptr(ptr))
}
func (v *ComboBoxText) Remove(position int) {
	C._gtk_combo_box_text_remove(C.to_GtkComboBoxText(v.GWidget), C.gint(position))
}
func (v *ComboBoxText) GetActiveText() string {
	return C.GoString(C.to_charptr(C._gtk_combo_box_text_get_active_text(C.to_GtkComboBoxText(v.GWidget))))
}

//-----------------------------------------------------------------------
// GtkComboBoxEntry
//-----------------------------------------------------------------------
type ComboBoxEntry struct {
	ComboBox
}

func NewComboBoxEntry() *ComboBoxEntry {
	deprecated_since(2, 24, 0, "gtk_combo_box_entry_new()")
	return &ComboBoxEntry{ComboBox{Bin{Container{Widget{
		C.gtk_combo_box_entry_new()}}}}}
}

// gtk_combo_box_entry_new_with_model

func NewComboBoxEntryNewText() *ComboBoxEntry {
	deprecated_since(2, 24, 0, "gtk_combo_box_entry_new_text()")
	return &ComboBoxEntry{ComboBox{Bin{Container{Widget{
		C.gtk_combo_box_entry_new_text()}}}}}
}
func (v *ComboBoxEntry) GetTextColumn() int {
	deprecated_since(2, 24, 0, "gtk_combo_box_entry_get_text_column()")
	return int(C.gtk_combo_box_entry_get_text_column(C.to_GtkComboBoxEntry(v.GWidget)))
}
func (v *ComboBoxEntry) SetTextColumn(text_column int) {
	deprecated_since(2, 24, 0, "gtk_combo_box_entry_set_text_column()")
	C.gtk_combo_box_entry_set_text_column(C.to_GtkComboBoxEntry(v.GWidget), C.gint(text_column))
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
func (v *Menu) Append(child WidgetLike) {
	C.gtk_menu_shell_append(C.to_GtkMenuShell(v.GWidget), child.ToNative())
}

//TODO remove when GtkMenuShell is done
func (v *Menu) Prepend(child WidgetLike) {
	C.gtk_menu_shell_prepend(C.to_GtkMenuShell(v.GWidget), child.ToNative())
}

//TODO remove when GtkMenuShell is done
func (v *Menu) Insert(child WidgetLike, position int) {
	C.gtk_menu_shell_insert(C.to_GtkMenuShell(v.GWidget), child.ToNative(), C.gint(position))
}

// void gtk_menu_reorder_child(GtkMenu *menu, GtkWidget *child, gint position);
// void gtk_menu_attach(GtkMenu *menu, GtkWidget *child, guint left_attach, guint right_attach, guint top_attach, guint bottom_attach);

func (v *Menu) Popup(parent_menu_shell, parent_menu_item WidgetLike, f MenuPositionFunc, data interface{}, button uint, active_item uint) {
	var pms, pmi *C.GtkWidget
	if parent_menu_shell != nil {
		pms = parent_menu_shell.ToNative()
	}
	if parent_menu_item != nil {
		pmi = parent_menu_item.ToNative()
	}
	C._gtk_menu_popup(v.GWidget, pms, pmi, unsafe.Pointer(&MenuPositionFuncInfo{v, f, data}), C.guint(button), C.guint32(active_item))
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
	return gboolean2bool(C.gtk_menu_get_tearoff_state(C.to_GtkMenu(v.GWidget)))
}
func (v *Menu) SetReserveToggleSize(b bool) {
	panic_if_version_older(2, 18, 0, "gtk_menu_set_reserve_toggle_size()")
	C._gtk_menu_set_reserve_toggle_size(C.to_GtkMenu(v.GWidget), bool2gboolean(b))
}
func (v *Menu) GetReserveToggleSize() bool {
	panic_if_version_older(2, 18, 0, "gtk_menu_get_reserve_toggle_size()")
	return gboolean2bool(C._gtk_menu_get_reserve_toggle_size(C.to_GtkMenu(v.GWidget)))
}
func (v *Menu) Popdown() {
	C.gtk_menu_popdown(C.to_GtkMenu(v.GWidget))
}
func (v *Menu) Reposition() {
	C.gtk_menu_reposition(C.to_GtkMenu(v.GWidget))
}
func (v *Menu) GetActive() *Widget {
	return &Widget{C.gtk_menu_get_active(C.to_GtkMenu(v.GWidget))}
}

// void gtk_menu_set_active (GtkMenu *menu, guint index_);

func (v *Menu) SetTearoffState(b bool) {
	C.gtk_menu_set_tearoff_state(C.to_GtkMenu(v.GWidget), bool2gboolean(b))
}

// void gtk_menu_attach_to_widget (GtkMenu *menu, GtkWidget *attach_widget, GtkMenuDetachFunc detacher);

func (v *Menu) Detach() {
	C.gtk_menu_detach(C.to_GtkMenu(v.GWidget))
}
func (v *Menu) GetAttachWidget() *Widget {
	return &Widget{C.gtk_menu_get_attach_widget(C.to_GtkMenu(v.GWidget))}
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
	gmpfigo := (*MenuPositionFuncInfo)(gmpfi.data)
	if gmpfigo.f == nil {
		return
	}
	x := int(gmpfi.x)
	y := int(gmpfi.y)
	push_in := gboolean2bool(gmpfi.push_in)
	gmpfigo.f(gmpfigo.menu, &x, &y, &push_in, gmpfigo.data)
	gmpfi.x = C.gint(x)
	gmpfi.y = C.gint(y)
	gmpfi.push_in = bool2gboolean(push_in)
}

//-----------------------------------------------------------------------
// GtkMenuBar
//-----------------------------------------------------------------------
type PackDirection int

const (
	PACK_DIRECTION_LTR PackDirection = 0
	PACK_DIRECTION_RTL PackDirection = 1
	PACK_DIRECTION_TTB PackDirection = 2
	PACK_DIRECTION_BTT PackDirection = 3
)

type MenuBar struct {
	Widget
}

func NewMenuBar() *MenuBar {
	return &MenuBar{Widget{C.gtk_menu_bar_new()}}
}
func (v *MenuBar) SetPackDirection(pack_dir PackDirection) {
	C.gtk_menu_bar_set_pack_direction(C.to_GtkMenuBar(v.GWidget), C.GtkPackDirection(pack_dir))
}
func (v *MenuBar) GetPackDirection() PackDirection {
	return PackDirection(C.gtk_menu_bar_get_pack_direction(C.to_GtkMenuBar(v.GWidget)))
}
func (v *MenuBar) SetChildPackDirection(pack_dir PackDirection) {
	C.gtk_menu_bar_set_child_pack_direction(C.to_GtkMenuBar(v.GWidget), C.GtkPackDirection(pack_dir))
}
func (v *MenuBar) GetChildPackDirection() PackDirection {
	return PackDirection(C.gtk_menu_bar_get_child_pack_direction(C.to_GtkMenuBar(v.GWidget)))
}

//TODO da rimuovere, creare GtkMenuShell e usarlo come anonymous field per GtkMenu
func (v *MenuBar) Append(child WidgetLike) {
	C.gtk_menu_shell_append(C.to_GtkMenuShell(v.GWidget), child.ToNative())
}

//TODO da rimuovere, creare GtkMenuShell e usarlo come anonymous field per GtkMenu
func (v *MenuBar) Prepend(child WidgetLike) {
	C.gtk_menu_shell_prepend(C.to_GtkMenuShell(v.GWidget), child.ToNative())
}

//TODO da rimuovere, creare GtkMenuShell e usarlo come anonymous field per GtkMenu
func (v *MenuBar) Insert(child WidgetLike, position int) {
	C.gtk_menu_shell_insert(C.to_GtkMenuShell(v.GWidget), child.ToNative(), C.gint(position))
}

//-----------------------------------------------------------------------
// GtkMenuItem
//-----------------------------------------------------------------------
type MenuItem struct {
	Item
}

func NewMenuItem() *MenuItem {
	return &MenuItem{Item{Bin{Container{Widget{
		C.gtk_menu_item_new()}}}}}
}
func NewMenuItemWithLabel(label string) *MenuItem {
	ptr := C.CString(label)
	defer C.free_string(ptr)
	return &MenuItem{Item{Bin{Container{Widget{
		C.gtk_menu_item_new_with_label(C.to_gcharptr(ptr))}}}}}
}
func NewMenuItemWithMnemonic(label string) *MenuItem {
	ptr := C.CString(label)
	defer C.free_string(ptr)
	return &MenuItem{Item{Bin{Container{Widget{
		C.gtk_menu_item_new_with_mnemonic(C.to_gcharptr(ptr))}}}}}
}
func (v *MenuItem) SetRightJustified(b bool) {
	C.gtk_menu_item_set_right_justified(C.to_GtkMenuItem(v.GWidget), bool2gboolean(b))
}
func (v *MenuItem) GetRightJustified() bool {
	return gboolean2bool(C.gtk_menu_item_get_right_justified(C.to_GtkMenuItem(v.GWidget)))
}

// G_CONST_RETURN gchar *gtk_menu_item_get_label(GtkMenuItem *menu_item);
// void gtk_menu_item_set_label(GtkMenuItem *menu_item, const gchar *label);

func (v *MenuItem) GetUseUnderline() bool {
	return gboolean2bool(C.gtk_menu_item_get_use_underline(C.to_GtkMenuItem(v.GWidget)))
}
func (v *MenuItem) SetUseUnderline(setting bool) {
	C.gtk_menu_item_set_use_underline(C.to_GtkMenuItem(v.GWidget), bool2gboolean(setting))
}
func (v *MenuItem) SetSubmenu(w WidgetLike) {
	C.gtk_menu_item_set_submenu(C.to_GtkMenuItem(v.GWidget), w.ToNative())
}
func (v *MenuItem) GetSubmenu() *Widget {
	return &Widget{C.gtk_menu_item_get_submenu(C.to_GtkMenuItem(v.GWidget))}
}

//Deprecated since 2.12. Use SetSubmenu() instead.
func (v *MenuItem) RemoveSubmenu() {
	deprecated_since(2, 12, 0, "gtk_menu_item_remove_submenu()")
	C.gtk_menu_item_remove_submenu(C.to_GtkMenuItem(v.GWidget))
}

// void gtk_menu_item_set_accel_path(GtkMenuItem *menu_item, const gchar *accel_path);
// G_CONST_RETURN gchar* gtk_menu_item_get_accel_path(GtkMenuItem *menu_item);

func (v *MenuItem) Select() {
	C.gtk_menu_item_select(C.to_GtkMenuItem(v.GWidget))
}
func (v *MenuItem) Deselect() {
	C.gtk_menu_item_deselect(C.to_GtkMenuItem(v.GWidget))
}
func (v *MenuItem) Activate() {
	C.gtk_menu_item_activate(C.to_GtkMenuItem(v.GWidget))
}
func (v *MenuItem) ToggleSizeRequest(i *int) {
	gi := C.gint(*i)
	C.gtk_menu_item_toggle_size_request(C.to_GtkMenuItem(v.GWidget), &gi)
	*i = int(gi)
}
func (v *MenuItem) ToggleSizeAllocate(i int) {
	C.gtk_menu_item_toggle_size_allocate(C.to_GtkMenuItem(v.GWidget), C.gint(i))
}

//-----------------------------------------------------------------------
// GtkImageMenuItem
//-----------------------------------------------------------------------

// gtk_image_menu_item_set_image
// gtk_image_menu_item_get_image
// gtk_image_menu_item_new
// gtk_image_menu_item_new_from_stock
// gtk_image_menu_item_new_with_label
// gtk_image_menu_item_new_with_mnemonic
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
	if group != nil {
		return &RadioMenuItem{CheckMenuItem{MenuItem{Item{Bin{Container{Widget{
			C.gtk_radio_menu_item_new(C.to_gslist(unsafe.Pointer(group.ToSList())))}}}}}}}
	}
	return &RadioMenuItem{CheckMenuItem{MenuItem{Item{Bin{Container{Widget{
		C.gtk_radio_menu_item_new(nil)}}}}}}}
}

func NewRadioMenuItemWithLabel(group *glib.SList, label string) *RadioMenuItem {
	ptr := C.CString(label)
	defer C.free_string(ptr)
	if group != nil {
		return &RadioMenuItem{CheckMenuItem{MenuItem{Item{Bin{Container{Widget{
			C.gtk_radio_menu_item_new_with_label(C.to_gslist(unsafe.Pointer(group.ToSList())), C.to_gcharptr(ptr))}}}}}}}
	}
	return &RadioMenuItem{CheckMenuItem{MenuItem{Item{Bin{Container{Widget{
		C.gtk_radio_menu_item_new_with_label(nil, C.to_gcharptr(ptr))}}}}}}}
}

// gtk_radio_menu_item_new_with_mnemonic
// gtk_radio_menu_item_new_from_widget
// gtk_radio_menu_item_new_with_label_from_widget
// gtk_radio_menu_item_new_with_mnemonic_from_widget
// gtk_radio_menu_item_group

func (v *RadioMenuItem) SetGroup(group *glib.SList) {
	if group != nil {
		C.gtk_radio_menu_item_set_group(C.to_GtkRadioMenuItem(v.GWidget), C.to_gslist(unsafe.Pointer(group)))
	} else {
		C.gtk_radio_menu_item_set_group(C.to_GtkRadioMenuItem(v.GWidget), nil)
	}
}

func (v *RadioMenuItem) GetGroup() *glib.SList {
	return glib.SListFromNative(unsafe.Pointer(C.gtk_radio_menu_item_get_group(C.to_GtkRadioMenuItem(v.GWidget))))
}

//-----------------------------------------------------------------------
// GtkCheckMenuItem
//-----------------------------------------------------------------------
type CheckMenuItem struct {
	MenuItem
}

func NewCheckMenuItem() *CheckMenuItem {
	return &CheckMenuItem{MenuItem{Item{Bin{Container{Widget{
		C.gtk_check_menu_item_new()}}}}}}
}
func NewCheckMenuItemWithLabel(label string) *CheckMenuItem {
	ptr := C.CString(label)
	defer C.free_string(ptr)
	return &CheckMenuItem{MenuItem{Item{Bin{Container{Widget{
		C.gtk_check_menu_item_new_with_label(C.to_gcharptr(ptr))}}}}}}
}
func NewCheckMenuItemWithMnemonic(label string) *CheckMenuItem {
	ptr := C.CString(label)
	defer C.free_string(ptr)
	return &CheckMenuItem{MenuItem{Item{Bin{Container{Widget{
		C.gtk_check_menu_item_new_with_mnemonic(C.to_gcharptr(ptr))}}}}}}
}
func (v *CheckMenuItem) GetActive() bool {
	return gboolean2bool(C.gtk_check_menu_item_get_active(C.to_GtkCheckMenuItem(v.GWidget)))
}
func (v *CheckMenuItem) SetActive(setting bool) {
	C.gtk_check_menu_item_set_active(C.to_GtkCheckMenuItem(v.GWidget), bool2gboolean(setting))
}
func (v *CheckMenuItem) Toggled() {
	C.gtk_check_menu_item_toggled(C.to_GtkCheckMenuItem(v.GWidget))
}
func (v *CheckMenuItem) GetInconsistent() bool {
	return gboolean2bool(C.gtk_check_menu_item_get_inconsistent(C.to_GtkCheckMenuItem(v.GWidget)))
}
func (v *CheckMenuItem) SetInconsistent(setting bool) {
	C.gtk_check_menu_item_set_inconsistent(C.to_GtkCheckMenuItem(v.GWidget), bool2gboolean(setting))
}
func (v *CheckMenuItem) SetDrawAsRadio(setting bool) {
	C.gtk_check_menu_item_set_draw_as_radio(C.to_GtkCheckMenuItem(v.GWidget), bool2gboolean(setting))
}
func (v *CheckMenuItem) GetDrawAsRadio() bool {
	return gboolean2bool(C.gtk_check_menu_item_get_draw_as_radio(C.to_GtkCheckMenuItem(v.GWidget)))
}

//-----------------------------------------------------------------------
// GtkSeparatorMenuItem
//-----------------------------------------------------------------------
type SeparatorMenuItem struct {
	MenuItem
}

func NewSeparatorMenuItem() *SeparatorMenuItem {
	return &SeparatorMenuItem{MenuItem{Item{Bin{Container{Widget{
		C.gtk_separator_menu_item_new()}}}}}}
}

//-----------------------------------------------------------------------
// GtkTearoffMenuItem
//-----------------------------------------------------------------------
type TearoffMenuItem struct {
	MenuItem
}

func NewTearoffMenuItem() *TearoffMenuItem {
	return &TearoffMenuItem{MenuItem{Item{Bin{Container{Widget{
		C.gtk_tearoff_menu_item_new()}}}}}}
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
	ORIENTATION_HORIZONTAL     = 0
	ORIENTATION_VERTICAL       = 1
)

type ToolbarStyle int

const (
	TOOLBAR_ICONS        = 0
	TOOLBAR_TEXT         = 1
	TOOLBAR_BOTH         = 2
	TOOLBAR_BOTH_HORIZ   = 3
)

type Toolbar struct {
	Container
	items map[*C.GtkToolItem]IToolItem
}

func NewToolbar() *Toolbar {
	return &Toolbar{Container{Widget{C.gtk_toolbar_new()}}, make(map[*C.GtkToolItem]IToolItem)}
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

func (v *Toolbar) Insert(item IToolItem, pos int) {
	p_tool_item := C.to_GtkToolItem(item.AsToolItem().GWidget)
	v.items[p_tool_item] = item
	C.gtk_toolbar_insert(C.to_GtkToolbar(v.GWidget), p_tool_item, C.gint(pos))
}
func (v *Toolbar) GetItemIndex(item IToolItem) int {
	return int(C.gtk_toolbar_get_item_index(C.to_GtkToolbar(v.GWidget), C.to_GtkToolItem(item.AsToolItem().GWidget)))
}
func (v *Toolbar) GetNItems() int {
	return int(C.gtk_toolbar_get_n_items(C.to_GtkToolbar(v.GWidget)))
}
func (v *Toolbar) GetNthItem(n int) IToolItem {	
	p_tool_item := C.gtk_toolbar_get_nth_item(C.to_GtkToolbar(v.GWidget), C.gint(n))
	if p_tool_item == nil {
		panic("Toolbar.GetNthItem: index out of bounds")
	}
	if _, ok := v.items[p_tool_item]; !ok {	
		panic("Toolbar.GetNthItem: interface not found in map")
	}
	return v.items[p_tool_item]	
}
func (v *Toolbar) GetDropIndex(x, y int) int {
	return int(C.gtk_toolbar_get_drop_index(C.to_GtkToolbar(v.GWidget), C.gint(x), C.gint(y)))
}
func (v *Toolbar) SetDropHighlightItem(item IToolItem, index int) {
	C.gtk_toolbar_set_drop_highlight_item(C.to_GtkToolbar(v.GWidget), C.to_GtkToolItem(item.AsToolItem().GWidget), C.gint(index))
}
func (v *Toolbar) SetShowArrow(show_arrow bool) {
	C.gtk_toolbar_set_show_arrow(C.to_GtkToolbar(v.GWidget), bool2gboolean(show_arrow))
}
func (v *Toolbar) SetOrientation(orientation Orientation) {
	C.gtk_toolbar_set_orientation(C.to_GtkToolbar(v.GWidget), C.GtkOrientation(orientation))
}
func (v *Toolbar) SetTooltips(enable bool) {
	C.gtk_toolbar_set_tooltips(C.to_GtkToolbar(v.GWidget), bool2gboolean(enable))
}
func (v *Toolbar) UnsetIconSize() {
	C.gtk_toolbar_unset_icon_size(C.to_GtkToolbar(v.GWidget))
}
func (v *Toolbar) GetShowArrow() bool {
	return gboolean2bool(C.gtk_toolbar_get_show_arrow(C.to_GtkToolbar(v.GWidget)))
}
func (v *Toolbar) GetOrientation() Orientation {
	return Orientation(C.gtk_toolbar_get_orientation(C.to_GtkToolbar(v.GWidget)))
}
func (v *Toolbar) GetStyle() ToolbarStyle {
	return ToolbarStyle(C.gtk_toolbar_get_style(C.to_GtkToolbar(v.GWidget)))
}
func (v *Toolbar) GetIconSize() IconSize {
	return IconSize(C.gtk_toolbar_get_icon_size(C.to_GtkToolbar(v.GWidget)))
}
func (v *Toolbar) GetTooltips() bool {
	return gboolean2bool(C.gtk_toolbar_get_tooltips(C.to_GtkToolbar(v.GWidget)))
}
func (v *Toolbar) GetReliefStyle() ReliefStyle {
	return ReliefStyle(C.gtk_toolbar_get_relief_style(C.to_GtkToolbar(v.GWidget)))
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
	C.gtk_toolbar_set_style(C.to_GtkToolbar(v.GWidget), C.GtkToolbarStyle(style))
}
// gtk_toolbar_insert_stock
func (v *Toolbar) SetIconSize(icon_size IconSize) {
	C.gtk_toolbar_set_icon_size(C.to_GtkToolbar(v.GWidget), C.GtkIconSize(icon_size))
}
// gtk_toolbar_remove_space
func (v *Toolbar) UnsetStyle() {
	C.gtk_toolbar_unset_style(C.to_GtkToolbar(v.GWidget))
}

//-----------------------------------------------------------------------
// GtkToolItem
//-----------------------------------------------------------------------

type IToolItem interface {
	AsToolItem() *ToolItem
}

type ToolItem struct {
	Bin
}

func NewToolItem() *ToolItem {		
	return &ToolItem{Bin{Container{Widget{C.to_GtkWidget(unsafe.Pointer(C.gtk_tool_item_new()))}}}}
}

func (v *ToolItem) AsToolItem() *ToolItem {
	return v
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
	C.gtk_tool_item_set_homogeneous(C.to_GtkToolItem(v.GWidget), bool2gboolean(homogeneous))
}
func (v *ToolItem) GetHomogeneous() bool {
	return gboolean2bool(C.gtk_tool_item_get_homogeneous(C.to_GtkToolItem(v.GWidget)))
}
func (v *ToolItem) SetExpand(expand bool) {	
	C.gtk_tool_item_set_expand(C.to_GtkToolItem(v.GWidget), bool2gboolean(expand))
}
func (v *ToolItem) GetExpand() bool {
	return gboolean2bool(C.gtk_tool_item_get_expand(C.to_GtkToolItem(v.GWidget)))
}
/*func (v *GtkToolItem) SetTooltip(tooltips *GtkTooltips, tip_text, tip_private string) { // FIXME
	p_tip_text := C.CString(tip_text)
	defer C.free_string(p_tip_text)	
	p_tip_private := C.CString(tip_private)
	defer C.free_string(p_tip_private)	
	C.gtk_tool_item_set_tooltip(C.to_GtkToolItem(v.GWidget), tooltips.Tooltip, C.to_gcharptr(p_tip_text), C.to_gcharptr(p_tip_private))
}*/
func (v *ToolItem) SetTooltipMarkup(markup string) {	
	p_markup := C.CString(markup)
	defer C.free_string(p_markup)		
	C.gtk_tool_item_set_tooltip_markup(C.to_GtkToolItem(v.GWidget), C.to_gcharptr(p_markup))
}
func (v *ToolItem) GetToolbarStyle() ToolbarStyle {
	return ToolbarStyle(C.gtk_tool_item_get_toolbar_style(C.to_GtkToolItem(v.GWidget)))
}
func (v *ToolItem) GetReliefStyle() ReliefStyle {
	return ReliefStyle(C.gtk_tool_item_get_relief_style(C.to_GtkToolItem(v.GWidget)))
}
func (v *ToolItem) GetTextAlignment() float64 {
	return float64(C.gtk_tool_item_get_text_alignment(C.to_GtkToolItem(v.GWidget)))
}
func (v *ToolItem) GetTextOrientation() Orientation {
	return Orientation(C.gtk_tool_item_get_text_orientation(C.to_GtkToolItem(v.GWidget)))
}
// gtk_tool_item_retrieve_proxy_menu_item
// gtk_tool_item_get_proxy_menu_item
// gtk_tool_item_set_proxy_menu_item
func (v *ToolItem) RebuildMenu() {		
	C.gtk_tool_item_rebuild_menu(C.to_GtkToolItem(v.GWidget))
}
// gtk_tool_item_toolbar_reconfigured
// gtk_tool_item_get_text_size_group

//-----------------------------------------------------------------------
// GtkToolPalette
//-----------------------------------------------------------------------

// gtk_tool_palette_new
// gtk_tool_palette_get_exclusive
// gtk_tool_palette_set_exclusive
// gtk_tool_palette_get_expand
// gtk_tool_palette_set_expand
// gtk_tool_palette_get_group_position
// gtk_tool_palette_set_group_position
// gtk_tool_palette_get_icon_size
// gtk_tool_palette_set_icon_size
// gtk_tool_palette_unset_icon_size
// gtk_tool_palette_get_style
// gtk_tool_palette_set_style
// gtk_tool_palette_unset_style
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

// gtk_tool_item_group_get_collapsed
// gtk_tool_item_group_get_drop_item
// gtk_tool_item_group_get_ellipsize
// gtk_tool_item_group_get_item_position
// gtk_tool_item_group_get_n_items
// gtk_tool_item_group_get_label
// gtk_tool_item_group_get_label_widget
// gtk_tool_item_group_get_nth_item
// gtk_tool_item_group_get_header_relief
// gtk_tool_item_group_insert
// gtk_tool_item_group_new
// gtk_tool_item_group_set_collapsed
// gtk_tool_item_group_set_ellipsize
// gtk_tool_item_group_set_item_position
// gtk_tool_item_group_set_label
// gtk_tool_item_group_set_label_widget
// gtk_tool_item_group_set_header_relief

//-----------------------------------------------------------------------
// GtkSeparatorToolItem
//-----------------------------------------------------------------------
type SeparatorToolItem struct {
	ToolItem
}

func NewSeparatorToolItem() *SeparatorToolItem {		
	return &SeparatorToolItem{ToolItem{Bin{Container{Widget{
		C.to_GtkWidget(unsafe.Pointer(C.gtk_separator_tool_item_new()))}}}}}
}

func (v *SeparatorToolItem) AsToolItem() *ToolItem {
	return &v.ToolItem
}

func (v *SeparatorToolItem) SetDraw(draw bool) {	
	C.gtk_separator_tool_item_set_draw(C.to_GtkSeparatorToolItem(v.GWidget), bool2gboolean(draw))
}
func (v *SeparatorToolItem) GetDraw() bool {
	return gboolean2bool(C.gtk_separator_tool_item_get_draw(C.to_GtkSeparatorToolItem(v.GWidget)))
}

//-----------------------------------------------------------------------
// GtkToolButton
//-----------------------------------------------------------------------

type ToolButton struct {
	ToolItem	
}

func NewToolButton(icon_widget *Widget, text string) *ToolButton {
	p_text := C.CString(text)
	defer C.free_string(p_text)	
	p_icon_widget := C.to_GtkWidget(unsafe.Pointer(icon_widget))
	p_tool_button_widget := C.to_GtkWidget(unsafe.Pointer(
		C.gtk_tool_button_new(p_icon_widget, C.to_gcharptr(p_text))))
	return &ToolButton{ToolItem{Bin{Container{Widget{p_tool_button_widget}}}}}
}
func NewToolButtonFromStock(stock_id string) *ToolButton {
	p_stock_id := C.CString(stock_id)
	defer C.free_string(p_stock_id)
	p_tool_button_widget := C.to_GtkWidget(unsafe.Pointer(
		C.gtk_tool_button_new_from_stock(C.to_gcharptr(p_stock_id))))
	return &ToolButton{ToolItem{Bin{Container{Widget{p_tool_button_widget}}}}}
}

func (v *ToolButton) AsToolItem() *ToolItem {
	return &v.ToolItem
}

func (v *ToolButton) OnClicked(onclick interface{}, datas ...interface{}) int {
	return v.Connect("clicked", onclick, datas...)
}

func (v *ToolButton) SetLabel(label string) {
	p_label := C.CString(label)
	defer C.free_string(p_label)	
	C.gtk_tool_button_set_label(C.to_GtkToolButton(v.GWidget), C.to_gcharptr(p_label))
}
func (v *ToolButton) GetLabel() string {
	return C.GoString(C.to_charptr(C.gtk_tool_button_get_label(C.to_GtkToolButton(v.GWidget))))
}
func (v *ToolButton) SetUseUnderline(use_underline bool) {	
	C.gtk_tool_button_set_use_underline(C.to_GtkToolButton(v.GWidget), bool2gboolean(use_underline))
}
func (v *ToolButton) GetUseUnderline() bool {
	return gboolean2bool(C.gtk_tool_button_get_use_underline(C.to_GtkToolButton(v.GWidget)))
}
func (v *ToolButton) SetStockId(stock_id string) {
	p_stock_id := C.CString(stock_id)
	defer C.free_string(p_stock_id)	
	C.gtk_tool_button_set_stock_id(C.to_GtkToolButton(v.GWidget), C.to_gcharptr(p_stock_id))
}
func (v *ToolButton) GetStockId() string {
	return C.GoString(C.to_charptr(C.gtk_tool_button_get_stock_id(C.to_GtkToolButton(v.GWidget))))
}
func (v *ToolButton) SetIconName(icon_name string) {
	p_icon_name := C.CString(icon_name)
	defer C.free_string(p_icon_name)	
	C.gtk_tool_button_set_icon_name(C.to_GtkToolButton(v.GWidget), C.to_gcharptr(p_icon_name))
}
func (v *ToolButton) GetIconName() string {
	return C.GoString(C.to_charptr(C.gtk_tool_button_get_icon_name(C.to_GtkToolButton(v.GWidget))))
}
func (v *ToolButton) SetIconWidget(icon_widget *Widget) {
	p_icon_widget := C.to_GtkWidget(unsafe.Pointer(icon_widget.GWidget))
	C.gtk_tool_button_set_icon_widget(C.to_GtkToolButton(v.GWidget), p_icon_widget)
}
var g_ToolButton_IconWidget *Widget
func (v *ToolButton) GetIconWidget() *Widget {
	if g_ToolButton_IconWidget == nil {
		g_ToolButton_IconWidget = &Widget{C.to_GtkWidget(unsafe.Pointer(
			C.gtk_tool_button_get_icon_widget(C.to_GtkToolButton(v.GWidget))))}	
	} else {
		g_ToolButton_IconWidget.GWidget = C.gtk_tool_button_get_icon_widget(C.to_GtkToolButton(v.GWidget))
	}
	return g_ToolButton_IconWidget
}
func (v *ToolButton) SetLabelWidget(label_widget *Widget) {
	p_label_widget := C.to_GtkWidget(unsafe.Pointer(label_widget.GWidget))
	C.gtk_tool_button_set_label_widget(C.to_GtkToolButton(v.GWidget), p_label_widget)
}
var g_ToolButton_LabelWidget *Widget
func (v *ToolButton) GetLabelWidget() *Widget {
	if g_ToolButton_LabelWidget == nil {
		g_ToolButton_LabelWidget = &Widget{C.to_GtkWidget(unsafe.Pointer(
			C.gtk_tool_button_get_label_widget(C.to_GtkToolButton(v.GWidget))))}	
	} else {
		g_ToolButton_LabelWidget.GWidget = C.gtk_tool_button_get_label_widget(C.to_GtkToolButton(v.GWidget))
	}
	return g_ToolButton_LabelWidget
}
//-----------------------------------------------------------------------
// GtkMenuToolButton
//-----------------------------------------------------------------------

// gtk_menu_tool_button_new
// gtk_menu_tool_button_new_from_stock
// gtk_menu_tool_button_set_menu
// gtk_menu_tool_button_get_menu
// gtk_menu_tool_button_set_arrow_tooltip
// gtk_menu_tool_button_set_arrow_tooltip_text
// gtk_menu_tool_button_set_arrow_tooltip_markup

//-----------------------------------------------------------------------
// GtkToggleToolButton
//-----------------------------------------------------------------------

type ToggleToolButton struct {
	ToolButton
}

func NewToggleToolButton() *ToggleToolButton {		
	return &ToggleToolButton{ToolButton{ToolItem{Bin{Container{Widget{
		C.to_GtkWidget(unsafe.Pointer(C.gtk_toggle_tool_button_new()))}}}}}}
}
func NewToggleToolButtonFromStock(stock_id string) *ToggleToolButton {		
	p_stock_id := C.CString(stock_id)
	defer C.free_string(p_stock_id)	
	return &ToggleToolButton{ToolButton{ToolItem{Bin{Container{Widget{
		C.to_GtkWidget(unsafe.Pointer(C.gtk_toggle_tool_button_new_from_stock(C.to_gcharptr(p_stock_id))))}}}}}}
}

func (v *ToggleToolButton) AsToolItem() *ToolItem {
	return &v.ToolButton.ToolItem
}

func (v *ToggleToolButton) OnToggled(onclick interface{}, datas ...interface{}) int {
	return v.Connect("toggled", onclick, datas...)
}

func (v *ToggleToolButton) SetActive(is_active bool) {	
	C.gtk_toggle_tool_button_set_active(C.to_GtkToggleToolButton(v.GWidget), bool2gboolean(is_active))
}
func (v *ToggleToolButton) GetActive() bool {
	return gboolean2bool(C.gtk_toggle_tool_button_get_active(C.to_GtkToggleToolButton(v.GWidget)))
}

//-----------------------------------------------------------------------
// GtkRadioToolButton
//-----------------------------------------------------------------------

type RadioToolButton struct {
	ToggleToolButton
}

func NewRadioToolButton(group *glib.SList) *RadioToolButton {		
	return &RadioToolButton{ToggleToolButton{ToolButton{ToolItem{Bin{Container{Widget{
		C.to_GtkWidget(unsafe.Pointer(C.gtk_radio_tool_button_new(
			C.to_gslist(unsafe.Pointer(group.ToSList())))))}}}}}}}
}
func NewRadioToolButtonFromStock(group *glib.SList, stock_id string) *RadioToolButton {		
	p_stock_id := C.CString(stock_id)
	defer C.free_string(p_stock_id)	
	return &RadioToolButton{ToggleToolButton{ToolButton{ToolItem{Bin{Container{Widget{
		C.to_GtkWidget(unsafe.Pointer(C.gtk_radio_tool_button_new_from_stock(
			C.to_gslist(unsafe.Pointer(group.ToSList())), C.to_gcharptr(p_stock_id))))}}}}}}}
}
// gtk_radio_tool_button_new_from_widget
// gtk_radio_tool_button_new_with_stock_from_widget
// gtk_radio_tool_button_get_group
// gtk_radio_tool_button_set_group

//-----------------------------------------------------------------------
// GtkUIManager
//-----------------------------------------------------------------------

// gtk_ui_manager_new
// gtk_ui_manager_set_add_tearoffs
// gtk_ui_manager_get_add_tearoffs
// gtk_ui_manager_insert_action_group
// gtk_ui_manager_remove_action_group
// gtk_ui_manager_get_action_groups
// gtk_ui_manager_get_accel_group
// gtk_ui_manager_get_widget
// gtk_ui_manager_get_toplevels
// gtk_ui_manager_get_action
// gtk_ui_manager_add_ui_from_string
// gtk_ui_manager_add_ui_from_file
// gtk_ui_manager_new_merge_id
// gtk_ui_manager_add_ui
// gtk_ui_manager_remove_ui
// gtk_ui_manager_get_ui
// gtk_ui_manager_ensure_update

//-----------------------------------------------------------------------
// GtkActionGroup
//-----------------------------------------------------------------------

// gtk_action_group_new
// gtk_action_group_get_name
// gtk_action_group_get_sensitive
// gtk_action_group_set_sensitive
// gtk_action_group_get_visible
// gtk_action_group_set_visible
// gtk_action_group_get_action
// gtk_action_group_list_actions
// gtk_action_group_add_action
// gtk_action_group_add_action_with_accel
// gtk_action_group_remove_action
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

// gtk_action_new
// gtk_action_get_name
// gtk_action_is_sensitive
// gtk_action_get_sensitive
// gtk_action_set_sensitive
// gtk_action_is_visible
// gtk_action_get_visible
// gtk_action_set_visible
// gtk_action_activate
// gtk_action_create_icon
// gtk_action_create_menu_item
// gtk_action_create_tool_item
// gtk_action_create_menu
// gtk_action_connect_proxy
// gtk_action_disconnect_proxy
// gtk_action_get_proxies
// gtk_action_connect_accelerator
// gtk_action_disconnect_accelerator
// gtk_action_block_activate
// gtk_action_unblock_activate
// gtk_action_block_activate_from
// gtk_action_unblock_activate_from
// gtk_action_get_always_show_image
// gtk_action_set_always_show_image
// gtk_action_get_accel_path
// gtk_action_set_accel_path
// gtk_action_get_accel_closure
// gtk_action_set_accel_group
// gtk_action_set_label
// gtk_action_get_label
// gtk_action_set_short_label
// gtk_action_get_short_label
// gtk_action_set_tooltip
// gtk_action_get_tooltip
// gtk_action_set_stock_id
// gtk_action_get_stock_id
// gtk_action_set_gicon
// gtk_action_get_gicon
// gtk_action_set_icon_name
// gtk_action_get_icon_name
// gtk_action_set_visible_horizontal
// gtk_action_get_visible_horizontal
// gtk_action_set_visible_vertical
// gtk_action_get_visible_vertical
// gtk_action_set_is_important
// gtk_action_get_is_important

//-----------------------------------------------------------------------
// GtkToggleAction
//-----------------------------------------------------------------------

// gtk_toggle_action_new
// gtk_toggle_action_toggled
// gtk_toggle_action_set_active
// gtk_toggle_action_get_active
// gtk_toggle_action_set_draw_as_radio
// gtk_toggle_action_get_draw_as_radio

//-----------------------------------------------------------------------
// GtkRadioAction
//-----------------------------------------------------------------------

// gtk_radio_action_new
// gtk_radio_action_get_group
// gtk_radio_action_set_group
// gtk_radio_action_get_current_value
// gtk_radio_action_set_current_value

//-----------------------------------------------------------------------
// GtkRecentAction
//-----------------------------------------------------------------------

// gtk_recent_action_new
// gtk_recent_action_new_for_manager
// gtk_recent_action_get_show_numbers
// gtk_recent_action_set_show_numbers

//-----------------------------------------------------------------------
// GtkActivatable
//-----------------------------------------------------------------------

// gtk_activatable_do_set_related_action
// gtk_activatable_get_related_action
// gtk_activatable_get_use_action_appearance
// gtk_activatable_sync_action_properties
// gtk_activatable_set_related_action
// gtk_activatable_set_use_action_appearance

//-----------------------------------------------------------------------
// GtkColorButton
//-----------------------------------------------------------------------

// gtk_color_button_new
// gtk_color_button_new_with_color
// gtk_color_button_set_color
// gtk_color_button_get_color
// gtk_color_button_set_alpha
// gtk_color_button_get_alpha
// gtk_color_button_set_use_alpha
// gtk_color_button_get_use_alpha
// gtk_color_button_set_title
// gtk_color_button_get_title

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
	FILE_CHOOSER_ACTION_OPEN          FileChooserAction = 0
	FILE_CHOOSER_ACTION_SAVE          FileChooserAction = 1
	FILE_CHOOSER_ACTION_SELECT_FOLDER FileChooserAction = 2
	FILE_CHOOSER_ACTION_CREATE_FOLDER FileChooserAction = 3
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
	C.gtk_file_chooser_set_local_only(v.GFileChooser, bool2gboolean(b))
}
func (v *FileChooser) GetLocalOnly() bool {
	return gboolean2bool(C.gtk_file_chooser_get_local_only(v.GFileChooser))
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
	return C.GoString(C.to_charptr(C.gtk_file_chooser_get_filename(v.GFileChooser)))
}
func (v *FileChooser) SetFilename(filename string) {
	ptr := C.CString(filename)
	defer C.free_string(ptr)
	C.gtk_file_chooser_set_filename(v.GFileChooser, ptr)
}

// gboolean gtk_file_chooser_select_filename(GtkFileChooser* chooser, const char* filename);
// void gtk_file_chooser_unselect_filename(GtkFileChooser* chooser, const char* filename);
// void gtk_file_chooser_select_all(GtkFileChooser* chooser);
// void gtk_file_chooser_unselect_all(GtkFileChooser* chooser);
// GSList*  gtk_file_chooser_get_filenames(GtkFileChooser* chooser);

func (v *FileChooser) SetCurrentFolder(f string) bool {
	cf := C.CString(f)
	defer C.free_string(cf)
	return gboolean2bool(C.gtk_file_chooser_set_current_folder(
		v.GFileChooser, C.to_gcharptr(cf)))
}
func (v *FileChooser) GetCurrentFolder() string {
	return C.GoString(C.to_charptr(C.gtk_file_chooser_get_current_folder(v.GFileChooser)))
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
		ret[i] = &FileFilter{C.to_GtkFileFilter(C.g_slist_nth_data(c_list, C.guint(i)))}
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
	defer C.free_string(ptitle)
	widget := Widget{
		C.gtk_file_chooser_button_new(C.to_gcharptr(ptitle), C.GtkFileChooserAction(action)), // FIXME
	}
	return &FileChooserButton{HBox{Box{Container{widget}}},
		FileChooser{C.to_GtkFileChooser(widget.GWidget)}, // FIXME
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
		FileChooser{C.to_GtkFileChooser(w.GWidget)},
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

func NewFileChooserDialog(title string, parent *Window, file_chooser_action FileChooserAction, button_text string, button_action ResponseType, buttons ...interface{}) *FileChooserDialog {
	ptitle := C.CString(title)
	defer C.free_string(ptitle)
	pbutton := C.CString(button_text)
	defer C.free_string(pbutton)
	widget := Widget{
		C._gtk_file_chooser_dialog_new(
			C.to_gcharptr(ptitle),
			parent.ToNative(),
			C.int(file_chooser_action),
			C.int(button_action),
			C.to_gcharptr(pbutton))}
	ret := &FileChooserDialog{
		Dialog{Window{Bin{Container{widget}}}},
		FileChooser{C.to_GtkFileChooser(widget.GWidget)}}
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
	w := Widget{C._gtk_file_chooser_widget_new(C.int(file_chooser_action))}
	return &FileChooserWidget{VBox{Box{Container{w}}},
		FileChooser{C.to_GtkFileChooser(w.GWidget)}, // FIXME
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
	defer C.free_string(ptr)
	C.gtk_file_filter_set_name(v.GFileFilter, C.to_gcharptr(ptr))
}
func (v *FileFilter) GetName() string {
	return C.GoString(C.to_charptr(C.gtk_file_filter_get_name(v.GFileFilter)))
}
func (v *FileFilter) AddMimeType(mimetype string) {
	ptr := C.CString(mimetype)
	defer C.free_string(ptr)
	C.gtk_file_filter_add_mime_type(v.GFileFilter, C.to_gcharptr(ptr))
}
func (v *FileFilter) AddPattern(pattern string) {
	ptr := C.CString(pattern)
	defer C.free_string(ptr)
	C.gtk_file_filter_add_pattern(v.GFileFilter, C.to_gcharptr(ptr))
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
	return &FontButton{Button{Bin{Container{Widget{
		C.gtk_font_button_new()}}}}}
}
func NewFontButtonWithFont(fontname string) *FontButton {
	ptr := C.CString(fontname)
	defer C.free_string(ptr)
	return &FontButton{Button{Bin{Container{Widget{
		C.gtk_font_button_new_with_font(C.to_gcharptr(ptr))}}}}}
}
func (v *FontButton) SetFontName(fontname string) {
	ptr := C.CString(fontname)
	defer C.free_string(ptr)
	C.gtk_font_button_set_font_name(C.to_GtkFontButton(v.GWidget), C.to_gcharptr(ptr))
}
func (v *FontButton) GetFontName() string {
	return C.GoString(C.to_charptr(C.gtk_font_button_get_font_name(C.to_GtkFontButton(v.GWidget))))
}

// gtk_font_button_set_show_style
// gtk_font_button_get_show_style

func (v *FontButton) SetShowSize(show_size bool) {
	C.gtk_font_button_set_show_size(C.to_GtkFontButton(v.GWidget), bool2gboolean(show_size))
}
func (v *FontButton) GetShowSize() bool {
	return gboolean2bool(C.gtk_font_button_get_show_size(C.to_GtkFontButton(v.GWidget)))
}

// gtk_font_button_set_use_font
// gtk_font_button_get_use_font

func (v *FontButton) SetUseSize(use_size bool) {
	C.gtk_font_button_set_use_size(C.to_GtkFontButton(v.GWidget), bool2gboolean(use_size))
}
func (v *FontButton) GetUseSize() bool {
	return gboolean2bool(C.gtk_font_button_get_use_size(C.to_GtkFontButton(v.GWidget)))
}
func (v *FontButton) SetTitle(title string) {
	ptr := C.CString(title)
	defer C.free_string(ptr)
	C.gtk_font_button_set_title(C.to_GtkFontButton(v.GWidget), C.to_gcharptr(ptr))
}
func (v *FontButton) GetTitle() string {
	return C.GoString(C.to_charptr(C.gtk_font_button_get_title(C.to_GtkFontButton(v.GWidget))))
}

//-----------------------------------------------------------------------
// GtkFontSelection
//-----------------------------------------------------------------------

// gtk_font_selection_new
// gtk_font_selection_get_font
// gtk_font_selection_get_font_name
// gtk_font_selection_set_font_name
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
	defer C.free_string(ptitle)
	return &FontSelectionDialog{Dialog{Window{Bin{Container{Widget{
		C.gtk_font_selection_dialog_new(C.to_gcharptr(ptitle))}}}}}}
}
func (v *FontSelectionDialog) GetFontName() string {
	return C.GoString(C.to_charptr(C.gtk_font_selection_dialog_get_font_name(C.to_GtkFontSelectionDialog(v.GWidget))))
}
func (v *FontSelectionDialog) SetFontName(font string) {
	pfont := C.CString(font)
	defer C.free_string(pfont)
	C.gtk_font_selection_dialog_set_font_name(C.to_GtkFontSelectionDialog(v.GWidget), C.to_gcharptr(pfont))
}

// gtk_font_selection_dialog_get_preview_text
// gtk_font_selection_dialog_set_preview_text
// gtk_font_selection_dialog_get_cancel_button
// gtk_font_selection_dialog_get_ok_button
// gtk_font_selection_dialog_get_font_selection //since 2.22

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
	C.gtk_alignment_set(C.to_GtkAlignment(v.GWidget), C.gfloat(xalign), C.gfloat(yalign), C.gfloat(xscale), C.gfloat(yscale))
}
func (v *Alignment) SetPadding(padding_top uint, padding_bottom uint, padding_left uint, padding_right uint) {
	C.gtk_alignment_set_padding(C.to_GtkAlignment(v.GWidget), C.guint(padding_top), C.guint(padding_bottom), C.guint(padding_left), C.guint(padding_right))
}
func (v *Alignment) GetPadding() (padding_top uint, padding_bottom uint, padding_left uint, padding_right uint) {
	var cpadding_top, cpadding_bottom, cpadding_left, cpadding_right C.guint
	C.gtk_alignment_get_padding(C.to_GtkAlignment(v.GWidget), &cpadding_top, &cpadding_bottom, &cpadding_left, &cpadding_right)
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

func NewHBox(homogeneous bool, spacing uint) *HBox {
	return &HBox{Box{Container{Widget{
		C.gtk_hbox_new(bool2gboolean(homogeneous), C.gint(spacing))}}}}
}

//-----------------------------------------------------------------------
// GtkVBox
//-----------------------------------------------------------------------
type VBox struct {
	Box
}

func NewVBox(homogeneous bool, spacing uint) *VBox {
	return &VBox{Box{Container{Widget{
		C.gtk_vbox_new(bool2gboolean(homogeneous), C.gint(spacing))}}}}
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
func (v *Fixed) Put(w WidgetLike, x, y int) {
	C.gtk_fixed_put(C.to_GtkFixed(v.GWidget), w.ToNative(), C.gint(x), C.gint(y))
}
func (v *Fixed) Move(w WidgetLike, x, y int) {
	C.gtk_fixed_move(C.to_GtkFixed(v.GWidget), w.ToNative(), C.gint(x), C.gint(y))
}

//Deprecated since 2.20. Use GtkWidget.GetHasWindow() instead.
/*GtkFixed gets GetHasWindow() from anonymous field so this method can be commented out.
func (v *GtkFixed) GetHasWindow() bool {
	deprecated_since(2,20,0,"gtk_fixed_get_has_window()")
	return gboolean2bool(C.gtk_fixed_get_has_window(C.to_GtkFixed(v.GWidget)))
}*/
//Deprecated since 2.20. Use GtkWidget.SetHasWindow() instead.
/*GtkFixed gets SetHasWindow() from anonymous field so this method can be commented out.
func (v *GtkFixed) SetHasWindow(has_window bool) {
	deprecated_since(2,20,0,"gtk_fixed_set_has_window()")
	C.gtk_fixed_set_has_window(C.to_GtkFixed(v.GWidget), bool2gboolean(has_window))
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

// gtk_layout_new
// gtk_layout_put
// gtk_layout_move
// gtk_layout_set_size
// gtk_layout_get_size
// gtk_layout_freeze
// gtk_layout_thaw
// gtk_layout_get_hadjustment
// gtk_layout_get_vadjustment
// gtk_layout_set_hadjustment
// gtk_layout_set_vadjustment
// gtk_layout_get_bin_window

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
func (v *Notebook) AppendPage(child WidgetLike, tab_label WidgetLike) int {
	return int(C.gtk_notebook_append_page(C.to_GtkNotebook(v.GWidget), child.ToNative(), tab_label.ToNative()))
}
func (v *Notebook) AppendPageMenu(child WidgetLike, tab_label WidgetLike, menu_label WidgetLike) int {
	return int(C.gtk_notebook_append_page_menu(C.to_GtkNotebook(v.GWidget), child.ToNative(), tab_label.ToNative(), menu_label.ToNative()))
}
func (v *Notebook) PrependPage(child WidgetLike, tab_label WidgetLike) int {
	return int(C.gtk_notebook_prepend_page(C.to_GtkNotebook(v.GWidget), child.ToNative(), tab_label.ToNative()))
}
func (v *Notebook) PrependPageMenu(child WidgetLike, tab_label WidgetLike, menu_label WidgetLike) int {
	return int(C.gtk_notebook_prepend_page_menu(C.to_GtkNotebook(v.GWidget), child.ToNative(), tab_label.ToNative(), menu_label.ToNative()))
}
func (v *Notebook) InsertPage(child WidgetLike, tab_label WidgetLike, position int) int {
	return int(C.gtk_notebook_insert_page(C.to_GtkNotebook(v.GWidget), child.ToNative(), tab_label.ToNative(), C.gint(position)))
}
func (v *Notebook) InsertPageMenu(child WidgetLike, tab_label WidgetLike, menu_label WidgetLike, position int) int {
	return int(C.gtk_notebook_insert_page_menu(C.to_GtkNotebook(v.GWidget), child.ToNative(), tab_label.ToNative(), menu_label.ToNative(), C.gint(position)))
}
func (v *Notebook) RemovePage(child WidgetLike, page_num int) {
	C.gtk_notebook_remove_page(C.to_GtkNotebook(v.GWidget), C.gint(page_num))
}
func (v *Notebook) PageNum(child WidgetLike) int {
	return int(C.gtk_notebook_page_num(C.to_GtkNotebook(v.GWidget), child.ToNative()))
}
func (v *Notebook) NextPage() {
	C.gtk_notebook_next_page(C.to_GtkNotebook(v.GWidget))
}
func (v *Notebook) PrevPage() {
	C.gtk_notebook_prev_page(C.to_GtkNotebook(v.GWidget))
}
func (v *Notebook) ReorderChild(child WidgetLike, position int) {
	C.gtk_notebook_reorder_child(C.to_GtkNotebook(v.GWidget), child.ToNative(), C.gint(position))
}
func (v *Notebook) SetTabPos(pos PositionType) {
	C.gtk_notebook_set_tab_pos(C.to_GtkNotebook(v.GWidget), C.GtkPositionType(pos))
}
func (v *Notebook) SetShowTabs(show_tabs bool) {
	C.gtk_notebook_set_show_tabs(C.to_GtkNotebook(v.GWidget), bool2gboolean(show_tabs))
}
func (v *Notebook) SetShowBorder(show_border bool) {
	C.gtk_notebook_set_show_border(C.to_GtkNotebook(v.GWidget), bool2gboolean(show_border))
}
func (v *Notebook) SetScrollable(scrollable bool) {
	C.gtk_notebook_set_scrollable(C.to_GtkNotebook(v.GWidget), bool2gboolean(scrollable))
}

//Deprecated.
func (v *Notebook) SetTabBorder(border_width uint) {
	deprecated_since(2, 0, 0, "gtk_notebook_set_tab_border()")
	C.gtk_notebook_set_tab_border(C.to_GtkNotebook(v.GWidget), C.guint(border_width))
}
func (v *Notebook) PopupEnable() {
	C.gtk_notebook_popup_enable(C.to_GtkNotebook(v.GWidget))
}
func (v *Notebook) PopupDisable() {
	C.gtk_notebook_popup_disable(C.to_GtkNotebook(v.GWidget))
}
func (v *Notebook) GetCurrentPage() int {
	return int(C.gtk_notebook_get_current_page(C.to_GtkNotebook(v.GWidget)))
}
func (v *Notebook) GetMenuLabel(child WidgetLike) *Widget {
	return &Widget{
		C.gtk_notebook_get_menu_label(C.to_GtkNotebook(v.GWidget), child.ToNative())}
}
func (v *Notebook) GetNthPage(page_num int) *Widget {
	return &Widget{
		C.gtk_notebook_get_nth_page(C.to_GtkNotebook(v.GWidget), C.gint(page_num))}
}
func (v *Notebook) GetNPages() int {
	return int(C.gtk_notebook_get_n_pages(C.to_GtkNotebook(v.GWidget)))
}
func (v *Notebook) GetTabLabel(child WidgetLike) *Widget {
	return &Widget{
		C.gtk_notebook_get_tab_label(C.to_GtkNotebook(v.GWidget), child.ToNative())}
}

//Deprecated since 2.20. Modify the "tab-expand" and "tab-fill" child properties instead.
func (v *Notebook) QueryTabLabelPacking(child WidgetLike, expand *bool, fill *bool, pack_type *uint) {
	deprecated_since(2, 20, 0, "gtk_notebook_query_tab_label_packing()")
	var cexpand, cfill C.gboolean
	var cpack_type C.GtkPackType
	C.gtk_notebook_query_tab_label_packing(C.to_GtkNotebook(v.GWidget), child.ToNative(), &cexpand, &cfill, &cpack_type)
	*expand = gboolean2bool(cexpand)
	*fill = gboolean2bool(cfill)
	*pack_type = uint(cpack_type)
}

//Deprecated.
func (v *Notebook) SetHomogeneousTabs(homogeneous bool) {
	deprecated_since(2, 0, 0, "gtk_notebook_set_homogeneous_tabs()")
	C.gtk_notebook_set_homogeneous_tabs(C.to_GtkNotebook(v.GWidget), bool2gboolean(homogeneous))
}
func (v *Notebook) SetMenuLabel(child WidgetLike, menu_label WidgetLike) {
	C.gtk_notebook_set_menu_label(C.to_GtkNotebook(v.GWidget), child.ToNative(), menu_label.ToNative())
}
func (v *Notebook) SetMenuLabelText(child WidgetLike, name string) {
	ptr := C.CString(name)
	defer C.free_string(ptr)
	C.gtk_notebook_set_menu_label_text(C.to_GtkNotebook(v.GWidget), child.ToNative(), C.to_gcharptr(ptr))
}

//Deprecated.
func (v *Notebook) SetTabHBorder(tab_hborder uint) {
	deprecated_since(2, 0, 0, "gtk_notebook_set_tab_hborder()")
	C.gtk_notebook_set_tab_hborder(C.to_GtkNotebook(v.GWidget), C.guint(tab_hborder))
}
func (v *Notebook) SetTabLabel(child WidgetLike, tab_label WidgetLike) {
	C.gtk_notebook_set_tab_label(C.to_GtkNotebook(v.GWidget), child.ToNative(), tab_label.ToNative())
}

//Deprecated since 2.20. Modify the "tab-expand" and "tab-fill" child properties instead.
func (v *Notebook) SetTabLabelPacking(child WidgetLike, expand bool, fill bool, pack_type uint) {
	deprecated_since(2, 20, 0, "gtk_notebook_set_tab_label_packing()")
	C.gtk_notebook_set_tab_label_packing(C.to_GtkNotebook(v.GWidget), child.ToNative(), bool2gboolean(expand), bool2gboolean(fill), C.GtkPackType(pack_type))
}
func (v *Notebook) SetTabLabelText(child WidgetLike, name string) {
	ptr := C.CString(name)
	defer C.free_string(ptr)
	C.gtk_notebook_set_tab_label_text(C.to_GtkNotebook(v.GWidget), child.ToNative(), C.to_gcharptr(ptr))
}

//Deprecated.
func (v *Notebook) SetTabVBorder(tab_vborder uint) {
	deprecated_since(2, 0, 0, "gtk_notebook_set_tab_vborder()")
	C.gtk_notebook_set_tab_vborder(C.to_GtkNotebook(v.GWidget), C.guint(tab_vborder))
}
func (v *Notebook) SetReorderable(child WidgetLike, reorderable bool) {
	C.gtk_notebook_set_tab_reorderable(C.to_GtkNotebook(v.GWidget), child.ToNative(), bool2gboolean(reorderable))
}
func (v *Notebook) SetTabDetachable(child WidgetLike, detachable bool) {
	C.gtk_notebook_set_tab_detachable(C.to_GtkNotebook(v.GWidget), child.ToNative(), bool2gboolean(detachable))
}
func (v *Notebook) GetMenuLabelText(child WidgetLike) string {
	return C.GoString(C.to_charptr(C.gtk_notebook_get_menu_label_text(C.to_GtkNotebook(v.GWidget), child.ToNative())))
}
func (v *Notebook) GetScrollable() bool {
	return gboolean2bool(C.gtk_notebook_get_scrollable(C.to_GtkNotebook(v.GWidget)))
}
func (v *Notebook) GetShowBorder() bool {
	return gboolean2bool(C.gtk_notebook_get_show_border(C.to_GtkNotebook(v.GWidget)))
}
func (v *Notebook) GetShowTabs() bool {
	return gboolean2bool(C.gtk_notebook_get_show_tabs(C.to_GtkNotebook(v.GWidget)))
}
func (v *Notebook) GetTabLabelText(child WidgetLike) string {
	return C.GoString(C.to_charptr(C.gtk_notebook_get_tab_label_text(C.to_GtkNotebook(v.GWidget), child.ToNative())))
}
func (v *Notebook) GetTabPos() uint {
	return uint(C.gtk_notebook_get_tab_pos(C.to_GtkNotebook(v.GWidget)))
}
func (v *Notebook) GetTabReorderable(child WidgetLike) bool {
	return gboolean2bool(C.gtk_notebook_get_tab_reorderable(C.to_GtkNotebook(v.GWidget), child.ToNative()))
}
func (v *Notebook) GetTabDetachable(child WidgetLike) bool {
	return gboolean2bool(C.gtk_notebook_get_tab_detachable(C.to_GtkNotebook(v.GWidget), child.ToNative()))
}

// gtk_notebook_get_tab_hborder //since 2.22
// gtk_notebook_get_tab_vborder //since 2.22

func (v *Notebook) SetCurrentPage(pageNum int) {
	C.gtk_notebook_set_current_page(C.to_GtkNotebook(v.GWidget), C.gint(pageNum))
}

//Deprecated since 2.12, use SetGroupName() instead
func (v *Notebook) SetGroupId(group_id int) {
	deprecated_since(2, 12, 0, "gtk_notebook_set_group_id()")
	C.gtk_notebook_set_group_id(C.to_GtkNotebook(v.GWidget), C.gint(group_id))
}

//Deprecated since 2.12, use GetGroupName() instead
func (v *Notebook) GetGroupId() int {
	deprecated_since(2, 12, 0, "gtk_notebook_get_group_id()")
	return int(C.gtk_notebook_get_group_id(C.to_GtkNotebook(v.GWidget)))
}

//Deprecated since 2.24, use SetGroupName() instead
func (v *Notebook) SetGroup(group unsafe.Pointer) {
	deprecated_since(2, 24, 0, "gtk_notebook_set_group()")
	C.gtk_notebook_set_group(C.to_GtkNotebook(v.GWidget), C.gpointer(group))
}

//Deprecated since 2.24, use GetGroupName() instead
func (v *Notebook) GetGroup() unsafe.Pointer {
	deprecated_since(2, 24, 0, "gtk_notebook_get_group()")
	return unsafe.Pointer(C.gtk_notebook_get_group(C.to_GtkNotebook(v.GWidget)))
}
func (v *Notebook) SetGroupName(group string) {
	panic_if_version_older(2, 24, 0, "gtk_notebook_set_group_name()")
	ptr := C.CString(group)
	defer C.free_string(ptr)
	C._gtk_notebook_set_group_name(C.to_GtkNotebook(v.GWidget), C.to_gcharptr(ptr))
}
func (v *Notebook) GetGroupName() string {
	panic_if_version_older(2, 24, 0, "gtk_notebook_get_group_name()")
	return C.GoString(C.to_charptr(C._gtk_notebook_get_group_name(C.to_GtkNotebook(v.GWidget))))
}

// gtk_notebook_set_action_widget //since 2.20
// gtk_notebook_get_action_widget //since 2.20
// void gtk_notebook_set_window_creation_hook (GtkNotebookWindowCreationFunc func, gpointer data, GDestroyNotify destroy); //deprecated in 2.24

//-----------------------------------------------------------------------
// GtkTable
//-----------------------------------------------------------------------
type AttachOptions int

const (
	EXPAND AttachOptions = 1 << 0
	SHRINK AttachOptions = 1 << 1
	FILL   AttachOptions = 1 << 2
)

type Table struct {
	Container
}

func NewTable(rows uint, columns uint, homogeneous bool) *Table {
	return &Table{Container{Widget{
		C.gtk_table_new(C.guint(rows), C.guint(columns), bool2gboolean(homogeneous))}}}
}
func (v *Table) Resize(rows uint, columns uint) {
	C.gtk_table_resize(C.to_GtkTable(v.GWidget), C.guint(rows), C.guint(columns))
}
func (v *Table) Attach(child WidgetLike, left_attach uint, right_attach uint, top_attach uint, bottom_attach uint, xoptions AttachOptions, yoptions AttachOptions, xpadding uint, ypadding uint) {
	C.gtk_table_attach(C.to_GtkTable(v.GWidget), child.ToNative(), C.guint(left_attach), C.guint(right_attach), C.guint(top_attach), C.guint(bottom_attach), C.GtkAttachOptions(xoptions), C.GtkAttachOptions(yoptions), C.guint(xpadding), C.guint(ypadding))
}
func (v *Table) AttachDefaults(child WidgetLike, left_attach uint, right_attach uint, top_attach uint, bottom_attach uint) {
	C.gtk_table_attach_defaults(C.to_GtkTable(v.GWidget), child.ToNative(), C.guint(left_attach), C.guint(right_attach), C.guint(top_attach), C.guint(bottom_attach))
}
func (v *Table) SetRowSpacing(child WidgetLike, row uint, spacing uint) {
	C.gtk_table_set_row_spacing(C.to_GtkTable(v.GWidget), C.guint(row), C.guint(spacing))
}
func (v *Table) SetColSpacing(child WidgetLike, column uint, spacing uint) {
	C.gtk_table_set_col_spacing(C.to_GtkTable(v.GWidget), C.guint(column), C.guint(spacing))
}
func (v *Table) SetRowSpacings(child WidgetLike, spacing uint) {
	C.gtk_table_set_row_spacings(C.to_GtkTable(v.GWidget), C.guint(spacing))
}
func (v *Table) SetColSpacings(child WidgetLike, spacing uint) {
	C.gtk_table_set_col_spacings(C.to_GtkTable(v.GWidget), C.guint(spacing))
}
func (v *Table) SetHomogeneous(child WidgetLike, homogeneous bool) {
	C.gtk_table_set_homogeneous(C.to_GtkTable(v.GWidget), bool2gboolean(homogeneous))
}
func (v *Table) GetDefaultRowSpacing(child WidgetLike) uint {
	return uint(C.gtk_table_get_default_row_spacing(C.to_GtkTable(v.GWidget)))
}
func (v *Table) GetHomogeneous(child WidgetLike) bool {
	return gboolean2bool(C.gtk_table_get_homogeneous(C.to_GtkTable(v.GWidget)))
}
func (v *Table) GetRowSpacing(child WidgetLike, row uint) uint {
	return uint(C.gtk_table_get_row_spacing(C.to_GtkTable(v.GWidget), C.guint(row)))
}
func (v *Table) GetColSpacing(child WidgetLike, column uint) uint {
	return uint(C.gtk_table_get_col_spacing(C.to_GtkTable(v.GWidget), C.guint(column)))
}
func (v *Table) GetDefaultColSpacing(child WidgetLike) uint {
	return uint(C.gtk_table_get_default_col_spacing(C.to_GtkTable(v.GWidget)))
}

// gtk_table_get_size //since 2.22

//-----------------------------------------------------------------------
// GtkExpander
//-----------------------------------------------------------------------
type Expander struct {
	Bin
}

func NewExpander(label string) *Expander {
	ptr := C.CString(label)
	defer C.free_string(ptr)
	return &Expander{Bin{Container{Widget{
		C.gtk_expander_new(C.to_gcharptr(ptr))}}}}
}
func NewExpanderWithMnemonic(label string) *Expander {
	ptr := C.CString(label)
	defer C.free_string(ptr)
	return &Expander{Bin{Container{Widget{
		C.gtk_expander_new_with_mnemonic(C.to_gcharptr(ptr))}}}}
}

func (v *Expander) SetExpanded(expanded bool) {
	C.gtk_expander_set_expanded(C.to_GtkExpander(v.GWidget), bool2gboolean(expanded))
}
func (v *Expander) GetExpanded() bool {
	return gboolean2bool(C.gtk_expander_get_expanded(C.to_GtkExpander(v.GWidget)))
}
func (v *Expander) SetSpacing(spacing int) {
	C.gtk_expander_set_spacing(C.to_GtkExpander(v.GWidget), C.gint(spacing))
}
func (v *Expander) GetSpacing() int {
	return int(C.gtk_expander_get_spacing(C.to_GtkExpander(v.GWidget)))
}
func (v *Expander) SetLabel(label string) {
	ptr := C.CString(label)
	defer C.free_string(ptr)
	C.gtk_expander_set_label(C.to_GtkExpander(v.GWidget), C.to_gcharptr(ptr))
}
func (v *Expander) GetLabel() string {
	return C.GoString(C.to_charptr(C.gtk_expander_get_label(C.to_GtkExpander(v.GWidget))))
}
func (v *Expander) SetUseUnderline(setting bool) {
	C.gtk_expander_set_use_underline(C.to_GtkExpander(v.GWidget), bool2gboolean(setting))
}
func (v *Expander) GetUseUnderline() bool {
	return gboolean2bool(C.gtk_expander_get_use_underline(C.to_GtkExpander(v.GWidget)))
}
func (v *Expander) SetUseMarkup(setting bool) {
	C.gtk_expander_set_use_markup(C.to_GtkExpander(v.GWidget), bool2gboolean(setting))
}
func (v *Expander) GetUseMarkup() bool {
	return gboolean2bool(C.gtk_expander_get_use_markup(C.to_GtkExpander(v.GWidget)))
}
func (v *Expander) SetLabelWidget(label_widget LabelLike) {
	C.gtk_expander_set_label_widget(C.to_GtkExpander(v.GWidget), label_widget.ToNative())
}
func (v *Expander) GetLabelWidget() LabelLike {
	return &Label{Widget{
		C.gtk_expander_get_label_widget(C.to_GtkExpander(v.GWidget))}}
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
	SHADOW_NONE       ShadowType = 0
	SHADOW_IN         ShadowType = 1
	SHADOW_OUT        ShadowType = 2
	SHADOW_ETCHED_IN  ShadowType = 3
	SHADOW_ETCHED_OUT ShadowType = 4
)

type Frame struct {
	Bin
}

func NewFrame(label string) *Frame {
	ptr := C.CString(label)
	defer C.free_string(ptr)
	return &Frame{Bin{Container{Widget{
		C.gtk_frame_new(C.to_gcharptr(ptr))}}}}
}
func (v *Frame) SetLabel(label string) {
	ptr := C.CString(label)
	defer C.free_string(ptr)
	C.gtk_frame_set_label(C.to_GtkFrame(v.GWidget), C.to_gcharptr(ptr))
}
func (v *Frame) SetLabelWidget(label_widget LabelLike) {
	C.gtk_frame_set_label_widget(C.to_GtkFrame(v.GWidget), label_widget.ToNative())
}
func (v *Frame) SetLabelAlign(xalign, yalign float64) {
	C.gtk_frame_set_label_align(C.to_GtkFrame(v.GWidget), C.gfloat(xalign), C.gfloat(yalign))
}
func (v *Frame) SetShadowType(shadow_type ShadowType) {
	C.gtk_frame_set_shadow_type(C.to_GtkFrame(v.GWidget), C.GtkShadowType(shadow_type))
}
func (v *Frame) GetLabel() string {
	return C.GoString(C.to_charptr(C.gtk_frame_get_label(C.to_GtkFrame(v.GWidget))))
}
func (v *Frame) GetLabelAlign() (xalign, yalign float64) {
	var xalign_, yalign_ C.gfloat
	C.gtk_frame_get_label_align(C.to_GtkFrame(v.GWidget), &xalign_, &yalign_)
	return float64(xalign_), float64(yalign_)
}
func (v *Frame) GetLabelWidget() LabelLike {
	return &Label{Widget{
		C.gtk_frame_get_label_widget(C.to_GtkFrame(v.GWidget))}}
}
func (v *Frame) GetShadowType() ShadowType {
	return ShadowType(C.gtk_frame_get_shadow_type(C.to_GtkFrame(v.GWidget)))
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
// GtkHScrollbar
//-----------------------------------------------------------------------

// gtk_hscrollbar_new

//-----------------------------------------------------------------------
// GtkVScrollbar
//-----------------------------------------------------------------------

// gtk_vscrollbar_new

//-----------------------------------------------------------------------
// GtkScrolledWindow
//-----------------------------------------------------------------------
type PolicyType int

const (
	POLICY_ALWAYS    = 0
	POLICY_AUTOMATIC = 1
	POLICY_NEVER     = 2
)

type CornerType int

const (
	CORNER_TOP_LEFT     CornerType = 0
	CORNER_BOTTOM_LEFT  CornerType = 1
	CORNER_TOP_RIGHT    CornerType = 2
	CORNER_BOTTOM_RIGHT CornerType = 3
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
		C.gtk_scrolled_window_get_hadjustment(C.to_GtkScrolledWindow(v.GWidget))}
}
func (v *ScrolledWindow) GetVAdjustment() *Adjustment {
	return &Adjustment{
		C.gtk_scrolled_window_get_vadjustment(C.to_GtkScrolledWindow(v.GWidget))}
}

// gtk_scrolled_window_get_hscrollbar
// gtk_scrolled_window_get_vscrollbar

func (v *ScrolledWindow) SetPolicy(hscrollbar_policy PolicyType, vscrollbar_policy PolicyType) {
	C.gtk_scrolled_window_set_policy(C.to_GtkScrolledWindow(v.GWidget), C.GtkPolicyType(hscrollbar_policy), C.GtkPolicyType(vscrollbar_policy))
}
func (v *ScrolledWindow) AddWithViewPort(w WidgetLike) {
	C.gtk_scrolled_window_add_with_viewport(C.to_GtkScrolledWindow(v.GWidget), w.ToNative())
}
func (v *ScrolledWindow) SetPlacement(window_placement CornerType) {
	C.gtk_scrolled_window_set_placement(C.to_GtkScrolledWindow(v.GWidget), C.GtkCornerType(window_placement))
}
func (v *ScrolledWindow) UnsetPlacement() {
	C.gtk_scrolled_window_unset_placement(C.to_GtkScrolledWindow(v.GWidget))
}
func (v *ScrolledWindow) SetShadowType(typ ShadowType) {
	C.gtk_scrolled_window_set_shadow_type(C.to_GtkScrolledWindow(v.GWidget), C.GtkShadowType(typ))
}
func (v *ScrolledWindow) SetHAdjustment(hadjustment *Adjustment) {
	C.gtk_scrolled_window_set_hadjustment(C.to_GtkScrolledWindow(v.GWidget), hadjustment.GAdjustment)
}
func (v *ScrolledWindow) SetVAdjustment(vadjustment *Adjustment) {
	C.gtk_scrolled_window_set_vadjustment(C.to_GtkScrolledWindow(v.GWidget), vadjustment.GAdjustment)
}
func (v *ScrolledWindow) GetPlacement() CornerType {
	return CornerType(C.gtk_scrolled_window_get_placement(C.to_GtkScrolledWindow(v.GWidget)))
}
func (v *ScrolledWindow) GetPolicy(hscrollbar_policy *PolicyType, vscrollbar_policy *PolicyType) {
	var chscrollbar_policy, cvscrollbar_policy C.GtkPolicyType
	C.gtk_scrolled_window_get_policy(C.to_GtkScrolledWindow(v.GWidget), &chscrollbar_policy, &cvscrollbar_policy)
	*hscrollbar_policy = PolicyType(chscrollbar_policy)
	*vscrollbar_policy = PolicyType(cvscrollbar_policy)
}
func (v *ScrolledWindow) GetShadowType() ShadowType {
	return ShadowType(C.gtk_scrolled_window_get_shadow_type(C.to_GtkScrolledWindow(v.GWidget)))
}

//-----------------------------------------------------------------------
// GtkPrintOperation
//-----------------------------------------------------------------------

type PrintOperation struct {
	GPrintOperation *C.GtkPrintOperation
}

type PrintOperationResult int

const (
	PRINT_OPERATION_RESULT_ERROR       PrintOperationResult = 0
	PRINT_OPERATION_RESULT_APPLY       PrintOperationResult = 1
	PRINT_OPERATION_RESULT_CANCEL      PrintOperationResult = 2
	PRINT_OPERATION_RESULT_IN_PROGRESS PrintOperationResult = 3
)

type PrintOperationAction int

const (
	PRINT_OPERATION_ACTION_PRINT_DIALOG PrintOperationAction = 0
	PRINT_OPERATION_ACTION_PRINT        PrintOperationAction = 1
	PRINT_OPERATION_ACTION_PREVIEW      PrintOperationAction = 2
	PRINT_OPERATION_ACTION_EXPOR        PrintOperationAction = 3
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
			C.to_GtkWindow(parent.GWidget),
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
	return gboolean2bool(C.gtk_print_operation_is_finished(v.GPrintOperation))
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
		C.to_GtkAdjustment(C.gtk_adjustment_new(C.gdouble(value), C.gdouble(lower), C.gdouble(upper),
			C.gdouble(step_increment), C.gdouble(page_increment), C.gdouble(page_size)))}
}
func (v *Adjustment) GetValue() float64 {
	return float64(C.gtk_adjustment_get_value(v.GAdjustment))	
}
func (v *Adjustment) SetValue(value float64) {
	C.gtk_adjustment_set_value(v.GAdjustment, C.gdouble(value))
}

// gtk_adjustment_clamp_page
// gtk_adjustment_changed
// gtk_adjustment_value_changed

func (v *Adjustment) Configure(value, lower, upper, step_increment, page_increment, page_size float64) {
	panic_if_version_older(2, 14, 0, "gtk_adjustment_configure()")
	C._gtk_adjustment_configure(v.GAdjustment, C.gdouble(value), C.gdouble(lower), C.gdouble(upper),
		C.gdouble(step_increment), C.gdouble(page_increment), C.gdouble(page_size))
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
	C._gtk_adjustment_set_lower(v.GAdjustment, C.gdouble(lower))
}
func (v *Adjustment) SetPageIncrement(page_increment float64) {
	panic_if_version_older(2, 14, 0, "gtk_adjustment_set_page_increment()")
	C._gtk_adjustment_set_page_increment(v.GAdjustment, C.gdouble(page_increment))
}
func (v *Adjustment) SetPageSize(page_size float64) {
	panic_if_version_older(2, 14, 0, "gtk_adjustment_set_page_size()")
	C._gtk_adjustment_set_page_size(v.GAdjustment, C.gdouble(page_size))
}
func (v *Adjustment) SetStepIncrement(step_increment float64) {
	panic_if_version_older(2, 14, 0, "gtk_adjustment_set_step_increment()")
	C._gtk_adjustment_set_step_increment(v.GAdjustment, C.gdouble(step_increment))
}
func (v *Adjustment) SetUpper(upper float64) {
	panic_if_version_older(2, 14, 0, "gtk_adjustment_set_upper()")
	C._gtk_adjustment_set_upper(v.GAdjustment, C.gdouble(upper))
}

//-----------------------------------------------------------------------
// GtkArrow
//-----------------------------------------------------------------------

// gtk_arrow_new
// gtk_arrow_set

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
	C.gtk_drawing_area_size(C.to_GtkDrawingArea(v.GWidget), C.gint(width), C.gint(height))
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
	SIZE_GROUP_NONE       SizeGroupMode = 0
	SIZE_GROUP_HORIZONTAL SizeGroupMode = 1
	SIZE_GROUP_VERTICAL   SizeGroupMode = 2
	SIZE_GROUP_BOTH       SizeGroupMode = 3
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
	C.gtk_size_group_set_ignore_hidden(v.GSizeGroup, bool2gboolean(ignore_hidden))
}
func (v *SizeGroup) GetIgnoreHidden() bool {
	return gboolean2bool(C.gtk_size_group_get_ignore_hidden(v.GSizeGroup))
}
func (v *SizeGroup) AddWidget(w WidgetLike) {
	C.gtk_size_group_add_widget(v.GSizeGroup, w.ToNative())
}
func (v *SizeGroup) RemoveWidget(w WidgetLike) {
	C.gtk_size_group_remove_widget(v.GSizeGroup, w.ToNative())
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
	defer C.free_string(ptr)
	C.gtk_tooltip_set_markup(v.GTooltip, C.to_gcharptr(ptr))
}
func (v *Tooltip) SetText(text string) {
	ptr := C.CString(text)
	defer C.free_string(ptr)
	C.gtk_tooltip_set_text(v.GTooltip, C.to_gcharptr(ptr))
}
func (v *Tooltip) SetIcon(pixbuf *gdkpixbuf.GdkPixbuf) {
	C.gtk_tooltip_set_icon(v.GTooltip, pixbuf.Pixbuf)
}
func (v *Tooltip) SetIconFromStock(stock_id string, size IconSize) {
	ptr := C.CString(stock_id)
	defer C.free_string(ptr)
	C.gtk_tooltip_set_icon_from_stock(v.GTooltip, C.to_gcharptr(ptr), C.GtkIconSize(size))
}
func (v *Tooltip) SetIconFromIconName(icon_name string, size IconSize) {
	ptr := C.CString(icon_name)
	defer C.free_string(ptr)
	C.gtk_tooltip_set_icon_from_icon_name(v.GTooltip, C.to_gcharptr(ptr), C.GtkIconSize(size))
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
	return &Adjustment{C.gtk_viewport_get_hadjustment(C.to_GtkViewport(v.GWidget))}
}
func (v *Viewport) GetVAdjustment() *Adjustment {
	return &Adjustment{C.gtk_viewport_get_vadjustment(C.to_GtkViewport(v.GWidget))}
}
func (v *Viewport) SetHAdjustment(ha *Adjustment) {
	C.gtk_viewport_set_hadjustment(C.to_GtkViewport(v.GWidget), ha.GAdjustment)
}
func (v *Viewport) SetVAdjustment(va *Adjustment) {
	C.gtk_viewport_set_vadjustment(C.to_GtkViewport(v.GWidget), va.GAdjustment)
}
func (v *Viewport) GetShadowType() ShadowType {
	return ShadowType(C.gtk_viewport_get_shadow_type(C.to_GtkViewport(v.GWidget)))
}
func (v *Viewport) SetShadowType(typ ShadowType) {
	C.gtk_viewport_set_shadow_type(C.to_GtkViewport(v.GWidget), C.GtkShadowType(typ))
}
func (v *Viewport) GetBinWindow() *Window {
	panic_if_version_older_auto(2, 20, 0)
	return &Window{Bin{Container{Widget{C.to_GtkWidget(
		unsafe.Pointer(C._gtk_viewport_get_bin_window(C.to_GtkViewport(v.GWidget))))}}}}
}
func (v *Viewport) GetViewWindow() *Window {
	panic_if_version_older_auto(2, 22, 0)
	return &Window{Bin{Container{Widget{
		C.to_GtkWidget(unsafe.Pointer(C._gtk_viewport_get_view_window(C.to_GtkViewport(v.GWidget))))}}}}
}

//-----------------------------------------------------------------------
// GtkAccessible
//-----------------------------------------------------------------------
type Accessible struct {
	glib.GObject
}

func (v *Accessible) ConnectWidgetDestroyed() {
	C.gtk_accessible_connect_widget_destroyed(
		C.to_GtkAccessible(unsafe.Pointer(v.Object)))
}
func (v *Accessible) SetWidget(w WidgetLike) {
	panic_if_version_older_auto(2, 22, 0)
	C._gtk_accessible_set_widget(
		C.to_GtkAccessible(unsafe.Pointer(v.Object)), w.ToNative())
}
func (v *Accessible) GetWidget() *Widget {
	panic_if_version_older_auto(2, 22, 0)
	return &Widget{C._gtk_accessible_get_widget(
		C.to_GtkAccessible(unsafe.Pointer(v.Object)))}
}

//-----------------------------------------------------------------------
// GtkBin
//-----------------------------------------------------------------------
type Bin struct {
	Container
}

func (v *Bin) GetChild() *Widget {
	return &Widget{C.gtk_bin_get_child(C.to_GtkBin(v.GWidget))}
}

//-----------------------------------------------------------------------
// GtkBox
//-----------------------------------------------------------------------
type PackType int

const (
	PACK_START PackType = 0
	PACK_END   PackType = 1
)

type BoxLike interface {
	ContainerLike
	PackStart(child WidgetLike, expand bool, fill bool, padding uint)
	PackEnd(child WidgetLike, expand bool, fill bool, padding uint)
}
type Box struct {
	Container
}

func (v *Box) PackStart(child WidgetLike, expand bool, fill bool, padding uint) {
	C.gtk_box_pack_start(C.to_GtkBox(v.GWidget), child.ToNative(), bool2gboolean(expand),
		bool2gboolean(fill), C.guint(padding))
}
func (v *Box) PackEnd(child WidgetLike, expand bool, fill bool, padding uint) {
	C.gtk_box_pack_end(C.to_GtkBox(v.GWidget), child.ToNative(), bool2gboolean(expand),
		bool2gboolean(fill), C.guint(padding))
}

//Deprecated since 2.14. Use PackStart() instead.
func (v *Box) PackStartDefaults(child WidgetLike) {
	deprecated_since(2, 14, 0, "gtk_box_pack_start_defaults()")
	C.gtk_box_pack_start_defaults(C.to_GtkBox(v.GWidget), child.ToNative())
}

//Deprecated since 2.14. Use PackEnd() instead.
func (v *Box) PackEndDefaults(child WidgetLike) {
	deprecated_since(2, 14, 0, "gtk_box_pack_end_defaults()")
	C.gtk_box_pack_end_defaults(C.to_GtkBox(v.GWidget), child.ToNative())
}
func (v *Box) GetHomogeneous() bool {
	return gboolean2bool(C.gtk_box_get_homogeneous(C.to_GtkBox(v.GWidget)))
}
func (v *Box) SetHomogeneous(homogeneous bool) {
	C.gtk_box_set_homogeneous(C.to_GtkBox(v.GWidget), bool2gboolean(homogeneous))
}
func (v *Box) GetSpacing() int {
	return int(C.gtk_box_get_spacing(C.to_GtkBox(v.GWidget)))
}
func (v *Box) SetSpacing(spacing int) {
	C.gtk_box_set_spacing(C.to_GtkBox(v.GWidget), C.gint(spacing))
}
func (v *Box) ReorderChild(child WidgetLike, position PackType) {
	C.gtk_box_reorder_child(C.to_GtkBox(v.GWidget), child.ToNative(), C.gint(position))
}
func (v *Box) QueryChildPacking(child WidgetLike, expand *bool, fill *bool, padding *uint, pack_type *PackType) {
	var cexpand, cfill C.gboolean
	var cpadding C.guint
	var cpack_type C.GtkPackType
	C.gtk_box_query_child_packing(C.to_GtkBox(v.GWidget), child.ToNative(), &cexpand, &cfill, &cpadding, &cpack_type)
	*expand = gboolean2bool(cexpand)
	*fill = gboolean2bool(cfill)
	*padding = uint(cpadding)
	*pack_type = PackType(cpack_type)
}
func (v *Box) SetChildPacking(child WidgetLike, expand, fill bool, padding uint, pack_type PackType) {
	C.gtk_box_set_child_packing(C.to_GtkBox(v.GWidget), child.ToNative(), bool2gboolean(expand), bool2gboolean(fill), C.guint(padding), C.GtkPackType(pack_type))
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
type ContainerLike interface {
	WidgetLike
	Add(w WidgetLike)
}
type Container struct {
	Widget
}

func (v *Container) Add(w WidgetLike) {
	C.gtk_container_add(C.to_GtkContainer(v.GWidget), w.ToNative())
}
func (v *Container) Remove(w WidgetLike) {
	C.gtk_container_remove(C.to_GtkContainer(v.GWidget), w.ToNative())
}

// gtk_container_add_with_properties
// gtk_container_get_resize_mode
// gtk_container_set_resize_mode

func (v *Container) CheckResize() {
	C.gtk_container_check_resize(C.to_GtkContainer(v.GWidget))
}

// gtk_container_foreach

func (v *Container) GetChildren() *glib.List {
	return glib.ListFromNative(unsafe.Pointer(C.gtk_container_get_children(C.to_GtkContainer(v.GWidget))))
}

// gtk_container_set_reallocate_redraws
// gtk_container_get_focus_child
// gtk_container_set_focus_child
// gtk_container_get_focus_vadjustment
// gtk_container_set_focus_vadjustment
// gtk_container_get_focus_hadjustment
// gtk_container_set_focus_hadjustment
// gtk_container_resize_children
// gtk_container_child_type
// gtk_container_child_get
// gtk_container_child_set
// gtk_container_child_get_property
// gtk_container_child_set_property
// gtk_container_child_get_valist
// gtk_container_child_set_valist
// gtk_container_forall

func (v *Container) GetBorderWidth() uint {
	return uint(C.gtk_container_get_border_width(C.to_GtkContainer(v.GWidget)))
}
func (v *Container) SetBorderWidth(border_width uint) {
	C.gtk_container_set_border_width(C.to_GtkContainer(v.GWidget), C.guint(border_width))
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
	C.gtk_item_select(C.to_GtkItem(v.GWidget))
}

//Deprecated since 2.22. Use GtkMenuItem.Deselect() instead.
func (v *Item) Deselect() {
	deprecated_since(2, 22, 0, "gtk_item_deselect()")
	C.gtk_item_deselect(C.to_GtkItem(v.GWidget))
}

//Deprecated since 2.22.
func (v *Item) Toggle() {
	deprecated_since(2, 22, 0, "gtk_item_select()")
	C.gtk_item_toggle(C.to_GtkItem(v.GWidget))
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

// gtk_misc_set_alignment
// gtk_misc_set_padding
// gtk_misc_get_alignment
// gtk_misc_get_padding

//-----------------------------------------------------------------------
// GtkObject
//-----------------------------------------------------------------------
type Object struct {
	glib.GObject
}

//deprecated since 2.20

//-----------------------------------------------------------------------
// GtkPaned
//-----------------------------------------------------------------------
type PanedLike interface {
	ContainerLike
	Add1(child WidgetLike)
	Add2(child WidgetLike)
	Pack1(child WidgetLike, expand bool, fill bool)
	Pack2(child WidgetLike, expand bool, fill bool)
}
type Paned struct {
	Container
}

func (v *Paned) Add1(child WidgetLike) {
	C.gtk_paned_add1(C.to_GtkPaned(v.GWidget), child.ToNative())
}
func (v *Paned) Add2(child WidgetLike) {
	C.gtk_paned_add2(C.to_GtkPaned(v.GWidget), child.ToNative())
}
func (v *Paned) Pack1(child WidgetLike, resize bool, shrink bool) {
	C.gtk_paned_pack1(C.to_GtkPaned(v.GWidget), child.ToNative(), bool2gboolean(resize), bool2gboolean(shrink))
}
func (v *Paned) Pack2(child WidgetLike, resize bool, shrink bool) {
	C.gtk_paned_pack2(C.to_GtkPaned(v.GWidget), child.ToNative(), bool2gboolean(resize), bool2gboolean(shrink))
}
func (v *Paned) GetChild1() *Widget {
	return &Widget{
		C.gtk_paned_get_child1(C.to_GtkPaned(v.GWidget))}
}
func (v *Paned) GetChild2() *Widget {
	return &Widget{
		C.gtk_paned_get_child2(C.to_GtkPaned(v.GWidget))}
}
func (v *Paned) SetPosition(position int) {
	C.gtk_paned_set_position(C.to_GtkPaned(v.GWidget), C.gint(position))
}
func (v *Paned) GetPosition() int {
	return int(C.gtk_paned_get_position(C.to_GtkPaned(v.GWidget)))
}

// gtk_paned_get_handle_window //since 2.20

//-----------------------------------------------------------------------
// GtkRange
//-----------------------------------------------------------------------
type Range struct {
	Widget
}

func (v *Range) GetFillLevel() float64 {
	r := C.gtk_range_get_fill_level(C.to_GtkRange(v.GWidget))
	return float64(r)
}
func (v *Range) GetRestrictToFillLevel() bool {
	return gboolean2bool(C.gtk_range_get_restrict_to_fill_level(C.to_GtkRange(v.GWidget)))
}
func (v *Range) GetShowFillLevel() bool {
	return gboolean2bool(C.gtk_range_get_show_fill_level(C.to_GtkRange(v.GWidget)))
}
func (v *Range) SetFillLevel(value float64) {
	C.gtk_range_set_fill_level(C.to_GtkRange(v.GWidget), C.gdouble(value))
}
func (v *Range) SetRestrictToFillLevel(b bool) {
	C.gtk_range_set_restrict_to_fill_level(C.to_GtkRange(v.GWidget), bool2gboolean(b))
}
func (v *Range) SetShowFillLevel(b bool) {
	C.gtk_range_set_show_fill_level(C.to_GtkRange(v.GWidget), bool2gboolean(b))
}
func (v *Range) GetAdjustment() *Adjustment {
	return &Adjustment{C.gtk_range_get_adjustment(C.to_GtkRange(v.GWidget))}
}

// void gtk_range_set_update_policy (GtkRange *range, GtkUpdateType policy); //deprecated in 2.24

func (v *Range) SetAdjustment(a *Adjustment) {
	C.gtk_range_set_adjustment(C.to_GtkRange(v.GWidget), a.GAdjustment)
}
func (v *Range) GetInverted() bool {
	return gboolean2bool(C.gtk_range_get_inverted(C.to_GtkRange(v.GWidget)))
}
func (v *Range) SetInverted(b bool) {
	C.gtk_range_set_inverted(C.to_GtkRange(v.GWidget), bool2gboolean(b))
}

// GtkUpdateType gtk_range_get_update_policy (GtkRange *range); //deprecated since 2.24

func (v *Range) SetIncrements(step, page float64) {
	C.gtk_range_set_increments(C.to_GtkRange(v.GWidget), C.gdouble(step), C.gdouble(page))
}
func (v *Range) SetRange(min, max float64) {
	C.gtk_range_set_range(C.to_GtkRange(v.GWidget), C.gdouble(min), C.gdouble(max))
}
func (v *Range) GetValue() float64 {
	return float64(C.gtk_range_get_value(C.to_GtkRange(v.GWidget))) //TODO test
	//var r C.gdouble
	//C._gtk_range_get_value(C.to_GtkRange(v.GWidget), &r)
	//return float64(r)
}
func (v *Range) SetValue(value float64) {
	C.gtk_range_set_value(C.to_GtkRange(v.GWidget), C.gdouble(value))
}

// gtk_range_get_round_digits //since 2.24
// gtk_range_set_round_digits //since 2.24
// void gtk_range_set_lower_stepper_sensitivity (GtkRange *range, GtkSensitivityType sensitivity);
// GtkSensitivityType gtk_range_get_lower_stepper_sensitivity (GtkRange *range);
// void gtk_range_set_upper_stepper_sensitivity (GtkRange *range, GtkSensitivityType sensitivity);
// GtkSensitivityType gtk_range_get_upper_stepper_sensitivity (GtkRange *range);

func (v *Range) GetFlippable() bool {
	panic_if_version_older(2, 18, 0, "gtk_range_get_flippable()")
	return gboolean2bool(C._gtk_range_get_flippable(C.to_GtkRange(v.GWidget)))
}
func (v *Range) SetFlippable(b bool) {
	panic_if_version_older(2, 18, 0, "gtk_range_set_flippable()")
	C._gtk_range_set_flippable(C.to_GtkRange(v.GWidget), bool2gboolean(b))
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
	POS_LEFT   PositionType = 0
	POS_RIGHT  PositionType = 1
	POS_TOP    PositionType = 2
	POS_BOTTOM PositionType = 3
)

type Scale struct {
	Range
}

func (v *Scale) SetDigits(digits int) {
	C.gtk_scale_set_digits(C.to_GtkScale(v.GWidget), C.gint(digits))
}
func (v *Scale) SetDrawValue(draw_value bool) {
	C.gtk_scale_set_draw_value(C.to_GtkScale(v.GWidget), bool2gboolean(draw_value))
}
func (v *Scale) SetValuePos(pos PositionType) {
	C.gtk_scale_set_value_pos(C.to_GtkScale(v.GWidget), C.GtkPositionType(pos))
}
func (v *Scale) GetDigits() int {
	return int(C.gtk_scale_get_digits(C.to_GtkScale(v.GWidget)))
}
func (v *Scale) GetDrawValue() bool {
	return gboolean2bool(C.gtk_scale_get_draw_value(C.to_GtkScale(v.GWidget)))
}
func (v *Scale) GetValuePos() PositionType {
	return PositionType(C.gtk_scale_get_value_pos(C.to_GtkScale(v.GWidget)))
}

// PangoLayout * gtk_scale_get_layout (GtkScale *scale);

func (v *Scale) GetLayoutOffsets(x *int, y *int) {
	var xx, yy C.gint
	C.gtk_scale_get_layout_offsets(C.to_GtkScale(v.GWidget), &xx, &yy)
	*x = int(xx)
	*y = int(yy)
}
func (v *Scale) AddMark(value float64, position PositionType, markup string) {
	ptr := C.CString(markup)
	defer C.free_string(ptr)
	C.gtk_scale_add_mark(C.to_GtkScale(v.GWidget), C.gdouble(value), C.GtkPositionType(position), C.to_gcharptr(ptr))
}
func (v *Scale) ClearMarks() {
	C.gtk_scale_clear_marks(C.to_GtkScale(v.GWidget))
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
	ACCEL_VISIBLE AccelFlags = 1 << 0
	ACCEL_LOCKED  AccelFlags = 1 << 1
	ACCEL_MASK    AccelFlags = 0x07
)

type StateType int

const (
	STATE_NORMAL      StateType = 0
	STATE_ACTIVE      StateType = 1
	STATE_PRELIGHT    StateType = 2
	STATE_SELECTED    StateType = 3
	STATE_INSENSITIVE StateType = 4
)

type WidgetLike interface {
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
}
type Widget struct {
	GWidget *C.GtkWidget
}

func WidgetFromNative(p unsafe.Pointer) *Widget {
	return &Widget{C.to_GtkWidget(p)}
}

//TODO GtkWidget will have GObject as anonymous field.
func WidgetFromObject(object *glib.GObject) *Widget {
	return &Widget{
		C.to_GtkWidget(unsafe.Pointer(object.Object))}
}
func (v *Widget) ToNative() *C.GtkWidget {
	if v == nil {
		return nil
	}
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

//Deprecated since 2.12. Use g_object_ref() instead. //TODO gobject
func (v *Widget) Ref() {
	deprecated_since(2, 12, 0, "gtk_widget_ref()")
	C.gtk_widget_ref(v.GWidget)
}

//Deprecated since 2.12. Use g_object_ref() instead. //TODO gobject
func (v *Widget) Unref() {
	deprecated_since(2, 12, 0, "gtk_widget_unref()")
	C.gtk_widget_unref(v.GWidget)
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

func (v *Widget) AddAccelerator(signal string, group *AccelGroup, key uint, mods gdk.GdkModifierType, flags AccelFlags) {
	csignal := C.CString(signal)
	defer C.free_string(csignal)
	C.gtk_widget_add_accelerator(v.GWidget, C.to_gcharptr(csignal), group.GAccelGroup, C.guint(key),
		C.GdkModifierType(mods), C.GtkAccelFlags(flags))
}

// gtk_widget_remove_accelerator
// gtk_widget_set_accel_path
// gtk_widget_list_accel_closures

func (v *Widget) CanActivateAccel(signal_id uint) bool {
	return gboolean2bool(C.gtk_widget_can_activate_accel(v.GWidget, C.guint(signal_id)))
}

// gtk_widget_event

func (v *Widget) Activate() {
	C.gtk_widget_activate(v.GWidget)
}
func (v *Widget) Reparent(parent WidgetLike) {
	C.gtk_widget_reparent(v.GWidget, parent.ToNative())
}

// gtk_widget_intersect

func (v *Widget) IsFocus() bool {
	return gboolean2bool(C.gtk_widget_is_focus(v.GWidget))
}
func (v *Widget) GrabFocus() {
	C.gtk_widget_grab_focus(v.GWidget)
}
func (v *Widget) GrabDefault() {
	C.gtk_widget_grab_default(v.GWidget)
}
func (v *Window) SetName(name string) {
	ptr := C.CString(name)
	defer C.free_string(ptr)
	C.gtk_widget_set_name(v.GWidget, C.to_gcharptr(ptr))
}
func (v *Window) GetName() string {
	return C.GoString(C.to_charptr(C.gtk_widget_get_name(v.GWidget)))
}
func (v *Widget) SetState(state StateType) {
	C.gtk_widget_set_state(v.GWidget, C.GtkStateType(state))
}
func (v *Widget) SetSensitive(setting bool) {
	C.gtk_widget_set_sensitive(v.GWidget, bool2gboolean(setting))
}
func (v *Widget) SetParent(parent WidgetLike) {
	C.gtk_widget_set_parent(v.GWidget, parent.ToNative())
}
func (v *Widget) SetParentWindow(parent *gdk.GdkWindow) {
	C.gtk_widget_set_parent_window(v.GWidget, C.to_GdkWindow(unsafe.Pointer(parent.Window)))
}
func (v *Widget) GetParentWindow() *gdk.GdkWindow {
	return gdk.WindowFromUnsafe(unsafe.Pointer(C.gtk_widget_get_parent_window(v.GWidget)))
}

//Deprecated since 2.2, use SetSizeRequest() instead
func (v *Widget) SetUSize(width int, height int) {
	deprecated_since(2, 2, 0, "gtk_widget_set_usize()")
	C.gtk_widget_set_usize(v.GWidget, C.gint(width), C.gint(height))
}
func (v *Widget) SetEvents(events int) {
	C.gtk_widget_set_events(v.GWidget, C.gint(events))
}
func (v *Widget) AddEvents(events int) {
	C.gtk_widget_add_events(v.GWidget, C.gint(events))
}

// gtk_widget_set_extension_events
// gtk_widget_get_extension_events

func (v *Widget) GetTopLevel() *Widget {
	return &Widget{C.gtk_widget_get_toplevel(v.GWidget)}
}

// gtk_widget_get_ancestor
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
// gtk_widget_get_style
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
// gtk_widget_modify_fg
// gtk_widget_modify_bg
// gtk_widget_modify_text
// gtk_widget_modify_base
// gtk_widget_modify_font
// gtk_widget_modify_cursor
// gtk_widget_create_pango_context
// gtk_widget_get_pango_context
// gtk_widget_create_pango_layout

func (v *Widget) RenderIcon(stock_id string, size IconSize, detail string) *gdkpixbuf.GdkPixbuf {
	pstock_id := C.CString(stock_id)
	defer C.free_string(pstock_id)
	pdetail := C.CString(detail)
	defer C.free_string(pdetail)
	return &gdkpixbuf.GdkPixbuf{
		C.gtk_widget_render_icon(v.GWidget, C.to_gcharptr(pstock_id), C.GtkIconSize(size), C.to_gcharptr(pdetail))}
}

// gtk_widget_pop_composite_child
// gtk_widget_push_composite_child

//Deprecated since 2.2. Use QueueDraw() instead.
func (v *Widget) QueueClear() {
	deprecated_since(2, 2, 0, "gtk_widget_queue_clear()")
	C.gtk_widget_queue_clear(v.GWidget)
}

// gtk_widget_queue_draw_area

func (v *Widget) SetAppPrintable(setting bool) {
	C.gtk_widget_set_app_paintable(v.GWidget, bool2gboolean(setting))
}
func (v *Widget) SetDoubleBuffered(setting bool) {
	C.gtk_widget_set_double_buffered(v.GWidget, bool2gboolean(setting))
}
func (v *Widget) SetRedrawOnAllocate(setting bool) {
	C.gtk_widget_set_redraw_on_allocate(v.GWidget, bool2gboolean(setting))
}

// gtk_widget_set_composite_name
// gtk_widget_set_scroll_adjustments

func (v *Widget) MnemonicActivate(group_cycling bool) bool {
	return gboolean2bool(C.gtk_widget_mnemonic_activate(v.GWidget, bool2gboolean(group_cycling)))
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
	return gboolean2bool(C.gtk_widget_get_child_visible(v.GWidget))
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
	C.gtk_widget_set_child_visible(v.GWidget, bool2gboolean(setting))
}
func (v *Widget) SetSizeRequest(width int, height int) {
	C.gtk_widget_set_size_request(v.GWidget, C.gint(width), C.gint(height))
}
func (v *Widget) SetNoShowAll(setting bool) {
	C.gtk_widget_set_no_show_all(v.GWidget, bool2gboolean(setting))
}
func (v *Widget) GetNoShowAll() bool {
	return gboolean2bool(C.gtk_widget_get_no_show_all(v.GWidget))
}

// gtk_widget_list_mnemonic_labels
// gtk_widget_add_mnemonic_label
// gtk_widget_remove_mnemonic_label

func (v *Widget) IsComposited() bool {
	return gboolean2bool(C.gtk_widget_is_composited(v.GWidget))
}

// gtk_widget_error_bell
// gtk_widget_keynav_failed

func (v *Widget) GetTooltipMarkup() string {
	return C.GoString(C.to_charptr(C.gtk_widget_get_tooltip_markup(v.GWidget)))
}
func (v *Widget) SetTooltipMarkup(markup string) {
	ptr := C.CString(markup)
	defer C.free_string(ptr)
	C.gtk_widget_set_tooltip_markup(v.GWidget, C.to_gcharptr(ptr))
}
func (v *Widget) GetTooltipText() string {
	return C.GoString(C.to_charptr(C.gtk_widget_get_tooltip_text(v.GWidget)))
}
func (v *Widget) SetTooltipText(text string) {
	ptr := C.CString(text)
	defer C.free_string(ptr)
	C.gtk_widget_set_tooltip_text(v.GWidget, C.to_gcharptr(ptr))
}
func (v *Widget) GetTooltipWindow() *Window {
	return &Window{Bin{Container{Widget{
		C.to_GtkWidget(unsafe.Pointer(C.gtk_widget_get_tooltip_window(v.GWidget)))}}}}
}
func (v *Widget) SetTooltipWindow(w *Window) {
	C.gtk_widget_set_tooltip_window(v.GWidget, C.to_GtkWindow(w.ToNative()))
}
func (v *Widget) GetHasTooltip() bool {
	return gboolean2bool(C.gtk_widget_get_has_tooltip(v.GWidget))
}
func (v *Widget) SetHasTooltip(setting bool) {
	C.gtk_widget_set_has_tooltip(v.GWidget, bool2gboolean(setting))
}

// gtk_widget_trigger_tooltip_query
// gtk_widget_get_snapshot

func (v *Widget) GetWindow() *gdk.GdkWindow {
	return gdk.WindowFromUnsafe(unsafe.Pointer(C.gtk_widget_get_window(v.GWidget)))
}

//TODO get should return something (this function mechanism is not intuitive)
func (v *Widget) GetAllocation(allocation *Allocation) {
	panic_if_version_older(2, 18, 0, "gtk_widget_get_allocation()")
	var _allocation C.GtkAllocation
	C._gtk_widget_get_allocation(v.GWidget, &_allocation)
	allocation.X = int(_allocation.x)
	allocation.Y = int(_allocation.y)
	allocation.Width = int(_allocation.width)
	allocation.Height = int(_allocation.height)
}
func (v *Widget) SetAllocation(allocation *Allocation) {
	panic_if_version_older(2, 18, 0, "gtk_widget_set_allocation()")
	var _allocation C.GtkAllocation
	_allocation.x = C.gint(allocation.X)
	_allocation.y = C.gint(allocation.Y)
	_allocation.width = C.gint(allocation.Width)
	_allocation.height = C.gint(allocation.Height)
	C._gtk_widget_set_allocation(v.GWidget, &_allocation)
}
func (v *Widget) GetAppPaintable() bool {
	panic_if_version_older(2, 18, 0, "gtk_widget_get_app_paintable()")
	return gboolean2bool(C._gtk_widget_get_app_paintable(v.GWidget))
}
func (v *Widget) GetCanDefault() bool {
	panic_if_version_older(2, 18, 0, "gtk_widget_get_can_default()")
	return gboolean2bool(C._gtk_widget_get_can_default(v.GWidget))
}
func (v *Widget) SetCanDefault(setting bool) {
	panic_if_version_older(2, 18, 0, "gtk_widget_set_can_default()")
	C._gtk_widget_set_can_default(v.GWidget, bool2gboolean(setting))
}
func (v *Widget) GetCanFocus() bool {
	panic_if_version_older(2, 18, 0, "gtk_widget_get_can_focus()")
	return gboolean2bool(C._gtk_widget_get_can_focus(v.GWidget))
}
func (v *Widget) SetCanFocus(setting bool) {
	panic_if_version_older(2, 18, 0, "gtk_widget_set_can_focus()")
	C._gtk_widget_set_can_focus(v.GWidget, bool2gboolean(setting))
}
func (v *Widget) GetDoubleBuffered() bool {
	panic_if_version_older(2, 18, 0, "gtk_widget_get_double_buffered()")
	return gboolean2bool(C._gtk_widget_get_double_buffered(v.GWidget))
}
func (v *Widget) GetHasWindow() bool {
	panic_if_version_older(2, 18, 0, "gtk_widget_get_has_window()")
	return gboolean2bool(C._gtk_widget_get_has_window(v.GWidget))
}
func (v *Widget) SetHasWindow(setting bool) {
	panic_if_version_older(2, 18, 0, "gtk_widget_set_has_window()")
	C._gtk_widget_set_has_window(v.GWidget, bool2gboolean(setting))
}
func (v *Widget) GetSensitive() bool {
	panic_if_version_older(2, 18, 0, "gtk_widget_get_sensitive()")
	return gboolean2bool(C._gtk_widget_get_sensitive(v.GWidget))
}
func (v *Widget) IsSensitive() bool {
	panic_if_version_older(2, 18, 0, "gtk_widget_is_sensitive()")
	return gboolean2bool(C._gtk_widget_is_sensitive(v.GWidget))
}
func (v *Widget) GetState() StateType {
	panic_if_version_older(2, 18, 0, "gtk_widget_get_state()")
	return StateType(C._gtk_widget_get_state(v.GWidget))
}
func (v *Widget) GetVisible() bool {
	panic_if_version_older(2, 18, 0, "gtk_widget_get_visible()")
	return gboolean2bool(C._gtk_widget_get_visible(v.GWidget))
}
func (v *Widget) SetVisible(setting bool) {
	panic_if_version_older(2, 18, 0, "gtk_widget_set_visible()")
	C._gtk_widget_set_visible(v.GWidget, bool2gboolean(setting))
}
func (v *Widget) HasDefault() bool {
	panic_if_version_older(2, 18, 0, "gtk_widget_has_default()")
	return gboolean2bool(C._gtk_widget_has_default(v.GWidget))
}
func (v *Widget) HasFocus() bool {
	panic_if_version_older(2, 18, 0, "gtk_widget_has_focus()")
	return gboolean2bool(C._gtk_widget_has_focus(v.GWidget))
}
func (v *Widget) HasGrab() bool {
	panic_if_version_older(2, 18, 0, "gtk_widget_has_grab()")
	return gboolean2bool(C._gtk_widget_has_grab(v.GWidget))
}

// gtk_widget_has_rc_style //since 2.20

func (v *Widget) IsDrawable() bool {
	panic_if_version_older(2, 18, 0, "gtk_widget_is_drawable()")
	return gboolean2bool(C._gtk_widget_is_drawable(v.GWidget))
}
func (v *Widget) IsToplevel() bool {
	panic_if_version_older(2, 18, 0, "gtk_widget_is_toplevel()")
	return gboolean2bool(C._gtk_widget_is_toplevel(v.GWidget))
}
func (v *Widget) SetWindow(window *gdk.GdkWindow) {
	panic_if_version_older(2, 18, 0, "gtk_widget_set_window()")
	C._gtk_widget_set_window(v.GWidget, C.to_GdkWindow(unsafe.Pointer(window.Window)))
}
func (v *Widget) SetReceivesDefault(setting bool) {
	panic_if_version_older(2, 18, 0, "gtk_widget_set_receives_default()")
	C._gtk_widget_set_receives_default(v.GWidget, bool2gboolean(setting))
}
func (v *Widget) GetReceivesDefault() bool {
	panic_if_version_older(2, 18, 0, "gtk_widget_get_receives_default()")
	return gboolean2bool(C._gtk_widget_get_receives_default(v.GWidget))
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
	defer C.free_string(pdesc)
	C.gtk_widget_modify_font(v.GWidget, C.pango_font_description_from_string(pdesc))
}

func (v *Widget) ModifyBG(state StateType, color *gdk.GdkColor) {
	C.gtk_widget_modify_bg(v.GWidget, C.GtkStateType(state), (*C.GdkColor)(unsafe.Pointer(&color.Color)))
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
	defer C.free_string(ptr)
	ret = uint(C.gtk_builder_add_from_file(v.GBuilder, C.to_gcharptr(ptr), &gerror))
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
	defer C.free_string(ptr)
	ret = uint(C.gtk_builder_add_from_string(v.GBuilder, C.to_gcharptr(ptr), C.gsize(C.strlen(ptr)), &gerror))
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
	defer C.free_string(ptr)
	return &glib.GObject{
		unsafe.Pointer(C.gtk_builder_get_object(v.GBuilder, C.to_gcharptr(ptr)))}
}
func (v *Builder) GetObjects() *glib.SList {
	return glib.SListFromNative(unsafe.Pointer(C.gtk_builder_get_objects(v.GBuilder)))
}
func (v *Builder) ConnectSignals(user_data interface{}) {
	C.gtk_builder_connect_signals(v.GBuilder, nil)
}

// void gtk_builder_connect_signals_full (GtkBuilder *builder, GtkBuilderConnectFunc func, gpointer user_data);

func (v *Builder) SetTranslationDomain(domain string) {
	ptr := C.CString(domain)
	defer C.free_string(ptr)
	C.gtk_builder_set_translation_domain(v.GBuilder, C.to_gcharptr(ptr))
}
func (v *Builder) GetTranslationDomain() string {
	return C.GoString(C.to_charptr(C.gtk_builder_get_translation_domain(v.GBuilder)))
}
func (v *Builder) GetTypeFromName(type_name string) int {
	ptr := C.CString(type_name)
	defer C.free_string(ptr)
	return int(C.gtk_builder_get_type_from_name(v.GBuilder, ptr))
}

// gboolean gtk_builder_value_from_string (GtkBuilder *builder, GParamSpec *pspec, const gchar *string, GValue *value, GError **error);
// gboolean gtk_builder_value_from_string_type (GtkBuilder *builder, GType type, const gchar *string, GValue *value, GError **error);
