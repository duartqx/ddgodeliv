package postgresql

import (
	"sync"
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
	mutex    *sync.Mutex
}

type SessionRepository struct {
	store *sessionStore
}

func GetNewSessionRepository() *SessionRepository {
	return &SessionRepository{
		store: &sessionStore{
			sessions: &map[string]a.ISession{},
			mutex:    &sync.Mutex{},
		},
	}
}

func (sr SessionRepository) Get(user a.ISessionUser) (a.ISession, error) {

	sr.store.mutex.Lock()
	defer sr.store.mutex.Unlock()

	session, found := (*sr.store.sessions)[user.GetEmail()]
	if !found {
		createdAt := time.Now()
		session = &Session{User: user, CreatedAt: createdAt}
		(*sr.store.sessions)[user.GetEmail()] = session
	}
	return session, nil
}

func (sr SessionRepository) Set(user a.ISessionUser, createdAt time.Time) error {

	sr.store.mutex.Lock()
	defer sr.store.mutex.Unlock()

	(*sr.store.sessions)[user.GetEmail()] = &Session{
		User: user, CreatedAt: createdAt,
	}
	return nil
}

func (sr SessionRepository) Delete(user a.ISessionUser) error {

	sr.store.mutex.Lock()
	defer sr.store.mutex.Unlock()

	delete((*sr.store.sessions), user.GetEmail())
	return nil
}
