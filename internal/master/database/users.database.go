package database

import (
	"github.com/harleywinston/x-manager/internal/master/models"
)

type UsersDB struct{}

func (db *UsersDB) AddUserToDB(user models.Users) error {
	return DB.Create(&user).Error
}

func (db *UsersDB) GetUserFromDB(user models.Users) (models.Users, error) {
	var res models.Users
	err := DB.First(&res, user).Error
	return res, err
}

func (db *UsersDB) GetAllUsersFromDB(user models.Users) ([]models.Users, error) {
	return []models.Users{}, nil
}

func (db *UsersDB) DeleteUserFromDB(user models.Users) error {
	return DB.Delete(&user, user).Error
}
