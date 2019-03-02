package handlers

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

// authMiddleware adds id and admin to Context keys.
// Add r.Use(authMiddleware()) to any protected endpoint
func authMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		jwtCookie, err := c.Cookie("Authorization")
		if err != nil {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		if jwtCookie == "" {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		token, err := jwt.Parse(jwtCookie, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, errors.New("Unexpected signing method")
			}
			// TODO: Use RSA or something better than totalShite
			return []byte("totalShite"), nil
		})

		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			admin, err := strconv.ParseBool(claims["admin"].(string))
			if err != nil {
				// TODO: should probably redirect to the login stage
				c.AbortWithStatus(http.StatusInternalServerError)
				return
			}

			c.Set("id", claims["id"])
			c.Set("admin", admin)
			c.Next()
			return
		}
		// TODO: should probably redirect to the login stage
		c.AbortWithStatus(http.StatusUnauthorized)
	}
}
