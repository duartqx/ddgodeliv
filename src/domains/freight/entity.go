package freight

import (
	"ddgodeliv/src/domains/driver"
	"ddgodeliv/src/domains/user"
	"time"
)

type IFreight interface {
	GetId() int
	SetId(id int) IFreight

	GetDestination() string
	SetDestination(destination string) IFreight

	GetDeadline() time.Time
	SetDeadline(deadline time.Time) IFreight

	GetDriverId() int
	SetDriverId(driverId int) IFreight

	GetSenderId() int
	SetSenderId(senderId int) IFreight

	GetDriver() driver.IDriver
	GetSender() user.IUser
}

type Freight struct {
	Id          int       `db:"id" json:"id"`
	Destination string    `db:"destination" json:"destination"`
	Deadline    time.Time `db:"deadline" json:"deadline"`
	DriverId    int       `db:"driver_id" json:"driver_id"`
	SenderId    int       `db:"sender_id" json:"sender_id"`

	Driver driver.IDriver `json:"driver"`
	Sender user.IUser     `json:"sender"`
}

func (f Freight) GetId() int {
	return f.Id
}

func (f *Freight) SetId(id int) IFreight {
	f.Id = id
	return f
}

func (f Freight) GetDestination() string {
	return f.Destination
}

func (f *Freight) SetDestination(destination string) IFreight {
	f.Destination = destination
	return f
}

func (f Freight) GetDeadline() time.Time {
	return f.Deadline
}

func (f *Freight) SetDeadline(deadline time.Time) IFreight {
	f.Deadline = deadline
	return f
}

func (f Freight) GetDriverId() int {
	return f.DriverId
}

func (f *Freight) SetDriverId(driverId int) IFreight {
	f.DriverId = driverId
	return f
}

func (f Freight) GetSenderId() int {
	return f.SenderId
}

func (f *Freight) SetSenderId(senderId int) IFreight {
	f.SenderId = senderId
	return f
}

func (f Freight) GetDriver() driver.IDriver {
	return f.Driver
}

func (f Freight) GetSender() user.IUser {
	return f.Sender
}
