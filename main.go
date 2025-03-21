package main

import (
	"example.com/sample-api/database"
	"example.com/sample-api/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	database.InitializeDatabase()
	server := gin.Default()
	routes.RegisterRoutes(server)

	err := server.Run(":8080") // HOST:PORT
	if err != nil {
		return
	}
}
