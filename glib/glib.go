package glib

/*
#include <glib.h>
static GSList* to_slist(void* sl) {
	return (GSList*)sl;
}
*/
import "C";
import "unsafe";

type SList struct {
	List *C.GSList
}
func FromSList(sl unsafe.Pointer) *SList {
	return &SList {
		C.to_slist(sl) };
}
func (v SList) ToSList() *C.GSList { return v.List }
