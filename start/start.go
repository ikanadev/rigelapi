package main

import (
	"context"
	"log"

	_ "github.com/lib/pq"
	"github.com/vmkevv/rigelapi/ent"
)

func main() {
	client, err := ent.Open("postgres", "host=127.0.0.1 port=5432 user=taylor dbname=enttest password=postgres")
	// client, err := ent.Open("postgres", "postgres://taylor:@127.0.0.1:5432/enttest?sslmode=disable")
	if err != nil {
		log.Fatalf("failed opening connection to postgres: %v", err)
	}
	ctx := context.Background()
	defer client.Close()
	err = client.Schema.Create(ctx)
	if err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}
	if err := client.Schema.Create(ctx); err != nil {
		log.Fatalf("failed creating uuid: %v", err)
	}

	// CreateUser(ctx, client)
	// CreateCars(ctx, client)
	// AddCars(ctx, client)
	// QueryCarUsers(ctx, client)
	log.Println("All right")
}
