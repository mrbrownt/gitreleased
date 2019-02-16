package main

import (
	"log"
	"os"

	"gitlab.com/mrbrownt/gitreleased.app/backend/handlers"
	"gitlab.com/mrbrownt/gitreleased.app/backend/models"

	"github.com/sirupsen/logrus"

	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	db, err := models.SetupDB()
	if err != nil {
		logrus.Panicln("falied to connect to database", err)
		// panic("failed to connect to database")
	}
	defer db.Close()

	// migrateDB(db)
	models.AutoMigrate(db)

	router := gin.Default()

	ug := router.Group("/user")

	handlers.UserHandler(ug)

	port := os.Getenv("PORT")
	router.Run("0.0.0.0:" + port)
}
