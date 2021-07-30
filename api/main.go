package main

import (
	"github.com/gin-gonic/gin"
	"github.com/kutay-celebi/gotodo/todo"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

func main() {
	log.Println("Starting server")

	// initialize routes.
	router := gin.Default()

	// initialize database
	db := initializeDB()

	todoRepository := todo.NewRepositoryImpl(db)
	InitializeRoutes(router, todo.NewController(todoRepository))
	router.Run()
}

func initializeDB() *gorm.DB {

	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  "user=todo password=todo port=5432 sslmode=disable",
		PreferSimpleProtocol: true, // disables implicit prepared statement usage
	}), &gorm.Config{})

	db.AutoMigrate(&todo.Todo{})

	if err != nil {
		log.Fatal(err)
	}
	return db
}
