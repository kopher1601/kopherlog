package controller

import (
	"github.com/gin-gonic/gin"
	"kopherlog/domain"
	"log"
	"net/http"
)

type PostController interface {
	PostCreate(ctx *gin.Context)
}

type postController struct {
}

func NewPostController() PostController {
	return &postController{}
}

func (p *postController) PostCreate(ctx *gin.Context) {
	request := &domain.PostCreate{}
	if err := ctx.ShouldBindJSON(request); err != nil {
		errorResponse := domain.ErrorResponse{
			Code:    http.StatusBadRequest,
			Message: "間違ったリクエストです。",
		}
		errorResponse.AddValidationErrors(err)
		ctx.JSON(http.StatusBadRequest, errorResponse)
		return
	}

	log.Println("request =>", request)

	ctx.String(http.StatusOK, "Hello World")
}
