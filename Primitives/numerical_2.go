package main

import (
	"fmt"
)

/*
Bitshifting is occurring here.
*/
func main() {
	//short declaration operator. This stores 8 in a variable, without declaring it
	a := 8              // 8^3
	fmt.Println(a << 3) //n times 2, x times
	fmt.Println(a >> 3) // "y divided by 2, z times"
}
