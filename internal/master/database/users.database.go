package database

import (
	"os"
	"strconv"

	"github.com/harleywinston/x-manager/internal/master/consts"
	"github.com/harleywinston/x-manager/internal/master/models"
)

type UsersDB struct{}

func (db *UsersDB) AddUserToDB(user models.Users) error {
	err := DB.Create(&user).Error
	if err != nil {
		return &consts.CustomError{
			Message: consts.ADD_DB_ERROR.Message,
			Code:    consts.ADD_DB_ERROR.Code,
			Detail:  err.Error(),
		}
	}
	return nil
}

func (db *UsersDB) GetUserFromDB(user models.Users) (models.Users, error) {
	var res models.Users
	err := DB.First(&res, user).Error
	if err != nil {
		return models.Users{}, &consts.CustomError{
			Message: consts.GET_DB_ERROR.Message,
			Code:    consts.GET_DB_ERROR.Code,
			Detail:  err.Error(),
		}
	}
	return res, nil
}

func (db *UsersDB) GetAllUsersFromDB(user models.Users) ([]models.Users, error) {
	return []models.Users{}, nil
}

func (db *UsersDB) DeleteUserFromDB(user models.Users) error {
	err := DB.Delete(&user, user).Error
	if err != nil {
		return &consts.CustomError{
			Message: consts.DELETE_DB_ERROR.Message,
			Code:    consts.DELETE_DB_ERROR.Code,
			Detail:  err.Error(),
		}
	}
	return nil
}

func (db *UsersDB) GetFreeGroupIDFromDB() (int, error) {
	type groupWithCount struct {
		GroupID int
		Count   int
	}
	var groupsWithCount []groupWithCount

	err := DB.Model(&models.Users{}).
		Select("group_id, count(*) as count").
		Group("group_id").
		Scan(&groupsWithCount).Error
	if err != nil {
		return 0, &consts.CustomError{
			Message: consts.GET_DB_ERROR.Message,
			Code:    consts.GET_DB_ERROR.Code,
			Detail:  err.Error(),
		}
	}

	var resGroups []int
	for _, g := range groupsWithCount {
		group_limit, err := strconv.ParseInt(os.Getenv("GROUPS_LIMIT"), 10, 64)
		if err != nil {
			return 0, &consts.CustomError{
				Message: consts.PARSE_INT_ERROR.Message,
				Code:    consts.PARSE_INT_ERROR.Code,
				Detail:  err.Error(),
			}
		}
		if g.Count < int(group_limit) {
			resGroups = append(resGroups, g.GroupID)
		}
	}

	if len(resGroups) < 1 {
		return 0, consts.GROUP_LIMIT_ERROR
	}

	return resGroups[0], nil
}

func (db *UsersDB) GetUsersRecourse(user models.Users) (models.Resources, error) {
	err := DB.First(&user, user).Error
	if err != nil {
		return models.Resources{}, &consts.CustomError{
			Message: consts.GET_DB_ERROR.Message,
			Code:    consts.GET_DB_ERROR.Code,
			Detail:  err.Error(),
		}
	}

	var group models.Groups
	err = DB.Model(&models.Groups{}).Where("ID = ?", user.GroupID).Scan(&group).Error
	if err != nil {
		return models.Resources{}, &consts.CustomError{
			Message: consts.GET_DB_ERROR.Message,
			Code:    consts.GET_DB_ERROR.Code,
			Detail:  err.Error(),
		}
	}

	var res models.Resources
	err = DB.Model(&models.Resources{}).Where("ID = ?", group.ResourcesID).Scan(&res).Error
	if err != nil {
		return models.Resources{}, &consts.CustomError{
			Message: consts.GET_DB_ERROR.Message,
			Code:    consts.GET_DB_ERROR.Code,
			Detail:  err.Error(),
		}
	}

	return res, nil
}
