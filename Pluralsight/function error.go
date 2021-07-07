package main

import (
	"fmt"
	"errors"
)
func main() {
	fmt.Println("Yo killa.")
	port := 2000
	err := startWebServer(port)
	fmt.Println(err)
}

func startWebServer(port int) error { //specify the return data
	fmt.Println("Starting server...")
	//Stuff here
	fmt.Println("Server started on port", port)
	return errors.New("Something went wrong") //return it if the value is true
}

