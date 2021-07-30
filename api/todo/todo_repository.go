package todo

import (
	"gorm.io/gorm"
)

type Repository interface {
	FindAll() (*[]Todo, error)
}

func NewRepositoryImpl(db *gorm.DB) Repository {
	return &RepositoryImpl{db}
}

type RepositoryImpl struct {
	db *gorm.DB
}

func (r *RepositoryImpl) FindAll() (*[]Todo, error) {
	var todos []Todo

	r.db.Find(&todos)
	return &todos, nil
}
