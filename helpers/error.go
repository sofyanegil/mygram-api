package helpers

import (
	"fmt"
	"net/http"

	"mygram-api/data/response"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func ErrorPanic(err error) {
	if err != nil {
		panic(err)
	}
}

func ErrorResponse(ctx *gin.Context, err error, msg string) {
	//HANDLE GORM ERROR
	ctx.JSON(http.StatusBadRequest, response.Response{
		Code:    http.StatusBadRequest,
		Status:  "Bad Request",
		Message: msg,
		Errors:  []string{err.Error()},
	})

}

func ErrorBinding(ctx *gin.Context, err error, statusCode int, msg string) {
	var errorMessages []string
	if validationErrors, ok := err.(validator.ValidationErrors); ok {
		for _, e := range validationErrors {
			errorMessage := fmt.Sprintf("ERROR: Invalid %s (%s)!", e.Field(), e.Error())
			errorMessages = append(errorMessages, errorMessage)
		}
	} else {
		errorMessages = append(errorMessages, err.Error())
	}

	ctx.JSON(http.StatusBadRequest, response.Response{
		Code:    statusCode,
		Status:  "Error",
		Message: msg,
		Errors:  errorMessages,
	})
}
