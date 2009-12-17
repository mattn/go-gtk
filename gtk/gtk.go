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

static long _gtk_signal_connect(GtkWidget* widget, gchar* name, int func_no, void* data) {
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

static const gchar* _gtk_window_get_title(GtkWidget* widget) {
	return gtk_window_get_title(GTK_WINDOW(widget));
}

static void _gtk_window_set_title(GtkWidget* widget, gchar* title) {
	gtk_window_set_title(GTK_WINDOW(widget), title);
}

static void _gtk_window_set_transient_for(GtkWidget* widget, GtkWidget *parent) {
	gtk_window_set_transient_for(GTK_WINDOW(widget), GTK_WINDOW(parent));
}

static int _gtk_dialog_run(GtkWidget* dialog) {
	return gtk_dialog_run(GTK_DIALOG(dialog));
}

static GtkWidget* _gtk_dialog_add_button(GtkWidget* dialog, const gchar* button_text, gint response_id) {
	return gtk_dialog_add_button(GTK_DIALOG(dialog), button_text, response_id);
}

static GtkWidget* _gtk_message_dialog_new(GtkWidget* parent, GtkDialogFlags flags, GtkMessageType type, GtkButtonsType buttons, gchar *message) {
	return gtk_message_dialog_new(
			GTK_WINDOW(parent),
			flags,
			type,
			buttons,
			message, NULL);
}

static GtkWidget* _gtk_file_chooser_dialog_new(const gchar* title, GtkWidget* parent, int action, const gchar* button) {
	return gtk_file_chooser_dialog_new(
			title,
			GTK_WINDOW(parent),
			action,
			button,
			NULL);
}

static gchar* _gtk_file_chooser_get_filename(GtkWidget* chooser) {
	return gtk_file_chooser_get_filename(GTK_FILE_CHOOSER(chooser));
}

static const gchar* _gtk_entry_get_text(GtkWidget* widget) {
	return gtk_entry_get_text(GTK_ENTRY(widget));
}

static void _gtk_entry_set_text(GtkWidget* widget, gchar* text) {
	gtk_entry_set_text(GTK_ENTRY(widget), text);
}

static const gchar* _gtk_label_get_text(GtkWidget* widget) {
	return gtk_label_get_text(GTK_LABEL(widget));
}

static void _gtk_label_set_text(GtkWidget* widget, gchar* text) {
	gtk_label_set_text(GTK_LABEL(widget), text);
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

static const gchar* _gtk_button_get_label(GtkWidget* widget) {
	return gtk_button_get_label(GTK_BUTTON(widget));
}

static void _gtk_button_set_label(GtkWidget* widget, gchar* label) {
	gtk_button_set_label(GTK_BUTTON(widget), label);
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

static gint _gtk_combo_box_get_wrap_width(GtkWidget* widget) {
	return gtk_combo_box_get_wrap_width(GTK_COMBO_BOX(widget));
}

static void _gtk_combo_box_set_wrap_width(GtkWidget* widget, gint width) {
	gtk_combo_box_set_wrap_width(GTK_COMBO_BOX(widget), width);
}

static void _gtk_combo_box_append_text(GtkWidget* widget, gchar* text) {
	gtk_combo_box_append_text(GTK_COMBO_BOX(widget), text);
}

static void _gtk_combo_box_insert_text(GtkWidget* widget, gint position, gchar* text) {
	gtk_combo_box_insert_text(GTK_COMBO_BOX(widget), position, text);
}

static void _gtk_combo_box_prepend_text(GtkWidget* widget, gchar* text) {
	gtk_combo_box_prepend_text(GTK_COMBO_BOX(widget), text);
}

static void _gtk_combo_box_remove_text(GtkWidget* widget, gint position) {
	gtk_combo_box_remove_text(GTK_COMBO_BOX(widget), position);
}

static gchar* _gtk_combo_box_get_active_text(GtkWidget* widget) {
	return gtk_combo_box_get_active_text(GTK_COMBO_BOX(widget));
}

static gint _gtk_combo_box_get_active(GtkWidget* widget) {
	return gtk_combo_box_get_active(GTK_COMBO_BOX(widget));
}

static void _gtk_combo_box_set_active(GtkWidget* widget, gint index_) {
	gtk_combo_box_set_active(GTK_COMBO_BOX(widget), index_);
}

static const gchar* _gtk_combo_box_get_title(GtkWidget* widget) {
	return gtk_combo_box_get_title(GTK_COMBO_BOX(widget));
}

static void _gtk_combo_box_set_title(GtkWidget* widget, gchar* title) {
	gtk_combo_box_set_title(GTK_COMBO_BOX(widget), title);
}

static GtkTreeModel* _gtk_combo_box_get_model(GtkWidget* widget) {
	return (GtkTreeModel*)gtk_combo_box_get_model(GTK_COMBO_BOX(widget));
}

static void _gtk_combo_box_set_model(GtkWidget* widget, GtkTreeModel* model) {
	gtk_combo_box_set_model(GTK_COMBO_BOX(widget), model);
}

static gboolean _gtk_combo_box_get_focus_on_click(GtkWidget* widget) {
	return gtk_combo_box_get_focus_on_click(GTK_COMBO_BOX(widget));
}

static void _gtk_combo_box_set_focus_on_click(GtkWidget* widget, gboolean focus_on_click) {
	gtk_combo_box_set_focus_on_click(GTK_COMBO_BOX(widget), focus_on_click);
}

static gboolean _gtk_combo_box_get_active_iter(GtkWidget* widget, GtkTreeIter* iter) {
	return gtk_combo_box_get_active_iter(GTK_COMBO_BOX(widget), iter);
}

static void _gtk_combo_box_set_active_iter(GtkWidget* widget, GtkTreeIter* iter) {
	gtk_combo_box_set_active_iter(GTK_COMBO_BOX(widget), iter);
}

static void _gtk_combo_box_popup(GtkWidget* widget) {
	return gtk_combo_box_popup(GTK_COMBO_BOX(widget));
}

static void _gtk_combo_box_popdown(GtkWidget* widget) {
	return gtk_combo_box_popdown(GTK_COMBO_BOX(widget));
}

static gboolean _gtk_combo_box_get_add_tearoffs(GtkWidget* widget) {
	return gtk_combo_box_get_add_tearoffs(GTK_COMBO_BOX(widget));
}

static void _gtk_combo_box_set_add_tearoffs(GtkWidget* widget, gboolean add_tearoffs) {
	gtk_combo_box_set_add_tearoffs(GTK_COMBO_BOX(widget), add_tearoffs);
}

static gint _gtk_combo_box_get_row_span_column(GtkWidget* widget) {
	return gtk_combo_box_get_row_span_column(GTK_COMBO_BOX(widget));
}

static void _gtk_combo_box_set_row_span_column(GtkWidget* widget, gint row_span) {
	gtk_combo_box_set_row_span_column(GTK_COMBO_BOX(widget), row_span);
}

static gint _gtk_combo_box_get_column_span_column(GtkWidget* widget) {
	return gtk_combo_box_get_column_span_column(GTK_COMBO_BOX(widget));
}

static void _gtk_combo_box_set_column_span_column(GtkWidget* widget, gint column_span) {
	gtk_combo_box_set_column_span_column(GTK_COMBO_BOX(widget), column_span);
}

static gint _gtk_combo_box_entry_get_text_column(GtkWidget* widget) {
	return gtk_combo_box_entry_get_text_column(GTK_COMBO_BOX_ENTRY(widget));
}

static void _gtk_combo_box_entry_set_text_column(GtkWidget* widget, gint text_column) {
	gtk_combo_box_entry_set_text_column(GTK_COMBO_BOX_ENTRY(widget), text_column);
}

static const gchar* _gtk_font_button_get_title(GtkWidget* widget) {
	return gtk_font_button_get_title(GTK_FONT_BUTTON(widget));
}

static void _gtk_font_button_set_title(GtkWidget* widget, gchar* title) {
	gtk_font_button_set_title(GTK_FONT_BUTTON(widget), title);
}

static gboolean _gtk_font_button_get_use_size(GtkWidget* widget) {
	return gtk_font_button_get_use_size(GTK_FONT_BUTTON(widget));
}

static void _gtk_font_button_set_use_size(GtkWidget* widget, gboolean use_size) {
	gtk_font_button_set_use_size(GTK_FONT_BUTTON(widget), use_size);
}

static const gchar* _gtk_font_button_get_font_name(GtkWidget* widget) {
	return gtk_font_button_get_font_name(GTK_FONT_BUTTON(widget));
}

static void _gtk_font_button_set_font_name(GtkWidget* widget, gchar* fontname) {
	gtk_font_button_set_font_name(GTK_FONT_BUTTON(widget), fontname);
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

static gboolean _gtk_tree_model_get_iter(GtkTreeModel* tree_model, GtkTreeIter* iter, void* path) {
	return gtk_tree_model_get_iter(tree_model, iter, (GtkTreePath*)path);
}

static GtkTreePath* _gtk_tree_model_get_path(GtkTreeModel* tree_model, GtkTreeIter* iter) {
	return gtk_tree_model_get_path(tree_model, iter);
}

static guint _gtk_statusbar_get_context_id(GtkWidget* widget, gchar* context_description) {
	return gtk_statusbar_get_context_id(GTK_STATUSBAR(widget), context_description);
}

static guint _gtk_statusbar_push(GtkWidget* widget, guint context_id, const gchar *text) {
	return gtk_statusbar_push(GTK_STATUSBAR(widget), context_id, text);
}

static void _gtk_statusbar_pop(GtkWidget* widget, guint context_id) {
	gtk_statusbar_pop(GTK_STATUSBAR(widget), context_id);
}

static void _gtk_statusbar_remove(GtkWidget* widget, guint context_id, guint message_id) {
	gtk_statusbar_remove(GTK_STATUSBAR(widget), context_id, message_id);
}

static GtkWidget* _gtk_radio_button_new_from_widget(GtkWidget *widget) {
	return gtk_radio_button_new_from_widget(GTK_RADIO_BUTTON(widget));
}

static GtkWidget* _gtk_radio_button_new_with_label_from_widget(GtkWidget *widget, const gchar* label) {
	return gtk_radio_button_new_with_label_from_widget(GTK_RADIO_BUTTON(widget), label);
}

static GtkWidget* _gtk_radio_button_new_with_mnemonic_from_widget(GtkWidget *widget, const gchar* label) {
	return gtk_radio_button_new_with_mnemonic_from_widget(GTK_RADIO_BUTTON(widget), label);
}

static GSList* _gtk_radio_button_get_group(GtkWidget *widget) {
	return gtk_radio_button_get_group(GTK_RADIO_BUTTON(widget));
}

static void _gtk_radio_button_set_group(GtkWidget* widget, GSList* group) {
	gtk_radio_button_set_group(GTK_RADIO_BUTTON(widget), group);
}

static const gchar* _gtk_frame_get_label(GtkWidget* widget) {
	return gtk_frame_get_label(GTK_FRAME(widget));
}

static void _gtk_frame_set_label(GtkWidget* widget, gchar* label) {
	gtk_frame_set_label(GTK_FRAME(widget), label);
}

static GtkWidget* _gtk_frame_get_label_widget(GtkWidget* widget) {
	return gtk_frame_get_label_widget(GTK_FRAME(widget));
}

static void _gtk_frame_set_label_widget(GtkWidget* widget, GtkWidget* label_widget) {
	gtk_frame_set_label_widget(GTK_FRAME(widget), label_widget);
}

static void _gtk_frame_get_label_align(GtkWidget* widget, gfloat *xalign, gfloat *yalign) {
	gtk_frame_get_label_align(GTK_FRAME(widget), xalign, yalign);
}

static void _gtk_frame_set_label_align(GtkWidget* widget, gfloat xalign, gfloat yalign) {
	gtk_frame_set_label_align(GTK_FRAME(widget), xalign, yalign);
}

static GtkShadowType _gtk_frame_get_shadow_type(GtkWidget* widget) {
	return gtk_frame_get_shadow_type(GTK_FRAME(widget));
}

static void _gtk_frame_set_shadow_type(GtkWidget* widget, GtkShadowType shadow_type) {
	gtk_frame_set_shadow_type(GTK_FRAME(widget), shadow_type);
}

static gdouble _gtk_adjustment_get_lower(GtkAdjustment* adjustment) {
#if GTK_CHECK_VERSION(2,14,0)
	return gtk_adjustment_get_lower(adjustment);
#else
	return 0f;
#endif
}

static void _gtk_adjustment_set_lower(GtkAdjustment* adjustment, gdouble lower) {
#if GTK_CHECK_VERSION(2,14,0)
	gtk_adjustment_set_lower(adjustment, lower);
#endif
}

static gdouble _gtk_adjustment_get_upper(GtkAdjustment* adjustment) {
#if GTK_CHECK_VERSION(2,14,0)
	return gtk_adjustment_get_upper(adjustment);
#else
	return 0f;
#endif
}

static void _gtk_adjustment_set_upper(GtkAdjustment* adjustment, gdouble upper) {
#if GTK_CHECK_VERSION(2,14,0)
	gtk_adjustment_set_upper(adjustment, upper);
#endif
}

static gdouble _gtk_adjustment_get_step_increment(GtkAdjustment* adjustment) {
#if GTK_CHECK_VERSION(2,14,0)
	return gtk_adjustment_get_step_increment(adjustment);
#else
	return 0f;
#endif
}

static void _gtk_adjustment_set_step_increment(GtkAdjustment* adjustment, gdouble step_increment) {
#if GTK_CHECK_VERSION(2,14,0)
	gtk_adjustment_set_step_increment(adjustment, step_increment);
#endif
}

static gdouble _gtk_adjustment_get_page_increment(GtkAdjustment* adjustment) {
#if GTK_CHECK_VERSION(2,14,0)
	return gtk_adjustment_get_page_increment(adjustment);
#else
	return 0f;
#endif
}

static void _gtk_adjustment_set_page_increment(GtkAdjustment* adjustment, gdouble page_increment) {
#if GTK_CHECK_VERSION(2,14,0)
	gtk_adjustment_set_page_increment(adjustment, page_increment);
#endif
}

static gdouble _gtk_adjustment_get_page_size(GtkAdjustment* adjustment) {
#if GTK_CHECK_VERSION(2,14,0)
	return gtk_adjustment_get_page_size(adjustment);
#else
	return 0f;
#endif
}

static void _gtk_adjustment_set_page_size(GtkAdjustment* adjustment, gdouble page_size) {
#if GTK_CHECK_VERSION(2,14,0)
	gtk_adjustment_set_page_size(adjustment, page_size);
#endif
}

static void _gtk_adjustment_configure(GtkAdjustment* adjustment, gdouble value, gdouble lower, gdouble upper, gdouble step_increment, gdouble page_increment, gdouble page_size) {
#if GTK_CHECK_VERSION(2,14,0)
	gtk_adjustment_configure(adjustment, value, lower, upper, step_increment, page_increment, page_size);
#endif
}

static void _gtk_scrolled_window_set_hadjustment(GtkWidget* widget, GtkAdjustment* hadjustment) {
	gtk_scrolled_window_set_hadjustment(GTK_SCROLLED_WINDOW(widget), hadjustment);
}

static void _gtk_scrolled_window_set_vadjustment(GtkWidget* widget, GtkAdjustment* vadjustment) {
	gtk_scrolled_window_set_vadjustment(GTK_SCROLLED_WINDOW(widget), vadjustment);
}

static GtkAdjustment* _gtk_scrolled_window_get_hadjustment(GtkWidget* widget) {
	return gtk_scrolled_window_get_hadjustment(GTK_SCROLLED_WINDOW(widget));
}

static GtkAdjustment* _gtk_scrolled_window_get_vadjustment(GtkWidget* widget) {
	return gtk_scrolled_window_get_vadjustment(GTK_SCROLLED_WINDOW(widget));
}

static void _gtk_text_tag_table_add(GtkTextTagTable* table, void* tag) {
	gtk_text_tag_table_add(table, (GtkTextTag*)tag);
}
 
static void _gtk_text_tag_table_remove(GtkTextTagTable* table, void* tag) {
	gtk_text_tag_table_remove(table, (GtkTextTag*)tag);
}
 
static void* _gtk_text_tag_table_lookup(GtkTextTagTable* table, const gchar* name) {
	return gtk_text_tag_table_lookup(table, name);
}
 
static void* _gtk_text_iter_get_buffer(GtkTextIter* iter) {
	return gtk_text_iter_get_buffer(iter);
}

static void* _gtk_text_buffer_new(GtkTextTagTable* tagtable) {
	return gtk_text_buffer_new(tagtable);
}
 
static gint _gtk_text_buffer_get_line_count(void* buffer) {
	return gtk_text_buffer_get_line_count(GTK_TEXT_BUFFER(buffer));
}
 
static gint _gtk_text_buffer_get_char_count(void* buffer) {
	return gtk_text_buffer_get_char_count(GTK_TEXT_BUFFER(buffer));
}

static GtkTextTagTable* _gtk_text_buffer_get_tag_table(void* buffer) {
	return gtk_text_buffer_get_tag_table(GTK_TEXT_BUFFER(buffer));
}
static void _gtk_text_buffer_insert(void* buffer, GtkTextIter* iter, const gchar* text, gint len) {
	gtk_text_buffer_insert(GTK_TEXT_BUFFER(buffer), iter, text, len);
}

static void _gtk_text_buffer_insert_at_cursor(void *buffer, const gchar *text, gint len) {
	gtk_text_buffer_insert_at_cursor(GTK_TEXT_BUFFER(buffer), text, len);
}

static gboolean _gtk_text_buffer_insert_interactive(void* buffer, GtkTextIter* iter, const gchar* text, gint len, gboolean default_editable) {
	return gtk_text_buffer_insert_interactive(GTK_TEXT_BUFFER(buffer), iter, text, len, default_editable);
}

static gboolean _gtk_text_buffer_insert_interactive_at_cursor(void* buffer, const gchar* text, gint len, gboolean default_editable) {
	return gtk_text_buffer_insert_interactive_at_cursor(GTK_TEXT_BUFFER(buffer), text, len, default_editable);
}

static void _gtk_text_buffer_insert_range(void* buffer, GtkTextIter* iter, const GtkTextIter* start, const GtkTextIter* end) {
	gtk_text_buffer_insert_range(GTK_TEXT_BUFFER(buffer), iter, start, end);
}

static gboolean _gtk_text_buffer_insert_range_interactive(void* buffer, GtkTextIter* iter, const GtkTextIter* start, const GtkTextIter* end, gboolean default_editable) {
	return gtk_text_buffer_insert_range_interactive(GTK_TEXT_BUFFER(buffer), iter, start, end, default_editable);
}

//static void _gtk_text_buffer_insert_with_tags(void* buffer, GtkTextIter* iter, const gchar* text, gint len, GtkTextTag* first_tag, ...);

//static void _gtk_text_buffer_insert_with_tags_by_name(void* buffer, GtkTextIter* iter, const gchar* text, gint len, const gchar* first_tag_name, ...);

static void _gtk_text_buffer_delete(void* buffer, GtkTextIter* start, GtkTextIter* end) {
	gtk_text_buffer_delete(GTK_TEXT_BUFFER(buffer), start, end);
}

static gboolean _gtk_text_buffer_delete_interactive(void* buffer, GtkTextIter* start_iter, GtkTextIter* end_iter, gboolean default_editable) {
	return gtk_text_buffer_delete_interactive(GTK_TEXT_BUFFER(buffer), start_iter, end_iter, default_editable);
}

static gboolean _gtk_text_buffer_backspace(void* buffer, GtkTextIter* iter, gboolean interactive, gboolean default_editable) {
	return gtk_text_buffer_backspace(GTK_TEXT_BUFFER(buffer), iter, interactive, default_editable);
}

static void _gtk_text_buffer_set_text(void* buffer, const gchar* text, gint len) {
	gtk_text_buffer_set_text(GTK_TEXT_BUFFER(buffer), text, len);
}

static gchar* _gtk_text_buffer_get_text(void* buffer, const GtkTextIter* start, const GtkTextIter* end, gboolean include_hidden_chars) {
	return gtk_text_buffer_get_text(GTK_TEXT_BUFFER(buffer), start, end, include_hidden_chars);
}

static gchar* _gtk_text_buffer_get_slice(void* buffer, const GtkTextIter* start, const GtkTextIter* end, gboolean include_hidden_chars) {
	return gtk_text_buffer_get_slice(GTK_TEXT_BUFFER(buffer), start, end, include_hidden_chars);
}

// static void _gtk_text_buffer_insert_pixbuf(void* buffer, GtkTextIter* iter, GdkPixbuf* pixbuf) {
// 	gtk_text_buffer_insert_pixbuf(GTK_TEXT_BUFFER(buffer), iter, pixbuf);
// }

// static void _gtk_text_buffer_insert_child_anchor(void* buffer, GtkTextIter* iter, GtkTextChildAnchor* anchor) {
// 	gtk_text_buffer_insert_child_anchor(GTK_TEXT_BUFFER(buffer), iter, anchor);
// }

// static GtkTextChildAnchor* _gtk_text_buffer_create_child_anchor(void* buffer, GtkTextIter* iter) {
// 	return gtk_text_buffer_create_child_anchor(GTK_TEXT_BUFFER(buffer), iter);
// }
//
static GtkTextMark* _gtk_text_buffer_create_mark(void* buffer, const gchar* mark_name, const GtkTextIter* where, gboolean left_gravity) {
	return gtk_text_buffer_create_mark(GTK_TEXT_BUFFER(buffer), mark_name, where, left_gravity);
}

static void _gtk_text_buffer_move_mark(void* buffer, GtkTextMark* mark, const GtkTextIter* where) {
	gtk_text_buffer_move_mark(GTK_TEXT_BUFFER(buffer), mark, where);
}

static void _gtk_text_buffer_move_mark_by_name(void* buffer, const gchar* name, const GtkTextIter* where) {
	gtk_text_buffer_move_mark_by_name(GTK_TEXT_BUFFER(buffer), name, where);
}

static void _gtk_text_buffer_add_mark(void* buffer, GtkTextMark* mark, const GtkTextIter* where) {
	gtk_text_buffer_add_mark(GTK_TEXT_BUFFER(buffer), mark, where);
}

static void _gtk_text_buffer_delete_mark(void* buffer, GtkTextMark* mark) {
	gtk_text_buffer_delete_mark(GTK_TEXT_BUFFER(buffer), mark);
}

static void _gtk_text_buffer_delete_mark_by_name(void* buffer, const gchar* name) {
	gtk_text_buffer_delete_mark_by_name(GTK_TEXT_BUFFER(buffer), name);
}

static GtkTextMark* _gtk_text_buffer_get_mark(void* buffer, const gchar* name) {
	return gtk_text_buffer_get_mark(GTK_TEXT_BUFFER(buffer), name);
}

static GtkTextMark* _gtk_text_buffer_get_insert(void* buffer) {
	return gtk_text_buffer_get_insert(GTK_TEXT_BUFFER(buffer));
}

static GtkTextMark* _gtk_text_buffer_get_selection_bound(void* buffer) {
	return gtk_text_buffer_get_selection_bound(GTK_TEXT_BUFFER(buffer));
}

static gboolean _gtk_text_buffer_get_has_selection(void* buffer) {
	return gtk_text_buffer_get_has_selection(GTK_TEXT_BUFFER(buffer));
}

static void _gtk_text_buffer_place_cursor(void* buffer, const GtkTextIter* where) {
	gtk_text_buffer_place_cursor(GTK_TEXT_BUFFER(buffer), where);
}

static void _gtk_text_buffer_select_range(void* buffer, const GtkTextIter* ins, const GtkTextIter* bound) {
	gtk_text_buffer_select_range(GTK_TEXT_BUFFER(buffer), ins, bound);
}

static void _gtk_text_buffer_apply_tag(void* buffer, void* tag, const GtkTextIter* start, const GtkTextIter* end) {
	gtk_text_buffer_apply_tag(GTK_TEXT_BUFFER(buffer), tag, start, end);
}

static void _gtk_text_buffer_remove_tag(void* buffer, void* tag, const GtkTextIter* start, const GtkTextIter* end) {
	gtk_text_buffer_remove_tag(GTK_TEXT_BUFFER(buffer), tag, start, end);
}

static void _gtk_text_buffer_apply_tag_by_name(void* buffer, const gchar* name, const GtkTextIter* start, const GtkTextIter* end) {
	gtk_text_buffer_apply_tag_by_name(GTK_TEXT_BUFFER(buffer), name, start, end);
}

static void _gtk_text_buffer_remove_tag_by_name(void* buffer, const gchar* name, const GtkTextIter* start, const GtkTextIter* end) {
	gtk_text_buffer_remove_tag_by_name(GTK_TEXT_BUFFER(buffer), name, start, end);
}

static void _gtk_text_buffer_remove_all_tags(void* buffer, const GtkTextIter* start, const GtkTextIter* end) {
	gtk_text_buffer_remove_all_tags(GTK_TEXT_BUFFER(buffer), start, end);
}

static void* _gtk_text_buffer_create_tag(void* buffer, const gchar* tag_name) {
	return gtk_text_buffer_create_tag(GTK_TEXT_BUFFER(buffer), tag_name, NULL);
}

static void _gtk_text_buffer_get_iter_at_line_offset(void* buffer, GtkTextIter* iter, gint line_number, gint char_offset) {
	gtk_text_buffer_get_iter_at_line_offset(GTK_TEXT_BUFFER(buffer), iter, line_number, char_offset);
}

static void _gtk_text_buffer_get_iter_at_offset(void* buffer, GtkTextIter* iter, gint char_offset) {
	gtk_text_buffer_get_iter_at_offset(GTK_TEXT_BUFFER(buffer), iter, char_offset);
}

static void _gtk_text_buffer_get_iter_at_line(void* buffer, GtkTextIter* iter, gint line_number) {
	gtk_text_buffer_get_iter_at_line(GTK_TEXT_BUFFER(buffer), iter, line_number);
}

static void _gtk_text_buffer_get_iter_at_line_index(void* buffer, GtkTextIter* iter, gint line_number, gint byte_index) {
	gtk_text_buffer_get_iter_at_line_index(GTK_TEXT_BUFFER(buffer), iter, line_number, byte_index);
}

static void _gtk_text_buffer_get_iter_at_mark(void* buffer, GtkTextIter* iter, GtkTextMark* mark) {
	gtk_text_buffer_get_iter_at_mark(GTK_TEXT_BUFFER(buffer), iter, mark);
}

static void _gtk_text_buffer_get_iter_at_child_anchor(void* buffer, GtkTextIter* iter, GtkTextChildAnchor* anchor) {
	gtk_text_buffer_get_iter_at_child_anchor(GTK_TEXT_BUFFER(buffer), iter, anchor);
}

static void _gtk_text_buffer_get_start_iter(void* buffer, GtkTextIter* iter) {
	gtk_text_buffer_get_start_iter(GTK_TEXT_BUFFER(buffer), iter);
}

static void _gtk_text_buffer_get_end_iter(void* buffer, GtkTextIter* iter) {
	gtk_text_buffer_get_end_iter(GTK_TEXT_BUFFER(buffer), iter);
}

static void _gtk_text_buffer_get_bounds(void* buffer, GtkTextIter* start, GtkTextIter* end) {
	gtk_text_buffer_get_bounds(GTK_TEXT_BUFFER(buffer), start, end);
}

static gboolean _gtk_text_buffer_get_modified(void* buffer) {
	return gtk_text_buffer_get_modified(GTK_TEXT_BUFFER(buffer));
}

static void _gtk_text_buffer_set_modified(void* buffer, gboolean setting) {
	gtk_text_buffer_set_modified(GTK_TEXT_BUFFER(buffer), setting);
}

static gboolean _gtk_text_buffer_delete_selection(void* buffer, gboolean interactive, gboolean default_editable) {
	return gtk_text_buffer_delete_selection(GTK_TEXT_BUFFER(buffer), interactive, default_editable);
}

// static void gtk_text_buffer_paste_clipboard(void* buffer, GtkClipboard* clipboard, void* override_location, gboolean default_editable);
// static void gtk_text_buffer_copy_clipboard(void* buffer, GtkClipboard* clipboard);
// static void gtk_text_buffer_cut_clipboard(void* buffer, GtkClipboard* clipboard, gboolean default_editable);
// static gboolean gtk_text_buffer_get_selection_bounds(void* buffer, GtkTextIter* start, GtkTextIter* end);
// static void gtk_text_buffer_begin_user_action(void* buffer);
// static void gtk_text_buffer_end_user_action(void* buffer);
// static void gtk_text_buffer_add_selection_clipboard(GtkTextBuffer* buffer, GtkClipboard* clipboard);
// static void gtk_text_buffer_remove_selection_clipboard(void* buffer, GtkClipboard* clipboard);
// enum GtkTextBufferTargetInfo;
// gboolean (*GtkTextBufferDeserializeFunc) (GtkTextBuffer *register_buffer, GtkTextBuffer *content_buffer, void *iter, const guint8 *data, gsize length, gboolean create_tags, gpointer user_data, GError **error);
// gboolean gtk_text_buffer_deserialize (GtkTextBuffer *register_buffer, GtkTextBuffer *content_buffer, GdkAtom format, void *iter, const guint8 *data, gsize length, GError **error);
// gboolean gtk_text_buffer_deserialize_get_can_create_tags (GtkTextBuffer *buffer, GdkAtom format);
// void gtk_text_buffer_deserialize_set_can_create_tags (GtkTextBuffer *buffer, GdkAtom format, gboolean can_create_tags); GtkTargetList* gtk_text_buffer_get_copy_target_list(GtkTextBuffer *buffer) { }
// GdkAtom * gtk_text_buffer_get_deserialize_formats (GtkTextBuffer *buffer, gint *n_formats);
// GtkTargetList * gtk_text_buffer_get_paste_target_list (GtkTextBuffer *buffer);
// GdkAtom * gtk_text_buffer_get_serialize_formats (GtkTextBuffer *buffer, gint *n_formats);
// GdkAtom gtk_text_buffer_register_deserialize_format (GtkTextBuffer *buffer, const gchar *mime_type, GtkTextBufferDeserializeFunc function, gpointer user_data, GDestroyNotify user_data_destroy);
// GdkAtom gtk_text_buffer_register_deserialize_tagset (GtkTextBuffer *buffer, const gchar *tagset_name); 
// GdkAtom gtk_text_buffer_register_serialize_format (GtkTextBuffer *buffer, const gchar *mime_type, GtkTextBufferSerializeFunc function, gpointer user_data, GDestroyNotify user_data_destroy);
// GdkAtom gtk_text_buffer_register_serialize_tagset (GtkTextBuffer *buffer, const gchar *tagset_name); static guint8* (*GtkTextBufferSerializeFunc) (GtkTextBuffer *register_buffer, GtkTextBuffer *content_buffer, const void *start, const void *end, gsize *length, gpointer user_data);

// static guint8* _gtk_text_buffer_serialize(void* register_buffer, void* content_buffer, GdkAtom format, const GtkTextIter* start, const GtkTextIter* end, gsize* length) {
// 	return gtk_text_buffer_serialize(GTK_TEXT_BUFFER(register_buffer), GTK_TEXT_BUFFER(content_buffer), format, start, end, length);
// }

// static void _gtk_text_buffer_unregister_deserialize_format(void* buffer, GdkAtom format) {
// 	gtk_text_buffer_unregister_deserialize_format(GTK_TEXT_BUFFER(buffer), format);
// }

// static void _gtk_text_buffer_unregister_serialize_format(void* buffer, GdkAtom format) {
// 	gtk_text_buffer_unregister_serialize_format(GTK_TEXT_BUFFER(buffer), format) {
// }

static GtkWidget* _gtk_text_view_new_with_buffer(void* buffer) {
	return gtk_text_view_new_with_buffer(GTK_TEXT_BUFFER(buffer));
}
 
static void _gtk_text_view_set_buffer(GtkWidget* textview, void* buffer) {
	gtk_text_view_set_buffer(GTK_TEXT_VIEW(textview), GTK_TEXT_BUFFER(buffer));
}

static void* _gtk_text_view_get_buffer(GtkWidget* textview) {
	return gtk_text_view_get_buffer(GTK_TEXT_VIEW(textview));
}
 
static void _append_tag(void* tag, const gchar* prop, const gchar* val) {
	GParamSpec *pspec;
	GValue fromvalue = { 0, };
	GValue tovalue = { 0, };
	pspec = g_object_class_find_property(G_OBJECT_GET_CLASS(tag), prop);
	if (!pspec) return;
	g_value_init(&fromvalue, G_TYPE_STRING);
	g_value_set_string(&fromvalue, val);
	g_value_init(&tovalue, G_PARAM_SPEC_VALUE_TYPE(pspec));
	g_value_transform(&fromvalue, &tovalue);
	g_object_set_property((GObject *)tag, prop, &tovalue);
	g_value_unset(&fromvalue);
	g_value_unset(&tovalue);
}

static const gchar* to_gcharptr(const char* s) { return (gchar*)s; }

static const char* to_charptr(const gchar* s) { return (char*)s; }

static void free_string(char* s) { free(s); }

static GtkAdjustment* to_GtkAdjustment(GtkObject* o) { return GTK_ADJUSTMENT(o); }

static GSList* to_gslist(void* gs) {
	return (GSList*)gs;
}

static int _check_version(int major, int minor, int micro) {
	return GTK_CHECK_VERSION(major, minor, micro);
}
*/
import "C";
import "glib";
import "unsafe";
import "runtime";
import "container/vector";

func bool2gboolean(b bool) C.gboolean {
	if b {
		return C.gboolean(1);
	}
	return C.gboolean(0);
}
func gboolean2bool(b C.gboolean) bool {
	if b != 0 {
		return true;
	}
	return false;
}
func panic_if_version_older(major int, minor int, micro int, message string) {
	if C._check_version(C.int(major), C.int(minor), C.int(micro)) == 0 {
		panic(message);
	}
}

//-----------------------------------------------------------------------
// GtkObject
//-----------------------------------------------------------------------
type ObjectLike interface {
}
type GtkObject struct {
	Object *C.GtkObject;
}

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
	ptr := C.CString(s);
	defer C.free_string(ptr);
	C._gtk_signal_connect(v.Widget, C.to_gcharptr(ptr), C.int(funcs.Len())-1, data);
}
func (v GtkWidget) GetTopLevel() *GtkWidget {
	return &GtkWidget {
		C.gtk_widget_get_toplevel(v.Widget) };
}
func (v GtkWidget) GetTopLevelAsWindow() *GtkWindow {
	return &GtkWindow { GtkContainer { GtkWidget {
		C.gtk_widget_get_toplevel(v.Widget) }}};
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
// TODO
// gtk_container_set_border_width
// gtk_container_get_border_width
// gtk_container_remove
// gtk_container_set_resize_mode
// gtk_container_get_resize_mode
// gtk_container_check_resize
// gtk_container_foreach
// gtk_container_foreach_full
// gtk_container_get_children
// gtk_container_propagate_expose
// gtk_container_set_focus_chain
// gtk_container_unset_focus_chain
// gtk_container_set_reallocate_redraws
// gtk_container_set_focus_child
// gtk_container_set_focus_vadjustment
// gtk_container_get_focus_vadjustment
// gtk_container_set_focus_hadjustment
// gtk_container_get_focus_hadjustment
// gtk_container_resize_children
// gtk_container_class_install_child_property
// gtk_container_class_find_child_property
// gtk_container_class_list_child_properties
// gtk_container_add_with_properties
// gtk_container_child_set
// gtk_container_child_get
// gtk_container_child_set_valist
// gtk_container_child_get_valist
// gtk_container_child_set_property
// gtk_container_child_get_property
// gtk_container_forall

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
		C.gtk_window_new(C.GtkWindowType(t)) }}};
}
func (v GtkWindow) GetTitle() string {
	return C.GoString(C.to_charptr(C._gtk_window_get_title(v.Widget)));
}
func (v GtkWindow) SetTitle(title string) {
	ptr := C.CString(title);
	defer C.free_string(ptr);
	C._gtk_window_set_title(v.Widget, C.to_gcharptr(ptr));
}
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
func (v GtkDialog) AddButton(button_text string, response_id int) *GtkButton {
	ptr := C.CString(button_text);
	defer C.free_string(ptr);
	return &GtkButton { GtkWidget {
		C._gtk_dialog_add_button(v.Widget, C.to_gcharptr(ptr), C.gint(response_id)) }};
}
// TODO
// gtk_dialog_new_with_buttons
// gtk_dialog_add_action_widget
// gtk_dialog_add_buttons
// gtk_dialog_set_response_sensitive
// gtk_dialog_set_default_response
// gtk_dialog_get_response_for_widget
// gtk_dialog_set_has_separator
// gtk_dialog_get_has_separator
// gtk_alternative_dialog_button_order
// gtk_dialog_set_alternative_button_order
// gtk_dialog_set_alternative_button_order_from_array
// gtk_dialog_get_action_area
// gtk_dialog_get_content_area

