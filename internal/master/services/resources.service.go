package services

import (
	"github.com/harleywinston/x-manager/internal/master/database"
	"github.com/harleywinston/x-manager/internal/master/models"
)

type ResourcesService struct {
	resourceDB database.ResourceDB
}

func (s *ResourcesService) AddResourcesService(resource models.Resources) error {
	return nil
}

func (s *ResourcesService) GetResourcesService(
	resource models.Resources,
) (models.Resources, error) {
	return models.Resources{}, nil
}

func (s *ResourcesService) DeleteResourcesService(resource models.Resources) error {
	return nil
}
