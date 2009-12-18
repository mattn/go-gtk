include $(GOROOT)/src/Make.$(GOARCH)

TARG     = gdk-pixbuf
CGOFILES = gdk-pixbuf.go

CGO_CFLAGS  = `pkg-config --cflags gdk-pixbuf-2.0`
CGO_LDFLAGS = `pkg-config --libs gdk-pixbuf-2.0`

include $(GOROOT)/src/Make.pkg

%: install %.go
	$(GC) $*.go
	$(LD) -o $@ $*.$O
