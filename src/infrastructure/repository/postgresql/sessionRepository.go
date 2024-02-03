package postgresql

import (
	"time"

	a "ddgodeliv/domains/auth"
)

type Session struct {
	User      a.ISessionUser
	CreatedAt time.Time
}

func (s Session) GetUser() a.ISessionUser {
	return s.User
}

func (s Session) GetCreatedAt() time.Time {
	return s.CreatedAt
}

type sessionStore struct {
	sessions *map[string]a.ISession
}

type SessionRepository struct {
	store *sessionStore
}

func GetNewSessionRepository() *SessionRepository {
	return &SessionRepository{
		store: &sessionStore{
			sessions: &map[string]a.ISession{},
		},
	}
}

func (sr SessionRepository) Get(user a.ISessionUser) (a.ISession, error) {
	session, found := (*sr.store.sessions)[user.GetEmail()]
	if !found {
		session = &Session{User: user, CreatedAt: time.Now()}
	}
	return session, nil
}

func (sr SessionRepository) Set(user a.ISessionUser, createdAt time.Time) error {
	(*sr.store.sessions)[user.GetEmail()] = &Session{
		User: user, CreatedAt: createdAt,
	}
	return nil
}

func (sr SessionRepository) Delete(user a.ISessionUser) error {
	delete((*sr.store.sessions), user.GetEmail())
	return nil
}
