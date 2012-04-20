all:
	cd pango && go get -x .
	cd glib && go get -x .
	cd gdk && go get -x .
	cd gdkpixbuf && go get -x .
	cd gtk && go get -x .
	cd gtksourceview && go get -x .

install:
	cd pango && go install -x
	cd glib && go install -x
	cd gdk && go install -x
	cd gdkpixbuf && go install -x
	cd gtk && go install -x
	cd gtksourceview && go install -x

fmt:
	cd pango && go fmt .
	cd glib && go fmt .
	cd gdk && go fmt .
	cd gdkpixbuf && go fmt .
	cd gtk && go fmt .
	cd gtksourceview && go fmt .
