package services

import (
	"fmt"
	"os"
	"strings"

	"github.com/harleywinston/x-manager/internal/master/database"
	"github.com/harleywinston/x-manager/internal/master/models"
)

type GroupsService struct {
	groupsDB database.GroupsDB
}

func checkGroupMode(mode string) error {
	env := os.Getenv("GROUP_MODES")
	groupModes := strings.Split(strings.ReplaceAll(env, " ", ""), ",")

	for _, gM := range groupModes {
		if mode == gM {
			return nil
		}
	}
	return fmt.Errorf(`Invalid group mode %v!`, mode)
}

func (s *GroupsService) checkResourceID(id int) error {
	if err := s.groupsDB.CheckResourceID(id); err != nil {
		return err
	}
	return nil
}

func (s *GroupsService) AddGroupsService(group models.Groups) error {
	if err := checkGroupMode(group.Mode); err != nil {
		return err
	}
	if err := s.checkResourceID(group.ResourcesID); err != nil {
		return err
	}

	return s.groupsDB.AddGroupToDB(group)
}

func (s *GroupsService) GetGroupsService(group models.Groups) (models.Groups, error) {
	if err := checkGroupMode(group.Mode); err != nil {
		return models.Groups{}, err
	}
	if err := s.checkResourceID(group.ResourcesID); err != nil {
		return models.Groups{}, err
	}

	return s.groupsDB.GetGroupFromDB(group)
}

func (s *GroupsService) DeleteGroupsService(group models.Groups) error {
	if err := checkGroupMode(group.Mode); err != nil {
		return err
	}
	if err := s.checkResourceID(group.ResourcesID); err != nil {
		return err
	}

	return s.groupsDB.DeleteGroupFromDB(group)
}
