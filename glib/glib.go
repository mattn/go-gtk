package glib

/*
#ifndef uintptr
//#define uintptr unsigned int*
#endif
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

static inline void free_string(char* s) { free(s); }

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
	gboolean ret;
	guint id;
} callback_info;

static uintptr_t callback_info_get_arg(callback_info* cbi, int idx) {
	return cbi->args[idx];
}
extern void _go_glib_callback(void* cbi);
static gboolean _callback(void *data, ...) {
	va_list ap;
	callback_info *cbi = (callback_info*) data;

	int i;
	cbi->args = (uintptr_t*)malloc(sizeof(uintptr_t)*cbi->args_no);
	va_start(ap, data);
	for (i = 0; i < cbi->args_no; i++) {
		cbi->args[i] = va_arg(ap, uintptr_t);
	}
	va_end(ap);

	_go_glib_callback((void*)cbi);

	free(cbi->args);

	return cbi->ret;
}

static void free_callback_info(gpointer data, GClosure *closure) {
	g_slice_free(callback_info, data);
}

static callback_info* _g_signal_connect(void* obj, gchar* name, int func_no) {
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
	cbi->id = g_signal_connect_data((gpointer)obj, name, G_CALLBACK(_callback), cbi, free_callback_info, G_CONNECT_SWAPPED);
	return cbi;
}
static void _g_signal_emit_by_name(gpointer instance, const gchar *detailed_signal) {
	g_signal_emit_by_name(instance, detailed_signal);
}
*/
// #cgo pkg-config: glib-2.0 gobject-2.0
import "C"
import "unsafe"
import "reflect"

var callback_contexts []*CallbackContext

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

func GPtrToString(p interface{}) string {
	return C.GoString(C.to_charptr_from_gpointer(p.(C.gpointer)))
}

//-----------------------------------------------------------------------
// Application
//-----------------------------------------------------------------------
func SetApplicationName(name string) {
	str := C.CString(name)
	defer C.free_string(str)
	C.g_set_application_name(C.to_gcharptr(str))
}

//-----------------------------------------------------------------------
// String Convert
//-----------------------------------------------------------------------
func Utf8Validate(str []byte, len int, bar **byte) bool {
	return gboolean2bool(C._g_utf8_validate(unsafe.Pointer(&str[0]),
		C.int(len), unsafe.Pointer(bar)))
}

func FilenameFromUri(uri string) (filename string, hostname string, err error) {
	str := C.CString(uri)
	defer C.free_string(str)
	var gerror *C.GError
	var ptr *C.gchar
	filename = C.GoString(C.to_charptr(C.g_filename_from_uri(C.to_gcharptr(str), &ptr, &gerror)))
	if unsafe.Pointer(gerror) != nil {
		err = ErrorFromNative(unsafe.Pointer(gerror))
	} else {
		err = nil
	}
	hostname = ""
	if ptr != nil {
		hostname = C.GoString(C.to_charptr(ptr))
	}
	return
}

func FilenameToUri(filename string, hostname string) (uri string, err error) {
	pfilename := C.CString(filename)
	defer C.free_string(pfilename)
	phostname := C.CString(hostname)
	defer C.free_string(phostname)
	var gerror *C.GError
	uri = C.GoString(C.to_charptr(C.g_filename_to_uri(C.to_gcharptr(pfilename), C.to_gcharptr(phostname), &gerror)))
	if unsafe.Pointer(gerror) != nil {
		err = ErrorFromNative(unsafe.Pointer(gerror))
	} else {
		err = nil
	}
	return
}

func LocaleToUtf8(opsysstring []byte) (ret string, bytes_read int, bytes_written int, err *Error) {
	var gerror *C.GError
	var cbytes_read, cbytes_written C.int
	str := C._g_locale_to_utf8(unsafe.Pointer(&opsysstring[0]), C.int(len(opsysstring)), &cbytes_read, &cbytes_written, &gerror)
	if unsafe.Pointer(str) != nil {
		defer C.free_string(C.to_charptr(str))
		ret = C.GoString(C.to_charptr(str))
	} else {
		ret = ""
	}
	bytes_read = int(cbytes_read)
	bytes_written = int(cbytes_written)
	if unsafe.Pointer(gerror) != nil {
		err = ErrorFromNative(unsafe.Pointer(gerror))
	} else {
		err = nil
	}
	return
}

