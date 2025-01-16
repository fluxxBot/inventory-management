package store

import (
	"git.jfrog.info/kanishkg/inventory-management/item"
)

type Store interface {
	AddItem(item item.Item) error
	RemoveItem(name string) error
	ListItems() ([]item.Item, error)
}
