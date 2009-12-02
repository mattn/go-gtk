package main

import (
  "os";
  "gtk";
  "unsafe";
  "path";
)

func main() {
	gtk.Init(&os.Args);
	window := gtk.Window(gtk.GTK_WINDOW_TOPLEVEL);
	window.SetTitle("GTK Go!");
	window.Connect("destroy", func(w *gtk.GtkWidget, args []unsafe.Pointer) {
		println("got destroy!");
		gtk.MainQuit();
	}, nil);

	//--------------------------------------------------------
	// GtkVBox
	//--------------------------------------------------------
	vbox := gtk.VBox(false, 1);

	//--------------------------------------------------------
	// GtkImage
	//--------------------------------------------------------
	dir, _ := path.Split(os.Args[0]);
	imagefile := path.Join(dir, "../data/go-gtk-logo.png"); 

	label := gtk.Label("Go Binding for GTK");
	vbox.PackStart(label, false, true, 0);

	entry := gtk.Entry();
	entry.SetText("Hello world");
	vbox.Add(entry);

	image := gtk.ImageFromFile(imagefile);
	vbox.Add(image);

	//--------------------------------------------------------
	// GtkHBox
	//--------------------------------------------------------
	buttons := gtk.HBox(false, 1);

	//--------------------------------------------------------
	// GtkButton
	//--------------------------------------------------------
	button := gtk.ButtonWithLabel("Button with label");
	button.Clicked(func(w *gtk.GtkWidget, args []unsafe.Pointer) {
		print("button clicked: ", button.GetLabel(), "\n");
		dialog := gtk.MessageDialog(
			button.GetTopLevelAsWindow(),
			gtk.GTK_DIALOG_MODAL,
			gtk.GTK_MESSAGE_INFO,
			gtk.GTK_BUTTONS_OK,
			entry.GetText()
		);
		dialog.Response(func(w *gtk.GtkWidget, args []unsafe.Pointer) {
			println("Dialog OK!")
		}, nil);
		dialog.Run();
		dialog.Destroy();
	}, nil);
	buttons.Add(button);

	//--------------------------------------------------------
	// GtkFontButton
	//--------------------------------------------------------
	fontbutton := gtk.FontButton();
	fontbutton.Connect("font-set", func(w *gtk.GtkWidget, args []unsafe.Pointer) {
		print("title: ", fontbutton.GetTitle(), "\n");
		print("fontname: ", fontbutton.GetFontName(), "\n");
		print("use_size: ", fontbutton.GetUseSize(), "\n");
		print("show_size: ", fontbutton.GetShowSize(), "\n");
	}, nil);
	buttons.Add(fontbutton);
	vbox.Add(buttons);

	buttons = gtk.HBox(false, 1);

	//--------------------------------------------------------
	// GtkToggleButton
	//--------------------------------------------------------
	togglebutton := gtk.ToggleButtonWithLabel("ToggleButton with label");
	togglebutton.Connect("toggled", func(w *gtk.GtkWidget, args []unsafe.Pointer) {
		if togglebutton.GetActive() {
			togglebutton.SetLabel("ToggleButton ON!");
		} else {
			togglebutton.SetLabel("ToggleButton OFF!");
		}
	}, nil);
	buttons.Add(togglebutton);

	//--------------------------------------------------------
	// GtkCheckButton
	//--------------------------------------------------------
	checkbutton := gtk.CheckButtonWithLabel("CheckButton with label");
	checkbutton.Connect("toggled", func(w *gtk.GtkWidget, args []unsafe.Pointer) {
		if checkbutton.GetActive() {
			checkbutton.SetLabel("CheckButton CHECKED!");
		} else {
			checkbutton.SetLabel("CheckButton UNCHECKED!");
		}
	}, nil);
	buttons.Add(checkbutton);
	vbox.Add(buttons);

	combos := gtk.HBox(false, 1);

	comboboxentry := gtk.ComboBoxEntryNewText();
	comboboxentry.Connect("changed", func(w *gtk.GtkWidget, args []unsafe.Pointer) {
		//comboboxentry...
	}, nil);
	combos.Add(comboboxentry);
	vbox.Add(combos);

	window.Add(vbox);

	window.ShowAll();
	gtk.Main();
}
