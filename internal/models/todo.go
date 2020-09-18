package models

type Todo struct {
	ID     int    `json:"id" gorm:AUTO_INCREMENT;PRIMARY_KEY"`
	Text   string `json:"text"`
	Done   bool   `json:"done"`
	UserID int `json:"user"`
}

type NewTodo struct {
	Text   string `json:"text" binding:"required"`
	UserID int    `json:"userId" binding:"required"`
}

type SearchTodo struct {
	Text   *string `json:"text"`
	Done   *bool   `json:"done"`
	UserID int     `json:"userId" binding:"required"`
}
