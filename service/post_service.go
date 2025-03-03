package service

import (
	"kopherlog/domain"
	"kopherlog/repository"
)

type PostService interface {
	Write(postCreate *domain.PostCreate) error
}

type postService struct {
	postRepository repository.PostRepository
}

func NewPostService(postRepository repository.PostRepository) PostService {
	return &postService{postRepository: postRepository}
}

func (p *postService) Write(postCreate *domain.PostCreate) error {
	post := &domain.Post{Title: postCreate.Title, Content: postCreate.Content}
	err := p.postRepository.Save(post)
	if err != nil {
		return err
	}
	return nil
}
