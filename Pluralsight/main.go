package main

import (
	"net/http"
)

func main() {
	controllers.RegisterController()
	//IP address we are listening on
	http.ListenAndServe(":3000", nil)
}
