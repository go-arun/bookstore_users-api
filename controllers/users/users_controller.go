package users

import (
	"net/http"

	//"encoding/json"
	//"fmt"
	//"io/ioutil"

	"github.com/arun/Documents/workspaces/golang/src/github.com/kcarun/bookstore_users-api/domain/users"
	"github.com/arun/Documents/workspaces/golang/src/github.com/kcarun/bookstore_users-api/services"
	errors "github.com/arun/Documents/workspaces/golang/src/github.com/kcarun/bookstore_users-api/utils"
	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
	var user users.User
	if err := c.ShouldBindJSON(&user); err != nil {
		restErr := errors.NewBadRequestError("invalid json body")
		c.JSON(restErr.Status, restErr)
		return
	}

	result, saveErr := services.CreateUser(user)
	if saveErr != nil {
		c.JSON(saveErr.Status, saveErr)
		return
	}
	c.JSON(http.StatusCreated, result)

}

//GetUser comment-todo
func GetUser(c *gin.Context) {
	c.String(http.StatusNotImplemented, "implement me!")
}

// func SearchUser(c *gin.Context) {
// 	c.String(http.StatusNotImplemented, "implement me!")
// }
