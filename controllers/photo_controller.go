package controllers

import (
	"net/http"

	"mygram-api/data/request"
	"mygram-api/data/response"
	"mygram-api/helpers"
	"mygram-api/models"
	"mygram-api/repository"

	"github.com/gin-gonic/gin"
)

type PhotoController struct {
	photoRepository repository.PhotoRepository
}

func NewPhotoController(repository repository.PhotoRepository) *PhotoController {
	return &PhotoController{photoRepository: repository}
}

func (controller *PhotoController) GetPhotos(ctx *gin.Context) {
	photos, err := controller.photoRepository.FindAll()
	if err != nil {
		helpers.ErrorResponse(ctx, err, "photos not found")
		return
	}

	webResponse := response.Response{
		Code:    200,
		Status:  "Ok",
		Message: "Successfully fetch all photos data!",
		Data:    photos,
	}

	ctx.JSON(http.StatusOK, webResponse)
}

func (controller *PhotoController) GetPhoto(ctx *gin.Context) {
	id := ctx.Param("id")
	photo, err := controller.photoRepository.FindById(id)
	if err != nil {
		helpers.ErrorResponse(ctx, err, "photos not found")
		return
	}

	ctx.JSON(http.StatusOK, response.Response{
		Code:    http.StatusOK,
		Status:  "Ok",
		Message: "Successfully fetch photo data!",
		Data:    photo,
	})
}

func (controller *PhotoController) CreatePhoto(ctx *gin.Context) {
	createPhotoRequest := request.CreatePhotoRequest{}
	userId := ctx.Value("userId")

	ctx.ShouldBindJSON(&createPhotoRequest)
	photo := models.Photo{
		Title:    createPhotoRequest.Title,
		Caption:  createPhotoRequest.Caption,
		PhotoUrl: createPhotoRequest.Photo_url,
		UserID:   userId.(string),
	}
	err := validate.Struct(createPhotoRequest)
	if err != nil {
		helpers.ErrorBinding(ctx, err, http.StatusBadRequest, "Upload Photo Failed!")
		return
	}

	newPhoto, err := controller.photoRepository.Save(photo)
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
		Message: "Successfully upload photo!",
		Data:    newPhoto,
	})
}

func (controller *PhotoController) UpdatePhoto(ctx *gin.Context) {
	updatePhotoRequest := request.UpdatePhotoRequest{}
	photoId := ctx.Param("id")

	ctx.ShouldBindJSON(&updatePhotoRequest)
	photo, err := controller.photoRepository.FindById(photoId)
	if err != nil {
		helpers.ErrorResponse(ctx, err, "Photo not found")
		return
	}

	photo.Title = updatePhotoRequest.Title
	photo.Caption = updatePhotoRequest.Caption
	photo.PhotoUrl = updatePhotoRequest.Photo_url

	err = validate.Struct(updatePhotoRequest)
	if err != nil {
		helpers.ErrorBinding(ctx, err, http.StatusBadRequest, "Update Photo Failed!")
		return
	}

	updatedPhoto, err := controller.photoRepository.Update(*photo)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, response.Response{
			Code:    http.StatusInternalServerError,
			Status:  "Internal Server Error",
			Message: "Update Failed!",
			Errors:  []string{err.Error()},
		})
		return
	}

	ctx.JSON(http.StatusOK, response.Response{
		Code:    http.StatusOK,
		Status:  "Ok",
		Message: "Successfully updated photo!",
		Data:    updatedPhoto,
	})
}

func (controller *PhotoController) DeletePhoto(ctx *gin.Context) {
	photoId := ctx.Param("id")

	err := controller.photoRepository.Delete(photoId)
	if err != nil {
		helpers.ErrorResponse(ctx, err, "Photo not found")
		return
	}

	ctx.JSON(http.StatusOK, response.Response{
		Code:    http.StatusOK,
		Status:  "Ok",
		Message: "Successfully deleted photo!",
	})
}
