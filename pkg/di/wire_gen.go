// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package di

import (
	"github.com/todo/clean/pkg/api"
	"github.com/todo/clean/pkg/api/handler"
	"github.com/todo/clean/pkg/config"
	"github.com/todo/clean/pkg/db"
	"github.com/todo/clean/pkg/repository"
	"github.com/todo/clean/pkg/usecase"
)

// Injectors from wire.go:

func InitializeAPI(cfg config.Config) (*api.ServerHttp, error) {
	gormDB, err := db.ConnectToDataBase(cfg)
	if err != nil {
		return nil, err
	}
	userRespository := repository.NewUserRepository(gormDB)
	userUseCase := usecase.NewUserUseCase(userRespository)
	userHandler := handler.NewUserHandler(userUseCase)
	serverHttp := api.NewServerHttp(userHandler)
	return serverHttp, nil
}
