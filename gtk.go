package gtk

/*
#include <gtk/gtk.h>
#include <unistd.h>

static void _gtk_init(void *argc, void *argv) {
	gtk_init((int*)argc, (char***)argv);
}
static void _gtk_container_add(GtkWidget *container, GtkWidget *widget) {
	gtk_container_add(GTK_CONTAINER(container), widget);
}
int _gtk_pipe = 0;
static void _callback(GtkWidget* w, void* data) {
	int n = (int)data;
	write(_gtk_pipe, &n, 1);
}
static long _gtk_signal_connect(GtkWidget *widget, char *name, char *func) {
	return gtk_signal_connect_full(GTK_OBJECT(widget), name, GTK_SIGNAL_FUNC(_callback), 0, 0, 0, 0, 0);
}
static gchar* to_gchar(char* s) { return (gchar*)s; }
*/
import "C";
import "os";
import "fmt";
import "unsafe";
import "container/list";

const (
	GTK_WINDOW_TOPLEVEL = 0
);

var funcs *list.List;
type GtkWidget struct {
	w *C.GtkWidget;
}

func Init(args *[]string) {
	var argc C.int = C.int(len(*args));
	cargs := make([]*C.char, argc);
	for i, arg := range *args { cargs[i] = C.CString(arg) }
	C._gtk_init(unsafe.Pointer(&argc), unsafe.Pointer(&cargs));
	goargs := make([]string, argc);
	for i := 0;i < int(argc); i++ { goargs[i] = C.GoString(cargs[i]); }
	*args = goargs;

	funcs = list.New();
}
func Main() {
	go pollEvents();
	C.gtk_main();
}

func (v *GtkWidget) Hide() { C.gtk_widget_hide(v.w) }
func (v *GtkWidget) HideAll() { C.gtk_widget_hide_all(v.w) }
func (v *GtkWidget) Show() { C.gtk_widget_show(v.w) }
func (v *GtkWidget) ShowAll() { C.gtk_widget_show_all(v.w) }
func (v *GtkWidget) ShowNow() { C.gtk_widget_show_now(v.w) }
func (v *GtkWidget) Destory() { C.gtk_widget_destroy(v.w) }
func (v *GtkWidget) Add(w *GtkWidget) { C._gtk_container_add(v.w, w.w) }
func (v *GtkWidget) Connect(s string, f func(widget *GtkWidget, data unsafe.Pointer)) {
	funcs.PushBack(f);
	ff := fmt.Sprintf("%d", f);
	C._gtk_signal_connect(v.w, C.CString(s), C.CString(ff));
}

func Window(t int) *GtkWidget {
	return &GtkWidget{ C.gtk_window_new(C.GtkWindowType(t)) }
}
func Button(label string) *GtkWidget {
	return &GtkWidget{ C.gtk_button_new_with_label(C.to_gchar(C.CString(label))) }
}

func pollEvents() {
/* TODO: not working */

	r, w, _ := os.Pipe();
	C._gtk_pipe = C.int(w.Fd());
	var b []byte;
	for true {
		n, _ := r.Read(b);
		if (n > 0) {
			println(b[0]);
		}
	}
}
