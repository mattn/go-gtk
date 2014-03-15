package glib

import (
	// #include "list.go.h"
	// #cgo pkg-config: glib-2.0 gobject-2.0
	"C"
	"unsafe"
)

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
