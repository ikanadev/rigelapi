package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/vmkevv/rigelapi/app"
	"github.com/vmkevv/rigelapi/app/auth"
	"github.com/vmkevv/rigelapi/app/class"
	"github.com/vmkevv/rigelapi/app/extra"
	"github.com/vmkevv/rigelapi/app/location"
	"github.com/vmkevv/rigelapi/app/sync"
	"github.com/vmkevv/rigelapi/config"
	"github.com/vmkevv/rigelapi/database"
)

func setupServer(server app.Server) error {
	auth.Setup(server)
	class.Setup(server)
	extra.Setup(server)
	location.Setup(server)
	sync.Setup(server)
	return nil
}

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

	server := app.NewServer(entClient, config, logger, dbCtx)
	setupServer(server)
	server.Run()
	err = server.App.Listen(fmt.Sprintf("0.0.0.0:%s", server.Config.App.Port))
	if err != nil {
		log.Fatalf("Error starting fiber: %v", err)
	}
}
