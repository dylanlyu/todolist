package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	generated1 "todolist/internal/graph/generated"
	models1 "todolist/internal/graph/models"
)

func (r *mutationResolver) CreateTodo(ctx context.Context, input models1.NewTodo) (*models1.Todo, error) {
	return CreateUpdateTodo(r, input, false)
}

func (r *queryResolver) Todos(ctx context.Context) ([]*models1.Todo, error) {
	return SelectTodos(r)
}

// Mutation returns generated1.MutationResolver implementation.
func (r *Resolver) Mutation() generated1.MutationResolver { return &mutationResolver{r} }

// Query returns generated1.QueryResolver implementation.
func (r *Resolver) Query() generated1.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
