package postgresql

import (
	"fmt"
	"strings"

	"github.com/jmoiron/sqlx"

	c "ddgodeliv/src/domains/company"
)

type CompanyRepository struct {
	db *sqlx.DB
}

func GetNewCompanyRepository(db *sqlx.DB) *CompanyRepository {
	return &CompanyRepository{db: db}
}

func (cr CompanyRepository) FindById(id int) (c.ICompany, error) {
	company := c.GetNewCompany()

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
