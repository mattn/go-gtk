package gtk

/*
#include <gtk/gtk.h>
#include <unistd.h>
#include <stdlib.h>

static gchar* to_gcharptr(char* s) { return (gchar*)s; }

typedef struct {
	int func_no;
	void* data;
} _callback_struct;

int _gtk_enter_callback = 0;
int _gtk_callback_func_no = -1;
GtkWidget* _gtk_callback_widget = NULL;
void* _gtk_callback_data = NULL;
static void _callback(GtkWidget* w, void* data) {
	_callback_struct* s = (_callback_struct*)data;
	_gtk_callback_func_no = s->func_no;
	_gtk_callback_widget = w;
	_gtk_callback_data = s->data;
	_gtk_enter_callback = 1;
}

static void _gtk_init(void* argc, void* argv) {
	gtk_init((int*)argc, (char***)argv);
}

static void _gtk_container_add(GtkWidget* container, GtkWidget* widget) {
	gtk_container_add(GTK_CONTAINER(container), widget);
}

static long _gtk_signal_connect(GtkWidget* widget, char* name, int func_no, void* data) {
	_callback_struct* s = (_callback_struct*)malloc(sizeof(_callback_struct));
	s->func_no = func_no;
	s->data = data;
	return gtk_signal_connect_full(GTK_OBJECT(widget), name, GTK_SIGNAL_FUNC(_callback), 0, s, 0, 0, 0);
}

static char* _gtk_window_get_title(GtkWidget* widget) {
	return (char*)gtk_window_get_title(GTK_WINDOW(widget));
}

static void _gtk_window_set_title(GtkWidget* widget, char* title) {
	gtk_window_set_title(GTK_WINDOW(widget), (gchar*)title);
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

static char* _gtk_button_get_label(GtkWidget* widget) {
	return (char*)gtk_button_get_label(GTK_BUTTON(widget));
}

static void _gtk_button_set_label(GtkWidget* widget, char* label) {
	gtk_button_set_label(GTK_BUTTON(widget), (gchar*)label);
}

static void _gtk_box_pack_start(GtkWidget* box, GtkWidget* child, int expand, int fill, int padding) {
	gtk_box_pack_start(GTK_BOX(box), child, expand, fill, padding);
}

static void _gtk_box_pack_end(GtkWidget* box, GtkWidget* child, int expand, int fill, int padding) {
	gtk_box_pack_end(GTK_BOX(box), child, expand, fill, padding);
}
*/
import "C";
import "time";
import "unsafe";
import "container/vector";

//-----------------------------------------------------------------------
// GtkWidget
//-----------------------------------------------------------------------
type GtkWidget struct {
	Widget *C.GtkWidget;
}
func (v *GtkWidget) Hide() { C.gtk_widget_hide(v.Widget) }
func (v *GtkWidget) HideAll() { C.gtk_widget_hide_all(v.Widget) }
func (v *GtkWidget) Show() { C.gtk_widget_show(v.Widget) }
func (v *GtkWidget) ShowAll() { C.gtk_widget_show_all(v.Widget) }
func (v *GtkWidget) ShowNow() { C.gtk_widget_show_now(v.Widget) }
func (v *GtkWidget) Destory() { C.gtk_widget_destroy(v.Widget) }
func (v *GtkWidget) Add(w *GtkWidget) { C._gtk_container_add(v.Widget, w.Widget) }
func (v *GtkWidget) Connect(s string, f func(widget *GtkWidget, data unsafe.Pointer), data unsafe.Pointer) {
	funcs.Push(&Callback{f});
	C._gtk_signal_connect(v.Widget, C.CString(s), C.int(funcs.Len())-1, unsafe.Pointer(data));
}

//-----------------------------------------------------------------------
// GtkWindow
//-----------------------------------------------------------------------
const (
	GTK_WINDOW_TOPLEVEL = 0
);
type GtkWindow GtkWidget;
func Window(t int) *GtkWidget {
	return &GtkWidget{ C.gtk_window_new(C.GtkWindowType(t)) };
}
func (v *GtkWindow) GetTitle() string { return C.GoString(C._gtk_window_get_title(v.Widget)); }
func (v *GtkWindow) SetTitle(title string) { C._gtk_window_set_title(v.Widget, C.CString(title)); }
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
// GtkBox
//-----------------------------------------------------------------------
type GtkBox GtkWidget;
func (v *GtkBox) PackStart(child *GtkWidget, expand int, fill int, padding int) {
	C._gtk_box_pack_start(v.Widget, child.Widget, C.int(expand), C.int(fill), C.int(padding));
}
func (v *GtkBox) PackEnd(child *GtkWidget, expand int, fill int, padding int) {
	C._gtk_box_pack_end(v.Widget, child.Widget, C.int(expand), C.int(fill), C.int(padding));
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
type GtkVBox GtkBox;
func VBox(homogeneous int, spacing int) *GtkWidget { return &GtkWidget{ C.gtk_vbox_new(C.gboolean(homogeneous), C.gint(spacing)) }; }

//-----------------------------------------------------------------------
// GtkHBox
//-----------------------------------------------------------------------
type GtkHBox GtkBox;
func HBox(homogeneous int, spacing int) *GtkWidget { return &GtkWidget{ C.gtk_hbox_new(C.gboolean(homogeneous), C.gint(spacing)) }; }

//-----------------------------------------------------------------------
// GtkEntry
//-----------------------------------------------------------------------
type GtkEntry GtkWidget;
func Entry() *GtkWidget { return &GtkWidget{ C.gtk_entry_new() }; }
func (v *GtkEntry) GetText() string { return C.GoString(C._gtk_entry_get_text(v.Widget)); }
func (v *GtkEntry) SetText(text string) { C._gtk_entry_set_text(v.Widget, C.CString(text)); }

//-----------------------------------------------------------------------
// GtkLabel
//-----------------------------------------------------------------------
type GtkLabel GtkWidget;
func Label(label string) *GtkWidget { return &GtkWidget{ C.gtk_label_new(C.to_gcharptr(C.CString(label))) }; }
func (v *GtkLabel) GetText() string { return C.GoString(C._gtk_label_get_text(v.Widget)); }
func (v *GtkLabel) SetText(label string) { C._gtk_label_set_text(v.Widget, C.CString(label)); }
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
// GtkButton
//-----------------------------------------------------------------------
type GtkButton GtkWidget;
func Button() *GtkWidget { return &GtkWidget{ C.gtk_button_new() }; }
func ButtonWithLabel(label string) *GtkWidget { return &GtkWidget{ C.gtk_button_new_with_label(C.to_gcharptr(C.CString(label))) }; }
func (v *GtkButton) GetLabel() string { return C.GoString(C._gtk_button_get_label(v.Widget)); }
func (v *GtkButton) SetLabel(label string) { C._gtk_button_set_label(v.Widget, C.CString(label)); }
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
type Callback struct {
	f func(w *GtkWidget, d unsafe.Pointer);
}
var funcs *vector.Vector;
func pollEvents() {
	for true {
		if (int(C._gtk_enter_callback) == 1) {
			elem := funcs.At(int(C._gtk_callback_func_no));
			f := elem.(*Callback);
			w := &GtkWidget{ C._gtk_callback_widget };
			d := unsafe.Pointer(C._gtk_callback_data);
			f.f(w, d);
			C._gtk_enter_callback = C.int(0);
		}
		time.Sleep(1);
	}
}

func Init(args *[]string) {
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
	go pollEvents();
	C.gtk_main();
}
