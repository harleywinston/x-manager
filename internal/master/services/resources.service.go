package services

import (
	"strings"

	"github.com/harleywinston/x-manager/internal/master/database"
	"github.com/harleywinston/x-manager/internal/master/models"
)

type ResourcesService struct {
	resourcesDB database.ResourceDB
}

func checkIp(ip string) error {
	return nil
}

func checkDomain(domain string) error {
	return nil
}

func (s *ResourcesService) AddResourcesService(resource models.Resources) error {
	if err := checkIp(resource.ServerIp); err != nil {
		return err
	}
	for _, x := range strings.Split(strings.ReplaceAll(resource.Domains, " ", ""), ",") {
		if err := checkDomain(x); err != nil {
			return err
		}
	}

	return s.resourcesDB.AddResourceToDB(resource)
}

func (s *ResourcesService) GetResourcesService(
	resource models.Resources,
) (models.Resources, error) {
	if err := checkIp(resource.ServerIp); err != nil {
		return models.Resources{}, err
	}
	for _, x := range strings.Split(strings.ReplaceAll(resource.Domains, " ", ""), ",") {
		if err := checkDomain(x); err != nil {
			return models.Resources{}, err
		}
	}

	return s.resourcesDB.GetResourceFromDB(resource)
}

func (s *ResourcesService) DeleteResourcesService(resource models.Resources) error {
	if err := checkIp(resource.ServerIp); err != nil {
		return err
	}
	for _, x := range strings.Split(strings.ReplaceAll(resource.Domains, " ", ""), ",") {
		if err := checkDomain(x); err != nil {
			return err
		}
	}

	return s.resourcesDB.DeleteResourceFromDB(resource)
}
