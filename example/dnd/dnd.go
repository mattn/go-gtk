package main

import (
	"os"
	"github.com/mattn/go-gtk/gdk"
	"github.com/mattn/go-gtk/glib"
	"github.com/mattn/go-gtk/gtk"
	"strings"
	"unsafe"
)

func main() {
	gtk.Init(&os.Args)
	window := gtk.Window(gtk.GTK_WINDOW_TOPLEVEL)
	window.SetTitle("GTK DrawingArea")
	window.Connect("destroy", gtk.MainQuit)

	vbox := gtk.VBox(true, 0)
	vbox.SetBorderWidth(5)

	targets := []gtk.GtkTargetEntry {
			{"text/uri-list", 0, 0 },
			{"STRING", 0, 1 },
			{"text/plain", 0, 2 },
		}
	dest := gtk.Label("drop me file")
	dest.DragDestSet(
		gtk.GTK_DEST_DEFAULT_MOTION |
		gtk.GTK_DEST_DEFAULT_HIGHLIGHT |
		gtk.GTK_DEST_DEFAULT_DROP,
		targets,
		gdk.GDK_ACTION_COPY)
	dest.DragDestAddUriTargets()
	dest.Connect("drag-data-received", func(ctx *glib.CallbackContext) {
		sdata := gtk.SelectionDataFromNative(unsafe.Pointer(ctx.Args(3)))
		if sdata != nil {
			a := (*[2000]uint8)(sdata.GetData())
			files := strings.Split(string(a[0:sdata.GetLength()-1]), "\n")
			for i := range files {
				filename, _, _ := glib.FilenameFromUri(files[i])
				files[i] = filename
			}
			dialog := gtk.MessageDialog(
				window,
				gtk.GTK_DIALOG_MODAL,
				gtk.GTK_MESSAGE_INFO,
				gtk.GTK_BUTTONS_OK,
				strings.Join(files, "\n"))
			dialog.SetTitle("D&D")
			dialog.Response(func() {
				dialog.Destroy()
			})
			dialog.Run()
		}
	})
	vbox.Add(dest)

	window.Add(vbox)

	window.SetSizeRequest(300, 100)
	window.ShowAll()
	gtk.Main()
}
