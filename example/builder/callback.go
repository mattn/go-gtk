package callback

import "C"
import "gtk"

//export on_imagemenuitem1_activate
func on_imagemenuitem1_activate() {
	// TODO: FIXME: this callback make a panic().
	gtk.MainQuit()
}
