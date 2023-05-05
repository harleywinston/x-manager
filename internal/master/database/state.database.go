package database

import (
	"github.com/harleywinston/x-manager/internal/master/consts"
	"github.com/harleywinston/x-manager/internal/master/models"
)

type StateDatabase struct{}

func (db *StateDatabase) GetUsers(groupID int) ([]models.Users, error) {
	var users []models.Users
	err := DB.Model(&models.Users{}).Where("groups_id = ?", groupID).Find(&users).Error
	if err != nil {
		return []models.Users{}, &consts.CustomError{
			Message: consts.INVALID_GROUP_ID_ERROR.Message,
			Code:    consts.INVALID_GROUP_ID_ERROR.Code,
			Detail:  err.Error(),
		}
	}

	return users, nil
}

func (db *StateDatabase) GetResourceID(serverIP string) (int, error) {
	var resource models.Resources
	err := DB.Model(&models.Resources{}).Where("server_ip = ?", serverIP).First(&resource).Error
	if err != nil {
		return 0, &consts.CustomError{
			Message: consts.INVALID_RESOURCE_IP_ERROR.Message,
			Code:    consts.INVALID_RESOURCE_IP_ERROR.Code,
			Detail:  err.Error(),
		}
	}

	return int(resource.ID), nil
}

func (db *StateDatabase) GetGroupID(resourceID int) (int, error) {
	var group models.Groups
	err := DB.Model(&models.Groups{}).Where("resources_id = ?", resourceID).First(&group).Error
	if err != nil {
		return 0, &consts.CustomError{
			Message: consts.RECOURSE_ID_NOT_VALID_ERROR.Message,
			Code:    consts.RECOURSE_ID_NOT_VALID_ERROR.Code,
			Detail:  err.Error(),
		}
	}

	return int(group.ID), nil
}
