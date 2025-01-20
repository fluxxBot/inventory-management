package item

type Category int32

type BaseItem struct {
	id           string
	Quantity     int64    `json:"quantity"`
	Price        float32  `json:"price"`
	Category     Category `json:"category"`
	categoryName string
}

type Item interface {
	GetName() string
	GetId() string
	GetQuantity() int64
	GetPrice() float32
	Create(id string) any
}
