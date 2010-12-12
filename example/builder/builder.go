package main

import (
	"os"
	"gtk"
	"callback"
)

func main() {
	callback.Init()

	gtk.Init(&os.Args)
	builder := gtk.Builder()
	builder.AddFromFile("hello.ui")
	builder.ConnectSignals(nil)
	obj := builder.GetObject("window1")

	window := gtk.WidgetFromObject(obj)
	window.Show()
	window.Connect("destroy", func() {
		gtk.MainQuit()
	},
		nil)

	gtk.Main()
}
