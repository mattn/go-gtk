package glib

/*
#include <glib.h>
#include <glib-object.h>
#include <stdlib.h>
#include <string.h>
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
*/
import "C"
import "unsafe"

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
// GList
//-----------------------------------------------------------------------
type GList struct {
	GList *C.GList
}

func (v GList) Data() interface{} {
	return v.GList.data
}
func (v GList) Append(data unsafe.Pointer) *GList {
	return &GList{C.g_list_append(v.GList, C.gpointer(data))}
}
func (v GList) Prepend(data unsafe.Pointer) *GList {
	return &GList{C.g_list_prepend(v.GList, C.gpointer(data))}
}
func (v GList) Insert(data unsafe.Pointer, pos int) *GList {
	return &GList{C.g_list_insert(v.GList, C.gpointer(data), C.gint(pos))}
}
func (v GList) InsertBefore(sib GList, data unsafe.Pointer) *GList {
	return &GList{C.g_list_insert_before(v.GList, sib.GList, C.gpointer(data))}
}
//GList*              g_list_insert_sorted                (GList *list,
//                                                         gpointer data,
//                                                         GCompareFunc func);
func (v GList) Remove(data unsafe.Pointer) *GList {
	return &GList{C.g_list_remove(v.GList, C.gconstpointer(data))}
}
func (v GList) RemoveLink(link GList) *GList {
	return &GList{C.g_list_remove_link(v.GList, link.GList)}
}
func (v GList) DeleteLink(link GList) *GList {
	return &GList{C.g_list_delete_link(v.GList, link.GList)}
}
func (v GList) RemoveAll(data unsafe.Pointer) *GList {
	return &GList{C.g_list_remove_all(v.GList, C.gconstpointer(data))}
}
func (v GList) Free() {
	C.g_list_free(v.GList)
}
func GListAlloc() *GList {
	return &GList{C.g_list_alloc()}
}
func (v GList) Free1() {
	C.g_list_free_1(v.GList)
}
func (v GList) Length() uint {
	return uint(C.g_list_length(v.GList))
}
func (v GList) Copy() *GList {
	return &GList{C.g_list_copy(v.GList)}
}
func (v GList) Reverse() *GList {
	return &GList{C.g_list_reverse(v.GList)}
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
func (v GList) Concat(link GList) *GList {
	return &GList{C.g_list_concat(v.GList, link.GList)}
}
func (v GList) ForEach(callback func(interface{}, interface{}), user_data interface{}) {
	l := v.First()
	for n := uint(0); n < l.Length(); n++ {
		callback(l.NthData(n), user_data)
	}
}
func (v GList) First() *GList {
	return &GList{C.g_list_first(v.GList)}
}
func (v GList) Last() *GList {
	return &GList{C.g_list_last(v.GList)}
}
func (v GList) Nth(n uint) *GList {
	return &GList{C.g_list_nth(v.GList, C.guint(n))}
}
func (v GList) NthData(n uint) interface{} {
	return C.g_list_nth_data(v.GList, C.guint(n))
}
func (v GList) NthPrev(n uint) *GList {
	return &GList{C.g_list_nth_prev(v.GList, C.guint(n))}
}
func (v GList) Find(data unsafe.Pointer) *GList {
	return &GList{C.g_list_find(v.GList, C.gconstpointer(data))}
}
//GList*              g_list_find_custom                  (GList *list,
//                                                         gconstpointer data,
//                                                         GCompareFunc func);
func (v GList) Position(link GList) int {
	return int(C.g_list_position(v.GList, link.GList))
}
func (v GList) Index(data unsafe.Pointer) int {
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
func (v SList) ForEach(callback func(interface{}, interface{}), user_data interface{}) {
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
}
type GObject struct {
	Object unsafe.Pointer
}

func ObjectFromUnsafe(object unsafe.Pointer) *GObject {
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
func (v *GObject) Set() {
	C.g_object_unref(C.gpointer(v.Object))
}

func Utf8Validate(str []byte, len int, bar **byte) bool {
	return gboolean2bool(C._g_utf8_validate(unsafe.Pointer(&str[0]),
		C.int(len), unsafe.Pointer(bar)))
}

func LocaleToUtf8(opsysstring []byte) (ret string, bytes_read int, bytes_written int, error *Error) {
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
		error = ErrorFromNative(unsafe.Pointer(gerror))
	} else {
		error = nil
	}
	return
}

func LocaleFromUtf8(utf8string string) (ret []byte, bytes_read int, bytes_written int, error *Error) {
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
		error = ErrorFromNative(unsafe.Pointer(gerror))
	} else {
		error = nil
	}
	return
}
