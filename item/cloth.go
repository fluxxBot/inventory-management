package item

import (
	"git.jfrog.info/kanishkg/inventory-management/constants"
	"git.jfrog.info/kanishkg/inventory-management/errors"
	"github.com/google/uuid"
)

type Cloth struct {
	Material string `json:"material"`
	BaseItem
}

type ClothStore struct {
	Clothes map[string]*Cloth
}

func (c Cloth) GetName() string {
	return c.Material
}

func (c Cloth) GetId() string {
	return c.id
}

func (c Cloth) GetQuantity() int64 {
	return c.Quantity
}

func (c Cloth) GetPrice() float32 {
	return c.Price
}

func (c Cloth) Create(id string) any {
	return &Cloth{Material: c.GetName(), BaseItem: BaseItem{
		Price:        c.GetPrice(),
		id:           id,
		Quantity:     c.GetQuantity(),
		Category:     c.Category,
		categoryName: inferClothCategoryName(c.Category)},
	}
}

func (c *ClothStore) AddItem(item Item) error {
	if item.GetName() == "" {
		return errors.ForbiddenError{Message: "Empty cloth material is " + constants.NotAllowedMessage, StatusCode: 403}
	}
	existingItem := c.Clothes[item.GetName()]
	if existingItem != nil {
		c.Clothes[item.GetName()].Quantity = c.Clothes[item.GetName()].Quantity + 1
	} else {
		c.Clothes[item.GetName()] = item.Create(uuid.NewString()).(*Cloth)
	}
	return nil
}

func (c *ClothStore) RemoveItem(name string) error {
	if c.Clothes[name] != nil {
		if c.Clothes[name].Quantity > 1 {
			c.Clothes[name].Quantity = c.Clothes[name].Quantity - 1
			return nil
		}
		delete(c.Clothes, name)
		return nil
	}
	return errors.NotFoundError{"Item " + constants.NotFoundMessage, 404}
}

func (c *ClothStore) ListItems() ([]Item, error) {
	var cloths []Item
	for _, cloth := range c.Clothes {
		cloths = append(cloths, *cloth)
	}
	if len(cloths) == 0 {
		return cloths, errors.NotFoundError{"Items " + constants.NotFoundMessage, 404}
	}
	return cloths, nil
}

func inferClothCategoryName(category Category) string {
	switch category {
	case 1:
		return "Cotton"
	case 2:
		return "Others"
	default:
		return "Unknown"
	}
}
