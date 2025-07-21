package config

import (
	"context"
	"database/sql"
	"entgo.io/ent/dialect"
	entsql "entgo.io/ent/dialect/sql"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/joho/godotenv"
	"github.com/sweetnebulae/go_ent/ent"
	"log"
	"os"
)

// ConnectDB Open new connection
func ConnectDB() *ent.Client {
	err := godotenv.Load("DATABASE_URL")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	databaseUrl := os.Getenv("DATABASE_URL")
	if databaseUrl == "" {
		log.Fatal("DATABASE_URL is not set in the environment")
	}
	db, err := sql.Open("pgx", databaseUrl)
	if err != nil {
		log.Fatal(err)
	}

	// Create an ent.Driver from `db`.
	drv := entsql.OpenDB(dialect.Postgres, db)
	client := ent.NewClient(ent.Driver(drv))
	ctx := context.Background()
	if err := client.Schema.Create(ctx); err != nil {
		log.Fatal("failed creating schema resources")
	}
	return ent.NewClient(ent.Driver(drv))
}

// DisconnectDB close connection
func DisconnectDB(client *ent.Client) {
	if err := client.Close(); err != nil {
		log.Fatal(err)
	}
}
