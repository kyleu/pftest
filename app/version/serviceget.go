// Content managed by Project Forge, see [projectforge.md] for details.
package version

import (
	"context"

	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"

	"github.com/kyleu/pftest/app/lib/database"
	"github.com/kyleu/pftest/app/lib/filter"
)

func (s *Service) List(ctx context.Context, tx *sqlx.Tx, params *filter.Params) (Versions, error) {
	params = filters(params)
	wc := ""
	q := database.SQLSelect(columnsString, tablesJoined, wc, params.OrderByString(), params.Limit, params.Offset)
	ret := dtos{}
	err := s.db.Select(ctx, &ret, q, tx)
	if err != nil {
		return nil, errors.Wrap(err, "unable to get versions")
	}
	return ret.ToVersions(), nil
}

func (s *Service) Get(ctx context.Context, tx *sqlx.Tx, id string) (*Version, error) {
	wc := defaultWC
	ret := &dto{}
	q := database.SQLSelectSimple(columnsString, tablesJoined, wc)
	err := s.db.Get(ctx, ret, q, tx, id)
	if err != nil {
		return nil, errors.Wrapf(err, "unable to get version by id [%v]", id)
	}
	return ret.ToVersion(), nil
}

func (s *Service) GetMultiple(ctx context.Context, tx *sqlx.Tx, ids ...string) (Versions, error) {
	if len(ids) == 0 {
		return Versions{}, nil
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
		return nil, errors.Wrapf(err, "unable to get Versions for [%d] ids", len(ids))
	}
	return ret.ToVersions(), nil
}
