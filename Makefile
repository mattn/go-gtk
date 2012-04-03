all:
	cd pango && go build -x .
	cd glib && go build -x .
	cd gdk && go build -x .
	cd gdkpixbuf && go build -x .
	cd gtk && go build -x .
	cd gtksourceview && go build -x .

install:
	cd pango && go install -x
	cd glib && go install -x
	cd gdk && go install -x
	cd gdkpixbuf && go install -x
	cd gtk && go install -x
	cd gtksourceview && go install -x

fmt_all:
	cd pango && go fmt .
	cd glib && go fmt .
	cd gdk && go fmt .
	cd gdkpixbuf && go fmt .
	cd gtk && go fmt .
	cd gtksourceview && go fmt .
