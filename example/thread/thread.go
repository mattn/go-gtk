package main

import (
	"runtime"
	"strconv"
	"time"

	"github.com/mattn/go-gtk/gdk"
	"github.com/mattn/go-gtk/glib"
	"github.com/mattn/go-gtk/gtk"
)

func main() {
	runtime.GOMAXPROCS(10)
	glib.ThreadInit(nil)
	gdk.ThreadsInit()
	gdk.ThreadsEnter()
	gtk.Init(nil)
	window := gtk.NewWindow(gtk.WINDOW_TOPLEVEL)
	window.Connect("destroy", gtk.MainQuit)

	vbox := gtk.NewVBox(false, 1)

	label1 := gtk.NewLabel("")
	vbox.Add(label1)
	label2 := gtk.NewLabel("")
	vbox.Add(label2)

	window.Add(vbox)

	window.SetSizeRequest(100, 100)
	window.ShowAll()
	time.Sleep(1000 * 1000 * 100)
	go (func() {
		for i := 0; i < 300000; i++ {
			gdk.ThreadsEnter()
			label1.SetLabel(strconv.Itoa(i))
			gdk.ThreadsLeave()
		}
		gtk.MainQuit()
	})()
	go (func() {
		for i := 300000; i >= 0; i-- {
			gdk.ThreadsEnter()
			label2.SetLabel(strconv.Itoa(i))
			gdk.ThreadsLeave()
		}
		gtk.MainQuit()
	})()
	gtk.Main()
}
