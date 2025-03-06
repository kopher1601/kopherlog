package repository

import (
	"context"
	"kopherlog/db"
	"kopherlog/domain"
	"kopherlog/ent"
	"kopherlog/ent/post"
	"log"
)

type PostRepository interface {
	Save(ctx context.Context, post *domain.Post) (*domain.Post, error)
	FindAll() ([]*ent.Post, error)
	DeleteAll(ctx context.Context) error
	FindByID(ctx context.Context, id int) (*ent.Post, error)
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

func (p *postRepository) Save(ctx context.Context, post *domain.Post) (*domain.Post, error) {
	var response domain.Post
	err := db.WithTx(ctx, p.ent, func(tx *ent.Tx) error {
		savedPost, err := tx.Post.Create().SetTitle(post.Title).SetContent(post.Content).Save(ctx)
		if err != nil {
			return err
		}
		response = domain.Post{
			ID:      savedPost.ID,
			Title:   savedPost.Title,
			Content: savedPost.Content,
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return &response, nil
}

func (p *postRepository) FindAll() ([]*ent.Post, error) {
	posts, err := p.ent.Post.Query().All(context.Background())
	if err != nil {
		return nil, err
	}
	return posts, nil
}

func (p *postRepository) FindByID(ctx context.Context, id int) (*ent.Post, error) {
	foundPost, err := p.ent.Post.Query().Where(post.ID(id)).First(ctx)
	if err != nil {
		log.Println("postRepository.FindByID:", err)
		return nil, err
	}
	return foundPost, nil
}
