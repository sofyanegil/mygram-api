package repository

import (
	"errors"
	"mygram-api/data/request"
	"mygram-api/data/response"
	"mygram-api/helpers"
	"mygram-api/models"

	"gorm.io/gorm"
)

type UserRepository interface {
	FindAll() *[]response.UserResponse
	FindById(usersId string) (response.UserResponse, error)
	FindByUsername(username string) (models.User, error)
	Save(users models.User) (*models.User, error)
	Update(users models.User)
	Delete(usersId int)
}

type UserRepositoryImpl struct {
	DB *gorm.DB
}

func NewUsersRepositoryImpl(DB *gorm.DB) UserRepository {
	return &UserRepositoryImpl{DB: DB}
}

func (u *UserRepositoryImpl) Save(users models.User) (*models.User, error) {
	result := u.DB.Create(&users)
	return &users, result.Error
}

func (u *UserRepositoryImpl) Update(users models.User) {
	var updateUsers = request.UpdateUserRequest{
		// Id:       users.ID,
		Username: users.Username,
		Email:    users.Email,
		Password: users.Password,
	}
	result := u.DB.Model(&users).Updates(updateUsers)
	helpers.ErrorPanic(result.Error)
}

func (u *UserRepositoryImpl) Delete(usersId int) {
	var users models.User
	result := u.DB.Where("id = ?", usersId).Delete(&users)
	helpers.ErrorPanic(result.Error)
}

func (u *UserRepositoryImpl) FindById(usersId string) (response.UserResponse, error) {
	var users models.User
	result := u.DB.Find(&users, usersId)
	if result != nil {
		return response.UserResponse{Id: users.ID.String(), Username: users.Username, Email: users.Email, Age: users.Age, CreatedAt: users.CreatedAt, UpdatedAt: users.UpdatedAt}, nil
	} else {
		return response.UserResponse{}, errors.New("users is not found")
	}
}

func (u *UserRepositoryImpl) FindAll() *[]response.UserResponse {
	var users []models.User
	results := u.DB.Preload("Photos").Find(&users)
	helpers.ErrorPanic(results.Error)

	userResponses := make([]response.UserResponse, len(users))
	for i, user := range users {
		userResponses[i] = response.UserResponse{Id: user.ID.String(), Username: user.Username, Email: user.Email, Age: user.Age, Photos: user.Photos, CreatedAt: user.CreatedAt, UpdatedAt: user.UpdatedAt}
	}

	return &userResponses
}

func (u *UserRepositoryImpl) FindByUsername(username string) (models.User, error) {
	var users models.User
	result := u.DB.First(&users, "username = ?", username)

	if result.Error != nil {
		return users, errors.New("invalid username or Password")
	}
	return users, nil
}