func LocaleFromUtf8(utf8string string) (ret []byte, bytes_read int, bytes_written int, err *Error) {
	var gerror *C.GError
	var cbytes_read, cbytes_written C.int
	src := C.CString(utf8string)
	defer C.free_string(src)
	str := C._g_locale_from_utf8(src, C.int(C.strlen(src)), &cbytes_read, &cbytes_written, &gerror)
	if unsafe.Pointer(str) != nil {
		defer C.free_string(C.to_charptr(str))
		ret = ([]byte)(C.GoString(C.to_charptr(str)))
	} else {
		ret = ([]byte)("")
	}
	bytes_read = int(cbytes_read)
	bytes_written = int(cbytes_written)
	if unsafe.Pointer(gerror) != nil {
		err = ErrorFromNative(unsafe.Pointer(gerror))
	} else {
		err = nil
	}
	return
}

//-----------------------------------------------------------------------
// List
//-----------------------------------------------------------------------
type List struct {
	GList *C.GList
}

func ListFromNative(l unsafe.Pointer) *List {
	return &List{
		C.to_list(l)}
}
func (v List) Data() interface{} {
	return v.GList.data
}
func (v List) Append(data unsafe.Pointer) *List {
	return &List{C.g_list_append(v.GList, C.gpointer(data))}
}
func (v List) Prepend(data unsafe.Pointer) *List {
	return &List{C.g_list_prepend(v.GList, C.gpointer(data))}
}
func (v List) Insert(data unsafe.Pointer, pos int) *List {
	return &List{C.g_list_insert(v.GList, C.gpointer(data), C.gint(pos))}
}
func (v List) InsertBefore(sib List, data unsafe.Pointer) *List {
	return &List{C.g_list_insert_before(v.GList, sib.GList, C.gpointer(data))}
}
//GList*              g_list_insert_sorted                (GList *list,
//                                                         gpointer data,
//                                                         GCompareFunc func);
func (v List) Remove(data unsafe.Pointer) *List {
	return &List{C.g_list_remove(v.GList, C.gconstpointer(data))}
}
func (v List) RemoveLink(link List) *List {
	return &List{C.g_list_remove_link(v.GList, link.GList)}
}
func (v List) DeleteLink(link List) *List {
	return &List{C.g_list_delete_link(v.GList, link.GList)}
}
func (v List) RemoveAll(data unsafe.Pointer) *List {
	return &List{C.g_list_remove_all(v.GList, C.gconstpointer(data))}
}
func (v List) Free() {
	C.g_list_free(v.GList)
}
func GListAlloc() *List {
	return &List{C.g_list_alloc()}
}
func (v List) Free1() {
	C.g_list_free_1(v.GList)
}
func (v List) Length() uint {
	return uint(C.g_list_length(v.GList))
}
func (v List) Copy() *List {
	return &List{C.g_list_copy(v.GList)}
}
func (v List) Reverse() *List {
	return &List{C.g_list_reverse(v.GList)}
}
//GList*              g_list_sort                         (GList *list,
//                                                         GCompareFunc compare_func);
//gint                (*GCompareFunc)                     (gconstpointer a,
//                                                         gconstpointer b);
//GList*              g_list_insert_sorted_with_data      (GList *list,
//                                                         gpointer data,
//                                                         GCompareDataFunc func,
//                                                         gpointer user_data);
//GList*              g_list_sort_with_data               (GList *list,
//                                                         GCompareDataFunc compare_func,
//                                                         gpointer user_data);
//gint                (*GCompareDataFunc)                 (gconstpointer a,
//                                                         gconstpointer b,
//                                                         gpointer user_data);
func (v List) Concat(link List) *List {
	return &List{C.g_list_concat(v.GList, link.GList)}
}
func (v List) ForEach(callback func(interface{}, interface{}), user_datas ...interface{}) {
	var user_data interface{}
	if len(user_datas) > 0 {
		user_data = user_datas[0]
	}
	l := v.First()
	for n := uint(0); n < l.Length(); n++ {
		callback(l.NthData(n), user_data)
	}
}
func (v List) First() *List {
	return &List{C.g_list_first(v.GList)}
}
func (v List) Last() *List {
	return &List{C.g_list_last(v.GList)}
}
func (v List) Nth(n uint) *List {
	return &List{C.g_list_nth(v.GList, C.guint(n))}
}
func (v List) NthData(n uint) interface{} {
	return C.g_list_nth_data(v.GList, C.guint(n))
}
func (v List) NthPrev(n uint) *List {
	return &List{C.g_list_nth_prev(v.GList, C.guint(n))}
}
func (v List) Find(data unsafe.Pointer) *List {
	return &List{C.g_list_find(v.GList, C.gconstpointer(data))}
}
//GList*              g_list_find_custom                  (GList *list,
//                                                         gconstpointer data,
//                                                         GCompareFunc func);
func (v List) Position(link List) int {
	return int(C.g_list_position(v.GList, link.GList))
}
func (v List) Index(data unsafe.Pointer) int {
	return int(C.g_list_index(v.GList, C.gconstpointer(data)))
}

