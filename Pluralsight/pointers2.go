package main

import "fmt"

func main() {
	firstName := "Arthur"
	fmt.Println(firstName)

	//pointer variable here points to the firstname &
	ptr:= &firstName
	fmt.Println(ptr, *ptr) //prints ptr variable, and the value the pointer stores

	firstName = "Tricia"
	fmt.Println(ptr, *ptr)
}

/*
Arthur
<pointer address> Arthur
<pointer address> Tricia
 */
