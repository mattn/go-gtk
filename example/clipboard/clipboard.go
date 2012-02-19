package main

import (
	"github.com/mattn/go-gtk/gdk"
	"github.com/mattn/go-gtk/gtk"
)

func main() {
	gtk.Init(nil)
	clipboard := gtk.ClipboardGetForDisplay(
		gdk.DisplayGetDefault(),
		gdk.GDK_SELECTION_CLIPBOARD)
	println(clipboard.WaitForText())
	clipboard.SetText("helloworld")
	gtk.MainIterationDo(true)
	clipboard.Store()
	gtk.MainIterationDo(true)
}
