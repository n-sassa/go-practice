package repository

import (
	"app/app/infrastructure/postgresql/models"
	"context"
)

type IUserRepository interface {
	FindByID(ctx context.Context, id int) (*models.User, error)
	Create(ctx context.Context, user *models.User) error
	// Update(ctx context.Context, user *models.User) error
	// Delete(ctx context.Context, id int) error
}
