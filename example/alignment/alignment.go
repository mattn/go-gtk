package main

import (
  "os";
  "gtk";
)

func main() {
	gtk.Init(&os.Args);
	window := gtk.Window(gtk.GTK_WINDOW_TOPLEVEL);
	window.SetTitle("Alignment");
	window.Connect("destroy", func() {
		gtk.MainQuit();
	}, nil);

	align := gtk.Alignment(0.5, 0.5, 0.5, 0.5)
	window.Add(align)

	button := gtk.ButtonWithLabel("Hello World!")
	align.Add(button)

	window.ShowAll();
	window.SetSizeRequest(200, 200)

	gtk.Main();
}
