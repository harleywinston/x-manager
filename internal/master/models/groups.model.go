package models

import "gorm.io/gorm"

type Credentials struct {
	gorm.Model
	ServerIp string
	Domains  []string
	GroupsID int
}

type Groups struct {
	gorm.Model
	CredentialsID int
	Mode          string
}
