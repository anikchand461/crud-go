package routes

import (
	"backend/handlers"
	"github.com/gin-gonic/gin"
)

func UserRoutes(r *gin.Engine) {
	r.GET("/users", handlers.GetUsers)
	r.POST("/users", handlers.CreateUser)
	r.GET("/users/:id", handlers.GetUser)
	r.DELETE("/users/:id", handlers.DeleteUser)
	r.PUT("/users/:id", handlers.UpdateUser)
}