package main

import (
	"github.com/mattn/go-gtk/gtk"
	"os"
)

func main() {
	gtk.Init(&os.Args)
	window := gtk.NewWindow(gtk.WINDOW_TOPLEVEL)
	window.SetTitle("Alignment")
	window.Connect("destroy", gtk.MainQuit)

	notebook := gtk.NewNotebook()
	window.Add(notebook)

	align := gtk.NewAlignment(0.5, 0.5, 0.5, 0.5)
	notebook.AppendPage(align, gtk.NewLabel("Alignment"))

	button := gtk.NewButtonWithLabel("Hello World!")
	align.Add(button)

	fixed := gtk.NewFixed()
	notebook.AppendPage(fixed, gtk.NewLabel("Fixed"))

	button2 := gtk.NewButtonWithLabel("Pulse")
	fixed.Put(button2, 30, 30)

	progress := gtk.NewProgressBar()
	fixed.Put(progress, 30, 70)

	button.Connect("clicked", func() {
		progress.SetFraction(0.1 + 0.9*progress.GetFraction()) //easter egg
	})
	button2.Connect("clicked", func() {
		progress.Pulse()
	})

	window.ShowAll()
	window.SetSizeRequest(200, 200)

	gtk.Main()
}
