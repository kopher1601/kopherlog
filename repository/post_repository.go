package repository

import (
	"context"
	"github.com/gin-gonic/gin"
	"kopherlog/db"
	"kopherlog/domain"
	"kopherlog/ent"
	"kopherlog/ent/post"
	"log"
)

type PostRepository interface {
	Save(ctx context.Context, post *domain.Post) error
	FindAll() ([]*ent.Post, error)
	DeleteAll(ctx context.Context) error
	FindByID(ctx *gin.Context, id int) (*ent.Post, error)
}

type postRepository struct {
	ent *ent.Client
}

func (p *postRepository) DeleteAll(ctx context.Context) error {
	_, err := p.ent.Post.Delete().Exec(ctx)
	if err != nil {
		return err
	}
	return nil
}

func NewPostRepository(ent *ent.Client) PostRepository {
	return &postRepository{ent: ent}
}

func (p *postRepository) Save(ctx context.Context, post *domain.Post) error {
	return db.WithTx(ctx, p.ent, func(tx *ent.Tx) error {
		_, err := tx.Post.Create().SetTitle(post.Title).SetContent(post.Content).Save(ctx)
		if err != nil {
			return err
		}
		return nil
	})
}

func (p *postRepository) FindAll() ([]*ent.Post, error) {
	posts, err := p.ent.Post.Query().All(context.Background())
	if err != nil {
		return nil, err
	}
	return posts, nil
}

func (p *postRepository) FindByID(ctx *gin.Context, id int) (*ent.Post, error) {
	foundPost, err := p.ent.Post.Query().Where(post.ID(id)).First(ctx)
	if err != nil {
		log.Println("postRepository.FindByID:", err)
		return nil, err
	}
	return foundPost, nil
}
