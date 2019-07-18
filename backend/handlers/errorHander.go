package handlers

import (
	"net/http"

	"github.com/getsentry/raven-go"
	"github.com/gin-gonic/gin"
)

func internalServerError(c *gin.Context, err error) {
	c.AbortWithStatus(http.StatusInternalServerError)
	raven.CaptureError(err, nil)
}
