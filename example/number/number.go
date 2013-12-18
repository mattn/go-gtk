package main

import (
	"fmt"
	"github.com/mattn/go-gtk/glib"
	"github.com/mattn/go-gtk/gtk"
	"os"
	"strconv"
	"unsafe"
)

func main() {
	gtk.Init(&os.Args)

	dialog := gtk.NewDialog()
	dialog.SetTitle("number input")

	vbox := dialog.GetVBox()

	label := gtk.NewLabel("Numnber:")
	vbox.Add(label)

	input := gtk.NewEntry()
	input.SetEditable(true)
	vbox.Add(input)

	input.Connect("insert-text", func(ctx *glib.CallbackContext) {
		a := (*[2000]uint8)(unsafe.Pointer(ctx.Args(0)))
		p := (*int)(unsafe.Pointer(ctx.Args(2)))
		i := 0
		for a[i] != 0 {
			i++
		}
		s := string(a[0:i])
		if s == "." {
			if *p == 0 {
				input.StopEmission("insert-text")
			}
		} else {
			_, err := strconv.ParseFloat(s, 64)
			if err != nil {
				input.StopEmission("insert-text")
			}
		}
	})

	button := gtk.NewButtonWithLabel("OK")
	button.Connect("clicked", func() {
		fmt.Println(input.GetText())
		gtk.MainQuit()
	})
	vbox.Add(button)

	dialog.ShowAll()
	gtk.Main()
}
