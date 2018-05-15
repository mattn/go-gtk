package main

import (
	"fmt"
	"github.com/mattn/go-gtk/gdk"
	"github.com/mattn/go-gtk/glib"
	"github.com/mattn/go-gtk/gtk"
	"unsafe"
)

func main() {
	gtk.Init(nil)

	window := gtk.NewWindow(gtk.WINDOW_TOPLEVEL)
	window.SetPosition(gtk.WIN_POS_CENTER)
	window.SetTitle("GTK Go!")
	window.SetIconName("textview")
	window.Connect("destroy", gtk.MainQuit)

	textview := gtk.NewTextView()
	textview.SetEditable(true)
	textview.SetCursorVisible(true)
	var iter gtk.TextIter
	buffer := textview.GetBuffer()

	buffer.GetStartIter(&iter)
	buffer.Insert(&iter, "Hello ")

	tag := buffer.CreateTag("bold", map[string]string{"background": "#FF0000", "weight": "700"})
	buffer.InsertWithTag(&iter, "Google!", tag)

	u := "http://www.google.com"
	tag.SetData("tag-name", unsafe.Pointer(&u))
	textview.Connect("event-after", func(ctx *glib.CallbackContext) {
		arg := ctx.Args(0)
		if ev := *(**gdk.EventAny)(unsafe.Pointer(&arg)); ev.Type != gdk.BUTTON_RELEASE {
			return
		}
		ev := *(**gdk.EventButton)(unsafe.Pointer(&arg))
		var iter gtk.TextIter
		textview.GetIterAtLocation(&iter, int(ev.X), int(ev.Y))
		tags := iter.GetTags()
		for n := uint(0); n < tags.Length(); n++ {
			vv := tags.NthData(n)
			tag := gtk.NewTextTagFromPointer(vv)
			u := *(*string)(tag.GetData("tag-name"))
			fmt.Println(u)
		}
	})

	window.Add(textview)
	window.SetSizeRequest(600, 600)
	window.ShowAll()
	gtk.Main()
}
