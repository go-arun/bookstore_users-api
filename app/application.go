package app

import(  "fmt"
	"github.com/gin-gonic/gin"
)

var (
	router = gin.Default()
)

//StartApplicaiton will call from main
func StartApplicaton() {
	mapUrls()
	router.Run(":8081")
	fmt.Println("Calling okay!!!")
}
