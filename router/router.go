package router

import (
	"github.com/gin-gonic/gin"
	"kopherlog/controller"
)

func Setup() *gin.Engine {
	router := gin.Default()

	postController := controller.NewPostController()
	router.POST("/posts", postController.PostCreate)

	return router
}
