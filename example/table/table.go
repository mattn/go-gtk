package main

import (
	"fmt"
	"github.com/mattn/go-gtk/gtk"
	"os"
)

func main() {
	gtk.Init(&os.Args)
	window := gtk.NewWindow(gtk.WINDOW_TOPLEVEL)
	window.SetTitle("GTK Table")
	window.Connect("destroy", gtk.MainQuit)

	swin := gtk.NewScrolledWindow(nil, nil)
	swin.SetPolicy(gtk.POLICY_AUTOMATIC, gtk.POLICY_AUTOMATIC)

	table := gtk.NewTable(5, 5, false)
	for y := uint(0); y < 5; y++ {
		for x := uint(0); x < 5; x++ {
			table.Attach(gtk.NewButtonWithLabel(fmt.Sprintf("%02d:%02d", x, y)), x, x+1, y, y+1, gtk.FILL, gtk.FILL, 5, 5)
		}
	}
	swin.AddWithViewPort(table)

	window.Add(swin)
	window.SetDefaultSize(200, 200)
	window.ShowAll()

	gtk.Main()
}
