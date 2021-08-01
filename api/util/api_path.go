package util

const (
	baseApi = "/api"
	IdParam = "/:id"

	todo         = baseApi + "/todo"
	TodoList     = todo + "/list"
	CreateTodo   = todo + "/create"
	CompleteTodo = todo + "/complete"
)
