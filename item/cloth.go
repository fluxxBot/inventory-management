package item

type Cloth struct {
	Material string
	id       string
}

func (c Cloth) GetName() string {
	return c.Material
}

func (c Cloth) GetId() string {
	return c.id
}

func (c Cloth) Create(id string) any {
	return Cloth{Material: c.GetName(), id: id}
}
