package users

import (
	"fmt"
	"github.com/arun/Documents/workspaces/golang/src/github.com/kcarun/bookstore_users-api/utils/date_utils"
	"github.com/arun/Documents/workspaces/golang/src/github.com/kcarun/bookstore_users-api/utils/errors"
	"github.com/arun/Documents/workspaces/golang/src/github.com/kcarun/bookstore_users-api/datasources/mysql/users_db"

)

var usersDB = make(map[int64]*User)

func (user *User) Get() *errors.RestErr {
	if err := users_db.Client.Ping(); err != nil {
		panic(err)
	}

	result := usersDB[user.Id]
	if result == nil {
		return errors.NewNotFoundError(fmt.Sprintf("user %d not found ", user.Id))
	}
	user.Id = result.Id
	user.FirstName = result.FirstName
	user.LastName = result.LastName
	user.DateCreated = result.DateCreated
	user.Email = result.Email
	return nil
}
func (user *User) Save() *errors.RestErr {
	current := usersDB[user.Id]

	if current != nil {
		// if current.Email == user.Email |  {
		// 	return errors.NewBadRequestError(fmt.Sprintf("email %s already exists", user.Email))
		// }

		return errors.NewBadRequestError(fmt.Sprintf("userId %d already exists", user.Id))
	}
	//Ar logic to avoid duplicate mail id
	for id, usr := range usersDB {
		if usr.Email == user.Email {
			return errors.NewBadRequestError(fmt.Sprintf("email %s already registred with user id %v", user.Email, id))
		}
	}

	user.DateCreated = date_utils.GetNowString()
	usersDB[user.Id] = user
	fmt.Println(usersDB) //debugCode
	return nil
}
