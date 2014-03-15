#ifndef GO_GLIB_OBJECT_H 
#define GO_GLIB_OBJECT_H

#include <glib.h>
#include <glib-object.h>
#include <stdint.h>
#include <stdlib.h>
#include <string.h>

static inline gchar* to_gcharptr(const char* s) { return (gchar*)s; }

static inline void free_string(char* s) { free(s); }

static inline GObject* to_GObject(void* o) { return G_OBJECT(o); }

static void _g_object_set_ptr(gpointer object, const gchar *property_name, void* value) {
	g_object_set(object, property_name, value, NULL);
}

static void _g_object_set_addr(gpointer object, const gchar *property_name, void* value) {
	g_object_set(object, property_name, *(gpointer**)value, NULL);
}

#endif
