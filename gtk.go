package gtk

/*
#include <gtk/gtk.h>
#include <unistd.h>
#include <stdlib.h>

int _gtk_callback_count = 0;
int _gtk_callback_index = 0;
int _gtk_callback_func_no[200];
GtkWidget* _gtk_callback_widget[200];
void* _gtk_callback_data[200];
int _gtk_callback_fire[200];
static void _callback(GtkWidget* w, void* data1, void* data2) {
	if (GTK_IS_DIALOG(w)) {
		// data1 is dialog response.
		_gtk_callback_index = (int)data2;
	} else {
		_gtk_callback_index = (int)data1;
	}
	_gtk_callback_fire[_gtk_callback_index-1] = 0;
	_gtk_callback_fire[1] = 0;
	_gtk_callback_fire[2] = 0;
	_gtk_callback_fire[3] = 0;
	printf("index=%d\n", (int)_gtk_callback_index);
}

static void _gtk_init(void* argc, void* argv) {
	gtk_init((int*)argc, (char***)argv);
}

static void _gtk_container_add(GtkWidget* container, GtkWidget* widget) {
	gtk_container_add(GTK_CONTAINER(container), widget);
}

static long _gtk_signal_connect(GtkWidget* widget, char* name, int func_no, void* data) {
	_gtk_callback_func_no[_gtk_callback_count++] = func_no;
	printf("count=%d\n", _gtk_callback_count);
	return gtk_signal_connect_full(GTK_OBJECT(widget), name, GTK_SIGNAL_FUNC(_callback), 0, (void*)_gtk_callback_count, 0, 0, 1);
}

static char* _gtk_window_get_title(GtkWidget* widget) {
	return (char*)gtk_window_get_title(GTK_WINDOW(widget));
}

static void _gtk_window_set_title(GtkWidget* widget, char* title) {
	gtk_window_set_title(GTK_WINDOW(widget), (gchar*)title);
}

static void _gtk_window_set_transient_for(GtkWidget* widget, GtkWidget *parent) {
	gtk_window_set_transient_for(GTK_WINDOW(widget), GTK_WINDOW(parent));
}

static int _gtk_dialog_run(GtkWidget* dialog) {
	return gtk_dialog_run(GTK_DIALOG(dialog));
}

static GtkWidget* _gtk_message_dialog_new(GtkWidget* parent, GtkDialogFlags flags, GtkMessageType type, GtkButtonsType buttons, char *message) {
	return gtk_message_dialog_new(
			GTK_WINDOW(parent),
			flags,
			type,
			buttons,
			(gchar*)message, NULL);
}

static char* _gtk_entry_get_text(GtkWidget* widget) {
	return (char*)gtk_entry_get_text(GTK_ENTRY(widget));
}

static void _gtk_entry_set_text(GtkWidget* widget, char* text) {
	gtk_entry_set_text(GTK_ENTRY(widget), (gchar*)text);
}

static char* _gtk_label_get_text(GtkWidget* widget) {
	return (char*)gtk_label_get_text(GTK_LABEL(widget));
}

static void _gtk_label_set_text(GtkWidget* widget, char* text) {
	gtk_label_set_text(GTK_LABEL(widget), (gchar*)text);
}

static GtkWidget* _gtk_accel_label_get_accel_widget(GtkWidget* widget) {
	return gtk_accel_label_get_accel_widget(GTK_ACCEL_LABEL(widget));
}

static guint _gtk_accel_label_get_accel_width(GtkWidget* widget) {
	return gtk_accel_label_get_accel_width(GTK_ACCEL_LABEL(widget));
}

static void _gtk_accel_label_set_accel_widget(GtkWidget* label, GtkWidget* widget) {
	gtk_accel_label_set_accel_widget(GTK_ACCEL_LABEL(label), widget);
}

static gboolean _gtk_accel_label_refetch(GtkWidget* widget) {
	return gtk_accel_label_refetch(GTK_ACCEL_LABEL(widget));
}

static char* _gtk_button_get_label(GtkWidget* widget) {
	return (char*)gtk_button_get_label(GTK_BUTTON(widget));
}

static void _gtk_button_set_label(GtkWidget* widget, char* label) {
	gtk_button_set_label(GTK_BUTTON(widget), (gchar*)label);
}

static void _gtk_box_pack_start(GtkWidget* box, GtkWidget* child, gboolean expand, gboolean fill, guint padding) {
	gtk_box_pack_start(GTK_BOX(box), child, expand, fill, padding);
}

static void _gtk_box_pack_end(GtkWidget* box, GtkWidget* child, gboolean expand, gboolean fill, guint padding) {
	gtk_box_pack_end(GTK_BOX(box), child, expand, fill, padding);
}
static gchar* to_gcharptr(char* s) { return (gchar*)s; }
*/
import "C";
import "unsafe";
import "runtime";
import "container/vector";

