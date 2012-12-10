package main

import (
	"github.com/mattn/go-gtk/glib"
	"github.com/mattn/go-gtk/gtk"
)

func main() {
	gtk.Init(nil)
	window := gtk.NewWindow(gtk.WINDOW_TOPLEVEL)
	window.SetPosition(gtk.WIN_POS_CENTER)
	window.SetTitle("GTK Go!")
	window.Connect("destroy", func(ctx *glib.CallbackContext) {
		gtk.MainQuit()
	}, "")

	box := gtk.NewHPaned()

	bnew := gtk.NewToolButtonFromStock(gtk.STOCK_NEW)
	bclose := gtk.NewToolButtonFromStock(gtk.STOCK_CLOSE)

	palette := gtk.NewToolPalette()
	group := gtk.NewToolItemGroup("Tools")
	group.Insert(bnew, -1)
	group.Insert(bclose, -1)
	palette.Add(group)	

	bcopy := gtk.NewToolButtonFromStock(gtk.STOCK_COPY)
	bcut := gtk.NewToolButtonFromStock(gtk.STOCK_CUT)
	bdelete := gtk.NewToolButtonFromStock(gtk.STOCK_DELETE)
	
	group = gtk.NewToolItemGroup("Stuff")
	group.Insert(bcopy, -1)
	group.Insert(bcut, -1)
	group.Insert(bdelete, -1)
	palette.Add(group)

	frame := gtk.NewVBox(false, 1)
	image := gtk.NewImageFromFile("./turkey.jpg")
	frame.Add(image)

	box.Pack1(palette, true, false)
	box.Pack2(frame, false, false)

	window.Add(box)
	window.SetSizeRequest(600, 600)
	window.ShowAll()
	gtk.Main()
}
