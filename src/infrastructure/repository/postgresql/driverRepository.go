package postgresql

import (
	"database/sql"
	"errors"
	"fmt"
	"slices"
	"strings"

	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"

	e "ddgodeliv/common/errors"
	d "ddgodeliv/domains/driver"
)

type DriverRepository struct {
	db *sqlx.DB
}

var driverRepository *DriverRepository

func GetDriverRepository(db *sqlx.DB) *DriverRepository {
	if driverRepository == nil {
		driverRepository = &DriverRepository{db: db}
	}
	return driverRepository
}

func (dr DriverRepository) baseJoinedQuery(where string) string {
	return fmt.Sprintf(
		`
			SELECT
				d.id AS "id",
				d.user_id AS "user_id",
				d.company_id AS "company_id",
				d.license_id AS "license_id",
				COALESCE(de.status, 0) AS "status",

				u.id AS "user.id",
				u.name AS "user.name",
				u.email AS "user.email",
				
				c.id AS "company.id",
				c.owner_id AS "company.owner_id",
				c.name AS "company.name"
			FROM drivers d
			INNER JOIN users u ON d.user_id = u.id
			INNER JOIN companies c ON d.company_id = c.id
			LEFT JOIN LATERAL (
				SELECT status, driver_id
				FROM deliveries
				WHERE driver_id = d.id AND status != 0
				ORDER BY created_at DESC LIMIT 1
			) de ON d.id = de.driver_id
			WHERE %s
		`,
		where,
	)
}

func (dr DriverRepository) FindById(driver d.IDriver) error {

	if err := dr.db.Get(
		driver, dr.baseJoinedQuery("d.id = $1 AND d.company_id = $2"),
		driver.GetId(), driver.GetCompanyId(),
	); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return e.NotFoundError
		}
		return err
	}

	return nil
}

func (dr DriverRepository) FindByUserId(id int) (d.IDriver, error) {
	driver := d.GetNewDriver()

	if err := dr.db.Get(driver, dr.baseJoinedQuery("d.user_id = $1"), id); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, e.NotFoundError
		}
		return nil, err
	}

	return driver, nil
}

func (dr DriverRepository) ExistsById(id, companyId int) (exists bool) {
	dr.db.QueryRow(
		`
			SELECT EXISTS (
				SELECT 1 FROM drivers
				WHERE id = $1 AND company_id = $2
			)
		`,
		id, companyId,
	).Scan(&exists)

	return exists
}

func (dr DriverRepository) ExistsByUserId(id int) (exists bool) {
	dr.db.QueryRow(
		`SELECT EXISTS (SELECT 1 FROM drivers WHERE user_id = $1)`, id,
	).Scan(&exists)

	return exists
}

func (dr DriverRepository) FindByCompanyId(id int) (*[]d.IDriver, error) {
	drivers := []d.IDriver{}

	rows, err := dr.db.Queryx(dr.baseJoinedQuery("d.company_id = $1"), id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, e.NotFoundError
		}
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
	if err := dr.db.Get(
		driver,
		// Inserts new driver and grabs the company information
		`
			WITH new_driver AS (
				INSERT INTO drivers (user_id, company_id, license_id)
				VALUES ($1, $2, $3)
				RETURNING id, company_id
			)
			SELECT
				d.id as "id",
				c.id as "company.id",
				c.name as "company.name",
				c.owner_id as "company.owner_id"
			FROM new_driver d
			INNER JOIN companies c ON c.id = d.company_id
		`,
		driver.GetUserId(),
		driver.GetCompanyId(),
		strings.ToLower(driver.GetLicenseId()),
	); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return e.NotFoundError
		}
		if pqErr, ok := err.(*pq.Error); ok {
			// 23503 => foreign_key_violation
			// 23505 => unique_violation
			if slices.Contains[[]pq.ErrorCode](
				[]pq.ErrorCode{"23503", "23505"}, pqErr.Code,
			) {
				return e.BadRequestError
			}
		}
		return err
	}

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
	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok && pqErr.Code == "23505" {
			return fmt.Errorf("%w: Invalid License Id", e.BadRequestError)
		}
	}
	return err
}

func (dr DriverRepository) Delete(driver d.IDriver) error {
	if driver.GetId() == 0 {
		return fmt.Errorf("Invalid Driver Id")
	}

	_, err := dr.db.Exec(
		`DELETE FROM users WHERE id IN (
			SELECT user_id FROM drivers WHERE id = $1
		)`,
		driver.GetId(),
	)

	return err
}

func (dr DriverRepository) ExistsByLicenseId(licenseId string) (exists bool) {
	dr.db.QueryRow(
		`SELECT EXISTS (SELECT 1 FROM drivers WHERE license_id = $1)`,
		licenseId,
	).Scan(&exists)

	return exists
}
