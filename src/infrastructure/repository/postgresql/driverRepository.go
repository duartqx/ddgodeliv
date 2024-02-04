package postgresql

import (
	"fmt"
	"strings"

	"github.com/jmoiron/sqlx"

	d "ddgodeliv/domains/driver"
)

type DriverRepository struct {
	db *sqlx.DB
}

func GetNewDriverRepository(db *sqlx.DB) *DriverRepository {
	return &DriverRepository{db: db}
}

func (dr DriverRepository) baseJoinedQuery(where string) string {
	return fmt.Sprintf(
		`
			SELECT
				d.id AS "id",
				d.user_id AS "user_id",
				d.company_id AS "company_id",
				d.license_id AS "licence_id",

				u.id AS "user.id",
				u.name AS "user.name",
				u.email AS "user.email",
				
				c.id AS "company.id",
				c.owner_id AS "company.owner_id",
				c.name AS "company.name"
			FROM drivers d
			INNER JOIN users u ON d.user_id = u.id
			INNER JOIN companies c ON d.company_id = c.id
			WHERE %s
		`,
		where,
	)
}

func (dr DriverRepository) FindById(id, companyId int) (d.IDriver, error) {

	driver := d.GetNewDriver()

	if err := dr.db.Get(
		driver, dr.baseJoinedQuery("id = $1 AND company_id = $2"), id, companyId,
	); err != nil {
		return nil, err
	}

	return driver, nil
}

func (dr DriverRepository) FindByUserId(id int) (d.IDriver, error) {
	driver := d.GetNewDriver()

	if err := dr.db.Get(
		driver, dr.baseJoinedQuery("user_id = $1"), id,
	); err != nil {
		return nil, err
	}

	return driver, nil
}

func (dr DriverRepository) ExistsByUserId(id int) (exists bool) {
	dr.db.QueryRow(
		"SELECT EXISTS (SELECT 1 FROM drivers WHERE user_id = $1)",
		id,
	).Scan(&exists)

	return exists
}

func (dr DriverRepository) FindByCompanyId(id int) (*[]d.IDriver, error) {
	drivers := []d.IDriver{}

	rows, err := dr.db.Queryx(dr.baseJoinedQuery("company_id = $1"), id)
	if err != nil {
		return nil, err
	}

	for rows.Next() {

		driver := d.GetNewDriver()

		if err := rows.StructScan(driver); err != nil {
			return nil, err
		}

		var castedDriver d.IDriver = driver

		drivers = append(drivers, castedDriver)
	}

	return &drivers, nil
}

func (dr DriverRepository) ExistsByCompanyId(id int) (exists bool) {
	dr.db.QueryRow(
		"SELECT EXISTS (SELECT 1 FROM drivers WHERE company_id = $1)",
		id,
	).Scan(&exists)

	return exists
}

func (dr DriverRepository) Create(driver d.IDriver) error {
	var id int

	if err := dr.db.QueryRow(
		`
			INSERT INTO drivers (user_id, company_id, license_id)
			VALUES ($1, $2, $3)
			RETURNING id
		`,
		driver.GetUserId(),
		driver.GetCompanyId(),
		strings.ToLower(driver.GetLicenseId()),
	).Scan(&id); err != nil {
		return err
	}

	driver.SetId(id)

	return nil
}

func (dr DriverRepository) Update(driver d.IDriver) error {
	if driver.GetId() == 0 {
		return fmt.Errorf("Invalid Driver Id")
	}

	_, err := dr.db.Exec(
		"UPDATE drivers SET license_id = $1 WHERE id = $2",
		strings.ToLower(driver.GetLicenseId()),
		driver.GetId(),
	)
	return err
}

func (dr DriverRepository) Delete(driver d.IDriver) error {
	if driver.GetId() == 0 {
		return fmt.Errorf("Invalid Driver Id")
	}

	_, err := dr.db.Exec("DELETE FROM drivers WHERE id = $1", driver.GetId())

	return err
}
