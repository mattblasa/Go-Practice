package main

import (
	"fmt"
	"strconv"
)

func main() {
	var i int = 60
	fmt.Printf("%v, %T\n", i, i)

	var k string
	k = strconv.Itoa(i)
	fmt.Printf("%v, %T\n", k, k)
}