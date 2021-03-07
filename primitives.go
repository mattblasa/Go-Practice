package main

import (
	"fmt"
)

func main() {
	//var defines a variable, n is the name, and bool is type
	var n bool = true
	//string is the template. Prinf means print format.
	//%v means default format, %T means type of value, \n means U+000A line feed or newline
	fmt.Printf("%v, %T\n", n, n)
}
