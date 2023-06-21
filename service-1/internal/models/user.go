package models

type User struct {
	ID       string `json:"id"`
	FullName string `json:"name"`
	GroupNum int    `json:"group_num"`
	Login    string `json:"login"`
}
