package config

import (
	"entgo.io/ent/dialect/sql"
	"fmt"
	"kopherlog/ent"
	"kopherlog/ent/enttest"
	"log"
	"os"
	"testing"
)

func SetupDB(t *testing.T) *ent.Client {
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

	drv := sql.OpenDB(os.Getenv("DB_DRIVER"), db.DB())
	if env == "test" {
		options := []enttest.Option{
			enttest.WithOptions(ent.Driver(drv)),
		}
		return enttest.NewClient(t, options...)
	}
	options := []ent.Option{
		ent.Driver(drv),
	}
	return ent.NewClient(options...)
}
