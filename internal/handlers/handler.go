package handlers

import (
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"todolist/internal/graph/generated"
	"todolist/internal/graph/resolver"
	"todolist/internal/repository/todo"
	userrepo "todolist/internal/repository/user"
	todoservice "todolist/internal/service/todo"
	userservice "todolist/internal/service/user"
)

// GraphqlHandler is gin handler to graph
func GraphqlHandler(config generated.Config) gin.HandlerFunc {
	// NewExecutableSchema and Config are in the generated.go file
	// Resolver is in the resolver.go file
	h := handler.NewDefaultServer(generated.NewExecutableSchema(config))

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

// PlaygroundHandler is gin handler to graph
func PlaygroundHandler() gin.HandlerFunc {
	h := playground.Handler("GraphQL", "/query")

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

func ResolverConfig(db *gorm.DB) generated.Config {
	todoRepo := todorepo.NewTodoRepo(db)
	todoService := todoservice.NewTodoService(todoRepo)

	userRepo := userrepo.NewUserRepo(db)
	userService := userservice.NewTodoService(userRepo)

	return generated.Config {
		Resolvers: &resolver.Resolver{
			TodoService: todoService,
			UserService: userService,
		},
	}
}