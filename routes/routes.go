package routes

import "github.com/gin-gonic/gin"

func RegisterRoutes(server *gin.Engine) {
	server.GET("events/:id", getEvent)       // Get an event
	server.GET("events", getEvents)          // Get all events
	server.POST("events", createEvent)       // Create an event
	server.PUT("events/:id", updateEvent)    // Update an event
	server.DELETE("events/:id", deleteEvent) // Delete an event
}
