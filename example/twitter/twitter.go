package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/mattn/go-gtk/gdk"
	"github.com/mattn/go-gtk/gdkpixbuf"
	"github.com/mattn/go-gtk/gtk"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

func readURL(url string) ([]byte, *http.Response) {
	r, err := http.Get(url)
	if err != nil {
		return nil, nil
	}
	var b []byte
	if b, err = ioutil.ReadAll(r.Body); err != nil {
		return nil, nil
	}
	return b, r
}

func bytes2pixbuf(data []byte, typ string) *gdkpixbuf.GdkPixbuf {
	var loader *gdkpixbuf.GdkPixbufLoader
	if strings.Index(typ, "jpeg") >= 0 {
		loader, _ = gdkpixbuf.PixbufLoaderWithMimeType("image/jpeg")
	} else {
		loader, _ = gdkpixbuf.PixbufLoaderWithMimeType("image/png")
	}
	loader.SetSize(24, 24)
	loader.Write(data)
	loader.Close()
	return loader.GetPixbuf()
}

func main() {
	gdk.ThreadsInit()
	gtk.Init(&os.Args)
	window := gtk.Window(gtk.GTK_WINDOW_TOPLEVEL)
	window.SetTitle("Twitter!")
	window.Connect("destroy", gtk.MainQuit)

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
					fmt.Println(err)
					return
				}
				var j interface{}
				err = json.NewDecoder(bytes.NewBuffer(b)).Decode(&j)
				if err != nil {
					fmt.Println(err)
					return
				}
				arr := j.([]interface{})
				for i := 0; i < len(arr); i++ {
					data := arr[i].(map[string]interface{})
					icon := data["user"].(map[string]interface{})["profile_image_url"].(string)
					var iter gtk.GtkTextIter
					pixbufbytes, resp := readURL(icon)
					gdk.ThreadsEnter()
					buffer.GetEndIter(&iter)
					if resp != nil {
						buffer.InsertPixbuf(&iter, bytes2pixbuf(pixbufbytes, resp.Header.Get("Content-Type")))
					}
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
			} else {
				fmt.Println(err)
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
