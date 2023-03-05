package usecase

import (
	"app/app/domain/service"
	"app/app/infrastructure/postgresql/models"
	"context"
)

type IUserUsecase interface {
	FindByID(ctx context.Context, id int) (*models.User, error)
	Create(ctx context.Context, user *models.User) error
	Login(ctx context.Context, name string, password string) (bool, error)
}

type userUsecase struct {
	svc service.IUserService
}

func NewUserUsecase(svc service.IUserService) IUserUsecase {
	return &userUsecase{
		svc: svc,
	}
}

func (u *userUsecase) FindByID(ctx context.Context, id int) (*models.User, error) {
	ms, err := u.svc.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}
	return ms, nil
}

func (u *userUsecase) Create(ctx context.Context, user *models.User) error {
	err := u.svc.Create(ctx, user)
	if err != nil {
		return err
	}
	return nil
}

func (u *userUsecase) Login(ctx context.Context, name string, password string) (bool, error) {
	isExists, err := u.svc.Login(ctx, name, password)
	if err != nil {
		return false, err
	}
	return isExists, nil
}
