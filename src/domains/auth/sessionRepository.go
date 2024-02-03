package auth

import "time"

type ISession interface {
	GetUser() ISessionUser
	GetCreatedAt() time.Time
}

type ISessionRepository interface {
	Get(user ISessionUser) (ISession, error)
	Set(user ISessionUser, createdAt time.Time) error
	Delete(user ISessionUser) error
}
