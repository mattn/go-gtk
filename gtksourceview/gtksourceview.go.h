#ifndef GO_GTKSOURCEVIEW_H
#define GO_GTKSOURCEVIEW_H

#include <gtksourceview/gtksourceview.h>
#include <gtksourceview/gtksourcebuffer.h>
#include <gtksourceview/gtksourcelanguage.h>
#include <gtksourceview/gtksourcelanguagemanager.h>
#include <stdlib.h>

static inline gchar** make_strings(int count) {
	return (gchar**)malloc(sizeof(gchar*) * count);
}

static inline void destroy_strings(gchar** strings) {
	free(strings);
}

static inline void set_string(gchar** strings, int n, gchar* str) {
	strings[n] = str;
}

static inline GObject* to_GObject(void* o) { return G_OBJECT(o); }
static inline gchar* to_gcharptr(const char* s) { return (gchar*)s; }
static inline char* to_charptr(const gchar* s) { return (char*)s; }
static inline gchar** next_gcharptr(gchar** s) { return (s+1); }
static inline void free_string(char* s) { free(s); }
static GtkSourceView* to_GtkSourceView(void* w) { return GTK_SOURCE_VIEW(w); }

#endif
