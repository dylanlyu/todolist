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

func (ur userRepo) Read() ([]*models.User,error) {
	var users []*models.User
	err := ur.db.Find(&users).Error
	if err != nil {
		return nil, err
	}
	return users,nil
}

func (ur userRepo) ReadByID(id int) (*models.User,error) {
	var user *models.User
	err := ur.db.First(&user,id).Error
	if err != nil {
		return nil, err
	}
	return user,nil
}

func (ur userRepo) ReadByName(name string) ([]*models.User,error) {
	var users []*models.User
	err := ur.db.Find(&users,"name LINK ?","%"+name+"%").Error
	if err != nil {
		return nil, err
	}
	return users,nil
}

func (ur userRepo) Update(user *models.User) (*models.User,error) {
	err := ur.db.Save(&user).Error
	if err != nil {
		return nil, err
	}
	return user,err
}

func (ur userRepo) Delete(user *models.User) (*models.User,error) {
	er := ur.db.Delete(&user).Error
	if er != nil {
		return nil, er
	}
	return user,nil
}