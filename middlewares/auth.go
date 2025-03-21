package middlewares

import (
	"github.com/gin-gonic/gin"
	"github.com/memmre/GoSampleAPI/utilities"
	"net/http"
)

func Authenticate(context *gin.Context) {
	token := context.GetHeader("Authorization")
	if token == "" {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Not authorized."})
		return
	}

	userID, err := utilities.VerifyToken(token)
	if err != nil {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Not authorized."})
		return
	}

	context.Set("userID", userID)
	context.Next()
}
