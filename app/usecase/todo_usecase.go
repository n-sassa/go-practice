package usecase

import (
	"app/app/domain/service"
	"app/app/infrastructure/postgresql/models"
	"context"
	"fmt"
)

type ITodoUsecase interface {
	FindAll(ctx context.Context) (models.TodoSlice, error)
	Create(ctx context.Context, todo *models.Todo) (*models.Todo, error)
	Update(ctx context.Context, todo *models.Todo) (*models.Todo, error)
	Delete(ctx context.Context, id int) error
}

type todoUsecase struct {
	svc service.ITodoService
}

func NewTodoUsecase(svc service.ITodoService) ITodoUsecase {
	return &todoUsecase{
		svc: svc,
	}
}

func (u *todoUsecase) FindAll(ctx context.Context) (models.TodoSlice, error) {
	ms, err := u.svc.FindAll(ctx)
	if err != nil {
		return nil, err
	}
	return ms, nil
}

func (u *todoUsecase) Create(ctx context.Context, todo *models.Todo) (*models.Todo, error) {
	err := u.svc.Create(ctx, todo)
	if err != nil {
		return nil, err
	}
	fmt.Println(todo)
	return todo, nil
}

func (u *todoUsecase) Update(ctx context.Context, todo *models.Todo) (*models.Todo, error) {
	id, err := u.svc.Update(ctx, todo)
	if err != nil {
		return nil, err
	}
	todo.ID = int(id)
	return todo, nil
}

func (u *todoUsecase) Delete(ctx context.Context, id int) error {
	_, err := u.svc.Delete(ctx, id)
	if err != nil {
		return err
	}
	return nil
}
