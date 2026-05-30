package handlers

import (
	"fmt"
	"backend/storage"
	"github.com/gin-gonic/gin"
)

func DeleteUser(c *gin.Context) {
	id := c.Param("id")

	// loop over the users 
	for index, user := range storage.Users {
		// if found the id .. delete the user and return 
		if fmt.Sprint(user.ID) == id {
			storage.Users = append(storage.Users[:index], storage.Users[index+1:]...)
		}
		c.JSON(200, gin.H{
			"messege" : "user successfully deleted",
		})
		return
	}

	c.JSON(400, gin.H{
		"error" : "user not found",
	})
}