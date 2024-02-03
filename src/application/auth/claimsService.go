package auth

import (
	"context"
	"fmt"

	"ddgodeliv/domains/driver"
)

type ClaimsService struct {
	driverRepository driver.IDriverRepository
}

func GetNewClaimsService(driverRepository driver.IDriverRepository) *ClaimsService {
	return &ClaimsService{driverRepository: driverRepository}
}

func (cs ClaimsService) GetClaimsUserFromContext(ctx context.Context) (*ClaimsUser, error) {
	user, ok := ctx.Value("user").(*ClaimsUser)
	if !ok {
		return nil, fmt.Errorf("User is not set")
	}
	return user, nil
}

func (cs ClaimsService) GetWithDriverInfo(user *ClaimsUser) error {
	if user.HasInvalidCompany() {
		driver, err := cs.driverRepository.FindByUserId(user.GetId())
		if err != nil || driver.GetCompanyId() == 0 {
			return err
		}
		user.SetDriver(driver)
	}
	return nil
}
