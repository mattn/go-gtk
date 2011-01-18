include $(GOROOT)/src/Make.inc

all:
	$(MAKE) -f Make.cgo
	$(MAKE) -f Make.cgo installlocal
	$(MAKE) -f Make.cmd
	$(MAKE) -f Make.cmd installlocal
	echo "#!/bin/sh" > builder &&\
	echo "cd `pwd`" >> builder &&\
	echo "LD_PRELOAD=./pkg/$(GOOS)_$(GOARCH)/callback.so ./pkg/$(GOOS)_$(GOARCH)/builder" >> builder &&\
	chmod +x builder

clean:
	$(MAKE) -f Make.cgo clean
	$(MAKE) -f Make.cmd clean
