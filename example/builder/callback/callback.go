package callback

import "C"
import "github.com/mattn/go-gtk/gtk"
import "unsafe"

var aboutdialog *gtk.AboutDialog

func Init(builder *gtk.Builder) {
	aboutdialog = &gtk.AboutDialog{
		*(*gtk.Dialog)(unsafe.Pointer(&builder.GetObject("aboutdialog1").Object))}
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
