package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/memmre/GoSampleAPI/middlewares"
)

func RegisterRoutes(server *gin.Engine) {
	server.POST("signIn", signIn)
	server.POST("signUp", signUp)
	server.GET("events/:id", getEvent) // Get an event
	server.GET("events", getEvents)    // Get all events

	// Authenticated Routes
	authenticated := server.Group("/")
	authenticated.Use(middlewares.Authenticate)
	authenticated.POST("events", createEvent)                       // Create an event
	authenticated.PUT("events/:id", updateEvent)                    // Update an event
	authenticated.DELETE("events/:id", deleteEvent)                 // Delete an event
	authenticated.POST("events/:id/register", createRegistration)   // Create a registration
	authenticated.DELETE("events/:id/register", deleteRegistration) // Delete a registration
}
