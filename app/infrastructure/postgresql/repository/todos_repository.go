package repository

import (
	"app/app/domain/repository"
	"app/app/infrastructure/postgresql/models"
	"context"
	"database/sql"

	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

type TodoRepository struct {
	*sql.DB
}

func NewTodoRepository(db *sql.DB) repository.ITodoRepository {
	return &TodoRepository{db}
}

func (u *TodoRepository) FindAll(ctx context.Context) (models.TodoSlice, error) {
	return models.Todos(qm.OrderBy("id ASC")).All(ctx, u.DB)
}

func (u *TodoRepository) Create(ctx context.Context, todo *models.Todo) error {
	return todo.Insert(ctx, u.DB, boil.Infer())
}

func (u *TodoRepository) Update(ctx context.Context, todo *models.Todo) (int64, error) {
	return todo.Update(ctx, u.DB, boil.Infer())
}

func (u *TodoRepository) Delete(ctx context.Context, id int) (int64, error) {
	return models.Todos(models.TodoWhere.ID.EQ(id)).DeleteAll(ctx, u.DB)
}