package gtk

/*
#include <gtk/gtk.h>
#include <unistd.h>
#include <stdlib.h>
#include <stdint.h>
#include <stdarg.h>
#include <string.h>

typedef struct {
	char name[256];
    int func_no;
	GtkWidget* widget;
    uintptr_t* data;
	uintptr_t** args;
	int args_no;
	int index;
    GSignalQuery query;
	int fire;
} callback_info;

callback_info* current_callback_info;
static uintptr_t* callback_info_get_arg(callback_info* cbi, int idx) {
	return cbi->args[idx];
}
static void callback_info_free_args(callback_info* cbi) {
	free(cbi->args);
}

#ifdef __x86_64__
static int _callback(void *data, ...);
asm(
".text\n"
"	.type _callback_amd64, @function\n"
"_callback_amd64:\n"
"	movq	$1, %rax\n"
"	jmp	_callback\n"
"	.size	_callback_amd64, . - _callback_amd64\n"
);
#else
static void _callback(void *data, ...) {
    va_list ap;
    callback_info *cbi = (callback_info*) data;
	int i;

	cbi->fire = 0;
	cbi->args = (uintptr_t**)malloc(sizeof(uintptr_t*)*cbi->args_no);
    va_start(ap, data);
    for (i = 0; i < cbi->args_no; i++) {
		cbi->args[i] = va_arg(ap, void*);
	}
    va_end(ap);
	current_callback_info = cbi;
}
#endif

static void _gtk_init(void* argc, void* argv) {
	gtk_init((int*)argc, (char***)argv);
}

static void _gtk_container_add(GtkWidget* container, GtkWidget* widget) {
	gtk_container_add(GTK_CONTAINER(container), widget);
}

static void free_callback_info(gpointer data, GClosure *closure) {
	g_slice_free(callback_info, data);
}

static long _gtk_signal_connect(GtkWidget* widget, char* name, int func_no, void* data) {
	static int index = 0;
    GSignalQuery query;
    callback_info* cbi;
    guint signal_id = g_signal_lookup(name, G_OBJECT_TYPE(widget));
    g_signal_query(signal_id, &query);
    cbi = g_slice_new(callback_info);
	strcpy(cbi->name, name);
	cbi->func_no = func_no;
	cbi->widget = widget;
	cbi->args_no = query.n_params;
	cbi->data = data;
	cbi->index = index;
	index++;
    return g_signal_connect_data(widget, name, GTK_SIGNAL_FUNC(_callback), cbi, free_callback_info, G_CONNECT_SWAPPED);
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

static gboolean _gtk_toggle_button_get_mode(GtkWidget* widget) {
	return gtk_toggle_button_get_mode(GTK_TOGGLE_BUTTON(widget));
}

static void _gtk_toggle_button_set_mode(GtkWidget* widget, gboolean draw_indicator) {
	gtk_toggle_button_set_mode(GTK_TOGGLE_BUTTON(widget), draw_indicator);
}

static gboolean _gtk_toggle_button_get_active(GtkWidget* widget) {
	return gtk_toggle_button_get_active(GTK_TOGGLE_BUTTON(widget));
}

static void _gtk_toggle_button_set_active(GtkWidget* widget, gboolean draw_indicator) {
	gtk_toggle_button_set_active(GTK_TOGGLE_BUTTON(widget), draw_indicator);
}

static gboolean _gtk_toggle_button_get_inconsistent(GtkWidget* widget) {
	return gtk_toggle_button_get_inconsistent(GTK_TOGGLE_BUTTON(widget));
}

static void _gtk_toggle_button_set_inconsistent(GtkWidget* widget, gboolean draw_indicator) {
	gtk_toggle_button_set_inconsistent(GTK_TOGGLE_BUTTON(widget), draw_indicator);
}

static gboolean _gtk_combo_box_entry_get_text_column(GtkWidget* widget) {
	return gtk_combo_box_entry_get_text_column(GTK_COMBO_BOX_ENTRY(widget));
}

static void _gtk_combo_box_entry_set_text_column(GtkWidget* widget, gint text_column) {
	gtk_combo_box_entry_set_text_column(GTK_COMBO_BOX_ENTRY(widget), text_column);
}

static char* _gtk_font_button_get_title(GtkWidget* widget) {
	return (char*)gtk_font_button_get_title(GTK_FONT_BUTTON(widget));
}

static void _gtk_font_button_set_title(GtkWidget* widget, char* title) {
	gtk_font_button_set_title(GTK_FONT_BUTTON(widget), (gchar*)title);
}

static gboolean _gtk_font_button_get_use_size(GtkWidget* widget) {
	return gtk_font_button_get_use_size(GTK_FONT_BUTTON(widget));
}

static void _gtk_font_button_set_use_size(GtkWidget* widget, gboolean use_size) {
	gtk_font_button_set_use_size(GTK_FONT_BUTTON(widget), use_size);
}

static char* _gtk_font_button_get_font_name(GtkWidget* widget) {
	return (char*)gtk_font_button_get_font_name(GTK_FONT_BUTTON(widget));
}

static void _gtk_font_button_set_font_name(GtkWidget* widget, char* fontname) {
	gtk_font_button_set_font_name(GTK_FONT_BUTTON(widget), (gchar*)fontname);
}

static gboolean _gtk_font_button_get_show_size(GtkWidget* widget) {
	return gtk_font_button_get_show_size(GTK_FONT_BUTTON(widget));
}

static void _gtk_font_button_set_show_size(GtkWidget* widget, gboolean show_size) {
	gtk_font_button_set_show_size(GTK_FONT_BUTTON(widget), show_size);
}

static void _gtk_box_pack_start(GtkWidget* box, GtkWidget* child, gboolean expand, gboolean fill, guint padding) {
	gtk_box_pack_start(GTK_BOX(box), child, expand, fill, padding);
}

static void _gtk_box_pack_end(GtkWidget* box, GtkWidget* child, gboolean expand, gboolean fill, guint padding) {
	gtk_box_pack_end(GTK_BOX(box), child, expand, fill, padding);
}
static gboolean _gtk_tree_model_get_iter(GtkTreeModel* tree_model, void* iter, void* path) {
	return gtk_tree_model_get_iter(tree_model, (GtkTreeIter*)iter, (GtkTreePath*)path);
}
static GtkTreePath* _gtk_tree_model_get_path(GtkTreeModel* tree_model, GtkTreeIter* iter) {
	return gtk_tree_model_get_path(tree_model, iter);
}
static gchar* to_gcharptr(char* s) { return (gchar*)s; }
static char* to_charptr(gchar* s) { return (char*)s; }
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
type WidgetLike interface {
	ToGtkWidget() *C.GtkWidget;
	Hide();
	HideAll();
	Show();
	ShowAll();
	ShowNow();
	Destroy();
	Connect(s string, f CallbackFunc, data unsafe.Pointer);
	GetTopLevel() *GtkWidget;
	GetTopLevelAsWindow() *GtkWindow;
	HideOnDelete();
	QueueResize();
}
type GtkWidget struct {
	Widget *C.GtkWidget;
}
func (v GtkWidget) ToGtkWidget() *C.GtkWidget { return v.Widget }
func (v GtkWidget) Hide() { C.gtk_widget_hide(v.Widget) }
func (v GtkWidget) HideAll() { C.gtk_widget_hide_all(v.Widget) }
func (v GtkWidget) Show() { C.gtk_widget_show(v.Widget) }
func (v GtkWidget) ShowAll() { C.gtk_widget_show_all(v.Widget) }
func (v GtkWidget) ShowNow() { C.gtk_widget_show_now(v.Widget) }
func (v GtkWidget) Destroy() { C.gtk_widget_destroy(v.Widget) }
func (v GtkWidget) Connect(s string, f CallbackFunc, data unsafe.Pointer) {
	funcs.Push(&Callback{f});
	C._gtk_signal_connect(v.Widget, C.CString(s),
		                    C.int(funcs.Len())-1, data);
}
func (v GtkWidget) GetTopLevel() *GtkWidget {
	return &GtkWidget {
		C.gtk_widget_get_toplevel(v.Widget)
	};
}
func (v GtkWidget) GetTopLevelAsWindow() *GtkWindow {
	return &GtkWindow { GtkContainer { GtkWidget {
		C.gtk_widget_get_toplevel(v.Widget)
	}}};
}
func (v GtkWidget) HideOnDelete() {
	C.gtk_widget_hide_on_delete(v.Widget);
}
func (v GtkWidget) QueueResize() {
	C.gtk_widget_queue_resize(v.Widget);
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
// GtkContainer
//-----------------------------------------------------------------------
type ContainerLike interface {
	WidgetLike;
	Add(w WidgetLike);
}
type GtkContainer struct { GtkWidget; }
func (v GtkContainer) Add(w WidgetLike) {
	C._gtk_container_add(v.ToGtkWidget(), w.ToGtkWidget());
}

//-----------------------------------------------------------------------
// GtkWindow
//-----------------------------------------------------------------------
const (
	GTK_WINDOW_TOPLEVEL = 0;
	GTK_WINDOW_POPUP = 1;
);
type WindowLike interface {
	ContainerLike;
	SetTransientFor(parent WindowLike);
	GetTitle() string;
	SetTitle(title string);
}
type GtkWindow struct { GtkContainer; }
func Window(t int) *GtkWindow {
	return &GtkWindow { GtkContainer { GtkWidget {
		C.gtk_window_new(C.GtkWindowType(t))
	}}};
}
func (v GtkWindow) GetTitle() string { return C.GoString(C._gtk_window_get_title(v.Widget)); }
func (v GtkWindow) SetTitle(title string) { C._gtk_window_set_title(v.Widget, C.CString(title)); }
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
type DialogLike interface {
	WidgetLike;
	Run() int;
	Response(CallbackFunc);
}
type GtkDialog struct {
	GtkWindow
};
func (v GtkDialog) Run() int {
	return int(C._gtk_dialog_run(v.Widget));
}
func (v GtkDialog) Response(response CallbackFunc, data unsafe.Pointer) {
	v.Connect("response", response, data);
}
// TODO

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
type GtkMessageDialog struct {
	GtkDialog
};
func MessageDialog(parent WindowLike, flag int, t int, button int,
                   message string) *GtkMessageDialog {
	return &GtkMessageDialog { GtkDialog { GtkWindow { GtkContainer { GtkWidget {
		C._gtk_message_dialog_new(
			parent.ToGtkWidget(),
			C.GtkDialogFlags(flag),
			C.GtkMessageType(t),
			C.GtkButtonsType(button),
			C.CString(message))
	}}}}};
}
// TODO
 
//-----------------------------------------------------------------------
// GtkBox
//-----------------------------------------------------------------------
type Packable interface {
	ContainerLike;
	PackStart(child WidgetLike, expand bool, fill bool, padding uint);
	PackEnd(child WidgetLike, expand bool, fill bool, padding uint);
}
type GtkBox struct {
	GtkContainer
};
func (v GtkBox) PackStart(child WidgetLike, expand bool, fill bool, padding uint) {
	C._gtk_box_pack_start(v.Widget, child.ToGtkWidget(), bool2gboolean(expand), bool2gboolean(fill), C.guint(padding));
}
func (v GtkBox) PackEnd(child WidgetLike, expand bool, fill bool, padding uint) {
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
func VBox(homogeneous bool, spacing uint) *GtkBox {
	return &GtkBox { GtkContainer { GtkWidget {
		C.gtk_vbox_new(bool2gboolean(homogeneous), C.gint(spacing))
	}}};
}
// TODO

//-----------------------------------------------------------------------
// GtkHBox
//-----------------------------------------------------------------------
type GtkHBox GtkBox;
func HBox(homogeneous bool, spacing uint) *GtkBox {
	return &GtkBox { GtkContainer { GtkWidget {
		C.gtk_hbox_new(bool2gboolean(homogeneous), C.gint(spacing))
	}}};
}
// TODO

//-----------------------------------------------------------------------
// GtkEntry
//-----------------------------------------------------------------------
type TextInputLike interface {
	WidgetLike;
	GetText() string;
	SetText(label string);
}
type GtkEntry struct {
	GtkWidget;
}
func Entry() *GtkEntry {
	return &GtkEntry { GtkWidget {
		C.gtk_entry_new()
	}};
}
func (v GtkEntry) GetText() string { return C.GoString(C._gtk_entry_get_text(v.Widget)); }
func (v GtkEntry) SetText(text string) { C._gtk_entry_set_text(v.Widget, C.CString(text)); }
// TODO

//-----------------------------------------------------------------------
// GtkImage
//-----------------------------------------------------------------------
const (
	GTK_ICON_SIZE_INVALID = 0;
	GTK_ICON_SIZE_MENU = 1;
	GTK_ICON_SIZE_SMALL_TOOLBAR = 2;
	GTK_ICON_SIZE_LARGE_TOOLBAR = 3;
	GTK_ICON_SIZE_BUTTON = 4;
	GTK_ICON_SIZE_DND = 5;
	GTK_ICON_SIZE_DIALOG = 6;
)
type ImageLike interface {
	WidgetLike;
}
type GtkImage struct {
	GtkWidget;
}
func Image() *GtkImage {
	return &GtkImage { GtkWidget {
		C.gtk_image_new()
	}};
}
func ImageFromFile(filename string) *GtkImage {
	return &GtkImage { GtkWidget {
		C.gtk_image_new_from_file(C.to_gcharptr(C.CString(filename)))
	}};
}
func ImageFromStock(stock_id string, size int) *GtkImage {
	return &GtkImage { GtkWidget {
		C.gtk_image_new_from_stock(C.to_gcharptr(C.CString(stock_id)), C.GtkIconSize(size))
	}};
}
// TODO

//-----------------------------------------------------------------------
// GtkLabel
//-----------------------------------------------------------------------
type LabelLike interface {
	WidgetLike;
	GetLabel() string;
	SetLabel(label string);
}
type GtkLabel struct {
	GtkWidget;
}
func Label(label string) *GtkLabel {
	return &GtkLabel { GtkWidget {
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
type AccelLabelLike interface {
	WidgetLike;
	GetAccelWidget() GtkWidget;
	SetAccelWidget(GtkWidget);
}
type GtkAccelLabel struct {
	GtkWidget;
}
func AccelLabel(label string) *GtkAccelLabel {
	return &GtkAccelLabel { GtkWidget {
		C.gtk_accel_label_new(C.to_gcharptr(C.CString(label)))
	}};
}
func (v GtkAccelLabel) GetAccelWidget() GtkWidget { return GtkWidget{ C._gtk_accel_label_get_accel_widget(v.Widget) }; }
func (v GtkAccelLabel) GetAccelWidth() uint { return uint(C._gtk_accel_label_get_accel_width(v.Widget)); }
func (v GtkAccelLabel) SetAccelWidget(w GtkWidget) { C._gtk_accel_label_set_accel_widget(v.Widget, w.Widget); }
func (v GtkAccelLabel) Refetch() bool { return gboolean2bool(C._gtk_accel_label_refetch(v.Widget)); }
// TODO
// gtk_accel_label_set_accel_closure

//-----------------------------------------------------------------------
// GtkButton
//-----------------------------------------------------------------------
type ButtonLike interface { // Buttons are LabelLike Widgets!
	LabelLike;
	// the following should be just Clickable; ...
	Clicked(CallbackFunc, unsafe.Pointer); // this is a very simple interface...
}
type Clickable interface {
	WidgetLike;
	Clicked(CallbackFunc, unsafe.Pointer); // this is a very simple interface...
}
func (v GtkButton) Clicked(onclick CallbackFunc, data unsafe.Pointer) {
	v.Connect("clicked", onclick, data);
}
type GtkButton struct {
	GtkWidget;
}
func Button() *GtkButton {
	return &GtkButton { GtkWidget {
		C.gtk_button_new()
	}};
}
func ButtonWithLabel(label string) *GtkButton {
	return &GtkButton { GtkWidget {
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
// GtkToggleButton
//-----------------------------------------------------------------------
type GtkToggleButton struct {
	GtkButton;
}
func ToggleButton() *GtkToggleButton {
	return &GtkToggleButton { GtkButton { GtkWidget {
		C.gtk_toggle_button_new()
	}}};
}
func ToggleButtonWithLabel(label string) *GtkToggleButton {
	return &GtkToggleButton { GtkButton { GtkWidget {
		C.gtk_toggle_button_new_with_label(C.to_gcharptr(C.CString(label)))
	}}};
}
func ToggleButtonWithMnemonic(label string) *GtkToggleButton {
	return &GtkToggleButton { GtkButton { GtkWidget {
		C.gtk_check_button_new_with_mnemonic(C.to_gcharptr(C.CString(label)))
	}}};
}
func (v GtkToggleButton) GetMode() bool {
	return gboolean2bool(C._gtk_toggle_button_get_mode(v.Widget));
}
func (v GtkToggleButton) SetMode(draw_indicator bool) {
	C._gtk_toggle_button_set_mode(v.Widget, bool2gboolean(draw_indicator));
}
func (v GtkToggleButton) GetActive() bool {
	return gboolean2bool(C._gtk_toggle_button_get_active(v.Widget));
}
func (v GtkToggleButton) SetActive(is_active bool) {
	C._gtk_toggle_button_set_active(v.Widget, bool2gboolean(is_active));
}
func (v GtkToggleButton) GetInconsistent() bool {
	return gboolean2bool(C._gtk_toggle_button_get_inconsistent(v.Widget));
}
func (v GtkToggleButton) SetInconsistent(setting bool) {
	C._gtk_toggle_button_set_inconsistent(v.Widget, bool2gboolean(setting));
}
// FINISH

//-----------------------------------------------------------------------
// GtkCheckButton
//-----------------------------------------------------------------------
type GtkCheckButton struct {
	GtkToggleButton;
}
func CheckButton() *GtkCheckButton {
	return &GtkCheckButton { GtkToggleButton { GtkButton { GtkWidget {
		C.gtk_check_button_new()
	}}}};
}
func CheckButtonWithLabel(label string) *GtkCheckButton {
	return &GtkCheckButton { GtkToggleButton { GtkButton { GtkWidget {
		C.gtk_check_button_new_with_label(C.to_gcharptr(C.CString(label)))
	}}}};
}
func CheckButtonWithMnemonic(label string) *GtkCheckButton {
	return &GtkCheckButton { GtkToggleButton { GtkButton { GtkWidget {
		C.gtk_check_button_new_with_mnemonic(C.to_gcharptr(C.CString(label)))
	}}}};
}
//func (v GtkCheckButton) GetProps() string { return C.GoString(C._gtk_font_button_get_title(v.Widget)); }
// TODO

//-----------------------------------------------------------------------
// GtkFontButton
//-----------------------------------------------------------------------
type GtkFontButton struct {
	GtkButton;
}
func FontButton() *GtkFontButton {
	return &GtkFontButton { GtkButton { GtkWidget {
		C.gtk_font_button_new()
	}}};
}
func FontButtonWithFont(fontname string) *GtkFontButton {
	return &GtkFontButton { GtkButton { GtkWidget {
		C.gtk_font_button_new_with_font(C.to_gcharptr(C.CString(fontname)))
	}}};
}
func (v GtkFontButton) GetTitle() string { return C.GoString(C._gtk_font_button_get_title(v.Widget)); }
func (v GtkFontButton) SetTitle(title string) { C._gtk_font_button_set_title(v.Widget, C.CString(title)); }
func (v GtkFontButton) GetUseSize() bool { return gboolean2bool(C._gtk_font_button_get_use_size(v.Widget)); }
func (v GtkFontButton) SetUseSize(use_size bool) { C._gtk_font_button_set_use_size(v.Widget, bool2gboolean(use_size)); }
func (v GtkFontButton) GetFontName() string { return C.GoString(C._gtk_font_button_get_font_name(v.Widget)); }
func (v GtkFontButton) SetFontName(fontname string) { C._gtk_font_button_set_font_name(v.Widget, C.CString(fontname)); }
func (v GtkFontButton) GetShowSize() bool { return gboolean2bool(C._gtk_font_button_get_show_size(v.Widget)); }
func (v GtkFontButton) SetShowSize(show_size bool) { C._gtk_font_button_set_show_size(v.Widget, bool2gboolean(show_size)); }
// FINISH

//-----------------------------------------------------------------------
// GtkComboBoxEntry
//-----------------------------------------------------------------------
type GtkComboBoxEntry struct {
	GtkWidget;
}
func ComboBoxEntry() *GtkComboBoxEntry {
	return &GtkComboBoxEntry { GtkWidget {
		C.gtk_combo_box_entry_new()
	}};
}
func ComboBoxEntryNewText() *GtkComboBoxEntry {
	return &GtkComboBoxEntry { GtkWidget {
		C.gtk_combo_box_entry_new_text()
	}};
}
func (v GtkComboBoxEntry) GetTextColumn() int {
	return int(C._gtk_combo_box_entry_get_text_column(v.Widget));
}
func (v GtkComboBoxEntry) SetTextColumn(text_column int) {
	C._gtk_combo_box_entry_set_text_column(v.Widget, C.gint(text_column));
}
// FINISH

//-----------------------------------------------------------------------
// GtkTreePath
//-----------------------------------------------------------------------
type GtkTreePath struct {
	TreePath *C.GtkTreePath;
}
func TreePath() *GtkTreePath {
	return &GtkTreePath {
		C.gtk_tree_path_new()
	};
}
func TreePathFromString(path string) *GtkTreePath {
	return &GtkTreePath {
		C.gtk_tree_path_new_from_string(C.to_gcharptr(C.CString(path)))
	};
}
func TreePathNewFirst() *GtkTreePath {
	return &GtkTreePath {
		C.gtk_tree_path_new_first()
	};
}
func (v GtkTreePath) String() string {
	return C.GoString(C.to_charptr(C.gtk_tree_path_to_string(v.TreePath)));
}
func (v GtkTreePath) AppendIndex(index int) {
	C.gtk_tree_path_append_index(v.TreePath, C.gint(index));
}
func (v GtkTreePath) PrependIndex(index int) {
	C.gtk_tree_path_prepend_index(v.TreePath, C.gint(index));
}
func (v GtkTreePath) GetDepth() int {
	return int(C.gtk_tree_path_get_depth(v.TreePath));
}
//func (v GtkTreePath) GetIndices() *int {
//	return *int(C.gtk_tree_path_get_indices(v.TreePath));
//}
func (v GtkTreePath) Free() {
	C.gtk_tree_path_free(v.TreePath);
}
func (v GtkTreePath) Copy() *GtkTreePath {
	return &GtkTreePath {
		C.gtk_tree_path_copy(v.TreePath)
	};
}
func (v GtkTreePath) Compare(w GtkTreePath) int {
	return int(C.gtk_tree_path_compare(v.TreePath, w.TreePath));
}
func (v GtkTreePath) Next() {
	C.gtk_tree_path_next(v.TreePath);
}
func (v GtkTreePath) Prev() bool {
	return gboolean2bool(C.gtk_tree_path_prev(v.TreePath));
}
func (v GtkTreePath) Up() bool {
	return gboolean2bool(C.gtk_tree_path_up(v.TreePath));
}
func (v GtkTreePath) Down() {
	C.gtk_tree_path_down(v.TreePath);
}
func (v GtkTreePath) IsAncestor(descendant GtkTreePath) bool {
	return gboolean2bool(C.gtk_tree_path_is_ancestor(v.TreePath, descendant.TreePath));
}
func (v GtkTreePath) IsDescendant(ancestor GtkTreePath) bool {
	return gboolean2bool(C.gtk_tree_path_is_descendant(v.TreePath, ancestor.TreePath));
}
// FINISH

//-----------------------------------------------------------------------
// GtkTreeIter
//-----------------------------------------------------------------------
type GtkTreeIter struct {
	TreeIter *C.GtkTreeIter;
}
func (v GtkTreeIter) Copy() *GtkTreeIter {
	return &GtkTreeIter { C.gtk_tree_iter_copy(v.TreeIter) };
}
func (v GtkTreeIter) Free() {
	C.gtk_tree_iter_free(v.TreeIter);
}

//-----------------------------------------------------------------------
// GtkTreeModel
//-----------------------------------------------------------------------
const (
	GTK_TREE_MODEL_ITERS_PERSIST = 1 << 0;
	GTK_TREE_MODEL_LIST_ONLY     = 1 << 1;
)
type GtkTreeModel struct {
	TreeModel *C.GtkTreeModel;
}
func (v GtkTreeModel) GetFlags() int {
	return int(C.gtk_tree_model_get_flags(v.TreeModel));
}
func (v GtkTreeModel) GetNColumns() int {
	return int(C.gtk_tree_model_get_n_columns(v.TreeModel));
}
func (v GtkTreeModel) GetIter(iter *GtkTreeIter, path *GtkTreePath) bool {
	var iter_ C.GtkTreeIter;
	var path_ C.GtkTreePath;
	ret := gboolean2bool(C._gtk_tree_model_get_iter(v.TreeModel, unsafe.Pointer(&iter_), unsafe.Pointer(&path_)));
	*iter = GtkTreeIter { &iter_ };
	*path = GtkTreePath { &path_ };
	return ret;
}
func (v GtkTreeModel) GetIterFromString(iter *GtkTreeIter, path_string string) bool {
	var iter_ C.GtkTreeIter;
	ret := gboolean2bool(C.gtk_tree_model_get_iter_from_string(v.TreeModel, iter.TreeIter, C.to_gcharptr(C.CString(path_string))));
	*iter = GtkTreeIter { &iter_ };
	return ret;
}
func (v GtkTreeModel) GetStringFromIter(iter *GtkTreeIter) string {
	return C.GoString(C.to_charptr(C.gtk_tree_model_get_string_from_iter(v.TreeModel, iter.TreeIter)));
}
func (v GtkTreeModel) GetIterFirst(iter *GtkTreeIter) bool {
	var iter_ C.GtkTreeIter;
	ret := gboolean2bool(C.gtk_tree_model_get_iter_first(v.TreeModel, iter.TreeIter));
	*iter = GtkTreeIter { &iter_ };
	return ret;
}
func (v GtkTreeModel) GetPath(iter *GtkTreeIter) *GtkTreePath {
	return &GtkTreePath { C._gtk_tree_model_get_path(v.TreeModel, iter.TreeIter) };
}
// TODO
// gtk_tree_model_get_value
// gtk_tree_model_iter_next
// gtk_tree_model_iter_children
// gtk_tree_model_iter_has_child
// gtk_tree_model_iter_n_children
// gtk_tree_model_iter_nth_child
// gtk_tree_model_iter_parent
// gtk_tree_model_ref_node
// gtk_tree_model_unref_node
// gtk_tree_model_get
// gtk_tree_model_get_valist
// gtk_tree_model_foreach

//-----------------------------------------------------------------------
// Events
//-----------------------------------------------------------------------
var use_gtk_main bool = false;

// the go-gtk Callback is simpler than the one in C, because we have
// full closures, so there is never a need to pass additional data via
// a void * pointer.  Where you might have wanted to do that, you can
// instead just use func () { ... using data } to pass the data in.
type CallbackFunc func(*GtkWidget, []unsafe.Pointer);
type Callback struct {
	f CallbackFunc;
}
var funcs *vector.Vector;
var main_loop bool = true;
func pollEvents() {
	for main_loop {
		if use_gtk_main == false {
			C.gtk_main_iteration_do(C.gboolean(1));
		}
		//index := C._gtk_callback_index;
		cbi := C.current_callback_info;
		if cbi != nil && cbi.fire == C.int(0) {
			elem := funcs.At(int(cbi.func_no));
			args := make([]unsafe.Pointer, cbi.args_no);
			for i := C.int(0); i < cbi.args_no; i++ {
				args[i] = unsafe.Pointer(C.callback_info_get_arg(cbi, C.int(i)));
			}
			f := elem.(*Callback);
			f.f(&GtkWidget{cbi.widget}, args);
			cbi.fire = C.int(1);
			C.callback_info_free_args(cbi);
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
