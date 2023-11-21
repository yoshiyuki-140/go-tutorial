package main

import "fmt"

func main() {

	// 宣言と代入を一緒にするパターン
	msg := "Super Hello"

	var count int = 100
	for i := 1; i < count; i++ {
		fmt.Println(msg)
	}

	var a, b int
	a, b = 10, 15

	// 複数の型違いの変数を同時に定義
	var (
		c int
		d string
	)

	c = 20
	d = "hoge"

	var (
		e = 20
		f = "hoge"
	)
}
