package user

type IUser interface {
	Clean() interface{}

	GetId() int
	SetId(id int) IUser

	GetName() string
	SetName(name string) IUser

	GetPassword() string
	SetPassword(password string) IUser

	GetEmail() string
	SetEmail(email string) IUser
}
