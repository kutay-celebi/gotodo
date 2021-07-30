package todo

import (
	"github.com/gin-gonic/gin"
)

type Controller struct {
	repository Repository
}

func NewController(repository Repository) *Controller {
	return &Controller{repository}
}

func (controller *Controller) List(c *gin.Context) {
	todos, _ := controller.repository.FindAll()
	c.JSON(200, todos)
}
