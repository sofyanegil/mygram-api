package middlewares

import (
	"net/http"
	"strings"

	"mygram-api/config"
	"mygram-api/repository"
	"mygram-api/utils"

	"github.com/gin-gonic/gin"
)

func DeserializeUser(userRepository repository.UserRepository) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var token string
		authorizationHeader := ctx.Request.Header.Get("Authorization")
		fields := strings.Fields(authorizationHeader)

		if len(fields) != 0 && fields[0] == "Bearer" {
			token = fields[1]
		}

		if token == "" {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"status": "fail", "message": "You are not logged in"})
			return
		}

		config, _ := config.LoadConfig(".")
		sub, err := utils.ValidateToken(token, config.TokenSecret)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"status": "fail", "message": err.Error()})
			return
		}

		result, err := userRepository.FindById(sub.(string))
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusForbidden, gin.H{"status": "fail", "message": "the user belonging to this token no logger exists"})
			return
		}

		ctx.Set("currentUser", result.Username)
		ctx.Set("userId", sub.(string))
		ctx.Next()
	}
}

func ProtectPhoto(photoRepository repository.PhotoRepository) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		userId := ctx.Value("userId")

		var _ = userId
		id := ctx.Param("id")
		photo, err := photoRepository.FindById(id)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusForbidden, gin.H{"status": "fail", "message": "Photo does't exist"})
			return
		}

		if userId != photo.UserID {
			ctx.AbortWithStatusJSON(http.StatusForbidden, gin.H{"status": "fail", "message": "Access Denied: You are not authorized to modify this data !"})
			return
		}

		ctx.Next()

	}
}

func ProtectComment(commentRepository repository.CommentRepository) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		userId := ctx.Value("userId")

		var _ = userId
		id := ctx.Param("id")
		photo, err := commentRepository.FindById(id)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusForbidden, gin.H{"status": "fail", "message": "Comment does't exist"})
			return
		}

		if userId != photo.UserID {
			ctx.AbortWithStatusJSON(http.StatusForbidden, gin.H{"status": "fail", "message": "Access Denied: You are not authorized to modify this data !"})
			return
		}

		ctx.Next()

	}
}

func ProtectSocialMedia(socialMediaRepository repository.SocialMediaRepository) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		userId := ctx.Value("userId")

		var _ = userId
		id := ctx.Param("id")
		photo, err := socialMediaRepository.FindById(id)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusForbidden, gin.H{"status": "fail", "message": "Comment does't exist"})
			return
		}

		if userId != photo.UserID {
			ctx.AbortWithStatusJSON(http.StatusForbidden, gin.H{"status": "fail", "message": "Access Denied: You are not authorized to modify this data !"})
			return
		}

		ctx.Next()

	}
}
