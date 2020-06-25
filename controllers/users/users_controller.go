package users

import (
	"fmt"
	"net/http"
	"strconv"

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
	//c.String(http.StatusNotImplemented, "implement me!")
	userId, userErr := strconv.ParseInt(c.Param("user_id"), 10, 64)
	fmt.Println(userId,c.Param("first_name")) //TestCode
	if userErr != nil {
		err := errors.NewBadRequestError("invalid user id- Should be a nummber")
		c.JSON(err.Status, err)
		return
	}

	user, getErr := services.GetUser(userId)
	if getErr != nil {
		c.JSON(getErr.Status, getErr)
		return
	}
	c.JSON(http.StatusOK, user)

}

