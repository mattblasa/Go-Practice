//API Test

package models

type User struct {
	ID			int
	FirstName 	string
	LastName 	string
}

var (
	users []*User //This is a pointer
	nextID = 1

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

func AddUser(u User) (User, error) { //returns user, otherwise Error
	u.ID = nextID
	nextID++
	users = append(users, &u) //&u is a reference to operator. Its a pointer (Explain this better)
	return u, nil
}