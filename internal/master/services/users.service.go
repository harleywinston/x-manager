package services

import (
	"github.com/harleywinston/x-manager/internal/master/database"
	"github.com/harleywinston/x-manager/internal/master/models"
)

type UsersService struct {
	usersDB database.UsersDB
}

func (s *UsersService) GetUserService(user models.UsersModel) (models.UsersModel, error) {
	res, err := s.usersDB.GetUserFromDB(user)
	return res, err
}

func (s *UsersService) AddUserService(user models.UsersModel) error {
	err := s.usersDB.AddUserToDB(user)
	return err
}

func (s *UsersService) DeleteUserService(user models.UsersModel) error {
	err := s.usersDB.DeleteUserFromDB(user)
	return err
}
