package main

import (
	"fmt"
	"github.com/mattn/go-gtk/gdk"
	"github.com/mattn/go-gtk/gdkpixbuf"
	"github.com/mattn/go-gtk/gtk"
	"github.com/garyburd/go-oauth"
	"github.com/garyburd/go-twitterstream"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
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

type tweet struct {
	Text       string
	Identifier string `json:"id_str"`
	Source     string
	User       struct {
		Name            string
		ScreenName      string `json:"screen_name"`
		FollowersCount  int    `json:"followers_count"`
		ProfileImageUrl string `json:"profile_image_url"`
	}
	Place *struct {
		Id       string
		FullName string `json:"full_name"`
	}
	Entities struct {
		HashTags []struct {
			Indices [2]int
			Text    string
		}
		UserMentions []struct {
			Indices    [2]int
			ScreenName string `json:"screen_name"`
		}    `json:"user_mentions"`
		Urls []struct {
			Indices     [2]int
			Url         string
			DisplayUrl  string  `json:"display_url"`
			ExpandedUrl *string `json:"expanded_url"`
		}
	}
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
		b, err := ioutil.ReadFile("settings.json")
		if err != nil {
			fmt.Println(`"settings.json" not found: `, err)
			return
		}
		var config map[string]string
		err = json.Unmarshal(b, &config)
		if err != nil {
			fmt.Println(`can't read "settings.json": `, err)
			return
		}
		client := &oauth.Client{
			Credentials: oauth.Credentials{
				config["ClientToken"], config["ClientSecret"]}}
		cred := &oauth.Credentials{
			config["AccessToken"], config["AccessSecret"]}

		gdk.ThreadsEnter()
		button.SetSensitive(false)
		gdk.ThreadsLeave()
		go func() {
			ts, err := twitterstream.Open(client, cred,
				"https://stream.twitter.com/1/statuses/filter.json",
				url.Values{"track": {"picplz,instagr"}})
			if err != nil {
				return
			}
			for ts.Err() == nil {
				t := tweet{}
				if err := ts.UnmarshalNext(&t); err != nil {
					fmt.Println("error reading tweet: ", err)
					continue
				}
				var iter gtk.GtkTextIter
				pixbufbytes, resp := readURL(t.User.ProfileImageUrl)
				gdk.ThreadsEnter()
				buffer.GetStartIter(&iter)
				if resp != nil {
					buffer.InsertPixbuf(&iter, bytes2pixbuf(pixbufbytes, resp.Header.Get("Content-Type")))
				}
				gdk.ThreadsLeave()
				gdk.ThreadsEnter()
				buffer.Insert(&iter, " ")
				buffer.InsertWithTag(&iter, t.User.ScreenName, tag)
				buffer.Insert(&iter, ":"+t.Text+"\n")
				gtk.MainIterationDo(false)
				gdk.ThreadsLeave()
			}
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
