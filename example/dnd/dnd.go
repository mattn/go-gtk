package main

import (
	"os"
	"strings"
	"unsafe"

	"github.com/mattn/go-gtk/gdk"
	"github.com/mattn/go-gtk/glib"
	"github.com/mattn/go-gtk/gtk"
)

func main() {
	gtk.Init(&os.Args)
	window := gtk.NewWindow(gtk.WINDOW_TOPLEVEL)
	window.SetTitle("GTK DND")
	window.Connect("destroy", gtk.MainQuit)

	vbox := gtk.NewVBox(true, 0)
	vbox.SetBorderWidth(5)

	targets := []gtk.TargetEntry{
		{"text/uri-list", 0, 0},
		{"STRING", 0, 1},
		{"text/plain", 0, 2},
	}
	dest := gtk.NewLabel("drop me file")
	dest.DragDestSet(
		gtk.DEST_DEFAULT_MOTION|
			gtk.DEST_DEFAULT_HIGHLIGHT|
			gtk.DEST_DEFAULT_DROP,
		targets,
		gdk.ACTION_COPY)
	dest.DragDestAddUriTargets()
	dest.Connect("drag-data-received", func(ctx *glib.CallbackContext) {
		sdata := gtk.NewSelectionDataFromNative(unsafe.Pointer(ctx.Args(3)))
		if sdata != nil {
			a := (*[2000]uint8)(sdata.GetData())
			files := strings.Split(string(a[0:sdata.GetLength()-1]), "\n")
			for i := range files {
				filename, _, _ := glib.FilenameFromUri(files[i])
				files[i] = filename
			}
			dialog := gtk.NewMessageDialog(
				window,
				gtk.DIALOG_MODAL,
				gtk.MESSAGE_INFO,
				gtk.BUTTONS_OK,
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
