package main

import (
	"fmt"
	"log"

	"github.com/todo/clean/pkg/config"
	"github.com/todo/clean/pkg/di"
)

func main() {
	fmt.Println("Hi")
	config, configerr := config.LoadConfig()
	if configerr != nil {
		log.Fatal("Cannot load config: ", configerr)
	}
	server, diErr := di.InitializeAPI(config)
	if diErr != nil {
		log.Fatal("cannot start server: ", diErr)
	} else {
		server.Start()
	}

}
