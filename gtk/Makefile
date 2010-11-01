include $(GOROOT)/src/Make.inc

TARG     = gtk
CGOFILES = gtk.go

CGO_CFLAGS  = `pkg-config --cflags gtk+-2.0 gtksourceview-2.0`
CGO_LDFLAGS = `pkg-config --libs gtk+-2.0 gtksourceview-2.0`
CGO_DEPS=
#CGO_LDFLAGS = `pkg-config --libs gtk+-2.0` _cgo_export.o
#CGO_DEPS=_cgo_export.o
GC:=$(GC) -I ../glib/_obj -I ../gdk/_obj -I ../gdkpixbuf/_obj

include $(GOROOT)/src/Make.pkg

%: install %.go
	$(GC) $*.go
	$(LD) -o $@ $*.$O
