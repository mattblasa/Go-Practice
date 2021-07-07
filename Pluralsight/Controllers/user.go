package Controllers

import (
	"net/http"
	"regexp"
)

//Custom data type userController
type userController struct{
	userIDPattern *regexp.Regexp
}

//method. userController is the type for uc 
func (uc userController) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello from Plano!") //Convert string to byte slice
}