include $(GOROOT)/src/Make.$(GOARCH)

TARG     = glib
CGOFILES = glib.go

CGO_CFLAGS  = `pkg-config --cflags glib-2.0`
CGO_LDFLAGS = `pkg-config --libs glib-2.0`

include $(GOROOT)/src/Make.pkg

%: install %.go
	$(GC) $*.go
	$(LD) -o $@ $*.$O
