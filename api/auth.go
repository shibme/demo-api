package api

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func apiAuthHandler(apiToken string) gin.HandlerFunc {
	return func(context *gin.Context) {
		authHeader := context.GetHeader("Authorization")
		if authHeader == "" {
			context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Authorization header is missing"})
			return
		}
		authHeaderValueParts := strings.Split(authHeader, " ")
		if len(authHeaderValueParts) != 2 || strings.ToLower(authHeaderValueParts[0]) != "bearer" {
			context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid Authorization header format"})
			return
		}
		if authHeaderValueParts[1] != apiToken {
			context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid Token"})
			return
		}
	}
}
