package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/kutay-celebi/gotodo/todo"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
)

func main() {
	log.Println("Starting server")

	// initialize routes.
	router := gin.Default()

	// initialize database
	db := initializeDB()

	todoRepository := todo.NewRepositoryImpl(db)

	InitializeRoutes(router, todo.NewController(todoRepository))
	config := cors.DefaultConfig()
	config.AddAllowHeaders("*")
	config.AllowAllOrigins = true
	router.Use(cors.New(config))
	router.Run()
}

func initializeDB() *gorm.DB {

	dbHost := os.Getenv("DB_HOST")
	if dbHost == "" {
		log.Fatalf("Please set the DB_HOST variable.")
	}

	dbUser := os.Getenv("DB_USER")
	if dbUser == "" {
		log.Fatalf("Please set the DB_USER variable.")
	}

	dbPass := os.Getenv("DB_PASS")
	if dbPass == "" {
		log.Fatalf("Please set the DB_PASS variable.")
	}

	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  "host=" + dbHost + " user=" + dbUser + " password=" + dbPass + " port=5432 dbname=tododb sslmode=disable",
		PreferSimpleProtocol: true, // disables implicit prepared statement usage
	}), &gorm.Config{})

	// create uuid-ossp extension if not exists.
	// https://stackoverflow.com/questions/22446478/extension-exists-but-uuid-generate-v4-fails
	db.Exec("create extension IF NOT EXISTS \"uuid-ossp\" schema pg_catalog version \"1.1\"")
	// auto migration on DB
	db.AutoMigrate(&todo.Todo{})

	if err != nil {
		log.Fatal(err)
	}
	return db
}
