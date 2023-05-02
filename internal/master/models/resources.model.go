package models

import "gorm.io/gorm"

type Resources struct {
	gorm.Model
	ServerIp string
	Domains  []string
	GroupsID int
}
