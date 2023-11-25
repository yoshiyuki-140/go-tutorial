package main

import "fmt"

func main() {
	a := []int{10, 20}
	fmt.Println(a, cap(a))

	b := append(a, 30)
	a[0] = 100
	// [10 20 30] 4
	fmt.Println(b, cap(b))

	c := append(b, 40)
	b[1] = 200
	fmt.Println(c, cap(c))
}
