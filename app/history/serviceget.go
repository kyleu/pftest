// Content managed by Project Forge, see [projectforge.md] for details.
package history

import (
	"context"
	"strings"

	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"

	"github.com/kyleu/pftest/app/lib/database"
	"github.com/kyleu/pftest/app/lib/filter"
)

func (s *Service) List(ctx context.Context, tx *sqlx.Tx, params *filter.Params) (Histories, error) {
	params = filters(params)
	wc := ""
	q := database.SQLSelect(columnsString, tableQuoted, wc, params.OrderByString(), params.Limit, params.Offset)
	ret := dtos{}
	err := s.db.Select(ctx, &ret, q, tx, s.logger)
	if err != nil {
		return nil, errors.Wrap(err, "unable to get histories")
	}
	return ret.ToHistories(), nil
}

func (s *Service) Count(ctx context.Context, tx *sqlx.Tx, whereClause string, args ...any) (int, error) {
	if strings.Contains(whereClause, "'") || strings.Contains(whereClause, ";") {
		return 0, errors.Errorf("invalid where clause [%s]", whereClause)
	}
	q := database.SQLSelectSimple(columnsString, tableQuoted, whereClause)
	ret, err := s.db.SingleInt(ctx, q, tx, s.logger, args...)
	if err != nil {
		return 0, errors.Wrap(err, "unable to get count of histories")
	}
	return int(ret), nil
}

func (s *Service) Get(ctx context.Context, tx *sqlx.Tx, id string) (*History, error) {
	wc := defaultWC(0)
	ret := &dto{}
	q := database.SQLSelectSimple(columnsString, tableQuoted, wc)
	err := s.db.Get(ctx, ret, q, tx, s.logger, id)
	if err != nil {
		return nil, errors.Wrapf(err, "unable to get history by id [%v]", id)
	}
	return ret.ToHistory(), nil
}

func (s *Service) GetMultiple(ctx context.Context, tx *sqlx.Tx, ids ...string) (Histories, error) {
	if len(ids) == 0 {
		return Histories{}, nil
	}
	wc := database.SQLInClause("id", len(ids), 0)
	ret := dtos{}
	q := database.SQLSelectSimple(columnsString, tableQuoted, wc)
	vals := make([]any, 0, len(ids))
	for _, x := range ids {
		vals = append(vals, x)
	}
	err := s.db.Select(ctx, &ret, q, tx, s.logger, vals...)
	if err != nil {
		return nil, errors.Wrapf(err, "unable to get Histories for [%d] ids", len(ids))
	}
	return ret.ToHistories(), nil
}

func (s *Service) ListSQL(ctx context.Context, tx *sqlx.Tx, sql string) (Histories, error) {
	ret := dtos{}
	err := s.db.Select(ctx, &ret, sql, tx, s.logger)
	if err != nil {
		return nil, errors.Wrap(err, "unable to get histories using custom SQL")
	}
	return ret.ToHistories(), nil
}
