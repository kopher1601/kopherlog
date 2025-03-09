package controller

import (
	"github.com/gin-gonic/gin"
	"kopherlog/domain"
)

type AuthController interface {
	Login(ctx *gin.Context)
}

type authController struct {
}

func NewAuthController() AuthController {
	return &authController{}
}

func (a authController) Login(ctx *gin.Context) {
	var login domain.Login
	err := ctx.ShouldBindJSON(&login)
	if err != nil {
		//
	}
}
