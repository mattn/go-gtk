#ifndef _GO_GTK_H
#define _GO_GTK_H

#include <gtk/gtk.h>

typedef struct {
	GtkMenu *menu;
	gint x;
	gint y;
	gboolean push_in;
	gpointer data;
} _gtk_menu_position_func_info;



extern void _go_gtk_menu_position_func(_gtk_menu_position_func_info* gmpfi);

#endif
