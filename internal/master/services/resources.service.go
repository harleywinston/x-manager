package services

import (
	"fmt"
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

func (s *ResourcesService) checkIp(ip string) error {
	if net.ParseIP(ip) == nil {
		return &consts.CustomError{
			Message: consts.INVALID_IP_ERROR.Message,
			Code:    consts.INVALID_IP_ERROR.Code,
			Detail:  ip,
		}
	}
	return nil
}

func (s *ResourcesService) checkDomain(domain string) error {
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

func (s *ResourcesService) checkBridge(bridge string) error {
	data := strings.Split(bridge, ":")
	if len(data) < 4 {
		return &consts.CustomError{
			Message: consts.INVALID_BRIDGE_DATA.Message,
			Code:    consts.INVALID_BRIDGE_DATA.Code,
			Detail:  "",
		}
	}
	if err := s.checkDomain(data[2]); err != nil {
		return err
	}
	if err := s.checkDomain(data[3]); err != nil {
		return err
	}
	if err := s.checkIp(data[0]); err != nil {
		return err
	}

	return nil
}

func (s *ResourcesService) checkDuplicateRecord(resource models.Resources) error {
	record, err := s.resourcesDB.GetResourceFromDB(resource)
	if err != nil {
		return nil
	}
	fmt.Println(record.ServerIp, resource.ServerIp)
	if record.ServerIp == resource.ServerIp {
		return &consts.CustomError{
			Message: consts.DUPLICATE_RECORD_ERROR.Message,
			Code:    consts.DUPLICATE_RECORD_ERROR.Code,
			Detail:  "",
		}
	}
	return nil
}

func (s *ResourcesService) AddResourcesService(resource models.Resources) error {
	if err := s.checkIp(resource.ServerIp); err != nil {
		return err
	}
	for _, bridge := range strings.Split(resource.Bridges, "|") {
		if err := s.checkBridge(bridge); err != nil {
			return err
		}
	}
	for _, bridge := range strings.Split(resource.ForeignBridges, "|") {
		if err := s.checkBridge(bridge); err != nil {
			return err
		}
	}
	for _, x := range strings.Split(strings.ReplaceAll(resource.Domains, " ", ""), "|") {
		if err := s.checkDomain(strings.Split(x, ":")[0]); err != nil {
			return err
		}
	}

	if err := s.checkDuplicateRecord(resource); err != nil {
		return err
	}

	return s.resourcesDB.AddResourceToDB(resource)
}

func (s *ResourcesService) GetResourcesService(
	resource models.Resources,
) (models.Resources, error) {
	if err := s.checkIp(resource.ServerIp); err != nil {
		return models.Resources{}, err
	}
	// for _, x := range strings.Split(strings.ReplaceAll(resource.Domains, " ", ""), ",") {
	// 	if err := s.checkDomain(x); err != nil {
	// 		return models.Resources{}, err
	// 	}
	// }

	return s.resourcesDB.GetResourceFromDB(resource)
}

func (s *ResourcesService) DeleteResourcesService(resource models.Resources) error {
	if err := s.checkIp(resource.ServerIp); err != nil {
		return err
	}
	// for _, x := range strings.Split(strings.ReplaceAll(resource.Domains, " ", ""), ",") {
	// 	if err := s.checkDomain(x); err != nil {
	// 		return err
	// 	}
	// }

	return s.resourcesDB.DeleteResourceFromDB(resource)
}
