package todo

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/kutay-celebi/gotodo/api/util"
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
		c.JSON(http.StatusInternalServerError, util.ErrorResponse{Message: "An error occured", ErrorUUID: uuid.New().String()})
		return
	}

	c.JSON(http.StatusOK, todos)
}

func (controller *Controller) Create(c *gin.Context) {
	var data Todo
	err := c.ShouldBindJSON(&data)

	if err != nil {
		c.JSON(http.StatusBadRequest, util.ErrorResponse{Message: "record could not create", ErrorUUID: uuid.New().String()})
		return
	}

	createdTodo, _ := controller.repository.Save(&data)
	c.JSON(http.StatusOK, createdTodo)
}

func (controller *Controller) Complete(c *gin.Context) {
	param := c.Param("id")

	if len(param) <= 0 {
		c.JSON(http.StatusBadRequest, util.ErrorResponse{Message: "record could not create", ErrorUUID: uuid.New().String()})
		return
	}

	findTodo, err := controller.repository.FindById(param)

	if err != nil {
		c.JSON(http.StatusInternalServerError, util.ErrorResponse{Message: "record not found", ErrorUUID: uuid.New().String()})
		return
	}

	findTodo.Completed = true
	updatedTodo, _ := controller.repository.Update(findTodo)
	c.JSON(http.StatusOK, updatedTodo)
}
