package main

import (
	"git.jfrog.info/kanishkg/inventory-management/item"
	"git.jfrog.info/kanishkg/inventory-management/pkg/cmd/add"
	"github.com/urfave/cli/v2"
)

var clothMaterial string
var clothQuantity int64
var clothCategory int64
var clothPrice float64

func GetClothCommands() []*cli.Command {
	clothCommand := []*cli.Command{
		{
			Name:                   "cloth",
			Usage:                  "Add a cloth to inventory",
			UseShortOptionHandling: true,
			Flags: []cli.Flag{
				&cli.StringFlag{
					Name:        "material",
					Usage:       "Material of cloth",
					Required:    true,
					Destination: &clothMaterial,
				},
				&cli.Int64Flag{
					Name:        "quantity",
					Usage:       "Quanity of the cloth",
					Required:    true,
					Destination: &clothQuantity,
				},
				&cli.Int64Flag{
					Name:        "category",
					Usage:       "Category to which cloth belongs",
					Required:    true,
					Destination: &clothCategory,
				},
				&cli.Float64Flag{
					Name:        "price",
					Usage:       "Cloth Price",
					Required:    true,
					Destination: &clothPrice,
				},
			},
			Action: addClothToCollection,
		},
	}
	return clothCommand
}

func addClothToCollection(ctx *cli.Context) error {
	reqBody := item.Cloth{
		Material: clothMaterial,
		BaseItem: item.BaseItem{
			Quantity: clothQuantity,
			Price:    float32(clothPrice),
			Category: item.Category(clothCategory),
		},
	}

	add.AddCloth(reqBody)
	return nil
}
