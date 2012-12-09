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

	vbox := gtk.NewVBox(false, 0)

	bnew := gtk.NewToolButtonFromStock(gtk.STOCK_NEW)
	bclose := gtk.NewToolButtonFromStock(gtk.STOCK_CLOSE)

	palette := gtk.NewToolPalette()
	group := gtk.NewToolItemGroup("Tools")
	group.Insert(bnew, -1)
	group.Insert(bclose, -1)
	palette.Add(group)

	vbox.PackStart(palette, true, true, 1)

	bcopy := gtk.NewToolButtonFromStock(gtk.STOCK_COPY)
	bcut := gtk.NewToolButtonFromStock(gtk.STOCK_CUT)
	bdelete := gtk.NewToolButtonFromStock(gtk.STOCK_DELETE)

	palette = gtk.NewToolPalette()
	group = gtk.NewToolItemGroup("Stuff")
	group.Insert(bcopy, -1)
	group.Insert(bcut, -1)
	group.Insert(bdelete, -1)
	palette.Add(group)

	vbox.PackStart(palette, true, true, 1)

	window.Add(vbox)
	window.SetSizeRequest(600, 600)
	window.ShowAll()
	gtk.Main()
}
