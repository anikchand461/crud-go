package handlers

import (
	"backend/storage"
	"fmt"
	"github.com/gin-gonic/gin"
)

func GetUser(c *gin.Context) {
	id := c.Param("id")
	// loop over the users 
	for _, user := range storage.Users {
		// if id found show the user details 
		if fmt.Sprint(user.ID) == id {
			c.JSON(200, user)
		}
		return
	}

	// error - user not found 
	c.JSON(400, gin.H{
		"error" : "id not found",
	})
}