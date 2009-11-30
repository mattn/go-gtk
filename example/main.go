package main

import (
  "os";
  "gtk";
)

func main() {
	gtk.Init(&os.Args);
	window := gtk.Window(gtk.GTK_WINDOW_TOPLEVEL);
	window.SetLabel("GTK Go!");
	gtk.Connect(window, "destroy", func() {
		println("got destroy!");
		gtk.MainQuit();
	});

	vbox := gtk.VBox(0,1);

	label := gtk.Label("Label");
	vbox.PackStart(label, 0, 1, 0);

	entry := gtk.Entry();
	entry.SetLabel("Hello world");
	vbox.Add(entry);

	buttons := gtk.HBox(0,1);

	button := gtk.ButtonWithLabel("Button with label");
	button.Clicked(func() {
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
		println(button.GetLabel());
		println("entry text is: ", entry.GetLabel());
	});
	buttons.Add(button);

	opendialog := gtk.ButtonWithLabel("Press button to see dialog bug");
	opendialog.Clicked(func () {
		println("testing a dialog...");
		d := gtk.MessageDialog(window,
			gtk.GTK_DIALOG_MODAL, gtk.GTK_MESSAGE_INFO, gtk.GTK_BUTTONS_OK,
			"Don't panic!");
		// d.Response(func () {println("You panicked!")});
		d.Run();
	});
	buttons.Add(opendialog);

	vbox.Add(buttons);
	window.Add(vbox);

	gtk.ShowAll(window);
	gtk.Main();
}
