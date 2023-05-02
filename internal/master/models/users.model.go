package models

import "gorm.io/gorm"

type Users struct {
	gorm.Model
	Email    string `json:"email"    gorm:"unique"`
	Username string `json:"username"`
	Passwd   string `json:"password"`
	GroupsID int    `json:"group_id"`
}
