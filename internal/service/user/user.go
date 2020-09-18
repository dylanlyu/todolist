package userservice

import (
	"todolist/internal/models"
	userRepo "todolist/internal/repository/user"
)

type UserService interface {
	CreateUser(newUser *models.NewUser) (*models.User, error)
	GetUser(id int) (*models.User, error)
	SearchUser(name string) ([]*models.User, error)
}

type userService struct {
	Repo userRepo.Repo
}

func NewTodoService(
	repo userRepo.Repo) UserService {
	return &userService{
		Repo: repo,
	}
}

func (us userService) CreateUser(newUser *models.NewUser) (*models.User, error) {
	user := &models.User{
		Name: newUser.Name,
	}
	return us.Repo.Create(user)
}

func (us userService) GetUser(id int) (*models.User, error) {
	return us.Repo.ReadByID(id)
}

func (us userService) SearchUser(name string) ([]*models.User, error) {
	return us.Repo.ReadByName(name)
}