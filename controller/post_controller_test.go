package controller

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/stretchr/testify/assert"
	"kopherlog/config"
	"kopherlog/domain"
	"kopherlog/ent"
	"kopherlog/repository"
	"kopherlog/service"
	"log"
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

func WithTxTest(ctx context.Context, client *ent.Client, fn func(tx *ent.Tx) error) {
	tx, err := client.Tx(ctx)
	if err != nil {
		fmt.Println(err)
	}
	defer func() {
		if v := recover(); v != nil {
			tx.Rollback()
			panic(v)
		}
	}()
	if err := fn(tx); err != nil {
		if rerr := tx.Rollback(); rerr != nil {
			err = fmt.Errorf("%w: rolling back transaction: %v", err, rerr)
		}
		fmt.Println(err)
	}
	if err := tx.Rollback(); err != nil {
		fmt.Println("committing transaction: %w", err)
	}
}

//func Test_PostController_Post_Save(t *testing.T) {
//	WithTxTest(context.Background(), config.SetupDB(t), func(tx *ent.Tx) error {
//		r := gin.Default()
//
//		postCreate := &domain.PostCreate{
//			Title:   "吉祥寺マンション",
//			Content: "吉祥寺マンション購入します。",
//		}
//		var buf bytes.Buffer
//		_ = json.NewEncoder(&buf).Encode(postCreate)
//
//		resp := httptest.NewRecorder()
//		req, _ := http.NewRequest(http.MethodPost, "/posts", &buf)
//		req.Header.Set("Content-Type", "application/json")
//		postRepository := repository.NewPostRepository(tx.Client(), tx)
//		postService := service.NewPostService(postRepository)
//		postController := NewPostController(postService)
//		r.POST("/posts", postController.PostCreate)
//		r.ServeHTTP(resp, req)
//		posts, err := postRepository.FindAll()
//		assert.NoError(t, err)
//		assert.Equal(t, http.StatusCreated, resp.Code)
//		assert.Equal(t, 1, len(posts))
//		return err
//	})
//}

func Test_PostController_Post_Save(t *testing.T) {
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
}

func Test_PostController_PostCreate_Title_Required(t *testing.T) {
	client := config.SetupDB(t)

	r := gin.Default()
	postCreate := &domain.PostCreate{
		Title:   "",
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

	log.Println(resp.Body.String())
	var errors domain.ErrorResponse
	_ = json.NewDecoder(resp.Body).Decode(&errors)
	assert.Equal(t, http.StatusBadRequest, resp.Code)
	assert.Equal(t, "Title", errors.Validations[0].Field)
	assert.Equal(t, "必須です。", errors.Validations[0].Message)
}

func Test_PostController_PostCreate_Content_Required(t *testing.T) {
	client := config.SetupDB(t)
	r := gin.Default()

	postCreate := &domain.PostCreate{
		Title:   "吉祥寺マンション",
		Content: "",
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

	var errors domain.ErrorResponse
	_ = json.NewDecoder(resp.Body).Decode(&errors)
	assert.Equal(t, http.StatusBadRequest, resp.Code)
	assert.Equal(t, "Content", errors.Validations[0].Field)
	assert.Equal(t, "必須です。", errors.Validations[0].Message)
}

func Test_PostController_PostCreate_Title_Content_Required(t *testing.T) {
	client := config.SetupDB(t)
	r := gin.Default()

	postCreate := &domain.PostCreate{
		Title:   "",
		Content: "",
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

	log.Println(resp.Body.String())
	var errors domain.ErrorResponse
	_ = json.NewDecoder(resp.Body).Decode(&errors)
	assert.Equal(t, http.StatusBadRequest, resp.Code)
	assert.Equal(t, "Title", errors.Validations[0].Field)
	assert.Equal(t, "必須です。", errors.Validations[0].Message)
	assert.Equal(t, "Content", errors.Validations[1].Field)
	assert.Equal(t, "必須です。", errors.Validations[1].Message)
}
