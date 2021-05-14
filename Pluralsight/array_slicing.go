package main

import (
	"fmt"
)

func main() {
	arr := [3]int{1, 2, 3} //This sits in memory

	slice := arr[:] // the colon in the middle means that I want slice from the beginning of the connection to the end
					// Slice does not sit in memory

	arr[1] = 42
	slice[2] = 27 // Slice is pointing to the array, which is a value type. So changes mirror each other

	fmt.Println(arr, slice)
}
