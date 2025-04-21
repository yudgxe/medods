package migration

import (
	"database/sql"
	"embed"
	"fmt"

	"github.com/go-pg/pg/v10"
	_ "github.com/lib/pq"
	"github.com/pressly/goose/v3"
)

//go:embed migrations/*.sql
var embedMigrations embed.FS

const LastVersion int64 = -1

const dir string = "migrations"

func Up(options pg.Options, version int64) error {

	db, err := sql.Open("postgres", getPostgreDBString(options))
	if err != nil {
		return err
	}

	goose.SetBaseFS(embedMigrations)

	if version == LastVersion {
		if err := goose.Up(db, dir); err != nil {
			return err
		}
	} else {
		if err := goose.UpTo(db, dir, version); err != nil {
			return err
		}
	}

	return nil
}

func Down(options pg.Options, version int64) error {
	db, err := sql.Open("postgres", getPostgreDBString(options))
	if err != nil {
		return err
	}

	goose.SetBaseFS(embedMigrations)

	if version == LastVersion {
		if err := goose.Down(db, dir); err != nil {
			return err
		}
	} else {
		if err := goose.DownTo(db, dir, version); err != nil {
			return err
		}
	}

	return nil
}

func getPostgreDBString(options pg.Options) string {
	return fmt.Sprintf("postgresql://%s:%s@%s/%s?sslmode=disable",
		options.User, options.Password, options.Addr, options.Database,
	)
}
