package repository

import (
	"context"
	"k3_api/internal/domain/model"
)

type UserRepository interface {
	Save(ctx context.Context, user *model.User) error
}
