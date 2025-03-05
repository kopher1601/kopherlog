package db

import (
	"context"
	"fmt"
	"kopherlog/ent"
	"os"
)

func WithTx(ctx context.Context, client *ent.Client, testTx *ent.Tx, fn func(tx *ent.Tx) error) error {
	if os.Getenv("ENV") == "TEST" {
		fn(testTx)
		return testTx.Rollback()
	}
	tx, err := client.Tx(ctx)
	if err != nil {
		return err
	}
	defer func() {
		if v := recover(); v != nil {
			tx.Rollback()
			panic(v)
		}
	}()

	if err := fn(tx); err != nil {
		if rerr := tx.Rollback(); rerr != nil {
			err = fmt.Errorf("%w: rolling back transaction: %v", err, rerr)
		}
		return err
	}

	//if os.Getenv("ENV") == "TEST" {
	//	return tx.Rollback()
	//}

	if err := tx.Commit(); err != nil {
		return fmt.Errorf("committing transaction: %w", err)
	}
	return nil
}
