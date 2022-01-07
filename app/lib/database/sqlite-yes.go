//go:build darwin || (!android && linux && 386) || (!android && linux && amd64) || (!android && linux && arm) || (!android && linux && arm64) || (windows && 386) || (windows && amd64)
// +build darwin !android,linux,386 !android,linux,amd64 !android,linux,arm !android,linux,arm64 windows,386 windows,amd64

package database

import (
	"context"

	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
	"go.uber.org/zap"

	// load sqlite driver.
	_ "modernc.org/sqlite"

	"github.com/kyleu/pftest/app/lib/telemetry"
)

const SQLiteEnabled = true

var typeSQLite = &DBType{Key: "sqlite", Title: "SQLite", Quote: `"`, Placeholder: "$", SupportsReturning: true}

func OpenSQLiteDatabase(ctx context.Context, key string, params *SQLiteParams, logger *zap.SugaredLogger) (*Service, error) {
	_, span := telemetry.StartSpan(ctx, "database", "open")
	defer span.End()
	if params.File == "" {
		return nil, errors.New("need filename for SQLite database")
	}
	db, err := sqlx.Open("sqlite", params.File)
	if err != nil {
		return nil, errors.Wrap(err, "error opening database")
	}
	var log *zap.SugaredLogger
	if params.Debug {
		log = logger
	}
	return NewService(typeSQLite, key, key, params.Schema, "sqlite", db, log)
}
