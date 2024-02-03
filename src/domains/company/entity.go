package company

type Company struct {
	Id      int    `db:"id" json:"id"`
	OwnerId int    `db:"owner_id" json:"owner_id" validate:"required,gt=0"`
	Name    string `db:"name" json:"name" validate:"required,min=3"`
}

func GetNewCompany() *Company {
	return &Company{}
}

func (c Company) GetId() int {
	return c.Id
}

func (c *Company) SetId(id int) ICompany {
	c.Id = id
	return c
}

func (c Company) GetName() string {
	return c.Name
}

func (c *Company) SetName(name string) ICompany {
	c.Name = name
	return c
}

func (c Company) GetOwnerId() int {
	return c.OwnerId
}

func (c *Company) SetOwnerId(id int) ICompany {
	c.OwnerId = id
	return c
}

func (c Company) HasInvalidId() bool {
	return c.GetId() == 0
}
