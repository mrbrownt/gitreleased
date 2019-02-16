package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gitlab.com/mrbrownt/gitreleased.app/backend/models"
)

// UserHandler propigates routes for users and whatnot
func UserHandler(r *gin.RouterGroup) {
	r.POST("/", create)
	r.GET("/", get)
}

// UserInput for new users
type UserInput struct {
	Email    string `json:"email"`
	GithubID string `json:"github_id"`
}

func create(c *gin.Context) {
	var json models.User
	db := models.GetDB()

	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := db.Create(&json).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	c.JSON(http.StatusOK, json)
}

func get(c *gin.Context) {
	id := c.Query("id")
	var user UserInput

	user.Email = "someid" + id
	c.JSON(http.StatusOK, user)
}
