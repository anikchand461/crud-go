package handlers

import (
	"backend/models"
	"backend/storage"

	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {

	var user models.User

	err := c.BindJSON(&user)

	if err != nil {

		c.JSON(400, gin.H{
			"error": "invalid json",
		})

		return
	}

	storage.Users = append(storage.Users, user)

	c.JSON(201, user)
}