// Content managed by Project Forge, see [projectforge.md] for details.
package trouble

import (
	"context"
	"strings"

	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"

	"github.com/kyleu/pftest/app/lib/database"
	"github.com/kyleu/pftest/app/lib/filter"
	"github.com/kyleu/pftest/app/util"
)

func (s *Service) List(ctx context.Context, tx *sqlx.Tx, params *filter.Params, includeDeleted bool, logger util.Logger) (Troubles, error) {
	params = filters(params)
	wc := ""
	if !includeDeleted {
		wc = "\"delete\" is null"
	}
	q := database.SQLSelect(columnsString, tablesJoined, wc, params.OrderByString(), params.Limit, params.Offset)
	ret := dtos{}
	err := s.db.Select(ctx, &ret, q, tx, logger)
	if err != nil {
		return nil, errors.Wrap(err, "unable to get troubles")
	}
	return ret.ToTroubles(), nil
}

func (s *Service) Count(ctx context.Context, tx *sqlx.Tx, whereClause string, includeDeleted bool, logger util.Logger, args ...any) (int, error) {
	if strings.Contains(whereClause, "'") || strings.Contains(whereClause, ";") {
		return 0, errors.Errorf("invalid where clause [%s]", whereClause)
	}
	if !includeDeleted {
		if whereClause == "" {
			whereClause = "\"delete\" is null"
		} else {
			whereClause += "and " + "\"delete\" is null"
		}
	}
	q := database.SQLSelectSimple(columnsString, tablesJoined, whereClause)
	ret, err := s.db.SingleInt(ctx, q, tx, logger, args...)
	if err != nil {
		return 0, errors.Wrap(err, "unable to get count of troubles")
	}
	return int(ret), nil
}

func (s *Service) Get(ctx context.Context, tx *sqlx.Tx, from string, where int, includeDeleted bool, logger util.Logger) (*Trouble, error) {
	wc := defaultWC(0)
	wc = addDeletedClause(wc, includeDeleted)
	ret := &dto{}
	q := database.SQLSelectSimple(columnsString, tablesJoined, wc)
	err := s.db.Get(ctx, ret, q, tx, logger, from, where)
	if err != nil {
		return nil, errors.Wrapf(err, "unable to get trouble by from [%v], where [%v]", from, where)
	}
	return ret.ToTrouble(), nil
}

func (s *Service) GetByFrom(ctx context.Context, tx *sqlx.Tx, from string, params *filter.Params, includeDeleted bool, logger util.Logger) (Troubles, error) {
	params = filters(params)
	wc := "\"from\" = $1"
	wc = addDeletedClause(wc, includeDeleted)
	q := database.SQLSelect(columnsString, tablesJoined, wc, params.OrderByString(), params.Limit, params.Offset)
	ret := dtos{}
	err := s.db.Select(ctx, &ret, q, tx, logger, from)
	if err != nil {
		return nil, errors.Wrapf(err, "unable to get troubles by from [%v]", from)
	}
	return ret.ToTroubles(), nil
}

func (s *Service) GetByWhere(ctx context.Context, tx *sqlx.Tx, where int, params *filter.Params, includeDeleted bool, logger util.Logger) (Troubles, error) {
	params = filters(params)
	wc := "\"where\" = $1"
	wc = addDeletedClause(wc, includeDeleted)
	q := database.SQLSelect(columnsString, tablesJoined, wc, params.OrderByString(), params.Limit, params.Offset)
	ret := dtos{}
	err := s.db.Select(ctx, &ret, q, tx, logger, where)
	if err != nil {
		return nil, errors.Wrapf(err, "unable to get troubles by where [%v]", where)
	}
	return ret.ToTroubles(), nil
}

func (s *Service) ListSQL(ctx context.Context, tx *sqlx.Tx, sql string, logger util.Logger) (Troubles, error) {
	ret := dtos{}
	err := s.db.Select(ctx, &ret, sql, tx, logger)
	if err != nil {
		return nil, errors.Wrap(err, "unable to get troubles using custom SQL")
	}
	return ret.ToTroubles(), nil
}
