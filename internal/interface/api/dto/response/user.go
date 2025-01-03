package response

import "k3_api/internal/domain/model"

type UserResponse struct {
	Id int `json:"id"`
	Name string `json:"name"`
}

func NewUserResponse(user *model.User) *UserResponse {
	return &UserResponse{
		Id: user.Id,
		Name: user.Name,
	}
}
