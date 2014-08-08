package main

import (
	"fmt"
	"github.com/mattn/go-gtk/glib"
	"time"
)

func main() {
	glib.IdleAdd(func() bool {
		fmt.Println("start")
		return false
	})
	glib.TimeoutAdd(1000, func() bool {
		fmt.Println(fmt.Sprintf("%v", time.Now()))
		return true
	})

	glib.NewMainLoop(nil, false).Run()
}
