package usecase

import "k3_api/internal/domain/models"

type 

type GetUserUsecase struct{}

func NewGetUserUsecase() *GetUserUsecase{
	return &GetUserUsecase{}
}

func (uc *GetUserUsecase) Execute(*[]models.User){}

