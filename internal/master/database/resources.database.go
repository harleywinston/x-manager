package database

import "github.com/harleywinston/x-manager/internal/master/models"

type ResourceDB struct{}

func (db *ResourceDB) AddResourceToDB(resource models.Resources) error {
	return DB.Create(&resource).Error
}

func (db *ResourceDB) GetResourceFromDB(resource models.Resources) (models.Resources, error) {
	var res models.Resources
	err := DB.First(&res, resource).Error
	return res, err
}

func (db *ResourceDB) DeleteResourceFromDB(resource models.Resources) error {
	return DB.Delete(&resource, resource).Error
}
