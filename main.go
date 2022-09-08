package main

import (
	"context"
	"log"

	"github.com/vmkevv/rigelapi/app"
	"github.com/vmkevv/rigelapi/config"
	"github.com/vmkevv/rigelapi/database"
)

func main() {
	config := config.GetConfig()
	dbCtx := context.Background()
	entClient, err := database.SetUpDB(dbCtx, config)
	if err != nil {
		log.Fatalf("Error setting up database: %v", err)
	}

	server := app.NewServer(entClient, config)
	server.Run()
}
