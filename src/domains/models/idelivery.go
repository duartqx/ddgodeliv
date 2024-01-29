package models

import "time"

type IDelivery interface {
	GetId() int
	SetId(id int) IDelivery

	GetOrigin() string
	SetOrigin(origin string) IDelivery

	GetDestination() string
	SetDestination(destination string) IDelivery

	GetDeadline() time.Time
	SetDeadline(deadline time.Time) IDelivery

	GetStatus() uint8
	GetStatusDisplay() string
	SetStatus(status uint8) IDelivery

	GetDriverId() int
	SetDriverId(driverId int) IDelivery

	GetSenderId() int
	SetSenderId(senderId int) IDelivery

	GetDriver() IDriver
	GetSender() IUser
}
