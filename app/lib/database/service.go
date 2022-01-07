package database

import (
	"context"
	"database/sql"
	"fmt"
	"strings"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
	semconv "go.opentelemetry.io/otel/semconv/v1.4.0"
	"go.opentelemetry.io/otel/trace"
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
	Type         *DBType `json:"type"`
	db           *sqlx.DB
	metrics      *dbmetrics.Metrics
	logger       *zap.SugaredLogger
}

func NewService(typ *DBType, key string, dbName string, schName string, username string, db *sqlx.DB, logger *zap.SugaredLogger) (*Service, error) {
	m, err := dbmetrics.NewMetrics(key, db)
	if err != nil {
		logger.Warnf("unable to register database metrics for [%s]: %+v", key, err)
	}

	ret := &Service{Key: key, DatabaseName: dbName, SchemaName: schName, Username: username, Type: typ, db: db, metrics: m, logger: logger}
	err = ret.Healthcheck(db)
	if err != nil {
		return nil, errors.Wrap(err, "unable to run healthcheck")
	}
	return ret, nil
}

func (s *Service) Healthcheck(db *sqlx.DB) error {
	res, err := db.Query(queries.Healthcheck())
	if err != nil || res.Err() != nil {
		if err == nil {
			err = res.Err()
		}
		if strings.Contains(err.Error(), "does not exist") {
			return errors.Wrapf(err, "database does not exist; run the following:\n"+schema.CreateDatabase())
		}
		return errors.Wrapf(err, "unable to run healthcheck [%s]", queries.Healthcheck())
	}
	defer func() { _ = res.Close() }()
	return nil
}

func (s *Service) StartTransaction() (*sqlx.Tx, error) {
	if s.logger != nil {
		s.logger.Info("opening transaction")
	}
	return s.db.Beginx()
}

func (s *Service) Conn(ctx context.Context) (*sql.Conn, error) {
	return s.db.Conn(ctx)
}

func (s *Service) Stats() sql.DBStats {
	return s.db.Stats()
}

func errMessage(t string, q string, values []interface{}) string {
	return fmt.Sprintf("error running %s sql [%s] with values [%s]", t, strings.TrimSpace(q), valueStrings(values))
}

func (s *Service) logQuery(msg string, q string, values []interface{}) {
	if s.logger != nil {
		s.logger.Infof("%s {\n  SQL: %s\n  Values: %s\n}", msg, strings.TrimSpace(q), valueStrings(values))
	}
}

func (s *Service) newSpan(ctx context.Context, name string, q string) (time.Time, context.Context, trace.Span) {
	if s.metrics != nil {
		s.metrics.IncStmt(q, name)
	}
	nc, span := telemetry.StartSpan(ctx, "database", name)
	span.SetAttributes(
		semconv.DBStatementKey.String(q),
		semconv.DBSystemPostgreSQL,
		semconv.DBNameKey.String(s.DatabaseName),
		semconv.DBUserKey.String(s.Username),
	)
	return time.Now(), nc, span
}

func (s *Service) complete(q string, op string, span trace.Span, started time.Time, err error) {
	if err != nil {
		span.RecordError(err)
	}
	span.End()
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
