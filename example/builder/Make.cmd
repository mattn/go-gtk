include $(GOROOT)/src/Make.inc

TARG=builder
GOFILES=builder.go
CGOFILES=builder.go
GC=${O}g -Ipkg/$(GOOS)_$(GOARCH)
LD=${O}l -Lpkg/$(GOOS)_$(GOARCH)

include $(GOROOT)/src/Make.cmd

installlocal: builder
	mkdir -p pkg/$(GOOS)_$(GOARCH)
	mv builder pkg/$(GOOS)_$(GOARCH)/.
