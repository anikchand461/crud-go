package handlers

import (
	"backend/storage"
	"github.com/gin-gonic/gin"
)

func GetUsers(c *gin.Context) {
	c.JSON(200, storage.Users)
}
