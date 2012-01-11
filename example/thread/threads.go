package main

import "github.com/mattn/go-gtk/gdk"
import "github.com/mattn/go-gtk/gtk"
import "strconv"
import "time"
import "runtime"

func main() {
	runtime.GOMAXPROCS(10)
	gdk.ThreadsInit()
	gdk.ThreadsEnter()
	gtk.Init(nil)
	window := gtk.Window(gtk.GTK_WINDOW_TOPLEVEL)
	window.Connect("destroy", gtk.MainQuit)

	vbox := gtk.VBox(false, 1)

	label1 := gtk.Label("")
	vbox.Add(label1)
	label2 := gtk.Label("")
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
