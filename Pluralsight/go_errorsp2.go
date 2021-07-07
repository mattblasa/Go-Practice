package main

import "fmt"
func main() {
	port := 2000
	port, err := startWebServer(port)
	fmt.Println(err)
}

func startWebServer(port int) (int, error) { //specify the return data
	fmt.Println("Starting server...")
	//Stuff here
	fmt.Println("Server started on port", port)
	return port, nil
}


