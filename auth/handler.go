package main

import (
	"github.com/gin-gonic/gin"
	"github.com/markbates/goth/gothic"
)

func authHandler(r *gin.RouterGroup) {
	r.GET("/callback/:provider", callback)
	r.GET("/", auth)
	r.GET("/logout", logout)
}

func auth(c *gin.Context) {
	gothic.BeginAuthHandler(c.Writer, c.Request)
}

func callback(c *gin.Context) {}

func logout(c *gin.Context) {}
