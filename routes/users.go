package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/memmre/GoSampleAPI/models"
	"github.com/memmre/GoSampleAPI/utilities"
	"net/http"
)

func signIn(context *gin.Context) {
	var user models.User

	err := context.ShouldBindJSON(&user)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse body."})
		return
	}

	err = user.ValidateCredentials()
	if err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "Could not authenticate user."})
		return
	}

	token, err := utilities.GenerateToken(user.ID, user.EmailAddress)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not authenticate user."})
		return
	}

	context.JSON(
		http.StatusOK,
		gin.H{
			"message": "Sign in successful.",
			"token":   token,
		},
	)
}

func signUp(context *gin.Context) {
	var user models.User

	err := context.ShouldBindJSON(&user)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse body."})
		return
	}

	err = user.Create()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not create user."})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "User created successfully."})
}
