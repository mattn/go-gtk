package glib

/*
#include <glib.h>
static GSList* to_slist(void* sl) {
	return (GSList*)sl;
}

static inline char* to_charptr(gpointer s) { return (char*)s; }
*/
import "C"
import "unsafe"

type SList struct {
	List *C.GSList
}

func FromSList(sl unsafe.Pointer) *SList {
	return &SList{
		C.to_slist(sl)}
}
func (v SList) ToSList() *C.GSList {
	return v.List
}
func (v SList) Data() interface{} {
	return v.List.data
}
func GSListAlloc() *SList {
	return &SList{C.g_slist_alloc()}
}
func (v SList) Free() {
	C.g_slist_free(v.List)
}
func (v SList) Free1() {
	C.g_slist_free1(v.List)
}
func (v SList) Append(data unsafe.Pointer) *SList {
	return &SList{C.g_slist_append(v.List, C.gpointer(data))}
}
func (v SList) Prepend(data unsafe.Pointer) *SList {
	return &SList{C.g_slist_prepend(v.List, C.gpointer(data))}
}
// GSList* g_slist_insert (GSList *list, gpointer data, gint position) G_GNUC_WARN_UNUSED_RESULT;
// GSList* g_slist_insert_sorted (GSList *list, gpointer data, GCompareFunc func) G_GNUC_WARN_UNUSED_RESULT;
// GSList* g_slist_insert_sorted_with_data (GSList *list, gpointer data, GCompareDataFunc func, gpointer user_data) G_GNUC_WARN_UNUSED_RESULT;
func (v SList) InsertBefore(sibling SList, data unsafe.Pointer) *SList {
	return &SList{C.g_slist_insert_before(v.List, sibling.List, C.gpointer(data))}
}
func (v SList) Concat(llink SList) *SList {
	return &SList{C.g_slist_concat(v.List, llink.List)}
}
func (v SList) Remove(data unsafe.Pointer) *SList {
	return &SList{C.g_slist_remove(v.List, C.gconstpointer(data))}
}
func (v SList) RemoveAll(data unsafe.Pointer) *SList {
	return &SList{C.g_slist_remove_all(v.List, C.gconstpointer(data))}
}
func (v SList) RemoveLink(llink SList) *SList {
	return &SList{C.g_slist_delete_link(v.List, llink.List)}
}
func (v SList) DeleteLink(llink SList) *SList {
	return &SList{C.g_slist_delete_link(v.List, llink.List)}
}
func (v SList) Reverse() *SList {
	return &SList{C.g_slist_reverse(v.List)}
}
func (v SList) Copy() *SList {
	return &SList{C.g_slist_copy(v.List)}
}
func (v SList) Nth(n uint) *SList {
	return &SList{C.g_slist_nth(v.List, C.guint(n))}
}
func (v SList) Find(data unsafe.Pointer) *SList {
	return &SList{C.g_slist_find(v.List, C.gconstpointer(data))}
}
// GSList* g_slist_find_custom (GSList *list, gconstpointer data, GCompareFunc func);
func (v SList) Position(llink SList) int {
	return int(C.g_slist_position(v.List, llink.List))
}
func (v SList) Index(data unsafe.Pointer) int {
	return int(C.g_slist_index(v.List, C.gconstpointer(data)))
}
func (v SList) Last() *SList {
	return &SList{C.g_slist_last(v.List)}
}
func (v SList) Length() uint {
	return uint(C.g_slist_length(v.List))
}
func (v SList) ForEach(callback func(interface{}, interface{}), user_data interface{}) {
	for n := uint(0); n < v.Length(); n++ {
		callback(v.Nth(n).Data(), user_data)
	}
}
// GSList* g_slist_sort (GSList *list, GCompareFunc compare_func) G_GNUC_WARN_UNUSED_RESULT;
// GSList* g_slist_sort_with_data (GSList *list, GCompareDataFunc compare_func, gpointer user_data) G_GNUC_WARN_UNUSED_RESULT;
func (v SList) NthData(n uint) interface{} {
	return C.g_slist_nth_data(v.List, C.guint(n))
}

func GPtrToString(p interface{}) string {
	return C.GoString(C.to_charptr(p.(C.gpointer)))
}
