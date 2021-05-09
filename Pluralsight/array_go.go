package main

import (
	"fmt"
)

func main() {
	/*
	In Go, we declare the array, the size and types. Unless you specify the data type,
	an array cant be intialized

	If you reach the end of the array, you'll get a compile time error. So arr[4] will not run in this case. 
	 */
	// variable, with the size of the array and the data type
	var arr [3]int
	arr[0] = 1
	arr[1] = 2
	arr[3] = 3
	fmt.Println(arr)
}