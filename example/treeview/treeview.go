package main

import (
	"github.com/mattn/go-gtk/gdkpixbuf"
	"github.com/mattn/go-gtk/glib"
	"github.com/mattn/go-gtk/gtk"
	"os"
	"strconv"
)

func main() {
	gtk.Init(&os.Args)
	window := gtk.NewWindow(gtk.WINDOW_TOPLEVEL)
	window.SetTitle("GTK Folder View")
	window.Connect("destroy", gtk.MainQuit)

	swin := gtk.NewScrolledWindow(nil, nil)

	store := gtk.NewTreeStore(gdkpixbuf.GetType(), glib.G_TYPE_STRING)
	treeview := gtk.NewTreeView()
	swin.Add(treeview)

	treeview.SetModel(store.ToTreeModel())
	treeview.AppendColumn(gtk.NewTreeViewColumnWithAttributes("pixbuf", gtk.NewCellRendererPixbuf(), "pixbuf", 0))
	treeview.AppendColumn(gtk.NewTreeViewColumnWithAttributes("text", gtk.NewCellRendererText(), "text", 1))

	for n := 1; n <= 10; n++ {
		var iter1, iter2, iter3 gtk.TreeIter
		store.Append(&iter1, nil)
		store.Set(&iter1, gtk.NewImage().RenderIcon(gtk.STOCK_DIRECTORY, gtk.ICON_SIZE_SMALL_TOOLBAR, "").GPixbuf, "Folder"+strconv.Itoa(n))
		store.Append(&iter2, &iter1)
		store.Set(&iter2, gtk.NewImage().RenderIcon(gtk.STOCK_DIRECTORY, gtk.ICON_SIZE_SMALL_TOOLBAR, "").GPixbuf, "SubFolder"+strconv.Itoa(n))
		store.Append(&iter3, &iter2)
		store.Set(&iter3, gtk.NewImage().RenderIcon(gtk.STOCK_FILE, gtk.ICON_SIZE_SMALL_TOOLBAR, "").GPixbuf, "File"+strconv.Itoa(n))
	}

	treeview.Connect("row_activated", func() {
		var path *gtk.TreePath
		var column *gtk.TreeViewColumn
		treeview.GetCursor(&path, &column)
		mes := "TreePath is: " + path.String()
		dialog := gtk.NewMessageDialog(
			treeview.GetTopLevelAsWindow(),
			gtk.DIALOG_MODAL,
			gtk.MESSAGE_INFO,
			gtk.BUTTONS_OK,
			mes)
		dialog.SetTitle("TreePath")
		dialog.Response(func() {
			dialog.Destroy()
		})
		dialog.Run()
	})

	window.Add(swin)
	window.SetSizeRequest(400, 200)
	window.ShowAll()

	gtk.Main()
}
