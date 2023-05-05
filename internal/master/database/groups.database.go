package database

import (
	"fmt"

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

func (db *GroupsDB) CheckResourceID(id int) error {
	var resources []models.Resources
	DB.Model(&models.Resources{}).Select("id").Scan(&resources)

	for _, r := range resources {
		if int(r.ID) == id {
			return nil
		}
	}
	return fmt.Errorf(`Resource id %v is not valid!`, id)
}
