package company

type ICompany interface {
	GetId() int
	SetId(id int) ICompany

	GetName() string
	SetName(name string) ICompany

	GetOwnerId() int
}
