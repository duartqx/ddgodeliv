package repository

import (
	"time"

	de "ddgodeliv/src/domains/delivery"
)

type IDeliveryRepository interface {
	FindById(id int) (de.IDelivery, error)

	FindByDriverId(id int) (*[]de.IDelivery, error)
	ExistsByDriverId(id int) *bool

	FindByDeadlineDateRange(start, end time.Time) (*[]de.IDelivery, error)
	FindByDeadlineDate(deadline time.Time) (*[]de.IDelivery, error)
	ExistsByDeadlineDate(deadline time.Time) *bool

	FindBySenderId(id int) (*[]de.IDelivery, error)
	ExistsBySenderId(id int) *bool

	FindByCompanyId(id int) (*[]de.IDelivery, error)
	ExistsByCompanyId(id int) *bool

	Create(freight de.IDelivery) error
	Update(freight de.IDelivery) error
	Delete(freight de.IDelivery) error
}
