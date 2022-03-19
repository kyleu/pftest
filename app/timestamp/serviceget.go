// Content managed by Project Forge, see [projectforge.md] for details.
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
	wc := defaultWC
	wc = addDeletedClause(wc, includeDeleted)
	ret := &dto{}
	q := database.SQLSelectSimple(columnsString, tableQuoted, wc)
	err := s.db.Get(ctx, ret, q, tx, id)
	if err != nil {
		return nil, errors.Wrapf(err, "unable to get timestamp by id [%v]", id)
	}
	return ret.ToTimestamp(), nil
}

func (s *Service) GetMultiple(ctx context.Context, tx *sqlx.Tx, includeDeleted bool, ids ...string) (Timestamps, error) {
	if len(ids) == 0 {
		return Timestamps{}, nil
	}
	wc := database.SQLInClause("id", len(ids), 0)
	ret := dtos{}
	q := database.SQLSelectSimple(columnsString, tableQuoted, wc)
	vals := make([]any, 0, len(ids))
	for _, x := range ids {
		vals = append(vals, x)
	}
	err := s.db.Select(ctx, &ret, q, tx, vals...)
	if err != nil {
		return nil, errors.Wrapf(err, "unable to get Timestamps for [%d] ids", len(ids))
	}
	return ret.ToTimestamps(), nil
}
