package delivery

import (
	"time"

	"ddgodeliv/domains/driver"
	"ddgodeliv/domains/user"
)

type IDelivery interface {
	GetId() int
	SetId(id int) IDelivery

	GetDestination() string
	SetDestination(destination string) IDelivery

	GetDeadline() time.Time
	SetDeadline(deadline time.Time) IDelivery

	GetStatus() uint8
	SetStatus(status uint8) IDelivery

	GetDriverId() int
	SetDriverId(driverId int) IDelivery

	GetSenderId() int
	SetSenderId(senderId int) IDelivery

	GetDriver() driver.IDriver
	GetSender() user.IUser
}

type status struct {
	Value uint8
	Label string
}

type deliveryStatus struct {
	Pending   *status
	InTransit *status
	Late      *status
	Completed *status
}

var DeliveryStatusChoices *deliveryStatus = &deliveryStatus{
	Pending:   &status{Value: 0, Label: "Pending"},
	InTransit: &status{Value: 1, Label: "InTransit"},
	Late:      &status{Value: 2, Label: "Late"},
	Completed: &status{Value: 3, Label: "Completed"},
}

type Delivery struct {
	Id          int       `db:"id" json:"id"`
	DriverId    int       `db:"driver_id" json:"driver_id"`
	SenderId    int       `db:"sender_id" json:"sender_id"`
	Destination string    `db:"destination" json:"destination"`
	Deadline    time.Time `db:"deadline" json:"deadline"`
	Status      uint8     `db:"status" json:"status"`

	Driver driver.IDriver `json:"driver"`
	Sender user.IUser     `json:"sender"`
}

func GetNewDelivery() *Delivery {
	return &Delivery{
		Driver: driver.GetNewDriver(),
		Sender: user.GetNewUser(),
	}
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

func (d Delivery) GetStatus() uint8 {
	return d.Status
}

func (d *Delivery) SetStatus(status uint8) IDelivery {
	d.Status = status
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