//-----------------------------------------------------------------------
// g_slist
//-----------------------------------------------------------------------
type SList struct {
	GSList *C.GSList
}

func SListFromNative(sl unsafe.Pointer) *SList {
	return &SList{
		C.to_slist(sl)}
}
func (v SList) ToSList() *C.GSList {
	return v.GSList
}
func (v SList) Data() interface{} {
	return v.GSList.data
}
func GSListAlloc() *SList {
	return &SList{C.g_slist_alloc()}
}
func (v SList) Free() {
	C.g_slist_free(v.GSList)
}
func (v SList) Free1() {
	C.g_slist_free1(v.GSList)
}
func (v SList) Append(data unsafe.Pointer) *SList {
	return &SList{C.g_slist_append(v.GSList, C.gpointer(data))}
}
func (v SList) Prepend(data unsafe.Pointer) *SList {
	return &SList{C.g_slist_prepend(v.GSList, C.gpointer(data))}
}
// GSList* g_slist_insert (GSList *list, gpointer data, gint position) G_GNUC_WARN_UNUSED_RESULT;
// GSList* g_slist_insert_sorted (GSList *list, gpointer data, GCompareFunc func) G_GNUC_WARN_UNUSED_RESULT;
// GSList* g_slist_insert_sorted_with_data (GSList *list, gpointer data, GCompareDataFunc func, gpointer user_data) G_GNUC_WARN_UNUSED_RESULT;
func (v SList) InsertBefore(sibling SList, data unsafe.Pointer) *SList {
	return &SList{C.g_slist_insert_before(v.GSList, sibling.GSList, C.gpointer(data))}
}
func (v SList) Concat(llink SList) *SList {
	return &SList{C.g_slist_concat(v.GSList, llink.GSList)}
}
func (v SList) Remove(data unsafe.Pointer) *SList {
	return &SList{C.g_slist_remove(v.GSList, C.gconstpointer(data))}
}
func (v SList) RemoveAll(data unsafe.Pointer) *SList {
	return &SList{C.g_slist_remove_all(v.GSList, C.gconstpointer(data))}
}
func (v SList) RemoveLink(llink SList) *SList {
	return &SList{C.g_slist_delete_link(v.GSList, llink.GSList)}
}
func (v SList) DeleteLink(llink SList) *SList {
	return &SList{C.g_slist_delete_link(v.GSList, llink.GSList)}
}
func (v SList) Reverse() *SList {
	return &SList{C.g_slist_reverse(v.GSList)}
}
func (v SList) Copy() *SList {
	return &SList{C.g_slist_copy(v.GSList)}
}
func (v SList) Nth(n uint) *SList {
	return &SList{C.g_slist_nth(v.GSList, C.guint(n))}
}
func (v SList) Find(data unsafe.Pointer) *SList {
	return &SList{C.g_slist_find(v.GSList, C.gconstpointer(data))}
}
// GSList* g_slist_find_custom (GSList *list, gconstpointer data, GCompareFunc func);
func (v SList) Position(llink SList) int {
	return int(C.g_slist_position(v.GSList, llink.GSList))
}
func (v SList) Index(data unsafe.Pointer) int {
	return int(C.g_slist_index(v.GSList, C.gconstpointer(data)))
}
func (v SList) Last() *SList {
	return &SList{C.g_slist_last(v.GSList)}
}
func (v SList) Length() uint {
	return uint(C.g_slist_length(v.GSList))
}
func (v SList) ForEach(callback func(interface{}, interface{}), user_datas ...interface{}) {
	var user_data interface{}
	if len(user_datas) > 0 {
		user_data = user_datas[0]
	}
	for n := uint(0); n < v.Length(); n++ {
		callback(v.Nth(n).Data(), user_data)
	}
}
// GSList* g_slist_sort (GSList *list, GCompareFunc compare_func) G_GNUC_WARN_UNUSED_RESULT;
// GSList* g_slist_sort_with_data (GSList *list, GCompareDataFunc compare_func, gpointer user_data) G_GNUC_WARN_UNUSED_RESULT;
func (v SList) NthData(n uint) interface{} {
	return C.g_slist_nth_data(v.GSList, C.guint(n))
}

