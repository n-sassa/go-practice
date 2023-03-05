package repository

import (
	"app/app/infrastructure/postgresql/models"
	"context"
)

type IUserRepository interface {
	FindByID(ctx context.Context, id int) (*models.User, error)
	Create(ctx context.Context, user *models.User) error
	Login(ctx context.Context, name string, password string) (bool, error)
}
