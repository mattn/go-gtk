#ifndef GO_GIO_H
#define GO_GIO_H

#include <gio/gio.h>

#endif

#include <stdlib.h>

static inline void freeCstr(char* s) { free(s); }
static inline char* toCstr(const gchar* s) { return (char*)s; }
