package handlers

import (
	"net/http"

	"github.com/dgrijalva/jwt-go"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/markbates/goth"
	"github.com/markbates/goth/gothic"
	"gitlab.com/mrbrownt/gitreleased.app/backend/models"
)

// AuthHandler routes auth thing
func AuthHandler(r *gin.RouterGroup) {
	r.GET("/callback/:provider", callback)
	r.GET("/", auth)
	r.GET("/logout", logout)
}

// Start basic auth, just redirects to provider and sends them on to the callback
func auth(c *gin.Context) {
	gothic.BeginAuthHandler(c.Writer, c.Request)
}

// Callback checks if the returned data is valid, checks if the user is in the
// system already and then sets up a JWT for future requests.
func callback(c *gin.Context) {
	provider := c.Param("provider")
	q := c.Request.URL.Query()
	q.Add("provider", provider)
	c.Request.URL.RawQuery = q.Encode()

	user, err := gothic.CompleteUserAuth(c.Writer, c.Request)
	if err != nil {
		c.AbortWithError(http.StatusNotImplemented, err)
		return
	}

	u := &models.User{}
	exist, err := doesUserExist(&user, u)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	if !exist {
		err = createUser(&user, u)
		if err != nil {
			c.AbortWithError(http.StatusInternalServerError, err)
			return
		}
	}

	returnUserAndJWT(c, u)
}

func logout(c *gin.Context) {
	c.SetCookie("Authorization", "", 0, "/", "localhost", false, true)
}

// Creates a user from the callback information we got
func createUser(callbackUser *goth.User, user *models.User) (err error) {
	user.GithubID = callbackUser.UserID
	user.GithubUserName = callbackUser.NickName
	user.Email = callbackUser.Email
	user.FirstName = callbackUser.FirstName
	user.LastName = callbackUser.LastName

	db := models.GetDB()
	err = db.Create(user).Error
	if err != nil {
		return err
	}

	return err
}

// Checks if the user exists and returns a bunch of shit, this should be
// refactored or renamed
func doesUserExist(callbackUser *goth.User, user *models.User) (exist bool, err error) {
	db := models.GetDB()
	err = db.Where(&models.User{GithubID: callbackUser.UserID}).First(user).Error

	if err == gorm.ErrRecordNotFound {
		return exist, nil
	}

	if err != nil {
		return exist, err
	}

	if user.GithubID == callbackUser.UserID {
		exist = true
	}
	return exist, err
}

// Creates a JWT for furture requests, should only be called after validating
// the callback is valid
func createJWT(user *models.User) (signedToken string, err error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":    user.ID,
		"admin": "false",
	})

	// TODO: use RSA or something better than totalShite
	return token.SignedString([]byte("totalShite"))
}

func returnUserAndJWT(c *gin.Context, u *models.User) {
	jwt, err := createJWT(u)
	if err != nil {
		c.AbortWithStatus(http.StatusNotImplemented)
	}
	c.SetCookie("Authorization", jwt, 300, "/", "localhost", false, true)
	c.Redirect(http.StatusSeeOther, "http://localhost:8081/#/user")
}
