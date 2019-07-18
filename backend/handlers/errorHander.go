package handlers

import (
	"net/http"

	"github.com/getsentry/raven-go"
	"github.com/gin-gonic/gin"
)

func internalServerError(c *gin.Context, err error) {
	c.AbortWithStatus(http.StatusInternalServerError)
	cookie, cookieErr := c.Cookie("_gothic_session")
	if cookieErr != nil {
		raven.CaptureError(err, nil)
	} else {
		raven.CaptureError(err, map[string]string{"_gothic_session: cookie": cookie})
	}
}
