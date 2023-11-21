package main

import "fmt"

func main() {
	a := [5]int{2, 101, 29, 123, 4}
	fmt.Println(a[2:4])
	fmt.Println(a[2:])
	fmt.Println(a[:4])
	fmt.Println(a[:])

	fmt.Println(len(a))
	fmt.Println(cap(a))

}
