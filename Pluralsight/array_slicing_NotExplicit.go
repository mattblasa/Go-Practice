package main

import "fmt"

func main() {
	slice := []int{1, 2, 3} //This creates the array, and slices at the same time.
	fmt.Println(slice) //slice is not a fixed sized entity.

	slice = append(slice, 4, 96, 22, 33)
	fmt.Println(slice) //slice is not a fixed sized entity.
}

