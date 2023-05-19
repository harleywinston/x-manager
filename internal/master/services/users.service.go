package services

import (
	"encoding/base64"
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/harleywinston/x-manager/internal/master/consts"
	"github.com/harleywinston/x-manager/internal/master/database"
	"github.com/harleywinston/x-manager/internal/master/helper"
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

func (s *UsersService) getUsersXuiConfigs(
	user models.Users,
	resource models.Resources,
) (string, error) {
	HTTPClient := &http.Client{}
	req, err := http.NewRequest(
		http.MethodGet,
		"https://"+strings.Split(strings.ReplaceAll(resource.Domains, " ", ""), "|")[0]+fmt.Sprintf(
			"/sub/%s",
			user.Email,
		),
		nil,
	)
	req.Header.Set(
		"Accept",
		"text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.7",
	)
	if err != nil {
		return "", &consts.CustomError{
			Message: consts.HTTP_CLIENT_ERROR.Message,
			Code:    consts.HTTP_CLIENT_ERROR.Code,
			Detail:  err.Error(),
		}
	}

	resp, err := HTTPClient.Do(req)
	if err != nil {
		return "", &consts.CustomError{
			Message: consts.HTTP_CLIENT_ERROR.Message,
			Code:    consts.HTTP_CLIENT_ERROR.Code,
			Detail:  err.Error(),
		}
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", &consts.CustomError{
			Message: consts.HTTP_CLIENT_ERROR.Message,
			Code:    consts.HTTP_CLIENT_ERROR.Code,
			Detail:  err.Error(),
		}
	}

	decodedBody, err := base64.StdEncoding.DecodeString(string(body))
	if err != nil {
		return "", &consts.CustomError{
			Message: consts.HTTP_CLIENT_ERROR.Message,
			Code:    consts.HTTP_CLIENT_ERROR.Code,
			Detail:  err.Error(),
		}
	}
	return string(decodedBody), nil
}

func (s *UsersService) GetUserConfigs(user models.Users) (string, error) {
	var res string
	user, err := s.usersDB.GetUserFromDB(user)
	if err != nil {
		return "", err
	}
	group, err := s.usersDB.GetUsersGroup(user)
	if err != nil {
		return "", err
	}
	resource, err := s.usersDB.GetUsersRecourse(user)
	if err != nil {
		return "", err
	}

	xuiConfs, err := s.getUsersXuiConfigs(user, resource)
	if err != nil {
		return "", err
	}
	res += xuiConfs

	sublinkHelper := helper.SublinkHelper{}
	if group.Mode == "InDirect" {
		for i, bridge := range strings.Split(resource.Bridges, "|") {
			data := strings.Split(bridge, ":")
			if len(data) < 4 {
				return "", &consts.CustomError{
					Message: consts.INVALID_BRIDGE_DATA.Message,
					Code:    consts.INVALID_BRIDGE_DATA.Code,
					Detail:  "",
				}
			}
			addr := data[0]
			port, err := strconv.ParseInt(data[1], 10, 64)
			if err != nil {
				return "", &consts.CustomError{
					Message: consts.PARSE_INT_ERROR.Message,
					Code:    consts.PARSE_INT_ERROR.Code,
					Detail:  err.Error(),
				}
			}
			hostName := data[2]
			sni := data[3]

			sublinkHelper.LinkSettings = append(sublinkHelper.LinkSettings, &helper.TrojanLink{
				Remark: fmt.Sprintf("Bridge %d", i+1),
				Addr:   addr,
				Port:   int(port),
				Passwd: user.Passwd,
				Path:   "/",
				Host:   hostName,
				SNI:    sni,
			})
		}
	} else {
		for i, bridge := range strings.Split(resource.ForeignBridges, "|") {
			data := strings.Split(bridge, ":")
			if len(data) < 4 {
				return "", &consts.CustomError{
					Message: consts.INVALID_BRIDGE_DATA.Message,
					Code:    consts.INVALID_BRIDGE_DATA.Code,
					Detail:  "",
				}
			}
			addr := data[0]
			port, err := strconv.ParseInt(data[1], 10, 64)
			if err != nil {
				return "", &consts.CustomError{
					Message: consts.PARSE_INT_ERROR.Message,
					Code:    consts.PARSE_INT_ERROR.Code,
					Detail:  err.Error(),
				}
			}
			hostName := data[2]
			sni := data[3]

			sublinkHelper.LinkSettings = append(sublinkHelper.LinkSettings, &helper.TrojanLink{
				Remark: fmt.Sprintf("CDN %d", i+1),
				Addr:   addr,
				Port:   int(port),
				Passwd: user.Passwd,
				Path:   "/",
				Host:   hostName,
				SNI:    sni,
			})
		}
	}
	res += sublinkHelper.GetConfigs()

	return res, nil
}

func (s *UsersService) DeleteUserService(user models.Users) error {
	if err := s.checkEmail(user.Email); err != nil {
		return err
	}
	return s.usersDB.DeleteUserFromDB(user)
}
