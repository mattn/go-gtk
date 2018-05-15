package callback

import (
	"unsafe"

	"github.com/mattn/go-gtk/glib"
	"github.com/mattn/go-gtk/gtk"
)

var aboutdialog *gtk.AboutDialog

func Init(builder *gtk.Builder) {
	aboutdialog = &gtk.AboutDialog{
		*(*gtk.Dialog)(unsafe.Pointer(&builder.GetObject("aboutdialog1").Object))}
	builder.ConnectSignalsFull(func(builder *gtk.Builder, obj *glib.GObject, sig, handler string, conn *glib.GObject, flags glib.ConnectFlags, user_data interface{}) {
		switch handler {
		case "on_imagemenuitem1_activate":
			obj.SignalConnect(sig, on_imagemenuitem1_activate, user_data, flags)
		case "on_show_aboutdialog_activate":
			obj.SignalConnect(sig, on_show_aboutdialog_activate, user_data, flags)
		case "gtk_widget_hide":
			obj.SignalConnect(sig, func(c *glib.CallbackContext) {
				gtk.WidgetFromObject(c.Target().(*glib.GObject)).Hide()
			}, nil, flags)
		}
	}, nil)
}

//export on_imagemenuitem1_activate
func on_imagemenuitem1_activate() {
	gtk.MainQuit()
}

//export on_show_aboutdialog_activate
func on_show_aboutdialog_activate() {
	aboutdialog.Run()
}
