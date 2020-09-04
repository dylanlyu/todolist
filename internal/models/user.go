package models

type User struct {
	ID   int    `json:"id" gorm:AUTO_INCREMENT;PRIMARY_KEY"`
	Name string `json:"name"`
}
