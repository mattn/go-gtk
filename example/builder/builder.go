package main

import (
	"os"

	"github.com/mattn/go-gtk/example/builder/callback"
	"github.com/mattn/go-gtk/gtk"
)

//"github.com/mattn/go-gtk/example/builder/callback"
func main() {
	gtk.Init(&os.Args)

	builder := gtk.NewBuilder()

	builder.AddFromFile("hello.ui")
	obj := builder.GetObject("window1")

	window := gtk.WidgetFromObject(obj)
	window.Show()
	window.Connect("destroy", gtk.MainQuit)

	callback.Init(builder)

	gtk.Main()
}
