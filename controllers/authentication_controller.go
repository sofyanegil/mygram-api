package controllers

import (
	"fmt"
	"net/http"

	"mygram-api/data/request"
	"mygram-api/data/response"
	"mygram-api/helpers"
	"mygram-api/services"

	"github.com/go-playground/validator/v10"

	"github.com/gin-gonic/gin"
)

var validate = validator.New()

type AuthenticationController struct {
	authenticationService services.AuthenticationService
}

func NewAuthenticationController(services services.AuthenticationService) *AuthenticationController {
	return &AuthenticationController{authenticationService: services}
}

func (controller *AuthenticationController) Login(ctx *gin.Context) {
	loginRequest := request.LoginRequest{}
	ctx.ShouldBindJSON(&loginRequest)
	err := validate.Struct(loginRequest)
	helpers.ErrorPanic(err)

	token, err_token := controller.authenticationService.Login(loginRequest)
	fmt.Println(err_token)
	if err_token != nil {
		webResponse := response.Response{
			Code:    http.StatusBadRequest,
			Status:  "Bad Request",
			Message: "Invalid username or password",
		}
		ctx.JSON(http.StatusBadRequest, webResponse)
		return
	}

	resp := response.LoginResponse{
		TokenType: "Bearer",
		Token:     token,
	}

	webResponse := response.Response{
		Code:    200,
		Status:  "Ok",
		Message: "Successfully log in!",
		Data:    resp,
	}

	// ctx.SetCookie("token", token, config.TokenMaxAge*60, "/", "localhost", false, true)
	ctx.JSON(http.StatusOK, webResponse)
}

func (controller *AuthenticationController) Register(ctx *gin.Context) {
	createUsersRequest := request.CreateUserRequest{}
	ctx.ShouldBindJSON(&createUsersRequest)
	err := validate.Struct(createUsersRequest)
	if err != nil {
		helpers.ErrorBinding(ctx, err, http.StatusBadRequest, "Registration Failed!")
		return
	}

	savedUser, err := controller.authenticationService.Register(createUsersRequest)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, response.Response{
			Code:    http.StatusInternalServerError,
			Status:  "Internal Server Error",
			Message: "Registration Failed!",
			Errors:  []string{err.Error()},
		})
		return
	}

	ctx.JSON(http.StatusOK, response.Response{
		Code:    http.StatusOK,
		Status:  "Ok",
		Message: "Successfully created user!",
		Data:    savedUser,
	})
}
