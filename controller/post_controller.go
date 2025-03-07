package controller

import (
	"github.com/gin-gonic/gin"
	"kopherlog/domain"
	"kopherlog/service"
	"net/http"
	"strconv"
)

type PostController interface {
	PostCreate(ctx *gin.Context)
	Get(ctx *gin.Context)
	GetAll(ctx *gin.Context)
}

type postController struct {
	postService service.PostService
}

func NewPostController(postService service.PostService) PostController {
	return &postController{postService: postService}
}

func (p *postController) PostCreate(ctx *gin.Context) {
	request := &domain.PostCreate{}
	if err := ctx.ShouldBindJSON(request); err != nil {
		errorResponse := &domain.ErrorResponse{
			Code:    http.StatusBadRequest,
			Message: "間違ったリクエストです。",
		}
		errorResponse.AddValidationErrors(err)
		ctx.JSON(http.StatusBadRequest, errorResponse)
		return
	}

	err := p.postService.Write(ctx, request)
	if err != nil {
		// TODO constructor 사용
		errorResponse := &domain.ErrorResponse{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse)
		return
	}
	ctx.Status(http.StatusCreated)
	return
}

func (p *postController) Get(ctx *gin.Context) {
	postID, err := strconv.Atoi(ctx.Param("postID"))
	if err != nil {
		errorResponse := &domain.ErrorResponse{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		}
		ctx.JSON(errorResponse.Code, errorResponse)
		return
	}
	result, err := p.postService.Get(ctx, postID)
	if err != nil {
		errorResponse := &domain.ErrorResponse{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		}
		ctx.JSON(errorResponse.Code, errorResponse)
		return
	}
	ctx.JSON(http.StatusOK, result)
}

func (p *postController) GetAll(ctx *gin.Context) {
	posts, err := p.postService.GetAll(ctx, nil)
	if err != nil {
		errorResponse := &domain.ErrorResponse{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		}
		ctx.JSON(errorResponse.Code, errorResponse)
		return
	}
	ctx.JSON(http.StatusOK, posts)
}
