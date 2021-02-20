package main

import (
	"fmt"
)

func main() {
	var i int = 60
	fmt.Printf("%v, %T\n", i, i)

	var j float32 = 180
	//j = float32(i)
	fmt.Printf("%v, %T\n", j, j)
}