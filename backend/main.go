package main

import (
	"backend/routes"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.Use(cors.Default())
	routes.UserRoutes(r)
	r.Run(":8080")
}