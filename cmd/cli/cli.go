package main

import (
	"fmt"
	"github.com/urfave/cli/v2"
	"log"
	"os"
)

func main() {
	app := &cli.App{
		Commands: []*cli.Command{
			{
				Name:  "inventory",
				Usage: "Perform operations on inventory",
				Action: func(context *cli.Context) error {
					fmt.Println("Working!!")
					return nil
				},
			},
			{
				Name:        "add",
				Usage:       "Add an item to inventory",
				Subcommands: append(GetBookCommands(), GetClothCommands()...),
			},
		},
	}
	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
