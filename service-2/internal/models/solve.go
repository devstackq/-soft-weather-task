package models

type Solve struct {
	Input  []int  `json:"input"`
	UserID string `json:"user_id"`
	TaskID string `json:"task_id"`
}
