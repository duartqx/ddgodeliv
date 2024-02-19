package postgresql

import (
	"database/sql"
	"errors"
	"fmt"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"

	e "ddgodeliv/common/errors"
	d "ddgodeliv/domains/delivery"
)

type DeliveryRepository struct {
	db *sqlx.DB
}

func GetDeliveryRepository(db *sqlx.DB) *DeliveryRepository {
	return &DeliveryRepository{db: db}
}

func (dr DeliveryRepository) baseJoinedQuery(where string) string {
	return fmt.Sprintf(
		`
			SELECT
				de.id AS id,
				de.loadout AS loadout,
				de.weight AS weight,
				COALESCE(de.driver_id, 0) AS driver_id,
				de.sender_id AS sender_id,
				de.origin AS origin,
				de.destination AS destination,
				de.created_at AS created_at,
				de.deadline AS deadline,
				de.status AS status,

				su.id AS "sender.id",
				su.email AS "sender.email",
				su.name AS "sender.name",

				-- Coalesce will avoid having to use sql.NullString / sql.NullInt
				COALESCE(dr.id, 0) AS "driver.id",
				COALESCE(dr.user_id, 0) AS "driver.user_id",
				COALESCE(dr.company_id, 0) AS "driver.company_id",
				COALESCE(dr.license_id, '') AS "driver.license_id",

				COALESCE(du.id, 0) AS "driver.user.id",
				COALESCE(du.name, '') AS "driver.user.name",
				COALESCE(du.email, '') AS "driver.user.email",

				COALESCE(c.id, 0) AS "driver.company.id",
				COALESCE(c.owner_id, 0) AS "driver.company.owner_id",
				COALESCE(c.name, '') AS "driver.company.name"
			FROM deliveries de
			LEFT JOIN drivers dr ON de.driver_id = dr.id
			LEFT JOIN users du ON dr.user_id = du.id
			LEFT JOIN companies c ON dr.company_id = c.id
			LEFT JOIN users su ON de.sender_id = su.id
			WHERE  %s
		`,
		where,
	)
}

func (dr DeliveryRepository) FindById(delivery d.IDelivery) error {
	if err := dr.db.Get(
		delivery, dr.baseJoinedQuery("de.id = $1"), delivery.GetId(),
	); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return e.NotFoundError
		}
		return err
	}
	return nil
}

func (dr DeliveryRepository) findMany(where string, args ...interface{}) (*[]d.IDelivery, error) {
	deliveries := []d.IDelivery{}

	rows, err := dr.db.Queryx(dr.baseJoinedQuery(where), args...)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		delivery := d.GetNewDelivery()

		if err := rows.StructScan(delivery); err != nil {
			return nil, err
		}

		var castedDelivery d.IDelivery = delivery

		deliveries = append(deliveries, castedDelivery)
	}

	return &deliveries, nil

}

func (dr DeliveryRepository) FindByDriverId(id int) (*[]d.IDelivery, error) {
	return dr.findMany("de.driver_id = $1", id)
}

func (dr DeliveryRepository) FindCurrentByDriverId(id int) (d.IDelivery, error) {

	var delivery d.Delivery

	if err := dr.db.Get(
		&delivery,
		dr.baseJoinedQuery("de.driver_id = $1 AND de.status != 0"),
		id,
	); err != nil {
		if !errors.Is(err, sql.ErrNoRows) {
			return nil, err
		}
	}
	return &delivery, nil
}

func (dr DeliveryRepository) ExistsByDriverId(id int) (exists bool) {
	dr.db.QueryRow(
		"SELECT EXISTS (SELECT 1 FROM deliveries WHERE driver_id = $1)",
		id,
	).Scan(&exists)

	return exists
}

func (dr DeliveryRepository) FindByStatusByDriverId(
	id int, status uint8,
) (*[]d.IDelivery, error) {
	return dr.findMany("de.driver_id = $1 AND de.status = $2", id, status)
}

func (dr DeliveryRepository) ExistsByStatusByDriverId(
	id int, status uint8,
) (exists bool) {
	dr.db.QueryRow(
		`SELECT EXISTS (
			SELECT 1 FROM deliveries WHERE driver_id = $1 AND status = $2
		)`,
		id, status,
	).Scan(&exists)

	return exists
}

func (dr DeliveryRepository) FindByDeadlineDateRange(
	start, end time.Time,
) (*[]d.IDelivery, error) {
	return dr.findMany("de.deadline BETWEEN $1 AND $2", start, end)
}

func (dr DeliveryRepository) FindByCreatedAtDate(
	createdAt time.Time,
) (*[]d.IDelivery, error) {
	return dr.findMany("de.created_at::date = $1", createdAt.Format("2006-01-02"))
}

func (dr DeliveryRepository) FindByDeadlineDate(
	deadline time.Time,
) (*[]d.IDelivery, error) {
	return dr.findMany("de.deadline::date = $1", deadline.Format("2006-01-02"))
}

