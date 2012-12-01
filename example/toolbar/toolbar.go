package main

import (
	"github.com/mattn/go-gtk/glib"
	"github.com/mattn/go-gtk/gtk"	
)

func main() {	
	gtk.Init(nil)
	window := gtk.Window(gtk.GTK_WINDOW_TOPLEVEL)
	window.SetPosition(gtk.GTK_WIN_POS_CENTER)
	window.SetTitle("GTK Go!")
	window.Connect("destroy", func(ctx *glib.CallbackContext) {		
		gtk.MainQuit()
	}, "")
	
	vbox := gtk.VBox(false, 1)

	//--------------------------------------------------------
	// GtkToolbar
	//--------------------------------------------------------
	toolbar := gtk.Toolbar()
	toolbar.SetStyle(gtk.GTK_TOOLBAR_ICONS)
	vbox.PackStart(toolbar, false, false, 5)

	btnnew := gtk.ToolButtonFromStock(gtk.GTK_STOCK_NEW);
	btnclose := gtk.ToolButtonFromStock(gtk.GTK_STOCK_CLOSE);
	toolbar.Insert(btnnew, -1)	
	toolbar.Insert(btnclose, -1)	
	
	window.Add(vbox)
	window.SetSizeRequest(600, 600)
	window.ShowAll()
	gtk.Main()
}
