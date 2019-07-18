package handlers

import (
	"fmt"
	"net/http"

	"github.com/getsentry/raven-go"
	"github.com/gin-gonic/gin"
)

func internalServerError(c *gin.Context, err error) {
	c.AbortWithStatus(http.StatusInternalServerError)
	cookie, cookieErr := c.Cookie("_gothic_session")
	if cookieErr == nil {
		raven.CaptureMessage(fmt.Sprintf("_gothic_session cookie: %s", cookie), nil)
	}
	raven.CaptureError(err, nil)
}
