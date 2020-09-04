package todorepo

import (
	"github.com/jinzhu/gorm"
	"todolist/internal/models"
)

type Repo interface {
	GetByID(id int) (*models.Todo,error)
}

type todoRepo struct {
	db *gorm.DB
}

func NewTodoRepo(db *gorm.DB) Repo {
	return &todoRepo{
		db: db,
	}
}

func (tr *todoRepo) GetByID(id int) (todo *models.Todo,err error) {
	if err = tr.db.Find(&todo,id).Error;err != nil {
		return nil, err
	}
	return todo,nil
}
