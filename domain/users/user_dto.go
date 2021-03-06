package users

import (
	"strings"
	"github.com/arun/Documents/workspaces/golang/src/github.com/kcarun/bookstore_users-api/utils/errors"
)

//User strcut
type User struct {
	Id          int64  `json:"id"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Email       string `json:"email"`
	DateCreated string `json:"data_created`
}

func (user *User) Validate() *errors.RestErr {
	user.Email = strings.TrimSpace(strings.ToLower(user.Email))
	if user.Email == "" {
		return errors.NewBadRequestError("Invalid email address")
	}
	return nil
}
