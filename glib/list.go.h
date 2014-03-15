#ifndef GO_GLIB_LIST_H
#define GO_GLIB_LIST_H

#include <glib.h>
#include <glib-object.h>
#include <stdint.h>
#include <stdlib.h>
#include <string.h>

static GList* to_list(void* l) {
	return (GList*)l;
}

#endif
