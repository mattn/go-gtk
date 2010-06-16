GC=${O}g -I../glib/_obj -I../gdk/_obj -I../gdkpixbuf/_obj -I$(GOROOT)/pkg/$(GOOS)_$(GOARCH)

all:
	cd glib && make
	cd gdk && make
	cd gdkpixbuf && make
	cd gtk && make

install:
	cd glib && make install
	cd gdk && make install
	cd gdkpixbuf && make install
	cd gtk && make install

clean:
	cd glib && make clean
	cd gdk && make clean
	cd gdkpixbuf && make clean
	cd gtk && make clean
	cd example && make clean

example: install
	cd example && make
