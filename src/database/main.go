package database

import (
	"database/sql"
	"log/slog"

	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/sqlitedialect"
	"github.com/uptrace/bun/driver/sqliteshim"
)

func New(logger *slog.Logger) *bun.DB {
	sqldb, err := sql.Open(sqliteshim.ShimName, "file:database.db?cache=shared")
	if err != nil {
		panic(err)
	}

	db := bun.NewDB(sqldb, sqlitedialect.New())

	return db
}
