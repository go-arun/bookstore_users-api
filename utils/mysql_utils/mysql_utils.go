package mysql_utils

import (
	"strings"

	"github.com/arun/Documents/workspaces/golang/src/github.com/kcarun/bookstore_users-api/utils/errors"
	"github.com/go-sql-driver/mysql"
)

const (
	noRowsInResult = "no rows in result set"
)

//ParseError func
func ParseError(err error) *errors.RestErr {
	sqlErr, ok := err.(*mysql.MySQLError)
	if !ok {
		if strings.Contains(err.Error(), noRowsInResult) {
			return errors.NewNotFoundError("No Matching Records !!!")
		}
		return errors.NewInternalServerError("Error Parsing SQL Request!!!")
	}
	switch sqlErr.Number {
	case 1062:
		{
			return errors.NewBadRequestError("This mail id %s alrady exists")
		}
	}
	return errors.NewInternalServerError("Error processing request")
}
