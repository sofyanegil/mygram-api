package services

import (
	"errors"
	"mygram-api/config"
	"mygram-api/data/request"
	"mygram-api/data/response"
	"mygram-api/helpers"
	"mygram-api/models"
	"mygram-api/repository"
	"mygram-api/utils"
)

type AuthenticationService interface {
	Login(users request.LoginRequest) (string, error)
	Register(users request.CreateUserRequest) (response.UserResponse, error)
}

type AuthenticationServiceImpl struct {
	UsersRepository repository.UserRepository
}

func NewAuthenticationServiceImpl(usersRepository repository.UserRepository) AuthenticationService {
	return &AuthenticationServiceImpl{
		UsersRepository: usersRepository,
	}
}

func (a *AuthenticationServiceImpl) Login(users request.LoginRequest) (string, error) {
	new_users, users_err := a.UsersRepository.FindByUsername(users.Username)
	if users_err != nil {
		return "", errors.New("invalid username or Password")
	}

	config, _ := config.LoadConfig(".")

	verify_error := utils.VerifyPassword(new_users.Password, users.Password)
	if verify_error != nil {
		return "", errors.New("invalid username or Password")
	}

	token, err_token := utils.GenerateToken(config.TokenExpiresIn, new_users.ID, config.TokenSecret)
	helpers.ErrorPanic(err_token)
	return token, nil

}

func (a *AuthenticationServiceImpl) Register(u request.CreateUserRequest) (response.UserResponse, error) {
	hashedPassword, err := utils.HashPassword(u.Password)
	helpers.ErrorPanic(err)

	newUser := models.User{
		Username: u.Username,
		Email:    u.Email,
		Password: hashedPassword,
		Age:      u.Age,
	}
	savedUser, err := a.UsersRepository.Save(newUser)

	users := response.UserResponse{
		Id:        savedUser.ID.String(),
		Username:  savedUser.Username,
		Email:     savedUser.Email,
		Age:       savedUser.Age,
		CreatedAt: savedUser.CreatedAt,
		UpdatedAt: savedUser.UpdatedAt,
	}

	return users, err
}
