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

func GetNewUserRepository(db *sqlx.DB) (*UserRepository, error) {
	if _, err := db.Exec(`
		CREATE TABLE IF NOT EXISTS Users (
			id SERIAL PRIMARY KEY,
			email VARCHAR(255) NOT NULL UNIQUE,
			password VARCHAR(255) NOT NULL,
			name VARCHAR(255) NOT NULL
		);
	`); err != nil {
		return nil, err
	}

	return &UserRepository{db: db}, nil
}

func (ur UserRepository) getUserModel() *u.User {
	return &u.User{}
}

func (ur UserRepository) FindByID(id int) (u.IUser, error) {

	user := ur.getUserModel()

	if err := ur.db.Get(user, "SELECT * FROM users WHERE id = $1", id); err != nil {
		return nil, err
	}

	return user, nil
}

func (ur UserRepository) FindByEmail(email string) (u.IUser, error) {

	user := ur.getUserModel()

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

	if _, err := ur.db.Exec(
		"UPDATE users SET email = $1, name = $2 WHERE id = $3",
		strings.ToLower(user.GetEmail()),
		user.GetName(),
		user.GetId(),
	); err != nil {
		return err
	}

	return nil
}

func (ur UserRepository) Delete(user u.IUser) error {

	if user.GetId() == 0 {
		return fmt.Errorf("Invalid User")
	}

	if _, err := ur.db.Exec("DELETE FROM users WHERE id = $1", user.GetId()); err != nil {
		return err
	}

	return nil
}
