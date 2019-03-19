package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"gitlab.com/mrbrownt/gitreleased.app/backend/config"
	"gitlab.com/mrbrownt/gitreleased.app/backend/handlers"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		logrus.Debugln("Not using .env")
	}

	gc, err := config.New()
	if err != nil {
		logrus.Fatalln(err)
	}
	defer gc.DB.Close()

	err = config.Goth()
	if err != nil {
		logrus.Fatalln(err)
	}

	router := gin.Default()

	api := router.Group("/api")

	handlers.AuthHandler(router.Group("/auth"))
	handlers.UserHandler(api.Group("/user"))

	router.Run("0.0.0.0:" + gc.Port)
}
