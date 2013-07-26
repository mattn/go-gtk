#ifndef GO_GDK_PIXBUF_H
#define GO_GDK_PIXBUF_H

#include <gdk-pixbuf/gdk-pixbuf.h>
#include <unistd.h>
#include <stdlib.h>
#include <stdint.h>
#include <stdarg.h>
#include <string.h>

//static const gchar* to_gcharptr(const char* s) { return (const gchar*)s; }
static guchar* to_gucharptr(void* s) { return (guchar*)s; }

static void free_string(char* s) { free(s); }

static gchar* to_gcharptr(char* s) { return (gchar*)s; }

static inline GdkPixbuf* toGdkPixbuf(void* p) { return GDK_PIXBUF(p); }

#endif
