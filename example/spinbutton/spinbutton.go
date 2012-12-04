package main

import (
	"github.com/mattn/go-gtk/glib"
	"github.com/mattn/go-gtk/gtk"
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
	fixed := gtk.Fixed()

	//--------------------------------------------------------
	// GtkSpinButton
	//--------------------------------------------------------
	spinbutton1 := gtk.SpinButtonWithRange(1.0, 10.0, 1.0)
	spinbutton1.SetDigits(3)
	spinbutton1.Spin(gtk.GTK_SPIN_STEP_FORWARD, 7.0)
	fixed.Put(spinbutton1, 40, 50)

	spinbutton1.OnValueChanged(func() {
		val := spinbutton1.GetValueAsInt()
		fval := spinbutton1.GetValue()
		println("SpinButton changed, new value: " + strconv.Itoa(val) + " | " + strconv.FormatFloat(fval, 'f', 2, 64))
		min, max := spinbutton1.GetRange()
		println("Range: " + strconv.FormatFloat(min, 'f', 2, 64) + " " + strconv.FormatFloat(max, 'f', 2, 64))
		println("Digits: " + strconv.Itoa(int(spinbutton1.GetDigits())))
	})

	adjustment := gtk.Adjustment(2.0, 1.0, 8.0, 2.0, 0.0, 0.0)
	spinbutton2 := gtk.SpinButton(adjustment, 1.0, 1)
	spinbutton2.SetRange(0.0, 20.0)
	spinbutton2.SetValue(18.0)
	spinbutton2.SetIncrements(2.0, 4.0)
	fixed.Put(spinbutton2, 150, 50)

	//--------------------------------------------------------
	// Event
	//--------------------------------------------------------
	window.Add(fixed)
	window.SetSizeRequest(600, 600)
	window.ShowAll()
	gtk.Main()
}
