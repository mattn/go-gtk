package main

import (
	"os"
	"github.com/mattn/go-gtk/gtk"
	"github.com/mattn/go-gtk/glib"
	"github.com/mattn/go-gtk/gdkpixbuf"
)

func main() {
	gtk.Init(&os.Args)
	window := gtk.Window(gtk.GTK_WINDOW_TOPLEVEL)
	window.SetTitle("GTK Icon View")
	window.Connect("destroy", gtk.MainQuit)

	swin := gtk.ScrolledWindow(nil, nil)

	store := gtk.ListStore(gdkpixbuf.GetGdkPixbufType(), glib.G_TYPE_STRING)
	iconview := gtk.IconViewWithModel(store)
	iconview.SetPixbufColumn(0)
	iconview.SetTextColumn(1)
	swin.Add(iconview)

	gtk.GtkStockListIDs().ForEach(func(d interface{}, v interface{}) {
		id := glib.GPtrToString(d)
		var iter gtk.GtkTreeIter
		store.Append(&iter)
		store.Set(&iter,
			gtk.Image().RenderIcon(id, gtk.GTK_ICON_SIZE_SMALL_TOOLBAR, "").Pixbuf,
			id)
	})

	window.Add(swin)
	window.SetSizeRequest(500, 200)
	window.ShowAll()

	gtk.Main()
}
