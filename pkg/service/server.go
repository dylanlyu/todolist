package service

import (
	"todolist/internal/database"
	"todolist/internal/handlers"
	"github.com/gin-gonic/gin"
	"log"
)

// RunGraphServer is gin server
func RunGraphServer() {

	db, err := database.GetDatabase()
	if err != nil {
		log.Println("Unable to connect to database", err)
	}
	defer db.Close()
	database.Migrate(db)

	r := gin.Default()
	r.POST("first", handlers.GraphqlHandler(db))
	r.GET("/", handlers.PlaygroundHandler())
	r.Run(":8080")
}
