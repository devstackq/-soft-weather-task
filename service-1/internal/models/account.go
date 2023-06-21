package models

type Account struct {
	ID      string `json:"id"`
	UserID  string `json:"user_id"`
	Debt    int    `json:"debt"`
	Balance int    `json:"balance"`
}
