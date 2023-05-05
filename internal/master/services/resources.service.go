package services

import (
	"fmt"
	"net"
	"regexp"
	"strings"

	"github.com/harleywinston/x-manager/internal/master/database"
	"github.com/harleywinston/x-manager/internal/master/models"
)

type ResourcesService struct {
	resourcesDB database.ResourceDB
}

func checkIp(ip string) error {
	if net.ParseIP(ip) == nil {
		return fmt.Errorf("Resource ip {%v} is not valid", ip)
	}
	return nil
}

func checkDomain(domain string) error {
	re := regexp.MustCompile(`^([a-zA-Z0-9]+(-[a-zA-Z0-9]+)*\\.)+[a-zA-Z]{2,}$`)
	if re.MatchString(domain) {
		return nil
	}
	return fmt.Errorf("Resource domain address {%v} is not valid", domain)
}

func (s *ResourcesService) AddResourcesService(resource models.Resources) error {
	if err := checkIp(resource.ServerIp); err != nil {
		return err
	}
	if err := checkDomain(resource.BrdigeDomain); err != nil {
		return err
	}
	for _, x := range strings.Split(strings.ReplaceAll(resource.Domains, " ", ""), ",") {
		if err := checkDomain(x); err != nil {
			return err
		}
	}
	for _, x := range strings.Split(strings.ReplaceAll(resource.CloudflareDomains, " ", ""), ",") {
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
