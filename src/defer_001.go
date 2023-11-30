package main

import "fmt"

func main() {
	defer fmt.Println("Golang")
	defer fmt.Println("Ruby")
	fmt.Println("JS")
}
