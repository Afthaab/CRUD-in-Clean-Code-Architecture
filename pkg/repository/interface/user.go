package interfaces

import (
	"context"

	"github.com/todo/clean/pkg/domain"
)

type UserRespository interface {
	FindAll(ctx context.Context) ([]domain.User, error)
	Save(ctx context.Context, user domain.User) (domain.User, error)
	Delete(ctx context.Context, id uint) error
	Find(ctx context.Context, id uint) (domain.User, error)
}
