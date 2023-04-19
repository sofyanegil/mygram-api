package routers

import (
	"mygram-api/controllers"
	"mygram-api/middlewares"
	"mygram-api/repository"

	"github.com/gin-gonic/gin"
)

func NewRouter(
	userRepository repository.UserRepository,
	userController *controllers.UserController,

	photoRepository repository.PhotoRepository,
	photoController *controllers.PhotoController,

	commentRepository repository.CommentRepository,
	commentController *controllers.CommentController,

	socialMediaRepository repository.SocialMediaRepository,
	socialMediaController *controllers.SocialMediaController,

	authenticationController *controllers.AuthenticationController,
) *gin.Engine {
	router := gin.Default()
	userRouter := router.Group("/users")
	{
		userRouter.POST("/register", authenticationController.Register)
		userRouter.POST("/login", authenticationController.Login)
		userRouter.GET("/list", middlewares.DeserializeUser(userRepository), userController.GetUsers)
	}

	photoRouter := router.Group("/photos")
	{
		photoRouter.GET("/", middlewares.DeserializeUser(userRepository), photoController.GetPhotos)
		photoRouter.GET("/:id", middlewares.DeserializeUser(userRepository), photoController.GetPhoto)
		photoRouter.POST("/", middlewares.DeserializeUser(userRepository), photoController.CreatePhoto)
		photoRouter.PUT("/:id", middlewares.DeserializeUser(userRepository), middlewares.ProtectPhoto(photoRepository), photoController.UpdatePhoto)
		photoRouter.DELETE("/:id", middlewares.DeserializeUser(userRepository), middlewares.ProtectPhoto(photoRepository), photoController.DeletePhoto)
	}

	commentRouter := router.Group("/comments")
	{
		commentRouter.GET("/", middlewares.DeserializeUser(userRepository), commentController.GetAllComments)
		commentRouter.GET("/:id", middlewares.DeserializeUser(userRepository), commentController.GetComment)
		commentRouter.POST("/", middlewares.DeserializeUser(userRepository), commentController.CreateComment)
		commentRouter.PUT("/:id", middlewares.DeserializeUser(userRepository), middlewares.ProtectComment(commentRepository), commentController.UpdateComment)
		commentRouter.DELETE("/:id", middlewares.DeserializeUser(userRepository), middlewares.ProtectComment(commentRepository), commentController.DeleteComment)
	}

	socialMediaRouter := router.Group("/socialmedia")
	{
		socialMediaRouter.GET("/", middlewares.DeserializeUser(userRepository), socialMediaController.GetSocialMedias)
		socialMediaRouter.GET("/:id", middlewares.DeserializeUser(userRepository), socialMediaController.GetSocialMedia)
		socialMediaRouter.POST("/", middlewares.DeserializeUser(userRepository), socialMediaController.CreateSocialMedia)
		socialMediaRouter.PUT("/:id", middlewares.DeserializeUser(userRepository), middlewares.ProtectSocialMedia(socialMediaRepository), socialMediaController.UpdateSocialMedia)
		socialMediaRouter.DELETE("/:id", middlewares.DeserializeUser(userRepository), middlewares.ProtectSocialMedia(socialMediaRepository), socialMediaController.DeleteSocialMedia)
	}

	return router
}
