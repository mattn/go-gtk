package main

import (
  "os";
  "gtk";
  "unsafe";
)

func main() {
	gtk.Init(&os.Args);
	w := gtk.Window(gtk.GTK_WINDOW_TOPLEVEL);
	b := gtk.Button("こんにちわ！こんにちわ！");
	b.Connect("clicked", func(w *gtk.GtkWidget, d unsafe.Pointer){
		print("clicked\n");
	});
	w.Add(b);
	w.ShowAll();
	gtk.Main();
}
