package main

import (
	"os"
	"github.com/mattn/go-gtk/gtk"
)

func main() {
	gtk.Init(&os.Args)
	builder := gtk.Builder()
	builder.AddFromFile("hello.ui")
	builder.ConnectSignals(nil)
	obj := builder.GetObject("window1")

	window := gtk.WidgetFromObject(obj)
	window.Show()
	window.Connect("destroy", func() {
		gtk.MainQuit()
	})

	gtk.Main()
}
