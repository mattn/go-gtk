include $(GOROOT)/src/Make.inc

TARG     = github.com/mattn/go-gtk/gtk/gtksourceview
CGOFILES = gtksourceview.go
GC:=$(GC) -I ../../glib/_obj -I ../../gdk/_obj -I ../../gdkpixbuf/_obj -I ../../pango/_obj -I ../_obj

include $(GOROOT)/src/Make.pkg