//-----------------------------------------------------------------------
// g_error
//-----------------------------------------------------------------------
type Error struct {
	GError *C.GError
}

func (v *Error) Error() string {
	return v.Message()
}

func (v *Error) Message() string {
	if unsafe.Pointer(v.GError) == nil || unsafe.Pointer(v.GError.message) == nil {
		return ""
	}
	return C.GoString(C.to_charptr(v.GError.message))
}

func ErrorFromNative(err unsafe.Pointer) *Error {
	return &Error{
		C.to_error(err)}
}

//-----------------------------------------------------------------------
// GObject
//-----------------------------------------------------------------------
type ObjectLike interface {
	Ref()
	Unref()
	Connect(s string, f interface{}, data ...interface{})
}
type GObject struct {
	Object unsafe.Pointer
}

func ObjectFromNative(object unsafe.Pointer) *GObject {
	//	return &GObject {
	//		C.to_GObject(object) }
	return &GObject{
		object}
}

func (v *GObject) Ref() {
	C.g_object_ref(C.gpointer(v.Object))
}
func (v *GObject) Unref() {
	C.g_object_unref(C.gpointer(v.Object))
}

func (v *GObject) Set(name string, value interface{}) {
	ptr := C.CString(name)
	defer C.free_string(ptr)

	if _, ok := value.(WrappedObject); ok {
		value = value.(WrappedObject).GetInternalValue()
	}
	if _, ok := value.(GObject); ok {
		value = value.(GObject).Object
	}
	if _, ok := value.(GValue); ok {
		value = value.(GValue).Value
	}

	switch value.(type) {
	case bool:
		bval := bool2gboolean(value.(bool))
		C._g_object_set_addr(C.gpointer(v.Object), C.to_gcharptr(ptr), unsafe.Pointer(&bval))
	case byte:
		bval := C.gchar(value.(byte))
		C._g_object_set_addr(C.gpointer(v.Object), C.to_gcharptr(ptr), unsafe.Pointer(&bval))
		//C._g_object_set(C.gpointer(v.Object), C.to_gcharptr(ptr), unsafe.Pointer(reflect.ValueOf(C.gchar(value.(byte))).UnsafeAddr()))
	case int:
		ival := C.int(value.(int))
		C._g_object_set_addr(C.gpointer(v.Object), C.to_gcharptr(ptr), unsafe.Pointer(&ival))
	case uint:
		uval := C.guint(value.(uint))
		C._g_object_set_addr(C.gpointer(v.Object), C.to_gcharptr(ptr), unsafe.Pointer(&uval))
		//C._g_object_set(C.gpointer(v.Object), C.to_gcharptr(ptr), unsafe.Pointer(reflect.ValueOf(C.guint(value.(uint))).UnsafeAddr()))
	case float32:
		f32val := C.gfloat(value.(float32))
		C._g_object_set_addr(C.gpointer(v.Object), C.to_gcharptr(ptr), unsafe.Pointer(&f32val))
		//C._g_object_set(C.gpointer(v.Object), C.to_gcharptr(ptr), unsafe.Pointer(reflect.ValueOf(C.gfloat(value.(float64))).UnsafeAddr()))
	case float64:
		f64val := C.gfloat(value.(float64))
		C._g_object_set_addr(C.gpointer(v.Object), C.to_gcharptr(ptr), unsafe.Pointer(&f64val))
		//C._g_object_set(C.gpointer(v.Object), C.to_gcharptr(ptr), unsafe.Pointer(reflect.ValueOf(C.gfloat(value.(float64))).UnsafeAddr()))
	case string:
		pval := C.CString(value.(string))
		defer C.free_string(pval)
		C._g_object_set_addr(C.gpointer(v.Object), C.to_gcharptr(ptr), unsafe.Pointer(&pval))
	default:
		if pv, ok := value.(*[0]uint8); ok {
			C._g_object_set_ptr(C.gpointer(v.Object), C.to_gcharptr(ptr), unsafe.Pointer(pv))
		} else {
			av := reflect.ValueOf(value)
			if av.CanAddr() {
				C._g_object_set_addr(C.gpointer(v.Object), C.to_gcharptr(ptr), unsafe.Pointer(reflect.ValueOf(value).UnsafeAddr()))
			} else {
				C._g_object_set_ptr(C.gpointer(v.Object), C.to_gcharptr(ptr), value.(unsafe.Pointer))
			}
		}
	}
}
func (v *GObject) SetProperty(name string, val *GValue) {
	str := C.CString(name)
	defer C.free_string(str)
	C.g_object_set_property(C.to_GObject(v.Object), C.to_gcharptr(str), &val.Value)
}

