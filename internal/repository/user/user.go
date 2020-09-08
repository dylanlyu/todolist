package userrepo

import (
	"github.com/jinzhu/gorm"
	"todolist/internal/models"
)

type Repo interface {
	Create(user *models.User) (*models.User,error)
	Read() ([]*models.User,error)
	ReadByID(id int) (*models.User,error)
	ReadByName(name string) ([]*models.User,error)
	Update(user *models.User) (*models.User,error)
	Delete(user *models.User) (*models.User,error)
}

type userRepo struct {
	db *gorm.DB
}

func NewUserRepo(db *gorm.DB) Repo {
	return &userRepo{
		db: db,
	}
}

func (ur userRepo) Create(user *models.User) (*models.User, error) {
	err := ur.db.Create(&user).Error
	if err != nil {
		return nil, err
	}
	return user,nil
}

func (ur userRepo) Read() (users []*models.User, err error) {

}