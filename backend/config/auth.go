package config

import (
	"os"

	"github.com/gin-gonic/contrib/sessions"
	"github.com/markbates/goth"
	"github.com/markbates/goth/gothic"
	"github.com/markbates/goth/providers/github"
)

// Goth providers for authentication
func Goth() (err error) {
	c, err := Get()
	if err != nil {
		return err
	}

	// This stupid thing is added just so goth will work, no cookies are set
	// by goth
	store := sessions.NewCookieStore([]byte(os.Getenv("SESSION_SECRET")))
	store.Options(sessions.Options{
		Secure:   false,
		HttpOnly: true,
		Path:     "/",
		// Domain:   c.BaseURL,
	})
	gothic.Store = store

	// Auth providers
	goth.UseProviders(
		github.New(
			os.Getenv("GITHUB_KEY"),
			os.Getenv("GITHUB_SECRET"),
			"http://"+c.BaseURL+"/auth/callback/github",
		),
	)

	return err
}