func bool2gboolean(b bool) C.gboolean { if b { return C.gboolean(1) } return C.gboolean(0) }
func gboolean2bool(b C.gboolean) bool { if b != 0 { return true } return false }

//-----------------------------------------------------------------------
// GtkWidget
//-----------------------------------------------------------------------
type Widget interface {
	ToGtkWidget() *C.GtkWidget;
	Add(w Widget);
}
type GtkWidget struct {
	Widget *C.GtkWidget;
}
func (v GtkWidget) ToGtkWidget() *C.GtkWidget { return v.Widget }
func Hide(v Widget) { C.gtk_widget_hide(v.ToGtkWidget()) }
func HideAll(v Widget) { C.gtk_widget_hide_all(v.ToGtkWidget()) }
func Show(v Widget) { C.gtk_widget_show(v.ToGtkWidget()) }
func ShowAll(v Widget) { C.gtk_widget_show_all(v.ToGtkWidget()) }
func ShowNow(v Widget) { C.gtk_widget_show_now(v.ToGtkWidget()) }
func Destroy(v Widget) { C.gtk_widget_destroy(v.ToGtkWidget()) }
func (v GtkWidget) Add(w Widget) {
	C._gtk_container_add(v.ToGtkWidget(), w.ToGtkWidget());
}
func Connect(v Widget, s string, f func()) {
	funcs.Push(&Callback{f});
	C._gtk_signal_connect(v.ToGtkWidget(), C.CString(s),
		                    C.int(funcs.Len())-1, nil);
} // GtkContainer
func GetTopLevel(v Widget) GtkWidget {
	return GtkWidget{ C.gtk_widget_get_toplevel(v.ToGtkWidget()) };
}
func HideOnDelete(v Widget) {
	C.gtk_widget_hide_on_delete(v.ToGtkWidget());
}
func QueueResize(v Widget) {
	C.gtk_widget_queue_resize(v.ToGtkWidget());
}

