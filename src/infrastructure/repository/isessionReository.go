package repository

import "time"

type ISession interface {
	GetUserId() int
	GetCreatedAt() time.Time
	GetToken() string
}

type ISessionRepository interface {
	Get(token string) (ISession, error)
	Set(token string, createdAt time.Time, userId int) error
	Delete(token string) error
}
