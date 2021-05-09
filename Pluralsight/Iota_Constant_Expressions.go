package main

import "fmt"

//Create Constant outside of the main function

//Constant isn't limited to the main function
const (
	first = iota
	second = iota
)
/*
Iota - every time iota is used, it increments its value by 1. Meaning you can build long chains of constants using
iota.

 */
func main() {
	fmt.Println(first, second)
}


