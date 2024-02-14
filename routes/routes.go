package routes

import (
	"example.com/event-booking-app/middlewares"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(server *gin.Engine) {
	server.GET("/events", getEvents)
	server.GET("/events/:id", getEvent)

	// middlewares can be attached to grouped routes like this
	authenticated := server.Group("/")
	authenticated.Use(middlewares.Authenticate)
	authenticated.POST("/events", middlewares.Authenticate, createEvent)
	authenticated.PUT("/events/:id", updateEvent)
	authenticated.DELETE("/events/:id", deleteEvent)

	// server.POST("/events", middlewares.Authenticate, createEvent) // middlewares can also be attached individually like this to routes
	server.POST("/signup", signup)
	server.POST("/login", login)
}