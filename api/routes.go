package main

import (
	"github.com/gin-gonic/gin"
	"github.com/kutay-celebi/gotodo/api/todo"
	"github.com/kutay-celebi/gotodo/api/util"
)

func InitializeRoutes(router *gin.Engine, todoController *todo.Controller) {
	router.Use(responseMiddleware)

	router.GET(util.TodoList, todoController.List)
	router.POST(util.CreateTodo, todoController.Create)
	router.GET(util.CompleteTodo, todoController.Complete)
	router.GET(util.CompleteTodo+util.IdParam, todoController.Complete)
}

/**
Access-Control-Allow-Origin should be added at the end of each request for axios.
*/
func responseMiddleware(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")
}
