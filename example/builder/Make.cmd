include $(GOROOT)/src/Make.inc

TARG=builder
GOFILES=builder.go
CGOFILES=builder.go
GC=${O}g -Ipkg/$(GOOS)_$(GOARCH)
LD=${O}l -Lpkg/$(GOOS)_$(GOARCH)

include $(GOROOT)/src/Make.cmd

installlocal: $(TARG)
	mkdir -p pkg/$(GOOS)_$(GOARCH)
	mv $(TARG) pkg/$(GOOS)_$(GOARCH)/.
