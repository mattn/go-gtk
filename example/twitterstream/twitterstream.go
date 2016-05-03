package main

import (
	"bufio"
	"compress/gzip"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"runtime"
	"strings"
	"sync"

	"github.com/garyburd/go-oauth/oauth"
	"github.com/mattn/go-gtk/gdkpixbuf"
	"github.com/mattn/go-gtk/gtk"
)

var (
	alive = true
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

func bytes2pixbuf(data []byte, typ string) *gdkpixbuf.Pixbuf {
	var loader *gdkpixbuf.Loader
	if strings.Index(typ, "jpeg") >= 0 {
		loader, _ = gdkpixbuf.NewLoaderWithMimeType("image/jpeg")
	} else {
		loader, _ = gdkpixbuf.NewLoaderWithMimeType("image/png")
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
		} `json:"user_mentions"`
		Urls []struct {
			Indices     [2]int
			Url         string
			DisplayUrl  string  `json:"display_url"`
			ExpandedUrl *string `json:"expanded_url"`
		}
	}
}

func stream(client *oauth.Client, cred *oauth.Credentials, f func(*tweet)) {
	param := make(url.Values)
	uri := "https://userstream.twitter.com/1.1/user.json"
	client.SignParam(cred, "GET", uri, param)
	uri = uri + "?" + param.Encode()
	res, err := http.Get(uri)
	if err != nil {
		log.Fatal("failed to get tweets:", err)
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		log.Fatal("failed to get tweets:", err)
	}
	var buf *bufio.Reader
	if res.Header.Get("Content-Encoding") == "gzip" {
		gr, err := gzip.NewReader(res.Body)
		if err != nil {
			log.Fatal("failed to make gzip decoder:", err)
		}
		buf = bufio.NewReader(gr)
	} else {
		buf = bufio.NewReader(res.Body)
	}
	var last []byte
	for alive {
		b, _, err := buf.ReadLine()
		last = append(last, b...)
		var t tweet
		err = json.Unmarshal(last, &t)
		if err != nil {
			continue
		}
		last = []byte{}
		if t.Text == "" {
			continue
		}
		f(&t)
	}
}

func post(client *oauth.Client, cred *oauth.Credentials, s string) {
	param := make(url.Values)
	param.Set("status", s)
	uri := "https://api.twitter.com/1.1/statuses/update.json"
	client.SignParam(cred, "POST", uri, param)
	res, err := http.PostForm(uri, url.Values(param))
	if err != nil {
		log.Println("failed to post tweet:", err)
		return
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		log.Println("failed to get timeline:", err)
		return
	}
}

func display(t *tweet, buffer *gtk.TextBuffer, tag *gtk.TextTag) {
	var iter gtk.TextIter
	pixbufbytes, resp := readURL(t.User.ProfileImageUrl)
	buffer.GetStartIter(&iter)
	if resp != nil {
		buffer.InsertPixbuf(&iter, bytes2pixbuf(pixbufbytes, resp.Header.Get("Content-Type")))
	}
	buffer.Insert(&iter, " ")
	buffer.InsertWithTag(&iter, t.User.ScreenName, tag)
	buffer.Insert(&iter, ":"+t.Text+"\n")
	gtk.MainIterationDo(false)
}

func show(client *oauth.Client, cred *oauth.Credentials, f func(t *tweet)) {
	param := make(url.Values)
	uri := "https://api.twitter.com/1.1/statuses/home_timeline.json"
	client.SignParam(cred, "GET", uri, param)
	uri = uri + "?" + param.Encode()
	res, err := http.Get(uri)
	if err != nil {
		return
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		return
	}
	var tweets []tweet
	json.NewDecoder(res.Body).Decode(&tweets)
	for _, t := range tweets {
		f(&t)
	}
}

func main() {
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
	client := &oauth.Client{Credentials: oauth.Credentials{config["ClientToken"], config["ClientSecret"]}}
	cred := &oauth.Credentials{config["AccessToken"], config["AccessSecret"]}

	runtime.LockOSThread()
	gtk.Init(&os.Args)
	window := gtk.NewWindow(gtk.WINDOW_TOPLEVEL)
	window.SetTitle("Twitter!")
	window.Connect("destroy", gtk.MainQuit)

	vbox := gtk.NewVBox(false, 1)

	scrolledwin := gtk.NewScrolledWindow(nil, nil)
	textview := gtk.NewTextView()
	textview.SetEditable(false)
	textview.SetCursorVisible(false)
	scrolledwin.Add(textview)
	vbox.Add(scrolledwin)

	buffer := textview.GetBuffer()
	tag := buffer.CreateTag("blue", map[string]string{"foreground": "#0000FF", "weight": "700"})

	hbox := gtk.NewHBox(false, 1)
	vbox.PackEnd(hbox, false, true, 5)

	label := gtk.NewLabel("Tweet")
	hbox.PackStart(label, false, false, 5)
	text := gtk.NewEntry()
	hbox.PackEnd(text, true, true, 5)

	text.Connect("activate", func() {
		t := text.GetText()
		text.SetText("")
		post(client, cred, t)
	})

	window.Add(vbox)
	window.SetSizeRequest(800, 500)
	window.ShowAll()

	var mutex sync.Mutex

	go func() {
		show(client, cred, func(t *tweet) {
			mutex.Lock()
			display(t, buffer, tag)
			mutex.Unlock()
		})

		stream(client, cred, func(t *tweet) {
			mutex.Lock()
			display(t, buffer, tag)
			var start, end gtk.TextIter
			buffer.GetIterAtLine(&start, buffer.GetLineCount()-2)
			buffer.GetEndIter(&end)
			buffer.Delete(&start, &end)
			mutex.Unlock()
		})
	}()

	for alive {
		mutex.Lock()
		alive = gtk.MainIterationDo(false)
		mutex.Unlock()
	}
}
