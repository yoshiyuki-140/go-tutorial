package main

import (
	"fmt"
)

func main() {
	var array1 [2]string
	var array2 [2]string = [2]string{"golang", "python"}
	var array3 = [...]string{"golang", "python"}

	fmt.Println(array1, array2, array3)
}
