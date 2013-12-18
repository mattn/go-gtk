package main

import (
	"fmt"
	"github.com/mattn/go-gtk/glib"
	"github.com/mattn/go-gtk/gtk"
	"os"
)

func main() {
	gtk.Init(&os.Args)

	glib.SetApplicationName("go-gtk-statusicon-example")

	mi := gtk.NewMenuItemWithLabel("Popup!")
	mi.Connect("activate", func() {
		gtk.MainQuit()
	})
	nm := gtk.NewMenu()
	nm.Append(mi)
	nm.ShowAll()

	si := gtk.NewStatusIconFromStock(gtk.STOCK_FILE)
	si.SetTitle("StatusIcon Example")
	si.SetTooltipMarkup("StatusIcon Example")
	si.Connect("popup-menu", func(cbx *glib.CallbackContext) {
		nm.Popup(nil, nil, gtk.StatusIconPositionMenu, si, uint(cbx.Args(0)), uint32(cbx.Args(1)))
	})

	fmt.Println(`
Can you see statusicon in systray?
If you don't see it and if you use 'unity', try following.

# gsettings set com.canonical.Unity.Panel systray-whitelist \
  "$(gsettings get com.canonical.Unity.Panel systray-whitelist \|
  sed -e "s/]$/, 'go-gtk-statusicon-example']/")"
`)

	gtk.Main()
}
