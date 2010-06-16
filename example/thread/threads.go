package main

import "gdk"
import "gtk"
import "strconv"
import "syscall"

func main() {
	gdk.ThreadsInit();
	gtk.Init(nil);
	w := gtk.Window(gtk.GTK_WINDOW_TOPLEVEL);

	v := gtk.VBox(false, 1)

	l1 := gtk.Label("");
	v.Add(l1);
	l2 := gtk.Label("");
	v.Add(l2);

	w.Add(v);

	w.SetSizeRequest(100, 100);
	w.ShowAll();
	syscall.Sleep(1000*1000*100);
	go (func() {
		for i := 0; i < 300000; i++ {
			gdk.ThreadsEnter();
			l1.SetLabel(strconv.Itoa(i));
			gdk.ThreadsLeave();
		}
		gtk.MainQuit();
	})();
	go (func() {
		for i := 300000; i >= 0; i-- {
			gdk.ThreadsEnter();
			l2.SetLabel(strconv.Itoa(i));
			gdk.ThreadsLeave();
		}
		gtk.MainQuit();
	})();
	gtk.Main();
}
