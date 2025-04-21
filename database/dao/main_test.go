package dao

import (
	"context"
	"medods/database"
	"medods/utils"
	"os"
	"testing"

	"github.com/go-pg/pg/v10"
)

func TestMain(m *testing.M) {
	// TODO: parse from config/env? setup image?

	db := pg.Connect(&pg.Options{
		Addr:     "0.0.0.0:5432",
		User:     "postgres",
		Password: "postgres",
		Database: "medods",
	})

	if err := db.Ping(context.Background()); err != nil {
		utils.Panicf("error on Ping - %v", err)
	}

	tx, err := db.Begin()
	if err != nil {
		utils.Panicf("error on Begin tx - %v", err)
	}
	defer func() {
		if err := tx.Rollback(); err != nil {
			utils.Panicf("error on Rollback tx - %v", err)
		}
	}()

	database.SetDatabase(tx)

	os.Exit(m.Run())
}

func fatalOnError(t *testing.T, err error) {
	if err != nil {
		t.Fatal(err)
	}
}
