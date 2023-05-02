package services

import (
	"math/rand"
	"strings"

	"github.com/harleywinston/x-manager/internal/master/database"
	"github.com/harleywinston/x-manager/internal/master/models"
)

type UsersService struct {
	UsersDB database.UsersDB
}

func (s *UsersService) GetUserService(user models.UsersModel) (models.UsersModel, error) {
	res, err := s.UsersDB.GetUserFromDB(user)
	return res, err
}

func generateUserPasswd() string {
	var password strings.Builder
	lowerCharSet := "abcdedfghijklmnopqrst"
	upperCharSet := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	specialCharSet := "!@#$%&*"
	numberSet := "0123456789"
	allCharSet := lowerCharSet + upperCharSet + specialCharSet + numberSet

	// Set special character
	for i := 0; i < 3; i++ {
		random := rand.Intn(len(specialCharSet))
		password.WriteString(string(specialCharSet[random]))
	}

	// Set numeric
	for i := 0; i < 4; i++ {
		random := rand.Intn(len(numberSet))
		password.WriteString(string(numberSet[random]))
	}

	// Set uppercase
	for i := 0; i < 2; i++ {
		random := rand.Intn(len(upperCharSet))
		password.WriteString(string(upperCharSet[random]))
	}

	for i := 0; i < 3; i++ {
		random := rand.Intn(len(allCharSet))
		password.WriteString(string(allCharSet[random]))
	}
	inRune := []rune(password.String())
	rand.Shuffle(len(inRune), func(i, j int) {
		inRune[i], inRune[j] = inRune[j], inRune[i]
	})
	return string(inRune)
}

func (s *UsersService) AddUserService(user models.UsersModel) error {
	user.Passwd = generateUserPasswd()
	err := s.UsersDB.AddUserToDB(user)
	return err
}

func (s *UsersService) DeleteUserService(user models.UsersModel) error {
	err := s.UsersDB.DeleteUserFromDB(user)
	return err
}
