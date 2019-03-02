package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"gitlab.com/mrbrownt/gitreleased.app/backend/config"
	"gitlab.com/mrbrownt/gitreleased.app/backend/handlers"
	"gitlab.com/mrbrownt/gitreleased.app/backend/models"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		logrus.Fatalln(err)
	}

	gc, err := config.New()
	if err != nil {
		logrus.Fatalln(err)
	}

	db, err := models.SetupDB()
	if err != nil {
		logrus.Panicln("falied to connect to database", err)
	}
	defer db.Close()

	// migrateDB(db)
	models.AutoMigrate()

	err = config.Goth()
	if err != nil {
		logrus.Fatalln(err)
	}

	router := gin.Default()

	handlers.AuthHandler(router.Group("/auth"))
	handlers.UserHandler(router.Group("/user"))

	router.Run("0.0.0.0:" + gc.Port)
}
