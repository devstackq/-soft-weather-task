package models

type Task struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Price       int    `json:"price"`
}
