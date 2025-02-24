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
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	log.Println("request =>", request)

	ctx.String(http.StatusOK, "Hello World")
}
