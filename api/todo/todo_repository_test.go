package todo

import (
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/gomagedon/expectate"
	"github.com/google/uuid"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"regexp"
	"testing"
)

func TestRepositoryFunciton(t *testing.T) {

	mockdb, mock, _ := sqlmock.New()

	db, _ := gorm.Open(postgres.New(postgres.Config{
		DSN:                  "sqlmock_db_0",
		DriverName:           "postgres",
		Conn:                 mockdb,
		PreferSimpleProtocol: true,
	}), &gorm.Config{})

	mockTodo := &Todo{
		ID:          "1",
		Title:       "test",
		Description: "test",
		Completed:   false,
	}

	newMockTodo := &Todo{
		Title:       "test",
		Description: "test",
		Completed:   false,
	}

	t.Run("Return empty list", func(t *testing.T) {

		mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "todos"`)).WillReturnRows(sqlmock.NewRows(nil))

		repository := NewRepositoryImpl(db)
		response, err := repository.FindAll()

		expect := expectate.Expect(t)
		expect(err).ToEqual(nil)
		expect(response).ToEqual(&[]Todo{})
	})

	t.Run("Return todo list", func(t *testing.T) {

		rows := sqlmock.
			NewRows([]string{"id", "title", "description", "completed"}).
			AddRow(mockTodo.ID, mockTodo.Title, mockTodo.Description, mockTodo.Completed)

		mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "todos"`)).WillReturnRows(rows)

		repository := NewRepositoryImpl(db)
		response, err := repository.FindAll()

		expect := expectate.Expect(t)
		expect(err).ToEqual(nil)
		expect(response).ToEqual(&[]Todo{*mockTodo})
	})

	t.Run("Create todo", func(t *testing.T) {
		const sqlInsert = `
					INSERT INTO "todos" ("title","description","completed") 
						VALUES ($1,$2,$3) RETURNING "id"`

		newId := uuid.New().String()

		mock.ExpectBegin() // start transaction
		mock.ExpectQuery(regexp.QuoteMeta(sqlInsert)).
			WithArgs(newMockTodo.Title, newMockTodo.Description, newMockTodo.Completed).
			WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(newId))

		mock.ExpectCommit() // commit transaction

		repository := NewRepositoryImpl(db)
		_, err := repository.Save(newMockTodo)

		expect := expectate.Expect(t)
		expect(err).ToEqual(nil)
	})

	t.Run("Update", func(t *testing.T) {
		const sqlUpdate = `
					UPDATE "todos" SET "title"=$1,"description"=$2,"completed"=$3 WHERE "id" = $4`
		const sqlSelectOne = `
					SELECT * FROM "todos" WHERE "id" = $1 ORDER BY "todos"."id" LIMIT 1`
		newUUID, _ := uuid.NewUUID()
		mockTodo.ID = newUUID.String()
		mock.ExpectBegin()
		mock.ExpectExec(regexp.QuoteMeta(sqlUpdate)).
			WithArgs(mockTodo.Title, mockTodo.Description, mockTodo.Completed, mockTodo.ID).
			WillReturnResult(sqlmock.NewResult(0, 0))
		mock.ExpectCommit()

		// select after update
		mock.ExpectQuery(regexp.QuoteMeta(sqlSelectOne)).
			WithArgs(mockTodo.ID).
			WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(mockTodo.ID))

		repository := NewRepositoryImpl(db)
		_, err := repository.Update(mockTodo)
		expect := expectate.Expect(t)
		expect(err).ToEqual(nil)
	})

	t.Run("Find By ID", func(t *testing.T) {
		rows := sqlmock.
			NewRows([]string{"id", "title", "description", "completed"}).
			AddRow(mockTodo.ID, mockTodo.Title, mockTodo.Description, mockTodo.Completed)

		sqlSelectOne := `SELECT * FROM "todos" 
				WHERE id = $1`

		mock.ExpectQuery(regexp.QuoteMeta(sqlSelectOne)).
			WithArgs(mockTodo.ID).
			WillReturnRows(rows)

		repository := NewRepositoryImpl(db)
		response, err := repository.FindById(mockTodo.ID)
		expect := expectate.Expect(t)
		expect(err).ToEqual(nil)
		expect(response).NotToEqual(nil)
		expect(response.ID).ToEqual(mockTodo.ID)
	})
}
