package models

import "gorm.io/gorm"

type Resources struct {
	gorm.Model
	ServerIp string `json:"server_ip"`
	Domains  string `json:"domains"`
}
