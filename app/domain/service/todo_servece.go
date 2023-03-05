package service

import (
	"app/app/domain/repository"
	"app/app/infrastructure/postgresql/models"
	"context"
)

type ITodoService interface {
	FindAll(ctx context.Context) (models.TodoSlice, error)
	Create(ctx context.Context, todo *models.Todo) error
	Update(ctx context.Context, todo *models.Todo) (int64, error)
	Delete(ctx context.Context, id int) (int64, error)
}

type todoService struct {
	repo repository.ITodoRepository
}

func NewTodoService(repo repository.ITodoRepository) ITodoService {
	return &todoService{repo: repo}
}

func (us *todoService) FindAll(ctx context.Context) (models.TodoSlice, error) {
	return us.repo.FindAll(ctx)
}

func (us *todoService) Create(ctx context.Context, todo *models.Todo) error {
	return us.repo.Create(ctx, todo)
}

func (us *todoService) Update(ctx context.Context, todo *models.Todo) (int64, error) {
	return us.repo.Update(ctx, todo)
}

func (us *todoService) Delete(ctx context.Context, id int) (int64, error) {
	return us.repo.Delete(ctx, id)
}
