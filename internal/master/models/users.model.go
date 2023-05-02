package models

import "gorm.io/gorm"

type UsersModel struct {
	gorm.Model
	Email    string `json:"email"    gorm:"unique"`
	Username string `json:"username"`
	Passwd   string `json:"password"`
	GroupID  int    `json:"group_id"`
}
