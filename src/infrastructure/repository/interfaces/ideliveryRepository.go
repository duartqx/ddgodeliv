package repository

import (
	"time"

	m "ddgodeliv/domains/models"
)

type IDeliveryRepository interface {
	FindById(id int) (m.IDelivery, error)

	FindByDriverId(id int) (*[]m.IDelivery, error)
	ExistsByDriverId(id int) *bool

	FindByStatusByDriverId(id int, status uint8) (*[]m.IDelivery, error)
	ExistsByStatusByDriverId(id int, status uint8) *bool

	FindByDeadlineDateRange(start, end time.Time) (*[]m.IDelivery, error)
	FindByDeadlineDate(deadline time.Time) (*[]m.IDelivery, error)
	ExistsByDeadlineDate(deadline time.Time) *bool

	FindBySenderId(id int) (*[]m.IDelivery, error)
	ExistsBySenderId(id int) *bool

	FindByCompanyId(id int) (*[]m.IDelivery, error)
	ExistsByCompanyId(id int) *bool

	Create(delivery m.IDelivery) error
	Update(delivery m.IDelivery) error
	Delete(delivery m.IDelivery) error
}
