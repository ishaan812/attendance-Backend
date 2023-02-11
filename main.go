package main

//	@title			Weber API
//	@version		1.001
//	@description	Weber Backend API

//	@host						localhost:9000
//	@BasePath					/
//	@query.collection.format	json

import (
	"os"
	"service/controllers"
	"service/database"
	"service/routes"

	"github.com/joho/godotenv"
	"gorm.io/gorm"
)

var db *gorm.DB

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		panic("Error loading .env file")
	}
	
	db := database.InitialMigration(os.Getenv("DATABASE_CONNECTION_STRING"))

	controllers.GetDB(db)
	routes.InitializeRouter()
}
