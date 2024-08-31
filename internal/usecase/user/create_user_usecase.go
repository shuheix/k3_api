package usecase

import (
	"k3_api/internal/infrastructure/repository"
)

type CreateUserUsecase interface {
	Execute()
}

type createUserUsecase struct{
	repo repository.UserRepository
}

func NewCreateUserUsecase() *CreateUserUsecase {
	return &CreateUserUsecase{}
}

func 