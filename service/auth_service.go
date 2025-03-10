package service

import (
	"context"
	"kopherlog/domain"
	"kopherlog/repository"
)

type AuthService interface {
	SignIn(ctx context.Context, signin *domain.SignIn) (string, error)
}

type authService struct {
	repo repository.UserRepository
}

func NewAuthService(repo repository.UserRepository) AuthService {
	return &authService{repo: repo}
}

func (a authService) SignIn(ctx context.Context, signin *domain.SignIn) (string, error) {
	foundUser, err := a.repo.FindByEmailAndPassword(ctx, signin)
	if err != nil {
		return "", err
	}
	return foundUser.Email, nil
}
