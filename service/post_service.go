package service

import (
	"context"
	"kopherlog/domain"
	"kopherlog/repository"
)

type PostService interface {
	Write(ctx context.Context, postCreate *domain.PostCreate) error
	Get(ctx context.Context, id int) (*domain.PostResponse, error)
}

type postService struct {
	postRepository repository.PostRepository
}

func NewPostService(postRepository repository.PostRepository) PostService {
	return &postService{postRepository: postRepository}
}

func (p *postService) Write(ctx context.Context, postCreate *domain.PostCreate) error {
	post := &domain.Post{Title: postCreate.Title, Content: postCreate.Content}
	_, err := p.postRepository.Save(ctx, post)
	if err != nil {
		return err
	}
	return nil
}

func (p *postService) Get(ctx context.Context, id int) (*domain.PostResponse, error) {
	post, err := p.postRepository.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}
	return &domain.PostResponse{
		Title:   post.Title,
		Content: post.Content,
	}, nil
}
