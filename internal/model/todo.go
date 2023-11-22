package model

type Todo struct {
	ID     string `json:"id" db:"id"`
	Text   string `json:"text" db:"text"`
	Done   bool   `json:"done" db:"done"`
	UserID string `json:"userId" db:"user_id"`
}
