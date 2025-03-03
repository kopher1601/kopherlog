package controller

import (
	"bytes"
	"context"
	"encoding/json"
	"entgo.io/ent/dialect/sql"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/stretchr/testify/assert"
	db "kopherlog/db"
	"kopherlog/domain"
	"kopherlog/ent"
	"kopherlog/ent/enttest"
	"kopherlog/repository"
	"kopherlog/service"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
)

func connectDB() *sql.Driver {
	database, err := sql.Open("mysql", "kopherlog:kopherlog@tcp(127.0.0.1:3306)/kopherlog?parseTime=true")
	if err != nil {
		log.Fatal(err)
	}
	return database
}

func getClient(t *testing.T, db *sql.Driver) *ent.Client {
	drv := sql.OpenDB("mysql", db.DB())
	options := []enttest.Option{
		enttest.WithOptions(ent.Log(t.Log), ent.Driver(drv)),
	}
	client := enttest.NewClient(t, options...)
	return client
}

func Test_PostController_Post_Save(t *testing.T) {
	context := context.Background()
	client := getClient(t, connectDB())
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

	db.WithTx(context, client, func(tx *ent.Tx) error {
		postRepository := repository.NewPostRepository(tx.Client())
		postService := service.NewPostService(postRepository)
		postController := NewPostController(postService)
		r.POST("/posts", postController.PostCreate)
		r.ServeHTTP(resp, req)
		posts, err := postRepository.FindAll()
		assert.NoError(t, err)
		assert.Equal(t, http.StatusCreated, resp.Code)
		assert.Equal(t, 1, len(posts))
		return nil
	})

}

func Test_PostController_PostCreate_Title_Required(t *testing.T) {
	client := getClient(t, connectDB())

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
	client := getClient(t, connectDB())
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
	client := getClient(t, connectDB())
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
