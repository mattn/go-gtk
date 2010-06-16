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
	// GtkMenuBar
	//--------------------------------------------------------
	menubar := gtk.MenuBar();
	vbox.PackStart(menubar, false, false, 0);

	//--------------------------------------------------------
	// GtkMenuItem
	//--------------------------------------------------------
	filemenu := gtk.MenuItemWithMnemonic("_File");
	menubar.Append(filemenu);
	filesubmenu := gtk.Menu();
	filemenu.SetSubmenu(filesubmenu);

		exitmenuitem := gtk.MenuItemWithMnemonic("E_xit");
		exitmenuitem.Connect("activate", func(w *gtk.GtkWidget, args []unsafe.Pointer) {
			gtk.MainQuit();
		}, nil);
		filesubmenu.Append(exitmenuitem);

	//--------------------------------------------------------
	// GtkFrame
	//--------------------------------------------------------
	frame := gtk.Frame("Demo");
	framebox := gtk.VBox(false, 1);
	frame.Add(framebox);
	vbox.Add(frame);

	//--------------------------------------------------------
	// GtkImage
	//--------------------------------------------------------
	dir, _ := path.Split(os.Args[0]);
	imagefile := path.Join(dir, "../../data/go-gtk-logo.png"); 

	label := gtk.Label("Go Binding for GTK");
	framebox.PackStart(label, false, true, 0);

	entry := gtk.Entry();
	entry.SetText("Hello world");
	framebox.Add(entry);

	image := gtk.ImageFromFile(imagefile);
	framebox.Add(image);

	//--------------------------------------------------------
	// GtkScale
	//--------------------------------------------------------
	scale := gtk.HScaleWithRange(0, 100, 1);
	scale.Connect("value-changed", func(w *gtk.GtkWidget, args []unsafe.Pointer) {
		print("scale: ", scale.GetValue(), "\n");
	}, nil);
	framebox.Add(scale);

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
			messagedialog := gtk.MessageDialog(
				button.GetTopLevelAsWindow(),
				gtk.GTK_DIALOG_MODAL,
				gtk.GTK_MESSAGE_INFO,
				gtk.GTK_BUTTONS_OK,
				entry.GetText());
			messagedialog.Response(func(w *gtk.GtkWidget, args []unsafe.Pointer) {
				println("Dialog OK!")

				//--------------------------------------------------------
				// GtkFileChooserDialog
				//--------------------------------------------------------
				filechooserdialog := gtk.FileChooserDialog(
					"Choose File...",
					button.GetTopLevelAsWindow(),
					gtk.GTK_FILE_CHOOSER_ACTION_OPEN,
					gtk.GTK_STOCK_OK);
				filechooserdialog.Response(func(w *gtk.GtkWidget, args []unsafe.Pointer) {
					println(filechooserdialog.GetFilename());
					filechooserdialog.Destroy();
				}, nil);
				filechooserdialog.Run();
			}, nil);
			messagedialog.Run();
			messagedialog.Destroy();
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
		framebox.PackStart(buttons, false, false, 0);

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

		//--------------------------------------------------------
		// GtkRadioButton
		//--------------------------------------------------------
		buttonbox := gtk.VBox(false, 1)
		radiofirst := gtk.RadioButtonWithLabel(nil, "Radio1");
		buttonbox.Add(radiofirst);
		buttonbox.Add(gtk.RadioButtonWithLabel(radiofirst.GetGroup(), "Radio2"));
		buttonbox.Add(gtk.RadioButtonWithLabel(radiofirst.GetGroup(), "Radio3"));
		buttons.Add(buttonbox);
		//radiobutton.SetMode(false);
		radiofirst.SetActive(true);

		framebox.PackStart(buttons, false, false, 0);

	//--------------------------------------------------------
	// GtkComboBoxEntry
	//--------------------------------------------------------
	combos := gtk.HBox(false, 1);
	comboboxentry := gtk.ComboBoxEntryNewText();
	comboboxentry.AppendText("Monkey");
	comboboxentry.AppendText("Tiger");
	comboboxentry.AppendText("Elephant");
	comboboxentry.Connect("changed", func(w *gtk.GtkWidget, args []unsafe.Pointer) {
		print("value: ", comboboxentry.GetActiveText(), "\n");
	}, nil);
	combos.Add(comboboxentry);

	//--------------------------------------------------------
	// GtkComboBox
	//--------------------------------------------------------
	combobox := gtk.ComboBoxNewText();
	combobox.AppendText("Peach");
	combobox.AppendText("Banana");
	combobox.AppendText("Apple");
	combobox.SetActive(1);
	combobox .Connect("changed", func(w *gtk.GtkWidget, args []unsafe.Pointer) {
		print("value: ", combobox.GetActiveText(), "\n");
	}, nil);
	combos.Add(combobox);

	framebox.PackStart(combos, false, false, 0);

	//--------------------------------------------------------
	// GtkTextView
	//--------------------------------------------------------
	swin := gtk.ScrolledWindow(nil, nil);
	swin.SetPolicy(gtk.GTK_POLICY_AUTOMATIC, gtk.GTK_POLICY_AUTOMATIC);
	swin.SetShadowType(gtk.GTK_SHADOW_IN);
	textview := gtk.TextView();
	var start, end gtk.GtkTextIter;
	buffer := textview.GetBuffer();
	buffer.GetStartIter(&start);
	buffer.Insert(&start, "Hello ");
	buffer.GetEndIter(&end);
	buffer.Insert(&end, "World!");
	tag := buffer.CreateTag("bold", map[string] string {
		"background": "#FF0000", "weight": "700" });
	buffer.GetStartIter(&start);
	buffer.GetEndIter(&end);
	buffer.ApplyTag(tag, &start, &end);
	swin.Add(textview);
	framebox.Add(swin);

	//--------------------------------------------------------
	// GtkStatusbar
	//--------------------------------------------------------
	statusbar := gtk.Statusbar();
	context_id := statusbar.GetContextId("go-gtk");
	statusbar.Push(context_id, "GTK binding for Go!");

	framebox.PackStart(statusbar, false, false, 0);

	//--------------------------------------------------------
	// Event
	//--------------------------------------------------------
	window.Add(vbox);
	window.SetSizeRequest(600, 600);
	window.ShowAll();
	gtk.Main();
}
