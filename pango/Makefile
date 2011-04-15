include $(GOROOT)/src/Make.inc

TARG     = pango
CGOFILES = pango.go

CGO_CFLAGS  = `pkg-config --cflags pango`
CGO_LDFLAGS = -lpthread `pkg-config --libs pango`
CGO_DEPS=
GC:=$(GC) #-I ../glib/_obj -I ../gdk/_obj -I ../gdkpixbuf/_obj

include $(GOROOT)/src/Make.pkg
