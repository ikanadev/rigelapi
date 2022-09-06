package main

import (
	"context"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	_ "github.com/lib/pq"
	"github.com/vmkevv/rigelapi/ent"
)

func main() {
	client, err := ent.Open("postgres", "host=127.0.0.1 port=5432 user=taylor dbname=enttest password=postgres")
	// client, err := ent.Open("postgres", "postgres://taylor:@127.0.0.1:5432/enttest?sslmode=disable")
	if err != nil {
		log.Fatalf("failed opening connection to postgres: %v", err)
	}
	defer client.Close()
	ctx := context.Background()
	err = client.Schema.Create(ctx)
	if err != nil {
		log.Fatalf("failed creating schema: %v", err)
	}
  if err := PopulateStaticJsonData(client, ctx); err != nil {
		log.Fatalf("failed populating static json data: %v", err)
  }
  if err := PopulateStaticData(client, ctx); err != nil {
		log.Fatalf("failed populating static data: %v", err)
  }

	app := fiber.New()
	app.Use(cors.New())

	app.Listen(":8000")
}
