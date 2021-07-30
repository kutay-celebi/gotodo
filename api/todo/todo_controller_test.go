package todo

import (
	"encoding/json"
	"github.com/gomagedon/expectate"
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
		todoList := []Todo{{1, "Test todo", "description"}}
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
}

// Mocks
type TestRepository struct {
	TodoList *[]Todo
}

func (t TestRepository) FindAll() (*[]Todo, error) {
	return t.TodoList, nil
}
