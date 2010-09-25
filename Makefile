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

example: install
	cd example && gomake
