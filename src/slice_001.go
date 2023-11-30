package main

import "fmt"

func main() {
	// 参照用配列
	var array [2]string = [2]string{"golang", "python"}

	// slice
	var slice1 []string = []string{array[0], array[1]}

	fmt.Println(slice1)

}
