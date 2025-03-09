package service

import (
	"context"
	"kopherlog/domain"
	"kopherlog/repository"
)

type PostService interface {
	Write(ctx context.Context, postCreate *domain.PostCreate) error
	Get(ctx context.Context, id int) (*domain.PostResponse, error)
	GetAll(ctx context.Context, search *domain.PostSearch) ([]*domain.PostResponse, error)
	Edit(ctx context.Context, id int, edit *domain.PostEdit) error
	Delete(ctx context.Context, id int) error
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
		ID:      post.ID,
		Title:   post.Title,
		Content: post.Content,
	}, nil
}

func (p *postService) GetAll(ctx context.Context, search *domain.PostSearch) ([]*domain.PostResponse, error) {
	posts, err := p.postRepository.FindAll(ctx, search)
	if err != nil {
		return nil, err
	}
	var postResponses []*domain.PostResponse
	for _, post := range posts {
		postResponses = append(postResponses, &domain.PostResponse{
			ID:      post.ID,
			Title:   post.Title,
			Content: post.Content,
		})
	}
	return postResponses, nil
}

func (p *postService) Edit(ctx context.Context, id int, edit *domain.PostEdit) error {
	target, err := p.postRepository.FindByID(ctx, id)
	if err != nil {
		return err
	}

	post := &domain.Post{
		Title:   coalesce(edit.Title, target.Title),
		Content: coalesce(edit.Content, target.Content),
	}
	err = p.postRepository.Update(ctx, target, post)
	if err != nil {
		return err
	}
	return nil
}

func coalesce(values ...string) string {
	for _, value := range values {
		if value != "" {
			return value
		}
	}
	return ""
}

func (p *postService) Delete(ctx context.Context, id int) error {
	_, err := p.postRepository.FindByID(ctx, id)
	if err != nil {
		return err
	}
	err = p.postRepository.Delete(ctx, id)
	if err != nil {
		return err
	}
	return nil
}