//-----------------------------------------------------------------------
// GtkFileChooser
//-----------------------------------------------------------------------
const (
	GTK_FILE_CHOOSER_ACTION_OPEN = 0;
	GTK_FILE_CHOOSER_ACTION_SAVE = 1;
	GTK_FILE_CHOOSER_ACTION_SELECT_FOLDER = 2;
	GTK_FILE_CHOOSER_ACTION_CREATE_FOLDER = 3;
)
type GtkFileChooser interface {
	GetFilename() string;
};
type GtkFileChooserWidget struct {
	GtkFileChooser;
	GtkWidget;
}
func (v GtkFileChooserWidget) GetFilename() string {
	return C.GoString(C.to_charptr(C._gtk_file_chooser_get_filename(v.Widget)));
}
// TODO
// void gtk_file_chooser_set_action(GtkFileChooser* chooser, GtkFileChooserAction action);
// GtkFileChooserAction gtk_file_chooser_get_action(GtkFileChooser* chooser);
// void gtk_file_chooser_set_local_only(GtkFileChooser* chooser, gboolean local_only);
// gboolean gtk_file_chooser_get_local_only(GtkFileChooser* chooser);
// void gtk_file_chooser_set_select_multiple(GtkFileChooser* chooser, gboolean select_multiple);
// gboolean gtk_file_chooser_get_select_multiple(GtkFileChooser* chooser);
// void gtk_file_chooser_set_show_hidden(GtkFileChooser* chooser, gboolean show_hidden);
// gboolean gtk_file_chooser_get_show_hidden(GtkFileChooser* chooser);
// void gtk_file_chooser_set_do_overwrite_confirmation(GtkFileChooser* chooser, gboolean do_overwrite_confirmation);
// gboolean gtk_file_chooser_get_do_overwrite_confirmation(GtkFileChooser* chooser);
// void gtk_file_chooser_set_create_folders(GtkFileChooser* chooser, gboolean create_folders);
// gboolean gtk_file_chooser_get_create_folders(GtkFileChooser* chooser);
// void gtk_file_chooser_set_current_name(GtkFileChooser* chooser, const gchar* name);
// gchar*  gtk_file_chooser_get_filename(GtkFileChooser* chooser);
// gboolean gtk_file_chooser_set_filename(GtkFileChooser* chooser, const char* filename);
// gboolean gtk_file_chooser_select_filename(GtkFileChooser* chooser, const char* filename);
// void gtk_file_chooser_unselect_filename(GtkFileChooser* chooser, const char* filename);
// void gtk_file_chooser_select_all(GtkFileChooser* chooser);
// void gtk_file_chooser_unselect_all(GtkFileChooser* chooser);
// GSList*  gtk_file_chooser_get_filenames(GtkFileChooser* chooser);
// gboolean gtk_file_chooser_set_current_folder(GtkFileChooser* chooser, const gchar* filename);
// gchar*  gtk_file_chooser_get_current_folder(GtkFileChooser* chooser);
// gchar*  gtk_file_chooser_get_uri(GtkFileChooser* chooser);
// gboolean gtk_file_chooser_set_uri(GtkFileChooser* chooser, const char* uri);
// gboolean gtk_file_chooser_select_uri(GtkFileChooser* chooser, const char* uri);
// void gtk_file_chooser_unselect_uri(GtkFileChooser* chooser, const char* uri);
// GSList*  gtk_file_chooser_get_uris(GtkFileChooser* chooser);
// gboolean gtk_file_chooser_set_current_folder_uri(GtkFileChooser* chooser, const gchar* uri);
// gchar*  gtk_file_chooser_get_current_folder_uri(GtkFileChooser* chooser);
// GFile*  gtk_file_chooser_get_file(GtkFileChooser* chooser);
// gboolean gtk_file_chooser_set_file(GtkFileChooser* chooser, GFile* file, GError* *error);
// gboolean gtk_file_chooser_select_file(GtkFileChooser* chooser, GFile* file, GError* *error);
// void gtk_file_chooser_unselect_file(GtkFileChooser* chooser, GFile* file);
// GSList*  gtk_file_chooser_get_files(GtkFileChooser* chooser);
// gboolean gtk_file_chooser_set_current_folder_file(GtkFileChooser* chooser, GFile* file, GError* *error);
// GFile*  gtk_file_chooser_get_current_folder_file(GtkFileChooser* chooser);
// void gtk_file_chooser_set_preview_widget(GtkFileChooser* chooser, GtkWidget* preview_widget);
// GtkWidget* gtk_file_chooser_get_preview_widget(GtkFileChooser* chooser);
// void gtk_file_chooser_set_preview_widget_active(GtkFileChooser* chooser, gboolean active);
// gboolean gtk_file_chooser_get_preview_widget_active(GtkFileChooser* chooser);
// void gtk_file_chooser_set_use_preview_label(GtkFileChooser* chooser, gboolean use_label);
// gboolean gtk_file_chooser_get_use_preview_label(GtkFileChooser* chooser);
// char* gtk_file_chooser_get_preview_filename(GtkFileChooser* chooser);
// char* gtk_file_chooser_get_preview_uri(GtkFileChooser* chooser);
// GFile* gtk_file_chooser_get_preview_file(GtkFileChooser* chooser);
// void gtk_file_chooser_set_extra_widget(GtkFileChooser* chooser, GtkWidget* extra_widget);
// GtkWidget* gtk_file_chooser_get_extra_widget(GtkFileChooser* chooser);
// void gtk_file_chooser_add_filter(GtkFileChooser* chooser, GtkFileFilter* filter);
// void gtk_file_chooser_remove_filter(GtkFileChooser* chooser, GtkFileFilter* filter);
// GSList* gtk_file_chooser_list_filters(GtkFileChooser* chooser);
// void gtk_file_chooser_set_filter(GtkFileChooser* chooser, GtkFileFilter* filter);
// GtkFileFilter* gtk_file_chooser_get_filter(GtkFileChooser* chooser);
// gboolean gtk_file_chooser_add_shortcut_folder(GtkFileChooser* chooser, const char* folder, GError* *error);
// gboolean gtk_file_chooser_remove_shortcut_folder(GtkFileChooser* chooser, const char* folder, GError* *error);
// GSList* gtk_file_chooser_list_shortcut_folders(GtkFileChooser* chooser);
// gboolean gtk_file_chooser_add_shortcut_folder_uri(GtkFileChooser* chooser, const char* uri, GError* *error);
// gboolean gtk_file_chooser_remove_shortcut_folder_uri(GtkFileChooser* chooser, const char* uri, GError* *error);
// GSList* gtk_file_chooser_list_shortcut_folder_uris(GtkFileChooser* chooser);

