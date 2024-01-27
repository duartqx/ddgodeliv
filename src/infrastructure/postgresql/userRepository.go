package postgresql

import (
	"fmt"
	"strings"

	"github.com/jmoiron/sqlx"

	u "ddgodeliv/src/domains/user"
)

type UserRepository struct {
	db *sqlx.DB
}

func GetNewUserRepository(db *sqlx.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (ur UserRepository) FindByID(id int) (u.IUser, error) {

	user := u.GetNewUser()

	if err := ur.db.Get(user, "SELECT * FROM users WHERE id = $1", id); err != nil {
		return nil, err
	}

	return user, nil
}

func (ur UserRepository) FindByEmail(email string) (u.IUser, error) {

	user := u.GetNewUser()

	if err := ur.db.Get(user, "SELECT * FROM users WHERE email = $1", email); err != nil {
		return nil, err
	}

	return user, nil
}

func (ur UserRepository) ExistsByEmail(email string) (exists *bool) {
	ur.db.QueryRow(
		"SELECT EXISTS (SELECT 1 FROM users WHERE email = $1)",
		email,
	).Scan(&exists)

	return exists
}

func (ur UserRepository) Create(user u.IUser) error {

	if user.GetEmail() == "" || user.GetName() == "" || user.GetPassword() == "" {
		return fmt.Errorf("Invalid User")
	}

	var id int

	if err := ur.db.QueryRow(
		`
			INSERT INTO users (email, name, password)
			VALUES ($1, $2, $3)
			RETURNING id
		`,
		strings.ToLower(user.GetEmail()),
		user.GetName(),
		user.GetPassword(),
	).Scan(&id); err != nil {
		return err
	}

	user.SetId(id)

	return nil
}

func (ur UserRepository) Update(user u.IUser) error {

	if user.GetId() == 0 {
		return fmt.Errorf("Invalid user id")
	}

	_, err := ur.db.Exec(
		"UPDATE users SET email = $1, name = $2 WHERE id = $3",
		strings.ToLower(user.GetEmail()),
		user.GetName(),
		user.GetId(),
	)

	return err
}

func (ur UserRepository) Delete(user u.IUser) error {

	if user.GetId() == 0 {
		return fmt.Errorf("Invalid User")
	}

	_, err := ur.db.Exec("DELETE FROM users WHERE id = $1", user.GetId())

	return err
}
