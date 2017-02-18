#ifndef GO_GLIB_H
#define GO_GLIB_H

#include <glib.h>
#include <glib-object.h>
#include <stdint.h>
#include <stdlib.h>
#include <string.h>
static GList* to_list(void* l) {
	return (GList*)l;
}
static GSList* to_slist(void* sl) {
	return (GSList*)sl;
}
static GError* to_error(void* err) {
	return (GError*)err;
}
static inline GObject* to_GObject(void* o) { return G_OBJECT(o); }

static inline gchar* to_gcharptr(const char* s) { return (gchar*)s; }

static inline char* to_charptr(const gchar* s) { return (char*)s; }

static inline char* to_charptr_from_gpointer(gpointer s) { return (char*)s; }

static inline char* to_charptr_voidp(const void* s) { return (char*)s; }

static inline void free_string(char* s) { free(s); }

static inline gchar* next_string(gchar* s) { return (gchar*)(s + strlen(s) + 1); }

static gboolean _g_utf8_validate(void* str, int len, void* ppbar) {
	return g_utf8_validate((const gchar*)str, (gssize)len, (const gchar**)ppbar);
}

static gchar* _g_locale_to_utf8(void* opsysstring, int len, int* bytes_read, int* bytes_written, GError** error) {
	return g_locale_from_utf8((const gchar*)opsysstring, (gssize)len, (gsize*)bytes_read, (gsize*)bytes_written, error);
}

static gchar* _g_locale_from_utf8(char* utf8string, int len, int* bytes_read, int* bytes_written, GError** error) {
	return g_locale_from_utf8((const gchar*)utf8string, (gssize)len, (gsize*)bytes_read, (gsize*)bytes_written, error);
}
static void _g_object_set_ptr(gpointer object, const gchar *property_name, void* value) {
	g_object_set(object, property_name, value, NULL);
}
static void _g_object_set_addr(gpointer object, const gchar *property_name, void* value) {
	g_object_set(object, property_name, *(gpointer**)value, NULL);
}
//static void _g_object_get(gpointer object, const gchar *property_name, void* value) {
//  g_object_get(object, property_name, value, NULL);
//}

static void g_value_init_int(GValue* gv) { g_value_init(gv, G_TYPE_INT); }
static void g_value_init_string(GValue* gv) { g_value_init(gv, G_TYPE_STRING); }

static GValue* init_gvalue_string_type() {
	GValue* gv = g_new0(GValue, 1);
	g_value_init(gv, G_TYPE_STRING);
	return gv;
}
static GValue* init_gvalue_string(gchar* val) {
	GValue* gv = init_gvalue_string_type();
	g_value_set_string(gv, val);
	return gv;
}

static GValue* init_gvalue_int_type() {
	GValue* gv = g_new0(GValue, 1);
	g_value_init(gv, G_TYPE_INT);
	return gv;
}
static GValue* init_gvalue_int(gint val) {
	GValue* gv = init_gvalue_int_type();
	g_value_set_int(gv, val);
	return gv;
}

static GValue* init_gvalue_uint(guint val) { GValue* gv = g_new0(GValue, 1); g_value_init(gv, G_TYPE_UINT); g_value_set_uint(gv, val); return gv; }
static GValue* init_gvalue_int64(gint64 val) { GValue* gv = g_new0(GValue, 1); g_value_init(gv, G_TYPE_INT64); g_value_set_int64(gv, val); return gv; }
static GValue* init_gvalue_double(gdouble val) { GValue* gv = g_new0(GValue, 1); g_value_init(gv, G_TYPE_DOUBLE); g_value_set_double(gv, val); return gv; }
static GValue* init_gvalue_byte(guchar val) { GValue* gv = g_new0(GValue, 1); g_value_init(gv, G_TYPE_UCHAR); g_value_set_uchar(gv, val); return gv; }
static GValue* init_gvalue_bool(gboolean val) { GValue* gv = g_new0(GValue, 1); g_value_init(gv, G_TYPE_BOOLEAN); g_value_set_boolean(gv, val); return gv; }
//static GValue* init_gvalue_pointer(gpointer val) { GValue* gv = g_new0(GValue, 1); g_value_init(gv, G_TYPE_POINTER); g_value_set_pointer(gv, val); return gv; }

typedef struct {
	char *name;
	int func_no;
	void* target;
	uintptr_t* args;
	int args_no;
	void* ret;
	gulong id;
} callback_info;

