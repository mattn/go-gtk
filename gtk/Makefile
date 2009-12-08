include $(GOROOT)/src/Make.$(GOARCH)

TARG     = gtk
CGOFILES = gtk.go

CGO_CFLAGS  = `pkg-config --cflags gtk+-2.0`
CGO_LDFLAGS = `pkg-config --libs gtk+-2.0`

include $(GOROOT)/src/Make.pkg

%: install %.go
	$(GC) $*.go
	$(LD) -o $@ $*.$O
