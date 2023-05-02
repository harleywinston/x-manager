package models

type UsersModel struct {
	ID       int    `json:"id"`
	Email    string `json:"email"`
	Username string `json:"username"`
	Passwd   string `json:"password"`
	Group_ID int    `json:"group_id"`
}
