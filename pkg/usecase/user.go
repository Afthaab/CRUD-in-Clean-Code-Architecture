package usecase

import (
	"context"

	"github.com/todo/clean/pkg/domain"
	interfaces "github.com/todo/clean/pkg/repository/interface"
	service "github.com/todo/clean/pkg/usecase/interface"
)

type userUseCase struct {
	userRepo interfaces.UserRespository
}

func NewUserUseCase(repo interfaces.UserRespository) service.UserUseCase {
	return &userUseCase{
		userRepo: repo,
	}
}

func (c *userUseCase) FindAll(ctx context.Context) ([]domain.User, error) {
	users, err := c.userRepo.FindAll(ctx)
	return users, err
}
func (c *userUseCase) Save(ctx context.Context, user domain.User) (domain.User, error) {
	user, err := c.userRepo.Save(ctx, user)
	return user, err
}
func (c *userUseCase) Delete(ctx context.Context, id uint) error {
	err := c.userRepo.Delete(ctx, id)
	return err
}
func (c *userUseCase) Find(ctx context.Context, id uint) (domain.User, error) {
	user, err := c.userRepo.Find(ctx, id)
	return user, err
}
