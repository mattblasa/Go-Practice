package main

import (
	"fmt"
)

func main() {
	//data type here is a string pointer data type. new method initalizes it to string data type
	var firstName *string = new(string)

	//0xC0000022 is the memory address for the pointer

	//de-reference a pointer
	*firstName = "Arthur"

	fmt.Println(firstName)
}

