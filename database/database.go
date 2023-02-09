package database

import (
	"context"
	"fmt"
	"log"

	_ "github.com/lib/pq"
	"github.com/vmkevv/rigelapi/config"
	"github.com/vmkevv/rigelapi/ent"
)

func SetUpDB(ctx context.Context, config config.Config) (*ent.Client, error) {
	dbConfig := config.DB
	// client, err := ent.Open("postgres", "host=127.0.0.1 port=5432 user=taylor dbname=enttest password=postgres")
	log.Println("Connecting to postgres...")
	client, err := ent.Open(
		"postgres",
		fmt.Sprintf(
			"host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
			dbConfig.Host,
			dbConfig.Port,
			dbConfig.User,
			dbConfig.Name,
			dbConfig.Password,
			dbConfig.SslMode,
		),
	)
	if err != nil {
		return nil, err
	}
	log.Println("creating schema...")
	err = client.Schema.Create(ctx)
	if err != nil {
		return nil, err
	}
	log.Println("populating static data...")
	if err := PopulateStaticData(client, ctx); err != nil {
		return nil, err
	}
	return client, nil
}
