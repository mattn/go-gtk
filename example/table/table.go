package main

import (
  "os";
  "gtk";
  "fmt";
)

func main() {
	gtk.Init(&os.Args);
	window := gtk.Window(gtk.GTK_WINDOW_TOPLEVEL);
	window.SetTitle("GTK Table");
	window.Connect("destroy", func() {
		gtk.MainQuit();
	}, nil);

	swin := gtk.ScrolledWindow(nil, nil);
	swin.SetPolicy(gtk.GTK_POLICY_AUTOMATIC, gtk.GTK_POLICY_AUTOMATIC);

	table := gtk.Table(5, 5, false);
	for y := uint(0); y < 5; y++ {
		for x := uint(0); x < 5; x++ {
			table.Attach(gtk.ButtonWithLabel(fmt.Sprintf("%02d:%02d", x, y)), x, x+1, y, y+1, gtk.GTK_FILL, gtk.GTK_FILL, 5, 5);
		}
	}
	swin.AddWithViewPort(table);

	window.Add(swin);
	window.SetDefaultSize(200, 200);
	window.ShowAll();

	gtk.Main();
}
