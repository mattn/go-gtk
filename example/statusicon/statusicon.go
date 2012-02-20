package main

import (
	"os"
	"github.com/mattn/go-gtk/gtk"
	"github.com/mattn/go-gtk/glib"
)

func main() {
	gtk.Init(&os.Args)

	glib.SetApplicationName("go-gtk-statusicon-example")

	mi := gtk.MenuItemWithLabel("Popup!")
	mi.Connect("activate", func() {
		gtk.MainQuit()
	})
	nm := gtk.Menu()
	nm.Append(mi)
	nm.ShowAll()

	si := gtk.StatusIconFromStock(gtk.GTK_STOCK_FILE)
	si.SetTitle("StatusIcon Example")
	si.SetTooltipMarkup("StatusIcon Example")
	si.Connect("popup-menu", func(cbx *glib.CallbackContext) {
		nm.Popup(nil, nil, gtk.GtkStatusIconPositionMenu, si, uint(cbx.Args(0)), uint(cbx.Args(1)))
	})

	println(`
Can you see statusicon in systray?
If you don't see it and if you use 'unity', try following.

# gsettings set com.canonical.Unity.Panel systray-whitelist \
  "$(gsettings get com.canonical.Unity.Panel systray-whitelist \|
  sed -e "s/]$/, 'go-gtk-statusicon-example']/")"
`)

	gtk.Main()
}
