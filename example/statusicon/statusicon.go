package main

import (
	"os"
	"github.com/mattn/go-gtk/gtk"
	"github.com/mattn/go-gtk/glib"
)

func main() {
	gtk.Init(&os.Args)

	mi := gtk.MenuItemWithLabel("Popup!")
	mi.Connect("activate", func() {
		gtk.MainQuit()
	})
	nm := gtk.Menu()
	nm.Append(mi)
	nm.ShowAll()

	si := gtk.StatusIconFromStock(gtk.GTK_STOCK_FILE)
	si.Connect("popup-menu", func(cbx *glib.CallbackContext) {
		nm.Popup(nil, nil, gtk.GtkStatusIconPositionMenu, si, uint(cbx.Args(0)), uint(cbx.Args(1)))
	})

	gtk.Main()
}