// TODO
// gtk_widget_destroyed
// gtk_widget_ref
// gtk_widget_unref
// gtk_widget_set
// gtk_widget_unparent
// gtk_widget_set_no_show_all
// gtk_widget_get_no_show_all
// gtk_widget_map
// gtk_widget_unmap
// gtk_widget_realize
// gtk_widget_unrealize
// gtk_widget_queue_draw
// gtk_widget_queue_draw_area
// gtk_widget_queue_clear
// gtk_widget_queue_clear_area
// gtk_widget_queue_resize
// gtk_widget_queue_resize_no_redraw
// gtk_widget_draw
// gtk_widget_size_request
// gtk_widget_size_allocate
// gtk_widget_get_child_requisition
// gtk_widget_add_accelerator
// gtk_widget_remove_accelerator
// gtk_widget_set_accel_path
// gtk_widget_list_accel_closures
// gtk_widget_can_activate_accel
// gtk_widget_mnemonic_activate
// gtk_widget_event
// gtk_widget_send_expose
// gtk_widget_activate
// gtk_widget_set_scroll_adjustments
// gtk_widget_reparent
// gtk_widget_intersect
// gtk_widget_region_intersect
// gtk_widget_set_can_focus
// gtk_widget_get_can_focus
// gtk_widget_has_focus
// gtk_widget_is_focus
// gtk_widget_grab_focus
// gtk_widget_set_can_default
// gtk_widget_get_can_default
// gtk_widget_has_default
// gtk_widget_grab_default
// gtk_widget_set_receives_default
// gtk_widget_get_receives_default
// gtk_widget_has_grab
// gtk_widget_set_name
// gtk_widget_get_name
// gtk_widget_set_state
// gtk_widget_get_state
// gtk_widget_set_sensitive
// gtk_widget_get_sensitive
// gtk_widget_is_sensitive
// gtk_widget_set_visible
// gtk_widget_get_visible
// gtk_widget_set_has_window
// gtk_widget_get_has_window
// gtk_widget_is_toplevel
// gtk_widget_is_drawable
// gtk_widget_set_app_paintable
// gtk_widget_get_app_paintable
// gtk_widget_set_double_buffered
// gtk_widget_get_double_buffered
// gtk_widget_set_redraw_on_allocate
// gtk_widget_set_parent
// gtk_widget_get_parent
// gtk_widget_set_parent_window
// gtk_widget_get_parent_window
// gtk_widget_set_child_visible
// gtk_widget_get_child_visible
// gtk_widget_set_window
// gtk_widget_get_window
// gtk_widget_get_allocation
// gtk_widget_set_allocation
// gtk_widget_child_focus
// gtk_widget_keynav_failed
// gtk_widget_error_bell
// gtk_widget_set_size_request
// gtk_widget_get_size_request
// gtk_widget_set_uposition
// gtk_widget_set_usize
// gtk_widget_set_events
// gtk_widget_add_events
// gtk_widget_set_extension_events
// gtk_widget_get_extension_events
// gtk_widget_get_ancestor
// gtk_widget_get_colormap
// gtk_widget_get_visual
// gtk_widget_get_screen
// gtk_widget_has_screen
// gtk_widget_get_display
// gtk_widget_get_root_window
// gtk_widget_get_settings
// gtk_widget_get_clipboard
// gtk_widget_get_snapshot
// gtk_widget_get_accessible
// gtk_widget_set_colormap
// gtk_widget_get_events
// gtk_widget_get_pointer
// gtk_widget_is_ancestor
// gtk_widget_translate_coordinates
// gtk_widget_hide_on_delete
// gtk_widget_set_style
// gtk_widget_ensure_style
// gtk_widget_get_style
// gtk_widget_modify_style
// gtk_widget_get_modifier_style
// gtk_widget_modify_fg
// gtk_widget_modify_bg
// gtk_widget_modify_text
// gtk_widget_modify_base
// gtk_widget_modify_cursor
// gtk_widget_modify_font
// gtk_widget_create_pango_context
// gtk_widget_get_pango_context
// gtk_widget_create_pango_layout
// gtk_widget_render_icon
// gtk_widget_set_composite_name
// gtk_widget_get_composite_name
// gtk_widget_reset_rc_styles
// gtk_widget_push_colormap
// gtk_widget_push_composite_child
// gtk_widget_pop_composite_child
// gtk_widget_pop_colormap
// gtk_widget_class_install_style_property
// gtk_widget_class_install_style_property_parser
// gtk_widget_class_find_style_property
// gtk_widget_class_list_style_properties
// gtk_widget_style_get_property
// gtk_widget_style_get_valist
// gtk_widget_style_get
// gtk_widget_set_default_colormap
// gtk_widget_get_default_style
// gtk_widget_get_default_colormap
// gtk_widget_get_default_visual
// gtk_widget_set_direction
// gtk_widget_get_direction
// gtk_widget_set_default_direction
// gtk_widget_get_default_direction
// gtk_widget_is_composited
// gtk_widget_shape_combine_mask
// gtk_widget_input_shape_combine_mask
// gtk_widget_reset_shapes
// gtk_widget_path
// gtk_widget_class_path
// gtk_widget_list_mnemonic_labels
// gtk_widget_add_mnemonic_label
// gtk_widget_remove_mnemonic_label
// gtk_widget_set_tooltip_window
// gtk_widget_get_tooltip_window
// gtk_widget_trigger_tooltip_query
// gtk_widget_set_tooltip_text
// gtk_widget_get_tooltip_text
// gtk_widget_set_tooltip_markup
// gtk_widget_get_tooltip_markup
// gtk_widget_set_has_tooltip
// gtk_widget_get_has_tooltip
// gtk_requisition_get_type
// gtk_requisition_copy
// gtk_requisition_free

