package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gofrs/uuid"
	"gitlab.com/mrbrownt/gitreleased.app/backend/models"
)

// UserHandler propigates routes for users and whatnot
func UserHandler(r *gin.RouterGroup) {
	r.Use(authMiddleware())
	r.GET("/", getUser)
}

func getUser(c *gin.Context) {
	id, exist := c.Get("id")
	if !exist {
		// Not sure if this is what we want to do here
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	user := models.User{}

	db := models.GetDB()
	db.Where("id = ?", id.(string)).First(&user)
	if user.ID == uuid.Nil {
		// Something really shitty has happened if you hit this!
		c.JSON(http.StatusNotFound, gin.H{"error": "record not found"})
		return
	}

	c.JSON(http.StatusOK, user)
}
