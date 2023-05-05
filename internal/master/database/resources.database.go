package database

import (
	"github.com/harleywinston/x-manager/internal/master/consts"
	"github.com/harleywinston/x-manager/internal/master/models"
)

type ResourceDB struct{}

func (db *ResourceDB) AddResourceToDB(resource models.Resources) error {
	err := DB.Create(&resource).Error
	if err != nil {
		return &consts.CustomError{
			Message: consts.ADD_DB_ERROR.Message,
			Code:    consts.ADD_DB_ERROR.Code,
			Detail:  err.Error(),
		}
	}
	return nil
}

func (db *ResourceDB) GetResourceFromDB(resource models.Resources) (models.Resources, error) {
	var res models.Resources
	err := DB.First(&res, resource).Error
	if err != nil {
		return models.Resources{}, &consts.CustomError{
			Message: consts.GET_DB_ERROR.Message,
			Code:    consts.GET_DB_ERROR.Code,
			Detail:  err.Error(),
		}
	}
	return res, err
}

func (db *ResourceDB) DeleteResourceFromDB(resource models.Resources) error {
	err := DB.Delete(&resource, resource).Error
	if err != nil {
		return &consts.CustomError{
			Message: consts.DELETE_DB_ERROR.Message,
			Code:    consts.DELETE_DB_ERROR.Code,
			Detail:  err.Error(),
		}
	}
	return nil
}
