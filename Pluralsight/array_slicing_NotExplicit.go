package main

import "fmt"

func main() {
	slice := []int{1, 2, 3} //This creates the array, and slices at the same time.
	fmt.Println(slice) //slice is not a fixed sized entity.

	slice = append(slice, 4, 96, 22, 33) //Go manages the resizing of the array
	fmt.Println(slice) //slice is not a fixed sized entity.

	s2 := slice[1:] //slice index 1
	s3 := slice[:2]
	s4 := slice[1:2]

	fmt.Println(s2, s3, s4)

}

