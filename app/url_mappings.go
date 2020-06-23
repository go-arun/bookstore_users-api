package app

import (
	"github.com/arun/Documents/workspaces/golang/src/github.com/kcarun/bookstore_users-api/controllers/ping"
	"github.com/arun/Documents/workspaces/golang/src/github.com/kcarun/bookstore_users-api/controllers/users"
)

func mapUrls() {
	router.GET("/ping", ping.Ping)
	router.GET("/users/:user_id", users.GetUser)
	router.POST("/users", users.CreateUser)
	//router.POST("/users/search", controllers.SearchUser)

}
