package auth

import (
	"context"
	"fmt"

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

func (cs SessionService) GetSessionUser(ctx context.Context) a.ISessionUser {
	user, ok := ctx.Value("user").(*a.SessionUser)
	if !ok {
		return nil
	}
	session, _ := cs.sessionRepository.Get(user)
	return session.GetUser()
}

func (cs SessionService) GetSessionUserWithCompany(ctx context.Context) a.ISessionUser {
	user, ok := ctx.Value("user").(*a.SessionUser)
	if !ok {
		return nil
	}
	session, _ := cs.sessionRepository.Get(user)

	if session.GetUser().HasInvalidCompany() {
		driver, err := cs.driverRepository.FindByUserId(user.GetId())
		if err != nil || driver.GetCompanyId() == 0 {
			return nil
		}
		// Updates the sessionUser with the Driver information
		cs.sessionRepository.Set(
			session.GetUser().SetDriver(driver), session.GetCreatedAt(),
		)
	}

	return session.GetUser()
}

func (cs SessionService) GetSessionUserWithoutCompany(ctx context.Context) (a.ISessionUser, error) {
	user, ok := ctx.Value("user").(*a.SessionUser)
	if !ok {
		return nil, fmt.Errorf("User not found: %w", e.ForbiddenError)
	}
	session, _ := cs.sessionRepository.Get(user)

	if session.GetUser().HasInvalidCompany() {
		return session.GetUser(), nil
	}

	return nil, fmt.Errorf("User has company: %w", e.ForbiddenError)
}
