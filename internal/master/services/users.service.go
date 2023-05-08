package services

import (
	"fmt"
	"math/rand"
	"regexp"
	"strings"
	"time"

	"github.com/harleywinston/x-manager/internal/master/consts"
	"github.com/harleywinston/x-manager/internal/master/database"
	"github.com/harleywinston/x-manager/internal/master/models"
)

type UsersService struct {
	usersDB database.UsersDB
}

func (s *UsersService) generateUserPasswd() string {
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

func (s *UsersService) getExpiryTime() int64 {
	now := time.Now()
	future := now.AddDate(0, 0, 30)
	return future.UnixNano() / int64(time.Millisecond)
}

func (s *UsersService) checkEmail(email string) error {
	emailRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	if emailRegex.MatchString(email) {
		return nil
	}
	return &consts.CustomError{
		Message: consts.INVALID_EMAIL_ERROR.Message,
		Code:    consts.INVALID_EMAIL_ERROR.Code,
		Detail:  fmt.Sprintf(`Email: %s`, email),
	}
}

func (s *UsersService) GetUserService(user models.Users) (models.Users, error) {
	if err := s.checkEmail(user.Email); err != nil {
		return models.Users{}, err
	}
	return s.usersDB.GetUserFromDB(user)
}

func (s *UsersService) AddUserService(user models.Users) error {
	user.Passwd = s.generateUserPasswd()
	user.ExpiryTime = s.getExpiryTime()

	groupID, err := s.usersDB.GetFreeGroupIDFromDB()
	if err != nil {
		return err
	}
	fmt.Println(groupID)
	user.GroupID = groupID

	if err := s.checkEmail(user.Email); err != nil {
		return err
	}

	if err := s.usersDB.AddUserToDB(user); err != nil {
		return err
	}
	return nil
}

func (s *UsersService) DeleteUserService(user models.Users) error {
	if err := s.checkEmail(user.Email); err != nil {
		return err
	}
	return s.usersDB.DeleteUserFromDB(user)
}
