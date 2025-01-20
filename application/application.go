package application

import (
	"git.jfrog.info/kanishkg/inventory-management/errors"
	"git.jfrog.info/kanishkg/inventory-management/item"
	"git.jfrog.info/kanishkg/inventory-management/store"
	"github.com/gin-gonic/gin"
	"sort"
)

var router = gin.Default()

func Run() {
	application()
	router.Run(":8090")
}

func application() {
	bookStore := item.BookStore{
		Books: make(map[string]*item.Book),
	}
	clothStore := item.ClothStore{
		Clothes: make(map[string]*item.Cloth),
	}

	router.POST("items/books", func(context *gin.Context) {
		bookToAdd := item.Book{}
		err := context.BindJSON(&bookToAdd)
		if err != nil {
			return
		}
		addItem(&bookStore, bookToAdd)
	})
	router.POST("items/cloth", func(context *gin.Context) {
		clothToAdd := item.Cloth{}
		err := context.BindJSON(&clothToAdd)
		if err != nil {
			return
		}
		addItem(&clothStore, clothToAdd)
	})

	router.GET("/items", func(context *gin.Context) {
		context.JSON(200, gin.H{
			"items": getTotalItems(&bookStore, &clothStore),
		})
		context.Header("Content-type", "application/json")
	})

	router.DELETE("/items/:name", func(context *gin.Context) {
		name := context.Param("name")
		err := bookStore.RemoveItem(name)
		if err != nil {
			err.Error()
		}
	})

	router.GET("/items/top", func(context *gin.Context) {
		context.JSON(200, gin.H{
			"items": getTopItems(&bookStore, &clothStore),
		})
		context.Header("Content-type", "application/json")
	})
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
