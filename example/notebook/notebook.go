package main

import (
	"os"
	"gtk"
	"strconv"
)

func main() {
	gtk.Init(&os.Args)
	window := gtk.Window(gtk.GTK_WINDOW_TOPLEVEL)
	window.SetTitle("GTK Notebook")
	window.Connect("destroy", func() {
		gtk.MainQuit()
	},
		nil)

	notebook := gtk.Notebook()
	for n := 1; n <= 10; n++ {
		page := gtk.Frame("demo" + strconv.Itoa(n))
		notebook.AppendPage(page, gtk.Label("demo"+strconv.Itoa(n)))

		vbox := gtk.HBox(false, 1)

		prev := gtk.ButtonWithLabel("go perv")
		prev.Clicked(func() {
			notebook.PrevPage()
		},
			nil)
		vbox.Add(prev)

		next := gtk.ButtonWithLabel("go next")
		next.Clicked(func() {
			notebook.NextPage()
		},
			nil)
		vbox.Add(next)

		page.Add(vbox)
	}

	window.Add(notebook)
	window.SetSizeRequest(400, 200)
	window.ShowAll()

	gtk.Main()
}
