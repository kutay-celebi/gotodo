package todo

import (
	"bytes"
	"encoding/json"
	"github.com/gomagedon/expectate"
	"github.com/google/uuid"
	util "github.com/kutay-celebi/gotodo/util"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetTodos(t *testing.T) {
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
		todoList := []Todo{{uuid.New(), "Test todo", "description"}}
		var controller = Controller{TestRepository{TodoList: &todoList}}
		router := util.GetRouter()
		router.GET(util.TodoList, controller.List)

		req := httptest.NewRequest(http.MethodGet, util.TodoList, nil)
		rr := httptest.NewRecorder()

		router.ServeHTTP(rr, req)

		expect := expectate.Expect(t)
		resp := rr.Body.String()
		results := []Todo{}
		json.Unmarshal([]byte(resp), &results)
		expect(len(results)).NotToEqual(0)
	})

	t.Run("Create todo", func(t *testing.T) {
		//todoList := []Todo{{uuid.New(), "Test todo", "description"}}
		var controller = Controller{TestRepository{TodoList: &[]Todo{}}}
		router := util.GetRouter()
		router.GET(util.CreateTodo, controller.Create)

		body, _ := json.Marshal(Todo{uuid.New(), "Test todo", "description"})
		req := httptest.NewRequest(http.MethodGet, util.CreateTodo, bytes.NewReader(body))
		rr := httptest.NewRecorder()

		router.ServeHTTP(rr, req)

		expect := expectate.Expect(t)
		resp := rr.Body.String()
		result := Todo{}
		json.Unmarshal([]byte(resp), &result)
		expect(result).NotToEqual(nil)
	})
}

// Mocks
type TestRepository struct {
	TodoList *[]Todo
}

func (r TestRepository) FindAll() (*[]Todo, error) {
	return r.TodoList, nil
}

func (r TestRepository) Create(todo Todo) (Todo, error) {
	hello := append(*r.TodoList, todo)
	r.TodoList = &hello
	return todo, nil
}
