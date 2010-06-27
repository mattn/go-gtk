include $(GOROOT)/src/Make.$(GOARCH) 

TARG=builder
GOFILES=builder.go
CGOFILES=builder.go
CLEANFILES+=_obj/callback.a *.cgo1.go *.cgo2.c _cgo_* _obj pkg

include $(GOROOT)/src/Make.cmd
GC=${O}g -I . -I _obj -I $(CGOPKGPATH)/pkg/$(GOOS)_$(GOARCH)
LD=${O}l -L . -L _obj -L $(CGOPKGPATH)/pkg/$(GOOS)_$(GOARCH) -r ./

_obj/callback.a: callback.go
	$(MAKE) -f Make.cgo installlocal

pkg/$(GOOS)_$(GOARCH)/builder: builder
	mv builder pkg/$(GOOS)_$(GOARCH)/.

_go_.8: _obj/callback.a

all: pkg/$(GOOS)_$(GOARCH)/builder
	echo "#!/bin/sh" > builder &&\
	echo "cd `pwd`" >> builder &&\
	echo "LD_LIBRARY_PATH=./pkg/$(GOOS)_$(GOARCH) ./pkg/$(GOOS)_$(GOARCH)/builder" >> builder &&\
	chmod +x builder
