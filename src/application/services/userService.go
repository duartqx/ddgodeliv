package services

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"

	v "ddgodeliv/application/validation"
	u "ddgodeliv/domains/user"
	r "ddgodeliv/infrastructure/repository"
)

type UserService struct {
	userRepository r.IUserRepository
	*v.Validator
}

func GetNewUserService(userRepository r.IUserRepository, validator *v.Validator) *UserService {
	return &UserService{
		userRepository: userRepository,
		Validator:      validator,
	}
}

func (us UserService) Validate(user u.IUser) error {
	if errs := us.Struct(user); errs != nil {
		return fmt.Errorf(string(*us.JSON(errs)))
	}
	return nil
}

func (us UserService) Create(user u.IUser) error {

	if err := us.Validate(user); err != nil {
		return err
	}

	if us.userRepository.ExistsByEmail(user.GetEmail()) {
		return fmt.Errorf("Invalid Email")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword(
		[]byte(user.GetPassword()), 10,
	)
	if err != nil {
		return fmt.Errorf("Invalid Password")
	}

	user.SetPassword(string(hashedPassword))

	if err := us.userRepository.Create(user); err != nil {
		return fmt.Errorf("Internal Error trying to create user")
	}

	return nil
}

func (us UserService) UpdatePassword(user u.IUser) error {
	if err := us.Validator.Var(user.GetPassword(), "required,min=8,max=200"); err != nil {
		return fmt.Errorf("Invalid Password")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.GetPassword()), 10)
	if err != nil {
		return fmt.Errorf("Invalid Password")
	}

	user.SetPassword(string(hashedPassword))

	if err := us.userRepository.Update(user); err != nil {
		return fmt.Errorf("Internal Error trying to update the password")
	}

	return nil
}

func (us UserService) FindById(id int) (u.IUser, error) {
	return us.userRepository.FindById(id)
}
