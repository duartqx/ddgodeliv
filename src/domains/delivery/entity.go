package delivery

import (
	"ddgodeliv/src/domains/driver"
	"ddgodeliv/src/domains/user"
	"time"
)

type IDelivery interface {
	GetId() int
	SetId(id int) IDelivery

	GetDestination() string
	SetDestination(destination string) IDelivery

	GetDeadline() time.Time
	SetDeadline(deadline time.Time) IDelivery

	GetDriverId() int
	SetDriverId(driverId int) IDelivery

	GetSenderId() int
	SetSenderId(senderId int) IDelivery

	GetDriver() driver.IDriver
	GetSender() user.IUser
}

type Delivery struct {
	Id          int       `db:"id" json:"id"`
	Destination string    `db:"destination" json:"destination"`
	Deadline    time.Time `db:"deadline" json:"deadline"`
	DriverId    int       `db:"driver_id" json:"driver_id"`
	SenderId    int       `db:"sender_id" json:"sender_id"`

	Driver driver.IDriver `json:"driver"`
	Sender user.IUser     `json:"sender"`
}

func (d Delivery) GetId() int {
	return d.Id
}

func (d *Delivery) SetId(id int) IDelivery {
	d.Id = id
	return d
}

func (d Delivery) GetDestination() string {
	return d.Destination
}

func (d *Delivery) SetDestination(destination string) IDelivery {
	d.Destination = destination
	return d
}

func (d Delivery) GetDeadline() time.Time {
	return d.Deadline
}

func (d *Delivery) SetDeadline(deadline time.Time) IDelivery {
	d.Deadline = deadline
	return d
}

func (d Delivery) GetDriverId() int {
	return d.DriverId
}

func (d *Delivery) SetDriverId(driverId int) IDelivery {
	d.DriverId = driverId
	return d
}

func (d Delivery) GetSenderId() int {
	return d.SenderId
}

func (d *Delivery) SetSenderId(senderId int) IDelivery {
	d.SenderId = senderId
	return d
}

func (d Delivery) GetDriver() driver.IDriver {
	return d.Driver
}

func (d Delivery) GetSender() user.IUser {
	return d.Sender
}
