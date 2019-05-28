package main

import (
	"github.com/ElegantCreationism/go-hoover/service"
	"log"
	"runtime/debug"
)

func main() {
	defer func() {
		if r := recover(); r != nil {
			log.Printf("%s: %s", r, string(debug.Stack()))
			log.Fatal(r)
		}
	}()

	service.Start()
}
