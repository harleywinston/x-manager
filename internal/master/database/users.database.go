package database

import (
	"gorm.io/gorm"

	"github.com/harleywinston/x-manager/internal/master/models"
)

type UsersDB struct {
	DB *gorm.DB
}

func (db *UsersDB) AddUserToDB(user models.UsersModel) error {
	err := db.DB.AutoMigrate(&models.UsersModel{})
	if err != nil {
		return err
	}

	db.DB.Create(&user)
	return nil
}

func (db *UsersDB) GetUserFromDB(user models.UsersModel) (models.UsersModel, error) {
	var res models.UsersModel
	db.DB.First(&res, user)
	return res, nil
}

func (db *UsersDB) GetAllUsersFromDB(user models.UsersModel) ([]models.UsersModel, error) {
	return []models.UsersModel{}, nil
}

func (db *UsersDB) DeleteUserFromDB(user models.UsersModel) error {
	db.DB.Delete(&user, user)
	return nil
}
