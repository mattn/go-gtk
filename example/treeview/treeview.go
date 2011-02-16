package main

import (
	"os"
	"gtk"
	"gdkpixbuf"
	"strconv"
)

func main() {
	gtk.Init(&os.Args)
	window := gtk.Window(gtk.GTK_WINDOW_TOPLEVEL)
	window.SetTitle("GTK Folder View")
	window.Connect("destroy", func() {
		gtk.MainQuit()
	})

	swin := gtk.ScrolledWindow(nil, nil)

	store := gtk.TreeStore(gdkpixbuf.GetGdkPixbufType(), gtk.TYPE_STRING)
	treeview := gtk.TreeView()
	swin.Add(treeview)

	treeview.SetModel(store.ToTreeModel())
	treeview.AppendColumn(gtk.TreeViewColumnWithAttributes("pixbuf", gtk.CellRendererPixbuf(), "pixbuf", 0))
	treeview.AppendColumn(gtk.TreeViewColumnWithAttributes("text", gtk.CellRendererText(), "text", 1))

	for n := 1; n <= 10; n++ {
		var iter1, iter2, iter3 gtk.GtkTreeIter
		store.Append(&iter1, nil)
		store.Set(&iter1, gtk.Image().RenderIcon(gtk.GTK_STOCK_DIRECTORY, gtk.GTK_ICON_SIZE_SMALL_TOOLBAR, "").Pixbuf, "Folder"+strconv.Itoa(n))
		store.Append(&iter2, &iter1)
		store.Set(&iter2, gtk.Image().RenderIcon(gtk.GTK_STOCK_DIRECTORY, gtk.GTK_ICON_SIZE_SMALL_TOOLBAR, "").Pixbuf, "SubFolder"+strconv.Itoa(n))
		store.Append(&iter3, &iter2)
		store.Set(&iter3, gtk.Image().RenderIcon(gtk.GTK_STOCK_FILE, gtk.GTK_ICON_SIZE_SMALL_TOOLBAR, "").Pixbuf, "File"+strconv.Itoa(n))
	}

	treeview.Connect("row_activated", func() {
		var path *gtk.GtkTreePath
		var column *gtk.GtkTreeViewColumn
		treeview.GetCursor(&path, &column)
		mes := "TreePath is: " + path.String()
		dialog := gtk.MessageDialog(
			treeview.GetTopLevelAsWindow(),
			gtk.GTK_DIALOG_MODAL,
			gtk.GTK_MESSAGE_INFO,
			gtk.GTK_BUTTONS_OK,
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