//-----------------------------------------------------------------------
// GtkWindow
//-----------------------------------------------------------------------
const (
	GTK_WINDOW_TOPLEVEL = 0;
	GTK_WINDOW_POPUP = 1;
);
type WindowLike interface {
	Labelled;
	SetTransientFor(parent WindowLike);
}
type GtkWindow struct { GtkWidget; }
func Window(t int) WindowLike {
	return GtkWindow{ GtkWidget { C.gtk_window_new(C.GtkWindowType(t)) } };
}
func (v GtkWindow) GetLabel() string { return C.GoString(C._gtk_window_get_title(v.Widget)); }
func (v GtkWindow) SetLabel(title string) { C._gtk_window_set_title(v.Widget, C.CString(title)); }
func (v GtkWindow) SetTransientFor(parent WindowLike) {
	C._gtk_window_set_transient_for(v.Widget, parent.ToGtkWidget());
}
// TODO
// gtk_window_set_wmclass
// gtk_window_set_role
// gtk_window_set_startup_id
// gtk_window_get_role
// gtk_window_add_accel_group
// gtk_window_remove_accel_group
// gtk_window_set_position
// gtk_window_activate_focus	
// gtk_window_set_focus
// gtk_window_get_focus
// gtk_window_set_default
// gtk_window_get_default_widget
// gtk_window_activate_default	
// gtk_window_set_transient_for
// gtk_window_get_transient_for
// gtk_window_set_opacity
// gtk_window_get_opacity
// gtk_window_set_type_hint
// gtk_window_get_type_hint
// gtk_window_set_skip_taskbar_hint
// gtk_window_get_skip_taskbar_hint
// gtk_window_set_skip_pager_hint
// gtk_window_get_skip_pager_hint
// gtk_window_set_urgency_hint
// gtk_window_get_urgency_hint
// gtk_window_set_accept_focus
// gtk_window_get_accept_focus
// gtk_window_set_focus_on_map
// gtk_window_get_focus_on_map
// gtk_window_set_destroy_with_parent
// gtk_window_get_destroy_with_parent
// gtk_window_set_resizable
// gtk_window_get_resizable
// gtk_window_set_gravity
// gtk_window_get_gravity
// gtk_window_set_geometry_hints
// gtk_window_set_screen	
// gtk_window_get_screen	
// gtk_window_is_active
// gtk_window_has_toplevel_focus
// gtk_window_set_has_frame
// gtk_window_get_has_frame
// gtk_window_set_frame_dimensions
// gtk_window_get_frame_dimensions
// gtk_window_set_decorated
// gtk_window_get_decorated
// gtk_window_set_deletable
// gtk_window_get_deletable
// gtk_window_set_icon_list
// gtk_window_get_icon_list
// gtk_window_set_icon
// gtk_window_set_icon_name
// gtk_window_set_icon_from_file
// gtk_window_get_icon
// gtk_window_get_icon_name
// gtk_window_set_default_icon_list
// gtk_window_get_default_icon_list
// gtk_window_set_default_icon
// gtk_window_set_default_icon_name
// gtk_window_get_default_icon_name
// gtk_window_set_default_icon_from_file
// gtk_window_set_auto_startup_notification
// gtk_window_set_modal
// gtk_window_get_modal
// gtk_window_list_toplevels
// gtk_window_add_mnemonic
// gtk_window_remove_mnemonic
// gtk_window_mnemonic_activate
// gtk_window_set_mnemonic_modifier
// gtk_window_get_mnemonic_modifier
// gtk_window_activate_key
// gtk_window_propagate_key_event
// gtk_window_present
// gtk_window_present_with_time
// gtk_window_iconify
// gtk_window_deiconify
// gtk_window_stick
// gtk_window_unstick
// gtk_window_maximize
// gtk_window_unmaximize
// gtk_window_fullscreen
// gtk_window_unfullscreen
// gtk_window_set_keep_above
// gtk_window_set_keep_below
// gtk_window_begin_resize_drag
// gtk_window_begin_move_drag
// gtk_window_set_policy
// gtk_window_set_default_size
// gtk_window_get_default_size
// gtk_window_resize
// gtk_window_get_size
// gtk_window_move
// gtk_window_get_position
// gtk_window_parse_geometry
// gtk_window_get_group
// gtk_window_reshow_with_initial_size
// gtk_window_group_get_type
// gtk_window_group_new
// gtk_window_group_add_window
// gtk_window_group_remove_window
// gtk_window_group_list_windows
// gtk_window_remove_embedded_xid
// gtk_window_add_embedded_xid

