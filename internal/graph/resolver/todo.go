package resolver

import (
	"todolist/internal/graph/models"
)

func CreateUpdateTodo(r *mutationResolver,input models.NewTodo, update bool, ids ...string) (*models.Todo, error) {
	todo := models.Todo{
		Text: input.Text,
	}

	err := r.DB.Create(&todo).Error
	if err != nil {
		return nil, err
	}

	return &todo, nil
}

func SelectTodos(r *queryResolver) ([]*models.Todo, error) {
	var todos []*models.Todo
	err := r.DB.Find(&todos).Error

	if err != nil {
		return nil, err
	}

	return todos, nil
}
