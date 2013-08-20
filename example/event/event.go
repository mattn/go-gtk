package main

import (
	"fmt"
	"os"
	"github.com/mattn/go-gtk/glib"	
	"github.com/mattn/go-gtk/gdk"
	"github.com/mattn/go-gtk/gtk"	
	"unsafe"
)

func main() {
	gtk.Init(&os.Args)
	window := gtk.NewWindow(gtk.WINDOW_TOPLEVEL)
	window.SetTitle("GTK Events")
	window.Connect("destroy", gtk.MainQuit)

	window.Connect("key-press-event", func(ctx *glib.CallbackContext) {
		arg := ctx.Args(0)
		kev := *(**gdk.EventKey)(unsafe.Pointer(&arg))
		fmt.Println("key-press-event:", kev.Keyval)
	})
	window.Connect("motion-notify-event", func(ctx *glib.CallbackContext) {
		arg := ctx.Args(0)
		mev := *(**gdk.EventMotion)(unsafe.Pointer(&arg))
		fmt.Println("motion-notify-event:", int(mev.X), int(mev.Y))
	})

	window.SetEvents(int(gdk.POINTER_MOTION_MASK | gdk.POINTER_MOTION_HINT_MASK | gdk.BUTTON_PRESS_MASK))
	window.SetSizeRequest(400, 400)
	window.ShowAll()

	gtk.Main()
}