//-----------------------------------------------------------------------
// GtkFileChooserDialog
//-----------------------------------------------------------------------
type GtkFileChooserDialog struct {
	GtkDialog;
};
func FileChooserDialog(title string, parent WindowLike, action int, button string) *GtkFileChooserDialog {
	ptitle := C.CString(title);
	defer C.free_string(ptitle);
	pbutton := C.CString(button);
	defer C.free_string(pbutton);
	return &GtkFileChooserDialog { GtkDialog { GtkWindow { GtkContainer { GtkWidget {
		C._gtk_file_chooser_dialog_new(
			C.to_gcharptr(ptitle),
			parent.ToGtkWidget(),
			C.int(action),
			C.to_gcharptr(pbutton)) }}}}};
}
func (v GtkFileChooserDialog) GetFilename() string {
	return C.GoString(C.to_charptr(C._gtk_file_chooser_get_filename(v.Widget)));
}
// FINISH

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
	ptr := C.CString(message);
	defer C.free_string(ptr);
	return &GtkMessageDialog { GtkDialog { GtkWindow { GtkContainer { GtkWidget {
		C._gtk_message_dialog_new(
			parent.ToGtkWidget(),
			C.GtkDialogFlags(flag),
			C.GtkMessageType(t),
			C.GtkButtonsType(button),
			C.to_gcharptr(ptr)) }}}}};
}
// TODO
// gtk_message_dialog_new_with_markup
// gtk_message_dialog_set_image
// gtk_message_dialog_get_image
// gtk_message_dialog_set_markup
// gtk_message_dialog_format_secondary_text
// gtk_message_dialog_format_secondary_markup
 
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
		C.gtk_vbox_new(bool2gboolean(homogeneous), C.gint(spacing)) }}};
}
// FINISH

