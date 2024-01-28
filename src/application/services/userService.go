package services

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"

	u "ddgodeliv/domains/user"
	r "ddgodeliv/infrastructure/repository"
)

type UserService struct {
	userRepository r.IUserRepository
}

func GetNewUserService(userRepository r.IUserRepository) *UserService {
	return &UserService{
		userRepository: userRepository,
	}
}

func (us UserService) Create(user u.IUser) error {

	if user.GetEmail() == "" || *us.userRepository.ExistsByEmail(user.GetEmail()) {
		return fmt.Errorf(`{"error": "Bad Request"}`)
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.GetPassword()), 10)
	if err != nil {
		return fmt.Errorf(`{"error": "Bad Request"}`)
	}

	user.SetPassword(string(hashedPassword))

	if err := us.userRepository.Create(user); err != nil {
		return fmt.Errorf(`{"error": "Bad Request"}`)
	}

	return nil
}
