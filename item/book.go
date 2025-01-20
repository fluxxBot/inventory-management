package item

import (
	"git.jfrog.info/kanishkg/inventory-management/constants"
	"git.jfrog.info/kanishkg/inventory-management/errors"
	"github.com/google/uuid"
)

type Book struct {
	Title    string `json:"title"`
	BaseItem `json:"baseItem"`
}

type BookStore struct {
	Books map[string]*Book
}

func (b Book) GetName() string {
	return b.Title
}

func (b Book) GetId() string {
	return b.id
}

func (b Book) Create(id string) any {
	return &Book{Title: b.GetName(), BaseItem: BaseItem{
		Price:        b.GetPrice(),
		id:           id,
		Quantity:     b.GetQuantity(),
		Category:     b.Category,
		categoryName: inferBookCategoryName(b.Category)},
	}
}

func (b Book) GetPrice() float32 {
	return b.Price
}

func (b Book) GetQuantity() int64 {
	return b.Quantity
}

func (b *BookStore) AddItem(item Item) error {
	if item.GetName() == "" {
		return errors.ForbiddenError{Message: "Empty book name is " + constants.NotAllowedMessage, StatusCode: 403}
	}
	existingItem := b.Books[item.GetName()]
	if existingItem != nil {
		b.Books[item.GetName()].Quantity = b.Books[item.GetName()].Quantity + 1
	} else {
		b.Books[item.GetName()] = item.Create(uuid.NewString()).(*Book)
	}
	return nil
}

func (b *BookStore) RemoveItem(name string) error {
	if b.Books[name] != nil {
		if b.Books[name].GetQuantity() > 1 {
			b.Books[name].Quantity = b.Books[name].Quantity - 1
			return nil
		}
		delete(b.Books, name)
		return nil
	}
	return errors.NotFoundError{"Item " + constants.NotFoundMessage, 404}
}

func (b *BookStore) ListItems() ([]Item, error) {
	var books []Item
	for _, book := range b.Books {
		books = append(books, *book)
	}
	if len(books) == 0 {
		return books, errors.NotFoundError{"Items " + constants.NotFoundMessage, 404}
	}
	return books, nil
}

func inferBookCategoryName(category Category) string {
	switch category {
	case 1:
		return "Literature"
	case 2:
		return "Arts"
	case 3:
		return "Science"
	default:
		return "Unknown"
	}
}
