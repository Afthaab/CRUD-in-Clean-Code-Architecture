package repository

import (
	"context"
	"fmt"

	"github.com/todo/clean/pkg/domain"
	interfaces "github.com/todo/clean/pkg/repository/interface"
	"gorm.io/gorm"
)

type userDatabase struct {
	DB *gorm.DB
}

func NewUserRepository(DB *gorm.DB) interfaces.UserRespository {
	return &userDatabase{DB: DB}
}
func (c *userDatabase) FindAll(ctx context.Context) ([]domain.User, error) {
	var users []domain.User
	err := c.DB.Find(&users).Error

	return users, err
}

func (c *userDatabase) Save(ctx context.Context, user domain.User) (domain.User, error) {
	err := c.DB.Save(&user).Error
	return user, err
}
func (c *userDatabase) Delete(ctx context.Context, id uint) error {
	var user domain.User
	fmt.Println("Hi")
	err := c.DB.Raw("delete from users where id = ?", id).Scan(&user).Error
	return err
}
func (c *userDatabase) Find(ctx context.Context, id uint) (domain.User, error) {
	var user domain.User
	err := c.DB.First(&user, id).Error
	return user, err
}
