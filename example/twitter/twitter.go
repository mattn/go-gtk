package main

import (
	"github.com/mattn/go-gtk/gtk"
	"github.com/mattn/go-gtk/gdk"
	"github.com/mattn/go-gtk/gdkpixbuf"
	"http"
	"json"
	"bytes"
	"io"
	"io/ioutil"
	"os"
	"strings"
)

func url2pixbuf(url string) *gdkpixbuf.GdkPixbuf {
	r, err := http.Get(url)
	if err != nil {
		return nil
	}
	t := r.Header.Get("Content-Type")
	b := make([]byte, r.ContentLength)
	if _, err = io.ReadFull(r.Body, b); err != nil {
		return nil
	}
	var loader *gdkpixbuf.GdkPixbufLoader
	if strings.Index(t, "jpeg") >= 0 {
		loader, _ = gdkpixbuf.PixbufLoaderWithMimeType("image/jpeg")
	} else {
		loader, _ = gdkpixbuf.PixbufLoaderWithMimeType("image/png")
	}
	loader.SetSize(24, 24)
	loader.Write(b)
	loader.Close()
	return loader.GetPixbuf()
}

func main() {
	gdk.ThreadsInit()
	gtk.Init(&os.Args)
	window := gtk.Window(gtk.GTK_WINDOW_TOPLEVEL)
	window.SetTitle("Twitter!")
	window.Connect("destroy", func() {
		gtk.MainQuit()
	})

	vbox := gtk.VBox(false, 1)

	scrolledwin := gtk.ScrolledWindow(nil, nil)
	textview := gtk.TextView()
	textview.SetEditable(false)
	textview.SetCursorVisible(false)
	scrolledwin.Add(textview)
	vbox.Add(scrolledwin)

	buffer := textview.GetBuffer()

	tag := buffer.CreateTag("blue", map[string]string{
		"foreground": "#0000FF", "weight": "700"})
	button := gtk.ButtonWithLabel("Update Timeline")
	button.SetTooltipMarkup("update <b>public timeline</b>")
	button.Clicked(func() {
		go func() {
			gdk.ThreadsEnter()
			button.SetSensitive(false)
			gdk.ThreadsLeave()
			r, err := http.Get("http://twitter.com/statuses/public_timeline.json")
			if err == nil {
				var b []byte
				if r.ContentLength == -1 {
					b, err = ioutil.ReadAll(r.Body)
				} else {
					b = make([]byte, r.ContentLength)
					_, err = io.ReadFull(r.Body, b)
				}
				if err != nil {
					println(err.String())
					return
				}
				var j interface{}
				json.NewDecoder(bytes.NewBuffer(b)).Decode(&j)
				arr := j.([]interface{})
				for i := 0; i < len(arr); i++ {
					data := arr[i].(map[string]interface{})
					icon := data["user"].(map[string]interface{})["profile_image_url"].(string)
					var iter gtk.GtkTextIter
					gdk.ThreadsEnter()
					buffer.GetStartIter(&iter)
					buffer.InsertPixbuf(&iter, url2pixbuf(icon))
					gdk.ThreadsLeave()
					name := data["user"].(map[string]interface{})["screen_name"].(string)
					text := data["text"].(string)
					gdk.ThreadsEnter()
					buffer.Insert(&iter, " ")
					buffer.InsertWithTag(&iter, name, tag)
					buffer.Insert(&iter, ":"+text+"\n")
					gtk.MainIterationDo(false)
					gdk.ThreadsLeave()
				}
			}
			button.SetSensitive(true)
		}()
	})
	vbox.PackEnd(button, false, false, 0)

	window.Add(vbox)
	window.SetSizeRequest(800, 500)
	window.ShowAll()
	gdk.ThreadsEnter()
	gtk.Main()
	gdk.ThreadsLeave()
}
