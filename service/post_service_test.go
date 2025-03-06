package service

import (
	"context"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
	"kopherlog/config"
	"kopherlog/domain"
	"kopherlog/ent"
	"kopherlog/repository"
	"log"
	"os"
	"testing"
)

var client *ent.Client

func TestMain(m *testing.M) {
	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatal("Error loading .env file", err)
	}
	client = config.SetupDB()
	code := m.Run()
	client.Close()
	os.Exit(code)
}

func TestPostService_Write(t *testing.T) {
	// given
	ctx := context.Background()
	postRepository := repository.NewPostRepository(client)
	postService := NewPostService(postRepository)
	postCreate := &domain.PostCreate{
		Title:   "吉祥寺マンション",
		Content: "吉祥寺マンション購入します。",
	}
	t.Cleanup(func() {
		postRepository.DeleteAll(ctx)
	})

	// when
	err := postService.Write(ctx, postCreate)

	// then
	assert.NoError(t, err)
	posts, _ := postRepository.FindAll()
	assert.Equal(t, 1, len(posts))
	assert.Equal(t, "吉祥寺マンション", posts[0].Title)
}

func TestPostService_Get(t *testing.T) {
	// given
	ctx := context.Background()
	postRepository := repository.NewPostRepository(client)
	postService := NewPostService(postRepository)
	postCreate := &domain.PostCreate{
		Title:   "吉祥寺マンション",
		Content: "吉祥寺マンション購入します。",
	}
	postService.Write(ctx, postCreate)
	posts, _ := postRepository.FindAll()
	t.Cleanup(func() {
		postRepository.DeleteAll(ctx)
	})

	// when
	foundPost, err := postService.Get(ctx, posts[0].ID)

	// then
	assert.NoError(t, err)
	assert.Equal(t, "吉祥寺マンション", foundPost.Title)
	assert.Equal(t, "吉祥寺マンション購入します。", foundPost.Content)
}
