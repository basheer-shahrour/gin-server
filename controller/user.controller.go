package controller

import (
	"net/http"
	"strconv"

	"github.com/basheer-shahrour/gin-server/entity"
	"github.com/basheer-shahrour/gin-server/service"
	"github.com/gin-gonic/gin"
)

type UserController interface {
	FindAll() []entity.User
	Save(ctx *gin.Context) error
	Update(ctx *gin.Context) error
	ShowAll(ctx *gin.Context)
}

type userController struct {
	service service.UserService
}

func NewUserController(service service.UserService) UserController {
	return &userController{service: service}
}

func (c *userController) FindAll() []entity.User {
	return c.service.FindAll()
}

func (c *userController) Save(ctx *gin.Context) error {
	var user entity.User
	err := ctx.ShouldBindJSON(&user)
	if err != nil {
		return err
	}
	c.service.Save(user)
	return nil
}

func (c *userController) ShowAll(ctx *gin.Context) {
	users := c.service.FindAll()
	data := map[string]interface{}{
		"title": "users page",
		"users": users,
	}
	ctx.HTML(http.StatusOK, "index.html", data)
}

func (c *userController) Update(ctx *gin.Context) error {
	var user entity.User
	err := ctx.ShouldBindJSON(&user)
	if err != nil {
		return err
	}

	id, err := strconv.ParseUint(ctx.Param("id"), 0, 0)
	if err != nil {
		return err
	}

	user.Id = id
	c.service.Update(user)
	return nil
}
