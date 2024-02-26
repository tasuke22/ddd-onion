package main

import (
	"context"
	"github.com/tasuke/go-onion/config"
	"github.com/tasuke/go-onion/infrastructure/db"
	"github.com/tasuke/go-onion/server"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	conf := config.GetConfig()
	db.NewMainDB(conf.DB)

	server.Run(ctx, conf)
}
