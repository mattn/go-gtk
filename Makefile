include example/example.mk

.DEFAULT_GOAL := all
.PHONY: all
all:
	cd pango && go get -x
	cd glib && go get -x
	cd gdk && go get -x
	cd gdkpixbuf && go get -x
	cd gtk && go get -x
	cd gtksourceview && go get -x
	#cd gtkgl && go get -x .
	#cd gtkspell && go get -x .

.PHONY: all
fmt:
	cd pango && go fmt .
	cd glib && go fmt .
	cd gdk && go fmt .
	cd gdkpixbuf && go fmt .
	cd gtk && go fmt .
	cd gtksourceview && go fmt .
	#cd gtkgl && go fmt .
	#cd gtkspell && go fmt .

.PHONY: clean
clean: clean-example
	@true
