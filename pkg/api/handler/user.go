package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
	"github.com/todo/clean/pkg/domain"
	service "github.com/todo/clean/pkg/usecase/interface"
)

type UserHandler struct {
	userUseCase service.UserUseCase
}
type Response struct {
	ID   uint   `copier:"must"`
	Name string `copier:"must"`
	Age  uint   `copier:"must"`
}

func NewUserHandler(usecase service.UserUseCase) *UserHandler {
	return &UserHandler{
		userUseCase: usecase,
	}
}

func (cr *UserHandler) FindAll(c *gin.Context) {
	users, err := cr.userUseCase.FindAll(c.Request.Context())
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		reponse := []Response{}
		copier.Copy(&reponse, &users)
		c.JSON(http.StatusOK, reponse)
	}
}
func (cr *UserHandler) Save(c *gin.Context) {
	var user domain.User
	if err := c.Bind(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Error": err,
		})
		return
	}
	user, err := cr.userUseCase.Save(c.Request.Context(), user)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		response := []Response{}
		copier.Copy(&response, &user)
		c.JSON(http.StatusOK, response)
	}
}

func (cr *UserHandler) DirectDelete(c *gin.Context) {
	id, err := strconv.Atoi(c.Query("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Errorr": err,
		})
		return
	}
	err1 := cr.userUseCase.Delete(c.Request.Context(), uint(id))
	if err1 != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Error": err1,
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"Message": "Success",
		})
	}

}
func (cr *UserHandler) Find(c *gin.Context) {
	id, err := strconv.Atoi(c.Query("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Errorr": err,
		})
		return
	}
	var user domain.User
	user, err1 := cr.userUseCase.Find(c.Request.Context(), uint(id))
	if err1 != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"Error": err1,
		})
		return
	}
	response := []Response{}
	copier.Copy(&response, &user)
	c.JSON(http.StatusOK, gin.H{
		"Message": "Success",
		"User":    response,
	})

}
