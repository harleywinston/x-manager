package models

import "gorm.io/gorm"

type Groups struct {
	gorm.Model
	ResourcesID int
	Mode        string
}
