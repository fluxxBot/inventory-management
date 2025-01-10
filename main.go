package main

import (
	"git.jfrog.info/kanishkg/inventory-management/item"
	"git.jfrog.info/kanishkg/inventory-management/store"
	"git.jfrog.info/kanishkg/inventory-management/utility"
)

func main() {
	bookItemHP := item.Book{Title: "Harry Potter"}
	bookItemB1 := item.Book{Title: "Geography"}
	clothItem := item.Cloth{Material: "Cotton"}
	err := store.AddItem(&bookItemHP)
	if err := store.AddItem(&bookItemHP); err == nil {
		if err := store.AddItem(&bookItemB1); err == nil {
			err = store.AddItem(&clothItem)
		}
	}

	store.ListItems()

	switch err.(type) {
	case utility.NotFoundError:
		println("Panic for NotFoundError")
		err.Error()
	case utility.ForbiddenError:
		println("Panic for ForbiddenError")
		err.Error()
	case nil:
		println("No Panic")
	default:
		println("Executed Successfully")

	}

	// below code is to remove item from store
	itemsPresentInStore, err := store.GetItems()
	if err != nil {
		err.Error()
	}
	err = store.RemoveItem(itemsPresentInStore[0].GetId())
	if err != nil {
		panic(err)
	}
	store.ListItems()
}
