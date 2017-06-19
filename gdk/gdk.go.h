#ifndef GO_GDK_H
#define GO_GDK_H

#include <gdk/gdk.h>
#include <unistd.h>
#include <stdlib.h>
#include <stdint.h>
#include <stdarg.h>
#include <string.h>

static gchar* toGstr(char* s) { return (gchar*)s; }

static void freeCstr(char* s) { free(s); }

static GdkWindow* toGdkWindow(void* w) { return GDK_WINDOW(w); }
static GdkDragContext* toGdkDragContext(void* l) { return (GdkDragContext*)l; }
static GdkFont* toGdkFont(void* l) { return (GdkFont*)l; }

static void* _gdk_display_get_default() {
	return (void*) gdk_display_get_default();
}

#endif
