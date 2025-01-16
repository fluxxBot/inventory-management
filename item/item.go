package item

type Category int32

type BaseItem struct {
	id           string
	Quantity     int64
	Price        float32
	Category     Category
	categoryName string
}

type Item interface {
	GetName() string
	GetId() string
	GetQuantity() int64
	GetPrice() float32
	Create(id string) any
}
