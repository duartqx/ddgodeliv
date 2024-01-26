package postgresql

import (
	c "ddgodeliv/src/domains/company"
	"fmt"
	"strings"

	"github.com/jmoiron/sqlx"
)

type CompanyRepository struct {
	db *sqlx.DB
}

func GetNewCompanyRepository(db *sqlx.DB) (*CompanyRepository, error) {
	if _, err := db.Exec(`
		CREATE TABLE IF NOT EXISTS Companies (
			id SERIAL PRIMARY KEY,
			name VARCHAR(255) NOT NULL UNIQUE,
		);
	`); err != nil {
		return nil, err
	}

	return &CompanyRepository{db: db}, nil
}

func (cr CompanyRepository) getModel() *c.Company {
	return &c.Company{}
}

func (cr CompanyRepository) FindById(id int) (c.ICompany, error) {
	company := cr.getModel()

	if err := cr.db.Get(company, "SELECT * FROM companies WHERE id = $1", id); err != nil {
		return nil, err
	}

	return company, nil
}

func (cr CompanyRepository) Create(company c.ICompany) error {
	if company.GetName() == "" {
		return fmt.Errorf("Invalid Company Name")
	}

	var id int

	if err := cr.db.QueryRow(
		"INSERT INTO companies (name) VALUES ($1) RETURNING id",
		strings.ToLower(company.GetName()),
	).Scan(&id); err != nil {
		return err
	}

	company.SetId(id)

	return nil
}

func (cr CompanyRepository) Delete(company c.ICompany) error {
	if company.GetId() == 0 {
		return fmt.Errorf("Invalid Company Id")
	}
	_, err := cr.db.Exec("DELETE FROM companies WHERE id = $1", company.GetId())

	return err
}
