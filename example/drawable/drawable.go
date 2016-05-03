package main

import (
	"github.com/mattn/go-gtk/gdk"
	"github.com/mattn/go-gtk/glib"
	"github.com/mattn/go-gtk/gtk"
	"os"
	"unsafe"
)

type point struct {
	x int
	y int
}

func main() {
	gtk.Init(&os.Args)
	window := gtk.NewWindow(gtk.WINDOW_TOPLEVEL)
	window.SetTitle("GTK DrawingArea")
	window.Connect("destroy", gtk.MainQuit)

	vbox := gtk.NewVBox(true, 0)
	vbox.SetBorderWidth(5)
	drawingarea := gtk.NewDrawingArea()

	var p1, p2 point
	var gdkwin *gdk.Window
	var pixmap *gdk.Pixmap
	var gc *gdk.GC
	p1.x = -1
	p1.y = -1
	colors := []string{
		"black",
		"gray",
		"blue",
		"purple",
		"red",
		"orange",
		"yellow",
		"green",
		"darkgreen",
	}

	drawingarea.Connect("configure-event", func() {
		if pixmap != nil {
			pixmap.Unref()
		}
		allocation := drawingarea.GetAllocation()
		pixmap = gdk.NewPixmap(drawingarea.GetWindow().GetDrawable(), allocation.Width, allocation.Height, 24)
		gc = gdk.NewGC(pixmap.GetDrawable())
		gc.SetRgbFgColor(gdk.NewColor("white"))
		pixmap.GetDrawable().DrawRectangle(gc, true, 0, 0, -1, -1)
		gc.SetRgbFgColor(gdk.NewColor(colors[0]))
		gc.SetRgbBgColor(gdk.NewColor("white"))
	})

	drawingarea.Connect("motion-notify-event", func(ctx *glib.CallbackContext) {
		arg := ctx.Args(0)
		mev := *(**gdk.EventMotion)(unsafe.Pointer(&arg))
		var mt gdk.ModifierType
		if mev.IsHint != 0 {
			gdkwin.GetPointer(&p2.x, &p2.y, &mt)
		} else {
			p2.x, p2.y = int(mev.X), int(mev.Y)
		}
		if p1.x != -1 && p2.x != -1 && (gdk.EventMask(mt)&gdk.BUTTON_PRESS_MASK) != 0 {
			pixmap.GetDrawable().DrawLine(gc, p1.x, p1.y, p2.x, p2.y)
			gdkwin.Invalidate(nil, false)
		}
		colors = append(colors[1:], colors[0])
		gc.SetRgbFgColor(gdk.NewColor(colors[0]))
		p1 = p2
	})

	drawingarea.Connect("expose-event", func() {
		if pixmap == nil {
			return
		}
		gdkwin.GetDrawable().DrawDrawable(gc, pixmap.GetDrawable(), 0, 0, 0, 0, -1, -1)
	})

	drawingarea.SetEvents(int(gdk.POINTER_MOTION_MASK | gdk.POINTER_MOTION_HINT_MASK | gdk.BUTTON_PRESS_MASK))
	vbox.Add(drawingarea)

	window.Add(vbox)
	window.SetSizeRequest(400, 400)
	window.ShowAll()

	gdkwin = drawingarea.GetWindow()

	gtk.Main()
}
