package models

type User struct {
	ID			int
	FirstName 	string
	LastName 	string
}

var (
	users []*User //This is a pointer
	nextID int32 = 1

)


func GetUsers() []*User {
	/*
	   Returns Users from the user structure

	   Parameters:
	   None

	   Returns:
	   User Structure
	*/
	return users
}