package api

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func Serve(apiToken, message string, port uint16) {
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()

	router.PUT("/data", apiAuthHandler(apiToken), putData)
	router.GET("/data", getData)
	router.GET("/message", getMessage(message))

	err := router.Run(":" + fmt.Sprintf("%d", port))
	if err != nil {
		panic(err)
	}
}
