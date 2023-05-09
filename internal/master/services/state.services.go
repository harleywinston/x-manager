package services

import (
	"github.com/harleywinston/x-manager/internal/master/database"
	"github.com/harleywinston/x-manager/internal/master/models"
)

type StateServices struct {
	stateDB database.StateDatabase
}

func (s *StateServices) GetStateService(ip string) (models.State, error) {
	resourceID, err := s.stateDB.GetResourceID(ip)
	if err != nil {
		return models.State{}, err
	}
	groupID, err := s.stateDB.GetGroupID(resourceID)
	if err != nil {
		return models.State{}, err
	}
	users, err := s.stateDB.GetUsers(groupID)
	if err != nil {
		return models.State{}, err
	}

	return models.State{Users: users}, nil
}
