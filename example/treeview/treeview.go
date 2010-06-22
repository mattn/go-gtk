package main

import (
  "os";
  "glib";
  "gtk";
  "gdkpixbuf";
  "unsafe";
)

func main() {
	gtk.Init(&os.Args);
	window := gtk.Window(gtk.GTK_WINDOW_TOPLEVEL);
	window.SetTitle("GTK Stock Icons");
	window.Connect("destroy", func(w *gtk.GtkWidget, args []unsafe.Pointer) {
		gtk.MainQuit();
	}, nil);

	swin := gtk.ScrolledWindow(nil, nil);

	store := gtk.ListStore(gtk.TYPE_STRING, gtk.TYPE_BOOL, gdkpixbuf.GetGdkPixbufType());
	treeview := gtk.TreeView();
	swin.Add(treeview);

	treeview.SetModel(store.ToTreeModel());

	treeview.AppendColumn(gtk.TreeViewColumnWithAttributes("name", gtk.CellRendererText(), "text", 0));
	treeview.AppendColumn(gtk.TreeViewColumnWithAttributes("check", gtk.CellRendererToggle(), "active", 1));
	treeview.AppendColumn(gtk.TreeViewColumnWithAttributes("icon", gtk.CellRendererPixbuf(), "pixbuf", 2));

	window.Add(swin);
	window.SetSizeRequest(400, 200);
	window.ShowAll();

	n := 0;
	gtk.GtkStockListIDs().ForEach(func(d interface{}, v interface{}) {
		id := glib.GPtrToString(d);
		var iter gtk.GtkTreeIter;
		store.Append(&iter);
		store.Set(&iter, id, (n == 1), gtk.Image().RenderIcon(id, gtk.GTK_ICON_SIZE_SMALL_TOOLBAR, "").Pixbuf);
		n = 1 - n;
	}, 0)

	gtk.Main();
}
