package user

type User struct {
	Id       int    `db:"id" json:"id"`
	Email    string `db:"email" json:"email" validate:"email,required"`
	Password string `db:"password" json:"-" validate:"required,min=8,max=200"`
	Name     string `db:"name" json:"name" validate:"required,min=3,max=50"`
}

func GetNewUser() *User {
	return &User{}
}

func (u User) GetId() int {
	return u.Id
}

func (u *User) SetId(id int) IUser {
	u.Id = id
	return u
}

func (u User) GetName() string {
	return u.Name
}

func (u *User) SetName(name string) IUser {
	u.Name = name
	return u
}

func (u User) GetPassword() string {
	return u.Password
}

func (u *User) SetPassword(password string) IUser {
	u.Password = password
	return u
}

func (u User) GetEmail() string {
	return u.Email
}

func (u *User) SetEmail(email string) IUser {
	u.Email = email
	return u
}
