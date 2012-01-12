include $(GOROOT)/src/Make.inc

TARG     = github.com/mattn/go-gtk/gtksourceview
CGOFILES = gtksourceview.go
GC:=$(GC) -I ../gtk/_obj -I ../glib/_obj -I ../gdk/_obj -I ../gdkpixbuf/_obj -I ../pango/_obj

include $(GOROOT)/src/Make.pkg
