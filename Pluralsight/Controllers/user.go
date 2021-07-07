package Controllers

import "net/http"

//Custom data type userController
type userController struct{

}

//method. userController is the type for uc 
func (uc userController) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello from Plano!") //Convert string to byte slice
}