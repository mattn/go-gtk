package main

import (
	"github.com/mattn/go-gtk/gtk"
	"os"
)

func createArrowButton(at gtk.ArrowType, st gtk.ShadowType) *gtk.Button {
	b := gtk.NewButton()
	a := gtk.NewArrow(at, st)

	b.Add(a)

	b.Show()
	a.Show()

	return b
}

func main() {
	gtk.Init(&os.Args)

	window := gtk.NewWindow(gtk.WINDOW_TOPLEVEL)
	window.SetTitle("Arrow Buttons")
	window.Connect("destroy", gtk.MainQuit)

	box := gtk.NewHBox(false, 0)
	box.Show()
	window.Add(box)

	up := createArrowButton(gtk.ARROW_UP, gtk.SHADOW_IN)
	down := createArrowButton(gtk.ARROW_DOWN, gtk.SHADOW_OUT)
	left := createArrowButton(gtk.ARROW_LEFT, gtk.SHADOW_ETCHED_IN)
	right := createArrowButton(gtk.ARROW_RIGHT, gtk.SHADOW_ETCHED_OUT)

	box.PackStart(up, false, false, 3)
	box.PackStart(down, false, false, 3)
	box.PackStart(left, false, false, 3)
	box.PackStart(right, false, false, 3)

	up.Clicked(func() { println("↑") })
	down.Clicked(func() { println("↓") })
	left.Clicked(func() { println("←") })
	right.Clicked(func() { println("→") })

	window.Show()
	gtk.Main()
}
