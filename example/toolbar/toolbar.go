package main

import (
	"github.com/mattn/go-gtk/glib"
	"github.com/mattn/go-gtk/gtk"	
	"fmt"
)

func main() {	
	gtk.Init(nil)
	window := gtk.NewWindow(gtk.WINDOW_TOPLEVEL)
	window.SetPosition(gtk.WIN_POS_CENTER)
	window.SetTitle("GTK Go!")
	window.Connect("destroy", func(ctx *glib.CallbackContext) {		
		gtk.MainQuit()
	}, "")
	
	vbox := gtk.NewVBox(false, 1)
	hbox := gtk.NewHBox(false, 1)

	//--------------------------------------------------------
	// GtkToolbar
	//--------------------------------------------------------
	toolbar := gtk.NewToolbar()
	toolbar.SetStyle(gtk.TOOLBAR_ICONS)
	vbox.PackStart(toolbar, false, false, 5)

	btnnew := gtk.NewToolButtonFromStock(gtk.STOCK_NEW)
	btnclose := gtk.NewToolButtonFromStock(gtk.STOCK_CLOSE)
	separator := gtk.NewSeparatorToolItem()
	btncustom := gtk.NewToolButton(nil, "Custom")	

	btnnew.OnClicked(onToolButtonClicked)	
	btnclose.OnClicked(onToolButtonClicked)	
	btncustom.OnClicked(onToolButtonClicked)	

	toolbar.Insert(btnnew, -1)	
	toolbar.Insert(btnclose, -1)
	toolbar.Insert(separator, -1)		
	toolbar.Insert(btncustom, -1)

	toolbar2 := gtk.NewToolbar()
	toolbar2.SetOrientation(gtk.ORIENTATION_VERTICAL)
	hbox.PackStart(toolbar2, false, false, 5)
	btnhelp := gtk.NewToolButtonFromStock(gtk.STOCK_HELP)
	btnhelp.OnClicked(onToolButtonClicked)	
	toolbar2.Insert(btnhelp, -1)	

	btntoggle := gtk.NewToggleToolButton()
	btntoggle2 := gtk.NewToggleToolButtonFromStock(gtk.STOCK_ITALIC)
	toolbar2.Insert(btntoggle, -1)	
	toolbar2.Insert(btntoggle2, -1)		

	for i := 0; i < toolbar.GetNItems(); i++ {
		gti := toolbar.GetNthItem(i)
		switch gti.(type) {
		case *gtk.ToolButton:
		    fmt.Printf("toolbar[%d] is a *gtk.ToolButton\n", i)		    
		case *gtk.ToggleToolButton:
		    fmt.Printf("toolbar[%d] is a *gtk.ToggleToolButton\n", i)	
		    gti.(*gtk.ToggleToolButton).SetActive(true)
		case *gtk.SeparatorToolItem:
		    fmt.Printf("toolbar[%d] is a *gtk.SeparatorToolItem\n", i)		    
		default:
		    fmt.Printf("toolbar: Item is of unknown type")	
		}
	}

	for i := 0; i < toolbar2.GetNItems(); i++ {
		gti := toolbar2.GetNthItem(i)
		switch gti.(type) {
		case *gtk.ToolButton:
		    fmt.Printf("toolbar2[%d] is a *gtk.ToolButton\n", i)		    		    
		case *gtk.ToggleToolButton:
		    fmt.Printf("toolbar2[%d] is a *gtk.ToggleToolButton\n", i)	
		    gti.(*gtk.ToggleToolButton).SetActive(true)
		case *gtk.SeparatorToolItem:
		    fmt.Printf("toolbar2[%d] is a *gtk.SeparatorToolItem\n", i)		    
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

func onToolButtonClicked() {
	println("ToolButton clicked")			
}