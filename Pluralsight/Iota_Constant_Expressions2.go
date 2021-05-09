package main

import "fmt"

//Create Constant outside of the main function

//Constant isn't limited to the main function
const (
	first = iota
	second
)
/*
Iota resets between cosntant blocks. Easier to use for different constant blocks.
 */

const (
	third = iota
	fourth
)




func main() {
	fmt.Println(first, second, third, fourth)
}


