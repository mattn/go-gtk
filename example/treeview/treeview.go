package main

import (
  "os";
  "gtk";
  "unsafe";
)

func main() {
	gtk.Init(&os.Args);
	window := gtk.Window(gtk.GTK_WINDOW_TOPLEVEL);
	window.SetTitle("TreeView!");
	window.Connect("destroy", func(w *gtk.GtkWidget, args []unsafe.Pointer) {
		gtk.MainQuit();
	}, nil);

	vbox := gtk.VBox(false, 1)

	store := gtk.ListStore(gtk.TYPE_STRING, gtk.TYPE_BOOL);
	treeview := gtk.TreeView();
	treeview.SetModel(store.ToTreeModel());

	treeview.AppendColumn(gtk.TreeViewColumnWithAttributes("text", gtk.CellRendererText(), "text", 0));
	treeview.AppendColumn(gtk.TreeViewColumnWithAttributes("toggle", gtk.CellRendererToggle(), "active", 1));

	var iter gtk.GtkTreeIter;
	store.Append(&iter);
	store.Set(&iter, "amachang", true)
	store.Append(&iter);
	store.Set(&iter, "otsune", false)
	store.Append(&iter);
	store.Set(&iter, "miyagawa", true)
	store.Append(&iter);
	store.Set(&iter, "tokuhirom", false)

	vbox.Add(treeview);

	window.Add(vbox);
	window.SetSizeRequest(200, 200);
	window.ShowAll();
	gtk.Main();
}
