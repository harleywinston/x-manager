package models

import "gorm.io/gorm"

type UsersModel struct {
	gorm.Model
	Email    string `json:"email"`
	Username string `json:"username"`
	Passwd   string `json:"password"`
	Group_ID int    `json:"group_id"`
}
