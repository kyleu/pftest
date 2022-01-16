package mixed_case

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
	err := s.db.Select(ctx, &ret, q, tx)
	if err != nil {
		return nil, errors.Wrap(err, "unable to get mixed cases")
	}
	return ret.ToMixedCases(), nil
}

func (s *Service) Get(ctx context.Context, tx *sqlx.Tx, id string) (*MixedCase, error) {
	wc := "\"id\" = $1"
	ret := &dto{}
	q := database.SQLSelectSimple(columnsString, tableQuoted, wc)
	err := s.db.Get(ctx, ret, q, tx, id)
	if err != nil {
		return nil, errors.Wrapf(err, "unable to get mixedCase by id [%v]", id)
	}
	return ret.ToMixedCase(), nil
}
