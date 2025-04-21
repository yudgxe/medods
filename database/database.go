package database

import (
	"context"
	"medods/utils"

	"github.com/go-pg/pg/v10"
	"github.com/go-pg/pg/v10/orm"
)

var database orm.DB

func MustInitDatabase(ctx context.Context, options pg.Options) {
	db := pg.Connect(&options)

	if err := db.Ping(ctx); err != nil {
		utils.Panicf("error on Ping db - %v", err)
	}

	database = db
}

func GetDatabase() orm.DB {
	if database != nil {
		return database
	}

	panic("error on GetDatabase - need setup db")
}

func SetDatabase(db orm.DB) { // for tests
	database = db
}
