package main

import (
	"fmt"
	"os"
)

func main() {
	var (
		s string
		i int
	)

	sep := ","

	for i < len(os.Args) {
		i++
		s += "[" + sep + os.Args[i] + "]"
		s += "\n"
	}
	fmt.Println(s)
}
