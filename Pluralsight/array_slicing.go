package main

import (
	"fmt"
)

func main() {
	arr := [3]int{1, 2, 3}

	slice := arr[:] // the colon in the middle means that I want slice from the beginning of the connection to the end

	fmt.Println(arr, slice)
}
