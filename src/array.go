package main

import "fmt"

func main() {
	var a [...]int
	count := 100
	for i := 0; i < count; i++ {
		a[i] = i
	}
	for i := 0; i < count; i++ {
		fmt.Println(a[i])
	}
}
