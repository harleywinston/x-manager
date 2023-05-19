package models

import "gorm.io/gorm"

type Users struct {
	gorm.Model
	Email      string `json:"email"       gorm:"unique" uri:"email" binding:"required"`
	Username   string `json:"username"`
	Passwd     string `json:"password"`
	ExpiryTime int64  `json:"expiryTime"`
	GroupID    int    `json:"group_id"`
	FuckedUser bool   `json:"fucked_user"`
}
