package main

import(
	"fmt"
)

func main() {
	fmt.Println("Yo killa.")
	port := 2000
	startWebServer(port,2)
}

func startWebServer(port int, numberOfRetries int) {
	fmt.Println("Starting server...")
	//Stuff here
	fmt.Println("Server started")
}