func (dr DeliveryRepository) ExistsByDeadlineDate(
	deadline time.Time,
) (exists bool) {

	dr.db.QueryRow(
		"SELECT EXISTS (SELECT 1 FROM deliveries WHERE deadline::date = $1)",
		deadline.Format("2006-01-02"),
	).Scan(&exists)

	return exists
}

func (dr DeliveryRepository) FindBySenderId(id int) (*[]d.IDelivery, error) {
	return dr.findMany("de.sender_id = $1", id)
}

func (dr DeliveryRepository) ExistsBySenderId(id int) (exists bool) {

	dr.db.QueryRow(
		"SELECT EXISTS (SELECT 1 FROM deliveries WHERE sender_id = $1)", id,
	).Scan(&exists)

	return exists
}

func (dr DeliveryRepository) FindByCompanyId(id int) (*[]d.IDelivery, error) {
	return dr.findMany("dr.company_id = $1", id)
}

func (dr DeliveryRepository) ExistsByCompanyId(id int) (exists bool) {

	dr.db.QueryRow(
		"SELECT EXISTS (SELECT 1 FROM deliveries WHERE company_id = $1)", id,
	).Scan(&exists)

	return exists
}

func (dr DeliveryRepository) Create(delivery d.IDelivery) error {
	if err := dr.db.Get(
		delivery,
		`
			WITH new_delivery AS (
				INSERT INTO deliveries (
					sender_id,
					loadout,
					weight,
					origin,
					destination,
					deadline,
					status
				)
				VALUES ($1, $2, $3, $4, $5, $6, $7)
				RETURNING id, created_at, sender_id
			)
			SELECT
				d.id AS "id",
				d.created_at AS "created_at",
				u.id AS "sender.id",
				u.name AS "sender.name",
				u.email AS "sender.email"
			FROM new_delivery d
			INNER JOIN users u ON u.id = d.sender_id
		`,
		delivery.GetSenderId(),
		delivery.GetLoadout(),
		delivery.GetWeight(),
		delivery.GetOrigin(),
		delivery.GetDestination(),
		delivery.GetDeadline(),
		delivery.GetStatus(),
	); err != nil {
		if pqErr, ok := err.(*pq.Error); ok && pqErr.Code == "23503" {
			// "foreign_key_violation"
			return fmt.Errorf(
				"Foreign Key Constrainst Violation: %w",
				e.BadRequestError,
			)
		}
		return err
	}

	return nil
}

func (dr DeliveryRepository) baseUpdateJoinedQuery(setWhere string) string {
	return fmt.Sprintf(
		`
			WITH updated AS (
				UPDATE deliveries
		        %s
				RETURNING id, driver_id
			)
			SELECT
				de.id AS "id",
				-- Coalesce will avoid having to use sql.NullString / sql.NullInt
				COALESCE(dr.id, 0) AS "driver.id",
				COALESCE(dr.user_id, 0) AS "driver.user_id",
				COALESCE(dr.company_id, 0) AS "driver.company_id",
				COALESCE(dr.license_id, '') AS "driver.license_id",

				COALESCE(du.id, 0) AS "driver.user.id",
				COALESCE(du.name, '') AS "driver.user.name",
				COALESCE(du.email, '') AS "driver.user.email",

				COALESCE(c.id, 0) AS "driver.company.id",
				COALESCE(c.owner_id, 0) AS "driver.company.owner_id",
				COALESCE(c.name, '') AS "driver.company.name"

				FROM updated de
				LEFT JOIN drivers dr ON de.driver_id = dr.id
				LEFT JOIN users du ON dr.user_id = du.id
				LEFT JOIN companies c ON dr.company_id = c.id
		`,
		setWhere,
	)
}

func (dr DeliveryRepository) AssignDriver(delivery d.IDelivery) error {
	if err := dr.db.Get(
		delivery,
		dr.baseUpdateJoinedQuery(
			`SET driver_id = $1, status = $2 WHERE id = $3`,
		),
		delivery.GetDriverId(), delivery.GetStatus(), delivery.GetId(),
	); err != nil {
		return fmt.Errorf("%w: %v", e.InternalError, err.Error())
	}
	return nil
}

func (dr DeliveryRepository) UpdateStatus(delivery d.IDelivery) error {
	if err := dr.db.Get(
		delivery,
		dr.baseUpdateJoinedQuery(`SET status = $1 WHERE id = $2`),
		delivery.GetStatus(), delivery.GetId(),
	); err != nil {
		return fmt.Errorf("%w: %v", e.InternalError, err.Error())
	}
	return nil
}

func (dr DeliveryRepository) Delete(delivery d.IDelivery) error {
	_, err := dr.db.NamedExec("DELETE FROM deliveries WHERE id = :id", delivery)

	return err
}

func (dr DeliveryRepository) FindPendingWithNoDriver() (*[]d.IDelivery, error) {
	return dr.findMany(
		"de.status = $1 AND de.driver_id IS NULL", d.StatusChoices.Pending,
	)
}
