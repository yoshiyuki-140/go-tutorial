package main

import (
	"fmt"
	"log"
	"runtime"
)

func hello() {
	fmt.Println("Hello")
}
func goodMorning() {
	fmt.Println("good morning")
}

func main() {
	go hello()
	go goodMorning()

	log.Println(runtime.NumGoroutine())
}
