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
	sql := database.SQLSelect(columnsString, tablesJoined, wc, params.OrderByString(), params.Limit, params.Offset)
	ret := dtos{}
	err := s.db.Select(ctx, &ret, sql, tx)
	if err != nil {
		return nil, errors.Wrap(err, "unable to get versions")
	}
	return ret.ToVersions(), nil
}

func (s *Service) Get(ctx context.Context, tx *sqlx.Tx, id string) (*Version, error) {
	wc := "\"id\" = $1"
	ret := &dto{}
	sql := database.SQLSelectSimple(columnsString, tablesJoined, wc)
	err := s.db.Get(ctx, ret, sql, tx, id)
	if err != nil {
		return nil, errors.Wrapf(err, "unable to get version by id [%s]", id)
	}
	return ret.ToVersion(), nil
}
