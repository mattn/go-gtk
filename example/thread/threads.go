package main

import "gdk"
import "gtk"
import "strconv"
import "syscall"

func main() {
	gtk.Init(nil);
	gdk.ThreadsInit();
	w := gtk.Window(gtk.GTK_WINDOW_TOPLEVEL);
	l := gtk.Label("");
	w.Add(l);
	w.ShowAll();
	syscall.Sleep(1000*1000*100);
	go (func() {
		for i := 0; i < 100000; i++ {
			gdk.ThreadsEnter();
			l.SetLabel(strconv.Itoa(i));
			gdk.ThreadsLeave();
		}
		gtk.MainQuit();
	})();
	gtk.Main();
}
