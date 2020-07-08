package main

import (
	"fmt"

	"github.com/arun/Documents/workspaces/golang/src/github.com/kcarun/bookstore_users-api/app"
)

func main() {
	fmt.Println("Message for checking , whether init func in database package runs before main..")
	app.StartApplicaton()
}
