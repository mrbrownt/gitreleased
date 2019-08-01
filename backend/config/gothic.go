package config

import (
	"fmt"
	"os"

	"github.com/gin-gonic/contrib/sessions"
	"github.com/markbates/goth"
	"github.com/markbates/goth/gothic"
	"github.com/markbates/goth/providers/github"
)

func configGoth() {

	// This stupid thing is added just so goth will work, uses a temp cookie
	store := sessions.NewCookieStore([]byte(os.Getenv("SESSION_SECRET")))

	var secure = false
	if globalConf.Environment == "production" {
		secure = true
	}

	store.Options(sessions.Options{
		Secure:   secure,
		HttpOnly: true,
		Path:     "/",
		Domain:   globalConf.BaseURL,
	})
	gothic.Store = store

	callbackURL := fmt.Sprintf("http://%s/auth/callback/github", globalConf.BaseURL)
	if secure {
		callbackURL = fmt.Sprintf("https://%s/auth/callback/github", globalConf.BaseURL)
	}

	// Auth providers
	goth.UseProviders(
		github.New(
			os.Getenv("GITHUB_KEY"),
			os.Getenv("GITHUB_SECRET"),
			callbackURL,
		),
	)
}
