package main

import "github.com/mattn/go-gtk/gtk"
import "github.com/mattn/go-gtk/glib"
import "fmt"
import "syscall"

func main() {
	gtk.SetLocale()

	bs, _, _, err := glib.LocaleFromUtf8("こんにちわ世界\n")
	if err == nil {
		syscall.Write(syscall.Stdout, bs)
	} else {
		fmt.Println(err.Message())
	}
}
