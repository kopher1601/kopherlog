package service

import (
	"context"
	_ "github.com/go-sql-driver/mysql"
	"github.com/stretchr/testify/assert"
	"kopherlog/config"
	"kopherlog/domain"
	"kopherlog/repository"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	config.Initialize()
	code := m.Run()
	os.Exit(code)
}

func TestPostService_Write(t *testing.T) {
	client := config.SetupDB(t)
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

	err := postService.Write(ctx, postCreate)
	assert.NoError(t, err)
	posts, _ := postRepository.FindAll()
	assert.Equal(t, 1, len(posts))
	assert.Equal(t, "吉祥寺マンション", posts[0].Title)
}
