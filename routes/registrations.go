package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/memmre/GoSampleAPI/models"
	"net/http"
	"strconv"
)

func createRegistration(context *gin.Context) {
	userID := context.GetInt64("userID")
	eventID, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse event id."})
		return
	}

	event, err := models.GetEvent(eventID)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch event."})
		return
	}

	err = event.CreateRegistration(userID)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not create registration."})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "Registration created successfully."})
}

func deleteRegistration(context *gin.Context) {
	userID := context.GetInt64("userID")
	eventID, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse event id."})
		return
	}

	var event models.Event
	event.ID = eventID

	err = event.DeleteRegistration(userID)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not delete registration."})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Registration deleted successfully."})
}
