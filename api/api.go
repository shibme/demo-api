package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

var data = ""

type PutDataRequest struct {
	Data string `json:"data"`
}

func putData(context *gin.Context) {
	var request PutDataRequest
	err := context.BindJSON(&request)
	if err != nil {
		context.AbortWithStatus(http.StatusBadRequest)
		return
	}
	data = request.Data
	context.JSON(http.StatusOK, gin.H{"data": data})
}

func getData(context *gin.Context) {
	if data == "" {
		context.AbortWithStatus(http.StatusNotFound)
		return
	}
	context.JSON(http.StatusOK, gin.H{"data": data})
}

func getMessage(message string) gin.HandlerFunc {
	return func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{"message": message})
	}
}
