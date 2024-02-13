package main

import (
	"example.com/event-booking-app/db"
	"example.com/event-booking-app/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB()
	server := gin.Default()

	routes.RegisterRoutes(server) // moved routes to a separate package and then used them here through this

	server.Run(":8080") //localhost:8080
}

