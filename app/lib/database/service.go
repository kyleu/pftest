// Content managed by Project Forge, see [projectforge.md] for details.
package database

import (
	"context"
	"database/sql"
	"fmt"
	"strings"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
	"go.uber.org/zap"

	"github.com/kyleu/pftest/app/lib/telemetry"
	"github.com/kyleu/pftest/app/lib/telemetry/dbmetrics"
	"github.com/kyleu/pftest/queries"
	"github.com/kyleu/pftest/queries/schema"
)

type DBType struct {
	Key               string `json:"key"`
	Title             string `json:"title"`
	Quote             string `json:"-"`
	Placeholder       string `json:"-"`
	SupportsReturning bool   `json:"-"`
}

type Service struct {
	Key          string  `json:"key"`
	DatabaseName string  `json:"database,omitempty"`
	SchemaName   string  `json:"schema,omitempty"`
	Username     string  `json:"username,omitempty"`
	Debug        bool    `json:"debug,omitempty"`
	Type         *DBType `json:"type"`
	db           *sqlx.DB
	metrics      *dbmetrics.Metrics
}

func NewService(typ *DBType, key string, dbName string, schName string, username string, debug bool, db *sqlx.DB, logger *zap.SugaredLogger) (*Service, error) {
	if logger == nil {
		return nil, errors.New("logger must be provided to database service")
	}
	logger = logger.With("database", dbName, "user", username)
	m, err := dbmetrics.NewMetrics(strings.ReplaceAll(key, "-", "_"), db)
	if err != nil {
		logger.Warnf(fmt.Sprintf("unable to register database metrics for [%s]: %+v", key, err))
	}

	ret := &Service{Key: key, DatabaseName: dbName, SchemaName: schName, Username: username, Debug: debug, Type: typ, db: db, metrics: m}
	err = ret.Healthcheck(dbName, db)
	if err != nil {
		return nil, errors.Wrap(err, "unable to run healthcheck")
	}
	return ret, nil
}

func (s *Service) Healthcheck(dbName string, db *sqlx.DB) error {
	q := queries.Healthcheck()
	res, err := db.Query(q)
	if err != nil || res.Err() != nil {
		if err == nil {
			err = res.Err()
		}
		if strings.Contains(err.Error(), "does not exist") {
			return errors.Wrapf(err, "database does not exist; run the following:\n"+schema.CreateDatabase())
		}
		return errors.Wrapf(err, "unable to run healthcheck [%s]", q)
	}
	defer func() { _ = res.Close() }()
	return nil
}

func (s *Service) StartTransaction(logger *zap.SugaredLogger) (*sqlx.Tx, error) {
	if s.Debug {
		logger.Debug("opening transaction")
	}
	return s.db.Beginx()
}

func (s *Service) Conn(ctx context.Context) (*sql.Conn, error) {
	return s.db.Conn(ctx)
}

func (s *Service) Stats() sql.DBStats {
	return s.db.Stats()
}

func errMessage(t string, q string, values []any) string {
	return fmt.Sprintf("error running %s sql [%s] with values [%s]", t, strings.TrimSpace(q), valueStrings(values))
}

func (s *Service) logQuery(ctx context.Context, msg string, q string, logger *zap.SugaredLogger, values []any) {
	if s.Debug {
		logger.Debugf("%s {\n  SQL: %s\n  Values: %s\n}", msg, strings.TrimSpace(q), valueStrings(values))
	}
}

func (s *Service) newSpan(
	ctx context.Context, name string, q string, logger *zap.SugaredLogger,
) (time.Time, context.Context, *telemetry.Span, *zap.SugaredLogger) {
	if s.metrics != nil {
		s.metrics.IncStmt(q, name)
	}
	nc, span, logger := telemetry.StartSpan(ctx, "database"+name, logger)
	span.Attributes(
		&telemetry.Attribute{Key: "db.statement", Value: q},
		&telemetry.Attribute{Key: "db.system", Value: s.db.DriverName()},
		&telemetry.Attribute{Key: "db.name", Value: s.DatabaseName},
		&telemetry.Attribute{Key: "db.user", Value: s.Username},
	)
	return time.Now(), nc, span, logger
}

func (s *Service) complete(q string, op string, span *telemetry.Span, started time.Time, logger *zap.SugaredLogger, err error) {
	if err != nil {
		span.OnError(err)
	}
	span.Complete()
	if s.metrics != nil {
		s.metrics.CompleteStmt(q, op, started)
	}
}

func (s *Service) Close() error {
	if s.metrics != nil {
		_ = s.metrics.Close()
	}
	return s.db.Close()
}
