package main

import (
	"fmt"
)

func main() {
	type user struct {
		ID int
		FirstName string
		LastName string
	}
	var u user  //Initalization syntax of an object
	u.ID = 1
	u.FirstName = "Arthur"
	u.LastName = "Dent"
	fmt.Println(u.FirstName)
}