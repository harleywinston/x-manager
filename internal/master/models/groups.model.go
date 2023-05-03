package models

import "gorm.io/gorm"

type Groups struct {
	gorm.Model
	ResourcesID int    `json:"resources_id"`
	Mode        string `json:"mode"`
}
