#ifndef GO_GTKGL_H
#define GO_GTKGL_H

#include <gtkgl/gdkgl.h>
#include <gtkgl/gtkglarea.h>
#include <stdlib.h>
#include <string.h>

void *make_area(int attr_count, int *attrs) {
	if (attr_count > 0 && attrs[attr_count-1] == 0) {
		// already null terminated; use the pointer directly
		return gtk_gl_area_new(attrs);
	} else {
		// make a null terminated copy of the attribute list
		GtkWidget *area = NULL;
		int *zattrs = malloc((attr_count+1)*sizeof(int));
		if (!zattrs) {
			return NULL;
		}
		memcpy(zattrs, attrs, attr_count * sizeof(attrs));
		zattrs[attr_count] = 0;
		area = gtk_gl_area_new(zattrs);
		free(zattrs);
		return area;
	}
}

#endif
