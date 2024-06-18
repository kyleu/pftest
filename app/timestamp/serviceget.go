package timestamp

import (
	"context"

	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
	"github.com/samber/lo"

	"github.com/kyleu/pftest/app/lib/database"
	"github.com/kyleu/pftest/app/lib/filter"
	"github.com/kyleu/pftest/app/util"
)

func (s *Service) Get(ctx context.Context, tx *sqlx.Tx, id string, includeDeleted bool, logger util.Logger) (*Timestamp, error) {
	wc := defaultWC(0)
	wc = addDeletedClause(wc, includeDeleted)
	ret := &row{}
	q := database.SQLSelectSimple(columnsString, tableQuoted, s.db.Type, wc)
	err := s.dbRead.Get(ctx, ret, q, tx, logger, id)
	if err != nil {
		return nil, errors.Wrapf(err, "unable to get timestamp by id [%v]", id)
	}
	return ret.ToTimestamp(), nil
}

//nolint:lll
func (s *Service) GetMultiple(ctx context.Context, tx *sqlx.Tx, params *filter.Params, includeDeleted bool, logger util.Logger, ids ...string) (Timestamps, error) {
	if len(ids) == 0 {
		return Timestamps{}, nil
	}
	params = filters(params)
	wc := database.SQLInClause("id", len(ids), 0, s.db.Type)
	wc = addDeletedClause(wc, includeDeleted)
	q := database.SQLSelect(columnsString, tableQuoted, wc, params.OrderByString(), params.Limit, params.Offset, s.db.Type)
	ret := rows{}
	err := s.dbRead.Select(ctx, &ret, q, tx, logger, lo.ToAnySlice(ids)...)
	if err != nil {
		return nil, errors.Wrapf(err, "unable to get Timestamps for [%d] ids", len(ids))
	}
	return ret.ToTimestamps(), nil
}

func (s *Service) Random(ctx context.Context, tx *sqlx.Tx, logger util.Logger) (*Timestamp, error) {
	ret := &row{}
	q := database.SQLSelect(columnsString, tableQuoted, "", "random()", 1, 0, s.db.Type)
	err := s.db.Get(ctx, ret, q, tx, logger)
	if err != nil {
		return nil, errors.Wrap(err, "unable to get random timestamps")
	}
	return ret.ToTimestamp(), nil
}
