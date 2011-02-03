include $(GOROOT)/src/Make.inc

TARG     = gdkpixbuf
CGOFILES = gdkpixbuf.go

CGO_CFLAGS  = `pkg-config --cflags gdk-pixbuf-2.0`
CGO_LDFLAGS = `pkg-config --libs gdk-pixbuf-2.0`

GC:=$(GC) -I ../glib/_obj

include $(GOROOT)/src/Make.pkg
