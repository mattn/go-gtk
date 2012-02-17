package callback

import "C"
import "github.com/mattn/go-gtk/gtk"
import "unsafe"

var aboutdialog *gtk.GtkAboutDialog

func Init(builder *gtk.GtkBuilder) {
	aboutdialog = &gtk.GtkAboutDialog{
		*(*gtk.GtkDialog)(unsafe.Pointer(&builder.GetObject("aboutdialog1").Object))}
}

//export on_imagemenuitem1_activate
func on_imagemenuitem1_activate() {
	gtk.MainQuit()
}

//export on_show_aboutdialog_activate
func on_show_aboutdialog_activate() {
	//gtk.MainQuit()
	aboutdialog.Run()
}
