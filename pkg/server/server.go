package server

import (
	"github.com/gin-gonic/gin"
	"log"
	"todolist/internal/handlers"
	"todolist/internal/repository"
)

// RunGraphServer is gin server
func RunGraphServer() {

	db, err := repository.GetDatabase()
	if err != nil {
		log.Println("Unable to connect to database", err)
	}
	defer db.Close()
	repository.Migrate(db)

	r := gin.Default()
	r.POST("first", handlers.GraphqlHandler())
	r.GET("/", handlers.PlaygroundHandler())
	r.Run(":8080")
}
