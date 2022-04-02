// Content managed by Project Forge, see [projectforge.md] for details.
package trouble

import (
	"context"

	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"

	"github.com/kyleu/pftest/app/lib/database"
	"github.com/kyleu/pftest/app/lib/filter"
)

func (s *Service) List(ctx context.Context, tx *sqlx.Tx, params *filter.Params, includeDeleted bool) (Troubles, error) {
	params = filters(params)
	wc := ""
	if !includeDeleted {
		wc = "\"delete\" is null"
	}
	q := database.SQLSelect(columnsString, tablesJoined, wc, params.OrderByString(), params.Limit, params.Offset)
	ret := dtos{}
	err := s.db.Select(ctx, &ret, q, tx, s.logger)
	if err != nil {
		return nil, errors.Wrap(err, "unable to get troubles")
	}
	return ret.ToTroubles(), nil
}

func (s *Service) Get(ctx context.Context, tx *sqlx.Tx, from string, where int, includeDeleted bool) (*Trouble, error) {
	wc := defaultWC(0)
	wc = addDeletedClause(wc, includeDeleted)
	ret := &dto{}
	q := database.SQLSelectSimple(columnsString, tablesJoined, wc)
	err := s.db.Get(ctx, ret, q, tx, s.logger, from, where)
	if err != nil {
		return nil, errors.Wrapf(err, "unable to get trouble by from [%v], where [%v]", from, where)
	}
	return ret.ToTrouble(), nil
}

func (s *Service) GetByFrom(ctx context.Context, tx *sqlx.Tx, from string, params *filter.Params, includeDeleted bool) (Troubles, error) {
	params = filters(params)
	wc := "\"from\" = $1"
	wc = addDeletedClause(wc, includeDeleted)
	q := database.SQLSelect(columnsString, tablesJoined, wc, params.OrderByString(), params.Limit, params.Offset)
	ret := dtos{}
	err := s.db.Select(ctx, &ret, q, tx, s.logger, from)
	if err != nil {
		return nil, errors.Wrapf(err, "unable to get troubles by from [%v]", from)
	}
	return ret.ToTroubles(), nil
}

func (s *Service) GetByWhere(ctx context.Context, tx *sqlx.Tx, where int, params *filter.Params, includeDeleted bool) (Troubles, error) {
	params = filters(params)
	wc := "\"where\" = $1"
	wc = addDeletedClause(wc, includeDeleted)
	q := database.SQLSelect(columnsString, tablesJoined, wc, params.OrderByString(), params.Limit, params.Offset)
	ret := dtos{}
	err := s.db.Select(ctx, &ret, q, tx, s.logger, where)
	if err != nil {
		return nil, errors.Wrapf(err, "unable to get troubles by where [%v]", where)
	}
	return ret.ToTroubles(), nil
}
