package todo

import (
	"github.com/kutay-celebi/gotodo/api/util"
	"gorm.io/gorm"
)

type Repository interface {
	FindAll() (*[]Todo, error)
	Save(todo *Todo) (*Todo, error)
	Update(todo *Todo) (*Todo, error)
	FindById(param string) (*Todo, error)
}

func NewRepositoryImpl(db *gorm.DB) Repository {
	return &RepositoryImpl{db}
}

type RepositoryImpl struct {
	db *gorm.DB
}

func (r *RepositoryImpl) Update(todo *Todo) (*Todo, error) {
	db := r.db.Save(&todo)

	if db.Error != nil {
		return nil, util.ErrInternal
	}

	return todo, nil
}

func (r *RepositoryImpl) FindById(param string) (*Todo, error) {
	var todo Todo
	find := r.db.Where("id = ?", param).Find(&todo)

	if find.Error != nil {
		return nil, find.Error
	}

	if find.RowsAffected > 0 {
		return &todo, nil
	}

	return nil, util.ErrRecordNotFound
}

func (r *RepositoryImpl) FindAll() (*[]Todo, error) {
	var todos []Todo
	find := r.db.Find(&todos)
	return &todos, find.Error
}

func (r *RepositoryImpl) Save(todo *Todo) (*Todo, error) {
	r.db.Create(&todo)
	return todo, nil
}
