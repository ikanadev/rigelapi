package main

import (
	"context"
	"log"

	"github.com/vmkevv/rigelapi/app"
	"github.com/vmkevv/rigelapi/config"
	"github.com/vmkevv/rigelapi/database"
)

// @title          Rigel API
// @version        1.0
// @description    Backend json services for Rigel WebApp
// @termsOfService http://swagger.io/terms/

// @contact.name  API Support
// @contact.url   https://t.me/vmkevv
// @contact.email vargaskevv@gmail.com

// @license.name Apache 2.0
// @license.url  http://www.apache.org/licenses/LICENSE-2.0.html

// @host     localhost:4000
// @BasePath /
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
