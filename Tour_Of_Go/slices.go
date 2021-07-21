package main

import "fmt"

func main() {
	names := [4]string{ //creates an array of 4, in a string
		"John",
		"Paul",
		"George",
		"Ringo",
	}
	fmt.Println(names)

	a := names[0:2] //slice John index 0, and Paul index
	b := names[1:3]
	fmt.Println(a, b)

	b[0] = "XXX"
	fmt.Println(a, b)
	fmt.Println(names)
}