//-----------------------------------------------------------------------
// GtkHBox
//-----------------------------------------------------------------------
type GtkHBox GtkBox;
func HBox(homogeneous bool, spacing uint) *GtkBox {
	return &GtkBox { GtkContainer { GtkWidget {
		C.gtk_hbox_new(bool2gboolean(homogeneous), C.gint(spacing)) }}};
}
// FINISH

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
		C.gtk_entry_new() }};
}
func (v GtkEntry) GetText() string {
	return C.GoString(C.to_charptr(C._gtk_entry_get_text(v.Widget)));
}
func (v GtkEntry) SetText(text string) {
	ptr := C.CString(text);
	defer C.free_string(ptr);
	C._gtk_entry_set_text(v.Widget, C.to_gcharptr(ptr));
}
// TODO
// gtk_entry_new_with_buffer
// gtk_entry_get_buffer
// gtk_entry_set_buffer
// gtk_entry_set_visibility
// gtk_entry_get_visibility
// gtk_entry_set_invisible_char
// gtk_entry_get_invisible_char
// gtk_entry_unset_invisible_char
// gtk_entry_set_has_frame
// gtk_entry_get_has_frame
// gtk_entry_set_inner_border
// gtk_entry_set_overwrite_mode
// gtk_entry_get_overwrite_mode
// gtk_entry_set_max_length
// gtk_entry_get_max_length
// gtk_entry_get_text_length
// gtk_entry_set_activates_default
// gtk_entry_get_activates_default
// gtk_entry_set_width_chars
// gtk_entry_get_width_chars
// gtk_entry_set_text
// gtk_entry_get_layout
// gtk_entry_get_layout_offsets
// gtk_entry_set_alignment
// gtk_entry_get_alignment
// gtk_entry_set_completion
// gtk_entry_layout_index_to_text_index
// gtk_entry_text_index_to_layout_index
// gtk_entry_set_cursor_hadjustment
// gtk_entry_get_cursor_hadjustment
// gtk_entry_set_progress_fraction
// gtk_entry_get_progress_fraction
// gtk_entry_set_progress_pulse_step
// gtk_entry_get_progress_pulse_step
// gtk_entry_progress_pulse
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
// gtk_entry_new_with_max_length
// gtk_entry_append_text
// gtk_entry_prepend_text
// gtk_entry_set_position
// gtk_entry_select_region
// gtk_entry_set_editable

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
		C.gtk_image_new() }};
}
func ImageFromFile(filename string) *GtkImage {
	ptr := C.CString(filename);
	defer C.free_string(ptr);
	return &GtkImage { GtkWidget {
		C.gtk_image_new_from_file(C.to_gcharptr(ptr)) }};
}
func ImageFromStock(stock_id string, size int) *GtkImage {
	ptr := C.CString(stock_id);
	defer C.free_string(ptr);
	return &GtkImage { GtkWidget {
		C.gtk_image_new_from_stock(C.to_gcharptr(ptr), C.GtkIconSize(size)) }};
}
// TODO
// gtk_image_new_from_pixmap
// gtk_image_new_from_image
// gtk_image_new_from_file
// gtk_image_new_from_pixbuf
// gtk_image_new_from_stock
// gtk_image_new_from_icon_set
// gtk_image_new_from_animation
// gtk_image_new_from_icon_name
// gtk_image_new_from_gicon
// gtk_image_clear
// gtk_image_set_from_pixmap
// gtk_image_set_from_image
// gtk_image_set_from_file
// gtk_image_set_from_pixbuf
// gtk_image_set_from_stock
// gtk_image_set_from_icon_set
// gtk_image_set_from_animation
// gtk_image_set_from_icon_name
// gtk_image_set_from_gicon
// gtk_image_set_pixel_size
// gtk_image_get_storage_type
// gtk_image_get_pixmap
// gtk_image_get_image
// gtk_image_get_pixbuf
// gtk_image_get_stock
// gtk_image_get_icon_set
// gtk_image_get_animation
// gtk_image_get_icon_name
// gtk_image_get_gicon
// gtk_image_get_pixel_size
// gtk_image_set
// gtk_image_get

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
	ptr := C.CString(label);
	defer C.free_string(ptr);
	return &GtkLabel { GtkWidget {
		C.gtk_label_new(C.to_gcharptr(ptr)) }};
}
func (v GtkLabel) GetLabel() string {
	return C.GoString(C.to_charptr(C._gtk_label_get_text(v.Widget)));
}
func (v GtkLabel) SetLabel(label string) {
	ptr := C.CString(label);
	defer C.free_string(ptr);
	C._gtk_label_set_text(v.Widget, C.to_gcharptr(ptr));
}
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
	ptr := C.CString(label);
	defer C.free_string(ptr);
	return &GtkAccelLabel { GtkWidget {
		C.gtk_accel_label_new(C.to_gcharptr(ptr)) }};
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
		C.gtk_button_new() }};
}
func ButtonWithLabel(label string) *GtkButton {
	ptr := C.CString(label);
	defer C.free_string(ptr);
	return &GtkButton { GtkWidget {
		C.gtk_button_new_with_label(C.to_gcharptr(ptr)) }};
}
func (v GtkButton) GetLabel() string {
	return C.GoString(C.to_charptr(C._gtk_button_get_label(v.Widget)));
}
func (v GtkButton) SetLabel(label string) {
	ptr := C.CString(label);
	defer C.free_string(ptr);
	C._gtk_button_set_label(v.Widget, C.to_gcharptr(ptr));
}
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
		C.gtk_toggle_button_new() }}};
}
func ToggleButtonWithLabel(label string) *GtkToggleButton {
	ptr := C.CString(label);
	defer C.free_string(ptr);
	return &GtkToggleButton { GtkButton { GtkWidget {
		C.gtk_toggle_button_new_with_label(C.to_gcharptr(ptr)) }}};
}
func ToggleButtonWithMnemonic(label string) *GtkToggleButton {
	ptr := C.CString(label);
	defer C.free_string(ptr);
	return &GtkToggleButton { GtkButton { GtkWidget {
		C.gtk_check_button_new_with_mnemonic(C.to_gcharptr(ptr)) }}};
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
		C.gtk_check_button_new() }}}};
}
func CheckButtonWithLabel(label string) *GtkCheckButton {
	ptr := C.CString(label);
	defer C.free_string(ptr);
	return &GtkCheckButton { GtkToggleButton { GtkButton { GtkWidget {
		C.gtk_check_button_new_with_label(C.to_gcharptr(ptr)) }}}};
}
func CheckButtonWithMnemonic(label string) *GtkCheckButton {
	ptr := C.CString(label);
	defer C.free_string(ptr);
	return &GtkCheckButton { GtkToggleButton { GtkButton { GtkWidget {
		C.gtk_check_button_new_with_mnemonic(C.to_gcharptr(ptr)) }}}};
}
// FINISH