//-----------------------------------------------------------------------
// GtkDialog
//-----------------------------------------------------------------------
const (
	GTK_DIALOG_MODAL               = 1 << 0; /* call gtk_window_set_modal (win, TRUE) */
	GTK_DIALOG_DESTROY_WITH_PARENT = 1 << 1; /* call gtk_window_set_destroy_with_parent () */
	GTK_DIALOG_NO_SEPARATOR        = 1 << 2; /* no separator bar above buttons */
)
type Dialog interface {
	Widget;
	Run() int;
	Response(func ());
}
type GtkDialog struct { GtkWidget };
func (v GtkDialog) Run() int {
	return int(C._gtk_dialog_run(v.Widget));
}
func (v GtkDialog) Response(response func ()) {
	Connect(v, "response", response);
}

//-----------------------------------------------------------------------
// GtkMessageDialog
//-----------------------------------------------------------------------
const (
	GTK_MESSAGE_INFO = 0;
	GTK_MESSAGE_WARNING = 1;
	GTK_MESSAGE_QUESTION = 2;
	GTK_MESSAGE_ERROR = 3;
	GTK_MESSAGE_OTHER = 4;
)
const (
	GTK_BUTTONS_NONE = 0;
	GTK_BUTTONS_OK = 1;
	GTK_BUTTONS_CLOSE = 2;
	GTK_BUTTONS_CANCEL = 3;
	GTK_BUTTONS_YES_NO = 4;
	GTK_BUTTONS_OK_CANCEL = 5;
)
type GtkMessageDialog GtkDialog;
func MessageDialog(parent WindowLike, flag int, t int, button int,
                   message string) Dialog {
	return GtkDialog { GtkWidget {
		C._gtk_message_dialog_new(
			parent.ToGtkWidget(),
			C.GtkDialogFlags(flag),
			C.GtkMessageType(t),
			C.GtkButtonsType(button),
			C.CString(message))
	}};
}
 
//-----------------------------------------------------------------------
// GtkBox
//-----------------------------------------------------------------------
type Packable interface {
	Widget;
	PackStart(child Widget, expand bool, fill bool, padding uint);
	PackEnd(child Widget, expand bool, fill bool, padding uint);
}
type GtkBox struct { GtkWidget };
func (v GtkBox) PackStart(child Widget, expand bool, fill bool, padding uint) {
	C._gtk_box_pack_start(v.Widget, child.ToGtkWidget(), bool2gboolean(expand), bool2gboolean(fill), C.guint(padding));
}
func (v GtkBox) PackEnd(child Widget, expand bool, fill bool, padding uint) {
	C._gtk_box_pack_end(v.Widget, child.ToGtkWidget(), bool2gboolean(expand), bool2gboolean(fill), C.guint(padding));
}
// TODO
// gtk_box_pack_start_defaults
// gtk_box_pack_end_defaults
// gtk_box_set_homogeneous
// gtk_box_get_homogeneous
// gtk_box_set_spacing
// gtk_box_get_spacing
// gtk_box_reorder_child
// gtk_box_query_child_packing
// gtk_box_set_child_packing

