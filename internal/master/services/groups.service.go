package services

import (
	"github.com/harleywinston/x-manager/internal/master/database"
	"github.com/harleywinston/x-manager/internal/master/models"
)

type GroupsService struct {
	groupsDB database.GroupsDB
}

func checkGroupMode(mode string) error {
	return nil
}

func checkResourceID(id int) error {
	return nil
}

func (s *GroupsService) AddGroupsService(group models.Groups) error {
	if err := checkGroupMode(group.Mode); err != nil {
		return err
	}
	if err := checkResourceID(group.ResourcesID); err != nil {
		return err
	}

	return s.groupsDB.AddGroupToDB(group)
}

func (s *GroupsService) GetGroupsService(group models.Groups) (models.Groups, error) {
	if err := checkGroupMode(group.Mode); err != nil {
		return models.Groups{}, err
	}
	if err := checkResourceID(group.ResourcesID); err != nil {
		return models.Groups{}, err
	}

	return s.groupsDB.GetGroupFromDB(group)
}

func (s *GroupsService) DeleteGroupsService(group models.Groups) error {
	if err := checkGroupMode(group.Mode); err != nil {
		return err
	}
	if err := checkResourceID(group.ResourcesID); err != nil {
		return err
	}

	return s.groupsDB.DeleteGroupFromDB(group)
}
