package timestamp

import (
	"context"

	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"

	"github.com/kyleu/pftest/app/lib/database"
	"github.com/kyleu/pftest/app/lib/filter"
)

func (s *Service) List(ctx context.Context, tx *sqlx.Tx, params *filter.Params, includeDeleted bool) (Timestamps, error) {
	params = filters(params)
	wc := ""
	if !includeDeleted {
		wc = "\"deleted\" is null"
	}
	q := database.SQLSelect(columnsString, tableQuoted, wc, params.OrderByString(), params.Limit, params.Offset)
	ret := dtos{}
	err := s.db.Select(ctx, &ret, q, tx)
	if err != nil {
		return nil, errors.Wrap(err, "unable to get timestamps")
	}
	return ret.ToTimestamps(), nil
}

func (s *Service) Get(ctx context.Context, tx *sqlx.Tx, id string, includeDeleted bool) (*Timestamp, error) {
	wc := "\"id\" = $1"
	wc = addDeletedClause(wc, includeDeleted)
	ret := &dto{}
	q := database.SQLSelectSimple(columnsString, tableQuoted, wc)
	err := s.db.Get(ctx, ret, q, tx, id)
	if err != nil {
		return nil, errors.Wrapf(err, "unable to get timestamp by id [%v]", id)
	}
	return ret.ToTimestamp(), nil
}