//-----------------------------------------------------------------------
// GtkRadioButton
//-----------------------------------------------------------------------
type GtkRadioButton struct {
	GtkCheckButton;
}
func RadioButton(group *glib.SList) *GtkRadioButton {
	if group != nil {
		return &GtkRadioButton { GtkCheckButton { GtkToggleButton { GtkButton { GtkWidget {
			C.gtk_radio_button_new(C.to_gslist(unsafe.Pointer(group.ToSList()))) }}}}};
	}
	return &GtkRadioButton { GtkCheckButton { GtkToggleButton { GtkButton { GtkWidget {
		C.gtk_radio_button_new(nil) }}}}};
}
func RadioButtonWithLabel(group *glib.SList, label string) *GtkRadioButton {
	ptr := C.CString(label);
	defer C.free_string(ptr);
	if group != nil {
		return &GtkRadioButton { GtkCheckButton { GtkToggleButton { GtkButton { GtkWidget {
			C.gtk_radio_button_new_with_label(C.to_gslist(unsafe.Pointer(group.ToSList())), C.to_gcharptr(ptr)) }}}}};
	}
	return &GtkRadioButton { GtkCheckButton { GtkToggleButton { GtkButton { GtkWidget {
		C.gtk_radio_button_new_with_label(nil, C.to_gcharptr(ptr)) }}}}};
}
func RadioButtonFromWidget(w *GtkRadioButton) *GtkRadioButton {
	return &GtkRadioButton { GtkCheckButton { GtkToggleButton { GtkButton { GtkWidget {
		C._gtk_radio_button_new_from_widget(w.Widget) }}}}};
}
func RadioButtonWithLabelFromWidget(w *GtkRadioButton, label string) *GtkRadioButton {
	ptr := C.CString(label);
	defer C.free_string(ptr);
	return &GtkRadioButton { GtkCheckButton { GtkToggleButton { GtkButton { GtkWidget {
		C._gtk_radio_button_new_with_label_from_widget(w.Widget, C.to_gcharptr(ptr)) }}}}};
}
func RadioButtonWithMnemonic(group *glib.SList, label string) *GtkRadioButton {
	ptr := C.CString(label);
	defer C.free_string(ptr);
	if group != nil {
		return &GtkRadioButton { GtkCheckButton { GtkToggleButton { GtkButton { GtkWidget {
			C.gtk_radio_button_new_with_mnemonic(C.to_gslist(unsafe.Pointer(group.ToSList())), C.to_gcharptr(ptr)) }}}}};
	}
	return &GtkRadioButton { GtkCheckButton { GtkToggleButton { GtkButton { GtkWidget {
		C.gtk_radio_button_new_with_label(nil, C.to_gcharptr(ptr)) }}}}};
}
func RadioButtonWithMnemonicFromWidget(w *GtkRadioButton, label string) *GtkRadioButton {
	ptr := C.CString(label);
	defer C.free_string(ptr);
	return &GtkRadioButton { GtkCheckButton { GtkToggleButton { GtkButton { GtkWidget {
		C._gtk_radio_button_new_with_mnemonic_from_widget(w.Widget, C.to_gcharptr(ptr)) }}}}};
}
func (v GtkRadioButton) GetGroup() *glib.SList {
	return glib.FromSList(unsafe.Pointer(C._gtk_radio_button_get_group(v.Widget)));
}
func (v GtkRadioButton) SetGroup(group *glib.SList) {
	if group != nil {
		C._gtk_radio_button_set_group(v.Widget, C.to_gslist(unsafe.Pointer(group)));
	} else {
		C._gtk_radio_button_set_group(v.Widget, nil);
	}
}
// FINISH

