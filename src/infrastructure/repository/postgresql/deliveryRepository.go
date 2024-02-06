package postgresql

import (
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

func GetNewDeliveryRepository(db *sqlx.DB) *DeliveryRepository {
	return &DeliveryRepository{db: db}
}

func (dr DeliveryRepository) FindById(delivery d.IDelivery) error {
	if err := dr.db.Get(
		delivery, "SELECT * FROM deliveries WHERE id = $1", delivery.GetId(),
	); err != nil {
		return err
	}
	return nil
}

func (dr DeliveryRepository) findMany(where string, args ...interface{}) (*[]d.IDelivery, error) {
	deliveries := []d.IDelivery{}

	query := fmt.Sprintf(`
		SELECT
			de.id AS id,
			de.loadout AS loadout,
			de.weight AS weight,
			de.driver_id AS driver_id,
			de.sender_id AS sender_id,
			de.origin AS origin,
			de.destination AS destination,
			de.created_at AS created_at,
			de.deadline AS deadline,
			de.status AS status,

			dr.id AS "driver.id",
			dr.user_id AS "driver.user_id",
			dr.company_id AS "driver.company_id",
			dr.license_id AS "driver.license_id",

			u.id AS "driver.user.id"
			u.name AS "driver.user.name"
			u.email AS "driver.user.email"

			c.id AS "company.id",
			c.owner_id AS "company.owner_id",
			c.name AS "company.name"
		FROM deliveries de
		INNER JOIN companies c ON de.company_id = c.id
		INNER JOIN drivers dr ON de.driver_id = dr.id
		INNER JOIN users u ON dr.user_id = u.id
		WHERE %s
	`, where)

	rows, err := dr.db.Queryx(query, args...)
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

func (dr DeliveryRepository) ExistsByDriverId(id int) (exists bool) {
	dr.db.QueryRow(
		"SELECT EXISTS (SELECT 1 FROM deliveries WHERE driver_id = $1)",
		id,
	).Scan(&exists)

	return exists
}

func (dr DeliveryRepository) FindByStatusByDriverId(id int, status uint8) (*[]d.IDelivery, error) {
	return dr.findMany("driver_id = $1 AND status = $2", id, status)
}

func (dr DeliveryRepository) ExistsByStatusByDriverId(id int, status uint8) (exists bool) {
	dr.db.QueryRow(
		"SELECT EXISTS (SELECT 1 FROM deliveries WHERE driver_id = $1 AND status = $2)",
		id, status,
	).Scan(&exists)

	return exists
}

func (dr DeliveryRepository) FindByDeadlineDateRange(start, end time.Time) (*[]d.IDelivery, error) {
	return dr.findMany("de.deadline BETWEEN $1 AND $2", start, end)
}

func (dr DeliveryRepository) FindByCreatedAtDate(createdAt time.Time) (*[]d.IDelivery, error) {
	return dr.findMany("de.created_at::date = $1", createdAt.Format("2006-01-02"))
}

func (dr DeliveryRepository) FindByDeadlineDate(deadline time.Time) (*[]d.IDelivery, error) {
	return dr.findMany("de.deadline::date = $1", deadline.Format("2006-01-02"))
}

func (dr DeliveryRepository) ExistsByDeadlineDate(deadline time.Time) (exists bool) {

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

func (dr DeliveryRepository) getDriverIdToQuery(id int) string {
	var deliveryDriverId string
	if id == 0 {
		deliveryDriverId = "NULL"
	} else {
		deliveryDriverId = fmt.Sprint(id)
	}
	return deliveryDriverId
}

func (dr DeliveryRepository) Create(delivery d.IDelivery) error {
	var id int

	if err := dr.db.QueryRow(
		`
			INSERT INTO deliveries (
				loadout,
				weight,
				driver_id,
				sender_id,
				origin,
				destination,
				deadline,
				status
			)
			VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
			RETURNING id
		`,
		dr.getDriverIdToQuery(delivery.GetDriverId()),
		delivery.GetLoadout(),
		delivery.GetWeight(),
		delivery.GetSenderId(),
		delivery.GetOrigin(),
		delivery.GetDestination(),
		delivery.GetDeadline(),
		delivery.GetStatus(),
	).Scan(&id); err != nil {
		if pqErr, ok := err.(*pq.Error); ok && pqErr.Code == "23503" {
			// "foreign_key_violation"
			return fmt.Errorf(
				"Foreign Key Constrainst Violation: %w",
				e.BadRequestError,
			)
		}
		return err
	}

	delivery.SetId(id)

	return nil
}

func (dr DeliveryRepository) Update(delivery d.IDelivery) error {
	_, err := dr.db.Exec(
		`
			UPDATE deliveries
			SET
				driver_id = $1,
				sender_id = $2,
				origin = $3,
				destination = $4,
				deadline = $5,
				status = $6,
				loadout = $7,
				wegith = $8
			WHERE id = $9
		`,
		dr.getDriverIdToQuery(delivery.GetDriverId()),
		delivery.GetSenderId(),
		delivery.GetOrigin(),
		delivery.GetDestination(),
		delivery.GetDeadline(),
		delivery.GetStatus(),
		delivery.GetLoadout(),
		delivery.GetWeight(),
		delivery.GetId(),
	)
	return err
}

func (dr DeliveryRepository) Delete(delivery d.IDelivery) error {
	_, err := dr.db.Exec("DELETE FROM deliveries WHERE id = $1", delivery.GetId())

	return err
}

func (dr DeliveryRepository) FindPendingWithNoDriver() (*[]d.IDelivery, error) {
	return dr.findMany("de.status = $1 AND de.driver_id = NULL", d.StatusChoices.Pending)
}
