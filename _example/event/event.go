package main

import (
	"fmt"
	"github.com/mattn/go-gtk/gdk"
	"github.com/mattn/go-gtk/glib"
	"github.com/mattn/go-gtk/gtk"
	"os"
	"unsafe"
)

func main() {
	gtk.Init(&os.Args)
	window := gtk.NewWindow(gtk.WINDOW_TOPLEVEL)
	window.SetTitle("GTK Events")
	window.Connect("destroy", gtk.MainQuit)

	event := make(chan interface{})

	window.Connect("key-press-event", func(ctx *glib.CallbackContext) {
		arg := ctx.Args(0)
		event <- *(**gdk.EventKey)(unsafe.Pointer(&arg))
	})
	window.Connect("motion-notify-event", func(ctx *glib.CallbackContext) {
		arg := ctx.Args(0)
		event <- *(**gdk.EventMotion)(unsafe.Pointer(&arg))
	})

	go func() {
		for {
			e := <-event
			switch ev := e.(type) {
			case *gdk.EventKey:
				fmt.Println("key-press-event:", ev.Keyval)
				break
			case *gdk.EventMotion:
				fmt.Println("motion-notify-event:", int(ev.X), int(ev.Y))
				break
			}
		}
	}()

	window.SetEvents(int(gdk.POINTER_MOTION_MASK | gdk.POINTER_MOTION_HINT_MASK | gdk.BUTTON_PRESS_MASK))
	window.SetSizeRequest(400, 400)
	window.ShowAll()

	gtk.Main()
}
