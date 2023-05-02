package database

import (
	"github.com/harleywinston/x-manager/internal/master/models"
)

type UsersDB struct{}

func (db *UsersDB) AddUserToDB(user models.UsersModel) error {
	err := DB.AutoMigrate(&models.UsersModel{})
	if err != nil {
		return err
	}

	return DB.Create(&user).Error
}

func (db *UsersDB) GetUserFromDB(user models.UsersModel) (models.UsersModel, error) {
	var res models.UsersModel
	err := DB.First(&res, user).Error
	return res, err
}

func (db *UsersDB) GetAllUsersFromDB(user models.UsersModel) ([]models.UsersModel, error) {
	return []models.UsersModel{}, nil
}

func (db *UsersDB) DeleteUserFromDB(user models.UsersModel) error {
	return DB.Delete(&user, user).Error
}
