package usecase

import (
	"context"
	"k3_api/internal/domain/model"
	"k3_api/internal/domain/repository"
	"k3_api/internal/interface/api/dto/command"
)

type CreateUser struct {
	userRepo repository.UserRepository
}

func NewCreateUserUsecase(repo repository.UserRepository) *CreateUser {
	return &CreateUser{userRepo: repo}
}

func (u *CreateUser) Execute(ctx context.Context, cmd api.CreateUserCommand) (*model.User, error) {
	user := model.NewUser(cmd.Name)
	if err := u.userRepo.Save(ctx, user); err != nil {
		return nil, err
	}

	return user, nil
}

