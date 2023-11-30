package main

import "fmt"

type Person struct {
	name string
	age  int
}

func main() {
	var mike Person
	mike.name = "Mike"
	mike.age = 23

	var bob = Person{"Bob", 35}

	var sam = Person{age: 89, name: "Sam"}

	fmt.Println(mike, bob, sam)
}
