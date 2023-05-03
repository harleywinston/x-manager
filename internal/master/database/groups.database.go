package database

import (
	"github.com/harleywinston/x-manager/internal/master/models"
)

type GroupsDB struct{}

func (db *GroupsDB) AddGroupToDB(group models.Groups) error {
	return DB.Create(&group).Error
}

func (db *GroupsDB) GetGroupFromDB(group models.Groups) (models.Groups, error) {
	var res models.Groups
	err := DB.First(&res, group).Error
	return res, err
}

func (db *GroupsDB) DeleteGroupFromDB(group models.Groups) error {
	return DB.Delete(&group, group).Error
}
