package users

import (
	"fmt"

	"github.com/arun/Documents/workspaces/golang/src/github.com/kcarun/bookstore_users-api/utils"
)

var usersDB = make(map[int64]*User)

func (user *User) Get() *errors.RestErr {
	result := usersDB[user.Id]
	if result == nil {
		return errors.NewNotFoundError(fmt.Sprintf("user %d not found ", user.Id))
	}
	user.Id = result.Id
	user.FirstName = result.FirstName
	user.LastName = result.LastName
	user.DateCreated = result.DateCreated
	return nil
}
func (user *User) Save() *errors.RestErr {
	current := usersDB[user.Id]
	if current != nil {
		if current.Email == user.Email {
			return errors.NewBadRequestError(fmt.Sprintf("email % id already exists", user.Email))
		}

		return errors.NewBadRequestError(fmt.Sprintf("user % id already exists", user.Id))
	}
	usersDB[user.Id] = user
	fmt.Println(usersDB) //debugCode
	return nil
}
