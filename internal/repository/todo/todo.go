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
	ReadByTextAndUserID(text *string,uid int) ([]*models.Todo, error)
	ReadByDoneAndUserID(done *bool,uid int) ([]*models.Todo, error)
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

func (tr *todoRepo) ReadByID(id int) (*models.Todo, error) {
	var todo *models.Todo
	err := tr.db.First(&todo,id).Error
	if err != nil {
		return nil, err
	}
	return todo,nil
}

func (tr todoRepo) ReadByText(text string) ([]*models.Todo, error) {
	var todos []*models.Todo
	err := tr.db.Find(&todos,"text LINK ?","%"+text+"%").Error
	if err != nil {
		return nil, err
	}
	return todos,nil
}

func (tr todoRepo) ReadDone(done bool) ([]*models.Todo, error) {
	var todos []*models.Todo
	err := tr.db.Find(&todos, models.Todo{Done: done}).Error
	if err != nil {
		return nil, err
	}
	return todos,nil
}

func (tr todoRepo) ReadByTextAndUserID(text *string,uid int) ([]*models.Todo, error) {
	var todos []*models.Todo
	err := tr.db.Where("text LINK ? and user_id = ?","%"+*text+"%",uid).Find(&todos).Error
	if err != nil {
		return nil, err
	}
	return todos,nil
}

func (tr todoRepo) ReadByDoneAndUserID(done *bool,uid int) ([]*models.Todo, error) {
	var todos []*models.Todo
	err := tr.db.Find(&todos, models.Todo{Done: *done,UserID: uid}).Error
	if err != nil {
		return nil, err
	}
	return todos,nil
}

func (tr todoRepo) ReadByUserID(uid int) ([]*models.Todo, error) {
	var todos []*models.Todo
	err := tr.db.Find(&todos,models.Todo{UserID: uid}).Error
	if err != nil {
		return nil, err
	}
	return todos,nil
}

func (tr todoRepo) Read() ([]*models.Todo, error) {
	var todos []*models.Todo
	err := tr.db.Find(&todos).Error
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