package main

import "testing"

func main() {
}

func IsPalidrome(str string) {
}

func BenchmarkIsPalidrome(b *testing.B) {
	for i := 0; i < b.N; i++ {
		IsPalidrome("A man, a plan, a canal: Panama")
	}
}
