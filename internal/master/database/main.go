package database

import (
	"gorm.io/gorm"

	"github.com/harleywinston/x-manager/internal/master/consts"
	"github.com/harleywinston/x-manager/internal/master/models"
)

var DB *gorm.DB

func InitModels() error {
	if err := DB.AutoMigrate(&models.Resources{}); err != nil {
		return &consts.CustomError{
			Message: consts.AUTO_MIGRATE_DB_ERROR.Message,
			Code:    consts.AUTO_MIGRATE_DB_ERROR.Code,
			Detail:  err.Error(),
		}
	}
	if err := DB.AutoMigrate(&models.Groups{}); err != nil {
		return &consts.CustomError{
			Message: consts.AUTO_MIGRATE_DB_ERROR.Message,
			Code:    consts.AUTO_MIGRATE_DB_ERROR.Code,
			Detail:  err.Error(),
		}
	}
	if err := DB.AutoMigrate(&models.Users{}); err != nil {
		return &consts.CustomError{
			Message: consts.AUTO_MIGRATE_DB_ERROR.Message,
			Code:    consts.AUTO_MIGRATE_DB_ERROR.Code,
			Detail:  err.Error(),
		}
	}
	return nil
}
