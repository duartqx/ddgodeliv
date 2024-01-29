package postgresql

import (
	"fmt"
	"strings"

	"github.com/jmoiron/sqlx"

	c "ddgodeliv/domains/company"
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

func (cr CompanyRepository) ExistsByName(name string) (exists bool) {
	cr.db.QueryRow(
		"SELECT EXISTS (SELECT 1 FROM companies WHERE name = $1)",
		name,
	).Scan(&exists)

	return exists
}

func (cr CompanyRepository) Create(company c.ICompany, licenseId string) error {
	var id int

	// Creates Company and Driver for User with id = ownerId
	if err := cr.db.QueryRow(
		`
			WITH new_company AS (
				INSERT INTO companies (name, owner_id)
				VALUES ($1, $2)
				RETURNING id
			)
			INSERT INTO drivers (user_id, license_id, company_id)
			SELECT $2, $3, id FROM new_company
			RETURNING company_id
		`,
		strings.ToLower(company.GetName()),
		company.GetOwnerId(),
		licenseId,
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
