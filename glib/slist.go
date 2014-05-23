package glib

import (
	// #include "slist.go.h"
	// #cgo pkg-config: glib-2.0 gobject-2.0
	"C"
	"unsafe"
)

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
