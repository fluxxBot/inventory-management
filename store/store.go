package store

import (
	"fmt"
	"git.jfrog.info/kanishkg/inventory-management/utility"
	"github.com/google/uuid"
)

type Item interface {
	GetName() string
	GetId() string
	Create(id string) any
}

var inventory = make([]Item, 0)

func AddItem(item Item) error {
	if item.GetName() == "" {
		return utility.ForbiddenError{Message: "Item Name is Empty", StatusCode: 403}
	}
	id := uuid.NewString()
	newItem := item.Create(id)
	inventory = append(inventory, newItem.(Item))
	return nil
}

func RemoveItem(id string) error {
	for i, item := range inventory {
		if item.GetId() == id {
			inventory = append(inventory[:i], inventory[i+1:]...)
			fmt.Println("Successfully Removed Item with id: ", id)
			return nil
		}
	}
	fmt.Println("Item Not Present")
	return utility.NotFoundError{"Item Not Found", 404}
}

func GetItems() ([]Item, error) {
	if len(inventory) == 0 {
		return nil, utility.NotFoundError{"No Items Found", 404}
	}
	return inventory, nil
}

func ListItems() {
	fmt.Println(getFormattedItem())
}

func getFormattedItem() (formattedItems []string) {
	for _, item := range inventory {
		formattedItems = append(formattedItems, "{Item Name: ", item.GetName(), ", Item Id: ",
			fmt.Sprint(item.GetId()), "}\n")
	}
	return formattedItems
}
