package models

type Todo struct {
	ID     int    `json:"id" gorm:AUTO_INCREMENT;PRIMARY_KEY"`
	Text   string `json:"text"`
	Done   bool   `json:"done"`
	UserID int `json:"user"`
}
