package service

import (
	"app/app/domain/repository"
	"app/app/infrastructure/postgresql/models"
	"context"
)

type IUserService interface {
	FindByID(ctx context.Context, id int) (*models.User, error)
	Create(ctx context.Context, user *models.User) error
	// Update(ctx context.Context, user *models.User) error
	// Delete(ctx context.Context, id int) error
}

type userService struct {
	repo repository.IUserRepository
}

func NewUserService(repo repository.IUserRepository) IUserService {
	return &userService{repo: repo}
}

func (us *userService) FindByID(ctx context.Context, id int) (*models.User, error) {
	return us.repo.FindByID(ctx, id)
}

func (us *userService) Create(ctx context.Context, user *models.User) error {
	return us.repo.Create(ctx, user)
}
