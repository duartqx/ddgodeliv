package delivery

import (
	"time"

	d "ddgodeliv/domains/driver"
	u "ddgodeliv/domains/user"
)

type IDelivery interface {
	GetId() int
	SetId(id int) IDelivery

	GetLoadout() string
	SetLoadout(loadout string) IDelivery

	GetWeight() int
	SetWeight(weight int) IDelivery

	GetOrigin() string
	SetOrigin(origin string) IDelivery

	GetDestination() string
	SetDestination(destination string) IDelivery

	GetCreatedAt() time.Time
	SetCreatedAt(deadline time.Time) IDelivery

	GetDeadline() time.Time
	SetDeadline(deadline time.Time) IDelivery

	GetStatus() uint8
	GetStatusDisplay() string
	SetStatus(status uint8) IDelivery

	GetDriverId() int
	SetDriverId(driverId int) IDelivery
	SetDriver(driver d.IDriver) IDelivery

	GetSenderId() int
	SetSenderId(senderId int) IDelivery

	GetDriver() d.IDriver
	GetSender() u.IUser

	HasInvalidId() bool
	DriverIsNull() bool
}