//-----------------------------------------------------------------------
// GValue
//-----------------------------------------------------------------------
func GValueFromNative(value interface{}) *C.GValue {
	var gv *C.GValue

	if _, ok := value.(WrappedObject); ok {
		value = value.(WrappedObject).GetInternalValue()
	}
	if _, ok := value.(GObject); ok {
		value = value.(GObject).Object
	}
	if _, ok := value.(GValue); ok {
		value = value.(GValue).Value
	}

	switch value.(type) {
	case bool:
		gv = C.init_gvalue_bool(bool2gboolean(value.(bool)))
		break
	case byte:
		gv = C.init_gvalue_byte(C.guchar(value.(byte)))
		break
	case int:
		gv = C.init_gvalue_int(C.gint(value.(int)))
		break
	case uint:
		gv = C.init_gvalue_uint(C.guint(value.(uint)))
		break
	case float32:
		gv = C.init_gvalue_double(C.gdouble(value.(float32)))
		break
	case float64:
		gv = C.init_gvalue_double(C.gdouble(value.(float64)))
		break
	case string:
		{
			pval := C.CString(value.(string))
			defer C.free_string(pval)
			gv = C.init_gvalue_string(C.to_gcharptr(pval))
		}
		break
	default:
		//gv = C.init_gvalue_pointer(C.gpointer(unsafe.Pointer(&value)));
		break
	}
	return gv
}
func ValueFromNative(val interface{}) *GValue {
	return &GValue{*GValueFromNative(val)}
}

