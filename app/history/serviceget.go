// Content managed by Project Forge, see [projectforge.md] for details.
package history

import (
	"context"

	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"

	"github.com/kyleu/pftest/app/lib/database"
	"github.com/kyleu/pftest/app/lib/filter"
)

func (s *Service) List(ctx context.Context, tx *sqlx.Tx, params *filter.Params) (Histories, error) {
	params = filters(params)
	wc := ""
	q := database.SQLSelect(columnsString, tableQuoted, wc, params.OrderByString(), params.Limit, params.Offset)
	ret := dtos{}
	err := s.db.Select(ctx, &ret, q, tx, s.logger)
	if err != nil {
		return nil, errors.Wrap(err, "unable to get histories")
	}
	return ret.ToHistories(), nil
}

func (s *Service) Get(ctx context.Context, tx *sqlx.Tx, id string) (*History, error) {
	wc := defaultWC(0)
	ret := &dto{}
	q := database.SQLSelectSimple(columnsString, tableQuoted, wc)
	err := s.db.Get(ctx, ret, q, tx, s.logger, id)
	if err != nil {
		return nil, errors.Wrapf(err, "unable to get history by id [%v]", id)
	}
	return ret.ToHistory(), nil
}

func (s *Service) GetMultiple(ctx context.Context, tx *sqlx.Tx, ids ...string) (Histories, error) {
	if len(ids) == 0 {
		return Histories{}, nil
	}
	wc := database.SQLInClause("id", len(ids), 0)
	ret := dtos{}
	q := database.SQLSelectSimple(columnsString, tableQuoted, wc)
	vals := make([]any, 0, len(ids))
	for _, x := range ids {
		vals = append(vals, x)
	}
	err := s.db.Select(ctx, &ret, q, tx, s.logger, vals...)
	if err != nil {
		return nil, errors.Wrapf(err, "unable to get Histories for [%d] ids", len(ids))
	}
	return ret.ToHistories(), nil
}
