package router

import (
	"github.com/gin-gonic/gin"
	"kopherlog/controller"
	"kopherlog/ent"
	"kopherlog/middleware"
	"kopherlog/repository"
	"kopherlog/service"
)

func Setup(client *ent.Client) *gin.Engine {
	router := gin.Default()

	postRepository := repository.NewPostRepository(client)
	postService := service.NewPostService(postRepository)
	postController := controller.NewPostController(postService)

	router.GET("/posts/:postID", postController.Get)
	router.PUT("/posts/:postID", postController.Edit)
	router.DELETE("/posts/:postID", postController.Delete)
	router.POST("/posts", postController.PostCreate)
	router.GET("/posts", middleware.ValidateQueryParams(), postController.GetAll)

	return router
}
