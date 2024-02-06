package delivery

import "time"

type IDeliveryRepository interface {
	FindById(delivery IDelivery) error

	FindPendingWithNoDriver() (*[]IDelivery, error)

	FindByDriverId(id int) (*[]IDelivery, error)
	ExistsByDriverId(id int) bool

	FindByStatusByDriverId(id int, status uint8) (*[]IDelivery, error)
	ExistsByStatusByDriverId(id int, status uint8) bool

	FindByCreatedAtDate(createdAt time.Time) (*[]IDelivery, error)

	FindByDeadlineDateRange(start, end time.Time) (*[]IDelivery, error)
	FindByDeadlineDate(deadline time.Time) (*[]IDelivery, error)
	ExistsByDeadlineDate(deadline time.Time) bool

	FindBySenderId(id int) (*[]IDelivery, error)
	ExistsBySenderId(id int) bool

	FindByCompanyId(id int) (*[]IDelivery, error)
	ExistsByCompanyId(id int) bool

	Create(delivery IDelivery) error
	Update(delivery IDelivery) error
	Delete(delivery IDelivery) error
}
