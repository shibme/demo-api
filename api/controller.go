package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

var data = ""

type PutDataRequest struct {
	Data string `json:"data"`
}

func PutData(context *gin.Context) {
	var request PutDataRequest
	err := context.BindJSON(&request)
	if err != nil {
		context.AbortWithStatus(http.StatusBadRequest)
		return
	}
	data = request.Data
	context.JSON(http.StatusOK, gin.H{"data": data})
}

func GetData(context *gin.Context) {
	if data == "" {
		context.AbortWithStatus(http.StatusNotFound)
		return
	}
	context.JSON(http.StatusOK, gin.H{"data": data})
}
