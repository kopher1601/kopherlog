package controller

import (
	"github.com/gin-gonic/gin"
	"kopherlog/domain"
	"kopherlog/service"
	"net/http"
)

type AuthController interface {
	Login(ctx *gin.Context)
}

type authController struct {
	svc service.AuthService
}

func NewAuthController(svc service.AuthService) AuthController {
	return &authController{svc: svc}
}

func (a authController) Login(ctx *gin.Context) {
	var login domain.Login
	err := ctx.ShouldBindJSON(&login)
	if err != nil {
		resp := &domain.ErrorResponse{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		}
		ctx.JSON(http.StatusBadRequest, resp)
		return
	}

	signin := &domain.SignIn{
		Email:    login.Email,
		Password: login.Password,
	}
	_, err = a.svc.SignIn(ctx, signin)

}
