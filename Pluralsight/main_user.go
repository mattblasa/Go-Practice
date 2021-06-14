package main

import (
	"github.com/pluralsight/webservice/models"
	"fmt"
)

func main() {
	u := models.User{
		ID : 2,
		FirstName: "Tricia",
		LastName: "McMillan",
	} //This imports from the github
	fmt.Println(u)
}
