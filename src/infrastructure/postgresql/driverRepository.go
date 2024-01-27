package postgresql

import (
	"fmt"
	"strings"

	"github.com/jmoiron/sqlx"

	d "ddgodeliv/src/domains/driver"
)

type DriverRepository struct {
	db *sqlx.DB
}

func GetNewDriverRepository(db *sqlx.DB) *DriverRepository {
	return &DriverRepository{db: db}
}

func (dr DriverRepository) simpleValidation(driver d.IDriver) error {
	if driver.GetUserId() == 0 || driver.GetCompanyId() == 0 || driver.GetLicenseId() == "" {
		return fmt.Errorf("Invalid Driver, missing user or license")
	}
	return nil
}

func (dr DriverRepository) FindById(id int) (d.IDriver, error) {

	driver := d.GetNewDriver()

	if err := dr.db.Get(driver, "SELECT * FROM drivers WHERE id = $1", id); err != nil {
		return nil, err
	}

	return driver, nil
}

func (dr DriverRepository) FindByUserId(id int) (d.IDriver, error) {
	driver := d.GetNewDriver()

	if err := dr.db.Get(driver, "SELECT * FROM drivers WHERE user_id = $1", id); err != nil {
		return nil, err
	}

	return driver, nil
}

func (dr DriverRepository) ExistsByUserId(id int) (exists *bool) {
	dr.db.QueryRow(
		"SELECT EXISTS (SELECT 1 FROM drivers WHERE user_id = $1)",
		id,
	).Scan(&exists)

	return exists
}

func (dr DriverRepository) FindByCompanyId(id int) (*[]d.IDriver, error) {
	drivers := []d.IDriver{}

	rows, err := dr.db.Query("SELECT * FROM drivers WHERE company_id = $1", id)
	if err != nil {
		return nil, err
	}

	for rows.Next() {

		driver := d.GetNewDriver()

		if err := rows.Scan(driver); err != nil {
			return nil, err
		}

		var castedDriver d.IDriver = driver

		drivers = append(drivers, castedDriver)
	}

	return &drivers, nil
}

func (dr DriverRepository) ExistsByCompanyId(id int) (exists *bool) {
	dr.db.QueryRow(
		"SELECT EXISTS (SELECT 1 FROM drivers WHERE company_id = $1)",
		id,
	).Scan(&exists)

	return exists
}

func (dr DriverRepository) Create(driver d.IDriver) error {
	if err := dr.simpleValidation(driver); err != nil {
		return err
	}

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
	if err := dr.simpleValidation(driver); err != nil {
		return err
	}

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
