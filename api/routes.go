package main

import (
	"github.com/gin-gonic/gin"
	"github.com/kutay-celebi/gotodo/todo"
	"github.com/kutay-celebi/gotodo/util"
)

func InitializeRoutes(router *gin.Engine, todoController *todo.Controller) {
	router.GET(util.TodoList, todoController.List)
}
