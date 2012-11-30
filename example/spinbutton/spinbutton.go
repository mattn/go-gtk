package main

import (
	"github.com/mattn/go-gtk/glib"
	"github.com/mattn/go-gtk/gtk"
)

func main() {
	gtk.Init(nil)
	window := gtk.Window(gtk.GTK_WINDOW_TOPLEVEL)
	window.SetPosition(gtk.GTK_WIN_POS_CENTER)
	window.SetTitle("GTK Go!")
	window.Connect("destroy", func(ctx *glib.CallbackContext) {
		println("got destroy!", ctx.Data().(string))
		gtk.MainQuit()
	}, "foo")

	//--------------------------------------------------------
	// GtkHBox
	//--------------------------------------------------------
	hbox := gtk.HBox(false, 1)

	//--------------------------------------------------------
	// GtkSpinButton
	//--------------------------------------------------------
	spinbutton1 := gtk.SpinButtonWithRange(1.0, 10.0, 1.0)
	hbox.Add(spinbutton1)

	adjustment := gtk.Adjustment(2.0, 1.0, 8.0, 2.0, 0.0, 0.0)

	spinbutton2 := gtk.SpinButton(adjustment, 1.0, 1)
	hbox.Add(spinbutton2)

	//--------------------------------------------------------
	// Event
	//--------------------------------------------------------
	window.Add(hbox)
	window.SetSizeRequest(600, 600)
	window.ShowAll()
	gtk.Main()
}
