package main

import (
	"os"
	"gtk"
	"gdkpixbuf"
	"path"
)

func main() {
	gtk.Init(nil)
	window := gtk.Window(gtk.GTK_WINDOW_TOPLEVEL)
	window.SetTitle("GTK Go!")
	window.Connect("destroy", func(w *gtk.GtkWidget, user_data string) {
		println("got destroy!", user_data)
		gtk.MainQuit()
	},
		"foo")

	//--------------------------------------------------------
	// GtkVBox
	//--------------------------------------------------------
	vbox := gtk.VBox(false, 1)

	//--------------------------------------------------------
	// GtkMenuBar
	//--------------------------------------------------------
	menubar := gtk.MenuBar()
	vbox.PackStart(menubar, false, false, 0)

	//--------------------------------------------------------
	// GtkVPaned
	//--------------------------------------------------------
	vpaned := gtk.VPaned()
	vbox.Add(vpaned)

	//--------------------------------------------------------
	// GtkMenuItem
	//--------------------------------------------------------
	cascademenu := gtk.MenuItemWithMnemonic("_File")
	menubar.Append(cascademenu)
	submenu := gtk.Menu()
	cascademenu.SetSubmenu(submenu)

	menuitem := gtk.MenuItemWithMnemonic("E_xit")
	menuitem.Connect("activate", func() {
		gtk.MainQuit()
	},
		nil)
	submenu.Append(menuitem)

	cascademenu = gtk.MenuItemWithMnemonic("_View")
	menubar.Append(cascademenu)
	submenu = gtk.Menu()
	cascademenu.SetSubmenu(submenu)

	checkmenuitem := gtk.CheckMenuItemWithMnemonic("_Disable")
	checkmenuitem.Connect("activate", func() {
		vpaned.SetSensitive(!checkmenuitem.GetActive())
	},
		nil)
	submenu.Append(checkmenuitem)

	cascademenu = gtk.MenuItemWithMnemonic("_Help")
	menubar.Append(cascademenu)
	submenu = gtk.Menu()
	cascademenu.SetSubmenu(submenu)

	menuitem = gtk.MenuItemWithMnemonic("_About")
	menuitem.Connect("activate", func() {
		dialog := gtk.AboutDialog()
		dialog.SetName("Go-Gtk Demo!")
		dialog.SetProgramName("demo")
		dialog.SetAuthors([]string{
			"Yasuhiro Matsumoto <mattn.jp@gmail.com>",
			"David Roundy <roundyd@physics.oregonstate.edu>"})
		dir, _ := path.Split(os.Args[0])
		imagefile := path.Join(dir, "../../data/mattn-logo.png")
		pixbuf, _ := gdkpixbuf.PixbufFromFile(imagefile)
		dialog.SetLogo(pixbuf)
		dialog.SetLicense("The library is available under the same terms and conditions as the Go, the BSD style license, and the LGPL (Lesser GNU Public License). The idea is that if you can use Go (and Gtk) in a project, you should also be able to use go-gtk.")
		dialog.SetWrapLicense(true)
		dialog.Run()
		dialog.Destroy()
	},
		nil)
	submenu.Append(menuitem)

	//--------------------------------------------------------
	// GtkFrame
	//--------------------------------------------------------
	frame1 := gtk.Frame("Demo")
	framebox1 := gtk.VBox(false, 1)
	frame1.Add(framebox1)

	frame2 := gtk.Frame("Demo")
	framebox2 := gtk.VBox(false, 1)
	frame2.Add(framebox2)

	vpaned.Add1(frame1)
	vpaned.Add2(frame2)

	//--------------------------------------------------------
	// GtkImage
	//--------------------------------------------------------
	dir, _ := path.Split(os.Args[0])
	imagefile := path.Join(dir, "../../data/go-gtk-logo.png")

	label := gtk.Label("Go Binding for GTK")
	label.ModifyFontEasy("DejaVu Serif 15")
	framebox1.PackStart(label, false, true, 0)

	entry := gtk.Entry()
	entry.SetText("Hello world")
	framebox1.Add(entry)

	image := gtk.ImageFromFile(imagefile)
	framebox1.Add(image)

	//--------------------------------------------------------
	// GtkScale
	//--------------------------------------------------------
	scale := gtk.HScaleWithRange(0, 100, 1)
	scale.Connect("value-changed", func() {
		print("scale: ", int(scale.GetValue()), "\n")
	},
		nil)
	framebox2.Add(scale)

	//--------------------------------------------------------
	// GtkHBox
	//--------------------------------------------------------
	buttons := gtk.HBox(false, 1)

	//--------------------------------------------------------
	// GtkButton
	//--------------------------------------------------------
	button := gtk.ButtonWithLabel("Button with label")
	button.Clicked(func() {
		print("button clicked: ", button.GetLabel(), "\n")
		messagedialog := gtk.MessageDialog(
			button.GetTopLevelAsWindow(),
			gtk.GTK_DIALOG_MODAL,
			gtk.GTK_MESSAGE_INFO,
			gtk.GTK_BUTTONS_OK,
			entry.GetText())
		messagedialog.Response(func() {
			println("Dialog OK!")

			//--------------------------------------------------------
			// GtkFileChooserDialog
			//--------------------------------------------------------
			filechooserdialog := gtk.FileChooserDialog(
				"Choose File...",
				button.GetTopLevelAsWindow(),
				gtk.GTK_FILE_CHOOSER_ACTION_OPEN,
				gtk.GTK_STOCK_OK,
				gtk.GTK_RESPONSE_ACCEPT)
			filechooserdialog.Response(func() {
				println(filechooserdialog.GetFilename())
				filechooserdialog.Destroy()
			},
				nil)
			filechooserdialog.Run()
		},
			nil)
		messagedialog.Run()
		messagedialog.Destroy()
	},
		nil)
	buttons.Add(button)

	//--------------------------------------------------------
	// GtkFontButton
	//--------------------------------------------------------
	fontbutton := gtk.FontButton()
	fontbutton.Connect("font-set", func() {
		print("title: ", fontbutton.GetTitle(), "\n")
		print("fontname: ", fontbutton.GetFontName(), "\n")
		print("use_size: ", fontbutton.GetUseSize(), "\n")
		print("show_size: ", fontbutton.GetShowSize(), "\n")
	},
		nil)
	buttons.Add(fontbutton)
	framebox2.PackStart(buttons, false, false, 0)

	buttons = gtk.HBox(false, 1)

	//--------------------------------------------------------
	// GtkToggleButton
	//--------------------------------------------------------
	togglebutton := gtk.ToggleButtonWithLabel("ToggleButton with label")
	togglebutton.Connect("toggled", func() {
		if togglebutton.GetActive() {
			togglebutton.SetLabel("ToggleButton ON!")
		} else {
			togglebutton.SetLabel("ToggleButton OFF!")
		}
	},
		nil)
	buttons.Add(togglebutton)

	//--------------------------------------------------------
	// GtkCheckButton
	//--------------------------------------------------------
	checkbutton := gtk.CheckButtonWithLabel("CheckButton with label")
	checkbutton.Connect("toggled", func() {
		if checkbutton.GetActive() {
			checkbutton.SetLabel("CheckButton CHECKED!")
		} else {
			checkbutton.SetLabel("CheckButton UNCHECKED!")
		}
	},
		nil)
	buttons.Add(checkbutton)

	//--------------------------------------------------------
	// GtkRadioButton
	//--------------------------------------------------------
	buttonbox := gtk.VBox(false, 1)
	radiofirst := gtk.RadioButtonWithLabel(nil, "Radio1")
	buttonbox.Add(radiofirst)
	buttonbox.Add(gtk.RadioButtonWithLabel(radiofirst.GetGroup(), "Radio2"))
	buttonbox.Add(gtk.RadioButtonWithLabel(radiofirst.GetGroup(), "Radio3"))
	buttons.Add(buttonbox)
	//radiobutton.SetMode(false);
	radiofirst.SetActive(true)

	framebox2.PackStart(buttons, false, false, 0)

	//--------------------------------------------------------
	// GtkComboBoxEntry
	//--------------------------------------------------------
	combos := gtk.HBox(false, 1)
	comboboxentry := gtk.ComboBoxEntryNewText()
	comboboxentry.AppendText("Monkey")
	comboboxentry.AppendText("Tiger")
	comboboxentry.AppendText("Elephant")
	comboboxentry.Connect("changed", func() {
		print("value: ", comboboxentry.GetActiveText(), "\n")
	},
		nil)
	combos.Add(comboboxentry)

	//--------------------------------------------------------
	// GtkComboBox
	//--------------------------------------------------------
	combobox := gtk.ComboBoxNewText()
	combobox.AppendText("Peach")
	combobox.AppendText("Banana")
	combobox.AppendText("Apple")
	combobox.SetActive(1)
	combobox.Connect("changed", func() {
		print("value: ", combobox.GetActiveText(), "\n")
	},
		nil)
	combos.Add(combobox)

	framebox2.PackStart(combos, false, false, 0)

	//--------------------------------------------------------
	// GtkTextView
	//--------------------------------------------------------
	swin := gtk.ScrolledWindow(nil, nil)
	swin.SetPolicy(gtk.GTK_POLICY_AUTOMATIC, gtk.GTK_POLICY_AUTOMATIC)
	swin.SetShadowType(gtk.GTK_SHADOW_IN)
	textview := gtk.TextView()
	var start, end gtk.GtkTextIter
	buffer := textview.GetBuffer()
	buffer.GetStartIter(&start)
	buffer.Insert(&start, "Hello ")
	buffer.GetEndIter(&end)
	buffer.Insert(&end, "World!")
	tag := buffer.CreateTag("bold", map[string]string{
		"background": "#FF0000", "weight": "700"})
	buffer.GetStartIter(&start)
	buffer.GetEndIter(&end)
	buffer.ApplyTag(tag, &start, &end)
	swin.Add(textview)
	framebox2.Add(swin)

	buffer.Connect("changed", func() {
		println("changed")
	},
		nil)

	//--------------------------------------------------------
	// GtkStatusbar
	//--------------------------------------------------------
	statusbar := gtk.Statusbar()
	context_id := statusbar.GetContextId("go-gtk")
	statusbar.Push(context_id, "GTK binding for Go!")

	framebox2.PackStart(statusbar, false, false, 0)

	//--------------------------------------------------------
	// Event
	//--------------------------------------------------------
	window.Add(vbox)
	window.SetSizeRequest(600, 600)
	window.ShowAll()
	gtk.Main()
}