type GValue struct {
	Value C.GValue
}

const (
	G_TYPE_CHAR    = 3 << 2
	G_TYPE_UCHAR   = 4 << 2
	G_TYPE_BOOL    = 5 << 2
	G_TYPE_INT     = 6 << 2
	G_TYPE_UINT    = 7 << 2
	G_TYPE_LONG    = 8 << 2
	G_TYPE_ULONG   = 9 << 2
	G_TYPE_INT64   = 10 << 2
	G_TYPE_UINT64  = 11 << 2
	G_TYPE_ENUM    = 12 << 2
	G_TYPE_FLAGS   = 13 << 2
	G_TYPE_FLOAT   = 14 << 2
	G_TYPE_DOUBLE  = 15 << 2
	G_TYPE_STRING  = 16 << 2
	G_TYPE_POINTER = 17 << 2
	G_TYPE_BOXED   = 18 << 2
)

func (v *GValue) Init(t int) {
	if t == G_TYPE_INT {
		C.g_value_init_int(&v.Value)
	} else if t == G_TYPE_STRING {
		C.g_value_init_string(&v.Value)
	}
}

func (v *GValue) GetString() string {
	return C.GoString(C.to_charptr(C.g_value_get_string(&v.Value)))
}

func (v *GValue) GetInt() int {
	return int(C.g_value_get_int(&v.Value))
}

//-----------------------------------------------------------------------
// WrappedObject
//-----------------------------------------------------------------------
type WrappedObject interface {
	GetInternalValue() unsafe.Pointer
}

//-----------------------------------------------------------------------
// Events
//-----------------------------------------------------------------------
// the go-gtk Callback is simpler than the one in C, because we have
// full closures, so there is never a need to pass additional data via
// a void * pointer.  Where you might have wanted to do that, you can
// instead just use func () { ... using data } to pass the data in.
type CallbackContext struct {
	f      interface{}
	cbi    unsafe.Pointer
	target reflect.Value
	data   reflect.Value
}

func (c *CallbackContext) Target() interface{} {
	return c.target.Interface()
}

func (c *CallbackContext) Data() interface{} {
	return c.data.Interface()
}

func (c *CallbackContext) Args(n int) uintptr {
	return uintptr(C.callback_info_get_arg((*C.callback_info)(c.cbi), C.int(n)))
}

//export _go_glib_callback
func _go_glib_callback(pcbi unsafe.Pointer) {
	cbi := (*C.callback_info)(pcbi)
	context := callback_contexts[int(cbi.func_no)]
	rf := reflect.ValueOf(context.f)
	t := rf.Type()
	fargs := make([]reflect.Value, t.NumIn())
	if len(fargs) > 0 {
		fargs[0] = reflect.ValueOf(context)
	}
	ret := rf.Call(fargs)
	if len(ret) > 0 {
		bret, _ := ret[0].Interface().(bool)
		cbi.ret = bool2gboolean(bret)
	}
}

func (v *GObject) Connect(s string, f interface{}, datas ...interface{}) {
	var data interface{}
	if len(datas) > 0 {
		data = datas[0]
	}
	ctx := &CallbackContext{f, nil, reflect.ValueOf(v), reflect.ValueOf(data)}
	ptr := C.CString(s)
	defer C.free_string(ptr)
	ctx.cbi = unsafe.Pointer(C._g_signal_connect(unsafe.Pointer(v.Object), C.to_gcharptr(ptr), C.int(len(callback_contexts))))
	callback_contexts = append(callback_contexts, ctx)
}

func (v *GObject) StopEmission(s string) {
	ptr := C.CString(s)
	defer C.free_string(ptr)
	C.g_signal_stop_emission_by_name((C.gpointer)(v.Object), C.to_gcharptr(ptr))
}

func (v *GObject) Emit(s string) {
	ptr := C.CString(s)
	defer C.free_string(ptr)
	C._g_signal_emit_by_name((C.gpointer)(v.Object), C.to_gcharptr(ptr))
}
