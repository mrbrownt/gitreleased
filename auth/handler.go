package main

import (
	"net/http"

	"github.com/markbates/goth"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/markbates/goth/gothic"
	"gitlab.com/mrbrownt/gitreleased.app/backend/models"
)

func authHandler(r *gin.RouterGroup) {
	r.GET("/callback/:provider", callback)
	r.GET("/", auth)
	r.GET("/logout", logout)
}

func auth(c *gin.Context) {
	gothic.BeginAuthHandler(c.Writer, c.Request)
}

func callback(c *gin.Context) {
	provider := c.Param("provider")

	query := c.Request.URL.Query()
	query.Add("provider", provider)
	c.Request.URL.RawQuery = query.Encode()

	user, err := gothic.CompleteUserAuth(c.Writer, c.Request)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	u, err := createUser(user)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	createJWT(c, u)
}

func createUser(githubUser goth.User) (user models.User, err error) {
	user.AccessToken = githubUser.AccessToken
	user.Email = githubUser.Email
	user.GithubID = githubUser.UserID
	user.GithubUserName = githubUser.NickName

	err = db.FirstOrCreate(&user, &models.User{GithubUserName: user.GithubUserName}).Error
	if err != nil {
		return user, err
	}

	return user, nil
}

func createJWT(c *gin.Context, u models.User) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":    u.ID,
		"admin": "false",
	})

	// TODO: use RSA or something better than totalShite
	userToken, err := token.SignedString([]byte("totalShite"))
	if err != nil {
		c.AbortWithStatus(http.StatusNotImplemented)
		return
	}
	c.SetCookie("Authorization", userToken, 3600, "/", "localhost", false, true)
	c.Redirect(http.StatusSeeOther, "/#/user")
}

func logout(c *gin.Context) {
	c.SetCookie("Authorization", "", 0, "/", "localhost", false, true)
	c.Redirect(http.StatusSeeOther, "/")
}
