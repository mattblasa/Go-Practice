package Controllers

import "net/http"

func RegisterControllers() {
	uc := newUserController()

	//handle requests through the /users route
	http.Handle("/users", *uc) //*uc is the user controller object. its pointing
	http.Handle("/users/", *uc)
}