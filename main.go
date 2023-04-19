package main

import (
	"log"

	"mygram-api/config"
	"mygram-api/models"
)

func main() {
	loadConfig, err := config.LoadConfig(".")
	if err != nil {
		log.Fatal("❌ Could not load environment variables", err)
	}

	//Database
	db := config.ConnectionDB(&loadConfig)
	db.Table("users").AutoMigrate(&models.User{})
	db.Table("photos").AutoMigrate(&models.Photo{})
	db.Table("comments").AutoMigrate(&models.Comment{})
	db.Table("social_media").AutoMigrate(&models.SocialMedia{})
}
