package controller

import (
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"kopherlog/domain"
	"net/http"
	"net/http/httptest"
	"testing"
)

func Test_PostController_Post(t *testing.T) {
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

	postController := NewPostController()
	r.POST("/posts", postController.PostCreate)
	r.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusOK, resp.Code)
	assert.Equal(t, "Hello World", resp.Body.String())
}

func Test_PostController_PostCreate_Title_Required(t *testing.T) {
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

	postController := NewPostController()
	r.POST("/posts", postController.PostCreate)
	r.ServeHTTP(resp, req)

	var errors []domain.ErrorResponse
	_ = json.NewDecoder(resp.Body).Decode(&errors)
	assert.Equal(t, http.StatusBadRequest, resp.Code)
	assert.Equal(t, "Title", errors[0].Field)
	assert.Equal(t, "必須です。", errors[0].Message)
}

func Test_PostController_PostCreate_Content_Required(t *testing.T) {
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

	postController := NewPostController()
	r.POST("/posts", postController.PostCreate)
	r.ServeHTTP(resp, req)

	var errors []domain.ErrorResponse
	_ = json.NewDecoder(resp.Body).Decode(&errors)
	assert.Equal(t, http.StatusBadRequest, resp.Code)
	assert.Equal(t, "Content", errors[0].Field)
	assert.Equal(t, "必須です。", errors[0].Message)
}

func Test_PostController_PostCreate_Title_Content_Required(t *testing.T) {
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

	postController := NewPostController()
	r.POST("/posts", postController.PostCreate)
	r.ServeHTTP(resp, req)

	var errors []domain.ErrorResponse
	_ = json.NewDecoder(resp.Body).Decode(&errors)
	assert.Equal(t, http.StatusBadRequest, resp.Code)
	assert.Equal(t, "Title", errors[0].Field)
	assert.Equal(t, "必須です。", errors[0].Message)
	assert.Equal(t, "Content", errors[1].Field)
	assert.Equal(t, "必須です。", errors[1].Message)
}
