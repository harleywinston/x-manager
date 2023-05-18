package models

import "gorm.io/gorm"

type Resources struct {
	gorm.Model
	ServerIp       string `json:"server_ip"`
	Domains        string `json:"domains"`
	Bridges        string `json:"bridges"`
	ForeignBridges string `json:"foreign_bridges"`
}
