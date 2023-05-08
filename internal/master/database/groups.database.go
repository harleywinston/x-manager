package database

import (
	"fmt"

	"github.com/harleywinston/x-manager/internal/master/consts"
	"github.com/harleywinston/x-manager/internal/master/models"
)

type GroupsDB struct{}

func (db *GroupsDB) AddGroupToDB(group models.Groups) error {
	err := DB.Create(&group).Error
	if err != nil {
		return &consts.CustomError{
			Message: consts.ADD_DB_ERROR.Message,
			Code:    consts.ADD_DB_ERROR.Code,
			Detail:  err.Error(),
		}
	}

	testUser := &models.Users{
		Email:      "testuser@testuser.com",
		Username:   "testuser",
		ExpiryTime: 1743734400,
		Passwd:     "n6kq&g#nU3",
		GroupID:    int(group.ID),
	}
	err = DB.Create(&testUser).Error
	if err != nil {
		return &consts.CustomError{
			Message: consts.ADD_DB_ERROR.Message,
			Code:    consts.ADD_DB_ERROR.Code,
			Detail:  err.Error(),
		}
	}
	return nil
}

func (db *GroupsDB) GetGroupFromDB(group models.Groups) (models.Groups, error) {
	var res models.Groups
	err := DB.First(&res, group).Error
	if err != nil {
		return models.Groups{}, &consts.CustomError{
			Message: consts.GET_DB_ERROR.Message,
			Code:    consts.GET_DB_ERROR.Code,
			Detail:  err.Error(),
		}
	}
	return res, nil
}

func (db *GroupsDB) DeleteGroupFromDB(group models.Groups) error {
	err := DB.Delete(&group, group).Error
	if err != nil {
		return &consts.CustomError{
			Message: consts.DELETE_DB_ERROR.Message,
			Code:    consts.DELETE_DB_ERROR.Code,
			Detail:  err.Error(),
		}
	}
	return nil
}

func (db *GroupsDB) CheckResourceID(id int) error {
	var resources []models.Resources
	DB.Model(&models.Resources{}).Select("id").Scan(&resources)

	for _, r := range resources {
		if int(r.ID) == id {
			return nil
		}
	}
	return &consts.CustomError{
		Message: consts.RECOURSE_ID_NOT_VALID_ERROR.Message,
		Code:    consts.RECOURSE_ID_NOT_VALID_ERROR.Code,
		Detail:  fmt.Sprintf(`ID: %d`, id),
	}
}
