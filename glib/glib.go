package glib

/*
#include <glib.h>
static GSList* to_gslist(void* gs) {
	return (GSList*)gs;
}
*/
import "C";
import "unsafe";

type GSList struct {
	List *C.GSList
}
func FromGSList(gs unsafe.Pointer) *GSList {
	return &GSList {
		C.to_gslist(gs)
	};
}
func (v GSList) ToGSList() *C.GSList { return v.List }
