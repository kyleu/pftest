// Content managed by Project Forge, see [projectforge.md] for details.
package softdel

import (
	"context"
	"strings"

	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"

	"github.com/kyleu/pftest/app/lib/database"
	"github.com/kyleu/pftest/app/lib/filter"
	"github.com/kyleu/pftest/app/util"
)

func (s *Service) List(ctx context.Context, tx *sqlx.Tx, params *filter.Params, includeDeleted bool, logger util.Logger) (Softdels, error) {
	params = filters(params)
	wc := ""
	if !includeDeleted {
		wc = "\"deleted\" is null"
	}
	q := database.SQLSelect(columnsString, tableQuoted, wc, params.OrderByString(), params.Limit, params.Offset)
	ret := dtos{}
	err := s.dbRead.Select(ctx, &ret, q, tx, logger)
	if err != nil {
		return nil, errors.Wrap(err, "unable to get softdels")
	}
	return ret.ToSoftdels(), nil
}

func (s *Service) Count(ctx context.Context, tx *sqlx.Tx, whereClause string, includeDeleted bool, logger util.Logger, args ...any) (int, error) {
	if strings.Contains(whereClause, "'") || strings.Contains(whereClause, ";") {
		return 0, errors.Errorf("invalid where clause [%s]", whereClause)
	}
	if !includeDeleted {
		if whereClause == "" {
			whereClause = "\"deleted\" is null"
		} else {
			whereClause += "and " + "\"deleted\" is null"
		}
	}
	q := database.SQLSelectSimple(columnsString, tableQuoted, whereClause)
	ret, err := s.dbRead.SingleInt(ctx, q, tx, logger, args...)
	if err != nil {
		return 0, errors.Wrap(err, "unable to get count of softdels")
	}
	return int(ret), nil
}

func (s *Service) Get(ctx context.Context, tx *sqlx.Tx, id string, includeDeleted bool, logger util.Logger) (*Softdel, error) {
	wc := defaultWC(0)
	wc = addDeletedClause(wc, includeDeleted)
	ret := &dto{}
	q := database.SQLSelectSimple(columnsString, tableQuoted, wc)
	err := s.dbRead.Get(ctx, ret, q, tx, logger, id)
	if err != nil {
		return nil, errors.Wrapf(err, "unable to get softdel by id [%v]", id)
	}
	return ret.ToSoftdel(), nil
}

func (s *Service) GetMultiple(ctx context.Context, tx *sqlx.Tx, includeDeleted bool, logger util.Logger, ids ...string) (Softdels, error) {
	if len(ids) == 0 {
		return Softdels{}, nil
	}
	wc := database.SQLInClause("id", len(ids), 0)
	wc = addDeletedClause(wc, includeDeleted)
	ret := dtos{}
	q := database.SQLSelectSimple(columnsString, tableQuoted, wc)
	vals := make([]any, 0, len(ids))
	for _, x := range ids {
		vals = append(vals, x)
	}
	err := s.dbRead.Select(ctx, &ret, q, tx, logger, vals...)
	if err != nil {
		return nil, errors.Wrapf(err, "unable to get Softdels for [%d] ids", len(ids))
	}
	return ret.ToSoftdels(), nil
}

func (s *Service) ListSQL(ctx context.Context, tx *sqlx.Tx, sql string, logger util.Logger) (Softdels, error) {
	ret := dtos{}
	err := s.dbRead.Select(ctx, &ret, sql, tx, logger)
	if err != nil {
		return nil, errors.Wrap(err, "unable to get softdels using custom SQL")
	}
	return ret.ToSoftdels(), nil
}
