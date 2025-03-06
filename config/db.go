package config

import (
	"context"
	"fmt"
	"kopherlog/ent"
	"kopherlog/ent/migrate"
	"log"
	"os"
)

func SetupDB() *ent.Client {
	env := os.Getenv("ENV")
	driver := os.Getenv("DB_DRIVER")
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
	)
	client, err := ent.Open(driver, dsn)
	if err != nil {
		log.Fatal("Fail connect DB", err)
	}
	ctx := context.Background()

	if env == "test" {
		log.Println("Running migration...")
		err = client.Schema.Create(
			ctx,
			migrate.WithDropIndex(true),
			migrate.WithDropColumn(true),
		)
		if err != nil {
			log.Fatalf("failed creating schema resources: %v", err)
		}
		return client
	}
	if err := client.Schema.Create(ctx); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}
	return client
}
