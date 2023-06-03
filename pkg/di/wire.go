//go:build wireinject
// +build wireinject

package di

import (
	"github.com/google/wire"
	http "github.com/todo/clean/pkg/api"
	handler "github.com/todo/clean/pkg/api/handler"
	config "github.com/todo/clean/pkg/config"
	db "github.com/todo/clean/pkg/db"
	repo "github.com/todo/clean/pkg/repository"
	usecase "github.com/todo/clean/pkg/usecase"
)

func InitializeAPI(cfg config.Config) (*http.ServerHttp, error) {
	wire.Build(db.ConnectToDataBase,
		repo.NewUserRepository,
		usecase.NewUserUseCase,
		handler.NewUserHandler,
		http.NewServerHttp)

	return &http.ServerHttp{}, nil
}
