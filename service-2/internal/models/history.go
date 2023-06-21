package models

type History struct {
	ID     string `json:"id"`
	UserID string `json:"user_id"`
	TaskID string `json:"task_id"`
}
