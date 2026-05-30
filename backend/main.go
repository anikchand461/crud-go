package main

import (
	"backend/routes"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()

	r.Use(cors.Default())

	routes.UserRoutes(r)

	port := os.Getenv("PORT")

	if port == "" {
		port = "8080"
	}

	r.Run(":" + port)
}