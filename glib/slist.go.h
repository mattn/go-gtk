#ifndef GO_GLIB_SLIST_H
#define GO_GLIB_SLIST_H

#include <glib.h>
#include <glib-object.h>
#include <stdint.h>
#include <stdlib.h>
#include <string.h>

static GSList* to_slist(void* sl) {
	return (GSList*)sl;
}

#endif
