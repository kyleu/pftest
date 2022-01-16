package softdel

import (
	"context"

	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"

	"github.com/kyleu/pftest/app/lib/database"
	"github.com/kyleu/pftest/app/lib/filter"
)

func (s *Service) List(ctx context.Context, tx *sqlx.Tx, params *filter.Params, includeDeleted bool) (Softdels, error) {
	params = filters(params)
	wc := ""
	if !includeDeleted {
		wc = "\"deleted\" is null"
	}
	sql := database.SQLSelect(columnsString, tableQuoted, wc, params.OrderByString(), params.Limit, params.Offset)
	ret := dtos{}
	err := s.db.Select(ctx, &ret, sql, tx)
	if err != nil {
		return nil, errors.Wrap(err, "unable to get softdels")
	}
	return ret.ToSoftdels(), nil
}

func (s *Service) Get(ctx context.Context, tx *sqlx.Tx, id string, includeDeleted bool) (*Softdel, error) {
	wc := "\"id\" = $1"
	wc = addDeletedClause(wc, includeDeleted)
	ret := &dto{}
	sql := database.SQLSelectSimple(columnsString, tableQuoted, wc)
	err := s.db.Get(ctx, ret, sql, tx, id)
	if err != nil {
		return nil, errors.Wrapf(err, "unable to get softdel by id [%s]", id)
	}
	return ret.ToSoftdel(), nil
}
