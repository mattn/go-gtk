package main

import (
	"github.com/mattn/go-gtk/glib"
	"github.com/mattn/go-gtk/gtk"		
	"fmt"
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
	hbox := gtk.HBox(false, 1)

	//--------------------------------------------------------
	// GtkToolbar
	//--------------------------------------------------------
	toolbar := gtk.Toolbar()
	toolbar.SetStyle(gtk.GTK_TOOLBAR_ICONS)
	vbox.PackStart(toolbar, false, false, 5)

	btnnew := gtk.ToolButtonFromStock(gtk.GTK_STOCK_NEW)
	btnclose := gtk.ToolButtonFromStock(gtk.GTK_STOCK_CLOSE)
	separator := gtk.SeparatorToolItem()
	btncustom := gtk.ToolButton(nil, "Custom")	

	btnnew.OnClicked(onToolButtonClicked)	
	btnclose.OnClicked(onToolButtonClicked)	
	btncustom.OnClicked(onToolButtonClicked)	

	toolbar.Insert(btnnew, -1)	
	toolbar.Insert(btnclose, -1)
	toolbar.Insert(separator, -1)		
	toolbar.Insert(btncustom, -1)

	toolbar2 := gtk.Toolbar()
	toolbar2.SetOrientation(gtk.GTK_ORIENTATION_VERTICAL)
	hbox.PackStart(toolbar2, false, false, 5)
	btnhelp := gtk.ToolButtonFromStock(gtk.GTK_STOCK_HELP)
	btnhelp.OnClicked(onToolButtonClicked)	
	toolbar2.Insert(btnhelp, -1)	

	btntoggle := gtk.ToggleToolButton()
	btntoggle2 := gtk.ToggleToolButtonFromStock(gtk.GTK_STOCK_ITALIC)
	toolbar2.Insert(btntoggle, -1)	
	toolbar2.Insert(btntoggle2, -1)		

	for i := 0; i < toolbar.GetNItems(); i++ {
		gti := toolbar.GetNthItem(i)
		switch gti.(type) {
		case *gtk.GtkToolButton:
		    fmt.Printf("toolbar[%d] is a *gtk.GtkToolButton\n", i)		    
		case *gtk.GtkToggleToolButton:
		    fmt.Printf("toolbar[%d] is a *gtk.GtkToggleToolButton\n", i)	
		    gti.(*gtk.GtkToggleToolButton).SetActive(true)
		case *gtk.GtkSeparatorToolItem:
		    fmt.Printf("toolbar[%d] is a *gtk.GtkSeparatorToolItem\n", i)		    
		default:
		    fmt.Printf("toolbar: Item is of unknown type")	
		}
	}

	for i := 0; i < toolbar2.GetNItems(); i++ {
		gti := toolbar2.GetNthItem(i)
		switch gti.(type) {
		case *gtk.GtkToolButton:
		    fmt.Printf("toolbar2[%d] is a *gtk.GtkToolButton\n", i)		    		    
		case *gtk.GtkToggleToolButton:
		    fmt.Printf("toolbar2[%d] is a *gtk.GtkToggleToolButton\n", i)	
		    gti.(*gtk.GtkToggleToolButton).SetActive(true)
		case *gtk.GtkSeparatorToolItem:
		    fmt.Printf("toolbar2[%d] is a *gtk.GtkSeparatorToolItem\n", i)		    
		default:
		    fmt.Printf("toolbar2: Item is of unknown type")	
		}
	}		

	vbox.Add(hbox)
	window.Add(vbox)	
	window.SetSizeRequest(600, 600)
	window.ShowAll()
	gtk.Main()
}

func onToolButtonClicked(args ...interface{}) {
	println("ToolButton clicked")			
}