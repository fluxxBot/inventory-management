package application

import (
	"fmt"
	"git.jfrog.info/kanishkg/inventory-management/errors"
	"git.jfrog.info/kanishkg/inventory-management/item"
	"git.jfrog.info/kanishkg/inventory-management/store"
	"sort"
)

func Application() {
	bookStore := item.BookStore{
		Books: make(map[string]*item.Book),
	}
	clothStore := item.ClothStore{
		Clothes: make(map[string]*item.Cloth),
	}

	seedData(&bookStore, &clothStore)

	totalItems := getTotalItems(&bookStore, &clothStore)
	fmt.Println("Items after insertion = ", totalItems)

	if len(totalItems) > 0 {
		err := bookStore.RemoveItem("Geography")
		if err != nil {
			panic(err)
		}
		fmt.Println("Updated list after removing element = ", getTotalItems(&bookStore, &clothStore))
	}

	topItems := getTopItems(&bookStore, &clothStore)
	fmt.Println("Top items = ", topItems)
}

func seedData(bookCollection store.Store, clothCollection store.Store) {
	bookItemHP := item.Book{Title: "Harry Potter", BaseItem: item.BaseItem{Price: 100.5, Quantity: 4, Category: 1}}
	bookItemGeo := item.Book{Title: "Geography", BaseItem: item.BaseItem{Price: 50, Quantity: 10, Category: 2}}
	bookItemHis := item.Book{Title: "History", BaseItem: item.BaseItem{Price: 40.3, Quantity: 25, Category: 2}}
	bookItemLit := item.Book{Title: "Physics", BaseItem: item.BaseItem{Price: 50, Quantity: 12, Category: 3}}

	clothItemWool := item.Cloth{Material: "Cotton", BaseItem: item.BaseItem{Price: 200.7, Quantity: 15, Category: 1}}
	clothItemSilk := item.Cloth{Material: "Silk", BaseItem: item.BaseItem{Price: 100.7, Quantity: 20, Category: 2}}
	clothItemNy := item.Cloth{Material: "Nylon", BaseItem: item.BaseItem{Price: 300.7, Quantity: 30, Category: 2}}

	addItem(bookCollection, bookItemHP)
	addItem(bookCollection, bookItemGeo)
	addItem(bookCollection, bookItemHis)
	addItem(bookCollection, bookItemLit)
	addItem(clothCollection, clothItemWool)
	addItem(clothCollection, clothItemSilk)
	addItem(clothCollection, clothItemNy)
}

func addItem(collection store.Store, item item.Item) {
	err := collection.AddItem(item)
	if err != nil {
		callRespectivePanics(err)
	}
}

func callRespectivePanics(err error) {
	switch err.(type) {
	case errors.NotFoundError:
		println("Panic for NotFoundError")
		err.Error()
	case errors.ForbiddenError:
		println("Panic for ForbiddenError")
		err.Error()
	case nil:
		println("No Panic")
	default:
		println("Executed Successfully")

	}
}

func getTotalItems(collection1 store.Store, collection2 store.Store) []item.Item {
	collectionList1, _ := collection1.ListItems()
	collectionList2, _ := collection2.ListItems()
	return append(collectionList1, collectionList2...)
}

func getTopItems(collection1 store.Store, collection2 store.Store) []item.Item {
	collectionList := getTotalItems(collection1, collection2)
	sort.Slice(collectionList, func(i, j int) bool {
		return collectionList[i].GetQuantity() > collectionList[j].GetQuantity()
	})
	return collectionList[:5]
}
