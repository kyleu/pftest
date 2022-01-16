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
	err := s.db.Select(ctx, &ret, q, tx)
	if err != nil {
		return nil, errors.Wrap(err, "unable to get troubles")
	}
	return ret.ToTroubles(), nil
}

func (s *Service) Get(ctx context.Context, tx *sqlx.Tx, from string, where int, includeDeleted bool) (*Trouble, error) {
	wc := "\"from\" = $1 and \"where\" = $2"
	wc = addDeletedClause(wc, includeDeleted)
	ret := &dto{}
	q := database.SQLSelectSimple(columnsString, tablesJoined, wc)
	err := s.db.Get(ctx, ret, q, tx, from, where)
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
	err := s.db.Select(ctx, &ret, q, tx, from)
	if err != nil {
		return nil, errors.Wrapf(err, "unable to get troubles by from [%s]", from)
	}
	return ret.ToTroubles(), nil
}

func (s *Service) GetByWhere(ctx context.Context, tx *sqlx.Tx, where int, params *filter.Params, includeDeleted bool) (Troubles, error) {
	params = filters(params)
	wc := "\"where\" = $1"
	wc = addDeletedClause(wc, includeDeleted)
	q := database.SQLSelect(columnsString, tablesJoined, wc, params.OrderByString(), params.Limit, params.Offset)
	ret := dtos{}
	err := s.db.Select(ctx, &ret, q, tx, where)
	if err != nil {
		return nil, errors.Wrapf(err, "unable to get troubles by where [%s]", where)
	}
	return ret.ToTroubles(), nil
}
