package services

import (
	"github.com/arun/Documents/workspaces/golang/src/github.com/kcarun/bookstore_users-api/domain/users"
	"github.com/arun/Documents/workspaces/golang/src/github.com/kcarun/bookstore_users-api/utils"
)

func CreateUser(user users.User) (*users.User, *errors.RestErr) {
	if err := user.Validate(); err != nil {
		return nil, err
	}

	return nil, nil
}
