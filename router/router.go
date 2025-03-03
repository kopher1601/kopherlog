package router

import (
	"github.com/gin-gonic/gin"
	"kopherlog/controller"
	"kopherlog/ent"
	"kopherlog/repository"
	"kopherlog/service"
)

func Setup(client *ent.Client) *gin.Engine {
	router := gin.Default()

	postRepository := repository.NewPostRepository(client)
	postService := service.NewPostService(postRepository)
	postController := controller.NewPostController(postService)
	router.POST("/posts", postController.PostCreate)

	return router
}
