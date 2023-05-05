package database

import (
	"fmt"
	"os"
	"strconv"

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

func (db *UsersDB) GetFreeGroupIDFromDB() (int, error) {
	type groupWithCount struct {
		GroupID int
		Count   int
	}
	var groupsWithCount []groupWithCount

	err := DB.Model(&models.Users{}).
		Select("group_id, count(*) as count").
		Group("group_id").
		Scan(&groupsWithCount).Error
	if err != nil {
		return 0, err
	}

	var resGroups []int
	for _, g := range groupsWithCount {
		group_limit, err := strconv.ParseInt(os.Getenv("GROUPS_LIMIT"), 10, 64)
		if err != nil {
			return 0, err
		}
		if g.Count < int(group_limit) {
			resGroups = append(resGroups, g.GroupID)
		}
	}

	if len(resGroups) < 1 {
		return 0, fmt.Errorf("All groups limit exceeded!")
	}

	return resGroups[0], nil
}
