package handlers

import (
	"net/http"

	"github.com/gofrs/uuid"
	"gitlab.com/mrbrownt/gitreleased.app/backend/config"
	"gitlab.com/mrbrownt/gitreleased.app/backend/models"

	"github.com/gin-gonic/gin"
)

// RepoHandler routes repo requests
func RepoHandler(r *gin.RouterGroup) {
	r.Use(authMiddleware())
	r.GET("/:id", getRepo)
}

func getRepo(c *gin.Context) {
	repoParam := c.Param("id")
	if repoParam == "" {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": "no query param set"})
		return
	}

	conf := config.Get()

	repoUUID, err := uuid.FromString(repoParam)
	if err != nil {
		c.AbortWithStatusJSON(
			http.StatusBadRequest,
			gin.H{"error": "UUID not formatted correctly"},
		)
		return
	}

	repo := models.Repository{}
	err = conf.DB.Where(models.Repository{ID: repoUUID}).First(&repo).Error
	if err != nil {
		c.AbortWithStatusJSON(
			http.StatusInternalServerError,
			gin.H{"error": "database error"},
		)
		return
	}

	c.JSON(http.StatusOK, repo)
}
