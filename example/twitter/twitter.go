package main

import "gtk"
import "gdkpixbuf"
import "unsafe"
import "http"
import "json"
import "io"
import "os"
import "strconv"
import "strings"

func url2pixbuf(url string) *gdkpixbuf.GdkPixbuf {
	if r, _, err := http.Get(url); err == nil {
		n, _ := strconv.Atoi64(r.GetHeader("Content-Length"));
		t := r.GetHeader("Content-Type");
		b := make([]byte, n);
		io.ReadFull(r.Body, b);
		var loader *gdkpixbuf.GdkPixbufLoader;
		if strings.Index(t, "jpeg") >= 0 {
			loader, _ = gdkpixbuf.PixbufLoaderWithMimeType("image/jpeg");
		} else {
			loader, _ = gdkpixbuf.PixbufLoaderWithMimeType("image/png");
		}
		loader.SetSize(24, 24);
		loader.Write(b);
		loader.Close();
		return loader.GetPixbuf();
	}
	return nil;
}

func main() {
	gtk.Init(&os.Args);
	window := gtk.Window(gtk.GTK_WINDOW_TOPLEVEL);
	window.SetTitle("Twitter!");
	window.Connect("destroy", func(w *gtk.GtkWidget, args []unsafe.Pointer) {
		gtk.MainQuit();
	}, nil);

	vbox := gtk.VBox(false, 1);

	scrolledwin := gtk.ScrolledWindow(nil, nil);
	textview := gtk.TextView();
	textview.SetEditable(false);
	textview.SetCursorVisible(false);
	scrolledwin.Add(textview);
	vbox.Add(scrolledwin);

	buffer := textview.GetBuffer();

	tag := buffer.CreateTag("blue", map[string] string {
		"foreground": "#0000FF", "weight": "700" });
	button := gtk.ButtonWithLabel("Update Timeline");
	button.SetTooltipMarkup("update <b>public timeline</b>");
	button.Clicked(func(w *gtk.GtkWidget, args []unsafe.Pointer) {
		button.SetSensitive(false);
		go func() {
			if r, _, err := http.Get("http://twitter.com/statuses/public_timeline.json"); err == nil {
				n, _ := strconv.Atoi64(r.GetHeader("Content-Length"));
				b := make([]byte, n);
				io.ReadFull(r.Body, b);
				j, _ := json.Decode(string(b));
				arr := j.([]interface{});
				for i := 0; i < len(arr); i++ {
					data := arr[i].(map[string]interface{});
					icon := data["user"].(map[string]interface{})["profile_image_url"].(string);
					var iter gtk.GtkTextIter;
					buffer.GetStartIter(&iter);
					buffer.InsertPixbuf(&iter, url2pixbuf(icon));
					name := data["user"].(map[string]interface{})["screen_name"].(string);
					text := data["text"].(string);
					buffer.Insert(&iter, " ");
					buffer.InsertWithTag(&iter, name, tag);
					buffer.Insert(&iter, ":" + text + "\n");
				}  
			}
			button.SetSensitive(true);
		}();
	}, nil);
	vbox.PackEnd(button, false, false, 0);

	window.Add(vbox);
	window.SetSizeRequest(300, 400);
	window.ShowAll();
	gtk.Main();
}
