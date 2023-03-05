package repository

import (
	"app/app/domain/repository"
	"app/app/infrastructure/postgresql/models"
	"context"
	"database/sql"

	"github.com/volatiletech/sqlboiler/v4/boil"
)

type UserRepository struct {
	*sql.DB
}

func NewUserRepository(db *sql.DB) repository.IUserRepository {
	return &UserRepository{db}
}

func (u *UserRepository) FindByID(ctx context.Context, id int) (*models.User, error) {
	return models.Users(models.UserWhere.ID.EQ(id)).One(ctx, u.DB)
}

func (u *UserRepository) Create(ctx context.Context, user *models.User) error {
	return user.Insert(ctx, u.DB, boil.Infer())
}

func (u *UserRepository) Login(ctx context.Context, name string, password string) (bool, error) {
	return models.Users(models.UserWhere.Name.EQ(name), models.UserWhere.Password.EQ(password)).Exists(ctx, u.DB)
}
