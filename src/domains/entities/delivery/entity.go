package delivery

import (
	"time"

	d "ddgodeliv/domains/entities/driver"
	u "ddgodeliv/domains/entities/user"
	m "ddgodeliv/domains/models"
)

type Delivery struct {
	Id          int       `db:"id" json:"id"`
	DriverId    int       `db:"driver_id" json:"driver_id"`
	SenderId    int       `db:"sender_id" json:"sender_id" validate:"required"`
	Origin      string    `db:"origin" json:"origin" validate:"required"`
	Destination string    `db:"destination" json:"destination" validate:"required"`
	Deadline    time.Time `db:"deadline" json:"deadline" validate:"future"`
	Status      uint8     `db:"status" json:"status" validate:"required,gte=0,lte=4"`

	Driver m.IDriver `json:"driver"`
	Sender m.IUser   `json:"sender"`
}

func GetNewDelivery() *Delivery {
	return &Delivery{
		Driver: d.GetNewDriver(),
		Sender: u.GetNewUser(),
	}
}

func (d Delivery) GetId() int {
	return d.Id
}

func (d *Delivery) SetId(id int) m.IDelivery {
	d.Id = id
	return d
}

func (d Delivery) GetOrigin() string {
	return d.Origin
}

func (d *Delivery) SetOrigin(origin string) m.IDelivery {
	d.Origin = origin
	return d
}

func (d Delivery) GetDestination() string {
	return d.Destination
}

func (d *Delivery) SetDestination(destination string) m.IDelivery {
	d.Destination = destination
	return d
}

func (d Delivery) GetDeadline() time.Time {
	return d.Deadline
}

func (d *Delivery) SetDeadline(deadline time.Time) m.IDelivery {
	d.Deadline = deadline
	return d
}

func (d Delivery) GetStatus() uint8 {
	return d.Status
}

func (d Delivery) GetStatusDisplay() string {
	return StatusChoices.GetDisplay(d.GetStatus())
}

func (d *Delivery) SetStatus(status uint8) m.IDelivery {
	d.Status = status
	return d
}

func (d Delivery) GetDriverId() int {
	return d.DriverId
}

func (d *Delivery) SetDriverId(driverId int) m.IDelivery {
	d.DriverId = driverId
	return d
}

func (d Delivery) GetSenderId() int {
	return d.SenderId
}

func (d *Delivery) SetSenderId(senderId int) m.IDelivery {
	d.SenderId = senderId
	return d
}

func (d Delivery) GetDriver() m.IDriver {
	return d.Driver
}

func (d Delivery) GetSender() m.IUser {
	return d.Sender
}
