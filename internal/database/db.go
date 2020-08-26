package database

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"todolist/configs"
	"todolist/internal/graph/models"
)

var config *configs.Configs

func init() {
	config = configs.GetConfig()
}

func getDatabaseUrl() string {
	return fmt.Sprintf(
		"host=%s port=%d user=%s dbname=%s password=%s sslmode=%s",
		config.PostgreSQL.Host, config.PostgreSQL.Port, config.PostgreSQL.User, config.PostgreSQL.DBName,
		config.PostgreSQL.Password, config.PostgreSQL.SSLMode)
}

func GetDatabase() (*gorm.DB, error) {
	db, err := gorm.Open("postgres", getDatabaseUrl())
	return db, err
}

func Migrate(db *gorm.DB) {
	db.AutoMigrate(&models.Todo{})
}
