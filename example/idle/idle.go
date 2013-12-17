package main

import (
	"fmt"
	"github.com/d2r2/go-gtk/glib"
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

	glib.MainLoopNew(nil, false).Run()
}
