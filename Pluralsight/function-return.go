package main

import(
	"fmt"
	"errors"
)

func main() {
	fmt.Println("Yo killa.")
	port := 2000
	isStarted := startWebServer(port)
	fmt.Println(isStarted)
}

/*
//Return True
func startWebServer(port int) bool { //specify the return data
	fmt.Println("Starting server...")
	//Stuff here
	fmt.Println("Server started on port", port)
	return true //return it if the value is true
}

 */
 //Error Type
func startWebServer(port int) error { //specify the return data
	fmt.Println("Starting server...")
	//Stuff here
	fmt.Println("Server started on port", port)
	return nil //return it if the value is true
}
