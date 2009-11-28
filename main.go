package main

import (
  "os";
  "gtk";
  "unsafe";
)

func main() {
	gtk.Init(&os.Args);
	window := gtk.Window(gtk.GTK_WINDOW_TOPLEVEL);
	(&gtk.GtkWindow{window.Widget}).SetTitle("GTK Go!");
	window.Connect("destroy", func(widget *gtk.GtkWidget, data unsafe.Pointer){
		gtk.MainQuit();
	}, nil);

	vbox := gtk.VBox(0, 1);

	label := gtk.Label("ハローワールド");
	(&gtk.GtkBox{vbox.Widget}).PackStart(label, 0, 1, 0);

	entry := gtk.Entry();
	(&gtk.GtkEntry{entry.Widget}).SetText("入力エリア！");
	vbox.Add(entry);

	button := gtk.ButtonWithLabel("こんにちわ！こんにちわ！");
	button.Connect("clicked", func(widget *gtk.GtkWidget, data unsafe.Pointer){
		println("button clicked");
		/*
		dialog := gtk.MessageDialog(
			&gtk.GtkWindow{widget.GetTopLevel().Widget},
			gtk.GTK_DIALOG_MODAL,
			gtk.GTK_MESSAGE_INFO,
			gtk.GTK_BUTTONS_OK,
			(&gtk.GtkEntry{entry.Widget}).GetText());
		(&gtk.GtkDialog{dialog.Widget}).Run();
		*/
		println((&gtk.GtkEntry{entry.Widget}).GetText());
	}, nil);
	vbox.Add(button);

	window.Add(vbox);

	window.ShowAll();
	gtk.Main();
}
