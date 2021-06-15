package main

import(
	"fmt"
)

func main() {
	fmt.Println("Yo killa.")
	port := 2000
	startWebServer(port)
}

func startWebServer(port int) {
	fmt.Println("Starting server...")
	//Stuff here
	fmt.Println("Server started")
}