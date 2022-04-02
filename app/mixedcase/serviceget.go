// Content managed by Project Forge, see [projectforge.md] for details.
package mixedcase

import (
	"context"

	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"

	"github.com/kyleu/pftest/app/lib/database"
	"github.com/kyleu/pftest/app/lib/filter"
)

func (s *Service) List(ctx context.Context, tx *sqlx.Tx, params *filter.Params) (MixedCases, error) {
	params = filters(params)
	wc := ""
	q := database.SQLSelect(columnsString, tableQuoted, wc, params.OrderByString(), params.Limit, params.Offset)
	ret := dtos{}
	err := s.db.Select(ctx, &ret, q, tx, s.logger)
	if err != nil {
		return nil, errors.Wrap(err, "unable to get mixed cases")
	}
	return ret.ToMixedCases(), nil
}

func (s *Service) Get(ctx context.Context, tx *sqlx.Tx, id string) (*MixedCase, error) {
	wc := defaultWC(0)
	ret := &dto{}
	q := database.SQLSelectSimple(columnsString, tableQuoted, wc)
	err := s.db.Get(ctx, ret, q, tx, s.logger, id)
	if err != nil {
		return nil, errors.Wrapf(err, "unable to get mixedCase by id [%v]", id)
	}
	return ret.ToMixedCase(), nil
}

func (s *Service) GetMultiple(ctx context.Context, tx *sqlx.Tx, ids ...string) (MixedCases, error) {
	if len(ids) == 0 {
		return MixedCases{}, nil
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
		return nil, errors.Wrapf(err, "unable to get MixedCases for [%d] ids", len(ids))
	}
	return ret.ToMixedCases(), nil
}
