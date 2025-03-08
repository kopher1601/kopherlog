package service

import (
	"context"
	"fmt"
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
	posts, _ := postRepository.FindAll(ctx, nil)
	assert.Equal(t, 1, len(posts))
	assert.Equal(t, "吉祥寺マンション", posts[0].Title)
}

func TestPostService_Get(t *testing.T) {
	// given
	ctx := context.Background()
	postRepository := repository.NewPostRepository(client)
	postService := NewPostService(postRepository)

	var postCreates []*domain.PostCreate
	for i := 0; i < 30; i++ {
		postCreates = append(postCreates, &domain.PostCreate{
			Title:   fmt.Sprintf("吉祥寺マンション %d", i),
			Content: fmt.Sprintf("吉祥寺マンション購入します。 %d", i),
		})
	}
	postRepository.SaveAll(ctx, postCreates)

	posts, _ := postRepository.FindAll(ctx, nil)
	t.Cleanup(func() {
		postRepository.DeleteAll(ctx)
	})

	// when
	foundPost, err := postService.Get(ctx, posts[0].ID)

	// then
	assert.NoError(t, err)
	assert.Equal(t, "吉祥寺マンション 29", foundPost.Title)
	assert.Equal(t, "吉祥寺マンション購入します。 29", foundPost.Content)
}

func TestPostService_GetAll_FirstPage(t *testing.T) {
	// given
	ctx := context.Background()
	postRepository := repository.NewPostRepository(client)
	postService := NewPostService(postRepository)
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

	// then
	search := &domain.PostSearch{
		Page: 2,
		Size: 10,
	}
	posts, _ := postService.GetAll(ctx, search)

	// then
	assert.Equal(t, "吉祥寺マンション 19", posts[0].Title)
	assert.Equal(t, "吉祥寺マンション購入します。 19", posts[0].Content)
}
