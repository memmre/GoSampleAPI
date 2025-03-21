package main

import (
	"github.com/gin-gonic/gin"
	"github.com/memmre/GoSampleAPI/database"
	"github.com/memmre/GoSampleAPI/routes"
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
