package app

import (
	"github.com/gin-gonic/gin"
)

var (
	router = gin.Default()
)

//StartApplicaton will call from main
func StartApplicaton() {
	mapUrls()
	router.Run(":8080")
}
