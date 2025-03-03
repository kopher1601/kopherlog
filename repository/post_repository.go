package repository

import (
	"context"
	"kopherlog/domain"
	"kopherlog/ent"
)

type PostRepository interface {
	Save(post *domain.Post) error
	FindAll() ([]*ent.Post, error)
}

type postRepository struct {
	ent *ent.Client
}

func NewPostRepository(ent *ent.Client) PostRepository {
	return &postRepository{ent: ent}
}

func (p *postRepository) Save(post *domain.Post) error {
	_, err := p.ent.Post.Create().SetTitle(post.Title).SetContent(post.Content).Save(context.Background())
	if err != nil {
		return err
	}
	return nil
}

func (p *postRepository) FindAll() ([]*ent.Post, error) {
	posts, err := p.ent.Post.Query().All(context.Background())
	if err != nil {
		return nil, err
	}
	return posts, nil
}
