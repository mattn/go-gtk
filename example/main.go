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

	vbox := gtk.VBox(false, true);

	label := gtk.Label("Label");
	vbox.PackStart(label, false, true, 0);

	entry := gtk.Entry();
	entry.SetLabel("Hello world");
	gtk.Add(vbox, entry);

	button := gtk.ButtonWithLabel("Button with label");
	button.Clicked(func() {
		print("button clicked: ", button.GetLabel(), "\n");
		dialog := gtk.MessageDialog(
			&gtk.GtkWindow{gtk.GetTopLevel(button).ToGtkWidget()},
			gtk.GTK_DIALOG_MODAL,
			gtk.GTK_MESSAGE_INFO,
			gtk.GTK_BUTTONS_OK,
			entry.GetLabel());
		gtk.HideOnDelete(dialog);
		(&gtk.GtkDialog{dialog.Widget}).Run();
		gtk.Destroy(dialog);
	});
	gtk.Add(vbox, button);
	gtk.Add(window, vbox);

	gtk.ShowAll(window);
	gtk.Main();
}
