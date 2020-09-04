package service

import (
	"todolist/internal/models"
)

type TodoControl struct {
	*dataBase
}

func (tc *repository.dataBase) CreateUpdateTodo(input models.NewTodo, update bool, ids ...string) (*models.Todo, error) {
	todo := &models.Todo{
		Text: input.Text,
		UserID: input.UserID,
	}
	err := tc.db.Create(&todo).Error
	if err != nil {
		return nil, err
	}

	return todo, nil
}

func (tc *TodoControl) SelectTodos() ([]*models.Todo, error) {
	var todos []*models.Todo
	err := tc.db.Find(&todos).Error
	if err != nil {
		return nil, err
	}

	return todos, nil
}
