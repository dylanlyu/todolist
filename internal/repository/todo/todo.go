package todorepo

import (
	"github.com/jinzhu/gorm"
	"todolist/internal/models"
)

type Repo interface {
	Create(todo *models.Todo) (*models.Todo,error)
	ReadByID(id int) (*models.Todo,error)
	ReadByText(text string) ([]*models.Todo, error)
	ReadDone(done bool) ([]*models.Todo, error)
	ReadByUserID(id int) ([]*models.Todo, error)
	Read()([]*models.Todo,error)
	Update(todo *models.Todo) (*models.Todo,error)
	Delete(todo *models.Todo) (*models.Todo,error)
}

type todoRepo struct {
	db *gorm.DB
}

func NewTodoRepo(db *gorm.DB) Repo {
	return &todoRepo{
		db: db,
	}
}

func (tr *todoRepo) ReadByID(id int) (todo *models.Todo,err error) {
	err = tr.db.First(&todo,id).Error
	if err != nil {
		return nil, err
	}
	return todo,nil
}

func (tr todoRepo) ReadByText(text string) (todos []*models.Todo, err error) {
	err = tr.db.Find(&todos,"text LINK ?","%"+text+"%").Error
	if err != nil {
		return nil, err
	}
	return todos,nil
}

func (tr todoRepo) ReadDone(done bool) (todos []*models.Todo, err error) {
	err = tr.db.Find(&todos, models.Todo{Done: done}).Error
	if err != nil {
		return nil, err
	}
	return todos,nil
}

func (tr todoRepo) ReadByUserID(uid int) (todos []*models.Todo,err error) {
	err = tr.db.Find(&todos,models.Todo{UserID: uid}).Error
	if err != nil {
		return nil, err
	}
	return todos,nil
}

func (tr todoRepo) Read() (todos []*models.Todo, err error) {
	err = tr.db.Find(&todos).Error
	if err != nil {
		return nil, err
	}
	return todos,nil
}

func (tr todoRepo) Create(todo *models.Todo) (*models.Todo,error) {
	err := tr.db.Create(&todo).Error
	if err != nil {
		return nil, err
	}
	return todo,nil
}

func (tr todoRepo) Update(todo *models.Todo) (*models.Todo,error) {
	err := tr.db.Save(&todo).Error
	if err != nil {
		return nil, err
	}
	return todo,nil
}

func (tr todoRepo) Delete(todo *models.Todo) (*models.Todo,error) {
	err:= tr.db.Delete(&todo).Error
	if err != nil {
		return nil, err
	}
	return todo,nil
}