// Content managed by Project Forge, see [projectforge.md] for details.
package capital

import (
	"context"
	"strings"

	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"

	"github.com/kyleu/pftest/app/lib/database"
	"github.com/kyleu/pftest/app/lib/filter"
	"github.com/kyleu/pftest/app/util"
)

func (s *Service) List(ctx context.Context, tx *sqlx.Tx, params *filter.Params, logger util.Logger) (Capitals, error) {
	params = filters(params)
	wc := ""
	q := database.SQLSelect(columnsString, tablesJoined, wc, params.OrderByString(), params.Limit, params.Offset)
	ret := dtos{}
	err := s.dbRead.Select(ctx, &ret, q, tx, logger)
	if err != nil {
		return nil, errors.Wrap(err, "unable to get capitals")
	}
	return ret.ToCapitals(), nil
}

func (s *Service) Count(ctx context.Context, tx *sqlx.Tx, whereClause string, logger util.Logger, args ...any) (int, error) {
	if strings.Contains(whereClause, "'") || strings.Contains(whereClause, ";") {
		return 0, errors.Errorf("invalid where clause [%s]", whereClause)
	}
	q := database.SQLSelectSimple(columnsString, tablesJoined, whereClause)
	ret, err := s.dbRead.SingleInt(ctx, q, tx, logger, args...)
	if err != nil {
		return 0, errors.Wrap(err, "unable to get count of capitals")
	}
	return int(ret), nil
}

func (s *Service) Get(ctx context.Context, tx *sqlx.Tx, id string, logger util.Logger) (*Capital, error) {
	wc := defaultWC(0)
	ret := &dto{}
	q := database.SQLSelectSimple(columnsString, tablesJoined, wc)
	err := s.dbRead.Get(ctx, ret, q, tx, logger, id)
	if err != nil {
		return nil, errors.Wrapf(err, "unable to get capital by id [%v]", id)
	}
	return ret.ToCapital(), nil
}

func (s *Service) GetMultiple(ctx context.Context, tx *sqlx.Tx, logger util.Logger, IDs ...string) (Capitals, error) {
	if len(IDs) == 0 {
		return Capitals{}, nil
	}
	wc := database.SQLInClause("ID", len(IDs), 0)
	ret := dtos{}
	q := database.SQLSelectSimple(columnsString, tablesJoined, wc)
	vals := make([]any, 0, len(IDs))
	for _, x := range IDs {
		vals = append(vals, x)
	}
	err := s.dbRead.Select(ctx, &ret, q, tx, logger, vals...)
	if err != nil {
		return nil, errors.Wrapf(err, "unable to get Capitals for [%d] IDs", len(IDs))
	}
	return ret.ToCapitals(), nil
}

func (s *Service) ListSQL(ctx context.Context, tx *sqlx.Tx, sql string, logger util.Logger, values ...any) (Capitals, error) {
	ret := dtos{}
	err := s.dbRead.Select(ctx, &ret, sql, tx, logger, values...)
	if err != nil {
		return nil, errors.Wrap(err, "unable to get capitals using custom SQL")
	}
	return ret.ToCapitals(), nil
}
