package postgresql

import (
	"fmt"
	"time"

	"github.com/jmoiron/sqlx"

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

func (dr DeliveryRepository) findMany(query string, args ...interface{}) (*[]d.IDelivery, error) {
	deliveries := []d.IDelivery{}

	rows, err := dr.db.Query(query, args...)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		delivery := d.GetNewDelivery()

		if err := rows.Scan(delivery); err != nil {
			return nil, err
		}

		var castedDelivery d.IDelivery = delivery

		deliveries = append(deliveries, castedDelivery)
	}

	return &deliveries, nil

}

func (dr DeliveryRepository) FindByDriverId(id int) (*[]d.IDelivery, error) {
	return dr.findMany("SELECT * FROM deliveries WHERE driver_id = $1", id)
}

func (dr DeliveryRepository) ExistsByDriverId(id int) (exists *bool) {
	dr.db.QueryRow(
		"SELECT EXISTS (SELECT 1 FROM deliveries WHERE driver_id = $1)",
		id,
	).Scan(&exists)

	return exists
}

func (dr DeliveryRepository) FindByStatusByDriverId(id int, status uint8) (*[]d.IDelivery, error) {
	return dr.findMany(
		"SELECT * FROM deliveries WHERE driver_id = $1 AND status = $2",
		id, status,
	)
}

func (dr DeliveryRepository) ExistsByStatusByDriverId(id int, status uint8) (exists *bool) {
	dr.db.QueryRow(
		"SELECT EXISTS (SELECT 1 FROM deliveries WHERE driver_id = $1 AND status = $2)",
		id, status,
	).Scan(&exists)

	return exists
}

func (dr DeliveryRepository) FindByDeadlineDateRange(start, end time.Time) (*[]d.IDelivery, error) {
	return dr.findMany(
		"SELECT * FROM deliveries WHERE deadline BETWEEN $1 AND $2", start, end,
	)
}

func (dr DeliveryRepository) FindByDeadlineDate(deadline time.Time) (*[]d.IDelivery, error) {
	return dr.findMany(
		"SELECT * FROM deliveries WHERE deadline::date = $1",
		deadline.Format("2006-01-02"),
	)
}

func (dr DeliveryRepository) ExistsByDeadlineDate(deadline time.Time) (exists *bool) {

	dr.db.QueryRow(
		"SELECT EXISTS (SELECT 1 FROM deliveries WHERE deadline::date = $1)",
		deadline.Format("2006-01-02"),
	).Scan(&exists)

	return exists
}

func (dr DeliveryRepository) FindBySenderId(id int) (*[]d.IDelivery, error) {
	return dr.findMany("SELECT * FROM deliveries WHERE sender_id = $1", id)
}

func (dr DeliveryRepository) ExistsBySenderId(id int) (exists *bool) {

	dr.db.QueryRow(
		"SELECT EXISTS (SELECT 1 FROM deliveries WHERE sender_id = $1)", id,
	).Scan(&exists)

	return exists
}

func (dr DeliveryRepository) FindByCompanyId(id int) (*[]d.IDelivery, error) {
	return dr.findMany("SELECT * FROM deliveries WHERE company_id = $1", id)
}

func (dr DeliveryRepository) ExistsByCompanyId(id int) (exists *bool) {

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
			INSERT INTO deliveries (driver_id, sender_id, origin, destination, deadline, status)
			VALUES ($1, $2, $3, $4, $5, $6)
			RETURNING id
		`,
		dr.getDriverIdToQuery(delivery.GetDriverId()),
		delivery.GetSenderId(),
		delivery.GetOrigin(),
		delivery.GetDestination(),
		delivery.GetDeadline(),
		delivery.GetStatus(),
	).Scan(&id); err != nil {
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
				status = $6
			WHERE id = $7
		`,
		dr.getDriverIdToQuery(delivery.GetDriverId()),
		delivery.GetSenderId(),
		delivery.GetOrigin(),
		delivery.GetDestination(),
		delivery.GetDeadline(),
		delivery.GetStatus(),
		delivery.GetId(),
	)
	return err
}

func (dr DeliveryRepository) Delete(delivery d.IDelivery) error {
	_, err := dr.db.Exec("DELETE FROM deliveries WHERE id = $1", delivery.GetId())

	return err
}

func (dr DeliveryRepository) FindPendingWithNoDriver() (*[]d.IDelivery, error) {
	return dr.findMany(
		"SELECT * FROM deliveries WHERE status = $1 AND driver_id = NULL",
		d.StatusChoices.Pending,
	)
}
