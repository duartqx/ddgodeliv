package postgresql

import (
	"fmt"
	"time"

	r "ddgodeliv/infrastructure/repository/interfaces"
)

type Session struct {
	UserId    int
	CreatedAt time.Time
	Token     string
}

func (s Session) GetUserId() int {
	return s.UserId
}

func (s Session) GetCreatedAt() time.Time {
	return s.CreatedAt
}

func (s Session) GetToken() string {
	return s.Token
}

type sessionStore struct {
	sessions *map[string]r.ISession
}

type SessionRepository struct {
	store *sessionStore
}

func GetNewSessionRepository() *SessionRepository {
	return &SessionRepository{
		store: &sessionStore{
			sessions: &map[string]r.ISession{},
		},
	}
}

func (sr SessionRepository) Get(token string) (r.ISession, error) {
	session, found := (*sr.store.sessions)[token]
	if !found {
		return nil, fmt.Errorf("Session Not Found!")
	}
	return session, nil
}

func (sr SessionRepository) Set(token string, createdAt time.Time, userId int) error {
	(*sr.store.sessions)[token] = &Session{
		UserId: userId, CreatedAt: createdAt, Token: token,
	}
	return nil
}

func (sr SessionRepository) Delete(token string) error {
	delete((*sr.store.sessions), token)
	return nil
}
