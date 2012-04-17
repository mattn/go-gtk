package main

import (
	"fmt"
	"github.com/mattn/go-gtk/glib"
	"time"
)

func main() {
	glib.IdleAdd(func() bool {
		println("start")
		return false
	})
	glib.TimeoutAdd(1000, func() bool {
		println(fmt.Sprintf("%v", time.Now()))
		return true
	})

	glib.MainLoopNew(nil, false).Run()
}
