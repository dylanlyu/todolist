package todoservice

import (
	"errors"
	"todolist/internal/models"
	todoRepo "todolist/internal/repository/todo"
)

type TodoService interface {
	CreateTodo(input *models.NewTodo) (*models.Todo, error)
	GetTodo(id int) (*models.Todo, error)
	SearchTodos(input *models.SearchTodo) ([]*models.Todo, error)
}

type todoService struct {
	Repo todoRepo.Repo
}

func NewTodoService(
	repo todoRepo.Repo) TodoService {
	return &todoService{
		Repo: repo,
	}
}

func (ts todoService) CreateTodo(input *models.NewTodo) (*models.Todo, error) {
	todo := &models.Todo{
		Text: input.Text,
		UserID: input.UserID,
		Done: false,
	}
	return ts.Repo.Create(todo)
}

func (ts todoService) GetTodo(id int) (*models.Todo, error) {
	return ts.Repo.ReadByID(id)
}

func (ts todoService) SearchTodos(input *models.SearchTodo) ([]*models.Todo, error) {
	if input.Text != nil {
		return ts.Repo.ReadByTextAndUserID(input.Text,input.UserID)
	}
	if input.Done != nil {
		return ts.Repo.ReadByDoneAndUserID(input.Done,input.UserID)
	}
	return nil, errors.New("pls key in effective value")
}