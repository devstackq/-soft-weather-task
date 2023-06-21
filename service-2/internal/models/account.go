package models

type Account struct {
	UserID  string `json:"user_id"`
	Debt    int    `json:"debt"`
	Balance int    `json:"balance"`
}
