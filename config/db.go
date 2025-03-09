package config

import (
	"context"
	"database/sql"
	"kopherlog/ent/migrate"
	"strings"

	entsql "entgo.io/ent/dialect/sql"

	"fmt"
	"kopherlog/ent"
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
	db, err := sql.Open(driver, dsn)
	if err != nil {
		log.Fatal("Fail connect DB", err)
	}
	drv := entsql.OpenDB(driver, db)
	client := ent.NewClient(ent.Driver(drv))
	ctx := context.Background()

	if env == "TEST" {
		log.Println("Running migration...")
		executeSQL(db, "config/drop_table.sql")
		err = client.Schema.Create(
			ctx,
			migrate.WithDropIndex(true),
			migrate.WithDropColumn(true),
		)
		if err != nil {
			log.Fatalf("failed creating schema resources: %v", err)
		}
		// Execute SQL
		executeSQL(db, "config/data.sql")
		return client
	}
	if err := client.Schema.Create(ctx); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}
	return client
}

func executeSQL(db *sql.DB, path string) {
	log.Printf("Executing %s ...", path)
	data, err := os.ReadFile(path)
	if err != nil {
		log.Fatalf("failed reading data.sql: %v", err)
	}
	queries := strings.Split(string(data), ";")
	for _, query := range queries {
		query = strings.TrimSpace(query)
		if query == "" {
			continue
		}
		_, err := db.Exec(query)
		if err != nil {
			log.Fatalf("failed executing query %q: %v", query, err)
		}
	}
}
