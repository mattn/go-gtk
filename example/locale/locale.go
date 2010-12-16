package main

import "gtk"
import "glib"
import "fmt"

func main() {
	gtk.SetLocale()

	bs := ([]byte)("こんにちわ世界")

	str, bytes_read, bytes_written, error := glib.LocaleToUtf8(bs)
	if error == nil {
		fmt.Printf("str=%s, bytes_read=%d, bytes_written=%d\n", str, bytes_read, bytes_written)
	} else {
		println(error.Message())
	}
}
