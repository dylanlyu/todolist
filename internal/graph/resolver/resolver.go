package resolver

import (
	todoservice "todolist/internal/service/todo"
	userservice "todolist/internal/service/user"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	TodoService todoservice.TodoService
	UserService userservice.UserService
}
