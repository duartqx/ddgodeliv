package company

type Company struct {
	Id   int    `db:"id" json:"id"`
	Name string `db:"name" json:"name"`
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
