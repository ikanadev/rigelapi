package main

import (
	"context"
	"log"
	"os"

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
	file, err := os.OpenFile("errors.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 644)
	if err != nil {
		log.Fatalf("Error setting up log file: %v", err)
	}
	defer file.Close()
	logger := log.New(file, "", log.LstdFlags)

	server := app.NewServer(entClient, config, logger)
	err = server.Run()
	if err != nil {
		log.Fatalf("Error starting fiber: %v", err)
	}
}
