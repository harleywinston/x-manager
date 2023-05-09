package models

import "gorm.io/gorm"

type Resources struct {
	gorm.Model
	ServerIp          string `json:"server_ip"`
	Domains           string `json:"domains"`
	BrdigeDomain      string `json:"bridge_domain"`
	CloudflareDomains string `json:"cloudflare_domains"`
}
