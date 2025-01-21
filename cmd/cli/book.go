package main

import (
	"git.jfrog.info/kanishkg/inventory-management/item"
	"git.jfrog.info/kanishkg/inventory-management/pkg/cmd/add"
	"github.com/urfave/cli/v2"
)

var bookTitle string
var bookQuantity int64
var bookCategory int64
var bookPrice float64

func GetBookCommands() []*cli.Command {
	bookCommand := []*cli.Command{
		{
			Name:                   "book",
			Usage:                  "Add a book to inventory",
			UseShortOptionHandling: true,
			Flags: []cli.Flag{
				&cli.StringFlag{
					Name:        "title",
					Usage:       "Title of the book",
					Required:    true,
					Destination: &bookTitle,
				},
				&cli.Int64Flag{
					Name:        "quantity",
					Usage:       "Quanity of the book",
					Required:    true,
					Destination: &bookQuantity,
				},
				&cli.Int64Flag{
					Name:        "category",
					Usage:       "Category to which book belongs",
					Required:    true,
					Destination: &bookCategory,
				},
				&cli.Float64Flag{
					Name:        "price",
					Usage:       "Book Price",
					Required:    true,
					Destination: &bookPrice,
				},
			},
			Action: addBookToCollection,
		},
	}
	return bookCommand
}

func addBookToCollection(ctx *cli.Context) error {
	reqBody := item.Book{
		Title: bookTitle,
		BaseItem: item.BaseItem{
			Quantity: bookQuantity,
			Price:    float32(bookPrice),
			Category: item.Category(bookCategory),
		},
	}

	add.AddBook(reqBody)
	return nil
}
