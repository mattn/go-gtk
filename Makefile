all:
	cd pango && go build && go install github.com/mattn/go-gtk/pango
	cd glib && go build && go install github.com/mattn/go-gtk/glib
	cd gdk && go build && go install github.com/mattn/go-gtk/gdk
	cd gdkpixbuf && go build && go install github.com/mattn/go-gtk/gdkpixbuf
	cd gtk && go build && go install github.com/mattn/go-gtk/gtk
	cd gtksourceview && go build && go install github.com/mattn/go-gtk/gtksourceview

fmt:
	cd pango && go fmt .
	cd glib && go fmt .
	cd gdk && go fmt .
	cd gdkpixbuf && go fmt .
	cd gtk && go fmt .
	cd gtksourceview && go fmt .