static uintptr_t callback_info_get_arg(callback_info* cbi, int idx) {
	return cbi->args[idx];
}
extern void _go_glib_callback(callback_info* cbi);
static uintptr_t _glib_callback(void *data, ...) {
	va_list ap;
	callback_info *cbi = (callback_info*) data;

	int i;
	cbi->args = (uintptr_t*)malloc(sizeof(uintptr_t)*cbi->args_no);
	va_start(ap, data);
	for (i = 0; i < cbi->args_no; i++) {
		cbi->args[i] = va_arg(ap, uintptr_t);
	}
	va_end(ap);

	_go_glib_callback(cbi);

	free(cbi->args);

	return *(uintptr_t*)(&cbi->ret);
}

static void free_callback_info(gpointer data, GClosure *closure) {
	g_slice_free(callback_info, data);
}

static callback_info* _g_signal_connect(void* obj, gchar* name, int func_no, int flags) {
	GSignalQuery query;
	callback_info* cbi;
	guint signal_id = g_signal_lookup(name, G_OBJECT_TYPE(obj));
	g_signal_query(signal_id, &query);
	cbi = g_slice_new0(callback_info);
	cbi->name = g_strdup(name);
	cbi->func_no = func_no;
	cbi->args = NULL;
	cbi->target = obj;
	cbi->args_no = query.n_params;
	if (((GConnectFlags)flags) == G_CONNECT_AFTER)
		cbi->id = g_signal_connect_data((gpointer)obj, name, G_CALLBACK(_glib_callback), (gpointer)cbi, free_callback_info, G_CONNECT_AFTER);
	else
		cbi->id = g_signal_connect_data((gpointer)obj, name, G_CALLBACK(_glib_callback), (gpointer)cbi, free_callback_info, G_CONNECT_SWAPPED);
	return cbi;
}

static void _g_signal_emit_by_name(gpointer instance, const gchar *detailed_signal) {
	g_signal_emit_by_name(instance, detailed_signal);
}

static gulong _g_signal_callback_id(callback_info* cbi) {
	return cbi->id;
}

typedef struct {
	int func_no;
	gboolean ret;
	guint id;
} sourcefunc_info;

extern void _go_glib_sourcefunc(sourcefunc_info* sfi);
static gboolean _go_sourcefunc(void *data) {
	sourcefunc_info *sfi = (sourcefunc_info*) data;

	_go_glib_sourcefunc(sfi);

	return sfi->ret;
}

static void free_sourcefunc_info(gpointer data) {
	g_slice_free(sourcefunc_info, data);
}

static sourcefunc_info* _g_idle_add(int func_no) {
	sourcefunc_info* sfi;
	sfi = g_slice_new0(sourcefunc_info);
	sfi->func_no = func_no;
	sfi->id = g_idle_add_full(G_PRIORITY_DEFAULT_IDLE, _go_sourcefunc, sfi, free_sourcefunc_info);
	return sfi;
}

static sourcefunc_info* _g_timeout_add(guint interval, int func_no) {
	sourcefunc_info* sfi;
	sfi = g_slice_new0(sourcefunc_info);
	sfi->func_no = func_no;
	sfi->id = g_timeout_add_full(G_PRIORITY_DEFAULT, interval, _go_sourcefunc, sfi, free_sourcefunc_info);
	return sfi;
}

static void _g_thread_init(GThreadFunctions *vtable) {
#if !GLIB_CHECK_VERSION(2,32,0)
#ifdef	G_THREADS_ENABLED
	g_thread_init(vtable);
#endif
#endif
}

#if !GLIB_CHECK_VERSION(2,28,0)
static const gchar *
g_get_user_runtime_dir (void)
{
#ifndef G_OS_WIN32
  static const gchar *runtime_dir;
  static gsize initialised;

  if (g_once_init_enter (&initialised)) {
    runtime_dir = g_strdup (getenv ("XDG_RUNTIME_DIR"));
    if (runtime_dir == NULL)
      g_warning ("XDG_RUNTIME_DIR variable not set. Falling back to XDG cache dir.");
    g_once_init_leave (&initialised, 1);
  }

  if (runtime_dir)
    return runtime_dir;

  /* Both fallback for UNIX and the default
   * in Windows: use the user cache directory.
   */
#endif

  return g_get_user_cache_dir ();
}
#endif

#endif
