package models

import "gorm.io/gorm"

type Users struct {
	gorm.Model
	Email      string `json:"email"       gorm:"unique"`
	Username   string `json:"username"`
	Passwd     string `json:"password"`
	ExpiryTime int64  `json:"expiryTime"`
	GroupsID   int    `json:"group_id"`
	FuckedUser bool   `json:"fucked_user"`
}
