package handlers

import (
	"github.com/gin-gonic/gin"
	"backend/storage"
	"fmt"
	"backend/models"
)

func UpdateUser(c *gin.Context) {
	// get the id of the user form the path params
	id := c.Param("id")

	// create a user tye variable
	var updateuser models.User

	// jsonify and error handling
	err := c.BindJSON(&updateuser)
	if err != nil {
		c.JSON(400, gin.H{
			"error" : "wrong json format",
		})
		return
	}

	// loop over the users
	for index, user := range storage.Users {
		// if the id found then update the user details and return
		if fmt.Sprint(user.ID) == id {
			storage.Users[index].Name = updateuser.Name
			storage.Users[index].Age = updateuser.Age
			c.JSON(200, storage.Users[index])

			return
		}
	}

	c.JSON(400, gin.H{
		"error" : "user not found",
	})
}
