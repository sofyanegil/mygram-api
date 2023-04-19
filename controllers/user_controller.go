package controllers

import (
	"net/http"

	"mygram-api/data/response"
	"mygram-api/repository"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	userRepository repository.UserRepository
}

func NewUserController(repository repository.UserRepository) *UserController {
	return &UserController{userRepository: repository}
}

func (controller *UserController) GetUsers(ctx *gin.Context) {
	users := controller.userRepository.FindAll()
	webResponse := response.Response{
		Code:    200,
		Status:  "Ok",
		Message: "Successfully fetch all user data!",
		Data:    users,
	}

	ctx.JSON(http.StatusOK, webResponse)
}
