package main

import (
	"fmt"
)

func main() {
	a := 10            // 1010
	b := 3             // 0011
	fmt.Println(a & b) // 0010 = 2
	fmt.Println(a | b) // OR 1011 =11
	fmt.Println(a ^ b) //
	fmt.Println(a &^ b)
}
