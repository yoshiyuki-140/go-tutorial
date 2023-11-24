package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func good_version() {
	fmt.Println(strings.Join(os.Args, " "))

}
func bad_version() {
	var s string
	for i := 0; i < len(os.Args); i++ {
		s += os.Args[i] + " "
	}
	fmt.Println(s)
}

func main() {
	// good version
	start := time.Now()
	good_version()
	fmt.Println(time.Since(start).Seconds())

	// bad version
	start = time.Now()
	bad_version()
	fmt.Println(time.Since(start).Seconds())
}
