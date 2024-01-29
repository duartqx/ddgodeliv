package user

type IUserRepository interface {
	FindById(id int) (IUser, error)

	FindByEmail(email string) (IUser, error)
	ExistsByEmail(email string) bool

	Create(user IUser) error
	Update(user IUser) error
	Delete(user IUser) error
}
