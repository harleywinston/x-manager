package services

import (
	"net"
	"regexp"
	"strings"

	"github.com/harleywinston/x-manager/internal/master/consts"
	"github.com/harleywinston/x-manager/internal/master/database"
	"github.com/harleywinston/x-manager/internal/master/models"
)

type ResourcesService struct {
	resourcesDB database.ResourceDB
}

func checkIp(ip string) error {
	if net.ParseIP(ip) == nil {
		return &consts.CustomError{
			Message: consts.INVALID_IP_ERROR.Message,
			Code:    consts.INVALID_IP_ERROR.Code,
			Detail:  ip,
		}
	}
	return nil
}

func checkDomain(domain string) error {
	re := regexp.MustCompile(`^[a-zA-Z0-9]+([-.][a-zA-Z0-9]+)*\.[a-zA-Z]{2,}$`)
	if re.MatchString(domain) {
		return nil
	}
	return &consts.CustomError{
		Message: consts.INVALID_DOMAIN_ERROR.Message,
		Code:    consts.INVALID_DOMAIN_ERROR.Code,
		Detail:  domain,
	}
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
