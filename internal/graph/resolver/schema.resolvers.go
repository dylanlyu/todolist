package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"encoding/json"
	"errors"
	"todolist/internal/graph/generated"
	"todolist/internal/middlewares"
	"todolist/internal/models"
)

func (r *mutationResolver) CreateTodo(ctx context.Context, input models.NewTodo) (*models.Todo, error) {
	data, _ := json.Marshal(&input)
	gc, err := middlewares.GinContextFromContext(ctx)
	gc = middlewares.SetGinRequestData(gc, data)
	err = gc.ShouldBind(&input)
	if err != nil {
		return nil, err
	}
	user, _ := r.UserService.GetUser(input.UserID)
	if user.Name == "" {
		return nil, errors.New("user not found")
	}
	return r.TodoService.CreateTodo(&input)
}

func (r *mutationResolver) CreateUser(ctx context.Context, input models.NewUser) (*models.User, error) {
	data, _ := json.Marshal(&input)
	gc, err := middlewares.GinContextFromContext(ctx)
	gc = middlewares.SetGinRequestData(gc, data)
	err = gc.ShouldBind(&input)
	if err != nil {
		return nil, err
	}
	return r.UserService.CreateUser(&input)
}

func (r *mutationResolver) DeleteTodo(ctx context.Context, id int) (*models.Todo, error) {
	return r.TodoService.DeleteTodo(id)
}

func (r *mutationResolver) DeleteUser(ctx context.Context, id int) (*models.User, error) {
	return r.UserService.DeleteUser(id)
}

func (r *queryResolver) SearchTodos(ctx context.Context, input models.SearchTodo) ([]*models.Todo, error) {
	data, _ := json.Marshal(&input)
	gc, err := middlewares.GinContextFromContext(ctx)
	gc = middlewares.SetGinRequestData(gc, data)
	err = gc.ShouldBind(&input)
	if err != nil {
		return nil, err
	}
	return r.TodoService.SearchTodos(&input)
}

func (r *queryResolver) SearchUsers(ctx context.Context, name string) ([]*models.User, error) {
	return r.UserService.SearchUser(name)
}

func (r *queryResolver) GetUser(ctx context.Context, id int) (*models.User, error) {
	return r.UserService.GetUser(id)
}

func (r *queryResolver) GetTodo(ctx context.Context, id int) (*models.Todo, error) {
	return r.TodoService.GetTodo(id)
}

func (r *todoResolver) User(ctx context.Context, obj *models.Todo) (*models.User, error) {
	return r.UserService.GetUser(obj.UserID)
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

// Todo returns generated.TodoResolver implementation.
func (r *Resolver) Todo() generated.TodoResolver { return &todoResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
type todoResolver struct{ *Resolver }
