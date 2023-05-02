package database

import "github.com/harleywinston/x-manager/internal/master/models"

type ResourceDB struct{}

func (db *ResourceDB) AddResourceToDB(resource models.Resources) error {
	return nil
}

func (db *ResourceDB) GetResourceFromDB(resource models.Resources) (models.Resources, error) {
	return models.Resources{}, nil
}

func (db *ResourceDB) DeleteResourceFromDB(resource models.Resources) error {
	return nil
}
