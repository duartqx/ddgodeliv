package auth

import (
	"context"
	"fmt"
	"time"

	e "ddgodeliv/application/errors"
	a "ddgodeliv/domains/auth"
	d "ddgodeliv/domains/driver"
)

type SessionService struct {
	driverRepository  d.IDriverRepository
	sessionRepository a.ISessionRepository
}

func GetNewSessionService(
	driverRepository d.IDriverRepository,
	sessionRepository a.ISessionRepository,
) *SessionService {
	return &SessionService{
		driverRepository:  driverRepository,
		sessionRepository: sessionRepository,
	}
}

func (ss SessionService) GetSessionUser(ctx context.Context) a.ISessionUser {
	user, ok := ctx.Value("user").(*a.SessionUser)
	if !ok {
		return nil
	}
	session, _ := ss.sessionRepository.Get(user)
	return session.GetUser()
}

func (ss SessionService) GetSessionUserWithCompany(ctx context.Context) a.ISessionUser {
	user, ok := ctx.Value("user").(*a.SessionUser)
	if !ok {
		return nil
	}
	session, _ := ss.sessionRepository.Get(user)

	if session.GetUser().HasInvalidCompany() {
		driver, err := ss.driverRepository.FindByUserId(user.GetId())
		if err != nil || driver.GetCompanyId() == 0 {
			return nil
		}
		// Updates the sessionUser with the Driver information
		ss.sessionRepository.Set(
			session.GetUser().SetDriver(driver), session.GetCreatedAt(),
		)
	}

	return session.GetUser()
}

func (ss SessionService) GetSessionUserWithoutCompany(ctx context.Context) (a.ISessionUser, error) {
	user, ok := ctx.Value("user").(*a.SessionUser)
	if !ok {
		return nil, fmt.Errorf("User not found: %w", e.ForbiddenError)
	}
	session, _ := ss.sessionRepository.Get(user)

	if !session.GetUser().HasInvalidCompany() {
		driver, _ := ss.driverRepository.FindByUserId(user.GetId())
		if driver != nil {
			return nil, fmt.Errorf("User has company: %w", e.ForbiddenError)
		}
	}

	return session.GetUser().ResetDriver(), nil
}

func (ss SessionService) SetSessionUserCompany(user a.ISessionUser) error {
	driver, _ := ss.driverRepository.FindByUserId(user.GetId())
	if driver != nil {
		return ss.sessionRepository.Set(user.SetDriver(driver), time.Now())
	}
	return nil
}

func (ss SessionService) SetSessionUserNoCompany(user a.ISessionUser) error {
	resetedUser := a.GetNewSessionUser().
		SetId(user.GetId()).
		SetName(user.GetName()).
		SetEmail(user.GetEmail())
	return ss.sessionRepository.Set(resetedUser, time.Now())
}
