package glib

import (
	// #include "object.go.h"
	// #cgo pkg-config: glib-2.0 gobject-2.0
	"C"
	"reflect"
	"unsafe"
)

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
		bval := gbool(value.(bool))
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
