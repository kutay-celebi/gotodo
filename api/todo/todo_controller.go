package todo

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Controller struct {
	repository Repository
}

func NewController(repository Repository) *Controller {
	return &Controller{repository}
}

func (controller *Controller) List(c *gin.Context) {
	todos, err := controller.repository.FindAll()

	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
	}

	c.JSON(http.StatusOK, todos)
}

func (controller *Controller) Create(c *gin.Context) {
	var data Todo
	err := c.ShouldBindJSON(&data)

	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
	}

	createdTodo, _ := controller.repository.Create(data)
	c.JSON(http.StatusOK, createdTodo)
}
