package database

import "github.com/harleywinston/x-manager/internal/master/models"

type UsersDB struct{}

func (db *UsersDB) AddUserToDB(user models.UsersModel) error {
	return nil
}

func (db *UsersDB) GetUserFromDB(user models.UsersModel) (models.UsersModel, error) {
	return models.UsersModel{}, nil
}

func (db *UsersDB) GetAllUsersFromDB(user models.UsersModel) ([]models.UsersModel, error) {
	return []models.UsersModel{}, nil
}

func (db *UsersDB) DeleteUserFromDB(user models.UsersModel) error {
	return nil
}
