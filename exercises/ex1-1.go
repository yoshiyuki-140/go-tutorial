package main

import (
	"fmt"
	"os"
	"strings"
)

// echo プログラム

func main() {
	fmt.Println(strings.Join(os.Args, " "))
}
