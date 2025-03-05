package service

import (
	"context"
	"kopherlog/domain"
	"kopherlog/repository"
)

type PostService interface {
	Write(ctx context.Context, postCreate *domain.PostCreate) error
}

type postService struct {
	ctx            context.Context
	postRepository repository.PostRepository
}

func NewPostService(ctx context.Context, postRepository repository.PostRepository) PostService {
	return &postService{ctx: ctx, postRepository: postRepository}
}

func (p *postService) Write(ctx context.Context, postCreate *domain.PostCreate) error {
	post := &domain.Post{Title: postCreate.Title, Content: postCreate.Content}
	err := p.postRepository.Save(ctx, post)
	if err != nil {
		return err
	}
	return nil
}
