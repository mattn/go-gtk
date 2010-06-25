package callback

/*
static inline void dummy_function() { }
*/
import "C"
import "gtk"

func Init() {
	C.dummy_function()
}

//export on_imagemenuitem1_activate
func on_imagemenuitem1_activate() {
	gtk.MainQuit()
}
