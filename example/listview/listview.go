package main

import (
	"os"
	"github.com/mattn/go-gtk/glib"
	"github.com/mattn/go-gtk/gtk"
	"github.com/mattn/go-gtk/gdkpixbuf"
)

func main() {
	gtk.Init(&os.Args)
	window := gtk.Window(gtk.GTK_WINDOW_TOPLEVEL)
	window.SetTitle("GTK Stock Icons")
	window.Connect("destroy", gtk.MainQuit)

	swin := gtk.ScrolledWindow(nil, nil)

	store := gtk.ListStore(glib.G_TYPE_STRING, glib.G_TYPE_BOOL, gdkpixbuf.GetGdkPixbufType())
	treeview := gtk.TreeView()
	swin.Add(treeview)

	treeview.SetModel(store)
	treeview.AppendColumn(gtk.TreeViewColumnWithAttributes("name", gtk.CellRendererText(), "text", 0))
	treeview.AppendColumn(gtk.TreeViewColumnWithAttributes("check", gtk.CellRendererToggle(), "active", 1))
	treeview.AppendColumn(gtk.TreeViewColumnWithAttributes("icon", gtk.CellRendererPixbuf(), "pixbuf", 2))
	n := 0
	gtk.GtkStockListIDs().ForEach(func(d interface{}, v interface{}) {
		id := glib.GPtrToString(d)
		var iter gtk.GtkTreeIter
		store.Append(&iter)
		store.Set(&iter, id, (n == 1), gtk.Image().RenderIcon(id, gtk.GTK_ICON_SIZE_SMALL_TOOLBAR, "").Pixbuf)
		n = 1 - n
	})

	window.Add(swin)
	window.SetSizeRequest(400, 200)
	window.ShowAll()

	gtk.Main()
}
