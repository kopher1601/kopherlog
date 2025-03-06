package controller

import (
	"bytes"
	"context"
	"encoding/json"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/stretchr/testify/assert"
	"kopherlog/config"
	"kopherlog/domain"
	"kopherlog/repository"
	"kopherlog/service"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	config.Initialize()
	code := m.Run()
	os.Exit(code)
}

func Test_PostController_Post_Save(t *testing.T) {
	ctx := context.Background()
	client := config.SetupDB(t)
	r := gin.Default()

	postCreate := &domain.PostCreate{
		Title:   "吉祥寺マンション",
		Content: "吉祥寺マンション購入します。",
	}
	var buf bytes.Buffer
	_ = json.NewEncoder(&buf).Encode(postCreate)

	resp := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPost, "/posts", &buf)
	req.Header.Set("Content-Type", "application/json")
	postRepository := repository.NewPostRepository(client)
	postService := service.NewPostService(postRepository)
	postController := NewPostController(postService)
	r.POST("/posts", postController.PostCreate)
	r.ServeHTTP(resp, req)
	posts, err := postRepository.FindAll()
	assert.NoError(t, err)
	assert.Equal(t, http.StatusCreated, resp.Code)
	assert.Equal(t, 1, len(posts))

	t.Cleanup(func() {
		postRepository.DeleteAll(ctx)
	})
}

//
//func Test_PostController_PostCreate_Title_Required(t *testing.T) {
//	client := config.SetupDB(t)
//
//	r := gin.Default()
//	postCreate := &domain.PostCreate{
//		Title:   "",
//		Content: "吉祥寺マンション購入します。",
//	}
//	var buf bytes.Buffer
//	_ = json.NewEncoder(&buf).Encode(postCreate)
//
//	resp := httptest.NewRecorder()
//	req, _ := http.NewRequest(http.MethodPost, "/posts", &buf)
//	req.Header.Set("Content-Type", "application/json")
//
//	postRepository := repository.NewPostRepository(client)
//	postService := service.NewPostService(postRepository)
//	postController := NewPostController(postService)
//	r.POST("/posts", postController.PostCreate)
//	r.ServeHTTP(resp, req)
//
//	log.Println(resp.Body.String())
//	var errors domain.ErrorResponse
//	_ = json.NewDecoder(resp.Body).Decode(&errors)
//	assert.Equal(t, http.StatusBadRequest, resp.Code)
//	assert.Equal(t, "Title", errors.Validations[0].Field)
//	assert.Equal(t, "必須です。", errors.Validations[0].Message)
//}
//
//func Test_PostController_PostCreate_Content_Required(t *testing.T) {
//	client := config.SetupDB(t)
//	r := gin.Default()
//
//	postCreate := &domain.PostCreate{
//		Title:   "吉祥寺マンション",
//		Content: "",
//	}
//	var buf bytes.Buffer
//	_ = json.NewEncoder(&buf).Encode(postCreate)
//
//	resp := httptest.NewRecorder()
//	req, _ := http.NewRequest(http.MethodPost, "/posts", &buf)
//	req.Header.Set("Content-Type", "application/json")
//
//	postRepository := repository.NewPostRepository(client)
//	postService := service.NewPostService(postRepository)
//	postController := NewPostController(postService)
//	r.POST("/posts", postController.PostCreate)
//	r.ServeHTTP(resp, req)
//
//	var errors domain.ErrorResponse
//	_ = json.NewDecoder(resp.Body).Decode(&errors)
//	assert.Equal(t, http.StatusBadRequest, resp.Code)
//	assert.Equal(t, "Content", errors.Validations[0].Field)
//	assert.Equal(t, "必須です。", errors.Validations[0].Message)
//}
//
//func Test_PostController_PostCreate_Title_Content_Required(t *testing.T) {
//	client := config.SetupDB(t)
//	r := gin.Default()
//
//	postCreate := &domain.PostCreate{
//		Title:   "",
//		Content: "",
//	}
//	var buf bytes.Buffer
//	_ = json.NewEncoder(&buf).Encode(postCreate)
//
//	resp := httptest.NewRecorder()
//	req, _ := http.NewRequest(http.MethodPost, "/posts", &buf)
//	req.Header.Set("Content-Type", "application/json")
//
//	postRepository := repository.NewPostRepository(client)
//	postService := service.NewPostService(postRepository)
//	postController := NewPostController(postService)
//	r.POST("/posts", postController.PostCreate)
//	r.ServeHTTP(resp, req)
//
//	log.Println(resp.Body.String())
//	var errors domain.ErrorResponse
//	_ = json.NewDecoder(resp.Body).Decode(&errors)
//	assert.Equal(t, http.StatusBadRequest, resp.Code)
//	assert.Equal(t, "Title", errors.Validations[0].Field)
//	assert.Equal(t, "必須です。", errors.Validations[0].Message)
//	assert.Equal(t, "Content", errors.Validations[1].Field)
//	assert.Equal(t, "必須です。", errors.Validations[1].Message)
//}
