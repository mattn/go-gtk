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

	vbox := gtk.VBox(false, 1);

	label := gtk.Label("Label");
	vbox.PackStart(label, false, true, 0);

	entry := gtk.Entry();
	entry.SetLabel("Hello world");
	vbox.Add(entry);

	buttons := gtk.HBox(false, 1);

	button := gtk.ButtonWithLabel("Button with label");
	button.Clicked(func() {
		print("button clicked: ", button.GetLabel(), "\n");
		dialog := gtk.MessageDialog(
			&gtk.GtkWindow{gtk.GetTopLevel(button)},
			gtk.GTK_DIALOG_MODAL,
			gtk.GTK_MESSAGE_INFO,
			gtk.GTK_BUTTONS_OK,
			entry.GetLabel());
		gtk.HideOnDelete(dialog);
		dialog.Run();
		gtk.Destroy(dialog);
	});
	buttons.Add(button);

	opendialog := gtk.ButtonWithLabel("Press button to see dialog bug");
	opendialog.Clicked(func () {
		println("testing a dialog...");
		dialog := gtk.MessageDialog(window,
			gtk.GTK_DIALOG_MODAL,
			gtk.GTK_MESSAGE_INFO,
			gtk.GTK_BUTTONS_OK,
			"Don't panic!");
		// dialog.Response(func () {println("You panicked!")});
		dialog.Run();
		gtk.Destroy(dialog);
	});
	buttons.Add(opendialog);

	vbox.Add(buttons);
	window.Add(vbox);

	gtk.ShowAll(window);
	gtk.Main();
}
