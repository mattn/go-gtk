/*
Go Bindings for Gtk+ 2. Support version 2.16 and later.
*/
package gtk

/*
#include <gtk/gtk.h>
#include <gdk/gdk.h>
#ifdef _WIN32
#include <windows.h>
#else
#include <gdk/gdkx.h>
#endif
#include <unistd.h>
#include <stdlib.h>
#include <stdint.h>
#include <stdarg.h>
#include <string.h>
#include <stdio.h>
#include <pthread.h>

static void _gtk_init(void* argc, void* argv) {
	gtk_init((int*)argc, (char***)argv);
}

static GtkClipboard* _gtk_clipboard_get_for_display(void* display, void* selection) {
	return gtk_clipboard_get_for_display((GdkDisplay*)display, (GdkAtom)selection);
}

static int _gtk_selection_data_get_length(void* selection_data) {
	return (int) gtk_selection_data_get_length((GtkSelectionData*)selection_data);
}

static void* _gtk_selection_data_get_data(void* selection_data) {
	return (char*) gtk_selection_data_get_data((GtkSelectionData*)selection_data);
}

static char* _gtk_selection_data_get_text(void* selection_data) {
	return (char*) gtk_selection_data_get_text((GtkSelectionData*)selection_data);
}

static void _gtk_drag_finish(void *context, gboolean success, gboolean del, guint32 time_) {
	gtk_drag_finish((GdkDragContext*) context, success, del, time_);
}

static GtkWidget* _gtk_message_dialog_new(GtkWidget* parent, GtkDialogFlags flags, GtkMessageType type, GtkButtonsType buttons, gchar *message) {
	return gtk_message_dialog_new(
			GTK_WINDOW(parent),
			flags,
			type,
			buttons,
			message, NULL);
}

static GtkWidget* _gtk_message_dialog_new_with_markup(GtkWidget* parent, GtkDialogFlags flags, GtkMessageType type, GtkButtonsType buttons, gchar *message) {
	return gtk_message_dialog_new_with_markup(
			GTK_WINDOW(parent),
			flags,
			type,
			buttons,
			message, NULL);
}

static GtkWidget* _gtk_file_chooser_dialog_new(const gchar* title,
		GtkWidget* parent, int file_chooser_action, int action, const gchar* button) {
	return gtk_file_chooser_dialog_new(
			title,
			GTK_WINDOW(parent),
			file_chooser_action,
			button,
			action,
			NULL);
}

static GtkWidget* _gtk_file_chooser_widget_new(int file_chooser_action) {
	return gtk_file_chooser_widget_new(file_chooser_action);
}

static GtkTreePath* _gtk_tree_model_get_path(GtkTreeModel* tree_model, GtkTreeIter* iter) {
	return gtk_tree_model_get_path(tree_model, iter);
}

static void _gtk_text_buffer_insert_with_tag(GtkTextBuffer* buffer, GtkTextIter* iter, const gchar* text, gint len, GtkTextTag* tag) {
	gtk_text_buffer_insert_with_tags(buffer, iter, text, len, tag, NULL);
}

//static void _gtk_text_buffer_insert_with_tags_by_name(GtkTextBuffer* buffer, GtkTextIter* iter, const gchar* text, gint len, const gchar* first_tag_name, ...);

static GtkTextTag* _gtk_text_buffer_create_tag(GtkTextBuffer* buffer, const gchar* tag_name) {
	return gtk_text_buffer_create_tag(buffer, tag_name, NULL);
}

static void _gtk_widget_hide_on_delete(GtkWidget* w) {
	g_signal_connect(GTK_WIDGET(w), "delete-event", G_CALLBACK(gtk_widget_hide_on_delete), NULL);
}

static void _gtk_text_iter_assign(GtkTextIter* one, GtkTextIter* two) {
	*one = *two;
}

static void _apply_property(void* obj, const gchar* prop, const gchar* val) {
	GParamSpec *pspec;
	GValue fromvalue = { 0, };
	GValue tovalue = { 0, };
	pspec = g_object_class_find_property(G_OBJECT_GET_CLASS(obj), prop);
	if (!pspec) return;
	g_value_init(&fromvalue, G_TYPE_STRING);
	g_value_set_string(&fromvalue, val);
	g_value_init(&tovalue, G_PARAM_SPEC_VALUE_TYPE(pspec));
	g_value_transform(&fromvalue, &tovalue);
	g_object_set_property((GObject *)obj, prop, &tovalue);
	g_value_unset(&fromvalue);
	g_value_unset(&tovalue);
}

static GtkTreeViewColumn* _gtk_tree_view_column_new_with_attribute(gchar* title, GtkCellRenderer* cell) {
	return gtk_tree_view_column_new_with_attributes(title, cell, NULL);
}

static GtkTreeViewColumn* _gtk_tree_view_column_new_with_attributes(gchar* title, GtkCellRenderer* cell, gchar* prop, gint column) {
	return gtk_tree_view_column_new_with_attributes(title, cell, prop, column, NULL);
}

static void _gtk_list_store_set_ptr(GtkListStore* list_store, GtkTreeIter* iter, gint column, void* data) {
	gtk_list_store_set(list_store, iter, column, data, -1);
}

static void _gtk_list_store_set_addr(GtkListStore* list_store, GtkTreeIter* iter, gint column, void* data) {
	gtk_list_store_set(list_store, iter, column, *(gpointer*)data, -1);
}

static void _gtk_tree_store_set_ptr(GtkTreeStore* tree_store, GtkTreeIter* iter, gint column, void* data) {
	gtk_tree_store_set(tree_store, iter, column, data, -1);
}

static void _gtk_tree_store_set_addr(GtkTreeStore* tree_store, GtkTreeIter* iter, gint column, void* data) {
	gtk_tree_store_set(tree_store, iter, column, *(gpointer*)data, -1);
}

//static void _gtk_range_get_value(GtkRange* range, gdouble* value) {
//	*value = gtk_range_get_value(range);
//}

typedef struct {
	GtkMenu *menu;
	gint x;
	gint y;
	gboolean push_in;
	gpointer data;
} _gtk_menu_position_func_info;

extern void _go_gtk_menu_position_func(_gtk_menu_position_func_info* gmpfi);
static void _c_gtk_menu_position_func(GtkMenu *menu, gint *x, gint *y, gboolean *push_in, gpointer user_data) {
	_gtk_menu_position_func_info gmpfi;
	gmpfi.menu = menu;
	gmpfi.x = *x;
	gmpfi.y = *y;
	gmpfi.push_in = *push_in;
	gmpfi.data = user_data;
	_go_gtk_menu_position_func(&gmpfi);
	*x = gmpfi.x;
	*y = gmpfi.y;
	*push_in = gmpfi.push_in;
#ifdef _WIN32
	RECT rect;
	SystemParametersInfo(SPI_GETWORKAREA, 0, &rect, 0);
	gint h = GTK_WIDGET(menu)->requisition.height;
	if (*y + h > rect.bottom) *y -= h;
#endif
}

static void _gtk_menu_popup(GtkWidget *menu, GtkWidget *parent_menu_shell, GtkWidget *parent_menu_item, void* data, guint button, guint32 activate_time) {
	gtk_menu_popup(GTK_MENU(menu), parent_menu_shell, parent_menu_item, _c_gtk_menu_position_func, (gpointer) data, button, activate_time);
}

static inline GType* make_gtypes(int count) {
	return g_new0(GType, count);
}

static inline void destroy_gtypes(GType* types) {
	g_free(types);
}

static inline void set_gtype(GType* types, int n, int type) {
	types[n] = (GType) type;
}

static inline gchar** make_strings(int count) {
	return (gchar**)malloc(sizeof(gchar*) * count);
}

static inline void destroy_strings(gchar** strings) {
	free(strings);
}

static inline gchar* get_string(gchar** strings, int n) {
	return strings[n];
}

static inline void set_string(gchar** strings, int n, gchar* str) {
	strings[n] = str;
}

static GSList* to_gslist(void* gs) {
	return (GSList*)gs;
}

static int _check_version(int major, int minor, int micro) {
	return GTK_CHECK_VERSION(major, minor, micro);
}

static void _gtk_tree_iter_assign(void* iter, void* to) {
	*(GtkTreeIter*)iter = *(GtkTreeIter*)to;
}

static GtkWidget* _gtk_dialog_get_vbox(GtkWidget* w) {
  return GTK_DIALOG(w)->vbox;
}

//////////////////////////////////////////////
// ############# Version Control #############
//////////////////////////////////////////////

#if GTK_CHECK_VERSION(2,14,0)
static gdouble _gtk_adjustment_get_lower(GtkAdjustment* adjustment) {
	return gtk_adjustment_get_lower(adjustment);
}
static void _gtk_adjustment_set_lower(GtkAdjustment* adjustment, gdouble lower) {
	gtk_adjustment_set_lower(adjustment, lower);
}
static gdouble _gtk_adjustment_get_upper(GtkAdjustment* adjustment) {
	return gtk_adjustment_get_upper(adjustment);
}
static void _gtk_adjustment_set_upper(GtkAdjustment* adjustment, gdouble upper) {
	gtk_adjustment_set_upper(adjustment, upper);
}
static gdouble _gtk_adjustment_get_step_increment(GtkAdjustment* adjustment) {
	return gtk_adjustment_get_step_increment(adjustment);
}
static void _gtk_adjustment_set_step_increment(GtkAdjustment* adjustment, gdouble step_increment) {
	gtk_adjustment_set_step_increment(adjustment, step_increment);
}
static gdouble _gtk_adjustment_get_page_increment(GtkAdjustment* adjustment) {
	return gtk_adjustment_get_page_increment(adjustment);
}
static void _gtk_adjustment_set_page_increment(GtkAdjustment* adjustment, gdouble page_increment) {
	gtk_adjustment_set_page_increment(adjustment, page_increment);
}
static gdouble _gtk_adjustment_get_page_size(GtkAdjustment* adjustment) {
	return gtk_adjustment_get_page_size(adjustment);
}
static void _gtk_adjustment_set_page_size(GtkAdjustment* adjustment, gdouble page_size) {
	gtk_adjustment_set_page_size(adjustment, page_size);
}
static void _gtk_adjustment_configure(GtkAdjustment* adjustment, gdouble value, gdouble lower, gdouble upper, gdouble step_increment, gdouble page_increment, gdouble page_size) {
	gtk_adjustment_configure(adjustment, value, lower, upper, step_increment, page_increment, page_size);
}
#else //GTK_CHECK_VERSION(2,14,0)
static gdouble _gtk_adjustment_get_lower(GtkAdjustment* adjustment) {
	return 0f;
}
static void _gtk_adjustment_set_lower(GtkAdjustment* adjustment, gdouble lower) {
}
static gdouble _gtk_adjustment_get_upper(GtkAdjustment* adjustment) {
	return 0f;
}
static void _gtk_adjustment_set_upper(GtkAdjustment* adjustment, gdouble upper) {
}
static gdouble _gtk_adjustment_get_step_increment(GtkAdjustment* adjustment) {
	return 0f;
}
static void _gtk_adjustment_set_step_increment(GtkAdjustment* adjustment, gdouble step_increment) {
}
static gdouble _gtk_adjustment_get_page_increment(GtkAdjustment* adjustment) {
	return 0f;
}
static void _gtk_adjustment_set_page_increment(GtkAdjustment* adjustment, gdouble page_increment) {
}
static gdouble _gtk_adjustment_get_page_size(GtkAdjustment* adjustment) {
	return 0f;
}
static void _gtk_adjustment_set_page_size(GtkAdjustment* adjustment, gdouble page_size) {
}
static void _gtk_adjustment_configure(GtkAdjustment* adjustment, gdouble value, gdouble lower, gdouble upper, gdouble step_increment, gdouble page_increment, gdouble page_size) {
}
#endif //GTK_CHECK_VERSION(2,14,0)

#if GTK_CHECK_VERSION(2,18,0)
static gboolean _gtk_cell_renderer_toggle_get_activatable(GtkCellRendererToggle *toggle) {
	return gtk_cell_renderer_toggle_get_activatable(toggle);
}
static void _gtk_cell_renderer_toggle_set_activatable(GtkCellRendererToggle *toggle, gboolean setting) {
	gtk_cell_renderer_toggle_set_activatable(toggle, setting);
}
static void _gtk_menu_set_reserve_toggle_size(GtkMenu *menu, gboolean reserve_toggle_size) {
	gtk_menu_set_reserve_toggle_size(menu, reserve_toggle_size);
}
static gboolean _gtk_menu_get_reserve_toggle_size(GtkMenu *menu) {
	return gtk_menu_get_reserve_toggle_size(menu);
}
static gboolean _gtk_range_get_flippable(GtkRange *range) {
	return gtk_range_get_flippable(range);
}
static void _gtk_range_set_flippable(GtkRange *range, gboolean flippable) {
	gtk_range_set_flippable(range, flippable);
}
static void _gtk_widget_get_allocation(GtkWidget *widget, GtkAllocation *allocation) {
	gtk_widget_get_allocation(widget, allocation);
}
static void _gtk_widget_set_allocation(GtkWidget *widget, const GtkAllocation *allocation) {
	gtk_widget_set_allocation(widget, allocation);
}
static gboolean _gtk_widget_get_app_paintable(GtkWidget *widget) {
	return gtk_widget_get_app_paintable(widget);
}
static gboolean _gtk_widget_get_can_default(GtkWidget *widget) {
	return gtk_widget_get_can_default(widget);
}
static void _gtk_widget_set_can_default(GtkWidget *widget, gboolean can_default) {
	gtk_widget_set_can_default(widget, can_default);
}
static gboolean _gtk_widget_get_can_focus(GtkWidget *widget) {
	return gtk_widget_get_can_focus(widget);
}
static void _gtk_widget_set_can_focus(GtkWidget *widget, gboolean can_focus) {
	gtk_widget_set_can_focus(widget, can_focus);
}
static gboolean _gtk_widget_get_double_buffered(GtkWidget *widget) {
	return gtk_widget_get_double_buffered(widget);
}
static gboolean _gtk_widget_get_has_window(GtkWidget *widget) {
	return gtk_widget_get_has_window(widget);
}
static void _gtk_widget_set_has_window(GtkWidget *widget, gboolean has_window) {
	gtk_widget_set_has_window(widget, has_window);
}
static gboolean _gtk_widget_get_sensitive(GtkWidget *widget) {
	return gtk_widget_get_sensitive(widget);
}
static gboolean _gtk_widget_is_sensitive(GtkWidget *widget) {
	return gtk_widget_is_sensitive(widget);
}
static GtkStateType _gtk_widget_get_state(GtkWidget *widget) {
	return gtk_widget_get_state(widget);
}
static gboolean _gtk_widget_get_visible(GtkWidget *widget) {
	return gtk_widget_get_visible(widget);
}
static void _gtk_widget_set_visible(GtkWidget *widget, gboolean visible) {
	gtk_widget_set_visible(widget, visible);
}
static gboolean _gtk_widget_has_default(GtkWidget *widget) {
	return gtk_widget_has_default(widget);
}
static gboolean _gtk_widget_has_focus(GtkWidget *widget) {
	return gtk_widget_has_focus(widget);
}
static gboolean _gtk_widget_has_grab(GtkWidget *widget) {
	return gtk_widget_has_grab(widget);
}
static gboolean _gtk_widget_is_drawable(GtkWidget *widget) {
	return gtk_widget_is_drawable(widget);
}
static gboolean _gtk_widget_is_toplevel(GtkWidget *widget) {
	return gtk_widget_is_toplevel(widget);
}
static void _gtk_widget_set_window(GtkWidget *widget, GdkWindow *window) {
	gtk_widget_set_window(widget, window);
}
static void _gtk_widget_set_receives_default(GtkWidget *widget, gboolean receives_default) {
	gtk_widget_set_receives_default(widget, receives_default);
}
static gboolean _gtk_widget_get_receives_default(GtkWidget *widget) {
	return gtk_widget_get_receives_default(widget);
}
static GtkWidget* _gtk_info_bar_new() {
	return gtk_info_bar_new();
}
static void _gtk_info_bar_add_action_widget(GtkInfoBar *info_bar, GtkWidget *child, gint response_id) {
	gtk_info_bar_add_action_widget(info_bar, child, response_id);
}
static GtkWidget* _gtk_info_bar_add_button(GtkInfoBar *info_bar, const gchar *button_text, gint response_id) {
	return gtk_info_bar_add_button(info_bar, button_text, response_id);
}
static void _gtk_info_bar_set_response_sensitive(GtkInfoBar *info_bar, gint response_id, gboolean setting) {
	gtk_info_bar_set_response_sensitive(info_bar, response_id, setting);
}
static void _gtk_info_bar_set_default_response(GtkInfoBar *info_bar, gint response_id) {
	gtk_info_bar_set_default_response(info_bar, response_id);
}
static void _gtk_info_bar_response(GtkInfoBar *info_bar, gint response_id) {
	gtk_info_bar_response(info_bar, response_id);
}
static void _gtk_info_bar_set_message_type(GtkInfoBar *info_bar, GtkMessageType message_type) {
	gtk_info_bar_set_message_type(info_bar, message_type);
}
static GtkMessageType _gtk_info_bar_get_message_type(GtkInfoBar *info_bar) {
	return gtk_info_bar_get_message_type(info_bar);
}
static GtkWidget* _gtk_info_bar_get_action_area(GtkInfoBar *info_bar) {
	return gtk_info_bar_get_action_area(info_bar);
}
static GtkWidget* _gtk_info_bar_get_content_area(GtkInfoBar *info_bar) {
	return gtk_info_bar_get_content_area(info_bar);
}
static GtkWidget* _gtk_entry_new_with_buffer(GtkEntryBuffer *buffer) {
	return gtk_entry_new_with_buffer(buffer);
}
static GtkEntryBuffer* _gtk_entry_get_buffer(GtkEntry *entry) {
	return gtk_entry_get_buffer(entry);
}
static void _gtk_entry_set_buffer(GtkEntry *entry,  GtkEntryBuffer *buffer) {
	gtk_entry_set_buffer(entry, buffer);
}
static GtkEntryBuffer* _gtk_entry_buffer_new(const gchar *initial_chars, gint n_initial_chars) {
	return gtk_entry_buffer_new(initial_chars, n_initial_chars);
}
static const gchar* _gtk_entry_buffer_get_text(GtkEntryBuffer *buffer) {
	return gtk_entry_buffer_get_text(buffer);
}
static void _gtk_entry_buffer_set_text(GtkEntryBuffer *buffer, const gchar *chars, gint n_chars) {
	gtk_entry_buffer_set_text(buffer, chars, n_chars);
}
//static gsize _gtk_entry_buffer_get_bytes(GtkEntryBuffer *buffer) {
//	return gtk_entry_buffer_get_bytes(buffer);
//}
static guint _gtk_entry_buffer_get_length(GtkEntryBuffer *buffer) {
	return gtk_entry_buffer_get_length(buffer);
}
static gint _gtk_entry_buffer_get_max_length(GtkEntryBuffer *buffer) {
	return gtk_entry_buffer_get_max_length(buffer);
}
static void _gtk_entry_buffer_set_max_length(GtkEntryBuffer *buffer, gint max_length) {
	gtk_entry_buffer_set_max_length(buffer, max_length);
}
static guint _gtk_entry_buffer_insert_text(GtkEntryBuffer *buffer, guint position, const gchar *chars, gint n_chars) {
	return gtk_entry_buffer_insert_text(buffer, position, chars, n_chars);
}
static guint _gtk_entry_buffer_delete_text(GtkEntryBuffer *buffer, guint position, gint n_chars) {
	return gtk_entry_buffer_delete_text(buffer, position, n_chars);
}
static void _gtk_status_icon_set_title(GtkStatusIcon *status_icon, const gchar *title) {
	gtk_status_icon_set_title(status_icon, title);
}
static const gchar* _gtk_status_icon_get_title(GtkStatusIcon *status_icon) {
	return gtk_status_icon_get_title(status_icon);
}
#else //!GTK_CHECK_VERSION(2,18,0)
static gboolean _gtk_cell_renderer_toggle_get_activatable(GtkCellRendererToggle *toggle) {
	return 0;
}
static void _gtk_cell_renderer_toggle_set_activatable(GtkCellRendererToggle *toggle, gboolean setting) {
}
static void _gtk_menu_set_reserve_toggle_size(GtkMenu *menu, gboolean reserve_toggle_size) {
}
static gboolean _gtk_menu_get_reserve_toggle_size(GtkMenu *menu) {
	return 0;
}
static gboolean _gtk_range_get_flippable(GtkRange *range) {
	return 0;
}
static void _gtk_range_set_flippable(GtkRange *range, gboolean flippable) {
}
static void _gtk_widget_get_allocation(GtkWidget *widget, GtkAllocation *allocation) {
}
static void _gtk_widget_set_allocation(GtkWidget *widget, const GtkAllocation *allocation) {
}
static gboolean _gtk_widget_get_app_paintable(GtkWidget *widget) {
	return 0;
}
static gboolean _gtk_widget_get_can_default(GtkWidget *widget) {
	return 0;
}
static void _gtk_widget_set_can_default(GtkWidget *widget, gboolean can_default) {
}
static gboolean _gtk_widget_get_can_focus(GtkWidget *widget) {
	return 0;
}
static void _gtk_widget_set_can_focus(GtkWidget *widget, gboolean can_focus) {
}
static gboolean _gtk_widget_get_double_buffered(GtkWidget *widget) {
	return 0;
}
static gboolean _gtk_widget_get_has_window(GtkWidget *widget) {
	return 0;
}
static void _gtk_widget_set_has_window(GtkWidget *widget, gboolean has_window) {
}
static gboolean _gtk_widget_get_sensitive(GtkWidget *widget) {
	return 0;
}
static gboolean _gtk_widget_is_sensitive(GtkWidget *widget) {
	return 0;
}
static GtkStateType _gtk_widget_get_state(GtkWidget *widget) {
	return 0;
}
static gboolean _gtk_widget_get_visible(GtkWidget *widget) {
	return 0;
}
static void _gtk_widget_set_visible(GtkWidget *widget, gboolean visible) {
}
static gboolean _gtk_widget_has_default(GtkWidget *widget) {
	return 0;
}
static gboolean _gtk_widget_has_focus(GtkWidget *widget) {
	return 0:
}
static gboolean _gtk_widget_has_grab(GtkWidget *widget) {
	return 0;
}
static gboolean _gtk_widget_is_drawable(GtkWidget *widget) {
	return 0:
}
static gboolean _gtk_widget_is_toplevel(GtkWidget *widget) {
	return 0;
}
static void _gtk_widget_set_window(GtkWidget *widget, GdkWindow *window) {
}
static void _gtk_widget_set_receives_default(GtkWidget *widget, gboolean receives_default) {
}
static gboolean _gtk_widget_get_receives_default(GtkWidget *widget) {
	return 0;
}
static GtkWidget* _gtk_info_bar_new() {
	return NULL;
}
static void _gtk_info_bar_add_action_widget(GtkInfoBar *info_bar, GtkWidget *child, gint response_id) {
}
static GtkWidget* _gtk_info_bar_add_button(GtkInfoBar *info_bar, const gchar *button_text, gint response_id) {
	return NULL;
}
static void _gtk_info_bar_set_response_sensitive(GtkInfoBar *info_bar, gint response_id, gboolean setting) {
}
static void _gtk_info_bar_set_default_response(GtkInfoBar *info_bar, gint response_id) {
}
static void _gtk_info_bar_response(GtkInfoBar *info_bar, gint response_id) {
}
static void _gtk_info_bar_set_message_type(GtkInfoBar *info_bar, GtkMessageType message_type) {
}
static GtkMessageType _gtk_info_bar_get_message_type(GtkInfoBar *info_bar) {
	return 0;
}
static GtkWidget* _gtk_info_bar_get_action_area(GtkInfoBar *info_bar) {
	return NULL;
}
static GtkWidget* _gtk_info_bar_get_content_area(GtkInfoBar *info_bar) {
	return NULL;
}
static GtkWidget* _gtk_entry_new_with_buffer(GtkEntryBuffer *buffer) {
	return NULL;
}
static GtkEntryBuffer* _gtk_entry_get_buffer(GtkEntry *entry) {
	return NULL;
}
static void _gtk_entry_set_buffer(GtkEntry *entry,  GtkEntryBuffer *buffer) {
}
static GtkEntryBuffer* _gtk_entry_buffer_new(const gchar *initial_chars, gint n_initial_chars) {
	return NULL;
}
static const gchar* _gtk_entry_buffer_get_text(GtkEntryBuffer *buffer) {
	return NULL;
}
static void _gtk_entry_buffer_set_text(GtkEntryBuffer *buffer, const gchar *chars, gint n_chars) {
}
//static gsize _gtk_entry_buffer_get_bytes(GtkEntryBuffer *buffer) {
//	return 0;
//}
static guint _gtk_entry_buffer_get_length(GtkEntryBuffer *buffer) {
	return 0;
}
static gint _gtk_entry_buffer_get_max_length(GtkEntryBuffer *buffer) {
	return 0;
}
static void _gtk_entry_buffer_set_max_length(GtkEntryBuffer *buffer, gint max_length) {
}
static guint _gtk_entry_buffer_insert_text(GtkEntryBuffer *buffer, guint position, const gchar *chars, gint n_chars) {
	return 0;
}
static guint _gtk_entry_buffer_delete_text(GtkEntryBuffer *buffer, guint position, gint n_chars) {
	return 0;
}
static void _gtk_status_icon_set_title(GtkStatusIcon *status_icon, const gchar *title) {
}
static const gchar* _gtk_status_icon_get_title(GtkStatusIcon *status_icon) {
	return NULL;
}
#endif //GTK_CHECK_VERSION(2,18,0)

#if GTK_CHECK_VERSION(2,20,0)
static GtkWidget* _gtk_dialog_get_widget_for_response(GtkDialog* dialog, gint id) {
	return gtk_dialog_get_widget_for_response(dialog, id);
}
static GdkWindow* _gtk_viewport_get_bin_window(GtkViewport *viewport) {
	return gtk_viewport_get_bin_window(viewport);
}
static void _gtk_status_icon_set_name(GtkStatusIcon *status_icon, const gchar *name) {
	gtk_status_icon_set_name(status_icon, name);
}
#else //GTK_CHECK_VERSION(2,20,0)
static GtkWidget* _gtk_dialog_get_widget_for_response(GtkDialog* dialog, gint id) {
	return NULL;
}
static GdkWindow* _gtk_viewport_get_bin_window(GtkViewport *viewport) {
	return NULL;
}
static void _gtk_status_icon_set_name(GtkStatusIcon *status_icon, const gchar *name) {
}
#endif //GTK_CHECK_VERSION(2,20,0)

#if GTK_CHECK_VERSION(2,22,0)
static GtkWidget* _gtk_accessible_get_widget(GtkAccessible *accessible) {
	return gtk_accessible_get_widget(accessible);
}
static void _gtk_accessible_set_widget(GtkAccessible *accessible, GtkWidget *widget) {
	gtk_accessible_set_widget(accessible, widget);
}
static GdkWindow* _gtk_viewport_get_view_window(GtkViewport *viewport) {
	return gtk_viewport_get_view_window(viewport);
}
#else //GTK_CHECK_VERSION(2,22,0)
static GtkWidget* _gtk_accessible_get_widget(GtkAccessible *accessible) {
	return NULL;
}
static void _gtk_accessible_set_widget(GtkAccessible *accessible, GtkWidget *widget) {
}
static GdkWindow* _gtk_viewport_get_view_window(GtkViewport *viewport) {
	return NULL;
}
#endif //GTK_CHECK_VERSION(2,22,0)

#if GTK_CHECK_VERSION(2,24,0)
static GtkWidget* _gtk_combo_box_new_with_entry(void) {
	return gtk_combo_box_new_with_entry();
}
static GtkWidget* _gtk_combo_box_new_with_model_and_entry(GtkTreeModel *model) {
	return gtk_combo_box_new_with_model_and_entry(model);
}
static GtkWidget* _gtk_combo_box_text_new(void) {
	return gtk_combo_box_text_new();
}
static GtkWidget* _gtk_combo_box_text_new_with_entry(void) {
	return gtk_combo_box_text_new_with_entry();
}
static void _gtk_combo_box_text_append_text(GtkComboBoxText *combo_box, const gchar *text) {
	gtk_combo_box_text_append_text(combo_box, text);
}
static void  _gtk_combo_box_text_insert_text(GtkComboBoxText *combo_box, gint position, const gchar *text) {
	gtk_combo_box_text_insert_text(combo_box, position, text);
}
static void _gtk_combo_box_text_prepend_text(GtkComboBoxText *combo_box, const gchar *text) {
	gtk_combo_box_text_prepend_text(combo_box, text);
}
static void _gtk_combo_box_text_remove(GtkComboBoxText *combo_box, gint position) {
	gtk_combo_box_text_remove(combo_box, position);
}
static gchar* _gtk_combo_box_text_get_active_text(GtkComboBoxText *combo_box) {
	return gtk_combo_box_text_get_active_text(combo_box);
}
static void _gtk_notebook_set_group_name(GtkNotebook* notebook, const gchar* group_name) {
	gtk_notebook_set_group_name(notebook, group_name);
}
static const gchar* _gtk_notebook_get_group_name(GtkNotebook* notebook) {
	return gtk_notebook_get_group_name(notebook);
}
#else //GTK_CHECK_VERSION(2,24,0)
static GtkWidget* _gtk_combo_box_new_with_entry(void) {
	return NULL;
}
static GtkWidget* _gtk_combo_box_new_with_model_and_entry(GtkTreeModel *model) {
	return NULL;
}
static GtkWidget* _gtk_combo_box_text_new(void) {
	return NULL;
}
static GtkWidget* _gtk_combo_box_text_new_with_entry(void) {
	return NULL;
}
static void _gtk_combo_box_text_append_text(GtkWidget *combo_box, const gchar *text) {
}
static void  _gtk_combo_box_text_insert_text(GtkWidget *combo_box, gint position, const gchar *text) {
}
static void _gtk_combo_box_text_prepend_text(GtkWidget *combo_box, const gchar *text) {
}
static void _gtk_combo_box_text_remove(GtkWidget *combo_box, gint position) {
}
static gchar* _gtk_combo_box_text_get_active_text(GtkWidget *combo_box) {
	return NULL;
}
static void _gtk_notebook_set_group_name(GtkNotebook* notebook, const gchar* group_name) {
}
static const gchar* _gtk_notebook_get_group_name(GtkNotebook* notebook) {
	return NULL;
}
#endif //GTK_CHECK_VERSION(2,24,0)

static GtkCellRenderer* _gtk_cell_renderer_spinner_new(void) {
#ifdef GTK_CELL_RENDERER_SPINNER //2.20
	return gtk_cell_renderer_spinner_new();
#else
	return gtk_cell_renderer_spin_new();
#endif
}

//////////////////////////////////////////////
// ################# Casting #################
//////////////////////////////////////////////

static inline GObject* to_GObject(void* o) { return G_OBJECT(o); }
static inline gchar* to_gcharptr(const char* s) { return (gchar*)s; }
static inline char* to_charptr(const gchar* s) { return (char*)s; }
static inline char* to_charptr_guchar(const guchar* s) { return (char*)s; }
static inline char* to_charptr_voidp(const void* s) { return (char*)s; }
static inline gchar** next_gcharptr(gchar** s) { return (s+1); }
static inline void free_string(char* s) { free(s); }

static GValue* to_GValueptr(void* s) { return (GValue*)s; }
static GtkWindow* to_GtkWindow(GtkWidget* w) { return GTK_WINDOW(w); }
static GtkDialog* to_GtkDialog(GtkWidget* w) { return GTK_DIALOG(w); }
static GtkAboutDialog* to_GtkAboutDialog(GtkWidget* w) { return GTK_ABOUT_DIALOG(w); }
static GtkContainer* to_GtkContainer(GtkWidget* w) { return GTK_CONTAINER(w); }
static GtkFileChooser* to_GtkFileChooser(GtkWidget* w) { return GTK_FILE_CHOOSER(w); }
static GtkFontSelectionDialog* to_GtkFontSelectionDialog(GtkWidget* w) { return GTK_FONT_SELECTION_DIALOG(w); }
static GtkLabel* to_GtkLabel(GtkWidget* w) { return GTK_LABEL(w); }
static GtkButton* to_GtkButton(GtkWidget* w) { return GTK_BUTTON(w); }
static GtkRadioButton* to_GtkRadioButton(GtkWidget* w) { return GTK_RADIO_BUTTON(w); }
static GtkFontButton* to_GtkFontButton(GtkWidget* w) { return GTK_FONT_BUTTON(w); }
static GtkLinkButton* to_GtkLinkButton(GtkWidget* w) { return GTK_LINK_BUTTON(w); }
static GtkComboBox* to_GtkComboBox(GtkWidget* w) { return GTK_COMBO_BOX(w); }
static GtkComboBoxEntry* to_GtkComboBoxEntry(GtkWidget* w) { return GTK_COMBO_BOX_ENTRY(w); }
static GtkMessageDialog* to_GtkMessageDialog(GtkWidget* w) { return GTK_MESSAGE_DIALOG(w); }

#if GTK_CHECK_VERSION(2,24,0)
static GtkComboBoxText* to_GtkComboBoxText(GtkWidget* w) { return GTK_COMBO_BOX_TEXT(w); }
#else
static GtkWidget* to_GtkComboBoxText(GtkWidget* w) { return w; }
#endif

static GtkAccessible* to_GtkAccessible(void* w) { return GTK_ACCESSIBLE(w); }
static GtkBin* to_GtkBin(GtkWidget* w) { return GTK_BIN(w); }
static GtkStatusbar* to_GtkStatusbar(GtkWidget* w) { return GTK_STATUSBAR(w); }
static GtkInfoBar* to_GtkInfoBar(GtkWidget* w) { return GTK_INFO_BAR(w); }
static GtkFrame* to_GtkFrame(GtkWidget* w) { return GTK_FRAME(w); }
static GtkBox* to_GtkBox(GtkWidget* w) { return GTK_BOX(w); }
static GtkPaned* to_GtkPaned(GtkWidget* w) { return GTK_PANED(w); }
static GtkToggleButton* to_GtkToggleButton(GtkWidget* w) { return GTK_TOGGLE_BUTTON(w); }
static GtkAccelLabel* to_GtkAccelLabel(GtkWidget* w) { return GTK_ACCEL_LABEL(w); }
static GtkEntry* to_GtkEntry(GtkWidget* w) { return GTK_ENTRY(w); }
static GtkAdjustment* to_GtkAdjustment(GtkObject* o) { return GTK_ADJUSTMENT(o); }
static GtkTextView* to_GtkTextView(GtkWidget* w) { return GTK_TEXT_VIEW(w); }
static GtkTextBuffer* to_GtkTextBuffer(void* w) { return GTK_TEXT_BUFFER(w); }
static GtkMenu* to_GtkMenu(GtkWidget* w) { return GTK_MENU(w); }
static GtkMenuBar* to_GtkMenuBar(GtkWidget* w) { return GTK_MENU_BAR(w); }
static GtkMenuShell* to_GtkMenuShell(GtkWidget* w) { return GTK_MENU_SHELL(w); }
static GtkMenuItem* to_GtkMenuItem(GtkWidget* w) { return GTK_MENU_ITEM(w); }
static GtkItem* to_GtkItem(GtkWidget* w) { return GTK_ITEM(w); }
static GtkScrolledWindow* to_GtkScrolledWindow(GtkWidget* w) { return GTK_SCROLLED_WINDOW(w); }
static GtkViewport* to_GtkViewport(GtkWidget* w) { return GTK_VIEWPORT(w); }
static GtkWidget* to_GtkWidget(void* w) { return GTK_WIDGET(w); }
static GdkWindow* to_GdkWindow(void* w) { return GDK_WINDOW(w); }
static GtkTreeView* to_GtkTreeView(GtkWidget* w) { return GTK_TREE_VIEW(w); }
static GtkIconView* to_GtkIconView(GtkWidget* w) { return GTK_ICON_VIEW(w); }
static GtkEditable* to_GtkEditable(GtkWidget* w) { return GTK_EDITABLE(w); }
static GtkCellRendererText* to_GtkCellRendererText(GtkCellRenderer* w) { return GTK_CELL_RENDERER_TEXT(w); }
static GtkCellRendererToggle* to_GtkCellRendererToggle(GtkCellRenderer* w) { return GTK_CELL_RENDERER_TOGGLE(w); }
static GtkScale* to_GtkScale(GtkWidget* w) { return GTK_SCALE(w); }
static GtkRange* to_GtkRange(GtkWidget* w) { return GTK_RANGE(w); }
static GtkTreeModel* to_GtkTreeModelFromListStore(GtkListStore* w) { return GTK_TREE_MODEL(w); }
static GtkTreeModel* to_GtkTreeModelFromTreeStore(GtkTreeStore* w) { return GTK_TREE_MODEL(w); }
static GtkImage* to_GtkImage(GtkWidget* w) { return GTK_IMAGE(w); }
static GtkNotebook* to_GtkNotebook(GtkWidget* w) { return GTK_NOTEBOOK(w); }
static GtkTable* to_GtkTable(GtkWidget* w) { return GTK_TABLE(w); }
static GtkDrawingArea* to_GtkDrawingArea(GtkWidget* w) { return GTK_DRAWING_AREA(w); }
static GtkAssistant* to_GtkAssistant(GtkWidget* w) { return GTK_ASSISTANT(w); }
static GtkExpander* to_GtkExpander(GtkWidget* w) { return GTK_EXPANDER(w); }
static GtkAlignment* to_GtkAlignment(GtkWidget* w) { return GTK_ALIGNMENT(w); }
static GtkProgressBar* to_GtkProgressBar(GtkWidget* w) { return GTK_PROGRESS_BAR(w); }
static GtkFixed* to_GtkFixed(GtkWidget* w) { return GTK_FIXED(w); }
static GtkCheckMenuItem* to_GtkCheckMenuItem(GtkWidget* w) { return GTK_CHECK_MENU_ITEM(w); }
static GtkRadioMenuItem* to_GtkRadioMenuItem(GtkWidget* w) { return GTK_RADIO_MENU_ITEM(w); }
static GtkFileFilter* to_GtkFileFilter(gpointer p) { return GTK_FILE_FILTER(p); }
*/
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
			case GtkResponseType:
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
type GtkAccelGroup struct {
	AccelGroup *C.GtkAccelGroup
}

func AccelGroup() *GtkAccelGroup {
	return &GtkAccelGroup{C.gtk_accel_group_new()}
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
type GtkClipboard struct {
	Clipboard *C.GtkClipboard
}

func ClipboardGetForDisplay(display *gdk.GdkDisplay, selection gdk.GdkAtom) *GtkClipboard {
	var cdisplay unsafe.Pointer
	if display != nil {
		cdisplay = display.Display
	}
	return &GtkClipboard{
		C._gtk_clipboard_get_for_display(cdisplay, unsafe.Pointer(selection))}
}

func (v *GtkClipboard) Clear() {
	C.gtk_clipboard_clear(v.Clipboard)
}

func (v *GtkClipboard) SetText(text string) {
	ptr := C.to_charptr_voidp(unsafe.Pointer(&([]byte(text))[0]))
	C.gtk_clipboard_set_text(v.Clipboard, C.to_gcharptr(ptr), C.gint(-1))
}

func (v *GtkClipboard) SetImage(pixbuf *gdkpixbuf.GdkPixbuf) {
	C.gtk_clipboard_set_image(v.Clipboard, pixbuf.Pixbuf)
}

func (v *GtkClipboard) Store() {
	C.gtk_clipboard_store(v.Clipboard)
}

func (v *GtkClipboard) WaitForText() string {
	return C.GoString(C.to_charptr(C.gtk_clipboard_wait_for_text(v.Clipboard)))
}

func (v *GtkClipboard) Connect(s string, f interface{}, datas ...interface{}) int {
	return glib.ObjectFromNative(unsafe.Pointer(C.to_GObject(unsafe.Pointer(v.Clipboard)))).Connect(s, f, datas...)
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
type GtkDestDefaults int

const (
	GTK_DEST_DEFAULT_MOTION    GtkDestDefaults = 1 << 0 /* respond to "drag_motion" */
	GTK_DEST_DEFAULT_HIGHLIGHT GtkDestDefaults = 1 << 1 /* auto-highlight */
	GTK_DEST_DEFAULT_DROP      GtkDestDefaults = 1 << 2 /* respond to "drag_drop" */
	GTK_DEST_DEFAULT_ALL       GtkDestDefaults = 0x07
)

type GtkTargetEntry struct {
	Target string
	Flags  uint
	Info   uint
}

func (v *GtkWidget) DragDestSet(flags GtkDestDefaults, targets []GtkTargetEntry, actions gdk.GdkDragAction) {
	ctargets := make([]C.GtkTargetEntry, len(targets))
	for i, target := range targets {
		ptr := C.CString(target.Target)
		defer C.free_string(ptr)
		ctargets[i].target = C.to_gcharptr(ptr)
		ctargets[i].flags = C.guint(target.Flags)
		ctargets[i].info = C.guint(target.Info)
	}
	C.gtk_drag_dest_set(v.Widget, C.GtkDestDefaults(flags), &ctargets[0], C.gint(len(targets)), C.GdkDragAction(actions))
}

func (v *GtkWidget) DragSourceSet(start_button_mask gdk.GdkModifierType, targets []GtkTargetEntry, actions gdk.GdkDragAction) {
	ctargets := make([]C.GtkTargetEntry, len(targets))
	for i, target := range targets {
		ptr := C.CString(target.Target)
		defer C.free_string(ptr)
		ctargets[i].target = C.to_gcharptr(ptr)
		ctargets[i].flags = C.guint(target.Flags)
		ctargets[i].info = C.guint(target.Info)
	}
	C.gtk_drag_source_set(v.Widget, C.GdkModifierType(start_button_mask), &ctargets[0], C.gint(len(targets)), C.GdkDragAction(actions))
}

func (v *GtkWidget) DragFinish(context *gdk.GdkDragContext, success bool, del bool, time uint) {
	C._gtk_drag_finish(unsafe.Pointer(context.DragContext), bool2gboolean(success), bool2gboolean(del), C.guint32(time))
}

func (v *GtkWidget) DragDestAddUriTargets() {
	C.gtk_drag_dest_add_uri_targets(v.Widget)
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
	GTK_STOCK_ABOUT                         = "gtk-about"
	GTK_STOCK_ADD                           = "gtk-add"
	GTK_STOCK_APPLY                         = "gtk-apply"
	GTK_STOCK_BOLD                          = "gtk-bold"
	GTK_STOCK_CANCEL                        = "gtk-cancel"
	GTK_STOCK_CAPS_LOCK_WARNING             = "gtk-caps-lock-warning"
	GTK_STOCK_CDROM                         = "gtk-cdrom"
	GTK_STOCK_CLEAR                         = "gtk-clear"
	GTK_STOCK_CLOSE                         = "gtk-close"
	GTK_STOCK_COLOR_PICKER                  = "gtk-color-picker"
	GTK_STOCK_CONVERT                       = "gtk-convert"
	GTK_STOCK_CONNECT                       = "gtk-connect"
	GTK_STOCK_COPY                          = "gtk-copy"
	GTK_STOCK_CUT                           = "gtk-cut"
	GTK_STOCK_DELETE                        = "gtk-delete"
	GTK_STOCK_DIALOG_AUTHENTICATION         = "gtk-dialog-authentication"
	GTK_STOCK_DIALOG_INFO                   = "gtk-dialog-info"
	GTK_STOCK_DIALOG_WARNING                = "gtk-dialog-warning"
	GTK_STOCK_DIALOG_ERROR                  = "gtk-dialog-error"
	GTK_STOCK_DIALOG_QUESTION               = "gtk-dialog-question"
	GTK_STOCK_DIRECTORY                     = "gtk-directory"
	GTK_STOCK_DISCARD                       = "gtk-discard"
	GTK_STOCK_DISCONNECT                    = "gtk-disconnect"
	GTK_STOCK_DND                           = "gtk-dnd"
	GTK_STOCK_DND_MULTIPLE                  = "gtk-dnd-multiple"
	GTK_STOCK_EDIT                          = "gtk-edit"
	GTK_STOCK_EXECUTE                       = "gtk-execute"
	GTK_STOCK_FILE                          = "gtk-file"
	GTK_STOCK_FIND                          = "gtk-find"
	GTK_STOCK_FIND_AND_REPLACE              = "gtk-find-and-replace"
	GTK_STOCK_FLOPPY                        = "gtk-floppy"
	GTK_STOCK_FULLSCREEN                    = "gtk-fullscreen"
	GTK_STOCK_GOTO_BOTTOM                   = "gtk-goto-bottom"
	GTK_STOCK_GOTO_FIRST                    = "gtk-goto-first"
	GTK_STOCK_GOTO_LAST                     = "gtk-goto-last"
	GTK_STOCK_GOTO_TOP                      = "gtk-goto-top"
	GTK_STOCK_GO_BACK                       = "gtk-go-back"
	GTK_STOCK_GO_DOWN                       = "gtk-go-down"
	GTK_STOCK_GO_FORWARD                    = "gtk-go-forward"
	GTK_STOCK_GO_UP                         = "gtk-go-up"
	GTK_STOCK_HARDDISK                      = "gtk-harddisk"
	GTK_STOCK_HELP                          = "gtk-help"
	GTK_STOCK_HOME                          = "gtk-home"
	GTK_STOCK_INDEX                         = "gtk-index"
	GTK_STOCK_INDENT                        = "gtk-indent"
	GTK_STOCK_INFO                          = "gtk-info"
	GTK_STOCK_UNINDENT                      = "gtk-unindent"
	GTK_STOCK_ITALIC                        = "gtk-italic"
	GTK_STOCK_JUMP_TO                       = "gtk-jump-to"
	GTK_STOCK_JUSTIFY_CENTER                = "gtk-justify-center"
	GTK_STOCK_JUSTIFY_FILL                  = "gtk-justify-fill"
	GTK_STOCK_JUSTIFY_LEFT                  = "gtk-justify-left"
	GTK_STOCK_JUSTIFY_RIGHT                 = "gtk-justify-right"
	GTK_STOCK_LEAVE_FULLSCREEN              = "gtk-leave-fullscreen"
	GTK_STOCK_MISSING_IMAGE                 = "gtk-missing-image"
	GTK_STOCK_MEDIA_FORWARD                 = "gtk-media-forward"
	GTK_STOCK_MEDIA_NEXT                    = "gtk-media-next"
	GTK_STOCK_MEDIA_PAUSE                   = "gtk-media-pause"
	GTK_STOCK_MEDIA_PLAY                    = "gtk-media-play"
	GTK_STOCK_MEDIA_PREVIOUS                = "gtk-media-previous"
	GTK_STOCK_MEDIA_RECORD                  = "gtk-media-record"
	GTK_STOCK_MEDIA_REWIND                  = "gtk-media-rewind"
	GTK_STOCK_MEDIA_STOP                    = "gtk-media-stop"
	GTK_STOCK_NETWORK                       = "gtk-network"
	GTK_STOCK_NEW                           = "gtk-new"
	GTK_STOCK_NO                            = "gtk-no"
	GTK_STOCK_OK                            = "gtk-ok"
	GTK_STOCK_OPEN                          = "gtk-open"
	GTK_STOCK_ORIENTATION_PORTRAIT          = "gtk-orientation-portrait"
	GTK_STOCK_ORIENTATION_LANDSCAPE         = "gtk-orientation-landscape"
	GTK_STOCK_ORIENTATION_REVERSE_LANDSCAPE = "gtk-orientation-reverse-landscape"
	GTK_STOCK_ORIENTATION_REVERSE_PORTRAIT  = "gtk-orientation-reverse-portrait"
	GTK_STOCK_PAGE_SETUP                    = "gtk-page-setup"
	GTK_STOCK_PASTE                         = "gtk-paste"
	GTK_STOCK_PREFERENCES                   = "gtk-preferences"
	GTK_STOCK_PRINT                         = "gtk-print"
	GTK_STOCK_PRINT_ERROR                   = "gtk-print-error"
	GTK_STOCK_PRINT_PAUSED                  = "gtk-print-paused"
	GTK_STOCK_PRINT_PREVIEW                 = "gtk-print-preview"
	GTK_STOCK_PRINT_REPORT                  = "gtk-print-report"
	GTK_STOCK_PRINT_WARNING                 = "gtk-print-warning"
	GTK_STOCK_PROPERTIES                    = "gtk-properties"
	GTK_STOCK_QUIT                          = "gtk-quit"
	GTK_STOCK_REDO                          = "gtk-redo"
	GTK_STOCK_REFRESH                       = "gtk-refresh"
	GTK_STOCK_REMOVE                        = "gtk-remove"
	GTK_STOCK_REVERT_TO_SAVED               = "gtk-revert-to-saved"
	GTK_STOCK_SAVE                          = "gtk-save"
	GTK_STOCK_SAVE_AS                       = "gtk-save-as"
	GTK_STOCK_SELECT_ALL                    = "gtk-select-all"
	GTK_STOCK_SELECT_COLOR                  = "gtk-select-color"
	GTK_STOCK_SELECT_FONT                   = "gtk-select-font"
	GTK_STOCK_SORT_ASCENDING                = "gtk-sort-ascending"
	GTK_STOCK_SORT_DESCENDING               = "gtk-sort-descending"
	GTK_STOCK_SPELL_CHECK                   = "gtk-spell-check"
	GTK_STOCK_STOP                          = "gtk-stop"
	GTK_STOCK_STRIKETHROUGH                 = "gtk-strikethrough"
	GTK_STOCK_UNDELETE                      = "gtk-undelete"
	GTK_STOCK_UNDERLINE                     = "gtk-underline"
	GTK_STOCK_UNDO                          = "gtk-undo"
	GTK_STOCK_YES                           = "gtk-yes"
	GTK_STOCK_ZOOM_100                      = "gtk-zoom-100"
	GTK_STOCK_ZOOM_FIT                      = "gtk-zoom-fit"
	GTK_STOCK_ZOOM_IN                       = "gtk-zoom-in"
	GTK_STOCK_ZOOM_OUT                      = "gtk-zoom-out"
)

type GtkStockItem struct {
	StockItem *C.GtkStockItem
}

func (v *GtkStockItem) Add(nitems uint) {
	C.gtk_stock_add(v.StockItem, C.guint(nitems))
}
func (v *GtkStockItem) AddStatic(nitems uint) {
	C.gtk_stock_add_static(v.StockItem, C.guint(nitems))
}
func GtkStockLookup(stock_id string, item *GtkStockItem) bool {
	ptr := C.CString(stock_id)
	defer C.free_string(ptr)
	return gboolean2bool(C.gtk_stock_lookup(C.to_gcharptr(ptr), item.StockItem))
}

// gtk_stock_item_copy
// gtk_stock_item_free

func GtkStockListIDs() *glib.SList {
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
type GtkSettings struct {
	Settings *C.GtkSettings
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

func (s *GtkSettings) SetStringProperty(name string, v_string string, origin string) {
	ptrn := C.CString(name)
	defer C.free_string(ptrn)
	ptrv := C.CString(v_string)
	defer C.free_string(ptrv)
	prts := C.CString(origin)
	defer C.free_string(prts)
	C.gtk_settings_set_string_property(s.Settings, C.to_gcharptr(ptrn), C.to_gcharptr(ptrv), C.to_gcharptr(prts))
}
func (s *GtkSettings) SetLongProperty(name string, v_long int32, origin string) {
	ptrn := C.CString(name)
	defer C.free_string(ptrn)
	prts := C.CString(origin)
	defer C.free_string(prts)
	C.gtk_settings_set_long_property(s.Settings, C.to_gcharptr(ptrn), C.glong(v_long), C.to_gcharptr(prts))
}
func (s *GtkSettings) SetDoubleProperty(name string, v_double float64, origin string) {
	ptrn := C.CString(name)
	defer C.free_string(ptrn)
	prts := C.CString(origin)
	defer C.free_string(prts)
	C.gtk_settings_set_double_property(s.Settings, C.to_gcharptr(ptrn), C.gdouble(v_double), C.to_gcharptr(prts))
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

type GtkSelectionData struct {
	SelectionData unsafe.Pointer
}

func SelectionDataFromNative(l unsafe.Pointer) *GtkSelectionData {
	return &GtkSelectionData{l}
}

func (v *GtkSelectionData) GetLength() int {
	return int(C._gtk_selection_data_get_length(v.SelectionData))
}

func (v *GtkSelectionData) GetData() unsafe.Pointer {
	return unsafe.Pointer(C._gtk_selection_data_get_data(v.SelectionData))
}

func (v *GtkSelectionData) GetText() string {
	return C.GoString(C._gtk_selection_data_get_text(v.SelectionData))
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
type GtkDialogFlags int

const (
	GTK_DIALOG_MODAL               GtkDialogFlags = 1 << 0 /* call gtk_window_set_modal (win, TRUE) */
	GTK_DIALOG_DESTROY_WITH_PARENT GtkDialogFlags = 1 << 1 /* call gtk_window_set_destroy_with_parent () */
	GTK_DIALOG_NO_SEPARATOR        GtkDialogFlags = 1 << 2 /* no separator bar above buttons */
)

type GtkResponseType int

const (
	GTK_RESPONSE_NONE         GtkResponseType = -1
	GTK_RESPONSE_REJECT       GtkResponseType = -2
	GTK_RESPONSE_ACCEPT       GtkResponseType = -3
	GTK_RESPONSE_DELETE_EVENT GtkResponseType = -4
	GTK_RESPONSE_OK           GtkResponseType = -5
	GTK_RESPONSE_CANCEL       GtkResponseType = -6
	GTK_RESPONSE_CLOSE        GtkResponseType = -7
	GTK_RESPONSE_YES          GtkResponseType = -8
	GTK_RESPONSE_NO           GtkResponseType = -9
	GTK_RESPONSE_APPLY        GtkResponseType = -10
	GTK_RESPONSE_HELP         GtkResponseType = -11
)

/*type DialogLike interface {
	WidgetLike
	Run() int
	Response(interface{}, ...interface{})
}*/

type GtkDialog struct {
	GtkWindow
}

func Dialog() *GtkDialog {
	return &GtkDialog{GtkWindow{GtkBin{GtkContainer{GtkWidget{C.gtk_dialog_new()}}}}}
}

// gtk_dialog_new_with_buttons

func (v *GtkDialog) Run() GtkResponseType {
	return GtkResponseType(C.gtk_dialog_run(C.to_GtkDialog(v.Widget)))
}
func (v *GtkDialog) Response(response interface{}, datas ...interface{}) {
	v.Connect("response", response, datas...)
}
func (v *GtkDialog) AddButton(button_text string, response_id int) *GtkButton {
	ptr := C.CString(button_text)
	defer C.free_string(ptr)
	return &GtkButton{GtkBin{GtkContainer{GtkWidget{
		C.gtk_dialog_add_button(C.to_GtkDialog(v.Widget), C.to_gcharptr(ptr), C.gint(response_id))}}}}
}

// gtk_dialog_add_buttons
// gtk_dialog_add_action_widget
// gtk_dialog_get_has_separator //deprecated since 2.22

//Deprecated since 2.22.
func (v *GtkDialog) SetHasSeparator(f bool) {
	deprecated_since(2, 22, 0, "gtk_dialog_set_has_separator()")
	C.gtk_dialog_set_has_separator(C.to_GtkDialog(v.Widget), bool2gboolean(f))
}
func (v *GtkDialog) SetDefaultResponse(id int) {
	C.gtk_dialog_set_default_response(C.to_GtkDialog(v.Widget), C.gint(id))
}

// gtk_dialog_set_has_separator //deprecated since 2.22
// gtk_dialog_set_response_sensitive

func (v *GtkDialog) GetResponseForWidget(w *GtkWidget) int {
	return int(C.gtk_dialog_get_response_for_widget(C.to_GtkDialog(v.Widget), w.Widget))
}
func (v *GtkDialog) GetWidgetForResponse(id int) *GtkWidget {
	panic_if_version_older(2, 20, 0, "gtk_dialog_get_widget_for_response()")
	return &GtkWidget{C._gtk_dialog_get_widget_for_response(C.to_GtkDialog(v.Widget), C.gint(id))}
}

// gtk_dialog_get_action_area
// gtk_dialog_get_content_area
// gtk_alternative_dialog_button_order
// gtk_dialog_set_alternative_button_order
// gtk_dialog_set_alternative_button_order_from_array

func (v *GtkDialog) GetVBox() *GtkVBox {
	return &GtkVBox{GtkBox{GtkContainer{GtkWidget{C._gtk_dialog_get_vbox(v.Widget)}}}}
}

//-----------------------------------------------------------------------
// GtkMessageDialog
//-----------------------------------------------------------------------
type GtkMessageType int

const (
	GTK_MESSAGE_INFO     = 0
	GTK_MESSAGE_WARNING  = 1
	GTK_MESSAGE_QUESTION = 2
	GTK_MESSAGE_ERROR    = 3
	GTK_MESSAGE_OTHER    = 4
)

type GtkButtonsType int

const (
	GTK_BUTTONS_NONE      = 0
	GTK_BUTTONS_OK        = 1
	GTK_BUTTONS_CLOSE     = 2
	GTK_BUTTONS_CANCEL    = 3
	GTK_BUTTONS_YES_NO    = 4
	GTK_BUTTONS_OK_CANCEL = 5
)

type GtkMessageDialog struct {
	GtkDialog
}

// TODO should be variadic function
func MessageDialog(parent *GtkWindow, flag GtkDialogFlags, t GtkMessageType, buttons GtkButtonsType, format string, args ...interface{}) *GtkMessageDialog {
	ptr := C.CString(strings.Replace(fmt.Sprintf(format, args...), "%", "%%", -1))
	defer C.free_string(ptr)
	return &GtkMessageDialog{GtkDialog{GtkWindow{GtkBin{GtkContainer{GtkWidget{
		C._gtk_message_dialog_new(
			parent.ToNative(),
			C.GtkDialogFlags(flag),
			C.GtkMessageType(t),
			C.GtkButtonsType(buttons),
			C.to_gcharptr(ptr))}}}}}}
}

func MessageDialogWithMarkup(parent *GtkWindow, flag GtkDialogFlags, t GtkMessageType, buttons GtkButtonsType, format string, args ...interface{}) *GtkMessageDialog {
	r := &GtkMessageDialog{GtkDialog{GtkWindow{GtkBin{GtkContainer{GtkWidget{
		C._gtk_message_dialog_new_with_markup(
			parent.ToNative(),
			C.GtkDialogFlags(flag),
			C.GtkMessageType(t),
			C.GtkButtonsType(buttons),
			nil)}}}}}}
	r.SetMarkup(fmt.Sprintf(format, args...))
	return r
}

func (v *GtkMessageDialog) SetMarkup(markup string) {
	ptr := C.CString(markup)
	defer C.free_string(ptr)
	C.gtk_message_dialog_set_markup(C.to_GtkMessageDialog(v.Widget), C.to_gcharptr(ptr))
}
func (v *GtkMessageDialog) SetImage(image WidgetLike) {
	C.gtk_message_dialog_set_image(C.to_GtkMessageDialog(v.Widget), image.ToNative())
}
func (v *GtkMessageDialog) GetImage() *GtkImage {
	return &GtkImage{GtkWidget{C.gtk_message_dialog_get_image(C.to_GtkMessageDialog(v.Widget))}}
}

// gtk_message_dialog_get_message_area //since 2.22
// gtk_message_dialog_format_secondary_text
// gtk_message_dialog_format_secondary_markup

//-----------------------------------------------------------------------
// GtkWindow
//-----------------------------------------------------------------------
type GtkWindowType int

const (
	GTK_WINDOW_TOPLEVEL GtkWindowType = 0
	GTK_WINDOW_POPUP    GtkWindowType = 1
)

type GtkWindowPosition int

const (
	GTK_WIN_POS_NONE             GtkWindowPosition = 0
	GTK_WIN_POS_CENTER           GtkWindowPosition = 1
	GTK_WIN_POS_MOUSE            GtkWindowPosition = 2
	GTK_WIN_POS_CENTER_ALWAYS    GtkWindowPosition = 3
	GTK_WIN_POS_CENTER_ON_PARENT GtkWindowPosition = 4
)

/*type WindowLike interface {
	ContainerLike
	SetTransientFor(parent WindowLike)
	GetTitle() string
	SetTitle(title string)
}*/

type GtkWindow struct {
	GtkBin
}

func Window(t GtkWindowType) *GtkWindow {
	return &GtkWindow{GtkBin{GtkContainer{GtkWidget{
		C.gtk_window_new(C.GtkWindowType(t))}}}}
}
func (v *GtkWindow) SetTitle(title string) {
	ptr := C.CString(title)
	defer C.free_string(ptr)
	C.gtk_window_set_title(C.to_GtkWindow(v.Widget), C.to_gcharptr(ptr))
}

// gtk_window_set_wmclass

func (v *GtkWindow) SetResizable(resizable bool) {
	C.gtk_window_set_resizable(C.to_GtkWindow(v.Widget), bool2gboolean(resizable))
}
func (v *GtkWindow) GetResizable() bool {
	return gboolean2bool(C.gtk_window_get_resizable(C.to_GtkWindow(v.Widget)))
}
func (v *GtkWindow) AddAccelGroup(group *GtkAccelGroup) {
	C.gtk_window_add_accel_group(C.to_GtkWindow(v.Widget), group.AccelGroup)
}

// gtk_window_remove_accel_group
// gtk_window_activate_focus
// gtk_window_activate_default

func (v *GtkWindow) SetModal(modal bool) {
	C.gtk_window_set_modal(C.to_GtkWindow(v.Widget), bool2gboolean(modal))
}
func (v *GtkWindow) SetDefaultSize(width int, height int) {
	C.gtk_window_set_default_size(C.to_GtkWindow(v.Widget), C.gint(width), C.gint(height))
}

// gtk_window_set_geometry_hints
// gtk_window_set_gravity
// gtk_window_get_gravity

func (v *GtkWindow) SetPosition(position GtkWindowPosition) {
	C.gtk_window_set_position(C.to_GtkWindow(v.Widget), C.GtkWindowPosition(position))
}
func (v *GtkWindow) SetTransientFor(parent *GtkWindow) {
	C.gtk_window_set_transient_for(C.to_GtkWindow(v.Widget), C.to_GtkWindow(parent.ToNative()))
}
func (v *GtkWindow) SetDestroyWithParent(setting bool) {
	C.gtk_window_set_destroy_with_parent(C.to_GtkWindow(v.Widget), bool2gboolean(setting))
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

func (v *GtkWindow) SetDefault(w *GtkWidget) {
	C.gtk_window_set_default(C.to_GtkWindow(v.Widget), w.Widget)
}
func (v *GtkWindow) Present() {
	C.gtk_window_present(C.to_GtkWindow(v.Widget))
}

// gtk_window_present_with_time
// gtk_window_iconify
// gtk_window_deiconify
// gtk_window_stick
// gtk_window_unstick

func (v *GtkWindow) Maximize() {
	C.gtk_window_maximize(C.to_GtkWindow(v.Widget))
}
func (v *GtkWindow) Unmaximize() {
	C.gtk_window_unmaximize(C.to_GtkWindow(v.Widget))
}

// gtk_window_fullscreen
// gtk_window_unfullscreen
// gtk_window_set_keep_above
// gtk_window_set_keep_below
// gtk_window_begin_resize_drag
// gtk_window_begin_move_drag
// gtk_window_set_decorated
// gtk_window_set_deletable
// gtk_window_set_frame_dimensions //deprecated since 2.24
// gtk_window_set_has_frame  //deprecated since 2.24
// gtk_window_set_mnemonic_modifier
// gtk_window_set_type_hint
// gtk_window_set_skip_taskbar_hint
// gtk_window_set_skip_pager_hint
// gtk_window_set_urgency_hint

func (v *GtkWindow) SetAcceptFocus(setting bool) {
	C.gtk_window_set_accept_focus(C.to_GtkWindow(v.Widget), bool2gboolean(setting))
}

// gtk_window_set_focus_on_map
// gtk_window_set_startup_id
// gtk_window_set_role
// gtk_window_get_decorated
// gtk_window_get_deletable
// gtk_window_get_default_icon_list
// gtk_window_get_default_icon_name

func (v *GtkWindow) GetDefaultSize(width *int, height *int) {
	var cwidth, cheight C.gint
	C.gtk_window_get_default_size(C.to_GtkWindow(v.Widget), &cwidth, &cheight)
	*width = int(cwidth)
	*height = int(cheight)
}
func (v *GtkWindow) GetDestroyWithParent() bool {
	return gboolean2bool(C.gtk_window_get_destroy_with_parent(C.to_GtkWindow(v.Widget)))
}

// gtk_window_get_frame_dimensions //deprecated since 2.24
// gtk_window_get_has_frame  //deprecated since 2.24
// gtk_window_get_icon
// gtk_window_get_icon_list
// gtk_window_get_icon_name
// gtk_window_get_mnemonic_modifier

func (v *GtkWindow) GetModal() bool {
	return gboolean2bool(C.gtk_window_get_modal(C.to_GtkWindow(v.Widget)))
}
func (v *GtkWindow) GetPosition(root_x *int, root_y *int) {
	var croot_x, croot_y C.gint
	C.gtk_window_get_position(C.to_GtkWindow(v.Widget), &croot_x, &croot_y)
	*root_x = int(croot_x)
	*root_y = int(croot_y)
}

// gtk_window_get_role

func (v *GtkWindow) GetSize(width *int, height *int) {
	var cwidth, cheight C.gint
	C.gtk_window_get_size(C.to_GtkWindow(v.Widget), &cwidth, &cheight)
	*width = int(cwidth)
	*height = int(cheight)
}
func (v *GtkWindow) GetTitle() string {
	return C.GoString(C.to_charptr(C.gtk_window_get_title(C.to_GtkWindow(v.Widget))))
}

// gtk_window_get_transient_for
// gtk_window_get_type_hint
// gtk_window_get_skip_taskbar_hint
// gtk_window_get_skip_pager_hint
// gtk_window_get_urgency_hint

func (v *GtkWindow) GetAcceptFocus() bool {
	return gboolean2bool(C.gtk_window_get_accept_focus(C.to_GtkWindow(v.Widget)))
}

// gtk_window_get_focus_on_map
// gtk_window_get_group
// gtk_window_has_group //since 2.22
// gtk_window_get_window_type //since 2.20

func (v *GtkWindow) Move(x int, y int) {
	C.gtk_window_move(C.to_GtkWindow(v.Widget), C.gint(x), C.gint(y))
}

// gtk_window_parse_geometry
// gtk_window_reshow_with_initial_size

func (v *GtkWindow) Resize(width int, height int) {
	C.gtk_window_resize(C.to_GtkWindow(v.Widget), C.gint(width), C.gint(height))
}

func (v *GtkWindow) XID() int32 {
	return gdk.WindowFromUnsafe(unsafe.Pointer(v.Widget.window)).GetNativeWindowID()
}

// gtk_window_set_default_icon_list
// gtk_window_set_default_icon
// gtk_window_set_default_icon_from_file
// gtk_window_set_default_icon_name
// gtk_window_set_icon
// gtk_window_set_icon_list

func (v *GtkWindow) SetIconFromFile(file string) {
	ptr := C.CString(file)
	defer C.free_string(ptr)
	C.gtk_window_set_icon_from_file(C.to_GtkWindow(v.Widget), C.to_gcharptr(ptr), nil) // last arg : GError **err
}

// gtk_window_set_icon_name
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
type GtkAboutDialog struct {
	GtkDialog
}

func AboutDialog() *GtkAboutDialog {
	return &GtkAboutDialog{GtkDialog{GtkWindow{GtkBin{GtkContainer{GtkWidget{
		C.gtk_about_dialog_new()}}}}}}
}
func (v *GtkAboutDialog) GetProgramName() string {
	return C.GoString(C.to_charptr(C.gtk_about_dialog_get_program_name(C.to_GtkAboutDialog(v.Widget))))
}
func (v *GtkAboutDialog) SetProgramName(name string) {
	ptr := C.CString(name)
	defer C.free_string(ptr)
	C.gtk_about_dialog_set_program_name(C.to_GtkAboutDialog(v.Widget), C.to_gcharptr(ptr))
}
func (v *GtkAboutDialog) GetVersion() string {
	return C.GoString(C.to_charptr(C.gtk_about_dialog_get_version(C.to_GtkAboutDialog(v.Widget))))
}
func (v *GtkAboutDialog) SetVersion(version string) {
	ptr := C.CString(version)
	defer C.free_string(ptr)
	C.gtk_about_dialog_set_version(C.to_GtkAboutDialog(v.Widget), C.to_gcharptr(ptr))
}
func (v *GtkAboutDialog) GetCopyright() string {
	return C.GoString(C.to_charptr(C.gtk_about_dialog_get_copyright(C.to_GtkAboutDialog(v.Widget))))
}
func (v *GtkAboutDialog) SetCopyright(copyright string) {
	ptr := C.CString(copyright)
	defer C.free_string(ptr)
	C.gtk_about_dialog_set_copyright(C.to_GtkAboutDialog(v.Widget), C.to_gcharptr(ptr))
}
func (v *GtkAboutDialog) GetComments() string {
	return C.GoString(C.to_charptr(C.gtk_about_dialog_get_comments(C.to_GtkAboutDialog(v.Widget))))
}
func (v *GtkAboutDialog) SetComments(comments string) {
	ptr := C.CString(comments)
	defer C.free_string(ptr)
	C.gtk_about_dialog_set_comments(C.to_GtkAboutDialog(v.Widget), C.to_gcharptr(ptr))
}
func (v *GtkAboutDialog) GetLicense() string {
	return C.GoString(C.to_charptr(C.gtk_about_dialog_get_license(C.to_GtkAboutDialog(v.Widget))))
}
func (v *GtkAboutDialog) SetLicense(license string) {
	ptr := C.CString(license)
	defer C.free_string(ptr)
	C.gtk_about_dialog_set_license(C.to_GtkAboutDialog(v.Widget), C.to_gcharptr(ptr))
}
func (v *GtkAboutDialog) GetWrapLicense() bool {
	return gboolean2bool(C.gtk_about_dialog_get_wrap_license(C.to_GtkAboutDialog(v.Widget)))
}
func (v *GtkAboutDialog) SetWrapLicense(wrap_license bool) {
	C.gtk_about_dialog_set_wrap_license(C.to_GtkAboutDialog(v.Widget), bool2gboolean(wrap_license))
}
func (v *GtkAboutDialog) GetWebsite() string {
	return C.GoString(C.to_charptr(C.gtk_about_dialog_get_website(C.to_GtkAboutDialog(v.Widget))))
}
func (v *GtkAboutDialog) SetWebsite(website string) {
	ptr := C.CString(website)
	defer C.free_string(ptr)
	C.gtk_about_dialog_set_website(C.to_GtkAboutDialog(v.Widget), C.to_gcharptr(ptr))
}
func (v *GtkAboutDialog) GetWebsiteLabel() string {
	return C.GoString(C.to_charptr(C.gtk_about_dialog_get_website_label(C.to_GtkAboutDialog(v.Widget))))
}
func (v *GtkAboutDialog) SetWebsiteLabel(website_label string) {
	ptr := C.CString(website_label)
	defer C.free_string(ptr)
	C.gtk_about_dialog_set_website_label(C.to_GtkAboutDialog(v.Widget), C.to_gcharptr(ptr))
}
func (v *GtkAboutDialog) GetAuthors() []string {
	var authors []string
	cauthors := C.gtk_about_dialog_get_authors(C.to_GtkAboutDialog(v.Widget))
	for {
		authors = append(authors, C.GoString(C.to_charptr(*cauthors)))
		cauthors = C.next_gcharptr(cauthors)
		if *cauthors == nil {
			break
		}
	}
	return authors
}
func (v *GtkAboutDialog) SetAuthors(authors []string) {
	cauthors := C.make_strings(C.int(len(authors) + 1))
	for i, author := range authors {
		ptr := C.CString(author)
		defer C.free_string(ptr)
		C.set_string(cauthors, C.int(i), C.to_gcharptr(ptr))
	}
	C.set_string(cauthors, C.int(len(authors)), nil)
	C.gtk_about_dialog_set_authors(C.to_GtkAboutDialog(v.Widget), cauthors)
	C.destroy_strings(cauthors)
}
func (v *GtkAboutDialog) GetArtists() []string {
	var artists []string
	cartists := C.gtk_about_dialog_get_artists(C.to_GtkAboutDialog(v.Widget))
	for {
		artists = append(artists, C.GoString(C.to_charptr(*cartists)))
		cartists = C.next_gcharptr(cartists)
		if *cartists == nil {
			break
		}
	}
	return artists
}
func (v *GtkAboutDialog) SetArtists(artists []string) {
	cartists := C.make_strings(C.int(len(artists)))
	for i, author := range artists {
		ptr := C.CString(author)
		defer C.free_string(ptr)
		C.set_string(cartists, C.int(i), C.to_gcharptr(ptr))
	}
	C.gtk_about_dialog_set_artists(C.to_GtkAboutDialog(v.Widget), cartists)
	C.destroy_strings(cartists)
}
func (v *GtkAboutDialog) GetDocumenters() []string {
	var documenters []string
	cdocumenters := C.gtk_about_dialog_get_documenters(C.to_GtkAboutDialog(v.Widget))
	for {
		documenters = append(documenters, C.GoString(C.to_charptr(*cdocumenters)))
		cdocumenters = C.next_gcharptr(cdocumenters)
		if *cdocumenters == nil {
			break
		}
	}
	return documenters
}
func (v *GtkAboutDialog) SetDocumenters(documenters []string) {
	cdocumenters := C.make_strings(C.int(len(documenters)))
	for i, author := range documenters {
		ptr := C.CString(author)
		defer C.free_string(ptr)
		C.set_string(cdocumenters, C.int(i), C.to_gcharptr(ptr))
	}
	C.gtk_about_dialog_set_documenters(C.to_GtkAboutDialog(v.Widget), cdocumenters)
	C.destroy_strings(cdocumenters)
}
func (v *GtkAboutDialog) GetTranslatorCredits() string {
	return C.GoString(C.to_charptr(C.gtk_about_dialog_get_translator_credits(C.to_GtkAboutDialog(v.Widget))))
}
func (v *GtkAboutDialog) SetTranslatorCredits(translator_credits string) {
	ptr := C.CString(translator_credits)
	defer C.free_string(ptr)
	C.gtk_about_dialog_set_translator_credits(C.to_GtkAboutDialog(v.Widget), C.to_gcharptr(ptr))
}
func (v *GtkAboutDialog) GetLogo() *gdkpixbuf.GdkPixbuf {
	return &gdkpixbuf.GdkPixbuf{
		C.gtk_about_dialog_get_logo(C.to_GtkAboutDialog(v.Widget))}
}
func (v *GtkAboutDialog) SetLogo(logo *gdkpixbuf.GdkPixbuf) {
	C.gtk_about_dialog_set_logo(C.to_GtkAboutDialog(v.Widget), logo.Pixbuf)
}
func (v *GtkAboutDialog) GetLogoIconName() string {
	return C.GoString(C.to_charptr(C.gtk_about_dialog_get_logo_icon_name(C.to_GtkAboutDialog(v.Widget))))
}
func (v *GtkAboutDialog) SetLogoIconName(icon_name string) {
	ptr := C.CString(icon_name)
	defer C.free_string(ptr)
	C.gtk_about_dialog_set_logo_icon_name(C.to_GtkAboutDialog(v.Widget), C.to_gcharptr(ptr))
}

// gtk_about_dialog_set_email_hook //deprecated since 2.24
// gtk_about_dialog_set_url_hook //deprecated since 2.24
// gtk_show_about_dialog

//-----------------------------------------------------------------------
// GtkAssistant
//-----------------------------------------------------------------------
type GtkAssistantPageType int

const (
	GTK_ASSISTANT_PAGE_CONTENT  GtkAssistantPageType = 0
	GTK_ASSISTANT_PAGE_INTRO    GtkAssistantPageType = 1
	GTK_ASSISTANT_PAGE_CONFIRM  GtkAssistantPageType = 2
	GTK_ASSISTANT_PAGE_SUMMARY  GtkAssistantPageType = 3
	GTK_ASSISTANT_PAGE_PROGRESS GtkAssistantPageType = 4
)

type GtkAssistant struct {
	GtkWidget
}

func Assistant() *GtkAssistant {
	return &GtkAssistant{GtkWidget{C.gtk_assistant_new()}}
}
func (v *GtkAssistant) GetCurrentPage() int {
	return int(C.gtk_assistant_get_current_page(C.to_GtkAssistant(v.Widget)))
}
func (v *GtkAssistant) SetCurrentPage(page_num int) {
	C.gtk_assistant_set_current_page(C.to_GtkAssistant(v.Widget), C.gint(page_num))
}
func (v *GtkAssistant) GetNPages() int {
	return int(C.gtk_assistant_get_n_pages(C.to_GtkAssistant(v.Widget)))
}
func (v *GtkAssistant) GetNthPage(page_num int) *GtkWidget {
	return &GtkWidget{
		C.gtk_assistant_get_nth_page(C.to_GtkAssistant(v.Widget), C.gint(page_num))}
}
func (v *GtkAssistant) PrependPage(page WidgetLike) int {
	return int(C.gtk_assistant_prepend_page(C.to_GtkAssistant(v.Widget), page.ToNative()))
}
func (v *GtkAssistant) AppendPage(page WidgetLike) int {
	return int(C.gtk_assistant_prepend_page(C.to_GtkAssistant(v.Widget), page.ToNative()))
}
func (v *GtkAssistant) InsertPage(page WidgetLike, position int) int {
	return int(C.gtk_assistant_insert_page(C.to_GtkAssistant(v.Widget), page.ToNative(), C.gint(position)))
}

// void gtk_assistant_set_forward_page_func (GtkAssistant *assistant, GtkAssistantPageFunc page_func, gpointer data, GDestroyNotify destroy);

func (v *GtkAssistant) SetPageType(page WidgetLike, t GtkAssistantPageType) {
	C.gtk_assistant_set_page_type(C.to_GtkAssistant(v.Widget), page.ToNative(), C.GtkAssistantPageType(t))
}
func (v *GtkAssistant) GetPageType(page WidgetLike) GtkAssistantPageType {
	return GtkAssistantPageType(C.gtk_assistant_get_page_type(C.to_GtkAssistant(v.Widget), page.ToNative()))
}
func (v *GtkAssistant) SetPageTitle(page WidgetLike, title string) {
	ptr := C.CString(title)
	defer C.free_string(ptr)
	C.gtk_assistant_set_page_title(C.to_GtkAssistant(v.Widget), page.ToNative(), C.to_gcharptr(ptr))
}
func (v *GtkAssistant) GetPageTitle(page WidgetLike) string {
	return C.GoString(C.to_charptr(C.gtk_assistant_get_page_title(C.to_GtkAssistant(v.Widget), page.ToNative())))
}
func (v *GtkAssistant) SetPageHeaderImage(page WidgetLike, pixbuf *gdkpixbuf.GdkPixbuf) {
	C.gtk_assistant_set_page_header_image(C.to_GtkAssistant(v.Widget), page.ToNative(), pixbuf.Pixbuf)
}
func (v *GtkAssistant) GetPageHeaderImage(page WidgetLike) *gdkpixbuf.GdkPixbuf {
	return &gdkpixbuf.GdkPixbuf{
		C.gtk_assistant_get_page_header_image(C.to_GtkAssistant(v.Widget), page.ToNative())}
}
func (v *GtkAssistant) SetPageSideImage(page WidgetLike, pixbuf *gdkpixbuf.GdkPixbuf) {
	C.gtk_assistant_set_page_side_image(C.to_GtkAssistant(v.Widget), page.ToNative(), pixbuf.Pixbuf)
}
func (v *GtkAssistant) GetPageSideImage(page WidgetLike) *gdkpixbuf.GdkPixbuf {
	return &gdkpixbuf.GdkPixbuf{
		C.gtk_assistant_get_page_side_image(C.to_GtkAssistant(v.Widget), page.ToNative())}
}
func (v *GtkAssistant) SetPageComplete(page WidgetLike, complete bool) {
	C.gtk_assistant_set_page_complete(C.to_GtkAssistant(v.Widget), page.ToNative(), bool2gboolean(complete))
}
func (v *GtkAssistant) GetPageComplete(page WidgetLike) bool {
	return gboolean2bool(C.gtk_assistant_get_page_complete(C.to_GtkAssistant(v.Widget), page.ToNative()))
}
func (v *GtkAssistant) AddActionWidget(child WidgetLike) {
	C.gtk_assistant_add_action_widget(C.to_GtkAssistant(v.Widget), child.ToNative())
}
func (v *GtkAssistant) RemoveActionWidget(child WidgetLike) {
	C.gtk_assistant_remove_action_widget(C.to_GtkAssistant(v.Widget), child.ToNative())
}
func (v *GtkAssistant) UpdateButtonsState() {
	C.gtk_assistant_update_buttons_state(C.to_GtkAssistant(v.Widget))
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
type GtkAccelLabel struct {
	GtkWidget
}

func AccelLabel(label string) *GtkAccelLabel {
	ptr := C.CString(label)
	defer C.free_string(ptr)
	return &GtkAccelLabel{GtkWidget{
		C.gtk_accel_label_new(C.to_gcharptr(ptr))}}
}

// gtk_accel_label_set_accel_closure

func (v *GtkAccelLabel) GetAccelWidget() GtkWidget {
	return GtkWidget{C.gtk_accel_label_get_accel_widget(C.to_GtkAccelLabel(v.Widget))}
}
func (v *GtkAccelLabel) SetAccelWidget(w WidgetLike) {
	C.gtk_accel_label_set_accel_widget(C.to_GtkAccelLabel(v.Widget), w.ToNative())
}
func (v *GtkAccelLabel) GetAccelWidth() uint {
	return uint(C.gtk_accel_label_get_accel_width(C.to_GtkAccelLabel(v.Widget)))
}
func (v *GtkAccelLabel) Refetch() bool {
	return gboolean2bool(C.gtk_accel_label_refetch(C.to_GtkAccelLabel(v.Widget)))
}

//-----------------------------------------------------------------------
// GtkImage
//-----------------------------------------------------------------------
type GtkIconSize int

const (
	GTK_ICON_SIZE_INVALID       GtkIconSize = 0
	GTK_ICON_SIZE_MENU          GtkIconSize = 1
	GTK_ICON_SIZE_SMALL_TOOLBAR GtkIconSize = 2
	GTK_ICON_SIZE_LARGE_TOOLBAR GtkIconSize = 3
	GTK_ICON_SIZE_BUTTON        GtkIconSize = 4
	GTK_ICON_SIZE_DND           GtkIconSize = 5
	GTK_ICON_SIZE_DIALOG        GtkIconSize = 6
)

/*type ImageLike interface {
	WidgetLike
}*/
type GtkImage struct {
	GtkWidget
}

// gtk_image_get_icon_set
// gtk_image_get_image

func (v *GtkImage) GetPixbuf() *gdkpixbuf.GdkPixbuf {
	return &gdkpixbuf.GdkPixbuf{
		C.gtk_image_get_pixbuf(C.to_GtkImage(v.Widget))}
}

// gtk_image_get_pixmap
// gtk_image_get_stock
// gtk_image_get_animation
// gtk_image_get_icon_name
// gtk_image_get_gicon
// gtk_image_get_storage_type

func ImageFromFile(filename string) *GtkImage {
	ptr := C.CString(filename)
	defer C.free_string(ptr)
	return &GtkImage{GtkWidget{
		C.gtk_image_new_from_file(C.to_gcharptr(ptr))}}
}

// gtk_image_new_from_icon_set
// gtk_image_new_from_image

func ImageFromPixbuf(pixbuf gdkpixbuf.GdkPixbuf) *GtkImage {
	return &GtkImage{GtkWidget{
		C.gtk_image_new_from_pixbuf(pixbuf.Pixbuf)}}
}

// gtk_image_new_from_pixmap

func ImageFromStock(stock_id string, size GtkIconSize) *GtkImage {
	ptr := C.CString(stock_id)
	defer C.free_string(ptr)
	return &GtkImage{GtkWidget{
		C.gtk_image_new_from_stock(C.to_gcharptr(ptr), C.GtkIconSize(size))}}
}

// gtk_image_new_from_animation
// gtk_image_new_from_icon_name
// gtk_image_new_from_gicon

func (v *GtkImage) SetFromFile(filename string) {
	ptr := C.CString(filename)
	defer C.free_string(ptr)
	C.gtk_image_set_from_file(C.to_GtkImage(v.Widget), C.to_gcharptr(ptr))
}

// gtk_image_set_from_icon_set
// gtk_image_set_from_image

func (v *GtkImage) SetFromPixbuf(pixbuf *gdkpixbuf.GdkPixbuf) {
	C.gtk_image_set_from_pixbuf(C.to_GtkImage(v.Widget), pixbuf.Pixbuf)
}

// gtk_image_set_from_pixmap

func (v *GtkImage) SetFromStock(stock_id string, size GtkIconSize) {
	ptr := C.CString(stock_id)
	defer C.free_string(ptr)
	C.gtk_image_set_from_stock(C.to_GtkImage(v.Widget), C.to_gcharptr(ptr), C.GtkIconSize(size))
}

// gtk_image_set_from_animation
// gtk_image_set_from_icon_name
// gtk_image_set_from_gicon

func (v *GtkImage) Clear() {
	C.gtk_image_clear(C.to_GtkImage(v.Widget))
}
func Image() *GtkImage {
	return &GtkImage{GtkWidget{
		C.gtk_image_new()}}
}

// gtk_image_set_pixel_size
// gtk_image_get_pixel_size

//-----------------------------------------------------------------------
// GtkLabel
//-----------------------------------------------------------------------
type GtkJustification int

const (
	GTK_JUSTIFY_LEFT   GtkJustification = 0
	GTK_JUSTIFY_RIGHT  GtkJustification = 1
	GTK_JUSTIFY_CENTER GtkJustification = 2
	GTK_JUSTIFY_FILL   GtkJustification = 3
)

type LabelLike interface {
	WidgetLike
	isLabelLike()
	GetLabel() string
	SetLabel(label string)
}
type GtkLabel struct {
	GtkWidget
}

func (GtkLabel) isLabelLike() {}

func Label(label string) *GtkLabel {
	ptr := C.CString(label)
	defer C.free_string(ptr)
	return &GtkLabel{GtkWidget{C.gtk_label_new(C.to_gcharptr(ptr))}}
}
func (v *GtkLabel) SetText(label string) {
	ptr := C.CString(label)
	defer C.free_string(ptr)
	C.gtk_label_set_text(C.to_GtkLabel(v.Widget), C.to_gcharptr(ptr))
}

func (v *GtkLabel) SetMnemonicWidget(widget WidgetLike) {
	C.gtk_label_set_mnemonic_widget(C.to_GtkLabel(v.Widget), widget.ToNative())
}

func (v *GtkLabel) SetMarkup(markup string) {
	ptr := C.CString(markup)
	defer C.free_string(ptr)
	C.gtk_label_set_markup(C.to_GtkLabel(v.Widget), C.to_gcharptr(ptr))
}

func (v *GtkLabel) GetMnemonicWidget() *GtkWidget {
	return &GtkWidget{C.gtk_label_get_mnemonic_widget(C.to_GtkLabel(v.Widget))}
}

func (v *GtkLabel) SetPattern(pattern string) {
	ptr := C.CString(pattern)
	defer C.free_string(ptr)
	C.gtk_label_set_pattern(C.to_GtkLabel(v.Widget), C.to_gcharptr(ptr))
}
func (v *GtkLabel) SetJustify(jtype GtkJustification) {
	C.gtk_label_set_justify(C.to_GtkLabel(v.Widget), C.GtkJustification(jtype))
}
func (v *GtkLabel) SetEllipsize(ellipsize pango.PangoEllipsizeMode) {
	C.gtk_label_set_ellipsize(C.to_GtkLabel(v.Widget), C.PangoEllipsizeMode(ellipsize))
}
func (v *GtkLabel) SetWidthChars(n_chars int) {
	C.gtk_label_set_width_chars(C.to_GtkLabel(v.Widget), C.gint(n_chars))
}
func (v *GtkLabel) SetMaxWidthChars(n_chars int) {
	C.gtk_label_set_max_width_chars(C.to_GtkLabel(v.Widget), C.gint(n_chars))
}
func (v *GtkLabel) SetLineWrap(setting bool) {
	C.gtk_label_set_line_wrap(C.to_GtkLabel(v.Widget), bool2gboolean(setting))
}
func (v *GtkLabel) SetUseLineWrapMode(wrap_mode pango.PangoWrapMode) {
	C.gtk_label_set_line_wrap_mode(C.to_GtkLabel(v.Widget), C.PangoWrapMode(wrap_mode))
}

// gtk_label_get_layout_offsets
// gtk_label_get_mnemonic_keyval

func (v *GtkLabel) GetSelectable() bool {
	return gboolean2bool(C.gtk_label_get_selectable(C.to_GtkLabel(v.Widget)))
}
func (v *GtkLabel) GetText() string {
	return C.GoString(C.to_charptr(C.gtk_label_get_text(C.to_GtkLabel(v.Widget))))
}
func LabelWithMnemonic(label string) *GtkLabel {
	ptr := C.CString(label)
	defer C.free_string(ptr)
	return &GtkLabel{GtkWidget{C.gtk_label_new_with_mnemonic(C.to_gcharptr(ptr))}}
}
func (v *GtkLabel) SelectRegion(start_offset int, end_offset int) {
	C.gtk_label_select_region(C.to_GtkLabel(v.Widget), C.gint(start_offset), C.gint(end_offset))
}

// gtk_label_set_mnemonic_widget

func (v *GtkLabel) SetSelectable(setting bool) {
	C.gtk_label_set_selectable(C.to_GtkLabel(v.Widget), bool2gboolean(setting))
}
func (v *GtkLabel) SetTextWithMnemonic(str string) {
	ptr := C.CString(str)
	defer C.free_string(ptr)
	C.gtk_label_set_text_with_mnemonic(C.to_GtkLabel(v.Widget), C.to_gcharptr(ptr))
}

// gtk_label_get_attributes
func (v *GtkLabel) GetJustify() GtkJustification {
	return GtkJustification(C.gtk_label_get_justify(C.to_GtkLabel(v.Widget)))
}
func (v *GtkLabel) GetEllipsize() pango.PangoEllipsizeMode {
	return pango.PangoEllipsizeMode(C.gtk_label_get_ellipsize(C.to_GtkLabel(v.Widget)))
}
func (v *GtkLabel) GetWidthChars() int {
	return int(C.gtk_label_get_width_chars(C.to_GtkLabel(v.Widget)))
}
func (v *GtkLabel) GetMaxWidthChars() int {
	return int(C.gtk_label_get_max_width_chars(C.to_GtkLabel(v.Widget)))
}
func (v *GtkLabel) GetLabel() string {
	return C.GoString(C.to_charptr(C.gtk_label_get_label(C.to_GtkLabel(v.Widget))))
}

// gtk_label_get_layout

func (v *GtkLabel) GetLineWrap() bool {
	return gboolean2bool(C.gtk_label_get_line_wrap(C.to_GtkLabel(v.Widget)))
}
func (v *GtkLabel) GetLineWrapMode() pango.PangoWrapMode {
	return pango.PangoWrapMode(C.gtk_label_get_line_wrap_mode(C.to_GtkLabel(v.Widget)))
}

// gtk_label_get_mnemonic_widget
func (v *GtkLabel) GetSelectionBounds(start *int, end *int) {
	var cstart, cend C.gint
	C.gtk_label_get_selection_bounds(C.to_GtkLabel(v.Widget), &cstart, &cend)
	*start = int(cstart)
	*end = int(cend)
}
func (v *GtkLabel) GetUseMarkup() bool {
	return gboolean2bool(C.gtk_label_get_use_markup(C.to_GtkLabel(v.Widget)))
}
func (v *GtkLabel) GetUseUnderline() bool {
	return gboolean2bool(C.gtk_label_get_use_underline(C.to_GtkLabel(v.Widget)))
}
func (v *GtkLabel) GetSingleLineMode() bool {
	return gboolean2bool(C.gtk_label_get_single_line_mode(C.to_GtkLabel(v.Widget)))
}
func (v *GtkLabel) GetAngle() float64 {
	r := C.gtk_label_get_angle(C.to_GtkLabel(v.Widget))
	return float64(r)
}
func (v *GtkLabel) SetLabel(label string) {
	ptr := C.CString(label)
	defer C.free_string(ptr)
	C.gtk_label_set_label(C.to_GtkLabel(v.Widget), C.to_gcharptr(ptr))
}
func (v *GtkLabel) SetUseMarkup(setting bool) {
	C.gtk_label_set_use_markup(C.to_GtkLabel(v.Widget), bool2gboolean(setting))
}
func (v *GtkLabel) SetUseUnderline(setting bool) {
	C.gtk_label_set_use_underline(C.to_GtkLabel(v.Widget), bool2gboolean(setting))
}
func (v *GtkLabel) SetSingleLineMode(single_line bool) {
	C.gtk_label_set_single_line_mode(C.to_GtkLabel(v.Widget), bool2gboolean(single_line))
}
func (v *GtkLabel) SetAngle(angle float64) {
	C.gtk_label_set_angle(C.to_GtkLabel(v.Widget), C.gdouble(angle))
}
func (v *GtkLabel) GetCurrentUri() string {
	panic_if_version_older(2, 18, 0, "gtk_label_get_current_uri()")
	return C.GoString(C.to_charptr(C.gtk_label_get_current_uri(C.to_GtkLabel(v.Widget))))
}
func (v *GtkLabel) SetTrackVisitedLinks(track_links bool) {
	panic_if_version_older(2, 18, 0, "gtk_label_set_track_visited_links()")
	C.gtk_label_set_track_visited_links(C.to_GtkLabel(v.Widget), bool2gboolean(track_links))
}
func (v *GtkLabel) GetTrackVisitedLinks() bool {
	panic_if_version_older(2, 18, 0, "gtk_label_get_track_visited_links()")
	return gboolean2bool(C.gtk_label_get_track_visited_links(C.to_GtkLabel(v.Widget)))
}

//-----------------------------------------------------------------------
// GtkProgressBar
//-----------------------------------------------------------------------
type GtkProgressBarOrientation int

const (
	GTK_PROGRESS_LEFT_TO_RIGHT GtkProgressBarOrientation = 0
	GTK_PROGRESS_RIGHT_TO_LEFT GtkProgressBarOrientation = 1
	GTK_PROGRESS_BOTTOM_TO_TOP GtkProgressBarOrientation = 2
	GTK_PROGRESS_TOP_TO_BOTTOM GtkProgressBarOrientation = 3
)

type GtkProgressBar struct {
	GtkWidget
}

func ProgressBar() *GtkProgressBar {
	return &GtkProgressBar{GtkWidget{C.gtk_progress_bar_new()}}
}
func (v *GtkProgressBar) Pulse() {
	C.gtk_progress_bar_pulse(C.to_GtkProgressBar(v.Widget))
}
func (v *GtkProgressBar) SetText(show_text string) {
	ptr := C.CString(show_text)
	defer C.free_string(ptr)
	C.gtk_progress_bar_set_text(C.to_GtkProgressBar(v.Widget), C.to_gcharptr(ptr))
}
func (v *GtkProgressBar) SetFraction(fraction float64) {
	C.gtk_progress_bar_set_fraction(C.to_GtkProgressBar(v.Widget), C.gdouble(fraction))
}
func (v *GtkProgressBar) SetPulseStep(fraction float64) {
	C.gtk_progress_bar_set_pulse_step(C.to_GtkProgressBar(v.Widget), C.gdouble(fraction))
}
func (v *GtkProgressBar) SetOrientation(i GtkProgressBarOrientation) {
	C.gtk_progress_bar_set_orientation(C.to_GtkProgressBar(v.Widget), C.GtkProgressBarOrientation(i))
}

// gtk_progress_bar_set_ellipsize

func (v *GtkProgressBar) GetText() string {
	return C.GoString(C.to_charptr(C.gtk_progress_bar_get_text(C.to_GtkProgressBar(v.Widget))))
}
func (v *GtkProgressBar) GetFraction() float64 {
	r := C.gtk_progress_bar_get_fraction(C.to_GtkProgressBar(v.Widget))
	return float64(r)
}
func (v *GtkProgressBar) GetPulseStep() float64 {
	r := C.gtk_progress_bar_get_pulse_step(C.to_GtkProgressBar(v.Widget))
	return float64(r)
}
func (v *GtkProgressBar) GetOrientation() GtkProgressBarOrientation {
	return GtkProgressBarOrientation(C.gtk_progress_bar_get_orientation(C.to_GtkProgressBar(v.Widget)))
}

// gtk_progress_bar_get_ellipsize

//-----------------------------------------------------------------------
// GtkStatusbar
//-----------------------------------------------------------------------
type GtkStatusbar struct {
	GtkHBox
}

func Statusbar() *GtkStatusbar {
	return &GtkStatusbar{GtkHBox{GtkBox{GtkContainer{GtkWidget{
		C.gtk_statusbar_new()}}}}}
}
func (v *GtkStatusbar) GetContextId(content_description string) uint {
	ptr := C.CString(content_description)
	defer C.free_string(ptr)
	return uint(C.gtk_statusbar_get_context_id(C.to_GtkStatusbar(v.Widget), C.to_gcharptr(ptr)))
}
func (v *GtkStatusbar) Push(context_id uint, text string) {
	ptr := C.CString(text)
	defer C.free_string(ptr)
	C.gtk_statusbar_push(C.to_GtkStatusbar(v.Widget), C.guint(context_id), C.to_gcharptr(ptr))
}
func (v *GtkStatusbar) Pop(context_id uint) {
	C.gtk_statusbar_pop(C.to_GtkStatusbar(v.Widget), C.guint(context_id))
}
func (v *GtkStatusbar) Remove(context_id uint, message_id uint) {
	C.gtk_statusbar_remove(C.to_GtkStatusbar(v.Widget), C.guint(context_id), C.guint(message_id))
}

// gtk_statusbar_remove_all //since 2.22

func (v *GtkStatusbar) SetHasResizeGrip(add_tearoffs bool) {
	C.gtk_statusbar_set_has_resize_grip(C.to_GtkStatusbar(v.Widget), bool2gboolean(add_tearoffs))
}
func (v *GtkStatusbar) GetHasResizeGrip() bool {
	return gboolean2bool(C.gtk_statusbar_get_has_resize_grip(C.to_GtkStatusbar(v.Widget)))
}

// gtk_statusbar_get_message_area //since 2.20

//-----------------------------------------------------------------------
// GtkInfoBar
//-----------------------------------------------------------------------
type GtkInfoBar struct {
	GtkHBox
}

func InfoBar() *GtkInfoBar {
	panic_if_version_older_auto(2, 18, 0)
	return &GtkInfoBar{GtkHBox{GtkBox{GtkContainer{GtkWidget{
		C._gtk_info_bar_new()}}}}}
}

func InfoBarWithButtons(buttons ...interface{}) *GtkInfoBar {
	panic_if_version_older_auto(2, 18, 0)
	infobar := InfoBar()
	text, res := variadicButtonsToArrays(buttons)
	for i := range text {
		infobar.AddButton(text[i], res[i])
	}
	return infobar
}

func (v *GtkInfoBar) AddActionWidget(child WidgetLike, responseId int) {
	panic_if_version_older_auto(2, 18, 0)
	C._gtk_info_bar_add_action_widget(C.to_GtkInfoBar(v.Widget),
		child.ToNative(), C.gint(responseId))
}

func (v *GtkInfoBar) AddButton(buttonText string, responseId int) *GtkWidget {
	panic_if_version_older_auto(2, 18, 0)
	ptr := C.CString(buttonText)
	defer C.free_string(ptr)
	return &GtkWidget{C._gtk_info_bar_add_button(C.to_GtkInfoBar(v.Widget),
		C.to_gcharptr(ptr), C.gint(responseId))}
}

func (v *GtkInfoBar) AddButtons(buttons ...interface{}) {
	panic_if_version_older_auto(2, 18, 0)
	text, res := variadicButtonsToArrays(buttons)
	for i := range text {
		v.AddButton(text[i], res[i])
	}
}

func (v *GtkInfoBar) SetResponseSensitive(responseId int, setting bool) {
	panic_if_version_older_auto(2, 18, 0)
	C._gtk_info_bar_set_response_sensitive(C.to_GtkInfoBar(v.Widget),
		C.gint(responseId), bool2gboolean(setting))
}

func (v *GtkInfoBar) SetDefaultResponse(responseId int) {
	panic_if_version_older_auto(2, 18, 0)
	C._gtk_info_bar_set_default_response(C.to_GtkInfoBar(v.Widget),
		C.gint(responseId))
}

func (v *GtkInfoBar) Response(responseId int) {
	panic_if_version_older_auto(2, 18, 0)
	C._gtk_info_bar_response(C.to_GtkInfoBar(v.Widget), C.gint(responseId))
}

func (v *GtkInfoBar) SetMessageType(messageType GtkMessageType) {
	panic_if_version_older_auto(2, 18, 0)
	C._gtk_info_bar_set_message_type(C.to_GtkInfoBar(v.Widget),
		C.GtkMessageType(messageType))
}

func (v *GtkInfoBar) GetMessageType() GtkMessageType {
	panic_if_version_older_auto(2, 18, 0)
	return GtkMessageType(C._gtk_info_bar_get_message_type(C.to_GtkInfoBar(v.Widget)))
}

func (v *GtkInfoBar) GetActionArea() *GtkWidget {
	panic_if_version_older_auto(2, 18, 0)
	return &GtkWidget{C._gtk_info_bar_get_action_area(C.to_GtkInfoBar(v.Widget))}
}

func (v *GtkInfoBar) GetContentArea() *GtkWidget {
	panic_if_version_older_auto(2, 18, 0)
	return &GtkWidget{C._gtk_info_bar_get_content_area(C.to_GtkInfoBar(v.Widget))}
}

//-----------------------------------------------------------------------
// GtkStatusIcon
//-----------------------------------------------------------------------
type GtkStatusIcon struct {
	StatusIcon *C.GtkStatusIcon
}

func StatusIcon() *GtkStatusIcon {
	return &GtkStatusIcon{
		C.gtk_status_icon_new()}
}
func StatusIconFromPixbuf(pixbuf *gdkpixbuf.GdkPixbuf) *GtkStatusIcon {
	return &GtkStatusIcon{
		C.gtk_status_icon_new_from_pixbuf(pixbuf.Pixbuf)}
}
func StatusIconFromFile(filename string) *GtkStatusIcon {
	ptr := C.CString(filename)
	defer C.free_string(ptr)
	return &GtkStatusIcon{
		C.gtk_status_icon_new_from_file(C.to_gcharptr(ptr))}
}
func StatusIconFromStock(stock_id string) *GtkStatusIcon {
	ptr := C.CString(stock_id)
	defer C.free_string(ptr)
	return &GtkStatusIcon{
		C.gtk_status_icon_new_from_stock(C.to_gcharptr(ptr))}
}
func StatusIconFromIconName(icon_name string) *GtkStatusIcon {
	ptr := C.CString(icon_name)
	defer C.free_string(ptr)
	return &GtkStatusIcon{
		C.gtk_status_icon_new_from_icon_name(C.to_gcharptr(ptr))}
}

//GtkStatusIcon *gtk_status_icon_new_from_gicon(GIcon *icon);

func (v *GtkStatusIcon) SetFromPixbuf(pixbuf *gdkpixbuf.GdkPixbuf) {
	C.gtk_status_icon_set_from_pixbuf(v.StatusIcon, pixbuf.Pixbuf)
}
func (v *GtkStatusIcon) SetFromFile(filename string) {
	ptr := C.CString(filename)
	defer C.free_string(ptr)
	C.gtk_status_icon_set_from_file(v.StatusIcon, C.to_gcharptr(ptr))
}
func (v *GtkStatusIcon) SetFromStock(stock_id string) {
	ptr := C.CString(stock_id)
	defer C.free_string(ptr)
	C.gtk_status_icon_set_from_stock(v.StatusIcon, C.to_gcharptr(ptr))
}
func (v *GtkStatusIcon) SetFromIconName(icon_name string) {
	ptr := C.CString(icon_name)
	defer C.free_string(ptr)
	C.gtk_status_icon_set_from_icon_name(v.StatusIcon, C.to_gcharptr(ptr))
}

//void gtk_status_icon_set_from_gicon (GtkStatusIcon *status_icon, GIcon *icon);
//GtkImageType gtk_status_icon_get_storage_type (GtkStatusIcon *status_icon);

func (v *GtkStatusIcon) GetPixbuf() *gdkpixbuf.GdkPixbuf {
	return &gdkpixbuf.GdkPixbuf{
		C.gtk_status_icon_get_pixbuf(v.StatusIcon)}
}
func (v *GtkStatusIcon) GetStock() string {
	return C.GoString(C.to_charptr(C.gtk_status_icon_get_stock(v.StatusIcon)))
}
func (v *GtkStatusIcon) GetIconName() string {
	return C.GoString(C.to_charptr(C.gtk_status_icon_get_icon_name(v.StatusIcon)))
}
func (v *GtkStatusIcon) SetName(name string) {
	panic_if_version_older(2, 20, 0, "gtk_status_icon_set_name()")
	ptr := C.CString(name)
	defer C.free_string(ptr)
	C._gtk_status_icon_set_name(v.StatusIcon, C.to_gcharptr(ptr))
}
func (v *GtkStatusIcon) SetTitle(title string) {
	panic_if_version_older(2, 18, 0, "gtk_status_icon_set_title()")
	ptr := C.CString(title)
	defer C.free_string(ptr)
	C._gtk_status_icon_set_title(v.StatusIcon, C.to_gcharptr(ptr))
}
func (v *GtkStatusIcon) GetTitle() string {
	panic_if_version_older(2, 18, 0, "gtk_status_icon_get_title()")
	return C.GoString(C.to_charptr(C._gtk_status_icon_get_title(v.StatusIcon)))
}
func (v *GtkStatusIcon) SetTooltipText(text string) {
	ptr := C.CString(text)
	defer C.free_string(ptr)
	C.gtk_status_icon_set_tooltip_text(v.StatusIcon, C.to_gcharptr(ptr))
}
func (v *GtkStatusIcon) GetTooltipText() string {
	return C.GoString(C.to_charptr(C.gtk_status_icon_get_tooltip_text(v.StatusIcon)))
}
func (v *GtkStatusIcon) SetTooltipMarkup(markup string) {
	ptr := C.CString(markup)
	defer C.free_string(ptr)
	C.gtk_status_icon_set_tooltip_markup(v.StatusIcon, C.to_gcharptr(ptr))
}
func (v *GtkStatusIcon) GetTooltipMarkup() string {
	return C.GoString(C.to_charptr(C.gtk_status_icon_get_tooltip_markup(v.StatusIcon)))
}
func (v *GtkStatusIcon) GetHasTooltip() bool {
	return gboolean2bool(C.gtk_status_icon_get_has_tooltip(v.StatusIcon))
}
func (v *GtkStatusIcon) SetHasTooltip(setting bool) {
	C.gtk_status_icon_set_has_tooltip(v.StatusIcon, bool2gboolean(setting))
}
func (v *GtkStatusIcon) GetVisible() bool {
	return gboolean2bool(C.gtk_status_icon_get_visible(v.StatusIcon))
}
func (v *GtkStatusIcon) SetVisible(setting bool) {
	C.gtk_status_icon_set_visible(v.StatusIcon, bool2gboolean(setting))
}
func GtkStatusIconPositionMenu(menu *GtkMenu, px, py *int, push_in *bool, data interface{}) {
	x := C.gint(*px)
	y := C.gint(*py)
	pi := bool2gboolean(*push_in)
	var pdata C.gpointer
	if sm, ok := data.(*GtkStatusIcon); ok {
		pdata = C.gpointer(unsafe.Pointer(sm.StatusIcon))
	}
	C.gtk_status_icon_position_menu(C.to_GtkMenu(menu.Widget), &x, &y, &pi, pdata)
	*px = int(x)
	*py = int(y)
	*push_in = gboolean2bool(pi)
}
func (v *GtkStatusIcon) Connect(s string, f interface{}, datas ...interface{}) int {
	return glib.ObjectFromNative(unsafe.Pointer(C.to_GObject(unsafe.Pointer(v.StatusIcon)))).Connect(s, f, datas...)
}

func PrintContextFromNative(l unsafe.Pointer) *GtkPrintContext {
	return &GtkPrintContext{(*C.GtkPrintContext)(l)}
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

type GtkReliefStyle int

const (
	GTK_RELIEF_NORMAL GtkReliefStyle = 0
	GTK_RELIEF_HALF   GtkReliefStyle = 1
	GTK_RELIEF_NONE   GtkReliefStyle = 2
)

type GtkButton struct {
	GtkBin
}

func (GtkButton) isLabelLike() {}

func Button() *GtkButton {
	return &GtkButton{GtkBin{GtkContainer{GtkWidget{
		C.gtk_button_new()}}}}
}
func ButtonWithLabel(label string) *GtkButton {
	ptr := C.CString(label)
	defer C.free_string(ptr)
	return &GtkButton{GtkBin{GtkContainer{GtkWidget{
		C.gtk_button_new_with_label(C.to_gcharptr(ptr))}}}}
}
func ButtonWithMnemonic(label string) *GtkButton {
	ptr := C.CString(label)
	defer C.free_string(ptr)
	return &GtkButton{GtkBin{GtkContainer{GtkWidget{
		C.gtk_button_new_with_mnemonic(C.to_gcharptr(ptr))}}}}
}

// gtk_button_new_from_stock
// gtk_button_pressed //deprecated since 2.20
// gtk_button_released //deprecated since 2.20

func (v *GtkButton) Clicked(onclick interface{}, datas ...interface{}) int {
	return v.Connect("clicked", onclick, datas...)
}

// gtk_button_enter //deprecated since 2.20
// gtk_button_leave //deprecated since 2.20
func (v *GtkButton) GetRelief() GtkReliefStyle {
	return GtkReliefStyle(C.gtk_button_get_relief(C.to_GtkButton(v.Widget)))
}
func (v *GtkButton) SetRelief(relief GtkReliefStyle) {
	C.gtk_button_set_relief(C.to_GtkButton(v.Widget), C.GtkReliefStyle(relief))
}
func (v *GtkButton) GetLabel() string {
	return C.GoString(C.to_charptr(C.gtk_button_get_label(C.to_GtkButton(v.Widget))))
}
func (v *GtkButton) SetLabel(label string) {
	ptr := C.CString(label)
	defer C.free_string(ptr)
	C.gtk_button_set_label(C.to_GtkButton(v.Widget), C.to_gcharptr(ptr))
}
func (v *GtkButton) GetUseStock() bool {
	return gboolean2bool(C.gtk_button_get_use_stock(C.to_GtkButton(v.Widget)))
}
func (v *GtkButton) SetUseStock(use bool) {
	C.gtk_button_set_use_stock(C.to_GtkButton(v.Widget), bool2gboolean(use))
}
func (v *GtkButton) GetUseUnderline() bool {
	return gboolean2bool(C.gtk_button_get_use_underline(C.to_GtkButton(v.Widget)))
}
func (v *GtkButton) SetUseUnderline(setting bool) {
	C.gtk_button_set_use_underline(C.to_GtkButton(v.Widget), bool2gboolean(setting))
}
func (v *GtkButton) GetFocusOnClick() bool {
	return gboolean2bool(C.gtk_button_get_focus_on_click(C.to_GtkButton(v.Widget)))
}
func (v *GtkButton) SetFocusOnClick(setting bool) {
	C.gtk_button_set_focus_on_click(C.to_GtkButton(v.Widget), bool2gboolean(setting))
}
func (v *GtkButton) SetAlignment(xalign, yalign float64) {
	C.gtk_button_set_alignment(C.to_GtkButton(v.Widget), C.gfloat(xalign), C.gfloat(yalign))
}
func (v *GtkButton) GetAlignment() (xalign, yalign float64) {
	var xalign_, yalign_ C.gfloat
	C.gtk_button_get_alignment(C.to_GtkButton(v.Widget), &xalign_, &yalign_)
	return float64(xalign_), float64(yalign_)
}
func (v *GtkButton) SetImage(image WidgetLike) {
	C.gtk_button_set_image(C.to_GtkButton(v.Widget), image.ToNative())
}
func (v *GtkButton) GetImage() *GtkImage {
	return &GtkImage{GtkWidget{C.gtk_button_get_image(C.to_GtkButton(v.Widget))}}
}

// gtk_button_set_image_position
// gtk_button_get_image_position
// gtk_button_get_event_window

//-----------------------------------------------------------------------
// GtkCheckButton
//-----------------------------------------------------------------------
type GtkCheckButton struct {
	GtkToggleButton
}

func CheckButton() *GtkCheckButton {
	return &GtkCheckButton{GtkToggleButton{GtkButton{GtkBin{GtkContainer{GtkWidget{
		C.gtk_check_button_new()}}}}}}
}
func CheckButtonWithLabel(label string) *GtkCheckButton {
	ptr := C.CString(label)
	defer C.free_string(ptr)
	return &GtkCheckButton{GtkToggleButton{GtkButton{GtkBin{GtkContainer{GtkWidget{
		C.gtk_check_button_new_with_label(C.to_gcharptr(ptr))}}}}}}
}
func CheckButtonWithMnemonic(label string) *GtkCheckButton {
	ptr := C.CString(label)
	defer C.free_string(ptr)
	return &GtkCheckButton{GtkToggleButton{GtkButton{GtkBin{GtkContainer{GtkWidget{
		C.gtk_check_button_new_with_mnemonic(C.to_gcharptr(ptr))}}}}}}
}

//-----------------------------------------------------------------------
// GtkRadioButton
//-----------------------------------------------------------------------
type GtkRadioButton struct {
	GtkCheckButton
}

func RadioButton(group *glib.SList) *GtkRadioButton {
	if group != nil {
		return &GtkRadioButton{GtkCheckButton{GtkToggleButton{GtkButton{GtkBin{GtkContainer{GtkWidget{
			C.gtk_radio_button_new(C.to_gslist(unsafe.Pointer(group.ToSList())))}}}}}}}
	}
	return &GtkRadioButton{GtkCheckButton{GtkToggleButton{GtkButton{GtkBin{GtkContainer{GtkWidget{
		C.gtk_radio_button_new(nil)}}}}}}}
}
func RadioButtonFromWidget(w *GtkRadioButton) *GtkRadioButton {
	return &GtkRadioButton{GtkCheckButton{GtkToggleButton{GtkButton{GtkBin{GtkContainer{GtkWidget{
		C.gtk_radio_button_new_from_widget(C.to_GtkRadioButton(w.Widget))}}}}}}}
}
func RadioButtonWithLabel(group *glib.SList, label string) *GtkRadioButton {
	ptr := C.CString(label)
	defer C.free_string(ptr)
	if group != nil {
		return &GtkRadioButton{GtkCheckButton{GtkToggleButton{GtkButton{GtkBin{GtkContainer{GtkWidget{
			C.gtk_radio_button_new_with_label(C.to_gslist(unsafe.Pointer(group.ToSList())), C.to_gcharptr(ptr))}}}}}}}
	}
	return &GtkRadioButton{GtkCheckButton{GtkToggleButton{GtkButton{GtkBin{GtkContainer{GtkWidget{
		C.gtk_radio_button_new_with_label(nil, C.to_gcharptr(ptr))}}}}}}}
}
func RadioButtonWithLabelFromWidget(w *GtkRadioButton, label string) *GtkRadioButton {
	ptr := C.CString(label)
	defer C.free_string(ptr)
	return &GtkRadioButton{GtkCheckButton{GtkToggleButton{GtkButton{GtkBin{GtkContainer{GtkWidget{
		C.gtk_radio_button_new_with_label_from_widget(C.to_GtkRadioButton(w.Widget), C.to_gcharptr(ptr))}}}}}}}
}
func RadioButtonWithMnemonic(group *glib.SList, label string) *GtkRadioButton {
	ptr := C.CString(label)
	defer C.free_string(ptr)
	if group != nil {
		return &GtkRadioButton{GtkCheckButton{GtkToggleButton{GtkButton{GtkBin{GtkContainer{GtkWidget{
			C.gtk_radio_button_new_with_mnemonic(C.to_gslist(unsafe.Pointer(group.ToSList())), C.to_gcharptr(ptr))}}}}}}}
	}
	return &GtkRadioButton{GtkCheckButton{GtkToggleButton{GtkButton{GtkBin{GtkContainer{GtkWidget{
		C.gtk_radio_button_new_with_mnemonic(nil, C.to_gcharptr(ptr))}}}}}}}
}
func RadioButtonWithMnemonicFromWidget(w *GtkRadioButton, label string) *GtkRadioButton {
	ptr := C.CString(label)
	defer C.free_string(ptr)
	return &GtkRadioButton{GtkCheckButton{GtkToggleButton{GtkButton{GtkBin{GtkContainer{GtkWidget{
		C.gtk_radio_button_new_with_mnemonic_from_widget(C.to_GtkRadioButton(w.Widget), C.to_gcharptr(ptr))}}}}}}}
}
func (v *GtkRadioButton) GetGroup() *glib.SList {
	return glib.SListFromNative(unsafe.Pointer(C.gtk_radio_button_get_group(C.to_GtkRadioButton(v.Widget))))
}
func (v *GtkRadioButton) SetGroup(group *glib.SList) {
	if group != nil {
		C.gtk_radio_button_set_group(C.to_GtkRadioButton(v.Widget), C.to_gslist(unsafe.Pointer(group)))
	} else {
		C.gtk_radio_button_set_group(C.to_GtkRadioButton(v.Widget), nil)
	}
}

//-----------------------------------------------------------------------
// GtkToggleButton
//-----------------------------------------------------------------------
type GtkToggleButton struct {
	GtkButton
}

func ToggleButton() *GtkToggleButton {
	return &GtkToggleButton{GtkButton{GtkBin{GtkContainer{GtkWidget{
		C.gtk_toggle_button_new()}}}}}
}
func ToggleButtonWithLabel(label string) *GtkToggleButton {
	ptr := C.CString(label)
	defer C.free_string(ptr)
	return &GtkToggleButton{GtkButton{GtkBin{GtkContainer{GtkWidget{
		C.gtk_toggle_button_new_with_label(C.to_gcharptr(ptr))}}}}}
}
func ToggleButtonWithMnemonic(label string) *GtkToggleButton {
	ptr := C.CString(label)
	defer C.free_string(ptr)
	return &GtkToggleButton{GtkButton{GtkBin{GtkContainer{GtkWidget{
		C.gtk_toggle_button_new_with_mnemonic(C.to_gcharptr(ptr))}}}}}
}
func (v *GtkToggleButton) SetMode(draw_indicator bool) {
	C.gtk_toggle_button_set_mode(C.to_GtkToggleButton(v.Widget), bool2gboolean(draw_indicator))
}
func (v *GtkToggleButton) GetMode() bool {
	return gboolean2bool(C.gtk_toggle_button_get_mode(C.to_GtkToggleButton(v.Widget)))
}
func (v *GtkToggleButton) GetActive() bool {
	return gboolean2bool(C.gtk_toggle_button_get_active(C.to_GtkToggleButton(v.Widget)))
}
func (v *GtkToggleButton) SetActive(is_active bool) {
	C.gtk_toggle_button_set_active(C.to_GtkToggleButton(v.Widget), bool2gboolean(is_active))
}
func (v *GtkToggleButton) GetInconsistent() bool {
	return gboolean2bool(C.gtk_toggle_button_get_inconsistent(C.to_GtkToggleButton(v.Widget)))
}
func (v *GtkToggleButton) SetInconsistent(setting bool) {
	C.gtk_toggle_button_set_inconsistent(C.to_GtkToggleButton(v.Widget), bool2gboolean(setting))
}

//-----------------------------------------------------------------------
// GtkLinkButton
//-----------------------------------------------------------------------
type GtkLinkButton struct {
	GtkButton
}

func LinkButton(uri string) *GtkLinkButton {
	ptr := C.CString(uri)
	defer C.free_string(ptr)
	return &GtkLinkButton{GtkButton{GtkBin{GtkContainer{GtkWidget{
		C.gtk_link_button_new(C.to_gcharptr(ptr))}}}}}
}
func LinkButtonWithLabel(uri string, label string) *GtkLinkButton {
	puri := C.CString(uri)
	defer C.free_string(puri)
	plabel := C.CString(label)
	defer C.free_string(plabel)
	return &GtkLinkButton{GtkButton{GtkBin{GtkContainer{GtkWidget{
		C.gtk_link_button_new_with_label(C.to_gcharptr(puri), C.to_gcharptr(plabel))}}}}}
}
func (v *GtkLinkButton) GetUri() string {
	return C.GoString(C.to_charptr(C.gtk_link_button_get_uri(C.to_GtkLinkButton(v.Widget))))
}
func (v *GtkLinkButton) SetUri(uri string) {
	ptr := C.CString(uri)
	defer C.free_string(ptr)
	C.gtk_link_button_set_uri(C.to_GtkLinkButton(v.Widget), C.to_gcharptr(ptr))
}

//gtk_link_button_set_uri_hook has been deprecated since 2.24. Use clicked signal instead. //TODO
//func (v GtkLinkButton) SetUriHook(f func(button *GtkLinkButton, link string, user_data unsafe.Pointer), ) {
// GtkLinkButtonUriFunc gtk_link_button_set_uri_hook (GtkLinkButtonUriFunc func, gpointer data, GDestroyNotify destroy);
//}
func (v *GtkLinkButton) GetVisited() bool {
	return gboolean2bool(C.gtk_link_button_get_visited(C.to_GtkLinkButton(v.Widget)))
}
func (v *GtkLinkButton) SetVisited(visited bool) {
	C.gtk_link_button_set_visited(C.to_GtkLinkButton(v.Widget), bool2gboolean(visited))
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
type TextInputLike interface {
	WidgetLike
	GetText() string
	SetText(label string)
}
type GtkEntry struct {
	GtkWidget
	GtkEditable
}

func Entry() *GtkEntry {
	widget := GtkWidget{C.gtk_entry_new()}
	return &GtkEntry{widget, GtkEditable{C.to_GtkEditable(widget.Widget)}}
}
func EntryWithBuffer(buffer *GtkEntryBuffer) *GtkEntry {
	panic_if_version_older_auto(2, 18, 0)
	widget := GtkWidget{C._gtk_entry_new_with_buffer(buffer.EntryBuffer)}
	return &GtkEntry{widget, GtkEditable{C.to_GtkEditable(widget.Widget)}}
}
func (v *GtkEntry) GetBuffer() *GtkEntryBuffer {
	panic_if_version_older_auto(2, 18, 0)
	return &GtkEntryBuffer{C._gtk_entry_get_buffer(C.to_GtkEntry(v.Widget))}
}
func (v *GtkEntry) SetBuffer(buffer *GtkEntryBuffer) {
	panic_if_version_older_auto(2, 18, 0)
	C._gtk_entry_set_buffer(C.to_GtkEntry(v.Widget), buffer.EntryBuffer)
}
func (v *GtkEntry) SetText(text string) {
	ptr := C.CString(text)
	defer C.free_string(ptr)
	C.gtk_entry_set_text(C.to_GtkEntry(v.Widget), C.to_gcharptr(ptr))
}
func (v *GtkEntry) GetText() string {
	return C.GoString(C.to_charptr(C.gtk_entry_get_text(C.to_GtkEntry(v.Widget))))
}
func (v *GtkEntry) GetTextLength() int {
	return int(C.gtk_entry_get_text_length(C.to_GtkEntry(v.Widget)))
}
func (v *GtkEntry) SetVisibility(setting bool) {
	C.gtk_entry_set_visibility(C.to_GtkEntry(v.Widget), bool2gboolean(setting))
}
func (v *GtkEntry) SetInvisibleChar(ch uint8) {
	C.gtk_entry_set_invisible_char(C.to_GtkEntry(v.Widget), C.gunichar(ch))
}
func (v *GtkEntry) UnsetInvisibleChar() {
	C.gtk_entry_unset_invisible_char(C.to_GtkEntry(v.Widget))
}
func (v *GtkEntry) SetMaxLength(i int) {
	C.gtk_entry_set_max_length(C.to_GtkEntry(v.Widget), C.gint(i))
}
func (v *GtkEntry) GetActivatesDefault() bool {
	return gboolean2bool(C.gtk_entry_get_activates_default(C.to_GtkEntry(v.Widget)))
}
func (v *GtkEntry) GetHasFrame() bool {
	return gboolean2bool(C.gtk_entry_get_has_frame(C.to_GtkEntry(v.Widget)))
}

// gtk_entry_get_inner_border

func (v *GtkEntry) GetWidthChars() int {
	return int(C.gtk_entry_get_width_chars(C.to_GtkEntry(v.Widget)))
}
func (v *GtkEntry) SetActivatesDefault(setting bool) {
	C.gtk_entry_set_activates_default(C.to_GtkEntry(v.Widget), bool2gboolean(setting))
}
func (v *GtkEntry) SetHasFrame(setting bool) {
	C.gtk_entry_set_has_frame(C.to_GtkEntry(v.Widget), bool2gboolean(setting))
}

// gtk_entry_set_inner_border

func (v *GtkEntry) SetWidthChars(i int) {
	C.gtk_entry_set_width_chars(C.to_GtkEntry(v.Widget), C.gint(i))
}
func (v *GtkEntry) GetInvisibleChar() uint8 {
	return uint8(C.gtk_entry_get_invisible_char(C.to_GtkEntry(v.Widget)))
}
func (v *GtkEntry) SetAlignment(xalign float64) {
	C.gtk_entry_set_alignment(C.to_GtkEntry(v.Widget), C.gfloat(xalign))
}
func (v *GtkEntry) GetAlignment() float64 {
	return float64(C.gtk_entry_get_alignment(C.to_GtkEntry(v.Widget)))
}

// gtk_entry_set_overwrite_mode
// gtk_entry_get_overwrite_mode
// gtk_entry_get_layout
// gtk_entry_get_layout_offsets
// gtk_entry_layout_index_to_text_index
// gtk_entry_text_index_to_layout_index

func (v *GtkEntry) GetMaxLength() int {
	return int(C.gtk_entry_get_max_length(C.to_GtkEntry(v.Widget)))
}
func (v *GtkEntry) GetVisibility() bool {
	return gboolean2bool(C.gtk_entry_get_visibility(C.to_GtkEntry(v.Widget)))
}
func (v *GtkEntry) SetCompletion(completion *GtkEntryCompletion) {
	C.gtk_entry_set_completion(C.to_GtkEntry(v.Widget), completion.EntryCompletion)
}
func (v *GtkEntry) GetCompletion() *GtkEntryCompletion {
	return &GtkEntryCompletion{C.gtk_entry_get_completion(C.to_GtkEntry(v.Widget))}
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
type GtkEntryBuffer struct {
	EntryBuffer *C.GtkEntryBuffer
}

func EntryBuffer(initialText string) *GtkEntryBuffer {
	panic_if_version_older_auto(2, 18, 0)
	if len(initialText) == 0 {
		return &GtkEntryBuffer{C._gtk_entry_buffer_new(nil, C.gint(-1))}
	}
	ptr := C.CString(initialText)
	defer C.free_string(ptr)
	return &GtkEntryBuffer{
		C._gtk_entry_buffer_new(C.to_gcharptr(ptr), C.gint(len(initialText)))}
}
func (v *GtkEntryBuffer) GetText() string {
	panic_if_version_older_auto(2, 18, 0)
	return C.GoString(C.to_charptr(C._gtk_entry_buffer_get_text(v.EntryBuffer)))
}
func (v *GtkEntryBuffer) SetText(text string) {
	panic_if_version_older_auto(2, 18, 0)
	if len(text) == 0 {
		C._gtk_entry_buffer_set_text(v.EntryBuffer, nil, C.gint(-1))
	}
	ptr := C.CString(text)
	defer C.free_string(ptr)
	C._gtk_entry_buffer_set_text(v.EntryBuffer,
		C.to_gcharptr(ptr), C.gint(len(text)))
}

// gtk_entry_buffer_get_bytes //since 2.18
/*func (v *GtkEntryBuffer) GetBytes() ? {
	panic_if_version_older_auto(2, 18, 0)
	//TODO(any) what is the equivalent type for gsize in go?
	return ?(C._gtk_entry_buffer_get_bytes(v.EntryBuffer))
}*/

func (v *GtkEntryBuffer) GetLength() uint {
	panic_if_version_older_auto(2, 18, 0)
	return uint(C._gtk_entry_buffer_get_length(v.EntryBuffer))
}
func (v *GtkEntryBuffer) GetMaxLength() int {
	panic_if_version_older_auto(2, 18, 0)
	return int(C._gtk_entry_buffer_get_max_length(v.EntryBuffer))
}
func (v *GtkEntryBuffer) SetMaxLength(maxLength int) {
	panic_if_version_older_auto(2, 18, 0)
	C._gtk_entry_buffer_set_max_length(v.EntryBuffer, C.gint(maxLength))
}
func (v *GtkEntryBuffer) InsertText(position uint, text string) uint {
	panic_if_version_older_auto(2, 18, 0)
	ptr := C.CString(text)
	defer C.free_string(ptr)
	return uint(C._gtk_entry_buffer_insert_text(v.EntryBuffer,
		C.guint(position), C.to_gcharptr(ptr), C.gint(len(text))))
}
func (v *GtkEntryBuffer) DeleteText(position uint, nChars int) uint {
	panic_if_version_older_auto(2, 18, 0)
	return uint(C._gtk_entry_buffer_delete_text(v.EntryBuffer,
		C.guint(position), C.gint(nChars)))
}

// gtk_entry_buffer_emit_deleted_text //since 2.18
// gtk_entry_buffer_emit_inserted_text //since 2.18

//-----------------------------------------------------------------------
// GtkEntryCompletion
//-----------------------------------------------------------------------

type GtkEntryCompletion struct {
	EntryCompletion *C.GtkEntryCompletion
}

func EntryCompletion() *GtkEntryCompletion {
	return &GtkEntryCompletion{C.gtk_entry_completion_new()}
}
func (v *GtkEntryCompletion) GetEntry() *GtkWidget {
	return &GtkWidget{C.gtk_entry_completion_get_entry(v.EntryCompletion)}
}
func (v *GtkEntryCompletion) SetModel(model *GtkTreeModel) {
	C.gtk_entry_completion_set_model(v.EntryCompletion, model.TreeModel)
}
func (v *GtkEntryCompletion) GetModel() *GtkTreeModel {
	return &GtkTreeModel{C.gtk_entry_completion_get_model(v.EntryCompletion)}
}

type GtkEntryCompletionMatchFunc func(completion *GtkEntryCompletion, key string, iter *GtkTreeIter, data interface{})

// gtk_entry_completion_set_match_func

func (v *GtkEntryCompletion) SetMinimumKeyLength(length int) {
	C.gtk_entry_completion_set_minimum_key_length(v.EntryCompletion, C.gint(length))
}
func (v *GtkEntryCompletion) GetMinimumKeyLength() int {
	return int(C.gtk_entry_completion_get_minimum_key_length(v.EntryCompletion))
}
func (v *GtkEntryCompletion) Complete() {
	C.gtk_entry_completion_complete(v.EntryCompletion)
}
func (v *GtkEntryCompletion) GetCompletionPrefix() string {
	return C.GoString(C.to_charptr(C.gtk_entry_completion_get_completion_prefix(v.EntryCompletion)))
}
func (v *GtkEntryCompletion) InsertPrefix() {
	C.gtk_entry_completion_insert_prefix(v.EntryCompletion)
}
func (v *GtkEntryCompletion) InsertActionText(index int, text string) {
	ptr := C.CString(text)
	defer C.free_string(ptr)
	C.gtk_entry_completion_insert_action_text(
		v.EntryCompletion, C.gint(index), C.to_gcharptr(ptr))
}
func (v *GtkEntryCompletion) InsertActionMarkup(index int, markup string) {
	ptr := C.CString(markup)
	defer C.free_string(ptr)
	C.gtk_entry_completion_insert_action_markup(
		v.EntryCompletion, C.gint(index), C.to_gcharptr(ptr))
}
func (v *GtkEntryCompletion) DeleteAction(index int) {
	C.gtk_entry_completion_delete_action(v.EntryCompletion, C.gint(index))
}
func (v *GtkEntryCompletion) SetTextColumn(column int) {
	C.gtk_entry_completion_set_text_column(v.EntryCompletion, C.gint(column))
}
func (v *GtkEntryCompletion) GetTextColumn() int {
	return int(C.gtk_entry_completion_get_text_column(v.EntryCompletion))
}
func (v *GtkEntryCompletion) SetInlineCompletion(inlineCompletion bool) {
	C.gtk_entry_completion_set_inline_completion(v.EntryCompletion,
		bool2gboolean(inlineCompletion))
}
func (v *GtkEntryCompletion) GetInlineCompletion() bool {
	return gboolean2bool(C.gtk_entry_completion_get_inline_completion(v.EntryCompletion))
}
func (v *GtkEntryCompletion) SetInlineSelection(inlineSelection bool) {
	C.gtk_entry_completion_set_inline_selection(v.EntryCompletion,
		bool2gboolean(inlineSelection))
}
func (v *GtkEntryCompletion) GetInlineSelection() bool {
	return gboolean2bool(C.gtk_entry_completion_get_inline_selection(v.EntryCompletion))
}
func (v *GtkEntryCompletion) SetPopupCompletion(popupCompletion bool) {
	C.gtk_entry_completion_set_popup_completion(v.EntryCompletion,
		bool2gboolean(popupCompletion))
}
func (v *GtkEntryCompletion) GetPopupCompletion() bool {
	return gboolean2bool(C.gtk_entry_completion_get_popup_completion(v.EntryCompletion))
}
func (v *GtkEntryCompletion) SetPopupSetWidth(popupSetWidth bool) {
	C.gtk_entry_completion_set_popup_set_width(v.EntryCompletion,
		bool2gboolean(popupSetWidth))
}
func (v *GtkEntryCompletion) GetPopupSetWidth() bool {
	return gboolean2bool(C.gtk_entry_completion_get_popup_set_width(v.EntryCompletion))
}
func (v *GtkEntryCompletion) SetPopupSingleMatch(popupSingleMatch bool) {
	C.gtk_entry_completion_set_popup_single_match(v.EntryCompletion,
		bool2gboolean(popupSingleMatch))
}
func (v *GtkEntryCompletion) GetPopupSingleMatch() bool {
	return gboolean2bool(C.gtk_entry_completion_get_popup_single_match(v.EntryCompletion))
}

//-----------------------------------------------------------------------
// GtkHScale
//-----------------------------------------------------------------------
func HScale(adjustment *GtkAdjustment) *GtkScale {
	return &GtkScale{GtkRange{GtkWidget{
		C.gtk_hscale_new(adjustment.Adjustment)}}}
}
func HScaleWithRange(min float64, max float64, step float64) *GtkScale {
	return &GtkScale{GtkRange{GtkWidget{
		C.gtk_hscale_new_with_range(C.gdouble(min), C.gdouble(max), C.gdouble(step))}}}
}

//-----------------------------------------------------------------------
// GtkVScale
//-----------------------------------------------------------------------
func VScale(adjustment *GtkAdjustment) *GtkScale {
	return &GtkScale{GtkRange{GtkWidget{
		C.gtk_vscale_new(adjustment.Adjustment)}}}
}
func VScaleWithRange(min float64, max float64, step float64) *GtkScale {
	return &GtkScale{GtkRange{GtkWidget{
		C.gtk_vscale_new_with_range(C.gdouble(min), C.gdouble(max), C.gdouble(step))}}}
}

//-----------------------------------------------------------------------
// GtkSpinButton
//-----------------------------------------------------------------------

// gtk_spin_button_configure
// gtk_spin_button_new
// gtk_spin_button_new_with_range
// gtk_spin_button_set_adjustment
// gtk_spin_button_get_adjustment
// gtk_spin_button_set_digits
// gtk_spin_button_set_increments
// gtk_spin_button_set_range
// gtk_spin_button_get_value_as_float
// gtk_spin_button_get_value_as_int
// gtk_spin_button_set_value
// gtk_spin_button_set_update_policy
// gtk_spin_button_set_numeric
// gtk_spin_button_spin
// gtk_spin_button_set_wrap
// gtk_spin_button_set_snap_to_ticks
// gtk_spin_button_update
// gtk_spin_button_get_digits
// gtk_spin_button_get_increments
// gtk_spin_button_get_numeric
// gtk_spin_button_get_range
// gtk_spin_button_get_snap_to_ticks
// gtk_spin_button_get_update_policy
// gtk_spin_button_get_value
// gtk_spin_button_get_wrap

//-----------------------------------------------------------------------
// GtkEditable
//-----------------------------------------------------------------------
type GtkEditable struct {
	Editable *C.GtkEditable
}

func (v *GtkEditable) SelectRegion(startPos int, endPos int) {
	C.gtk_editable_select_region(v.Editable, C.gint(startPos), C.gint(endPos))
}
func (v *GtkEditable) GetSelectionBounds() (isSelected bool,
	startPos int, endPos int) {
	var s, e C.gint
	return gboolean2bool(C.gtk_editable_get_selection_bounds(v.Editable, &s, &e)),
		int(s), int(e)
}
func (v *GtkEditable) InsertText(newText string, position int) int {
	ptr := C.CString(newText)
	defer C.free_string(ptr)
	gpos := (C.gint)(position)
	C.gtk_editable_insert_text(v.Editable, C.to_gcharptr(ptr),
		C.gint(len(newText)), &gpos)
	return int(gpos)
}
func (v *GtkEditable) DeleteText(startPos int, endPos int) {
	C.gtk_editable_delete_text(v.Editable, C.gint(startPos), C.gint(endPos))
}
func (v *GtkEditable) GetChars(startPos int, endPos int) string {
	return C.GoString(C.to_charptr(C.gtk_editable_get_chars(v.Editable,
		C.gint(startPos), C.gint(endPos))))
}
func (v *GtkEditable) CutClipboard() {
	C.gtk_editable_cut_clipboard(v.Editable)
}
func (v *GtkEditable) CopyClipboard() {
	C.gtk_editable_copy_clipboard(v.Editable)
}
func (v *GtkEditable) PasteClipboard() {
	C.gtk_editable_paste_clipboard(v.Editable)
}
func (v *GtkEditable) DeleteSelection() {
	C.gtk_editable_delete_selection(v.Editable)
}
func (v *GtkEditable) SetPosition(position int) {
	C.gtk_editable_set_position(v.Editable, C.gint(position))
}
func (v *GtkEditable) GetPosition() int {
	return int(C.gtk_editable_get_position(v.Editable))
}
func (v *GtkEditable) SetEditable(isEditable bool) {
	C.gtk_editable_set_editable(v.Editable, bool2gboolean(isEditable))
}
func (v *GtkEditable) GetEditable() bool {
	return gboolean2bool(C.gtk_editable_get_editable(v.Editable))
}

//-----------------------------------------------------------------------
// GtkTextIter
//-----------------------------------------------------------------------
type GtkTextIter struct {
	TextIter C.GtkTextIter
}

type GtkTextSearchFlags int

const (
	GTK_TEXT_SEARCH_VISIBLE_ONLY     GtkTextSearchFlags = 1 << 0
	GTK_TEXT_SEARCH_TEXT_ONLY        GtkTextSearchFlags = 1 << 1
	GTK_TEXT_SEARCH_CASE_INSENSITIVE GtkTextSearchFlags = 1 << 2
)

func (v *GtkTextIter) GetBuffer() *GtkTextBuffer {
	return newTextBuffer(C.gtk_text_iter_get_buffer(&v.TextIter))
}
func (v *GtkTextIter) Copy() *GtkTextIter {
	return &GtkTextIter{
		*C.gtk_text_iter_copy(&v.TextIter)}
}
func (v *GtkTextIter) Free() {
	C.gtk_text_iter_free(&v.TextIter)
}
func (v *GtkTextIter) GetOffset() int {
	return int(C.gtk_text_iter_get_offset(&v.TextIter))
}
func (v *GtkTextIter) GetLine() int {
	return int(C.gtk_text_iter_get_line(&v.TextIter))
}
func (v *GtkTextIter) GetLineOffset() int {
	return int(C.gtk_text_iter_get_line_offset(&v.TextIter))
}
func (v *GtkTextIter) GetLineIndex() int {
	return int(C.gtk_text_iter_get_line_index(&v.TextIter))
}
func (v *GtkTextIter) GetVisibleLineIndex() int {
	return int(C.gtk_text_iter_get_visible_line_index(&v.TextIter))
}
func (v *GtkTextIter) GetVisibleLineOffset() int {
	return int(C.gtk_text_iter_get_visible_line_offset(&v.TextIter))
}
func (v *GtkTextIter) GetChar() int {
	return int(C.gtk_text_iter_get_char(&v.TextIter))
}
func (v *GtkTextIter) GetSlice(end *GtkTextIter) string {
	pchar := C.to_charptr(C.gtk_text_iter_get_slice(&v.TextIter, &end.TextIter))
	defer C.free(unsafe.Pointer(pchar))
	return C.GoString(pchar)
}
func (v *GtkTextIter) GetText(end *GtkTextIter) string {
	pchar := C.to_charptr(C.gtk_text_iter_get_text(&v.TextIter, &end.TextIter))
	defer C.free(unsafe.Pointer(pchar))
	return C.GoString(pchar)
}
func (v *GtkTextIter) GetVisibleSlice(end *GtkTextIter) string {
	return C.GoString(C.to_charptr(C.gtk_text_iter_get_visible_slice(&v.TextIter, &end.TextIter)))
}
func (v *GtkTextIter) GetVisibleText(end *GtkTextIter) string {
	return C.GoString(C.to_charptr(C.gtk_text_iter_get_visible_text(&v.TextIter, &end.TextIter)))
}

// gtk_text_iter_get_pixbuf

func (v *GtkTextIter) GetMarks() *glib.SList {
	return glib.SListFromNative(unsafe.Pointer(C.gtk_text_iter_get_marks(&v.TextIter)))
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

func (v *GtkTextIter) ForwardChar() bool {
	return gboolean2bool(C.gtk_text_iter_forward_char(&v.TextIter))
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

func (v *GtkTextIter) ForwardSearch(str string, flags GtkTextSearchFlags, start *GtkTextIter, end *GtkTextIter, limit *GtkTextIter) bool {
	cstr := C.CString(str)
	defer C.free_string(cstr)
	return gboolean2bool(C.gtk_text_iter_forward_search(&v.TextIter,
		C.to_gcharptr(cstr), C.GtkTextSearchFlags(flags), &start.TextIter,
		&end.TextIter, &limit.TextIter))
}
func (v *GtkTextIter) BackwardSearch(str string, flags GtkTextSearchFlags, start *GtkTextIter, end *GtkTextIter, limit *GtkTextIter) bool {
	cstr := C.CString(str)
	defer C.free_string(cstr)
	return gboolean2bool(C.gtk_text_iter_backward_search(&v.TextIter,
		C.to_gcharptr(cstr), C.GtkTextSearchFlags(flags), &start.TextIter,
		&end.TextIter, &limit.TextIter))
}

// gtk_text_iter_equal
// gtk_text_iter_compare
// gtk_text_iter_in_range
// gtk_text_iter_order

func (v *GtkTextIter) Assign(iter *GtkTextIter) {
	C._gtk_text_iter_assign(&v.TextIter, &iter.TextIter)
}

//-----------------------------------------------------------------------
// GtkTextMark
//-----------------------------------------------------------------------
type GtkTextMark struct {
	TextMark *C.GtkTextMark
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

type GtkTextBuffer struct {
	TextBuffer *C.GtkTextBuffer
	*glib.GObject
}

func newTextBuffer(buffer *C.GtkTextBuffer) *GtkTextBuffer {
	return &GtkTextBuffer{
		TextBuffer: buffer,
		GObject:    glib.ObjectFromNative(unsafe.Pointer(buffer)),
	}
}
func (v *GtkTextBuffer) GetNativeBuffer() unsafe.Pointer {
	return unsafe.Pointer(v.TextBuffer)
}
func TextBufferFromPointer(v unsafe.Pointer) GtkTextBuffer {
	return *newTextBuffer(C.to_GtkTextBuffer(v))
}
func TextBuffer(tagtable *GtkTextTagTable) *GtkTextBuffer {
	return newTextBuffer(C.gtk_text_buffer_new(tagtable.TextTagTable))
}
func (v *GtkTextBuffer) GetLineCount() int {
	return int(C.gtk_text_buffer_get_line_count(v.TextBuffer))
}
func (v *GtkTextBuffer) GetCharCount() int {
	return int(C.gtk_text_buffer_get_char_count(v.TextBuffer))
}
func (v *GtkTextBuffer) GetTagTable() *GtkTextTagTable {
	return &GtkTextTagTable{
		C.gtk_text_buffer_get_tag_table(v.TextBuffer)}
}
func (v *GtkTextBuffer) Insert(iter *GtkTextIter, text string) {
	ptr := C.CString(text)
	defer C.free_string(ptr)
	l := C.strlen(ptr)
	C.gtk_text_buffer_insert(v.TextBuffer, &iter.TextIter, C.to_gcharptr(ptr), C.gint(l))
}
func (v *GtkTextBuffer) InsertAtCursor(text string) {
	ptr := C.CString(text)
	defer C.free_string(ptr)
	l := C.strlen(ptr)
	C.gtk_text_buffer_insert_at_cursor(v.TextBuffer, C.to_gcharptr(ptr), C.gint(l))
}
func (v *GtkTextBuffer) InsertInteractive(iter *GtkTextIter, text string, default_editable bool) bool {
	ptr := C.CString(text)
	defer C.free_string(ptr)
	l := C.strlen(ptr)
	return gboolean2bool(C.gtk_text_buffer_insert_interactive(v.TextBuffer, &iter.TextIter, C.to_gcharptr(ptr), C.gint(l), bool2gboolean(default_editable)))
}
func (v *GtkTextBuffer) InsertInteractiveAtCursor(text string, default_editable bool) bool {
	ptr := C.CString(text)
	defer C.free_string(ptr)
	l := C.strlen(ptr)
	return gboolean2bool(C.gtk_text_buffer_insert_interactive_at_cursor(v.TextBuffer, C.to_gcharptr(ptr), C.gint(l), bool2gboolean(default_editable)))
}
func (v *GtkTextBuffer) InsertRange(iter *GtkTextIter, start *GtkTextIter, end *GtkTextIter) {
	C.gtk_text_buffer_insert_range(v.TextBuffer, &iter.TextIter, &start.TextIter, &end.TextIter)
}
func (v *GtkTextBuffer) InsertRangeInteractive(iter *GtkTextIter, start *GtkTextIter, end *GtkTextIter, default_editable bool) bool {
	return gboolean2bool(C.gtk_text_buffer_insert_range_interactive(v.TextBuffer, &iter.TextIter, &start.TextIter, &end.TextIter, bool2gboolean(default_editable)))
}
func (v *GtkTextBuffer) InsertWithTag(iter *GtkTextIter, text string, tag *GtkTextTag) {
	ptr := C.CString(text)
	defer C.free_string(ptr)
	l := C.strlen(ptr)
	C._gtk_text_buffer_insert_with_tag(v.TextBuffer, &iter.TextIter, C.to_gcharptr(ptr), C.gint(l), tag.TextTag)
}

//func (v GtkTextBuffer) InsertWithTags(iter *GtkTextIter, start *GtkTextIter, end *GtkTextIter, default_editable bool) bool {
//	return gboolean2bool(C._gtk_text_buffer_insert_range_interactive(v.TextBuffer, &iter.TextIter, &start.TextIter, &end.TextIter, bool2gboolean(default_editable)));
//}
// gtk_text_buffer_insert_with_tags_by_name

func (v *GtkTextBuffer) Delete(start *GtkTextIter, end *GtkTextIter) {
	C.gtk_text_buffer_delete(v.TextBuffer, &start.TextIter, &end.TextIter)
}
func (v *GtkTextBuffer) DeleteInteractive(start *GtkTextIter, end *GtkTextIter, default_editable bool) bool {
	return gboolean2bool(C.gtk_text_buffer_delete_interactive(v.TextBuffer, &start.TextIter, &end.TextIter, bool2gboolean(default_editable)))
}
func (v *GtkTextBuffer) Backspace(iter *GtkTextIter, interactive bool, default_editable bool) bool {
	return gboolean2bool(C.gtk_text_buffer_backspace(v.TextBuffer, &iter.TextIter, bool2gboolean(interactive), bool2gboolean(default_editable)))
}
func (v *GtkTextBuffer) SetText(text string) {
	ptr := C.CString(text)
	defer C.free_string(ptr)
	l := C.strlen(ptr)
	C.gtk_text_buffer_set_text(v.TextBuffer, C.to_gcharptr(ptr), C.gint(l))
}
func (v *GtkTextBuffer) GetText(start *GtkTextIter, end *GtkTextIter, include_hidden_chars bool) string {
	pchar := C.to_charptr(C.gtk_text_buffer_get_text(v.TextBuffer, &start.TextIter, &end.TextIter, bool2gboolean(include_hidden_chars)))
	defer C.free(unsafe.Pointer(pchar))
	return C.GoString(pchar)
}
func (v *GtkTextBuffer) GetSlice(start *GtkTextIter, end *GtkTextIter, include_hidden_chars bool) string {
	pchar := C.to_charptr(C.gtk_text_buffer_get_slice(v.TextBuffer, &start.TextIter, &end.TextIter, bool2gboolean(include_hidden_chars)))
	defer C.free(unsafe.Pointer(pchar))
	return C.GoString(pchar)
}
func (v *GtkTextBuffer) InsertPixbuf(iter *GtkTextIter, pixbuf *gdkpixbuf.GdkPixbuf) {
	C.gtk_text_buffer_insert_pixbuf(v.TextBuffer, &iter.TextIter, pixbuf.Pixbuf)
}

// gtk_text_buffer_insert_child_anchor
// gtk_text_buffer_create_child_anchor

func (v *GtkTextBuffer) CreateMark(mark_name string, where *GtkTextIter, left_gravity bool) *GtkTextMark {
	ptr := C.CString(mark_name)
	defer C.free_string(ptr)
	return &GtkTextMark{C.gtk_text_buffer_create_mark(v.TextBuffer, C.to_gcharptr(ptr), &where.TextIter, bool2gboolean(left_gravity))}
}
func (v *GtkTextBuffer) MoveMark(mark *GtkTextMark, where *GtkTextIter) {
	C.gtk_text_buffer_move_mark(v.TextBuffer, mark.TextMark, &where.TextIter)
}
func (v *GtkTextBuffer) MoveMarkByName(name string, where *GtkTextIter) {
	ptr := C.CString(name)
	defer C.free_string(ptr)
	C.gtk_text_buffer_move_mark_by_name(v.TextBuffer, C.to_gcharptr(ptr), &where.TextIter)
}
func (v *GtkTextBuffer) AddMark(mark *GtkTextMark, where *GtkTextIter) {
	C.gtk_text_buffer_add_mark(v.TextBuffer, mark.TextMark, &where.TextIter)
}
func (v *GtkTextBuffer) DeleteMark(mark *GtkTextMark) {
	C.gtk_text_buffer_delete_mark(v.TextBuffer, mark.TextMark)
}
func (v *GtkTextBuffer) DeleteMarkByName(name string) {
	ptr := C.CString(name)
	defer C.free_string(ptr)
	C.gtk_text_buffer_delete_mark_by_name(v.TextBuffer, C.to_gcharptr(ptr))
}
func (v *GtkTextBuffer) GetMark(name string) *GtkTextMark {
	ptr := C.CString(name)
	defer C.free_string(ptr)
	return &GtkTextMark{
		C.gtk_text_buffer_get_mark(v.TextBuffer, C.to_gcharptr(ptr))}
}
func (v *GtkTextBuffer) GetInsert() *GtkTextMark {
	return &GtkTextMark{
		C.gtk_text_buffer_get_insert(v.TextBuffer)}
}
func (v *GtkTextBuffer) GetSelectionBound() *GtkTextMark {
	return &GtkTextMark{
		C.gtk_text_buffer_get_selection_bound(v.TextBuffer)}
}
func (v *GtkTextBuffer) GetHasSelection() bool {
	return gboolean2bool(C.gtk_text_buffer_get_has_selection(v.TextBuffer))
}
func (v *GtkTextBuffer) PlaceCursor(where *GtkTextIter) {
	C.gtk_text_buffer_place_cursor(v.TextBuffer, &where.TextIter)
}
func (v *GtkTextBuffer) SelectRange(ins *GtkTextIter, bound *GtkTextIter) {
	C.gtk_text_buffer_select_range(v.TextBuffer, &ins.TextIter, &bound.TextIter)
}
func (v *GtkTextBuffer) ApplyTag(tag *GtkTextTag, start *GtkTextIter, end *GtkTextIter) {
	C.gtk_text_buffer_apply_tag(v.TextBuffer, tag.TextTag, &start.TextIter, &end.TextIter)
}
func (v *GtkTextBuffer) RemoveTag(tag *GtkTextTag, start *GtkTextIter, end *GtkTextIter) {
	C.gtk_text_buffer_remove_tag(v.TextBuffer, tag.TextTag, &start.TextIter, &end.TextIter)
}
func (v *GtkTextBuffer) ApplyTagByName(name string, start *GtkTextIter, end *GtkTextIter) {
	ptr := C.CString(name)
	defer C.free_string(ptr)
	C.gtk_text_buffer_apply_tag_by_name(v.TextBuffer, C.to_gcharptr(ptr), &start.TextIter, &end.TextIter)
}
func (v *GtkTextBuffer) RemoveTagByName(name string, start *GtkTextIter, end *GtkTextIter) {
	ptr := C.CString(name)
	defer C.free_string(ptr)
	C.gtk_text_buffer_remove_tag_by_name(v.TextBuffer, C.to_gcharptr(ptr), &start.TextIter, &end.TextIter)
}
func (v *GtkTextBuffer) RemoveAllTags(start *GtkTextIter, end *GtkTextIter) {
	C.gtk_text_buffer_remove_all_tags(v.TextBuffer, &start.TextIter, &end.TextIter)
}
func (v *GtkTextBuffer) CreateTag(tag_name string, props map[string]string) *GtkTextTag {
	ptr := C.CString(tag_name)
	defer C.free_string(ptr)
	tag := C._gtk_text_buffer_create_tag(v.TextBuffer, C.to_gcharptr(ptr))
	for prop, val := range props {
		pprop := C.CString(prop)
		pval := C.CString(val)
		C._apply_property(unsafe.Pointer(tag), C.to_gcharptr(pprop), C.to_gcharptr(pval))
		C.free_string(pprop)
		C.free_string(pval)
	}
	return &GtkTextTag{tag}
}
func (v *GtkTextBuffer) GetIterAtLineOffset(iter *GtkTextIter, line_number int, char_offset int) {
	C.gtk_text_buffer_get_iter_at_line_offset(v.TextBuffer, &iter.TextIter, C.gint(line_number), C.gint(char_offset))
}
func (v *GtkTextBuffer) GetIterAtOffset(iter *GtkTextIter, char_offset int) {
	C.gtk_text_buffer_get_iter_at_offset(v.TextBuffer, &iter.TextIter, C.gint(char_offset))
}
func (v *GtkTextBuffer) GetIterAtLine(iter *GtkTextIter, line_number int) {
	C.gtk_text_buffer_get_iter_at_line(v.TextBuffer, &iter.TextIter, C.gint(line_number))
}
func (v *GtkTextBuffer) GetIterAtLineIndex(iter *GtkTextIter, line_number int, byte_index int) {
	C.gtk_text_buffer_get_iter_at_line_index(v.TextBuffer, &iter.TextIter, C.gint(line_number), C.gint(byte_index))
}
func (v *GtkTextBuffer) GetIterAtMark(iter *GtkTextIter, mark *GtkTextMark) {
	C.gtk_text_buffer_get_iter_at_mark(v.TextBuffer, &iter.TextIter, mark.TextMark)
}
func (v *GtkTextBuffer) GetIterAtChildAnchor(iter *GtkTextIter, anchor *GtkTextChildAnchor) {
	C.gtk_text_buffer_get_iter_at_child_anchor(v.TextBuffer, &iter.TextIter, anchor.TextChildAnchor)
}
func (v *GtkTextBuffer) GetStartIter(iter *GtkTextIter) {
	C.gtk_text_buffer_get_start_iter(v.TextBuffer, &iter.TextIter)
}
func (v *GtkTextBuffer) GetEndIter(iter *GtkTextIter) {
	C.gtk_text_buffer_get_end_iter(v.TextBuffer, &iter.TextIter)
}
func (v *GtkTextBuffer) GetBounds(start *GtkTextIter, end *GtkTextIter) {
	C.gtk_text_buffer_get_bounds(v.TextBuffer, &start.TextIter, &end.TextIter)
}
func (v *GtkTextBuffer) GetModified() bool {
	return gboolean2bool(C.gtk_text_buffer_get_modified(v.TextBuffer))
}
func (v *GtkTextBuffer) SetModified(setting bool) {
	C.gtk_text_buffer_set_modified(v.TextBuffer, bool2gboolean(setting))
}
func (v *GtkTextBuffer) DeleteSelection(interactive bool, default_editable bool) {
	C.gtk_text_buffer_delete_selection(v.TextBuffer, bool2gboolean(interactive), bool2gboolean(default_editable))
}

// gtk_text_buffer_paste_clipboard
// gtk_text_buffer_copy_clipboard
// gtk_text_buffer_cut_clipboard

func (v *GtkTextBuffer) GetSelectionBounds(be, en *GtkTextIter) bool {
	return gboolean2bool(C.gtk_text_buffer_get_selection_bounds(v.TextBuffer, &be.TextIter, &en.TextIter))
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
type GtkTextTag struct {
	TextTag *C.GtkTextTag
}

func TextTag(name string) *GtkTextTag {
	ptr := C.CString(name)
	defer C.free_string(ptr)
	return &GtkTextTag{
		C.gtk_text_tag_new(C.to_gcharptr(ptr))}
}
func (v *GtkTextTag) SetPriority(priority int) {
	C.gtk_text_tag_set_priority(v.TextTag, C.gint(priority))
}
func (v *GtkTextTag) GetPriority() int {
	return int(C.gtk_text_tag_get_priority(v.TextTag))
}

// gtk_text_tag_event

//-----------------------------------------------------------------------
// GtkTextAttributes
//-----------------------------------------------------------------------
type GtkTextAttributes struct {
	TextAttributes *C.GtkTextAttributes
}

func TextAttributes() *GtkTextAttributes {
	return &GtkTextAttributes{C.gtk_text_attributes_new()}
}
func (v *GtkTextAttributes) Copy() *GtkTextAttributes {
	return &GtkTextAttributes{C.gtk_text_attributes_copy(v.TextAttributes)}
}
func (v *GtkTextAttributes) CopyValues(values *GtkTextAttributes) {
	C.gtk_text_attributes_copy_values(v.TextAttributes, values.TextAttributes)
}
func (v *GtkTextAttributes) Unref() {
	C.gtk_text_attributes_unref(v.TextAttributes)
}
func (v *GtkTextAttributes) Ref() *GtkTextAttributes {
	return &GtkTextAttributes{C.gtk_text_attributes_ref(v.TextAttributes)}
}

//-----------------------------------------------------------------------
// GtkTextTagTable
//-----------------------------------------------------------------------
type GtkTextTagTable struct {
	TextTagTable *C.GtkTextTagTable
}

func TextTagTable() *GtkTextTagTable {
	return &GtkTextTagTable{C.gtk_text_tag_table_new()}
}
func (v *GtkTextTagTable) Add(tag *GtkTextTag) {
	C.gtk_text_tag_table_add(v.TextTagTable, tag.TextTag)
}
func (v *GtkTextTagTable) Remove(tag *GtkTextTag) {
	C.gtk_text_tag_table_remove(v.TextTagTable, tag.TextTag)
}
func (v *GtkTextTagTable) Lookup(name string) *GtkTextTag {
	ptr := C.CString(name)
	defer C.free_string(ptr)
	return &GtkTextTag{
		C.gtk_text_tag_table_lookup(v.TextTagTable, C.to_gcharptr(ptr))}
}

// gtk_text_tag_table_foreach

func (v *GtkTextTagTable) GetSize() int {
	return int(C.gtk_text_tag_table_get_size(v.TextTagTable))
}

//-----------------------------------------------------------------------
// GtkTextView
//-----------------------------------------------------------------------
type GtkWrapMode int

const (
	GTK_WRAP_NONE      GtkWrapMode = 0
	GTK_WRAP_CHAR      GtkWrapMode = 1
	GTK_WRAP_WORD      GtkWrapMode = 2
	GTK_WRAP_WORD_CHAR GtkWrapMode = 3
)

type GtkTextChildAnchor struct {
	TextChildAnchor *C.GtkTextChildAnchor
}

// gtk_text_child_anchor_new
// gtk_text_child_anchor_get_widgets
// gtk_text_child_anchor_get_deleted
type GtkTextView struct {
	GtkContainer
}

func TextView() *GtkTextView {
	return &GtkTextView{GtkContainer{GtkWidget{C.gtk_text_view_new()}}}
}
func TextViewWithBuffer(b GtkTextBuffer) *GtkTextView {
	return &GtkTextView{GtkContainer{GtkWidget{
		C.gtk_text_view_new_with_buffer(b.TextBuffer)}}}
}
func (v *GtkTextView) SetBuffer(b TextBufferLike) {
	C.gtk_text_view_set_buffer(C.to_GtkTextView(v.Widget), C.to_GtkTextBuffer(b.GetNativeBuffer()))
}
func (v *GtkTextView) GetBuffer() *GtkTextBuffer {
	return newTextBuffer(C.gtk_text_view_get_buffer(C.to_GtkTextView(v.Widget)))
}
func (v *GtkTextView) ScrollToMark(mark *GtkTextMark, wm float64, ua bool, xa float64, ya float64) {
	C.gtk_text_view_scroll_to_mark(C.to_GtkTextView(v.Widget),
		mark.TextMark, C.gdouble(wm), bool2gboolean(ua), C.gdouble(xa), C.gdouble(ya))
}
func (v *GtkTextView) ScrollToIter(iter *GtkTextIter, wm float64, ua bool, xa float64, ya float64) bool {
	return gboolean2bool(C.gtk_text_view_scroll_to_iter(C.to_GtkTextView(v.Widget),
		&iter.TextIter, C.gdouble(wm), bool2gboolean(ua), C.gdouble(xa), C.gdouble(ya)))
}

// void gtk_text_view_scroll_mark_onscreen(GtkTextView* text_view, GtkTextMark* mark);
// gboolean gtk_text_view_move_mark_onscreen(GtkTextView* text_view, GtkTextMark* mark);
// gboolean gtk_text_view_place_cursor_onscreen(GtkTextView* text_view);
// gtk_text_view_get_visible_rect
// void gtk_text_view_get_iter_location(GtkTextView* text_view, const GtkTextIter* iter, GdkRectangle* location);
// void gtk_text_view_get_line_at_y(GtkTextView* text_view, GtkTextIter* target_iter, gint y, gint* line_top);

func (v *GtkTextView) GetLineYrange(iter *GtkTextIter, y *int, h *int) {
	var yy, hh C.gint
	C.gtk_text_view_get_line_yrange(C.to_GtkTextView(v.Widget), &iter.TextIter, &yy, &hh)
	*y = int(yy)
	*h = int(hh)
}

// void gtk_text_view_get_iter_at_location(GtkTextView* text_view, GtkTextIter* iter, gint x, gint y);

func (v *GtkTextView) GetIterAtPosition(iter *GtkTextIter, trailing *int, x int, y int) {
	if nil != trailing {
		var tt C.gint
		C.gtk_text_view_get_iter_at_position(C.to_GtkTextView(v.Widget), &iter.TextIter, &tt, C.gint(x), C.gint(y))
		*trailing = int(tt)
	} else {
		C.gtk_text_view_get_iter_at_position(C.to_GtkTextView(v.Widget), &iter.TextIter, nil, C.gint(x), C.gint(y))
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

func (v *GtkTextView) SetWrapMode(mode GtkWrapMode) {
	C.gtk_text_view_set_wrap_mode(C.to_GtkTextView(v.Widget), C.GtkWrapMode(mode))
}
func (v *GtkTextView) GetWrapMode() GtkWrapMode {
	return GtkWrapMode(C.gtk_text_view_get_wrap_mode(C.to_GtkTextView(v.Widget)))
}
func (v *GtkTextView) SetEditable(setting bool) {
	C.gtk_text_view_set_editable(C.to_GtkTextView(v.Widget), bool2gboolean(setting))
}
func (v *GtkTextView) GetEditable() bool {
	return gboolean2bool(C.gtk_text_view_get_editable(C.to_GtkTextView(v.Widget)))
}
func (v *GtkTextView) SetCursorVisible(setting bool) {
	C.gtk_text_view_set_cursor_visible(C.to_GtkTextView(v.Widget), bool2gboolean(setting))
}
func (v *GtkTextView) GetCursorVisible() bool {
	return gboolean2bool(C.gtk_text_view_get_cursor_visible(C.to_GtkTextView(v.Widget)))
}
func (v *GtkTextView) SetOverwrite(overwrite bool) {
	C.gtk_text_view_set_overwrite(C.to_GtkTextView(v.Widget), bool2gboolean(overwrite))
}
func (v *GtkTextView) GetOverwrite() bool {
	return gboolean2bool(C.gtk_text_view_get_overwrite(C.to_GtkTextView(v.Widget)))
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

func (v *GtkTextView) SetAcceptsTab(accepts_tab bool) {
	C.gtk_text_view_set_accepts_tab(C.to_GtkTextView(v.Widget), bool2gboolean(accepts_tab))
}
func (v *GtkTextView) GetAcceptsTab() bool {
	return gboolean2bool(C.gtk_text_view_get_accepts_tab(C.to_GtkTextView(v.Widget)))
}

// GtkTextAttributes* gtk_text_view_get_default_attributes(GtkTextView* text_view);
// gtk_text_view_im_context_filter_keypress //since 2.22
// gtk_text_view_reset_im_context //since 2.22
// gtk_text_view_get_hadjustment //since 2.22
// gtk_text_view_get_vadjustment //since 2.22

//-----------------------------------------------------------------------
// GtkTreePath
//-----------------------------------------------------------------------
type GtkTreePath struct {
	TreePath *C.GtkTreePath
}

func TreePath() *GtkTreePath {
	return &GtkTreePath{C.gtk_tree_path_new()}
}
func TreePathFromString(path string) *GtkTreePath {
	ptr := C.CString(path)
	defer C.free_string(ptr)
	return &GtkTreePath{
		C.gtk_tree_path_new_from_string(C.to_gcharptr(ptr))}
}

// gtk_tree_path_new_from_indices

func (v *GtkTreePath) String() string {
	return C.GoString(C.to_charptr(C.gtk_tree_path_to_string(v.TreePath)))
}
func TreePathNewFirst() *GtkTreePath {
	return &GtkTreePath{C.gtk_tree_path_new_first()}
}
func (v *GtkTreePath) AppendIndex(index int) {
	C.gtk_tree_path_append_index(v.TreePath, C.gint(index))
}
func (v *GtkTreePath) PrependIndex(index int) {
	C.gtk_tree_path_prepend_index(v.TreePath, C.gint(index))
}
func (v *GtkTreePath) GetDepth() int {
	return int(C.gtk_tree_path_get_depth(v.TreePath))
}

// gtk_tree_path_get_indices
// gtk_tree_path_get_indices_with_depth //since 2.22

func (v *GtkTreePath) Free() {
	C.gtk_tree_path_free(v.TreePath)
}
func (v *GtkTreePath) Copy() *GtkTreePath {
	return &GtkTreePath{C.gtk_tree_path_copy(v.TreePath)}
}
func (v *GtkTreePath) Compare(w GtkTreePath) int {
	return int(C.gtk_tree_path_compare(v.TreePath, w.TreePath))
}
func (v *GtkTreePath) Next() {
	C.gtk_tree_path_next(v.TreePath)
}
func (v *GtkTreePath) Prev() bool {
	return gboolean2bool(C.gtk_tree_path_prev(v.TreePath))
}
func (v *GtkTreePath) Up() bool {
	return gboolean2bool(C.gtk_tree_path_up(v.TreePath))
}
func (v *GtkTreePath) Down() {
	C.gtk_tree_path_down(v.TreePath)
}
func (v *GtkTreePath) IsAncestor(descendant GtkTreePath) bool {
	return gboolean2bool(C.gtk_tree_path_is_ancestor(v.TreePath, descendant.TreePath))
}
func (v *GtkTreePath) IsDescendant(ancestor GtkTreePath) bool {
	return gboolean2bool(C.gtk_tree_path_is_descendant(v.TreePath, ancestor.TreePath))
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
type GtkTreeIter struct {
	TreeIter C.GtkTreeIter
}

func (v *GtkTreeIter) Assign(to *GtkTreeIter) {
	C._gtk_tree_iter_assign(unsafe.Pointer(&v.TreeIter), unsafe.Pointer(&to.TreeIter))
}

//-----------------------------------------------------------------------
// GtkTreeModel
//-----------------------------------------------------------------------
type GtkTreeModelFlags int

const (
	GTK_TREE_MODEL_ITERS_PERSIST GtkTreeModelFlags = 1 << 0
	GTK_TREE_MODEL_LIST_ONLY     GtkTreeModelFlags = 1 << 1
)

type GtkTreeModelLike interface {
	cTreeModel() *C.GtkTreeModel
}

type GtkTreeModel struct {
	TreeModel *C.GtkTreeModel
}

func (v GtkTreeModel) cTreeModel() *C.GtkTreeModel {
	return v.TreeModel
}

func (v *GtkTreeModel) GetFlags() GtkTreeModelFlags {
	return GtkTreeModelFlags(C.gtk_tree_model_get_flags(v.TreeModel))
}
func (v *GtkTreeModel) GetNColumns() int {
	return int(C.gtk_tree_model_get_n_columns(v.TreeModel))
}

// gtk_tree_model_get_column_type
func (v *GtkTreeModel) GetIter(iter *GtkTreeIter, path *GtkTreePath) bool {
	return gboolean2bool(C.gtk_tree_model_get_iter(v.TreeModel, &iter.TreeIter, path.TreePath))
}
func (v *GtkTreeModel) GetIterFromString(iter *GtkTreeIter, path_string string) bool {
	ptr := C.CString(path_string)
	defer C.free_string(ptr)
	ret := gboolean2bool(C.gtk_tree_model_get_iter_from_string(v.TreeModel, &iter.TreeIter, C.to_gcharptr(ptr)))
	return ret
}
func (v *GtkTreeModel) GetIterFirst(iter *GtkTreeIter) bool {
	return gboolean2bool(C.gtk_tree_model_get_iter_first(v.TreeModel, &iter.TreeIter))
}
func (v *GtkTreeModel) GetPath(iter *GtkTreeIter) *GtkTreePath {
	return &GtkTreePath{C._gtk_tree_model_get_path(v.TreeModel, &iter.TreeIter)}
}
func (v *GtkTreeModel) GetValue(iter *GtkTreeIter, col int, val *glib.GValue) {
	C.gtk_tree_model_get_value(v.TreeModel, &iter.TreeIter, C.gint(col), C.to_GValueptr(unsafe.Pointer(&val.Value)))
}
func (v *GtkTreeModel) IterNext(iter *GtkTreeIter) bool {
	return gboolean2bool(C.gtk_tree_model_iter_next(v.TreeModel, &iter.TreeIter))
}
func (v *GtkTreeModel) IterChildren(iter *GtkTreeIter, parent *GtkTreeIter) bool {
	return gboolean2bool(C.gtk_tree_model_iter_children(v.TreeModel, &iter.TreeIter, &parent.TreeIter))
}
func (v *GtkTreeModel) IterHasChild(iter *GtkTreeIter) bool {
	return gboolean2bool(C.gtk_tree_model_iter_has_child(v.TreeModel, &iter.TreeIter))
}
func (v *GtkTreeModel) IterNChildren(iter *GtkTreeIter) int {
	return int(C.gtk_tree_model_iter_n_children(v.TreeModel, &iter.TreeIter))
}
func (v *GtkTreeModel) IterNthChild(iter *GtkTreeIter, parent *GtkTreeIter, n int) bool {
	return gboolean2bool(C.gtk_tree_model_iter_nth_child(v.TreeModel, &iter.TreeIter, &parent.TreeIter, C.gint(n)))
}
func (v *GtkTreeModel) IterParent(iter *GtkTreeIter, child *GtkTreeIter) bool {
	return gboolean2bool(C.gtk_tree_model_iter_parent(v.TreeModel, &iter.TreeIter, &child.TreeIter))
}
func (v *GtkTreeModel) GetStringFromIter(iter *GtkTreeIter) string {
	return C.GoString(C.to_charptr(C.gtk_tree_model_get_string_from_iter(v.TreeModel, &iter.TreeIter)))
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
type GtkTreeSelection struct {
	TreeSelection *C.GtkTreeSelection
}

type GtkSelectionMode int

const (
	GTK_SELECTION_NONE     GtkSelectionMode = 0
	GTK_SELECTION_SINGLE   GtkSelectionMode = 1
	GTK_SELECTION_BROWSE   GtkSelectionMode = 2
	GTK_SELECTION_MULTIPLE GtkSelectionMode = 3
	GTK_SELECTION_EXTENDED                  = GTK_SELECTION_MULTIPLE
)

func (v *GtkTreeSelection) Connect(s string, f interface{}, datas ...interface{}) int {
	return glib.ObjectFromNative(unsafe.Pointer(v.TreeSelection)).Connect(s, f, datas...)
}

func (v *GtkTreeSelection) SetMode(m GtkSelectionMode) {
	C.gtk_tree_selection_set_mode(v.TreeSelection, C.GtkSelectionMode(m))
}

func (v *GtkTreeSelection) GetMode() GtkSelectionMode {
	return GtkSelectionMode(C.gtk_tree_selection_get_mode(v.TreeSelection))
}

//gtk_tree_selection_set_select_function (GtkTreeSelection *selection, GtkTreeSelectionFunc func, gpointer data, GDestroyNotify destroy);
//gtk_tree_selection_get_select_function (GtkTreeSelection *selection);
//gtk_tree_selection_get_tree_view (GtkTreeSelection *selection);

func (v *GtkTreeSelection) GetSelected(i *GtkTreeIter) bool {
	return gboolean2bool(C.gtk_tree_selection_get_selected(v.TreeSelection, nil, &i.TreeIter))
}

//gtk_tree_selection_selected_foreach (GtkTreeSelection *selection, GtkTreeSelectionForeachFunc func, gpointer data);
//gtk_tree_selection_get_selected_rows (GtkTreeSelection *selection, GtkTreeModel **model);

func (v *GtkTreeSelection) CountSelectedRows() int {
	return int(C.gtk_tree_selection_count_selected_rows(v.TreeSelection))
}
func (v *GtkTreeSelection) SelectPath(path *GtkTreePath) {
	C.gtk_tree_selection_select_path(v.TreeSelection, path.TreePath)
}
func (v *GtkTreeSelection) UnselectPath(path *GtkTreePath) {
	C.gtk_tree_selection_unselect_path(v.TreeSelection, path.TreePath)
}
func (v *GtkTreeSelection) PathIsSelected(path *GtkTreePath) bool {
	return gboolean2bool(C.gtk_tree_selection_path_is_selected(v.TreeSelection, path.TreePath))
}
func (v *GtkTreeSelection) SelectIter(iter *GtkTreeIter) {
	C.gtk_tree_selection_select_iter(v.TreeSelection, &iter.TreeIter)
}
func (v *GtkTreeSelection) UnselectIter(iter *GtkTreeIter) {
	C.gtk_tree_selection_unselect_iter(v.TreeSelection, &iter.TreeIter)
}
func (v *GtkTreeSelection) IterIsSelected(iter *GtkTreeIter) bool {
	return gboolean2bool(C.gtk_tree_selection_iter_is_selected(v.TreeSelection, &iter.TreeIter))
}
func (v *GtkTreeSelection) SelectAll() {
	C.gtk_tree_selection_select_all(v.TreeSelection)
}
func (v *GtkTreeSelection) UnselectAll() {
	C.gtk_tree_selection_unselect_all(v.TreeSelection)
}
func (v *GtkTreeSelection) SelectRange(start_path *GtkTreePath, end_path *GtkTreePath) {
	C.gtk_tree_selection_select_range(v.TreeSelection, start_path.TreePath, end_path.TreePath)
}
func (v *GtkTreeSelection) UnselectRange(start_path *GtkTreePath, end_path *GtkTreePath) {
	C.gtk_tree_selection_unselect_range(v.TreeSelection, start_path.TreePath, end_path.TreePath)
}

//-----------------------------------------------------------------------
// GtkTreeViewColumn
//-----------------------------------------------------------------------
type GtkTreeViewColumnSizing int

const (
	GTK_TREE_VIEW_COLUMN_GROW_ONLY GtkTreeViewColumnSizing = 0
	GTK_TREE_VIEW_COLUMN_AUTOSIZE  GtkTreeViewColumnSizing = 1
	GTK_TREE_VIEW_COLUMN_FIXED     GtkTreeViewColumnSizing = 2
)

type GtkTreeViewColumn struct {
	TreeViewColumn *C.GtkTreeViewColumn
	*glib.GObject
}

func newTreeViewColumn(column *C.GtkTreeViewColumn) *GtkTreeViewColumn {
	return &GtkTreeViewColumn{
		TreeViewColumn: column,
		GObject:        glib.ObjectFromNative(unsafe.Pointer(column)),
	}
}

func TreeViewColumn() *GtkTreeViewColumn {
	return newTreeViewColumn(C.gtk_tree_view_column_new())
}

//TODO test
func TreeViewColumnWithAttributes2(title string, cell CellRendererLike, attributes ...interface{}) *GtkTreeViewColumn {
	if len(attributes)%2 != 0 {
		log.Panic("Error in gtk.TreeViewColumnWithAttributes: last attribute isn't associated to a value, len(attributes) must be even")
	}
	ptrTitle := C.CString(title)
	defer C.free_string(ptrTitle)
	ret := newTreeViewColumn(C._gtk_tree_view_column_new_with_attribute(
		C.to_gcharptr(ptrTitle), cell.ToGtkCellRenderer()))
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
		C.gtk_tree_view_column_add_attribute(ret.TreeViewColumn,
			cell.ToGtkCellRenderer(), C.to_gcharptr(ptrAttribute), C.gint(column))
	}
	return ret
}

func TreeViewColumnWithAttribute(title string, cell CellRendererLike) *GtkTreeViewColumn {
	ptitle := C.CString(title)
	defer C.free_string(ptitle)
	return newTreeViewColumn(
		C._gtk_tree_view_column_new_with_attribute(C.to_gcharptr(ptitle), cell.ToGtkCellRenderer()))
}
func TreeViewColumnWithAttributes(title string, cell CellRendererLike, attribute string, column int) *GtkTreeViewColumn {
	ptitle := C.CString(title)
	defer C.free_string(ptitle)
	pattribute := C.CString(attribute)
	defer C.free_string(pattribute)
	return newTreeViewColumn(
		C._gtk_tree_view_column_new_with_attributes(C.to_gcharptr(ptitle), cell.ToGtkCellRenderer(), C.to_gcharptr(pattribute), C.gint(column)))
}
func (v *GtkTreeViewColumn) PackStart(cell CellRendererLike, expand bool) {
	C.gtk_tree_view_column_pack_start(v.TreeViewColumn, cell.ToGtkCellRenderer(), bool2gboolean(expand))
}
func (v *GtkTreeViewColumn) PackEnd(cell CellRendererLike, expand bool) {
	C.gtk_tree_view_column_pack_end(v.TreeViewColumn, cell.ToGtkCellRenderer(), bool2gboolean(expand))
}
func (v *GtkTreeViewColumn) Clear() {
	C.gtk_tree_view_column_clear(v.TreeViewColumn)
}

// gtk_tree_view_column_get_cell_renderers //deprecated since 2.18

func (v *GtkTreeViewColumn) AddAttribute(cell CellRendererLike, attribute string, column int) {
	ptr := C.CString(attribute)
	defer C.free_string(ptr)
	C.gtk_tree_view_column_add_attribute(v.TreeViewColumn, cell.ToGtkCellRenderer(), C.to_gcharptr(ptr), C.gint(column))
}

//void gtk_tree_view_column_set_attributes (GtkTreeViewColumn *tree_column, GtkCellRenderer *cell_renderer, ...) G_GNUC_NULL_TERMINATED;
//void gtk_tree_view_column_set_cell_data_func (GtkTreeViewColumn *tree_column, GtkCellRenderer *cell_renderer, GtkTreeCellDataFunc func, gpointer func_data, GDestroyNotify destroy);
//void gtk_tree_view_column_clear_attributes (GtkTreeViewColumn *tree_column, GtkCellRenderer *cell_renderer);

func (v *GtkTreeViewColumn) SetSpacing(spacing int) {
	C.gtk_tree_view_column_set_spacing(v.TreeViewColumn, C.gint(spacing))
}
func (v *GtkTreeViewColumn) GetSpacing() int {
	return int(C.gtk_tree_view_column_get_spacing(v.TreeViewColumn))
}
func (v *GtkTreeViewColumn) SetVisible(resizable bool) {
	C.gtk_tree_view_column_set_visible(v.TreeViewColumn, bool2gboolean(resizable))
}
func (v *GtkTreeViewColumn) GetVisible() bool {
	return gboolean2bool(C.gtk_tree_view_column_get_visible (v.TreeViewColumn))
}
func (v *GtkTreeViewColumn) SetResizable(resizable bool) {
	C.gtk_tree_view_column_set_resizable(v.TreeViewColumn, bool2gboolean(resizable))
}
func (v *GtkTreeViewColumn) GetResizable() bool {
	return gboolean2bool(C.gtk_tree_view_column_get_resizable(v.TreeViewColumn))
}
func (v *GtkTreeViewColumn) SetSizing(s GtkTreeViewColumnSizing) {
	C.gtk_tree_view_column_set_sizing(v.TreeViewColumn, C.GtkTreeViewColumnSizing(s))
}
func (v *GtkTreeViewColumn) GetSizing() GtkTreeViewColumnSizing {
	return GtkTreeViewColumnSizing(C.gtk_tree_view_column_get_sizing(v.TreeViewColumn))
}
func (v *GtkTreeViewColumn) GetWidth() int {
	return int(C.gtk_tree_view_column_get_width(v.TreeViewColumn))
}
func (v *GtkTreeViewColumn) GetFixedWidth() int {
	return int(C.gtk_tree_view_column_get_fixed_width(v.TreeViewColumn))
}
func (v *GtkTreeViewColumn) SetFixedWidth(w int) {
	C.gtk_tree_view_column_set_fixed_width(v.TreeViewColumn, C.gint(w))
}
func (v *GtkTreeViewColumn) SetMinWidth(w int) {
	C.gtk_tree_view_column_set_min_width(v.TreeViewColumn, C.gint(w))
}
func (v *GtkTreeViewColumn) GetMinWidth() int {
	return int(C.gtk_tree_view_column_get_min_width(v.TreeViewColumn))
}
func (v *GtkTreeViewColumn) SetMaxWidth(w int) {
	C.gtk_tree_view_column_set_max_width(v.TreeViewColumn, C.gint(w))
}
func (v *GtkTreeViewColumn) GetMaxWidth() int {
	return int(C.gtk_tree_view_column_get_max_width(v.TreeViewColumn))
}
func (v *GtkTreeViewColumn) Clicked() {
	C.gtk_tree_view_column_clicked(v.TreeViewColumn)
}
func (v *GtkTreeViewColumn) SetTitle(title string) {
	ptr := C.CString(title)
	defer C.free_string(ptr)
	C.gtk_tree_view_column_set_title(v.TreeViewColumn, C.to_gcharptr(ptr))

}
func (v *GtkTreeViewColumn) GetTitle() string {
	return C.GoString(C.to_charptr(C.gtk_tree_view_column_get_title(v.TreeViewColumn)))
}

func (v *GtkTreeViewColumn) SetExpand(expand bool) {
	C.gtk_tree_view_column_set_expand(v.TreeViewColumn, bool2gboolean(expand))
}
func (v *GtkTreeViewColumn) GetExpand() bool {
	return gboolean2bool(C.gtk_tree_view_column_get_expand(v.TreeViewColumn))
}

func (v *GtkTreeViewColumn) SetClickable(clickable bool) {
	C.gtk_tree_view_column_set_clickable(v.TreeViewColumn, bool2gboolean(clickable))
}
func (v *GtkTreeViewColumn) GetClickable() bool {
	return gboolean2bool(C.gtk_tree_view_column_get_clickable(v.TreeViewColumn))
}
//void gtk_tree_view_column_set_widget (GtkTreeViewColumn *tree_column, GtkWidget *widget);
//GtkWidget *gtk_tree_view_column_get_widget (GtkTreeViewColumn *tree_column);
//void gtk_tree_view_column_set_alignment (GtkTreeViewColumn *tree_column, gfloat xalign);
//gfloat gtk_tree_view_column_get_alignment (GtkTreeViewColumn *tree_column);

func (v *GtkTreeViewColumn) SetReorderable(reorderable bool) {
	C.gtk_tree_view_column_set_reorderable(v.TreeViewColumn, bool2gboolean(reorderable))
}
func (v *GtkTreeViewColumn) GetReorderable() bool {
	return gboolean2bool(C.gtk_tree_view_column_get_reorderable(v.TreeViewColumn))
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
type GtkTreeView struct {
	GtkContainer
}

func TreeView() *GtkTreeView {
	return &GtkTreeView{GtkContainer{GtkWidget{
		C.gtk_tree_view_new()}}}
}

//gint gtk_tree_view_get_level_indentation (GtkTreeView *tree_view);
//gboolean gtk_tree_view_get_show_expanders (GtkTreeView *tree_view);
//void gtk_tree_view_set_level_indentation (GtkTreeView *tree_view, gint indentation);
//void gtk_tree_view_set_show_expanders (GtkTreeView *tree_view, gboolean enabled);
//GtkWidget *gtk_tree_view_new_with_model (GtkTreeModel *model);
//GtkTreeModel *gtk_tree_view_get_model (GtkTreeView *tree_view);

func (v *GtkTreeView) SetModel(model GtkTreeModelLike) {
	var tm *C.GtkTreeModel
	if model != nil {
		tm = model.cTreeModel()
	}
	C.gtk_tree_view_set_model(C.to_GtkTreeView(v.Widget), tm)
}
func (v *GtkTreeView) GetSelection() *GtkTreeSelection {
	return &GtkTreeSelection{C.gtk_tree_view_get_selection(C.to_GtkTreeView(v.Widget))}
}

//GtkAdjustment *gtk_tree_view_get_hadjustment (GtkTreeView *tree_view);
//void gtk_tree_view_set_hadjustment (GtkTreeView *tree_view, GtkAdjustment *adjustment);
//GtkAdjustment *gtk_tree_view_get_vadjustment (GtkTreeView *tree_view);
//void gtk_tree_view_set_vadjustment (GtkTreeView *tree_view, GtkAdjustment *adjustment);
//gboolean gtk_tree_view_get_headers_visible (GtkTreeView *tree_view);

func (v *GtkTreeView) SetHeadersVisible(flag bool) {
	C.gtk_tree_view_set_headers_visible(C.to_GtkTreeView(v.Widget), bool2gboolean(flag))
}

//void gtk_tree_view_columns_autosize (GtkTreeView *tree_view);
//gboolean gtk_tree_view_get_headers_clickable (GtkTreeView *tree_view);
//void gtk_tree_view_set_headers_clickable (GtkTreeView *tree_view, gboolean setting);
//void gtk_tree_view_set_rules_hint (GtkTreeView *tree_view, gboolean setting);
//gboolean gtk_tree_view_get_rules_hint (GtkTreeView *tree_view);

func (v *GtkTreeView) AppendColumn(column *GtkTreeViewColumn) int {
	return int(C.gtk_tree_view_append_column(C.to_GtkTreeView(v.Widget), column.TreeViewColumn))
}

//gint gtk_tree_view_remove_column (GtkTreeView *tree_view, GtkTreeViewColumn *column);
//gint gtk_tree_view_insert_column (GtkTreeView *tree_view, GtkTreeViewColumn *column, gint position);
//gint gtk_tree_view_insert_column_with_attributes (GtkTreeView *tree_view, gint position, const gchar *title, GtkCellRenderer *cell, ...) G_GNUC_NULL_TERMINATED;
//gint gtk_tree_view_insert_column_with_data_func (GtkTreeView *tree_view, gint position, const gchar *title, GtkCellRenderer *cell, GtkTreeCellDataFunc func, gpointer data, GDestroyNotify dnotify);

func (v *GtkTreeView) GetColumn(n int) *GtkTreeViewColumn {
	return newTreeViewColumn(C.gtk_tree_view_get_column(C.to_GtkTreeView(v.Widget), C.gint(n)))
}

//GList *gtk_tree_view_get_columns (GtkTreeView *tree_view);
//void gtk_tree_view_move_column_after (GtkTreeView *tree_view, GtkTreeViewColumn *column, GtkTreeViewColumn *base_column);
//void gtk_tree_view_set_expander_column (GtkTreeView *tree_view, GtkTreeViewColumn *column);
//GtkTreeViewColumn *gtk_tree_view_get_expander_column (GtkTreeView *tree_view);
//void gtk_tree_view_set_column_drag_function (GtkTreeView *tree_view, GtkTreeViewColumnDropFunc func, gpointer user_data, GDestroyNotify destroy);
//void gtk_tree_view_scroll_to_point (GtkTreeView *tree_view, gint tree_x, gint tree_y);

func (v *GtkTreeView) ScrollToCell(path *GtkTreePath, col *GtkTreeViewColumn, use bool, ra float64, ca float64) {
	var pcol *C.GtkTreeViewColumn
	if nil == col {
		pcol = nil
	} else {
		pcol = col.TreeViewColumn
	}
	C.gtk_tree_view_scroll_to_cell(C.to_GtkTreeView(v.Widget), path.TreePath,
		pcol, bool2gboolean(use), C.gfloat(ra), C.gfloat(ca))
}
func (v *GtkTreeView) SetCursor(path *GtkTreePath, col *GtkTreeViewColumn, se bool) {
	var pcol *C.GtkTreeViewColumn
	if nil == col {
		pcol = nil
	} else {
		pcol = col.TreeViewColumn
	}
	C.gtk_tree_view_set_cursor(C.to_GtkTreeView(v.Widget), path.TreePath,
		pcol, bool2gboolean(se))
}

//void gtk_tree_view_set_cursor_on_cell (GtkTreeView *tree_view, GtkTreePath *path, GtkTreeViewColumn *focus_column, GtkCellRenderer *focus_cell, gboolean start_editing);

func (v *GtkTreeView) GetCursor(path **GtkTreePath, focus_column **GtkTreeViewColumn) {
	*path = &GtkTreePath{nil}
	if nil != focus_column {
		*focus_column = &GtkTreeViewColumn{nil, nil}
		C.gtk_tree_view_get_cursor(C.to_GtkTreeView(v.Widget), &(*path).TreePath, &(*focus_column).TreeViewColumn)
	} else {
		C.gtk_tree_view_get_cursor(C.to_GtkTreeView(v.Widget), &(*path).TreePath, nil)
	}
}

//void gtk_tree_view_row_activated (GtkTreeView *tree_view, GtkTreePath *path, GtkTreeViewColumn *column);

func (v *GtkTreeView) ExpandAll() {
	C.gtk_tree_view_expand_all(C.to_GtkTreeView(v.Widget))
}
func (v *GtkTreeView) CollapseAll() {
	C.gtk_tree_view_collapse_all(C.to_GtkTreeView(v.Widget))
}

//void gtk_tree_view_expand_to_path (GtkTreeView *tree_view, GtkTreePath *path);

func (v *GtkTreeView) ExpandRow(path *GtkTreePath, openall bool) bool {
	return gboolean2bool(C.gtk_tree_view_expand_row(C.to_GtkTreeView(v.Widget), path.TreePath, bool2gboolean(openall)))
}
func (v *GtkTreeView) CollapseRow(path *GtkTreePath) bool {
	return gboolean2bool(C.gtk_tree_view_collapse_row(C.to_GtkTreeView(v.Widget), path.TreePath))
}

//void gtk_tree_view_map_expanded_rows (GtkTreeView *tree_view, GtkTreeViewMappingFunc func, gpointer data);

func (v *GtkTreeView) RowExpanded(path *GtkTreePath) bool {
	return gboolean2bool(C.gtk_tree_view_row_expanded(C.to_GtkTreeView(v.Widget), path.TreePath))
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
type GtkIconView struct {
	GtkContainer
}

func IconView() *GtkIconView {
	return &GtkIconView{GtkContainer{GtkWidget{
		C.gtk_icon_view_new()}}}
}
func IconViewWithModel(model GtkTreeModelLike) *GtkIconView {
	var tm *C.GtkTreeModel
	if model != nil {
		tm = model.cTreeModel()
	}
	return &GtkIconView{GtkContainer{GtkWidget{
		C.gtk_icon_view_new_with_model(tm)}}}
}
func (v *GtkIconView) GetModel() *GtkTreeModel {
	return &GtkTreeModel{
		C.gtk_icon_view_get_model(C.to_GtkIconView(v.Widget))}
}
func (v *GtkIconView) SetModel(model GtkTreeModelLike) {
	var tm *C.GtkTreeModel
	if model != nil {
		tm = model.cTreeModel()
	}
	C.gtk_icon_view_set_model(C.to_GtkIconView(v.Widget), tm)
}
func (v *GtkIconView) GetTextColumn() int {
	return int(C.gtk_icon_view_get_text_column(C.to_GtkIconView(v.Widget)))
}
func (v *GtkIconView) SetTextColumn(text_column int) {
	C.gtk_icon_view_set_text_column(C.to_GtkIconView(v.Widget), C.gint(text_column))
}
func (v *GtkIconView) GetMarkupColumn() int {
	return int(C.gtk_icon_view_get_markup_column(C.to_GtkIconView(v.Widget)))
}
func (v *GtkIconView) SetMarkupColumn(markup_column int) {
	C.gtk_icon_view_set_markup_column(C.to_GtkIconView(v.Widget), C.gint(markup_column))
}
func (v *GtkIconView) GetPixbufColumn() int {
	return int(C.gtk_icon_view_get_pixbuf_column(C.to_GtkIconView(v.Widget)))
}
func (v *GtkIconView) SetPixbufColumn(pixbuf_column int) {
	C.gtk_icon_view_set_pixbuf_column(C.to_GtkIconView(v.Widget), C.gint(pixbuf_column))
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

func (v *GtkIconView) ScrollToPath(path *GtkTreePath, use bool, ra float64, ca float64) {
	C.gtk_icon_view_scroll_to_path(C.to_GtkIconView(v.Widget), path.TreePath,
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
	ToGtkCellRenderer() *C.GtkCellRenderer
}
type GtkCellRenderer struct {
	CellRenderer *C.GtkCellRenderer
	CellRendererLike
}

func (v *GtkCellRenderer) ToGtkCellRenderer() *C.GtkCellRenderer {
	return v.CellRenderer
}

func (v *GtkCellRenderer) Connect(s string, f interface{}, datas ...interface{}) int {
	return glib.ObjectFromNative(unsafe.Pointer(v.CellRenderer)).Connect(s, f, datas...)
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
type GtkCellRendererAccel struct {
	GtkCellRenderer
}

func CellRendererAccel() *GtkCellRendererAccel {
	return &GtkCellRendererAccel{GtkCellRenderer{
		C.gtk_cell_renderer_accel_new(), nil}}
}

//-----------------------------------------------------------------------
// GtkCellRendererCombo
//-----------------------------------------------------------------------
type GtkCellRendererCombo struct {
	GtkCellRenderer
}

func CellRendererCombo() *GtkCellRendererCombo {
	return &GtkCellRendererCombo{GtkCellRenderer{
		C.gtk_cell_renderer_combo_new(), nil}}
}

//-----------------------------------------------------------------------
// GtkCellRendererPixbuf
//-----------------------------------------------------------------------
type GtkCellRendererPixbuf struct {
	GtkCellRenderer
}

func CellRendererPixbuf() *GtkCellRendererPixbuf {
	return &GtkCellRendererPixbuf{GtkCellRenderer{
		C.gtk_cell_renderer_pixbuf_new(), nil}}
}

//-----------------------------------------------------------------------
// GtkCellRendererProgress
//-----------------------------------------------------------------------
type GtkCellRendererProgress struct {
	GtkCellRenderer
}

func CellRendererProgress() *GtkCellRendererProgress {
	return &GtkCellRendererProgress{GtkCellRenderer{
		C.gtk_cell_renderer_progress_new(), nil}}
}

//-----------------------------------------------------------------------
// GtkCellRendererSpin
//-----------------------------------------------------------------------
type GtkCellRendererSpin struct {
	GtkCellRenderer
}

func CellRendererSpin() *GtkCellRendererSpin {
	return &GtkCellRendererSpin{GtkCellRenderer{
		C.gtk_cell_renderer_spin_new(), nil}}
}

//-----------------------------------------------------------------------
// GtkCellRendererText
//-----------------------------------------------------------------------
type GtkCellRendererText struct {
	GtkCellRenderer
}

func CellRendererText() *GtkCellRendererText {
	return &GtkCellRendererText{GtkCellRenderer{
		C.gtk_cell_renderer_text_new(), nil}}
}
func (v *GtkCellRendererText) SetFixedHeightFromFont(number_of_rows int) {
	C.gtk_cell_renderer_text_set_fixed_height_from_font(C.to_GtkCellRendererText(v.CellRenderer), C.gint(number_of_rows))
}

//-----------------------------------------------------------------------
// GtkCellRendererToggle
//-----------------------------------------------------------------------
type GtkCellRendererToggle struct {
	GtkCellRenderer
}

func CellRendererToggle() *GtkCellRendererToggle {
	return &GtkCellRendererToggle{GtkCellRenderer{
		C.gtk_cell_renderer_toggle_new(), nil}}
}
func (v *GtkCellRendererToggle) GetRadio() bool {
	return gboolean2bool(C.gtk_cell_renderer_toggle_get_radio(C.to_GtkCellRendererToggle(v.CellRenderer)))
}
func (v *GtkCellRendererToggle) SetRadio(radio bool) {
	C.gtk_cell_renderer_toggle_set_radio(C.to_GtkCellRendererToggle(v.CellRenderer), bool2gboolean(radio))
}
func (v *GtkCellRendererToggle) GetActive() bool {
	return gboolean2bool(C.gtk_cell_renderer_toggle_get_active(C.to_GtkCellRendererToggle(v.CellRenderer)))
}
func (v *GtkCellRendererToggle) SetActive(active bool) {
	C.gtk_cell_renderer_toggle_set_active(C.to_GtkCellRendererToggle(v.CellRenderer), bool2gboolean(active))
}
func (v *GtkCellRendererToggle) GetActivatable() bool {
	panic_if_version_older(2, 18, 0, "gtk_cell_renderer_toggle_get_activatable()")
	return gboolean2bool(C._gtk_cell_renderer_toggle_get_activatable(C.to_GtkCellRendererToggle(v.CellRenderer)))
}
func (v *GtkCellRendererToggle) SetActivatable(activatable bool) {
	panic_if_version_older(2, 18, 0, "gtk_cell_renderer_toggle_set_activatable()")
	C._gtk_cell_renderer_toggle_set_activatable(C.to_GtkCellRendererToggle(v.CellRenderer), bool2gboolean(activatable))
}

//-----------------------------------------------------------------------
// GtkCellRendererSpinner
//-----------------------------------------------------------------------
type GtkCellRendererSpinner struct {
	GtkCellRenderer
}

func CellRendererSpinner() *GtkCellRendererSpinner {
	panic_if_version_older(2, 20, 0, "gtk_cell_renderer_spinner_new()")
	return &GtkCellRendererSpinner{GtkCellRenderer{
		C._gtk_cell_renderer_spinner_new(), nil}}
}

//-----------------------------------------------------------------------
// GtkListStore
//-----------------------------------------------------------------------
const (
	GTK_TYPE_CHAR    = glib.G_TYPE_CHAR
	GTK_TYPE_UCHAR   = glib.G_TYPE_UCHAR
	GTK_TYPE_BOOL    = glib.G_TYPE_BOOL
	GTK_TYPE_INT     = glib.G_TYPE_INT
	GTK_TYPE_UINT    = glib.G_TYPE_UINT
	GTK_TYPE_LONG    = glib.G_TYPE_LONG
	GTK_TYPE_ULONG   = glib.G_TYPE_ULONG
	GTK_TYPE_FLOAT   = glib.G_TYPE_FLOAT
	GTK_TYPE_DOUBLE  = glib.G_TYPE_DOUBLE
	GTK_TYPE_STRING  = glib.G_TYPE_STRING
	GTK_TYPE_BOXED   = glib.G_TYPE_BOXED
	GTK_TYPE_POINTER = glib.G_TYPE_POINTER
	GTK_TYPE_PIXBUF  = GTK_TYPE_POINTER
)

type GtkListStore struct {
	GtkTreeModel
	ListStore *C.GtkListStore
}

func ListStore(v ...interface{}) *GtkListStore {
	types := C.make_gtypes(C.int(len(v)))
	for n := range v {
		C.set_gtype(types, C.int(n), C.int(v[n].(int)))
	}
	defer C.destroy_gtypes(types)
	cliststore := C.gtk_list_store_newv(C.gint(len(v)), types)
	return &GtkListStore{GtkTreeModel{
		C.to_GtkTreeModelFromListStore(cliststore)}, cliststore}
}

//void gtk_list_store_set_column_types (GtkListStore *list_store, gint n_columns, GType *types);

func (v *GtkListStore) Set(iter *GtkTreeIter, a ...interface{}) {
	for r := range a {
		v.SetValue(iter, r, a[r])
	}
}
func (v *GtkListStore) SetValue(iter *GtkTreeIter, column int, a interface{}) {
	gv := glib.GValueFromNative(a)
	if gv != nil {
		C.gtk_list_store_set_value(v.ListStore, &iter.TreeIter, C.gint(column), C.to_GValueptr(unsafe.Pointer(gv)))
	} else {
		if pv, ok := a.(*[0]uint8); ok {
			C._gtk_list_store_set_ptr(v.ListStore, &iter.TreeIter, C.gint(column), unsafe.Pointer(pv))
		} else {
			av := reflect.ValueOf(a)
			if av.CanAddr() {
				C._gtk_list_store_set_addr(v.ListStore, &iter.TreeIter, C.gint(column), unsafe.Pointer(av.UnsafeAddr()))
			} else {
				C._gtk_list_store_set_addr(v.ListStore, &iter.TreeIter, C.gint(column), unsafe.Pointer(&a))
			}
		}
	}
}
func (v *GtkListStore) Remove(iter *GtkTreeIter) bool {
	return gboolean2bool(C.gtk_list_store_remove(v.ListStore, &iter.TreeIter))
}
func (v *GtkListStore) Insert(iter *GtkTreeIter, position int) {
	C.gtk_list_store_insert(v.ListStore, &iter.TreeIter, C.gint(position))
}
func (v *GtkListStore) InsertBefore(iter *GtkTreeIter, sibling *GtkTreeIter) {
	C.gtk_list_store_insert_before(v.ListStore, &iter.TreeIter, &sibling.TreeIter)
}
func (v *GtkListStore) InsertAfter(iter *GtkTreeIter, sibling *GtkTreeIter) {
	C.gtk_list_store_insert_after(v.ListStore, &iter.TreeIter, &sibling.TreeIter)
}

//void gtk_list_store_insert_with_values (GtkListStore *list_store, GtkTreeIter *iter, gint position, ...);
//void gtk_list_store_insert_with_valuesv (GtkListStore *list_store, GtkTreeIter *iter, gint position, gint *columns, GValue *values, gint n_values);

func (v *GtkListStore) Prepend(iter *GtkTreeIter) {
	C.gtk_list_store_prepend(v.ListStore, &iter.TreeIter)
}
func (v *GtkListStore) Append(iter *GtkTreeIter) {
	C.gtk_list_store_append(v.ListStore, &iter.TreeIter)
}
func (v *GtkListStore) Clear() {
	C.gtk_list_store_clear(v.ListStore)
}
func (v *GtkListStore) IterIsValid(iter *GtkTreeIter) bool {
	log.Println("Warning: GtkListStore.IterIsValid: This function is slow. Only use it for debugging and/or testing purposes.")
	return gboolean2bool(C.gtk_list_store_iter_is_valid(v.ListStore, &iter.TreeIter))
}
func (v *GtkListStore) Reorder(i *int) {
	gi := C.gint(*i)
	C.gtk_list_store_reorder(v.ListStore, &gi)
	*i = int(gi)
}
func (v *GtkListStore) Swap(a *GtkTreeIter, b *GtkTreeIter) {
	C.gtk_list_store_swap(v.ListStore, &a.TreeIter, &b.TreeIter)
}
func (v *GtkListStore) MoveBefore(iter *GtkTreeIter, position *GtkTreeIter) {
	C.gtk_list_store_move_before(v.ListStore, &iter.TreeIter, &position.TreeIter)
}
func (v *GtkListStore) MoveAfter(iter *GtkTreeIter, position *GtkTreeIter) {
	C.gtk_list_store_move_after(v.ListStore, &iter.TreeIter, &position.TreeIter)
}

//TODO instead of using this methods to change between treemodel and liststore, is better to usa an interface GtkTreeModelLike
//nb: ListStore e TreeStore sono un TreeModel (implementano GtkTreeModel!)
/*func (v *GtkListStore) ToTreeModel() *GtkTreeModel {
	return &GtkTreeModel{
		C.to_GtkTreeModelFromListStore(v.ListStore)}
}*/
/*func (v *GtkTreeModel) ToListStore() *GtkListStore {
	return &GtkListStore{
		C.to_GtkListStoreFromTreeModel(v.TreeModel)}
}*/

//-----------------------------------------------------------------------
// GtkTreeStore
//-----------------------------------------------------------------------
type GtkTreeStore struct {
	GtkTreeModel
	TreeStore *C.GtkTreeStore
}

func TreeStore(v ...interface{}) *GtkTreeStore {
	types := C.make_gtypes(C.int(len(v)))
	for n := range v {
		C.set_gtype(types, C.int(n), C.int(v[n].(int)))
	}
	defer C.destroy_gtypes(types)
	ctreestore := C.gtk_tree_store_newv(C.gint(len(v)), types)
	return &GtkTreeStore{GtkTreeModel{C.to_GtkTreeModelFromTreeStore(ctreestore)}, ctreestore}
}

// void gtk_tree_store_set_column_types (GtkTreeStore *tree_store, gint n_columns, GType *types); void gtk_tree_store_set_value (GtkTreeStore *tree_store, GtkTreeIter *iter, gint column, GValue *value);

func (v *GtkTreeStore) Set(iter *GtkTreeIter, a ...interface{}) {
	for r := range a {
		v.SetValue(iter, r, a[r])
	}
}
func (v *GtkTreeStore) SetValue(iter *GtkTreeIter, column int, a interface{}) {
	gv := glib.GValueFromNative(a)
	if gv != nil {
		C.gtk_tree_store_set_value(v.TreeStore, &iter.TreeIter, C.gint(column), C.to_GValueptr(unsafe.Pointer(gv)))
	} else {
		if pv, ok := a.(*[0]uint8); ok {
			C._gtk_tree_store_set_ptr(v.TreeStore, &iter.TreeIter, C.gint(column), unsafe.Pointer(pv))
		} else {
			av := reflect.ValueOf(a)
			if av.CanAddr() {
				C._gtk_tree_store_set_addr(v.TreeStore, &iter.TreeIter, C.gint(column), unsafe.Pointer(av.UnsafeAddr()))
			} else {
				C._gtk_tree_store_set_addr(v.TreeStore, &iter.TreeIter, C.gint(column), unsafe.Pointer(&a))
			}
		}
	}
}
func (v *GtkTreeStore) Remove(iter *GtkTreeIter) bool {
	return gboolean2bool(C.gtk_tree_store_remove(v.TreeStore, &iter.TreeIter))
}
func (v *GtkTreeStore) Insert(iter *GtkTreeIter, parent *GtkTreeIter, position int) {
	C.gtk_tree_store_insert(v.TreeStore, &iter.TreeIter, &parent.TreeIter, C.gint(position))
}
func (v *GtkTreeStore) InsertBefore(iter *GtkTreeIter, parent *GtkTreeIter, sibling *GtkTreeIter) {
	C.gtk_tree_store_insert_before(v.TreeStore, &iter.TreeIter, &parent.TreeIter, &sibling.TreeIter)
}
func (v *GtkTreeStore) InsertAfter(iter *GtkTreeIter, parent *GtkTreeIter, sibling *GtkTreeIter) {
	C.gtk_tree_store_insert_after(v.TreeStore, &iter.TreeIter, &parent.TreeIter, &sibling.TreeIter)
}

// void gtk_tree_store_insert_with_values (GtkTreeStore *tree_store, GtkTreeIter *iter, GtkTreeIter *parent, gint position, ...);
// void gtk_tree_store_insert_with_valuesv (GtkTreeStore *tree_store, GtkTreeIter *iter, GtkTreeIter *parent, gint position, gint *columns, GValue *values, gint n_values);

func (v *GtkTreeStore) Prepend(iter *GtkTreeIter, parent *GtkTreeIter) {
	if parent == nil {
		C.gtk_tree_store_prepend(v.TreeStore, &iter.TreeIter, nil)
	} else {
		C.gtk_tree_store_prepend(v.TreeStore, &iter.TreeIter, &parent.TreeIter)
	}
}
func (v *GtkTreeStore) Append(iter *GtkTreeIter, parent *GtkTreeIter) {
	if parent == nil {
		C.gtk_tree_store_append(v.TreeStore, &iter.TreeIter, nil)
	} else {
		C.gtk_tree_store_append(v.TreeStore, &iter.TreeIter, &parent.TreeIter)
	}
}

// gtk_tree_store_is_ancestor

func (v *GtkTreeStore) ToTreeModel() *GtkTreeModel {
	return &GtkTreeModel{
		C.to_GtkTreeModelFromTreeStore(v.TreeStore)}
}
func (v *GtkTreeStore) IterDepth(iter *GtkTreeIter) int {
	return int(C.gtk_tree_store_iter_depth(v.TreeStore, &iter.TreeIter))
}
func (v *GtkTreeStore) Clear() {
	C.gtk_tree_store_clear(v.TreeStore)
}
func (v *GtkTreeStore) IterIsValid(iter *GtkTreeIter) bool {
	log.Println("Warning: GtkTreeStore.IterIsValid: This function is slow. Only use it for debugging and/or testing purposes.")
	return gboolean2bool(C.gtk_tree_store_iter_is_valid(v.TreeStore, &iter.TreeIter))
}
func (v *GtkTreeStore) Reorder(iter *GtkTreeIter, i *int) {
	gi := C.gint(*i)
	C.gtk_tree_store_reorder(v.TreeStore, &iter.TreeIter, &gi)
	*i = int(gi)
}
func (v *GtkTreeStore) Swap(a *GtkTreeIter, b *GtkTreeIter) {
	C.gtk_tree_store_swap(v.TreeStore, &a.TreeIter, &b.TreeIter)
}
func (v *GtkTreeStore) MoveBefore(iter *GtkTreeIter, position *GtkTreeIter) {
	C.gtk_tree_store_move_before(v.TreeStore, &iter.TreeIter, &position.TreeIter)
}
func (v *GtkTreeStore) MoveAfter(iter *GtkTreeIter, position *GtkTreeIter) {
	C.gtk_tree_store_move_after(v.TreeStore, &iter.TreeIter, &position.TreeIter)
}

//-----------------------------------------------------------------------
// GtkComboBox
//-----------------------------------------------------------------------
type GtkComboBox struct {
	GtkBin
}

func ComboBox() *GtkComboBox {
	return &GtkComboBox{GtkBin{GtkContainer{GtkWidget{
		C.gtk_combo_box_new()}}}}
}
func ComboBoxWithEntry() *GtkComboBox {
	deprecated_since(2, 24, 0, "gtk_combo_box_new_with_entry()")
	return &GtkComboBox{GtkBin{GtkContainer{GtkWidget{
		C._gtk_combo_box_new_with_entry()}}}}
}
func ComboBoxWithModel(model *GtkTreeModel) *GtkComboBox {
	return &GtkComboBox{GtkBin{GtkContainer{GtkWidget{
		C.gtk_combo_box_new_with_model(model.TreeModel)}}}}
}
func ComboBoxWithModelAndEntry(model *GtkTreeModel) *GtkComboBox {
	deprecated_since(2, 24, 0, "gtk_combo_box_new_with_model_and_entry()")
	return &GtkComboBox{GtkBin{GtkContainer{GtkWidget{
		C._gtk_combo_box_new_with_model_and_entry(model.TreeModel)}}}}
}
func (v *GtkComboBox) GetWrapWidth() int {
	return int(C.gtk_combo_box_get_wrap_width(C.to_GtkComboBox(v.Widget)))
}
func (v *GtkComboBox) SetWrapWidth(width int) {
	C.gtk_combo_box_set_wrap_width(C.to_GtkComboBox(v.Widget), C.gint(width))
}
func (v *GtkComboBox) GetRowSpanColumn() int {
	return int(C.gtk_combo_box_get_row_span_column(C.to_GtkComboBox(v.Widget)))
}
func (v *GtkComboBox) SetRowSpanColumn(row_span int) {
	C.gtk_combo_box_set_row_span_column(C.to_GtkComboBox(v.Widget), C.gint(row_span))
}
func (v *GtkComboBox) GetColumnSpanColumn() int {
	return int(C.gtk_combo_box_get_column_span_column(C.to_GtkComboBox(v.Widget)))
}
func (v *GtkComboBox) SetColumnSpanColumn(column_span int) {
	C.gtk_combo_box_set_column_span_column(C.to_GtkComboBox(v.Widget), C.gint(column_span))
}
func (v *GtkComboBox) GetActive() int {
	return int(C.gtk_combo_box_get_active(C.to_GtkComboBox(v.Widget)))
}
func (v *GtkComboBox) SetActive(width int) {
	C.gtk_combo_box_set_active(C.to_GtkComboBox(v.Widget), C.gint(width))
}
func (v *GtkComboBox) GetActiveIter(iter *GtkTreeIter) bool {
	return gboolean2bool(C.gtk_combo_box_get_active_iter(C.to_GtkComboBox(v.Widget), &iter.TreeIter))
}
func (v *GtkComboBox) SetActiveIter(iter *GtkTreeIter) {
	C.gtk_combo_box_set_active_iter(C.to_GtkComboBox(v.Widget), &iter.TreeIter)
}
func (v *GtkComboBox) GetModel() *GtkTreeModel {
	return &GtkTreeModel{
		C.gtk_combo_box_get_model(C.to_GtkComboBox(v.Widget))}
}
func (v *GtkComboBox) SetModel(model *GtkTreeModel) {
	C.gtk_combo_box_set_model(C.to_GtkComboBox(v.Widget), model.TreeModel)
}

//Deprecated since 2.24. Use GtkComboBoxText.
func ComboBoxNewText() *GtkComboBox {
	deprecated_since(2, 24, 0, "gtk_combo_box_new_text()")
	return &GtkComboBox{GtkBin{GtkContainer{GtkWidget{
		C.gtk_combo_box_new_text()}}}}
}

//Deprecated since 2.24. Use GtkComboBoxText.
func (v *GtkComboBox) AppendText(text string) {
	deprecated_since(2, 24, 0, "gtk_combo_box_append_text()")
	ptr := C.CString(text)
	defer C.free_string(ptr)
	C.gtk_combo_box_append_text(C.to_GtkComboBox(v.Widget), C.to_gcharptr(ptr))
}

//Deprecated since 2.24. Use GtkComboBoxText.
func (v *GtkComboBox) InsertText(text string, position int) {
	deprecated_since(2, 24, 0, "gtk_combo_box_insert_text()")
	ptr := C.CString(text)
	defer C.free_string(ptr)
	C.gtk_combo_box_insert_text(C.to_GtkComboBox(v.Widget), C.gint(position), C.to_gcharptr(ptr))
}

//Deprecated since 2.24. Use GtkComboBoxText.
func (v *GtkComboBox) PrependText(text string) {
	deprecated_since(2, 24, 0, "gtk_combo_box_prepend_text()")
	ptr := C.CString(text)
	defer C.free_string(ptr)
	C.gtk_combo_box_prepend_text(C.to_GtkComboBox(v.Widget), C.to_gcharptr(ptr))
}

//Deprecated since 2.24. Use GtkComboBoxText.
func (v *GtkComboBox) RemoveText(position int) {
	deprecated_since(2, 24, 0, "gtk_combo_box_remove_text()")
	C.gtk_combo_box_remove_text(C.to_GtkComboBox(v.Widget), C.gint(position))
}

//Deprecated since 2.24. Use GtkComboBoxText or, if combo box contains an entry,
// get text directly from GtkEntry.
func (v *GtkComboBox) GetActiveText() string {
	deprecated_since(2, 24, 0, "gtk_combo_box_get_active_text()")
	return C.GoString(C.to_charptr(C.gtk_combo_box_get_active_text(C.to_GtkComboBox(v.Widget))))
}
func (v *GtkComboBox) Popup() {
	C.gtk_combo_box_popup(C.to_GtkComboBox(v.Widget))
}
func (v *GtkComboBox) Popdown() {
	C.gtk_combo_box_popdown(C.to_GtkComboBox(v.Widget))
}

// gtk_combo_box_get_popup_accessible
// gtk_combo_box_get_row_separator_func
// gtk_combo_box_set_row_separator_func

func (v *GtkComboBox) SetAddTearoffs(add_tearoffs bool) {
	C.gtk_combo_box_set_add_tearoffs(C.to_GtkComboBox(v.Widget), bool2gboolean(add_tearoffs))
}
func (v *GtkComboBox) GetAddTearoffs() bool {
	return gboolean2bool(C.gtk_combo_box_get_add_tearoffs(C.to_GtkComboBox(v.Widget)))
}
func (v *GtkComboBox) SetTitle(title string) {
	ptr := C.CString(title)
	defer C.free_string(ptr)
	C.gtk_combo_box_set_title(C.to_GtkComboBox(v.Widget), C.to_gcharptr(ptr))
}
func (v *GtkComboBox) GetTitle() string {
	return C.GoString(C.to_charptr(C.gtk_combo_box_get_title(C.to_GtkComboBox(v.Widget))))
}
func (v *GtkComboBox) SetFocusOnClick(focus_on_click bool) {
	C.gtk_combo_box_set_focus_on_click(C.to_GtkComboBox(v.Widget), bool2gboolean(focus_on_click))
}
func (v *GtkComboBox) GetFocusOnClick() bool {
	return gboolean2bool(C.gtk_combo_box_get_focus_on_click(C.to_GtkComboBox(v.Widget)))
}

// gtk_combo_box_set_button_sensitivity
// gtk_combo_box_get_button_sensitivity
// gtk_combo_box_get_has_entry //since 2.24
// gtk_combo_box_set_entry_text_column //since 2.24
// gtk_combo_box_get_entry_text_column //since 2.24

//-----------------------------------------------------------------------
// GtkComboBoxText
//-----------------------------------------------------------------------
type GtkComboBoxText struct {
	GtkComboBox
}

func ComboBoxText() *GtkComboBoxText {
	return &GtkComboBoxText{GtkComboBox{GtkBin{GtkContainer{GtkWidget{
		C._gtk_combo_box_text_new()}}}}}
}
func ComboBoxTextWithEntry() *GtkComboBoxText {
	return &GtkComboBoxText{GtkComboBox{GtkBin{GtkContainer{GtkWidget{
		C._gtk_combo_box_text_new_with_entry()}}}}}
}
func (v *GtkComboBoxText) AppendText(text string) {
	ptr := C.CString(text)
	defer C.free_string(ptr)
	C._gtk_combo_box_text_append_text(C.to_GtkComboBoxText(v.Widget), C.to_gcharptr(ptr))
}
func (v *GtkComboBoxText) InsertText(position int, text string) {
	ptr := C.CString(text)
	defer C.free_string(ptr)
	C._gtk_combo_box_text_insert_text(C.to_GtkComboBoxText(v.Widget), C.gint(position), C.to_gcharptr(ptr))
}
func (v *GtkComboBoxText) PrependText(text string) {
	ptr := C.CString(text)
	defer C.free_string(ptr)
	C._gtk_combo_box_text_prepend_text(C.to_GtkComboBoxText(v.Widget), C.to_gcharptr(ptr))
}
func (v *GtkComboBoxText) Remove(position int) {
	C._gtk_combo_box_text_remove(C.to_GtkComboBoxText(v.Widget), C.gint(position))
}
func (v *GtkComboBoxText) GetActiveText() string {
	return C.GoString(C.to_charptr(C._gtk_combo_box_text_get_active_text(C.to_GtkComboBoxText(v.Widget))))
}

//-----------------------------------------------------------------------
// GtkComboBoxEntry
//-----------------------------------------------------------------------
type GtkComboBoxEntry struct {
	GtkComboBox
}

func ComboBoxEntry() *GtkComboBoxEntry {
	deprecated_since(2, 24, 0, "gtk_combo_box_entry_new()")
	return &GtkComboBoxEntry{GtkComboBox{GtkBin{GtkContainer{GtkWidget{
		C.gtk_combo_box_entry_new()}}}}}
}

// gtk_combo_box_entry_new_with_model

func ComboBoxEntryNewText() *GtkComboBoxEntry {
	deprecated_since(2, 24, 0, "gtk_combo_box_entry_new_text()")
	return &GtkComboBoxEntry{GtkComboBox{GtkBin{GtkContainer{GtkWidget{
		C.gtk_combo_box_entry_new_text()}}}}}
}
func (v *GtkComboBoxEntry) GetTextColumn() int {
	deprecated_since(2, 24, 0, "gtk_combo_box_entry_get_text_column()")
	return int(C.gtk_combo_box_entry_get_text_column(C.to_GtkComboBoxEntry(v.Widget)))
}
func (v *GtkComboBoxEntry) SetTextColumn(text_column int) {
	deprecated_since(2, 24, 0, "gtk_combo_box_entry_set_text_column()")
	C.gtk_combo_box_entry_set_text_column(C.to_GtkComboBoxEntry(v.Widget), C.gint(text_column))
}

//-----------------------------------------------------------------------
// GtkMenu
//-----------------------------------------------------------------------
type GtkMenu struct {
	GtkContainer
	//TODO GtkMenuShell
}

func Menu() *GtkMenu {
	return &GtkMenu{GtkContainer{GtkWidget{C.gtk_menu_new()}}}
}

// void gtk_menu_set_screen (GtkMenu *menu, GdkScreen *screen);

//TODO remove when GtkMenuShell is done
func (v *GtkMenu) Append(child WidgetLike) {
	C.gtk_menu_shell_append(C.to_GtkMenuShell(v.Widget), child.ToNative())
}

//TODO remove when GtkMenuShell is done
func (v *GtkMenu) Prepend(child WidgetLike) {
	C.gtk_menu_shell_prepend(C.to_GtkMenuShell(v.Widget), child.ToNative())
}

//TODO remove when GtkMenuShell is done
func (v *GtkMenu) Insert(child WidgetLike, position int) {
	C.gtk_menu_shell_insert(C.to_GtkMenuShell(v.Widget), child.ToNative(), C.gint(position))
}

// void gtk_menu_reorder_child(GtkMenu *menu, GtkWidget *child, gint position);
// void gtk_menu_attach(GtkMenu *menu, GtkWidget *child, guint left_attach, guint right_attach, guint top_attach, guint bottom_attach);

func (v *GtkMenu) Popup(parent_menu_shell, parent_menu_item WidgetLike, f GtkMenuPositionFunc, data interface{}, button uint, active_item uint) {
	var pms, pmi *C.GtkWidget
	if parent_menu_shell != nil {
		pms = parent_menu_shell.ToNative()
	}
	if parent_menu_item != nil {
		pmi = parent_menu_item.ToNative()
	}
	C._gtk_menu_popup(v.Widget, pms, pmi, unsafe.Pointer(&GtkMenuPositionFuncInfo{v, f, data}), C.guint(button), C.guint32(active_item))
}

// void gtk_menu_set_accel_group (GtkMenu *menu, GtkAccelGroup *accel_group);
// GtkAccelGroup* gtk_menu_get_accel_group (GtkMenu *menu);
// void gtk_menu_set_accel_path(GtkMenu *menu, const gchar *accel_path);
// const gchar* gtk_menu_get_accel_path(GtkMenu *menu);
// void gtk_menu_set_title(GtkMenu *menu, const gchar *title);
// G_CONST_RETURN gchar *gtk_menu_get_title(GtkMenu *menu);
// void gtk_menu_set_monitor(GtkMenu *menu, gint monitor_num);
// gint gtk_menu_get_monitor(GtkMenu *menu);

func (v *GtkMenu) GetTearoffState() bool {
	return gboolean2bool(C.gtk_menu_get_tearoff_state(C.to_GtkMenu(v.Widget)))
}
func (v *GtkMenu) SetReserveToggleSize(b bool) {
	panic_if_version_older(2, 18, 0, "gtk_menu_set_reserve_toggle_size()")
	C._gtk_menu_set_reserve_toggle_size(C.to_GtkMenu(v.Widget), bool2gboolean(b))
}
func (v *GtkMenu) GetReserveToggleSize() bool {
	panic_if_version_older(2, 18, 0, "gtk_menu_get_reserve_toggle_size()")
	return gboolean2bool(C._gtk_menu_get_reserve_toggle_size(C.to_GtkMenu(v.Widget)))
}
func (v *GtkMenu) Popdown() {
	C.gtk_menu_popdown(C.to_GtkMenu(v.Widget))
}
func (v *GtkMenu) Reposition() {
	C.gtk_menu_reposition(C.to_GtkMenu(v.Widget))
}
func (v *GtkMenu) GetActive() *GtkWidget {
	return &GtkWidget{C.gtk_menu_get_active(C.to_GtkMenu(v.Widget))}
}

// void gtk_menu_set_active (GtkMenu *menu, guint index_);

func (v *GtkMenu) SetTearoffState(b bool) {
	C.gtk_menu_set_tearoff_state(C.to_GtkMenu(v.Widget), bool2gboolean(b))
}

// void gtk_menu_attach_to_widget (GtkMenu *menu, GtkWidget *attach_widget, GtkMenuDetachFunc detacher);

func (v *GtkMenu) Detach() {
	C.gtk_menu_detach(C.to_GtkMenu(v.Widget))
}
func (v *GtkMenu) GetAttachWidget() *GtkWidget {
	return &GtkWidget{C.gtk_menu_get_attach_widget(C.to_GtkMenu(v.Widget))}
}

// GList* gtk_menu_get_for_attach_widget(GtkWidget *widget);

type GtkMenuPositionFunc func(menu *GtkMenu, px, py *int, push_in *bool, data interface{})
type GtkMenuPositionFuncInfo struct {
	menu *GtkMenu
	f    GtkMenuPositionFunc
	data interface{}
}

//export _go_gtk_menu_position_func
func _go_gtk_menu_position_func(gmpfi *C._gtk_menu_position_func_info) {
	if gmpfi == nil {
		return
	}
	gmpfigo := (*GtkMenuPositionFuncInfo)(gmpfi.data)
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
type GtkPackDirection int

const (
	GTK_PACK_DIRECTION_LTR GtkPackDirection = 0
	GTK_PACK_DIRECTION_RTL GtkPackDirection = 1
	GTK_PACK_DIRECTION_TTB GtkPackDirection = 2
	GTK_PACK_DIRECTION_BTT GtkPackDirection = 3
)

type GtkMenuBar struct {
	GtkWidget
}

func MenuBar() *GtkMenuBar {
	return &GtkMenuBar{GtkWidget{C.gtk_menu_bar_new()}}
}
func (v *GtkMenuBar) SetPackDirection(pack_dir GtkPackDirection) {
	C.gtk_menu_bar_set_pack_direction(C.to_GtkMenuBar(v.Widget), C.GtkPackDirection(pack_dir))
}
func (v *GtkMenuBar) GetPackDirection() GtkPackDirection {
	return GtkPackDirection(C.gtk_menu_bar_get_pack_direction(C.to_GtkMenuBar(v.Widget)))
}
func (v *GtkMenuBar) SetChildPackDirection(pack_dir GtkPackDirection) {
	C.gtk_menu_bar_set_child_pack_direction(C.to_GtkMenuBar(v.Widget), C.GtkPackDirection(pack_dir))
}
func (v *GtkMenuBar) GetChildPackDirection() GtkPackDirection {
	return GtkPackDirection(C.gtk_menu_bar_get_child_pack_direction(C.to_GtkMenuBar(v.Widget)))
}

//TODO da rimuovere, creare GtkMenuShell e usarlo come anonymous field per GtkMenu
func (v *GtkMenuBar) Append(child WidgetLike) {
	C.gtk_menu_shell_append(C.to_GtkMenuShell(v.Widget), child.ToNative())
}

//TODO da rimuovere, creare GtkMenuShell e usarlo come anonymous field per GtkMenu
func (v *GtkMenuBar) Prepend(child WidgetLike) {
	C.gtk_menu_shell_prepend(C.to_GtkMenuShell(v.Widget), child.ToNative())
}

//TODO da rimuovere, creare GtkMenuShell e usarlo come anonymous field per GtkMenu
func (v *GtkMenuBar) Insert(child WidgetLike, position int) {
	C.gtk_menu_shell_insert(C.to_GtkMenuShell(v.Widget), child.ToNative(), C.gint(position))
}

//-----------------------------------------------------------------------
// GtkMenuItem
//-----------------------------------------------------------------------
type GtkMenuItem struct {
	GtkItem
}

func MenuItem() *GtkMenuItem {
	return &GtkMenuItem{GtkItem{GtkBin{GtkContainer{GtkWidget{
		C.gtk_menu_item_new()}}}}}
}
func MenuItemWithLabel(label string) *GtkMenuItem {
	ptr := C.CString(label)
	defer C.free_string(ptr)
	return &GtkMenuItem{GtkItem{GtkBin{GtkContainer{GtkWidget{
		C.gtk_menu_item_new_with_label(C.to_gcharptr(ptr))}}}}}
}
func MenuItemWithMnemonic(label string) *GtkMenuItem {
	ptr := C.CString(label)
	defer C.free_string(ptr)
	return &GtkMenuItem{GtkItem{GtkBin{GtkContainer{GtkWidget{
		C.gtk_menu_item_new_with_mnemonic(C.to_gcharptr(ptr))}}}}}
}
func (v *GtkMenuItem) SetRightJustified(b bool) {
	C.gtk_menu_item_set_right_justified(C.to_GtkMenuItem(v.Widget), bool2gboolean(b))
}
func (v *GtkMenuItem) GetRightJustified() bool {
	return gboolean2bool(C.gtk_menu_item_get_right_justified(C.to_GtkMenuItem(v.Widget)))
}

// G_CONST_RETURN gchar *gtk_menu_item_get_label(GtkMenuItem *menu_item);
// void gtk_menu_item_set_label(GtkMenuItem *menu_item, const gchar *label);

func (v *GtkMenuItem) GetUseUnderline() bool {
	return gboolean2bool(C.gtk_menu_item_get_use_underline(C.to_GtkMenuItem(v.Widget)))
}
func (v *GtkMenuItem) SetUseUnderline(setting bool) {
	C.gtk_menu_item_set_use_underline(C.to_GtkMenuItem(v.Widget), bool2gboolean(setting))
}
func (v *GtkMenuItem) SetSubmenu(w WidgetLike) {
	C.gtk_menu_item_set_submenu(C.to_GtkMenuItem(v.Widget), w.ToNative())
}
func (v *GtkMenuItem) GetSubmenu() *GtkWidget {
	return &GtkWidget{C.gtk_menu_item_get_submenu(C.to_GtkMenuItem(v.Widget))}
}

//Deprecated since 2.12. Use SetSubmenu() instead.
func (v *GtkMenuItem) RemoveSubmenu() {
	deprecated_since(2, 12, 0, "gtk_menu_item_remove_submenu()")
	C.gtk_menu_item_remove_submenu(C.to_GtkMenuItem(v.Widget))
}

// void gtk_menu_item_set_accel_path(GtkMenuItem *menu_item, const gchar *accel_path);
// G_CONST_RETURN gchar* gtk_menu_item_get_accel_path(GtkMenuItem *menu_item);

func (v *GtkMenuItem) Select() {
	C.gtk_menu_item_select(C.to_GtkMenuItem(v.Widget))
}
func (v *GtkMenuItem) Deselect() {
	C.gtk_menu_item_deselect(C.to_GtkMenuItem(v.Widget))
}
func (v *GtkMenuItem) Activate() {
	C.gtk_menu_item_activate(C.to_GtkMenuItem(v.Widget))
}
func (v *GtkMenuItem) ToggleSizeRequest(i *int) {
	gi := C.gint(*i)
	C.gtk_menu_item_toggle_size_request(C.to_GtkMenuItem(v.Widget), &gi)
	*i = int(gi)
}
func (v *GtkMenuItem) ToggleSizeAllocate(i int) {
	C.gtk_menu_item_toggle_size_allocate(C.to_GtkMenuItem(v.Widget), C.gint(i))
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
type GtkRadioMenuItem struct {
	GtkCheckMenuItem
}

func RadioMenuItem(group *glib.SList) *GtkRadioMenuItem {
	if group != nil {
		return &GtkRadioMenuItem{GtkCheckMenuItem{GtkMenuItem{GtkItem{GtkBin{GtkContainer{GtkWidget{
			C.gtk_radio_menu_item_new(C.to_gslist(unsafe.Pointer(group.ToSList())))}}}}}}}
	}
	return &GtkRadioMenuItem{GtkCheckMenuItem{GtkMenuItem{GtkItem{GtkBin{GtkContainer{GtkWidget{
		C.gtk_radio_menu_item_new(nil)}}}}}}}
}

func RadioMenuItemWithLabel(group *glib.SList, label string) *GtkRadioMenuItem {
	ptr := C.CString(label)
	defer C.free_string(ptr)
	if group != nil {
		return &GtkRadioMenuItem{GtkCheckMenuItem{GtkMenuItem{GtkItem{GtkBin{GtkContainer{GtkWidget{
			C.gtk_radio_menu_item_new_with_label(C.to_gslist(unsafe.Pointer(group.ToSList())), C.to_gcharptr(ptr))}}}}}}}
	}
	return &GtkRadioMenuItem{GtkCheckMenuItem{GtkMenuItem{GtkItem{GtkBin{GtkContainer{GtkWidget{
		C.gtk_radio_menu_item_new_with_label(nil, C.to_gcharptr(ptr))}}}}}}}
}

// gtk_radio_menu_item_new_with_mnemonic
// gtk_radio_menu_item_new_from_widget
// gtk_radio_menu_item_new_with_label_from_widget
// gtk_radio_menu_item_new_with_mnemonic_from_widget
// gtk_radio_menu_item_group

func (v *GtkRadioMenuItem) SetGroup(group *glib.SList) {
	if group != nil {
		C.gtk_radio_menu_item_set_group(C.to_GtkRadioMenuItem(v.Widget), C.to_gslist(unsafe.Pointer(group)))
	} else {
		C.gtk_radio_menu_item_set_group(C.to_GtkRadioMenuItem(v.Widget), nil)
	}
}

func (v *GtkRadioMenuItem) GetGroup() *glib.SList {
	return glib.SListFromNative(unsafe.Pointer(C.gtk_radio_menu_item_get_group(C.to_GtkRadioMenuItem(v.Widget))))
}

//-----------------------------------------------------------------------
// GtkCheckMenuItem
//-----------------------------------------------------------------------
type GtkCheckMenuItem struct {
	GtkMenuItem
}

func CheckMenuItem() *GtkCheckMenuItem {
	return &GtkCheckMenuItem{GtkMenuItem{GtkItem{GtkBin{GtkContainer{GtkWidget{
		C.gtk_check_menu_item_new()}}}}}}
}
func CheckMenuItemWithLabel(label string) *GtkCheckMenuItem {
	ptr := C.CString(label)
	defer C.free_string(ptr)
	return &GtkCheckMenuItem{GtkMenuItem{GtkItem{GtkBin{GtkContainer{GtkWidget{
		C.gtk_check_menu_item_new_with_label(C.to_gcharptr(ptr))}}}}}}
}
func CheckMenuItemWithMnemonic(label string) *GtkCheckMenuItem {
	ptr := C.CString(label)
	defer C.free_string(ptr)
	return &GtkCheckMenuItem{GtkMenuItem{GtkItem{GtkBin{GtkContainer{GtkWidget{
		C.gtk_check_menu_item_new_with_mnemonic(C.to_gcharptr(ptr))}}}}}}
}
func (v *GtkCheckMenuItem) GetActive() bool {
	return gboolean2bool(C.gtk_check_menu_item_get_active(C.to_GtkCheckMenuItem(v.Widget)))
}
func (v *GtkCheckMenuItem) SetActive(setting bool) {
	C.gtk_check_menu_item_set_active(C.to_GtkCheckMenuItem(v.Widget), bool2gboolean(setting))
}
func (v *GtkCheckMenuItem) Toggled() {
	C.gtk_check_menu_item_toggled(C.to_GtkCheckMenuItem(v.Widget))
}
func (v *GtkCheckMenuItem) GetInconsistent() bool {
	return gboolean2bool(C.gtk_check_menu_item_get_inconsistent(C.to_GtkCheckMenuItem(v.Widget)))
}
func (v *GtkCheckMenuItem) SetInconsistent(setting bool) {
	C.gtk_check_menu_item_set_inconsistent(C.to_GtkCheckMenuItem(v.Widget), bool2gboolean(setting))
}
func (v *GtkCheckMenuItem) SetDrawAsRadio(setting bool) {
	C.gtk_check_menu_item_set_draw_as_radio(C.to_GtkCheckMenuItem(v.Widget), bool2gboolean(setting))
}
func (v *GtkCheckMenuItem) GetDrawAsRadio() bool {
	return gboolean2bool(C.gtk_check_menu_item_get_draw_as_radio(C.to_GtkCheckMenuItem(v.Widget)))
}

//-----------------------------------------------------------------------
// GtkSeparatorMenuItem
//-----------------------------------------------------------------------
type GtkSeparatorMenuItem struct {
	GtkMenuItem
}

func SeparatorMenuItem() *GtkSeparatorMenuItem {
	return &GtkSeparatorMenuItem{GtkMenuItem{GtkItem{GtkBin{GtkContainer{GtkWidget{
		C.gtk_separator_menu_item_new()}}}}}}
}

//-----------------------------------------------------------------------
// GtkTearoffMenuItem
//-----------------------------------------------------------------------
type GtkTearoffMenuItem struct {
	GtkMenuItem
}

func TearoffMenuItem() *GtkTearoffMenuItem {
	return &GtkTearoffMenuItem{GtkMenuItem{GtkItem{GtkBin{GtkContainer{GtkWidget{
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

// gtk_toolbar_new
// gtk_toolbar_insert
// gtk_toolbar_get_item_index
// gtk_toolbar_get_n_items
// gtk_toolbar_get_nth_item
// gtk_toolbar_get_drop_index
// gtk_toolbar_set_drop_highlight_item
// gtk_toolbar_set_show_arrow
// gtk_toolbar_set_orientation
// gtk_toolbar_set_tooltips
// gtk_toolbar_unset_icon_size
// gtk_toolbar_get_show_arrow
// gtk_toolbar_get_orientation
// gtk_toolbar_get_style
// gtk_toolbar_get_icon_size
// gtk_toolbar_get_tooltips
// gtk_toolbar_get_relief_style
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
// gtk_toolbar_set_style
// gtk_toolbar_insert_stock
// gtk_toolbar_set_icon_size
// gtk_toolbar_remove_space
// gtk_toolbar_unset_style

//-----------------------------------------------------------------------
// GtkToolItem
//-----------------------------------------------------------------------

// gtk_tool_item_new
// gtk_tool_item_set_homogeneous
// gtk_tool_item_get_homogeneous
// gtk_tool_item_set_expand
// gtk_tool_item_get_expand
// gtk_tool_item_set_tooltip
// gtk_tool_item_set_tooltip_text
// gtk_tool_item_set_tooltip_markup
// gtk_tool_item_set_use_drag_window
// gtk_tool_item_get_use_drag_window
// gtk_tool_item_set_visible_horizontal
// gtk_tool_item_get_visible_horizontal
// gtk_tool_item_set_visible_vertical
// gtk_tool_item_get_visible_vertical
// gtk_tool_item_set_is_important
// gtk_tool_item_get_is_important
// gtk_tool_item_get_ellipsize_mode
// gtk_tool_item_get_icon_size
// gtk_tool_item_get_orientation
// gtk_tool_item_get_toolbar_style
// gtk_tool_item_get_relief_style
// gtk_tool_item_get_text_alignment
// gtk_tool_item_get_text_orientation
// gtk_tool_item_retrieve_proxy_menu_item
// gtk_tool_item_get_proxy_menu_item
// gtk_tool_item_set_proxy_menu_item
// gtk_tool_item_rebuild_menu
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

// gtk_separator_tool_item_new
// gtk_separator_tool_item_set_draw
// gtk_separator_tool_item_get_draw

//-----------------------------------------------------------------------
// GtkToolButton
//-----------------------------------------------------------------------

// gtk_tool_button_new
// gtk_tool_button_new_from_stock
// gtk_tool_button_set_label
// gtk_tool_button_get_label
// gtk_tool_button_set_use_underline
// gtk_tool_button_get_use_underline
// gtk_tool_button_set_stock_id
// gtk_tool_button_get_stock_id
// gtk_tool_button_set_icon_name
// gtk_tool_button_get_icon_name
// gtk_tool_button_set_icon_widget
// gtk_tool_button_get_icon_widget
// gtk_tool_button_set_label_widget
// gtk_tool_button_get_label_widget

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

// gtk_toggle_tool_button_new
// gtk_toggle_tool_button_new_from_stock
// gtk_toggle_tool_button_set_active
// gtk_toggle_tool_button_get_active

//-----------------------------------------------------------------------
// GtkRadioToolButton
//-----------------------------------------------------------------------

// gtk_radio_tool_button_new
// gtk_radio_tool_button_new_from_stock
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
type GtkFileChooserAction int

const (
	GTK_FILE_CHOOSER_ACTION_OPEN          GtkFileChooserAction = 0
	GTK_FILE_CHOOSER_ACTION_SAVE          GtkFileChooserAction = 1
	GTK_FILE_CHOOSER_ACTION_SELECT_FOLDER GtkFileChooserAction = 2
	GTK_FILE_CHOOSER_ACTION_CREATE_FOLDER GtkFileChooserAction = 3
)

type GtkFileChooser struct {
	w *C.GtkFileChooser
}

func (v *GtkFileChooser) SetAction(action GtkFileChooserAction) {
	C.gtk_file_chooser_set_action(v.w, C.GtkFileChooserAction(action))
}
func (v *GtkFileChooser) GetAction() GtkFileChooserAction {
	return GtkFileChooserAction(C.gtk_file_chooser_get_action(v.w))
}
func (v *GtkFileChooser) SetLocalOnly(b bool) {
	C.gtk_file_chooser_set_local_only(v.w, bool2gboolean(b))
}
func (v *GtkFileChooser) GetLocalOnly() bool {
	return gboolean2bool(C.gtk_file_chooser_get_local_only(v.w))
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

func (v *GtkFileChooser) GetFilename() string {
	return C.GoString(C.to_charptr(C.gtk_file_chooser_get_filename(v.w)))
}
func (v *GtkFileChooser) SetFilename(filename string) {
	ptr := C.CString(filename)
	defer C.free_string(ptr)
	C.gtk_file_chooser_set_filename(v.w, ptr)
}

// gboolean gtk_file_chooser_select_filename(GtkFileChooser* chooser, const char* filename);
// void gtk_file_chooser_unselect_filename(GtkFileChooser* chooser, const char* filename);
// void gtk_file_chooser_select_all(GtkFileChooser* chooser);
// void gtk_file_chooser_unselect_all(GtkFileChooser* chooser);
// GSList*  gtk_file_chooser_get_filenames(GtkFileChooser* chooser);

func (v *GtkFileChooser) SetCurrentFolder(f string) bool {
	cf := C.CString(f)
	defer C.free_string(cf)
	return gboolean2bool(C.gtk_file_chooser_set_current_folder(v.w, C.to_gcharptr(cf)))
}
func (v *GtkFileChooser) GetCurrentFolder() string {
	return C.GoString(C.to_charptr(C.gtk_file_chooser_get_current_folder(v.w)))
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

func (v *GtkFileChooser) AddFilter(filter *GtkFileFilter) {
	C.gtk_file_chooser_add_filter(v.w, filter.FileFilter)
}
func (v *GtkFileChooser) RemoveFilter(filter *GtkFileFilter) {
	C.gtk_file_chooser_remove_filter(v.w, filter.FileFilter)
}
func (v *GtkFileChooser) ListFilters() []*GtkFileFilter {
	c_list := C.gtk_file_chooser_list_filters(v.w)
	defer C.g_slist_free(c_list)
	n := int(C.g_slist_length(c_list))
	ret := make([]*GtkFileFilter, n)
	for i := 0; i < n; i++ {
		ret[i] = &GtkFileFilter{C.to_GtkFileFilter(C.g_slist_nth_data(c_list, C.guint(i)))}
	}
	return ret
}
func (v *GtkFileChooser) SetFilter(filter *GtkFileFilter) {
	C.gtk_file_chooser_set_filter(v.w, filter.FileFilter)
}
func (v *GtkFileChooser) GetFilter() *GtkFileFilter {
	return &GtkFileFilter{C.gtk_file_chooser_get_filter(v.w)}
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
type GtkFileChooserButton struct {
	GtkHBox
	GtkFileChooser
}

// gtk_file_chooser_button_new
func FileChooserButton(title string, action GtkFileChooserAction) *GtkFileChooserButton {
	ptitle := C.CString(title)
	defer C.free_string(ptitle)
	widget := GtkWidget{
		C.gtk_file_chooser_button_new(C.to_gcharptr(ptitle), C.GtkFileChooserAction(action)),
	}
	return &GtkFileChooserButton{
		GtkHBox{GtkBox{GtkContainer{widget}}},
		GtkFileChooser{C.to_GtkFileChooser(widget.Widget)},
	}
}

// gtk_file_chooser_button_new_with_backend

// gtk_file_chooser_button_new_with_dialog
func FileChooserButtonWithDialog(dialog *GtkFileChooserDialog) *GtkFileChooserButton {
	widget := GtkWidget{
		C.gtk_file_chooser_button_new_with_dialog(dialog.Widget),
	}
	return &GtkFileChooserButton{
		GtkHBox{GtkBox{GtkContainer{widget}}},
		GtkFileChooser{C.to_GtkFileChooser(widget.Widget)},
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
type GtkFileChooserDialog struct {
	GtkDialog
	GtkFileChooser
}

func FileChooserDialog(title string, parent *GtkWindow, file_chooser_action GtkFileChooserAction, button_text string, button_action GtkResponseType, buttons ...interface{}) *GtkFileChooserDialog {
	ptitle := C.CString(title)
	defer C.free_string(ptitle)
	pbutton := C.CString(button_text)
	defer C.free_string(pbutton)
	widget := GtkWidget{
		C._gtk_file_chooser_dialog_new(
			C.to_gcharptr(ptitle),
			parent.ToNative(),
			C.int(file_chooser_action),
			C.int(button_action),
			C.to_gcharptr(pbutton))}
	ret := &GtkFileChooserDialog{
		GtkDialog{GtkWindow{GtkBin{GtkContainer{widget}}}},
		GtkFileChooser{C.to_GtkFileChooser(widget.Widget)}}
	text, res := variadicButtonsToArrays(buttons)
	for i := range text {
		ret.AddButton(text[i], res[i])
	}
	return ret
}

//-----------------------------------------------------------------------
// GtkFileChooserWidget
//-----------------------------------------------------------------------
type GtkFileChooserWidget struct {
	GtkVBox
	GtkFileChooser
}

func FileChooserWidget(file_chooser_action GtkFileChooserAction) *GtkFileChooserWidget {
	widget := GtkWidget{C._gtk_file_chooser_widget_new(C.int(file_chooser_action))}
	return &GtkFileChooserWidget{
		GtkVBox{GtkBox{GtkContainer{widget}}},
		GtkFileChooser{C.to_GtkFileChooser(widget.Widget)},
	}
}

// gtk_file_chooser_widget_new_with_backend

//-----------------------------------------------------------------------
// GtkFileFilter
//-----------------------------------------------------------------------
type GtkFileFilter struct {
	FileFilter *C.GtkFileFilter
}

func FileFilter() *GtkFileFilter {
	return &GtkFileFilter{C.gtk_file_filter_new()}
}
func (v *GtkFileFilter) SetName(name string) {
	ptr := C.CString(name)
	defer C.free_string(ptr)
	C.gtk_file_filter_set_name(v.FileFilter, C.to_gcharptr(ptr))
}
func (v *GtkFileFilter) GetName() string {
	return C.GoString(C.to_charptr(C.gtk_file_filter_get_name(v.FileFilter)))
}
func (v *GtkFileFilter) AddMimeType(mimetype string) {
	ptr := C.CString(mimetype)
	defer C.free_string(ptr)
	C.gtk_file_filter_add_mime_type(v.FileFilter, C.to_gcharptr(ptr))
}
func (v *GtkFileFilter) AddPattern(pattern string) {
	ptr := C.CString(pattern)
	defer C.free_string(ptr)
	C.gtk_file_filter_add_pattern(v.FileFilter, C.to_gcharptr(ptr))
}

//void gtk_file_filter_add_pixbuf_formats (GtkFileFilter *filter);
//void gtk_file_filter_add_custom (GtkFileFilter *filter, GtkFileFilterFlags needed, GtkFileFilterFunc func, gpointer data, GDestroyNotify notify);
// gtk_file_filter_get_needed  //for use in the implementation of GtkFileChooser
// gtk_file_filter_filter  //for use in the implementation of GtkFileChooser

//-----------------------------------------------------------------------
// GtkFontButton
//-----------------------------------------------------------------------
type GtkFontButton struct {
	GtkButton
}

func FontButton() *GtkFontButton {
	return &GtkFontButton{GtkButton{GtkBin{GtkContainer{GtkWidget{
		C.gtk_font_button_new()}}}}}
}
func FontButtonWithFont(fontname string) *GtkFontButton {
	ptr := C.CString(fontname)
	defer C.free_string(ptr)
	return &GtkFontButton{GtkButton{GtkBin{GtkContainer{GtkWidget{
		C.gtk_font_button_new_with_font(C.to_gcharptr(ptr))}}}}}
}
func (v *GtkFontButton) SetFontName(fontname string) {
	ptr := C.CString(fontname)
	defer C.free_string(ptr)
	C.gtk_font_button_set_font_name(C.to_GtkFontButton(v.Widget), C.to_gcharptr(ptr))
}
func (v *GtkFontButton) GetFontName() string {
	return C.GoString(C.to_charptr(C.gtk_font_button_get_font_name(C.to_GtkFontButton(v.Widget))))
}

// gtk_font_button_set_show_style
// gtk_font_button_get_show_style

func (v *GtkFontButton) SetShowSize(show_size bool) {
	C.gtk_font_button_set_show_size(C.to_GtkFontButton(v.Widget), bool2gboolean(show_size))
}
func (v *GtkFontButton) GetShowSize() bool {
	return gboolean2bool(C.gtk_font_button_get_show_size(C.to_GtkFontButton(v.Widget)))
}

// gtk_font_button_set_use_font
// gtk_font_button_get_use_font

func (v *GtkFontButton) SetUseSize(use_size bool) {
	C.gtk_font_button_set_use_size(C.to_GtkFontButton(v.Widget), bool2gboolean(use_size))
}
func (v *GtkFontButton) GetUseSize() bool {
	return gboolean2bool(C.gtk_font_button_get_use_size(C.to_GtkFontButton(v.Widget)))
}
func (v *GtkFontButton) SetTitle(title string) {
	ptr := C.CString(title)
	defer C.free_string(ptr)
	C.gtk_font_button_set_title(C.to_GtkFontButton(v.Widget), C.to_gcharptr(ptr))
}
func (v *GtkFontButton) GetTitle() string {
	return C.GoString(C.to_charptr(C.gtk_font_button_get_title(C.to_GtkFontButton(v.Widget))))
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
type GtkFontSelectionDialog struct {
	GtkDialog
}

func FontSelectionDialog(title string) *GtkFontSelectionDialog {
	ptitle := C.CString(title)
	defer C.free_string(ptitle)
	return &GtkFontSelectionDialog{GtkDialog{GtkWindow{GtkBin{GtkContainer{GtkWidget{
		C.gtk_font_selection_dialog_new(C.to_gcharptr(ptitle))}}}}}}
}
func (v *GtkFontSelectionDialog) GetFontName() string {
	return C.GoString(C.to_charptr(C.gtk_font_selection_dialog_get_font_name(C.to_GtkFontSelectionDialog(v.Widget))))
}
func (v *GtkFontSelectionDialog) SetFontName(font string) {
	pfont := C.CString(font)
	defer C.free_string(pfont)
	C.gtk_font_selection_dialog_set_font_name(C.to_GtkFontSelectionDialog(v.Widget), C.to_gcharptr(pfont))
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
type GtkAlignment struct {
	GtkBin
}

func Alignment(xalign float64, yalign float64, xscale float64, yscale float64) *GtkAlignment {
	return &GtkAlignment{GtkBin{GtkContainer{GtkWidget{
		C.gtk_alignment_new(C.gfloat(xalign), C.gfloat(yalign), C.gfloat(xscale), C.gfloat(yscale))}}}}
}
func (v *GtkAlignment) Set(xalign float64, yalign float64, xscale float64, yscale float64) {
	C.gtk_alignment_set(C.to_GtkAlignment(v.Widget), C.gfloat(xalign), C.gfloat(yalign), C.gfloat(xscale), C.gfloat(yscale))
}
func (v *GtkAlignment) SetPadding(padding_top uint, padding_bottom uint, padding_left uint, padding_right uint) {
	C.gtk_alignment_set_padding(C.to_GtkAlignment(v.Widget), C.guint(padding_top), C.guint(padding_bottom), C.guint(padding_left), C.guint(padding_right))
}
func (v *GtkAlignment) GetPadding() (padding_top uint, padding_bottom uint, padding_left uint, padding_right uint) {
	var cpadding_top, cpadding_bottom, cpadding_left, cpadding_right C.guint
	C.gtk_alignment_get_padding(C.to_GtkAlignment(v.Widget), &cpadding_top, &cpadding_bottom, &cpadding_left, &cpadding_right)
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
type GtkHBox struct {
	GtkBox
}

func HBox(homogeneous bool, spacing uint) *GtkHBox {
	return &GtkHBox{GtkBox{GtkContainer{GtkWidget{
		C.gtk_hbox_new(bool2gboolean(homogeneous), C.gint(spacing))}}}}
}

//-----------------------------------------------------------------------
// GtkVBox
//-----------------------------------------------------------------------
type GtkVBox struct {
	GtkBox
}

func VBox(homogeneous bool, spacing uint) *GtkVBox {
	return &GtkVBox{GtkBox{GtkContainer{GtkWidget{
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
type GtkFixed struct {
	GtkContainer
}

func Fixed() *GtkFixed {
	return &GtkFixed{GtkContainer{GtkWidget{C.gtk_fixed_new()}}}
}
func (v *GtkFixed) Put(w WidgetLike, x, y int) {
	C.gtk_fixed_put(C.to_GtkFixed(v.Widget), w.ToNative(), C.gint(x), C.gint(y))
}
func (v *GtkFixed) Move(w WidgetLike, x, y int) {
	C.gtk_fixed_move(C.to_GtkFixed(v.Widget), w.ToNative(), C.gint(x), C.gint(y))
}

//Deprecated since 2.20. Use GtkWidget.GetHasWindow() instead.
/*GtkFixed gets GetHasWindow() from anonymous field so this method can be commented out.
func (v *GtkFixed) GetHasWindow() bool {
	deprecated_since(2,20,0,"gtk_fixed_get_has_window()")
	return gboolean2bool(C.gtk_fixed_get_has_window(C.to_GtkFixed(v.Widget)))
}*/
//Deprecated since 2.20. Use GtkWidget.SetHasWindow() instead.
/*GtkFixed gets SetHasWindow() from anonymous field so this method can be commented out.
func (v *GtkFixed) SetHasWindow(has_window bool) {
	deprecated_since(2,20,0,"gtk_fixed_set_has_window()")
	C.gtk_fixed_set_has_window(C.to_GtkFixed(v.Widget), bool2gboolean(has_window))
}*/

//-----------------------------------------------------------------------
// GtkHPaned
//-----------------------------------------------------------------------
type GtkHPaned struct {
	GtkPaned
}

func HPaned() *GtkHPaned {
	return &GtkHPaned{GtkPaned{GtkContainer{GtkWidget{
		C.gtk_hpaned_new()}}}}
}

//-----------------------------------------------------------------------
// GtkVPaned
//-----------------------------------------------------------------------
type GtkVPaned struct {
	GtkPaned
}

func VPaned() *GtkVPaned {
	return &GtkVPaned{GtkPaned{GtkContainer{GtkWidget{
		C.gtk_vpaned_new()}}}}
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
type GtkNotebook struct {
	GtkContainer
}

func Notebook() *GtkNotebook {
	return &GtkNotebook{GtkContainer{GtkWidget{
		C.gtk_notebook_new()}}}
}
func (v *GtkNotebook) AppendPage(child WidgetLike, tab_label WidgetLike) int {
	return int(C.gtk_notebook_append_page(C.to_GtkNotebook(v.Widget), child.ToNative(), tab_label.ToNative()))
}
func (v *GtkNotebook) AppendPageMenu(child WidgetLike, tab_label WidgetLike, menu_label WidgetLike) int {
	return int(C.gtk_notebook_append_page_menu(C.to_GtkNotebook(v.Widget), child.ToNative(), tab_label.ToNative(), menu_label.ToNative()))
}
func (v *GtkNotebook) PrependPage(child WidgetLike, tab_label WidgetLike) int {
	return int(C.gtk_notebook_prepend_page(C.to_GtkNotebook(v.Widget), child.ToNative(), tab_label.ToNative()))
}
func (v *GtkNotebook) PrependPageMenu(child WidgetLike, tab_label WidgetLike, menu_label WidgetLike) int {
	return int(C.gtk_notebook_prepend_page_menu(C.to_GtkNotebook(v.Widget), child.ToNative(), tab_label.ToNative(), menu_label.ToNative()))
}
func (v *GtkNotebook) InsertPage(child WidgetLike, tab_label WidgetLike, position int) int {
	return int(C.gtk_notebook_insert_page(C.to_GtkNotebook(v.Widget), child.ToNative(), tab_label.ToNative(), C.gint(position)))
}
func (v *GtkNotebook) InsertPageMenu(child WidgetLike, tab_label WidgetLike, menu_label WidgetLike, position int) int {
	return int(C.gtk_notebook_insert_page_menu(C.to_GtkNotebook(v.Widget), child.ToNative(), tab_label.ToNative(), menu_label.ToNative(), C.gint(position)))
}
func (v *GtkNotebook) RemovePage(child WidgetLike, page_num int) {
	C.gtk_notebook_remove_page(C.to_GtkNotebook(v.Widget), C.gint(page_num))
}
func (v *GtkNotebook) PageNum(child WidgetLike) int {
	return int(C.gtk_notebook_page_num(C.to_GtkNotebook(v.Widget), child.ToNative()))
}
func (v *GtkNotebook) NextPage() {
	C.gtk_notebook_next_page(C.to_GtkNotebook(v.Widget))
}
func (v *GtkNotebook) PrevPage() {
	C.gtk_notebook_prev_page(C.to_GtkNotebook(v.Widget))
}
func (v *GtkNotebook) ReorderChild(child WidgetLike, position int) {
	C.gtk_notebook_reorder_child(C.to_GtkNotebook(v.Widget), child.ToNative(), C.gint(position))
}
func (v *GtkNotebook) SetTabPos(pos GtkPositionType) {
	C.gtk_notebook_set_tab_pos(C.to_GtkNotebook(v.Widget), C.GtkPositionType(pos))
}
func (v *GtkNotebook) SetShowTabs(show_tabs bool) {
	C.gtk_notebook_set_show_tabs(C.to_GtkNotebook(v.Widget), bool2gboolean(show_tabs))
}
func (v *GtkNotebook) SetShowBorder(show_border bool) {
	C.gtk_notebook_set_show_border(C.to_GtkNotebook(v.Widget), bool2gboolean(show_border))
}
func (v *GtkNotebook) SetScrollable(scrollable bool) {
	C.gtk_notebook_set_scrollable(C.to_GtkNotebook(v.Widget), bool2gboolean(scrollable))
}

//Deprecated.
func (v *GtkNotebook) SetTabBorder(border_width uint) {
	deprecated_since(2, 0, 0, "gtk_notebook_set_tab_border()")
	C.gtk_notebook_set_tab_border(C.to_GtkNotebook(v.Widget), C.guint(border_width))
}
func (v *GtkNotebook) PopupEnable() {
	C.gtk_notebook_popup_enable(C.to_GtkNotebook(v.Widget))
}
func (v *GtkNotebook) PopupDisable() {
	C.gtk_notebook_popup_disable(C.to_GtkNotebook(v.Widget))
}
func (v *GtkNotebook) GetCurrentPage() int {
	return int(C.gtk_notebook_get_current_page(C.to_GtkNotebook(v.Widget)))
}
func (v *GtkNotebook) GetMenuLabel(child WidgetLike) *GtkWidget {
	return &GtkWidget{
		C.gtk_notebook_get_menu_label(C.to_GtkNotebook(v.Widget), child.ToNative())}
}
func (v *GtkNotebook) GetNthPage(page_num int) *GtkWidget {
	return &GtkWidget{
		C.gtk_notebook_get_nth_page(C.to_GtkNotebook(v.Widget), C.gint(page_num))}
}
func (v *GtkNotebook) GetNPages() int {
	return int(C.gtk_notebook_get_n_pages(C.to_GtkNotebook(v.Widget)))
}
func (v *GtkNotebook) GetTabLabel(child WidgetLike) *GtkWidget {
	return &GtkWidget{
		C.gtk_notebook_get_tab_label(C.to_GtkNotebook(v.Widget), child.ToNative())}
}

//Deprecated since 2.20. Modify the "tab-expand" and "tab-fill" child properties instead.
func (v *GtkNotebook) QueryTabLabelPacking(child WidgetLike, expand *bool, fill *bool, pack_type *uint) {
	deprecated_since(2, 20, 0, "gtk_notebook_query_tab_label_packing()")
	var cexpand, cfill C.gboolean
	var cpack_type C.GtkPackType
	C.gtk_notebook_query_tab_label_packing(C.to_GtkNotebook(v.Widget), child.ToNative(), &cexpand, &cfill, &cpack_type)
	*expand = gboolean2bool(cexpand)
	*fill = gboolean2bool(cfill)
	*pack_type = uint(cpack_type)
}

//Deprecated.
func (v *GtkNotebook) SetHomogeneousTabs(homogeneous bool) {
	deprecated_since(2, 0, 0, "gtk_notebook_set_homogeneous_tabs()")
	C.gtk_notebook_set_homogeneous_tabs(C.to_GtkNotebook(v.Widget), bool2gboolean(homogeneous))
}
func (v *GtkNotebook) SetMenuLabel(child WidgetLike, menu_label WidgetLike) {
	C.gtk_notebook_set_menu_label(C.to_GtkNotebook(v.Widget), child.ToNative(), menu_label.ToNative())
}
func (v *GtkNotebook) SetMenuLabelText(child WidgetLike, name string) {
	ptr := C.CString(name)
	defer C.free_string(ptr)
	C.gtk_notebook_set_menu_label_text(C.to_GtkNotebook(v.Widget), child.ToNative(), C.to_gcharptr(ptr))
}

//Deprecated.
func (v *GtkNotebook) SetTabHBorder(tab_hborder uint) {
	deprecated_since(2, 0, 0, "gtk_notebook_set_tab_hborder()")
	C.gtk_notebook_set_tab_hborder(C.to_GtkNotebook(v.Widget), C.guint(tab_hborder))
}
func (v *GtkNotebook) SetTabLabel(child WidgetLike, tab_label WidgetLike) {
	C.gtk_notebook_set_tab_label(C.to_GtkNotebook(v.Widget), child.ToNative(), tab_label.ToNative())
}

//Deprecated since 2.20. Modify the "tab-expand" and "tab-fill" child properties instead.
func (v *GtkNotebook) SetTabLabelPacking(child WidgetLike, expand bool, fill bool, pack_type uint) {
	deprecated_since(2, 20, 0, "gtk_notebook_set_tab_label_packing()")
	C.gtk_notebook_set_tab_label_packing(C.to_GtkNotebook(v.Widget), child.ToNative(), bool2gboolean(expand), bool2gboolean(fill), C.GtkPackType(pack_type))
}
func (v *GtkNotebook) SetTabLabelText(child WidgetLike, name string) {
	ptr := C.CString(name)
	defer C.free_string(ptr)
	C.gtk_notebook_set_tab_label_text(C.to_GtkNotebook(v.Widget), child.ToNative(), C.to_gcharptr(ptr))
}

//Deprecated.
func (v *GtkNotebook) SetTabVBorder(tab_vborder uint) {
	deprecated_since(2, 0, 0, "gtk_notebook_set_tab_vborder()")
	C.gtk_notebook_set_tab_vborder(C.to_GtkNotebook(v.Widget), C.guint(tab_vborder))
}
func (v *GtkNotebook) SetReorderable(child WidgetLike, reorderable bool) {
	C.gtk_notebook_set_tab_reorderable(C.to_GtkNotebook(v.Widget), child.ToNative(), bool2gboolean(reorderable))
}
func (v *GtkNotebook) SetTabDetachable(child WidgetLike, detachable bool) {
	C.gtk_notebook_set_tab_detachable(C.to_GtkNotebook(v.Widget), child.ToNative(), bool2gboolean(detachable))
}
func (v *GtkNotebook) GetMenuLabelText(child WidgetLike) string {
	return C.GoString(C.to_charptr(C.gtk_notebook_get_menu_label_text(C.to_GtkNotebook(v.Widget), child.ToNative())))
}
func (v *GtkNotebook) GetScrollable() bool {
	return gboolean2bool(C.gtk_notebook_get_scrollable(C.to_GtkNotebook(v.Widget)))
}
func (v *GtkNotebook) GetShowBorder() bool {
	return gboolean2bool(C.gtk_notebook_get_show_border(C.to_GtkNotebook(v.Widget)))
}
func (v *GtkNotebook) GetShowTabs() bool {
	return gboolean2bool(C.gtk_notebook_get_show_tabs(C.to_GtkNotebook(v.Widget)))
}
func (v *GtkNotebook) GetTabLabelText(child WidgetLike) string {
	return C.GoString(C.to_charptr(C.gtk_notebook_get_tab_label_text(C.to_GtkNotebook(v.Widget), child.ToNative())))
}
func (v *GtkNotebook) GetTabPos() uint {
	return uint(C.gtk_notebook_get_tab_pos(C.to_GtkNotebook(v.Widget)))
}
func (v *GtkNotebook) GetTabReorderable(child WidgetLike) bool {
	return gboolean2bool(C.gtk_notebook_get_tab_reorderable(C.to_GtkNotebook(v.Widget), child.ToNative()))
}
func (v *GtkNotebook) GetTabDetachable(child WidgetLike) bool {
	return gboolean2bool(C.gtk_notebook_get_tab_detachable(C.to_GtkNotebook(v.Widget), child.ToNative()))
}

// gtk_notebook_get_tab_hborder //since 2.22
// gtk_notebook_get_tab_vborder //since 2.22

func (v *GtkNotebook) SetCurrentPage(pageNum int) {
	C.gtk_notebook_set_current_page(C.to_GtkNotebook(v.Widget), C.gint(pageNum))
}

//Deprecated since 2.12, use SetGroupName() instead
func (v *GtkNotebook) SetGroupId(group_id int) {
	deprecated_since(2, 12, 0, "gtk_notebook_set_group_id()")
	C.gtk_notebook_set_group_id(C.to_GtkNotebook(v.Widget), C.gint(group_id))
}

//Deprecated since 2.12, use GetGroupName() instead
func (v *GtkNotebook) GetGroupId() int {
	deprecated_since(2, 12, 0, "gtk_notebook_get_group_id()")
	return int(C.gtk_notebook_get_group_id(C.to_GtkNotebook(v.Widget)))
}

//Deprecated since 2.24, use SetGroupName() instead
func (v *GtkNotebook) SetGroup(group unsafe.Pointer) {
	deprecated_since(2, 24, 0, "gtk_notebook_set_group()")
	C.gtk_notebook_set_group(C.to_GtkNotebook(v.Widget), C.gpointer(group))
}

//Deprecated since 2.24, use GetGroupName() instead
func (v *GtkNotebook) GetGroup() unsafe.Pointer {
	deprecated_since(2, 24, 0, "gtk_notebook_get_group()")
	return unsafe.Pointer(C.gtk_notebook_get_group(C.to_GtkNotebook(v.Widget)))
}
func (v *GtkNotebook) SetGroupName(group string) {
	panic_if_version_older(2, 24, 0, "gtk_notebook_set_group_name()")
	ptr := C.CString(group)
	defer C.free_string(ptr)
	C._gtk_notebook_set_group_name(C.to_GtkNotebook(v.Widget), C.to_gcharptr(ptr))
}
func (v *GtkNotebook) GetGroupName() string {
	panic_if_version_older(2, 24, 0, "gtk_notebook_get_group_name()")
	return C.GoString(C.to_charptr(C._gtk_notebook_get_group_name(C.to_GtkNotebook(v.Widget))))
}

// gtk_notebook_set_action_widget //since 2.20
// gtk_notebook_get_action_widget //since 2.20
// void gtk_notebook_set_window_creation_hook (GtkNotebookWindowCreationFunc func, gpointer data, GDestroyNotify destroy); //deprecated in 2.24

//-----------------------------------------------------------------------
// GtkTable
//-----------------------------------------------------------------------
type GtkAttachOptions int

const (
	GTK_EXPAND GtkAttachOptions = 1 << 0
	GTK_SHRINK GtkAttachOptions = 1 << 1
	GTK_FILL   GtkAttachOptions = 1 << 2
)

type GtkTable struct {
	GtkContainer
}

func Table(rows uint, columns uint, homogeneous bool) *GtkTable {
	return &GtkTable{GtkContainer{GtkWidget{
		C.gtk_table_new(C.guint(rows), C.guint(columns), bool2gboolean(homogeneous))}}}
}
func (v *GtkTable) Resize(rows uint, columns uint) {
	C.gtk_table_resize(C.to_GtkTable(v.Widget), C.guint(rows), C.guint(columns))
}
func (v *GtkTable) Attach(child WidgetLike, left_attach uint, right_attach uint, top_attach uint, bottom_attach uint, xoptions GtkAttachOptions, yoptions GtkAttachOptions, xpadding uint, ypadding uint) {
	C.gtk_table_attach(C.to_GtkTable(v.Widget), child.ToNative(), C.guint(left_attach), C.guint(right_attach), C.guint(top_attach), C.guint(bottom_attach), C.GtkAttachOptions(xoptions), C.GtkAttachOptions(yoptions), C.guint(xpadding), C.guint(ypadding))
}
func (v *GtkTable) AttachDefaults(child WidgetLike, left_attach uint, right_attach uint, top_attach uint, bottom_attach uint) {
	C.gtk_table_attach_defaults(C.to_GtkTable(v.Widget), child.ToNative(), C.guint(left_attach), C.guint(right_attach), C.guint(top_attach), C.guint(bottom_attach))
}
func (v *GtkTable) SetRowSpacing(child WidgetLike, row uint, spacing uint) {
	C.gtk_table_set_row_spacing(C.to_GtkTable(v.Widget), C.guint(row), C.guint(spacing))
}
func (v *GtkTable) SetColSpacing(child WidgetLike, column uint, spacing uint) {
	C.gtk_table_set_col_spacing(C.to_GtkTable(v.Widget), C.guint(column), C.guint(spacing))
}
func (v *GtkTable) SetRowSpacings(child WidgetLike, spacing uint) {
	C.gtk_table_set_row_spacings(C.to_GtkTable(v.Widget), C.guint(spacing))
}
func (v *GtkTable) SetColSpacings(child WidgetLike, spacing uint) {
	C.gtk_table_set_col_spacings(C.to_GtkTable(v.Widget), C.guint(spacing))
}
func (v *GtkTable) SetHomogeneous(child WidgetLike, homogeneous bool) {
	C.gtk_table_set_homogeneous(C.to_GtkTable(v.Widget), bool2gboolean(homogeneous))
}
func (v *GtkTable) GetDefaultRowSpacing(child WidgetLike) uint {
	return uint(C.gtk_table_get_default_row_spacing(C.to_GtkTable(v.Widget)))
}
func (v *GtkTable) GetHomogeneous(child WidgetLike) bool {
	return gboolean2bool(C.gtk_table_get_homogeneous(C.to_GtkTable(v.Widget)))
}
func (v *GtkTable) GetRowSpacing(child WidgetLike, row uint) uint {
	return uint(C.gtk_table_get_row_spacing(C.to_GtkTable(v.Widget), C.guint(row)))
}
func (v *GtkTable) GetColSpacing(child WidgetLike, column uint) uint {
	return uint(C.gtk_table_get_col_spacing(C.to_GtkTable(v.Widget), C.guint(column)))
}
func (v *GtkTable) GetDefaultColSpacing(child WidgetLike) uint {
	return uint(C.gtk_table_get_default_col_spacing(C.to_GtkTable(v.Widget)))
}

// gtk_table_get_size //since 2.22

//-----------------------------------------------------------------------
// GtkExpander
//-----------------------------------------------------------------------
type GtkExpander struct {
	GtkBin
}

func Expander(label string) *GtkExpander {
	ptr := C.CString(label)
	defer C.free_string(ptr)
	return &GtkExpander{GtkBin{GtkContainer{GtkWidget{
		C.gtk_expander_new(C.to_gcharptr(ptr))}}}}
}
func ExpanderWithMnemonic(label string) *GtkExpander {
	ptr := C.CString(label)
	defer C.free_string(ptr)
	return &GtkExpander{GtkBin{GtkContainer{GtkWidget{
		C.gtk_expander_new_with_mnemonic(C.to_gcharptr(ptr))}}}}
}

func (v *GtkExpander) SetExpanded(expanded bool) {
	C.gtk_expander_set_expanded(C.to_GtkExpander(v.Widget), bool2gboolean(expanded))
}
func (v *GtkExpander) GetExpanded() bool {
	return gboolean2bool(C.gtk_expander_get_expanded(C.to_GtkExpander(v.Widget)))
}
func (v *GtkExpander) SetSpacing(spacing int) {
	C.gtk_expander_set_spacing(C.to_GtkExpander(v.Widget), C.gint(spacing))
}
func (v *GtkExpander) GetSpacing() int {
	return int(C.gtk_expander_get_spacing(C.to_GtkExpander(v.Widget)))
}
func (v *GtkExpander) SetLabel(label string) {
	ptr := C.CString(label)
	defer C.free_string(ptr)
	C.gtk_expander_set_label(C.to_GtkExpander(v.Widget), C.to_gcharptr(ptr))
}
func (v *GtkExpander) GetLabel() string {
	return C.GoString(C.to_charptr(C.gtk_expander_get_label(C.to_GtkExpander(v.Widget))))
}
func (v *GtkExpander) SetUseUnderline(setting bool) {
	C.gtk_expander_set_use_underline(C.to_GtkExpander(v.Widget), bool2gboolean(setting))
}
func (v *GtkExpander) GetUseUnderline() bool {
	return gboolean2bool(C.gtk_expander_get_use_underline(C.to_GtkExpander(v.Widget)))
}
func (v *GtkExpander) SetUseMarkup(setting bool) {
	C.gtk_expander_set_use_markup(C.to_GtkExpander(v.Widget), bool2gboolean(setting))
}
func (v *GtkExpander) GetUseMarkup() bool {
	return gboolean2bool(C.gtk_expander_get_use_markup(C.to_GtkExpander(v.Widget)))
}
func (v *GtkExpander) SetLabelWidget(label_widget LabelLike) {
	C.gtk_expander_set_label_widget(C.to_GtkExpander(v.Widget), label_widget.ToNative())
}
func (v *GtkExpander) GetLabelWidget() LabelLike {
	return &GtkLabel{GtkWidget{
		C.gtk_expander_get_label_widget(C.to_GtkExpander(v.Widget))}}
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
type GtkShadowType int

const (
	GTK_SHADOW_NONE       GtkShadowType = 0
	GTK_SHADOW_IN         GtkShadowType = 1
	GTK_SHADOW_OUT        GtkShadowType = 2
	GTK_SHADOW_ETCHED_IN  GtkShadowType = 3
	GTK_SHADOW_ETCHED_OUT GtkShadowType = 4
)

type GtkFrame struct {
	GtkBin
}

func Frame(label string) *GtkFrame {
	ptr := C.CString(label)
	defer C.free_string(ptr)
	return &GtkFrame{GtkBin{GtkContainer{GtkWidget{
		C.gtk_frame_new(C.to_gcharptr(ptr))}}}}
}
func (v *GtkFrame) SetLabel(label string) {
	ptr := C.CString(label)
	defer C.free_string(ptr)
	C.gtk_frame_set_label(C.to_GtkFrame(v.Widget), C.to_gcharptr(ptr))
}
func (v *GtkFrame) SetLabelWidget(label_widget LabelLike) {
	C.gtk_frame_set_label_widget(C.to_GtkFrame(v.Widget), label_widget.ToNative())
}
func (v *GtkFrame) SetLabelAlign(xalign, yalign float64) {
	C.gtk_frame_set_label_align(C.to_GtkFrame(v.Widget), C.gfloat(xalign), C.gfloat(yalign))
}
func (v *GtkFrame) SetShadowType(shadow_type GtkShadowType) {
	C.gtk_frame_set_shadow_type(C.to_GtkFrame(v.Widget), C.GtkShadowType(shadow_type))
}
func (v *GtkFrame) GetLabel() string {
	return C.GoString(C.to_charptr(C.gtk_frame_get_label(C.to_GtkFrame(v.Widget))))
}
func (v *GtkFrame) GetLabelAlign() (xalign, yalign float64) {
	var xalign_, yalign_ C.gfloat
	C.gtk_frame_get_label_align(C.to_GtkFrame(v.Widget), &xalign_, &yalign_)
	return float64(xalign_), float64(yalign_)
}
func (v *GtkFrame) GetLabelWidget() LabelLike {
	return &GtkLabel{GtkWidget{
		C.gtk_frame_get_label_widget(C.to_GtkFrame(v.Widget))}}
}
func (v *GtkFrame) GetShadowType() GtkShadowType {
	return GtkShadowType(C.gtk_frame_get_shadow_type(C.to_GtkFrame(v.Widget)))
}

//-----------------------------------------------------------------------
// GtkHSeparator
//-----------------------------------------------------------------------
type GtkHSeparator struct {
	GtkSeparator
}

func HSeparator() *GtkHSeparator {
	return &GtkHSeparator{GtkSeparator{GtkWidget{
		C.gtk_hseparator_new()}}}
}

//-----------------------------------------------------------------------
// GtkVSeparator
//-----------------------------------------------------------------------
type GtkVSeparator struct {
	GtkSeparator
}

func VSeparator() *GtkVSeparator {
	return &GtkVSeparator{GtkSeparator{GtkWidget{
		C.gtk_vseparator_new()}}}
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
type GtkPolicyType int

const (
	GTK_POLICY_ALWAYS    = 0
	GTK_POLICY_AUTOMATIC = 1
	GTK_POLICY_NEVER     = 2
)

type GtkCornerType int

const (
	GTK_CORNER_TOP_LEFT     GtkCornerType = 0
	GTK_CORNER_BOTTOM_LEFT  GtkCornerType = 1
	GTK_CORNER_TOP_RIGHT    GtkCornerType = 2
	GTK_CORNER_BOTTOM_RIGHT GtkCornerType = 3
)

type GtkScrolledWindow struct {
	GtkBin
}

func ScrolledWindow(hadjustment *GtkAdjustment, vadjustment *GtkAdjustment) *GtkScrolledWindow {
	var had, vad *C.GtkAdjustment
	if hadjustment != nil {
		had = hadjustment.Adjustment
	}
	if vadjustment != nil {
		vad = vadjustment.Adjustment
	}
	return &GtkScrolledWindow{GtkBin{GtkContainer{GtkWidget{
		C.gtk_scrolled_window_new(had, vad)}}}}
}
func (v *GtkScrolledWindow) GetHAdjustment() *GtkAdjustment {
	return &GtkAdjustment{
		C.gtk_scrolled_window_get_hadjustment(C.to_GtkScrolledWindow(v.Widget))}
}
func (v *GtkScrolledWindow) GetVAdjustment() *GtkAdjustment {
	return &GtkAdjustment{
		C.gtk_scrolled_window_get_vadjustment(C.to_GtkScrolledWindow(v.Widget))}
}

// gtk_scrolled_window_get_hscrollbar
// gtk_scrolled_window_get_vscrollbar

func (v *GtkScrolledWindow) SetPolicy(hscrollbar_policy GtkPolicyType, vscrollbar_policy GtkPolicyType) {
	C.gtk_scrolled_window_set_policy(C.to_GtkScrolledWindow(v.Widget), C.GtkPolicyType(hscrollbar_policy), C.GtkPolicyType(vscrollbar_policy))
}
func (v *GtkScrolledWindow) AddWithViewPort(w WidgetLike) {
	C.gtk_scrolled_window_add_with_viewport(C.to_GtkScrolledWindow(v.Widget), w.ToNative())
}
func (v *GtkScrolledWindow) SetPlacement(window_placement GtkCornerType) {
	C.gtk_scrolled_window_set_placement(C.to_GtkScrolledWindow(v.Widget), C.GtkCornerType(window_placement))
}
func (v *GtkScrolledWindow) UnsetPlacement() {
	C.gtk_scrolled_window_unset_placement(C.to_GtkScrolledWindow(v.Widget))
}
func (v *GtkScrolledWindow) SetShadowType(typ GtkShadowType) {
	C.gtk_scrolled_window_set_shadow_type(C.to_GtkScrolledWindow(v.Widget), C.GtkShadowType(typ))
}
func (v *GtkScrolledWindow) SetHAdjustment(hadjustment *GtkAdjustment) {
	C.gtk_scrolled_window_set_hadjustment(C.to_GtkScrolledWindow(v.Widget), hadjustment.Adjustment)
}
func (v *GtkScrolledWindow) SetVAdjustment(vadjustment *GtkAdjustment) {
	C.gtk_scrolled_window_set_vadjustment(C.to_GtkScrolledWindow(v.Widget), vadjustment.Adjustment)
}
func (v *GtkScrolledWindow) GetPlacement() GtkCornerType {
	return GtkCornerType(C.gtk_scrolled_window_get_placement(C.to_GtkScrolledWindow(v.Widget)))
}
func (v *GtkScrolledWindow) GetPolicy(hscrollbar_policy *GtkPolicyType, vscrollbar_policy *GtkPolicyType) {
	var chscrollbar_policy, cvscrollbar_policy C.GtkPolicyType
	C.gtk_scrolled_window_get_policy(C.to_GtkScrolledWindow(v.Widget), &chscrollbar_policy, &cvscrollbar_policy)
	*hscrollbar_policy = GtkPolicyType(chscrollbar_policy)
	*vscrollbar_policy = GtkPolicyType(cvscrollbar_policy)
}
func (v *GtkScrolledWindow) GetShadowType() GtkShadowType {
	return GtkShadowType(C.gtk_scrolled_window_get_shadow_type(C.to_GtkScrolledWindow(v.Widget)))
}

//-----------------------------------------------------------------------
// GtkPrintOperation
//-----------------------------------------------------------------------

type GtkPrintOperation struct {
	PrintOperation *C.GtkPrintOperation
}

type GtkPrintOperationResult int

const (
	GTK_PRINT_OPERATION_RESULT_ERROR       GtkPrintOperationResult = 0
	GTK_PRINT_OPERATION_RESULT_APPLY       GtkPrintOperationResult = 1
	GTK_PRINT_OPERATION_RESULT_CANCEL      GtkPrintOperationResult = 2
	GTK_PRINT_OPERATION_RESULT_IN_PROGRESS GtkPrintOperationResult = 3
)

type GtkPrintOperationAction int

const (
	GTK_PRINT_OPERATION_ACTION_PRINT_DIALOG GtkPrintOperationAction = 0
	GTK_PRINT_OPERATION_ACTION_PRINT        GtkPrintOperationAction = 1
	GTK_PRINT_OPERATION_ACTION_PREVIEW      GtkPrintOperationAction = 2
	GTK_PRINT_OPERATION_ACTION_EXPOR        GtkPrintOperationAction = 3
)

func PrintOperation() *GtkPrintOperation {
	return &GtkPrintOperation{C.gtk_print_operation_new()}
}

func (v *GtkPrintOperation) Run(action GtkPrintOperationAction, parent *GtkWindow) (result GtkPrintOperationResult, err error) {
	var gerror *C.GError
	ret := GtkPrintOperationResult(
		C.gtk_print_operation_run(
			v.PrintOperation,
			C.GtkPrintOperationAction(action),
			C.to_GtkWindow(parent.Widget),
			&gerror))
	if gerror != nil {
		err = glib.ErrorFromNative(unsafe.Pointer(gerror))
	}
	return GtkPrintOperationResult(ret), err
}

func (v *GtkPrintOperation) Cancel() {
	C.gtk_print_operation_cancel(v.PrintOperation)
}

func (v *GtkPrintOperation) IsFinished() bool {
	return gboolean2bool(C.gtk_print_operation_is_finished(v.PrintOperation))
}

func (v *GtkPrintOperation) Connect(s string, f interface{}, datas ...interface{}) int {
	return glib.ObjectFromNative(unsafe.Pointer(v.PrintOperation)).Connect(s, f, datas...)
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
type GtkPrintContext struct {
	PrintContext *C.GtkPrintContext
}

func (v *GtkPrintContext) GetCairoContext() *C.cairo_t {
	return C.gtk_print_context_get_cairo_context(v.PrintContext)
}

func (v *GtkPrintContext) SetCairoContext(cairo *C.cairo_t, dpi_x float64, dpi_y float64) {
	C.gtk_print_context_set_cairo_context(v.PrintContext, cairo, C.double(dpi_x), C.double(dpi_y))
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
type GtkAdjustment struct {
	Adjustment *C.GtkAdjustment
}

func Adjustment(value float64, lower float64, upper float64, step_increment float64, page_increment float64, page_size float64) *GtkAdjustment {
	return &GtkAdjustment{
		C.to_GtkAdjustment(C.gtk_adjustment_new(C.gdouble(value), C.gdouble(lower), C.gdouble(upper), C.gdouble(step_increment), C.gdouble(page_increment), C.gdouble(page_size)))}
}
func (v *GtkAdjustment) GetValue() float64 {
	r := C.gtk_adjustment_get_value(v.Adjustment)
	return float64(r)
}
func (v *GtkAdjustment) SetValue(value float64) {
	C.gtk_adjustment_set_value(v.Adjustment, C.gdouble(value))
}

// gtk_adjustment_clamp_page
// gtk_adjustment_changed
// gtk_adjustment_value_changed

func (v *GtkAdjustment) Configure(value float64, lower float64, upper float64, step_increment float64, page_increment float64, page_size float64) {
	panic_if_version_older(2, 14, 0, "gtk_adjustment_configure()")
	C._gtk_adjustment_configure(v.Adjustment, C.gdouble(value), C.gdouble(lower), C.gdouble(upper), C.gdouble(step_increment), C.gdouble(page_increment), C.gdouble(page_size))
}
func (v *GtkAdjustment) GetLower() float64 {
	panic_if_version_older(2, 14, 0, "gtk_adjustment_get_lower()")
	r := C._gtk_adjustment_get_lower(v.Adjustment)
	return float64(r)
}
func (v *GtkAdjustment) GetPageIncrement() float64 {
	panic_if_version_older(2, 14, 0, "gtk_adjustment_get_page_increment()")
	r := C._gtk_adjustment_get_page_increment(v.Adjustment)
	return float64(r)
}
func (v *GtkAdjustment) GetPageSize() float64 {
	panic_if_version_older(2, 14, 0, "gtk_adjustment_get_page_size()")
	r := C._gtk_adjustment_get_page_size(v.Adjustment)
	return float64(r)
}
func (v *GtkAdjustment) GetStepIncrement() float64 {
	panic_if_version_older(2, 14, 0, "gtk_adjustment_get_step_increment()")
	r := C._gtk_adjustment_get_step_increment(v.Adjustment)
	return float64(r)
}
func (v *GtkAdjustment) GetUpper() float64 {
	panic_if_version_older(2, 14, 0, "gtk_adjustment_get_upper()")
	r := C._gtk_adjustment_get_upper(v.Adjustment)
	return float64(r)
}
func (v *GtkAdjustment) SetLower(lower float64) {
	panic_if_version_older(2, 14, 0, "gtk_adjustment_set_lower()")
	C._gtk_adjustment_set_lower(v.Adjustment, C.gdouble(lower))
}
func (v *GtkAdjustment) SetPageIncrement(page_increment float64) {
	panic_if_version_older(2, 14, 0, "gtk_adjustment_set_page_increment()")
	C._gtk_adjustment_set_page_increment(v.Adjustment, C.gdouble(page_increment))
}
func (v *GtkAdjustment) SetPageSize(page_size float64) {
	panic_if_version_older(2, 14, 0, "gtk_adjustment_set_page_size()")
	C._gtk_adjustment_set_page_size(v.Adjustment, C.gdouble(page_size))
}
func (v *GtkAdjustment) SetStepIncrement(step_increment float64) {
	panic_if_version_older(2, 14, 0, "gtk_adjustment_set_step_increment()")
	C._gtk_adjustment_set_step_increment(v.Adjustment, C.gdouble(step_increment))
}
func (v *GtkAdjustment) SetUpper(upper float64) {
	panic_if_version_older(2, 14, 0, "gtk_adjustment_set_upper()")
	C._gtk_adjustment_set_upper(v.Adjustment, C.gdouble(upper))
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
type GtkDrawingArea struct {
	GtkWidget
}

func DrawingArea() *GtkDrawingArea {
	return &GtkDrawingArea{GtkWidget{C.gtk_drawing_area_new()}}
}

//Deprecated. Use GtkWidget.SetSizeRequest() instead.
func (v *GtkDrawingArea) GetSizeRequest(width int, height int) {
	deprecated_since(2, 0, 0, "gtk_drawing_area_size()")
	C.gtk_drawing_area_size(C.to_GtkDrawingArea(v.Widget), C.gint(width), C.gint(height))
}

//-----------------------------------------------------------------------
// GtkEventBox
//-----------------------------------------------------------------------
type GtkEventBox struct {
	GtkBin
}

func EventBox() *GtkEventBox {
	return &GtkEventBox{GtkBin{GtkContainer{GtkWidget{
		C.gtk_event_box_new()}}}}
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
type GtkSizeGroupMode int

const (
	GTK_SIZE_GROUP_NONE       GtkSizeGroupMode = 0
	GTK_SIZE_GROUP_HORIZONTAL GtkSizeGroupMode = 1
	GTK_SIZE_GROUP_VERTICAL   GtkSizeGroupMode = 2
	GTK_SIZE_GROUP_BOTH       GtkSizeGroupMode = 3
)

type GtkSizeGroup struct {
	SizeGroup *C.GtkSizeGroup
}

func SizeGroup(mode GtkSizeGroupMode) *GtkSizeGroup {
	return &GtkSizeGroup{C.gtk_size_group_new(C.GtkSizeGroupMode(mode))}
}
func (v *GtkSizeGroup) SetMode(mode GtkSizeGroupMode) {
	C.gtk_size_group_set_mode(v.SizeGroup, C.GtkSizeGroupMode(mode))
}
func (v *GtkSizeGroup) GetMode() GtkSizeGroupMode {
	return GtkSizeGroupMode(C.gtk_size_group_get_mode(v.SizeGroup))
}
func (v *GtkSizeGroup) SetIgnoreHidden(ignore_hidden bool) {
	C.gtk_size_group_set_ignore_hidden(v.SizeGroup, bool2gboolean(ignore_hidden))
}
func (v *GtkSizeGroup) GetIgnoreHidden() bool {
	return gboolean2bool(C.gtk_size_group_get_ignore_hidden(v.SizeGroup))
}
func (v *GtkSizeGroup) AddWidget(w WidgetLike) {
	C.gtk_size_group_add_widget(v.SizeGroup, w.ToNative())
}
func (v *GtkSizeGroup) RemoveWidget(w WidgetLike) {
	C.gtk_size_group_remove_widget(v.SizeGroup, w.ToNative())
}
func (v *GtkSizeGroup) GetWidgets() *glib.SList {
	return glib.SListFromNative(unsafe.Pointer(C.gtk_size_group_get_widgets(v.SizeGroup)))
}

//-----------------------------------------------------------------------
// GtkTooltip
//-----------------------------------------------------------------------
type GtkTooltip struct {
	Tooltip *C.GtkTooltip
}

func (v *GtkTooltip) SetMarkup(markup string) {
	ptr := C.CString(markup)
	defer C.free_string(ptr)
	C.gtk_tooltip_set_markup(v.Tooltip, C.to_gcharptr(ptr))
}
func (v *GtkTooltip) SetText(text string) {
	ptr := C.CString(text)
	defer C.free_string(ptr)
	C.gtk_tooltip_set_text(v.Tooltip, C.to_gcharptr(ptr))
}
func (v *GtkTooltip) SetIcon(pixbuf *gdkpixbuf.GdkPixbuf) {
	C.gtk_tooltip_set_icon(v.Tooltip, pixbuf.Pixbuf)
}
func (v *GtkTooltip) SetIconFromStock(stock_id string, size GtkIconSize) {
	ptr := C.CString(stock_id)
	defer C.free_string(ptr)
	C.gtk_tooltip_set_icon_from_stock(v.Tooltip, C.to_gcharptr(ptr), C.GtkIconSize(size))
}
func (v *GtkTooltip) SetIconFromIconName(icon_name string, size GtkIconSize) {
	ptr := C.CString(icon_name)
	defer C.free_string(ptr)
	C.gtk_tooltip_set_icon_from_icon_name(v.Tooltip, C.to_gcharptr(ptr), C.GtkIconSize(size))
}

// gtk_tooltip_set_icon_from_gicon //since 2.20
// gtk_tooltip_set_custom
// gtk_tooltip_trigger_tooltip_query
// gtk_tooltip_set_tip_area

func TooltipFromNative(l unsafe.Pointer) *GtkTooltip {
	return &GtkTooltip{(*C.GtkTooltip)(l)}
}

//-----------------------------------------------------------------------
// GtkViewport
//-----------------------------------------------------------------------
type GtkViewport struct {
	GtkBin
}

func Viewport(hadjustment *GtkAdjustment, vadjustment *GtkAdjustment) *GtkViewport {
	var had, vad *C.GtkAdjustment
	if hadjustment != nil {
		had = hadjustment.Adjustment
	}
	if vadjustment != nil {
		vad = vadjustment.Adjustment
	}
	return &GtkViewport{GtkBin{GtkContainer{GtkWidget{
		C.gtk_viewport_new(had, vad)}}}}
}
func (v *GtkViewport) GetHAdjustment() *GtkAdjustment {
	return &GtkAdjustment{
		C.gtk_viewport_get_hadjustment(C.to_GtkViewport(v.Widget))}
}
func (v *GtkViewport) GetVAdjustment() *GtkAdjustment {
	return &GtkAdjustment{
		C.gtk_viewport_get_vadjustment(C.to_GtkViewport(v.Widget))}
}
func (v *GtkViewport) SetHAdjustment(hadjustment *GtkAdjustment) {
	C.gtk_viewport_set_hadjustment(C.to_GtkViewport(v.Widget), hadjustment.Adjustment)
}
func (v *GtkViewport) SetVAdjustment(vadjustment *GtkAdjustment) {
	C.gtk_viewport_set_vadjustment(C.to_GtkViewport(v.Widget), vadjustment.Adjustment)
}
func (v *GtkViewport) GetShadowType() GtkShadowType {
	return GtkShadowType(C.gtk_viewport_get_shadow_type(C.to_GtkViewport(v.Widget)))
}
func (v *GtkViewport) SetShadowType(typ GtkShadowType) {
	C.gtk_viewport_set_shadow_type(C.to_GtkViewport(v.Widget), C.GtkShadowType(typ))
}
func (v *GtkViewport) GetBinWindow() *GtkWindow {
	panic_if_version_older_auto(2, 20, 0)
	return &GtkWindow{GtkBin{GtkContainer{GtkWidget{
		C.to_GtkWidget(unsafe.Pointer(C._gtk_viewport_get_bin_window(C.to_GtkViewport(v.Widget))))}}}}
}
func (v *GtkViewport) GetViewWindow() *GtkWindow {
	panic_if_version_older_auto(2, 22, 0)
	return &GtkWindow{GtkBin{GtkContainer{GtkWidget{
		C.to_GtkWidget(unsafe.Pointer(C._gtk_viewport_get_view_window(C.to_GtkViewport(v.Widget))))}}}}
}

//-----------------------------------------------------------------------
// GtkAccessible
//-----------------------------------------------------------------------
type GtkAccessible struct {
	glib.GObject
}

func (v *GtkAccessible) ConnectWidgetDestroyed() {
	C.gtk_accessible_connect_widget_destroyed(
		C.to_GtkAccessible(unsafe.Pointer(v.Object)))
}
func (v *GtkAccessible) SetWidget(w WidgetLike) {
	panic_if_version_older_auto(2, 22, 0)
	C._gtk_accessible_set_widget(
		C.to_GtkAccessible(unsafe.Pointer(v.Object)), w.ToNative())
}
func (v *GtkAccessible) GetWidget() *GtkWidget {
	panic_if_version_older_auto(2, 22, 0)
	return &GtkWidget{C._gtk_accessible_get_widget(
		C.to_GtkAccessible(unsafe.Pointer(v.Object)))}
}

//-----------------------------------------------------------------------
// GtkBin
//-----------------------------------------------------------------------
type GtkBin struct {
	GtkContainer
}

func (v *GtkBin) GetChild() *GtkWidget {
	return &GtkWidget{C.gtk_bin_get_child(C.to_GtkBin(v.Widget))}
}

//-----------------------------------------------------------------------
// GtkBox
//-----------------------------------------------------------------------
type GtkPackType int

const (
	GTK_PACK_START GtkPackType = 0
	GTK_PACK_END   GtkPackType = 1
)

type BoxLike interface {
	ContainerLike
	PackStart(child WidgetLike, expand bool, fill bool, padding uint)
	PackEnd(child WidgetLike, expand bool, fill bool, padding uint)
}
type GtkBox struct {
	GtkContainer
}

func (v *GtkBox) PackStart(child WidgetLike, expand bool, fill bool, padding uint) {
	C.gtk_box_pack_start(C.to_GtkBox(v.Widget), child.ToNative(), bool2gboolean(expand), bool2gboolean(fill), C.guint(padding))
}
func (v *GtkBox) PackEnd(child WidgetLike, expand bool, fill bool, padding uint) {
	C.gtk_box_pack_end(C.to_GtkBox(v.Widget), child.ToNative(), bool2gboolean(expand), bool2gboolean(fill), C.guint(padding))
}

//Deprecated since 2.14. Use PackStart() instead.
func (v *GtkBox) PackStartDefaults(child WidgetLike) {
	deprecated_since(2, 14, 0, "gtk_box_pack_start_defaults()")
	C.gtk_box_pack_start_defaults(C.to_GtkBox(v.Widget), child.ToNative())
}

//Deprecated since 2.14. Use PackEnd() instead.
func (v *GtkBox) PackEndDefaults(child WidgetLike) {
	deprecated_since(2, 14, 0, "gtk_box_pack_end_defaults()")
	C.gtk_box_pack_end_defaults(C.to_GtkBox(v.Widget), child.ToNative())
}
func (v *GtkBox) GetHomogeneous() bool {
	return gboolean2bool(C.gtk_box_get_homogeneous(C.to_GtkBox(v.Widget)))
}
func (v *GtkBox) SetHomogeneous(homogeneous bool) {
	C.gtk_box_set_homogeneous(C.to_GtkBox(v.Widget), bool2gboolean(homogeneous))
}
func (v *GtkBox) GetSpacing() int {
	return int(C.gtk_box_get_spacing(C.to_GtkBox(v.Widget)))
}
func (v *GtkBox) SetSpacing(spacing int) {
	C.gtk_box_set_spacing(C.to_GtkBox(v.Widget), C.gint(spacing))
}
func (v *GtkBox) ReorderChild(child WidgetLike, position GtkPackType) {
	C.gtk_box_reorder_child(C.to_GtkBox(v.Widget), child.ToNative(), C.gint(position))
}
func (v *GtkBox) QueryChildPacking(child WidgetLike, expand *bool, fill *bool, padding *uint, pack_type *GtkPackType) {
	var cexpand, cfill C.gboolean
	var cpadding C.guint
	var cpack_type C.GtkPackType
	C.gtk_box_query_child_packing(C.to_GtkBox(v.Widget), child.ToNative(), &cexpand, &cfill, &cpadding, &cpack_type)
	*expand = gboolean2bool(cexpand)
	*fill = gboolean2bool(cfill)
	*padding = uint(cpadding)
	*pack_type = GtkPackType(cpack_type)
}
func (v *GtkBox) SetChildPacking(child WidgetLike, expand bool, fill bool, padding uint, pack_type GtkPackType) {
	C.gtk_box_set_child_packing(C.to_GtkBox(v.Widget), child.ToNative(), bool2gboolean(expand), bool2gboolean(fill), C.guint(padding), C.GtkPackType(pack_type))
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
type GtkContainer struct {
	GtkWidget
}

func (v *GtkContainer) Add(w WidgetLike) {
	C.gtk_container_add(C.to_GtkContainer(v.Widget), w.ToNative())
}
func (v *GtkContainer) Remove(w WidgetLike) {
	C.gtk_container_remove(C.to_GtkContainer(v.Widget), w.ToNative())
}

// gtk_container_add_with_properties
// gtk_container_get_resize_mode
// gtk_container_set_resize_mode

func (v *GtkContainer) CheckResize() {
	C.gtk_container_check_resize(C.to_GtkContainer(v.Widget))
}

// gtk_container_foreach

func (v *GtkContainer) GetChildren() *glib.List {
	return glib.ListFromNative(unsafe.Pointer(C.gtk_container_get_children(C.to_GtkContainer(v.Widget))))
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

func (v *GtkContainer) GetBorderWidth() uint {
	return uint(C.gtk_container_get_border_width(C.to_GtkContainer(v.Widget)))
}
func (v *GtkContainer) SetBorderWidth(border_width uint) {
	C.gtk_container_set_border_width(C.to_GtkContainer(v.Widget), C.guint(border_width))
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
type GtkItem struct {
	GtkBin
}

//Deprecated since 2.22. Use GtkMenuItem.Select() instead.
func (v *GtkItem) Select() {
	deprecated_since(2, 22, 0, "gtk_item_select()")
	C.gtk_item_select(C.to_GtkItem(v.Widget))
}

//Deprecated since 2.22. Use GtkMenuItem.Deselect() instead.
func (v *GtkItem) Deselect() {
	deprecated_since(2, 22, 0, "gtk_item_deselect()")
	C.gtk_item_deselect(C.to_GtkItem(v.Widget))
}

//Deprecated since 2.22.
func (v *GtkItem) Toggle() {
	deprecated_since(2, 22, 0, "gtk_item_select()")
	C.gtk_item_toggle(C.to_GtkItem(v.Widget))
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
type GtkObject struct {
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
type GtkPaned struct {
	GtkContainer
}

func (v *GtkPaned) Add1(child WidgetLike) {
	C.gtk_paned_add1(C.to_GtkPaned(v.Widget), child.ToNative())
}
func (v *GtkPaned) Add2(child WidgetLike) {
	C.gtk_paned_add2(C.to_GtkPaned(v.Widget), child.ToNative())
}
func (v *GtkPaned) Pack1(child WidgetLike, resize bool, shrink bool) {
	C.gtk_paned_pack1(C.to_GtkPaned(v.Widget), child.ToNative(), bool2gboolean(resize), bool2gboolean(shrink))
}
func (v *GtkPaned) Pack2(child WidgetLike, resize bool, shrink bool) {
	C.gtk_paned_pack2(C.to_GtkPaned(v.Widget), child.ToNative(), bool2gboolean(resize), bool2gboolean(shrink))
}
func (v *GtkPaned) GetChild1() *GtkWidget {
	return &GtkWidget{
		C.gtk_paned_get_child1(C.to_GtkPaned(v.Widget))}
}
func (v *GtkPaned) GetChild2() *GtkWidget {
	return &GtkWidget{
		C.gtk_paned_get_child2(C.to_GtkPaned(v.Widget))}
}
func (v *GtkPaned) SetPosition(position int) {
	C.gtk_paned_set_position(C.to_GtkPaned(v.Widget), C.gint(position))
}
func (v *GtkPaned) GetPosition() int {
	return int(C.gtk_paned_get_position(C.to_GtkPaned(v.Widget)))
}

// gtk_paned_get_handle_window //since 2.20

//-----------------------------------------------------------------------
// GtkRange
//-----------------------------------------------------------------------
type GtkRange struct {
	GtkWidget
}

func (v *GtkRange) GetFillLevel() float64 {
	r := C.gtk_range_get_fill_level(C.to_GtkRange(v.Widget))
	return float64(r)
}
func (v *GtkRange) GetRestrictToFillLevel() bool {
	return gboolean2bool(C.gtk_range_get_restrict_to_fill_level(C.to_GtkRange(v.Widget)))
}
func (v *GtkRange) GetShowFillLevel() bool {
	return gboolean2bool(C.gtk_range_get_show_fill_level(C.to_GtkRange(v.Widget)))
}
func (v *GtkRange) SetFillLevel(value float64) {
	C.gtk_range_set_fill_level(C.to_GtkRange(v.Widget), C.gdouble(value))
}
func (v *GtkRange) SetRestrictToFillLevel(b bool) {
	C.gtk_range_set_restrict_to_fill_level(C.to_GtkRange(v.Widget), bool2gboolean(b))
}
func (v *GtkRange) SetShowFillLevel(b bool) {
	C.gtk_range_set_show_fill_level(C.to_GtkRange(v.Widget), bool2gboolean(b))
}
func (v *GtkRange) GetAdjustment() *GtkAdjustment {
	return &GtkAdjustment{
		C.gtk_range_get_adjustment(C.to_GtkRange(v.Widget))}
}

// void gtk_range_set_update_policy (GtkRange *range, GtkUpdateType policy); //deprecated in 2.24

func (v *GtkRange) SetAdjustment(adjustment *GtkAdjustment) {
	C.gtk_range_set_adjustment(C.to_GtkRange(v.Widget), adjustment.Adjustment)
}
func (v *GtkRange) GetInverted() bool {
	return gboolean2bool(C.gtk_range_get_inverted(C.to_GtkRange(v.Widget)))
}
func (v *GtkRange) SetInverted(b bool) {
	C.gtk_range_set_inverted(C.to_GtkRange(v.Widget), bool2gboolean(b))
}

// GtkUpdateType gtk_range_get_update_policy (GtkRange *range); //deprecated since 2.24

func (v *GtkRange) SetIncrements(step, page float64) {
	C.gtk_range_set_increments(C.to_GtkRange(v.Widget), C.gdouble(step), C.gdouble(page))
}
func (v *GtkRange) SetRange(min, max float64) {
	C.gtk_range_set_range(C.to_GtkRange(v.Widget), C.gdouble(min), C.gdouble(max))
}
func (v *GtkRange) GetValue() float64 {
	return float64(C.gtk_range_get_value(C.to_GtkRange(v.Widget))) //TODO test
	//var r C.gdouble
	//C._gtk_range_get_value(C.to_GtkRange(v.Widget), &r)
	//return float64(r)
}
func (v *GtkRange) SetValue(value float64) {
	C.gtk_range_set_value(C.to_GtkRange(v.Widget), C.gdouble(value))
}

// gtk_range_get_round_digits //since 2.24
// gtk_range_set_round_digits //since 2.24
// void gtk_range_set_lower_stepper_sensitivity (GtkRange *range, GtkSensitivityType sensitivity);
// GtkSensitivityType gtk_range_get_lower_stepper_sensitivity (GtkRange *range);
// void gtk_range_set_upper_stepper_sensitivity (GtkRange *range, GtkSensitivityType sensitivity);
// GtkSensitivityType gtk_range_get_upper_stepper_sensitivity (GtkRange *range);

func (v *GtkRange) GetFlippable() bool {
	panic_if_version_older(2, 18, 0, "gtk_range_get_flippable()")
	return gboolean2bool(C._gtk_range_get_flippable(C.to_GtkRange(v.Widget)))
}
func (v *GtkRange) SetFlippable(b bool) {
	panic_if_version_older(2, 18, 0, "gtk_range_set_flippable()")
	C._gtk_range_set_flippable(C.to_GtkRange(v.Widget), bool2gboolean(b))
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
type GtkPositionType int

const (
	GTK_POS_LEFT   GtkPositionType = 0
	GTK_POS_RIGHT  GtkPositionType = 1
	GTK_POS_TOP    GtkPositionType = 2
	GTK_POS_BOTTOM GtkPositionType = 3
)

type GtkScale struct {
	GtkRange
}

func (v *GtkScale) SetDigits(digits int) {
	C.gtk_scale_set_digits(C.to_GtkScale(v.Widget), C.gint(digits))
}
func (v *GtkScale) SetDrawValue(draw_value bool) {
	C.gtk_scale_set_draw_value(C.to_GtkScale(v.Widget), bool2gboolean(draw_value))
}
func (v *GtkScale) SetValuePos(pos GtkPositionType) {
	C.gtk_scale_set_value_pos(C.to_GtkScale(v.Widget), C.GtkPositionType(pos))
}
func (v *GtkScale) GetDigits() int {
	return int(C.gtk_scale_get_digits(C.to_GtkScale(v.Widget)))
}
func (v *GtkScale) GetDrawValue() bool {
	return gboolean2bool(C.gtk_scale_get_draw_value(C.to_GtkScale(v.Widget)))
}
func (v *GtkScale) GetValuePos() GtkPositionType {
	return GtkPositionType(C.gtk_scale_get_value_pos(C.to_GtkScale(v.Widget)))
}

// PangoLayout * gtk_scale_get_layout (GtkScale *scale);

func (v *GtkScale) GetLayoutOffsets(x *int, y *int) {
	var xx, yy C.gint
	C.gtk_scale_get_layout_offsets(C.to_GtkScale(v.Widget), &xx, &yy)
	*x = int(xx)
	*y = int(yy)
}
func (v *GtkScale) AddMark(value float64, position GtkPositionType, markup string) {
	ptr := C.CString(markup)
	defer C.free_string(ptr)
	C.gtk_scale_add_mark(C.to_GtkScale(v.Widget), C.gdouble(value), C.GtkPositionType(position), C.to_gcharptr(ptr))
}
func (v *GtkScale) ClearMarks() {
	C.gtk_scale_clear_marks(C.to_GtkScale(v.Widget))
}

//-----------------------------------------------------------------------
// GtkSeparator
//-----------------------------------------------------------------------
type GtkSeparator struct {
	GtkWidget
}

//-----------------------------------------------------------------------
// GtkWidget
//-----------------------------------------------------------------------
type GtkAllocation gdk.Rectangle

type GtkAccelFlags int

const (
	GTK_ACCEL_VISIBLE GtkAccelFlags = 1 << 0
	GTK_ACCEL_LOCKED  GtkAccelFlags = 1 << 1
	GTK_ACCEL_MASK    GtkAccelFlags = 0x07
)

type GtkStateType int

const (
	GTK_STATE_NORMAL      GtkStateType = 0
	GTK_STATE_ACTIVE      GtkStateType = 1
	GTK_STATE_PRELIGHT    GtkStateType = 2
	GTK_STATE_SELECTED    GtkStateType = 3
	GTK_STATE_INSENSITIVE GtkStateType = 4
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
	GetTopLevel() *GtkWidget
	GetTopLevelAsWindow() *GtkWindow
	HideOnDelete()
	QueueResize()
}
type GtkWidget struct {
	Widget *C.GtkWidget
}

func WidgetFromNative(p unsafe.Pointer) *GtkWidget {
	return &GtkWidget{C.to_GtkWidget(p)}
}

//TODO GtkWidget will have GObject as anonymous field.
func WidgetFromObject(object *glib.GObject) *GtkWidget {
	return &GtkWidget{
		C.to_GtkWidget(unsafe.Pointer(object.Object))}
}
func (v *GtkWidget) ToNative() *C.GtkWidget {
	if v == nil {
		return nil
	}
	return v.Widget
}
func (v *GtkWidget) Connect(s string, f interface{}, datas ...interface{}) int {
	return glib.ObjectFromNative(unsafe.Pointer(v.Widget)).Connect(s, f, datas...)
}
func (v *GtkWidget) StopEmission(s string) {
	glib.ObjectFromNative(unsafe.Pointer(v.Widget)).StopEmission(s)
}
func (v *GtkWidget) Emit(s string) {
	glib.ObjectFromNative(unsafe.Pointer(v.Widget)).Emit(s)
}

func (v *GtkWidget) HandlerBlock(id int) {
	glib.ObjectFromNative(unsafe.Pointer(v.Widget)).HandlerBlock(id)
}

func (v *GtkWidget) HandlerUnblock(id int) {
	glib.ObjectFromNative(unsafe.Pointer(v.Widget)).HandlerUnblock(id)
}

func (v *GtkWidget) HandlerDisconnect(id int) {
	glib.ObjectFromNative(unsafe.Pointer(v.Widget)).HandlerDisconnect(id)
}

// gtk_widget_new

//Deprecated since 2.12. Use g_object_ref() instead. //TODO gobject
func (v *GtkWidget) Ref() {
	deprecated_since(2, 12, 0, "gtk_widget_ref()")
	C.gtk_widget_ref(v.Widget)
}

//Deprecated since 2.12. Use g_object_ref() instead. //TODO gobject
func (v *GtkWidget) Unref() {
	deprecated_since(2, 12, 0, "gtk_widget_unref()")
	C.gtk_widget_unref(v.Widget)
}
func (v *GtkWidget) Destroy() {
	C.gtk_widget_destroy(v.Widget)
}

// gtk_widget_destroyed

func (v *GtkWidget) Unparent() {
	C.gtk_widget_unparent(v.Widget)
}
func (v *GtkWidget) Show() {
	C.gtk_widget_show(v.Widget)
}
func (v *GtkWidget) ShowNow() {
	C.gtk_widget_show_now(v.Widget)
}
func (v *GtkWidget) Hide() {
	C.gtk_widget_hide(v.Widget)
}
func (v *GtkWidget) ShowAll() {
	C.gtk_widget_show_all(v.Widget)
}

//Deprecated since 2.24. Use Hide() instead.
func (v *GtkWidget) HideAll() {
	deprecated_since(2, 24, 0, "gtk_widget_hide_all()")
	C.gtk_widget_hide_all(v.Widget)
}
func (v *GtkWidget) Map() {
	C.gtk_widget_map(v.Widget)
}
func (v *GtkWidget) Unmap() {
	C.gtk_widget_unmap(v.Widget)
}
func (v *GtkWidget) Realize() {
	C.gtk_widget_realize(v.Widget)
}
func (v *GtkWidget) Unrealize() {
	C.gtk_widget_unrealize(v.Widget)
}
func (v *GtkWidget) QueueDraw() {
	C.gtk_widget_queue_draw(v.Widget)
}
func (v *GtkWidget) QueueResize() {
	C.gtk_widget_queue_resize(v.Widget)
}
func (v *GtkWidget) QueueResizeNoRedraw() {
	C.gtk_widget_queue_resize_no_redraw(v.Widget)
}

// gtk_widget_size_request
// gtk_widget_get_child_requisition
// gtk_widget_size_allocate

func (v *GtkWidget) AddAccelerator(signal string, group *GtkAccelGroup, key uint, mods gdk.GdkModifierType, flags GtkAccelFlags) {
	csignal := C.CString(signal)
	defer C.free_string(csignal)
	C.gtk_widget_add_accelerator(v.Widget, C.to_gcharptr(csignal), group.AccelGroup, C.guint(key), C.GdkModifierType(mods), C.GtkAccelFlags(flags))
}

// gtk_widget_remove_accelerator
// gtk_widget_set_accel_path
// gtk_widget_list_accel_closures

func (v *GtkWidget) CanActivateAccel(signal_id uint) bool {
	return gboolean2bool(C.gtk_widget_can_activate_accel(v.Widget, C.guint(signal_id)))
}

// gtk_widget_event

func (v *GtkWidget) Activate() {
	C.gtk_widget_activate(v.Widget)
}
func (v *GtkWidget) Reparent(parent WidgetLike) {
	C.gtk_widget_reparent(v.Widget, parent.ToNative())
}

// gtk_widget_intersect

func (v *GtkWidget) IsFocus() bool {
	return gboolean2bool(C.gtk_widget_is_focus(v.Widget))
}
func (v *GtkWidget) GrabFocus() {
	C.gtk_widget_grab_focus(v.Widget)
}
func (v *GtkWidget) GrabDefault() {
	C.gtk_widget_grab_default(v.Widget)
}
func (v *GtkWindow) SetName(name string) {
	ptr := C.CString(name)
	defer C.free_string(ptr)
	C.gtk_widget_set_name(v.Widget, C.to_gcharptr(ptr))
}
func (v *GtkWindow) GetName() string {
	return C.GoString(C.to_charptr(C.gtk_widget_get_name(v.Widget)))
}
func (v *GtkWidget) SetState(state GtkStateType) {
	C.gtk_widget_set_state(v.Widget, C.GtkStateType(state))
}
func (v *GtkWidget) SetSensitive(setting bool) {
	C.gtk_widget_set_sensitive(v.Widget, bool2gboolean(setting))
}
func (v *GtkWidget) SetParent(parent WidgetLike) {
	C.gtk_widget_set_parent(v.Widget, parent.ToNative())
}
func (v *GtkWidget) SetParentWindow(parent *gdk.GdkWindow) {
	C.gtk_widget_set_parent_window(v.Widget, C.to_GdkWindow(unsafe.Pointer(parent.Window)))
}
func (v *GtkWidget) GetParentWindow() *gdk.GdkWindow {
	return gdk.WindowFromUnsafe(unsafe.Pointer(C.gtk_widget_get_parent_window(v.Widget)))
}

//Deprecated since 2.2, use SetSizeRequest() instead
func (v *GtkWidget) SetUSize(width int, height int) {
	deprecated_since(2, 2, 0, "gtk_widget_set_usize()")
	C.gtk_widget_set_usize(v.Widget, C.gint(width), C.gint(height))
}
func (v *GtkWidget) SetEvents(events int) {
	C.gtk_widget_set_events(v.Widget, C.gint(events))
}
func (v *GtkWidget) AddEvents(events int) {
	C.gtk_widget_add_events(v.Widget, C.gint(events))
}

// gtk_widget_set_extension_events
// gtk_widget_get_extension_events

func (v *GtkWidget) GetTopLevel() *GtkWidget {
	return &GtkWidget{
		C.gtk_widget_get_toplevel(v.Widget)}
}

// gtk_widget_get_ancestor
// gtk_widget_get_colormap
// gtk_widget_set_colormap
// gtk_widget_get_visual
// gtk_widget_get_events
// gtk_widget_get_pointer
// gtk_widget_is_ancestor
// gtk_widget_translate_coordinates

func (v *GtkWidget) HideOnDelete() {
	C._gtk_widget_hide_on_delete(v.Widget)
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

func (v *GtkWidget) RenderIcon(stock_id string, size GtkIconSize, detail string) *gdkpixbuf.GdkPixbuf {
	pstock_id := C.CString(stock_id)
	defer C.free_string(pstock_id)
	pdetail := C.CString(detail)
	defer C.free_string(pdetail)
	return &gdkpixbuf.GdkPixbuf{
		C.gtk_widget_render_icon(v.Widget, C.to_gcharptr(pstock_id), C.GtkIconSize(size), C.to_gcharptr(pdetail))}
}

// gtk_widget_pop_composite_child
// gtk_widget_push_composite_child

//Deprecated since 2.2. Use QueueDraw() instead.
func (v *GtkWidget) QueueClear() {
	deprecated_since(2, 2, 0, "gtk_widget_queue_clear()")
	C.gtk_widget_queue_clear(v.Widget)
}

// gtk_widget_queue_draw_area

func (v *GtkWidget) SetAppPrintable(setting bool) {
	C.gtk_widget_set_app_paintable(v.Widget, bool2gboolean(setting))
}
func (v *GtkWidget) SetDoubleBuffered(setting bool) {
	C.gtk_widget_set_double_buffered(v.Widget, bool2gboolean(setting))
}
func (v *GtkWidget) SetRedrawOnAllocate(setting bool) {
	C.gtk_widget_set_redraw_on_allocate(v.Widget, bool2gboolean(setting))
}

// gtk_widget_set_composite_name
// gtk_widget_set_scroll_adjustments

func (v *GtkWidget) MnemonicActivate(group_cycling bool) bool {
	return gboolean2bool(C.gtk_widget_mnemonic_activate(v.Widget, bool2gboolean(group_cycling)))
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

func (v *GtkWidget) GetChildVisible() bool {
	return gboolean2bool(C.gtk_widget_get_child_visible(v.Widget))
}
func (v *GtkWidget) GetParent() *GtkWidget {
	return &GtkWidget{C.gtk_widget_get_parent(v.Widget)}
}
func (v *GtkWidget) GetSettings() *GtkSettings {
	return &GtkSettings{C.gtk_widget_get_settings(v.Widget)}
}

// gtk_widget_get_clipboard
// gtk_widget_get_display
// gtk_widget_get_root_window
// gtk_widget_get_screen
// gtk_widget_has_screen

//TODO go can have multiple return, adapt the function!
func (v *GtkWidget) GetSizeRequest(width *int, height *int) {
	var w, h C.gint
	C.gtk_widget_get_size_request(v.Widget, &w, &h)
	*width = int(w)
	*height = int(h)
}
func (v *GtkWidget) SetChildVisible(setting bool) {
	C.gtk_widget_set_child_visible(v.Widget, bool2gboolean(setting))
}
func (v *GtkWidget) SetSizeRequest(width int, height int) {
	C.gtk_widget_set_size_request(v.Widget, C.gint(width), C.gint(height))
}
func (v *GtkWidget) SetNoShowAll(setting bool) {
	C.gtk_widget_set_no_show_all(v.Widget, bool2gboolean(setting))
}
func (v *GtkWidget) GetNoShowAll() bool {
	return gboolean2bool(C.gtk_widget_get_no_show_all(v.Widget))
}

// gtk_widget_list_mnemonic_labels
// gtk_widget_add_mnemonic_label
// gtk_widget_remove_mnemonic_label

func (v *GtkWidget) IsComposited() bool {
	return gboolean2bool(C.gtk_widget_is_composited(v.Widget))
}

// gtk_widget_error_bell
// gtk_widget_keynav_failed

func (v *GtkWidget) GetTooltipMarkup() string {
	return C.GoString(C.to_charptr(C.gtk_widget_get_tooltip_markup(v.Widget)))
}
func (v *GtkWidget) SetTooltipMarkup(markup string) {
	ptr := C.CString(markup)
	defer C.free_string(ptr)
	C.gtk_widget_set_tooltip_markup(v.Widget, C.to_gcharptr(ptr))
}
func (v *GtkWidget) GetTooltipText() string {
	return C.GoString(C.to_charptr(C.gtk_widget_get_tooltip_text(v.Widget)))
}
func (v *GtkWidget) SetTooltipText(text string) {
	ptr := C.CString(text)
	defer C.free_string(ptr)
	C.gtk_widget_set_tooltip_text(v.Widget, C.to_gcharptr(ptr))
}
func (v *GtkWidget) GetTooltipWindow() *GtkWindow {
	return &GtkWindow{GtkBin{GtkContainer{GtkWidget{
		C.to_GtkWidget(unsafe.Pointer(C.gtk_widget_get_tooltip_window(v.Widget)))}}}}
}
func (v *GtkWidget) SetTooltipWindow(w *GtkWindow) {
	C.gtk_widget_set_tooltip_window(v.Widget, C.to_GtkWindow(w.ToNative()))
}
func (v *GtkWidget) GetHasTooltip() bool {
	return gboolean2bool(C.gtk_widget_get_has_tooltip(v.Widget))
}
func (v *GtkWidget) SetHasTooltip(setting bool) {
	C.gtk_widget_set_has_tooltip(v.Widget, bool2gboolean(setting))
}

// gtk_widget_trigger_tooltip_query
// gtk_widget_get_snapshot

func (v *GtkWidget) GetWindow() *gdk.GdkWindow {
	return gdk.WindowFromUnsafe(unsafe.Pointer(C.gtk_widget_get_window(v.Widget)))
}

//TODO get should return something (this function mechanism is not intuitive)
func (v *GtkWidget) GetAllocation(allocation *GtkAllocation) {
	panic_if_version_older(2, 18, 0, "gtk_widget_get_allocation()")
	var _allocation C.GtkAllocation
	C._gtk_widget_get_allocation(v.Widget, &_allocation)
	allocation.X = int(_allocation.x)
	allocation.Y = int(_allocation.y)
	allocation.Width = int(_allocation.width)
	allocation.Height = int(_allocation.height)
}
func (v *GtkWidget) SetAllocation(allocation *GtkAllocation) {
	panic_if_version_older(2, 18, 0, "gtk_widget_set_allocation()")
	var _allocation C.GtkAllocation
	_allocation.x = C.gint(allocation.X)
	_allocation.y = C.gint(allocation.Y)
	_allocation.width = C.gint(allocation.Width)
	_allocation.height = C.gint(allocation.Height)
	C._gtk_widget_set_allocation(v.Widget, &_allocation)
}
func (v *GtkWidget) GetAppPaintable() bool {
	panic_if_version_older(2, 18, 0, "gtk_widget_get_app_paintable()")
	return gboolean2bool(C._gtk_widget_get_app_paintable(v.Widget))
}
func (v *GtkWidget) GetCanDefault() bool {
	panic_if_version_older(2, 18, 0, "gtk_widget_get_can_default()")
	return gboolean2bool(C._gtk_widget_get_can_default(v.Widget))
}
func (v *GtkWidget) SetCanDefault(setting bool) {
	panic_if_version_older(2, 18, 0, "gtk_widget_set_can_default()")
	C._gtk_widget_set_can_default(v.Widget, bool2gboolean(setting))
}
func (v *GtkWidget) GetCanFocus() bool {
	panic_if_version_older(2, 18, 0, "gtk_widget_get_can_focus()")
	return gboolean2bool(C._gtk_widget_get_can_focus(v.Widget))
}
func (v *GtkWidget) SetCanFocus(setting bool) {
	panic_if_version_older(2, 18, 0, "gtk_widget_set_can_focus()")
	C._gtk_widget_set_can_focus(v.Widget, bool2gboolean(setting))
}
func (v *GtkWidget) GetDoubleBuffered() bool {
	panic_if_version_older(2, 18, 0, "gtk_widget_get_double_buffered()")
	return gboolean2bool(C._gtk_widget_get_double_buffered(v.Widget))
}
func (v *GtkWidget) GetHasWindow() bool {
	panic_if_version_older(2, 18, 0, "gtk_widget_get_has_window()")
	return gboolean2bool(C._gtk_widget_get_has_window(v.Widget))
}
func (v *GtkWidget) SetHasWindow(setting bool) {
	panic_if_version_older(2, 18, 0, "gtk_widget_set_has_window()")
	C._gtk_widget_set_has_window(v.Widget, bool2gboolean(setting))
}
func (v *GtkWidget) GetSensitive() bool {
	panic_if_version_older(2, 18, 0, "gtk_widget_get_sensitive()")
	return gboolean2bool(C._gtk_widget_get_sensitive(v.Widget))
}
func (v *GtkWidget) IsSensitive() bool {
	panic_if_version_older(2, 18, 0, "gtk_widget_is_sensitive()")
	return gboolean2bool(C._gtk_widget_is_sensitive(v.Widget))
}
func (v *GtkWidget) GetState() GtkStateType {
	panic_if_version_older(2, 18, 0, "gtk_widget_get_state()")
	return GtkStateType(C._gtk_widget_get_state(v.Widget))
}
func (v *GtkWidget) GetVisible() bool {
	panic_if_version_older(2, 18, 0, "gtk_widget_get_visible()")
	return gboolean2bool(C._gtk_widget_get_visible(v.Widget))
}
func (v *GtkWidget) SetVisible(setting bool) {
	panic_if_version_older(2, 18, 0, "gtk_widget_set_visible()")
	C._gtk_widget_set_visible(v.Widget, bool2gboolean(setting))
}
func (v *GtkWidget) HasDefault() bool {
	panic_if_version_older(2, 18, 0, "gtk_widget_has_default()")
	return gboolean2bool(C._gtk_widget_has_default(v.Widget))
}
func (v *GtkWidget) HasFocus() bool {
	panic_if_version_older(2, 18, 0, "gtk_widget_has_focus()")
	return gboolean2bool(C._gtk_widget_has_focus(v.Widget))
}
func (v *GtkWidget) HasGrab() bool {
	panic_if_version_older(2, 18, 0, "gtk_widget_has_grab()")
	return gboolean2bool(C._gtk_widget_has_grab(v.Widget))
}

// gtk_widget_has_rc_style //since 2.20

func (v *GtkWidget) IsDrawable() bool {
	panic_if_version_older(2, 18, 0, "gtk_widget_is_drawable()")
	return gboolean2bool(C._gtk_widget_is_drawable(v.Widget))
}
func (v *GtkWidget) IsToplevel() bool {
	panic_if_version_older(2, 18, 0, "gtk_widget_is_toplevel()")
	return gboolean2bool(C._gtk_widget_is_toplevel(v.Widget))
}
func (v *GtkWidget) SetWindow(window *gdk.GdkWindow) {
	panic_if_version_older(2, 18, 0, "gtk_widget_set_window()")
	C._gtk_widget_set_window(v.Widget, C.to_GdkWindow(unsafe.Pointer(window.Window)))
}
func (v *GtkWidget) SetReceivesDefault(setting bool) {
	panic_if_version_older(2, 18, 0, "gtk_widget_set_receives_default()")
	C._gtk_widget_set_receives_default(v.Widget, bool2gboolean(setting))
}
func (v *GtkWidget) GetReceivesDefault() bool {
	panic_if_version_older(2, 18, 0, "gtk_widget_get_receives_default()")
	return gboolean2bool(C._gtk_widget_get_receives_default(v.Widget))
}

// gtk_widget_set_realized //since 2.20
// gtk_widget_get_realized //since 2.20
// gtk_widget_set_mapped //since 2.20
// gtk_widget_get_mapped //since 2.20
// gtk_widget_get_requisition //since 2.20
// gtk_requisition_copy
// gtk_requisition_free

//Convenience functions

func (v *GtkWidget) GetTopLevelAsWindow() *GtkWindow {
	return &GtkWindow{GtkBin{GtkContainer{GtkWidget{
		C.gtk_widget_get_toplevel(v.Widget)}}}}
}
func (v *GtkWidget) ModifyFontEasy(desc string) {
	pdesc := C.CString(desc)
	defer C.free_string(pdesc)
	C.gtk_widget_modify_font(v.Widget, C.pango_font_description_from_string(pdesc))
}

func (v *GtkWidget) ModifyBG(state GtkStateType, color *gdk.GdkColor) {
	C.gtk_widget_modify_bg(v.Widget, C.GtkStateType(state), (*C.GdkColor)(unsafe.Pointer(&color.Color)))
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
type GtkBuilder struct {
	Builder *C.GtkBuilder
}

func Builder() *GtkBuilder {
	return &GtkBuilder{
		C.gtk_builder_new()}
}
func (v *GtkBuilder) AddFromFile(filename string) (ret uint, error *glib.Error) {
	var gerror *C.GError
	ptr := C.CString(filename)
	defer C.free_string(ptr)
	ret = uint(C.gtk_builder_add_from_file(v.Builder, C.to_gcharptr(ptr), &gerror))
	if gerror != nil {
		error = glib.ErrorFromNative(unsafe.Pointer(gerror))
	} else {
		error = nil
	}
	return
}
func (v *GtkBuilder) AddFromString(buffer string) (ret uint, error *glib.Error) {
	var gerror *C.GError
	ptr := C.CString(buffer)
	defer C.free_string(ptr)
	ret = uint(C.gtk_builder_add_from_string(v.Builder, C.to_gcharptr(ptr), C.gsize(C.strlen(ptr)), &gerror))
	if gerror != nil {
		error = glib.ErrorFromNative(unsafe.Pointer(gerror))
	} else {
		error = nil
	}
	return
}

// guint gtk_builder_add_objects_from_file (GtkBuilder *builder, const gchar *filename, gchar **object_ids, GError **error);
// guint gtk_builder_add_objects_from_string (GtkBuilder *builder, const gchar *buffer, gsize length, gchar **object_ids, GError **error);

func (v *GtkBuilder) GetObject(name string) *glib.GObject {
	ptr := C.CString(name)
	defer C.free_string(ptr)
	return &glib.GObject{
		unsafe.Pointer(C.gtk_builder_get_object(v.Builder, C.to_gcharptr(ptr)))}
}
func (v *GtkBuilder) GetObjects() *glib.SList {
	return glib.SListFromNative(unsafe.Pointer(C.gtk_builder_get_objects(v.Builder)))
}
func (v *GtkBuilder) ConnectSignals(user_data interface{}) {
	C.gtk_builder_connect_signals(v.Builder, nil)
}

// void gtk_builder_connect_signals_full (GtkBuilder *builder, GtkBuilderConnectFunc func, gpointer user_data);

func (v *GtkBuilder) SetTranslationDomain(domain string) {
	ptr := C.CString(domain)
	defer C.free_string(ptr)
	C.gtk_builder_set_translation_domain(v.Builder, C.to_gcharptr(ptr))
}
func (v *GtkBuilder) GetTranslationDomain() string {
	return C.GoString(C.to_charptr(C.gtk_builder_get_translation_domain(v.Builder)))
}
func (v *GtkBuilder) GetTypeFromName(type_name string) int {
	ptr := C.CString(type_name)
	defer C.free_string(ptr)
	return int(C.gtk_builder_get_type_from_name(v.Builder, ptr))
}

// gboolean gtk_builder_value_from_string (GtkBuilder *builder, GParamSpec *pspec, const gchar *string, GValue *value, GError **error);
// gboolean gtk_builder_value_from_string_type (GtkBuilder *builder, GType type, const gchar *string, GValue *value, GError **error);