//-----------------------------------------------------------------------
// GtkFontButton
//-----------------------------------------------------------------------
type GtkFontButton struct {
	GtkButton;
}
func FontButton() *GtkFontButton {
	return &GtkFontButton { GtkButton { GtkWidget {
		C.gtk_font_button_new() }}};
}
func FontButtonWithFont(fontname string) *GtkFontButton {
	ptr := C.CString(fontname);
	defer C.free_string(ptr);
	return &GtkFontButton { GtkButton { GtkWidget {
		C.gtk_font_button_new_with_font(C.to_gcharptr(ptr)) }}};
}
func (v GtkFontButton) GetTitle() string {
	return C.GoString(C.to_charptr(C._gtk_font_button_get_title(v.Widget)));
}
func (v GtkFontButton) SetTitle(title string) {
	ptr := C.CString(title);
	defer C.free_string(ptr);
	C._gtk_font_button_set_title(v.Widget, C.to_gcharptr(ptr));
}
func (v GtkFontButton) GetUseSize() bool {
	return gboolean2bool(C._gtk_font_button_get_use_size(v.Widget));
}
func (v GtkFontButton) SetUseSize(use_size bool) {
	C._gtk_font_button_set_use_size(v.Widget, bool2gboolean(use_size));
}
func (v GtkFontButton) GetFontName() string {
	return C.GoString(C.to_charptr(C._gtk_font_button_get_font_name(v.Widget)));
}
func (v GtkFontButton) SetFontName(fontname string) {
	ptr := C.CString(fontname);
	defer C.free_string(ptr);
	C._gtk_font_button_set_font_name(v.Widget, C.to_gcharptr(ptr));
}
func (v GtkFontButton) GetShowSize() bool {
	return gboolean2bool(C._gtk_font_button_get_show_size(v.Widget));
}
func (v GtkFontButton) SetShowSize(show_size bool) {
	C._gtk_font_button_set_show_size(v.Widget, bool2gboolean(show_size));
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
		C.gtk_tree_path_new() };
}
func TreePathFromString(path string) *GtkTreePath {
	ptr := C.CString(path);
	defer C.free_string(ptr);
	return &GtkTreePath {
		C.gtk_tree_path_new_from_string(C.to_gcharptr(ptr)) };
}
func TreePathNewFirst() *GtkTreePath {
	return &GtkTreePath {
		C.gtk_tree_path_new_first() };
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
func (v GtkTreePath) Free() {
	C.gtk_tree_path_free(v.TreePath);
}
func (v GtkTreePath) Copy() *GtkTreePath {
	return &GtkTreePath {
		C.gtk_tree_path_copy(v.TreePath) };
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
// TODO
// gtk_tree_path_get_indices

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
// FINISH

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
	ret := gboolean2bool(C._gtk_tree_model_get_iter(v.TreeModel, &iter_, unsafe.Pointer(&path_)));
	iter.TreeIter = &iter_;
	path.TreePath = &path_;
	return ret;
}
func (v GtkTreeModel) GetIterFromString(iter *GtkTreeIter, path_string string) bool {
	var iter_ C.GtkTreeIter;
	ptr := C.CString(path_string);
	defer C.free_string(ptr);
	ret := gboolean2bool(C.gtk_tree_model_get_iter_from_string(v.TreeModel, iter.TreeIter, C.to_gcharptr(ptr)));
	iter.TreeIter = &iter_;
	return ret;
}
func (v GtkTreeModel) GetStringFromIter(iter *GtkTreeIter) string {
	return C.GoString(C.to_charptr(C.gtk_tree_model_get_string_from_iter(v.TreeModel, iter.TreeIter)));
}
func (v GtkTreeModel) GetIterFirst(iter *GtkTreeIter) bool {
	var iter_ C.GtkTreeIter;
	ret := gboolean2bool(C.gtk_tree_model_get_iter_first(v.TreeModel, iter.TreeIter));
	iter.TreeIter = &iter_;
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
// GtkComboBox
//-----------------------------------------------------------------------
type GtkComboBox struct {
	GtkContainer;
}
func ComboBox() *GtkComboBox {
	return &GtkComboBox { GtkContainer { GtkWidget {
		C.gtk_combo_box_new() }}};
}
func ComboBoxWithModel(model *GtkTreeModel) *GtkComboBox {
	return &GtkComboBox { GtkContainer { GtkWidget {
		C.gtk_combo_box_new_with_model(model.TreeModel) }}};
}
func ComboBoxNewText() *GtkComboBox {
	return &GtkComboBox { GtkContainer { GtkWidget {
		C.gtk_combo_box_new_text() }}};
}
func (v GtkComboBox) GetWrapWidth() int {
	return int(C._gtk_combo_box_get_wrap_width(v.Widget));
}
func (v GtkComboBox) SetWrapWidth(width int) {
	C._gtk_combo_box_set_wrap_width(v.Widget, C.gint(width));
}
func (v GtkComboBox) AppendText(text string) {
	ptr := C.CString(text);
	defer C.free_string(ptr);
	C._gtk_combo_box_append_text(v.Widget, C.to_gcharptr(ptr));
}
func (v GtkComboBox) InsertText(text string, position int) {
	ptr := C.CString(text);
	defer C.free_string(ptr);
	C._gtk_combo_box_insert_text(v.Widget, C.gint(position), C.to_gcharptr(ptr));
}
func (v GtkComboBox) PrependText(text string) {
	ptr := C.CString(text);
	defer C.free_string(ptr);
	C._gtk_combo_box_prepend_text(v.Widget, C.to_gcharptr(ptr));
}
func (v GtkComboBox) RemoveText(position int) {
	C._gtk_combo_box_remove_text(v.Widget, C.gint(position));
}
func (v GtkComboBox) GetActiveText() string {
	return C.GoString(C.to_charptr(C._gtk_combo_box_get_active_text(v.Widget)));
}
func (v GtkComboBox) GetActive() int {
	return int(C._gtk_combo_box_get_active(v.Widget));
}
func (v GtkComboBox) SetActive(width int) {
	C._gtk_combo_box_set_active(v.Widget, C.gint(width));
}
func (v GtkComboBox) GetTitle() string {
	return C.GoString(C.to_charptr(C._gtk_combo_box_get_title(v.Widget)));
}
func (v GtkComboBox) SetTitle(title string) {
	ptr := C.CString(title);
	defer C.free_string(ptr);
	C._gtk_combo_box_set_title(v.Widget, C.to_gcharptr(ptr));
}
func (v GtkComboBox) GetModel() *GtkTreeModel {
	return &GtkTreeModel {
		C._gtk_combo_box_get_model(v.Widget) };
}
func (v GtkComboBox) SetModel(model *GtkTreeModel) {
	C._gtk_combo_box_set_model(v.Widget, model.TreeModel);
}
func (v GtkComboBox) GetFocusOnClick() bool {
	return gboolean2bool(C._gtk_combo_box_get_focus_on_click(v.Widget));
}
func (v GtkComboBox) SetFocusOnClick(focus_on_click bool) {
	C._gtk_combo_box_set_focus_on_click(v.Widget, bool2gboolean(focus_on_click));
}
func (v GtkComboBox) GetActiveIter(iter *GtkTreeIter) bool {
	var iter_ C.GtkTreeIter;
	ret := gboolean2bool(C._gtk_combo_box_get_active_iter(v.Widget, &iter_));
	iter.TreeIter = &iter_;
	return ret;
}
func (v GtkComboBox) SetActiveIter(iter *GtkTreeIter) {
	C._gtk_combo_box_set_active_iter(v.Widget, iter.TreeIter);
}
func (v GtkComboBox) Popup() {
	C._gtk_combo_box_popup(v.Widget);
}
func (v GtkComboBox) Popdown() {
	C._gtk_combo_box_popdown(v.Widget);
}
func (v GtkComboBox) GetAddTearoffs() bool {
	return gboolean2bool(C._gtk_combo_box_get_add_tearoffs(v.Widget));
}
func (v GtkComboBox) SetAddTearoffs(add_tearoffs bool) {
	C._gtk_combo_box_set_add_tearoffs(v.Widget, bool2gboolean(add_tearoffs));
}
func (v GtkComboBox) GetRowSpanColumn() int {
	return int(C._gtk_combo_box_get_row_span_column(v.Widget));
}
func (v GtkComboBox) SetRowSpanColumn(row_span int) {
	C._gtk_combo_box_set_row_span_column(v.Widget, C.gint(row_span));
}
func (v GtkComboBox) GetColumnSpanColumn() int {
	return int(C._gtk_combo_box_get_column_span_column(v.Widget));
}
func (v GtkComboBox) SetColumnSpanColumn(column_span int) {
	C._gtk_combo_box_set_column_span_column(v.Widget, C.gint(column_span));
}
// TODO
// gtk_combo_box_get_popup_accessible
// gtk_combo_box_get_row_separator_func
// gtk_combo_box_set_row_separator_func
// gtk_combo_box_set_button_sensitivity
// gtk_combo_box_get_button_sensitivity

//-----------------------------------------------------------------------
// GtkComboBoxEntry
//-----------------------------------------------------------------------
type GtkComboBoxEntry struct {
	GtkComboBox;
}
func ComboBoxEntry() *GtkComboBoxEntry {
	return &GtkComboBoxEntry { GtkComboBox { GtkContainer { GtkWidget {
		C.gtk_combo_box_entry_new() }}}};
}
func ComboBoxEntryNewText() *GtkComboBoxEntry {
	return &GtkComboBoxEntry { GtkComboBox { GtkContainer { GtkWidget {
		C.gtk_combo_box_entry_new_text() }}}};
}
func (v GtkComboBoxEntry) GetTextColumn() int {
	return int(C._gtk_combo_box_entry_get_text_column(v.Widget));
}
func (v GtkComboBoxEntry) SetTextColumn(text_column int) {
	C._gtk_combo_box_entry_set_text_column(v.Widget, C.gint(text_column));
}
// FINISH

//-----------------------------------------------------------------------
// GtkStatusbar
//-----------------------------------------------------------------------
type GtkStatusbar struct {
	GtkHBox;
}
func Statusbar() *GtkStatusbar {
	return &GtkStatusbar { GtkHBox { GtkContainer { GtkWidget {
		C.gtk_statusbar_new() }}}};
}
func (v GtkStatusbar) GetContextId(content_description string) uint {
	ptr := C.CString(content_description);
	defer C.free_string(ptr);
	return uint(C._gtk_statusbar_get_context_id(v.Widget, C.to_gcharptr(ptr)));
}
func (v GtkStatusbar) Push(context_id uint, text string) {
	ptr := C.CString(text);
	defer C.free_string(ptr);
	C._gtk_statusbar_push(v.Widget, C.guint(context_id), C.to_gcharptr(ptr));
}
func (v GtkStatusbar) Pop(context_id uint) {
	C._gtk_statusbar_pop(v.Widget, C.guint(context_id));
}
func (v GtkStatusbar) Remove(context_id uint, message_id uint) {
	C._gtk_statusbar_remove(v.Widget, C.guint(context_id), C.guint(message_id));
}
func (v GtkStatusbar) GetHasResizeGrip() bool {
	return gboolean2bool(C._gtk_combo_box_get_add_tearoffs(v.Widget));
}
func (v GtkStatusbar) SetHasResizeGrip(add_tearoffs bool) {
	C._gtk_combo_box_set_add_tearoffs(v.Widget, bool2gboolean(add_tearoffs));
}
// FINISH

//-----------------------------------------------------------------------
// GtkFrame
//-----------------------------------------------------------------------
const (
	GTK_SHADOW_NONE = 0;
	GTK_SHADOW_IN = 1;
	GTK_SHADOW_OUT = 2;
	GTK_SHADOW_ETCHED_IN = 3;
	GTK_SHADOW_ETCHED_OUT = 4;
)
type GtkFrame struct {
	GtkContainer;
}
func Frame(label string) *GtkFrame {
	ptr := C.CString(label);
	defer C.free_string(ptr);
	return &GtkFrame { GtkContainer { GtkWidget {
		C.gtk_frame_new(C.to_gcharptr(ptr)) }}};
}
func (v GtkFrame) GetLabel() string {
	return C.GoString(C.to_charptr(C._gtk_frame_get_label(v.Widget)));
}
func (v GtkFrame) SetLabel(label string) {
	ptr := C.CString(label);
	defer C.free_string(ptr);
	C._gtk_frame_set_label(v.Widget, C.to_gcharptr(ptr));
}
func (v GtkFrame) GetLabelWidget() LabelLike {
	return &GtkLabel { GtkWidget {
		C._gtk_frame_get_label_widget(v.Widget) }};
}
func (v GtkFrame) SetLabelWidget(label_widget *LabelLike) {
	C._gtk_frame_set_label_widget(v.Widget, label_widget.ToGtkWidget());
}
func (v GtkFrame) GetLabelAlign() (xalign, yalign float) {
	var xalign_, yalign_ C.gfloat;
	C._gtk_frame_get_label_align(v.Widget, &xalign_, &yalign_);
	return float(xalign_), float(yalign_);
}
func (v GtkFrame) SetLabelAlign(xalign, yalign float) {
	C._gtk_frame_set_label_align(v.Widget, C.gfloat(xalign), C.gfloat(yalign));
}
func (v GtkFrame) GetShadowType() int {
	return int(C._gtk_frame_get_shadow_type(v.Widget));
}
func (v GtkFrame) SetShadowType(shadow_type int) {
	C._gtk_frame_set_shadow_type(v.Widget, C.GtkShadowType(shadow_type));
}
// FINISH

//-----------------------------------------------------------------------
// GtkAdjustment
//-----------------------------------------------------------------------
type GtkAdjustment struct {
	Adjustment *C.GtkAdjustment;
}
func Adjustment(value float64, lower float64, upper float64, step_increment float64, page_increment float64, page_size float64) *GtkAdjustment {
	return &GtkAdjustment {
		C.to_GtkAdjustment(C.gtk_adjustment_new(C.gdouble(value), C.gdouble(lower), C.gdouble(upper), C.gdouble(step_increment), C.gdouble(page_increment), C.gdouble(page_size))) };
}
func (v GtkAdjustment) GetValue() float64 {
	return float64(C.gtk_adjustment_get_value(v.Adjustment));
}
func (v GtkAdjustment) SetValue(value float64) {
	C.gtk_adjustment_set_value(v.Adjustment, C.gdouble(value));
}
func (v GtkAdjustment) GetLower() float64 {
	panic_if_version_older(2,14,0,"gtk_adjustment_get_lower() is not provided on your GTK");
	return float64(C._gtk_adjustment_get_lower(v.Adjustment));
}
func (v GtkAdjustment) SetLower(lower float64) {
	panic_if_version_older(2,14,0,"gtk_adjustment_set_lower() is not provided on your GTK");
	C._gtk_adjustment_set_lower(v.Adjustment, C.gdouble(lower));
}
func (v GtkAdjustment) GetUpper() float64 {
	panic_if_version_older(2,14,0,"gtk_adjustment_get_upper() is not provided on your GTK");
	return float64(C._gtk_adjustment_get_upper(v.Adjustment));
}
func (v GtkAdjustment) SetUpper(upper float64) {
	panic_if_version_older(2,14,0,"gtk_adjustment_set_upper() is not provided on your GTK");
	C._gtk_adjustment_set_upper(v.Adjustment, C.gdouble(upper));
}
func (v GtkAdjustment) GetStepIncrement() float64 {
	panic_if_version_older(2,14,0,"gtk_adjustment_get_step_increment() is not provided on your GTK");
	return float64(C._gtk_adjustment_get_step_increment(v.Adjustment));
}
func (v GtkAdjustment) SetStepIncrement(step_increment float64) {
	panic_if_version_older(2,14,0,"gtk_adjustment_set_step_increment() is not provided on your GTK");
	C._gtk_adjustment_set_step_increment(v.Adjustment, C.gdouble(step_increment));
}
func (v GtkAdjustment) GetPageIncrement() float64 {
	panic_if_version_older(2,14,0,"gtk_adjustment_get_page_increment() is not provided on your GTK");
	return float64(C._gtk_adjustment_get_page_increment(v.Adjustment));
}
func (v GtkAdjustment) SetPageIncrement(page_increment float64) {
	panic_if_version_older(2,14,0,"gtk_adjustment_set_page_increment() is not provided on your GTK");
	C._gtk_adjustment_set_page_increment(v.Adjustment, C.gdouble(page_increment));
}
func (v GtkAdjustment) GetPageSize() float64 {
	panic_if_version_older(2,14,0,"gtk_adjustment_get_page_size() is not provided on your GTK");
	return float64(C._gtk_adjustment_get_page_size(v.Adjustment));
}
func (v GtkAdjustment) SetPageSize(page_size float64) {
	panic_if_version_older(2,14,0,"gtk_adjustment_set_page_size() is not provided on your GTK");
	C._gtk_adjustment_set_page_size(v.Adjustment, C.gdouble(page_size));
}
func (v GtkAdjustment) Configure(value float64, lower float64, upper float64, step_increment float64, page_increment float64, page_size float64) {
	panic_if_version_older(2,14,0,"gtk_adjustment_configure() is not provided on your GTK");
	C._gtk_adjustment_configure(v.Adjustment, C.gdouble(value), C.gdouble(lower), C.gdouble(upper), C.gdouble(step_increment), C.gdouble(page_increment), C.gdouble(page_size));
}
// FINISH

//-----------------------------------------------------------------------
// GtkScrolledWindow
//-----------------------------------------------------------------------
type GtkScrolledWindow struct {
	GtkContainer;
}
func ScrolledWindow(hadjustment *GtkAdjustment, vadjustment *GtkAdjustment) *GtkScrolledWindow {
	return &GtkScrolledWindow { GtkContainer { GtkWidget {
		C.gtk_scrolled_window_new(hadjustment.Adjustment, vadjustment.Adjustment) }}};
}
func (v GtkScrolledWindow) SetHAdjustment(hadjustment *GtkAdjustment) {
	C._gtk_scrolled_window_set_hadjustment(v.Widget, hadjustment.Adjustment);
}
func (v GtkScrolledWindow) SetVAdjustment(vadjustment *GtkAdjustment) {
	C._gtk_scrolled_window_set_vadjustment(v.Widget, vadjustment.Adjustment);
}
func (v GtkScrolledWindow) GetHAdjustment() *GtkAdjustment {
	return &GtkAdjustment {
		C._gtk_scrolled_window_get_hadjustment(v.Widget) };
}
func (v GtkScrolledWindow) GetVAdjustment() *GtkAdjustment {
	return &GtkAdjustment {
		C._gtk_scrolled_window_get_vadjustment(v.Widget) };
}
// TODO
// gtk_scrolled_window_get_hscrollbar
// gtk_scrolled_window_get_vscrollbar
// gtk_scrolled_window_set_policy
// gtk_scrolled_window_get_policy
// gtk_scrolled_window_set_placement
// gtk_scrolled_window_unset_placement
// gtk_scrolled_window_get_placement
// gtk_scrolled_window_set_shadow_type
// gtk_scrolled_window_get_shadow_type
// gtk_scrolled_window_add_with_viewport

//-----------------------------------------------------------------------
// GtkTextTagTable
//-----------------------------------------------------------------------
type GtkTextTagTable struct {
	TextTagTable *C.GtkTextTagTable;
}
func TextTagTable() *GtkTextTagTable {
	return &GtkTextTagTable {
		C.gtk_text_tag_table_new() };
}
func (v GtkTextTagTable) Add(tag *GtkTextTag) {
	C._gtk_text_tag_table_add(v.TextTagTable, unsafe.Pointer(tag.TextTag));
}
func (v GtkTextTagTable) Remove(tag *GtkTextTag) {
	C._gtk_text_tag_table_remove(v.TextTagTable, unsafe.Pointer(tag.TextTag));
}
func (v GtkTextTagTable) Lookup(name string) *GtkTextTag {
	ptr := C.CString(name);
	defer C.free_string(ptr);
	return &GtkTextTag {
		C._gtk_text_tag_table_lookup(v.TextTagTable, C.to_gcharptr(ptr)) };
}
func (v GtkTextTagTable) GetSize() int {
	return int(C.gtk_text_tag_table_get_size(v.TextTagTable));
}
// TODO
// gtk_text_tag_table_foreach

//-----------------------------------------------------------------------
// GtkTextChildAnchor
//-----------------------------------------------------------------------
type GtkTextChildAnchor struct {
	TextChildAnchor *C.GtkTextChildAnchor;
}

//-----------------------------------------------------------------------
// GtkTextMark
//-----------------------------------------------------------------------
type GtkTextMark struct {
	TextMark *C.GtkTextMark;
}

//-----------------------------------------------------------------------
// GtkTextIter
//-----------------------------------------------------------------------
type GtkTextIter struct {
	TextIter C.GtkTextIter;
}
func (v GtkTextIter) GetBuffer() *GtkTextBuffer {
	return &GtkTextBuffer {
		C._gtk_text_iter_get_buffer(&v.TextIter) };
}
func (v GtkTextIter) Copy() *GtkTextIter {
	return &GtkTextIter {
		*C.gtk_text_iter_copy(&v.TextIter) };
}
func (v GtkTextIter) Free() {
	C.gtk_text_iter_free(&v.TextIter);
}
func (v GtkTextIter) GetOffset() int {
	return int(C.gtk_text_iter_get_offset(&v.TextIter));
}
func (v GtkTextIter) GetLine() int {
	return int(C.gtk_text_iter_get_line(&v.TextIter));
}
func (v GtkTextIter) GetLineOffset() int {
	return int(C.gtk_text_iter_get_line_offset(&v.TextIter));
}
func (v GtkTextIter) GetLineIndex() int {
	return int(C.gtk_text_iter_get_line_index(&v.TextIter));
}
func (v GtkTextIter) GetVisibleLineOffset() int {
	return int(C.gtk_text_iter_get_visible_line_offset(&v.TextIter));
}
func (v GtkTextIter) GetVisibleLineIndex() int {
	return int(C.gtk_text_iter_get_visible_line_index(&v.TextIter));
}
func (v GtkTextIter) GetChar() int {
	return int(C.gtk_text_iter_get_char(&v.TextIter));
}
func (v GtkTextIter) GetSlice(end *GtkTextIter) string {
	return C.GoString(C.to_charptr(C.gtk_text_iter_get_slice(&v.TextIter, &end.TextIter)));
}
func (v GtkTextIter) GetText(end *GtkTextIter) string {
	return C.GoString(C.to_charptr(C.gtk_text_iter_get_text(&v.TextIter, &end.TextIter)));
}
func (v GtkTextIter) GetVisibleSlice(end *GtkTextIter) string {
	return C.GoString(C.to_charptr(C.gtk_text_iter_get_visible_slice(&v.TextIter, &end.TextIter)));
}
func (v GtkTextIter) GetVisibleText(end *GtkTextIter) string {
	return C.GoString(C.to_charptr(C.gtk_text_iter_get_visible_text(&v.TextIter, &end.TextIter)));
}
func (v GtkTextIter) GetMarks() *glib.SList {
	return glib.FromSList(unsafe.Pointer(C.gtk_text_iter_get_marks(&v.TextIter)));
}
// TODO
// gtk_text_iter_get_pixbuf
// gtk_text_iter_get_marks
// gtk_text_iter_get_child_anchor
// gtk_text_iter_get_toggled_tags
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
// gtk_text_iter_starts_sentence
// gtk_text_iter_ends_sentence
// gtk_text_iter_inside_sentence
// gtk_text_iter_starts_line
// gtk_text_iter_ends_line
// gtk_text_iter_is_cursor_position
// gtk_text_iter_get_chars_in_line
// gtk_text_iter_get_bytes_in_line
// gtk_text_iter_get_attributes
// gtk_text_iter_get_language
// gtk_text_iter_is_end
// gtk_text_iter_is_start
// gtk_text_iter_forward_char
// gtk_text_iter_backward_char
// gtk_text_iter_forward_chars
// gtk_text_iter_backward_chars
// gtk_text_iter_forward_line
// gtk_text_iter_backward_line
// gtk_text_iter_forward_lines
// gtk_text_iter_backward_lines
// gtk_text_iter_forward_word_end
// gtk_text_iter_backward_word_start
// gtk_text_iter_forward_word_ends
// gtk_text_iter_backward_word_starts
// gtk_text_iter_forward_visible_line
// gtk_text_iter_backward_visible_line
// gtk_text_iter_forward_visible_lines
// gtk_text_iter_backward_visible_lines
// gtk_text_iter_forward_visible_word_end
// gtk_text_iter_backward_visible_word_start
// gtk_text_iter_forward_visible_word_ends
// gtk_text_iter_backward_visible_word_starts
// gtk_text_iter_forward_sentence_end
// gtk_text_iter_backward_sentence_start
// gtk_text_iter_forward_sentence_ends
// gtk_text_iter_backward_sentence_starts
// gtk_text_iter_forward_cursor_position
// gtk_text_iter_backward_cursor_position
// gtk_text_iter_forward_cursor_positions
// gtk_text_iter_backward_cursor_positions
// gtk_text_iter_forward_visible_cursor_position
// gtk_text_iter_backward_visible_cursor_position
// gtk_text_iter_forward_visible_cursor_positions
// gtk_text_iter_backward_visible_cursor_positions
// gtk_text_iter_set_offset
// gtk_text_iter_set_line
// gtk_text_iter_set_line_offset
// gtk_text_iter_set_line_index
// gtk_text_iter_forward_to_end
// gtk_text_iter_forward_to_line_end
// gtk_text_iter_set_visible_line_offset
// gtk_text_iter_set_visible_line_index
// gtk_text_iter_forward_to_tag_toggle
// gtk_text_iter_backward_to_tag_toggle
// gtk_text_iter_forward_find_char
// gtk_text_iter_backward_find_char
// gtk_text_iter_forward_search
// gtk_text_iter_backward_search
// gtk_text_iter_equal
// gtk_text_iter_compare
// gtk_text_iter_in_range
// gtk_text_iter_order

//-----------------------------------------------------------------------
// GtkTextTag
//-----------------------------------------------------------------------
type GtkTextTag struct {
	TextTag unsafe.Pointer;
}

//-----------------------------------------------------------------------
// GtkTextBuffer
//-----------------------------------------------------------------------
type GtkTextBuffer struct {
	TextBuffer unsafe.Pointer;
}
func TextBuffer(tagtable *GtkTextTagTable) *GtkTextBuffer {
	return &GtkTextBuffer {
		C._gtk_text_buffer_new(tagtable.TextTagTable) };
}
func (v GtkTextBuffer) GetLineCount() int {
	return int(C._gtk_text_buffer_get_line_count(v.TextBuffer));
}
func (v GtkTextBuffer) GetCharCount() int {
	return int(C._gtk_text_buffer_get_char_count(v.TextBuffer));
}
func (v GtkTextBuffer) GetTextTagTable() *GtkTextTagTable {
	return &GtkTextTagTable {
		C._gtk_text_buffer_get_tag_table(v.TextBuffer) };
}
func (v GtkTextBuffer) Insert(iter *GtkTextIter, text string) {
	ptr := C.CString(text);
	defer C.free_string(ptr);
	len := C.strlen(ptr);
	C._gtk_text_buffer_insert(v.TextBuffer, &iter.TextIter, C.to_gcharptr(ptr), C.gint(len));
}
func (v GtkTextBuffer) InsertInteractive(iter *GtkTextIter, text string, default_editable bool) bool {
	ptr := C.CString(text);
	defer C.free_string(ptr);
	len := C.strlen(ptr);
	return gboolean2bool(C._gtk_text_buffer_insert_interactive(v.TextBuffer, &iter.TextIter, C.to_gcharptr(ptr), C.gint(len), bool2gboolean(default_editable)));
}
func (v GtkTextBuffer) InsertAtCursor(text string) {
	ptr := C.CString(text);
	defer C.free_string(ptr);
	len := C.strlen(ptr);
	C._gtk_text_buffer_insert_at_cursor(v.TextBuffer, C.to_gcharptr(ptr), C.gint(len));
}
func (v GtkTextBuffer) InsertInteractiveAtCursor(text string, default_editable bool) bool {
	ptr := C.CString(text);
	defer C.free_string(ptr);
	len := C.strlen(ptr);
	return gboolean2bool(C._gtk_text_buffer_insert_interactive_at_cursor(v.TextBuffer, C.to_gcharptr(ptr), C.gint(len), bool2gboolean(default_editable)));
}
func (v GtkTextBuffer) InsertRange(iter *GtkTextIter, start *GtkTextIter, end *GtkTextIter) {
	C._gtk_text_buffer_insert_range(v.TextBuffer, &iter.TextIter, &start.TextIter, &end.TextIter);
}
func (v GtkTextBuffer) InsertRangeInteractive(iter *GtkTextIter, start *GtkTextIter, end *GtkTextIter, default_editable bool) bool {
	return gboolean2bool(C._gtk_text_buffer_insert_range_interactive(v.TextBuffer, &iter.TextIter, &start.TextIter, &end.TextIter, bool2gboolean(default_editable)));
}
func (v GtkTextBuffer) Delete(start *GtkTextIter, end *GtkTextIter) {
	C._gtk_text_buffer_delete(v.TextBuffer, &start.TextIter, &end.TextIter);
}
func (v GtkTextBuffer) DeleteInteractive(start *GtkTextIter, end *GtkTextIter, default_editable bool) bool {
	return gboolean2bool(C._gtk_text_buffer_delete_interactive(v.TextBuffer, &start.TextIter, &end.TextIter, bool2gboolean(default_editable)));
}
func (v GtkTextBuffer) Backspace(iter *GtkTextIter, interactive bool, default_editable bool) bool {
	return gboolean2bool(C._gtk_text_buffer_backspace(v.TextBuffer, &iter.TextIter, bool2gboolean(interactive), bool2gboolean(default_editable)));
}
func (v GtkTextBuffer) SetText(iter *GtkTextIter, text string) {
	ptr := C.CString(text);
	defer C.free_string(ptr);
	len := C.strlen(ptr);
	C._gtk_text_buffer_set_text(v.TextBuffer, C.to_gcharptr(ptr), C.gint(len));
}
func (v GtkTextBuffer) GetText(iter *GtkTextIter, start *GtkTextIter, end *GtkTextIter, include_hidden_chars bool) string {
	return C.GoString(C.to_charptr(C._gtk_text_buffer_get_text(v.TextBuffer, &start.TextIter, &end.TextIter, bool2gboolean(include_hidden_chars))));
}
func (v GtkTextBuffer) GetSlice(iter *GtkTextIter, start *GtkTextIter, end *GtkTextIter, include_hidden_chars bool) string {
	return C.GoString(C.to_charptr(C._gtk_text_buffer_get_slice(v.TextBuffer, &start.TextIter, &end.TextIter, bool2gboolean(include_hidden_chars))));
}
func (v GtkTextBuffer) CreateMark(mark *GtkTextMark, mark_name string, where *GtkTextIter, left_gravity bool) {
	ptr := C.CString(mark_name);
	defer C.free_string(ptr);
	C._gtk_text_buffer_create_mark(v.TextBuffer, C.to_gcharptr(ptr), &where.TextIter, bool2gboolean(left_gravity));
}
func (v GtkTextBuffer) MoveMark(mark *GtkTextMark, where *GtkTextIter) {
	C._gtk_text_buffer_move_mark(v.TextBuffer, mark.TextMark, &where.TextIter);
}
func (v GtkTextBuffer) MoveMarkByName(mark *GtkTextMark, name string, where *GtkTextIter) {
	ptr := C.CString(name);
	defer C.free_string(ptr);
	C._gtk_text_buffer_move_mark_by_name(v.TextBuffer, C.to_gcharptr(ptr), &where.TextIter);
}
func (v GtkTextBuffer) AddMark(mark *GtkTextMark, where *GtkTextIter) {
	C._gtk_text_buffer_add_mark(v.TextBuffer, mark.TextMark, &where.TextIter);
}
func (v GtkTextBuffer) DeleteMark(mark *GtkTextMark) {
	C._gtk_text_buffer_delete_mark(v.TextBuffer, mark.TextMark);
}
func (v GtkTextBuffer) DeleteMarkByName(name string) {
	ptr := C.CString(name);
	defer C.free_string(ptr);
	C._gtk_text_buffer_delete_mark_by_name(v.TextBuffer, C.to_gcharptr(ptr));
}
func (v GtkTextBuffer) GetMark(name string) *GtkTextMark {
	ptr := C.CString(name);
	defer C.free_string(ptr);
	return &GtkTextMark {
		C._gtk_text_buffer_get_mark(v.TextBuffer, C.to_gcharptr(ptr)) };
}
func (v GtkTextBuffer) GetInsert() *GtkTextMark {
	return &GtkTextMark {
		C._gtk_text_buffer_get_insert(v.TextBuffer) };
}
func (v GtkTextBuffer) GetSelectionBound() *GtkTextMark {
	return &GtkTextMark {
		C._gtk_text_buffer_get_selection_bound(v.TextBuffer) };
}
func (v GtkTextBuffer) GetHasSelection() bool {
	return gboolean2bool(C._gtk_text_buffer_get_has_selection(v.TextBuffer));
}
func (v GtkTextBuffer) PlaceCursor(where *GtkTextIter) {
	C._gtk_text_buffer_place_cursor(v.TextBuffer, &where.TextIter);
}
func (v GtkTextBuffer) SelectRange(ins *GtkTextIter, bound *GtkTextIter) {
	C._gtk_text_buffer_select_range(v.TextBuffer, &ins.TextIter, &bound.TextIter);
}
func (v GtkTextBuffer) ApplyTag(tag *GtkTextTag, start *GtkTextIter, end *GtkTextIter) {
	C._gtk_text_buffer_apply_tag(v.TextBuffer, tag.TextTag, &start.TextIter, &end.TextIter);
}
func (v GtkTextBuffer) RemoveTag(tag *GtkTextTag, start *GtkTextIter, end *GtkTextIter) {
	C._gtk_text_buffer_remove_tag(v.TextBuffer, tag.TextTag, &start.TextIter, &end.TextIter);
}
func (v GtkTextBuffer) ApplyTagByName(name string, start *GtkTextIter, end *GtkTextIter) {
	ptr := C.CString(name);
	defer C.free_string(ptr);
	C._gtk_text_buffer_apply_tag_by_name(v.TextBuffer, C.to_gcharptr(ptr), &start.TextIter, &end.TextIter);
}
func (v GtkTextBuffer) RemoveTagByName(name string, start *GtkTextIter, end *GtkTextIter) {
	ptr := C.CString(name);
	defer C.free_string(ptr);
	C._gtk_text_buffer_remove_tag_by_name(v.TextBuffer, C.to_gcharptr(ptr), &start.TextIter, &end.TextIter);
}
func (v GtkTextBuffer) RemoveAllTags(start *GtkTextIter, end *GtkTextIter) {
	C._gtk_text_buffer_remove_all_tags(v.TextBuffer, &start.TextIter, &end.TextIter);
}
func (v GtkTextBuffer) CreateTag(tag_name string, props map[string] string) *GtkTextTag {
	ptr := C.CString(tag_name);
	defer C.free_string(ptr);
	tag := C._gtk_text_buffer_create_tag(v.TextBuffer, C.to_gcharptr(ptr));
	for prop, val := range props {
		pprop := C.CString(prop);
		pval := C.CString(val);
		C._append_tag(tag, C.to_gcharptr(pprop), C.to_gcharptr(pval));
		C.free_string(pprop);
		C.free_string(pval);
	}
	return &GtkTextTag { tag };
}
		/* C.g_value_init(&value, C.GType(C.G_TYPE_STRING)); */

func (v GtkTextBuffer) GetIterAtLineOffset(iter *GtkTextIter, line_number int, char_offset int) {
	C._gtk_text_buffer_get_iter_at_line_offset(v.TextBuffer, &iter.TextIter, C.gint(line_number), C.gint(char_offset));
}
func (v GtkTextBuffer) GetIterAtOffset(iter *GtkTextIter, char_offset int) {
	C._gtk_text_buffer_get_iter_at_offset(v.TextBuffer, &iter.TextIter, C.gint(char_offset));
}
func (v GtkTextBuffer) GetIterAtLine(iter *GtkTextIter, line_number int) {
	C._gtk_text_buffer_get_iter_at_line(v.TextBuffer, &iter.TextIter, C.gint(line_number));
}
func (v GtkTextBuffer) GetIterAtLineIndex(iter *GtkTextIter, line_number int, byte_index int) {
	C._gtk_text_buffer_get_iter_at_line_index(v.TextBuffer, &iter.TextIter, C.gint(line_number), C.gint(byte_index));
}
func (v GtkTextBuffer) GetIterAtMark(iter *GtkTextIter, mark *GtkTextMark) {
	C._gtk_text_buffer_get_iter_at_mark(v.TextBuffer, &iter.TextIter, mark.TextMark);
}
func (v GtkTextBuffer) GetIterAtChildAnchor(iter *GtkTextIter, anchor *GtkTextChildAnchor) {
	C._gtk_text_buffer_get_iter_at_child_anchor(v.TextBuffer, &iter.TextIter, anchor.TextChildAnchor);
}
func (v GtkTextBuffer) GetStartIter(iter *GtkTextIter) {
	C._gtk_text_buffer_get_start_iter(v.TextBuffer, &iter.TextIter);
}
func (v GtkTextBuffer) GetEndIter(iter *GtkTextIter) {
	C._gtk_text_buffer_get_end_iter(v.TextBuffer, &iter.TextIter)
}
func (v GtkTextBuffer) GetBounds(start *GtkTextIter, end *GtkTextIter) {
	C._gtk_text_buffer_get_bounds(v.TextBuffer, &start.TextIter, &end.TextIter)
}
func (v GtkTextBuffer) GetModified() bool {
	return gboolean2bool(C._gtk_text_buffer_get_modified(v.TextBuffer));
}
func (v GtkTextBuffer) SetModified(setting bool) {
	C._gtk_text_buffer_set_modified(v.TextBuffer, bool2gboolean(setting));
}
func (v GtkTextBuffer) DeleteSelection(interactive bool, default_editable bool) {
	C._gtk_text_buffer_delete_selection(v.TextBuffer, bool2gboolean(interactive), bool2gboolean(default_editable));
}
// TODO

//-----------------------------------------------------------------------
// GtkTextView
//-----------------------------------------------------------------------
type GtkTextView struct { GtkContainer; }
func TextView() *GtkTextView {
	return &GtkTextView { GtkContainer { GtkWidget {
		C.gtk_text_view_new() }}};
}
func TextViewWithBuffer(b GtkTextBuffer) *GtkTextView {
	return &GtkTextView { GtkContainer { GtkWidget {
		C._gtk_text_view_new_with_buffer(b.TextBuffer) }}};
}
func (v GtkTextView) SetBuffer(b GtkTextBuffer) {
	C._gtk_text_view_set_buffer(v.Widget, b.TextBuffer);
}
func (v GtkTextView) GetBuffer() *GtkTextBuffer {
	return &GtkTextBuffer {
		C._gtk_text_view_get_buffer(v.Widget) };
}
// TODO
// gboolean gtk_text_view_scroll_to_iter(GtkTextView* text_view, GtkTextIter* iter, gdouble within_margin, gboolean use_align, gdouble xalign, gdouble yalign);
// void gtk_text_view_scroll_to_mark(GtkTextView* text_view, GtkTextMark* mark, gdouble within_margin, gboolean use_align, gdouble xalign, gdouble yalign);
// void gtk_text_view_scroll_mark_onscreen(GtkTextView* text_view, GtkTextMark* mark);
// gboolean gtk_text_view_move_mark_onscreen(GtkTextView* text_view, GtkTextMark* mark);
// gboolean gtk_text_view_place_cursor_onscreen(GtkTextView* text_view);
// void gtk_text_view_get_visible_rect(GtkTextView* text_view, GdkRectangle* visible_rect);
// void gtk_text_view_set_cursor_visible(GtkTextView* text_view, gboolean setting);
// gboolean gtk_text_view_get_cursor_visible(GtkTextView* text_view); 
// void gtk_text_view_get_iter_location(GtkTextView* text_view, const GtkTextIter* iter, GdkRectangle* location);
// void gtk_text_view_get_iter_at_location(GtkTextView* text_view, GtkTextIter* iter, gint x, gint y);
// void gtk_text_view_get_iter_at_position(GtkTextView* text_view, GtkTextIter* iter, gint* trailing, gint x, gint y);
// void gtk_text_view_get_line_yrange(GtkTextView* text_view, const GtkTextIter* iter, gint* y, gint* height);
// void gtk_text_view_get_line_at_y(GtkTextView* text_view, GtkTextIter* target_iter, gint y, gint* line_top);
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
// void gtk_text_view_set_wrap_mode(GtkTextView* text_view, GtkWrapMode wrap_mode);
// GtkWrapMode gtk_text_view_get_wrap_mode(GtkTextView* text_view);
// void gtk_text_view_set_editable(GtkTextView* text_view, gboolean setting);
// gboolean gtk_text_view_get_editable(GtkTextView* text_view);
// void gtk_text_view_set_overwrite(GtkTextView* text_view, gboolean overwrite);
// gboolean gtk_text_view_get_overwrite(GtkTextView* text_view);
// void gtk_text_view_set_accepts_tab(GtkTextView *text_view, gboolean accepts_tab);
// gboolean gtk_text_view_get_accepts_tab(GtkTextView *text_view);
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
// GtkTextAttributes* gtk_text_view_get_default_attributes(GtkTextView* text_view);

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
	for i := 0;i < int(argc); i++ { C.free_string(cargs[i]); }
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
