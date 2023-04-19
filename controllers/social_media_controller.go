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

type SocialMediaController struct {
	socialMediaRepository repository.SocialMediaRepository
}

func NewSocialMediaController(repository repository.SocialMediaRepository) *SocialMediaController {
	return &SocialMediaController{socialMediaRepository: repository}
}

func (controller *SocialMediaController) GetSocialMedias(ctx *gin.Context) {
	socialMedia, err := controller.socialMediaRepository.FindAll()
	if err != nil {
		helpers.ErrorResponse(ctx, err, "SocialMedia not found")
		return
	}

	webResponse := response.Response{
		Code:    200,
		Status:  "Ok",
		Message: "Successfully fetch all SocialMedia data!",
		Data:    socialMedia,
	}

	ctx.JSON(http.StatusOK, webResponse)
}

func (controller *SocialMediaController) CreateSocialMedia(ctx *gin.Context) {
	createSocialMediaRequest := request.CreateSocialMediaRequest{}
	userId := ctx.Value("userId")

	ctx.ShouldBindJSON(&createSocialMediaRequest)
	socialMedia := models.SocialMedia{
		Name:           createSocialMediaRequest.Name,
		SocialMediaUrl: createSocialMediaRequest.SocialMediaUrl,
		UserID:         userId.(string),
	}
	err := validate.Struct(createSocialMediaRequest)
	if err != nil {
		helpers.ErrorBinding(ctx, err, http.StatusBadRequest, "Create SocialMedia Failed!")
		return
	}

	newSocialMedia, err := controller.socialMediaRepository.Save(socialMedia)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, response.Response{
			Code:    http.StatusInternalServerError,
			Status:  "Internal Server Error",
			Message: "Create Media Failed!",
			Errors:  []string{err.Error()},
		})
		return
	}

	ctx.JSON(http.StatusOK, response.Response{
		Code:    http.StatusOK,
		Status:  "Ok",
		Message: "Successfully created SocialMedia!",
		Data:    newSocialMedia,
	})

}

func (controller *SocialMediaController) UpdateSocialMedia(ctx *gin.Context) {
	updateSocialMediaRequest := request.UpdateSocialMediaRequest{}
	socialMediaId := ctx.Param("id")

	ctx.ShouldBindJSON(&updateSocialMediaRequest)
	socialMedia, err := controller.socialMediaRepository.FindById(socialMediaId)
	if err != nil {
		helpers.ErrorResponse(ctx, err, "SocialMedia not found")
		return
	}

	if updateSocialMediaRequest.Name != nil {
		socialMedia.Name = *updateSocialMediaRequest.Name
	}
	if updateSocialMediaRequest.SocialMediaUrl != nil {
		socialMedia.SocialMediaUrl = *updateSocialMediaRequest.SocialMediaUrl
	}

	err = validate.Struct(&updateSocialMediaRequest)
	if err != nil {
		helpers.ErrorBinding(ctx, err, http.StatusBadRequest, "Update SocialMedia Failed!")
		return
	}

	updatedMedia, err := controller.socialMediaRepository.Update(*socialMedia)
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
		Message: "Successfully updated socialMedia!",
		Data:    updatedMedia,
	})
}

func (controller *SocialMediaController) DeleteSocialMedia(ctx *gin.Context) {
	socialMediaId := ctx.Param("id")

	err := controller.socialMediaRepository.Delete(socialMediaId)
	if err != nil {
		helpers.ErrorResponse(ctx, err, "Media not found")
		return
	}
	ctx.JSON(http.StatusOK, response.Response{
		Code:    http.StatusOK,
		Status:  "Ok",
		Message: "Successfully deleted SocialMedia!",
	})
}

func (controller *SocialMediaController) GetSocialMedia(ctx *gin.Context) {
	id := ctx.Param("id")
	socialMedia, err := controller.socialMediaRepository.FindById(id)
	if err != nil {
		helpers.ErrorResponse(ctx, err, "SocialMedia not found")
		return
	}

	ctx.JSON(http.StatusOK, response.Response{
		Code:    http.StatusOK,
		Status:  "Ok",
		Message: "Successfully fetch SocialMedia data!",
		Data:    socialMedia,
	})
}
