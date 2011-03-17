package main

import (
	"os"
	"gtk"
	"gdk"
	"unsafe"
)

type GdkEventKey struct {
  t int
  w unsafe.Pointer
  send_event int8
  time uint32
  state uint
  keyval uint
  length int
  s *uint8
  hardware_keycode uint16
  group uint8
  is_modifier uint
}

type GdkEventMotion struct {
  t int
  w unsafe.Pointer
  send_event int8
  time uint32
  x float64
  y float64
  axes *float64
  state uint
  is_hint uint16
  device uintptr;
  x_root float64
  y_root float64
}

func main() {
	gtk.Init(&os.Args)
	window := gtk.Window(gtk.GTK_WINDOW_TOPLEVEL)
	window.SetTitle("GTK Events")
	window.Connect("destroy", func() {
		gtk.MainQuit()
	})

	window.Connect("key-press-event", func(ctx *gtk.CallbackContext) {
		arg := ctx.Args(0)
		kev := *(**GdkEventKey)(unsafe.Pointer(&arg))
		println("key-press-event:", kev.keyval)
	})
	window.Connect("motion-notify-event", func(ctx *gtk.CallbackContext) {
		arg := ctx.Args(0)
		mev := *(**GdkEventMotion)(unsafe.Pointer(&arg))
		println("motion-notify-event:", int(mev.x), int(mev.y))
	})

	window.SetEvents(gdk.GDK_POINTER_MOTION_MASK | gdk.GDK_POINTER_MOTION_HINT_MASK | gdk.GDK_BUTTON_PRESS_MASK)
	window.SetSizeRequest(400, 400)
	window.ShowAll()

	gtk.Main()
}
