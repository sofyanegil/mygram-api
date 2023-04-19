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

type CommentController struct {
	commentRepository repository.CommentRepository
}

func NewCommentController(repository repository.CommentRepository) *CommentController {
	return &CommentController{commentRepository: repository}
}

func (controller *CommentController) GetAllComments(ctx *gin.Context) {
	comments, err := controller.commentRepository.FindAll()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, response.Response{
			Code:    http.StatusInternalServerError,
			Status:  "Internal Server Error",
			Message: "Failed to fetch comments!",
			Errors:  []string{err.Error()},
		})
		return
	}

	ctx.JSON(http.StatusOK, response.Response{
		Code:    http.StatusOK,
		Status:  "Ok",
		Message: "Successfully fetched comments!",
		Data:    comments,
	})
}

func (controller *CommentController) GetComment(ctx *gin.Context) {
	id := ctx.Param("id")
	comment, err := controller.commentRepository.FindById(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, response.Response{
			Code:    http.StatusNotFound,
			Status:  "Not Found",
			Message: "Comment not found!",
			Errors:  []string{err.Error()},
		})
		return
	}

	ctx.JSON(http.StatusOK, response.Response{
		Code:    http.StatusOK,
		Status:  "Ok",
		Message: "Successfully fetched comment!",
		Data:    comment,
	})
}

func (controller *CommentController) CreateComment(ctx *gin.Context) {
	createCommentRequest := request.CreateCommentRequest{}
	ctx.ShouldBindJSON(&createCommentRequest)
	err := validate.Struct(createCommentRequest)
	if err != nil {
		helpers.ErrorBinding(ctx, err, http.StatusBadRequest, "Comment Failed!")
		return
	}
	userId := ctx.Value("userId")

	comment := models.Comment{
		UserID:  userId.(string),
		PhotoID: createCommentRequest.PhotoId,
		Message: createCommentRequest.Message,
	}

	newComment, err := controller.commentRepository.Save(comment)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, response.Response{
			Code:    http.StatusInternalServerError,
			Status:  "Internal Server Error",
			Message: "Failed to create comment!",
			Errors:  []string{err.Error()},
		})
		return
	}

	ctx.JSON(http.StatusOK, response.Response{
		Code:    http.StatusOK,
		Status:  "Ok",
		Message: "Successfully created comment!",
		Data:    newComment,
	})
}

func (controller *CommentController) UpdateComment(ctx *gin.Context) {
	updateCommentRequest := request.UpdateCommentRequest{}
	id := ctx.Param("id")

	ctx.ShouldBindJSON(&updateCommentRequest)
	comment, err := controller.commentRepository.FindById(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, response.Response{
			Code:    http.StatusNotFound,
			Status:  "Not Found",
			Message: "Comment not found!",
			Errors:  []string{err.Error()},
		})
		return
	}

	comment.Message = updateCommentRequest.Message
	err = validate.Struct(updateCommentRequest)
	if err != nil {
		helpers.ErrorBinding(ctx, err, http.StatusBadRequest, "Update Comment Failed!")
		return
	}

	updatedComment, err := controller.commentRepository.Update(*comment)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, response.Response{
			Code:    http.StatusInternalServerError,
			Status:  "Internal Server Error",
			Message: "Failed to update comment!",
			Errors:  []string{err.Error()},
		})
		return
	}

	ctx.JSON(http.StatusOK, response.Response{
		Code:    http.StatusOK,
		Status:  "Ok",
		Message: "Successfully updated comment!",
		Data:    updatedComment,
	})
}

func (controller *CommentController) DeleteComment(ctx *gin.Context) {
	id := ctx.Param("id")
	_, err := controller.commentRepository.FindById(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, response.Response{
			Code:    http.StatusNotFound,
			Status:  "Not Found",
			Message: "Comment not found!",
			Errors:  []string{err.Error()},
		})
		return
	}

	err = controller.commentRepository.Delete(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, response.Response{
			Code:    http.StatusInternalServerError,
			Status:  "Internal Server Error",
			Message: "Failed to delete comment!",
			Errors:  []string{err.Error()},
		})
		return
	}

	ctx.JSON(http.StatusOK, response.Response{
		Code:    http.StatusOK,
		Status:  "Ok",
		Message: "Successfully deleted comment!",
	})
}
