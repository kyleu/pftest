// Content managed by Project Forge, see [projectforge.md] for details.
package capital

import (
	"context"

	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"

	"github.com/kyleu/pftest/app/lib/database"
	"github.com/kyleu/pftest/app/lib/filter"
)

func (s *Service) List(ctx context.Context, tx *sqlx.Tx, params *filter.Params) (Capitals, error) {
	params = filters(params)
	wc := ""
	q := database.SQLSelect(columnsString, tablesJoined, wc, params.OrderByString(), params.Limit, params.Offset)
	ret := dtos{}
	err := s.db.Select(ctx, &ret, q, tx)
	if err != nil {
		return nil, errors.Wrap(err, "unable to get capitals")
	}
	return ret.ToCapitals(), nil
}

func (s *Service) Get(ctx context.Context, tx *sqlx.Tx, id string) (*Capital, error) {
	wc := defaultWC
	ret := &dto{}
	q := database.SQLSelectSimple(columnsString, tablesJoined, wc)
	err := s.db.Get(ctx, ret, q, tx, id)
	if err != nil {
		return nil, errors.Wrapf(err, "unable to get capital by id [%v]", id)
	}
	return ret.ToCapital(), nil
}

func (s *Service) GetMultiple(ctx context.Context, tx *sqlx.Tx, IDs ...string) (Capitals, error) {
	if len(IDs) == 0 {
		return Capitals{}, nil
	}
	wc := database.SQLInClause("ID", len(IDs), 0)
	ret := dtos{}
	q := database.SQLSelectSimple(columnsString, tableQuoted, wc)
	vals := make([]any, 0, len(IDs))
	for _, x := range IDs {
		vals = append(vals, x)
	}
	err := s.db.Select(ctx, &ret, q, tx, vals...)
	if err != nil {
		return nil, errors.Wrapf(err, "unable to get Capitals for [%d] IDs", len(IDs))
	}
	return ret.ToCapitals(), nil
}
