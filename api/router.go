package api

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func Serve(apiToken string, port uint16) {
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()

	router.PUT("/data", apiAuthHandler(apiToken), PutData)
	router.GET("/data", GetData)

	err := router.Run(":" + fmt.Sprintf("%d", port))
	if err != nil {
		panic(err)
	}
}
