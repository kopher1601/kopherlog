package repository

import (
	"context"
	"kopherlog/domain"
	"kopherlog/ent"
	"kopherlog/ent/user"
)

type UserRepository interface {
	FindByEmailAndPassword(ctx context.Context, signin *domain.SignIn) (*domain.User, error)
}

type userRepository struct {
	ent *ent.Client
}

func NewUserRepository(ent *ent.Client) UserRepository {
	return &userRepository{ent: ent}
}

func (u userRepository) FindByEmailAndPassword(ctx context.Context, signin *domain.SignIn) (*domain.User, error) {
	foundUser, err := u.ent.User.Query().Where(user.And(user.EmailEQ(signin.Email), user.PasswordEQ(signin.Password))).First(ctx)
	if err != nil {
		return nil, err
	}
	return &domain.User{
		ID:        foundUser.ID,
		Name:      foundUser.Name,
		Email:     foundUser.Email,
		Password:  foundUser.Password,
		CreatedAt: foundUser.CreatedAt,
	}, nil
}
