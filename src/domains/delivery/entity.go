package delivery

import (
	"time"

	d "ddgodeliv/domains/driver"
	u "ddgodeliv/domains/user"
)

type Delivery struct {
	Id          int       `db:"id" json:"id"`
	Loadout     string    `db:"loadout" json:"loadout" validate:"required"`
	Weight      int       `db:"weight" json:"weight"` // milligrams
	DriverId    int       `db:"driver_id" json:"-"`
	SenderId    int       `db:"sender_id" json:"-" validate:"required,gt=0"`
	Origin      string    `db:"origin" json:"origin" validate:"required,min=2"`
	Destination string    `db:"destination" json:"destination" validate:"required,min=2"`
	CreatedAt   time.Time `db:"created_at" json:"created_at"`
	Deadline    time.Time `db:"deadline" json:"deadline" validate:"future"`
	Status      uint8     `db:"status" json:"status" validate:"gte=0,lte=4"`

	Driver d.Driver `db:"driver" json:"driver" validate:"-"`
	Sender u.User   `db:"sender" json:"sender" validate:"-"`
}

func GetNewDelivery() *Delivery {
	return &Delivery{
		Driver: *d.GetNewDriver(),
		Sender: *u.GetNewUser(),
	}
}

func (d Delivery) GetId() int {
	return d.Id
}

func (d *Delivery) SetId(id int) IDelivery {
	d.Id = id
	return d
}

func (d Delivery) GetLoadout() string {
	return d.Loadout
}

func (d *Delivery) SetLoadout(loadout string) IDelivery {
	d.Loadout = loadout
	return d
}

func (d Delivery) GetWeight() int {
	return d.Weight
}

func (d *Delivery) SetWeight(weight int) IDelivery {
	d.Weight = weight
	return d
}

func (d Delivery) GetOrigin() string {
	return d.Origin
}

func (d *Delivery) SetOrigin(origin string) IDelivery {
	d.Origin = origin
	return d
}

func (d Delivery) GetDestination() string {
	return d.Destination
}

func (d *Delivery) SetDestination(destination string) IDelivery {
	d.Destination = destination
	return d
}

func (d Delivery) GetCreatedAt() time.Time {
	return d.CreatedAt
}

func (d *Delivery) SetCreatedAt(createdAt time.Time) IDelivery {
	d.CreatedAt = createdAt
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

func (d Delivery) GetStatusDisplay() string {
	return StatusChoices.GetDisplay(d.GetStatus())
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

func (d *Delivery) SetDriver(driver d.IDriver) IDelivery {

	d.SetDriverId(driver.GetId())

	d.Driver.
		SetId(driver.GetId()).
		SetUserId(driver.GetUserId()).
		SetCompanyId(driver.GetCompanyId())

	d.Driver.User.Id = driver.GetUserId()
	d.Driver.User.Name = driver.GetUser().GetName()
	d.Driver.User.Email = driver.GetUser().GetEmail()

	return d
}

func (d Delivery) GetSenderId() int {
	return d.SenderId
}

func (d *Delivery) SetSenderId(senderId int) IDelivery {
	d.SenderId = senderId
	return d
}

func (d Delivery) GetDriver() d.IDriver {
	return &d.Driver
}

func (d Delivery) GetSender() u.IUser {
	return &d.Sender
}

func (d Delivery) HasInvalidId() bool {
	return d.GetId() == 0
}

func (d Delivery) DriverIsNull() bool {
	return d.GetDriverId() == 0
}

func (d Delivery) IsPending() bool {
	return d.Status == StatusChoices.Pending
}
