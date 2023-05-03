package database

import (
	"gorm.io/gorm"

	"github.com/harleywinston/x-manager/internal/master/models"
)

var DB *gorm.DB

func InitModels() error {
	if err := DB.AutoMigrate(&models.Resources{}); err != nil {
		return err
	}
	if err := DB.AutoMigrate(&models.Groups{}); err != nil {
		return err
	}
	if err := DB.AutoMigrate(&models.Users{}); err != nil {
		return err
	}
	return nil
}
