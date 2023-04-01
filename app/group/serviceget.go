// Content managed by Project Forge, see [projectforge.md] for details.
package group

import (
	"context"
	"strings"

	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"

	"github.com/kyleu/pftest/app/lib/database"
	"github.com/kyleu/pftest/app/lib/filter"
	"github.com/kyleu/pftest/app/util"
)

func (s *Service) List(ctx context.Context, tx *sqlx.Tx, params *filter.Params, logger util.Logger) (Groups, error) {
	params = filters(params)
	wc := ""
	q := database.SQLSelect(columnsString, tableQuoted, wc, params.OrderByString(), params.Limit, params.Offset)
	ret := rows{}
	err := s.dbRead.Select(ctx, &ret, q, tx, logger)
	if err != nil {
		return nil, errors.Wrap(err, "unable to get groups")
	}
	return ret.ToGroups(), nil
}

func (s *Service) Count(ctx context.Context, tx *sqlx.Tx, whereClause string, logger util.Logger, args ...any) (int, error) {
	if strings.Contains(whereClause, "'") || strings.Contains(whereClause, ";") {
		return 0, errors.Errorf("invalid where clause [%s]", whereClause)
	}
	q := database.SQLSelectSimple("count(*) as x", tableQuoted, whereClause)
	ret, err := s.dbRead.SingleInt(ctx, q, tx, logger, args...)
	if err != nil {
		return 0, errors.Wrap(err, "unable to get count of groups")
	}
	return int(ret), nil
}

func (s *Service) Get(ctx context.Context, tx *sqlx.Tx, id string, logger util.Logger) (*Group, error) {
	wc := defaultWC(0)
	ret := &row{}
	q := database.SQLSelectSimple(columnsString, tableQuoted, wc)
	err := s.dbRead.Get(ctx, ret, q, tx, logger, id)
	if err != nil {
		return nil, errors.Wrapf(err, "unable to get group by id [%v]", id)
	}
	return ret.ToGroup(), nil
}

func (s *Service) GetMultiple(ctx context.Context, tx *sqlx.Tx, logger util.Logger, ids ...string) (Groups, error) {
	if len(ids) == 0 {
		return Groups{}, nil
	}
	wc := database.SQLInClause("id", len(ids), 0, "")
	ret := rows{}
	q := database.SQLSelectSimple(columnsString, tableQuoted, wc)
	vals := make([]any, 0, len(ids))
	for _, x := range ids {
		vals = append(vals, x)
	}
	err := s.dbRead.Select(ctx, &ret, q, tx, logger, vals...)
	if err != nil {
		return nil, errors.Wrapf(err, "unable to get Groups for [%d] ids", len(ids))
	}
	return ret.ToGroups(), nil
}

func (s *Service) GetChildren(ctx context.Context, tx *sqlx.Tx, logger util.Logger) ([]*util.KeyValInt, error) {
	wc := ""
	q := database.SQLSelectGrouped("\"child\" as key, count(*) as val", tableQuoted, wc, "\"child\"", "\"child\"", 0, 0)
	var ret []*util.KeyValInt
	err := s.dbRead.Select(ctx, &ret, q, tx, logger)
	if err != nil {
		return nil, errors.Wrap(err, "unable to get groups by child")
	}
	return ret, nil
}

func (s *Service) GetByChild(ctx context.Context, tx *sqlx.Tx, child string, params *filter.Params, logger util.Logger) (Groups, error) {
	params = filters(params)
	wc := "\"child\" = $1"
	q := database.SQLSelect(columnsString, tableQuoted, wc, params.OrderByString(), params.Limit, params.Offset)
	ret := rows{}
	err := s.dbRead.Select(ctx, &ret, q, tx, logger, child)
	if err != nil {
		return nil, errors.Wrapf(err, "unable to get groups by child [%v]", child)
	}
	return ret.ToGroups(), nil
}

func (s *Service) ListSQL(ctx context.Context, tx *sqlx.Tx, sql string, logger util.Logger, values ...any) (Groups, error) {
	ret := rows{}
	err := s.dbRead.Select(ctx, &ret, sql, tx, logger, values...)
	if err != nil {
		return nil, errors.Wrap(err, "unable to get groups using custom SQL")
	}
	return ret.ToGroups(), nil
}
