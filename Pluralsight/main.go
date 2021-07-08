package main

import (
	"fmt"
	"net/http"

	"github.com/pluralsight/webservice/models"
)

func main() {
	controllers.RegisterController()
	//IP address we are listening on
	http.ListenAndServe(":3000", nil)
}
