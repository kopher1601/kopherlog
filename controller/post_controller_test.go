package controller

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
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

var client *ent.Client
var postRepository repository.PostRepository
var postService service.PostService
var tPostController PostController

func TestMain(m *testing.M) {
	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatal("Error loading .env file", err)
	}
	client = config.SetupDB()

	// injection
	postRepository = repository.NewPostRepository(client)
	postService = service.NewPostService(postRepository)
	tPostController = NewPostController(postService)

	code := m.Run()
	client.Close()
	os.Exit(code)
}

func Test_PostController_Post_Save(t *testing.T) {
	// given
	ctx := context.Background()
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

	// when
	r.POST("/posts", tPostController.PostCreate)
	r.ServeHTTP(resp, req)

	// then
	posts, err := postRepository.FindAll(ctx, nil)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusCreated, resp.Code)
	assert.Equal(t, 1, len(posts))

	t.Cleanup(func() {
		postRepository.DeleteAll(ctx)
	})
}

func Test_PostController_PostCreate_Title_Required(t *testing.T) {
	// given
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

	// when
	r.POST("/posts", tPostController.PostCreate)
	r.ServeHTTP(resp, req)

	// then
	var errors domain.ErrorResponse
	_ = json.NewDecoder(resp.Body).Decode(&errors)
	assert.Equal(t, http.StatusBadRequest, resp.Code)
	assert.Equal(t, "Title", errors.Validations[0].Field)
	assert.Equal(t, "必須です。", errors.Validations[0].Message)
}

func Test_PostController_PostCreate_Content_Required(t *testing.T) {
	// given
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

	// when
	r.POST("/posts", tPostController.PostCreate)
	r.ServeHTTP(resp, req)

	// then
	var errors domain.ErrorResponse
	_ = json.NewDecoder(resp.Body).Decode(&errors)
	assert.Equal(t, http.StatusBadRequest, resp.Code)
	assert.Equal(t, "Content", errors.Validations[0].Field)
	assert.Equal(t, "必須です。", errors.Validations[0].Message)
}

func Test_PostController_PostCreate_Title_Content_Required(t *testing.T) {
	// given
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

	// when
	r.POST("/posts", tPostController.PostCreate)
	r.ServeHTTP(resp, req)

	// then
	var errors domain.ErrorResponse
	_ = json.NewDecoder(resp.Body).Decode(&errors)
	assert.Equal(t, http.StatusBadRequest, resp.Code)
	assert.Equal(t, "Title", errors.Validations[0].Field)
	assert.Equal(t, "必須です。", errors.Validations[0].Message)
	assert.Equal(t, "Content", errors.Validations[1].Field)
	assert.Equal(t, "必須です。", errors.Validations[1].Message)
}

func TestPostController_Get(t *testing.T) {
	// given
	r := gin.Default()
	ctx := context.Background()
	post := &domain.Post{
		Title:   "吉祥寺マンション",
		Content: "吉祥寺マンション購入します。",
	}
	t.Cleanup(func() {
		postRepository.DeleteAll(ctx)
	})
	savedPost, _ := postRepository.Save(ctx, post)

	resp := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPost, fmt.Sprintf("/posts/%d", savedPost.ID), nil)
	req.Header.Set("Content-Type", "application/json")

	// when
	r.POST("/posts/:postID", tPostController.Get)
	r.ServeHTTP(resp, req)

	// then
	var postResponse domain.PostResponse
	_ = json.NewDecoder(resp.Body).Decode(&postResponse)
	assert.Equal(t, http.StatusOK, resp.Code)
	assert.Equal(t, "吉祥寺マンション", postResponse.Title)
	assert.Equal(t, "吉祥寺マンション購入します。", postResponse.Content)
}

func TestPostController_GetAll(t *testing.T) {
	// given
	r := gin.Default()
	ctx := context.Background()
	t.Cleanup(func() {
		postRepository.DeleteAll(ctx)
	})
	var postCreates []*domain.PostCreate
	for i := 0; i < 30; i++ {
		postCreates = append(postCreates, &domain.PostCreate{
			Title:   fmt.Sprintf("吉祥寺マンション %d", i),
			Content: fmt.Sprintf("吉祥寺マンション購入します。 %d", i),
		})
	}
	postRepository.SaveAll(ctx, postCreates)

	resp := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/posts?page=1&size=10", nil)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	// when
	r.GET("/posts", tPostController.GetAll)
	r.ServeHTTP(resp, req)

	// then
	log.Println(resp)
	var postResponse []*domain.PostResponse
	_ = json.NewDecoder(resp.Body).Decode(&postResponse)
	assert.Equal(t, http.StatusOK, resp.Code)
	assert.Equal(t, "吉祥寺マンション 29", postResponse[0].Title)
	assert.Equal(t, "吉祥寺マンション購入します。 29", postResponse[0].Content)
	assert.Equal(t, "吉祥寺マンション 28", postResponse[1].Title)
	assert.Equal(t, "吉祥寺マンション購入します。 28", postResponse[1].Content)
}
