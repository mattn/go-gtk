include $(GOROOT)/src/Make.inc

GC=${O}g -I../glib/_obj -I../gdk/_obj -I../gdkpixbuf/_obj -I$(GOROOT)/pkg/$(GOOS)_$(GOARCH)

all:
	cd glib && gomake
	cd gdk && gomake
	cd gdkpixbuf && gomake
	cd gtk && gomake

install:
	cd glib && gomake install
	cd gdk && gomake install
	cd gdkpixbuf && gomake install
	cd gtk && gomake install

clean:
	cd glib && gomake clean
	cd gdk && gomake clean
	cd gdkpixbuf && gomake clean
	cd gtk && gomake clean
	cd example && gomake clean
	cd gtksourceview && gomake clean

fmt_all:
	gofmt ./gdk/gdk.go  > ./gdk/gdk.go.fmt
	gofmt ./gtk/gtk.go > ./gtk/gtk.go.fmt
	gofmt ./glib/glib.go > ./glib/glib.go.fmt
	mv ./gtk/gtk.go.fmt ./gtk/gtk.go
	mv ./gdk/gdk.go.fmt ./gdk/gdk.go
	mv ./glib/glib.go.fmt ./glib/glib.go

example: install
	cd example && gomake
