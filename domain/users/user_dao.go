package users

import (


	"github.com/arun/Documents/workspaces/golang/src/github.com/kcarun/bookstore_users-api/datasources/mysql/users_db"
	"github.com/arun/Documents/workspaces/golang/src/github.com/kcarun/bookstore_users-api/utils/date_utils"
	"github.com/arun/Documents/workspaces/golang/src/github.com/kcarun/bookstore_users-api/utils/errors"
	"github.com/arun/Documents/workspaces/golang/src/github.com/kcarun/bookstore_users-api/utils/mysql_utils"
)

const (
	queryInsertUser  = "INSERT INTO users(first_name,last_name,email,date_created) VALUES(?,?,?,?);"
	queryGetUser     = "SELECT id,first_name,last_name,email,date_created FROM users WHERE id=?"
	indexUniqueEmail = "email_UNIQUE"
	noRowsInResult   = "no rows in result set"
)

//Get function
func (user *User) Get() *errors.RestErr {

	if err := users_db.Client.Ping(); err != nil {
		panic(err)
	}

	stmt, err := users_db.Client.Prepare(queryGetUser)
	if err != nil {
		return mysql_utils.ParseError(err)
	}
	defer stmt.Close()

	result := stmt.QueryRow(user.Id)
	if err := result.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.DateCreated); err != nil {
		return mysql_utils.ParseError(err)

	}

	return nil
}

//Save function
func (user *User) Save() *errors.RestErr {
	stmt, err := users_db.Client.Prepare(queryInsertUser)

	//fmt.Println("err.Error():", err.Error())
	if err != nil {
		return mysql_utils.ParseError(err)
	}
	defer stmt.Close()
	user.DateCreated = date_utils.GetNowString()
	insertResult, saveErr := stmt.Exec(
		user.FirstName,
		user.LastName,
		user.Email,
		user.DateCreated,
	)
	if saveErr != nil {

		return mysql_utils.ParseError(saveErr)

	}
	userId, err := insertResult.LastInsertId()
	if err != nil {
		return mysql_utils.ParseError(err)
	}
	user.Id = userId
	return nil
}
