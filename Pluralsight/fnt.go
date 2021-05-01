package main

import (
	"fmt"
)

func main() {
	var i int
	i = 42
	fmt.Println(i)

	var f float32 = 3.14
	fmt.Println(f)

	firstName := "Arthur"
	fmt.Println(firstName)

	b := true
	fmt.Println(b)

	//Store imaginary numbers using the complex method
	c := complex(3,4)
	fmt.Println(c)

	//Declare real and imaginary values of variable c 
	r, im := real(c), imag(c)
	fmt.Println(r, im)
}