//-----------------------------------------------------------------------
// GtkVBox
//-----------------------------------------------------------------------
func VBox(homogeneous bool, spacing uint) Packable {
	return GtkBox{ GtkWidget {
		C.gtk_vbox_new(bool2gboolean(homogeneous), C.gint(spacing))
	}};
}

//-----------------------------------------------------------------------
// GtkHBox
//-----------------------------------------------------------------------
type GtkHBox GtkBox;
func HBox(homogeneous bool, spacing uint) Packable {
	return GtkBox{ GtkWidget {
		C.gtk_hbox_new(bool2gboolean(homogeneous), C.gint(spacing))
	}};
}

//-----------------------------------------------------------------------
// GtkEntry
//-----------------------------------------------------------------------
type GtkEntry struct { GtkWidget; }
func Entry() Labelled { return GtkEntry{ GtkWidget { C.gtk_entry_new()} }; }
func (v GtkEntry) GetLabel() string { return C.GoString(C._gtk_entry_get_text(v.Widget)); }
func (v GtkEntry) SetLabel(text string) { C._gtk_entry_set_text(v.Widget, C.CString(text)); }

//-----------------------------------------------------------------------
// GtkLabel
//-----------------------------------------------------------------------
type Labelled interface {
	Widget; // labels are Widgets!
	GetLabel() string;
	SetLabel(label string);
}
type GtkLabel struct { GtkWidget; }
func Label(label string) Labelled {
	return GtkLabel{ GtkWidget {
		C.gtk_label_new(C.to_gcharptr(C.CString(label)))
	}};
}

func (v GtkLabel) GetLabel() string { return C.GoString(C._gtk_label_get_text(v.Widget)); }
func (v GtkLabel) SetLabel(label string) { C._gtk_label_set_text(v.Widget, C.CString(label)); }
// TODO
// gtk_label_new_with_mnemonic
// gtk_label_set_attributes
// gtk_label_get_attributes
// gtk_label_set_markup
// gtk_label_set_use_markup
// gtk_label_get_use_markup
// gtk_label_set_use_underline
// gtk_label_get_use_underline
// gtk_label_set_markup_with_mnemonic
// gtk_label_get_mnemonic_keyval
// gtk_label_set_mnemonic_widget
// gtk_label_get_mnemonic_widget
// gtk_label_set_text_with_mnemonic
// gtk_label_set_justify
// gtk_label_get_justify
// gtk_label_set_ellipsize		
// gtk_label_get_ellipsize
// gtk_label_set_width_chars		
// gtk_label_get_width_chars
// gtk_label_set_max_width_chars
// gtk_label_get_max_width_chars
// gtk_label_set_pattern
// gtk_label_set_line_wrap
// gtk_label_get_line_wrap
// gtk_label_set_line_wrap_mode
// gtk_label_get_line_wrap_mode
// gtk_label_set_selectable
// gtk_label_get_selectable
// gtk_label_set_angle
// gtk_label_get_angle
// gtk_label_select_region
// gtk_label_get_selection_bounds
// gtk_label_get_layout
// gtk_label_get_layout_offsets
// gtk_label_set_single_line_mode
// gtk_label_get_single_line_mode
// gtk_label_get_current_uri
// gtk_label_set_track_visited_links
// gtk_label_get_track_visited_links
// gtk_label_get
// gtk_label_parse_uline

