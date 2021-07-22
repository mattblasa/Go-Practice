package main

import "fmt"

//Declare array
var array = []int{3, 5, -4, 8, 11, 1, -1, 6}
var target int = 10

func main() {
	for i := 0; i < len(array)-1; i++ {
		//This prints each number in the array
		firstNum := array[i]
		fmt.Println(firstNum)
	}
}

