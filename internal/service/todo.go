package service

import (
	todoRepo "todolist/internal/repository/todo"
)

type TodoService interface {

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