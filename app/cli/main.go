package main

import (
	"github.com/tasuke/go-onion/infrastructure/db/schema"
	"log"
	"os"

	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Commands = []cli.Command{
		{
			Name:   "migration",
			Usage:  "DBマイグレーション",
			Action: schema.Migrate,
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
