package postgresql

import (
	"fmt"
	"time"

	"github.com/jmoiron/sqlx"

	de "ddgodeliv/domains/delivery"
)

type DeliveryRepository struct {
	db *sqlx.DB
}

func GetNewDeliveryRepository(db *sqlx.DB) *DeliveryRepository {
	return &DeliveryRepository{db: db}
}

func (dr DeliveryRepository) FindById(id int) (de.IDelivery, error) {
	delivery := de.GetNewDelivery()

	if err := dr.db.Get(delivery, "SELECT * FROM deliveries WHERE id = $1", id); err != nil {
		return nil, err
	}

	return delivery, nil
}

func (dr DeliveryRepository) FindByDriverId(id int) (*[]de.IDelivery, error) {
	deliveries := []de.IDelivery{}

	rows, err := dr.db.Query("SELECT * FROM deliveries WHERE driver_id = $1", id)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		delivery := de.GetNewDelivery()

		if err := rows.Scan(delivery); err != nil {
			return nil, err
		}

		var castedDelivery de.IDelivery = delivery

		deliveries = append(deliveries, castedDelivery)
	}

	return &deliveries, nil
}

func (dr DeliveryRepository) ExistsByDriverId(id int) (exists *bool) {
	dr.db.QueryRow(
		"SELECT EXISTS (SELECT 1 FROM deliveries WHERE driver_id = $1)",
		id,
	).Scan(&exists)

	return exists
}

func (dr DeliveryRepository) FindByStatusByDriverId(id int, status uint8) (*[]de.IDelivery, error) {
	deliveries := []de.IDelivery{}

	rows, err := dr.db.Query(
		"SELECT * FROM deliveries WHERE driver_id = $1 AND status = $2",
		id, status,
	)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		delivery := de.GetNewDelivery()

		if err := rows.Scan(delivery); err != nil {
			return nil, err
		}

		var castedDelivery de.IDelivery = delivery

		deliveries = append(deliveries, castedDelivery)
	}

	return &deliveries, nil
}

func (dr DeliveryRepository) ExistsByCompletionByDriverId(id int, status uint8) (exists *bool) {
	dr.db.QueryRow(
		"SELECT EXISTS (SELECT 1 FROM deliveries WHERE driver_id = $1 AND status = $2)",
		id, status,
	).Scan(&exists)

	return exists
}

func (dr DeliveryRepository) FindByDeadlineDateRange(start, end time.Time) (*[]de.IDelivery, error) {
	deliveries := []de.IDelivery{}

	rows, err := dr.db.Query(
		"SELECT * FROM deliveries WHERE deadline BETWEEN $1 AND $2", start, end,
	)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		delivery := de.GetNewDelivery()

		if err := rows.Scan(delivery); err != nil {
			return nil, err
		}

		var castedDelivery de.IDelivery = delivery

		deliveries = append(deliveries, castedDelivery)
	}

	return &deliveries, nil
}

func (dr DeliveryRepository) FindByDeadlineDate(deadline time.Time) (*[]de.IDelivery, error) {

	deliveries := []de.IDelivery{}

	rows, err := dr.db.Query(
		"SELECT * FROM deliveries WHERE deadline::date = $1",
		deadline.Format("2006-01-02"),
	)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		delivery := de.GetNewDelivery()

		if err := rows.Scan(delivery); err != nil {
			return nil, err
		}

		var castedDelivery de.IDelivery = delivery

		deliveries = append(deliveries, castedDelivery)
	}

	return &deliveries, nil
}

func (dr DeliveryRepository) ExistsByDeadlineDate(deadline time.Time) (exists *bool) {

	dr.db.QueryRow(
		"SELECT EXISTS (SELECT 1 FROM deliveries WHERE deadline::date = $1)",
		deadline.Format("2006-01-02"),
	).Scan(&exists)

	return exists
}

func (dr DeliveryRepository) FindBySenderId(id int) (*[]de.IDelivery, error) {

	deliveries := []de.IDelivery{}

	rows, err := dr.db.Query("SELECT * FROM deliveries WHERE sender_id = $1", id)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		delivery := de.GetNewDelivery()

		if err := rows.Scan(delivery); err != nil {
			return nil, err
		}

		var castedDelivery de.IDelivery = delivery

		deliveries = append(deliveries, castedDelivery)
	}

	return &deliveries, nil
}

func (dr DeliveryRepository) ExistsBySenderId(id int) (exists *bool) {

	dr.db.QueryRow(
		"SELECT EXISTS (SELECT 1 FROM deliveries WHERE sender_id = $1)", id,
	).Scan(&exists)

	return exists
}

func (dr DeliveryRepository) FindByCompanyId(id int) (*[]de.IDelivery, error) {

	deliveries := []de.IDelivery{}

	rows, err := dr.db.Query("SELECT * FROM deliveries WHERE company_id = $1", id)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		delivery := de.GetNewDelivery()

		if err := rows.Scan(delivery); err != nil {
			return nil, err
		}

		var castedDelivery de.IDelivery = delivery

		deliveries = append(deliveries, castedDelivery)
	}

	return &deliveries, nil
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

func (dr DeliveryRepository) Create(delivery de.IDelivery) error {
	if delivery.GetSenderId() == 0 || delivery.GetDestination() == "" {
		return fmt.Errorf("Invalid Delivery: Missing Sender or Destination")
	}

	var id int

	if err := dr.db.QueryRow(
		`
			INSERT INTO deliveries (driver_id, sender_id, destination, deadline)
			VALUES ($1, $2, $3, $4)
			RETURNING id
		`,
		dr.getDriverIdToQuery(delivery.GetDriverId()),
		delivery.GetSenderId(),
		delivery.GetDestination(),
		delivery.GetDeadline(),
	).Scan(&id); err != nil {
		return err
	}

	delivery.SetId(id)

	return nil
}

func (dr DeliveryRepository) Update(delivery de.IDelivery) error {

	if delivery.GetSenderId() == 0 || delivery.GetDestination() == "" {
		return fmt.Errorf("Invalid Delivery: Missing Sender or Destination")
	}

	_, err := dr.db.Exec(
		`
			UPDATE deliveries
			SET driver_id = $1, sender_id = $2, destination = $3, deadline = $4
			WHERE id = $5
		`,
		dr.getDriverIdToQuery(delivery.GetDriverId()),
		delivery.GetSenderId(),
		delivery.GetDestination(),
		delivery.GetDeadline(),
	)
	return err
}

func (dr DeliveryRepository) Delete(delivery de.IDelivery) error {
	if delivery.GetId() == 0 {
		return fmt.Errorf("Invalid Delivery Id")
	}

	_, err := dr.db.Exec("DELETE FROM deliveries WHERE id = $1", delivery.GetId())

	return err
}