//-----------------------------------------------------------------------
// GtkAccelLabel
//-----------------------------------------------------------------------
type AccelLabelled interface {
	Widget; // labels are Widgets!
	GetAccelWidget() GtkWidget;
	SetAccelWidget(GtkWidget);
}
type GtkAccelLabel struct {
	GtkWidget;
}
func AccelLabel(label string) AccelLabelled {
	return GtkAccelLabel{ GtkWidget {
		C.gtk_accel_label_new(C.to_gcharptr(C.CString(label)))
	}};
}
func (v GtkAccelLabel) GetAccelWidget() GtkWidget { return GtkWidget{ C._gtk_accel_label_get_accel_widget(v.Widget) }; }
func (v GtkAccelLabel) GetAccelWidth() uint { return uint(C._gtk_accel_label_get_accel_width(v.Widget)); }
func (v GtkAccelLabel) SetAccelWidget(w GtkWidget) { C._gtk_accel_label_set_accel_widget(v.Widget, w.Widget); }
func (v GtkAccelLabel) Refetch() bool { return gboolean2bool(C._gtk_accel_label_refetch(v.Widget)); }

//-----------------------------------------------------------------------
// GtkButton
//-----------------------------------------------------------------------
type ButtonLike interface { // Buttons are Labelled Widgets!
	Widget;
	// the following should be just Clickable; ...
	Clicked(func()); // this is a very simple interface...
	GetLabel() string;
	SetLabel(label string);
}
type Clickable interface {
	Widget;
	Clicked(func()); // this is a very simple interface...
}
func (v GtkButton) Clicked(onclick func()) {
	Connect(v, "clicked", onclick);
}
type GtkButton struct { GtkWidget; }
func Button() ButtonLike { return GtkButton{ GtkWidget {C.gtk_button_new()} }; }
func ButtonWithLabel(label string) ButtonLike {
	return GtkButton{ GtkWidget {
		C.gtk_button_new_with_label(C.to_gcharptr(C.CString(label)))
	}};
}
func (v GtkButton) GetLabel() string { return C.GoString(C._gtk_button_get_label(v.Widget)); }
func (v GtkButton) SetLabel(label string) { C._gtk_button_set_label(v.Widget, C.CString(label)); }
// TODO
// gtk_button_new_from_stock
// gtk_button_new_with_mnemonic
// gtk_button_set_relief
// gtk_button_get_relief
// gtk_button_set_use_underline
// gtk_button_get_use_underline
// gtk_button_set_use_stock
// gtk_button_get_use_stock
// gtk_button_set_focus_on_click
// gtk_button_get_focus_on_click
// gtk_button_set_alignment
// gtk_button_get_alignment
// gtk_button_set_image
// gtk_button_get_image
// gtk_button_set_image_position
// gtk_button_get_image_position

//-----------------------------------------------------------------------
// Events
var use_gtk_main bool = false;

// the go-gtk Callback is simpler than the one in C, because we have
// full closures, so there is never a need to pass additional data via
// a void * pointer.  Where you might have wanted to do that, you can
// instead just use func () { ... using data } to pass the data in.
type Callback struct {
	f func();
}
var funcs *vector.Vector;
var main_loop bool = true;
func pollEvents() {
	for main_loop {
		if use_gtk_main == false {
			C.gtk_main_iteration_do(C.gboolean(1));
		}
		index := C._gtk_callback_index;
		if int(index) > 0 && C._gtk_callback_fire[index-1] == 0 {
			elem := funcs.At(int(C._gtk_callback_func_no[index-1]));
			f := elem.(*Callback);
			f.f();
			C._gtk_callback_fire[index-1] = C.int(1);
		}
	}
}

func Init(args *[]string) {
	runtime.GOMAXPROCS(10);
	var argc C.int = C.int(len(*args));
	cargs := make([]*C.char, argc);
	for i, arg := range *args { cargs[i] = C.CString(arg) }
	C._gtk_init(unsafe.Pointer(&argc), unsafe.Pointer(&cargs));
	goargs := make([]string, argc);
	for i := 0;i < int(argc); i++ { goargs[i] = C.GoString(cargs[i]); }
	*args = goargs;

	funcs = new(vector.Vector);
}

func Main() {
	if use_gtk_main {
		go pollEvents();
		C.gtk_main();
	} else {
		pollEvents();
	}
}
func MainQuit() {
	main_loop = false;
	if use_gtk_main {
		C.gtk_main_quit();
	}
}
