package app

import (
	"fmt"
	"os"

	"coba-go/database/connection"

	"github.com/gin-gonic/gin"
)

var (
	router = gin.Default()
)

func StartApplication() {

	mapUrls()

	connection.Connect()

	port := os.Getenv("APP_PORT")
	fmt.Printf("running at port: %s", port)
	router.Run(port)
}
