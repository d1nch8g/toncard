package main

import (
	"fmt"
	"log"

	"golang.org/x/mobile/app"
	"golang.org/x/mobile/event/lifecycle"
	"golang.org/x/mobile/event/paint"
)

func main() {
	app.Main(func(a app.App) {
		for e := range a.Events() {
			switch e := a.Filter(e).(type) {
			case lifecycle.Event:
				fmt.Println(e.String())
			case paint.Event:
				log.Print("Call OpenGL here.")
				a.Publish()
			}
		}
	})
}
