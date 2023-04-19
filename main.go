package main

import (
	"log"
	"net/http"
	"time"

	"mygram-api/config"
	"mygram-api/controllers"
	"mygram-api/helpers"
	"mygram-api/models"
	"mygram-api/repository"
	"mygram-api/routers"
	"mygram-api/services"
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

	//Init Repository
	userRepository := repository.NewUsersRepositoryImpl(db)
	photoRepository := repository.NewPhotoRepositoryImpl(db)
	commentRepository := repository.NewCommentRepositoryImpl(db)
	socialMediaRepository := repository.NewSocialMediaRepositoryImpl(db)

	//Init Service
	authenticationService := services.NewAuthenticationServiceImpl(userRepository)

	//Init controller
	authenticationController := controllers.NewAuthenticationController(authenticationService)
	userController := controllers.NewUserController(userRepository)
	photoController := controllers.NewPhotoController(photoRepository)
	commentController := controllers.NewCommentController(commentRepository)
	socialMediaController := controllers.NewSocialMediaController(socialMediaRepository)

	routes := routers.NewRouter(
		userRepository,
		userController,
		photoRepository,
		photoController,
		commentRepository,
		commentController,
		socialMediaRepository,
		socialMediaController,
		authenticationController,
	)
	server := &http.Server{
		Addr:           ":" + loadConfig.ServerPort,
		Handler:        routes,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	server_err := server.ListenAndServe()
	helpers.ErrorPanic(server_err)
}
