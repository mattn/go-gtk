package main

import (
	"fmt"
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

	palette := gtk.NewToolPalette()
	group := gtk.NewToolItemGroup("Tools")
	b := gtk.NewToolButtonFromStock(gtk.STOCK_NEW)
	b.OnClicked(func() { fmt.Println("You clicked new!") })
	group.Insert(b, -1)
	b = gtk.NewToolButtonFromStock(gtk.STOCK_CLOSE)
	group.Insert(b, -1)
	b = gtk.NewToolButtonFromStock(gtk.STOCK_REDO)
	group.Insert(b, -1)
	b = gtk.NewToolButtonFromStock(gtk.STOCK_REFRESH)
	group.Insert(b, -1)
	b = gtk.NewToolButtonFromStock(gtk.STOCK_QUIT)
	group.Insert(b, -1)
	b = gtk.NewToolButtonFromStock(gtk.STOCK_YES)
	group.Insert(b, -1)
	b = gtk.NewToolButtonFromStock(gtk.STOCK_NO)
	group.Insert(b, -1)
	b = gtk.NewToolButtonFromStock(gtk.STOCK_PRINT)
	group.Insert(b, -1)
	b = gtk.NewToolButtonFromStock(gtk.STOCK_NETWORK)
	group.Insert(b, -1)
	b = gtk.NewToolButtonFromStock(gtk.STOCK_INFO)
	group.Insert(b, -1)
	b = gtk.NewToolButtonFromStock(gtk.STOCK_HOME)
	group.Insert(b, -1)
	b = gtk.NewToolButtonFromStock(gtk.STOCK_INDEX)
	group.Insert(b, -1)
	b = gtk.NewToolButtonFromStock(gtk.STOCK_FIND)
	group.Insert(b, -1)
	b = gtk.NewToolButtonFromStock(gtk.STOCK_FILE)
	group.Insert(b, -1)
	b = gtk.NewToolButtonFromStock(gtk.STOCK_EXECUTE)
	group.Insert(b, -1)
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
	align := gtk.NewAlignment(0, 0, 0, 0)
	image := gtk.NewImageFromFile("./turkey.jpg")
	align.Add(image)
	frame.Add(align)

	box.Pack1(palette, true, false)
	box.Pack2(frame, false, false)

	window.Add(box)
	window.SetSizeRequest(600, 600)
	window.ShowAll()
	gtk.Main()
}
