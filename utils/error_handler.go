package utils

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mhmmdrivaldhi/go-social-media-api/model/dto"
)

func HandleError(ctx *gin.Context, err error) {
	var statusCode int
	
	switch err.(type) {
	case *NotFoundError:
		statusCode = http.StatusNotFound
	case *BadRequestError:
		statusCode = http.StatusBadRequest
	case *InternalServerError:
		statusCode = http.StatusInternalServerError
	case *UnauthorizedError:
		statusCode = http.StatusUnauthorized
	}

	response := Response(dto.ResponseParam{
		StatusCode: statusCode,
		Message:    err.Error(),
	})

	ctx.JSON(statusCode, response)
}