package main

import (
	"fmt"
)

func main() {
	//This is a short variable declaration
	n := 1 == 1
	m := 1 == 2
	//\n makes a line skip 
	fmt.Printf("%v, %T\n", n, n )
	fmt.Printf("%v, %T", m, m)
}