package main

import (
	"github.com/mattn/go-gtk/gtk"
	"os"
	"strconv"
)

func main() {
	gtk.Init(&os.Args)
	window := gtk.NewWindow(gtk.WINDOW_TOPLEVEL)
	window.SetTitle("GTK Notebook")
	window.Connect("destroy", gtk.MainQuit)

	notebook := gtk.NewNotebook()
	for n := 1; n <= 10; n++ {
		page := gtk.NewFrame("demo" + strconv.Itoa(n))
		notebook.AppendPage(page, gtk.NewLabel("demo"+strconv.Itoa(n)))

		vbox := gtk.NewHBox(false, 1)

		prev := gtk.NewButtonWithLabel("go prev")
		prev.Clicked(func() {
			notebook.PrevPage()
		})
		vbox.Add(prev)

		next := gtk.NewButtonWithLabel("go next")
		next.Clicked(func() {
			notebook.NextPage()
		})
		vbox.Add(next)

		page.Add(vbox)
	}

	window.Add(notebook)
	window.SetSizeRequest(400, 200)
	window.ShowAll()

	gtk.Main()
}
