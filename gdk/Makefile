include $(GOROOT)/src/Make.inc

TARG     = gdk
CGOFILES = gdk.go

CGO_CFLAGS  = `pkg-config --cflags gdk-2.0`
CGO_LDFLAGS = `pkg-config --libs gdk-2.0 gthread-2.0`

include $(GOROOT)/src/Make.pkg

%: install %.go
	$(GC) $*.go
	$(LD) -o $@ $*.$O
