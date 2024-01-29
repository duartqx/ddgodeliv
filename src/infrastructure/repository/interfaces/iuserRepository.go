package repository

import (
	m "ddgodeliv/domains/models"
)

type IUserRepository interface {
	FindById(id int) (m.IUser, error)

	FindByEmail(email string) (m.IUser, error)
	ExistsByEmail(email string) bool

	Create(user m.IUser) error
	Update(user m.IUser) error
	Delete(user m.IUser) error
}
