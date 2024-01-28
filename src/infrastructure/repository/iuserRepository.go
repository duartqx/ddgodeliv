package repository

import (
	u "ddgodeliv/domains/user"
)

type IUserRepository interface {
	FindById(id int) (u.IUser, error)

	FindByEmail(email string) (u.IUser, error)
	ExistsByEmail(email string) bool

	Create(user u.IUser) error
	Update(user u.IUser) error
	Delete(user u.IUser) error
}
