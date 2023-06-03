package api

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/todo/clean/pkg/api/handler"
)

type ServerHttp struct {
	engine *gin.Engine
}

func NewServerHttp(userHandler *handler.UserHandler) *ServerHttp {
	engine := gin.New()

	//use logger from gin
	engine.Use(gin.Logger())

	engine.GET("user/findall", userHandler.FindAll)
	engine.POST("user/save", userHandler.Save)
	engine.DELETE("user/direct/delete", userHandler.DirectDelete)
	engine.GET("user/direct/find", userHandler.Find)

	return &ServerHttp{
		engine: engine,
	}
}

func (ser *ServerHttp) Start() {
	fmt.Println("Server Started at :1234")
	ser.engine.Run(":1234")
}
