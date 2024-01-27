package repository

import (
	"time"

	de "ddgodeliv/domains/delivery"
)

type IDeliveryRepository interface {
	FindById(id int) (de.IDelivery, error)

	FindByDriverId(id int) (*[]de.IDelivery, error)
	ExistsByDriverId(id int) *bool

	FindByCompletionByDriverId(id int, completed bool) (*[]de.IDelivery, error)
	ExistsByCompletionByDriverId(id int, completed bool) *bool

	FindByDeadlineDateRange(start, end time.Time) (*[]de.IDelivery, error)
	FindByDeadlineDate(deadline time.Time) (*[]de.IDelivery, error)
	ExistsByDeadlineDate(deadline time.Time) *bool

	FindBySenderId(id int) (*[]de.IDelivery, error)
	ExistsBySenderId(id int) *bool

	FindByCompanyId(id int) (*[]de.IDelivery, error)
	ExistsByCompanyId(id int) *bool

	Create(delivery de.IDelivery) error
	Update(delivery de.IDelivery) error
	Delete(delivery de.IDelivery) error
}
