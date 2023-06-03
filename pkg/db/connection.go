package db

import (
	"fmt"

	"github.com/todo/clean/pkg/config"
	"github.com/todo/clean/pkg/domain"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectToDataBase(cfg config.Config) (*gorm.DB, error) {
	psqlinfo := fmt.Sprintf("host=%s user=%s dbname=%s port=%s password=%s", cfg.DBHost, cfg.DBUser, cfg.DBName, cfg.DBPort, cfg.DBPassword)
	db, dberr := gorm.Open(postgres.Open(psqlinfo), &gorm.Config{
		SkipDefaultTransaction: true,
	})
	db.AutoMigrate(&domain.User{})
	return db, dberr

}
