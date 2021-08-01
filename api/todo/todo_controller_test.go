package todo

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gomagedon/expectate"
	"github.com/google/uuid"
	util "github.com/kutay-celebi/gotodo/util"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestTodoController(t *testing.T) {
	t.Run("No todo found", func(t *testing.T) {
		controller := NewController(nil)
		router := util.GetRouter()
		router.GET(util.TodoList, controller.List)

		req := httptest.NewRequest(http.MethodGet, util.TodoList, nil)
		rr := httptest.NewRecorder()

		router.ServeHTTP(rr, req)

		expect := expectate.Expect(t)
		resp := rr.Body.String()
		results := []Todo{}
		json.Unmarshal([]byte(resp), &results)
		expect(len(results)).ToEqual(0)
	})

	t.Run("Find all todo", func(t *testing.T) {
		todoList := []Todo{{uuid.New().String(), "Test todo", "description", false}}
		var controller = Controller{TestRepository{TodoList: &todoList}}
		router := util.GetRouter()
		router.GET(util.TodoList, controller.List)

		req := httptest.NewRequest(http.MethodGet, util.TodoList, nil)
		rr := httptest.NewRecorder()

		router.ServeHTTP(rr, req)

		expect := expectate.Expect(t)
		resp := rr.Body.String()
		var results []Todo
		json.Unmarshal([]byte(resp), &results)
		expect(len(results) > 0).ToEqual(true)
	})

	t.Run("Create todo", func(t *testing.T) {
		var controller = Controller{TestRepository{TodoList: &[]Todo{}}}
		router := util.GetRouter()
		router.POST(util.CreateTodo, controller.Create)

		body, _ := json.Marshal(Todo{uuid.New().String(), "Test todo", "description", false})
		req := httptest.NewRequest(http.MethodPost, util.CreateTodo, bytes.NewReader(body))
		rr := httptest.NewRecorder()

		router.ServeHTTP(rr, req)

		expect := expectate.Expect(t)
		resp := rr.Body.String()
		result := Todo{}
		json.Unmarshal([]byte(resp), &result)
		expect(result).NotToEqual(nil)
	})

	t.Run("The wrong todo object should not be given.", func(t *testing.T) {
		//todoList := []Todo{{uuid.New().String(), "Test todo", "description"}}
		var controller = Controller{TestRepository{TodoList: &[]Todo{}}}
		router := util.GetRouter()
		router.POST(util.CreateTodo, controller.Create)

		body, _ := json.Marshal("wrong todos")
		req := httptest.NewRequest(http.MethodPost, util.CreateTodo, bytes.NewReader(body))
		rr := httptest.NewRecorder()

		router.ServeHTTP(rr, req)

		expect := expectate.Expect(t)
		expect(rr.Code).ToEqual(http.StatusBadRequest)
	})

	t.Run("Complete todo", func(t *testing.T) {
		todoList := []Todo{{uuid.New().String(), "Test todo", "description", false}}
		var controller = Controller{TestRepository{TodoList: &todoList}}
		router := util.GetRouter()
		router.GET(util.CompleteTodo+util.IdParam, controller.Complete)

		req := httptest.NewRequest(http.MethodGet, util.CompleteTodo+"/"+todoList[0].ID, nil)
		rr := httptest.NewRecorder()

		router.ServeHTTP(rr, req)

		expect := expectate.Expect(t)
		resp := rr.Body.String()
		result := Todo{}
		json.Unmarshal([]byte(resp), &result)
		expect(result).NotToEqual(nil)
		expect(result.Completed).ToEqual(true)
	})

	t.Run("Path variable validation error", func(t *testing.T) {
		todoList := []Todo{{uuid.New().String(), "Test todo", "description", false}}
		var controller = Controller{TestRepository{TodoList: &todoList}}
		router := util.GetRouter()
		router.GET(util.CompleteTodo+util.IdParam, controller.Complete)
		router.GET(util.CompleteTodo, controller.Complete)

		req := httptest.NewRequest(http.MethodGet, util.CompleteTodo, nil)
		rr := httptest.NewRecorder()

		router.ServeHTTP(rr, req)

		expect := expectate.Expect(t)
		expect(rr.Code).ToEqual(400)
	})

	t.Run("Error updating non-existent record", func(t *testing.T) {
		todoList := []Todo{{uuid.New().String(), "Test todo", "description", false}}
		var controller = Controller{TestRepository{TodoList: &todoList}}
		router := util.GetRouter()
		router.GET(util.CompleteTodo+util.IdParam, controller.Complete)
		router.GET(util.CompleteTodo, controller.Complete)

		req := httptest.NewRequest(http.MethodGet, util.CompleteTodo+"/555", nil)
		rr := httptest.NewRecorder()

		router.ServeHTTP(rr, req)

		expect := expectate.Expect(t)
		expect(rr.Code).ToEqual(500)
	})
}

// Mocks
type TestRepository struct {
	TodoList *[]Todo
}

func (r TestRepository) FindAll() (*[]Todo, error) {
	return r.TodoList, nil
}

func (r TestRepository) Save(todo *Todo) (*Todo, error) {
	hello := append(*r.TodoList, *todo)
	r.TodoList = &hello
	return todo, nil
}

func (r TestRepository) Update(todo *Todo) (*Todo, error) {

	if *r.TodoList != nil && todo != nil {
		for i, v := range *r.TodoList {
			if v.ID == todo.ID {
				(*r.TodoList)[i] = *todo
			}
		}
	}
	return todo, nil
}

func (r TestRepository) FindById(param string) (*Todo, error) {

	if *r.TodoList != nil && len(param) >= 0 {
		for i, v := range *r.TodoList {
			if v.ID == param {
				return &(*r.TodoList)[i], nil
			}
		}
	}

	return nil, fmt.Errorf("todo not found")
}
