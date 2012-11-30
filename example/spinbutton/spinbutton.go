package main

import (
	"github.com/mattn/go-gtk/glib"
	"go-gtk/gtk"
	"strconv"
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
	box := gtk.VBox(false, 1)

	//--------------------------------------------------------
	// GtkSpinButton
	//--------------------------------------------------------
	spinbutton1 := gtk.SpinButtonWithRange(1.0, 10.0, 1.0)
	box.Add(spinbutton1)

	spinbutton1.ValueChanged(func() {
		val := spinbutton1.GetValueAsInt()
		fval := spinbutton1.GetValue()
		println("SpinButton changed, new value: " + strconv.Itoa(val) + " | " + strconv.FormatFloat(fval, 'f', 2, 64))
	})

	adjustment := gtk.Adjustment(2.0, 1.0, 8.0, 2.0, 0.0, 0.0)
	spinbutton2 := gtk.SpinButton(adjustment, 1.0, 1)
	spinbutton2.SetRange(0.0, 20.0)
	spinbutton2.SetValue(18.0)
	spinbutton2.SetIncrements(2.0, 4.0)
	box.Add(spinbutton2)

	//--------------------------------------------------------
	// Event
	//--------------------------------------------------------
	window.Add(box)
	window.SetSizeRequest(600, 600)
	window.ShowAll()
	gtk.Main()
}
