package repository

import (
	"app/app/infrastructure/postgresql/models"
	"context"
)

type ITodoRepository interface {
	FindAll(ctx context.Context) (models.TodoSlice, error)
	Create(ctx context.Context, todo *models.Todo) error
	Update(ctx context.Context, todo *models.Todo) (int64, error)
	Delete(ctx context.Context, id int) (int64, error)
}
